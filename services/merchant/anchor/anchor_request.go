package anchor

// MerchantAnchorInfoQueryRequest is the query request model for alipay.merchant.anchor.info.query.
type MerchantAnchorInfoQueryRequest struct {
	// UserId 是蚂蚁统一会员ID
	UserId *string `json:"user_id,omitempty" query:"user_id,omitempty"`
	// OpenId 是蚂蚁统一会员ID
	OpenId *string `json:"open_id,omitempty" query:"open_id,omitempty"`
	// PublicId 是生活号ID
	PublicId *string `json:"public_id,omitempty" query:"public_id,omitempty"`
}
