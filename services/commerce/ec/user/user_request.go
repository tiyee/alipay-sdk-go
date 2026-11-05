package user

// EcUserEnterpriseQueryRequest is the query request model for alipay.commerce.ec.user.enterprise.query.
type EcUserEnterpriseQueryRequest struct {
	// UserId 是用户id
	UserId *string `json:"user_id,omitempty" query:"user_id,omitempty"`
	// OpenId 是用户open_id
	OpenId *string `json:"open_id,omitempty" query:"open_id,omitempty"`
}
