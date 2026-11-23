package trans

import (
	"github.com/tiyee/alipay-sdk-go/core"
)

// OrderQueryRequest 是 alipay.fund.trans.order.query（转账订单查询，
// V3 GET）的请求：order_id 与 out_biz_no 二选一，同时提供时优先
// 使用 order_id。
type OrderQueryRequest struct {
	// OrderID 是支付宝转账单据号；与 OutBizNo 不能同时为空。
	OrderID *string `query:"order_id,omitempty"`
	// OutBizNo 是商户转账唯一订单号；与 OrderID 不能同时为空。
	OutBizNo *string `query:"out_biz_no,omitempty"`
}

// CommonQueryRequest 是 alipay.fund.trans.common.query（转账单据通用
// 查询，V3 GET）的请求：out_biz_no、order_id、pay_fund_order_id
// 三者不能同时为空；传入 out_biz_no 时 product_code 与 biz_scene
// 必传。三者同时传入时优先级：pay_fund_order_id > order_id > out_biz_no。
type CommonQueryRequest struct {
	// OrderID 是支付宝转账单据号。
	OrderID *string `query:"order_id,omitempty"`
	// PayFundOrderID 是支付宝支付资金流水号。
	PayFundOrderID *string `query:"pay_fund_order_id,omitempty"`
	// OutBizNo 是商户转账唯一订单号。
	OutBizNo *string `query:"out_biz_no,omitempty"`
	// ProductCode 是销售产品码；传入 out_biz_no 时必传。
	ProductCode *string `query:"product_code,omitempty"`
	// BizScene 是业务场景；传入 out_biz_no 时必传。
	BizScene *string `query:"biz_scene,omitempty"`
}

// PayeeBindQueryRequest 是 alipay.fund.trans.payee.bind.query
// （收款账号绑定查询，V3 GET）的请求。
type PayeeBindQueryRequest struct {
	// Identity 是证件号。
	Identity string `query:"identity"`
	// IdentityType 是证件类型；暂仅支持 IDENTITY_CARD（身份证）。
	IdentityType *string `query:"identity_type,omitempty"`
}

func (r OrderQueryRequest) validate() error {
	if (r.OrderID == nil || *r.OrderID == "") && (r.OutBizNo == nil || *r.OutBizNo == "") {
		return &core.ValidationError{Field: "order_id",
			Reason: "order_id or out_biz_no is required"}
	}
	return nil
}

func (r CommonQueryRequest) validate() error {
	if r.OutBizNo != nil && *r.OutBizNo != "" {
		if r.ProductCode == nil || *r.ProductCode == "" {
			return &core.ValidationError{Field: "product_code",
				Reason: "product_code is required when out_biz_no is provided"}
		}
		if r.BizScene == nil || *r.BizScene == "" {
			return &core.ValidationError{Field: "biz_scene",
				Reason: "biz_scene is required when out_biz_no is provided"}
		}
		return nil
	}
	if r.OrderID != nil && *r.OrderID != "" {
		return nil
	}
	if r.PayFundOrderID != nil && *r.PayFundOrderID != "" {
		return nil
	}
	return &core.ValidationError{Field: "order_id",
		Reason: "one of out_biz_no, order_id or pay_fund_order_id is required"}
}

func (r PayeeBindQueryRequest) validate() error {
	if r.Identity == "" {
		return &core.ValidationError{Field: "identity", Reason: "identity is required"}
	}
	return nil
}

