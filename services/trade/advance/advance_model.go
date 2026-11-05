package advance

// AdvanceSceneOrderRiskEvaluation 是 alipay.trade.advance.consult 的
// 单笔订单风险预评估咨询场景，取值依据官方 V3 文档。
const AdvanceSceneOrderRiskEvaluation = "ORDER_RISK_EVALUATION"

// UserRiskPrediction 是用户风险预测结果。
type UserRiskPrediction struct {
	// PhoneRecycleRiskLevel 是用户绑定手机号被二次放号风险等级。
	PhoneRecycleRiskLevel *string `json:"phone_recycle_risk_leve,omitempty"`
	// RefusedPaymentRiskLevel 是用户拒付风险等级。
	RefusedPaymentRiskLevel *string `json:"refused_payment_risk_level,omitempty"`
}

// WaitRepaymentOrderInfo 是待还订单信息。
type WaitRepaymentOrderInfo struct {
	// AdvanceOrderID 是垫资单 id。
	AdvanceOrderID *string `json:"advance_order_id,omitempty"`
	// AlipayUserID 是买家的支付宝用户 id。
	AlipayUserID *string `json:"alipay_user_id,omitempty"`
	// OpenID 是买家的支付宝用户唯一标识。
	OpenID *string `json:"open_id,omitempty"`
	// BizProduct 是商户签约的收单产品码。
	BizProduct *string `json:"biz_product,omitempty"`
	// OrigBizOrderID 是原始业务单号，通常为支付宝交易号。
	OrigBizOrderID *string `json:"orig_biz_order_id,omitempty"`
	// WaitRepaymentAmount 是垫资金额。
	WaitRepaymentAmount *string `json:"wait_repayment_amount,omitempty"`
}
