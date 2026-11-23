package shop

// OfflineMarketShopBatchQueryRequest 是支付宝官方 V3 OpenAPI 模型。
type OfflineMarketShopBatchQueryRequest struct {
	// BizChannel 是表示接口查询门店的业务渠道限制：ALL、 POS、不传。不传代表只查询普通门店，传入POS代表只查询简易门店，传入ALL代表查询普通门店和简易门店。
	BizChannel *string `json:"biz_channel,omitempty"`
	// PageNo 是页码，第一页传入"1"，默认500个结果为一页。此参数必须是大于0的正整数，为0时将查询报错。
	PageNo *string `json:"page_no,omitempty"`
}

// OfflineMarketShopCreateRequest 是支付宝官方 V3 OpenAPI 模型。
type OfflineMarketShopCreateRequest struct {
	// Address 是门店详细地址，地址字符长度在4-50个字符，注：不含省市区。门店详细地址按规范格式填写地址，以免影响门店搜索及活动报名：例1：道路+门牌号，“人民东路18号”；例2：道路+门牌号+标志性建筑+楼层，“四川北路1552号欢乐广场1楼”。
	Address *string `json:"address,omitempty"`
	// AuditImages 是门店审核时需要的图片；至少包含一张门头照片，两张内景照片，必须反映真实的门店情况，审核才能够通过；多个图片之间以英文逗号分隔。
	AuditImages *string `json:"audit_images,omitempty"`
	// AuthLetter 是门店授权函,营业执照与签约账号主体不一致时需要。
	AuthLetter *string `json:"auth_letter,omitempty"`
	// BizVersion 是店铺接口业务版本号，新接入的ISV，请统一传入2.0。
	BizVersion *string `json:"biz_version,omitempty"`
	// Box 是门店是否有包厢，T表示有，F表示没有，不传在客户端不作展示。
	Box *string `json:"box,omitempty"`
	// BranchShopName 是分店名称，比如：万塘路店，与主门店名合并在客户端显示为：肯德基(万塘路店)。
	BranchShopName *string `json:"branch_shop_name,omitempty"`
	// BrandLogo 是品牌LOGO; 图片ID，不填写则默认为门店首图main_image。
	BrandLogo *string `json:"brand_logo,omitempty"`
	// BrandName 是品牌名，不填写则默认为“其它品牌”。
	BrandName *string `json:"brand_name,omitempty"`
	// BusinessCertificate 是许可证，各行业所需的证照资质参见商户入驻要求；该字段只能上传一张许可证，一张以外的许可证、除营业执照和许可证之外其他证照请放在其他资质字段上传。
	BusinessCertificate *string `json:"business_certificate,omitempty"`
	// BusinessCertificateExpires 是许可证有效期，格式：2020-03-20或长期。严格按照格式填写。
	BusinessCertificateExpires *string `json:"business_certificate_expires,omitempty"`
	// BusinessTime 是请严格按"周一-周五 09:00-20:00,周六-周日 10:00-22:00"的格式进行填写，时间段不能重复，最多支持两个时间段，24小时营业请填写"00:00-23:59"
	BusinessTime *string `json:"business_time,omitempty"`
	// CategoryId 是类目id，请参考商户入驻要求。
	CategoryId *string `json:"category_id,omitempty"`
	// CityCode 是城市编码，国标码，详见国家统计局数据 点此下载。
	CityCode *string `json:"city_code,omitempty"`
	// ContactNumber 是门店电话号码；支持座机和手机，只支持数字和+-号，在客户端对用户展现， 支持多个电话， 以英文逗号分隔。
	ContactNumber *string `json:"contact_number,omitempty"`
	// Creator 是（支付宝内部参数）小二的支付宝账号。
	Creator *string `json:"creator,omitempty"`
	// DistrictCode 是区县编码，国标码，详见国家统计局数据 点此下载。
	DistrictCode *string `json:"district_code,omitempty"`
	// EnterpriseLogonId 是（支付宝内部参数）企业支付宝账号。
	EnterpriseLogonId *string `json:"enterprise_logon_id,omitempty"`
	// EnterpriseName 是（支付宝内部参数）企业支付宝账户名称。
	EnterpriseName *string `json:"enterprise_name,omitempty"`
	// ImplementId 是机具号，多个之间以英文逗号分隔。
	ImplementId *string `json:"implement_id,omitempty"`
	// IsOperatingOnline 是是否在其他平台开店，T表示有开店，F表示未开店。
	IsOperatingOnline *string `json:"is_operating_online,omitempty"`
	// IsvUid 是ISV返佣id，门店创建、或者门店交易的返佣将通过此账号反给ISV，如果有口碑签订了返佣协议，则该字段作为返佣数据提取的依据。此字段必须是个合法uid，2088开头的16位支付宝会员账号，如果传入错误将无法创建门店。
	IsvUid *string `json:"isv_uid,omitempty"`
	// Latitude 是纬度；最长15位字符（包括小数点）， 注：高德坐标系。经纬度是门店搜索和活动推荐的重要参数，录入时请确保经纬度参数准确。高德经纬度查询：http://lbs.amap.com/console/show/picker
	Latitude *string `json:"latitude,omitempty"`
	// LeadsId 是（支付宝内部参数）leads的编号。
	LeadsId *string `json:"leads_id,omitempty"`
	// Licence 是门店营业执照图片，各行业所需的证照资质参见商户入驻要求。
	Licence *string `json:"licence,omitempty"`
	// LicenceCode 是门店营业执照编号，只支持输入中文，英文和数字，营业执照信息与is_operating_online至少填一项。
	LicenceCode *string `json:"licence_code,omitempty"`
	// LicenceExpires 是营业执照过期时间。格式：2020-10-20或长期。严格按照格式填写。
	LicenceExpires *string `json:"licence_expires,omitempty"`
	// LicenceName 是门店营业执照名称。
	LicenceName *string `json:"licence_name,omitempty"`
	// Longitude 是经度；最长15位字符（包括小数点）， 注：高德坐标系。经纬度是门店搜索和活动推荐的重要参数，录入时请确保经纬度参数准确。高德经纬度查询：http://lbs.amap.com/console/show/picker
	Longitude *string `json:"longitude,omitempty"`
	// MainImage 是门店首图，非常重要，推荐尺寸2000*1500。
	MainImage *string `json:"main_image,omitempty"`
	// MainShopName 是主门店名  比如：肯德基；主店名里不要包含分店名，如“万塘路店”。主店名长度不能超过20个字符。
	MainShopName *string `json:"main_shop_name,omitempty"`
	// NoSmoking 是是否有无烟区，T表示有无烟区，F表示没有无烟区，不传在客户端不展示。
	NoSmoking *string `json:"no_smoking,omitempty"`
	// NotifyMobile 是门店店长电话号码；用于接收门店状态变更通知，收款成功通知等通知消息， 不在客户端展示。
	NotifyMobile *string `json:"notify_mobile,omitempty"`
	// NotifyUrl 是通知地址
	NotifyUrl *string `json:"notify_url,omitempty"`
	// OnlineImage 是废弃字段，请使用online_url字段替代。
	OnlineImage *string `json:"online_image,omitempty"`
	// OnlineUrl 是其他平台开店的店铺链接url，多个url使用英文逗号隔开,isv迁移到新接口使用此字段，与is_operating_online=T配套使用。
	OnlineUrl *string `json:"online_url,omitempty"`
	// OpRole 是表示以系统集成商的身份开店，开放平台现在统一传入ISV。
	OpRole *string `json:"op_role,omitempty"`
	// OperateNotifyUrl 是当商户的门店审核状态发生变化时，会向该地址推送消息。
	OperateNotifyUrl *string `json:"operate_notify_url,omitempty"`
	// OtherAuthorization 是其他资质。用于上传营业证照、许可证照外的其他资质，除已上传许可证外的其他许可证也可以在该字段上传。
	OtherAuthorization *string `json:"other_authorization,omitempty"`
	// Parking 是门店是否支持停车，T表示支持，F表示不支持，不传在客户端不作展示。
	Parking *string `json:"parking,omitempty"`
	// PartnerId 是（支付宝内部参数）服务商要操作的商户PID。
	PartnerId *string `json:"partner_id,omitempty"`
	// PayType 是（支付宝内部参数）付款方式：code_scanned_pay：付款码  	online_pay：在线买单。ISV不可以指定此字段，ISV泛行业开店默认为在线买单，云纵开店可以指定支付方式。
	PayType *string `json:"pay_type,omitempty"`
	// ProvinceCode 是省份编码，国标码，详见国家统计局数据 点此下载。
	ProvinceCode *string `json:"province_code,omitempty"`
	// RefApplyId 是（支付宝内部参数）leads开店重试支付宝流水ID。
	RefApplyId *string `json:"ref_apply_id,omitempty"`
	// RequestId 是支持英文字母和数字，由开发者自行定义（不允许重复），在门店notify消息中也会带有该参数，以此标明本次notify消息是对哪个请求的回应。
	RequestId *string `json:"request_id,omitempty"`
	// StoreId 是外部门店编号；最长32位字符，该编号将作为收单接口的入参， 请开发者自行确保其唯一性。
	StoreId *string `json:"store_id,omitempty"`
	// ValueAdded 是门店其他的服务，门店与用户线下兑现。
	ValueAdded *string `json:"value_added,omitempty"`
	// Version 是废弃字段，使用biz_version字段替代。
	Version *string `json:"version,omitempty"`
	// Wifi 是门店是否支持WIFI，T表示支持，F表示不支持，不传在客户端不作展示。
	Wifi *string `json:"wifi,omitempty"`
}

