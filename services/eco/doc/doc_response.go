// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package doc

// DocTemplateCreateResponse 是支付宝官方 V3 OpenAPI 模型。
type DocTemplateCreateResponse struct {
	// TemplateId 是模板id（请记录模板ID（templateId），后续发起合同签署需要使用到）
	TemplateId *string `json:"template_id,omitempty"`
	// UploadUrl 是文件直传地址，需要用此上传地址使用PUT方式上传文件，只有文件上传后模板才可用
	UploadUrl *string `json:"upload_url,omitempty"`
}
