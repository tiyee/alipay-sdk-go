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
	OutTradeNo     string        `json:"out_trade_no"`
	TotalAmount    string        `json:"total_amount"`
	Subject        string        `json:"subject"`
	ProductCode    string        `json:"product_code"`
	Body           *string       `json:"body,omitempty"`
	GoodsDetail    []GoodsDetail `json:"goods_detail,omitempty"`
	TimeExpire     *string       `json:"time_expire,omitempty"`
	TimeoutExpress *string       `json:"timeout_express,omitempty"`
	QRPayMode      *QRPayMode    `json:"qr_pay_mode,omitempty"`
	SettleInfo     *SettleInfo   `json:"settle_info,omitempty"`
	ExtendParams   *ExtendParams `json:"extend_params,omitempty"`
	BusinessParams *string       `json:"business_params,omitempty"`
	PassbackParams *string       `json:"passback_params,omitempty"`
	QRCodeWidth    *int64        `json:"qrcode_width,omitempty"`
	QuitURL        *string       `json:"quit_url,omitempty"`
	ExtUserInfo    *ExtUserInfo  `json:"ext_user_info,omitempty"`
	SubMerchant    *SubMerchant  `json:"sub_merchant,omitempty"`
	InvoiceInfo    *InvoiceInfo  `json:"invoice_info,omitempty"`
	StoreID        *string       `json:"store_id,omitempty"`
	RoyaltyInfo    *RoyaltyInfo  `json:"royalty_info,omitempty"`
}

