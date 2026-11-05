// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package servicerecord

// CcmServiceRecordCreateRequest 是支付宝官方 V3 OpenAPI 模型。
type CcmServiceRecordCreateRequest struct {
	// Acid 是来电唯一标识
	Acid *string `json:"acid,omitempty"`
	// Aid 是客服坐席号(软电话中定义的)
	Aid *string `json:"aid,omitempty"`
	// Ani 是主叫号码
	Ani *string `json:"ani,omitempty"`
	// Asid 是恒生专用来电唯一标识
	Asid *string `json:"asid,omitempty"`
	// CategoryList 是类目信息(详细信息请参考具体类型描述)
	CategoryList *[]ServiceRecordCatRequest `json:"category_list,omitempty"`
	// ChatBeginTime 是标准时间格式：yyyy-MM-dd HH:mm:ss
	ChatBeginTime *string `json:"chat_begin_time,omitempty"`
	// ChatEndTime 是标准时间格式：yyyy-MM-dd HH:mm:ss
	ChatEndTime *string `json:"chat_end_time,omitempty"`
	// City 是来电地址
	City *string `json:"city,omitempty"`
	// CreatorId 是创建者ID
	CreatorId *string `json:"creator_id,omitempty"`
	// DisconSymbol 是先挂断方	AGENT:客服	CUSTOMER:客户
	DisconSymbol *string `json:"discon_symbol,omitempty"`
	// Dnis 是被叫号码
	Dnis *string `json:"dnis,omitempty"`
	// FileSize 是文件大小
	FileSize *int64 `json:"file_size,omitempty"`
	// GmtCreate 是创建时间(yyyy-MM-dd HH:mm:ss)
	GmtCreate *string `json:"gmt_create,omitempty"`
	// GmtModified 是修改时间
	GmtModified *string `json:"gmt_modified,omitempty"`
	// Memo 是服务记录备注信息
	Memo *string `json:"memo,omitempty"`
	// ModifierId 是修改人ID
	ModifierId *string `json:"modifier_id,omitempty"`
	// OutboundTaskId 是外呼任务ID
	OutboundTaskId *string `json:"outbound_task_id,omitempty"`
	// Satisfaction 是满意度	0:非常满意	1:满意	2:一般	3:不满意	4:非常不满意
	Satisfaction *string `json:"satisfaction,omitempty"`
	// SatisfactionMemo 是满意度补充信息
	SatisfactionMemo *string `json:"satisfaction_memo,omitempty"`
	// ServiceSource 是1:客户来电	4:电话回访	5:网点服务	6:在线服务	7:外呼任务	8:手动外呼	9:在线离线留言
	ServiceSource *string `json:"service_source,omitempty"`
	// ServiceTime 是服务时长(单位:秒)
	ServiceTime *int64 `json:"service_time,omitempty"`
	// SkillgroupId 是技能组ID
	SkillgroupId *string `json:"skillgroup_id,omitempty"`
	// SkillgroupName 是技能组名称(技能组ID对应的技能组名称)
	SkillgroupName *string `json:"skillgroup_name,omitempty"`
	// TntInstId 是业务租户
	TntInstId *string `json:"tnt_inst_id,omitempty"`
	// UserId 是来访用户唯一标识ID
	UserId *string `json:"user_id,omitempty"`
	// VerifyResult 是IVR验证结果	ROLE_N:未定位到用户	CERTNO_Y_PWD_Y:证件号码+交易密码验证通过	CERTNO_Y_PWD_N:证件号码验证通过+交易密码不通过	CERTNO_Y_PWD_UNKNOW:证件号码验证通过	PHONE_Y_PWD_Y:来电手机+交易密码验证通过	PHONE_Y_PWD_N:来电手机+交易密码未通过	PHONE_Y_PWD_UNKNOW:来电手机账户	SEARCH_Y_PWD_Y:搜索定位用户+推屏交易密码通过	SEARCH_Y_PWD_N:搜索定位用户+推屏交易密码未通过	SEARCH_Y_PWD_UNKNOW:搜索定位用户
	VerifyResult *string `json:"verify_result,omitempty"`
}

