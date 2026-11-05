// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package merchant

import (
	"github.com/tiyee/alipay-sdk-go/core"
)

// ExpandAppRecommendAccountCreateRequest is the official V3 AntMerchantExpandApprecommendAccountCreateModel schema.
type ExpandAppRecommendAccountCreateRequest struct {
	// AccNo represents acc_no. 待绑定账号PID
	AccNo *string `json:"acc_no,omitzero"`
	// AppNo represents app_no. 待绑定小程序的app_id
	AppNo *string `json:"app_no,omitzero"`
}

// ExpandItemCreateRequest is the official V3 AntMerchantExpandItemCreateModel schema.
type ExpandItemCreateRequest struct {
	// DetailUrl represents detail_url. 详情地址
	DetailUrl *string `json:"detail_url,omitzero"`
	// ExtInfo represents ext_info. 商品扩展信息：可以解析成 Map<String, String> 的 json string
	ExtInfo *string `json:"ext_info,omitzero"`
	// ExternalItemId represents external_item_id. 外部商品ID
	ExternalItemId *string `json:"external_item_id,omitzero"`
	// FrontCategoryId represents front_category_id. 前台类目id：target_type + target_id 和 front_category_id 二选一
	FrontCategoryId *string `json:"front_category_id,omitzero"`
	// LabelList represents label_list. 商品标签列表
	LabelList *[]ItemLabelCreateInfo `json:"label_list,omitzero"`
	// MainPic represents main_pic. 主图地址
	MainPic *string `json:"main_pic,omitzero"`
	// Name represents name. 商品名称
	Name *string `json:"name,omitzero"`
	// Scene represents scene. 场景：GAS_CHARGE（加油）
	Scene *string `json:"scene,omitzero"`
	// SkuList represents sku_list. 商品sku列表，至少有一个
	SkuList *[]ItemSkuCreateInfo `json:"sku_list,omitzero"`
	// TargetId represents target_id. 商户归属主体id
	TargetId *string `json:"target_id,omitzero"`
	// TargetType represents target_type. 商品归属主体类型：target_type + target_id 和 front_category_id 二选一 商品归属主体类型: 5: 店铺 4: 主站MID 3: 参与者 2: 角色 1: 联系人
	TargetType *string `json:"target_type,omitzero"`
}

// ExpandItemModifyRequest is the official V3 AntMerchantExpandItemModifyModel schema.
type ExpandItemModifyRequest struct {
	// DetailUrl represents detail_url. 商品详情地址
	DetailUrl *string `json:"detail_url,omitzero"`
	// ExtInfo represents ext_info. 商品扩展信息：可以解析成 Map<String, String> 的 json string
	ExtInfo *string `json:"ext_info,omitzero"`
	// ExternalItemId represents external_item_id. 外部商品ID（item_id不为空时作为普通更新项，item_id为空时作为更新KEY）
	ExternalItemId *string `json:"external_item_id,omitzero"`
	// FrontCategoryId represents front_category_id. 前台类目id, null表示未分类
	FrontCategoryId *string `json:"front_category_id,omitzero"`
	// ItemId represents item_id. 商品id(和external_item_id不能同时为空)
	ItemId *string `json:"item_id,omitzero"`
	// LabelList represents label_list. 商品标签列表，null表示不修改标签列表，空list表示清空标签
	LabelList *[]ItemLabelModifyInfo `json:"label_list,omitzero"`
	// MainPic represents main_pic. 商品主图
	MainPic *string `json:"main_pic,omitzero"`
	// Name represents name. 商品名称
	Name *string `json:"name,omitzero"`
	// OpTimestamp represents op_timestamp. 更新时间戳（只处理时间戳最大的一次请求）
	OpTimestamp *int64 `json:"op_timestamp,omitzero"`
	// Scene represents scene. 场景：GAS_CHARGE（加油）
	Scene *string `json:"scene,omitzero"`
	// SkuList represents sku_list. SKU列表（会和商品已存在SKU做差异化比较后做增删改操作）
	SkuList *[]ItemSkuModifyInfo `json:"sku_list,omitzero"`
	// Status represents status. 商品状态： EFFECT（有效）、 INVALID（无效）
	Status *string `json:"status,omitzero"`
}

// ExpandItemOpenBatchQueryRequest is the official V3 AntMerchantExpandItemOpenBatchqueryModel schema.
type ExpandItemOpenBatchQueryRequest struct {
	// ItemIdList represents item_id_list. 商品ID列表
	ItemIdList *[]string `json:"item_id_list,omitzero"`
}

