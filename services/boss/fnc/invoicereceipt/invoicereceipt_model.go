package invoicereceipt

import (
	"github.com/tiyee/alipay-sdk-go/services/boss"
)

// ArMonthlyBillDTO 是支付宝官方 V3 OpenAPI 模型。
type ArMonthlyBillDTO struct {
	// AccountDay 是出账日
	AccountDay *string `json:"account_day,omitempty"`
	// AccruedDate 是应收日期(yyyyMMdd)
	AccruedDate *string                         `json:"accrued_date,omitempty"`
	AdjustAmt   *boss.MultiCurrencyMoneyOpenAPI `json:"adjust_amt,omitempty"`
	// AnalysisDmsn1 是分析维度1
	AnalysisDmsn1 *string `json:"analysis_dmsn_1,omitempty"`
	// AnalysisDmsn2 是分析维度2
	AnalysisDmsn2 *string `json:"analysis_dmsn_2,omitempty"`
	// AnalysisDmsn3 是分析维度3
	AnalysisDmsn3 *string `json:"analysis_dmsn_3,omitempty"`
	// AnalysisDmsn4 是分析维度4
	AnalysisDmsn4 *string `json:"analysis_dmsn_4,omitempty"`
	// ArrangementNo 是合约号，商户签约产品对应的合约号
	ArrangementNo *string                         `json:"arrangement_no,omitempty"`
	BillAmt       *boss.MultiCurrencyMoneyOpenAPI `json:"bill_amt,omitempty"`
	// BillEndDate 是账单结束日期
	BillEndDate *string `json:"bill_end_date,omitempty"`
	// BillMonth 是账单月
	BillMonth *string `json:"bill_month,omitempty"`
	// BillNo 是月账单号,唯一标识月账单的ID
	BillNo *string `json:"bill_no,omitempty"`
	// BillStartDate 是账单开始日期
	BillStartDate *string `json:"bill_start_date,omitempty"`
	// BillType 是账单类型 	21：应收月账单 	目前仅支持该类型的账单类型
	BillType *string `json:"bill_type,omitempty"`
	// ChargeItemCode 是收费项编码
	ChargeItemCode *string `json:"charge_item_code,omitempty"`
	// ChargeType 是计费类型 	01：云在线计费，02：离线计费，03：老主站计收费，04：芝麻在线计收费
	ChargeType *string `json:"charge_type,omitempty"`
	// CheckStatus 是出账状态 	01：账单生成，02：待确认出账，03：已出账，04：客户已对账，05：对账差异，06：部分出账
	CheckStatus *string                         `json:"check_status,omitempty"`
	CheckedAmt  *boss.MultiCurrencyMoneyOpenAPI `json:"checked_amt,omitempty"`
	CheckingAmt *boss.MultiCurrencyMoneyOpenAPI `json:"checking_amt,omitempty"`
	// ClcnBasicAmt 是计算依据
	ClcnBasicAmt *int64 `json:"clcn_basic_amt,omitempty"`
	// ClcnBasicType 是计算依据类型 	01：交易金额，02：交易笔数，03：交易张数，04：资费依据扩展，05：其他
	ClcnBasicType *string `json:"clcn_basic_type,omitempty"`
	// ClcnMethod 是计算方法 	01：差额累进，02：全额累进，03：按传入金额收费，04：价格包量，05：包时间周期，06：按时间全额累进，07：包时间周期差额累进，08：期间差额累进，09：期间全额累进，10：包时间周期单笔，11：标准单笔计费，12：包时间周期单笔-按照传入金额计费，13：期间靠档差额累进，靠档计息使用，14：自定义
	ClcnMethod *string                         `json:"clcn_method,omitempty"`
	FreezeAmt  *boss.MultiCurrencyMoneyOpenAPI `json:"freeze_amt,omitempty"`
	// GmtCreate 是创建时间（创建备份时需要）
	GmtCreate *string `json:"gmt_create,omitempty"`
	// GmtModified 是修改时间（创建备份时需要）
	GmtModified *string `json:"gmt_modified,omitempty"`
	// GmtPay 是付款时间
	GmtPay *string `json:"gmt_pay,omitempty"`
	// InstId 是分支机构号
	InstId      *string                         `json:"inst_id,omitempty"`
	InvoiceAmt  *boss.MultiCurrencyMoneyOpenAPI `json:"invoice_amt,omitempty"`
	InvoicedAmt *boss.MultiCurrencyMoneyOpenAPI `json:"invoiced_amt,omitempty"`
	// IpId 是结算对象参与者标识
	IpId *string `json:"ip_id,omitempty"`
	// IpName 是参与者名字
	IpName *string `json:"ip_name,omitempty"`
	// IpRoleId 是结算对象参与者角色标识
	IpRoleId *string `json:"ip_role_id,omitempty"`
	// LastModer 是最后修改人
	LastModer      *string                         `json:"last_moder,omitempty"`
	LinkInvoiceAmt *boss.MultiCurrencyMoneyOpenAPI `json:"link_invoice_amt,omitempty"`
	// PayOriginal 是付款来源 	1：主账务，2：包量，3：预收，4：预存，5：银行，6：积分，7：统一支付，8：支付受理
	PayOriginal *string `json:"pay_original,omitempty"`
	// PayStatus 是付款状态 	01：待结算；02：部分结算；03：结算完成
	PayStatus *string `json:"pay_status,omitempty"`
	// PayWay 是付款方式 	1：资金；2：包量
	PayWay *string `json:"pay_way,omitempty"`
	// PayeeAccount 是收款方账号
	PayeeAccount *string `json:"payee_account,omitempty"`
	// PayerAccount 是付款方账号
	PayerAccount *string `json:"payer_account,omitempty"`
	// PayerIpRoleId 是扣款对象pid
	PayerIpRoleId *string `json:"payer_ip_role_id,omitempty"`
	// ProdCode 是产品码
	ProdCode    *string                         `json:"prod_code,omitempty"`
	ReceivedAmt *boss.MultiCurrencyMoneyOpenAPI `json:"received_amt,omitempty"`
	// RecordId 是记录id，账单全局唯一号
	RecordId   *string                         `json:"record_id,omitempty"`
	ServiceAmt *boss.MultiCurrencyMoneyOpenAPI `json:"service_amt,omitempty"`
	// SettleType 是结算类型 	01：实时；02：预收；03：后收；04：周期结算；05：按日汇总结算；09：延期结算
	SettleType *string `json:"settle_type,omitempty"`
	// SignIpId 是签约对象参与者标识
	SignIpId *string `json:"sign_ip_id,omitempty"`
	// SignIpRoleId 是签约对象参与者角色标识
	SignIpRoleId *string `json:"sign_ip_role_id,omitempty"`
	// StlIpRoleId 是结算对象PID
	StlIpRoleId *string                         `json:"stl_ip_role_id,omitempty"`
	TaxAmt      *boss.MultiCurrencyMoneyOpenAPI `json:"tax_amt,omitempty"`
	// TaxRate 是税率
	TaxRate *int64 `json:"tax_rate,omitempty"`
	// TaxType 是税收类型，01：增值税，02：营业税 	营业税已经在大部分行业废弃，要传入需咨询后再确认
	TaxType *string `json:"tax_type,omitempty"`
	// ToWriteoffDetailCount 是待核销明细数量
	ToWriteoffDetailCount *int64 `json:"to_writeoff_detail_count,omitempty"`
	// Type 是账单分类 1应收，2返点
	Type *string `json:"type,omitempty"`
	// WriteoffDetailDount 是核销明细数量
	WriteoffDetailDount *int64                          `json:"writeoff_detail_dount,omitempty"`
	WritingoffAmt       *boss.MultiCurrencyMoneyOpenAPI `json:"writingoff_amt,omitempty"`
}
