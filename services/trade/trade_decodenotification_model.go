package trade

import (
	"context"
	"encoding/json"
	"net/url"

	"github.com/tiyee/alipay-sdk-go/core"
)

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

// TradeFundBill 是交易支付使用的资金渠道明细。
type TradeFundBill struct {
	// FundChannel 是交易使用的资金渠道；完整取值见官方支付渠道列表。
	FundChannel FundChannel `json:"fund_channel"`
	// Amount 是该支付工具类型所使用的金额。
	Amount string `json:"amount"`
	// RealAmount 是渠道实际付款金额。
	RealAmount *string `json:"real_amount,omitempty"`
	// FundType 是渠道所使用的资金类型，仅银行卡渠道返回。
	FundType *FundType `json:"fund_type,omitempty"`
	// BankCode 表示银行卡支付时的银行代码。
	BankCode *string `json:"bank_code,omitempty"`
}

// GoodsDetail 是预下单与支付的商品明细。
type GoodsDetail struct {
	// GoodsID 是商品编号（最小粒度 ID）。
	GoodsID string `json:"goods_id"`
	// GoodsName 是商品名称。
	GoodsName string `json:"goods_name"`
	// Quantity 是商品数量。
	Quantity int64 `json:"quantity"`
	// Price 是商品单价，单位为元。
	Price string `json:"price"`
	// GoodsCategory 是商品类目。
	GoodsCategory *string `json:"goods_category,omitempty"`
	// CategoriesTree 是商品类目树，类目 id 用 | 分割。
	CategoriesTree *string `json:"categories_tree,omitempty"`
	// ShowURL 是商品的展示地址。
	ShowURL *string `json:"show_url,omitempty"`
	// OutItemID 是商家侧小程序商品 ID。
	OutItemID *string `json:"out_item_id,omitempty"`
	// OutSkuID 是商家侧小程序商品 SKU ID。
	OutSkuID *string `json:"out_sku_id,omitempty"`
	// AlipayGoodsID 表示支付宝定义的统一商品编号。
	AlipayGoodsID *string `json:"alipay_goods_id,omitempty"`
	// Body 表示商品描述信息。
	Body *string `json:"body,omitempty"`
	// GoodsEANID 表示商品69条码。
	GoodsEANID *string `json:"goods_ean_id,omitempty"`
}

// SettleInfo 描述结算信息；目前只支持一条结算明细。
type SettleInfo struct {
	// SettleDetailInfos 是结算详细信息，目前只支持一条。
	SettleDetailInfos []SettleDetailInfo `json:"settle_detail_infos"`
	// SettlePeriodTime 表示该笔订单的超期自动确认结算时间，到达期限后，将自动确认结算。
	SettlePeriodTime *string `json:"settle_period_time,omitempty"`
}

// SettleDetailInfo 是单条结算明细。
type SettleDetailInfo struct {
	// TransInType 是结算收款方的账户类型。
	TransInType AccountType `json:"trans_in_type"`
	// TransIn 是结算收款方；类型为 defaultSettle 时不能传值。
	TransIn *string `json:"trans_in,omitempty"`
	// Amount 是结算的金额，创建订单时必须与交易金额相同。
	Amount *string `json:"amount,omitempty"`
	// SummaryDimension 是结算汇总维度，需与 cardAliasNo 配合使用。
	SummaryDimension *string `json:"summary_dimension,omitempty"`
	// SettleEntityID 是结算主体标识。
	SettleEntityID *string `json:"settle_entity_id,omitempty"`
	// SettleEntityType 是结算主体类型：SecondMerchant 或 Store。
	SettleEntityType *string `json:"settle_entity_type,omitempty"`
	// ActualAmount 仅用于直付通账期模式的部分确认结算。
	ActualAmount *string `json:"actual_amount,omitempty"`
}

