package trade

import (
	"github.com/tiyee/alipay-sdk-go/core/notify"
)

// 页面与唤端接口的方法名。这些接口不走 V3 REST 调用：
// 产物按传统网关契约构建后由用户浏览器或支付宝客户端 SDK 提交。
// 依据官方 V3 文档与 GenericExecuteApi 的 pageExecute/sdkExecute。
const (
	methodPagePay = "alipay.trade.page.pay"
	methodWapPay  = "alipay.trade.wap.pay"
	methodAppPay  = "alipay.trade.app.pay"
)

// pageBizContent 是 page.pay/wap.pay 的 biz_content JSON 结构，
// 字段名依据官方接口文档。
type pageBizContent struct {
	OutTradeNo     string          `json:"out_trade_no"`
	TotalAmount    string          `json:"total_amount"`
	Subject        string          `json:"subject"`
	ProductCode    string          `json:"product_code"`
	SellerID       *string         `json:"seller_id,omitempty"`
	Body           *string         `json:"body,omitempty"`
	GoodsDetail    []GoodsDetail   `json:"goods_detail,omitempty"`
	TimeExpire     *string         `json:"time_expire,omitempty"`
	TimeoutExpress *string         `json:"timeout_express,omitempty"`
	QRPayMode      *string         `json:"qr_pay_mode,omitempty"`
	QrPayTimeLimit *string         `json:"qr_pay_time_limit,omitempty"`
	SettleInfo     *SettleInfo     `json:"settle_info,omitempty"`
	ExtendParams   *ExtendParams   `json:"extend_params,omitempty"`
	BusinessParams *BusinessParams `json:"business_params,omitempty"`
	PassbackParams *string         `json:"passback_params,omitempty"`
}

// appBizContent 是 app.pay 的 biz_content JSON 结构，
// 字段名依据官方接口文档与商家扣款场景文档。
type appBizContent struct {
	OutTradeNo     string          `json:"out_trade_no"`
	TotalAmount    string          `json:"total_amount"`
	Subject        string          `json:"subject"`
	ProductCode    string          `json:"product_code"`
	SellerID       *string         `json:"seller_id,omitempty"`
	Body           *string         `json:"body,omitempty"`
	GoodsDetail    []GoodsDetail   `json:"goods_detail,omitempty"`
	TimeExpire     *string         `json:"time_expire,omitempty"`
	TimeoutExpress *string         `json:"timeout_express,omitempty"`
	SettleInfo     *SettleInfo     `json:"settle_info,omitempty"`
	RoyaltyInfo    *RoyaltyInfo    `json:"royalty_info,omitempty"`
	ExtendParams   *ExtendParams   `json:"extend_params,omitempty"`
	BusinessParams *BusinessParams `json:"business_params,omitempty"`
	SignParams     *SignParams     `json:"sign_params,omitempty"`
	PassbackParams *string         `json:"passback_params,omitempty"`
}

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

// PreCreateQueryOption 是 alipay.trade.precreate 的返回参数选项，
// 取值依据预下单文档。
type PreCreateQueryOption string

