// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package file

// FilePathQueryRequest is the query request model for alipay.eco.file.path.query.
type FilePathQueryRequest struct {
	// TargetAppId 是目标isv应用ID
	TargetAppId *string `json:"target_app_id,omitempty" query:"target_app_id,omitempty"`
	// ContentMd5 是先计算文件md5值，在对该md5值进行base64编码
	ContentMd5 *string `json:"content_md_5,omitempty" query:"content_md_5,omitempty"`
	// ContentType 是目标文件的MIME类型
	ContentType *string `json:"content_type,omitempty" query:"content_type,omitempty"`
	// FileName 是文件名称。	注意：必须带上文件扩展名，不然会导致后续发起流程校验不通过。示例：合同.pdf。
	FileName *string `json:"file_name,omitempty" query:"file_name,omitempty"`
	// FileSize 是文件大小，单位byte。最大允许上传30Mb
	FileSize *int64 `json:"file_size,omitempty" query:"file_size,omitempty"`
}