// ExtendParams 是预下单的业务扩展参数。
type ExtendParams struct {
	// SysServiceProviderID 是系统商编号（返佣数据提取依据）。
	SysServiceProviderID *string `json:"sys_service_provider_id,omitempty"`
	// CardType 是卡类型。
	CardType *string `json:"card_type,omitempty"`
	// SpecifiedSellerName 允许商户指定交易展示的卖家名称。
	SpecifiedSellerName *string `json:"specified_seller_name,omitempty"`
	// RoyaltyFreeze 为 true 时进行资金冻结，用于后续分账。
	RoyaltyFreeze *string `json:"royalty_freeze,omitempty"`
	// TCInstallmentOrderID 是公域商品交易分期单 ID。
	TCInstallmentOrderID *string `json:"tc_installment_order_id,omitempty"`
	// TradeComponentOrderID 是公域商品交易业务订单 ID。
	TradeComponentOrderID *string `json:"trade_component_order_id,omitempty"`
	// CreditExtInfo 表示信用参数，可选，如有需要请与芝麻约定后传入。
	CreditExtInfo *string `json:"credit_ext_info,omitempty"`
	// FqChannels 表示分期渠道。
	FqChannels *string `json:"fq_channels,omitempty"`
	// FqNum 表示使用花呗分期的期数。
	FqNum *string `json:"fq_num,omitempty"`
	// FqSellerPercent 表示花呗分期卖家承担收费比例。
	FqSellerPercent *string `json:"fq_seller_percent,omitempty"`
	// HbFqNum 表示使用花呗分期要进行的分期数。
	HbFqNum *string `json:"hb_fq_num,omitempty"`
	// HbFqSellerPercent 表示使用花呗分期需要卖家承担的手续费比例的百分值，传入100代表100%。
	HbFqSellerPercent *string `json:"hb_fq_seller_percent,omitempty"`
	// IndustryRefluxInfo 表示行业数据回流信息, 详见：地铁支付接口参数补充说明。
	IndustryRefluxInfo *string `json:"industry_reflux_info,omitempty"`
}

// BusinessParams 是商户与支付宝约定的直传业务信息。
type BusinessParams struct {
	// CampusCard 是校园卡编号。
	CampusCard *string `json:"campus_card,omitempty"`
	// CardType 是虚拟卡卡类型。
	CardType *string `json:"card_type,omitempty"`
	// ActualOrderTime 是实际订单时间（乘车码场景为刷码乘车时间）。
	ActualOrderTime *string `json:"actual_order_time,omitempty"`
	// GoodTaxes 是商户传入的交易税费。
	GoodTaxes *string `json:"good_taxes,omitempty"`
	// EnterprisePayInfo 是因公付业务信息（JSON 文本）。
	EnterprisePayInfo *string `json:"enterprise_pay_info,omitempty"`
	// EnterprisePayAmount 是因公付金额。
	EnterprisePayAmount *string `json:"enterprise_pay_amount,omitempty"`
	// MCCreateTradeIP 是商户端创建订单的 IP（ipv4/ipv6）。
	MCCreateTradeIP *string `json:"mc_create_trade_ip,omitempty"`
	// TinyAppMerchantBizType 是小程序细分业务场景。
	TinyAppMerchantBizType *string `json:"tiny_app_merchant_biz_type,omitempty"`
	// BuyerMemo 表示买家备注信息。
	BuyerMemo *string `json:"buyer_memo,omitempty"`
	// ForcePwd 表示收银台强制密码核身参数，实际收银侧会根据这个参数为true并结合业务场景来判断是否密码核身。
	ForcePwd *string `json:"force_pwd,omitempty"`
	// PriorityUnfreezeAsset 表示信用预授权转交易阶段，优先解冻的用户资产。
	PriorityUnfreezeAsset *string `json:"priority_unfreeze_asset,omitempty"`
}

// BkAgentReqInfo 是间联交易下由收单机构上送的信息。
type BkAgentReqInfo struct {
	// AcqCode 是收单机构在清算组织登记或分配的机构代码。
	AcqCode *string `json:"acq_code,omitempty"`
	// MerchCode 是收单机构在清算组织登记的商户编码。
	MerchCode *string `json:"merch_code,omitempty"`
	// DeviceType 是终端设备类型（如 01 ATM、02 传统 POS、03 mPOS）。
	DeviceType *string `json:"device_type,omitempty"`
	// SerialNum 是终端序列号；特定设备类型下必填。
	SerialNum *string `json:"serial_num,omitempty"`
	// Location 是终端设备实时经纬度，格式为纬度/经度。
	Location *string `json:"location,omitempty"`
}

