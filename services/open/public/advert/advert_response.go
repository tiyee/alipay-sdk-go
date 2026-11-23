// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package advert

// PublicAdvertBatchQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type PublicAdvertBatchQueryResponse struct {
	// AdvertList 是广告位列表 。目前只有一个广告位。。
	AdvertList *[]Advert `json:"advert_list,omitempty"`
	// Count 是广告组数量。
	Count *int64 `json:"count,omitempty"`
}

// PublicAdvertCreateResponse 是支付宝官方 V3 OpenAPI 模型。
type PublicAdvertCreateResponse struct {
	// AdvertGroup 是广告分组标识。
	AdvertGroup *string `json:"advert_group,omitempty"`
	// AdvertId 是广告位id。
	AdvertId *string `json:"advert_id,omitempty"`
}

// PublicAdvertDeleteResponse 是支付宝官方 V3 OpenAPI 模型。
type PublicAdvertDeleteResponse struct {
}

// PublicAdvertModifyResponse 是支付宝官方 V3 OpenAPI 模型。
type PublicAdvertModifyResponse struct {
}
