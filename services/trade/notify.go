package trade

import (
	"encoding/json"
	"net/url"

	"github.com/tiyee/go-alipay/core"
	"github.com/tiyee/go-alipay/core/notify"
)

// 普通支付通知（统一收单交易支付完成异步通知）的领域模型与映射。
// 参数名与必选性依据官方《异步通知说明》
// https://opendocs.alipay.com/open/194/103296（2026-11-05 核验）；
// 验签规则依据 https://opendocs.alipay.com/common/02mse7.md，
// 由 core/notify 统一执行，本包只负责把通过验签的参数映射为具名 struct。

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

func decodeNotification(values url.Values) (Notification, error) {
	var n Notification
	var err error
	for _, pair := range []struct {
		target *string
		name   string
	}{
		{&n.NotifyTime, notifyParamNotifyTime},
		{&n.NotifyID, notifyParamNotifyID},
		{&n.TradeNo, notifyParamTradeNo},
		{&n.AppID, notifyParamAppID},
		{&n.AuthAppID, notifyParamAuthAppID},
		{&n.OutTradeNo, notifyParamOutTradeNo},
		{&n.TotalAmount, notifyParamTotalAmount},
		{&n.ReceiptAmount, notifyParamReceipt},
	} {
		if *pair.target, err = requiredParam(values, pair.name); err != nil {
			return Notification{}, err
		}
	}
	if err := requiredEnum(values, notifyParamTradeStatus, &n.TradeStatus); err != nil {
		return Notification{}, err
	}
	if err := requiredEnum(values, notifyParamNotifyType, &n.NotifyType); err != nil {
		return Notification{}, err
	}

	n.OutBizNo = optionalParam(values, notifyParamOutBizNo)
	n.BuyerID = optionalParam(values, notifyParamBuyerID)
	n.BuyerOpenID = optionalParam(values, notifyParamBuyerOpenID)
	n.BuyerLogonID = optionalParam(values, notifyParamBuyerLogonID)
	n.SellerID = optionalParam(values, notifyParamSellerID)
	n.SellerEmail = optionalParam(values, notifyParamSellerEmail)
	n.InvoiceAmount = optionalParam(values, notifyParamInvoice)
	n.BuyerPayAmount = optionalParam(values, notifyParamBuyerPay)
	n.PointAmount = optionalParam(values, notifyParamPointAmount)
	n.RefundFee = optionalParam(values, notifyParamRefundFee)
	n.SendBackFee = optionalParam(values, notifyParamSendBackFee)
	n.Subject = optionalParam(values, notifyParamSubject)
	n.Body = optionalParam(values, notifyParamBody)
	n.GmtCreate = optionalParam(values, notifyParamGmtCreate)
	n.GmtPayment = optionalParam(values, notifyParamGmtPayment)
	n.GmtRefund = optionalParam(values, notifyParamGmtRefund)
	n.GmtClose = optionalParam(values, notifyParamGmtClose)
	if mode := values.Get(notifyParamSettleMode); mode != "" {
		bizSettleMode := BizSettleMode(mode)
		n.BizSettleMode = &bizSettleMode
	}
	if raw := values.Get(notifyParamFundBillList); raw != "" {
		if err := json.Unmarshal([]byte(raw), &n.FundBillList); err != nil {
			return Notification{}, &core.ValidationError{Field: notifyParamFundBillList,
				Reason: "is not a valid JSON array"}
		}
	}
	if raw := values.Get(notifyParamVouchers); raw != "" {
		if err := json.Unmarshal([]byte(raw), &n.VoucherDetailList); err != nil {
			return Notification{}, &core.ValidationError{Field: notifyParamVouchers,
				Reason: "is not a valid JSON array"}
		}
	}
	return n, nil
}

// requiredParam 读取必选通知参数；缺失时返回 *core.ValidationError。
func requiredParam(values url.Values, name string) (string, error) {
	value := values.Get(name)
	if value == "" {
		return "", &core.ValidationError{Field: name, Reason: "required notify param is missing"}
	}
	return value, nil
}

// requiredEnum 读取必选枚举参数；未知响应值保留原始字符串。
func requiredEnum[T ~string](values url.Values, name string, target *T) error {
	value, err := requiredParam(values, name)
	if err != nil {
		return err
	}
	*target = T(value)
	return nil
}

// optionalParam 读取可选通知参数；空值视为未传。
func optionalParam(values url.Values, name string) *string {
	if value := values.Get(name); value != "" {
		return &value
	}
	return nil
}
