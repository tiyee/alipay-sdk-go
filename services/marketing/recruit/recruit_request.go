// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package recruit

// RecruitEnrollCloseRequest 是支付宝官方 V3 OpenAPI 模型。
type RecruitEnrollCloseRequest struct {
	// EnrollId 是活动报名ID
	EnrollId *string `json:"enroll_id,omitempty"`
}

// RecruitEnrollCreateRequest 是支付宝官方 V3 OpenAPI 模型。
type RecruitEnrollCreateRequest struct {
	EnrollInfo *RecruitEnrollInfo `json:"enroll_info,omitempty"`
	// EnrollSceneType 是活动报名场景配置。	枚举值：	VOUCHER：券报名场景的报名；	MINI_APP：小程序报名场景的报名；
	EnrollSceneType *string `json:"enroll_scene_type,omitempty"`
	// OutBizNo 是外部操作流水号。由商家/ISV 自定义，仅支持字母、数字、下划线且需保证每次操作唯一。
	OutBizNo *string `json:"out_biz_no,omitempty"`
	// PlanId 是要报名的招商方案ID
	PlanId *string `json:"plan_id,omitempty"`
}

// RecruitPlanQueryRequest 是支付宝官方 V3 OpenAPI 模型。
type RecruitPlanQueryRequest struct {
	EnrollMerchant *RecruitEnrollMerchant `json:"enroll_merchant,omitempty"`
	// PlanId 是招商方案ID
	PlanId *string `json:"plan_id,omitempty"`
}

// RecruitPlanListQueryRequest 是支付宝官方 V3 OpenAPI 模型。
type RecruitPlanListQueryRequest struct {
	EnrollMerchant *RecruitEnrollMerchant `json:"enroll_merchant,omitempty"`
	// EnrollSceneType 是活动报名场景配置。 	枚举值：	VOUCHER：券报名场景查询；	MINI_APP：小程序报名场景查询；	VOUCHER,MINI_APP：同时进行券报名场景和小程序报名场景的查询；
	EnrollSceneType *string `json:"enroll_scene_type,omitempty"`
	// PageNum 是第几页，默认1（需要在1和999之间）
	PageNum *int64 `json:"page_num,omitempty"`
	// PageSize 是每页记录条数，默认20，需要在1和100之间
	PageSize *int64 `json:"page_size,omitempty"`
}

// RecruitEnrollQueryRequest is the query request model for alipay.marketing.recruit.enroll.query.
type RecruitEnrollQueryRequest struct {
	// OutBizNo 是外部操作流水号，创建招商报名时传入。由商家/ISV 自定义，仅支持字母、数字、下划线且需保证每次操作唯一。
	OutBizNo *string `json:"out_biz_no,omitempty" query:"out_biz_no,omitempty"`
	// EnrollId 是报名ID，此参数和out_biz_no至少传一个，优先取enroll_id
	EnrollId *string `json:"enroll_id,omitempty" query:"enroll_id,omitempty"`
}
