package trade_test

import (
	"context"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"errors"
	stdhtml "html"
	"net/url"
	"strings"
	"testing"

	"github.com/tiyee/alipay-sdk-go/core"
	"github.com/tiyee/alipay-sdk-go/core/option"
	"github.com/tiyee/alipay-sdk-go/internal/testutil"
	"github.com/tiyee/alipay-sdk-go/services/trade"
)

// 页面/唤端产物测试：用独立实现的签名核对（testutil.SignForm 语义）
// 验证产物签名可被网关复算；并核对产物形态与 HTML 转义。

func newPaymentProducts(t *testing.T, opts ...core.ClientOption) *trade.PaymentProducts {
	t.Helper()
	// 产物网关与 REST 网关是两项独立配置；测试指向同一受控地址。
	opts = append([]core.ClientOption{
		option.WithAppID(testutil.TestAppID),
		option.WithGateway("https://openapi.example.com/gateway.do"),
		option.WithArtifactGateway("https://openapi.example.com/gateway.do"),
	}, opts...)
	products, err := trade.NewPaymentProducts(testutil.NewCredentials(t), opts...)
	if err != nil {
		t.Fatalf("NewPaymentProducts() error = %v", err)
	}
	return products
}

func pagePayRequest() trade.PagePayRequest {
	return trade.PagePayRequest{
		OutTradeNo:   "20261105001",
		TotalAmount:  "88.88",
		Subject:      `Iphone6 16G "特惠"`,
		ProductCode:  "FAST_INSTANT_TRADE_PAY",
		ReturnURL:    ptr("https://merchant.example.com/return"),
		ExtendParams: &trade.ExtendParams{SysServiceProviderID: ptr("2088000000000000")},
	}
}

// verifyArtifactValues 按网关规则独立复算产物签名：
// 剔除 sign 与空值、按名升序 k=v 用 & 连接，用应用公钥验证。
func verifyArtifactValues(t *testing.T, values url.Values) {
	t.Helper()
	if err := testutil.VerifyWithAppKey([]byte(testutil.SignForm(values)), values.Get("sign")); err != nil {
		t.Fatalf("artifact signature mismatch: %v", err)
	}
}

func TestPagePayFormArtifact(t *testing.T) {
	products := newPaymentProducts(t)
	html, err := products.PagePayForm(context.Background(), pagePayRequest())
	if err != nil {
		t.Fatalf("PagePayForm() error = %v", err)
	}
	// 表单形态与官方一致：punchout_form POST，公共参数在 action query。
	if !strings.HasPrefix(html, `<form name="punchout_form" method="post" action="https://openapi.example.com/gateway.do?`) {
		t.Fatalf("form prefix = %q", html[:80])
	}
	if !strings.Contains(html, `<script>document.forms[0].submit();</script>`) {
		t.Error("auto submit script missing")
	}
	action := html[strings.Index(html, `action="`)+len(`action="`):]
	action = action[:strings.Index(action, `"`)]
	actionURL, err := url.Parse(stdhtml.UnescapeString(action))
	if err != nil {
		t.Fatal(err)
	}
	actionQuery, err := url.ParseQuery(actionURL.RawQuery)
	if err != nil {
		t.Fatal(err)
	}
	for key, want := range map[string]string{
		"method":    "alipay.trade.page.pay",
		"app_id":    testutil.TestAppID,
		"format":    "json",
		"charset":   "UTF-8",
		"version":   "1.0",
		"sign_type": "RSA2",
	} {
		if got := actionQuery.Get(key); got != want {
			t.Errorf("action %s = %q, want %q", key, got, want)
		}
	}
	// biz_content 与 return_url 属于业务参数，不进入 action query，
	// 以隐藏域提交；网关将 query 与表单合并后验签。
	if actionQuery.Has("biz_content") || actionQuery.Has("return_url") {
		t.Error("biz params must be hidden fields, not action query params")
	}
	// 隐藏域 biz_content 的值经过双引号转义，解码后为合法 JSON。
	hidden := `<input type="hidden" name="biz_content" value="`
	i := strings.Index(html, hidden)
	if i < 0 {
		t.Fatal("biz_content hidden field missing")
	}
	value := html[i+len(hidden):]
	value = value[:strings.Index(value, `">`)]
	if strings.Contains(value, `"`) {
		t.Errorf("double quotes must be escaped: %q", value)
	}
	// 隐藏域值中的双引号经官方规则转义；解码后为合法 JSON。
	decoded := stdhtml.UnescapeString(value)
	var biz map[string]any
	if err := json.Unmarshal([]byte(decoded), &biz); err != nil {
		t.Fatalf("hidden biz_content is not valid json: %v", err)
	}
	if biz["total_amount"] != "88.88" || biz["product_code"] != "FAST_INSTANT_TRADE_PAY" {
		t.Fatalf("biz_content = %v", biz)
	}

	// 合并 query 与隐藏域后按网关规则复算签名；return_url 位于隐藏域。
	merged := url.Values{}
	for key, vs := range actionQuery {
		merged[key] = vs
	}
	merged.Set("biz_content", decoded)
	merged.Set("return_url", "https://merchant.example.com/return")
	verifyArtifactValues(t, merged)
}

