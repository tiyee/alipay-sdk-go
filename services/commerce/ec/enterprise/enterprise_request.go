package enterprise

// EcEnterpriseAddressAddRequest 是支付宝官方 V3 OpenAPI 模型。
type EcEnterpriseAddressAddRequest struct {
	// AccountId 是通过企业码1.0接口签约的共同账户，和agreement_no搭配使用。
	AccountId *string `json:"account_id,omitempty"`
	// Address 是详细地址最长50个字符
	Address *string `json:"address,omitempty"`
	// AgreementNo 是可通过签约消息获取。配合共同账户id使用，当填写企业共同账户id时，此字段必填。
	AgreementNo *string `json:"agreement_no,omitempty"`
	// CityCode 是市(国家统一行政规划编码)
	CityCode *string `json:"city_code,omitempty"`
	// CityName 是市
	CityName *string `json:"city_name,omitempty"`
	// Community 是小区/楼宇
	Community *string `json:"community,omitempty"`
	// EnterpriseId 是通过企业码2.0签约接口签约，只填写企业id，无需填写共同账户id和授权签约协议号。
	EnterpriseId *string `json:"enterprise_id,omitempty"`
	// Latitude 是纬度
	Latitude *string `json:"latitude,omitempty"`
	// Longitude 是经度
	Longitude *string `json:"longitude,omitempty"`
	// Mark 是备注最长50个字符
	Mark *string `json:"mark,omitempty"`
	// PoiId 是高德地图poi
	PoiId *string `json:"poi_id,omitempty"`
}

// EcEnterpriseAddressModifyRequest 是支付宝官方 V3 OpenAPI 模型。
type EcEnterpriseAddressModifyRequest struct {
	// AccountId 是通过企业码1.0接口签约的共同账户，和agreement_no搭配使用。
	AccountId *string `json:"account_id,omitempty"`
	// Address 是详细地址最长50个字符
	Address *string `json:"address,omitempty"`
	// AddressId 是地址id
	AddressId *string `json:"address_id,omitempty"`
	// AgreementNo 是可通过签约消息获取。配合共同账户id使用，当填写企业共同账户id时，此字段必填。
	AgreementNo *string `json:"agreement_no,omitempty"`
	// CityCode 是市(国家统一行政规划编码)
	CityCode *string `json:"city_code,omitempty"`
	// CityName 是城市名称
	CityName *string `json:"city_name,omitempty"`
	// Community 是小区/楼宇
	Community *string `json:"community,omitempty"`
	// EnterpriseId 是通过企业码2.0签约接口签约，只填写企业id，无需填写共同账户id和授权签约协议号。
	EnterpriseId *string `json:"enterprise_id,omitempty"`
	// Latitude 是纬度
	Latitude *string `json:"latitude,omitempty"`
	// Longitude 是经度
	Longitude *string `json:"longitude,omitempty"`
	// Mark 是备注
	Mark *string `json:"mark,omitempty"`
	// PoiId 是高德地图poi
	PoiId *string `json:"poi_id,omitempty"`
	// Status 是状态(生效/失效)
	Status *string `json:"status,omitempty"`
}

// EcEnterpriseCreateRequest 是支付宝官方 V3 OpenAPI 模型。
type EcEnterpriseCreateRequest struct {
	// BizScene 是用于区分资金因公付协议、以及收银台文案等差异；
	BizScene *string `json:"biz_scene,omitempty"`
	// CreateFundAccount 是是否设置管理员支付宝为企业出资账户
	CreateFundAccount *bool `json:"create_fund_account,omitempty"`
	// CreateIotGroup 是是否创建企业人脸库，适用于对接团餐刷脸付、门禁刷脸通行等场景
	CreateIotGroup *bool `json:"create_iot_group,omitempty"`
	// EnterpriseAlias 是企业简称
	EnterpriseAlias *string `json:"enterprise_alias,omitempty"`
	// EnterpriseName 是企业名称
	EnterpriseName *string `json:"enterprise_name,omitempty"`
	// GroupAppId 是订购【一脸通行开通插件】的小程序appId
	GroupAppId *string `json:"group_app_id,omitempty"`
	// Identity 是管理员身份标识
	Identity *string `json:"identity,omitempty"`
	// IdentityName 是管理员姓名
	IdentityName *string `json:"identity_name,omitempty"`
	// IdentityOpenId 是管理员身份openId
	IdentityOpenId *string `json:"identity_open_id,omitempty"`
	// IdentityType 是管理员身份类型
	IdentityType *string `json:"identity_type,omitempty"`
	// OutBizNo 是外部业务号
	OutBizNo *string `json:"out_biz_no,omitempty"`
	// SignReturnUrl 是企业码签约后回跳地址。不传默认为空，签约后停留在当前页。
	SignReturnUrl *string `json:"sign_return_url,omitempty"`
	// SignTerminal 是管理员签约终端类型(字段已废弃，不需要服务商传入)
	SignTerminal *string `json:"sign_terminal,omitempty"`
}

