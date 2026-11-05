package bill

const (
	// BillFileCodeEmptyDataWithFile 表示无数据但存在账单文件。
	BillFileCodeEmptyDataWithFile BillFileCode = "EMPTY_DATA_WITH_BILL_FILE"
)

// BillFileCode 是账单文件结果码的文档已知值；未知值原样保留。
type BillFileCode string

// AccountLogItem 是一条账务明细。
type AccountLogItem struct {
	// AccountLogID 是支付宝账务流水号。
	AccountLogID *string `json:"account_log_id,omitempty"`
	// AlipayOrderNo 是支付宝订单号。
	AlipayOrderNo *string `json:"alipay_order_no,omitempty"`
	// MerchantOrderNo 是商户订单号。
	MerchantOrderNo *string `json:"merchant_order_no,omitempty"`
	// MerchantOutRefundNo 是商户退款请求号。
	MerchantOutRefundNo *string `json:"merchant_out_refund_no,omitempty"`
	// TransAmount 是金额，单位为元；正数为收入、负数为支出。
	TransAmount *string `json:"trans_amount,omitempty"`
	// Balance 是该笔变动后的余额，仅供参考；不保证连续，最终一致。
	Balance *string `json:"balance,omitempty"`
	// Direction 是收入/支出描述。
	Direction *string `json:"direction,omitempty"`
	// TransDt 是入账时间。
	TransDt *string `json:"trans_dt,omitempty"`
	// Type 是账务记录的类型，仅供参考（可能变化）。
	Type *string `json:"type,omitempty"`
	// BizOrigNo 是业务基础订单号。
	BizOrigNo *string `json:"biz_orig_no,omitempty"`
	// BizNos 是业务订单号明细（M：平台交易主单号，S：子单号，
	// O：业务系统单据号）。
	BizNos *string `json:"biz_nos,omitempty"`
	// BizDesc 是业务描述。
	BizDesc *string `json:"biz_desc,omitempty"`
	// BillSource 是业务账单来源。
	BillSource *string `json:"bill_source,omitempty"`
	// OtherAccount 是对方账户。
	OtherAccount *string `json:"other_account,omitempty"`
	// StoreName 是门店信息。
	StoreName *string `json:"store_name,omitempty"`
	// TransMemo 是账务备注；不可依赖此字段进行对账。
	TransMemo *string `json:"trans_memo,omitempty"`
}

// BailDetailResult 是支付宝官方 V3 OpenAPI 模型。
type BailDetailResult struct {
	// Amount 是保证金收支金额。
	Amount *string `json:"amount,omitempty"`
	// BailType 是保证金类型描述，仅供参考。
	BailType *string `json:"bail_type,omitempty"`
	// Balance 是本次操作后的保证金余额。字段数据展示为"--"，表明数据暂未更新，请稍等1分钟后重试。。
	Balance *string `json:"balance,omitempty"`
	// BizDesc 是业务描述，资金收支对应的详细业务场景信息。
	BizDesc *string `json:"biz_desc,omitempty"`
	// BizOrigNo 是业务基础订单号，资金收支对应的原始业务订单唯一识别编号。
	BizOrigNo *string `json:"biz_orig_no,omitempty"`
	// Memo 是保证金说明。
	Memo *string `json:"memo,omitempty"`
	// TransDt 是业务发生时间。
	TransDt *string `json:"trans_dt,omitempty"`
	// TransLogId 是保证金业务流水号。
	TransLogId *string `json:"trans_log_id,omitempty"`
}