// appBizContent 是 app.pay 的 biz_content JSON 结构，
// 字段名依据官方接口文档与商家扣款场景文档。
type appBizContent struct {
	OutTradeNo          string        `json:"out_trade_no"`
	TotalAmount         string        `json:"total_amount"`
	Subject             string        `json:"subject"`
	ProductCode         ProductCode   `json:"product_code,omitempty"`
	Body                *string       `json:"body,omitempty"`
	GoodsDetail         []GoodsDetail `json:"goods_detail,omitempty"`
	TimeExpire          *string       `json:"time_expire,omitempty"`
	TimeoutExpress      *string       `json:"timeout_express,omitempty"`
	ExtendParams        *ExtendParams `json:"extend_params,omitempty"`
	AgreementSignParams *SignParams   `json:"agreement_sign_params,omitempty"`
	PassbackParams      *string       `json:"passback_params,omitempty"`
	ExtUserInfo         *ExtUserInfo  `json:"ext_user_info,omitempty"`
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

// PaymentType 指定预下单的支付方式。
type PaymentType string

// PaymentTypeAgentPay 指定智能体 SDK 支付。
const PaymentTypeAgentPay PaymentType = "agent_pay"

// ProductCode 是交易产品码，取值由具体接口决定。
type ProductCode string

const (
	// ProductCodeFaceToFacePayment 表示当面付产品，为默认值。
	ProductCodeFaceToFacePayment ProductCode = "FACE_TO_FACE_PAYMENT"
	// ProductCodeFastInstantTradePay 用于电脑网站支付。
	ProductCodeFastInstantTradePay ProductCode = "FAST_INSTANT_TRADE_PAY"
	// ProductCodeQuickWapWay 用于手机网站支付。
	ProductCodeQuickWapWay ProductCode = "QUICK_WAP_WAY"
	// ProductCodeQuickMSecurityPay 是普通 App 支付的默认产品码。
	ProductCodeQuickMSecurityPay ProductCode = "QUICK_MSECURITY_PAY"
	// ProductCodeCyclePayAuth 用于 App 支付并签约周期扣款。
	ProductCodeCyclePayAuth ProductCode = "CYCLE_PAY_AUTH"
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

// PaymentRoyaltyType 是支付请求 royalty_info 的分账类型。
type PaymentRoyaltyType string

// PaymentRoyaltyOrdinary 表示支付请求中的普通分账。
const PaymentRoyaltyOrdinary PaymentRoyaltyType = "ROYALTY"

// CloseExtendParams 是关闭交易的扩展参数。
type CloseExtendParams struct {
	// IsFulfilled 表示标识是否通过其他方式履约。
	IsFulfilled *string `json:"is_fulfilled,omitempty"`
}

// SpecifiedChannelParams 是指定优先扣款渠道。
type SpecifiedChannelParams struct {
	// AssetTypeCode 表示资产类型编码，具体传参值请联系技术进行分配。
	AssetTypeCode *string `json:"asset_type_code,omitempty"`
	// InstID 表示机构描述。
	InstID *string `json:"inst_id,omitempty"`
}

// SpecifiedDeductRuleParams 是商户指定扣款规则。
type SpecifiedDeductRuleParams struct {
	// SpecifiedMonthLimitAmount 表示用户自定义单月扣款额度，单位元，用于自动充场景下单月最大累计充值金额管控，商户侧传入。
	SpecifiedMonthLimitAmount *string `json:"specified_month_limit_amount,omitempty"`
	// SpecifiedPayAmount 表示签约时指定代扣金额，单位元，传入后后续代扣环节的金额必须等于该金额。
	SpecifiedPayAmount *string `json:"specified_pay_amount,omitempty"`
	// SpecifiedPayTriggerAmount 表示扣款条件额度数值，用于自动充场景下低于该额度时触发扣款，商户侧管控。
	SpecifiedPayTriggerAmount *string `json:"specified_pay_trigger_amount,omitempty"`
	// SpecifiedPayTriggerUnit 表示元/积分/游戏币。
	SpecifiedPayTriggerUnit *string `json:"specified_pay_trigger_unit,omitempty"`
}

// SignMerchantParams 是协议签约的子商户信息。
type SignMerchantParams struct {
	// SubMerchantID 表示子商户的商户id。
	SubMerchantID *string `json:"sub_merchant_id,omitempty"`
	// SubMerchantName 表示子商户的商户名称。
	SubMerchantName *string `json:"sub_merchant_name,omitempty"`
	// SubMerchantServiceDescription 表示子商户的服务描述。
	SubMerchantServiceDescription *string `json:"sub_merchant_service_description,omitempty"`
	// SubMerchantServiceName 表示子商户的服务名称。
	SubMerchantServiceName *string `json:"sub_merchant_service_name,omitempty"`
}

// SceneRuleParams 是周期扣款的场景规则。
type SceneRuleParams struct {
	// DiscountPeriod 表示在周期扣场景化模板中配置优惠类型为优惠期玩法时需要该参数，表示后续有多少期扣款可享受优惠，值为自然数代表周期。
	DiscountPeriod *string `json:"discount_period,omitempty"`
	// LowPricePeriod 表示在周期扣场景化模板中配置优惠类型为低价玩法时需要该参数，表示代扣低价期持续的时间。
	LowPricePeriod *string `json:"low_price_period,omitempty"`
}

// ExtUserInfo 是外部指定买家的身份信息。
type ExtUserInfo struct {
	// CertNo 表示买家证件号。
	CertNo *string `json:"cert_no,omitempty"`
	// CertType 表示指定买家证件类型。
	CertType *string `json:"cert_type,omitempty"`
	// FixBuyer 表示是否强制校验买家身份。
	FixBuyer *string `json:"fix_buyer,omitempty"`
	// IdentityHash 表示买家加密身份信息。
	IdentityHash *string `json:"identity_hash,omitempty"`
	// MinAge 表示允许的最小买家年龄。
	MinAge *string `json:"min_age,omitempty"`
	// Mobile 表示指定买家手机号。
	Mobile *string `json:"mobile,omitempty"`
	// Name 表示指定买家姓名。
	Name *string `json:"name,omitempty"`
	// NeedCheckInfo 表示是否强制校验买家信息。
	NeedCheckInfo *string `json:"need_check_info,omitempty"`
}

// LogisticsDetail 是订单物流信息。
type LogisticsDetail struct {
	// LogisticsType 表示物流类型。
	LogisticsType *string `json:"logistics_type,omitempty"`
}

// ReceiverAddressInfo 是订单收货人及其地址。
type ReceiverAddressInfo struct {
	// Address 表示收货地址。
	Address *string `json:"address,omitempty"`
	// DivisionCode 表示中国标准城市区域码。
	DivisionCode *string `json:"division_code,omitempty"`
	// Mobile 表示收货人手机号。
	Mobile *string `json:"mobile,omitempty"`
	// Name 表示收货人的姓名。
	Name *string `json:"name,omitempty"`
	// Zip 表示收货地址邮编。
	Zip *string `json:"zip,omitempty"`
}

// SubMerchant 是支付宝二级商户信息。
type SubMerchant struct {
	// MerchantID 是支付宝二级商户编号；直付通和机构间连场景必传。
	MerchantID *string `json:"merchant_id,omitempty"`
	// MerchantType 是二级商户编号类型；目前仅支持 alipay。
	MerchantType *string `json:"merchant_type,omitempty"`
}

// EnterprisePayInfo 是因公支付返回信息。
type EnterprisePayInfo struct {
	// BizInfo 表示因公付业务信息。
	BizInfo *string `json:"biz_info,omitempty"`
	// InvoiceAmount 表示开票金额。
	InvoiceAmount *string `json:"invoice_amount,omitempty"`
	// IsUseEnterprisePay 表示是否包含因公付资产。
	IsUseEnterprisePay *bool `json:"is_use_enterprise_pay,omitempty"`
}

// PayParams 是支付请求的异步和重试设置。
type PayParams struct {
	// AsyncType 表示普通异步支付, 传入该参数时，如果满足受理条件，会先同步受理支付，然后在异步调度推进支付。
	AsyncType *string `json:"async_type,omitempty"`
	// RetryType 表示重试类型，当async_type传入NORMAL_ASYNC时，可以设置该参数，选择是否要重试，retry_type 可选，不设置时，可重试。
	RetryType *string `json:"retry_type,omitempty"`
}

// PromoParams 是支付优惠参数。
type PromoParams struct {
	// ActualOrderTime 表示存在延迟扣款这一类的场景，用这个时间表明用户发生交易的时间，比如说，在公交地铁场景，用户刷码出站的时间，和商户上送交易的时间是不一样的。
	ActualOrderTime *string `json:"actual_order_time,omitempty"`
}

// BkAgentRespInfo 是间联机构交易返回信息。
type BkAgentRespInfo struct {
	// BindclrissrID 表示枚举值，01 银联；02 网联；03 连通等。
	BindclrissrID *string `json:"bindclrissr_id,omitempty"`
	// BindpyeracctbkID 表示付款机构在清算组织登记或分配的机构代码。
	BindpyeracctbkID *string `json:"bindpyeracctbk_id,omitempty"`
	// BindtrxID 表示原快捷交易流水号。
	BindtrxID *string `json:"bindtrx_id,omitempty"`
	// BkpyeruserCode 表示用户在银行付款账号的标记化处理编号。
	BkpyeruserCode *string `json:"bkpyeruser_code,omitempty"`
	// EstterLocation 表示设备推测位置。
	EstterLocation *string `json:"estter_location,omitempty"`
}

// ChargeInfo 是交易收费信息。
type ChargeInfo struct {
	// ChargeFee 表示实收费用。
	ChargeFee *string `json:"charge_fee,omitempty"`
	// ChargeType 表示收单手续费trade，花呗分期手续hbfq，其他手续费charge。
	ChargeType *string `json:"charge_type,omitempty"`
	// IsRatingOnSwitch 表示是否合约指定收费账号出资，值为"Y"或"N"。
	IsRatingOnSwitch *string `json:"is_rating_on_switch,omitempty"`
	// IsRatingOnTradeReceiver 表示是否收款账号出资，值为"Y"或"N"。
	IsRatingOnTradeReceiver *string `json:"is_rating_on_trade_receiver,omitempty"`
	// OriginalChargeFee 表示原始费用。
	OriginalChargeFee *string `json:"original_charge_fee,omitempty"`
	// SubFeeDetailList 表示组合支付收费明细。
	SubFeeDetailList []SubFee `json:"sub_fee_detail_list,omitempty"`
	// SwitchFeeRate 表示签约费率。
	SwitchFeeRate *string `json:"switch_fee_rate,omitempty"`
}

// SubFee 是收费分项明细。
type SubFee struct {
	// ChargeFee 表示实收费用。
	ChargeFee *string `json:"charge_fee,omitempty"`
	// OriginalChargeFee 表示原始费用。
	OriginalChargeFee *string `json:"original_charge_fee,omitempty"`
	// SwitchFeeRate 表示签约费率。
	SwitchFeeRate *string `json:"switch_fee_rate,omitempty"`
}

// TapPayInfo 是碰一下支付的支付介质和优惠信息。
type TapPayInfo struct {
	// PaymentMediumType 表示碰一下支付的支付介质类型，标识当前支付为碰一下支付。
	PaymentMediumType *string `json:"payment_medium_type,omitempty"`
	// TotalDiscountAmount 表示碰一下支付的碰一下共减优惠金额。
	TotalDiscountAmount *string `json:"total_discount_amount,omitempty"`
	// TotalDiscountName 表示碰一下支付的碰一下共减金额展示名称。
	TotalDiscountName *string `json:"total_discount_name,omitempty"`
}

// ExternalPayerInfo 是外部付款方身份信息。
type ExternalPayerInfo struct {
	// OutInstCardNo 表示机构卡信息。
	OutInstCardNo *string `json:"out_inst_card_no,omitempty"`
	// OutInstPayerInstID 表示机构编码。
	OutInstPayerInstID *string `json:"out_inst_payer_inst_id,omitempty"`
	// OutPayerAccountName 表示代理企业支付中，企业的名称。
	OutPayerAccountName *string `json:"out_payer_account_name,omitempty"`
}

// FulfillmentDetail 是订单履约信息。
type FulfillmentDetail struct {
	// FulfillmentAmount 表示履约金额。
	FulfillmentAmount *string `json:"fulfillment_amount,omitempty"`
	// GmtPayment 表示履约支付时间。
	GmtPayment *string `json:"gmt_payment,omitempty"`
	// OutRequestNo 表示商户发起履约请求时，传入的out_request_no，标识一次请求的唯一id。
	OutRequestNo *string `json:"out_request_no,omitempty"`
}

// HbFqPayInfo 是花呗分期付款信息。
type HbFqPayInfo struct {
	// FqAmount 表示用户使用花呗分期支付的金额数。
	FqAmount *string `json:"fq_amount,omitempty"`
	// FqInstID 表示提供分期服务的机构ID。
	FqInstID *string `json:"fq_inst_id,omitempty"`
	// UserInstallNum 表示用户使用花呗分期支付的分期数。
	UserInstallNum *string `json:"user_install_num,omitempty"`
}

// IntactChargeInfo 是订单的完整收费记录。
type IntactChargeInfo struct {
	// ActualAmount 表示实际收费金额，单位元。
	ActualAmount *string `json:"actual_amount,omitempty"`
	// BillType 表示收费类型。
	BillType *string `json:"bill_type,omitempty"`
	// GmtPay 表示收费时间,时间精确到秒。
	GmtPay *string `json:"gmt_pay,omitempty"`
	// IsRefund 表示是否退费。
	IsRefund *bool `json:"is_refund,omitempty"`
	// OutBizNo 表示外部请求号。
	OutBizNo *string `json:"out_biz_no,omitempty"`
	// PlanAmount 表示应收费金额，单位元。
	PlanAmount *string `json:"plan_amount,omitempty"`
	// ProductName 表示收费产品。
	ProductName *string `json:"product_name,omitempty"`
	// ServiceTarget 表示收费唯一id。
	ServiceTarget *string `json:"service_target,omitempty"`
	// ServiceType 表示收费类型。
	ServiceType *string `json:"service_type,omitempty"`
	// Status 表示状态。
	Status *string `json:"status,omitempty"`
	// TargetAccountNo 表示收费目标账号。
	TargetAccountNo *string `json:"target_account_no,omitempty"`
	// TargetUserID 表示收费目标uid。
	TargetUserID *string `json:"target_user_id,omitempty"`
}

// PaymentInfoWithID 是支付类型及其支付单号。
type PaymentInfoWithID struct {
	// PaymentIDs 表示支付单id。
	PaymentIDs []string `json:"payment_ids,omitempty"`
	// Type 表示支付资金类型。
	Type *string `json:"type,omitempty"`
}

// QRPayMode 是电脑网站支付的扫码方式。
type QRPayMode string

const (
	// QRPaySimple 使用简约前置二维码。
	QRPaySimple QRPayMode = "0"
	// QRPayEmbedded 使用前置二维码。
	QRPayEmbedded QRPayMode = "1"
	// QRPayRedirect 跳转支付宝扫码页面。
	QRPayRedirect QRPayMode = "2"
	// QRPayMini 使用迷你前置二维码。
	QRPayMini QRPayMode = "3"
	// QRPayCustomWidth 允许商户设置二维码宽度。
	QRPayCustomWidth QRPayMode = "4"
)

// InvoiceInfo 是电脑网站支付的开票信息。
type InvoiceInfo struct {
	// KeyInfo 是商户开票信息。
	KeyInfo InvoiceKeyInfo `json:"key_info"`
	// Details 是支付宝约定格式的发票明细字符串。
	Details string `json:"details"`
}

// InvoiceKeyInfo 是商户的开票标识和税号。
type InvoiceKeyInfo struct {
	// IsSupportInvoice 表示商户是否支持开票。
	IsSupportInvoice bool `json:"is_support_invoice"`
	// InvoiceMerchantName 是开票商户名称。
	InvoiceMerchantName string `json:"invoice_merchant_name"`
	// TaxNum 是商户纳税人识别号。
	TaxNum string `json:"tax_num"`
}
