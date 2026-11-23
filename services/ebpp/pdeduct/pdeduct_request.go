package pdeduct

// PdeductAsyncPayRequest 是支付宝官方 V3 OpenAPI 模型。
type PdeductAsyncPayRequest struct {
	// AgentChannel 是分配给外部机构发起扣款时的渠道码。朗新为LANGXIN
	AgentChannel *string `json:"agent_channel,omitempty"`
	// AgentCode 是二级渠道码，预留字段
	AgentCode *string `json:"agent_code,omitempty"`
	// AgreementId 是支付宝代扣协议Id
	AgreementId *string `json:"agreement_id,omitempty"`
	// BillDate 是账期
	BillDate *string `json:"bill_date,omitempty"`
	// BillKey 是户号
	BillKey *string `json:"bill_key,omitempty"`
	// ExtendField 是扩展字段
	ExtendField *string `json:"extend_field,omitempty"`
	// FineAmount 是滞纳金
	FineAmount *string `json:"fine_amount,omitempty"`
	// Memo 是备注信息
	Memo *string `json:"memo,omitempty"`
	// OpenId 是用户UserId在应用AppId下的唯一用户标识
	OpenId *string `json:"open_id,omitempty"`
	// OutOrderNo 是商户外部业务流水号
	OutOrderNo *string `json:"out_order_no,omitempty"`
	// PayAmount 是扣款金额，支付总金额，包含滞纳金
	PayAmount *string `json:"pay_amount,omitempty"`
	// Pid 是商户partnerId
	Pid *string `json:"pid,omitempty"`
	// UserId 是用户ID
	UserId *string `json:"user_id,omitempty"`
}

// PdeductPayRequest 是支付宝官方 V3 OpenAPI 模型。
type PdeductPayRequest struct {
	// AgentChannel 是渠道码，如用户通过机构通过服务窗进来签约则是PUBLICFORM，此值可随意传，只要不空就行
	AgentChannel *string `json:"agent_channel,omitempty"`
	// AgentCode 是二级渠道码，预留字段
	AgentCode *string `json:"agent_code,omitempty"`
	// AgreementId 是支付宝代扣协议Id
	AgreementId *string `json:"agreement_id,omitempty"`
	// BillDate 是账期
	BillDate *string `json:"bill_date,omitempty"`
	// BillKey 是户号，缴费单位用于标识每一户的唯一性的
	BillKey *string `json:"bill_key,omitempty"`
	// ExtendField 是扩展参数。必须以key value形式定义，	转为json为格式：{"key1":"value1","key2":"value2",	"key3":"value3","key4":"value4"}	 后端会直接转换为MAP对象，转换异常会报参数格式错误。	instUniqueNo,机构唯一流水号
	ExtendField *string `json:"extend_field,omitempty"`
	// FineAmount 是滞纳金
	FineAmount *string `json:"fine_amount,omitempty"`
	// Memo 是备注信息
	Memo *string `json:"memo,omitempty"`
	// OpenId 是用户UserId在应用AppId下的唯一用户标识
	OpenId *string `json:"open_id,omitempty"`
	// OutOrderNo 是商户外部业务流水号
	OutOrderNo *string `json:"out_order_no,omitempty"`
	// PayAmount 是扣款金额，支付总金额，包含滞纳金
	PayAmount *string `json:"pay_amount,omitempty"`
	// Pid 是商户PartnerId
	Pid *string `json:"pid,omitempty"`
	// UserId 是用户ID
	UserId *string `json:"user_id,omitempty"`
}

// PdeductBillPayStatusRequest is the query request model for alipay.ebpp.pdeduct.bill.pay.status.
type PdeductBillPayStatusRequest struct {
	// AgreementId 是支付宝用户ID
	AgreementId *string `json:"agreement_id,omitempty" query:"agreement_id,omitempty"`
	// OpenId 是用户UserId在应用AppId下的唯一用户标识
	OpenId *string `json:"open_id,omitempty" query:"open_id,omitempty"`
	// OutOrderNo 是商户代扣业务流水
	OutOrderNo *string `json:"out_order_no,omitempty" query:"out_order_no,omitempty"`
}
