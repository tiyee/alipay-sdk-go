// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package taxno

// InvoiceTaxNoBatchQueryRequest 是支付宝官方 V3 OpenAPI 模型。
type InvoiceTaxNoBatchQueryRequest struct {
	// AuthToken 在 URL query 中传递。用户授权令牌
	AuthToken *string `query:"auth_token" json:"-"`
	// EnableTradeOut 是默认值为false。true为输出交易信息，false为不输出交易信息。
	EnableTradeOut *string `json:"enable_trade_out,omitempty"`
	// EndInvoiceDate 是查询结束时间，精确到天（按开票日期查询）	start_invoice_date和end_invoice_date传值要求	1.同时为空时，返回最近半年200条数据	2.必须同时为空 或 同时不为空	3.结束日期不能大于当前日期	4.开始时间和结束时间跨度不能超过6个月
	EndInvoiceDate *string `json:"end_invoice_date,omitempty"`
	// InvoiceKindList 是查询票种列表
	InvoiceKindList *[]string `json:"invoice_kind_list,omitempty"`
	// LimitSize 是查询结果上限笔数，最大值20
	LimitSize *int64 `json:"limit_size,omitempty"`
	// PageNum 是当前页码，为空时默认第一页
	PageNum *int64 `json:"page_num,omitempty"`
	// Scene 是发票要素获取应用场景。"拉"模式报销 固定为 INVOICE_EXPENSE 表示发票报销。
	Scene *string `json:"scene,omitempty"`
	// StartInvoiceDate 是查询起始时间，精确到天（按开票日期查询）	start_invoice_date和end_invoice_date传值要求	1.同时为空时，返回最近半年200条数据	2.必须同时为空 或 同时不为空	3.结束日期不能大于当前日期	4.开始时间和结束时间跨度不能超过6个月
	StartInvoiceDate *string `json:"start_invoice_date,omitempty"`
	// TaxNo 是企业税号
	TaxNo *string `json:"tax_no,omitempty"`
}