// BizFundReportResult 是支付宝官方 V3 OpenAPI 模型。
type BizFundReportResult struct {
	// AlipayAmount 是实付金额（元）。
	AlipayAmount *string `json:"alipay_amount,omitempty"`
	// BatchType 是是否批量转账。
	BatchType *string `json:"batch_type,omitempty"`
	// BizType 是业务类型。
	BizType *string `json:"biz_type,omitempty"`
	// BizTypeDesc 是业务类型描述。
	BizTypeDesc *string `json:"biz_type_desc,omitempty"`
	// ChargeFee 是服务费（元）。
	ChargeFee *string `json:"charge_fee,omitempty"`
	// GmtCreate 是交易时间。
	GmtCreate *string `json:"gmt_create,omitempty"`
	// InstructionId 是银行单据号。
	InstructionId *string `json:"instruction_id,omitempty"`
	// Memo 是出借。
	Memo *string `json:"memo,omitempty"`
	// NoMbillEncrypt 是是否脱敏。非空就是不脱敏，空就是脱敏。脱敏规则根据发生的业务定义。
	NoMbillEncrypt *string `json:"no_mbill_encrypt,omitempty"`
	// OrderId 是业务订单号。
	OrderId *string `json:"order_id,omitempty"`
	// PayAmount 是实付金额（元）。
	PayAmount *string `json:"pay_amount,omitempty"`
	// PayeeCardNo 是收款方银行卡号。
	PayeeCardNo *string `json:"payee_card_no,omitempty"`
	// PayeeFullName 是收款方全名。
	PayeeFullName *string `json:"payee_full_name,omitempty"`
	// PayeeFundType 是收款方资金类型。
	PayeeFundType *string `json:"payee_fund_type,omitempty"`
	// PayeeFundTypeDesc 是支付宝余额。
	PayeeFundTypeDesc *string `json:"payee_fund_type_desc,omitempty"`
	// PayeeInstName 是收款银行中文简称。
	PayeeInstName *string `json:"payee_inst_name,omitempty"`
	// PayeeLoginEmail 是收款方登录Id。
	PayeeLoginEmail *string `json:"payee_login_email,omitempty"`
	// PayeeName 是收款方姓名。
	PayeeName *string `json:"payee_name,omitempty"`
	// PayerFundType 是付款方资金类型。
	PayerFundType *string `json:"payer_fund_type,omitempty"`
	// PayerFundTypeDesc 是付款方资金类型描述。
	PayerFundTypeDesc *string `json:"payer_fund_type_desc,omitempty"`
	// RefundAmount 是退款金额（元）。
	RefundAmount *string `json:"refund_amount,omitempty"`
	// ShowDetail 是是否展示细节，默认为true。
	ShowDetail *string `json:"show_detail,omitempty"`
	// ShowVoucher 是是否出示凭证。
	ShowVoucher *string `json:"show_voucher,omitempty"`
	// Source 是数据来源。
	Source *string `json:"source,omitempty"`
	// Status 是订单状态。
	Status *string `json:"status,omitempty"`
	// StatusDesc 是单据状态描述。
	StatusDesc *string `json:"status_desc,omitempty"`
	// SubBizType 是普通转账。
	SubBizType *string `json:"sub_biz_type,omitempty"`
	// SubBizTypeDesc 是业务子类型描述。
	SubBizTypeDesc *string `json:"sub_biz_type_desc,omitempty"`
}

// TradeItemResult 是支付宝官方 V3 OpenAPI 模型。
type TradeItemResult struct {
	// AlipayOrderNo 是支付宝订单号。对账使用，不脱敏。
	AlipayOrderNo *string `json:"alipay_order_no,omitempty"`
	// GmtCreate 是交易创建时间。
	GmtCreate *string `json:"gmt_create,omitempty"`
	// GmtPay 是交易支付时间。支付后会有值。
	GmtPay *string `json:"gmt_pay,omitempty"`
	// GmtRefund 是交易退款时间。
	GmtRefund *string `json:"gmt_refund,omitempty"`
	// GoodsMemo 是商品备注信息。
	GoodsMemo *string `json:"goods_memo,omitempty"`
	// GoodsTitle 是商品名称。
	GoodsTitle *string `json:"goods_title,omitempty"`
	// MerchantOrderNo 是商户订单号，创建支付宝交易时传入的信息。对账使用，不脱敏。
	MerchantOrderNo *string `json:"merchant_order_no,omitempty"`
	// NetMdiscount 是商家优惠金额，人民币元。
	NetMdiscount *string `json:"net_mdiscount,omitempty"`
	// OtherAccount 是对方账户。
	OtherAccount *string `json:"other_account,omitempty"`
	// RefundAmount 是订单退款金额，人民币元。
	RefundAmount *string `json:"refund_amount,omitempty"`
	// ServiceFee 是服务费金额，人民币元。
	ServiceFee *string `json:"service_fee,omitempty"`
	// StoreName 是门店名称。
	StoreName *string `json:"store_name,omitempty"`
	// StoreNo 是门店编号。
	StoreNo *string `json:"store_no,omitempty"`
	// TotalAmount 是订单金额，人民币元。
	TotalAmount *string `json:"total_amount,omitempty"`
	// TradeStatus 是订单状态(待付款,成功,关闭,待发货,待确认收货,已预付,进行中)。
	TradeStatus *string `json:"trade_status,omitempty"`
	// TradeType 是业务类型，帮助商户作为对账参考。
	TradeType *string `json:"trade_type,omitempty"`
}

