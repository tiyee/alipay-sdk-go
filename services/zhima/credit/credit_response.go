// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package credit

// PayAfterUseCreditAgreementQueryResponse is the official V3 ZhimaCreditPayafteruseCreditagreementQueryResponseModel schema.
type PayAfterUseCreditAgreementQueryResponse struct {
	// AgreementStatus represents agreement_status. 开通/授权状态，VALID: 有效，INVALID: 无效。不可开通CANNOT_OPEN
	AgreementStatus *string `json:"agreement_status,omitzero"`
	// AlipayUserId represents alipay_user_id. 蚂蚁统一会员ID
	AlipayUserId *string `json:"alipay_user_id,omitzero"`
	// BizTime represents biz_time. 开通时间，agreement_status为VALID有效。
	BizTime *string `json:"biz_time,omitzero"`
	// CreditAgreementId represents credit_agreement_id. 芝麻开通/授权协议号
	CreditAgreementId *string `json:"credit_agreement_id,omitzero"`
	// ExtInfo represents ext_info. 该字段只在特殊业务场景下，根据业务方约定返回；一般业务场景下不需要消费该字段
	ExtInfo *string `json:"ext_info,omitzero"`
	// OpenId represents open_id. 开放ID
	OpenId *string `json:"open_id,omitzero"`
	// OutAgreementNo represents out_agreement_no. 商户外部协议号
	OutAgreementNo *string `json:"out_agreement_no,omitzero"`
}

// PayAfterUseCreditAgreementTransferResponse is the official V3 ZhimaCreditPayafteruseCreditagreementTransferResponseModel schema.
type PayAfterUseCreditAgreementTransferResponse struct {
	// CreditAgreementId represents credit_agreement_id. 芝麻开通协议号
	CreditAgreementId *string `json:"credit_agreement_id,omitzero"`
	// DeductAgreementNo represents deduct_agreement_no. 当operation_type==AGREEMENT_TRANSFER时，并且是迁移到小程序商家版时有值，其他情况下不返回
	DeductAgreementNo *string `json:"deduct_agreement_no,omitzero"`
	// OutAgreementNo represents out_agreement_no. 商户外部协议号
	OutAgreementNo *string `json:"out_agreement_no,omitzero"`
}

// PayAfterUseCreditBizOrderFinishResponse is the official V3 ZhimaCreditPayafteruseCreditbizorderFinishResponseModel schema.
type PayAfterUseCreditBizOrderFinishResponse struct {
	// CreditBizOrderId represents credit_biz_order_id. 信用服务订单号
	CreditBizOrderId *string `json:"credit_biz_order_id,omitzero"`
	// OrderStatus represents order_status. 信用服务订单状态.INIT: 下单状态; TRADE_CLOSED: 订单取消或者交易全额退款； TRADE_FINISHED：扣款成功状态
	OrderStatus *string `json:"order_status,omitzero"`
	// OutRequestNo represents out_request_no. 商户外部请求号
	OutRequestNo *string `json:"out_request_no,omitzero"`
}

// PayAfterUseCreditBizOrderOrderResponse is the official V3 ZhimaCreditPayafteruseCreditbizorderOrderResponseModel schema.
type PayAfterUseCreditBizOrderOrderResponse struct {
	// CreditBizOrderId represents credit_biz_order_id. 信用服务订单号
	CreditBizOrderId *string `json:"credit_biz_order_id,omitzero"`
	// OutOrderNo represents out_order_no. 商户传入的外部订单号
	OutOrderNo *string `json:"out_order_no,omitzero"`
}

