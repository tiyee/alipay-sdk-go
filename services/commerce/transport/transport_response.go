// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package transport

// TransportChargerChargerBindInfoSyncResponse 是支付宝官方 V3 OpenAPI 模型。
type TransportChargerChargerBindInfoSyncResponse struct {
	// EquipId 是商户同步的充电桩设备编码
	EquipId *string `json:"equip_id,omitempty"`
	// FailReason 是绑定关系结果接收失败原因
	FailReason *string `json:"fail_reason,omitempty"`
	// ResultStatus 是0：成功	1：拒绝
	ResultStatus *string `json:"result_status,omitempty"`
}

// TransportParkingEnterInfoSyncResponse 是支付宝官方 V3 OpenAPI 模型。
type TransportParkingEnterInfoSyncResponse struct {
	// AgreementScene 是用户签约的代扣场景字符集，多个英文逗号分割，当车场为ORC识别车牌的场景，返回值包含PLATE_PAY时，表示代扣协议可用。 当车场为ETC设备识别车牌，返回值包含ETC_PAY时表示协议可用。当用户未签约任何代扣场景时返回NO_AGREEMENT_SCENE
	AgreementScene *string `json:"agreement_scene,omitempty"`
	// BizCode 是具体错误code
	BizCode *string `json:"biz_code,omitempty"`
	// BizMsg 是具体错误描述
	BizMsg *string `json:"biz_msg,omitempty"`
	// SerialNo 是支付宝停车纪录流水号，串联一次停车行为；
	SerialNo *string `json:"serial_no,omitempty"`
}

// TransportParkingExitInfoSyncResponse 是支付宝官方 V3 OpenAPI 模型。
type TransportParkingExitInfoSyncResponse struct {
	// BizCode 是具体错误码；
	BizCode *string `json:"biz_code,omitempty"`
	// BizMsg 是具体错误原因
	BizMsg *string `json:"biz_msg,omitempty"`
}

// TransportParkingPaymentInfoSyncResponse 是支付宝官方 V3 OpenAPI 模型。
type TransportParkingPaymentInfoSyncResponse struct {
	// BizCode 是具体错误码	参数异常：INVALID_PARAMETER	没有权限操作：NO_PERMISSION_ERROR	请求受理结果未知：BIZ_UNKNOWN	请求受理失败：BIZ_FAIL	系统繁忙：SYSTEM_ERROR
	BizCode *string `json:"biz_code,omitempty"`
	// BizMsg 是具体错误原因
	BizMsg *string `json:"biz_msg,omitempty"`
	// SendParkingMessageFlag 是当前行程是否会给用户发送停车消息
	SendParkingMessageFlag *string `json:"send_parking_message_flag,omitempty"`
}
