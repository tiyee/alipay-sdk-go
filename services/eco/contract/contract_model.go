// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package contract

// Attachment 是支付宝官方 V3 OpenAPI 模型。
type Attachment struct {
	// AttachmentName 是流程附件名称
	AttachmentName *string `json:"attachment_name,omitempty"`
	// FileId 是流程附件id
	FileId *string `json:"file_id,omitempty"`
}

// AttachmentExplainBean 是支付宝官方 V3 OpenAPI 模型。
type AttachmentExplainBean struct {
	// Description 是页面展示的任务描（collect_attachement值为true时必填）
	Description *string `json:"description,omitempty"`
	// Details 是附件列表，用于定义每个附件code的描述及页面展示顺序
	Details *[]DetailBean `json:"details,omitempty"`
	// Title 是页面展示的任务标题（collect_attachement值为true时必填）
	Title *string `json:"title,omitempty"`
}

// ConfigInfoBean 是支付宝官方 V3 OpenAPI 模型。
type ConfigInfoBean struct {
	AttachmentExplain *AttachmentExplainBean `json:"attachment_explain,omitempty"`
	// AutoArchive 是签署完成是否自动归档，默认true，如果false，则归档需要调用签署流程归档接口
	AutoArchive *bool `json:"auto_archive,omitempty"`
	// CollectAttachement 是是否收集附件（需签署人上传身份证或者其他文件的，需设置为true）
	CollectAttachement *bool `json:"collect_attachement,omitempty"`
	// CompanyNo 是物流公司编号（目前仅支持顺丰"SF"）	注：避免影响生产订单，联调及测试环境请传入SF_TEST。
	CompanyNo *string `json:"company_no,omitempty"`
	// ContractValidity 是合同过期时间：该参数设置的时间若到期，则会触发流程文件过期通知
	ContractValidity *int64 `json:"contract_validity,omitempty"`
	// JumpUrl 是签署完成后跳转商户小程序的地址（signModel字段值为2时 选填）
	JumpUrl *string `json:"jump_url,omitempty"`
	// MerchantMiniSignUrl 是商户小程序签署地址（signModel字段值为1时 必填）
	MerchantMiniSignUrl *string `json:"merchant_mini_sign_url,omitempty"`
	// NoticeDeveloperUrl 是https://esign.cn/callback
	NoticeDeveloperUrl *string `json:"notice_developer_url,omitempty"`
	// PlatformOrderNo 是平台订单号，可将商户订单与电子合同进行关联。用户通过支付宝官方小程序 合同管家 查看合同时，可快速进入商家的小程序及对应的订单详情页，为商家的小程序提供了流量入口，提高用户的活跃度和留存率；
	PlatformOrderNo *string `json:"platform_order_no,omitempty"`
	// SerialNo 是12或15位物流单号
	SerialNo *string `json:"serial_no,omitempty"`
	// SignModel 是快递扫码签收方式：	1-商户小程序：扫码后跳转商户小程序进行签收；	0-e签宝小程序：如商户没有支付宝小程序，可使用e签宝小程序完成签收。
	SignModel *int64 `json:"sign_model,omitempty"`
	// SpecifyAttachmentInfos 是指定签署人需要上传的附件列表，例如：A、B、C三方签署，A需上传附件1，B需指定附件2，C需上传附件1,2,3
	SpecifyAttachmentInfos *[]SpecifyAttachmentInfo `json:"specify_attachment_infos,omitempty"`
}

// DetailBean 是支付宝官方 V3 OpenAPI 模型。
type DetailBean struct {
	// Code 是自定义附件标识，后续下载附件使用此标识区分。
	Code *string `json:"code,omitempty"`
	// Description 是页面展示的附件中文描述（collectAttachement值为true时必填）
	Description *string `json:"description,omitempty"`
	// Order 是展示顺序（collectAttachement值为true时必填）
	Order *int64 `json:"order,omitempty"`
}

// FillContent 是支付宝官方 V3 OpenAPI 模型。
type FillContent struct {
	// StructKey 是模板组件自定义key
	StructKey *string `json:"struct_key,omitempty"`
	// Value 是模板值,该值长度取决于配置模板时该字段的限制
	Value *string `json:"value,omitempty"`
}

