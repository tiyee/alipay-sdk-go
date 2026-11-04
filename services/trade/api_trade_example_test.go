package trade_test

import (
	"context"
	"encoding/json/jsontext"
	"fmt"
	"net/http"

	"github.com/tiyee/go-alipay/core"
	"github.com/tiyee/go-alipay/services/trade"
)

// exampleExecutor 是示例用的模拟执行器：返回固定成功响应，
// 使示例可离线编译运行，不发出真实请求。
type exampleExecutor struct {
	node jsontext.Value
}

func (e *exampleExecutor) Do(_ context.Context, req core.WireRequest) (*core.APIResult, error) {
	fmt.Printf("POST %s method=%s biz_content=%s\n",
		"https://openapi.alipay.com/gateway.do", req.Spec.Method, string(req.Params.JSON))
	return &core.APIResult{
		HTTPStatus:        http.StatusOK,
		SignatureVerified: true,
		Node:              e.node,
	}, nil
}

// ExampleService_Query 展示交易查询：查询成功不代表支付成功，
// 业务结果必须依据 trade_status 判断。
func ExampleService_Query() {
	executor := &exampleExecutor{node: jsontext.Value(`{"code":"10000","msg":"Success",` +
		`"trade_no":"2013112011001004330000121536","out_trade_no":"6823789339978248",` +
		`"trade_status":"TRADE_SUCCESS","total_amount":"88.88"}`)}
	svc, err := trade.NewService(executor)
	if err != nil {
		panic(err)
	}
	resp, err := svc.Query(context.Background(), trade.QueryRequest{
		OutTradeNo: "6823789339978248",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.Data.TradeStatus, resp.Data.TotalAmount)
	// Output:
	// POST https://openapi.alipay.com/gateway.do method=alipay.trade.query biz_content={"out_trade_no":"6823789339978248"}
	// TRADE_SUCCESS 88.88
}

// ExampleService_Precreate 展示当面付预下单：二维码有效时间 2 小时，
// qr_code 缺失时按业务响应处理。
func ExampleService_Precreate() {
	executor := &exampleExecutor{node: jsontext.Value(`{"code":"10000","msg":"Success",` +
		`"out_trade_no":"20150320010101001","qr_code":"https://qr.alipay.com/bavh4wjlxf12tper3a"}`)}
	svc, err := trade.NewService(executor)
	if err != nil {
		panic(err)
	}
	resp, err := svc.Precreate(context.Background(), trade.PrecreateRequest{
		OutTradeNo:  "20150320010101001",
		TotalAmount: "88.88",
		Subject:     "Iphone6 16G",
	}, core.WithNotifyURL("https://merchant.example/notify"))
	if err != nil {
		panic(err)
	}
	fmt.Println(*resp.Data.QRCode)
	// Output:
	// POST https://openapi.alipay.com/gateway.do method=alipay.trade.precreate biz_content={"out_trade_no":"20150320010101001","total_amount":"88.88","subject":"Iphone6 16G"}
	// https://qr.alipay.com/bavh4wjlxf12tper3a
}

// ExampleService_Refund 展示退款：code=10000 仅代表请求成功，
// fund_change=Y 才表示退款成功，否则需查询退款确认。
func ExampleService_Refund() {
	executor := &exampleExecutor{node: jsontext.Value(`{"code":"10000","msg":"Success",` +
		`"trade_no":"2013112011001004330000121536","out_trade_no":"6823789339978248",` +
		`"refund_fee":"12.33","fund_change":"Y"}`)}
	svc, err := trade.NewService(executor)
	if err != nil {
		panic(err)
	}
	resp, err := svc.Refund(context.Background(), trade.RefundRequest{
		OutTradeNo:   "6823789339978248",
		RefundAmount: "12.33",
		OutRequestNo: strPtr("REFUND-20260101-001"),
	})
	if err != nil {
		panic(err)
	}
	if resp.Data.FundChange != nil && *resp.Data.FundChange == trade.FundChangeYes {
		fmt.Println("refunded", resp.Data.RefundFee)
	} else {
		fmt.Println("confirm via refund query")
	}
	// Output:
	// POST https://openapi.alipay.com/gateway.do method=alipay.trade.refund biz_content={"refund_amount":"12.33","out_trade_no":"6823789339978248","out_request_no":"REFUND-20260101-001"}
	// refunded 12.33
}

// ExampleService_QueryRefund 展示退款查询：code=10000 仅代表查询成功，
// 仅 refund_status=REFUND_SUCCESS 才表示退款成功。
func ExampleService_QueryRefund() {
	executor := &exampleExecutor{node: jsontext.Value(`{"code":"10000","msg":"Success",` +
		`"out_request_no":"REFUND-20260101-001","refund_status":"REFUND_SUCCESS",` +
		`"refund_amount":"12.33","total_amount":"88.88"}`)}
	svc, err := trade.NewService(executor)
	if err != nil {
		panic(err)
	}
	resp, err := svc.QueryRefund(context.Background(), trade.QueryRefundRequest{
		OutTradeNo:   "6823789339978248",
		OutRequestNo: "REFUND-20260101-001",
	})
	if err != nil {
		panic(err)
	}
	if resp.Data.RefundStatus != nil && *resp.Data.RefundStatus == trade.RefundStatusSuccess {
		fmt.Println("refund confirmed", *resp.Data.RefundAmount)
	} else {
		fmt.Println("refund not confirmed yet")
	}
	// Output:
	// POST https://openapi.alipay.com/gateway.do method=alipay.trade.fastpay.refund.query biz_content={"out_request_no":"REFUND-20260101-001","out_trade_no":"6823789339978248"}
	// refund confirmed 12.33
}

// ExampleService_Close 展示未支付订单关闭：先用查询确认交易仍处于
// WAIT_BUYER_PAY，再发起关单；已支付订单绝不进入关闭流程
// （已支付交易应走退款）。
func ExampleService_Close() {
	// 第一步：查询订单当前状态。
	statusNode := jsontext.Value(`{"code":"10000","msg":"Success",` +
		`"trade_no":"2013112011001004330000121536","out_trade_no":"6823789339978248",` +
		`"trade_status":"WAIT_BUYER_PAY","total_amount":"88.88"}`)
	svc, err := trade.NewService(&exampleExecutor{node: statusNode})
	if err != nil {
		panic(err)
	}
	query, err := svc.Query(context.Background(), trade.QueryRequest{
		OutTradeNo: "6823789339978248",
	})
	if err != nil {
		panic(err)
	}
	if query.Data.TradeStatus != trade.TradeStatusWaitBuyerPay {
		fmt.Println("not closable:", query.Data.TradeStatus)
		return
	}

	// 第二步：仅未支付订单发起关单。
	closedNode := jsontext.Value(`{"code":"10000","msg":"Success",` +
		`"trade_no":"2013112011001004330000121536","out_trade_no":"6823789339978248"}`)
	closeSvc, err := trade.NewService(&exampleExecutor{node: closedNode})
	if err != nil {
		panic(err)
	}
	resp, err := closeSvc.Close(context.Background(), trade.CloseRequest{
		OutTradeNo: "6823789339978248",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("closed", *resp.Data.OutTradeNo)
	// Output:
	// POST https://openapi.alipay.com/gateway.do method=alipay.trade.query biz_content={"out_trade_no":"6823789339978248"}
	// POST https://openapi.alipay.com/gateway.do method=alipay.trade.close biz_content={"out_trade_no":"6823789339978248"}
	// closed 6823789339978248
}

func strPtr(s string) *string { return &s }
