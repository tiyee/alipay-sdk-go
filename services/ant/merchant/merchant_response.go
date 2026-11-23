// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package merchant

// ExpandAppRecommendAccountCreateResponse is the official V3 AntMerchantExpandApprecommendAccountCreateResponseModel schema.
type ExpandAppRecommendAccountCreateResponse struct {
}

// ExpandAppRecommendAccountDeleteResponse is the official V3 AntMerchantExpandApprecommendAccountDeleteResponseModel schema.
type ExpandAppRecommendAccountDeleteResponse struct {
}

// ExpandAppRecommendAccountQueryResponse is the official V3 AntMerchantExpandApprecommendAccountQueryResponseModel schema.
type ExpandAppRecommendAccountQueryResponse struct {
	// List represents list. 账号列表
	List *[]string `json:"list,omitzero"`
	// PageNumber represents page_number. 当前页码
	PageNumber *int64 `json:"page_number,omitzero"`
	// PageSize represents page_size. 单页行数，单位：行
	PageSize *int64 `json:"page_size,omitzero"`
	// TotalPages represents total_pages. 总页数，单位：页
	TotalPages *int64 `json:"total_pages,omitzero"`
	// TotalSize represents total_size. 总行数， 单位：行
	TotalSize *string `json:"total_size,omitzero"`
}

// ExpandAppRecommendAvailableQueryResponse is the official V3 AntMerchantExpandApprecommendAvailableQueryResponseModel schema.
type ExpandAppRecommendAvailableQueryResponse struct {
	// List represents list. 总页数
	List *[]string `json:"list,omitzero"`
	// PageNumber represents page_number. 当前页码
	PageNumber *int64 `json:"page_number,omitzero"`
	// PageSize represents page_size. 单页行数
	PageSize *int64 `json:"page_size,omitzero"`
	// TotalPages represents total_pages. 总页数
	TotalPages *int64 `json:"total_pages,omitzero"`
	// TotalSize represents total_size. 总行数
	TotalSize *int64 `json:"total_size,omitzero"`
}

// ExpandInfoQueryResponse is the official V3 AntMerchantExpandInfoQueryResponseModel schema.
type ExpandInfoQueryResponse struct {
	// MerchantQueryResult represents merchant_query_result.
	MerchantQueryResult *MerchantQueryResult `json:"merchant_query_result,omitzero"`
}

// ExpandItemCreateResponse is the official V3 AntMerchantExpandItemCreateResponseModel schema.
type ExpandItemCreateResponse struct {
	// ItemId represents item_id. 商品id
	ItemId *string `json:"item_id,omitzero"`
}

// ExpandItemDeleteResponse is the official V3 AntMerchantExpandItemDeleteResponseModel schema.
type ExpandItemDeleteResponse struct {
	// ItemId represents item_id. 商品id
	ItemId *string `json:"item_id,omitzero"`
}

// ExpandItemModifyResponse is the official V3 AntMerchantExpandItemModifyResponseModel schema.
type ExpandItemModifyResponse struct {
	// ItemId represents item_id. 商品id
	ItemId *string `json:"item_id,omitzero"`
}

// ExpandItemOpenBatchQueryResponse is the official V3 AntMerchantExpandItemOpenBatchqueryResponseModel schema.
type ExpandItemOpenBatchQueryResponse struct {
	// ItemList represents item_list. 商品列表。
	ItemList *[]CmItemInfo `json:"item_list,omitzero"`
}

// ExpandItemOpenCreateResponse is the official V3 AntMerchantExpandItemOpenCreateResponseModel schema.
type ExpandItemOpenCreateResponse struct {
	// ItemId represents item_id. 商品ID
	ItemId *string `json:"item_id,omitzero"`
}

// ExpandItemOpenDeleteResponse is the official V3 AntMerchantExpandItemOpenDeleteResponseModel schema.
type ExpandItemOpenDeleteResponse struct {
}

// ExpandItemOpenModifyResponse is the official V3 AntMerchantExpandItemOpenModifyResponseModel schema.
type ExpandItemOpenModifyResponse struct {
}

// ExpandItemOpenQueryResponse is the official V3 AntMerchantExpandItemOpenQueryResponseModel schema.
type ExpandItemOpenQueryResponse struct {
	// ItemList represents item_list. 商品列表
	ItemList *[]CmItemInfo `json:"item_list,omitzero"`
}

// ExpandItemQueryResponse is the official V3 AntMerchantExpandItemQueryResponseModel schema.
type ExpandItemQueryResponse struct {
	// ItemList represents item_list. 商品列表
	ItemList *[]Item `json:"item_list,omitzero"`
}

// ExpandItemStatusModifyResponse is the official V3 AntMerchantExpandItemStatusModifyResponseModel schema.
type ExpandItemStatusModifyResponse struct {
	// ItemId represents item_id. 商品id
	ItemId *string `json:"item_id,omitzero"`
}

