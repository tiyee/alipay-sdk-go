// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package cityfacilitator

// CityFacilitatorVoucherBatchQueryRequest 是支付宝官方 V3 OpenAPI 模型。
type CityFacilitatorVoucherBatchQueryRequest struct {
	// CityCode 是城市编码请参考查询 中华人民共和国行政区划代码。	已支持城市：广州 440100，深圳 440300，杭州330100。
	CityCode *string `json:"city_code,omitempty"`
	// TradeNos 是支付宝交易号列表
	TradeNos *[]string `json:"trade_nos,omitempty"`
}

// CityFacilitatorVoucherGenerateRequest 是支付宝官方 V3 OpenAPI 模型。
type CityFacilitatorVoucherGenerateRequest struct {
	// CityCode 是城市编码请参考查询 中华人民共和国行政区划代码。	已支持城市：广州 440100，深圳 440300，杭州330100。
	CityCode *string `json:"city_code,omitempty"`
	// SiteBegin 是起点站站点编码
	SiteBegin *string `json:"site_begin,omitempty"`
	// SiteEnd 是终点站站点编码
	SiteEnd *string `json:"site_end,omitempty"`
	// TicketNum 是地铁票购票数量
	TicketNum *string `json:"ticket_num,omitempty"`
	// TicketPrice 是单张票价，元为单价
	TicketPrice *string `json:"ticket_price,omitempty"`
	// TicketType 是地铁票种类，枚举支持：	*oneway。
	TicketType *string `json:"ticket_type,omitempty"`
	// TotalFee 是订单总金额，元为单位
	TotalFee *string `json:"total_fee,omitempty"`
	// TradeNo 是支付宝交易号（交易支付时，必须通过指定sellerId：2088121612215201，将钱支付到指定的中间户中）
	TradeNo *string `json:"trade_no,omitempty"`
}

// CityFacilitatorVoucherRefundRequest 是支付宝官方 V3 OpenAPI 模型。
type CityFacilitatorVoucherRefundRequest struct {
	// CityCode 是城市编码请参考查询 中华人民共和国行政区划代码。	已支持城市：广州 440100，深圳 440300，杭州330100。
	CityCode *string `json:"city_code,omitempty"`
	// TradeNo 是支付宝交易号
	TradeNo *string `json:"trade_no,omitempty"`
}

// CityFacilitatorStationQueryRequest is the query request model for alipay.commerce.cityfacilitator.station.query.
type CityFacilitatorStationQueryRequest struct {
	// CityCode 是城市编码请参考查询 中华人民共和国行政区划代码。	已支持城市：广州 440100，深圳 440300，杭州330100。
	CityCode *string `json:"city_code,omitempty" query:"city_code,omitempty"`
}
