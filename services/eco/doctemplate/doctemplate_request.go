// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package doctemplate

// DocTemplateSettingURLQueryRequest is the query request model for alipay.eco.doctemplate.settingurl.query.
type DocTemplateSettingURLQueryRequest struct {
	// TargetAppId 是目标isv应用ID
	TargetAppId *string `json:"target_app_id,omitempty" query:"target_app_id,omitempty"`
	// TemplateId 是合同模板id。通过创建合同模版（alipay.eco.doc.template.create）接口获取。
	TemplateId *string `json:"template_id,omitempty" query:"template_id,omitempty"`
}
