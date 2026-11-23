// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package info

// PublicInfoModifyResponse 是支付宝官方 V3 OpenAPI 模型。
type PublicInfoModifyResponse struct {
	// AuditDesc 是服务窗审核状态描述。
	AuditDesc *string `json:"audit_desc,omitempty"`
	// AuditStatus 是服务窗审核状态，申请成功后返回AUDITING，等待风控审核。
	AuditStatus *string `json:"audit_status,omitempty"`
}

// PublicInfoQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type PublicInfoQueryResponse struct {
	// AppName 是生活号名称。
	AppName *string `json:"app_name,omitempty"`
	// AuditDesc 是最新审核状态描述，如果审核驳回则有相关的驳回理由。
	AuditDesc *string `json:"audit_desc,omitempty"`
	// AuditStatus 是最新审核状态，对于系统商而言，只有四个状态，AUDITING：审核中，AUDIT_FAILED：审核驳回，AUDIT_SUCCESS：审核通过，AUDIT_NORMAL：无审核状态（当前没有处于审核过程的工单）。
	AuditStatus *string `json:"audit_status,omitempty"`
	// AuditStatusList 是审核状态表，目前支持名称、头像、名称与头像、简介审核状态。
	AuditStatusList *[]PublicAuditStatus `json:"audit_status_list,omitempty"`
	// BackgroundUrl 是背景图片地址。
	BackgroundUrl *string `json:"background_url,omitempty"`
	// Introduction 是生活号简介。
	Introduction *string `json:"introduction,omitempty"`
	// IsOnline 是生活号是否上线，T表示上线，F表示未上线。
	IsOnline *string `json:"is_online,omitempty"`
	// IsRelease 是生活号是否上架，T表示上架，上架即可在支付宝客户端被搜索到，F表示未上架。
	IsRelease *string `json:"is_release,omitempty"`
	// LogoUrl 是生活号头像地址。
	LogoUrl *string `json:"logo_url,omitempty"`
	// MccCodeDesc 是商家经营类目，详情参见 <a href="https://opendocs.alipay.com/open/01n22g">商家经营类目</a>。
	MccCodeDesc *string `json:"mcc_code_desc,omitempty"`
	// PublicGreeting 是欢迎语。
	PublicGreeting *string `json:"public_greeting,omitempty"`
}
