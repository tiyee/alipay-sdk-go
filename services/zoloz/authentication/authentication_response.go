// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package authentication

// CustomerFaceManageCreateResponse is the official V3 ZolozAuthenticationCustomerFacemanageCreateResponseModel schema.
type CustomerFaceManageCreateResponse struct {
	// Result represents result. 业务结果
	Result *string `json:"result,omitzero"`
	// Retcode represents retcode. 返回码
	Retcode *string `json:"retcode,omitzero"`
	// Retcodesub represents retcodesub. 返回详细码
	Retcodesub *string `json:"retcodesub,omitzero"`
	// Retmessage represents retmessage. 返回信息
	Retmessage *string `json:"retmessage,omitzero"`
	// Retmessagesub represents retmessagesub. 返回详细信息
	Retmessagesub *string `json:"retmessagesub,omitzero"`
}

// CustomerFaceManageDeleteResponse is the official V3 ZolozAuthenticationCustomerFacemanageDeleteResponseModel schema.
type CustomerFaceManageDeleteResponse struct {
	// Result represents result. 返回值
	Result *string `json:"result,omitzero"`
	// Retcode represents retcode. 返回码
	Retcode *string `json:"retcode,omitzero"`
	// Retcodesub represents retcodesub. 返回详细码
	Retcodesub *string `json:"retcodesub,omitzero"`
	// Retmessage represents retmessage. 返回信息
	Retmessage *string `json:"retmessage,omitzero"`
	// Retmessagesub represents retmessagesub. 返回详细信息
	Retmessagesub *string `json:"retmessagesub,omitzero"`
}

// CustomerFTokenQueryResponse is the official V3 ZolozAuthenticationCustomerFtokenQueryResponseModel schema.
type CustomerFTokenQueryResponse struct {
	// AgeCheckResult represents age_check_result. 年龄是否在指定范围内，未指定范围则返回空，true/false
	AgeCheckResult *string `json:"age_check_result,omitzero"`
	// AuthimgBase64 represents authimg_base_64. 图片base64 encodeString
	AuthimgBase64 *string `json:"authimg_base_64,omitzero"`
	// CertName represents cert_name. 证件姓名
	CertName *string `json:"cert_name,omitzero"`
	// CertNo represents cert_no. 身份证号码
	CertNo *string `json:"cert_no,omitzero"`
	// FaceId represents face_id. 由ISV定义的对自然人唯一编码，举例可以是身份证号码和姓名的MD5值，或者是其他编码方式，要求脱敏、随机且在ISV可以唯一说明一个自然人
	FaceId *string `json:"face_id,omitzero"`
	// OpenId represents open_id. 支付宝用户open_id
	OpenId *string `json:"open_id,omitzero"`
	// Uid represents uid. 支付宝uid
	Uid *string `json:"uid,omitzero"`
	// UidTelPairList represents uid_tel_pair_list. 用户名信息返回的列表
	UidTelPairList *[]ZhubUidTelPair `json:"uid_tel_pair_list,omitzero"`
}

// CustomerSmilePayInitializeResponse is the official V3 ZolozAuthenticationCustomerSmilepayInitializeResponseModel schema.
type CustomerSmilePayInitializeResponse struct {
	// Result represents result. 返回值
	Result *string `json:"result,omitzero"`
}

// SmilePayInitializeResponse is the official V3 ZolozAuthenticationSmilepayInitializeResponseModel schema.
type SmilePayInitializeResponse struct {
	// RetCodeSub represents ret_code_sub. 返回详细码
	RetCodeSub *string `json:"ret_code_sub,omitzero"`
	// RetMessageSub represents ret_message_sub. 返回详细信息
	RetMessageSub *string `json:"ret_message_sub,omitzero"`
	// ZimId represents zim_id. ZIM上下文ID
	ZimId *string `json:"zim_id,omitzero"`
	// ZimInitClientData represents zim_init_client_data. 客户端协议
	ZimInitClientData *string `json:"zim_init_client_data,omitzero"`
}
