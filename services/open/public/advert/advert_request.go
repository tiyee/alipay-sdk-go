// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package advert

// PublicAdvertCreateRequest 是支付宝官方 V3 OpenAPI 模型。
type PublicAdvertCreateRequest struct {
	// AdvertItems 是广告位轮播内容列表。数量限制：大于1个，小于5个。广告位轮播内容顺序：与接口传入的顺序排列一致。。
	AdvertItems *[]AdvertItem `json:"advert_items,omitempty"`
}

// PublicAdvertModifyRequest 是支付宝官方 V3 OpenAPI 模型。
type PublicAdvertModifyRequest struct {
	// AdvertId 是广告位id，通过<a href="https://opendocs.alipay.com/apis/api_6/alipay.open.public.advert.create">alipay.open.public.advert.create</a>(生活号广告位添加接口)添加广告位获取。。
	AdvertId *string `json:"advert_id,omitempty"`
	// AdvertItems 是广告位轮播内容列表。数量限制：大于1个，小于5个。广告位轮播内容顺序：与接口传入的顺序排列一致。。
	AdvertItems *[]AdvertItem `json:"advert_items,omitempty"`
}

// PublicAdvertBatchQueryRequest is the request type for alipay.open.public.advert.batchquery.
type PublicAdvertBatchQueryRequest struct{}

// PublicAdvertDeleteRequest is the request type for alipay.open.public.advert.delete.
type PublicAdvertDeleteRequest struct {
	// AdvertId 是待删除的广告位id，删除通投广告位需传递此参数
	AdvertId *string `json:"advert_id,omitempty" query:"advert_id,omitempty"`
	// AdvertGroup 是待删除的广告位分组标识，删除个性化广告位需传递此参数。
	AdvertGroup *string `json:"advert_group,omitempty" query:"advert_group,omitempty"`
}
