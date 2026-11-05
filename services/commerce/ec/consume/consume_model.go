package consume

// EcConsumeInfo 是支付宝官方 V3 OpenAPI 模型。
type EcConsumeInfo struct {
	// AccountId 是共同账户ID
	AccountId *string `json:"account_id,omitempty"`
	// AgreementPeerPayerId 是协议出资支付宝账号
	AgreementPeerPayerId *string `json:"agreement_peer_payer_id,omitempty"`
	// BenefitAmount 是员工优惠金额，单位：元
	BenefitAmount *string `json:"benefit_amount,omitempty"`
	// BizOutNo 是外部交易流水号
	BizOutNo *string `json:"biz_out_no,omitempty"`
	// CategoryName 是账单分类名称，如餐饮等
	CategoryName *string `json:"category_name,omitempty"`
	// ConsumeAmount 是账单金额，单位：元，不包含营销资产
	ConsumeAmount *string `json:"consume_amount,omitempty"`
	// ConsumeBizType 是账单业务类型	-因公支付：EC_PAY	-因公收款：EC_CLLCT
	ConsumeBizType *string `json:"consume_biz_type,omitempty"`
	// ConsumeCategory 是员工消费记账分类：	"cater": "餐饮"	"purchase": "采购"	"trip": "出行"	"train": "培训"	"entertainment": "娱乐"	"hotel": "酒店"	"pay": "缴费"	"rent": "房租"	"other": "其他"
	ConsumeCategory *string `json:"consume_category,omitempty"`
	// ConsumeFeeWithDiscount 是订单原价，单位：元，包含营销资产
	ConsumeFeeWithDiscount *string `json:"consume_fee_with_discount,omitempty"`
	// ConsumeMemo 是员工消费记账备注
	ConsumeMemo *string `json:"consume_memo,omitempty"`
	// ConsumeTitle 是商户收款时传的商品备注说明，在账单内进行表达
	ConsumeTitle *string `json:"consume_title,omitempty"`
	// ConsumeType 是账单类型	-消费账单：CONSUME	-退款账单：REFUND	-转账账单：TRANSFER
	ConsumeType *string `json:"consume_type,omitempty"`
	// EmployeeId 是员工账号ID
	EmployeeId *string `json:"employee_id,omitempty"`
	// EnterpriseId 是企业ID
	EnterpriseId *string `json:"enterprise_id,omitempty"`
	// ExpenseRuleGroupId 是使用规则ID
	ExpenseRuleGroupId *string `json:"expense_rule_group_id,omitempty"`
	// ExpenseSceneCode 是费用场景/一级场景
	ExpenseSceneCode *string `json:"expense_scene_code,omitempty"`
	// ExpenseType 是费用类型
	ExpenseType *string `json:"expense_type,omitempty"`
	// ExpenseTypeSubCategory 是费用类型子类目/二级场景
	ExpenseTypeSubCategory *string `json:"expense_type_sub_category,omitempty"`
	// ExtInfos 是账单扩展信息，JsonMap格式，以下信息可组合返回，"SUMMARY_INFO"表示"账单汇总信息"，"MERCHANT_EXTEND_INFO"表示"商户信息"，"TERMINAL_INFO"表示"终端设备信息"，"BENEFIT_INFO"表示"员工优惠信息"，"CONSUMPTION_LOCATION"表示"消费上报位置"，"THIRD_PARTY_PAYMENT_INFO"表示"代收款详情"，"REIMBURSEMENT_INFO"表示"审批单信息"
	ExtInfos *string `json:"ext_infos,omitempty"`
	// FundBizType 是出资模式	个人出资：PERSONAL	企业出资：ENTERPRISE	三方垫资合作伙伴出资：TP	记账：ACCOUNTING
	FundBizType *string `json:"fund_biz_type,omitempty"`
	// GmtBizCreate 是账单创建时间，格式：yyyy-MM-dd HH:mm:ss
	GmtBizCreate *string `json:"gmt_biz_create,omitempty"`
	// GmtReceivePay 是账单支付时间，格式：yyyy-MM-dd HH:mm:ss
	GmtReceivePay *string `json:"gmt_receive_pay,omitempty"`
	// InvoiceOpenMode 是开票模式	企业汇总开：ENTERPRISE_AUTO_BATCH
	InvoiceOpenMode *string `json:"invoice_open_mode,omitempty"`
	// MerchantId 是商户ID
	MerchantId *string `json:"merchant_id,omitempty"`
	// MerchantName 是商户名称
	MerchantName *string `json:"merchant_name,omitempty"`
	// OpenId 是员工支付宝UID
	OpenId *string `json:"open_id,omitempty"`
	// OppositeFullName 是收款方全称（如果是个人会显示脱敏后的名称）
	OppositeFullName *string `json:"opposite_full_name,omitempty"`
	// OrderCompleteLabel 是订单完结标识，0标识“未完结”，1标识“已完结”
	OrderCompleteLabel *string `json:"order_complete_label,omitempty"`
	// OrderCompleteTime 是2022-01-01 01:01:02	订单完结时间
	OrderCompleteTime *string `json:"order_complete_time,omitempty"`
	// PayNo 是交易流水号
	PayNo *string `json:"pay_no,omitempty"`
	// PayerCardNo 是付款方卡号
	PayerCardNo *string `json:"payer_card_no,omitempty"`
	// PayerLogonId 是付款方登录账号
	PayerLogonId *string `json:"payer_logon_id,omitempty"`
	// PayerName 是付款方名称
	PayerName *string `json:"payer_name,omitempty"`
	// PeerPayAmount 是企业代付金额，单位：元
	PeerPayAmount *string `json:"peer_pay_amount,omitempty"`
	// PeerPayerCardName 是实际出资人名称
	PeerPayerCardName *string `json:"peer_payer_card_name,omitempty"`
	// PeerPayerCardNo 是实际出资支付宝账号
	PeerPayerCardNo *string `json:"peer_payer_card_no,omitempty"`
	// PeerRefundAmount 是消费账单企业代付部分退款金额，单位：元。退款账单该值无意义，值为0。
	PeerRefundAmount *string `json:"peer_refund_amount,omitempty"`
	// PeerRefundStatus 是该字段只对因公付交易生效，默认为INIT。当消费账单有退款，该值会变为REFUND_PART或REFUND_FULL；退款账单该值无意义，为初始值INIT。	未退款：INIT	部分退款：REFUND_PART，	全额退款：REFUND_FULL
	PeerRefundStatus *string `json:"peer_refund_status,omitempty"`
	// RefundAmount 是消费账单退款金额，单位：元。退款账单该值无意义，值为0。
	RefundAmount *string `json:"refund_amount,omitempty"`
	// RefundStatus 是当消费账单有退款，该值会变为REFUND_PART或REFUND_FULL；退款账单该值无意义，为初始值INIT。
	RefundStatus *string `json:"refund_status,omitempty"`
	// RelatedPayNo 是退款账单关联的消费账单交易流水号，退款账单才有值
	RelatedPayNo *string `json:"related_pay_no,omitempty"`
	// SceneCode 是账单场景	TAKE_AWAY：外卖	METRO：地铁	OTHER：其他
	SceneCode *string `json:"scene_code,omitempty"`
	// SellerId 是卖家ID
	SellerId *string `json:"seller_id,omitempty"`
	// ShopId 是门店ID
	ShopId *string `json:"shop_id,omitempty"`
	// StoreId 是外部门店ID
	StoreId *string `json:"store_id,omitempty"`
	// SummaryApplyId 是汇总批次ID
	SummaryApplyId *string `json:"summary_apply_id,omitempty"`
	// UserId 是员工支付宝UID
	UserId *string `json:"user_id,omitempty"`
}

