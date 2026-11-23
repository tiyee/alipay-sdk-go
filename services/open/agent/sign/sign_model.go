package sign

// AgentCommonSign 是支付宝官方 V3 OpenAPI 模型。
type AgentCommonSign struct {
	// AlipayLifeName 是支付宝生活号(原服务窗)名称（1 app_name、app_demo；2 web_sites；3 alipay_life_name。1、2、3少选择一个填写）。
	AlipayLifeName *string `json:"alipay_life_name,omitempty"`
	// AppMarket 是投放应用市场。
	AppMarket *string `json:"app_market,omitempty"`
	// AppName 是商户的APP应用名称（1 app_name、app_demo；2 web_sites；3 alipay_life_name；4 wechat_official_account_name。1、2、3、4至少选择一个填写）。
	AppName *string `json:"app_name,omitempty"`
	// AppStatus 是应用状态。
	AppStatus *string `json:"app_status,omitempty"`
	// AppTestAccount 是可以登录该应用的测试账户。
	AppTestAccount *string `json:"app_test_account,omitempty"`
	// AppTestAccountPassword 是可以登录此应用的账户的密码。对app_test_account的登录密码。。
	AppTestAccountPassword *string `json:"app_test_account_password,omitempty"`
	// BatchNo 是代商户操作事务编号，通过alipay.open.agent.create接口进行创建。。
	BatchNo *string `json:"batch_no,omitempty"`
	// BusinessLicenseNo 是营业执照号码。
	BusinessLicenseNo *string `json:"business_license_no,omitempty"`
	// DateLimitation 是营业期限。
	DateLimitation *string `json:"date_limitation,omitempty"`
	// DownloadLink 是应用下载链接。
	DownloadLink *string `json:"download_link,omitempty"`
	// H5Memo 是网站扩展信息文本。
	H5Memo *string `json:"h_5_memo,omitempty"`
	// H5Sites 是H5网站地址。
	H5Sites *string `json:"h_5_sites,omitempty"`
	// H5Status 是H5网站状态。
	H5Status *string `json:"h_5_status,omitempty"`
	// LifeAppId 是请输入被签约商户的支付宝生活号ID编号。
	LifeAppId *string `json:"life_app_id,omitempty"`
	// LongTerm 是营业期限是否长期有效。
	LongTerm *bool `json:"long_term,omitempty"`
	// MccCode 是<a href="https://opendocs.alipay.com/common/02khjv">商家经营类目2.0</a> 中的“经营类目编码”。
	MccCode *string `json:"mcc_code,omitempty"`
	// MiniAppId 是请输入被签约商家下任意小程序APPID ID编号。
	MiniAppId *string `json:"mini_app_id,omitempty"`
	// MiniAppName 是小程序名称。
	MiniAppName *string `json:"mini_app_name,omitempty"`
	// MobileType 是移动设备类型，当开通电脑网站支付产品时不需赋值；当开通手机网站支付时，传H5；当开通JSAPI时，传APP；当开通预授权支付产品，且交易场景选择“mobile”(移动设备)时，需要从APP,H5中至少选择一个，可多选，字符串英文逗号分隔。。
	MobileType *string `json:"mobile_type,omitempty"`
	// ProductCode 是isv要代商户签约产品码，产品码是支付宝内部对产品的唯一标识。
	ProductCode *string `json:"product_code,omitempty"`
	// Rate 是服务费率（%），例如：请输入0.38~0.6 之间（小数点后两位，可取0.38%及0.6%），直播买卖通涉及JSAPI支付产品，交易手续费为 0.38%-0.6%，后续费率变化可参考支付宝标准服务费。 当签约且授权标识 sign_and_auth=true 时，该费率信息必填。。
	Rate *string `json:"rate,omitempty"`
	// ShopCityCode 是门店所在城市。请按照<a href="https://mdn.alipayobjects.com/huamei_sm7gq8/afts/file/A*blT9RqSR9_gAAAAAAAAAAAAADuKQAQ/2022%E8%9A%82%E8%9A%81%E9%87%91%E6%9C%8D%E5%9C%B0%E5%8C%BA%E7%A0%81.xlsx">地区码表格</a> 中内容填写。。
	ShopCityCode *string `json:"shop_city_code,omitempty"`
	// ShopDetailAddress 是门店详细地址。
	ShopDetailAddress *string `json:"shop_detail_address,omitempty"`
	// ShopDistrictCode 是门店所在区/县。请按照<a href="https://mdn.alipayobjects.com/huamei_sm7gq8/afts/file/A*blT9RqSR9_gAAAAAAAAAAAAADuKQAQ/2022%E8%9A%82%E8%9A%81%E9%87%91%E6%9C%8D%E5%9C%B0%E5%8C%BA%E7%A0%81.xlsx">地区码表格</a> 中内容填写。。
	ShopDistrictCode *string `json:"shop_district_code,omitempty"`
	// ShopName 是门店名称。
	ShopName *string `json:"shop_name,omitempty"`
	// ShopProvinceCode 是门店所在省份。请按照<a href="https://mdn.alipayobjects.com/huamei_sm7gq8/afts/file/A*blT9RqSR9_gAAAAAAAAAAAAADuKQAQ/2022%E8%9A%82%E8%9A%81%E9%87%91%E6%9C%8D%E5%9C%B0%E5%8C%BA%E7%A0%81.xlsx">地区码表格</a> 中内容填写。。
	ShopProvinceCode *string `json:"shop_province_code,omitempty"`
	// SpecialLicenseType 是合约特殊资质的类型，不传默认99。
	SpecialLicenseType *string `json:"special_license_type,omitempty"`
	// TradeScene 是交易场景，当开通产品为电脑网站支付时，传pc；当开通产品为手机网站支付时，传mobile；当开通产品为JSAPI时，传tiny_app；当开通产品为预授权支付时，下面括号中三个至少传一个（tiny_app,mobile,offline_store），可多选，字符串用英文逗号分隔。
	TradeScene *string `json:"trade_scene,omitempty"`
	// WebSites 是接入网址信息（1 app_name、app_demo；2 web_sites；3 alipay_life_name；4 wechat_official_account_name。1、2、3、4至少选择一个填写）。
	WebSites *[]string `json:"web_sites,omitempty"`
	// WebStatus 是网站状态，枚举值为：ONLINE，OFFLINE。
	WebStatus *string `json:"web_status,omitempty"`
	// WebTestAccount 是可以登录此网站的测试账户。。
	WebTestAccount *string `json:"web_test_account,omitempty"`
	// WebTestAccountPassword 是可以登录此网站的账户的密码。对应web_test_account的登录密码。
	WebTestAccountPassword *string `json:"web_test_account_password,omitempty"`
	// WechatOfficialAccountName 是1 app_name、app_demo；2 web_sites；3 alipay_life_name；4 wechat_official_account_name。1、2、3、4至少选择一个填写(该字段已经不在使用，废弃，不用填写)。
	WechatOfficialAccountName *string `json:"wechat_official_account_name,omitempty"`
}

