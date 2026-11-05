// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package merchant

// ExpandIndirectImageUpload is the official V3 AntMerchantExpandIndirectImageUploadModel schema.
type ExpandIndirectImageUpload struct {
	// ImageType represents image_type. 图片格式，支持格式：bmp、jpg、jpeg、png、gif.
	ImageType *string `json:"image_type,omitzero"`
}

// MerchantQueryResult is the official V3 MerchantQueryResult schema.
type MerchantQueryResult struct {
	// AliasName represents alias_name. 商户的别名，用户商户对客展示的名称
	AliasName *string `json:"alias_name,omitzero"`
	// CertNo represents cert_no. 商户认证证件号，企业营业执照号
	CertNo *string `json:"cert_no,omitzero"`
	// City represents city. 市名称
	City *string `json:"city,omitzero"`
	// DetailAddress represents detail_address. 商户经营详细地址
	DetailAddress *string `json:"detail_address,omitzero"`
	// Distinct represents distinct. 区县名称
	Distinct *string `json:"distinct,omitzero"`
	// MccCode represents mcc_code. 商户新版mcc code
	MccCode *string `json:"mcc_code,omitzero"`
	// MerchantType represents merchant_type. 描述商户类型，个人-P/企业-B
	MerchantType *string `json:"merchant_type,omitzero"`
	// Name represents name. 商户认证名称信息
	Name *string `json:"name,omitzero"`
	// Province represents province. 省名称
	Province *string `json:"province,omitzero"`
}

// ItemLabelCreateInfo is the official V3 ItemLabelCreateInfo schema.
type ItemLabelCreateInfo struct {
	// LabelKey represents label_key. 标签键 OIL_NUM：油号
	LabelKey *string `json:"label_key,omitzero"`
	// LabelValue represents label_value. 标签值： OIL_NUM对应的值为：95# 92# 98# 等等
	LabelValue *string `json:"label_value,omitzero"`
}

// ItemSkuCreateInfo is the official V3 ItemSkuCreateInfo schema.
type ItemSkuCreateInfo struct {
	// CostPrice represents cost_price. 成本价，单位分
	CostPrice *int64 `json:"cost_price,omitzero"`
	// ExtInfo represents ext_info. SKU扩展信息
	ExtInfo *[]ItemExtInfo `json:"ext_info,omitzero"`
	// ExternalSkuId represents external_sku_id. 外部SKU ID
	ExternalSkuId *string `json:"external_sku_id,omitzero"`
	// Inventory represents inventory. 库存
	Inventory *int64 `json:"inventory,omitzero"`
	// OriginalPrice represents original_price. 标价，单位分
	OriginalPrice *int64 `json:"original_price,omitzero"`
	// Price represents price. 售价，单位分
	Price *int64 `json:"price,omitzero"`
}

// ItemLabelModifyInfo is the official V3 ItemLabelModifyInfo schema.
type ItemLabelModifyInfo struct {
	// LabelKey represents label_key. 标签键 OIL_NUM：油号
	LabelKey *string `json:"label_key,omitzero"`
	// LabelValue represents label_value. 标签值： OIL_NUM对应的值为：95# 92# 98# 等等
	LabelValue *string `json:"label_value,omitzero"`
}

// ItemSkuModifyInfo is the official V3 ItemSkuModifyInfo schema.
type ItemSkuModifyInfo struct {
	// CostPrice represents cost_price. 成本价，单位分
	CostPrice *int64 `json:"cost_price,omitzero"`
	// ExtInfo represents ext_info. SKU扩展信息
	ExtInfo *[]ItemExtInfo `json:"ext_info,omitzero"`
	// ExternalSkuId represents external_sku_id. 外部skuId（sku_id不为空时作为普通更新项,sku_id为空时作为更新key,sku_id和该字段同时为空则新增该条数据）
	ExternalSkuId *string `json:"external_sku_id,omitzero"`
	// Inventory represents inventory. 库存
	Inventory *int64 `json:"inventory,omitzero"`
	// OriginalPrice represents original_price. 标价，单位分
	OriginalPrice *int64 `json:"original_price,omitzero"`
	// Price represents price. 售价，单位分
	Price *int64 `json:"price,omitzero"`
	// SkuId represents sku_id. skuId
	SkuId *string `json:"sku_id,omitzero"`
	// Status represents status. 商品SKU状态：EFFECT、INVALID
	Status *string `json:"status,omitzero"`
}

