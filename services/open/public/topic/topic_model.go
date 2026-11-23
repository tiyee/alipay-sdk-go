// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package topic

// Topic 是支付宝官方 V3 OpenAPI 模型。
type Topic struct {
	// ImgUrl 是营销位图片url。
	ImgUrl *string `json:"img_url,omitempty"`
	// LinkType 是跳转类型，网页:HTTP、小程序:APP。
	LinkType *string `json:"link_type,omitempty"`
	// LinkUrl 是营销位跳转地址，点击营销位头图跳到的链接url。。
	LinkUrl *string `json:"link_url,omitempty"`
	// SubTitle 是营销位描述。
	SubTitle *string `json:"sub_title,omitempty"`
	// Title 是营销位名称。
	Title *string `json:"title,omitempty"`
	// TopicId 是营销位id。
	TopicId *string `json:"topic_id,omitempty"`
	// TopicItems 是营销位内容列表。
	TopicItems *[]TopicItem `json:"topic_items,omitempty"`
}

// TopicItem 是支付宝官方 V3 OpenAPI 模型。
type TopicItem struct {
	// ImgUrl 是内容图片url, 尺寸为300*300，最大不超过3M，支持格式:.jpg、.png ，请先调用<a href="https://docs.open.alipay.com/api_3/alipay.offline.material.image.upload"> 图片上传接口</a>获得图片url。
	ImgUrl *string `json:"img_url,omitempty"`
	// LinkType 是跳转类型，网页:HTTP、小程序:APP，不传默认HTTP。
	LinkType *string `json:"link_type,omitempty"`
	// LinkUrl 是跳转链接，点击营销位内容图片跳到的链接url。
	LinkUrl *string `json:"link_url,omitempty"`
	// SubTitle 是内容说明。
	SubTitle *string `json:"sub_title,omitempty"`
	// Title 是内容标题。
	Title *string `json:"title,omitempty"`
}
