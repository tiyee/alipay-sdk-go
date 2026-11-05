// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package merchant

// InvoiceMerchantEnterStatusQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type InvoiceMerchantEnterStatusQueryResponse struct {
	// Status 是工单状态，审核中：AUDITING，验收中：ACCEPTING，成功：SUCCESS，失败：FAIL
	Status *string `json:"status,omitempty"`
}
