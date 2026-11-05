// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package iot

// IoTDeviceReportUploadResponse 是支付宝官方 V3 OpenAPI 模型。
type IoTDeviceReportUploadResponse struct {
	// Status 是是否上传成功
	Status *bool `json:"status,omitempty"`
}

// IoTSdarttoolMessageQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type IoTSdarttoolMessageQueryResponse struct {
	// ConsumeErrCode 是消息消费返回码，不用的消息类型返回码不一样
	ConsumeErrCode *string `json:"consume_err_code,omitempty"`
	// ConsumeErrMsg 是消息消费错误原因
	ConsumeErrMsg *string `json:"consume_err_msg,omitempty"`
	// ConsumeState 是消息消费状态1-未消费; 2-消费成功；3-消费失败
	ConsumeState *int64 `json:"consume_state,omitempty"`
	// DeliverCnt 是消息投递次数
	DeliverCnt *int64 `json:"deliver_cnt,omitempty"`
	// DeliverStatus 是消息投递状态	1: 未投递 2: 投递中 3: 投递成功 4: 投递失败 5: 消息失效不再发送
	DeliverStatus *string `json:"deliver_status,omitempty"`
	// LastDeliverAt 是消息最后一次发送时间
	LastDeliverAt *int64 `json:"last_deliver_at,omitempty"`
	// MsgAckTime 是消息确认消费时间(ms)
	MsgAckTime *int64 `json:"msg_ack_time,omitempty"`
	// MsgBida 是消息必达类型 1-非必达; 2-必达.
	MsgBida *int64 `json:"msg_bida,omitempty"`
	// MsgExpireTime 是消息过期时间(ms)
	MsgExpireTime *int64 `json:"msg_expire_time,omitempty"`
	// MsgGmtCreate 是消息创建时间(ms)
	MsgGmtCreate *int64 `json:"msg_gmt_create,omitempty"`
	// MsgPriority 是消息投递优先级
	MsgPriority *int64 `json:"msg_priority,omitempty"`
}

// IoTSdarttoolMessageSendResponse 是支付宝官方 V3 OpenAPI 模型。
type IoTSdarttoolMessageSendResponse struct {
	// MessageNo 是消息ID
	MessageNo *string `json:"message_no,omitempty"`
}
