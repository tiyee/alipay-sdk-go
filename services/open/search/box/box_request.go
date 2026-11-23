package box

// SearchBoxApplyRequest 是支付宝官方 V3 OpenAPI 模型。
type SearchBoxApplyRequest struct {
	// BoxDesc 是品牌介绍，5-15个中文字符。 小程序直达不支持设置此项。
	BoxDesc *string `json:"box_desc,omitempty"`
	// BrandId 是品牌id，参考<a href="https://opendocs.alipay.com/rules/029uy4"> 品牌认证说明 </a>。
	BrandId *string `json:"brand_id,omitempty"`
	// BusinessBenefitSwitch 是可通过配置来开启商圈权益模块，关闭后搜索直达不展示商圈权益模块。
	BusinessBenefitSwitch *bool `json:"business_benefit_switch,omitempty"`
	// BusinessDistrictIds 是小程序已关联商圈时，可添加商圈id（目前仅对品牌直达开放，小程序直达暂未开放）。
	BusinessDistrictIds *[]string `json:"business_district_ids,omitempty"`
	// CustomKeywords 是自定义触发词，最多可配置10个，限1-8个中文字符。 小程序直达不支持设置此项。
	CustomKeywords *[]string `json:"custom_keywords,omitempty"`
	// ImageId 是氛围图片id，调用 <a href="https://opendocs.alipay.com/mini/03hvl1?ref=api">支付宝文件上传接口</a> 上传图片获取图片id(bizCode：search_box_atmosphere)。 <a href="https://opendocs.alipay.com/b/03al6f">图片规范</a> 小程序直达不支持设置此项。。
	ImageId *string `json:"image_id,omitempty"`
	// ImageName 是氛围图片名。 小程序直达不支持设置此项。
	ImageName *string `json:"image_name,omitempty"`
	// MerchantId 是商户id，代运营模式下传入。代运营模式，需要服务商已获得商家"运营支付宝小程序"授权。。
	MerchantId *string `json:"merchant_id,omitempty"`
	// RelatedAccounts 是关联账号信息，1-3个。 内部字段均需设置。当为品牌直达时，数组中的第1个账号会被设置为"账号1"，也就是将作为搜索直达专区头部的跳转地址。
	RelatedAccounts *[]SearchBoxAppInfo `json:"related_accounts,omitempty"`
	// ServiceInfos 是"服务信息，服务必须审核通过才能申请搜索直达。品牌直达最多可配置同一品牌认证下的小程序4个，小程序直达最多可配置2个。 内部字段均需设置。"。
	ServiceInfos *[]SearchBoxServiceInfo `json:"service_infos,omitempty"`
	// TargetAppid 是小程序id，小程序直达时必传，需要和申请的商户主体保持一致，且符合<a href="https://opendocs.alipay.com/b/03al6e"> 准入类目 </a>。
	TargetAppid *string `json:"target_appid,omitempty"`
}

// SearchBoxBatchQueryRequest 是支付宝官方 V3 OpenAPI 模型。
type SearchBoxBatchQueryRequest struct {
	// MerchantId 是商户id，代运营模式下传入。代运营模式，需要服务商已获得商家"运营支付宝小程序"授权。。
	MerchantId *string `json:"merchant_id,omitempty"`
	// PageNumber 是分页查询的当前页号,从1开始。
	PageNumber *int64 `json:"page_number,omitempty"`
	// PageSize 是每页查询的数量，默认10，不得超过50。
	PageSize *int64 `json:"page_size,omitempty"`
}