// ExpandItemOpenCreateRequest is the official V3 AntMerchantExpandItemOpenCreateModel schema.
type ExpandItemOpenCreateRequest struct {
	// Description represents description. 商品描述
	Description *string `json:"description,omitzero"`
	// ExtInfo represents ext_info. 商品扩展信息（具体KEY请参见产品文档）。 小程序订单中心场景接入参见 <a href="https://opendocs.alipay.com/mini/024hj4">接入指南</a>。
	ExtInfo *[]ItemExtInfo `json:"ext_info,omitzero"`
	// MaterialList represents material_list. 素材列表（最多3个），素材内容为素材key（素材key为<a href="https://opendocs.alipay.com/apis/api_4/alipay.merchant.item.file.upload">alipay.merchant.item.file.upload</a> 接口返回的 material_key）
	MaterialList *[]MaterialCreateInfo `json:"material_list,omitzero"`
	// Name represents name. 商品名称
	Name *string `json:"name,omitzero"`
	// PropertyList represents property_list. 商品属性列表
	PropertyList *[]ItemPropertyInfo `json:"property_list,omitzero"`
	// Scene represents scene. 业务场景码。
	Scene *string `json:"scene,omitzero"`
	// SkuList represents sku_list. 商品SKU列表（至少1个，最多20个）
	SkuList *[]SkuCreateInfo `json:"sku_list,omitzero"`
	// StandardCategoryId represents standard_category_id. 商品所属标准类目ID，标品及非标品都需传入。 请填入 <a href="https://opendocs.alipay.com/mini/011lxt">小程序商品类目表</a> 中三级类目ID。
	StandardCategoryId *string `json:"standard_category_id,omitzero"`
	// TargetId represents target_id. 商品归属主体ID 例：商品归属主体类型target_type为店铺，则商品归属主体ID为店铺ID（支付宝侧店铺ID）；归属主体类型target_type为小程序，则归属主体ID为小程序ID
	TargetId *string `json:"target_id,omitzero"`
	// TargetType represents target_type. 商品归属主体类型。枚举如下： 5：店铺。 8：小程序。
	TargetType *string `json:"target_type,omitzero"`
	// Type represents type. 商品类型。枚举支持： STANDARD_GOODS：标品，一般是具有明确、标准规格、型号、参数的商品，如：手机、数码产品、大多数的家电 。 NON_STANDARD_GOODS：非标品，则是在这些方面没有统一的行业标准和参数规格的商品，如：服装、鞋袜等。
	Type *string `json:"type,omitzero"`
}

// ExpandItemOpenModifyRequest is the official V3 AntMerchantExpandItemOpenModifyModel schema.
type ExpandItemOpenModifyRequest struct {
	// Description represents description. 商品描述
	Description *string `json:"description,omitzero"`
	// ExtInfo represents ext_info. 商品扩展信息（覆盖商品已存在扩展信息、具体KEY请参见产品文档）
	ExtInfo *[]ItemExtInfo `json:"ext_info,omitzero"`
	// ItemId represents item_id. 商品ID
	ItemId *string `json:"item_id,omitzero"`
	// MaterialList represents material_list. 商品素材列表（会和商品已存在素材做差异化比较后做增删改操作）
	MaterialList *[]MaterialModifyInfo `json:"material_list,omitzero"`
	// Name represents name. 商品名称
	Name *string `json:"name,omitzero"`
	// PropertyList represents property_list. 商品属性列表（覆盖商品已存在属性）
	PropertyList *[]ItemPropertyInfo `json:"property_list,omitzero"`
	// SkuList represents sku_list. SKU列表（会和商品已存在SKU做差异化比较后做增删改操作）
	SkuList *[]SkuModifyInfo `json:"sku_list,omitzero"`
	// StandardCategoryId represents standard_category_id. 商品所属标准类目ID，当 type 为STANDARD_GOODS（标品）时必填。 请填入 https://opendocs.alipay.com/mini/011lxt 小程序商品类目表中三级类目ID。
	StandardCategoryId *string `json:"standard_category_id,omitzero"`
	// Status represents status. 商品状态
	Status *string `json:"status,omitzero"`
	// Type represents type. 商品类型：STANDARD_GOODS（标品）、NON_STANDARD_GOODS（非标品）
	Type *string `json:"type,omitzero"`
}

// ExpandItemStatusModifyRequest is the official V3 AntMerchantExpandItemStatusModifyModel schema.
type ExpandItemStatusModifyRequest struct {
	// ItemId represents item_id. 商品ID
	ItemId *string `json:"item_id,omitzero"`
	// Status represents status. 商品状态：EFFECT、INVALID
	Status *string `json:"status,omitzero"`
}