func TestPagePayURLArtifact(t *testing.T) {
	products := newPaymentProducts(t)
	rawURL, err := products.PagePayURL(context.Background(), pagePayRequest())
	if err != nil {
		t.Fatalf("PagePayURL() error = %v", err)
	}
	if !strings.HasPrefix(rawURL, "https://openapi.example.com/gateway.do?") {
		t.Fatalf("url = %q", rawURL)
	}
	parsed, err := url.Parse(rawURL)
	if err != nil {
		t.Fatal(err)
	}
	values, err := url.ParseQuery(parsed.RawQuery)
	if err != nil {
		t.Fatal(err)
	}
	if values.Get("method") != "alipay.trade.page.pay" || values.Get("biz_content") == "" {
		t.Fatalf("query missing required params: %v", values)
	}
	verifyArtifactValues(t, values)

	// 篡改检测：修改任意参数值后，独立复算必须失败。
	values.Set("total_amount", "0.01")
	if biz := values.Get("biz_content"); strings.Contains(biz, `"total_amount":"88.88"`) {
		values.Set("biz_content", strings.Replace(biz, `"total_amount":"88.88"`, `"total_amount":"0.01"`, 1))
	} else {
		t.Fatalf("biz_content missing total_amount: %q", biz)
	}
	if err := testutil.VerifyWithAppKey([]byte(testutil.SignForm(values)), values.Get("sign")); err == nil {
		t.Fatal("tampered artifact must fail verification")
	}
}

func TestWapPayUsesOwnMethod(t *testing.T) {
	products := newPaymentProducts(t)
	req := pagePayRequest()
	req.ProductCode = "QUICK_WAP_WAY"
	rawURL, err := products.WapPayURL(context.Background(), req)
	if err != nil {
		t.Fatalf("WapPayURL() error = %v", err)
	}
	parsed, _ := url.Parse(rawURL)
	values, _ := url.ParseQuery(parsed.RawQuery)
	if values.Get("method") != "alipay.trade.wap.pay" {
		t.Fatalf("method = %q", values.Get("method"))
	}
	verifyArtifactValues(t, values)
}

func TestAppPayOrderString(t *testing.T) {
	products := newPaymentProducts(t)
	order, err := products.AppPayOrderString(context.Background(), trade.AppPayRequest{
		OutTradeNo:  "20261105002",
		TotalAmount: "9.90",
		Subject:     "订单标题",
		ProductCode: trade.ProductCodeCyclePayAuth,
		AgreementSignParams: &trade.SignParams{
			PersonalProductCode: ptr("CYCLE_PAY_AUTH_P"),
			AccessParams:        &trade.SignAccessParams{Channel: ptr("ALIPAYAPP")},
			PeriodRuleParams: &trade.PeriodRuleParams{
				PeriodType: ptr("MONTH"), Period: ptr(int64(1)), ExecuteTime: ptr("2026-12-01"), SingleAmount: ptr("9.90"),
			},
			ProductCode: ptr("GENERAL_WITHHOLDING"),
			SignScene:   ptr("DEFAULT|DEFAULT"),
		},
	})
	if err != nil {
		t.Fatalf("AppPayOrderString() error = %v", err)
	}
	// 订单串不含网关地址；参数对可解析且签名可独立复算。
	if strings.Contains(order, "http") {
		t.Fatalf("order string must not contain gateway url: %q", order)
	}
	values, err := url.ParseQuery(order)
	if err != nil {
		t.Fatalf("parse order string: %v", err)
	}
	if values.Get("method") != "alipay.trade.app.pay" {
		t.Fatalf("method = %q", values.Get("method"))
	}
	verifyArtifactValues(t, values)
	var biz struct {
		ObsoleteKey json.RawMessage `json:"sign_params"`
		SignParams  *struct {
			ProductCode string `json:"product_code"`
		} `json:"agreement_sign_params"`
	}
	if err := json.Unmarshal([]byte(values.Get("biz_content")), &biz); err != nil {
		t.Fatal(err)
	}
	if biz.ObsoleteKey != nil || biz.SignParams == nil || biz.SignParams.ProductCode != "GENERAL_WITHHOLDING" {
		t.Fatalf("agreement_sign_params missing: %v", biz)
	}
}