// ExpandShopConsultResponse is the official V3 AntMerchantExpandShopConsultResponseModel schema.
type ExpandShopConsultResponse struct {
	// AccountAudit represents account_audit. 资金账户是否审核。当调用没有报错时，本值才有参考意义
	AccountAudit *bool `json:"account_audit,omitzero"`
	// OrderId represents order_id. 咨询申请单id
	OrderId *string `json:"order_id,omitzero"`
	// RiskAudit represents risk_audit. 是否风控审核。当调用没有报错时，本值才有参考意义
	RiskAudit *bool `json:"risk_audit,omitzero"`
}

// ExpandShopReceiptAccountSaveResponse is the official V3 AntMerchantExpandShopReceiptaccountSaveResponseModel schema.
type ExpandShopReceiptAccountSaveResponse struct {
}

// ExpandIndirectImageUploadResponse is the official V3 AntMerchantExpandIndirectImageUploadResponseModel schema.
type ExpandIndirectImageUploadResponse struct {
	// ImageId represents image_id. 图片在文件存储平台的标识
	ImageId *string `json:"image_id,omitzero"`
}

// ExpandMCCQueryResponse is the official V3 AntMerchantExpandMccQueryResponseModel schema.
type ExpandMCCQueryResponse struct {
	// MccInfoList represents mcc_info_list. mcc查询信息结果列表
	MccInfoList *[]MCCQueryInfo `json:"mcc_info_list,omitzero"`
}

// ExpandOrderQueryResponse is the official V3 AntMerchantExpandOrderQueryResponseModel schema.
type ExpandOrderQueryResponse struct {
	// ApplyTime represents apply_time. 申请单创建时间
	ApplyTime *string `json:"apply_time,omitzero"`
	// ExtInfo represents ext_info. 返回申请单相关参数。当前返回内容有cardAliasNo，smid
	ExtInfo *string `json:"ext_info,omitzero"`
	// IpRoleId represents ip_role_id. 商户角色id。间连商户或平台商二级商户业务中，本字段表示smid。
	IpRoleId *[]string `json:"ip_role_id,omitzero"`
	// MerchantName represents merchant_name. 商户名称
	MerchantName *string `json:"merchant_name,omitzero"`
	// Status represents status. 申请单状态。
	Status *string `json:"status,omitzero"`
}