// ExpandShopConsultRequest is the official V3 AntMerchantExpandShopConsultModel schema.
type ExpandShopConsultRequest struct {
	// BizCards represents biz_cards. "门店结算卡信息。本业务当前只允许传入一张结算卡。 说明：本参数仅直付通业务使用，其余业务无需关注。"
	BizCards *[]SettleCardInfo `json:"biz_cards,omitzero"`
	// BrandId represents brand_id. 品牌id，非加油站等特殊门店无需关注。
	BrandId *string `json:"brand_id,omitzero"`
	// BusinessAddress represents business_address.
	BusinessAddress *AddressInfo `json:"business_address,omitzero"`
	// BusinessTime represents business_time. 店铺经营时间。
	BusinessTime *[]ShopBusinessTime `json:"business_time,omitzero"`
	// CertImage represents cert_image. 营业执照图片KEY。通过<a href="https://opendocs.alipay.com/apis/api_1/ant.merchant.expand.indirect.image.upload">ant.merchant.expand.indirect.image.upload</a> 接口上传图片后得到的 image_id
	CertImage *string `json:"cert_image,omitzero"`
	// CertName represents cert_name. 营业执照名称，填写值为营业执照或统一社会信用代码证上的名称。
	CertName *string `json:"cert_name,omitzero"`
	// CertNo represents cert_no. 证件号码，请填写店铺营业执照号。
	CertNo *string `json:"cert_no,omitzero"`
	// CertType represents cert_type. 证件类型
	CertType *string `json:"cert_type,omitzero"`
	// ContactInfos represents contact_infos. 联系人信息。如果填写，其中名称必填，手机、固话、email 三选一必填。
	ContactInfos *[]ContactInfo `json:"contact_infos,omitzero"`
	// ContactMobile represents contact_mobile. 店铺联系手机，与店铺联系固话二选一必填
	ContactMobile *string `json:"contact_mobile,omitzero"`
	// ContactPhone represents contact_phone. 店铺的联系固话，和店铺联系手机二选一必填
	ContactPhone *string `json:"contact_phone,omitzero"`
	// ExtInfos represents ext_infos. 扩展信息列表。key值需要向对应行业的bd进行申请。
	ExtInfos *[]ShopExtInfo `json:"ext_infos,omitzero"`
	// IpRoleId represents ip_role_id. 商户角色id，表示将要开的店属于哪个商户角色。对于直连开店场景，填写商户pid；对于间连开店场景（线上、线下、直付通），填写商户smid。特别说明：IoT设备三绑定场景统一填写商户pid
	IpRoleId *string `json:"ip_role_id,omitzero"`
	// LegalCertNo represents legal_cert_no. 法人身份证号
	LegalCertNo *string `json:"legal_cert_no,omitzero"`
	// LegalName represents legal_name. 法人名称。
	LegalName *string `json:"legal_name,omitzero"`
	// LicenseAuthLetterImage represents license_auth_letter_image. 营业执照授权函。需传入<a href="https://opendocs.alipay.com/apis/api_1/ant.merchant.expand.indirect.image.upload">ant.merchant.expand.indirect.image.upload</a> 接口上传图片后得到的 image_id。 蚂蚁门店管理场景可选。
	LicenseAuthLetterImage *string `json:"license_auth_letter_image,omitzero"`
	// Memo represents memo. 备注
	Memo *string `json:"memo,omitzero"`
	// OutDoorImages represents out_door_images. 门头照 id，需传入 <a href="https://opendocs.alipay.com/apis/api_1/ant.merchant.expand.indirect.image.upload">ant.merchant.expand.indirect.image.upload</a> 接口上传图片后得到的 image_id
	OutDoorImages *string `json:"out_door_images,omitzero"`
	// Qualifications represents qualifications. 行业特殊资质。
	Qualifications *[]IndustryQualificationInfo `json:"qualifications,omitzero"`
	// Scene represents scene. 场景，非加油站等特殊门店无需关注。
	Scene *string `json:"scene,omitzero"`
	// SettleAlipayLogonId represents settle_alipay_logon_id. 结算支付宝账号的登录号，需传入小程序归属账号。若需增加收款账号，可在创建门店后登录<a href="https://b.alipay.com/index2.htm">商家中心</a> ，在 账户中心 > 经营信息 > 门店管理 中选择对应门店选择 更多 > 编辑 > 新增收款账号。
	SettleAlipayLogonId *string `json:"settle_alipay_logon_id,omitzero"`
	// ShopCategory represents shop_category. 新版门店类目标准二级类目code。类目标准及与原类目映射关系参见 <a href="https://ur.alipay.com/4cmn0n">支付宝门店类目-最新</a> 表格。
	ShopCategory *string `json:"shop_category,omitzero"`
	// ShopName represents shop_name. 店铺名称。直连开店要保证全局店铺名称+地址唯一，间连开店要保证服务商pid下店铺名称+地址唯一
	ShopName *string `json:"shop_name,omitzero"`
	// ShopType represents shop_type. 店铺经营类型，01表示直营，02表示加盟
	ShopType *string `json:"shop_type,omitzero"`
	// StoreId represents store_id. 门店编号，表示该门店在该商户角色id(直连pid，间连smid)下，由商户自己定义的外部门店编号
	StoreId *string `json:"store_id,omitzero"`
}

// ExpandShopReceiptAccountSaveRequest is the official V3 AntMerchantExpandShopReceiptaccountSaveModel schema.
type ExpandShopReceiptAccountSaveRequest struct {
	// Promise represents promise. 是否承诺收单账号信息准确。具体承诺信息可查看 <a href="https://gw.alipayobjects.com/os/bmw-prod/922bafa8-a712-4f79-aa32-6f08d7359a5c.docx">门店信息承诺函</a>。 Y 是 N 否
	Promise *string `json:"promise,omitzero"`
	// ReceiptAccountId represents receipt_account_id. 收单账号
	ReceiptAccountId *string `json:"receipt_account_id,omitzero"`
	// ShopId represents shop_id. 店铺ID
	ShopId *string `json:"shop_id,omitzero"`
}

