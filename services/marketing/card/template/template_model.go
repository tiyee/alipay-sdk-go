package template

// McardTemplate 是支付宝官方 V3 OpenAPI 模型。
type McardTemplate struct {
	// CardType 是卡类型
	CardType *string `json:"card_type,omitempty"`
	// GmtCreate 是会员卡模板创建时间
	GmtCreate *string `json:"gmt_create,omitempty"`
	// GmtModified 是会员卡模板修改时间
	GmtModified *string `json:"gmt_modified,omitempty"`
	// TemplateId 是会员卡模板ID
	TemplateId        *string               `json:"template_id,omitempty"`
	TemplateStyleInfo *TemplateStyleInfoDTO `json:"template_style_info,omitempty"`
}

// MoreInfoDTO 是支付宝官方 V3 OpenAPI 模型。
type MoreInfoDTO struct {
	// Descs 是描述，当operate_type为openNative必填
	Descs *[]string `json:"descs,omitempty"`
	// Params 是扩展参数，需要URL地址回带的值，JSON格式(openweb时填)
	Params *string `json:"params,omitempty"`
	// Title 是二级页面标题，只有当operate_type为openNative时有效。
	Title *string `json:"title,omitempty"`
	// Url 是超链接(选择openweb的时候必须填写url参数内容)
	Url *string `json:"url,omitempty"`
}

// PaidOuterCardCycleSellConfDTO 是支付宝官方 V3 OpenAPI 模型。
type PaidOuterCardCycleSellConfDTO struct {
	// CycleSellingUrl 是用户开通连续购买地址
	CycleSellingUrl *string `json:"cycle_selling_url,omitempty"`
	// CycleType 是周期购买类型。
	CycleType *[]string `json:"cycle_type,omitempty"`
	// SupportCycleSell 是是否支持用户连续购买
	SupportCycleSell *bool `json:"support_cycle_sell,omitempty"`
}

// PaidOuterCardManageURLConfDTO 是支付宝官方 V3 OpenAPI 模型。
type PaidOuterCardManageURLConfDTO struct {
	// CycleManageUrl 是连续购买管理地址。包括查看当前用户连续购买详情，关闭连续购买等功能
	CycleManageUrl *string `json:"cycle_manage_url,omitempty"`
	// DowngradeUrl 是付费外卡降级地址
	DowngradeUrl *string `json:"downgrade_url,omitempty"`
	// RefundUrl 是续费外卡退款地址
	RefundUrl *string `json:"refund_url,omitempty"`
	// RenewUrl 是付费外卡续费地址
	RenewUrl *string `json:"renew_url,omitempty"`
	// UpgradeUrl 是付费外卡升级地址
	UpgradeUrl *string `json:"upgrade_url,omitempty"`
}

// PaidOuterCardPriceDetailDTO 是支付宝官方 V3 OpenAPI 模型。
type PaidOuterCardPriceDetailDTO struct {
	// Desc 是方案描述
	Desc *string `json:"desc,omitempty"`
	// Price 是售卖具体金额。单位为元。支持两位小数。
	Price *string `json:"price,omitempty"`
	// PriceType 是售卖金额类型
	PriceType *string `json:"price_type,omitempty"`
	// Worth 是价值。单位为元。支持两位小数
	Worth *string `json:"worth,omitempty"`
}

// PaidOuterCardSellingConfDTO 是支付宝官方 V3 OpenAPI 模型。
type PaidOuterCardSellingConfDTO struct {
	// EndDate 是售卖结束时间。	格式：yyyy-MM-dd HH:mm:ss	如果永久有效则不传。
	EndDate *string `json:"end_date,omitempty"`
	// PriceDetail 是售卖方案列表
	PriceDetail *[]PaidOuterCardPriceDetailDTO `json:"price_detail,omitempty"`
	// SellingUrl 是售卖地址
	SellingUrl *string `json:"selling_url,omitempty"`
	// StartDate 是售卖开始时间。格式：yyyy-MM-dd HH:mm:ss
	StartDate *string `json:"start_date,omitempty"`
}