// CcmServiceRecordModifyRequest 是支付宝官方 V3 OpenAPI 模型。
type CcmServiceRecordModifyRequest struct {
	// Acid 是来电唯一标识
	Acid *string `json:"acid,omitempty"`
	// Aid 是客服坐席号(软电话中定义的)
	Aid *string `json:"aid,omitempty"`
	// Ani 是主叫号码
	Ani *string `json:"ani,omitempty"`
	// Asid 是恒生专用来电唯一标识
	Asid         *string                  `json:"asid,omitempty"`
	CategoryList *ServiceRecordCatRequest `json:"category_list,omitempty"`
	// ChatBeginTime 是标准时间格式：yyyy-MM-dd HH:mm:ss
	ChatBeginTime *string `json:"chat_begin_time,omitempty"`
	// ChatEndTime 是标准时间格式：yyyy-MM-dd HH:mm:ss
	ChatEndTime *string `json:"chat_end_time,omitempty"`
	// City 是来电地址
	City *string `json:"city,omitempty"`
	// CreatorId 是创建者ID
	CreatorId *string `json:"creator_id,omitempty"`
	// DisconSymbol 是先挂断方 AGENT:客服 CUSTOMER:客户
	DisconSymbol *string `json:"discon_symbol,omitempty"`
	// Dnis 是被叫号码
	Dnis *string `json:"dnis,omitempty"`
	// FileSize 是文件大小
	FileSize *int64 `json:"file_size,omitempty"`
	// GmtCreate 是创建时间(yyyy-MM-dd HH:mm:ss)
	GmtCreate *string `json:"gmt_create,omitempty"`
	// GmtModified 是修改时间
	GmtModified *string `json:"gmt_modified,omitempty"`
	// Id 是云客服服务记录ID
	Id *string `json:"id,omitempty"`
	// Memo 是服务记录备注信息
	Memo *string `json:"memo,omitempty"`
	// ModifierId 是修改人ID
	ModifierId *string `json:"modifier_id,omitempty"`
	// OutboundTaskId 是外呼任务ID
	OutboundTaskId *string `json:"outbound_task_id,omitempty"`
	// Satisfaction 是满意度 0:非常满意 1:满意 2:一般 3:不满意 4:非常不满意
	Satisfaction *string `json:"satisfaction,omitempty"`
	// SatisfactionMemo 是满意度补充信息
	SatisfactionMemo *string `json:"satisfaction_memo,omitempty"`
	// ServiceSource 是1:客户来电 4:电话回访 5:网点服务 6:在线服务 7:外呼任务 8:手动外呼 9:在线离线留言
	ServiceSource *string `json:"service_source,omitempty"`
	// ServiceTime 是服务时长(单位:秒)
	ServiceTime *int64 `json:"service_time,omitempty"`
	// SkillgroupId 是技能组ID
	SkillgroupId *string `json:"skillgroup_id,omitempty"`
	// SkillgroupName 是技能组名称(技能组ID对应的技能组名称)
	SkillgroupName *string `json:"skillgroup_name,omitempty"`
	// Status 是1:待处理	2:暂存	3:完结	4:废除	5:无效	6:未知
	Status *string `json:"status,omitempty"`
	// TntInstId 是业务租户
	TntInstId *string `json:"tnt_inst_id,omitempty"`
	// UserId 是来访用户唯一标识ID
	UserId *string `json:"user_id,omitempty"`
	// VerifyResult 是IVR验证结果 ROLE_N:未定位到用户 CERTNO_Y_PWD_Y:证件号码+交易密码验证通过 CERTNO_Y_PWD_N:证件号码验证通过+交易密码不通过 CERTNO_Y_PWD_UNKNOW:证件号码验证通过 PHONE_Y_PWD_Y:来电手机+交易密码验证通过 PHONE_Y_PWD_N:来电手机+交易密码未通过 PHONE_Y_PWD_UNKNOW:来电手机账户 SEARCH_Y_PWD_Y:搜索定位用户+推屏交易密码通过 SEARCH_Y_PWD_N:搜索定位用户+推屏交易密码未通过 SEARCH_Y_PWD_UNKNOW:搜索定位用户
	VerifyResult *string `json:"verify_result,omitempty"`
}

// ServiceRecordCatRequest 是支付宝官方 V3 OpenAPI 模型。
type ServiceRecordCatRequest struct {
	// CatName 是类目名称(需要每一级用;分隔封装)
	CatName *string `json:"cat_name,omitempty"`
	// FirstCat 是一级类目编码
	FirstCat *string `json:"first_cat,omitempty"`
	// FourthCat 是四级类目编码
	FourthCat *string `json:"fourth_cat,omitempty"`
	// InstenceCode 是类目树编码
	InstenceCode *string `json:"instence_code,omitempty"`
	// SecondCat 是二级类目编码
	SecondCat *string `json:"second_cat,omitempty"`
	// ThirdCat 是三级类目编码
	ThirdCat *string `json:"third_cat,omitempty"`
}
