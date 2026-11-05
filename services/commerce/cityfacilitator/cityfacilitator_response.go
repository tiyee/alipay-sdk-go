// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package cityfacilitator

import (
	"encoding/json"
)

// CityFacilitatorStationQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type CityFacilitatorStationQueryResponse struct {
	// SupportStarts 是支持设为起点的站点列表
	SupportStarts *[]StationDetailInfo `json:"support_starts,omitempty"`
}

// CityFacilitatorVoucherBatchQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type CityFacilitatorVoucherBatchQueryResponse struct {
	// Tickets 是查询到的订单信息列表
	Tickets *[]TicketDetailInfo `json:"tickets,omitempty"`
}

// CityFacilitatorVoucherGenerateResponse 是支付宝官方 V3 OpenAPI 模型。
type CityFacilitatorVoucherGenerateResponse struct {
	// ExpiredDate 是核销码过期时间
	ExpiredDate *string `json:"expired_date,omitempty"`
	// QrCodeNo 是地铁购票二维码编码，可自定义
	QrCodeNo *string `json:"qr_code_no,omitempty"`
	// TicketNo 是地铁购票的核销码
	TicketNo *string `json:"ticket_no,omitempty"`
}

// CityFacilitatorVoucherRefundResponse 是支付宝官方 V3 OpenAPI 模型。
type CityFacilitatorVoucherRefundResponse = json.RawMessage
