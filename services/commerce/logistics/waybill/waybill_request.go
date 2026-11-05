package waybill

// LogisticsWaybillIstdDetailQueryRequest is the query request model for alipay.commerce.logistics.waybill.istddetail.query.
type LogisticsWaybillIstdDetailQueryRequest struct {
	// ShopNo 是商家门店编号
	ShopNo *string `json:"shop_no,omitempty" query:"shop_no,omitempty"`
	// OutOrderNo 是商家订单号
	OutOrderNo *string `json:"out_order_no,omitempty" query:"out_order_no,omitempty"`
	// OrderNo 是支付宝订单流水号
	OrderNo *string `json:"order_no,omitempty" query:"order_no,omitempty"`
	// LogisticsCode 是即时配送公司编码
	LogisticsCode *string `json:"logistics_code,omitempty" query:"logistics_code,omitempty"`
	// WaybillNo 是即时配送运单编号
	WaybillNo *string `json:"waybill_no,omitempty" query:"waybill_no,omitempty"`
}