// PaidOuterCardTemplateConfDTO 是支付宝官方 V3 OpenAPI 模型。
type PaidOuterCardTemplateConfDTO struct {
	CycleSellingConf *PaidOuterCardCycleSellConfDTO `json:"cycle_selling_conf,omitempty"`
	ManageUrlConf    *PaidOuterCardManageURLConfDTO `json:"manage_url_conf,omitempty"`
	OpenSellingConf  *PaidOuterCardSellingConfDTO   `json:"open_selling_conf,omitempty"`
}

// PubChannelDTO 是支付宝官方 V3 OpenAPI 模型。
type PubChannelDTO struct {
	// ExtInfo 是扩展信息，无需配置
	ExtInfo *string `json:"ext_info,omitempty"`
	// PubChannel 是投放渠道类型（选择需要推广投放的渠道）	SHOP_DETAIL：店铺详情页
	PubChannel *string `json:"pub_channel,omitempty"`
}

// TemplateActionInfoDTO 是支付宝官方 V3 OpenAPI 模型。
type TemplateActionInfoDTO struct {
	// Code 是行动点业务CODE，商户自定义
	Code       *string                      `json:"code,omitempty"`
	MiniAppUrl *TemplateActionMiniAppURLDTO `json:"mini_app_url,omitempty"`
	// Text 是行动点展示文案
	Text *string `json:"text,omitempty"`
	// Url 是行动点跳转链接，当url_type填"url"或不填时必填，支持http(s)和支付宝schema地址等
	Url *string `json:"url,omitempty"`
	// UrlType 是跳转链接类型，不填则默认为url类型：		url：对应填写url参数		miniAppUrl: 对应填写mini_app_url参数，跳转至指定的支付宝小程序页面
	UrlType *string `json:"url_type,omitempty"`
}

// TemplateActionMiniAppURLDTO 是支付宝官方 V3 OpenAPI 模型。
type TemplateActionMiniAppURLDTO struct {
	// DisplayOnList 是行动点按钮是否在列表页展示，true 或 false，不填则默认false；		列表页可以配置展示最多2项小程序跳转行动点。
	DisplayOnList *string `json:"display_on_list,omitempty"`
	// MiniAppId 是小程序appId
	MiniAppId *string `json:"mini_app_id,omitempty"`
	// MiniPageParam 是指定小程序页面，不填则默认跳转至对应小程序首页
	MiniPageParam *string `json:"mini_page_param,omitempty"`
	// MiniQueryParam 是query参数，商户自定义传参，填入值将以"query"为参数名提交至指定小程序页面
	MiniQueryParam *string `json:"mini_query_param,omitempty"`
}

// TemplateBenefitInfoDTO 是支付宝官方 V3 OpenAPI 模型。
type TemplateBenefitInfoDTO struct {
	// BenefitDesc 是权益描述信息
	BenefitDesc *[]string `json:"benefit_desc,omitempty"`
	// EndDate 是权益结束时间
	EndDate *string `json:"end_date,omitempty"`
	// StartDate 是权益开始时间
	StartDate *string `json:"start_date,omitempty"`
	// Title 是权益描述
	Title *string `json:"title,omitempty"`
}

// TemplateCardLevelConfDTO 是支付宝官方 V3 OpenAPI 模型。
type TemplateCardLevelConfDTO struct {
	// Level 是会员级别 该级别和开卡接口中的level要一致
	Level *string `json:"level,omitempty"`
	// LevelDesc 是会员级别描述
	LevelDesc *string `json:"level_desc,omitempty"`
	// LevelIcon 是会员级别对应icon图片，		会员信息预留字段，暂不在用户端展示；		通过接口（alipay.offline.material.image.upload）上传图片
	LevelIcon *string `json:"level_icon,omitempty"`
	// LevelShowName 是会员级别显示名称
	LevelShowName *string `json:"level_show_name,omitempty"`
}