// RoyaltyDetailInfo 是分账/退分账明细（OpenApiRoyaltyDetailInfoPojo）。
type RoyaltyDetailInfo struct {
	// RoyaltyType 是分账类型；为空默认为普通分账 transfer。
	RoyaltyType *RoyaltyType `json:"royalty_type,omitempty"`
	// TransOut 是支出方账户。
	TransOut *string `json:"trans_out,omitempty"`
	// TransOutType 是支出方账户类型。
	TransOutType *AccountType `json:"trans_out_type,omitempty"`
	// TransInType 是收入方账户类型。
	TransInType *AccountType `json:"trans_in_type,omitempty"`
	// TransIn 是收入方账户。
	TransIn *string `json:"trans_in,omitempty"`
	// TransInID 是分账转入方的支付宝 userId。
	TransInID *string `json:"trans_in_id,omitempty"`
	// Amount 是分账的金额。
	Amount *string `json:"amount,omitempty"`
	// AmountPercentage 是分账百分比，取值范围 (0,100] 的整数。
	AmountPercentage *int64 `json:"amount_percentage,omitempty"`
	// Desc 是分账描述。
	Desc *string `json:"desc,omitempty"`
	// RoyaltyScene 是分账场景（达人佣金、平台服务费等）。
	RoyaltyScene *string `json:"royalty_scene,omitempty"`
	// TransInName 是分账收款方姓名；上送则进行一致性校验。
	TransInName *string `json:"trans_in_name,omitempty"`
}

// VoucherDetail 是优惠券信息。
type VoucherDetail struct {
	// ID 是券 id。
	ID string `json:"id"`
	// Name 是券名称。
	Name string `json:"name"`
	// Type 是券类型；官方注明可能新增类型，未知值原样保留。
	Type VoucherType `json:"type"`
	// Amount 是优惠券面额，应等于商家出资加上其他出资方出资。
	Amount string `json:"amount"`
	// MerchantContribute 是商家出资金额。
	MerchantContribute *string `json:"merchant_contribute,omitempty"`
	// OtherContribute 是其他出资方出资金额。
	OtherContribute *string `json:"other_contribute,omitempty"`
	// Memo 是优惠券备注信息。
	Memo *string `json:"memo,omitempty"`
	// TemplateID 是券模板 id。
	TemplateID *string `json:"template_id,omitempty"`
	// OtherContributeDetail 是优惠券的其他出资方明细。
	OtherContributeDetail []ContributeDetail `json:"other_contribute_detail,omitempty"`
	// PurchaseBuyerContribute 是用户购买该券实际付款金额。
	PurchaseBuyerContribute *string `json:"purchase_buyer_contribute,omitempty"`
	// PurchaseMerchantContribute 是购买该券时商户优惠金额。
	PurchaseMerchantContribute *string `json:"purchase_merchant_contribute,omitempty"`
	// PurchaseAntContribute 是购买该券时平台优惠金额。
	PurchaseAntContribute *string `json:"purchase_ant_contribute,omitempty"`
}

// ContributeDetail 是优惠券的出资方明细。
type ContributeDetail struct {
	// ContributeType 是出资方类型；官方注明可能新增类型。
	ContributeType ContributeType `json:"contribute_type"`
	// ContributeAmount 是出资方金额。
	ContributeAmount string `json:"contribute_amount"`
}

// AgreementParams 是代扣/周期扣款场景的协议参数。
type AgreementParams struct {
	// AgreementNo 是支付宝系统中唯一标识用户签约记录的协议号。
	AgreementNo string `json:"agreement_no"`
	// ApplyToken 是鉴权申请 token；需要支付鉴权校验时不能为空。
	ApplyToken *string `json:"apply_token,omitempty"`
	// AuthConfirmNo 是鉴权确认码；需要支付鉴权校验时不能为空。
	AuthConfirmNo *string `json:"auth_confirm_no,omitempty"`
	// DeductPermission 是商户代扣扣款许可。
	DeductPermission *string `json:"deduct_permission,omitempty"`
}

// RoyaltyInfo 描述分账信息。
type RoyaltyInfo struct {
	// RoyaltyType 是分账类型；目前只支持 ROYALTY。
	RoyaltyType *PaymentRoyaltyType `json:"royalty_type,omitempty"`
	// RoyaltyDetailInfos 是分账明细信息，可描述多条分账指令。
	RoyaltyDetailInfos []PaymentRoyaltyDetail `json:"royalty_detail_infos,omitempty"`
}

