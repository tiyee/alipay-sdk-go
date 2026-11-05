package audit

// MiniInnerVersionAuditCancelRequest 是支付宝官方 V3 OpenAPI 模型。
type MiniInnerVersionAuditCancelRequest struct {
	// AppOrigin 是业务来源。
	AppOrigin *string `json:"app_origin,omitempty"`
	// AppVersion 是小程序版本。
	AppVersion *string `json:"app_version,omitempty"`
	// BundleId 是端信息。
	BundleId *string `json:"bundle_id,omitempty"`
	// MiniAppId 是小程序ID，仅特殊场景使用，普通业务方无需关注该参数。
	MiniAppId *string `json:"mini_app_id,omitempty"`
	// Pid 是小程序所属PID。
	Pid *string `json:"pid,omitempty"`
}

// MiniInnerVersionAuditSubmitRequest 是支付宝官方 V3 OpenAPI 模型。
type MiniInnerVersionAuditSubmitRequest struct {
	// AppCategoryIds 是小程序类目，格式为 第一个一级类目_第一个二级类目;第二个一级类目_第二个二级类目，详细类目可以参考<a href=‘https://opendocs.alipay.com/b/03al2m’>开放服务类目</a>，如果不填默认采用当前小程序应用类目。使用默认应用类目后不需要再次上传营业执照号、营业执照名、营业执照截图、营业执照有效期。。
	AppCategoryIds *string `json:"app_category_ids,omitempty"`
	// AppDesc 是小程序应用描述，20-200个字，如果不填默认采用当前小程序的应用描述。
	AppDesc *string `json:"app_desc,omitempty"`
	// AppEnglishName 是小程序应用英文名称，如果不填默认采用当前小程序应用英文名称，3～30个字符。插件不填。
	AppEnglishName *string `json:"app_english_name,omitempty"`
	// AppLogo 是小程序logo图标，图片格式必须为：png、jpeg、jpg，比例必须为1:1，建议上传像素为180*180，不超过256kb， 如果不填默认采用当前小程序应用logo图标。
	AppLogo *string `json:"app_logo,omitempty"`
	// AppName 是小程序应用名称，如果不填默认采用当前小程序应用名称，名称3-40个字符，一个中文算两个字符。
	AppName *string `json:"app_name,omitempty"`
	// AppOrigin 是来源类型，新接入方需要向支付宝申请专用来源，否则不予接入，申请方式请参见接入手册。。
	AppOrigin *string `json:"app_origin,omitempty"`
	// AppSlogan 是小程序应用简介，一句话描述小程序功能，如果不填默认采用当前小程序应用简介，10~32个字符。
	AppSlogan *string `json:"app_slogan,omitempty"`
	// AppVersion 是需要提交审核的小程序版本号，格式为: x.y.z，其中x、y、z均为整型数字，一个应用最多只能有一个审核中、审核通过或者审核驳回的版本。。
	AppVersion *string `json:"app_version,omitempty"`
	// BundleId 是端ID，多端场景下区分不同端。
	BundleId    *string           `json:"bundle_id,omitempty"`
	LicenseInfo *AuditLicenseInfo `json:"license_info,omitempty"`
	// Memo 是小程序备注。
	Memo *string `json:"memo,omitempty"`
	// MiniAppId 是小程序ID，特殊场景专用，普通业务方无需关注该参数。。
	MiniAppId *string `json:"mini_app_id,omitempty"`
	// MiniCategoryIds 是新小程序前台类目，格式为 第一个一级类目_第一个二级类目;第二个一级类目_第二个二级类目_第二个三级类目，详细类目可以通过 <a href='https://opendocs.alipay.com/mini/03l8c6'>alipay.open.mini.category.query</a>（小程序类目树查询接口）查询mini_category_list，如果不填默认采用当前小程序应用类目。使用默认应用类目后不需要再次上传营业执照号、营业执照名、营业执照截图、营业执照有效期。使用后不再读取app_category_ids值，老前台类目将废弃。
	MiniCategoryIds *string `json:"mini_category_ids,omitempty"`
	// OpenId 是小程序所属PID。
	OpenId *string `json:"open_id,omitempty"`
	// OutDoorPic 是门头照图片地址，部分小程序类目需要提交，参照 <a href=‘https://opendocs.alipay.com/b/03al2m’>开放服务类目</a> 中是否需要营业执照信息，如果不填默认采用当前小程序门头照图片。
	OutDoorPic *string `json:"out_door_pic,omitempty"`
	// Pid 是小程序所属PID。
	Pid *string `json:"pid,omitempty"`
	// ScreenShotList 是小程序应用截图列表，逗号分割，需要2-5张图片，单张图片不能超过4MB，图片格式只支持jpg，png；通过模板实例化小程序不需要传递此参数。。
	ScreenShotList *[]string `json:"screen_shot_list,omitempty"`
	// ServiceEmail 是小程序客服邮箱，如果不填默认采用当前小程序的应用客服邮箱，和客服电话至少填写一个。
	ServiceEmail *string `json:"service_email,omitempty"`
	// ServicePhone 是小程序客服电话，如果不填默认采用当前小程序的应用客服电话，和客服邮箱至少填写一个。插件不填。
	ServicePhone *string `json:"service_phone,omitempty"`
	// SpecialLicensePicList 是特殊资质图片地址列表，逗号分隔；部分类目需要特殊资质，如果需要特殊资质，最少一张，最多三张。模板和插件不需要特殊资质。
	SpecialLicensePicList *string `json:"special_license_pic_list,omitempty"`
	// VersionDesc 是小程序版本变更描述，30-500个字符，区分于app_desc。
	VersionDesc *string `json:"version_desc,omitempty"`
}