// TransferAccountBookDetailResult 是支付宝官方 V3 OpenAPI 模型。
type TransferAccountBookDetailResult struct {
	// Account 是付款/收款账户。充值记录中是付款账户。提现、转账记录中是收款账户。支付宝名称及账号脱敏；银行账户的户名脱敏，银行账户显示银行名称+银行卡号后四位。
	Account *string `json:"account,omitempty"`
	// Amount 是金额。
	Amount *string `json:"amount,omitempty"`
	// BizDesc 是业务类型描述。可用于区分业务类型，和不同子账本。
	BizDesc *string `json:"biz_desc,omitempty"`
	// FundDesc 是资金来源/去向类型。在充值记录中，表示资金来源类型，在转账和提现类型中，表示去向类型。
	FundDesc *string `json:"fund_desc,omitempty"`
	// InstructionId 是银行单据号。对账使用，无需脱敏。
	InstructionId *string `json:"instruction_id,omitempty"`
	// Memo 是备注信息。
	Memo *string `json:"memo,omitempty"`
	// OrderNo 是业务订单号。该笔业务单据的唯一识别编号。
	OrderNo *string `json:"order_no,omitempty"`
	// ServiceFee 是服务费金额。
	ServiceFee *string `json:"service_fee,omitempty"`
	// Status 是资金状态。
	Status *string `json:"status,omitempty"`
	// SubTypeDesc 是子类型。“充值类型”，普通充值、大额充值。“转账类型”，暂无实现。转账至支付宝账户、转账至银行卡、批量转账支付宝账户、批量转账至银行卡、批量付款。“提现类型”，暂无实现。普通提现、批量委托提现。对账使用，无需脱敏。
	SubTypeDesc *string `json:"sub_type_desc,omitempty"`
	// TransDt 是业务发生时间。
	TransDt *string `json:"trans_dt,omitempty"`
	// TypeDesc 是查询类型描述：充值、转账、提现。
	TypeDesc *string `json:"type_desc,omitempty"`
}

// TransferDetailResult 是支付宝官方 V3 OpenAPI 模型。
type TransferDetailResult struct {
	// Account 是付款/收款账户。充值记录中是付款账户。提现、转账记录中是收款账户。支付宝名称及账号脱敏；银行账户的户名脱敏，银行账户显示银行名称+银行卡号后四位。
	Account *string `json:"account,omitempty"`
	// Amount 是金额，人民币元。
	Amount *string `json:"amount,omitempty"`
	// FundDesc 是资金来源/去向类型。在充值记录中，表示资金来源类型，在转账和提现类型中，表示去向类型。
	FundDesc *string `json:"fund_desc,omitempty"`
	// InstructionId 是银行单据号。对账使用，无需脱敏。
	InstructionId *string `json:"instruction_id,omitempty"`
	// Memo 是备注信息。
	Memo *string `json:"memo,omitempty"`
	// OrderNo 是业务订单号。该笔业务单据的唯一识别编号。
	OrderNo *string `json:"order_no,omitempty"`
	// ServiceFee 是服务费金额，人民币元。
	ServiceFee *string `json:"service_fee,omitempty"`
	// Status 是资金状态。
	Status *string `json:"status,omitempty"`
	// SubTypeDesc 是子类型。“充值类型”，普通充值、大额充值。“转账类型”，暂无实现。转账至支付宝账户、转账至银行卡、批量转账支付宝账户、批量转账至银行卡、批量付款。“提现类型”，暂无实现。普通提现、批量委托提现。对账使用，无需脱敏。
	SubTypeDesc *string `json:"sub_type_desc,omitempty"`
	// TransDt 是业务发生时间。
	TransDt *string `json:"trans_dt,omitempty"`
	// TypeDesc 是查询类型描述：充值、转账、提现。
	TypeDesc *string `json:"type_desc,omitempty"`
}