// PayAfterUseCreditBizOrderQueryResponse is the official V3 ZhimaCreditPayafteruseCreditbizorderQueryResponseModel schema.
type PayAfterUseCreditBizOrderQueryResponse struct {
	// CreateTime represents create_time. 订单创建时间
	CreateTime *string `json:"create_time,omitzero"`
	// CreditAgreementId represents credit_agreement_id. 芝麻开通协议号
	CreditAgreementId *string `json:"credit_agreement_id,omitzero"`
	// CreditBizOrderId represents credit_biz_order_id. 信用服务订单号
	CreditBizOrderId *string `json:"credit_biz_order_id,omitzero"`
	// OrderStatus represents order_status. 信用服务订单状态，区分不同产品： （1）针对产品「先用后付」，涉及状态如下：INIT: 下单状态；TRADE_CLOSED: 订单取消或者交易全额退款； TRADE_FINISHED：扣款成功状态。 （2）针对产品「芝麻风险评估与召回」，状态定义如下： [INIT: 下单状态]；[WAIT_FULFILL：待守约]；[OVERDUE：已逾期]；[TRADE_FINISHED:已守约]； [TRADE_CLOSED：已取消]。
	OrderStatus *string `json:"order_status,omitzero"`
	// ProductCode represents product_code. 产品码
	ProductCode *string `json:"product_code,omitzero"`
	// TotalAmount represents total_amount. 信用服务订单金额
	TotalAmount *string `json:"total_amount,omitzero"`
	// TotalCreditPayAmount represents total_credit_pay_amount. 信用服务订单中信用支付金额，不包含支付宝营销金额
	TotalCreditPayAmount *string `json:"total_credit_pay_amount,omitzero"`
	// TradeNo represents trade_no. 支付宝交易号。 先用后付产品下必传；芝麻风险评估与召回产品下不传。
	TradeNo *string `json:"trade_no,omitzero"`
	// ZmServiceId represents zm_service_id. 芝麻信用服务ID
	ZmServiceId *string `json:"zm_service_id,omitzero"`
}

// PeZmGoAgreementQueryResponse is the official V3 ZhimaCreditPeZmgoAgreementQueryResponseModel schema.
type PeZmGoAgreementQueryResponse struct {
	// AgreementId represents agreement_id. 支付宝系统中用以唯一标识用户签约记录的编号，即花芝轻会员协议号。
	AgreementId *string `json:"agreement_id,omitzero"`
	// AgreementName represents agreement_name. 协议名称
	AgreementName *string `json:"agreement_name,omitzero"`
	// AgreementStatus represents agreement_status. 协议状态。Y表示状态有效，P表示状态失效中，N表示状态已失效
	AgreementStatus *string `json:"agreement_status,omitzero"`
	// AlipayUserId represents alipay_user_id. 支付宝用户userId
	AlipayUserId *string `json:"alipay_user_id,omitzero"`
	// BizType represents biz_type. 代表签约芝麻GO所属业务类型
	BizType *string `json:"biz_type,omitzero"`
	// OpenId represents open_id. 支付宝用户userId
	OpenId *string `json:"open_id,omitzero"`
	// SignTime represents sign_time. 该条芝麻GO协议签约时间
	SignTime *string `json:"sign_time,omitzero"`
	// StartTime represents start_time. 该芝麻GO协议开始时间
	StartTime *string `json:"start_time,omitzero"`
}

// PeZmGoAgreementUnsignResponse is the official V3 ZhimaCreditPeZmgoAgreementUnsignResponseModel schema.
type PeZmGoAgreementUnsignResponse struct {
	// AgreementId represents agreement_id. 支付宝系统中用以唯一标识用户签约记录的编号。
	AgreementId *string `json:"agreement_id,omitzero"`
	// WithholdPlanNo represents withhold_plan_no. 调用芝麻GO结算受理接口时，需要传入该值
	WithholdPlanNo *string `json:"withhold_plan_no,omitzero"`
}

// PeZmGoBizOptCloseResponse is the official V3 ZhimaCreditPeZmgoBizoptCloseResponseModel schema.
type PeZmGoBizOptCloseResponse struct {
	// BizOptNo represents biz_opt_no. 芝麻GO签约申请时生成的签约申请单据号
	BizOptNo *string `json:"biz_opt_no,omitzero"`
	// OpenId represents open_id. 蚂蚁统一会员ID
	OpenId *string `json:"open_id,omitzero"`
	// OutRequestNo represents out_request_no. 商户本次操作的请求流水号
	OutRequestNo *string `json:"out_request_no,omitzero"`
	// PartnerId represents partner_id. 商户ID
	PartnerId *string `json:"partner_id,omitzero"`
	// UserId represents user_id. 蚂蚁统一会员ID
	UserId *string `json:"user_id,omitzero"`
}

