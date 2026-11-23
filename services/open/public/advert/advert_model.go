// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package advert

// Advert 是支付宝官方 V3 OpenAPI 模型。
type Advert struct {
	// AdvertGroup 是广告组标识。
	AdvertGroup *string `json:"advert_group,omitempty"`
	// AdvertId 是广告位id。
	AdvertId *string `json:"advert_id,omitempty"`
	// AdvertItems 是广告位轮播内容列表。
	AdvertItems *[]AdvertItem `json:"advert_items,omitempty"`
	// GroupId 是投放人群分组id。
	GroupId *string `json:"group_id,omitempty"`
	// MobileClientType 是手机客户端类型，iphone、android、wp,空为不区分机型。
	MobileClientType *string `json:"mobile_client_type,omitempty"`
}

// AdvertItem 是支付宝官方 V3 OpenAPI 模型。
type AdvertItem struct {
	// ImgUrl 是广告图片url, 尺寸为996*240，最大不超过5M，支持格式:.jpg、.png ，请先调用<a href="https://docs.open.alipay.com/api_3/alipay.offline.material.image.upload"> 图片上传接口</a>获得图片url。
	ImgUrl *string `json:"img_url,omitempty"`
	// LinkType 是跳转类型，网页:HTTP、小程序:APP，不传默认HTTP。
	LinkType *string `json:"link_type,omitempty"`
	// LinkUrl 是跳转链接，点击广告图片跳到的链接url。
	LinkUrl *string `json:"link_url,omitempty"`
	// Name 是广告帧名称，可选。
	Name *string `json:"name,omitempty"`
}
