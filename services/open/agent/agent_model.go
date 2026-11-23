package agent

// AgentMiniCreate 是支付宝官方 V3 OpenAPI 模型。
type AgentMiniCreate struct {
	// AppCategoryIds 是小程序类目，选填，可以支持最多三个类目，最少一个，参数格式：一级类目1_二级类目1;一级类目2_二级类目2。 类目可以通过以下接口查询：alipay.open.mini.category.query。
	AppCategoryIds *string `json:"app_category_ids,omitempty"`
	// AppDesc 是商家小程序描述信息，简要描述小程序主要功能（20-200个字），例：xx小程序提供了xx功能，主要解决了XX问题。。
	AppDesc *string `json:"app_desc,omitempty"`
	// AppEnglishName 是小程序英文名称，长度3~20个字符。
	AppEnglishName *string `json:"app_english_name,omitempty"`
	// AppName 是代商家创建的小程序应用名称。名称可以由中文、数字、英文及下划线组成，长度在3-20个字符之间，一个中文字等于2个字符，更多名称规则见：https://docs.alipay.com/mini/operation/name。
	AppName *string `json:"app_name,omitempty"`
	// AppSlogan 是代商家创建的小程序的简介，请用一句话简要描述小程序提供的服务；应用上架后一个自然月可修改5次（10~32个字符）。
	AppSlogan *string `json:"app_slogan,omitempty"`
	// BatchNo 是ISV 代商家操作事务编号，通过事务开启接口alipay.open.agent.create调用返回。。
	BatchNo *string `json:"batch_no,omitempty"`
	// MiniCategoryIds 是新小程序前台类目，格式为 第一个一级类目_第一个二级类目;第二个一级类目_第二个二级类目_第二个三级类目，详细类目可以通过 https://docs.open.alipay.com/api_49/alipay.open.mini.category.query接口查询mini_category_list。使用后不再读取app_category_ids值，老前台类目将废弃。
	MiniCategoryIds *string `json:"mini_category_ids,omitempty"`
	// ServiceEmail 是商家小程序客服邮箱 商家小程序客服电话和邮箱，可以二选一填写，但不能同时为空。
	ServiceEmail *string `json:"service_email,omitempty"`
	// ServicePhone 是商家小程序的客服电话，推荐填写 商家小程序客服电话和邮箱，可以二选一填写，但不能同时为空。
	ServicePhone *string `json:"service_phone,omitempty"`
}

// Contact 是支付宝官方 V3 OpenAPI 模型。
type Contact struct {
	// ContactEmail 是联系人邮箱。
	ContactEmail *string `json:"contact_email,omitempty"`
	// ContactMobile 是联系人手机号码。
	ContactMobile *string `json:"contact_mobile,omitempty"`
	// ContactName 是联系人名称。
	ContactName *string `json:"contact_name,omitempty"`
}

// ProductAgentStatusInfo 是支付宝官方 V3 OpenAPI 模型。
type ProductAgentStatusInfo struct {
	// ProductCode 是产品编码。
	ProductCode *string `json:"product_code,omitempty"`
	// ProductName 是产品名称。
	ProductName *string `json:"product_name,omitempty"`
	// RejectReason 是产品被驳回的原因。
	RejectReason *string `json:"reject_reason,omitempty"`
	// Status 是none:未签约，表示还没有签约该产品 valid:已生效，表示合约已经生效，不需要再签约了 restrictValid:受限生效，表示合约已经生效，但是资料不全，功能受限 audit:审核中，已经有合约在审核中，请等待审核完成 waitConfirm:待商户确认协议，合约已经审核通过，需要商户确认后合约才生效 auditReject:审核未通过 invalid:合约失效，曾经签过合约，但已经失效了，可以重新发起签约 restrictInvalid:受限失效，受限合约失效了。
	Status *string `json:"status,omitempty"`
}

// SignRestrictInfo 是支付宝官方 V3 OpenAPI 模型。
type SignRestrictInfo struct {
	// RestrictProduct 是产品码，支付宝定义的产品码。
	RestrictProduct *string `json:"restrict_product,omitempty"`
	// RestrictReason 是受限的原因。
	RestrictReason *string `json:"restrict_reason,omitempty"`
	// RestrictReasonCode 是受限原因编码。
	RestrictReasonCode *string `json:"restrict_reason_code,omitempty"`
	// RestrictVersion 是受限版本，02-受限版 04-极速版。
	RestrictVersion *string `json:"restrict_version,omitempty"`
}