// UniTransferRequest 是 alipay.fund.trans.uni.transfer（单笔转账接口
// 或现金红包）的请求。out_biz_no、trans_amount、product_code、
// payee_info 必填；产品码决定取值范围。
type UniTransferRequest struct {
	// OutBizNo 是商家侧唯一订单号；不同转账请求需保证唯一。
	OutBizNo string `json:"out_biz_no"`
	// TransAmount 是订单总金额，单位为元，精确到小数点后两位：
	// TRANS_ACCOUNT_NO_PWD 取值范围 [0.1,100000000]，
	// STD_RED_PACKET 取值范围 [0.01,100000000]。
	TransAmount string `json:"trans_amount"`
	// ProductCode 是业务产品码：单笔无密转账到支付宝账户固定为
	// TRANS_ACCOUNT_NO_PWD，收发现金红包固定为 STD_RED_PACKET。
	ProductCode string `json:"product_code"`
	// BizScene 描述特定的业务场景：DIRECT_TRANSFER（单笔无密转账到
	// 支付宝、B2C 现金红包）、PERSONAL_COLLECTION（C2C 现金红包-领红包）。
	BizScene *string `json:"biz_scene,omitempty"`
	// PayeeInfo 是收款方信息。
	PayeeInfo *Participant `json:"payee_info,omitempty"`
	// PayerInfo 是付款方信息。
	PayerInfo *Participant `json:"payer_info,omitempty"`
	// OrderTitle 是转账业务的标题，用于在支付宝用户的账单里显示。
	OrderTitle *string `json:"order_title,omitempty"`
	// Remark 是业务备注。
	Remark *string `json:"remark,omitempty"`
	// BusinessParams 是转账业务请求的扩展参数（JSON 文本，
	// 如子业务场景 sub_biz_scene=REDPACKET）。
	BusinessParams *string `json:"business_params,omitempty"`
	// PassbackParams 是公用回传参数，异步通知时原样回传。
	PassbackParams *string `json:"passback_params,omitempty"`
	// OriginalOrderID 是原支付宝业务单号。
	OriginalOrderID *string `json:"original_order_id,omitempty"`
	// TransferSceneName 是该笔转账使用的转账场景，见商家平台
	// 资金管理-转账场景。
	TransferSceneName *string `json:"transfer_scene_name,omitempty"`
	// TransferSceneReportInfos 是转账场景信息；26 年及以后新接入
	// 商户必须传入，字段说明见官方转账场景信息字段说明。
	TransferSceneReportInfos []TransferSceneReportInfo `json:"transfer_scene_report_infos,omitempty"`
	// MultipleCurrencyDetail 是多币种转账金额与币种明细。
	MultipleCurrencyDetail *MultipleCurrencyDetail `json:"mutiple_currency_detail,omitempty"`
	// SignData 是转账请求的原始签名信息。
	SignData *SignData `json:"sign_data,omitempty"`
}

// ToAccountTransferRequest 是 alipay.fund.trans.toaccount.transfer
// （单笔转账到支付宝账户）的请求。out_biz_no、amount、payee_account、
// payee_type 必填；付款方和收款方不能是同一个账户。
type ToAccountTransferRequest struct {
	// OutBizNo 是商户转账唯一订单号；同一来源方必须保证唯一，
	// 只支持半角英文、数字及 “-”“_”。
	OutBizNo string `json:"out_biz_no"`
	// TransAmount 是转账金额，单位为元，只支持 2 位小数，
	// 金额必须大于等于 0.1 元。
	TransAmount string `json:"amount"`
	// PayeeAccount 是收款方账户，与 PayeeType 配合使用。
	PayeeAccount string `json:"payee_account"`
	// PayeeType 是收款方账户类型。
	PayeeType PayeeAccountType `json:"payee_type"`
	// PayeeRealName 是收款方真实姓名；非空时校验收款账户实名一致。
	PayeeRealName *string `json:"payee_real_name,omitempty"`
	// PayerRealName 是付款方真实姓名；非空时校验付款账户实名一致。
	PayerRealName *string `json:"payer_real_name,omitempty"`
	// PayerShowName 是付款方显示名称，展示在收款方的账单详情页。
	PayerShowName *string `json:"payer_show_name,omitempty"`
	// Remark 是转账备注；企业账户单笔 ≥5 万元时不能为空。
	Remark *string `json:"remark,omitempty"`
	// ExtParam 是扩展参数（JSON 字符串，目前仅支持 order_title）。
	ExtParam *string `json:"ext_param,omitempty"`
}

// PayRequest 是 alipay.fund.trans.pay（资金支付/收银台转账）的请求。
// out_biz_no、trans_amount、product_code、payee_info 必填。
type PayRequest struct {
	// OutBizNo 是商户端的唯一订单号；同一笔转账请求需保证唯一。
	OutBizNo string `json:"out_biz_no"`
	// TransAmount 是订单总金额，单位为元，精确到小数点后两位，
	// 取值范围 [0.01,100000000]。
	TransAmount string `json:"trans_amount"`
	// ProductCode 是销售产品码。
	ProductCode string `json:"product_code"`
	// PayeeInfo 是收款方信息。
	PayeeInfo *Participant `json:"payee_info,omitempty"`
	// PayerInfo 是付款方信息。
	PayerInfo *Participant `json:"payer_info,omitempty"`
	// OrderTitle 是订单标题，用于在收银台和消费记录展示。
	OrderTitle *string `json:"order_title,omitempty"`
	// BizScene 是业务场景（如群收款、红包等）。
	BizScene *string `json:"biz_scene,omitempty"`
	// AuthInfo 是协议授权信息。
	AuthInfo *AuthInfo `json:"auth_info,omitempty"`
	// TimeExpire 是绝对超时时间（yyyy-MM-dd HH:mm）。
	TimeExpire *string `json:"time_expire,omitempty"`
	// RefundTimeExpire 是退款超时时间（yyyy-MM-dd HH:mm）；
	// 到期后系统自动原路退款。
	RefundTimeExpire *string `json:"refund_time_expire,omitempty"`
	// Remark 是备注。
	Remark *string `json:"remark,omitempty"`
	// BusinessParams 是业务扩展参数（JSON 文本，使用前与支付宝确认）。
	BusinessParams *string `json:"business_params,omitempty"`
	// PassbackParams 是回传参数，异步通知时原样回传。
	PassbackParams *string `json:"passback_params,omitempty"`
}

