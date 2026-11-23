// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package ecorder

// InvoiceEcOrderOrderQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type InvoiceEcOrderOrderQueryResponse struct {
	OrderInfo *EcOrderItem `json:"order_info,omitempty"`
	// SubOrderList 是企业码子订单详情列表
	SubOrderList *[]EcOrderItem `json:"sub_order_list,omitempty"`
}
