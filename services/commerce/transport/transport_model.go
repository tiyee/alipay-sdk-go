// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package transport

// ParkingChargeInfo 是支付宝官方 V3 OpenAPI 模型。
type ParkingChargeInfo struct {
	// ChargeFee 是计费金额；单位元；
	ChargeFee *string `json:"charge_fee,omitempty"`
	// ChargeNode 是START_CHARGING：开始计费	FEE_INCREASE：费用增加	TO_TOP：日封顶
	ChargeNode *string `json:"charge_node,omitempty"`
	// ChargeTime 是开始计费或者费用增加的时间；传东八区时间；
	ChargeTime *string `json:"charge_time,omitempty"`
}

// ParkingMembershipInfo 是支付宝官方 V3 OpenAPI 模型。
type ParkingMembershipInfo struct {
	// MembershipBalance 是周卡、月卡：剩余天数（包括今天），单位日；	次卡：剩余次数（包括本次），单位次；	储值：余额，单位元；	折扣：折扣力度，代表0-10折；0折即为免费；
	MembershipBalance *string `json:"membership_balance,omitempty"`
	// MembershipType 是会员类型；	WEEK_CARD：周卡	MONTH_CARD：月卡	SEASON_CARD：季卡	YEAR_CARD：年卡	MULTIPLE_PASS_CARD：次卡	MONEY_STORED_CARD：储值	DISCOUNT_CARD：折扣
	MembershipType *string `json:"membership_type,omitempty"`
}

// ParkingPaymentDiscountInfo 是支付宝官方 V3 OpenAPI 模型。
type ParkingPaymentDiscountInfo struct {
	// DiscountAmount 是折扣金额；单位元；
	DiscountAmount *string `json:"discount_amount,omitempty"`
	// DiscountType 是折扣类型；	COMPREHENSIVE：综合；	DISCOUNT：打折；	VOUCHER：券；	POINTS：积分；
	DiscountType *string `json:"discount_type,omitempty"`
}
