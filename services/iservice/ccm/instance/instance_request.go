// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package instance

// CcmInstanceCreateRequest 是支付宝官方 V3 OpenAPI 模型。
type CcmInstanceCreateRequest struct {
	// CreatorId 是创建人id
	CreatorId *string `json:"creator_id,omitempty"`
	// Description 是租户实例（数据权限）描述信息
	Description *string `json:"description,omitempty"`
	// ExternalId 是外部id
	ExternalId *string `json:"external_id,omitempty"`
	// Name 是租户实例（数据权限）名称，如果名称已经存在，将创建失败
	Name *string `json:"name,omitempty"`
}

// CcmInstanceGetRequest is the query request model for alipay.iservice.ccm.instance.get.
type CcmInstanceGetRequest struct {
	// Id 是部门id（即租户实例ID、数据权限ID）
	Id *string `json:"id,omitempty" query:"id,omitempty"`
	// ExternalId 是外部id
	ExternalId *string `json:"external_id,omitempty" query:"external_id,omitempty"`
}

// CcmInstanceQueryRequest is the query request model for alipay.iservice.ccm.instance.query.
type CcmInstanceQueryRequest struct {
	// PageNum 是查询结果的页码，起始值为 1，默认值为 1
	PageNum *string `json:"page_num,omitempty" query:"page_num,omitempty"`
	// PageSize 是分页查询时设置的每页记录数，最大值 100 行，默认为 10
	PageSize *string `json:"page_size,omitempty" query:"page_size,omitempty"`
}
