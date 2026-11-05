// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package doc

// DocTemplateCreateRequest 是支付宝官方 V3 OpenAPI 模型。
type DocTemplateCreateRequest struct {
	// TargetAppId 在 URL query 中传递。目标isv应用ID
	TargetAppId *string `query:"target_app_id" json:"-"`
	// ContentMd5 是模版文件的md5值。示例代码参见 创建合同模板
	ContentMd5 *string `json:"content_md_5,omitempty"`
	// ContentType 是目标文件的MIME类型。示例代码参见 创建合同模板
	ContentType *string `json:"content_type,omitempty"`
	// ConvertToPdf 是是否需要转成pdf，如果模板文件为.doc/.docx 传true，为pdf传false
	ConvertToPdf *string `json:"convert_to_pdf,omitempty"`
	// FileName 是文件名称，必须带上文件扩展名如 .pdf、.doc、.docx，否则会导致后续发起流程校验异常。
	FileName *string `json:"file_name,omitempty"`
}