func TestAppPayDefaultsAndSigningConstraints(t *testing.T) {
	products := newPaymentProducts(t)
	req := trade.AppPayRequest{OutTradeNo: "app_order", TotalAmount: "9.90", Subject: "item"}
	order, err := products.AppPayOrderString(t.Context(), req)
	if err != nil {
		t.Fatal(err)
	}
	values, err := url.ParseQuery(order)
	if err != nil {
		t.Fatal(err)
	}
	verifyArtifactValues(t, values)
	var biz map[string]json.RawMessage
	if err := json.Unmarshal([]byte(values.Get("biz_content")), &biz); err != nil {
		t.Fatal(err)
	}
	if biz["product_code"] != nil || biz["agreement_sign_params"] != nil {
		t.Fatalf("unexpected fields in ordinary payment: %v", biz)
	}

	checkField := func(field string) {
		t.Helper()
		_, err := products.AppPayOrderString(t.Context(), req)
		var invalid *core.ValidationError
		if !errors.As(err, &invalid) || invalid.Field != field {
			t.Fatalf("got %v, want error for %s", err, field)
		}
	}
	req.ProductCode = trade.ProductCodeCyclePayAuth
	checkField("agreement_sign_params")
	sign := &trade.SignParams{ProductCode: ptr("GENERAL_WITHHOLDING"), PersonalProductCode: ptr("CYCLE_PAY_AUTH_P"), SignScene: ptr("DEFAULT|DEFAULT")}
	req.AgreementSignParams = sign
	checkField("agreement_sign_params.access_params.channel")
	sign.AccessParams = &trade.SignAccessParams{Channel: ptr("ALIPAYAPP")}
	checkField("agreement_sign_params.period_rule_params")
	period := &trade.PeriodRuleParams{SingleAmount: ptr("9.90")}
	sign.PeriodRuleParams = period
	checkField("agreement_sign_params.period_rule_params.period_type")
	period.PeriodType = ptr("MONTH")
	checkField("agreement_sign_params.period_rule_params.period")
	period.Period = ptr(int64(1))
	period.ExecuteTime = ptr("2026-12-31")
	checkField("agreement_sign_params.period_rule_params.execute_time")
}

func TestArtifactCertificateModeParams(t *testing.T) {
	suite := testutil.NewCertSuite(t)
	products, err := trade.NewPaymentProducts(suite.NewCertificateCredentials(t),
		option.WithAppID(testutil.TestAppID),
		option.WithGateway("https://openapi.example.com/gateway.do"),
		option.WithArtifactGateway("https://openapi.example.com/gateway.do"))
	if err != nil {
		t.Fatalf("NewPaymentProducts() error = %v", err)
	}
	rawURL, err := products.PagePayURL(context.Background(), pagePayRequest())
	if err != nil {
		t.Fatalf("PagePayURL() error = %v", err)
	}
	parsed, _ := url.Parse(rawURL)
	values, _ := url.ParseQuery(parsed.RawQuery)
	if values.Get("app_cert_sn") == "" || values.Get("alipay_root_cert_sn") == "" {
		t.Fatalf("certificate params missing: %v", values)
	}
	// 证书模式签名使用应用证书私钥（与测试应用私钥配对），
	// 用应用公钥独立复算。
	if err := testutil.VerifyWithAppKey([]byte(testutil.SignForm(values)), values.Get("sign")); err != nil {
		t.Fatalf("cert artifact signature mismatch: %v", err)
	}
	// 证书 SN 与套件一致。
	sn := testutil.CertSN(t, suite.AppCert)
	if values.Get("app_cert_sn") != sn {
		t.Fatalf("app_cert_sn = %q, want %q", values.Get("app_cert_sn"), sn)
	}
	_ = x509.ErrUnsupportedAlgorithm // 保持 x509 引用（套件类型经由 testutil 使用）
	_ = base64.StdEncoding
}

func TestArtifactValidations(t *testing.T) {
	products := newPaymentProducts(t)
	if _, err := products.PagePayURL(context.Background(), trade.PagePayRequest{
		OutTradeNo: "T1", TotalAmount: "1.00", Subject: "标题", ProductCode: ""}); err == nil {
		t.Fatal("product_code is required")
	}
	if _, err := products.AppPayOrderString(context.Background(), trade.AppPayRequest{
		OutTradeNo: "T1", TotalAmount: "0.00", Subject: "标题", ProductCode: "QUICK_MSECURITY_PAY"}); err == nil {
		t.Fatal("total_amount must be a positive decimal")
	}
}