// EcOrderInfo 是支付宝官方 V3 OpenAPI 模型。
type EcOrderInfo struct {
	OrderInfo *EcOrderItem `json:"order_info,omitempty"`
	// SubOrderList 是子订单详情列表
	SubOrderList *[]EcOrderItem `json:"sub_order_list,omitempty"`
}

// EcOrderItem 是支付宝官方 V3 OpenAPI 模型。
type EcOrderItem struct {
	// AccountId 是共同账户ID
	AccountId *string `json:"account_id,omitempty"`
	// BizOutNo 是外部业务号
	BizOutNo *string `json:"biz_out_no,omitempty"`
	// EmployeeId 是员工ID
	EmployeeId *string `json:"employee_id,omitempty"`
	// EnterpriseId 是企业ID
	EnterpriseId *string `json:"enterprise_id,omitempty"`
	// GmtCreate 是创建时间
	GmtCreate *string `json:"gmt_create,omitempty"`
	// GmtModified 是修改时间
	GmtModified *string `json:"gmt_modified,omitempty"`
	// OpenId 是用户支付宝UID
	OpenId *string `json:"open_id,omitempty"`
	// OrderContent 是订单内容，JSON格式。
	OrderContent *string `json:"order_content,omitempty"`
	// OrderId 是订单ID
	OrderId *string `json:"order_id,omitempty"`
	// OrderType 是订单类型	METRO：地铁	TAKE_AWAY：外卖	OTHER：其他	REACH_SHOP：到店
	OrderType *string `json:"order_type,omitempty"`
	// PartnerId 是合作伙伴ID
	PartnerId *string `json:"partner_id,omitempty"`
	// PayNo 是关联支付宝交易号
	PayNo *string `json:"pay_no,omitempty"`
	// UserId 是用户支付宝UID
	UserId *string `json:"user_id,omitempty"`
}

// EcVoucherInfo 是支付宝官方 V3 OpenAPI 模型。
type EcVoucherInfo struct {
	// AccountId 是共同账户ID
	AccountId *string `json:"account_id,omitempty"`
	// EmployeeId 是员工ID，汇总发票该字段无效
	EmployeeId *string `json:"employee_id,omitempty"`
	// EnterpriseId 是企业ID
	EnterpriseId                 *string                       `json:"enterprise_id,omitempty"`
	InvoiceComplianceCheckResult *InvoiceComplianceCheckResult `json:"invoice_compliance_check_result,omitempty"`
	// OpenId 是员工支付宝UID
	OpenId *string `json:"open_id,omitempty"`
	// UserId 是员工支付宝UID
	UserId *string `json:"user_id,omitempty"`
	// VoucherContent 是凭证内容
	VoucherContent *string `json:"voucher_content,omitempty"`
	// VoucherDate 是凭证创建时间，格式：yyyy-MM-dd HH:mm:ss
	VoucherDate *string `json:"voucher_date,omitempty"`
	// VoucherId 是凭证ID，幂等用
	VoucherId *string `json:"voucher_id,omitempty"`
	// VoucherSource 是凭证来源
	VoucherSource *string `json:"voucher_source,omitempty"`
	// VoucherType 是凭证类型
	VoucherType *string `json:"voucher_type,omitempty"`
}

// InvoiceComplianceCheckResult 是支付宝官方 V3 OpenAPI 模型。
type InvoiceComplianceCheckResult struct {
	// MatchResultType 是发票和账单匹配的合规检查结果：	0 - 模糊匹配	1 - 精准匹配
	MatchResultType *string `json:"match_result_type,omitempty"`
}
