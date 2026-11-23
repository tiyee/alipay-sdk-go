package blueseaactivity

// SpBlueSeaActivityCreateRequest 是支付宝官方 V3 OpenAPI 模型。
type SpBlueSeaActivityCreateRequest struct {
	// Address 是详细地址。 该资质是否必传请参见<a href='https://opendocs.alipay.com/open/01hd83'>报名资质要求</a>。。
	Address *string `json:"address,omitempty"`
	// BizScene 是蓝海活动的场景，支持场景如下： * 直连餐饮（BLUE_SEA_FOOD_APPLY）； * 直连快消（BLUE_SEA_FMCG_APPLY）； * 间连餐饮（BLUE_SEA_FOOD_INDIRECT_APPLY）； * 间连快消（BLUE_SEA_FMCG_INDIRECT_APPLY）。。
	BizScene *string `json:"biz_scene,omitempty"`
	// BusinessLic 是营业执照，要求证件文本信息清晰可见，图片文件大小在 50K-5M 之间，不限制长宽，支持 png、bmp、gif、jpg、jpeg 格式。 请传入<a href='https://opendocs.alipay.com/apis/01ea4t'>alipay.open.sp.image.upload</a>(图片上传接口) 返回的 image_id。 该资质是否必传请参见<a href='https://opendocs.alipay.com/open/01hd83'>报名资质要求</a>。。
	BusinessLic *string `json:"business_lic,omitempty"`
	// CityCode 是城市编码。请按照<a href='https://gw.alipayobjects.com/os/basement_prod/253c4dcb-b8a4-4a1e-8be2-79e191a9b6db.xlsx'>表格</a> 中内容填写。 （请参考资料：<a href='http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/'> 统计用区划和城乡划分代码</a>）。
	CityCode *string `json:"city_code,omitempty"`
	// DistrictCode 是区县编码。请按照<a href='https://gw.alipayobjects.com/os/basement_prod/253c4dcb-b8a4-4a1e-8be2-79e191a9b6db.xlsx'>表格</a> 中内容填写。 （请参考资料：<a href='http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/'> 统计用区划和城乡划分代码</a>）。
	DistrictCode *string `json:"district_code,omitempty"`
	// FoodBusinessLic 是食品经营许可证，要求证件文本信息清晰可见，图片文件大小在 50K-5M 之间，不限制长宽，支持 png、bmp、gif、jpg、jpeg 格式。 请传入<a href='https://opendocs.alipay.com/apis/01ea4t'>alipay.open.sp.image.upload</a>(图片上传接口) 返回的 image_id。 该资质是否必传请参见<a href='https://opendocs.alipay.com/open/01hd83'>报名资质要求</a>。。
	FoodBusinessLic *string `json:"food_business_lic,omitempty"`
	// FoodCirculateLic 是食品流通许可证，要求证件文本信息清晰可见，图片文件大小在 50K-5M 之间，不限制长宽，支持 png、bmp、gif、jpg、jpeg 格式。 请传入<a href='https://opendocs.alipay.com/apis/01ea4t'>alipay.open.sp.image.upload</a>(图片上传接口) 返回的 image_id。 该资质是否必传请参见<a href='https://opendocs.alipay.com/open/01hd83'>报名资质要求</a>。。
	FoodCirculateLic *string `json:"food_circulate_lic,omitempty"`
	// FoodHealthLic 是食品卫生许可证，要求证件文本信息清晰可见，图片文件大小在 50K-5M 之间，不限制长宽，支持 png、bmp、gif、jpg、jpeg 格式。 请传入<a href='https://opendocs.alipay.com/apis/01ea4t'>alipay.open.sp.image.upload</a>(图片上传接口) 返回的 image_id。 该资质是否必传请参见<a href='https://opendocs.alipay.com/open/01hd83'>报名资质要求</a>。。
	FoodHealthLic *string `json:"food_health_lic,omitempty"`
	// FoodProductionLic 是食品生产许可证，要求证件文本信息清晰可见，图片文件大小在 50K-5M 之间，不限制长宽，支持 png、bmp、gif、jpg、jpeg 格式。 请传入<a href='https://opendocs.alipay.com/apis/01ea4t'>alipay.open.sp.image.upload</a>(图片上传接口) 返回的 image_id。 该资质是否必传请参见<a href='https://opendocs.alipay.com/open/01hd83'>报名资质要求</a>。。
	FoodProductionLic *string `json:"food_production_lic,omitempty"`
	// FoodServiceLic 是餐饮服务许可证，要求证件文本信息清晰可见，图片文件大小在 50K-5M 之间，不限制长宽，支持 png、bmp、gif、jpg、jpeg 格式。 请传入<a href='https://opendocs.alipay.com/apis/01ea4t'>alipay.open.sp.image.upload</a>(图片上传接口) 返回的 image_id。 该资质是否必传请参见<a href='https://opendocs.alipay.com/open/01hd83'>报名资质要求</a>。。
	FoodServiceLic *string `json:"food_service_lic,omitempty"`
	// IndoorPic 是内景照，要求图片清晰可见，图片文件大小在 50K-5M 之间，不限制长宽，支持 png、bmp、gif、jpg、jpeg 格式。 请传入<a href='https://opendocs.alipay.com/apis/01ea4t'>alipay.open.sp.image.upload</a>(图片上传接口) 返回的 image_id。 该资质是否必传请参见<a href='https://opendocs.alipay.com/open/01hd83'>报名资质要求</a>。。
	IndoorPic *string `json:"indoor_pic,omitempty"`
	// MerchantLogon 是参与蓝海活动的商户支付宝账号，只有当参与直连蓝海活动场景（BLUE_SEA_FOOD_APPLY/BLUE_SEA_FMCG_APPLY）时必填，间连场景可空。。
	MerchantLogon *string `json:"merchant_logon,omitempty"`
	// ProvinceCode 是省份编码。请按照<a href='https://gw.alipayobjects.com/os/basement_prod/253c4dcb-b8a4-4a1e-8be2-79e191a9b6db.xlsx'>表格</a> 中内容填写。 （请参考资料：<a href='http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/'> 统计用区划和城乡划分代码</a>）。
	ProvinceCode *string `json:"province_code,omitempty"`
	// ShopEntrancePic 是门头照，要求图片清晰可见，图片文件大小在 50K-5M 之间，不限制长宽，支持 png、bmp、gif、jpg、jpeg 格式。 请传入<a href='https://opendocs.alipay.com/apis/01ea4t'>alipay.open.sp.image.upload</a>(图片上传接口) 返回的 image_id。 该资质是否必传请参见<a href='https://opendocs.alipay.com/open/01hd83'>报名资质要求</a>。。
	ShopEntrancePic *string `json:"shop_entrance_pic,omitempty"`
	// SubMerchantId 是参与蓝海活动的间连商户账号，只有当参与间连蓝海活动场景（BLUE_SEA_FOOD_INDIRECT_APPLY/BLUE_SEA_FMCG_INDIRECT_APPLY）时必填，直连场景可空。。
	SubMerchantId *string `json:"sub_merchant_id,omitempty"`
	// TobaccoLic 是烟草专卖零售许可证，要求证件文本信息清晰可见，图片文件大小在 50K-5M 之间，不限制长宽，支持 png、bmp、gif、jpg、jpeg 格式。 请传入<a href='https://opendocs.alipay.com/apis/01ea4t'>alipay.open.sp.image.upload</a>(图片上传接口) 返回的 image_id。 该资质是否必传请参见<a href='https://opendocs.alipay.com/open/01hd83'>报名资质要求</a>。。
	TobaccoLic *string `json:"tobacco_lic,omitempty"`
}

