package order

// ServiceMarketOrderCreate 是支付宝官方 V3 OpenAPI 模型。
type ServiceMarketOrderCreate struct {
	// AppCategoryIds 是11_12;12_13。小程序类目，格式为 第一个一级类目_第一个二级类目;第二个一级类目_第二个二级类目，详细类目可以参考https://docs.alipay.com/isv/10325。
	AppCategoryIds *string `json:"app_category_ids,omitempty"`
	// AppDesc 是小程序官方示例Demo，展示已支持的接口能力及组件。。
	AppDesc *string `json:"app_desc,omitempty"`
	// AppEnglishName 是小程序应用英文名称。
	AppEnglishName *string `json:"app_english_name,omitempty"`
	// AppName 是小程序应用名称。
	AppName *string `json:"app_name,omitempty"`
	// AppOrigin 是来源的业务方，需要申请。
	AppOrigin *string `json:"app_origin,omitempty"`
	// AppSlogan 是小程序应用简介，一句话描述小程序功能。
	AppSlogan *string `json:"app_slogan,omitempty"`
	// MarketCode 是订购的服务商品ID所在的市场编码。新接入场景必须传递，具体值请联系产品分配。。
	MarketCode *string `json:"market_code,omitempty"`
	// MerchandiseId 是订购的服务商品ID。
	MerchandiseId *string `json:"merchandise_id,omitempty"`
	// MerchantPid 是商户PID。
	MerchantPid *string `json:"merchant_pid,omitempty"`
	// MiniAppId 是一二方支持传入appId。
	MiniAppId *string `json:"mini_app_id,omitempty"`
	// MiniCategoryIds 是新小程序前台类目，格式为 第一个一级类目_第一个二级类目;第二个一级类目_第二个二级类目_第二个三级类目，详细类目可以通过 https://docs.open.alipay.com/api_49/alipay.open.mini.category.query接口查询mini_category_list，如果不填默认采用当前小程序应用类目。使用默认应用类目后不需要再次上传营业执照号、营业执照名、营业执照截图、营业执照有效期。使用后不再读取app_category_ids值，老前台类目将废弃。
	MiniCategoryIds *string `json:"mini_category_ids,omitempty"`
	// OutBizNo 是比如某种业务标准外部订单号,比如交易外部订单号，代表商户端自己订单号。
	OutBizNo *string `json:"out_biz_no,omitempty"`
	// ServiceEmail 是小程序客服邮箱。
	ServiceEmail *string `json:"service_email,omitempty"`
	// ServicePhone 是小程序客服电话。
	ServicePhone *string `json:"service_phone,omitempty"`
}

// OrderItem 是支付宝官方 V3 OpenAPI 模型。
type OrderItem struct {
	// Address 是店铺所在具体位置。
	Address *string `json:"address,omitempty"`
	// BrandName 是品牌名称。
	BrandName *string `json:"brand_name,omitempty"`
	// Category 是店铺品类。
	Category *string `json:"category,omitempty"`
	// City 是店铺所在的市。
	City *string `json:"city,omitempty"`
	// CityCode 是城市编码。
	CityCode *string `json:"city_code,omitempty"`
	// CommodityId 是订购的服务商品ID。
	CommodityId *string `json:"commodity_id,omitempty"`
	// Contacts 是订单联系人，在下单是传入则会返回。
	Contacts *string `json:"contacts,omitempty"`
	// Creator 是门店创建人(已删除)。
	Creator *string `json:"creator,omitempty"`
	// ExpireDate 是过期时间。
	ExpireDate *string `json:"expire_date,omitempty"`
	// MerchantName 是商户名称。
	MerchantName *string `json:"merchant_name,omitempty"`
	// MerchantPid 是商户PID。
	MerchantPid *string `json:"merchant_pid,omitempty"`
	// MiniAppId 是小程序APPID，订购实体为小程序的场景该字段必填。
	MiniAppId *string `json:"mini_app_id,omitempty"`
	// MiniAppName 是小程序应用名称，订购实体为小程序的场景该字段必填。
	MiniAppName *string `json:"mini_app_name,omitempty"`
	// OnlineTime 是上架时间。
	OnlineTime *string `json:"online_time,omitempty"`
	// OpenIdModel 是应用用户标识模式。
	OpenIdModel *string `json:"open_id_model,omitempty"`
	// OrderStatus 是TO_DO-未实施,DOING-实施中,TO_CONFIRM-待商户确认,CONFIRMED-商户已确认,DONE-已完成,MERCHANT_REJECTED-商户已回绝,MERCHANT_CANCELLED-商户已取消,ISV_REJECTED-服务商已回绝,ISV_CANCELLED-服务商已取消。
	OrderStatus *string `json:"order_status,omitempty"`
	// PhoneNo 是订单所属人联系方式（手机或者座机）。
	PhoneNo *string `json:"phone_no,omitempty"`
	// Province 是店铺所在的省份。
	Province *string `json:"province,omitempty"`
	// ServiceEffectDate 是订购的服务有效期生效时间。
	ServiceEffectDate *string `json:"service_effect_date,omitempty"`
	// ServiceExpireDate 是服务有效期截止时间。
	ServiceExpireDate *string `json:"service_expire_date,omitempty"`
	// ShopId 是店铺ID，订购实体为口碑门店的场景该字段必填。
	ShopId *string `json:"shop_id,omitempty"`
	// ShopName 是店铺名称，订购实体为口碑门店的场景该字段必填。
	ShopName *string `json:"shop_name,omitempty"`
	// ShopStatus 是店铺状态（ONLINE--已上架 OFFLINE--未上架 AVAILABLE--已开通 INIT--未开通 EXPIRED--已过期）。
	ShopStatus *string `json:"shop_status,omitempty"`
	// Status 是待服务商接单。
	Status *string `json:"status,omitempty"`
}
