// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package bizrisks

// FactCheckRequest is the official V3 GrandsecurityBizrisksFactCheckModel schema.
type FactCheckRequest struct {
	// Query represents query. 用户输入的谣言/疑问文本，必填，最大 5000 字符
	Query *string `json:"query,omitzero"`
}