// EcEnterpriseDeleteRequest 是支付宝官方 V3 OpenAPI 模型。
type EcEnterpriseDeleteRequest struct {
	// EnterpriseId 是企业id
	EnterpriseId *string `json:"enterprise_id,omitempty"`
}

// EcEnterpriseInfoModifyRequest 是支付宝官方 V3 OpenAPI 模型。
type EcEnterpriseInfoModifyRequest struct {
	// EnterpriseAlias 是企业简称
	EnterpriseAlias *string `json:"enterprise_alias,omitempty"`
	// EnterpriseId 是企业id
	EnterpriseId *string `json:"enterprise_id,omitempty"`
	// EnterpriseName 是企业名称
	EnterpriseName   *string                        `json:"enterprise_name,omitempty"`
	ReliableProfiles *ReliableEnterpriseProfilesDTO `json:"reliable_profiles,omitempty"`
}

// EcEnterpriseAddressQueryRequest is the query request model for alipay.commerce.ec.enterprise.address.query.
type EcEnterpriseAddressQueryRequest struct {
	// EnterpriseId 是通过企业码2.0签约接口签约，只填写企业id，无需填写共同账户id和授权签约协议号。
	EnterpriseId *string `json:"enterprise_id,omitempty" query:"enterprise_id,omitempty"`
	// AccountId 是通过企业码1.0接口签约的共同账户，和agreement_no搭配使用。
	AccountId *string `json:"account_id,omitempty" query:"account_id,omitempty"`
	// AgreementNo 是可通过签约消息获取。配合共同账户id使用，当填写企业共同账户id时，此字段必填。
	AgreementNo *string `json:"agreement_no,omitempty" query:"agreement_no,omitempty"`
	// AddressId 是地址id
	AddressId *string `json:"address_id,omitempty" query:"address_id,omitempty"`
	// CityCode 是市
	CityCode *string `json:"city_code,omitempty" query:"city_code,omitempty"`
	// PageNum 是页码从1开始
	PageNum *int64 `json:"page_num,omitempty" query:"page_num,required"`
	// PageSize 是每页数据
	PageSize *int64 `json:"page_size,omitempty" query:"page_size,required"`
}

// EcEnterpriseAgreementQueryRequest is the query request model for alipay.commerce.ec.enterprise.agreement.query.
type EcEnterpriseAgreementQueryRequest struct {
	// EnterpriseId 是企业id；enterprise_id与account_id至少传一个，如果都传了优先使用enterprise_id
	EnterpriseId *string `json:"enterprise_id,omitempty" query:"enterprise_id,omitempty"`
	// AccountId 是资金共同账号id；enterprise_id与account_id至少传一个，如果都传了优先使用enterprise_id(为兼容企业码1.0老接口保留的字段，已过时，新接客户请使用enterprise_id)
	AccountId *string `json:"account_id,omitempty" query:"account_id,omitempty"`
}

// EcEnterpriseInfoQueryRequest is the query request model for alipay.commerce.ec.enterprise.info.query.
type EcEnterpriseInfoQueryRequest struct {
	// EnterpriseId 是企业id,从企业注册等接口获取，和out_biz_no参数二选一，都传优先取企业id
	EnterpriseId *string `json:"enterprise_id,omitempty" query:"enterprise_id,omitempty"`
	// OutBizNo 是服务商生成的请求唯一流水号/业务幂等号，和enterprise_id参数二选一，都传优先取企业id
	OutBizNo *string `json:"out_biz_no,omitempty" query:"out_biz_no,omitempty"`
}

// EcEnterpriseUnsignRequest is the query request model for alipay.commerce.ec.enterprise.unsign.
type EcEnterpriseUnsignRequest struct {
	// EnterpriseId 是通过企业码2.0签约接口签约，只填写企业id，无需填写共同账户id和授权签约协议号。
	EnterpriseId *string `json:"enterprise_id,omitempty" query:"enterprise_id,omitempty"`
	// AccountId 是通过企业码1.0接口签约的共同账户，和agreement_no搭配使用。(为兼容企业码1.0老接口的参数，已过时，新接客户请传企业ID)
	AccountId *string `json:"account_id,omitempty" query:"account_id,omitempty"`
	// AgreementNo 是可通过签约消息获取。配合共同账户id使用，当填写企业共同账户id时，此字段必填。(为兼容企业码1.0老接口的参数，已过时，新接客户请传企业ID)
	AgreementNo *string `json:"agreement_no,omitempty" query:"agreement_no,omitempty"`
}