// SearchBoxConsultRequest 是支付宝官方 V3 OpenAPI 模型。
type SearchBoxConsultRequest struct {
	// AreaKeywords 是小程序直达配置的常用服务中带有门店信息时，可添加简称触发词。
	AreaKeywords *string `json:"area_keywords,omitempty"`
	// BoxDesc 是品牌介绍，5-15个中文字符。 小程序直达不支持设置此项。
	BoxDesc *string `json:"box_desc,omitempty"`
	// BrandId 是品牌id，参考<a href="https://opendocs.alipay.com/rules/029uy4">品牌认证说明</a>。
	BrandId *string `json:"brand_id,omitempty"`
	// BusinessBenefitSwitch 是可通过配置来开启商圈权益模块，关闭后搜索直达不展示商圈权益模块。
	BusinessBenefitSwitch *bool `json:"business_benefit_switch,omitempty"`
	// BusinessDistrictIds 是小程序已关联商圈时，可添加商圈id（目前仅对品牌直达开放，小程序直达暂未开放）。
	BusinessDistrictIds *[]string `json:"business_district_ids,omitempty"`
	// CustomKeywords 是自定义触发词，最多可配置6个，限1-8个中文字符。 小程序直达不支持设置此项。
	CustomKeywords *[]string `json:"custom_keywords,omitempty"`
	// ImageId 是氛围图片id，调用 <a href="https://opendocs.alipay.com/mini/03hvl1?ref=api">支付宝文件上传接口</a> 上传图片获取图片id(bizCode：search_box_atmosphere)。 <a href="https://opendocs.alipay.com/b/03al6f">图片规范</a> 小程序直达不支持设置此项。。
	ImageId *string `json:"image_id,omitempty"`
	// ImageName 是氛围图片名。 小程序直达不支持设置此项。
	ImageName *string `json:"image_name,omitempty"`
	// MerchantId 是商户id，代运营模式下传入。代运营模式，需要服务商已获得商家"运营支付宝小程序"授权。。
	MerchantId *string `json:"merchant_id,omitempty"`
	// RelatedAccounts 是关联账号信息，1-2个。 内部字段均需设置。。
	RelatedAccounts *[]SearchBoxAppInfo `json:"related_accounts,omitempty"`
	// ServiceInfos 是服务信息，服务必须审核通过才能申请搜索直达。 内部字段均需设置。。
	ServiceInfos *[]SearchBoxServiceInfo `json:"service_infos,omitempty"`
	// TargetAppid 是小程序id，小程序直达时必传，需要和申请的商户主体保持一致，且符合<a href="https://opendocs.alipay.com/b/03al6e">准入类目</a>。
	TargetAppid *string `json:"target_appid,omitempty"`
}

// SearchBoxModifyRequest 是支付宝官方 V3 OpenAPI 模型。
type SearchBoxModifyRequest struct {
	// BoxDesc 是品牌介绍，5-15个中文字符。当修改品牌介绍模块(module_type=BOX_EXCLUSIVE_BASE)时传入。 小程序直达时不支持设置。
	BoxDesc *string `json:"box_desc,omitempty"`
	// BoxId 是搜索直达配置id。
	BoxId *string `json:"box_id,omitempty"`
	// BrandId 是品牌id，参考<a href="https://opendocs.alipay.com/rules/029uy4"> 品牌认证说明 </a>。
	BrandId *string `json:"brand_id,omitempty"`
	// BusinessBenefitSwitch 是可通过配置来开启商圈权益模块，关闭后搜索直达不展示商圈权益模块。
	BusinessBenefitSwitch *bool `json:"business_benefit_switch,omitempty"`
	// BusinessDistrictIds 是小程序已关联商圈时，可添加商圈id（目前仅对品牌直达开放，小程序直达暂未开放）。
	BusinessDistrictIds *[]string `json:"business_district_ids,omitempty"`
	// CustomKeywords 是自定义触发词。
	CustomKeywords *string `json:"custom_keywords,omitempty"`
	// ImageId 是氛围图片id，调用 <a href="https://opendocs.alipay.com/mini/03hvl1?ref=api">支付宝文件上传接口</a> 上传图片获取图片id(bizCode：search_box_atmosphere)。 <a href="https://opendocs.alipay.com/b/03al6f">图片规范</a> 小程序直达不支持设置此项。。
	ImageId *string `json:"image_id,omitempty"`
	// ImageName 是氛围图片名，当修改氛围图模块(module_type=BOX_ATMOSPHERE_IMAGE)时传入。 小程序直达不支持设置。
	ImageName *string `json:"image_name,omitempty"`
	// MerchantId 是商户id，代运营模式下传入。代运营模式，需要服务商已获得商家"运营支付宝小程序"授权。。
	MerchantId *string `json:"merchant_id,omitempty"`
	// ModuleId 是搜索直达模块配置id，新增模块时可不传，更新模块时必传。从alipay.open.search.box.query详情接口查询可获取已存在模块的配置id。
	ModuleId *string `json:"module_id,omitempty"`
	// ModuleType 是搜索直达模块类型，参考<a href="https://opendocs.alipay.com/mini/03fjba#%E4%BD%BF%E7%94%A8%E8%AF%B4%E6%98%8E"> 搜索直达模块说明 </a>。。
	ModuleType *string `json:"module_type,omitempty"`
	// RelatedAccounts 是"关联账号信息，可配置1-2个。传入账号需归属于商家主体。品牌直达修改官方账号(module_type=BOX_EXCLUSIVE_ACCOUNTS)时传入，小程序直达修改常用服务模块时也需传值。 当为品牌直达时，数组中的第1个账号会被设置为"账号1"，也就是将作为搜索直达专区头部的跳转地址。"。
	RelatedAccounts *[]SearchBoxAppInfo `json:"related_accounts,omitempty"`
	// ServiceInfos 是服务信息，服务必须审核通过才能申请搜索直达，可配置1-4个。当修改常用服务模块(module_type=BOX_EXCLUSIVE_FUNCTIONS)时传入。
	ServiceInfos *[]SearchBoxServiceInfo `json:"service_infos,omitempty"`
	// TargetAppid 是小程序id，小程序直达时必传，需要和申请的商户主体保持一致，且符合<a href="https://opendocs.alipay.com/b/03al6e"> 准入类目 </a>。
	TargetAppid *string `json:"target_appid,omitempty"`
}

