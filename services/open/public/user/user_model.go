// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package user

// UserAnalysisData 是支付宝官方 V3 OpenAPI 模型。
type UserAnalysisData struct {
	// CancelUserCnt 是取消关注人数。
	CancelUserCnt *string `json:"cancel_user_cnt,omitempty"`
	// CumulateUserCnt 是累积关注人数。
	CumulateUserCnt *string `json:"cumulate_user_cnt,omitempty"`
	// Date 是日期。
	Date *string `json:"date,omitempty"`
	// GrowUserCnt 是净增关注人数。
	GrowUserCnt *string `json:"grow_user_cnt,omitempty"`
	// NewUserCnt 是新关注人数。
	NewUserCnt *string `json:"new_user_cnt,omitempty"`
}
