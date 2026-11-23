// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package agent

// AgentChatInfo 是支付宝官方 V3 OpenAPI 模型。
type AgentChatInfo struct {
	// CcsInstanceId 是数据权限id(租户实例id)
	CcsInstanceId *string `json:"ccs_instance_id,omitempty"`
	// ExtendedGroupIds 是在线扩展技能组id列表
	ExtendedGroupIds *[]string `json:"extended_group_ids,omitempty"`
	// GroupId 是在线技能组id
	GroupId *string `json:"group_id,omitempty"`
	// LevelId 是在线服务等级
	LevelId *string `json:"level_id,omitempty"`
}

// AgentHotlineInfo 是支付宝官方 V3 OpenAPI 模型。
type AgentHotlineInfo struct {
	// CcsInstanceId 是数据权限id(租户实例id)
	CcsInstanceId *string `json:"ccs_instance_id,omitempty"`
	// GroupIds 是热线技能组id
	GroupIds *[]string `json:"group_ids,omitempty"`
}

// AgentScheduleLog 是支付宝官方 V3 OpenAPI 模型。
type AgentScheduleLog struct {
	// AgentId 是客服id
	AgentId *string `json:"agent_id,omitempty"`
	// AgentName 是客服名称
	AgentName *string `json:"agent_name,omitempty"`
	// CreateTime 是状态变更发生时间,采用UTC时间，按照ISO8601标准表示，格式为：yyyy-MM-dd'T'HH:mm:ss'Z'
	CreateTime *string `json:"create_time,omitempty"`
	// Duration 是状态持续时长,单位秒
	Duration *int64 `json:"duration,omitempty"`
	// EndTime 是状态变更结束时间,采用UTC时间，按照ISO8601标准表示，格式为：yyyy-MM-dd'T'HH:mm:ss'Z'
	EndTime *string `json:"end_time,omitempty"`
	// ExternalUserNo 是isv或商户系统中对应的客服工号
	ExternalUserNo *string `json:"external_user_no,omitempty"`
	// Id 是客服状态变更流水id
	Id *string `json:"id,omitempty"`
	// LastStatus 是变更前状态
	LastStatus *string `json:"last_status,omitempty"`
	// StartTime 是状态变更开始时间,采用UTC时间，按照ISO8601标准表示，格式为：yyyy-MM-dd'T'HH:mm:ss'Z'
	StartTime *string `json:"start_time,omitempty"`
	// Status 是变更后状态
	Status *string `json:"status,omitempty"`
}

// AgentVO 是支付宝官方 V3 OpenAPI 模型。
type AgentVO struct {
	// AnsweringMode 是热线接入方式
	AnsweringMode *string `json:"answering_mode,omitempty"`
	// Avatar 是头像的oss file key
	Avatar *string `json:"avatar,omitempty"`
	// CcsInstanceIds 是客服归属的部门id列表
	CcsInstanceIds *[]string `json:"ccs_instance_ids,omitempty"`
	// ChatConfig 是在线技能组信息
	ChatConfig *[]AgentChatInfo `json:"chat_config,omitempty"`
	// CreateTime 是创建时间， 标准时间格式：yyyy-MM-dd HH:mm:ss
	CreateTime *string `json:"create_time,omitempty"`
	// CreatorId 是创建人id
	CreatorId *string `json:"creator_id,omitempty"`
	// DingtalkUserId 是钉钉user_id（绑定钉钉的用户该字段非空）
	DingtalkUserId *string `json:"dingtalk_user_id,omitempty"`
	// Email 是客服邮箱
	Email *string `json:"email,omitempty"`
	// ExternalId 是外部系统用户,比如金融云的用户id、支付宝账号id
	ExternalId *string `json:"external_id,omitempty"`
	// HotlineConfig 是热线技能组信息
	HotlineConfig *[]AgentHotlineInfo `json:"hotline_config,omitempty"`
	// Id 是客服唯一标识id
	Id *string `json:"id,omitempty"`
	// JobNumber 是客服工号
	JobNumber *string `json:"job_number,omitempty"`
	// LastLoginTime 是最后登录时间,标准时间格式：yyyy-MM-dd HH:mm:ss
	LastLoginTime *string `json:"last_login_time,omitempty"`
	// Mobile 是客服手机号
	Mobile *string `json:"mobile,omitempty"`
	// NickName 是客服昵称
	NickName *string `json:"nick_name,omitempty"`
	// Profile 是个人简介
	Profile *string `json:"profile,omitempty"`
	// RealName 是客服姓名
	RealName *string `json:"real_name,omitempty"`
	// RoleIds 是客服角色id列表
	RoleIds *[]RoleId `json:"role_ids,omitempty"`
	// Status 是状态：NORMAL，DELETE
	Status *string `json:"status,omitempty"`
	// Type 是用户类型：NORMAL（普通客服），ADMIN（超级管理员，不能被删除）
	Type *string `json:"type,omitempty"`
	// UpdateTime 是最后修改时间， 标准时间格式：yyyy-MM-dd HH:mm:ss
	UpdateTime *string `json:"update_time,omitempty"`
	// UpdaterId 是最后修改人id
	UpdaterId *string `json:"updater_id,omitempty"`
}

// RoleId 是支付宝官方 V3 OpenAPI 模型。
type RoleId struct {
	// CcsInstanceId 是部门id（即租户实例ID、数据权限ID）
	CcsInstanceId *string `json:"ccs_instance_id,omitempty"`
	// RoleId 是角色ID
	RoleId *string `json:"role_id,omitempty"`
}
