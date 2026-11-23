// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package query

// ActivityQueryRequest is the query request model for alipay.marketing.activity.query.
type ActivityQueryRequest struct {
	// ActivityId 是活动id
	ActivityId string `json:"-" path:"activity_id"`
	// MerchantId 是商户PID,默认为当前接口调用商户。
	MerchantId *string `json:"merchant_id,omitempty" query:"merchant_id,omitempty"`
	// MerchantAccessMode 是商户接入模式。
	MerchantAccessMode *string `json:"merchant_access_mode,omitempty" query:"merchant_access_mode,omitempty"`
}
