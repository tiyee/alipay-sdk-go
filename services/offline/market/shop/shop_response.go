package shop

// OfflineMarketShopBatchQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type OfflineMarketShopBatchQueryResponse struct {
	// CurrentPageno 是当前页码
	CurrentPageno *string `json:"current_pageno,omitempty"`
	// ShopIds 是门店列表ID，逗号分隔
	ShopIds *[]string `json:"shop_ids,omitempty"`
	// TotalPageno 是总页码数目
	TotalPageno *string `json:"total_pageno,omitempty"`
}

// OfflineMarketShopCategoryQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type OfflineMarketShopCategoryQueryResponse struct {
	// ShopCategoryConfigInfos 是门店类目配置信息，包括能够开店的叶子节点类目信息，以及类目约束配置信息。
	ShopCategoryConfigInfos *[]ShopCategoryConfigInfo `json:"shop_category_config_infos,omitempty"`
}

// OfflineMarketShopCreateResponse 是支付宝官方 V3 OpenAPI 模型。
type OfflineMarketShopCreateResponse struct {
	// ApplyId 是开店请求受理成功后返回的支付宝流水ID，根据此ID调用接口 	alipay.offline.market.applyorder.batchquery，能够获取当前开店请求审核状态、驳回原因等信息。
	ApplyId *string `json:"apply_id,omitempty"`
	// AuditDesc 是废弃字段。
	AuditDesc *string `json:"audit_desc,omitempty"`
	// AuditStatus 是同步请求如果支付宝受理成功，将返回AUDITING状态。
	AuditStatus *string `json:"audit_status,omitempty"`
	// IsOnline 是废弃字段，T表示上架,F表示未上架。
	IsOnline *string `json:"is_online,omitempty"`
	// IsShow 是废弃字段，T表示显示，F表示隐藏。
	IsShow *string `json:"is_show,omitempty"`
	// Rate 是门店费率值，指定的ISV可用。创建门店时，为符合情况的门店，进行费率单独设置。若费率不符合门店类目的设置，则会创建门店失败。
	Rate *string `json:"rate,omitempty"`
	// ResultCode 是开店请求结果码： 	WAIT_MERCHANT_CONFIRM：等待商户确认 	当开店需要商户确认时返回此结果码，商户需要登录到商家中心e.alipay.com进行开店确认。例如，ISV帮商户开店时，出现需要改签口碑当面付主体，接口同步返回此错误码，并且支付宝发送短信告知商户登录商家中心进行改签确认，确认后进入到后续审核流程。
	ResultCode *string `json:"result_code,omitempty"`
	// ShopId 是废弃字段，支付宝门店ID。
	ShopId *string `json:"shop_id,omitempty"`
}

// OfflineMarketShopModifyResponse 是支付宝官方 V3 OpenAPI 模型。
type OfflineMarketShopModifyResponse struct {
	// ApplyId 是修改门店请求受理成功后返回的支付宝流水ID，根据此ID调用接口 	alipay.offline.market.applyorder.batchquery，能够获取当前修改店铺请求审核状态、驳回原因等信息。
	ApplyId *string `json:"apply_id,omitempty"`
	// AuditDesc 是废弃字段。
	AuditDesc *string `json:"audit_desc,omitempty"`
	// AuditStatus 是同步请求如果支付宝受理成功，将返回AUDITING状态。
	AuditStatus *string `json:"audit_status,omitempty"`
	// IsOnline 是废弃字段，T表示上架,F表示未上架。
	IsOnline *string `json:"is_online,omitempty"`
	// IsShow 是废弃字段，T表示显示，F表示隐藏。
	IsShow *string `json:"is_show,omitempty"`
}