// CmItemInfo is the official V3 CmItemInfo schema.
type CmItemInfo struct {
	// Description represents description. 商品描述
	Description *string `json:"description,omitzero"`
	// ExtInfo represents ext_info. 商品扩展信息
	ExtInfo *[]ItemExtInfo `json:"ext_info,omitzero"`
	// FrontCategoryIdList represents front_category_id_list. 商品所属前台类目ID列表
	FrontCategoryIdList *[]string `json:"front_category_id_list,omitzero"`
	// GmtCreate represents gmt_create. 商品创建时间
	GmtCreate *string `json:"gmt_create,omitzero"`
	// GmtModified represents gmt_modified. 商品更新时间
	GmtModified *string `json:"gmt_modified,omitzero"`
	// ItemId represents item_id. 商品ID
	ItemId *string `json:"item_id,omitzero"`
	// MaterialList represents material_list. 商品素材列表
	MaterialList *[]MaterialInfo `json:"material_list,omitzero"`
	// Name represents name. 商品名称
	Name *string `json:"name,omitzero"`
	// PropertyList represents property_list. 商品属性列表
	PropertyList *[]ItemPropertyInfo `json:"property_list,omitzero"`
	// SkuList represents sku_list. 商品SKU列表
	SkuList *[]CmItemSkuInfo `json:"sku_list,omitzero"`
	// StandardCategoryId represents standard_category_id. 商品所属标准类目ID（具体值请参见产品文档）
	StandardCategoryId *string `json:"standard_category_id,omitzero"`
	// Status represents status. 商品状态
	Status *string `json:"status,omitzero"`
	// TargetId represents target_id. 商品归属主体ID 例：商品归属主体类型为店铺，则商品归属主体ID为店铺ID；归属主体为小程序，则归属主体ID为小程序ID
	TargetId *string `json:"target_id,omitzero"`
	// TargetType represents target_type. 商品归属主体类型: 5（店铺） 8（小程序）
	TargetType *string `json:"target_type,omitzero"`
	// Type represents type. 商品类型： STANDARD_GOODS（标品）、NON_STANDARD_GOODS（非标品）
	Type *string `json:"type,omitzero"`
}

// ItemExtInfo is the official V3 ItemExtInfo schema.
type ItemExtInfo struct {
	// ExtKey represents ext_key. 扩展信息的key
	ExtKey *string `json:"ext_key,omitzero"`
	// ExtValue represents ext_value. 扩展信息的值
	ExtValue *string `json:"ext_value,omitzero"`
}

// MaterialCreateInfo is the official V3 MaterialCreateInfo schema.
type MaterialCreateInfo struct {
	// Content represents content. 素材内容（素材地址或素材KEY）
	Content *string `json:"content,omitzero"`
	// Type represents type. 素材类型
	Type *string `json:"type,omitzero"`
}

// ItemPropertyInfo is the official V3 ItemPropertyInfo schema.
type ItemPropertyInfo struct {
	// PropertyKey represents property_key. 属性名称
	PropertyKey *string `json:"property_key,omitzero"`
	// PropertyValueList represents property_value_list. 属性值列表
	PropertyValueList *[]string `json:"property_value_list,omitzero"`
}

// SkuCreateInfo is the official V3 SkuCreateInfo schema.
type SkuCreateInfo struct {
	// Inventory represents inventory. 库存
	Inventory *int64 `json:"inventory,omitzero"`
	// MaterialList represents material_list. SKU素材列表（最多3个）
	MaterialList *[]MaterialCreateInfo `json:"material_list,omitzero"`
	// OriginalPrice represents original_price. 标价，单位分
	OriginalPrice *int64 `json:"original_price,omitzero"`
	// Price represents price. 售价，单位分
	Price *int64 `json:"price,omitzero"`
	// PropertyList represents property_list. SKU属性列表
	PropertyList *[]ItemSkuPropertyInfo `json:"property_list,omitzero"`
}

