// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package apply

// InvoiceItem 是支付宝官方 V3 OpenAPI 模型。
type InvoiceItem struct {
	// Amount 是含税总金额 (等于sum_price和tax之和)，	单位：元，格式为2位小数，精度2位小数
	Amount *string `json:"amount,omitempty"`
	// ItemName 是发票项目名称（或商品名称）
	ItemName *string `json:"item_name,omitempty"`
	// ItemNo 是发票项目编号（或商品编号）
	ItemNo *string `json:"item_no,omitempty"`
	// Price 是单价（不含税），格式为2位小数。最大支持6位小数，不足2位小数时需转化为2位小数格式。	折扣行此参数不能传
	Price *string `json:"price,omitempty"`
	// Quantity 是数量，须大于0，最多6位小数。	折扣行此参数不能传
	Quantity *string `json:"quantity,omitempty"`
	// RowType 是发票行性质。0表示正常行，1表示折扣行，2表示被折扣行。	比如充电器单价100元，折扣10元，则明细为2行，充电器行性质为2，折扣行性质为1。如果充电器没有折扣，则值应为0。
	RowType *int64 `json:"row_type,omitempty"`
	// Specification 是规格型号
	Specification *string `json:"specification,omitempty"`
	// SumPrice 是不含税总金额，格式为2位小数	单位：元，精度2位小数
	SumPrice *string `json:"sum_price,omitempty"`
	// Tax 是税额，	格式为2位小数
	Tax *string `json:"tax,omitempty"`
	// TaxRate 是税率。格式为2位小数，如：0.00, 0.03, 0.13等等
	TaxRate *string `json:"tax_rate,omitempty"`
	// Unit 是单位，折扣行不能传
	Unit *string `json:"unit,omitempty"`
	// ZeroRateFlag 是0税率标识，只有税率为0的情况才有值，0=出口零税率，1=免税，2=不征收，3=普通零税率
	ZeroRateFlag *string `json:"zero_rate_flag,omitempty"`
}

// InvoiceUkDTO 是支付宝官方 V3 OpenAPI 模型。
type InvoiceUkDTO struct {
	// InvoiceCode 是发票代码
	InvoiceCode *string `json:"invoice_code,omitempty"`
	// InvoiceNo 是发票号码
	InvoiceNo *string `json:"invoice_no,omitempty"`
}