// Org 是支付宝官方 V3 OpenAPI 模型。
type Org struct {
	// OrgIdNumber 是企业证件号
	OrgIdNumber *string `json:"org_id_number,omitempty"`
	// OrgIdType 是企业证件类型，（证件号不为空时必填，详见企业证件类型）
	OrgIdType *string `json:"org_id_type,omitempty"`
	// OrgLegalIdNumber 是法定代表人证件号
	OrgLegalIdNumber *string `json:"org_legal_id_number,omitempty"`
	// OrgLegalIdType 是法定代表人证件类型
	OrgLegalIdType *string `json:"org_legal_id_type,omitempty"`
	// OrgLegalName 是法定代表人名称
	OrgLegalName *string `json:"org_legal_name,omitempty"`
	// OrgName 是企业名称
	OrgName *string `json:"org_name,omitempty"`
	// ThirdPartyUserId 是企业唯一标识：可传入平台机构id、企业证件号、企业邮箱等（个人用户与企业的唯一标识不可重复）
	ThirdPartyUserId *string `json:"third_party_user_id,omitempty"`
}

// SignFieldBean 是支付宝官方 V3 OpenAPI 模型。
type SignFieldBean struct {
	// SignFieldType 是签署类型：		platform-平台自动签：无需指定签署人。创建流程后，系统将自动盖上商户的默认企业印章。		person-个人签署：需要指定签署人个人信息。创建流程后，需通过签署插件完成签署。	org-企业签署：需要指定签署企业信息与经办人个人信息。创建流程后，需经办人代企业完成签署。（企业暂不支持插件签署，可通过获取签署地址接口获取企业签署地址）
	SignFieldType *string     `json:"sign_field_type,omitempty"`
	Signer        *SignerBean `json:"signer,omitempty"`
	// StructKey 是签署区key值
	StructKey *string `json:"struct_key,omitempty"`
}

// SignerBean 是支付宝官方 V3 OpenAPI 模型。
type SignerBean struct {
	// Email 是个人邮箱
	Email *string `json:"email,omitempty"`
	// IdNumber 是个人证件号
	IdNumber *string `json:"id_number,omitempty"`
	// IdType 是个人证件类型（证件号不为空时必填，详见个人证件类型说明 ）
	IdType *string `json:"id_type,omitempty"`
	// Mobile 是个人手机号
	Mobile *string `json:"mobile,omitempty"`
	// Name 是个人姓名
	Name *string `json:"name,omitempty"`
	Org  *Org    `json:"org,omitempty"`
	// ThirdPartyUserId 是个人唯一标识：可传入平台的个人用户id、支付宝userid、证件号、手机号、邮箱等。
	ThirdPartyUserId *string `json:"third_party_user_id,omitempty"`
}

// SpecifyAttachmentInfo 是支付宝官方 V3 OpenAPI 模型。
type SpecifyAttachmentInfo struct {
	// AllowMoreUploads 是是否允许上传更多附件，默认值为false。开启后，除了“collect_common_codes”中指定的附件，签署人可自由上传至多10张图片。
	AllowMoreUploads *bool `json:"allow_more_uploads,omitempty"`
	// AllowMoreUploadsMaxCount 是允许上传更多附件时，可上传的最大附件数，一次流程中累计的上传附件数至多30（含“附件列表”数量）
	AllowMoreUploadsMaxCount *int64 `json:"allow_more_uploads_max_count,omitempty"`
	// CollectCertTypes 是需要签署人上传的证件照片，可进行ORC识别，目前仅支持IDCARD
	CollectCertTypes *[]string `json:"collect_cert_types,omitempty"`
	// CollectCommonCodes 是需要签署人上传的其他附件列表，即attachment_explain.details中定义的code列表
	CollectCommonCodes *[]string `json:"collect_common_codes,omitempty"`
	// ThirdPartyUserId 是个人唯一标识，对应签署区中指定的签署人
	ThirdPartyUserId *string `json:"third_party_user_id,omitempty"`
}

// TemplateInfoBean 是支付宝官方 V3 OpenAPI 模型。
type TemplateInfoBean struct {
	// FillContents 是模板填充项，根据组件key值传入填写内容
	FillContents *[]FillContent `json:"fill_contents,omitempty"`
	// Name 是签署文件名称
	Name *string `json:"name,omitempty"`
	// Signfields 是签署区，根据签署区key值传入对应的签署人信息
	Signfields *[]SignFieldBean `json:"signfields,omitempty"`
	// TemplateId 是模板id，通过创建合同模板获取
	TemplateId *string `json:"template_id,omitempty"`
}