// ExpandShopQueryResponse is the official V3 AntMerchantExpandShopQueryResponseModel schema.
type ExpandShopQueryResponse struct {
	// AlipayPoiid represents alipay_poiid. 门店地址库 ID，按照一定的平台规则识别出的线下真实存在、真实经营的蚂蚁门店地址库 ID，将作用于服务商的返佣激励、商品/券等权益的公域分发。如平台未返回alipay_poiid，请在确认门店信息真实有效后，稍后再进行查询。
	AlipayPoiid *string `json:"alipay_poiid,omitzero"`
	// BrandId represents brand_id. 品牌id
	BrandId *string `json:"brand_id,omitzero"`
	// BusinessAddress represents business_address.
	BusinessAddress *AddressInfo `json:"business_address,omitzero"`
	// BusinessTime represents business_time. 门店经营时间。
	BusinessTime *[]ShopBusinessTime `json:"business_time,omitzero"`
	// CardAccountNo represents card_account_no. 银行卡号
	CardAccountNo *string `json:"card_account_no,omitzero"`
	// CertImage represents cert_image. 营业执照图片url。返回值为一个有访问时限的链接
	CertImage *string `json:"cert_image,omitzero"`
	// CertName represents cert_name. 营业执照名称，值为营业执照或统一社会信用代码证上的名称。
	CertName *string `json:"cert_name,omitzero"`
	// CertNo represents cert_no. 证件号码
	CertNo *string `json:"cert_no,omitzero"`
	// CertType represents cert_type. 证件类型，取值范围：201：营业执照；2011:多证合一(统一社会信用代码)。
	CertType *string `json:"cert_type,omitzero"`
	// ContactInfos represents contact_infos. 联系人信息。
	ContactInfos *[]ContactInfo `json:"contact_infos,omitzero"`
	// ContactMobile represents contact_mobile. 门店联系手机
	ContactMobile *string `json:"contact_mobile,omitzero"`
	// ContactPhone represents contact_phone. 门店的联系固话
	ContactPhone *string `json:"contact_phone,omitzero"`
	// ExtInfos represents ext_infos. 扩展信息列表。key值需要向对应行业的bd进行申请。
	ExtInfos *[]ShopExtInfo `json:"ext_infos,omitzero"`
	// IndustryInfo represents industry_info. 查询门店详情时，如果需要同时返回行业信息，需要设置入参中的need_industry_info=1
	IndustryInfo *[]MerchantShopIndustryInfo `json:"industry_info,omitzero"`
	// IndustryLicense represents industry_license. 查询门店详情时，如果需要同时返回行业资质，需要设置入参中的need_industry_license=1
	IndustryLicense *[]ShopIndustryLicense `json:"industry_license,omitzero"`
	// IpRoleId represents ip_role_id. 商户角色id，门店属于哪个商户角色。对于直连开店场景，为商户pid；对于间连开店场景（线上、线下、直付通），为商户smid。
	IpRoleId *string `json:"ip_role_id,omitzero"`
	// LegalCertNo represents legal_cert_no. 法人身份证号。
	LegalCertNo *string `json:"legal_cert_no,omitzero"`
	// LegalName represents legal_name. 法人名称。
	LegalName *string `json:"legal_name,omitzero"`
	// LicenseAuthLetterImage represents license_auth_letter_image. 营业执照授权函。返回值为一个有访问时限的链接
	LicenseAuthLetterImage *string `json:"license_auth_letter_image,omitzero"`
	// Memo represents memo. 备注
	Memo *string `json:"memo,omitzero"`
	// NewBusinessTime represents new_business_time. 门店营业时间，支持一周7天，支持1天内多时间段
	NewBusinessTime *[]ComplexBusinessTime `json:"new_business_time,omitzero"`
	// NewShopCategory represents new_shop_category. 新版门店类目标准二级类目code
	NewShopCategory *string `json:"new_shop_category,omitzero"`
	// OutDoorImages represents out_door_images. 门头照，返回值为一个有访问时限的链接
	OutDoorImages *[]string `json:"out_door_images,omitzero"`
	// Qualifications represents qualifications. 行业特殊资质。
	Qualifications *[]IndustryQualificationInfo `json:"qualifications,omitzero"`
	// Scene represents scene. 场景
	Scene *string `json:"scene,omitzero"`
	// SettleAlipayLogonId represents settle_alipay_logon_id. 结算支付宝账号的登录号
	SettleAlipayLogonId *string `json:"settle_alipay_logon_id,omitzero"`
	// ShopCategory represents shop_category. 门店类目，取值参见文件中的三级门店类目<a herf="https://mif-pub.alipayobjects.com/ShopCategory.xlsx"></a>
	ShopCategory *string `json:"shop_category,omitzero"`
	// ShopId represents shop_id. 支付宝侧蚂蚁门店 id
	ShopId *string `json:"shop_id,omitzero"`
	// ShopInfoStatus represents shop_info_status. 当前名称、地址、经纬度信息准确一致，可用于数字化经营场景消费(如商品、券、消费圈等场域的分发)，不影响门店支付结算
	ShopInfoStatus *string `json:"shop_info_status,omitzero"`
	// ShopMainType represents shop_main_type. 主要针对医疗行业门店主体类型进件使用，不同门店认证主体需要相应的主体资质证书。 企业营业执照：ENTERPRISE; 事业单位法人证书：INST_RGST_CTF; 民办非企业单位登记证书：PRIVATE_NON_ENTERPRISE; 企业（澳门）：AOMEN_ENTERPRISE
	ShopMainType *string `json:"shop_main_type,omitzero"`
	// ShopName represents shop_name. 门店名称。
	ShopName *string `json:"shop_name,omitzero"`
	// ShopRecommendInfo represents shop_recommend_info.
	ShopRecommendInfo *ShopRecommendInfo `json:"shop_recommend_info,omitzero"`
	// ShopType represents shop_type. 门店经营类型
	ShopType *string `json:"shop_type,omitzero"`
	// StoreId represents store_id. 商户侧门店编号。表示该门店在该商户角色id(直连pid，间连smid)下，由商户自己定义的外部门店编号。若未传入 shop_id 则本参数与与ip_role_id均必填。
	StoreId *string `json:"store_id,omitzero"`
}

// ExpandShopModifyResponse is the official V3 AntMerchantExpandShopModifyResponseModel schema.
type ExpandShopModifyResponse struct {
	// OrderId represents order_id. 申请单id
	OrderId *string `json:"order_id,omitzero"`
}

// ExpandShopCreateResponse is the official V3 AntMerchantExpandShopCreateResponseModel schema.
type ExpandShopCreateResponse struct {
	// ExistShopId represents exist_shop_id. 触发幂等时返回已存在的门店shopId
	ExistShopId *string `json:"exist_shop_id,omitzero"`
	// OrderId represents order_id. 申请单id
	OrderId *string `json:"order_id,omitzero"`
}

// ExpandShopCloseResponse is the official V3 AntMerchantExpandShopCloseResponseModel schema.
type ExpandShopCloseResponse struct {
}

// ExpandShopPageQueryResponse is the official V3 AntMerchantExpandShopPageQueryResponseModel schema.
type ExpandShopPageQueryResponse struct {
	// AlipayPoiid represents alipay_poiid. 门店地址库 ID，按照一定的平台规则识别出的线下真实存在、真实经营的蚂蚁门店地址库 ID，将作用于服务商的返佣激励、商品/券等权益的公域分发。如平台未返回alipay_poiid，请在确认门店信息真实有效后，稍后再进行查询
	AlipayPoiid *string `json:"alipay_poiid,omitzero"`
	// ShopInfos represents shop_infos. 门店详情
	ShopInfos *[]ShopQueryOpenAPIVO `json:"shop_infos,omitzero"`
	// TotalPages represents total_pages. 总页数
	TotalPages *int64 `json:"total_pages,omitzero"`
}
