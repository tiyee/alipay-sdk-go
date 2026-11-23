package blueseaactivity

// SpBlueSeaActivityCreateResponse 是支付宝官方 V3 OpenAPI 模型。
type SpBlueSeaActivityCreateResponse struct {
	// OrderId 是申请单Id。
	OrderId *string `json:"order_id,omitempty"`
}

// SpBlueSeaActivityModifyResponse 是支付宝官方 V3 OpenAPI 模型。
type SpBlueSeaActivityModifyResponse struct {
}

// SpBlueSeaActivityQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type SpBlueSeaActivityQueryResponse struct {
	// Address 是详细地址。
	Address *string `json:"address,omitempty"`
	// BizScene 是蓝海活动的场景，包括直连餐饮（BLUE_SEA_FOOD_APPLY）、直连快消（BLUE_SEA_FMCG_APPLY）、间连餐饮（BLUE_SEA_FOOD_INDIRECT_APPLY）、间连快消（BLUE_SEA_FMCG_INDIRECT_APPLY）场景。
	BizScene *string `json:"biz_scene,omitempty"`
	// BusinessLic 是营业执照，要求证件文本信息清晰可见。 资质信息请参见<a href="https://opendocs.alipay.com/open/01hd83">报名资质要求</a>。
	BusinessLic *string `json:"business_lic,omitempty"`
	// CityCode 是城市编码。请按照https://gw.alipayobjects.com/os/basement_prod/253c4dcb-b8a4-4a1e-8be2-79e191a9b6db.xlsx 表格中内容填写。 （参考资料： http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/）。
	CityCode *string `json:"city_code,omitempty"`
	// DistrictCode 是区县编码。请按照https://gw.alipayobjects.com/os/basement_prod/253c4dcb-b8a4-4a1e-8be2-79e191a9b6db.xlsx 表格中内容填写。 （参考资料： http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/）。
	DistrictCode *string `json:"district_code,omitempty"`
	// FoodBusinessLic 是食品经营许可证，要求证件文本信息清晰可见。资质信息请参见<a href="https://opendocs.alipay.com/open/01hd83">报名资质要求</a>。
	FoodBusinessLic *string `json:"food_business_lic,omitempty"`
	// FoodCirculateLic 是食品流通许可证，要求证件文本信息清晰可见。资质信息请参见<a href="https://opendocs.alipay.com/open/01hd83">报名资质要求</a>。
	FoodCirculateLic *string `json:"food_circulate_lic,omitempty"`
	// FoodHealthLic 是食品卫生许可证，要求证件文本信息清晰可见。资质信息请参见<a href="https://opendocs.alipay.com/open/01hd83">报名资质要求</a>。
	FoodHealthLic *string `json:"food_health_lic,omitempty"`
	// FoodProductionLic 是食品生产许可证，要求证件文本信息清晰可见。资质信息请参见<a href="https://opendocs.alipay.com/open/01hd83">报名资质要求</a>。
	FoodProductionLic *string `json:"food_production_lic,omitempty"`
	// FoodServiceLic 是餐饮服务许可证，要求证件文本信息清晰可见。资质信息请参见<a href="https://opendocs.alipay.com/open/01hd83">报名资质要求</a>。
	FoodServiceLic *string `json:"food_service_lic,omitempty"`
	// IndoorPic 是门头照，要求内景照片清晰可见。资质信息请参见<a href="https://opendocs.alipay.com/open/01hd83">报名资质要求</a>。
	IndoorPic *string `json:"indoor_pic,omitempty"`
	// Memo 是审核驳回原因，在订单状态为失败时有效。
	Memo *string `json:"memo,omitempty"`
	// MerchantLogon 是参与蓝海活动的商户支付宝账号，只有当参与直连蓝海活动场景（BLUE_SEA_FOOD_APPLY/BLUE_SEA_FMCG_APPLY）时才返回。
	MerchantLogon *string `json:"merchant_logon,omitempty"`
	// ProvinceCode 是省份编码。请按照https://gw.alipayobjects.com/os/basement_prod/253c4dcb-b8a4-4a1e-8be2-79e191a9b6db.xlsx 表格中内容填写。 （参考资料： http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/）。
	ProvinceCode *string `json:"province_code,omitempty"`
	// ShopEntrancePic 是门头照，要求店铺外观照片清晰可见。资质信息请参见<a href="https://opendocs.alipay.com/open/01hd83">报名资质要求</a>。
	ShopEntrancePic *string `json:"shop_entrance_pic,omitempty"`
	// Status 是申请单状态，状态机参考（AUDITING:审核中，FAIL:报名失败，PASS:报名成功）。
	Status *string `json:"status,omitempty"`
	// SubMerchantId 是参与蓝海活动的间连商户账号，只有当参与间连蓝海活动场景（BLUE_SEA_FOOD_INDIRECT_APPLY/BLUE_SEA_FMCG_INDIRECT_APPLY）时才返回。
	SubMerchantId *string `json:"sub_merchant_id,omitempty"`
	// TobaccoLic 是烟草专卖零售许可证，要求证件文本信息清晰可见。资质信息请参见<a href="https://opendocs.alipay.com/open/01hd83">报名资质要求</a>。
	TobaccoLic *string `json:"tobacco_lic,omitempty"`
}
