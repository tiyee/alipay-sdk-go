package search

// SearchAbilityBatchQueryRequest 是支付宝官方 V3 OpenAPI 模型。
type SearchAbilityBatchQueryRequest struct {
	BizData *SearchOperPageQueryRequest `json:"biz_data,omitempty"`
	// BizType 是表示修改内容的业务类型。
	BizType *string `json:"biz_type,omitempty"`
	// OptType 是进行具体操作的类型比如：query。
	OptType *string `json:"opt_type,omitempty"`
}

// SearchOperPageQueryRequest 是支付宝官方 V3 OpenAPI 模型。
type SearchOperPageQueryRequest struct {
	// AccessType 是申请类型 BASE：基础信息， BRAND_BOX：品牌直达，SERVICE_BOX服务直达。
	AccessType *string `json:"access_type,omitempty"`
	// Appid 是小程序ID。
	Appid *string `json:"appid,omitempty"`
	// PageNum 是当前页。
	PageNum *string `json:"page_num,omitempty"`
	// PageSize 是每页显示条数。
	PageSize *string `json:"page_size,omitempty"`
	// SceneCode 是场景码。
	SceneCode *string `json:"scene_code,omitempty"`
	// SpecCode 是服务类型 SP_MINI_APP 小程序 SP_PUBLIC_APP 生活号。
	SpecCode *string `json:"spec_code,omitempty"`
}
