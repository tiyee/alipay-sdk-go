// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package ols

// ChatMessage 是支付宝官方 V3 OpenAPI 模型。
type ChatMessage struct {
	// MessageContent 是消息内容
	MessageContent *string `json:"message_content,omitempty"`
	// MessageType 是消息内容类型,默认TEXT,包括: TEXT(文本消息),IMAGE(图片消息),FILE(文件消息)
	MessageType *string `json:"message_type,omitempty"`
	// SendTime 是消息发送时间，按照ISO8601标准表示，比如：2018-12-24T15:38:17.824+08:00
	SendTime *string `json:"send_time,omitempty"`
	// UserName 是发送者名称
	UserName *string `json:"user_name,omitempty"`
	// UserType 是发送者类型: VISITOR(访客) ,AGENT(在线客服) ,SYSTEM(系统)
	UserType *string `json:"user_type,omitempty"`
}