const (
	// PreCreateQueryOptionEnterprisePayInfo 请求返回因公付金额信息。
	PreCreateQueryOptionEnterprisePayInfo PreCreateQueryOption = "enterprise_pay_info"
	// PreCreateQueryOptionHybAmount 请求返回惠营宝回票金额信息。
	PreCreateQueryOptionHybAmount PreCreateQueryOption = "hyb_amount"
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

// PayScene 是 alipay.trade.pay 的支付场景，取值依据官方 V3 文档。
type PayScene string

const (
	// PaySceneBarCode 表示当面付条码支付，auth_code 为买家付款码。
	PaySceneBarCode PayScene = "bar_code"
	// PaySceneSecurityCode 表示当面付刷脸支付，auth_code 为 fp 开头的刷脸标识串。
	PaySceneSecurityCode PayScene = "security_code"
)

// IsKnown 报告场景是否属于文档已知值。
func (s PayScene) IsKnown() bool {
	return s == PaySceneBarCode || s == PaySceneSecurityCode
}

// PayQueryOption 是 alipay.trade.pay / create 的 query_options 已知取值；
// 官方注明支持列表可能扩展，未知值由调用方原样传递。
type PayQueryOption string

const (
	// PayQueryOptionFundBillList 请求返回资金渠道明细。
	PayQueryOptionFundBillList PayQueryOption = "fund_bill_list"
	// PayQueryOptionVoucherDetailList 请求返回优惠券明细。
	PayQueryOptionVoucherDetailList PayQueryOption = "voucher_detail_list"
	// PayQueryOptionEnterprisePayInfo 请求返回因公付信息。
	PayQueryOptionEnterprisePayInfo PayQueryOption = "enterprise_pay_info"
	// PayQueryOptionDiscountGoodsDetail 请求返回单品券优惠商品信息。
	PayQueryOptionDiscountGoodsDetail PayQueryOption = "discount_goods_detail"
	// PayQueryOptionHybAmount 请求返回惠营宝回票金额。
	PayQueryOptionHybAmount PayQueryOption = "hyb_amount"
	// PayQueryOptionAsyncPayInfo 请求返回异步支付受理信息。
	PayQueryOptionAsyncPayInfo PayQueryOption = "async_pay_info"
)

// CancelAction 是撤销接口触发的交易动作，取值依据官方 V3 文档。
type CancelAction string

const (
	// CancelActionClose 表示交易未支付，触发关闭交易动作，无退款。
	CancelActionClose CancelAction = "close"
	// CancelActionRefund 表示交易已支付，触发交易退款动作。
	CancelActionRefund CancelAction = "refund"
)

// BatchResultCode 是批量结算、请款分账与分账关系操作的业务结果码，
// 取值依据官方 V3 文档。
type BatchResultCode string

const (
	// BatchResultCodeAcceptSuccess 表示受理成功。
	BatchResultCodeAcceptSuccess BatchResultCode = "ACCEPT_SUCCESS"
	// BatchResultCodeAcceptFail 表示受理失败。
	BatchResultCodeAcceptFail BatchResultCode = "ACCEPT_FAIL"
	// BatchResultCodeSuccess 表示处理成功。
	BatchResultCodeSuccess BatchResultCode = "SUCCESS"
	// BatchResultCodeFail 表示处理失败。
	BatchResultCodeFail BatchResultCode = "FAIL"
	// BatchResultCodeSystemException 表示系统异常。
	BatchResultCodeSystemException BatchResultCode = "SYSTEM_EXCEPTION"
)

// 通知参数名，严格匹配官方通知参数表。
const (
	notifyParamNotifyTime   = "notify_time"
	notifyParamNotifyType   = "notify_type"
	notifyParamNotifyID     = "notify_id"
	notifyParamTradeNo      = "trade_no"
	notifyParamAppID        = "app_id"
	notifyParamAuthAppID    = "auth_app_id"
	notifyParamOutTradeNo   = "out_trade_no"
	notifyParamOutBizNo     = "out_biz_no"
	notifyParamBuyerID      = "buyer_id"
	notifyParamBuyerOpenID  = "buyer_open_id"
	notifyParamBuyerLogonID = "buyer_logon_id"
	notifyParamSellerID     = "seller_id"
	notifyParamSellerEmail  = "seller_email"
	notifyParamTradeStatus  = "trade_status"
	notifyParamTotalAmount  = "total_amount"
	notifyParamReceipt      = "receipt_amount"
	notifyParamInvoice      = "invoice_amount"
	notifyParamBuyerPay     = "buyer_pay_amount"
	notifyParamPointAmount  = "point_amount"
	notifyParamRefundFee    = "refund_fee"
	notifyParamSendBackFee  = "send_back_fee"
	notifyParamSubject      = "subject"
	notifyParamBody         = "body"
	notifyParamGmtCreate    = "gmt_create"
	notifyParamGmtPayment   = "gmt_payment"
	notifyParamGmtRefund    = "gmt_refund"
	notifyParamGmtClose     = "gmt_close"
	notifyParamFundBillList = "fund_bill_list"
	notifyParamVouchers     = "voucher_detail_list"
	notifyParamSettleMode   = "biz_settle_mode"
)

// Notification 是统一收单交易支付完成异步通知（普通支付通知，
// notify_type=trade_status_sync）的领域模型。
// 金额保持官方通知的十进制字符串；时间为官方格式的字符串，SDK 不解析
// 时间格式；响应中未知枚举值保留原始值。空值通知参数视为未传。
// 仅当 TradeStatus 为 TRADE_SUCCESS 或 TRADE_FINISHED 时表示买家付款成功；
// 其他状态不能当作支付成功处理。
type Notification struct {
	// NotifyTime 是通知的发送时间（yyyy-MM-dd HH:mm:ss）。
	NotifyTime string
	// NotifyType 是通知类型，官方枚举为 trade_status_sync。
	NotifyType NotifyType
	// NotifyID 是通知校验 ID；同一条通知重发时该值不变，
	// 可作为商户幂等去重的依据。
	NotifyID string
	// TradeNo 是支付宝交易号。
	TradeNo string
	// AppID 是开发者的应用 APPID；商户必须校验其为本应用的 APPID。
	AppID string
	// AuthAppID 是授权方应用 APPID，在服务商调用场景下与 AppID 可能不同。
	AuthAppID string
	// OutTradeNo 是商户订单号；商户必须校验其为系统中创建的订单号。
	OutTradeNo string
	// OutBizNo 是商家业务号，主要是退款通知中返回退款申请的流水号。
	OutBizNo *string
	// BuyerID 是买家支付宝用户号；新商户建议使用 BuyerOpenID。
	BuyerID *string
	// BuyerOpenID 是买家支付宝用户唯一标识。
	BuyerOpenID *string
	// BuyerLogonID 是买家支付宝账号（官方脱敏形式）。
	BuyerLogonID *string
	// SellerID 是卖家支付宝用户号；多收款账户的商户必须校验其为
	// 该订单对应的收款方。
	SellerID *string
	// SellerEmail 是卖家支付宝账号。
	SellerEmail *string
	// TradeStatus 是交易状态。
	TradeStatus TradeStatus
	// TotalAmount 是本次交易支付的订单金额；商户必须校验其为订单的
	// 实际金额（订单创建时金额）。
	TotalAmount string
	// ReceiptAmount 是商家在交易中实际收到的款项。
	ReceiptAmount string
	// InvoiceAmount 是用户在交易中支付的可开发票金额。
	InvoiceAmount *string
	// BuyerPayAmount 是用户在交易中支付的金额。
	BuyerPayAmount *string
	// PointAmount 是使用集分宝支付的金额。
	PointAmount *string
	// RefundFee 是退款通知中的总退款金额。
	RefundFee *string
	// SendBackFee 是商家实际退款给用户的金额。
	SendBackFee *string
	// Subject 是订单标题，请求时参数原样通知回来。
	Subject *string
	// Body 是商品描述，请求时参数原样通知回来。
	Body *string
	// GmtCreate 是交易创建时间。
	GmtCreate *string
	// GmtPayment 是买家付款时间。
	GmtPayment *string
	// GmtRefund 是退款时间（官方格式 yyyy-MM-dd HH:mm:ss.SS）。
	GmtRefund *string
	// GmtClose 是交易结束时间。
	GmtClose *string
	// FundBillList 是支付成功的各个渠道金额信息，
	// 由通知中的 JSON 文本解析；通知内字段名为官方规定的驼峰形式。
	FundBillList []NotifyFundBill
	// VoucherDetailList 是本交易支付时所使用的所有优惠券信息，
	// 由通知中的 JSON 文本解析。
	VoucherDetailList []NotifyVoucherDetail
	// BizSettleMode 是账期结算标识，目前在小程序交易组件场景返回。
	BizSettleMode *BizSettleMode
}

// Paid 报告通知是否表示买家付款成功：依据官方说明，仅 TRADE_SUCCESS 与
// TRADE_FINISHED 属于付款成功。其他状态（含未知新状态）返回 false，
// 调用方仍应依据 TradeStatus 的具体语义处理业务。
func (n Notification) Paid() bool {
	return n.TradeStatus == TradeStatusSuccess || n.TradeStatus == TradeStatusFinished
}

// NotifyFundBill 是通知 fund_bill_list JSON 文本中的单条资金渠道信息。
// 通知内字段名为官方规定的驼峰形式（fundChannel），与同步响应
// fund_bill_list 的下划线形式（fund_channel）不同，不能复用同一 struct。
type NotifyFundBill struct {
	// FundChannel 是支付渠道；完整取值见官方支付渠道说明，未知值原样保留。
	FundChannel FundChannel `json:"fundChannel"`
	// Amount 是使用该渠道支付的金额。
	Amount string `json:"amount"`
}

// NotifyVoucherDetail 是通知 voucher_detail_list JSON 文本中的单条
// 优惠券信息。字段名为官方规定的驼峰形式，与同步响应的 VoucherDetail
// 不同；未知优惠券类型保留原值。
type NotifyVoucherDetail struct {
	// VoucherID 是券 ID。
	VoucherID string `json:"voucherId"`
	// TemplateID 是券模板 ID。
	TemplateID *string `json:"templateId,omitempty"`
	// Name 是券名称。
	Name string `json:"name"`
	// Type 是优惠券类型；官方注明可能新增类型，未知值原样保留。
	Type NotifyVoucherType `json:"type"`
	// Amount 是优惠金额中由商家出资的金额。
	Amount string `json:"amount"`
	// MerchantContribute 是商家出资金额。
	MerchantContribute *string `json:"merchantContribute,omitempty"`
	// OtherContribute 是其他出资方出资金额。
	OtherContribute *string `json:"otherContribute,omitempty"`
	// OtherContributeDetail 是优惠券的其他出资方明细。
	OtherContributeDetail []NotifyContributeDetail `json:"otherContributeDetail,omitempty"`
	// Memo 是优惠券备注信息。
	Memo *string `json:"memo,omitempty"`
}

// NotifyContributeDetail 是通知优惠券的出资方明细（驼峰字段名）。
type NotifyContributeDetail struct {
	// ContributeType 是出资方类型；官方注明可能新增类型。
	ContributeType ContributeType `json:"contributeType"`
	// ContributeAmount 是出资方金额。
	ContributeAmount string `json:"contributeAmount"`
}

// NotificationDecoder 返回普通支付通知的领域映射器，供 notify.Parse 使用。
// 官方标注为必选的参数缺失时返回 *core.ValidationError；
// fund_bill_list 与 voucher_detail_list 存在但不是合法 JSON 时同样报错。
func NotificationDecoder() notify.Decoder[Notification] {
	return notify.DecoderFunc[Notification](decodeNotification)
}

// TradeSettleInfo 是 alipay.trade.query 的结算信息。
type TradeSettleInfo struct {
	// TradeSettleDetailList 是已发生的结算操作。
	TradeSettleDetailList []TradeSettleDetail `json:"trade_settle_detail_list,omitempty"`
	// TradeUnsettledAmount 是剩余待结算金额。
	TradeUnsettledAmount *string `json:"trade_unsettled_amount,omitempty"`
}

// TradeSettleDetail 是一项分账、补差或结算操作。
type TradeSettleDetail struct {
	// Amount 是本次操作的金额，单位为元。
	Amount *string `json:"amount,omitempty"`
	// OperationDt 是操作日期。
	OperationDt *string `json:"operation_dt,omitempty"`
	// OperationSerialNo 是商户发起操作时的外部请求号。
	OperationSerialNo *string `json:"operation_serial_no,omitempty"`
	// OperationType 是结算操作类型，保留网关返回的新值。
	OperationType *SettleOperationType `json:"operation_type,omitempty"`
	// OriTransIn 是商户请求的转入账号。
	OriTransIn *string `json:"ori_trans_in,omitempty"`
	// OriTransOut 是商户请求的转出账号。
	OriTransOut *string `json:"ori_trans_out,omitempty"`
	// TransIn 是实际转入账号。
	TransIn *string `json:"trans_in,omitempty"`
	// TransOut 是实际转出账号。
	TransOut *string `json:"trans_out,omitempty"`
}

// SettleOperationType 标识交易结算明细的操作类型。
type SettleOperationType string

const (
	// SettleOperationReplenish 表示补差。
	SettleOperationReplenish SettleOperationType = "replenish"
	// SettleOperationReplenishRefund 表示退补差。
	SettleOperationReplenishRefund SettleOperationType = "replenish_refund"
	// SettleOperationTransfer 表示分账。
	SettleOperationTransfer SettleOperationType = "transfer"
	// SettleOperationTransferRefund 表示退分账。
	SettleOperationTransferRefund SettleOperationType = "transfer_refund"
	// SettleOperationSettle 表示结算。
	SettleOperationSettle SettleOperationType = "settle"
	// SettleOperationSettleRefund 表示退结算。
	SettleOperationSettleRefund SettleOperationType = "settle_refund"
	// SettleOperationOnSettle 表示待结算。
	SettleOperationOnSettle SettleOperationType = "on_settle"
)
