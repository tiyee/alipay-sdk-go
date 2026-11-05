package refund

import (
	"github.com/tiyee/alipay-sdk-go/services/trade"
)

// RefundStatus 是退款状态，仅出现在退款查询响应中；
// 该枚举只有 REFUND_SUCCESS 一个文档值，字段未返回表示
// 退款请求未收到或退款失败。依据 alipay.trade.fastpay.refund.query 文档。
type RefundStatus string

const (
	// RefundStatusSuccess 表示退款处理成功。
	RefundStatusSuccess RefundStatus = "REFUND_SUCCESS"
)

// FundChange 表示本次退款是否发生了资金变化，
// 取值依据 alipay.trade.refund 文档：Y 表示退款成功；
// N 或未返回需通过退款查询接口确认。
type FundChange string

const (
	// FundChangeYes 表示本次退款发生了资金变化（退款成功）。
	FundChangeYes FundChange = "Y"
	// FundChangeNo 表示本次退款未发生资金变化，需退款查询确认。
	FundChangeNo FundChange = "N"
)

// DepositBackStatus 是银行卡冲退状态，取值依据退款查询文档：
// 冲退失败时资金自动转入用户支付宝余额。
type DepositBackStatus string

const (
	// DepositBackSuccess 表示冲退成功。
	DepositBackSuccess DepositBackStatus = "S"
	// DepositBackFailed 表示冲退失败。
	DepositBackFailed DepositBackStatus = "F"
	// DepositBackProcessing 表示冲退处理中。
	DepositBackProcessing DepositBackStatus = "P"
)

// RoyaltyResultCode 是退分账结果码的已知值；其他结果码原样保留。
type RoyaltyResultCode string

const (
	// RoyaltyResultSuccess 表示退分账成功。
	RoyaltyResultSuccess RoyaltyResultCode = "SUCCESS"
)

// ChargeType 是退费手续费类型，取值依据退款文档。
type ChargeType string

const (
	// ChargeTypeTrade 表示收单手续费。
	ChargeTypeTrade ChargeType = "trade"
	// ChargeTypeHBFQ 表示花呗分期手续费。
	ChargeTypeHBFQ ChargeType = "hbfq"
	// ChargeTypeOther 表示其他手续费。
	ChargeTypeOther ChargeType = "charge"
)

// RefundQueryOption 是退款与退款查询接口的查询选项，
// 取值依据 alipay.trade.refund 与 alipay.trade.fastpay.refund.query 文档。
type RefundQueryOption string

const (
	// RefundQueryOptionRefundDetailItemList 请求返回本次退款使用的资金渠道。
	RefundQueryOptionRefundDetailItemList RefundQueryOption = "refund_detail_item_list"
	// RefundQueryOptionDepositBackInfo 请求返回银行卡冲退信息。
	RefundQueryOptionDepositBackInfo RefundQueryOption = "deposit_back_info"
	// RefundQueryOptionRefundVoucherDetailList 请求返回本次退款退的券信息。
	RefundQueryOptionRefundVoucherDetailList RefundQueryOption = "refund_voucher_detail_list"
	// RefundQueryOptionGmtRefundPay 请求返回退款执行成功的时间（仅退款查询）。
	RefundQueryOptionGmtRefundPay RefundQueryOption = "gmt_refund_pay"
)

// RefundApplyStatus 是异步退款受理接口的退款状态，取值依据官方 V3 文档。
type RefundApplyStatus string

const (
	// RefundApplyStatusProcessing 表示退款处理中（非终态，需继续查询）。
	RefundApplyStatusProcessing RefundApplyStatus = "REFUND_PROCESSING"
	// RefundApplyStatusSuccess 表示退款处理成功。
	RefundApplyStatusSuccess RefundApplyStatus = "REFUND_SUCCESS"
	// RefundApplyStatusFail 表示退款失败。
	RefundApplyStatusFail RefundApplyStatus = "REFUND_FAIL"
)