// RefundRequest 是 alipay.fund.trans.refund（资金退回）的请求。
// order_id、biz_scene、out_request_no、refund_amount 必填；退回失败重新提交
// 必须沿用原 out_request_no。写操作：SDK 不自动重试。
type RefundRequest struct {
	// OrderID 是原转账的支付宝订单号。
	OrderID string `json:"order_id"`
	// OutRequestNo 是资金退回请求号；重新提交必须沿用原单号。
	OutRequestNo string `json:"out_request_no"`
	// RefundAmount 是退款金额，不能大于订单金额，单位为元。
	RefundAmount string `json:"refund_amount"`
	// Remark 是资金退回备注。
	Remark *string `json:"remark,omitempty"`
	// BizScene 是场景码。
	BizScene string `json:"biz_scene"`
	// ProductCode 是产品码。
	ProductCode *string `json:"product_code,omitempty"`
	// PassbackParams 是公用回传参数。
	PassbackParams *string `json:"passback_params,omitempty"`
	// BusinessParams 表示JSON格式，传递业务扩展参数，使用前请与支付宝工程师联系！。
	BusinessParams *string `json:"business_params,omitempty"`
}

func (r UniTransferRequest) validate() error {
	if r.OutBizNo == "" {
		return &core.ValidationError{Field: "out_biz_no", Reason: "out_biz_no is required"}
	}
	if err := validAmount(r.TransAmount, "trans_amount"); err != nil {
		return err
	}
	if r.ProductCode == "" {
		return &core.ValidationError{Field: "product_code", Reason: "product_code is required"}
	}
	if r.PayeeInfo == nil {
		return &core.ValidationError{Field: "payee_info", Reason: "payee_info is required"}
	}
	return validateParticipant(r.PayeeInfo)
}

func (r ToAccountTransferRequest) validate() error {
	if r.OutBizNo == "" {
		return &core.ValidationError{Field: "out_biz_no", Reason: "out_biz_no is required"}
	}
	if err := validAmount(r.TransAmount, "amount"); err != nil {
		return err
	}
	// 官方要求金额必须大于等于 0.1 元。
	if amountLessThan(r.TransAmount, "0.1") {
		return &core.ValidationError{Field: "amount", Reason: "amount must be at least 0.1 yuan"}
	}
	if r.PayeeAccount == "" {
		return &core.ValidationError{Field: "payee_account", Reason: "payee_account is required"}
	}
	if r.PayeeType == "" {
		return &core.ValidationError{Field: "payee_type", Reason: "payee_type is required"}
	}
	if !r.PayeeType.IsKnown() {
		return &core.ValidationError{Field: "payee_type",
			Reason: "payee_type must be ALIPAY_USERID, ALIPAY_LOGONID or ALIPAY_OPENID"}
	}
	return nil
}

func (r PayRequest) validate() error {
	if r.OutBizNo == "" {
		return &core.ValidationError{Field: "out_biz_no", Reason: "out_biz_no is required"}
	}
	if err := validAmount(r.TransAmount, "trans_amount"); err != nil {
		return err
	}
	if r.ProductCode == "" {
		return &core.ValidationError{Field: "product_code", Reason: "product_code is required"}
	}
	if r.PayeeInfo == nil {
		return &core.ValidationError{Field: "payee_info", Reason: "payee_info is required"}
	}
	return validateParticipant(r.PayeeInfo)
}

func (r RefundRequest) validate() error {
	if r.OrderID == "" {
		return &core.ValidationError{Field: "order_id", Reason: "order_id is required"}
	}
	if r.OutRequestNo == "" {
		return &core.ValidationError{Field: "out_request_no", Reason: "out_request_no is required"}
	}
	if r.BizScene == "" {
		return &core.ValidationError{Field: "biz_scene", Reason: "biz_scene is required"}
	}
	return validAmount(r.RefundAmount, "refund_amount")
}
