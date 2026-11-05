// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package file

// InvoiceFileOutputQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type InvoiceFileOutputQueryResponse struct {
	// FileType 是发票文件类型pdf或ofd
	FileType *string `json:"file_type,omitempty"`
	// InvoiceFileContent 是发票文件内容,需要先base64解码，获取原始byte[]类型文件流
	InvoiceFileContent *string `json:"invoice_file_content,omitempty"`
}
