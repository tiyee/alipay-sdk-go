// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package authentication

// CustomerFaceManageCreateRequest is the official V3 ZolozAuthenticationCustomerFacemanageCreateModel schema.
type CustomerFaceManageCreateRequest struct {
	// Areacode represents areacode. 地域编码
	Areacode *string `json:"areacode,omitzero"`
	// BizType represents biz_type. 人脸产品能力，预热场景，需要传college，k12，scenic，subway，traffic。具体值视具体场景
	BizType *string `json:"biz_type,omitzero"`
	// Bizscale represents bizscale. 业务量规模
	Bizscale *string `json:"bizscale,omitzero"`
	// Brandcode represents brandcode. 商户品牌
	Brandcode *string `json:"brandcode,omitzero"`
	// Devicenum represents devicenum. 商户机具唯一编码，关键参数 人脸预热，传logicGroupId
	Devicenum *string `json:"devicenum,omitzero"`
	// Extinfo represents extinfo. 拓展参数
	Extinfo *string `json:"extinfo,omitzero"`
	// Facetype represents facetype. 入库类型 IDCARD:身份证 ALIPAY_USER:支付宝用户id, ALIPAY_TEL:手机号入库 CUSTOMER:自定义 DIRECT_CLIENT_PRE：本地预写入 DIRECT_CLIENT：本地客户端库 ENCLOSED：封闭场景
	Facetype *string `json:"facetype,omitzero"`
	// Faceval represents faceval. 入库用户信息 人脸预热，alipayUid必填
	Faceval *string `json:"faceval,omitzero"`
	// Group represents group. 分组5
	Group *string `json:"group,omitzero"`
	// Storecode represents storecode. 门店编码
	Storecode *string `json:"storecode,omitzero"`
	// Validtimes represents validtimes. 有效期天数，如7天、30天、365天
	Validtimes *string `json:"validtimes,omitzero"`
}

// CustomerFaceManageDeleteRequest is the official V3 ZolozAuthenticationCustomerFacemanageDeleteModel schema.
type CustomerFaceManageDeleteRequest struct {
	// Areacode represents areacode. 地域编码
	Areacode *string `json:"areacode,omitzero"`
	// BizType represents biz_type. 人脸产品能力
	BizType *string `json:"biz_type,omitzero"`
	// Bizscale represents bizscale. 业务量
	Bizscale *string `json:"bizscale,omitzero"`
	// Brandcode represents brandcode. 商户品牌
	Brandcode *string `json:"brandcode,omitzero"`
	// Devicenum represents devicenum. 商户机具唯一编码，关键参数
	Devicenum *string `json:"devicenum,omitzero"`
	// Extinfo represents extinfo. 拓展参数
	Extinfo *string `json:"extinfo,omitzero"`
	// Facetype represents facetype. 入库类型 IDCARD:身份证 ALIPAY_USER:支付宝用户id, ALIPAY_TEL:手机号入库 CUSTOMER:自定义
	Facetype *string `json:"facetype,omitzero"`
	// Faceval represents faceval. 入库用户信息
	Faceval *string `json:"faceval,omitzero"`
	// Group represents group. 分组
	Group *string `json:"group,omitzero"`
	// Storecode represents storecode. 门店编码
	Storecode *string `json:"storecode,omitzero"`
	// Validtimes represents validtimes. 有效期天数，如7天、30天、365天
	Validtimes *string `json:"validtimes,omitzero"`
}

// CustomerFTokenQueryRequest is the official V3 ZolozAuthenticationCustomerFtokenQueryModel schema.
type CustomerFTokenQueryRequest struct {
	// BizType represents biz_type. 1、1：1人脸验证能力 2、1：n人脸搜索能力（支付宝uid入库） 3、1：n人脸搜索能力（支付宝手机号入库） 4、手机号和人脸识别综合能力
	BizType *string `json:"biz_type,omitzero"`
	// ExtInfo represents ext_info.
	ExtInfo *FaceExtInfo `json:"ext_info,omitzero"`
	// Ftoken represents ftoken. 人脸token
	Ftoken *string `json:"ftoken,omitzero"`
	// ZimId represents zim_id. 刷脸初始化流程中产生的zimId值
	ZimId *string `json:"zim_id,omitzero"`
}

// CustomerSmilePayInitializeRequest is the official V3 ZolozAuthenticationCustomerSmilepayInitializeModel schema.
type CustomerSmilePayInitializeRequest struct {
	// ServiceId represents service_id. 刷脸服务Id
	ServiceId *string `json:"service_id,omitzero"`
	// ServiceParams represents service_params. 商户可传入刷脸服务所需的扩展参数，json格式
	ServiceParams *string `json:"service_params,omitzero"`
	// Zimmetainfo represents zimmetainfo. { "apdidToken": "设备指纹", "appName": "应用名称", "appVersion": "应用版本", "bioMetaInfo": "生物信息如2.3.0:3,-4" }
	Zimmetainfo *string `json:"zimmetainfo,omitzero"`
}

// SmilePayInitializeRequest is the official V3 ZolozAuthenticationSmilepayInitializeModel schema.
type SmilePayInitializeRequest struct {
	// ApdidToken represents apdid_token. 设备指纹，用于唯一标识一台设备
	ApdidToken *string `json:"apdid_token,omitzero"`
	// AppName represents app_name. 人脸识别应用名称
	AppName *string `json:"app_name,omitzero"`
	// AppVersion represents app_version. 人脸识别应用版本号
	AppVersion *string `json:"app_version,omitzero"`
	// BaseVer represents base_ver. 基础包版本号
	BaseVer *string `json:"base_ver,omitzero"`
	// BioMetaInfo represents bio_meta_info. 生物识别元信息
	BioMetaInfo *string `json:"bio_meta_info,omitzero"`
	// DeviceModel represents device_model. 设备型号
	DeviceModel *string `json:"device_model,omitzero"`
	// DeviceType represents device_type. 设备类型
	DeviceType *string `json:"device_type,omitzero"`
	// ExtInfo represents ext_info.
	ExtInfo *FaceExtParams `json:"ext_info,omitzero"`
	// MachineInfo represents machine_info.
	MachineInfo *FaceMachineInfo `json:"machine_info,omitzero"`
	// MerchantInfo represents merchant_info.
	MerchantInfo *FaceMerchantInfo `json:"merchant_info,omitzero"`
	// OsVersion represents os_version. 操作系统版本号
	OsVersion *string `json:"os_version,omitzero"`
	// RemoteLogId represents remote_log_id. 业务ID
	RemoteLogId *string `json:"remote_log_id,omitzero"`
	// ZimVer represents zim_ver. ZIM版本号
	ZimVer *string `json:"zim_ver,omitzero"`
}
