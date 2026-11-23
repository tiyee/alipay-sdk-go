// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package iot

// IoTDeviceReportUploadRequest 是支付宝官方 V3 OpenAPI 模型。
type IoTDeviceReportUploadRequest struct {
	// ReportContent 是设备检测报告详情
	ReportContent *string `json:"report_content,omitempty"`
	// Sn 是设备唯一标识
	Sn *string `json:"sn,omitempty"`
}

// IoTSdarttoolMessageSendRequest 是支付宝官方 V3 OpenAPI 模型。
type IoTSdarttoolMessageSendRequest struct {
	// BiDa 是消息是否必达(immediate_msg为false时此字段设置有效)，消息过期时间顺延3*24h. 建议使用expire_time设置消息过期时间
	BiDa *bool `json:"bi_da,omitempty"`
	// DeviceQueryType 是设备查询条件类型 ；SUPPLIERID_SN:supplierid+sn ;	 ITEMID_SN:itemid + sn;	SN:sn
	DeviceQueryType *string `json:"device_query_type,omitempty"`
	// ImmediateMsg 是是否即时消息; true-校验设备是否在线，false-不校验设备是否在线
	ImmediateMsg *bool `json:"immediate_msg,omitempty"`
	// ItemId 是产品ID
	ItemId *string `json:"item_id,omitempty"`
	// MsgContent 是消息内容
	MsgContent *string `json:"msg_content,omitempty"`
	// MsgContentType 是消息内容格式
	MsgContentType *string `json:"msg_content_type,omitempty"`
	// MsgExpire 是消息过期时间戳(ms)， 默认为当前时间顺延24h有效。当设备在线时消息服务过期之前尝试推送。最大过期时间顺延3*24h
	MsgExpire *int64 `json:"msg_expire,omitempty"`
	// MsgPriority 是消息优先级（immediate_msg为false此字段设置有效）数字越大优先发送。
	MsgPriority *int64 `json:"msg_priority,omitempty"`
	// MsgType 是消息类型：	xpaas_common-小程序消息；	ruyi_ordermsg-如意订单消息
	MsgType *string `json:"msg_type,omitempty"`
	// ServiceId 是服务ID，当小程序消费消息时必传(传入的是小程序ID)
	ServiceId *string `json:"service_id,omitempty"`
	// Sn 是设备sn号
	Sn *string `json:"sn,omitempty"`
	// SupplierId 是厂商ID
	SupplierId *string `json:"supplier_id,omitempty"`
}

// IoTSdarttoolMessageQueryRequest is the query request model for alipay.commerce.iot.sdarttool.message.query.
type IoTSdarttoolMessageQueryRequest struct {
	// MessageNo 是消息ID
	MessageNo *string `json:"message_no,omitempty" query:"message_no,omitempty"`
}