// MaterialModifyInfo is the official V3 MaterialModifyInfo schema.
type MaterialModifyInfo struct {
	// Content represents content. 素材内容（素材地址或素材KEY）
	Content *string `json:"content,omitzero"`
	// Type represents type. 素材类型
	Type *string `json:"type,omitzero"`
}

// SkuModifyInfo is the official V3 SkuModifyInfo schema.
type SkuModifyInfo struct {
	// Inventory represents inventory. 库存
	Inventory *int64 `json:"inventory,omitzero"`
	// MaterialList represents material_list. SKU素材列表（会和SKU已存在素材做差异化比较后做增删改操作）
	MaterialList *[]MaterialModifyInfo `json:"material_list,omitzero"`
	// OriginalPrice represents original_price. 标价，单位分
	OriginalPrice *int64 `json:"original_price,omitzero"`
	// Price represents price. 售价，单位分
	Price *int64 `json:"price,omitzero"`
	// PropertyList represents property_list. SKU属性列表（覆盖SKU已存在属性）
	PropertyList *[]ItemSkuPropertyInfo `json:"property_list,omitzero"`
	// SkuId represents sku_id. SKU ID（传空表示新增对应SKU）
	SkuId *string `json:"sku_id,omitzero"`
}

// Item is the official V3 ItemModel schema.
type Item struct {
	// CreatorId represents creator_id. 创建者id
	CreatorId *string `json:"creator_id,omitzero"`
	// CreatorType represents creator_type. 创建者类型
	CreatorType *string `json:"creator_type,omitzero"`
	// ExtInfo represents ext_info. 商品扩展信息：可以解析成 Map<String, String> 的 json string
	ExtInfo *string `json:"ext_info,omitzero"`
	// FrontCategoryId represents front_category_id. 前台类目id
	FrontCategoryId *string `json:"front_category_id,omitzero"`
	// GmtCreate represents gmt_create. 商品创建时间
	GmtCreate *string `json:"gmt_create,omitzero"`
	// GmtModified represents gmt_modified. 商品更新时间
	GmtModified *string `json:"gmt_modified,omitzero"`
	// ItemId represents item_id. 商品id
	ItemId *string `json:"item_id,omitzero"`
	// Name represents name. 商品名称
	Name *string `json:"name,omitzero"`
	// SkuList represents sku_list. 商品sku列表
	SkuList *[]ItemSkuInfo `json:"sku_list,omitzero"`
	// Sort represents sort. 商品排序
	Sort *int64 `json:"sort,omitzero"`
	// Status represents status. 商品状态
	Status *string `json:"status,omitzero"`
	// TargetId represents target_id. 商户归属主体id
	TargetId *string `json:"target_id,omitzero"`
	// TargetType represents target_type. 商品归属主体类型：5(商铺)
	TargetType *string `json:"target_type,omitzero"`
	// Type represents type. 商品类型
	Type *string `json:"type,omitzero"`
}

// SettleCardInfo is the official V3 SettleCardInfo schema.
type SettleCardInfo struct {
	// AccountBranchName represents account_branch_name. 开户支行名
	AccountBranchName *string `json:"account_branch_name,omitzero"`
	// AccountHolderName represents account_holder_name. 卡户名
	AccountHolderName *string `json:"account_holder_name,omitzero"`
	// AccountInstCity represents account_inst_city. 开户行所在地-市
	AccountInstCity *string `json:"account_inst_city,omitzero"`
	// AccountInstId represents account_inst_id. 开户行简称缩写
	AccountInstId *string `json:"account_inst_id,omitzero"`
	// AccountInstName represents account_inst_name. 银行名称
	AccountInstName *string `json:"account_inst_name,omitzero"`
	// AccountInstProvince represents account_inst_province. 开户行所在地-省
	AccountInstProvince *string `json:"account_inst_province,omitzero"`
	// AccountNo represents account_no. 银行卡号
	AccountNo *string `json:"account_no,omitzero"`
	// AccountType represents account_type. 卡类型 借记卡-DC 信用卡-CC
	AccountType *string `json:"account_type,omitzero"`
	// BankCode represents bank_code. 联行号，详细需调用<a href="https://opendocs.alipay.com/apis/0853v1">‘联行号关联分支行查询’</a>API查看查询结果.
	BankCode *string `json:"bank_code,omitzero"`
	// UsageType represents usage_type. 账号使用类型 对公-01 对私-02
	UsageType *string `json:"usage_type,omitzero"`
}