// OfflineMarketShopModifyRequest 是支付宝官方 V3 OpenAPI 模型。
type OfflineMarketShopModifyRequest struct {
	// Address 是门店详细地址，地址字符长度在4-50个字符。门店详细地址，格式（不含省市区）：例1：道路+门牌号，“人民东路18号”；例2：道路+门牌号+标志性建筑+楼层；注：门店详细地址按规范格式填写地址，以免影响门店搜索及活动报名
	Address *string `json:"address,omitempty"`
	// AuditImages 是门店审核时需要的图片; 至少包含一张门头照片，两张内景照片，必须反映真实的门店情况，审核才能够通过， 多个以英文逗号分隔。
	AuditImages *string `json:"audit_images,omitempty"`
	// AuthLetter 是门店授权函。
	AuthLetter *string `json:"auth_letter,omitempty"`
	// AvgPrice 是人均消费价格，最少1元，最大不超过99999元，请按实际情况填写，单位元。
	AvgPrice *string `json:"avg_price,omitempty"`
	// BizVersion 是店铺接口业务版本号，新接入的ISV，请统一传入2.0。
	BizVersion *string `json:"biz_version,omitempty"`
	// Box 是包厢支持，T表示有包厢，F表示无包厢；不传值默认F。
	Box *string `json:"box,omitempty"`
	// BranchShopName 是分店名称，比如：万塘路店，与主门店名合并在客户端显示为：肯德基(万塘路店)。分店名长度需在2~20个字符之间。
	BranchShopName *string `json:"branch_shop_name,omitempty"`
	// BrandLogo 是品牌LOGO; 图片ID，不填写则默认为门店首图main_image。
	BrandLogo *string `json:"brand_logo,omitempty"`
	// BrandName 是品牌名称；不填写则默认为“其它品牌”。
	BrandName *string `json:"brand_name,omitempty"`
	// BusinessCertificate 是许可证，各行业所需的证照资质参见商户入驻要求；该字段只能上传一张许可证，一张以外的许可证、除营业执照和许可证之外其他证照请放在其他资质字段上传。
	BusinessCertificate *string `json:"business_certificate,omitempty"`
	// BusinessCertificateExpires 是许可证有效期，格式：2020-03-20或长期。严格按照格式填写。
	BusinessCertificateExpires *string `json:"business_certificate_expires,omitempty"`
	// BusinessTime 是请严格按"周一-周五 09:00-20:00,周六-周日 10:00-22:00"的格式进行填写，时间段不能重复，最多支持两个时间段，24小时营业请填写"00:00-23:59"
	BusinessTime *string `json:"business_time,omitempty"`
	// CategoryId 是废弃字段，不支持修改类目。 	类目id，请参考商户入驻要求。
	CategoryId *string `json:"category_id,omitempty"`
	// CityCode 是城市编码，国标码，详见国家统计局数据 点此下载。
	CityCode *string `json:"city_code,omitempty"`
	// ContactNumber 是门店电话号码；支持座机和手机，在客户端对用户展现，支持多个电话，以英文逗号分隔。
	ContactNumber *string `json:"contact_number,omitempty"`
	// DistrictCode 是区县编码，国标码，详见国家统计局数据 点此下载。
	DistrictCode *string `json:"district_code,omitempty"`
	// ImplementId 是店铺使用的机具编号，多个以英文逗号分隔。
	ImplementId *string `json:"implement_id,omitempty"`
	// IsOperatingOnline 是是否在其他平台开店，T表示有开店，F表示未开店。
	IsOperatingOnline *string `json:"is_operating_online,omitempty"`
	// IsShow 是废弃字段，T表示显示，F表示隐藏，默认为T。
	IsShow *string `json:"is_show,omitempty"`
	// Latitude 是纬度，注：高德坐标系。经纬度是门店搜索和活动推荐的重要参数，录入时请确保经纬度参数准确。高德经纬度查询：http://lbs.amap.com/console/show/picker
	Latitude *string `json:"latitude,omitempty"`
	// Licence 是门店营业执照图片，各行业所需的证照资质参见：https://doc.open.alipay.com/doc2/detail.htm?spm=a219a.7629140.0.0.NBvQVP&treeId=78&articleId=104497&docType=1。
	Licence *string `json:"licence,omitempty"`
	// LicenceCode 是门店营业执照编号。只支持输入中文，英文和数字。
	LicenceCode *string `json:"licence_code,omitempty"`
	// LicenceExpires 是营业执照过期时间。格式：2020-10-20或长期。严格按照格式填写。
	LicenceExpires *string `json:"licence_expires,omitempty"`
	// LicenceName 是门店营业执照名称。
	LicenceName *string `json:"licence_name,omitempty"`
	// Longitude 是经度，注：高德坐标系。经纬度是门店搜索和活动推荐的重要参数，录入时请确保经纬度参数准确。高德经纬度查询：http://lbs.amap.com/console/show/picker
	Longitude *string `json:"longitude,omitempty"`
	// MainImage 是门店首图；非常重要，推荐尺寸2000*1500。
	MainImage *string `json:"main_image,omitempty"`
	// MainShopName 是主门店名  比如：肯德基；主店名里不要包含分店名，如“万塘路店”。主店名长度不能超过20个字符。【xxx店】、(xxx店)、（xxx店）、[xxx店]、、xxx店，等类似的主店名都是不合法的，如果需要录入分店，请填写到branch_shop_name字段中。
	MainShopName *string `json:"main_shop_name,omitempty"`
	// NoSmoking 是无烟区支持，T表示禁烟，F表示不禁烟；不传值默认F。
	NoSmoking *string `json:"no_smoking,omitempty"`
	// NotifyMobile 是门店店长电话号码；用于接收门店状态变更通知，收款成功通知等通知消息，不在客户端展示；多个以引文逗号分隔。
	NotifyMobile *string `json:"notify_mobile,omitempty"`
	// NotifyUrl 是通知地址
	NotifyUrl *string `json:"notify_url,omitempty"`
	// OnlineImage 是废弃字段，请使用online_url字段替代。
	OnlineImage *string `json:"online_image,omitempty"`
	// OnlineUrl 是其他平台开店的店铺链接url,多个url使用英文逗号隔开，isv迁移到新接口使用此字段，与is_operating_online=T配套使用。
	OnlineUrl *string `json:"online_url,omitempty"`
	// OpId 是（支付宝内部参数）操作员的支付账号ID(服务商ID、城市经理ID)。
	OpId *string `json:"op_id,omitempty"`
	// OpRole 是表示以系统集成商的身份开店，开放平台现在统一传入ISV。
	OpRole *string `json:"op_role,omitempty"`
	// OperateNotifyUrl 是通知发送url;当商户的门店审核状态发生变化时，会向该地址推送消息。
	OperateNotifyUrl *string `json:"operate_notify_url,omitempty"`
	// OtherAuthorization 是其他资质。用于上传营业证照、许可证照外的其他资质，除已上传许可证外的其他许可证也可以在该字段上传。
	OtherAuthorization *string `json:"other_authorization,omitempty"`
	// Parking 是门店是否支持停车，T表示支持，F表示不支持，不传在客户端不作展示。
	Parking *string `json:"parking,omitempty"`
	// PartnerId 是（支付宝内部参数）服务商要操作的商户PID。
	PartnerId *string `json:"partner_id,omitempty"`
	// PayType 是（支付宝内部参数）付款方式：code_scanned_pay：付款码  	online_pay：在线买单。ISV不可以指定此字段，ISV泛行业开店默认为在线买单，云纵开店可以指定支付方式。
	PayType *string `json:"pay_type,omitempty"`
	// ProvinceCode 是省份编码，国标码，详见国家统计局数据 点此下载”。
	ProvinceCode *string `json:"province_code,omitempty"`
	// RequestId 是外部请求ID; 标识ISV本次修改的请求，由开发者自定义，不同的请求使用不同的ID，在门店notify消息中也会带有该参数，以此标明本次notify消息是对哪个请求的回应
	RequestId *string `json:"request_id,omitempty"`
	// ShopId 是支付宝门店ID。
	ShopId *string `json:"shop_id,omitempty"`
	// StoreId 是外部门店编号；最长32位字符，该编号将作为收单接口的入参， 请开发者自行确保其唯一性。
	StoreId *string `json:"store_id,omitempty"`
	// ValueAdded 是门店其他的服务，门店与用户线下兑现。
	ValueAdded *string `json:"value_added,omitempty"`
	// Version 是废弃字段，使用biz_version字段替代。
	Version *string `json:"version,omitempty"`
	// Wifi 是门店是否支持WIFI，T表示支持，F表示不支持，不传在客户端不作展示。
	Wifi *string `json:"wifi,omitempty"`
}