// ExpandShopModifyRequest is the official V3 AntMerchantExpandShopModifyModel schema.
type ExpandShopModifyRequest struct {
	// BizCards represents biz_cards. "门店结算卡信息。本业务当前只允许传入一张结算卡。 说明：本参数仅直付通业务使用，其余业务无需关注。"
	BizCards *[]SettleCardInfo `json:"biz_cards,omitzero"`
	// BrandId represents brand_id. 品牌id。若无需更新本信息项，可以不填写。可以从B/P站品牌工具箱获取品牌id
	BrandId *string `json:"brand_id,omitzero"`
	// BusinessAddress represents business_address.
	BusinessAddress *AddressInfo `json:"business_address,omitzero"`
	// BusinessTime represents business_time. 门店经营时间。若无需更新本信息项，可以不填写
	BusinessTime *[]ShopBusinessTime `json:"business_time,omitzero"`
	// CertImage represents cert_image. 营业执照图片KEY
	CertImage *string `json:"cert_image,omitzero"`
	// CertName represents cert_name. 营业执照名称，填写值为营业执照或统一社会信用代码证上的名称。
	CertName *string `json:"cert_name,omitzero"`
	// CertNo represents cert_no. 证件号码。请填写门店营业执照号。
	CertNo *string `json:"cert_no,omitzero"`
	// CertType represents cert_type. 证件类型
	CertType *string `json:"cert_type,omitzero"`
	// ContactInfos represents contact_infos.
	ContactInfos *ContactInfo `json:"contact_infos,omitzero"`
	// ContactMobile represents contact_mobile. 门店联系手机
	ContactMobile *string `json:"contact_mobile,omitzero"`
	// ContactPhone represents contact_phone. 门店的联系固话
	ContactPhone *string `json:"contact_phone,omitzero"`
	// Cover represents cover. 封面图，其值为使用ant.merchant.expand.indirect.image.upload上传图片得到的一串oss key。若无需更新本信息项，可以不填写。若填写则会整体覆盖原有的封面图
	Cover *string `json:"cover,omitzero"`
	// ExpireTime represents expire_time. 证件过期时间,当是神券团购门店（ext_infos传了biz_source=APLUS_SHOP）并且license_type=3118或者A521时，必填
	ExpireTime *string `json:"expire_time,omitzero"`
	// ExtInfos represents ext_infos. 扩展信息列表。key值需要向对应行业的bd进行申请。传入本项时，按key进行更新或补充。 如果是针对神券团购门店的操作，该值必传，且无需向bd申请。
	ExtInfos *[]ShopExtInfo `json:"ext_infos,omitzero"`
	// IndustryInfo represents industry_info. 行业信息，例如所属集团、特殊标签
	IndustryInfo *[]MerchantShopIndustryInfo `json:"industry_info,omitzero"`
	// IndustryLicense represents industry_license. 门店资质信息
	IndustryLicense *[]ShopIndustryLicense `json:"industry_license,omitzero"`
	// IpRoleId represents ip_role_id. 商户角色id，表示将要修改的门店属于哪个商户角色，和store_id一起定位待修改的门店。对于直连场景，填写商户pid；对于间连场景（线上、线下、直付通），填写商户smid。本信息项不可修改
	IpRoleId *string `json:"ip_role_id,omitzero"`
	// LegalCertNo represents legal_cert_no. 法人身份证号。如果要更新，必须和legal_name成对传入
	LegalCertNo *string `json:"legal_cert_no,omitzero"`
	// LegalName represents legal_name. 法人名称。如果要更新，必须和legal_cert_no成对传入
	LegalName *string `json:"legal_name,omitzero"`
	// LicenseAuthLetterImage represents license_auth_letter_image. 营业执照授权函。其值为使用ant.merchant.expand.indirect.image.upload上传图片得到的一串oss key。若无需更新本信息项，可以不填写
	LicenseAuthLetterImage *string `json:"license_auth_letter_image,omitzero"`
	// Memo represents memo. 备注
	Memo *string `json:"memo,omitzero"`
	// NewBusinessTime represents new_business_time. 门店营业时间，支持一周7天，支持1天内多时间段
	NewBusinessTime *[]ComplexBusinessTime `json:"new_business_time,omitzero"`
	// OutDoorImages represents out_door_images. 门头照，其值为使用ant.merchant.expand.indirect.image.upload上传图片得到的一串oss key。若无需更新本信息项，可以不填写。若填写则会整体覆盖原有的门头照列表
	OutDoorImages *[]string `json:"out_door_images,omitzero"`
	// Qualifications represents qualifications. 行业特殊资质。若无需更新本信息项，可以不填写。填写会按类型进行更新或补充
	Qualifications *[]IndustryQualificationInfo `json:"qualifications,omitzero"`
	// Scene represents scene. 场景。若无需更新本信息项，可以不填写
	Scene *string `json:"scene,omitzero"`
	// SettleAlipayLogonId represents settle_alipay_logon_id. 结算支付宝账号的登录号。若无需更新本信息项，可以不填写
	SettleAlipayLogonId *string `json:"settle_alipay_logon_id,omitzero"`
	// ShopCategory represents shop_category. 新版门店类目标准二级类目code。类目标准及与原类目映射关系参见 <a href="https://gw.alipayobjects.com/os/bmw-prod/4b3f82df-e53e-4b84-bc41-fe025101e726.xlsx">支付宝门店类目-最新</a> 表格。
	ShopCategory *string `json:"shop_category,omitzero"`
	// ShopId represents shop_id. 蚂蚁门店id，用于查出待修改的门店。填写本参数的话，store_id和ip_role_id可以不填。本信息项不可修改
	ShopId *string `json:"shop_id,omitzero"`
	// ShopMainType represents shop_main_type. 主要针对医疗行业门店主体类型进件使用，不同门店认证主体需要相应的主体资质证书。 企业营业执照：ENTERPRISE; 事业单位法人证书：INST_RGST_CTF; 民办非企业单位登记证书：PRIVATE_NON_ENTERPRISE; 企业（澳门）：AOMEN_ENTERPRISE
	ShopMainType *string `json:"shop_main_type,omitzero"`
	// ShopName represents shop_name. 门店名称。直连开店要保证全局门店名称+地址唯一，间连开店要保证服务商pid下门店名称+地址唯一。若无需更新本信息项，可以不填写
	ShopName *string `json:"shop_name,omitzero"`
	// StoreId represents store_id. 门店编号，表示该门店在该商户角色id(直连pid，间连smid)下，由商户自己定义的外部门店编号。用于查出待修改的门店。
	StoreId *string `json:"store_id,omitzero"`
}