// SignParams 是支付并签约场景的签约参数（官方 SignParams 模型子集）。
// 字段与必填关系依据官方商家扣款场景文档。
type SignParams struct {
	// ProductCode 是商家和支付宝签约的产品码；商家扣款场景传
	// GENERAL_WITHHOLDING。
	ProductCode *string `json:"product_code,omitempty"`
	// PersonalProductCode 是个人签约产品码。
	PersonalProductCode *string `json:"personal_product_code,omitempty"`
	// SignScene 是协议签约场景。
	SignScene *string `json:"sign_scene,omitempty"`
	// ExternalAgreementNo 是商户签约号，代扣协议中标示用户的唯一签约号。
	ExternalAgreementNo *string `json:"external_agreement_no,omitempty"`
	// ExternalLogonID 是用户在商户网站的登录账号。
	ExternalLogonID *string `json:"external_logon_id,omitempty"`
	// AccessParams 是签约接入参数（如 ALIPAYAPP、QRCODE）。
	AccessParams *SignAccessParams `json:"access_params,omitempty"`
	// PeriodRuleParams 是周期扣款规则参数。
	PeriodRuleParams *PeriodRuleParams `json:"period_rule_params,omitempty"`
	// SignNotifyURL 是签约成功后接收异步通知的地址。
	SignNotifyURL *string `json:"sign_notify_url,omitempty"`
	// EffectTime 是签约请求的有效时间（秒）。
	EffectTime *string `json:"effect_time,omitempty"`
	// AllowHuazhiDegrade 表示是否允许花芝GO降级成原代扣（即销售方案指定的代扣产品），在花芝GO场景下才会使用该值。
	AllowHuazhiDegrade *bool `json:"allow_huazhi_degrade,omitempty"`
	// RecommendSortChannelParams 表示商户指定优先渠道扣款参数。
	RecommendSortChannelParams []SpecifiedChannelParams `json:"recommend_sort_channel_params,omitempty"`
	// SpecifiedDeductRuleParams 是商户指定的扣款规则。
	SpecifiedDeductRuleParams *SpecifiedDeductRuleParams `json:"specified_deduct_rule_params,omitempty"`
	// SubMerchant 是二级商户信息。
	SubMerchant *SignMerchantParams `json:"sub_merchant,omitempty"`
}

// SignAccessParams 描述签约渠道。
type SignAccessParams struct {
	// Channel 是签约渠道，如 ALIPAYAPP（钱包 h5 页面签约）、
	// QRCODE（扫码签约）、QRCODEORSMS（扫码或短信签约）。
	Channel *string `json:"channel,omitempty"`
}

// PeriodRuleParams 是周期扣款产品的扣款周期参数。
type PeriodRuleParams struct {
	// PeriodType 是周期类型：DAY 或 MONTH。
	PeriodType *string `json:"period_type,omitempty"`
	// Period 是周期数；与 PeriodType 组合确定扣款周期。
	Period *int64 `json:"period,omitempty"`
	// ExecuteTime 是首次执行时间（yyyy-MM-dd）；周期扣款产品必填。
	ExecuteTime *string `json:"execute_time,omitempty"`
	// SingleAmount 是单次扣款最大金额，单位为元。
	SingleAmount *string `json:"single_amount,omitempty"`
	// TotalAmount 是总金额限制，单位为元。
	TotalAmount *string `json:"total_amount,omitempty"`
	// TotalPayments 是总扣款次数。
	TotalPayments *int64 `json:"total_payments,omitempty"`
	// IsTrialMode 标识本次签约是否有试用期（Y/N）。
	IsTrialMode *string `json:"is_trial_mode,omitempty"`
	// SceneRuleParams 是周期扣款的场景规则。
	SceneRuleParams *SceneRuleParams `json:"scene_rule_params,omitempty"`
}

// PaymentProducts 构建电脑网站支付（page.pay）、手机网站支付
// （wap.pay）与 App 支付（app.pay）的签名产物。构建过程不发起
// HTTP 请求；支付结果以异步通知与查询接口为准，跳转页与客户端
// 返回不构成支付完成的可信证据。
//
// 通过 NewPaymentProducts 以凭证构建；使用与 [core.Client] 相同的
// 公共选项选择环境：WithEnvironment 会成对设置 REST 网关与产物
// 提交网关，产物网关也可经 option.WithArtifactGateway 单独指定。
type PaymentProducts struct {
	builder *core.ArtifactBuilder
}

// PagePayURL 构建电脑网站支付的跳转 URL（GET 方式）：
// 公共参数与 biz_content 全部置于 query，用户浏览器打开后跳转
// 支付宝收银台。
//
// 契约：官方 V3 文档
// https://opendocs.alipay.com/open-v3/2423fad5_alipay.trade.page.pay。
func (p *PaymentProducts) PagePayURL(ctx context.Context, req PagePayRequest,
	opts ...core.CallOption) (string, error) {
	artifact, err := p.build(ctx, methodPagePay, req, opts...)
	if err != nil {
		return "", err
	}
	return artifact.RedirectURL(), nil
}

