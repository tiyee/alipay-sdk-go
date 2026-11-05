// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package sync

// InvoicePDFSyn 是支付宝官方 V3 OpenAPI 模型。
type InvoicePDFSyn struct {
	// ApplyId 是支付宝开票申请id
	ApplyId *string `json:"apply_id,omitempty"`
	// ExtendFields 是预留的扩展字段，格式如：key1=value1\nkey2=value2\nkey3=value3，字段之间以\n(换行符)分隔。
	ExtendFields *string `json:"extend_fields,omitempty"`
	// FileBase 是base64化的字符串，与file_download_url不可同时为空。
	FileBase *string `json:"file_base,omitempty"`
	// FileDownloadType 是文件类型，必填。	PDF类型文件填写PDF，	OFD类型文件填写ofd，	JPG类型文件填写JPG(JPG文件请先询问对接人当前是否支持)
	FileDownloadType *string `json:"file_download_type,omitempty"`
	// FileDownloadUrl 是发票文件下载地址，和file_base不可同时为空。
	FileDownloadUrl *string `json:"file_download_url,omitempty"`
	// LoginId 是支付宝登录账号
	LoginId *string `json:"login_id,omitempty"`
	// OpenId 是支付宝用户id，当apply_id为空时，open_id必填
	OpenId *string `json:"open_id,omitempty"`
	// OutInvoiceId 是外部ISV的唯一发票标识
	OutInvoiceId *string `json:"out_invoice_id,omitempty"`
	// UserId 是支付宝用户userId
	UserId *string `json:"user_id,omitempty"`
	// Zip 是压缩方式
	Zip *string `json:"zip,omitempty"`
}
