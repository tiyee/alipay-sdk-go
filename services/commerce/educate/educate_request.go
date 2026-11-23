// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package educate

// EducateCampusInstitutionsAddRequest 是支付宝官方 V3 OpenAPI 模型。
type EducateCampusInstitutionsAddRequest struct {
	// CardPictUrl 是事业单位法人证书或办学许可证的图片。入驻学校时，如果流入人工审核，会根据此图片进行辅助验证。
	CardPictUrl *string `json:"card_pict_url,omitempty"`
	// CityCode 是市
	CityCode *string `json:"city_code,omitempty"`
	// InstName 是入驻的学校名称，必须是完整的学校全称
	InstName *string `json:"inst_name,omitempty"`
	// InstStdCode 是学校外标，使用统一社会信用编码
	InstStdCode *string `json:"inst_std_code,omitempty"`
	// LearningStage 是办学阶段。	枚举值如下：	KINDERGARTEN（幼儿园）	PRIMARY_SCHOOL（小学） 	MIDDLE_SCHOOL（初中） 	HIGH_SCHOOL（高中）	SECONDARY_VOCATIONAL_SCHOOL（中职中专） 	注意：如果学校兼有多种属性，可以连写用英文逗号拆分，如：MIDDLE_SCHOOL,HIGH_SCHOOL 代表兼有初中部和高中部；
	LearningStage *string `json:"learning_stage,omitempty"`
	// ProvinceCode 是省份
	ProvinceCode *string `json:"province_code,omitempty"`
	// SchoolProperty 是学校性质.枚举值如下：	 1：公立 	2：民办
	SchoolProperty *string `json:"school_property,omitempty"`
}

// EducateSchoolCardOrderSyncRequest 是支付宝官方 V3 OpenAPI 模型。
type EducateSchoolCardOrderSyncRequest struct {
	// ActualAmount 是实际金额（总支付金额），单位为【元】
	ActualAmount *string `json:"actual_amount,omitempty"`
	// AppletAppId 是小程序appid
	AppletAppId *string `json:"applet_app_id,omitempty"`
	// CardBalance 是校园卡余额
	CardBalance *string `json:"card_balance,omitempty"`
	// CardNo 是128
	CardNo *string `json:"card_no,omitempty"`
	// CardType 是校园卡类型
	CardType *string `json:"card_type,omitempty"`
	// CreateTime 是该笔订单真实的创建时间，需精确到毫秒。
	CreateTime *string `json:"create_time,omitempty"`
	// DiscountAmount 是优惠金额
	DiscountAmount *string     `json:"discount_amount,omitempty"`
	GoodsOrders    *GoodsOrder `json:"goods_orders,omitempty"`
	// MerchantName 是商家名称，不传默认展示学校名称
	MerchantName *string `json:"merchant_name,omitempty"`
	// ModifiedTime 是订单修改时间
	ModifiedTime *string `json:"modified_time,omitempty"`
	// OpenId 是用户open_id
	OpenId *string `json:"open_id,omitempty"`
	// OrderAmount 是订单金额
	OrderAmount *string `json:"order_amount,omitempty"`
	// OrderDetailUrl 是订单详情链接
	OrderDetailUrl *string `json:"order_detail_url,omitempty"`
	// OrderStatus 是订单状态
	OrderStatus *string `json:"order_status,omitempty"`
	// OutBizNo 是外部业务号，由商家自定义，128个字符以内，仅支持字母、数字、下划线且需保证在商户端不重复。
	OutBizNo *string `json:"out_biz_no,omitempty"`
	// PayAddress 是支付地点
	PayAddress *string `json:"pay_address,omitempty"`
	// PayMode 是付款方式，不传默认展示学校名称+校园卡+（卡号后四位）
	PayMode *string `json:"pay_mode,omitempty"`
	// RakeBackPid 是系统商编号。该参数作为系统商返佣数据提取的依据，请填写系统商签约协议的PID
	RakeBackPid *string `json:"rake_back_pid,omitempty"`
	// SchoolId 是学校内标，录入学校接口返回的参数
	SchoolId *string `json:"school_id,omitempty"`
	// SchoolPid 是学校收款账号
	SchoolPid *string `json:"school_pid,omitempty"`
	// ServiceProviderName 是服务提供者名称
	ServiceProviderName *string `json:"service_provider_name,omitempty"`
	// Type 是业务类型
	Type *string `json:"type,omitempty"`
	// UserId 是买家支付宝用户ID。	2088开头的16位纯数字，小程序场景下获取用户ID请参考：用户授权;	其它场景下获取用户ID请参考：网页授权获取用户信息;
	UserId *string `json:"user_id,omitempty"`
}

// EducateCampusInstitutionsQueryRequest is the query request model for alipay.commerce.educate.campus.institutions.query.
type EducateCampusInstitutionsQueryRequest struct {
	// InstName 是学校名称
	InstName *string `json:"inst_name,omitempty" query:"inst_name,omitempty"`
	// ProvinceCode 是省份编码，编码标准为中华人民共和国行政区划代码，参见 2020年12月中华人民共和国县以上行政区划代码。
	ProvinceCode *string `json:"province_code,omitempty" query:"province_code,omitempty"`
	// CityCode 是城市编码，编码标准为中华人民共和国行政区划代码，参见 2020年12月中华人民共和国县以上行政区划代码。
	CityCode *string `json:"city_code,omitempty" query:"city_code,omitempty"`
	// InstId 是学校内标：支付宝内部学校唯一编号。使用内标查询时，默认使用内标进行精确匹配。
	InstId *string `json:"inst_id,omitempty" query:"inst_id,omitempty"`
	// InstStdCode 是学校外标：统一社会信用编码或教育部提供的学校标识码。使用学校外标查询时，默认使用外标进行精确匹配。
	InstStdCode *string `json:"inst_std_code,omitempty" query:"inst_std_code,omitempty"`
	// LikeProperty 是是否使用学校名称模糊匹配进行查询，默认精确匹配。	枚举值如下：	1：精确匹配；	0：模糊匹配
	LikeProperty *string `json:"like_property,omitempty" query:"like_property,omitempty"`
}
