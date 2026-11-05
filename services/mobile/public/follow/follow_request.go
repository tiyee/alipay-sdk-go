package follow

// MobilePublicFollowListRequest is the query request model for alipay.mobile.public.follow.list.
type MobilePublicFollowListRequest struct {
	// BizContent 是当nextUserId为空时,代表查询第一组,如果有值时以当前值为准查询下一组 	详情请见
	BizContent *string `json:"biz_content,omitempty" query:"biz_content,omitempty"`
}
