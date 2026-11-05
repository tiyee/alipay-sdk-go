// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package violation

// AuditEvidenceInfo 是支付宝官方 V3 OpenAPI 模型。
type AuditEvidenceInfo struct {
	// FileType 是文件类型。包括.jpg .png .jpeg .pdf等。
	FileType *string `json:"file_type,omitempty"`
	// Url 是文件url。
	Url *string `json:"url,omitempty"`
}

// ReplyRecord 是支付宝官方 V3 OpenAPI 模型。
type ReplyRecord struct {
	// Content 是回复内容。
	Content *string `json:"content,omitempty"`
	// GmtCreate 是回复时间。
	GmtCreate *string `json:"gmt_create,omitempty"`
	// Images 是回复图片。
	Images *[]string `json:"images,omitempty"`
	// ReplierRole 是回复角色 商家：MERCHANT 审核小二：AUDITOR。
	ReplierRole *string `json:"replier_role,omitempty"`
	// Status 是状态。只在replier_role = ‘AUDITOR’时有值。
	Status *string `json:"status,omitempty"`
}

// ViolationEvent 是支付宝官方 V3 OpenAPI 模型。
type ViolationEvent struct {
	// CanAppeal 是商家是否可以申诉。
	CanAppeal *bool `json:"can_appeal,omitempty"`
	// CanRectify 是商家是否可以整改。
	CanRectify *bool `json:"can_rectify,omitempty"`
	// PunishAction 是处罚动作及有效期。
	PunishAction *string `json:"punish_action,omitempty"`
	// Status 是违规工单状态枚举。
	Status *string `json:"status,omitempty"`
	// TargetId 是违规对象ID。
	TargetId *string `json:"target_id,omitempty"`
	// TargetName 是违规对象名称。
	TargetName *string `json:"target_name,omitempty"`
	// TargetType 是违规对象类型 小程序ID:APPID 生活号ID:PUBLICID。
	TargetType *string `json:"target_type,omitempty"`
	// ViolationRecordId 是支付宝侧生成的违规记录唯一标识。
	ViolationRecordId *string `json:"violation_record_id,omitempty"`
	// ViolationTime 是违规时间，格式为 yyyy-MM-dd HH:mm:ss。
	ViolationTime *string `json:"violation_time,omitempty"`
	// ViolationType 是即平台依据平台规范/规则，判定商户的违规类型。
	ViolationType *string `json:"violation_type,omitempty"`
}