// ExpandShopCreateRequest is the official V3 AntMerchantExpandShopCreateModel schema.
type ExpandShopCreateRequest struct {
	// BizCards represents biz_cards. 门店结算卡信息。本业务当前只允许传入一张结算卡。 说明：本参数仅直付通业务使用，其余业务无需关注。
	BizCards *[]SettleCardInfo `json:"biz_cards,omitzero"`
	// BrandId represents brand_id. 品牌id，非加油站等特殊门店无需关注。通过BP站品牌工具箱可以获取品牌id
	BrandId *string `json:"brand_id,omitzero"`
	// BusinessAddress represents business_address.
	BusinessAddress *AddressInfo `json:"business_address,omitzero"`
	// BusinessTime represents business_time. 门店经营时间。
	BusinessTime *[]ShopBusinessTime `json:"business_time,omitzero"`
	// CertImage represents cert_image. 营业执照图片KEY
	CertImage *string `json:"cert_image,omitzero"`
	// CertName represents cert_name. 营业执照名称，填写值为营业执照或统一社会信用代码证上的名称。
	CertName *string `json:"cert_name,omitzero"`
	// CertNo represents cert_no. 证件号码，请填写门店营业执照号。
	CertNo *string `json:"cert_no,omitzero"`
	// CertStartExpireTime represents cert_start_expire_time. 营业执照的开始时间
	CertStartExpireTime *string `json:"cert_start_expire_time,omitzero"`
	// CertType represents cert_type. 证件类型
	CertType *string `json:"cert_type,omitzero"`
	// ContactInfos represents contact_infos. 联系人信息
	ContactInfos *[]ContactInfo `json:"contact_infos,omitzero"`
	// ContactMobile represents contact_mobile. 门店联系手机号
	ContactMobile *string `json:"contact_mobile,omitzero"`
	// ContactPhone represents contact_phone. 门店的联系固定电话
	ContactPhone *string `json:"contact_phone,omitzero"`
	// Cover represents cover. 需传入<a href="https://opendocs.alipay.com/apis/api_1/ant.merchant.expand.indirect.image.upload">ant.merchant.expand.indirect.image.upload</a> 接口上传图片后得到的 image_id。
	Cover *string `json:"cover,omitzero"`
	// ExpireTime represents expire_time. 营业执照过期时间
	ExpireTime *string `json:"expire_time,omitzero"`
	// ExtInfos represents ext_infos. 扩展信息列表。key值需要向对应行业的bd进行申请。 如果是针对神券团购门店的操作，该值必传，且无需向bd申请。
	ExtInfos *[]ShopExtInfo `json:"ext_infos,omitzero"`
	// FundProofs represents fund_proofs. 需传入<a href="https://opendocs.alipay.com/apis/api_1*ant.merchant.expand.indirect.image.upload">ant.merchant.expand.indirect.image.upload</a> 接口上传图片后得到的 image_id。
	FundProofs *[]string `json:"fund_proofs,omitzero"`
	// FundType represents fund_type. 资金关系说明
	FundType *string `json:"fund_type,omitzero"`
	// IdentityBackUrl represents identity_back_url. 需传入<a href="https://opendocs.alipay.com/apis/api_1*ant.merchant.expand.indirect.image.upload">ant.merchant.expand.indirect.image.upload</a> 接口上传图片后得到的 image_id。
	IdentityBackUrl *string `json:"identity_back_url,omitzero"`
	// IdentityEndTime represents identity_end_time. 身份证有效结束时间,国徽面信息
	IdentityEndTime *string `json:"identity_end_time,omitzero"`
	// IdentityFrontUrl represents identity_front_url. 需传入<a href="https://opendocs.alipay.com/apis/api_1*ant.merchant.expand.indirect.image.upload">ant.merchant.expand.indirect.image.upload</a> 接口上传图片后得到的 image_id。
	IdentityFrontUrl *string `json:"identity_front_url,omitzero"`
	// IdentityStartTime represents identity_start_time. 身份证的有效开始时间，国徽面信息
	IdentityStartTime *string `json:"identity_start_time,omitzero"`
	// IndustryInfo represents industry_info. 行业信息，例如所属集团、特殊标签
	IndustryInfo *[]MerchantShopIndustryInfo `json:"industry_info,omitzero"`
	// IndustryLicense represents industry_license. 门店资质信息
	IndustryLicense *[]ShopIndustryLicense `json:"industry_license,omitzero"`
	// IpRoleId represents ip_role_id. 当前门店归属的商户角色 ID，表示当前进件的门店的归属商户。对于直连开店场景，填写商户 PID。对于间连开店场景（线上、线下、直付通），商户SMID
	IpRoleId *string `json:"ip_role_id,omitzero"`
	// LegalCertNo represents legal_cert_no. 法人身份证号
	LegalCertNo *string `json:"legal_cert_no,omitzero"`
	// LegalName represents legal_name. 法人名称
	LegalName *string `json:"legal_name,omitzero"`
	// LicenseAuthLetterImage represents license_auth_letter_image. 营业执照授权函图片id，通过 ant.merchant.expand.indirect.image.upload 接口上传图片后得到的 image_id
	LicenseAuthLetterImage *string `json:"license_auth_letter_image,omitzero"`
	// Memo represents memo. 备注
	Memo *string `json:"memo,omitzero"`
	// NewBusinessTime represents new_business_time. 门店营业时间，支持一周7天，支持1天内多时间段
	NewBusinessTime *[]ComplexBusinessTime `json:"new_business_time,omitzero"`
	// OutDoorImages represents out_door_images. 需传入<a href="https://opendocs.alipay.com/apis/api_1/ant.merchant.expand.indirect.image.upload">ant.merchant.expand.indirect.image.upload</a> 接口上传图片后得到的 image_id。
	OutDoorImages *[]string `json:"out_door_images,omitzero"`
	// Qualifications represents qualifications. 行业特殊资质。
	Qualifications *[]IndustryQualificationInfo `json:"qualifications,omitzero"`
	// Scene represents scene. 场景，非加油站/酒店等特殊门店无需关注。
	Scene *string `json:"scene,omitzero"`
	// SettleAlipayLogonId represents settle_alipay_logon_id. 门店结算收款的支付宝账号
	SettleAlipayLogonId *string `json:"settle_alipay_logon_id,omitzero"`
	// ShopCategory represents shop_category. 新版门店类目标准二级类目code。类目标准及与原类目映射关系参见 <a href="https://gw.alipayobjects.com/os/bmw-prod/4b3f82df-e53e-4b84-bc41-fe025101e726.xlsx">支付宝门店类目-最新</a> 表格。
	ShopCategory *string `json:"shop_category,omitzero"`
	// ShopMainType represents shop_main_type. 主要针对医疗行业门店主体类型进件使用，不同门店认证主体需要相应的主体资质证书。 企业营业执照：ENTERPRISE; 事业单位法人证书：INST_RGST_CTF; 民办非企业单位登记证书：PRIVATE_NON_ENTERPRISE; 企业（澳门）：AOMEN_ENTERPRISE
	ShopMainType *string `json:"shop_main_type,omitzero"`
	// ShopName represents shop_name. 门店名称, 由商户定义在支付宝内的门店名称
	ShopName *string `json:"shop_name,omitzero"`
	// ShopType represents shop_type. 门店经营类型，01表示直营，02表示加盟
	ShopType *string `json:"shop_type,omitzero"`
	// StoreId represents store_id. 门店编号，表示该门店在该商户角色id(直连pid，间连smid)下，由商户自己定义的外部门店编号
	StoreId *string `json:"store_id,omitzero"`
}

