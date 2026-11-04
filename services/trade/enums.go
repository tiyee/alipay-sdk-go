// Package trade 封装支付宝传统网关的交易领域接口。
// 请求、响应与枚举严格依据各接口官方文档建模；
// 响应中的未知枚举值保留原始值，不会改写为已知状态。
package trade

// TradeStatus 是交易状态，取值依据 alipay.trade.query 文档
// （https://opendocs.alipay.com/open/02ekfq.md）。
type TradeStatus string

const (
	// TradeStatusWaitBuyerPay 表示交易创建，等待买家付款。
	TradeStatusWaitBuyerPay TradeStatus = "WAIT_BUYER_PAY"
	// TradeStatusClosed 表示未付款交易超时关闭，或支付完成后全额退款。
	TradeStatusClosed TradeStatus = "TRADE_CLOSED"
	// TradeStatusSuccess 表示交易支付成功。
	TradeStatusSuccess TradeStatus = "TRADE_SUCCESS"
	// TradeStatusFinished 表示交易结束，不可退款。
	TradeStatusFinished TradeStatus = "TRADE_FINISHED"
)

// IsKnown 报告状态是否属于文档已知值；
// 未知值可能是平台新增状态，调用方应在 default 分支处理。
func (s TradeStatus) IsKnown() bool {
	switch s {
	case TradeStatusWaitBuyerPay, TradeStatusClosed, TradeStatusSuccess, TradeStatusFinished:
		return true
	}
	return false
}

// BuyerUserType 是买家用户类型，取值依据 alipay.trade.query 文档。
type BuyerUserType string

const (
	// BuyerUserTypeCorporate 表示企业用户。
	BuyerUserTypeCorporate BuyerUserType = "CORPORATE"
	// BuyerUserTypePrivate 表示个人用户。
	BuyerUserTypePrivate BuyerUserType = "PRIVATE"
)

// FundChannel 是资金渠道标识（如 ALIPAYACCOUNT、BANKCARD）。
// 官方支付渠道列表独立维护且可能扩充，本 SDK 不硬编码完整清单；
// 响应中的未知渠道值原样保留。
type FundChannel string

// FundType 是资金类型，仅在资金渠道为银行卡（BANKCARD）时返回，
// 取值依据 alipay.trade.refund 文档。
type FundType string

