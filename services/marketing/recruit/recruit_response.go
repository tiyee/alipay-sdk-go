// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package recruit

import (
	"encoding/json"
)

// RecruitEnrollCloseResponse 是支付宝官方 V3 OpenAPI 模型。
type RecruitEnrollCloseResponse = json.RawMessage

// RecruitEnrollCreateResponse 是支付宝官方 V3 OpenAPI 模型。
type RecruitEnrollCreateResponse struct {
	// EnrollId 是报名ID
	EnrollId *string `json:"enroll_id,omitempty"`
	// OutBizNo 是外部操作流水号。由商家/ISV 自定义，仅支持字母、数字、下划线且需保证每次操作唯一。
	OutBizNo *string `json:"out_biz_no,omitempty"`
}

// RecruitEnrollQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type RecruitEnrollQueryResponse struct {
	// EnrollId 是报名ID
	EnrollId   *string            `json:"enroll_id,omitempty"`
	EnrollInfo *RecruitEnrollInfo `json:"enroll_info,omitempty"`
	// EnrollSceneType 是报名所属的活动报名场景。	枚举值： VOUCHER：属于券报名场景；	MINI_APP：属于小程序报名场景；
	EnrollSceneType *string `json:"enroll_scene_type,omitempty"`
	// EnrollTime 是报名时间
	EnrollTime *string `json:"enroll_time,omitempty"`
	// Memo 是备注字段
	Memo *string `json:"memo,omitempty"`
	// OutBizNo 是外部操作流水号。由商家/ISV 自定义，仅支持字母、数字、下划线且需保证每次操作唯一。
	OutBizNo *string `json:"out_biz_no,omitempty"`
	// PlanId 是招商方案ID
	PlanId *string `json:"plan_id,omitempty"`
	// Status 是报名状态： AUDITING 审核中； ENABLED 通过 ； REJECTED 不通过 ； CLOSING 下线审核中 ； CLOSED 已下线；
	Status *string `json:"status,omitempty"`
}

// RecruitPlanQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type RecruitPlanQueryResponse struct {
	// Description 是招商方案描述：富文本内容，可能包含html标签
	Description *string `json:"description,omitempty"`
	// EnrollEndTime 是报名结束时间
	EnrollEndTime *string `json:"enroll_end_time,omitempty"`
	// EnrollRules 是"招商报名规则。 说明： 1. 查询到的规则将用于 alipay.marketing.recruit.enroll.create(招商报名提交)接口。 2. enroll_rules.required 为 true 时，enroll_rules.type 对应的内容类型在报名提交时必填。"
	EnrollRules *[]RecruitEnrollRule `json:"enroll_rules,omitempty"`
	// EnrollSceneType 是招商方案所属报名场景 枚举值： VOUCHER：券报名场景； MINI_APP：小程序报名场景；
	EnrollSceneType *string `json:"enroll_scene_type,omitempty"`
	// EnrollStartTime 是报名开始时间
	EnrollStartTime *string `json:"enroll_start_time,omitempty"`
	// Logo 是方案头图链接，150x150 以上尺寸的方形图片
	Logo *string `json:"logo,omitempty"`
	// PlanId 是招商方案ID
	PlanId *string `json:"plan_id,omitempty"`
	// PlanName 是招商方案名称
	PlanName *string `json:"plan_name,omitempty"`
	// Status 是方案状态: 不可用:DISABLED;可用:ENABLED
	Status *string `json:"status,omitempty"`
}

// RecruitPlanListQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type RecruitPlanListQueryResponse struct {
	// Data 是方案列表
	Data *[]RecruitPlanLight `json:"data,omitempty"`
	// PageNum 是第几页，默认1（从1开始计数）
	PageNum *int64 `json:"page_num,omitempty"`
	// PageSize 是每页记录条数，默认20
	PageSize *int64 `json:"page_size,omitempty"`
	// Total 是总数
	Total *int64 `json:"total,omitempty"`
}