// ExpandShopCloseRequest is the official V3 AntMerchantExpandShopCloseModel schema.
type ExpandShopCloseRequest struct {
	// ExtInfos represents ext_infos. 扩展信息列表。key值需要向对应行业的bd进行申请。 如果是针对神券团购门店的操作，该值必传，且无需向bd申请。
	ExtInfos *[]ShopExtInfo `json:"ext_infos,omitzero"`
	// IpRoleId represents ip_role_id. 商户角色id，表示将要开的店属于哪个商户角色。对于直连开店场景，填写商户pid；对于间连开店场景（线上、线下、直付通），填写商户smid。若未传入shop_id 则本参数与store_id均必填。
	IpRoleId *string `json:"ip_role_id,omitzero"`
	// ShopId represents shop_id. 支付宝侧蚂蚁门店 id。传入本参数后可不填 store_id 及 ip_role_id。
	ShopId *string `json:"shop_id,omitzero"`
	// StoreId represents store_id. 商户侧门店编号。表示该门店在该商户角色id(直连pid，间连smid)下，由商户自己定义的外部门店编号。若未传入 shop_id 则本参数与与ip_role_id均必填。
	StoreId *string `json:"store_id,omitzero"`
}

// ExpandAppRecommendAccountDeleteRequest contains parameters for ant.merchant.expand.apprecommend.account.delete.
type ExpandAppRecommendAccountDeleteRequest struct {
	// AppNo represents app_no. 待解绑小程序的app_id
	AppNo *string `query:"app_no"`
	// AccNo represents acc_no. 待解绑账号PID
	AccNo *string `query:"acc_no"`
}

