// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package templatemsg

// MiniTemplateMsgMaketingBatchQueryRequest 是支付宝官方 V3 OpenAPI 模型。
type MiniTemplateMsgMaketingBatchQueryRequest struct {
	// PageNum 是分页查询页码。
	PageNum *int64 `json:"page_num,omitempty"`
	// PageSize 是每页展示条数，最大支持50个。
	PageSize *int64 `json:"page_size,omitempty"`
}

// MiniTemplateMsgTinyPaySwitchConfirmRequest is the request type for alipay.open.mini.templatemsg.tinypayswitch.confirm.
type MiniTemplateMsgTinyPaySwitchConfirmRequest struct{}