// AddressInfo is the official V3 AddressInfo schema.
type AddressInfo struct {
	// Address represents address. 地址。商户详细经营地址或人员所在地点
	Address *string `json:"address,omitzero"`
	// CityCode represents city_code. 城市编码。 蚂蚁店铺请按照<a href="https://mdn.alipayobjects.com/huamei_sm7gq8/afts/file/A*blT9RqSR9_gAAAAAAAAAAAAADuKQAQ/2022%E8%9A%82%E8%9A%81%E9%87%91%E6%9C%8D%E5%9C%B0%E5%8C%BA%E7%A0%81.xlsx" target="_blank">蚂蚁店铺地区码</a> 表格中填写。 直付通商户请按照<a href="https://gw.alipayobjects.
	CityCode *string `json:"city_code,omitzero"`
	// DistrictCode represents district_code. 区县编码。 蚂蚁店铺请按照<a href="https://mdn.alipayobjects.com/huamei_sm7gq8/afts/file/A*blT9RqSR9_gAAAAAAAAAAAAADuKQAQ/2022%E8%9A%82%E8%9A%81%E9%87%91%E6%9C%8D%E5%9C%B0%E5%8C%BA%E7%A0%81.xlsx" target="_blank">蚂蚁店铺地区码</a> 表格中填写。 直付通商户请按照<a href="https://gw.alipayobjects.
	DistrictCode *string `json:"district_code,omitzero"`
	// Latitude represents latitude. 纬度，浮点型,小数点后最多保留6位 如需要录入经纬度，请以高德坐标系为准，录入时请确保经纬度参数准确。<a href="http://lbs.amap.com/console/show/picker">高德经纬度查询</a>
	Latitude *string `json:"latitude,omitzero"`
	// Longitude represents longitude. 经度，浮点型, 小数点后最多保留6位。 如需要录入经纬度，请以高德坐标系为准，录入时请确保经纬度参数准确。<a href="http://lbs.amap.com/console/show/picker">高德经纬度查询</a>
	Longitude *string `json:"longitude,omitzero"`
	// Poiid represents poiid. 高德poiid
	Poiid *string `json:"poiid,omitzero"`
	// ProvinceCode represents province_code. 省份编码。 蚂蚁店铺请按照<a href="https://mdn.alipayobjects.com/huamei_sm7gq8/afts/file/A*blT9RqSR9_gAAAAAAAAAAAAADuKQAQ/2022%E8%9A%82%E8%9A%81%E9%87%91%E6%9C%8D%E5%9C%B0%E5%8C%BA%E7%A0%81.xlsx" target="_blank">蚂蚁店铺地区码</a> 表格中填写。 直付通商户请按照<a href="https://gw.alipayobjects.
	ProvinceCode *string `json:"province_code,omitzero"`
	// Type represents type. 地址类型。取值范围：BUSINESS_ADDRESS：经营地址（默认）
	Type *string `json:"type,omitzero"`
}

// ShopBusinessTime is the official V3 ShopBusinessTime schema.
type ShopBusinessTime struct {
	// CloseTime represents close_time. 关门时间 格式：HH:mm
	CloseTime *string `json:"close_time,omitzero"`
	// OpenTime represents open_time. 开门时间 格式：HH:mm
	OpenTime *string `json:"open_time,omitzero"`
	// WeekDay represents week_day. 本对象表示周几的营业时间。1~6表示周一到周六，7表示周日
	WeekDay *int64 `json:"week_day,omitzero"`
}

