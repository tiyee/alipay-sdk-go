// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package violation

// ViolationViolationDetailQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type ViolationViolationDetailQueryResponse struct {
	// AppealDeadLine 是申诉截止日期。 即商户可进行申诉的截止日期，若超过该时间未申诉或申诉不通过，平台不提供申诉机会，维持原处罚结果；若该时间内完成申诉且申诉通过，平台将撤销对商户的处罚。
	AppealDeadLine *string `json:"appeal_dead_line,omitempty"`
	// AppealReplyRecords 是商家申诉记录。
	AppealReplyRecords *[]ReplyRecord `json:"appeal_reply_records,omitempty"`
	// CanAppeal 是商家是否可以申诉。
	CanAppeal *bool `json:"can_appeal,omitempty"`
	// CanRectify 是商家是否可以整改。
	CanRectify *bool `json:"can_rectify,omitempty"`
	// PunishAction 是处罚动作及有效期。
	PunishAction *[]string `json:"punish_action,omitempty"`
	// RectifyDeadLine 是截止整改时间。即商户可进行整改的截止日期，若超过该时间未整改或整改不通过，平台将对商户进行处罚，若该时间内完成整改且整改通过，平台对商户不处罚。
	RectifyDeadLine *string `json:"rectify_dead_line,omitempty"`
	// RectifyReplyRecords 是商家整改记录。
	RectifyReplyRecords *[]ReplyRecord `json:"rectify_reply_records,omitempty"`
	// Status 是违规工单状态枚举。
	Status *string `json:"status,omitempty"`
	// SurplusAppealCnt 是剩余申诉次数。 即在申诉截止日期内，平台允许商家申诉的次数，当商户提交申诉但申诉未通过时，剩余申诉次数减1，当整改次数为0，不管是否在整改截止日期内，平台将对商户进行处罚。
	SurplusAppealCnt *string `json:"surplus_appeal_cnt,omitempty"`
	// SurplusRectifyCnt 是剩余整改次数。即在整改截止日期内，平台允许商家整改的次数，当商户提交整改但整改未通过时，剩余整改次数减1，当整改次数为0，不管是否在整改截止日期内，平台将对商户进行处罚。
	SurplusRectifyCnt *string `json:"surplus_rectify_cnt,omitempty"`
	// TargetId 是违规对象ID。
	TargetId *string `json:"target_id,omitempty"`
	// TargetName 是违规对象名称。
	TargetName *string `json:"target_name,omitempty"`
	// TargetType 是违规对象类型 小程序ID:APPID 生活号ID:PUBLICID。
	TargetType *string `json:"target_type,omitempty"`
	// ViolationEvidence 是平台判定商家违规凭证。
	ViolationEvidence *[]AuditEvidenceInfo `json:"violation_evidence,omitempty"`
	// ViolationReason 是违规原因。
	ViolationReason *string `json:"violation_reason,omitempty"`
	// ViolationRecordId 是支付宝侧生成的违规记录唯一标识。
	ViolationRecordId *string `json:"violation_record_id,omitempty"`
	// ViolationTime 是违规时间，格式为 yyyy-MM-dd HH:mm:ss。
	ViolationTime *string `json:"violation_time,omitempty"`
	// ViolationType 是即平台依据平台规范/规则，判定商户的违规类型。
	ViolationType *string `json:"violation_type,omitempty"`
}

// ViolationViolationEventBatchQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type ViolationViolationEventBatchQueryResponse struct {
	ViolationRecordInfos *ViolationEvent `json:"violation_record_infos,omitempty"`
}
