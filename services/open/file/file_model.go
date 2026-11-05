// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package file

// FileUpload 是支付宝官方 V3 OpenAPI 模型。
type FileUpload struct {
	// BizCode 是素材使用场景，请关联其他开放平台接口联合使用。如果素材用于内容发布，请必传入content_creation。
	BizCode         *string               `json:"biz_code,omitempty"`
	ExtraParameters *FileUploadExtraParam `json:"extra_parameters,omitempty"`
	// Md5 是文件的md5值，如果不传则不校验md5。
	Md5 *string `json:"md5,omitempty"`
}

// FileUploadExtraParam 是支付宝官方 V3 OpenAPI 模型。
type FileUploadExtraParam struct {
	// ExternUpload 是是否同步上传到其他平台。素材类型为视频且需要同步优酷，请必传入参数youku，素材接口将代上传到优酷账号；素材类型为图片，无需传入参数。
	ExternUpload *string `json:"extern_upload,omitempty"`
	// FileEncryptType 是文件加密类型。
	FileEncryptType *string `json:"file_encrypt_type,omitempty"`
}
