package info

// MobilePublicInfoQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type MobilePublicInfoQueryResponse struct {
	// AppName 是服务窗名称
	AppName *string `json:"app_name,omitempty"`
	// AuditDesc 是服务窗审核状态描述，如果审核驳回则有相关的驳回理由
	AuditDesc *string `json:"audit_desc,omitempty"`
	// AuditStatus 是服务窗审核状态，对于系统商而言，只有三个状态，AUDITING：审核中，AUDIT_FAILED：审核驳回，AUDIT_SUCCESS：审核通过
	AuditStatus *string `json:"audit_status,omitempty"`
	// IsOnline 是服务窗是否上线，T表示上线，F表示未上线
	IsOnline *string `json:"is_online,omitempty"`
	// IsRelease 是服务窗是否上架，T表示上架，上架即可在支付宝客户端被搜索到，F表示未上架
	IsRelease *string `json:"is_release,omitempty"`
	// LogoUrl 是服务窗头像地址
	LogoUrl *string `json:"logo_url,omitempty"`
	// PublicGreeting 是服务窗欢迎语
	PublicGreeting *string `json:"public_greeting,omitempty"`
}