// TemplateColumnInfoDTO 是支付宝官方 V3 OpenAPI 模型。
type TemplateColumnInfoDTO struct {
	// Code 是标准栏位：行为由支付宝统一定，同时已经分配标准Code		BALANCE：会员卡余额		POINT：积分		LEVEL：等级		TELEPHONE：联系方式		自定义栏位：行为由商户定义，自定义Code码（只要无重复）
	Code *string `json:"code,omitempty"`
	// GroupTitle 是若template_style_info.column_info_layout	的值为grid，此项为宫格项所属分组标题。可空。如果需要展示该项，还需支付宝内部进行特殊配置。
	GroupTitle *string `json:"group_title,omitempty"`
	// IconId 是当template_style_info.column_info_layout	的值为grid时，此参数必填。此项为宫格项的展示icon。通过接口（alipay.offline.material.image.upload）上传图片。
	IconId   *string      `json:"icon_id,omitempty"`
	MoreInfo *MoreInfoDTO `json:"more_info,omitempty"`
	// OperateType 是1、openNative：打开二级页面，展现 more中descs		2、openWeb：打开URL		3、staticinfo：静态信息		注意：		不填则默认staticinfo；		标准code尽量使用staticinfo，例如TELEPHONE商家电话栏位就只支持staticinfo；
	OperateType *string `json:"operate_type,omitempty"`
	// Tag 是只有当template_style_info.column_info_layout	的值为grid时，此参数有效。此项为宫格项标签，最多只会展示一个标签。
	Tag *string `json:"tag,omitempty"`
	// Title 是栏目标题
	Title *string `json:"title,omitempty"`
	// Value 是卡包详情页面，卡栏位右边展现的值				TELEPHONE栏位的商家联系电话号码由此value字段传入
	Value *string `json:"value,omitempty"`
}

// TemplateFieldRuleDTO 是支付宝官方 V3 OpenAPI 模型。
type TemplateFieldRuleDTO struct {
	// FieldName 是模板字段规则	字段名称(用于定义会员卡开卡接口时卡相应的参数)	余额：Balance	积分：Point	等级：Level	开卡日期：OpenDate	过期日期：ValidDate
	FieldName *string `json:"field_name,omitempty"`
	// RuleName 是规则名		1、ASSIGN_FROM_REQUEST:   以rule_value为key值，表示该栏位的值从会员卡开卡接口中获取，会员卡开卡接口的card_info中获取对应参数值		2、DATE_IN_FUTURE: 生成一个未来的日期（格式YYYY-MM-DD)，当选择DATE_IN_FUTURE的时候，field_name  必须是OpenDate或ValidDate， 值为(10m或10d 分别表示10个月或10天)		3、CONST: 常量，会员卡开卡接口进行开卡的时候使用模板创建时候设置的值，即取rule_value的值
	RuleName *string `json:"rule_name,omitempty"`
	// RuleValue 是规则值，根据rule_name，采取相应取值策略		CONST：直接取rule_value作为卡属性值		DATE_IN_FUTURE：10m或10d 分别表示10个月或10天		ASSIGN_FROM_REQUEST：在开卡Reuqest请求中按rule_value取值，现在和field_name对应的为（OpenDate、ValidDate、Level、Point、Balance）
	RuleValue *string `json:"rule_value,omitempty"`
}

// TemplateFormConfig 是支付宝官方 V3 OpenAPI 模型。
type TemplateFormConfig struct {
	Fields *TemplateFormFields `json:"fields,omitempty"`
	// OpenCardMiniAppId 是入会时承接入会弹层的小程序appId
	OpenCardMiniAppId *string `json:"open_card_mini_app_id,omitempty"`
}

// TemplateFormFields 是支付宝官方 V3 OpenAPI 模型。
type TemplateFormFields struct {
	// Optional 是开卡表单可填字段，可选字段选项不能与必填字段required重复
	Optional *[]string `json:"optional,omitempty"`
	// Required 是开卡表单必填字段，必填字段必须有一个值，并且不能与可选字段optional里面重复
	Required *[]string `json:"required,omitempty"`
}

// TemplateMdcodeNotifyConfDTO 是支付宝官方 V3 OpenAPI 模型。
type TemplateMdcodeNotifyConfDTO struct {
	// ExtParams 是扩展参数信息；		格式为key-value键值对；		支付宝POST请求指定url时，除BizCardNo等固定参数外，将带上ext_params中配置的所有key-value参数。
	ExtParams *string `json:"ext_params,omitempty"`
	// Url 是商户接收发码通知的地址链接；		只支持https地址；		用户打开会员卡时，支付宝提交POST请求此url地址，通知商户发码。
	Url *string `json:"url,omitempty"`
}

