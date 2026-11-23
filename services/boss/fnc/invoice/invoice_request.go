package invoice

import (
	"github.com/tiyee/alipay-sdk-go/services/boss"
)

// BossFncInvoiceApplyRequest 是支付宝官方 V3 OpenAPI 模型。
type BossFncInvoiceApplyRequest struct {
	// BizId 是调用方身份识别标记(应用名称) 与业务单据号联合唯一
	BizId *string `json:"biz_id,omitempty"`
	// BizNo 是业务单据号，业务流水号
	BizNo      *string                         `json:"biz_no,omitempty"`
	InvoiceAmt *boss.MultiCurrencyMoneyOpenAPI `json:"invoice_amt,omitempty"`
	// InvoiceType 是申请开票类型，01：专票；02：普票
	InvoiceType *string `json:"invoice_type,omitempty"`
	// Memo 是备注
	Memo *string `json:"memo,omitempty"`
	// MthtlyBillNos 是需开票的应收月账单号
	MthtlyBillNos *[]string `json:"mthtly_bill_nos,omitempty"`
	// Operator 是申请人的uid号
	Operator *string `json:"operator,omitempty"`
	// OutBizType 是01主站，02芝麻，03金融云
	OutBizType *string `json:"out_biz_type,omitempty"`
}

// BossFncInvoiceBatchQueryRequest 是支付宝官方 V3 OpenAPI 模型。
type BossFncInvoiceBatchQueryRequest struct {
	// ApplyOrderId 是开票申请ID，唯一标识开票申请表的ID
	ApplyOrderId *string `json:"apply_order_id,omitempty"`
	// BuyerInvoiceTitle 是买方发票抬头，用于打印在发票上
	BuyerInvoiceTitle *string `json:"buyer_invoice_title,omitempty"`
	// BuyerIpid 是对应发票商户的ipId，在主站标示mid，在其他环境标示ip_id。
	BuyerIpid *string `json:"buyer_ipid,omitempty"`
	// BuyerIpids 是购方商户ID集合
	BuyerIpids *[]string `json:"buyer_ipids,omitempty"`
	// BuyerIproleId 是商户结算的ipRoleId，表示商户的结算对象，在主站表示pid。
	BuyerIproleId *string `json:"buyer_iprole_id,omitempty"`
	// BuyerIproleIds 是购方PID集合
	BuyerIproleIds *[]string `json:"buyer_iprole_ids,omitempty"`
	// InstId 是对应发票所在OU的id，用于标示该发票所述的OU。
	InstId *string `json:"inst_id,omitempty"`
	// InvoiceChannel 是开票渠道 01 线上;02 线下
	InvoiceChannel *string `json:"invoice_channel,omitempty"`
	// InvoiceCode 是发票代码，税务部门给予发票的编码
	InvoiceCode *string `json:"invoice_code,omitempty"`
	// InvoiceCreateDateBegin 是发票创建日期【起始日期，格式yyyyMMdd】
	InvoiceCreateDateBegin *string `json:"invoice_create_date_begin,omitempty"`
	// InvoiceCreateDateEnd 是发票创建日期【结束日期，格式yyyyMMdd】
	InvoiceCreateDateEnd *string `json:"invoice_create_date_end,omitempty"`
	// InvoiceIds 是发票ID集合，发票ID，唯一标示一张发票的ID
	InvoiceIds *[]string `json:"invoice_ids,omitempty"`
	// InvoiceMaterial 是开票介质 01电子 02纸质
	InvoiceMaterial *string `json:"invoice_material,omitempty"`
	// InvoiceNo 是发票号码，税务部门给予发票的编码
	InvoiceNo *string `json:"invoice_no,omitempty"`
	// InvoiceOpenDateBegin 是发票开票日期【起始日期，格式yyyyMMdd】
	InvoiceOpenDateBegin *string `json:"invoice_open_date_begin,omitempty"`
	// InvoiceOpenDateEnd 是发票开票日期【结束日期，格式yyyyMMdd】
	InvoiceOpenDateEnd *string `json:"invoice_open_date_end,omitempty"`
	// InvoiceStatus 是发票状态：用于标示发票当前的状态，状态类型以及相关含义：01 待开票;02 开票中;03 已开票;04 作废中;05 已作废;06 红冲中;07 已红冲;08 部分红冲;09 无效
	InvoiceStatus *[]string `json:"invoice_status,omitempty"`
	// InvoiceTypes 是发票类型,01专票 02普票 03营业税发票 04国际形式发票 05其它发票
	InvoiceTypes *[]string `json:"invoice_types,omitempty"`
	// IsOnline 是是否能在线开票 Y能 N不能
	IsOnline *string `json:"is_online,omitempty"`
	// IsRed 是是否红字 Y红字 N蓝字
	IsRed *string `json:"is_red,omitempty"`
	// MailStatus 是邮寄状态 01 待邮寄;02 已邮寄;03 退回中;04 已退回
	MailStatus *[]string `json:"mail_status,omitempty"`
	// MonthlyBillNo 是月账单号
	MonthlyBillNo *string `json:"monthly_bill_no,omitempty"`
	// PageNo 是起始页，分页时必须提供
	PageNo *int64 `json:"page_no,omitempty"`
	// PageSize 是每页显示数，分页时必须提供
	PageSize *int64 `json:"page_size,omitempty"`
	// TrackingNo 是快递单号
	TrackingNo *string `json:"tracking_no,omitempty"`
}

// BossFncInvoiceQueryRequest is the query request model for alipay.boss.fnc.invoice.query.
type BossFncInvoiceQueryRequest struct {
	// InvoiceId 是查询ID：用于标示一张发票，通过发票综合查询可以查询到该ID。
	InvoiceId *string `json:"invoice_id,omitempty" query:"invoice_id,omitempty"`
}
