// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package life

// PublicLifeAboardApplyResponse 是支付宝官方 V3 OpenAPI 模型。
type PublicLifeAboardApplyResponse struct {
	// Result 是上架成功后返回的提示。
	Result *string `json:"result,omitempty"`
}

// PublicLifeAgentCreateQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type PublicLifeAgentCreateQueryResponse struct {
	// LifeAppId 是只有审核通过，且商户在支付宝发送的授权邮件中确认授权，此时生活号才会正式创建，查询才会返回该值。
	LifeAppId *string `json:"life_app_id,omitempty"`
	// MerchantPid 是商户pid。
	MerchantPid *string `json:"merchant_pid,omitempty"`
	// OrderStatusBizDesc 是支付宝商户入驻申请单状态，申请单状态包括：暂存、审核中、待商户确认、成功、失败。 说明：暂存是审核前的中间状态，如出现暂存状态请再次提交代创建生活号申请。。
	OrderStatusBizDesc *string `json:"order_status_biz_desc,omitempty"`
	// OutBizNo 是由开发者创建的外部入驻申请单据号。
	OutBizNo *string `json:"out_biz_no,omitempty"`
	// RefusedReason 是只有审核失败才会返回该值。
	RefusedReason *string `json:"refused_reason,omitempty"`
}

// PublicLifeDebarkApplyResponse 是支付宝官方 V3 OpenAPI 模型。
type PublicLifeDebarkApplyResponse struct {
	// Result 是下架成功后返回的提示。
	Result *string `json:"result,omitempty"`
}

// PublicLifeLabelBatchQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type PublicLifeLabelBatchQueryResponse struct {
	// Labels 是标签列表。
	Labels *[]LifeLabel `json:"labels,omitempty"`
}

// PublicLifeLabelCreateResponse 是支付宝官方 V3 OpenAPI 模型。
type PublicLifeLabelCreateResponse struct {
	// LabelId 是标签id。
	LabelId *string `json:"label_id,omitempty"`
}

// PublicLifeLabelDeleteResponse 是支付宝官方 V3 OpenAPI 模型。
type PublicLifeLabelDeleteResponse struct {
}

// PublicLifeLabelModifyResponse 是支付宝官方 V3 OpenAPI 模型。
type PublicLifeLabelModifyResponse struct {
}

// PublicLifeMsgRecallResponse 是支付宝官方 V3 OpenAPI 模型。
type PublicLifeMsgRecallResponse struct {
}

// PublicLifeMsgSendResponse 是支付宝官方 V3 OpenAPI 模型。
type PublicLifeMsgSendResponse struct {
	// AlipayMsgId 是生活号消息唯一标识。
	AlipayMsgId *string `json:"alipay_msg_id,omitempty"`
}
