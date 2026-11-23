package ec

// EnterpriseInfoDTO 是支付宝官方 V3 OpenAPI 模型。
type EnterpriseInfoDTO struct {
	// AccountId 是共同账户id
	AccountId *string `json:"account_id,omitempty"`
	// Activate 是因公签约状态
	Activate *string `json:"activate,omitempty"`
	// AuthLevel 是企业认证等级
	AuthLevel *string `json:"auth_level,omitempty"`
	// AuthStatus 是认证状态
	AuthStatus *string `json:"auth_status,omitempty"`
	// AuthTime 是认证时间
	AuthTime *string `json:"auth_time,omitempty"`
	// EnterpriseAlias 是企业简称
	EnterpriseAlias *string `json:"enterprise_alias,omitempty"`
	// EnterpriseAliasAuditStatus 是企业简称审核状态	PASSED： 通过	REJECTED： 驳回	AUDITING： 审核中
	EnterpriseAliasAuditStatus *string `json:"enterprise_alias_audit_status,omitempty"`
	// EnterpriseCode 是统一社会信用码
	EnterpriseCode *string `json:"enterprise_code,omitempty"`
	// EnterpriseEmail 是企业邮箱
	EnterpriseEmail *string `json:"enterprise_email,omitempty"`
	// EnterpriseId 是企业id
	EnterpriseId *string `json:"enterprise_id,omitempty"`
	// EnterpriseName 是企业名称
	EnterpriseName *string `json:"enterprise_name,omitempty"`
	// GmtCreate 是创建时间
	GmtCreate *string `json:"gmt_create,omitempty"`
	// GmtModified 是更新时间
	GmtModified *string `json:"gmt_modified,omitempty"`
	// IotGroupId 是封闭场景（如班车）的人脸库id
	IotGroupId *string `json:"iot_group_id,omitempty"`
	// IotLogicGroupId 是开放场景（如团餐）的人脸库id
	IotLogicGroupId *string `json:"iot_logic_group_id,omitempty"`
	// OutBizNo 是服务商生成的请求唯一流水号/业务幂等号
	OutBizNo *string `json:"out_biz_no,omitempty"`
	// PlatformCode 是外部平台code
	PlatformCode *string `json:"platform_code,omitempty"`
	// PlatformOutId 是外部平台id
	PlatformOutId *string `json:"platform_out_id,omitempty"`
	// SignStatus 是企业码服务协议签约状态
	SignStatus *string `json:"sign_status,omitempty"`
	// Status 是企业当前状态
	Status *string `json:"status,omitempty"`
}