// ExpandAppRecommendAccountQueryRequest contains parameters for ant.merchant.expand.apprecommend.account.query.
type ExpandAppRecommendAccountQueryRequest struct {
	// AppNo represents app_no. 已关联小程序的app_id
	AppNo *string `query:"app_no"`
	// PageSize represents page_size. 单页行数，不能超过100行，单位：行
	PageSize *int64 `query:"page_size"`
	// PageNumber represents page_number. 页码
	PageNumber *int64 `query:"page_number"`
}

// ExpandAppRecommendAvailableQueryRequest contains parameters for ant.merchant.expand.apprecommend.available.query.
type ExpandAppRecommendAvailableQueryRequest struct {
	// PageSize represents page_size. 单页行数，不能超过100行
	PageSize *int64 `query:"page_size"`
	// PageNumber represents page_number. 页码
	PageNumber *int64 `query:"page_number"`
}

// ExpandInfoQueryRequest contains parameters for ant.merchant.expand.info.query.
type ExpandInfoQueryRequest struct {
	// MerchantNo represents merchant_no. 商户号，目前是商户pid
	MerchantNo *string `query:"merchant_no"`
}

// ExpandItemDeleteRequest contains parameters for ant.merchant.expand.item.delete.
type ExpandItemDeleteRequest struct {
	// ItemId represents item_id. 商品ID（item_id不为空则以item_id作为删除KEY,item_id为空则以external_item_id作为删除KEY,两者不能同时为空）
	ItemId *string `query:"item_id"`
	// ExternalItemId represents external_item_id. 外部商品ID（item_id不为空则以item_id作为删除KEY,item_id为空则以external_item_id作为删除KEY,两者不能同时为空）
	ExternalItemId *string `query:"external_item_id"`
}

// ExpandItemOpenDeleteRequest contains parameters for ant.merchant.expand.item.open.delete.
type ExpandItemOpenDeleteRequest struct {
	// ItemId represents item_id. 商品ID
	ItemId *string `query:"item_id"`
}

// ExpandItemOpenQueryRequest contains parameters for ant.merchant.expand.item.open.query.
type ExpandItemOpenQueryRequest struct {
	// TargetId represents target_id. 商品归属主体ID 例：商品归属主体类型target_type为店铺，则商品归属主体ID为店铺ID（支付宝侧店铺ID）；归属主体类型target_type为小程序，则归属主体ID为小程序ID
	TargetId *string `query:"target_id"`
	// Scene represents scene. 场景码（具体值请参见产品文档）。
	Scene *string `query:"scene"`
	// TargetType represents target_type. 商品归属主体类型。枚举如下： 5：店铺。 8：小程序。
	TargetType *string `query:"target_type"`
	// Status represents status. 商品状态
	Status *string `query:"status"`
}

// ExpandItemQueryRequest contains parameters for ant.merchant.expand.item.query.
type ExpandItemQueryRequest struct {
	// TargetType represents target_type. 商品归属主体类型: 5: 店铺 4: 主站MID 3: 参与者 2: 角色 1: 联系人
	TargetType *string `query:"target_type"`
	// TargetId represents target_id. 商品归属主体id
	TargetId *string `query:"target_id"`
	// UndefinedCategory represents undefined_category. 未分类类目: true 的时候 front_category_id 必须不填（null）
	UndefinedCategory *bool `query:"undefined_category"`
	// FrontCategoryId represents front_category_id. 商品前台类目id: null 表示所有分类
	FrontCategoryId *string `query:"front_category_id"`
	// Status represents status. 商品状态：EFFECT、INVALID、PAUSE；不填返回所有状态商品
	Status *string `query:"status"`
}

// ExpandIndirectImageUploadRequest contains parameters for ant.merchant.expand.indirect.image.upload.
type ExpandIndirectImageUploadRequest struct {
	// Data is the data multipart field.
	Data *ExpandIndirectImageUpload `multipart:"data,json"`
	// ImageContent is the image_content multipart field.
	ImageContent *core.MultipartFile `multipart:"image_content,file"`
}

// ExpandMCCQueryRequest contains parameters for ant.merchant.expand.mcc.query.
type ExpandMCCQueryRequest struct {
	// MccCodeList represents mcc_code_list. 传入需要查询的code列表，以英文逗号分隔
	MccCodeList *string `query:"mcc_code_list"`
}

// ExpandOrderQueryRequest contains parameters for ant.merchant.expand.order.query.
type ExpandOrderQueryRequest struct {
	// OrderId represents order_id. 申请单 id。 1.蚂蚁门店管理场景：通过 <a href="https://opendocs.alipay.com/open/05afbc4a_ant.merchant.expand.shop.create?scene=common&pathHash=bf443b73"> ant.merchant.expand.shop.create</a>(蚂蚁门店创建)接口进件时返回的order_id。 2.直付通二级商户进件场景：通过<a href="https://opendocs.alipay.com/open/028
	OrderId string `path:"order_id"`
	// BizSource represents biz_source. 业务来源 团购门店：APLUS_SHOP
	BizSource *string `query:"biz_source"`
}

