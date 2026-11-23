// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package user

// ActivityUserBatchQueryVoucherRequest is the query request model for alipay.marketing.activity.user.batchqueryvoucher.
type ActivityUserBatchQueryVoucherRequest struct {
	// UserId 是支付宝用户 id
	UserId *string `json:"user_id,omitempty" query:"user_id,omitempty"`
	// OpenId 是支付宝用户openId
	OpenId *string `json:"open_id,omitempty" query:"open_id,omitempty"`
	// ActivityId 是活动 id
	ActivityId *string `json:"activity_id,omitempty" query:"activity_id,omitempty"`
	// BelongMerchantId 是券归属商户
	BelongMerchantId *string `json:"belong_merchant_id,omitempty" query:"belong_merchant_id,omitempty"`
	// SenderMerchantId 是券发放商户
	SenderMerchantId *string `json:"sender_merchant_id,omitempty" query:"sender_merchant_id,omitempty"`
	// VoucherStatus 是券状态，其中已过期状态可查询6个月内的数据
	VoucherStatus *string `json:"voucher_status,omitempty" query:"voucher_status,omitempty"`
	// PageNum 是分页查询页码
	PageNum *int64 `json:"page_num,omitempty" query:"page_num,required"`
	// MerchantAccessMode 是商户接入模式
	MerchantAccessMode *string `json:"merchant_access_mode,omitempty" query:"merchant_access_mode,omitempty"`
	// PageSize 是分页查询单页数据条数
	PageSize *int64 `json:"page_size,omitempty" query:"page_size,required"`
}

// ActivityUserQueryVoucherRequest is the query request model for alipay.marketing.activity.user.queryvoucher.
type ActivityUserQueryVoucherRequest struct {
	// MerchantId 是商户 PID，默认为当前接口调用商户。
	MerchantId *string `json:"merchant_id,omitempty" query:"merchant_id,omitempty"`
	// UserId 是支付宝user_id账号
	UserId *string `json:"user_id,omitempty" query:"user_id,omitempty"`
	// OpenId 是领券的支付宝用户openId
	OpenId *string `json:"open_id,omitempty" query:"open_id,omitempty"`
	// ActivityId 是活动id
	ActivityId *string `json:"activity_id,omitempty" query:"activity_id,omitempty"`
	// VoucherId 是用户券 id。支付宝为用户优惠券唯一分配的 id。
	VoucherId *string `json:"voucher_id,omitempty" query:"voucher_id,omitempty"`
	// VoucherCode 是用户领取的商家券券码。
	VoucherCode *string `json:"voucher_code,omitempty" query:"voucher_code,omitempty"`
	// MerchantAccessMode 是商户接入模式。
	MerchantAccessMode *string `json:"merchant_access_mode,omitempty" query:"merchant_access_mode,omitempty"`
}
