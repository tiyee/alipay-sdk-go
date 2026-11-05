// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package instance

// Instance 是支付宝官方 V3 OpenAPI 模型。
type Instance struct {
	// CreateTime 是?创建时间
	CreateTime *string `json:"create_time,omitempty"`
	// CreatorId 是创建人id
	CreatorId *string `json:"creator_id,omitempty"`
	// Description 是部门描述
	Description *string `json:"description,omitempty"`
	// ExternalId 是外部id
	ExternalId *string `json:"external_id,omitempty"`
	// Id 是部门id（即租户实例ID、数据权限ID）
	Id *string `json:"id,omitempty"`
	// Name 是租户实例（数据权限）名称
	Name *string `json:"name,omitempty"`
	// Status 是租户实例（数据权限）状态，所有可能的状态如下：INIT（初始化）、STARTED（准备完成）、START_FAILED（准备失败）、STOPPING（停用中）、STOPPED（已停用）、RELEASING（释放中）、RELEASED（已释放）、RESUMING（重新启动）、MODIFING（修改规格中）、DELETED（已删除）。
	Status *string `json:"status,omitempty"`
	// UpdateTime 是最后修改时间
	UpdateTime *string `json:"update_time,omitempty"`
	// UpdaterId 是最后修改人id
	UpdaterId *string `json:"updater_id,omitempty"`
}