// TemplateOpenCardConfDTO 是支付宝官方 V3 OpenAPI 模型。
type TemplateOpenCardConfDTO struct {
	// CardRights 是领卡权益信息
	CardRights *[]TemplateRightsContentDTO `json:"card_rights,omitempty"`
	// Conf 是配置，预留字段，暂时不用
	Conf *string `json:"conf,omitempty"`
	// OpenCardSourceType 是开卡渠道类型	外部系统：ISV （代理运营）	商户直连对接：MER （商家自运营）
	OpenCardSourceType *string `json:"open_card_source_type,omitempty"`
	// OpenCardUrl 是开卡链接，必须http、https开头
	OpenCardUrl *string `json:"open_card_url,omitempty"`
	// SourceAppId 是渠道APPID，提供领卡页面的服务提供方
	SourceAppId *string `json:"source_app_id,omitempty"`
}

// TemplateRightsContentDTO 是支付宝官方 V3 OpenAPI 模型。
type TemplateRightsContentDTO struct {
	// Detail 是权益内容详情
	Detail *string `json:"detail,omitempty"`
	// LogoId 是权益icon图片Id（目前用于开卡授权页使用），通过接口（alipay.offline.material.image.upload）上传图片
	LogoId *string `json:"logo_id,omitempty"`
	// Title 是权益内容标题
	Title *string `json:"title,omitempty"`
}

// TemplateStyleInfoDTO 是支付宝官方 V3 OpenAPI 模型。
type TemplateStyleInfoDTO struct {
	// BackgroundId 是背景图片Id，通过接口（alipay.offline.material.image.upload）上传图片				图片说明：2M以内，格式：bmp、png、jpeg、jpg、gif；		尺寸不小于1020*643px；		图片不得有圆角，不得拉伸变形
	BackgroundId *string `json:"background_id,omitempty"`
	// BannerImgId 是banner图片id。 通过接口（alipay.offline.material.image.upload）上传图片。
	BannerImgId *string `json:"banner_img_id,omitempty"`
	// BannerUrl 是banner跳转地址。
	BannerUrl *string `json:"banner_url,omitempty"`
	// BgColor 是字体颜色（非背景色），只影响卡详情中部信息区域字体颜色
	BgColor *string `json:"bg_color,omitempty"`
	// BrandName 是品牌商名称
	BrandName *string `json:"brand_name,omitempty"`
	// CardShowName 是钱包端显示名称
	CardShowName *string `json:"card_show_name,omitempty"`
	// Color 是注意：此字段已废弃。
	Color *string `json:"color,omitempty"`
	// ColumnInfoLayout 是如果为空则默认为list。
	ColumnInfoLayout *string `json:"column_info_layout,omitempty"`
	// FeatureDescriptions 是特色信息，用于领卡预览
	FeatureDescriptions *[]string `json:"feature_descriptions,omitempty"`
	// FrontImageEnable 是设置是否在卡面展示（个人头像）图片信息，默认不展示；		当前仅用于身份验证信息类型的个人头像图片；		图片id随创建卡/更新卡时传入；		详见会员卡产品文档。
	FrontImageEnable *bool `json:"front_image_enable,omitempty"`
	// FrontTextListEnable 是设置是否在卡面展示文案信息，默认不展示；		文案信息分行展示，最多展示4行文案，每行文案分为label和value两部分；		文案实际内容随创建卡/更新卡时传入；		详见会员卡产品说明文档。
	FrontTextListEnable *bool `json:"front_text_list_enable,omitempty"`
	// LogoId 是logo的图片ID，通过接口（alipay.offline.material.image.upload）上传图片				图片说明：1M以内，格式bmp、png、jpeg、jpg、gif；		尺寸不小于500*500px的正方形；		请优先使用商家LOGO；
	LogoId *string `json:"logo_id,omitempty"`
	// Slogan 是标语
	Slogan *string `json:"slogan,omitempty"`
	// SloganImgId 是标语图片id， 通过接口（alipay.offline.material.image.upload）上传图片
	SloganImgId *string `json:"slogan_img_id,omitempty"`
}
