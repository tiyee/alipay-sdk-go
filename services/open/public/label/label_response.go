// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package label

// PublicLabelCreateResponse 是支付宝官方 V3 OpenAPI 模型。
type PublicLabelCreateResponse struct {
	// Id 是支付宝返回的标签ID，开发者后续通过该ID进行标签的查询、修改、删除等操作。
	Id *int64 `json:"id,omitempty"`
	// Name 是开发者传入name的值。
	Name *string `json:"name,omitempty"`
}

// PublicLabelQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type PublicLabelQueryResponse struct {
	// LabelList 是该服务窗拥有的标签列表。
	LabelList *[]PublicLabel `json:"label_list,omitempty"`
}

// PublicLabelUserQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type PublicLabelUserQueryResponse struct {
	// LabelIds 是用户标签id列表，以英文逗号分隔。
	LabelIds *string `json:"label_ids,omitempty"`
}