// OfflineMarketShopQueryDetailResponse 是支付宝官方 V3 OpenAPI 模型。
type OfflineMarketShopQueryDetailResponse struct {
	// Address 是门店详细地址，注：不含省市区
	Address *string `json:"address,omitempty"`
	// AuditDesc 是门店审核状态描述，如果审核驳回则有相关的驳回理由
	AuditDesc *string `json:"audit_desc,omitempty"`
	// AuditImages 是门店审核时需要的图片; 至少包含一张门头照片，两张内景照片，必须反映真实的门店情况，审核才能够通过， 多个以英文逗号分隔
	AuditImages *string `json:"audit_images,omitempty"`
	// AuditStatus 是门店审核状态，对于系统商而言，只有三个状态，AUDITING：审核中，AUDIT_FAILED：审核驳回，AUDIT_SUCCESS：审核通过。第一次审核通过会触发门店上架。
	AuditStatus *string `json:"audit_status,omitempty"`
	// AuthLetter 是授权函图片
	AuthLetter *string `json:"auth_letter,omitempty"`
	// AvgPrice 是人均消费价格，最少1元，最大不超过99999元，请按实际情况填写；单位元；
	AvgPrice *string `json:"avg_price,omitempty"`
	// Box 是门店是否有包厢，T表示有，F表示没有，不传在客户端不作展示
	Box *string `json:"box,omitempty"`
	// BranchShopName 是分店名称，比如：万塘路店，与主门店名合并在客户端显示为：肯德基(万塘路店)
	BranchShopName *string `json:"branch_shop_name,omitempty"`
	// BrandLogo 是品牌LOGO; 图片ID，不填写则默认为门店首图main_image
	BrandLogo *string `json:"brand_logo,omitempty"`
	// BrandName 是品牌名称；不填写则默认为“其它品牌”
	BrandName *string `json:"brand_name,omitempty"`
	// BusinessCertificate 是经营许可证，只有餐饮类目需要
	BusinessCertificate *string `json:"business_certificate,omitempty"`
	// BusinessCertificateExpires 是经营许可证有效期，格式：2020-03-20，只有餐饮类目需要
	BusinessCertificateExpires *string `json:"business_certificate_expires,omitempty"`
	// BusinessTime 是营业时间;支持分段营业时间，以英文逗号分隔
	BusinessTime *string `json:"business_time,omitempty"`
	// CategoryId 是类目ID，类目初始数据由口碑提供
	CategoryId *string `json:"category_id,omitempty"`
	// CityCode 是城市编码，国标码，详见国家统计局数据
	CityCode *string `json:"city_code,omitempty"`
	// ContactNumber 是门店电话号码；支持座机和手机，在客户端对用户展现，支持多个电话，以英文逗号分隔
	ContactNumber *string `json:"contact_number,omitempty"`
	// CreateChannel 是门店创建来源；如：开放平台、支付宝客户端、口碑商家app、商家自主开店、服务商开店、全民开店-支付宝客户端、全民开店-商户app、其它
	CreateChannel *string `json:"create_channel,omitempty"`
	// DistrictCode 是区县编码，国标码，详见国家统计局数据
	DistrictCode *string `json:"district_code,omitempty"`
	// GmtShopCreate 是门店创建时间
	GmtShopCreate *string `json:"gmt_shop_create,omitempty"`
	// GmtShopModified 是门店修改时间
	GmtShopModified *string `json:"gmt_shop_modified,omitempty"`
	// ImplementId 是店铺使用的机具编号，多个以英文逗号分隔
	ImplementId *string `json:"implement_id,omitempty"`
	// IsOnline 是门店是否上架，T表示上架,F表示未上架，第一次门店审核通过后会触发上架
	IsOnline *string `json:"is_online,omitempty"`
	// IsOperatingOnline 是是否在其他平台开店，T表示有开店，F表示未开店，用于证明其开店资质
	IsOperatingOnline *string `json:"is_operating_online,omitempty"`
	// IsShow 是门店是否在客户端显示，T表示显示，F表示隐藏
	IsShow *string `json:"is_show,omitempty"`
	// IsvUid 是开发者返佣ID，重要：如果有口碑签订了返佣协议，则该字段作为返佣数据提取的依据，需要与签约协议的PID保持一致， 该字段只能在创建接口中传入，不能在修改接口中被修改
	IsvUid *string `json:"isv_uid,omitempty"`
	// Latitude 是纬度，最长15位字符（包括小数点）， 注：高德坐标系
	Latitude *string `json:"latitude,omitempty"`
	// Licence 是门店营业执照图片
	Licence *string `json:"licence,omitempty"`
	// LicenceCode 是门店营业执照编号
	LicenceCode *string `json:"licence_code,omitempty"`
	// LicenceExpires 是营业执照过期时间
	LicenceExpires *string `json:"licence_expires,omitempty"`
	// LicenceName 是门店营业执照名称
	LicenceName *string `json:"licence_name,omitempty"`
	// Longitude 是经度，最长15位字符（包括小数点）， 注：高德坐标系
	Longitude *string `json:"longitude,omitempty"`
	// MainImage 是门店首图；非常重要，推荐尺寸2000*1500
	MainImage *string `json:"main_image,omitempty"`
	// MainShopName 是主店名；比如：肯德基
	MainShopName *string `json:"main_shop_name,omitempty"`
	// NoSmoking 是是否有无烟区，T表示有无烟区，F表示没有无烟区，不传在客户端不展示
	NoSmoking *string `json:"no_smoking,omitempty"`
	// NotifyMobile 是门店店长电话号码；用于接收门店状态变更通知，收款成功通知等通知消息，不在客户端展示；多个以引文逗号分隔
	NotifyMobile *string `json:"notify_mobile,omitempty"`
	// OnlineImage 是在其他平台的开店图片，支持多张，逗号分隔
	OnlineImage *string `json:"online_image,omitempty"`
	// OperateNotifyUrl 是通知发送url;当商户的门店审核状态发生变化时，会向该地址推送消息
	OperateNotifyUrl *string `json:"operate_notify_url,omitempty"`
	// OtherAuthImages 是其它资质证明图片集；支持多张，逗号分隔
	OtherAuthImages *string `json:"other_auth_images,omitempty"`
	// Parking 是门店是否支持停车，T表示支持，F表示不支持，不传在客户端不作展示
	Parking *string `json:"parking,omitempty"`
	// PartnerId 是门店的签约ID
	PartnerId *string `json:"partner_id,omitempty"`
	// PayType 是默认付款类型；如：付款码、扫码付、声波支付、在线买单、其它
	PayType *string `json:"pay_type,omitempty"`
	// PaymentAccount 是门店收款账户，门店收款账户只能被查询，不能通过接口修改。如果为空，则表示门店收款账户为商户签约账户
	PaymentAccount *string `json:"payment_account,omitempty"`
	// PicColl 是图片集，是map转化成的json串，key是图片id,value是图片url
	PicColl *string `json:"pic_coll,omitempty"`
	// ProcessedQrCode 是经过加工后的门店收款二维码
	ProcessedQrCode *string `json:"processed_qr_code,omitempty"`
	// ProviderXiaoerUid 是门店运营归属人uid
	ProviderXiaoerUid *string `json:"provider_xiaoer_uid,omitempty"`
	// ProvinceCode 是省份编码，国标码，详见国家统计局数据
	ProvinceCode *string `json:"province_code,omitempty"`
	// QrCode 是门店收款二维码裸码
	QrCode *string `json:"qr_code,omitempty"`
	// ShopTags 是门店标签；JSON格式。包括：keyMerchant-是否重点商户（T/F）；isHallMeal-堂食（T/F）；注：若标签 key 不存在，则门店无对应的标签
	ShopTags *string `json:"shop_tags,omitempty"`
	// StoreId 是外部门店编号；最长54位字符，该编号将作为收单接口的入参， 请开发者自行确保其唯一性
	StoreId *string `json:"store_id,omitempty"`
	// ValueAdded 是门店其他的服务，门店与用户线下兑现
	ValueAdded *string `json:"value_added,omitempty"`
	// Wifi 是门店是否支持WIFI，T表示支持，F表示不支持，不传在客户端不作展示
	Wifi *string `json:"wifi,omitempty"`
}

// OfflineMarketShopSummaryBatchQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type OfflineMarketShopSummaryBatchQueryResponse struct {
	// CurrentPageNo 是当前页码
	CurrentPageNo *string `json:"current_page_no,omitempty"`
	// PageSize 是每页记录数
	PageSize *string `json:"page_size,omitempty"`
	// ShopSummaryInfos 是支付宝门店摘要信息列表
	ShopSummaryInfos *[]ShopSummaryQueryResponse `json:"shop_summary_infos,omitempty"`
	// TotalItems 是总记录数
	TotalItems *string `json:"total_items,omitempty"`
	// TotalPageNo 是总页码数目
	TotalPageNo *string `json:"total_page_no,omitempty"`
}

// ShopSummaryQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type ShopSummaryQueryResponse struct {
	// Address 是门店地址
	Address *string `json:"address,omitempty"`
	// BranchShopName 是分店名
	BranchShopName *string `json:"branch_shop_name,omitempty"`
	// BrandName 是品牌名，不填写则默认为其它品牌
	BrandName *string `json:"brand_name,omitempty"`
	// BusinessTime 是经营时间
	BusinessTime *string `json:"business_time,omitempty"`
	// CategoryInfos 是门店类目列表
	CategoryInfos *[]ShopCategoryInfo `json:"category_infos,omitempty"`
	// CityCode 是城市编码，国标码，详见国家统计局数据 点此下载
	CityCode *string `json:"city_code,omitempty"`
	// DistrictCode 是区县编码，国标码，详见国家统计局数据 点此下载
	DistrictCode *string `json:"district_code,omitempty"`
	// ExtInfo 是json 字符串表示额外信息; 	 	order_biz_mode: 点餐营业模式, REGULAR-正餐、LIGHT_FAST-轻快餐; 	pre_order: 预点餐服务可用状态, 1-服务可用、0-服务不可用 	order: 到店点餐服务可用状态：1-服务可用、0-服务不可用
	ExtInfo *string `json:"ext_info,omitempty"`
	// GmtCreate 是创建时间
	GmtCreate *string `json:"gmt_create,omitempty"`
	// IsShow 是门店是否在客户端显示，T表示显示，F表示隐藏
	IsShow *string `json:"is_show,omitempty"`
	// Latitude 是纬度，只有在query_type=KB_PROMOTER非空
	Latitude *string `json:"latitude,omitempty"`
	// Longitude 是经度，只有在query_type=KB_PROMOTER非空
	Longitude *string `json:"longitude,omitempty"`
	// MainImage 是门店首图
	MainImage *string `json:"main_image,omitempty"`
	// MainShopName 是主门店名
	MainShopName *string `json:"main_shop_name,omitempty"`
	// PerPay 是人均价格。单位是分。例如：2000，表示20元
	PerPay *string `json:"per_pay,omitempty"`
	// PicColl 是图片集，是map转化成的json串，key是图片id,value是图片url
	PicColl *string `json:"pic_coll,omitempty"`
	// ProvinceCode 是省份编码，国标码，详见国家统计局数据 点此下载
	ProvinceCode    *string          `json:"province_code,omitempty"`
	ShopCommentInfo *ShopCommentInfo `json:"shop_comment_info,omitempty"`
	// ShopId 是门店ID
	ShopId *string `json:"shop_id,omitempty"`
	// ShopType 是COMMON（普通门店）、MALL（商圈）
	ShopType *string `json:"shop_type,omitempty"`
	// Status 是门店状态，OPEN：营业中、PAUSE：暂停营业、FREEZE：已冻结、CLOSE:门店已关闭
	Status *string `json:"status,omitempty"`
}
