package api

// AppAPIFieldApply 是支付宝官方 V3 OpenAPI 模型。
type AppAPIFieldApply struct {
	AuthFieldApply *AuthFieldApply `json:"auth_field_apply,omitempty"`
}

// AuthAPIDTO 是支付宝官方 V3 OpenAPI 模型。
type AuthAPIDTO struct {
	// ApiName 是敏感信息可申请接口名。
	ApiName *string `json:"api_name,omitempty"`
	// FieldName 是敏感信息申请字段。
	FieldName *string `json:"field_name,omitempty"`
	// PackageCode 是敏感信息申请能力code值。
	PackageCode *string `json:"package_code,omitempty"`
}

// AuthFieldApply 是支付宝官方 V3 OpenAPI 模型。
type AuthFieldApply struct {
	// ApiName 是接口英文名称，通过alipay.open.app.api.query接口查询获得。。
	ApiName *string `json:"api_name,omitempty"`
	// CustomerAnswer 是贵公司是否有自己的客服渠道，能及时响应和处理舆情数量是多少？。
	CustomerAnswer *string `json:"customer_answer,omitempty"`
	// FieldName 是字段英文名称，通过alipay.open.app.api.query接口查询获得。。
	FieldName *string `json:"field_name,omitempty"`
	// Memo 是备注。
	Memo *string `json:"memo,omitempty"`
	// PackageCode 是功能code，通过alipay.open.app.api.query接口查询获得。
	PackageCode *string `json:"package_code,omitempty"`
	// QpsAnswer 是接入后一年内预计接口秒级调用量峰值是多少？（最高峰值：1000QPS）。
	QpsAnswer *string `json:"qps_answer,omitempty"`
	// SceneCode 是场景code，alipay.open.app.api.scene.query 接口查询获得。
	SceneCode *string `json:"scene_code,omitempty"`
	// TinyAppTemplateId 是当为使用使用模板的小程序申请时可传入所使用的小程序模板id。
	TinyAppTemplateId *string `json:"tiny_app_template_id,omitempty"`
}

// AuthFieldDTO 是支付宝官方 V3 OpenAPI 模型。
type AuthFieldDTO struct {
	// ApiName 是接口英文名称。
	ApiName *string `json:"api_name,omitempty"`
	// FieldName 是接口字段英文名称。
	FieldName *string `json:"field_name,omitempty"`
	// PackageCode 是接口归属的功能code。
	PackageCode *string `json:"package_code,omitempty"`
	// Reason 是驳回原因。
	Reason *string `json:"reason,omitempty"`
	// Status 是AUDIT 审核中，AGREE通过，REJECT驳回，INVALID无效（isv代申请场景)。
	Status *string `json:"status,omitempty"`
	// UserAppId 是用户应用app_id。
	UserAppId *string `json:"user_app_id,omitempty"`
}

// AuthFieldSceneDTO 是支付宝官方 V3 OpenAPI 模型。
type AuthFieldSceneDTO struct {
	// SceneCode 是场景CODE。
	SceneCode *string `json:"scene_code,omitempty"`
	// SceneDesc 是场景code描述。
	SceneDesc *string `json:"scene_desc,omitempty"`
}