func TestWapPayChecksEchoedFieldsBeforeSigning(t *testing.T) {
	products := newPaymentProducts(t)
	for _, tt := range []struct{ field, value string }{
		{"passback_params", "note%26buyer_id%3Dother"}, {"passback_params", "note&buyer_id=other"}, {"passback_params", "bad%"},
		{"out_trade_no", "order&out_z=value"}, {"body", "memo&seller_id=second"}, {"subject", "title&total_amount=1"},
	} {
		req := pagePayRequest()
		req.ProductCode = "QUICK_WAP_WAY"
		switch tt.field {
		case "passback_params":
			req.PassbackParams = &tt.value
		case "out_trade_no":
			req.OutTradeNo = tt.value
		case "body":
			req.Body = &tt.value
		case "subject":
			req.Subject = tt.value
		}
		artifact, err := products.WapPayURL(t.Context(), req)
		var v *core.ValidationError
		if !errors.As(err, &v) || v.Field != tt.field || artifact != "" {
			t.Fatalf("%+v: %v", tt, err)
		}
	}
	req := pagePayRequest()
	req.ProductCode = "QUICK_WAP_WAY"
	passback := url.QueryEscape("A&B 中文")
	req.PassbackParams = &passback
	artifact, err := products.WapPayURL(t.Context(), req)
	if err != nil {
		t.Fatal(err)
	}
	target, err := url.Parse(artifact)
	if err != nil {
		t.Fatal(err)
	}
	var biz struct {
		Passback string `json:"passback_params"`
	}
	if err := json.Unmarshal([]byte(target.Query().Get("biz_content")), &biz); err != nil || biz.Passback != passback {
		t.Fatalf("passback %q, err %v", biz.Passback, err)
	}
}

func TestPageAndWapSpecificPaymentParameters(t *testing.T) {
	products := newPaymentProducts(t)
	req := pagePayRequest()
	req.QRPayMode = ptr(trade.QRPayCustomWidth)
	req.QRCodeWidth = ptr(int64(180))
	req.StoreID = ptr("store")
	req.BusinessParams = ptr(`{"feature":"value"}`)
	for _, method := range []string{"page", "wap"} {
		var raw string
		var err error
		if method == "page" {
			raw, err = products.PagePayURL(t.Context(), req)
		} else {
			req.ProductCode = trade.ProductCodeQuickWapWay
			if _, err = products.WapPayURL(t.Context(), req); err == nil {
				t.Fatal("WAP accepted PC fields")
			}
			req = pagePayRequest()
			req.QuitURL = ptr("https://merchant.example.com/quit")
			if _, err = products.PagePayURL(t.Context(), req); err == nil {
				t.Fatal("PC accepted quit_url")
			}
			req.ProductCode = trade.ProductCodeQuickWapWay
			raw, err = products.WapPayURL(t.Context(), req)
		}
		if err != nil {
			t.Fatal(err)
		}
		u, err := url.Parse(raw)
		if err != nil {
			t.Fatal(err)
		}
		values := u.Query()
		verifyArtifactValues(t, values)
		var biz map[string]json.RawMessage
		if err = json.Unmarshal([]byte(values.Get("biz_content")), &biz); err != nil {
			t.Fatal(err)
		}
		if method == "page" && (string(biz["qrcode_width"]) != "180" || string(biz["qr_pay_mode"]) != `"4"` || string(biz["business_params"]) != `"{\"feature\":\"value\"}"` || string(biz["store_id"]) != `"store"`) {
			t.Fatalf("unexpected PC fields: %v", biz)
		}
		if method == "wap" && (string(biz["quit_url"]) != `"https://merchant.example.com/quit"` || values.Has("quit_url")) {
			t.Fatalf("unexpected WAP fields: %v", biz)
		}
	}
}

func TestArtifactsRespectNestedParameterScope(t *testing.T) {
	products := newPaymentProducts(t)
	req := pagePayRequest()
	req.GoodsDetail = []trade.GoodsDetail{{
		GoodsID: "001", GoodsName: "item", Price: "1.00", Quantity: 1, GoodsEANID: ptr("6900000000000"),
	}}
	if _, err := products.PagePayURL(t.Context(), req); err == nil {
		t.Fatal("PC accepted WAP goods_ean_id")
	}
	req.ProductCode = trade.ProductCodeQuickWapWay
	if _, err := products.WapPayForm(t.Context(), req); err != nil {
		t.Fatal(err)
	}
	req.ExtendParams = &trade.ExtendParams{FqNum: ptr("3")}
	if _, err := products.WapPayURL(t.Context(), req); err == nil {
		t.Fatal("WAP accepted REST fq_num")
	}
	app := trade.AppPayRequest{
		OutTradeNo: "app_order", TotalAmount: "1.00", Subject: "item",
		ExtUserInfo: &trade.ExtUserInfo{FixBuyer: ptr("T")},
	}
	if _, err := products.AppPayOrderString(t.Context(), app); err == nil {
		t.Fatal("App accepted REST fix_buyer")
	}
}