// PeZmGoCumulationSyncResponse is the official V3 ZhimaCreditPeZmgoCumulationSyncResponseModel schema.
type PeZmGoCumulationSyncResponse struct {
	// AgreementNo represents agreement_no. 芝麻go协议号，唯一标识一个芝麻go协议。 累计成功后返回此字段。
	AgreementNo *string `json:"agreement_no,omitzero"`
	// FailReason represents fail_reason. 数据回传失败原因描述。 累计失败后返回此字段。
	FailReason *string `json:"fail_reason,omitzero"`
	// OpenId represents open_id. 用户userId，蚂蚁统一会员ID。 累计成功后返回此字段。
	OpenId *string `json:"open_id,omitzero"`
	// OutBizNo represents out_biz_no. 外部业务号。需保证唯一，若接口调用失败，提示信息为更换幂等号，商户更换此字段重新发起调用。 累计成功后返回此字段。
	OutBizNo *string `json:"out_biz_no,omitzero"`
	// UserId represents user_id. 用户userId，蚂蚁统一会员ID。 累计成功后返回此字段。
	UserId *string `json:"user_id,omitzero"`
}

// PeZmGoPreOrderCreateResponse is the official V3 ZhimaCreditPeZmgoPreorderCreateResponseModel schema.
type PeZmGoPreOrderCreateResponse struct {
	// BizType represents biz_type. 业务类型
	BizType *string `json:"biz_type,omitzero"`
	// PartnerId represents partner_id. 商户ID
	PartnerId *string `json:"partner_id,omitzero"`
	// PreorderNo represents preorder_no. 芝麻GO预创建单号
	PreorderNo *string `json:"preorder_no,omitzero"`
}

// PeZmGoSettleApplyResponse is the official V3 ZhimaCreditPeZmgoSettleApplyResponseModel schema.
type PeZmGoSettleApplyResponse struct {
	// AgreementId represents agreement_id. 支付宝系统中用以唯一标识用户签约记录的编号。
	AgreementId *string `json:"agreement_id,omitzero"`
	// FailReason represents fail_reason. 根据商户传入的轻会员协议结算周期内累计消费数据与支付宝端轻会员协议结算周期内消费数据比对之后，若数据不一致，则会返回对应的异常数据比对结果。
	FailReason *string `json:"fail_reason,omitzero"`
	// OutRequestNo represents out_request_no. 商户本次操作的请求流水号，用于标示请求流水的唯一性，不能包含除中文、英文、数字以外的字符，需要保证在商户端不重复。
	OutRequestNo *string `json:"out_request_no,omitzero"`
	// SettleStatus represents settle_status. Y("Y", "支付成功"), P("P", "支付中"), N("N", "支付失败"),
	SettleStatus *string `json:"settle_status,omitzero"`
	// WithholdPlanNo represents withhold_plan_no. 芝麻GO结算时，需要传入的扣款单据号。来源于协议到期通知消息
	WithholdPlanNo *string `json:"withhold_plan_no,omitzero"`
}

// PeZmGoSettleRefundResponse is the official V3 ZhimaCreditPeZmgoSettleRefundResponseModel schema.
type PeZmGoSettleRefundResponse struct {
	// FailReason represents fail_reason. 退款操作错误原因描述
	FailReason *string `json:"fail_reason,omitzero"`
	// OutRequestNo represents out_request_no. 外部请求号
	OutRequestNo *string `json:"out_request_no,omitzero"`
	// RefundAmount represents refund_amount. 本次请求退款金额，单位为：元（人民币），精确到小数点后两位
	RefundAmount *string `json:"refund_amount,omitzero"`
	// RefundOptNo represents refund_opt_no. 芝麻GO退款生成的内部唯一流水，外部商户可以拿这个参数进行对账
	RefundOptNo *string `json:"refund_opt_no,omitzero"`
	// Retry represents retry. 此字段含义，商家接入时需咨询对接技术。
	Retry *bool `json:"retry,omitzero"`
	// WithholdPlanNo represents withhold_plan_no. 本次退款的扣款计划单号
	WithholdPlanNo *string `json:"withhold_plan_no,omitzero"`
}

// PeZmGoSettleUnfreezeResponse is the official V3 ZhimaCreditPeZmgoSettleUnfreezeResponseModel schema.
type PeZmGoSettleUnfreezeResponse struct {
	// FailReaseon represents fail_reaseon. 失败原因
	FailReaseon *string `json:"fail_reaseon,omitzero"`
	// Retry represents retry. 是否可以重试
	Retry *string `json:"retry,omitzero"`
	// UnfreezeAmount represents unfreeze_amount. 解冻金额字符串
	UnfreezeAmount *string `json:"unfreeze_amount,omitzero"`
	// UnfreezeStatus represents unfreeze_status. 解冻状态
	UnfreezeStatus *string `json:"unfreeze_status,omitzero"`
}
