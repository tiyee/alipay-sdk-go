package members

// AppMembersCreateResponse 是支付宝官方 V3 OpenAPI 模型。
type AppMembersCreateResponse struct {
}

// AppMembersDeleteResponse 是支付宝官方 V3 OpenAPI 模型。
type AppMembersDeleteResponse struct {
}

// AppMembersQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type AppMembersQueryResponse struct {
	// AppMemberInfoList 是小程序成员模型。
	AppMemberInfoList *[]AppMemberInfo `json:"app_member_info_list,omitempty"`
}