const (
	// FundTypeDebitCard 表示借记卡。
	FundTypeDebitCard FundType = "DEBIT_CARD"
	// FundTypeCreditCard 表示信用卡。
	FundTypeCreditCard FundType = "CREDIT_CARD"
	// FundTypeMixedCard 表示借贷合一卡。
	FundTypeMixedCard FundType = "MIXED_CARD"
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

// VoucherType 是优惠券类型，取值依据 alipay.trade.refund 文档；
// 官方注明不排除将来新增类型，响应未知值原样保留。
type VoucherType string

const (
	// VoucherTypeFix 表示全场代金券。
	VoucherTypeFix VoucherType = "ALIPAY_FIX_VOUCHER"
	// VoucherTypeDiscount 表示折扣券。
	VoucherTypeDiscount VoucherType = "ALIPAY_DISCOUNT_VOUCHER"
	// VoucherTypeItem 表示单品优惠券。
	VoucherTypeItem VoucherType = "ALIPAY_ITEM_VOUCHER"
	// VoucherTypeCash 表示现金抵价券。
	VoucherTypeCash VoucherType = "ALIPAY_CASH_VOUCHER"
	// VoucherTypeBiz 表示商家全场券。
	VoucherTypeBiz VoucherType = "ALIPAY_BIZ_VOUCHER"
)

// ContributeType 是优惠券出资方类型，取值依据退款文档；
// 官方注明可能新增类型。
type ContributeType string

const (
	// ContributeTypePlatform 表示平台出资。
	ContributeTypePlatform ContributeType = "PLATFORM"
	// ContributeTypeBrand 表示品牌商出资。
	ContributeTypeBrand ContributeType = "BRAND"
	// ContributeTypeMall 表示商圈出资。
	ContributeTypeMall ContributeType = "MALL"
)

// RoyaltyType 是分账类型，取值依据退款文档：为空默认为 transfer。
type RoyaltyType string

const (
	// RoyaltyTypeTransfer 表示普通分账。
	RoyaltyTypeTransfer RoyaltyType = "transfer"
	// RoyaltyTypeReplenish 表示补差分账。
	RoyaltyTypeReplenish RoyaltyType = "replenish"
)

// RoyaltyResultCode 是退分账结果码的已知值；其他结果码原样保留。
type RoyaltyResultCode string

const (
	// RoyaltyResultSuccess 表示退分账成功。
	RoyaltyResultSuccess RoyaltyResultCode = "SUCCESS"
)

// AccountType 是分账/结算账户类型，取值依据预下单与退款文档。
type AccountType string

const (
	// AccountTypeUserID 表示支付宝用户号（2088 开头 16 位数字）。
	AccountTypeUserID AccountType = "userId"
	// AccountTypeLoginName 表示支付宝登录号。
	AccountTypeLoginName AccountType = "loginName"
	// AccountTypeCardAliasNo 表示用户在支付宝绑定的卡编号。
	AccountTypeCardAliasNo AccountType = "cardAliasNo"
	// AccountTypeDefaultSettle 表示默认结算主体；此时 trans_in 不能传值。
	AccountTypeDefaultSettle AccountType = "defaultSettle"
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

// ProductCode 是预下单产品码，取值依据 alipay.trade.precreate 文档。
type ProductCode string

const (
	// ProductCodeFaceToFacePayment 表示当面付产品，为默认值。
	ProductCodeFaceToFacePayment ProductCode = "FACE_TO_FACE_PAYMENT"
)

// CodeType 是预下单码类型，取值依据预下单文档；为空默认为订单码。
type CodeType string

const (
	// CodeTypeWaveCode 表示声波码。
	CodeTypeWaveCode CodeType = "wave_code"
	// CodeTypeQRCode 表示二维码。
	CodeTypeQRCode CodeType = "qr_code"
	// CodeTypeBarCode 表示条形码。
	CodeTypeBarCode CodeType = "bar_code"
)

// QueryOption 是 alipay.trade.query 的查询选项，
// 用于定制同步额外返回的字段，取值依据查询文档。
type QueryOption string

const (
	// QueryOptionTradeSettleInfo 请求返回交易结算信息。
	QueryOptionTradeSettleInfo QueryOption = "trade_settle_info"
	// QueryOptionFundBillList 请求返回交易支付使用的资金渠道。
	QueryOptionFundBillList QueryOption = "fund_bill_list"
	// QueryOptionVoucherDetailList 请求返回交易支付时使用的所有优惠券信息。
	QueryOptionVoucherDetailList QueryOption = "voucher_detail_list"
	// QueryOptionDiscountGoodsDetail 请求返回单品券优惠的商品优惠信息。
	QueryOptionDiscountGoodsDetail QueryOption = "discount_goods_detail"
	// QueryOptionMDiscountAmount 请求返回商家优惠金额。
	QueryOptionMDiscountAmount QueryOption = "mdiscount_amount"
	// QueryOptionMedicalInsuranceInfo 请求返回医保信息。
	QueryOptionMedicalInsuranceInfo QueryOption = "medical_insurance_info"
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

// PrecreateQueryOption 是 alipay.trade.precreate 的返回参数选项，
// 取值依据预下单文档。
type PrecreateQueryOption string

const (
	// PrecreateQueryOptionEnterprisePayInfo 请求返回因公付金额信息。
	PrecreateQueryOptionEnterprisePayInfo PrecreateQueryOption = "enterprise_pay_info"
	// PrecreateQueryOptionHybAmount 请求返回惠营宝回票金额信息。
	PrecreateQueryOptionHybAmount PrecreateQueryOption = "hyb_amount"
)

// NotifyType 是异步通知类型，取值依据官方《异步通知说明》
// （https://opendocs.alipay.com/open/194/103296）。
type NotifyType string

const (
	// NotifyTypeTradeStatusSync 表示统一收单交易支付完成（交易状态同步）通知。
	NotifyTypeTradeStatusSync NotifyType = "trade_status_sync"
)

// NotifyVoucherType 是异步通知中优惠券的类型，取值依据官方《异步通知说明》；
// 官方注明不排除新增类型，未知值原样保留。
type NotifyVoucherType string

const (
	// NotifyVoucherTypeBiz 表示商家全场券。
	NotifyVoucherTypeBiz NotifyVoucherType = "ALIPAY_BIZ_VOUCHER"
	// NotifyVoucherTypeCommonItem 表示商家单品券。
	NotifyVoucherTypeCommonItem NotifyVoucherType = "ALIPAY_COMMON_ITEM_VOUCHER"
	// NotifyVoucherTypeCash 表示平台优惠券，支付宝或第三方出资。
	NotifyVoucherTypeCash NotifyVoucherType = "ALIPAY_CASH_VOUCHER"
	// NotifyVoucherTypeCreditIntFree 表示花呗分期券，仅减免订单外分期费用。
	NotifyVoucherTypeCreditIntFree NotifyVoucherType = "ALICREDIT_INTFREE_VOUCHER"
)

// BizSettleMode 是异步通知中的账期结算标识，
// 取值依据官方《异步通知说明》。
type BizSettleMode string

const (
	// BizSettleModePeriod 表示已完成支付的订单进行账期管控，不会实时结算。
	BizSettleModePeriod BizSettleMode = "PERIOD"
)