// ContactInfo is the official V3 ContactInfo schema.
type ContactInfo struct {
	// Email represents email. 电子邮箱
	Email *string `json:"email,omitzero"`
	// IdCardNo represents id_card_no. 身份证号
	IdCardNo *string `json:"id_card_no,omitzero"`
	// Mobile represents mobile. 手机号
	Mobile *string `json:"mobile,omitzero"`
	// Name represents name. 联系人名字
	Name *string `json:"name,omitzero"`
	// Phone represents phone. 电话
	Phone *string `json:"phone,omitzero"`
	// Tag represents tag. 商户联系人业务标识枚举，表示商户联系人的职责
	Tag *[]string `json:"tag,omitzero"`
	// Type represents type. 联系人类型，取值范围：LEGAL_PERSON：法人；CONTROLLER：实际控制人；AGENT：代理人；OTHER：其他
	Type *string `json:"type,omitzero"`
}

// ShopExtInfo is the official V3 ShopExtInfo schema.
type ShopExtInfo struct {
	// KeyName represents key_name. KV数据对的key，表示该kv对象表示什么含义。
	KeyName *string `json:"key_name,omitzero"`
	// Value represents value. kv对象的值
	Value *string `json:"value,omitzero"`
}

// IndustryQualificationInfo is the official V3 IndustryQualificationInfo schema.
type IndustryQualificationInfo struct {
	// IndustryQualificationImage represents industry_qualification_image. 商户行业资质图片。其值为通过ant.merchant.expand.indirect.image.upload上传图片得到的image_id
	IndustryQualificationImage *string `json:"industry_qualification_image,omitzero"`
	// IndustryQualificationType represents industry_qualification_type. <a href="https://gw.alipayobjects.com/os/bmw-prod/7aa3a36b-2bc2-4d57-815f-08edd55ef67e.xlsx">商户行业资质类型，具体选值参见文档</a>
	IndustryQualificationType *string `json:"industry_qualification_type,omitzero"`
}

// MCCQueryInfo is the official V3 MCCQueryInfo schema.
type MCCQueryInfo struct {
	// IsSpecial represents is_special. 是否特殊行业
	IsSpecial *bool `json:"is_special,omitzero"`
	// MccLevel1 represents mcc_level_1. 一级类目code
	MccLevel1 *string `json:"mcc_level_1,omitzero"`
	// MccLevel1Name represents mcc_level_1_name. 商户一级类目名称
	MccLevel1Name *string `json:"mcc_level_1_name,omitzero"`
	// MccLevel2 represents mcc_level_2. 二级类目code
	MccLevel2 *string `json:"mcc_level_2,omitzero"`
	// MccLevel2Name represents mcc_level_2_name. 二级类目名称
	MccLevel2Name *string `json:"mcc_level_2_name,omitzero"`
	// MccRequirements represents mcc_requirements. 特殊行业需要上传的资质
	MccRequirements *string `json:"mcc_requirements,omitzero"`
	// SpecialQualRequired represents special_qual_required. 是否需要特殊资质
	SpecialQualRequired *bool `json:"special_qual_required,omitzero"`
}

// MerchantShopIndustryInfo is the official V3 MerchantShopIndustryInfo schema.
type MerchantShopIndustryInfo struct {
	// InfoCode represents info_code. 行业信息CODE
	InfoCode *string `json:"info_code,omitzero"`
	// InfoValue represents info_value. 行业信息值，具体值根据info_code决定。 当info_code=SHOP_MEDICAL_LOGO/SHOP_MEDICAL_IN_PIC这类文件信息时， 需要传入ant.merchant.expand.indirect.image.upload接口上传的文件key；当info_code= SHOP_MEDICAL_LEVEL_TAG时，需要<a href="https://mdn.alipayobjects.com/mrcheshop/afts/file/ZuDmQKgBifgAAAAAAAAAAAAADi
	InfoValue *[]string `json:"info_value,omitzero"`
}

