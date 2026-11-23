package invoicereceipt

import (
	"encoding/json"

	"github.com/tiyee/alipay-sdk-go/services/boss"
)

// BossFncInvoiceReceiptBatchQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type BossFncInvoiceReceiptBatchQueryResponse struct {
	// ResultSet 是返回结果对象：可开票单据
	ResultSet           *[]ArInvoiceReceiptOpenAPIResponse `json:"result_set,omitempty"`
	TotalInvAmt         *boss.MultiCurrencyMoneyOpenAPI    `json:"total_inv_amt,omitempty"`
	TotalInvedAmt       *boss.MultiCurrencyMoneyOpenAPI    `json:"total_inved_amt,omitempty"`
	TotalLinkInvoiceAmt *boss.MultiCurrencyMoneyOpenAPI    `json:"total_link_invoice_amt,omitempty"`
}

// BossFncInvoiceReceiptCreateResponse 是支付宝官方 V3 OpenAPI 模型。
type BossFncInvoiceReceiptCreateResponse = json.RawMessage

// ArInvoiceReceiptOpenAPIResponse 是支付宝官方 V3 OpenAPI 模型。
type ArInvoiceReceiptOpenAPIResponse struct {
	// ArrangementNo 是合约号
	ArrangementNo *string `json:"arrangement_no,omitempty"`
	// Id 是可开票单据主键ID
	Id *string `json:"id,omitempty"`
	// InstId 是机构ID
	InstId *string `json:"inst_id,omitempty"`
	// InvDt 是开票时间 格式：yyyymm
	InvDt *string `json:"inv_dt,omitempty"`
	// InvMode 是开票模式 	01：实收开票， 02：应收开票， 03：累计实收开票
	InvMode     *string                         `json:"inv_mode,omitempty"`
	InvoiceAmt  *boss.MultiCurrencyMoneyOpenAPI `json:"invoice_amt,omitempty"`
	InvoicedAmt *boss.MultiCurrencyMoneyOpenAPI `json:"invoiced_amt,omitempty"`
	// IpId 是结算ip_id
	IpId *string `json:"ip_id,omitempty"`
	// IpRoleId 是结算对象ip_role_id
	IpRoleId *string `json:"ip_role_id,omitempty"`
	// LastModer 是最后修改人
	LastModer      *string                         `json:"last_moder,omitempty"`
	LinkInvoiceAmt *boss.MultiCurrencyMoneyOpenAPI `json:"link_invoice_amt,omitempty"`
	// OutBizNo 是外部单据号，对应开票记录的月账单号
	OutBizNo *string `json:"out_biz_no,omitempty"`
	// OutBizType 是可开票单据来源，01：主站，02：芝麻，03：金融云，04：微贷
	OutBizType *string `json:"out_biz_type,omitempty"`
	// PayWay 是付款方式，1：资金，5：走量
	PayWay *string `json:"pay_way,omitempty"`
	// ProdCode 是产品码
	ProdCode *string `json:"prod_code,omitempty"`
	// SettleType 是结算类型 ，01：实时，02：预收，03：后收，04：周期性，05：按日汇总，09：延期结算
	SettleType *string `json:"settle_type,omitempty"`
	// Status 是开票金额消耗状态：01未开票，02部分开票，03：已开票
	Status *string `json:"status,omitempty"`
	// TaxRate 是税率
	TaxRate *int64 `json:"tax_rate,omitempty"`
	// TaxType 是税收类型01：增值税，02：营业税
	TaxType *string `json:"tax_type,omitempty"`
	// TntInstId 是租户ID
	TntInstId *string `json:"tnt_inst_id,omitempty"`
	// Type 是类型，1：应收，2：返点
	Type *string `json:"type,omitempty"`
}
