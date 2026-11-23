// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package agent

// CcmAgentCreateRequest 是支付宝官方 V3 OpenAPI 模型。
type CcmAgentCreateRequest struct {
	// AnsweringMode 是热线接入方式，取值0，1，2
	AnsweringMode *string `json:"answering_mode,omitempty"`
	// CcsInstanceIds 是部门id列表
	CcsInstanceIds *[]string `json:"ccs_instance_ids,omitempty"`
	// ChatConfigs 是在线技能组设置，技能组部门间隔离，此设置应与ccs_instance_ids对应
	ChatConfigs *[]AgentChatInfo `json:"chat_configs,omitempty"`
	// CreatorId 是操作人客服id
	CreatorId *string `json:"creator_id,omitempty"`
	// Email 是客服邮箱
	Email *string `json:"email,omitempty"`
	// ExternalUserId 是客服账号id，当user_channel=ALIPAY时，external_user_id为支付宝账号id
	ExternalUserId *string `json:"external_user_id,omitempty"`
	// HotlineConfigs 是热线技能组设置，技能组部门间隔离，此设置应与ccs_instance_ids对应
	HotlineConfigs *[]AgentHotlineInfo `json:"hotline_configs,omitempty"`
	// JobNumber 是客服工号，新增后不可变更。4位数字，不可重复
	JobNumber *string `json:"job_number,omitempty"`
	// Mobile 是客服手机号
	Mobile *string `json:"mobile,omitempty"`
	// NickName 是客服昵称
	NickName *string `json:"nick_name,omitempty"`
	// RealName 是客服姓名
	RealName *string `json:"real_name,omitempty"`
	// RoleIds 是角色id列表
	RoleIds *[]RoleId `json:"role_ids,omitempty"`
	// UserChannel 是客服账号渠道 ALIPAY:支付宝账号
	UserChannel *string `json:"user_channel,omitempty"`
}

// CcmAgentModifyRequest 是支付宝官方 V3 OpenAPI 模型。
type CcmAgentModifyRequest struct {
	// AnsweringMode 是热线接入方式，取值0，1，2
	AnsweringMode *string `json:"answering_mode,omitempty"`
	// CcsInstanceIds 是部门id列表
	CcsInstanceIds *[]string `json:"ccs_instance_ids,omitempty"`
	// ChatConfigs 是在线技能组设置，技能组部门间隔离，此设置应与ccs_instance_ids对应
	ChatConfigs *[]AgentChatInfo `json:"chat_configs,omitempty"`
	// Email 是客服邮箱
	Email *string `json:"email,omitempty"`
	// HotlineConfigs 是热线技能组设置，技能组部门间隔离，此设置应与ccs_instance_ids对应
	HotlineConfigs *[]AgentHotlineInfo `json:"hotline_configs,omitempty"`
	// Id 是客服id
	Id *string `json:"id,omitempty"`
	// Mobile 是客服手机号
	Mobile *string `json:"mobile,omitempty"`
	// NickName 是客服昵称
	NickName *string `json:"nick_name,omitempty"`
	// RealName 是客服姓名
	RealName *string `json:"real_name,omitempty"`
	// RoleIds 是客服角色id列表
	RoleIds *[]RoleId `json:"role_ids,omitempty"`
	// UpdaterId 是操作人客服id
	UpdaterId *string `json:"updater_id,omitempty"`
}

// CcmAgentQueryRequest 是支付宝官方 V3 OpenAPI 模型。
type CcmAgentQueryRequest struct {
	// AnsweringMode 是热线接入方式
	AnsweringMode *string `json:"answering_mode,omitempty"`
	// CcsInstanceId 是部门id
	CcsInstanceId *string `json:"ccs_instance_id,omitempty"`
	// ChatExtGroupIds 是在线扩展技能组id
	ChatExtGroupIds *[]string `json:"chat_ext_group_ids,omitempty"`
	// ChatGroupIds 是在线技能组id
	ChatGroupIds *[]string `json:"chat_group_ids,omitempty"`
	// ChatLevelIds 是在线服务等级
	ChatLevelIds *[]string `json:"chat_level_ids,omitempty"`
	// Email 是客服邮箱
	Email *string `json:"email,omitempty"`
	// HotlineGroupIds 是热线扩展技能组id
	HotlineGroupIds *[]string `json:"hotline_group_ids,omitempty"`
	// PageNum 是查询结果的页码，起始值为 1，默认值为 1
	PageNum *int64 `json:"page_num,omitempty"`
	// PageSize 是分页查询时设置的每页记录数，最大值 100 行，默认为 10
	PageSize *int64 `json:"page_size,omitempty"`
	// RealName 是客服姓名
	RealName *string `json:"real_name,omitempty"`
	// RoleIds 是客服角色id
	RoleIds *[]string `json:"role_ids,omitempty"`
}

// CcmAgentScheduleLogQueryRequest 是支付宝官方 V3 OpenAPI 模型。
type CcmAgentScheduleLogQueryRequest struct {
	// AgentIds 是客服id集合,限制最多100个id,不传查询部门下所有客服的状态变更流水日志
	AgentIds *[]string `json:"agent_ids,omitempty"`
	// CcsInstanceId 是部门id（即租户实例ID、数据权限ID）
	CcsInstanceId *string `json:"ccs_instance_id,omitempty"`
	// EndTime 是客服状态变更结束时间,开始时间与结束时间间隔不能超过1天,采用UTC时间，按照ISO8601标准表示，格式为：yyyy-MM-dd'T'HH:mm:ss'Z'
	EndTime *string `json:"end_time,omitempty"`
	// Limit 是查询条数,最大100, 不传默认100
	Limit *int64 `json:"limit,omitempty"`
	// PageNum 是查询结果的页码，起始值为 1，默认值为 1
	PageNum *int64 `json:"page_num,omitempty"`
	// PageSize 是分页查询时设置的每页记录数，最大值 100 行，默认为 100
	PageSize *int64 `json:"page_size,omitempty"`
	// StartId 是起始id,分页导出的起始客服状态变更流水id,第一页传0,翻页时传上一页结果的最大id
	StartId *int64 `json:"start_id,omitempty"`
	// StartTime 是客服状态变更开始时间,采用UTC时间，按照ISO8601标准表示，格式为：yyyy-MM-dd'T'HH:mm:ss'Z'
	StartTime *string `json:"start_time,omitempty"`
}

// CcmAgentDeleteRequest is the query request model for alipay.iservice.ccm.agent.delete.
type CcmAgentDeleteRequest struct {
	// Id 是客服id
	Id *string `json:"id,omitempty" query:"id,omitempty"`
}

// CcmAgentGetRequest is the query request model for alipay.iservice.ccm.agent.get.
type CcmAgentGetRequest struct {
	// Id 是客服id
	Id *string `json:"id,omitempty" query:"id,omitempty"`
	// JobNumber 是客服工号
	JobNumber *string `json:"job_number,omitempty" query:"job_number,omitempty"`
	// UserChannel 是客服账号渠道, ALIPAY:支付宝账号	user_channel和external_user_id必须配对使用
	UserChannel *string `json:"user_channel,omitempty" query:"user_channel,omitempty"`
	// ExternalUserId 是客服账号id，当user_channel=ALIPAY时，此id为支付宝2088id	external_user_id和user_channel必须配对使用
	ExternalUserId *string `json:"external_user_id,omitempty" query:"external_user_id,omitempty"`
}
