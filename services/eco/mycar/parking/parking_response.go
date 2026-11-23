package parking

import (
	"encoding/json"
)

// MyCarParkingChargeInfoSyncResponse 是支付宝官方 V3 OpenAPI 模型。
type MyCarParkingChargeInfoSyncResponse struct {
	// SyncResult 是同步结果：0 成功，1 失败
	SyncResult *string `json:"sync_result,omitempty"`
}

// MyCarParkingChargingInfoSyncResponse 是支付宝官方 V3 OpenAPI 模型。
type MyCarParkingChargingInfoSyncResponse = json.RawMessage

// MyCarParkingConfigQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type MyCarParkingConfigQueryResponse struct {
	// AccountNo 是签约支付宝账号，开发者通过ISV系统配置信息注册接口(alipay.eco.mycar.parking.config.set)传入的参数值
	AccountNo         *string            `json:"account_no,omitempty"`
	InterfaceInfoList *InterfaceInfoList `json:"interface_info_list,omitempty"`
	// MerchantLogo 是商户在停车平台首页露出的LOGO链接地址，开发者通过ISV系统配置信息注册接口(alipay.eco.mycar.parking.config.set)调用自动生成该链接
	MerchantLogo *string `json:"merchant_logo,omitempty"`
	// MerchantName 是商户简称，开发者通过ISV系统配置信息注册接口(alipay.eco.mycar.parking.config.set)传入的参数值
	MerchantName *string `json:"merchant_name,omitempty"`
	// MerchantServicePhone 是商户客服电话，开发者通过ISV系统配置信息注册接口(alipay.eco.mycar.parking.config.set)传入的参数值
	MerchantServicePhone *string `json:"merchant_service_phone,omitempty"`
}

// MyCarParkingConfigSetResponse 是支付宝官方 V3 OpenAPI 模型。
type MyCarParkingConfigSetResponse = json.RawMessage

// MyCarParkingEnterInfoSyncResponse 是支付宝官方 V3 OpenAPI 模型。
type MyCarParkingEnterInfoSyncResponse struct {
	// AgreementScene 是用户签约的代扣场景字符集，多个英文逗号分割，当车场为ORC识别车牌的场景，返回值包含PLATE_PAY时，表示代扣协议可用。 当车场为ETC设备识别车牌，返回值包含ETC_PAY时表示协议可用。当用户未签约任何代扣场景时返回NO_AGREEMENT_SCENE
	AgreementScene *string `json:"agreement_scene,omitempty"`
	// AgreementStatus 是该字段已废弃,请使用agreement_scene字段(该字段已废弃,请使用agreement_scene字段)
	AgreementStatus *string `json:"agreement_status,omitempty"`
	// CarStatus 是当前车辆在支付宝侧的信息，是否可以使用无感支付，可选返回项为： NORMAL：正常缴费用户 NON-SENSE-PAY：可无感支付用户
	CarStatus *string `json:"car_status,omitempty"`
	// CurrentTime 是当前时间戳(查询签约状态和是否支持无感支付仅代表当前时间点查询结果，不作为最后代扣结果，应在发起代扣前再次查询)
	CurrentTime *string `json:"current_time,omitempty"`
	// SerialNo 是支付宝业务流水号，用于记录车辆从驶入到驶出的全流程
	SerialNo *string `json:"serial_no,omitempty"`
}

// MyCarParkingExitInfoSyncResponse 是支付宝官方 V3 OpenAPI 模型。
type MyCarParkingExitInfoSyncResponse = json.RawMessage

// MyCarParkingOrderSyncResponse 是支付宝官方 V3 OpenAPI 模型。
type MyCarParkingOrderSyncResponse = json.RawMessage

// MyCarParkingOrderUpdateResponse 是支付宝官方 V3 OpenAPI 模型。
type MyCarParkingOrderUpdateResponse = json.RawMessage

// MyCarParkingOvertimeChargingInfoSyncResponse 是支付宝官方 V3 OpenAPI 模型。
type MyCarParkingOvertimeChargingInfoSyncResponse = json.RawMessage

// MyCarParkingParkingLotBizTransferResponse 是支付宝官方 V3 OpenAPI 模型。
type MyCarParkingParkingLotBizTransferResponse = json.RawMessage

// MyCarParkingParkingLotInfoCreateResponse 是支付宝官方 V3 OpenAPI 模型。
type MyCarParkingParkingLotInfoCreateResponse struct {
	// BizCode 是具体错误码 参数异常：INVALID_PARAMETER 没有权限操作：NO_PERMISSION_ERROR 请求受理结果未知：BIZ_UNKNOWN 请求受理失败：BIZ_FAIL 系统繁忙：SYSTEM_ERROR
	BizCode *string `json:"biz_code,omitempty"`
	// BizMsg 是具体错误原因
	BizMsg *string `json:"biz_msg,omitempty"`
	// ParkingId 是支付宝返回停车场id。成功不为空，失败返回空
	ParkingId *string `json:"parking_id,omitempty"`
}

// MyCarParkingParkingLotInfoQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type MyCarParkingParkingLotInfoQueryResponse struct {
	// AddressId 是地区编码
	AddressId *string `json:"address_id,omitempty"`
	// AgentId 是服务商ID（2088开头的16位纯数字），由服务商提供给ISV
	AgentId *string `json:"agent_id,omitempty"`
	// BusinessIsv 是车场业务归属列表
	BusinessIsv *[]BusinessItem `json:"business_isv,omitempty"`
	// ChargingRule 是停车场详细计费规则
	ChargingRule *[]ParkingLotChargingRuleInfo `json:"charging_rule,omitempty"`
	// CityId 是城市编码
	CityId *string `json:"city_id,omitempty"`
	// MapPoiAddress 是高德检索地址
	MapPoiAddress *string `json:"map_poi_address,omitempty"`
	// MapPoiName 是高德检索名称
	MapPoiName *string `json:"map_poi_name,omitempty"`
	// MchntId 是收款方ID（2088开头的16位纯数字），由停车场收款的业主方提供给ISV，该字段暂用于机具和物料申领
	MchntId *string `json:"mchnt_id,omitempty"`
	// OutParkingId 是ISV停车场ID，由ISV提供，同一个isv或商户范围内唯一
	OutParkingId *string `json:"out_parking_id,omitempty"`
	// ParkingAddress 是停车场地址
	ParkingAddress *string `json:"parking_address,omitempty"`
	// ParkingFeeDescription 是收费说明
	ParkingFeeDescription *string `json:"parking_fee_description,omitempty"`
	// ParkingFeeDescriptionImg 是商户在停车平台露出的停车价格图片地址
	ParkingFeeDescriptionImg *string `json:"parking_fee_description_img,omitempty"`
	// ParkingId 是支付宝返回停车场id，系统唯一
	ParkingId *string `json:"parking_id,omitempty"`
	// ParkingLatitude 是停车场位置经度（优先高德坐标获取结果）单位：度
	ParkingLatitude *string `json:"parking_latitude,omitempty"`
	// ParkingLongitude 是停车场位置经度（优先高德坐标获取结果）单位：度
	ParkingLongitude *string `json:"parking_longitude,omitempty"`
	// ParkingLotType 是停车场类型，1为居民小区、2为商圈停车场（购物中心商业广场商场等）、3为路侧停车、4为公园景点（景点乐园公园老街古镇等）、5为商务楼宇（酒店写字楼商务楼园区等）、6为其他、7为交通枢纽（机场火车站汽车站码头港口等）、8为市政设施（体育场博物图书馆医院学校等）、9独立园区（办公工业物流园区等）、10经营场所（4S店、门市餐饮等
	ParkingLotType *string `json:"parking_lot_type,omitempty"`
	// ParkingLotTypeCode 是停车场类型，COMMUNITY为居民小区、BUSINESS_AREA为商圈停车场（购物中心商业广场商场等）、ROADSIDE为路侧停车、PARK_SCENIC为公园景点（景点乐园公园老街古镇等）、OFFICE_BUILDING为商务楼宇（酒店写字楼商务楼园区等）、OTHER为其他、TRANSPORTATION为交通枢纽（机场火车站汽车站码头港口等）、PUBLIC_FACILITIES为市政设施（体育场博物图书馆医院学校等）、TERRITORY独立园区（办公工业物流园区等）、BUSINESS_PLACE经营场所（4S店、门市餐饮等
	ParkingLotTypeCode *string `json:"parking_lot_type_code,omitempty"`
	// ParkingMobile 是停车场客服电话
	ParkingMobile *string `json:"parking_mobile,omitempty"`
	// ParkingName 是停车场名称，由ISV定义，尽量与高德地图上的一致
	ParkingName *string `json:"parking_name,omitempty"`
	// ParkingPoiid 是高德地图唯一标识
	ParkingPoiid *string `json:"parking_poiid,omitempty"`
	// PayType 是支付方式（1为支付宝在线缴费，2为支付宝代扣缴费，3当面付 4现金)，如支持多种方式以','进行间隔
	PayType *string `json:"pay_type,omitempty"`
	// ProvinceId 是省份编码
	ProvinceId *string `json:"province_id,omitempty"`
	// SerivceUrl 是提供用户查看停车状态、缴费功能
	SerivceUrl *string `json:"serivce_url,omitempty"`
	// ServiceList 是停车场服务信息列表
	ServiceList *[]ParkingLotServiceInfo `json:"service_list,omitempty"`
	// ShopingmallId 是商圈id
	ShopingmallId *string `json:"shopingmall_id,omitempty"`
	// TimeOut 是用户支付未离场的超时时间(以分钟为单位)
	TimeOut *string `json:"time_out,omitempty"`
}

// MyCarParkingParkingLotInfoUpdateResponse 是支付宝官方 V3 OpenAPI 模型。
type MyCarParkingParkingLotInfoUpdateResponse struct {
	// BizCode 是具体错误码 参数异常：INVALID_PARAMETER 没有权限操作：NO_PERMISSION_ERROR 请求受理结果未知：BIZ_UNKNOWN 请求受理失败：BIZ_FAIL 系统繁忙：SYSTEM_ERROR
	BizCode *string `json:"biz_code,omitempty"`
	// BizMsg 是具体错误原因
	BizMsg *string `json:"biz_msg,omitempty"`
}

// MyCarParkingPaymentInfoSyncResponse 是支付宝官方 V3 OpenAPI 模型。
type MyCarParkingPaymentInfoSyncResponse = json.RawMessage

// MyCarParkingSpaceInfoSyncResponse 是支付宝官方 V3 OpenAPI 模型。
type MyCarParkingSpaceInfoSyncResponse struct {
	// SyncResult 是同步结果：0 成功，1 失败
	SyncResult *string `json:"sync_result,omitempty"`
}

// MyCarParkingVehicleQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type MyCarParkingVehicleQueryResponse struct {
	// CarNumber 是车牌信息（utf-8编码）
	CarNumber *string `json:"car_number,omitempty"`
}
