// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package social

// SocialBaseContentLibStandardContentBatchQueryRequest 是支付宝官方 V3 OpenAPI 模型。
type SocialBaseContentLibStandardContentBatchQueryRequest struct {
	// NeedDetail 是指示是否返回内容详情, 默认为false	false: 仅返回内容ID列表	true:返回 content_details
	NeedDetail *bool `json:"need_detail,omitempty"`
	// PageNum 是当页分页,默认为1
	PageNum *int64 `json:"page_num,omitempty"`
	// PageSize 是分页记录数，最大20,默认20
	PageSize *int64 `json:"page_size,omitempty"`
	// PublicId 是生活号+号ID，只有MCN机构为达人号ID查询内容才需要传入
	PublicId *string `json:"public_id,omitempty"`
	// Status 是内容状态，默认查询所有状态
	Status *string `json:"status,omitempty"`
}