// AgentFaceToFaceSign 是支付宝官方 V3 OpenAPI 模型。
type AgentFaceToFaceSign struct {
	// BatchNo 是代商户操作事务编号，通过<a href="https://opendocs.alipay.com/apis/api_50/alipay.open.agent.create">接口地址</a> (开启代商户签约、创建应用事务)接口进行事务创建后获取。。
	BatchNo *string `json:"batch_no,omitempty"`
	// BusinessLicenseMobile 是上传非经营者本人（签约账号）的营业执照时必填，填写营业执照法人支付宝账号对应的手机号。。
	BusinessLicenseMobile *string `json:"business_license_mobile,omitempty"`
	// BusinessLicenseNo 是营业执照号码。
	BusinessLicenseNo *string `json:"business_license_no,omitempty"`
	// DateLimitation 是营业期限。
	DateLimitation *string `json:"date_limitation,omitempty"`
	// LongTerm 是营业期限是否长期有效。
	LongTerm *bool `json:"long_term,omitempty"`
	// MccCode 是商家经营类目编码。参见 <a href="https://opendocs.alipay.com/open/01n22g#%E5%95%86%E5%AE%B6%E7%BB%8F%E8%90%A5%E7%B1%BB%E7%9B%AE">商家经营类目2.0</a> 中的“一级类目code_二级类目code”。。
	MccCode *string `json:"mcc_code,omitempty"`
	// Rate 是服务费率（%），0.38~0.6 之间（小数点后两位，可取0.38%及0.6%）。 当签约且授权标识 sign_and_auth=true 时，该费率信息必填。。
	Rate        *string          `json:"rate,omitempty"`
	ShopAddress *SignAddressInfo `json:"shop_address,omitempty"`
	// ShopName 是店铺名称。
	ShopName *string `json:"shop_name,omitempty"`
	// SignAndAuth 是签约且授权标识，默认为false，只进行签约操作； 如果设置为true，则表示签约成功后，会自动进行应用授权操作。。
	SignAndAuth *bool `json:"sign_and_auth,omitempty"`
}