// ExpandShopQueryRequest contains parameters for ant.merchant.expand.shop.query.
type ExpandShopQueryRequest struct {
	// ShopId represents shop_id. 蚂蚁门店id
	ShopId *string `query:"shop_id"`
	// StoreId represents store_id. 门店编号，表示该门店在该商户角色id(直连pid，间连smid)下，由商户自己定义的外部门店编号
	StoreId *string `query:"store_id"`
	// IpRoleId represents ip_role_id. 商户角色id，表示将要开的店属于哪个商户角色。对于直连开店场景，填写商户pid；对于间连开店场景（线上、线下、直付通），填写商户smid
	IpRoleId *string `query:"ip_role_id"`
	// AddressVersion represents address_version. 行政区划版本，当前可传空值(取默认版本)、2022Q2、UPTODATE(取最新版本)，其中空值默认为：2020Q1版本（ address_version=''或null），想要查看版本是2022年2季度版本则传入:(address_version='2022Q2')，想要获取最新版本则传入:(address_version ='UPTODATE')
	AddressVersion *string `query:"address_version"`
	// NeedRecommend represents need_recommend. 门店不置信时，是否需要返回shop_recommend_info
	NeedRecommend *string `query:"need_recommend"`
	// NeedIndustryInfo represents need_industry_info. need_industry_info=0时不返回行业信息；need_industry_info=1时返回不需要审核的行业信息、审核通过的行业信息
	NeedIndustryInfo *string `query:"need_industry_info"`
	// NeedIndustryLicense represents need_industry_license. need_industry_license=0时不返回行业资质；need_industry_license=1时返回审核通过的行业资质
	NeedIndustryLicense *string `query:"need_industry_license"`
	// ShopType represents shop_type. 门店经营类型，01表示直营，02表示加盟
	ShopType *string `query:"shop_type"`
	// BizSource represents biz_source. 团购门店业务类型来源
	BizSource *string `query:"biz_source"`
}

// ExpandShopPageQueryRequest contains parameters for ant.merchant.expand.shop.page.query.
type ExpandShopPageQueryRequest struct {
	// IpRoleId represents ip_role_id. 商户角色id，表示将要开的店属于哪个商户角色。对于直连开店场景，填写商户pid；对于间连开店场景（线上、线下、直付通），填写商户smid。特别说明：IoT设备三绑定场景统一填写商户pid
	IpRoleId *string `query:"ip_role_id"`
	// PageNum represents page_num. 查询页数
	PageNum *int64 `query:"page_num"`
	// PageSize represents page_size. 每页查询大小，限制100以内
	PageSize *int64 `query:"page_size"`
	// AddressVersion represents address_version. 行政区划版本，当前可传空值(取默认版本)、2022Q2、UPTODATE(取最新版本)，其中空值默认为：2020Q1版本(address_version=''或null)，想要查看版本是2022年2季度版本则传入:(address_version='2022Q2')，想要获取最新版本则传入:(address_version ='UPTODATE')
	AddressVersion *string `query:"address_version"`
	// ProvinceCode represents province_code. 省份编码。 蚂蚁店铺请按照<a href="https://mdn.alipayobjects.com/huamei_sm7gq8/afts/file/A*blT9RqSR9_gAAAAAAAAAAAAADuKQAQ/2022%E8%9A%82%E8%9A%81%E9%87%91%E6%9C%8D%E5%9C%B0%E5%8C%BA%E7%A0%81.xlsx" target="_blank">蚂蚁店铺地区码</a> 表格中填写。 直付通商户请按照<a href="https://gw.alipayobjects.
	ProvinceCode *string `query:"province_code"`
	// CityCode represents city_code. 城市编码。 蚂蚁店铺请按照<a href="https://mdn.alipayobjects.com/huamei_sm7gq8/afts/file/A*blT9RqSR9_gAAAAAAAAAAAAADuKQAQ/2022%E8%9A%82%E8%9A%81%E9%87%91%E6%9C%8D%E5%9C%B0%E5%8C%BA%E7%A0%81.xlsx" target="_blank">蚂蚁店铺地区码</a> 表格中填写。 直付通商户请按照<a href="https://gw.alipayobjects.
	CityCode *string `query:"city_code"`
	// DistrictCode represents district_code. 区县编码。 蚂蚁店铺请按照<a href="https://mdn.alipayobjects.com/huamei_sm7gq8/afts/file/A*blT9RqSR9_gAAAAAAAAAAAAADuKQAQ/2022%E8%9A%82%E8%9A%81%E9%87%91%E6%9C%8D%E5%9C%B0%E5%8C%BA%E7%A0%81.xlsx" target="_blank">蚂蚁店铺地区码</a> 表格中填写。 直付通商户请按照<a href="https://gw.alipayobjects.
	DistrictCode *string `query:"district_code"`
}
