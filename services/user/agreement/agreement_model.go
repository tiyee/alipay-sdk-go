package agreement

// BusinessParamsMap 是支付宝官方 V3 OpenAPI 模型。
type BusinessParamsMap struct {
	// ChangeTime 是服务变更时间
	ChangeTime *string `json:"change_time,omitempty"`
	// NewAmount 是服务变更后新金额
	NewAmount *string `json:"new_amount,omitempty"`
	// OriAmount 是原始价格
	OriAmount *string `json:"ori_amount,omitempty"`
}

// ExecutionPlan 是支付宝官方 V3 OpenAPI 模型。
type ExecutionPlan struct {
	// ExecuteTime 是周期扣预期执行时间，格式为YYYY-MM-DD
	ExecuteTime *string `json:"execute_time,omitempty"`
	// LatestExecuteTime 是周期扣执行计划最晚执行时间，格式为YYYY-MM-DD
	LatestExecuteTime *string `json:"latest_execute_time,omitempty"`
	// PeriodId 是该值为自然数，表示周期扣期数。
	PeriodId *string `json:"period_id,omitempty"`
	// SingleAmount 是周期扣中单笔金额，单位是元
	SingleAmount *string `json:"single_amount,omitempty"`
}

// PeriodRuleParams 是支付宝官方 V3 OpenAPI 模型。
type PeriodRuleParams struct {
	// ExecuteTime 是首次执行时间execute_time是周期扣款产品必填，即商户发起首次扣款的时间。精确到日，格式为yyyy-MM-dd	结合其他必填的扣款周期参数，会确定商户以后的扣款计划。发起扣款的时间需符合这里的扣款计划。
	ExecuteTime *string `json:"execute_time,omitempty"`
	// IsTrialMode 是若签约中传该参数，且值为 Y ，则表示本次签约有试用期。若传值为 N 或不传该参数，则表示本次签约无试用期。
	IsTrialMode *string `json:"is_trial_mode,omitempty"`
	// Period 是周期数period是周期扣款产品必填。与另一参数period_type组合使用确定扣款周期，例如period_type为DAY，period=90，则扣款周期为90天。
	Period *int64 `json:"period,omitempty"`
	// PeriodType 是周期类型period_type是周期扣款产品必填，枚举值为DAY和MONTH。	DAY即扣款周期按天计，MONTH代表扣款周期按自然月。	与另一参数period组合使用确定扣款周期，例如period_type为DAY，period=30，则扣款周期为30天；period_type为MONTH，period=3，则扣款周期为3个自然月。	自然月是指，不论这个月有多少天，周期都计算到月份中的同一日期。例如1月3日到2月3日为一个自然月，1月3日到4月3日为三个自然月。注意周期类型使用MONTH的时候，计划扣款时间execute_time不允许传28日之后的日期（可以传28日），以此避免有些月份可能不存在对应日期的情况。
	PeriodType      *string          `json:"period_type,omitempty"`
	SceneRuleParams *SceneRuleParams `json:"scene_rule_params,omitempty"`
	// SingleAmount 是单次扣款最大金额single_amount是周期扣款产品必填，即每次发起扣款时限制的最大金额，单位为元。商户每次发起扣款都不允许大于此金额。
	SingleAmount *string `json:"single_amount,omitempty"`
	// TotalAmount 是总金额限制，单位为元。如果传入此参数，商户多次扣款的累计金额不允许超过此金额。
	TotalAmount *string `json:"total_amount,omitempty"`
	// TotalPayments 是总扣款次数。如果传入此参数，则商户成功扣款的次数不能超过此次数限制（扣款失败不计入）。
	TotalPayments *int64 `json:"total_payments,omitempty"`
}

// SceneRuleParams 是支付宝官方 V3 OpenAPI 模型。
type SceneRuleParams struct {
	// DiscountPeriod 是在周期扣场景化模板中配置优惠类型为优惠期玩法时需要该参数，表示后续有多少期扣款可享受优惠，值为自然数代表周期
	DiscountPeriod *string `json:"discount_period,omitempty"`
	// LowPricePeriod 是在周期扣场景化模板中配置优惠类型为低价玩法时需要该参数，表示代扣低价期持续的时间。单位是天，该值为自然数
	LowPricePeriod *string `json:"low_price_period,omitempty"`
}
