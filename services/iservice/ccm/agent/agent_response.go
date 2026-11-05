// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package agent

import (
	"encoding/json"
)

// CcmAgentCreateResponse 是支付宝官方 V3 OpenAPI 模型。
type CcmAgentCreateResponse struct {
	// Id 是CCM客服id
	Id *string `json:"id,omitempty"`
}

// CcmAgentDeleteResponse 是支付宝官方 V3 OpenAPI 模型。
type CcmAgentDeleteResponse = json.RawMessage

// CcmAgentGetResponse 是支付宝官方 V3 OpenAPI 模型。
type CcmAgentGetResponse struct {
	// AnsweringMode 是热线接入方式： 0：话机 1：电脑耳机；2：webrtc
	AnsweringMode *string `json:"answering_mode,omitempty"`
	// Avatar 是头像的oss file key
	Avatar *string `json:"avatar,omitempty"`
	// CcsInstanceIds 是客服关联的数据权限id列表
	CcsInstanceIds *[]string `json:"ccs_instance_ids,omitempty"`
	// ChatConfig 是客服在线信息
	ChatConfig *[]AgentChatInfo `json:"chat_config,omitempty"`
	// CreateTime 是创建时间，采用UTC时间，按照ISO8601标准表示，格式为：yyyy-MM-dd'T'HH:mm:ss'Z
	CreateTime *string `json:"create_time,omitempty"`
	// CreatorId 是创建人id
	CreatorId *string `json:"creator_id,omitempty"`
	// DingtalkUserId 是钉钉user_id
	DingtalkUserId *string `json:"dingtalk_user_id,omitempty"`
	// Email 是邮箱
	Email *string `json:"email,omitempty"`
	// ExternalId 是外部系统用户id,比如：金融云的用户id
	ExternalId *string `json:"external_id,omitempty"`
	// HotlineConfig 是客服热线信息
	HotlineConfig *[]AgentHotlineInfo `json:"hotline_config,omitempty"`
	// Id 是客服id
	Id *string `json:"id,omitempty"`
	// JobNumber 是工号
	JobNumber *string `json:"job_number,omitempty"`
	// LastLoginTime 是客服上次登录时间	采用UTC时间，按照ISO8601标准表示，格式为：yyyy-MM-dd'T'HH:mm:ss'Z'
	LastLoginTime *string `json:"last_login_time,omitempty"`
	// Mobile 是手机号
	Mobile *string `json:"mobile,omitempty"`
	// NickName 是客服昵称
	NickName *string `json:"nick_name,omitempty"`
	// Profile 是个人简介
	Profile *string `json:"profile,omitempty"`
	// RealName 是客服姓名
	RealName *string `json:"real_name,omitempty"`
	// RoleIds 是角色id列表
	RoleIds *[]RoleId `json:"role_ids,omitempty"`
	// Status 是客服状态：NORMAL，DELETE
	Status *string `json:"status,omitempty"`
	// Type 是用户类型：NORMAL（普通客服），ADMIN（超级管理员，不能被删除）
	Type *string `json:"type,omitempty"`
	// UpdateTime 是最后修改时间，采用UTC时间，按照ISO8601标准表示，格式为：yyyy-MM-dd'T'HH:mm:ss'Z'
	UpdateTime *string `json:"update_time,omitempty"`
	// UpdaterId 是最后修改人id
	UpdaterId *string `json:"updater_id,omitempty"`
}

// CcmAgentModifyResponse 是支付宝官方 V3 OpenAPI 模型。
type CcmAgentModifyResponse = json.RawMessage

// CcmAgentQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type CcmAgentQueryResponse struct {
	// Agents 是客服列表
	Agents *[]AgentVO `json:"agents,omitempty"`
	// PageNum 是查询结果的页码，起始值为 1，默认值为 1
	PageNum *int64 `json:"page_num,omitempty"`
	// PageSize 是分页查询时设置的每页记录数，最大值 100 行，默认为 10
	PageSize *int64 `json:"page_size,omitempty"`
	// TotalCount 是总条目数
	TotalCount *int64 `json:"total_count,omitempty"`
}

// CcmAgentScheduleLogQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type CcmAgentScheduleLogQueryResponse struct {
	AgentScheduleLogs *AgentScheduleLog `json:"agent_schedule_logs,omitempty"`
	// PageNum 是查询结果的页码
	PageNum *int64 `json:"page_num,omitempty"`
	// PageSize 是分页查询时设置的每页记录数
	PageSize *int64 `json:"page_size,omitempty"`
	// TotalCount 是总条目数
	TotalCount *int64 `json:"total_count,omitempty"`
}
