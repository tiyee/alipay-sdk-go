// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package topic

// PublicTopicCreateRequest 是支付宝官方 V3 OpenAPI 模型。
type PublicTopicCreateRequest struct {
	// ImgUrl 是营销位图片url, 尺寸为996*450，最大不超过5M，支持格式:.jpg、.png ，请先调用<a href="https://docs.open.alipay.com/api_3/alipay.offline.material.image.upload"> 图片上传接口</a>获得图片url。营销位需要展示头图时，必须填写该参数。。
	ImgUrl *string `json:"img_url,omitempty"`
	// LinkType 是跳转类型，网页:HTTP、小程序:APP，不传默认HTTP。
	LinkType *string `json:"link_type,omitempty"`
	// LinkUrl 是营销位跳转地址，点击营销位头图跳到的链接url。营销位需要展示头图时，必须填写该参数。。
	LinkUrl *string `json:"link_url,omitempty"`
	// SubTitle 是营销位描述。营销位需要展示头图时，必须填写该参数。。
	SubTitle *string `json:"sub_title,omitempty"`
	// Title 是营销位名称。
	Title *string `json:"title,omitempty"`
	// TopicItems 是营销位内容，数量限制：大于4个，小于8个。
	TopicItems *[]TopicItem `json:"topic_items,omitempty"`
}

// PublicTopicModifyRequest 是支付宝官方 V3 OpenAPI 模型。
type PublicTopicModifyRequest struct {
	// ImgUrl 是营销位图片url, 尺寸为996*450，最大不超过5M，支持格式:.jpg、.png ，请先调用<a href="https://docs.open.alipay.com/api_3/alipay.offline.material.image.upload"> 图片上传接口</a>获得图片url。营销位需要展示头图时，必须填写该参数。。
	ImgUrl *string `json:"img_url,omitempty"`
	// LinkType 是跳转类型，网页:HTTP、小程序:APP，不传默认HTTP。
	LinkType *string `json:"link_type,omitempty"`
	// LinkUrl 是营销位跳转地址，点击营销位头图跳到的链接url。营销位需要展示头图时，必须填写该参数。。
	LinkUrl *string `json:"link_url,omitempty"`
	// SubTitle 是营销位描述。营销位需要展示头图时，必须填写该参数。
	SubTitle *string `json:"sub_title,omitempty"`
	// Title 是营销位名称。
	Title *string `json:"title,omitempty"`
	// TopicId 是营销位id。
	TopicId    *string    `json:"topic_id,omitempty"`
	TopicItems *TopicItem `json:"topic_items,omitempty"`
}

// PublicTopicBatchQueryRequest is the request type for alipay.open.public.topic.batchquery.
type PublicTopicBatchQueryRequest struct{}

// PublicTopicDeleteRequest is the request type for alipay.open.public.topic.delete.
type PublicTopicDeleteRequest struct {
	// TopicId 是营销位id
	TopicId *string `json:"topic_id,omitempty" query:"topic_id,omitempty"`
}