// AgentMobilePaySign 是支付宝官方 V3 OpenAPI 模型。
type AgentMobilePaySign struct {
	// AppMarket 是应用在哪些市场上架，枚举值为：苹果,应用宝,华为,360,小米,豌豆荚,其他。
	AppMarket *[]string `json:"app_market,omitempty"`
	// AppName 是商户的APP应用名称。
	AppName *string `json:"app_name,omitempty"`
	// AppStatus 是应用上架状态，枚举值为 已上线，未上线。
	AppStatus *string `json:"app_status,omitempty"`
	// AppTestAccount 是可以登录该应用的测试账户。
	AppTestAccount *string `json:"app_test_account,omitempty"`
	// AppTestAccountPassword 是可以登录此应用的账户的密码。对应app_test_account的登录密码。
	AppTestAccountPassword *string `json:"app_test_account_password,omitempty"`
	// AppType 是应用类型，枚举值为：IOS，安卓。
	AppType *[]string `json:"app_type,omitempty"`
	// BatchNo 是代商户操作事务编号，通过alipay.open.isv.agent.create接口进行创建。。
	BatchNo *string `json:"batch_no,omitempty"`
	// BusinessLicenseMobile 是营业执照法人手机号码。
	BusinessLicenseMobile *string `json:"business_license_mobile,omitempty"`
	// BusinessLicenseNo 是营业执照号码。。
	BusinessLicenseNo *string `json:"business_license_no,omitempty"`
	// DateLimitation 是营业期限。
	DateLimitation *string `json:"date_limitation,omitempty"`
	// DownloadLink 是应用下载链接。
	DownloadLink *string `json:"download_link,omitempty"`
	// LongTerm 是营业期限是否长期有效。
	LongTerm *bool `json:"long_term,omitempty"`
	// MccCode 是所属MCCCode，详情可参考 <a href="https://opendocs.alipay.com/open/01n22g#%E5%95%86%E5%AE%B6%E7%BB%8F%E8%90%A5%E7%B1%BB%E7%9B%AE">商家经营类目</a> 中的“经营类目编码”。
	MccCode *string `json:"mcc_code,omitempty"`
	// MobileType 是传参：APP，代表设备类型是APP。
	MobileType *string `json:"mobile_type,omitempty"`
	// TradeScene 是传参：mobile，代表交易场景是移动设备。
	TradeScene *string `json:"trade_scene,omitempty"`
}

