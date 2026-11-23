// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package sign

// Attachment 是支付宝官方 V3 OpenAPI 模型。
type Attachment struct {
	// AttachmentName 是流程附件名称
	AttachmentName *string `json:"attachment_name,omitempty"`
	// FileId 是流程附件id
	FileId *string `json:"file_id,omitempty"`
}

// AttachmentExplain 是支付宝官方 V3 OpenAPI 模型。
type AttachmentExplain struct {
	// AllowMoreUploads 是允许上传更多附件。开启后，除了“附件列表”中指定的内容，签署人可自由上传之多10张图片。默认值为false
	AllowMoreUploads *bool `json:"allow_more_uploads,omitempty"`
	// Description 是描述
	Description *string `json:"description,omitempty"`
	// Details 是附件列表
	Details *[]Detail `json:"details,omitempty"`
	// Title 是附件标题（collect_attachement值为true时必填
	Title *string `json:"title,omitempty"`
}

// ConfigInfo 是支付宝官方 V3 OpenAPI 模型。
type ConfigInfo struct {
	AttachmentExplain *AttachmentExplain `json:"attachment_explain,omitempty"`
	// CollectAttachement 是是否收集附件
	CollectAttachement *bool `json:"collect_attachement,omitempty"`
	// CollectCertTypes 是收集证件类型列表
	CollectCertTypes *[]string `json:"collect_cert_types,omitempty"`
	// CompanyNo 是公司编号-SF
	CompanyNo *string `json:"company_no,omitempty"`
	// ContractValidity 是文档过期时间
	ContractValidity *int64 `json:"contract_validity,omitempty"`
	// JumpUrl 是支付宝小程序跳转
	JumpUrl *string `json:"jump_url,omitempty"`
	// MerchantMiniSignUrl 是商户小程序签署地址（signModel字段值为1时 必填）
	MerchantMiniSignUrl *string `json:"merchant_mini_sign_url,omitempty"`
	// NoticeDeveloperUrl 是回调地址，签署的过程和签署完成都会回调。
	NoticeDeveloperUrl *string `json:"notice_developer_url,omitempty"`
	// PlatformOrderNo 是平台订单号
	PlatformOrderNo *string `json:"platform_order_no,omitempty"`
	// SerialNo 是物流单号
	SerialNo *string `json:"serial_no,omitempty"`
	// SignModel 是签署方式  0-e签宝小程序签署 1-商户小程序签署  默认0
	SignModel *int64 `json:"sign_model,omitempty"`
}

// Detail 是支付宝官方 V3 OpenAPI 模型。
type Detail struct {
	// Code 是英文描述（collectAttachement值为true时必填）	以下是系统固定code,对上传的图片做相应的校验。	身份证校验情况：	身份证正面照片：ESIGN_IDCARD	身份证背面照片：ESIGN_BACK_IDCARD
	Code *string `json:"code,omitempty"`
	// Description 是请上传手机正面照
	Description *string `json:"description,omitempty"`
	// Order 是顺序（collectAttachement值为true时必填）
	Order *int64 `json:"order,omitempty"`
}

// FillContent 是支付宝官方 V3 OpenAPI 模型。
type FillContent struct {
	// StructKey 是模板组件自定义key
	StructKey *string `json:"struct_key,omitempty"`
	// Value 是模板值,该值长度取决于配置模板时该字段的限制
	Value *string `json:"value,omitempty"`
}

// FlowSigner 是支付宝官方 V3 OpenAPI 模型。
type FlowSigner struct {
	// OrgThirdPartyUserId 是创建流程时指定的企业唯一标识，仅签署主体为企业时返回
	OrgThirdPartyUserId *string `json:"org_third_party_user_id,omitempty"`
	// SignStatus 是签署状态：0-待签, 1-未签, 2-已签 3-待审批 4-拒签
	SignStatus *int64 `json:"sign_status,omitempty"`
	// SignerAccountType 是签署主体类型：0-个人, 1-企业
	SignerAccountType *int64 `json:"signer_account_type,omitempty"`
	// ThirdPartyUserId 是创建流程时指定个人唯一标识
	ThirdPartyUserId *string `json:"third_party_user_id,omitempty"`
}

// SignField 是支付宝官方 V3 OpenAPI 模型。
type SignField struct {
	// AutoExecute 是平台自动签
	AutoExecute *string `json:"auto_execute,omitempty"`
	Signer      *Signer `json:"signer,omitempty"`
	// StructKey 是模板组件id
	StructKey *string `json:"struct_key,omitempty"`
}

// Signer 是支付宝官方 V3 OpenAPI 模型。
type Signer struct {
	// Email 是邮箱
	Email *string `json:"email,omitempty"`
	// IdNumber 是证件号
	IdNumber *string `json:"id_number,omitempty"`
	// IdType 是证件类型,详见个人证件类型说明（http://open.esign.cn/docs/xy/%E6%8E%A5%E5%8F%A3%E6%96%87%E6%A1%A3/%E9%99%84%E5%BD%95/%E4%B8%AA%E4%BA%BA%E8%AF%81%E4%BB%B6%E7%B1%BB%E5%9E%8B.html），默认CRED_PSN_CH_IDCARD (证件号不为空，则必填)
	IdType *string `json:"id_type,omitempty"`
	// Mobile 是手机号
	Mobile *string `json:"mobile,omitempty"`
	// Name 是姓名
	Name *string `json:"name,omitempty"`
	// UserId 是支付宝userId
	UserId *string `json:"user_id,omitempty"`
}

// TemplateInfo 是支付宝官方 V3 OpenAPI 模型。
type TemplateInfo struct {
	// FillContents 是模板填充项
	FillContents *[]FillContent `json:"fill_contents,omitempty"`
	// Name 是文件名称
	Name       *string    `json:"name,omitempty"`
	Signfields *SignField `json:"signfields,omitempty"`
	// TemplateId 是模板id
	TemplateId *string `json:"template_id,omitempty"`
}
