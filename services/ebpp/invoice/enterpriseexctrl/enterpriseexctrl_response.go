// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package enterpriseexctrl

// InvoiceEnterpriseExCtrlEmployerTitleBatchQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type InvoiceEnterpriseExCtrlEmployerTitleBatchQueryResponse struct {
	// Code 是返回码	10000-成功	其他都是失败
	Code *string `json:"code,omitempty"`
	// Msg 是返回消息
	Msg *string `json:"msg,omitempty"`
	// TitleCode 是抬头编码
	TitleCode *string `json:"title_code,omitempty"`
	// TitleInfoList 是企业抬头信息
	TitleInfoList *[]EnterpriseTitleInfo `json:"title_info_list,omitempty"`
}

// InvoiceEnterpriseExCtrlEmployerTitleCreateResponse 是支付宝官方 V3 OpenAPI 模型。
type InvoiceEnterpriseExCtrlEmployerTitleCreateResponse struct {
	// Code 是返回码	10000-成功	其他都是失败
	Code *string `json:"code,omitempty"`
	// Msg 是返回消息
	Msg *string `json:"msg,omitempty"`
	// Success 是是否成功
	Success *bool `json:"success,omitempty"`
	// TitleId 是抬头ID
	TitleId *string `json:"title_id,omitempty"`
}

// InvoiceEnterpriseExCtrlEmployerTitleModifyResponse 是支付宝官方 V3 OpenAPI 模型。
type InvoiceEnterpriseExCtrlEmployerTitleModifyResponse struct {
	// Code 是返回码	10000-成功	其他都是失败
	Code *string `json:"code,omitempty"`
	// Msg 是返回信息
	Msg *string `json:"msg,omitempty"`
	// Success 是是否成功
	Success *bool `json:"success,omitempty"`
	// TitleId 是抬头ID
	TitleId *string `json:"title_id,omitempty"`
}

// InvoiceEnterpriseExCtrlEmployerTitleQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type InvoiceEnterpriseExCtrlEmployerTitleQueryResponse struct {
	// Code 是返回码	10000-成功	其他都是失败
	Code *string `json:"code,omitempty"`
	// Msg 是返回消息
	Msg *string `json:"msg,omitempty"`
	// TitleCode 是抬头编码
	TitleCode *string              `json:"title_code,omitempty"`
	TitleInfo *EnterpriseTitleInfo `json:"title_info,omitempty"`
}