// SpBlueSeaActivityModifyRequest 是支付宝官方 V3 OpenAPI 模型。
type SpBlueSeaActivityModifyRequest struct {
	// Address 是详细地址。
	Address *string `json:"address,omitempty"`
	// BusinessLic 是营业执照，要求证件文本信息清晰可见。 请上传照片的image_id，传参明细请参见<a href="https://opendocs.alipay.com/open/01hd83">报名资质要求</a>。
	BusinessLic *string `json:"business_lic,omitempty"`
	// CityCode 是城市编码。请按照https://gw.alipayobjects.com/os/basement_prod/253c4dcb-b8a4-4a1e-8be2-79e191a9b6db.xlsx 表格中内容填写。 （参考资料： http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/）。
	CityCode *string `json:"city_code,omitempty"`
	// DistrictCode 是区县编码。请按照https://gw.alipayobjects.com/os/basement_prod/253c4dcb-b8a4-4a1e-8be2-79e191a9b6db.xlsx 表格中内容填写。 （参考资料： http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/）。
	DistrictCode *string `json:"district_code,omitempty"`
	// FoodBusinessLic 是食品经营许可证，要求证件文本信息清晰可见。请上传照片的image_id，传参明细请参见<a href="https://opendocs.alipay.com/open/01hd83">报名资质要求</a>。
	FoodBusinessLic *string `json:"food_business_lic,omitempty"`
	// FoodCirculateLic 是食品流通许可证，要求证件文本信息清晰可见。请上传照片的image_id，传参明细请参见<a href="https://opendocs.alipay.com/open/01hd83">报名资质要求</a>。
	FoodCirculateLic *string `json:"food_circulate_lic,omitempty"`
	// FoodHealthLic 是食品卫生许可证，要求证件文本信息清晰可见。请上传照片的image_id，传参明细请参见<a href="https://opendocs.alipay.com/open/01hd83">报名资质要求</a>。
	FoodHealthLic *string `json:"food_health_lic,omitempty"`
	// FoodProductionLic 是食品生产许可证，要求证件文本信息清晰可见。请上传照片的image_id，传参明细请参见<a href="https://opendocs.alipay.com/open/01hd83">报名资质要求</a>。
	FoodProductionLic *string `json:"food_production_lic,omitempty"`
	// FoodServiceLic 是餐饮服务许可证，要求证件文本信息清晰可见。请上传照片的image_id，传参明细请参见<a href="https://opendocs.alipay.com/open/01hd83">报名资质要求</a>。
	FoodServiceLic *string `json:"food_service_lic,omitempty"`
	// IndoorPic 是门头照，要求内景照片清晰可见。请上传照片的image_id，传参明细请参见<a href="https://opendocs.alipay.com/open/01hd83">报名资质要求</a>。
	IndoorPic *string `json:"indoor_pic,omitempty"`
	// OrderId 是申请单 id。通过 <a href="https://opendocs.alipay.com/apis/01ebig">alipay.open.sp.blueseaactivity.create</a>接口获取。。
	OrderId *string `json:"order_id,omitempty"`
	// ProvinceCode 是省份编码。请按照https://gw.alipayobjects.com/os/basement_prod/253c4dcb-b8a4-4a1e-8be2-79e191a9b6db.xlsx 表格中内容填写。 （参考资料： http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/）。
	ProvinceCode *string `json:"province_code,omitempty"`
	// ShopEntrancePic 是门头照，要求店铺外观照片清晰可见。请上传照片的image_id，传参明细请参见<a href="https://opendocs.alipay.com/open/01hd83">报名资质要求</a>。
	ShopEntrancePic *string `json:"shop_entrance_pic,omitempty"`
	// TobaccoLic 是烟草专卖零售许可证，要求证件文本信息清晰可见。请上传照片的image_id，传参明细请参见<a href="https://opendocs.alipay.com/open/01hd83">报名资质要求</a>。
	TobaccoLic *string `json:"tobacco_lic,omitempty"`
}

// SpBlueSeaActivityQueryRequest is the request type for alipay.open.sp.blueseaactivity.query.
type SpBlueSeaActivityQueryRequest struct {
	// OrderId 是申请单 id。通过 alipay.open.sp.blueseaactivity.create接口获取
	OrderId *string `json:"order_id,omitempty" query:"order_id,omitempty"`
}
