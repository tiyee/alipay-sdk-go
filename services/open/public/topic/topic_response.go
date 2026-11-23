// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package topic

// PublicTopicBatchQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type PublicTopicBatchQueryResponse struct {
	// TopicList 是营销位列表。
	TopicList *[]Topic `json:"topic_list,omitempty"`
}

// PublicTopicCreateResponse 是支付宝官方 V3 OpenAPI 模型。
type PublicTopicCreateResponse struct {
	// TopicId 是营销位id。
	TopicId *string `json:"topic_id,omitempty"`
}

// PublicTopicDeleteResponse 是支付宝官方 V3 OpenAPI 模型。
type PublicTopicDeleteResponse struct {
}

// PublicTopicModifyResponse 是支付宝官方 V3 OpenAPI 模型。
type PublicTopicModifyResponse struct {
}