// SearchBoxOfflineRequest 是支付宝官方 V3 OpenAPI 模型。
type SearchBoxOfflineRequest struct {
	// BoxId 是搜索直达配置id。
	BoxId *string `json:"box_id,omitempty"`
	// MerchantId 是商户id，代运营模式下传入。代运营模式，需要服务商已获得商家"运营支付宝小程序"授权。。
	MerchantId *string `json:"merchant_id,omitempty"`
}

// SearchBoxOnlineRequest 是支付宝官方 V3 OpenAPI 模型。
type SearchBoxOnlineRequest struct {
	// BoxId 是搜索直达配置id。
	BoxId *string `json:"box_id,omitempty"`
	// MerchantId 是商户id，代运营模式下传入。代运营模式，需要服务商已获得商家"运营支付宝小程序"授权。。
	MerchantId *string `json:"merchant_id,omitempty"`
}

// SearchBoxActivityApplyRequest 是支付宝官方 V3 OpenAPI 模型。
type SearchBoxActivityApplyRequest struct {
	// ActionUrl 是活动链接，必须为关联的小程序的页面链接（链接以 alipays 开头），可参考<a href="https://opendocs.alipay.com/support/01rb18"> 小程序scheme链接介绍 </a>。
	ActionUrl *string `json:"action_url,omitempty"`
	// BackgroundWord 是创建banner活动时，可添加底纹推广词，有机会在首页搜索框内展示，提升banner活动曝光机会。
	BackgroundWord *string `json:"background_word,omitempty"`
	// BoxId 是搜索直达配置id。
	BoxId *string `json:"box_id,omitempty"`
	// EndTime 是活动结束时间，开始和结束时间之间不得超过90天。
	EndTime *string `json:"end_time,omitempty"`
	// MaterialId 是当material_type="IMAGE"时，为图片id；当material_type="VIDEO"时，为视频id。小程序直达仅支持图片类型。图片id可以通过调用<a href="https://opendocs.alipay.com/mini/03hvl1?ref=api"> 支付宝文件上传接口 </a>上传图片，获取图片id(bizCode：search_box_banner)。<a href="https://opendocs.alipay.com/b/03al6f"> 图片规范 </a>。
	MaterialId *string `json:"material_id,omitempty"`
	// MaterialType 是IMAGE-图片/VIDEO-视频。
	MaterialType *string `json:"material_type,omitempty"`
	// MerchantId 是商户id，代运营模式下传入。代运营模式，需要服务商已获得商家"运营支付宝小程序"授权。。
	MerchantId   *string                   `json:"merchant_id,omitempty"`
	OrientedRule *DeliveryOrientedRuleInfo `json:"oriented_rule,omitempty"`
	// StartTime 是活动开始时间，开始和结束时间之间不得超过90天。
	StartTime *string `json:"start_time,omitempty"`
	// TargetAppid 是跳转应用ID。
	TargetAppid *string `json:"target_appid,omitempty"`
	// TargetAppname 是跳转小程序名称。
	TargetAppname *string `json:"target_appname,omitempty"`
	// TargetRegions 是投放目标区域列表。
	TargetRegions *[]DeliveryTargetRegion `json:"target_regions,omitempty"`
	// Title 是活动标题。
	Title     *string                     `json:"title,omitempty"`
	VideoInfo *SearchBoxActivityVideoInfo `json:"video_info,omitempty"`
}

// SearchBoxQueryRequest is the request type for alipay.open.search.box.query.
type SearchBoxQueryRequest struct {
	// BoxId 是搜索直达配置id
	BoxId *string `json:"box_id,omitempty" query:"box_id,omitempty"`
	// MerchantId 是商户id，代运营模式下传入。代运营模式，需要服务商已获得商家"运营支付宝小程序"授权。
	MerchantId *string `json:"merchant_id,omitempty" query:"merchant_id,omitempty"`
}

// SearchBoxActivityQueryRequest is the request type for alipay.open.search.boxactivity.query.
type SearchBoxActivityQueryRequest struct {
	// BoxActivityId 是搜索直达活动id
	BoxActivityId *string `json:"box_activity_id,omitempty" query:"box_activity_id,omitempty"`
	// MerchantId 是商户id，代运营模式下传入。代运营模式，需要服务商已获得商家"运营支付宝小程序"授权。
	MerchantId *string `json:"merchant_id,omitempty" query:"merchant_id,omitempty"`
}