// AgentOfflinePaymentSign 是支付宝官方 V3 OpenAPI 模型。
type AgentOfflinePaymentSign struct {
	// BatchNo 是代商户操作事务编号，通过alipay.open.agent.create接口进行创建。。
	BatchNo *string `json:"batch_no,omitempty"`
	// BusinessLicenseMobile 是被邀请授权的营业执照法人手机号码，上传非同人营业执照时必填。
	BusinessLicenseMobile *string `json:"business_license_mobile,omitempty"`
	// BusinessLicenseNo 是营业执照号码。若填写，请与以下营业执照图片、期限、一起提供。。
	BusinessLicenseNo *string `json:"business_license_no,omitempty"`
	// DateLimitation 是营业期限。
	DateLimitation  *string          `json:"date_limitation,omitempty"`
	DeliveryAddress *SignAddressInfo `json:"delivery_address,omitempty"`
	// LongTerm 是营业期限是否长期有效。
	LongTerm *bool `json:"long_term,omitempty"`
	// MccCode 是所属MCCCode，可参考 <a href="https://opendocs.alipay.com/open/01n22g#%E5%95%86%E5%AE%B6%E7%BB%8F%E8%90%A5%E7%B1%BB%E7%9B%AE">商家经营类目</a> 中的“经营类目编码”。
	MccCode *string `json:"mcc_code,omitempty"`
	// Rate 是服务费率（%），0.38~3之间，精确到0.01。
	Rate        *string          `json:"rate,omitempty"`
	ShopAddress *SignAddressInfo `json:"shop_address,omitempty"`
	// ShopName 是店铺名称。
	ShopName *string `json:"shop_name,omitempty"`
}

// ProductSignStatusInfo 是支付宝官方 V3 OpenAPI 模型。
type ProductSignStatusInfo struct {
	// ProductCode 是产品编码。
	ProductCode *string `json:"product_code,omitempty"`
	// ProductName 是产品名称。
	ProductName *string `json:"product_name,omitempty"`
	// Status 是签约状态。
	Status *string `json:"status,omitempty"`
}

// SignAddressInfo 是支付宝官方 V3 OpenAPI 模型。
type SignAddressInfo struct {
	// CityCode 是城市编码。请按照<a href="https://mdn.alipayobjects.com/huamei_sm7gq8/afts/file/A*blT9RqSR9_gAAAAAAAAAAAAADuKQAQ/2022%E8%9A%82%E8%9A%81%E9%87%91%E6%9C%8D%E5%9C%B0%E5%8C%BA%E7%A0%81.xlsx">地区码表格</a> 中内容填写。 （参考资料： <a href="https://www.mca.gov.cn/n156/n186/index.html">参考资料</a>）。
	CityCode *string `json:"city_code,omitempty"`
	// CountryCode 是国家编码，中国默认：156。
	CountryCode *string `json:"country_code,omitempty"`
	// DetailAddress 是详细地址。
	DetailAddress *string `json:"detail_address,omitempty"`
	// DistrictCode 是区县编码。请按照<a href="https://mdn.alipayobjects.com/huamei_sm7gq8/afts/file/A*blT9RqSR9_gAAAAAAAAAAAAADuKQAQ/2022%E8%9A%82%E8%9A%81%E9%87%91%E6%9C%8D%E5%9C%B0%E5%8C%BA%E7%A0%81.xlsx">地区码表格</a> 中内容填写。 （参考资料： <a href="https://www.mca.gov.cn/n156/n186/index.html">参考资料</a>）。
	DistrictCode *string `json:"district_code,omitempty"`
	// Latitude 是纬度，浮点型,小数点后最多保留6位 如需要录入经纬度，请以高德坐标系为准，录入时请确保经纬度参数准确。高德经纬度查询：<a href="http://lbs.amap.com/console/show/picker">查询地址</a>。
	Latitude *string `json:"latitude,omitempty"`
	// Longitude 是经度，浮点型, 小数点后最多保留6位。 如需要录入经纬度，请以高德坐标系为准，录入时请确保经纬度参数准确。高德经纬度查询：<a href="http://lbs.amap.com/console/show/picker">查询地址</a>。
	Longitude *string `json:"longitude,omitempty"`
	// ProvinceCode 是省份编码。请按照<a href="https://mdn.alipayobjects.com/huamei_sm7gq8/afts/file/A*blT9RqSR9_gAAAAAAAAAAAAADuKQAQ/2022%E8%9A%82%E8%9A%81%E9%87%91%E6%9C%8D%E5%9C%B0%E5%8C%BA%E7%A0%81.xlsx">地区码表格</a> 中内容填写。 （参考资料： <a href="https://www.mca.gov.cn/n156/n186/index.html">参考资料</a>）。
	ProvinceCode *string `json:"province_code,omitempty"`
}