// ShopIndustryLicense is the official V3 ShopIndustryLicense schema.
type ShopIndustryLicense struct {
	// ExpireTime represents expire_time. 证件过期时间
	ExpireTime *string `json:"expire_time,omitzero"`
	// LicenseNo represents license_no. 资质证件编号
	LicenseNo *string `json:"license_no,omitzero"`
	// LicensePic represents license_pic. 资质图片信息
	LicensePic *string `json:"license_pic,omitzero"`
	// LicenseType represents license_type. 行业资质类型，需使用<a href="https://mdn.alipayobjects.com/mrcheshop/afts/file/mS30TKabvCMAAAAAQKAAAAgAeiWmAQBr?af_fileName=%E8%B5%84%E8%B4%A8%E5%88%97%E8%A1%A8.xlsx">资质列表</a>中类型
	LicenseType *string `json:"license_type,omitzero"`
}

// ComplexBusinessTime is the official V3 ComplexBusinessTime schema.
type ComplexBusinessTime struct {
	// TimePeriod represents time_period. 营业时间段，支持多个时间段，例如08:00-11:00，14:00-16:00
	TimePeriod *[]ComplexBusinessTimePeriod `json:"time_period,omitzero"`
	// WeekDay represents week_day. 本对象表示周几的营业时间。1~6表示周一到周六，7表示周日
	WeekDay *int64 `json:"week_day,omitzero"`
}

// ShopRecommendInfo is the official V3 ShopRecommendInfo schema.
type ShopRecommendInfo struct {
	// Recommend represents recommend. 门店修改建议
	Recommend *string `json:"recommend,omitzero"`
	// RecommendAddress represents recommend_address. 推荐详细地址
	RecommendAddress *string `json:"recommend_address,omitzero"`
	// RecommendLatitude represents recommend_latitude. 推荐纬度，单位度
	RecommendLatitude *string `json:"recommend_latitude,omitzero"`
	// RecommendLongtitude represents recommend_longtitude. 推荐经度，单位为度
	RecommendLongtitude *string `json:"recommend_longtitude,omitzero"`
	// RecommendName represents recommend_name. 推荐门店名称
	RecommendName *string `json:"recommend_name,omitzero"`
	// UnconfidenceReason represents unconfidence_reason. 门店不置信原因
	UnconfidenceReason *string `json:"unconfidence_reason,omitzero"`
}

// ShopQueryOpenAPIVO is the official V3 ShopQueryOpenAPIVO schema.
type ShopQueryOpenAPIVO struct {
	// BusinessAddress represents business_address.
	BusinessAddress *AddressInfo `json:"business_address,omitzero"`
	// BusinessTime represents business_time. 店铺经营时间
	BusinessTime *[]ShopBusinessTime `json:"business_time,omitzero"`
	// ContactMobile represents contact_mobile. 店铺联系手机
	ContactMobile *string `json:"contact_mobile,omitzero"`
	// ContactPhone represents contact_phone. 店铺的联系固话
	ContactPhone *string `json:"contact_phone,omitzero"`
	// NewBusinessTime represents new_business_time. 门店营业时间，支持一周7天，支持1天内多时间段
	NewBusinessTime *[]ComplexBusinessTime `json:"new_business_time,omitzero"`
	// NewShopCategory represents new_shop_category. 新版门店类目标准二级类目code。类目标准及与原shop_category映射关系参见文档https://gw.alipayobjects.com/os/bmw-prod/4b3f82df-e53e-4b84-bc41-fe025101e726.xlsx
	NewShopCategory *string `json:"new_shop_category,omitzero"`
	// ShopCategory represents shop_category. 店铺类目.取值参见文件<a href="https://mif-pub.alipayobjects.com/ShopCategory.xlsx">文件</a>中的三级门店类目
	ShopCategory *string `json:"shop_category,omitzero"`
	// ShopId represents shop_id. 蚂蚁店铺id
	ShopId *string `json:"shop_id,omitzero"`
	// ShopInfoStatus represents shop_info_status. 当前名称、地址、经纬度信息准确一致，可用于数字化经营场景消费(如商品、券、消费圈等场域的分发)，不影响门店支付结算
	ShopInfoStatus *string `json:"shop_info_status,omitzero"`
	// ShopName represents shop_name. 店铺名称
	ShopName *string `json:"shop_name,omitzero"`
	// ShopStatus represents shop_status. 门店状态
	ShopStatus *string `json:"shop_status,omitzero"`
	// ShopType represents shop_type. 店铺经营类型
	ShopType *string `json:"shop_type,omitzero"`
	// StoreId represents store_id. 门店编号，表示该门店在该商户角色id(直连pid，间连smid)下，由商户自己定义的外部门店编号
	StoreId *string `json:"store_id,omitzero"`
}

