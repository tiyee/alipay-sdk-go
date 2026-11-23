package advance

import (
	"github.com/tiyee/alipay-sdk-go/core"
	"github.com/tiyee/alipay-sdk-go/internal/tradevalidate"
)

// AdvanceConsultRequest 是 alipay.trade.advance.consult（订单咨询，
// V3 GET）的请求：垫资咨询场景 alipay_user_id 必选；单笔订单风险
// 预评估场景（consult_scene=ORDER_RISK_EVALUATION）需传入协议号等。
type AdvanceConsultRequest struct {
	// AlipayUserID 是支付宝用户 id（2088 开头 16 位数字）；
	// 除订单风险预评估场景外必选。
	AlipayUserID *string `query:"alipay_user_id,omitempty"`
	// ConsultScene 是订单咨询类型；默认为垫资咨询，
	// ORDER_RISK_EVALUATION 表示单笔订单风险预评估。
	ConsultScene *string `query:"consult_scene,omitempty"`
	// IndustryProductCode 是行业产品信息。
	IndustryProductCode *string `query:"industry_product_code,omitempty"`
	// SubMerchantID 是子商户 id。
	SubMerchantID *string `query:"sub_merchant_id,omitempty"`
	// SubMerchantType 是子商户类型。
	SubMerchantType *string `query:"sub_merchant_type,omitempty"`
	// AgreementNo 是用户签约协议号；单笔订单风险预评估时必传。
	AgreementNo *string `query:"agreement_no,omitempty"`
	// EstimatedOrderAmount 是预估订单总金额；单笔订单风险预评估时使用。
	EstimatedOrderAmount *string `query:"estimated_order_amount,omitempty"`
	// OutTradeNo 是商户请求时的外部订单号。
	OutTradeNo *string `query:"out_trade_no,omitempty"`
	// NeedUserRiskPrediction 标识是否需要进行用户风险预测。
	NeedUserRiskPrediction *string `query:"need_user_risk_prediction,omitempty"`
	// TradeDigest 是交易信息摘要；需要用户风险预测时可传。
	TradeDigest *string `query:"trade_digest,omitempty"`
	// UserOutBindingPhone 是用户在外部域账号系统绑定的手机号；
	// 需要用户风险预测时必传，按官方要求加密后传输。
	UserOutBindingPhone *string `query:"user_out_binding_phone,omitempty"`
}

func (r AdvanceConsultRequest) validate() error {
	riskEval := r.ConsultScene != nil && *r.ConsultScene == AdvanceSceneOrderRiskEvaluation
	if riskEval && (r.AgreementNo == nil || *r.AgreementNo == "") {
		return &core.ValidationError{Field: "agreement_no", Reason: "agreement_no is required for ORDER_RISK_EVALUATION"}
	}
	if !riskEval && (r.AlipayUserID == nil || *r.AlipayUserID == "") {
		return &core.ValidationError{Field: "alipay_user_id",
			Reason: "alipay_user_id is required unless consult_scene is ORDER_RISK_EVALUATION"}
	}
	if r.EstimatedOrderAmount != nil {
		return tradevalidate.OrderAmount(*r.EstimatedOrderAmount, "estimated_order_amount")
	}
	return nil
}
