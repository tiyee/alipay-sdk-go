// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package cityfacilitator

// StationDetailInfo 是支付宝官方 V3 OpenAPI 模型。
type StationDetailInfo struct {
	// Code 是站点编码
	Code *string `json:"code,omitempty"`
	// ExtCode 是站点外部编码
	ExtCode *string `json:"ext_code,omitempty"`
	// Name 是站点中文名称
	Name *string `json:"name,omitempty"`
}

// TicketDetailInfo 是支付宝官方 V3 OpenAPI 模型。
type TicketDetailInfo struct {
	// Amount 是总金额，元为单位
	Amount *string `json:"amount,omitempty"`
	// EndStation 是终点站编码
	EndStation *string `json:"end_station,omitempty"`
	// EndStationName 是终点站中文名称
	EndStationName *string `json:"end_station_name,omitempty"`
	// Quantity 是票数量
	Quantity *string `json:"quantity,omitempty"`
	// StartStation 是起点站编码
	StartStation *string `json:"start_station,omitempty"`
	// StartStationName 是起点站中文名称
	StartStationName *string `json:"start_station_name,omitempty"`
	// Status 是订单状态
	Status *string `json:"status,omitempty"`
	// TicketPrice 是单价，元为单位
	TicketPrice *string `json:"ticket_price,omitempty"`
	// TicketType 是票类型
	TicketType *string `json:"ticket_type,omitempty"`
	// TradeNo 是支付宝交易号
	TradeNo *string `json:"trade_no,omitempty"`
}