// OfflineMarketShopSummaryBatchQueryRequest 是支付宝官方 V3 OpenAPI 模型。
type OfflineMarketShopSummaryBatchQueryRequest struct {
	// BizChannel 是表示接口查询门店的业务渠道限制：ALL、 POS、不传。不传代表只查询普通门店，传入POS代表只查询简易门店，传入ALL代表查询普通门店和简易门店。
	BizChannel *string `json:"biz_channel,omitempty"`
	// BrandName 是品牌名
	BrandName *string `json:"brand_name,omitempty"`
	// CityCode 是城市编码，国标码，如：120100表示天津市
	CityCode *string `json:"city_code,omitempty"`
	// DistrictCode 是区域编码，如：120104表示南开区
	DistrictCode *string `json:"district_code,omitempty"`
	// OpRole 是表示接口业务的调用方身份：ISV、 服务商身份标识。传入ISV代表系统集成商身份。传入PROVIDER代表服务商。
	OpRole *string `json:"op_role,omitempty"`
	// PageNo 是页码，留空标示第一页，默认 20个结果为一页
	PageNo *int64 `json:"page_no,omitempty"`
	// PageSize 是每页记录数，默认20，最大 100
	PageSize *int64 `json:"page_size,omitempty"`
	// ProvinceCode 是省份编码，国标码，如：120000表示天津
	ProvinceCode *string `json:"province_code,omitempty"`
	// QueryType 是门店数据查询类型，根据类型可以返回指定的门店数据，目前支持的类型如下： 	BRAND_RELATION ： 品牌商关联店铺 	MALL_SELF ：MALL自己的门店 	MALL_RELATION：MALL关联下的门店 	MERCHANT_SELF:商户自己的门店 	KB_PROMOTER：口碑客推广者
	QueryType *string `json:"query_type,omitempty"`
	// RelatedPartnerId 是query_type查询类型下所关联的商户PID
	RelatedPartnerId *string `json:"related_partner_id,omitempty"`
	// ShopId 是门店ID
	ShopId *string `json:"shop_id,omitempty"`
	// ShopStatus 是门店状态，传入多个状态，多个状态使用英文逗号隔开，例如：PAUSED,OPEN 	店铺状态：OPEN（营业）、PAUSED（暂停）、INIT（初始）、FREEZE（冻结）、CLOSED（关店）
	ShopStatus *string `json:"shop_status,omitempty"`
}

// OfflineMarketShopCategoryQueryRequest is the query request model for alipay.offline.market.shop.category.query.
type OfflineMarketShopCategoryQueryRequest struct {
	// CategoryId 是类目ID，如果为空则查询全部类目。
	CategoryId *string `json:"category_id,omitempty" query:"category_id,omitempty"`
	// OpRole 是表示接口业务的调用方身份,默认不填标识为ISV。
	OpRole *string `json:"op_role,omitempty" query:"op_role,omitempty"`
}

// OfflineMarketShopQueryDetailRequest is the query request model for alipay.offline.market.shop.querydetail.
type OfflineMarketShopQueryDetailRequest struct {
	// ShopId 是支付宝门店ID
	ShopId *string `json:"shop_id,omitempty" query:"shop_id,omitempty"`
	// OpRole 是服务商及商户调用情况下务必传递。操作人角色，默认商户操作:MERCHANT；服务商操作：PROVIDER；ISV: 不需要填写
	OpRole *string `json:"op_role,omitempty" query:"op_role,omitempty"`
}