// MaterialInfo is the official V3 MaterialInfo schema.
type MaterialInfo struct {
	// Content represents content. 素材内容（素材地址或素材KEY）
	Content *string `json:"content,omitzero"`
	// MaterialId represents material_id. 素材ID
	MaterialId *string `json:"material_id,omitzero"`
	// Type represents type. 素材类型
	Type *string `json:"type,omitzero"`
}

// CmItemSkuInfo is the official V3 CmItemSkuInfo schema.
type CmItemSkuInfo struct {
	// GmtCreate represents gmt_create. 创建时间
	GmtCreate *string `json:"gmt_create,omitzero"`
	// GmtModified represents gmt_modified. 更新时间
	GmtModified *string `json:"gmt_modified,omitzero"`
	// Inventory represents inventory. 库存
	Inventory *int64 `json:"inventory,omitzero"`
	// ItemId represents item_id. 商品ID
	ItemId *string `json:"item_id,omitzero"`
	// MaterialList represents material_list. 素材列表
	MaterialList *[]MaterialInfo `json:"material_list,omitzero"`
	// OriginalPrice represents original_price. 标价，单位分
	OriginalPrice *int64 `json:"original_price,omitzero"`
	// Price represents price. 售价，单位分
	Price *int64 `json:"price,omitzero"`
	// RemainInventory represents remain_inventory. 剩余库存
	RemainInventory *int64 `json:"remain_inventory,omitzero"`
	// SkuId represents sku_id. 商品SKU ID
	SkuId *string `json:"sku_id,omitzero"`
	// SkuPropertyList represents sku_property_list. SKU属性列表
	SkuPropertyList *[]ItemSkuPropertyInfo `json:"sku_property_list,omitzero"`
	// Status represents status. 商品SKU状态
	Status *string `json:"status,omitzero"`
}

// ItemSkuPropertyInfo is the official V3 ItemSkuPropertyInfo schema.
type ItemSkuPropertyInfo struct {
	// PropertyKey represents property_key. 属性
	PropertyKey *string `json:"property_key,omitzero"`
	// PropertyValue represents property_value. 属性值
	PropertyValue *string `json:"property_value,omitzero"`
}

// ItemSkuInfo is the official V3 ItemSkuInfo schema.
type ItemSkuInfo struct {
	// CostPrice represents cost_price. 成本价，单位分
	CostPrice *int64 `json:"cost_price,omitzero"`
	// ExtInfo represents ext_info. SKU扩展信息
	ExtInfo *[]ItemExtInfo `json:"ext_info,omitzero"`
	// GmtCreate represents gmt_create. 创建时间
	GmtCreate *string `json:"gmt_create,omitzero"`
	// GmtModified represents gmt_modified. 更新时间
	GmtModified *string `json:"gmt_modified,omitzero"`
	// ItemId represents item_id. 商品id
	ItemId *string `json:"item_id,omitzero"`
	// OriginalPrice represents original_price. 标价，单位分
	OriginalPrice *int64 `json:"original_price,omitzero"`
	// Price represents price. 售价，单位分
	Price *int64 `json:"price,omitzero"`
	// SkuId represents sku_id. 商品sku id
	SkuId *string `json:"sku_id,omitzero"`
	// Status represents status. 商品SKU状态：EFFECT、INVALID
	Status *string `json:"status,omitzero"`
}

// ComplexBusinessTimePeriod is the official V3 ComplexBusinessTimePeriod schema.
type ComplexBusinessTimePeriod struct {
	// CloseTime represents close_time. 关门时间 格式：HH:mm
	CloseTime *string `json:"close_time,omitzero"`
	// OpenTime represents open_time. 开门时间 格式：HH:mm
	OpenTime *string `json:"open_time,omitzero"`
}