// PagePayForm 构建电脑网站支付的自动提交 HTML 表单（POST 方式）：
// 公共参数位于 action 的 query，业务参数（biz_content 等）作为
// 隐藏域；表单渲染后自动提交。
//
// 契约：官方 V3 文档
// https://opendocs.alipay.com/open-v3/2423fad5_alipay.trade.page.pay。
func (p *PaymentProducts) PagePayForm(ctx context.Context, req PagePayRequest,
	opts ...core.CallOption) (string, error) {
	artifact, err := p.build(ctx, methodPagePay, req, opts...)
	if err != nil {
		return "", err
	}
	return artifact.FormHTML(), nil
}

// WapPayURL 构建手机网站支付的跳转 URL（GET 方式）。
//
// 契约：官方 V3 文档
// https://opendocs.alipay.com/open-v3/1a957be0_alipay.trade.wap.pay。
func (p *PaymentProducts) WapPayURL(ctx context.Context, req PagePayRequest,
	opts ...core.CallOption) (string, error) {
	artifact, err := p.build(ctx, methodWapPay, req, opts...)
	if err != nil {
		return "", err
	}
	return artifact.RedirectURL(), nil
}

// WapPayForm 构建手机网站支付的自动提交 HTML 表单（POST 方式）。
//
// 契约：官方 V3 文档
// https://opendocs.alipay.com/open-v3/1a957be0_alipay.trade.wap.pay。
func (p *PaymentProducts) WapPayForm(ctx context.Context, req PagePayRequest,
	opts ...core.CallOption) (string, error) {
	artifact, err := p.build(ctx, methodWapPay, req, opts...)
	if err != nil {
		return "", err
	}
	return artifact.FormHTML(), nil
}

// AppPayOrderString 构建 App 支付订单串（orderStr）：
// 调用方将订单串交给支付宝客户端 SDK 唤起支付。
// 支付并签约场景经 AgreementSignParams 传入签约参数，字段与必填关系依据
// 官方商家扣款场景文档（https://opendocs.alipay.com/open-v3/0jf33m）。
//
// 契约：官方 V3 文档
// https://opendocs.alipay.com/open-v3/429e4d75_alipay.trade.app.pay。
func (p *PaymentProducts) AppPayOrderString(ctx context.Context, req AppPayRequest,
	opts ...core.CallOption) (string, error) {
	if err := req.validate(); err != nil {
		return "", err
	}
	bizParams, err := req.artifactParams()
	if err != nil {
		return "", err
	}
	artifact, err := p.builder.BuildArtifact(ctx, methodAppPay, bizParams, core.NewCallSettings(opts...))
	if err != nil {
		return "", err
	}
	return artifact.OrderString(), nil
}

func (p *PaymentProducts) build(ctx context.Context, method string, req PagePayRequest,
	opts ...core.CallOption) (*core.Artifact, error) {
	if err := req.validate(method); err != nil {
		return nil, err
	}
	bizParams, err := req.artifactParams()
	if err != nil {
		return nil, err
	}
	return p.builder.BuildArtifact(ctx, method, bizParams, core.NewCallSettings(opts...))
}

// PaymentRoyaltyDetail 是支付分账指令（官方 RoyaltyDetailInfos）。
type PaymentRoyaltyDetail struct {
	// Amount 是分账金额，单位为元。
	Amount *string `json:"amount,omitempty"`
	// AmountPercentage 是分账比例，例如 "20" 表示 20%。
	AmountPercentage *string `json:"amount_percentage,omitempty"`
	// BatchNo 是配合 bankIndex 使用的分账批次号。
	BatchNo *string `json:"batch_no,omitempty"`
	// Desc 是分账说明。
	Desc *string `json:"desc,omitempty"`
	// OutRelationID 是商户为该笔分账分配的唯一关联号。
	OutRelationID *string `json:"out_relation_id,omitempty"`
	// SerialNo 是分账执行顺序，取正整数。
	SerialNo *int64 `json:"serial_no,omitempty"`
	// TransIn 是接收分账金额的账户。
	TransIn *string `json:"trans_in,omitempty"`
	// TransInType 是转入账户类型。
	TransInType *AccountType `json:"trans_in_type,omitempty"`
	// TransOut 是转出分账金额的账户。
	TransOut *string `json:"trans_out,omitempty"`
	// TransOutType 是转出账户类型。
	TransOutType *AccountType `json:"trans_out_type,omitempty"`
}
