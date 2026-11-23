package waybill

// LogisticsWaybillIstdDetailQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type LogisticsWaybillIstdDetailQueryResponse struct {
	// ReachDuration 是预计还剩多久送达 单位：秒
	ReachDuration *int64 `json:"reach_duration,omitempty"`
	// RiderLat 是骑手位置纬度
	RiderLat *string `json:"rider_lat,omitempty"`
	// RiderLng 是骑手位置经度
	RiderLng *string `json:"rider_lng,omitempty"`
	// RiderMobileNo 是骑手电话
	RiderMobileNo *string `json:"rider_mobile_no,omitempty"`
	// RiderName 是骑手姓名
	RiderName *string `json:"rider_name,omitempty"`
	// RiderPoiLink 是骑手实时定位H5链接字段，在骑手已接单后，将可展示骑手实时位置的H5页面链接同步商户
	RiderPoiLink *string `json:"rider_poi_link,omitempty"`
	// Status 是即时配送运单状态
	Status *string `json:"status,omitempty"`
}
