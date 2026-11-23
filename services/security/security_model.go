// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package security

// SecurityProdSssQuery 是支付宝官方 V3 OpenAPI 模型。
type SecurityProdSssQuery struct {
	// Tesst 是stst
	Tesst *[]PromiseDetail `json:"tesst,omitempty"`
	Xxx   *JinyouTestFive  `json:"xxx,omitempty"`
}

// JinyouTestFive 是支付宝官方 V3 OpenAPI 模型。
type JinyouTestFive struct {
	F1F *JinyouTestFour  `json:"f_1_f,omitempty"`
	F2F *JinyouTestThree `json:"f_2_f,omitempty"`
	// F3Openid 是1
	F3Openid *string `json:"f_3_openid,omitempty"`
	// F3Y 是1
	F3Y *string `json:"f_3_y,omitempty"`
	// F4N 是3
	F4N *string `json:"f_4_n,omitempty"`
}

// JinyouTestFour 是支付宝官方 V3 OpenAPI 模型。
type JinyouTestFour struct {
	// F1Openid 是1
	F1Openid *string `json:"f_1_openid,omitempty"`
	// F1Y 是1
	F1Y *string          `json:"f_1_y,omitempty"`
	F2F *JinyouTestThree `json:"f_2_f,omitempty"`
	// F3Openid 是1
	F3Openid *string `json:"f_3_openid,omitempty"`
	// F3Y 是1
	F3Y *int64 `json:"f_3_y,omitempty"`
}

// JinyouTestOne 是支付宝官方 V3 OpenAPI 模型。
type JinyouTestOne struct {
	// O1N 是1
	O1N *string `json:"o_1_n,omitempty"`
	// O2Openid 是o_2_y
	O2Openid *string `json:"o_2_openid,omitempty"`
	// O2Y 是o_2_y
	O2Y *string `json:"o_2_y,omitempty"`
	// O3Openid 是o_3_y
	O3Openid *string `json:"o_3_openid,omitempty"`
	// O3Y 是o_3_y
	O3Y *int64 `json:"o_3_y,omitempty"`
}

// JinyouTestThree 是支付宝官方 V3 OpenAPI 模型。
type JinyouTestThree struct {
	Th1F *JinyouTestTwo `json:"th_1_f,omitempty"`
	// Th2N 是无枚举值
	Th2N *string `json:"th_2_n,omitempty"`
	// Th3N 是1
	Th3N *bool `json:"th_3_n,omitempty"`
}

// JinyouTestTwo 是支付宝官方 V3 OpenAPI 模型。
type JinyouTestTwo struct {
	// T1Openid 是1
	T1Openid *string `json:"t_1_openid,omitempty"`
	// T1Y 是1
	T1Y *string        `json:"t_1_y,omitempty"`
	T2F *JinyouTestOne `json:"t_2_f,omitempty"`
	// T3N 是1
	T3N *string `json:"t_3_n,omitempty"`
}

// PromiseDetail 是支付宝官方 V3 OpenAPI 模型。
type PromiseDetail struct {
	// AuthStatus 是授权状态
	AuthStatus *bool `json:"auth_status,omitempty"`
	// CreateTime 是任务创建时间
	CreateTime *string `json:"create_time,omitempty"`
	// EndTime 是任务结束时间
	EndTime *string `json:"end_time,omitempty"`
	// FinalTime 是任务到达终态的时间
	FinalTime *string `json:"final_time,omitempty"`
	// FinishPeriods 是任务完成期数
	FinishPeriods *int64 `json:"finish_periods,omitempty"`
	// MerchantId 是芝麻侧的商户id
	MerchantId *string `json:"merchant_id,omitempty"`
	// MerchantLogo 是商户logo
	MerchantLogo *string `json:"merchant_logo,omitempty"`
	// MerchantName 是商户名称
	MerchantName *string `json:"merchant_name,omitempty"`
	// OutBizNo 是生活记录加入时的外部业务号
	OutBizNo *string `json:"out_biz_no,omitempty"`
	// PeriodType 是周期类型
	PeriodType *string `json:"period_type,omitempty"`
	// PromiseName 是生活记录模板名称
	PromiseName *string `json:"promise_name,omitempty"`
	// RecordId 是生活记录主记录id
	RecordId *string `json:"record_id,omitempty"`
	// RecordStatus 是主任务状态
	RecordStatus *string `json:"record_status,omitempty"`
	// StartTime 是任务开始时间
	StartTime *string `json:"start_time,omitempty"`
	// SubRecordStatus 是子记录状态
	SubRecordStatus *string `json:"sub_record_status,omitempty"`
	// SubTitle 是副标题
	SubTitle *string `json:"sub_title,omitempty"`
	// TemplateId 是生活记录模板id
	TemplateId *string `json:"template_id,omitempty"`
	// TotalPeriods 是任务总期数
	TotalPeriods *int64 `json:"total_periods,omitempty"`
}
