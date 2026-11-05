// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package group

// PublicGroupCreateRequest 是支付宝官方 V3 OpenAPI 模型。
type PublicGroupCreateRequest struct {
	// LabelRule 是标签规则，满足该规则的粉丝将被圈定，标签id不能重复。
	LabelRule *[]ComplexLabelRule `json:"label_rule,omitempty"`
	// Name 是分组名称，仅支持中文、字母、数字、下划线的组合。。
	Name *string `json:"name,omitempty"`
}

// PublicGroupCrowdQueryRequest 是支付宝官方 V3 OpenAPI 模型。
type PublicGroupCrowdQueryRequest struct {
	// LabelRule 是用户分组的规则项列表，规则项之间元素是与的逻辑，每个规则项内部用多个值表示或的逻辑。
	LabelRule *[]ComplexLabelRule `json:"label_rule,omitempty"`
}

// PublicGroupModifyRequest 是支付宝官方 V3 OpenAPI 模型。
type PublicGroupModifyRequest struct {
	// GroupId 是分组ID，通过<a href="https://opendocs.alipay.com/apis/api_6/alipay.open.public.group.create">alipay.open.public.group.create</a>（用户分组创建接口) 创建用户分组获取。。
	GroupId *string `json:"group_id,omitempty"`
	// LabelRule 是标签规则，满足该规则的粉丝将被圈定，标签id不能重复。
	LabelRule *[]ComplexLabelRule `json:"label_rule,omitempty"`
	// Name 是分组名称，仅支持中文、字母、数字、下划线的组合。。
	Name *string `json:"name,omitempty"`
}

// PublicGroupBatchQueryRequest is the request type for alipay.open.public.group.batchquery.
type PublicGroupBatchQueryRequest struct{}

// PublicGroupDeleteRequest is the request type for alipay.open.public.group.delete.
type PublicGroupDeleteRequest struct {
	// GroupId 是分组ID，通过alipay.open.public.group.create（用户分组创建接口)	创建用户分组获取。
	GroupId *string `json:"group_id,omitempty" query:"group_id,omitempty"`
}
