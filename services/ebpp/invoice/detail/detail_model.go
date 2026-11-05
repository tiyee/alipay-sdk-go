// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package detail

// InvoiceItemContent 是支付宝官方 V3 OpenAPI 模型。
type InvoiceItemContent struct {
	// ItemAmount 是价税合计。(等于sumPrice和tax之和)
	ItemAmount *string `json:"item_amount,omitempty"`
	// ItemName 是发票项目名称（或商品名称）
	ItemName *string `json:"item_name,omitempty"`
	// ItemNo 是商品编号
	ItemNo *string `json:"item_no,omitempty"`
	// ItemPrice 是单价，格式：100.00。新版电子发票，折扣行此参数不能传，非折扣行必传
	ItemPrice *string `json:"item_price,omitempty"`
	// ItemQuantity 是数量。新版电子发票，折扣行此参数不能传，非折扣行必传
	ItemQuantity *int64 `json:"item_quantity,omitempty"`
	// ItemSumPrice 是单项总价，格式：100.00
	ItemSumPrice *string `json:"item_sum_price,omitempty"`
	// ItemTaxPrice 是税额
	ItemTaxPrice *string `json:"item_tax_price,omitempty"`
	// ItemTaxRate 是税率
	ItemTaxRate *string `json:"item_tax_rate,omitempty"`
	// ItemUnit 是台
	ItemUnit *string `json:"item_unit,omitempty"`
	// RowType 是发票行性质。0表示正常行，1表示折扣行，2表示被折扣行。比如充电器单价100元，折扣10元，则明细为2行，充电器行性质为2，折扣行性质为1。如果充电器没有折扣，则值应为0
	RowType *int64 `json:"row_type,omitempty"`
}