// RefundGoodsDetail 是退款包含的单个商品信息。
type RefundGoodsDetail struct {
	// GoodsID 是商品编号，对应支付时传入的 goods_id。
	GoodsID string `json:"goods_id"`
	// RefundAmount 是该商品的退款总金额。
	RefundAmount string `json:"refund_amount"`
	// OutCertificateNoList 是外部商品凭证编号列表。
	OutCertificateNoList []string `json:"out_certificate_no_list,omitempty"`
	// OutItemID 是商家侧小程序商品 ID。
	OutItemID *string `json:"out_item_id,omitempty"`
	// OutSkuID 是商家侧小程序商品 SKU ID。
	OutSkuID *string `json:"out_sku_id,omitempty"`
}

// RefundChargeInfo 是单条退费信息。
type RefundChargeInfo struct {
	// RefundChargeFee 是实退费用。
	RefundChargeFee *string `json:"refund_charge_fee,omitempty"`
	// SwitchFeeRate 是签约费率。
	SwitchFeeRate *string `json:"switch_fee_rate,omitempty"`
	// ChargeType 是手续费类型。
	ChargeType *ChargeType `json:"charge_type,omitempty"`
	// RefundSubFeeDetailList 是组合支付退费明细。
	RefundSubFeeDetailList []RefundSubFee `json:"refund_sub_fee_detail_list,omitempty"`
}

// RefundSubFee 是组合支付的退费明细子项。
type RefundSubFee struct {
	// RefundChargeFee 是实退费用。
	RefundChargeFee *string `json:"refund_charge_fee,omitempty"`
	// SwitchFeeRate 是签约费率。
	SwitchFeeRate *string `json:"switch_fee_rate,omitempty"`
}

// RefundRoyaltyResult 是退分账结果（直付通产品返回）。
type RefundRoyaltyResult struct {
	// RefundAmount 是退分账金额。
	RefundAmount string `json:"refund_amount"`
	// ResultCode 是退分账结果码；未知值原样保留。
	ResultCode RoyaltyResultCode `json:"result_code"`
	// RoyaltyType 是分账类型；为空默认为 transfer。
	RoyaltyType *trade.RoyaltyType `json:"royalty_type,omitempty"`
	// TransOut 是转出人支付宝账号对应用户 ID。
	TransOut *string `json:"trans_out,omitempty"`
	// TransOutEmail 是转出人支付宝账号。
	TransOutEmail *string `json:"trans_out_email,omitempty"`
	// TransIn 是转入人支付宝账号对应用户 ID。
	TransIn *string `json:"trans_in,omitempty"`
	// TransInEmail 是转入人支付宝账号。
	TransInEmail *string `json:"trans_in_email,omitempty"`
	// OriTransOut 是商户请求的转出账号。
	OriTransOut *string `json:"ori_trans_out,omitempty"`
	// OriTransIn 是商户请求的转入账号。
	OriTransIn *string `json:"ori_trans_in,omitempty"`
}

// DepositBackInfo 是银行卡冲退信息。
type DepositBackInfo struct {
	// HasDepositBack 标识是否存在银行卡冲退信息（"true"/"false"）。
	HasDepositBack *string `json:"has_deposit_back,omitempty"`
	// DbackStatus 是银行卡冲退状态；冲退失败时资金自动转入用户余额。
	DbackStatus *DepositBackStatus `json:"dback_status,omitempty"`
	// DbackAmount 是银行卡冲退金额。
	DbackAmount *string `json:"dback_amount,omitempty"`
	// BankAckTime 是银行响应时间。
	BankAckTime *string `json:"bank_ack_time,omitempty"`
	// EstBankReceiptTime 是预估银行到账时间。
	EstBankReceiptTime *string `json:"est_bank_receipt_time,omitempty"`
}

// RefundFundDetail 是退款资金明细。
type RefundFundDetail struct {
	// Type 描述资金明细类型（如 paySerialNo）。
	Type *string `json:"type,omitempty"`
	// Funds 是退款资金明细内容；格式依 Type 而定。
	Funds []string `json:"funds,omitempty"`
	// TransIn 是收入方账户；为空则原路退回。
	TransIn *string `json:"trans_in,omitempty"`
	// TransInType 是收入方账户类型。
	TransInType *string `json:"trans_in_type,omitempty"`
}
