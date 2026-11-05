package ordervoucher

// ActivityBaseInfo 是支付宝官方 V3 OpenAPI 模型。
type ActivityBaseInfo struct {
	ActivityDeliveryExtInfo *ActivityDeliveryExtInfo `json:"activity_delivery_ext_info,omitempty"`
	// ActivityId 是活动id
	ActivityId *string `json:"activity_id,omitempty"`
	// ActivityName 是活动名称。
	ActivityName *string `json:"activity_name,omitempty"`
	// ActivityOperationStatus 是活动操作状态。
	ActivityOperationStatus *string `json:"activity_operation_status,omitempty"`
	// ActivityProductType 是券活动产品类型
	ActivityProductType *string `json:"activity_product_type,omitempty"`
	// ActivityStatus 是活动状态 。 ACTIVE:活动已激活，表示活动已经生效，等到活动开始(publish_start_time)之后用户就可以参与活动。  PAUSE:活动已暂停，表示商户临时暂停该活动，该状态下用户不能参与活动。
	ActivityStatus     *string             `json:"activity_status,omitempty"`
	BelongMerchantInfo *BelongMerchantInfo `json:"belong_merchant_info,omitempty"`
	// CodeMode 是码模式。   MERCHANT_UPLOAD：商户上传自定义code，发券时系统随机选取上传的券code发放  MERCHANT_API：发奖时指定券码发奖，此模式无须提前上传券码。
	CodeMode *string `json:"code_mode,omitempty"`
	// OutActivityId 是外部活动ID
	OutActivityId *string `json:"out_activity_id,omitempty"`
}

// ActivityDeliveryExtInfo 是支付宝官方 V3 OpenAPI 模型。
type ActivityDeliveryExtInfo struct {
	// AutoDelivery 是是否自动推广
	AutoDelivery *bool `json:"auto_delivery,omitempty"`
	// ChannelList 是推广渠道列表
	ChannelList *[]string `json:"channel_list,omitempty"`
}

// MarketingActivityOrderVoucherAppend 是支付宝官方 V3 OpenAPI 模型。
type MarketingActivityOrderVoucherAppend struct {
	// MerchantAccessMode 是商户接入模式
	MerchantAccessMode *string `json:"merchant_access_mode,omitempty"`
	// OutBizNo 是外部业务单号，用作幂等控制。		幂等作用：	参数不变的情况下，再次请求返回与上一次相同的结果。		外部接入方需保证业务单号唯一。
	OutBizNo *string `json:"out_biz_no,omitempty"`
	// VoucherQuantity 是发行券的数量。		限制：	券总预算<= 999999		特别说明：	该字段的含义是追加到的数量。	不可减少，只能增加。		code_mode=	MERCHANT_UPLOAD 的情况下，voucher_quantity 需要与导入码的总数量保持一致（即历史导入数量 + 追加导入数量）。
	VoucherQuantity *int64 `json:"voucher_quantity,omitempty"`
}

// MarketingActivityOrderVoucherCodeDeposit 是支付宝官方 V3 OpenAPI 模型。
type MarketingActivityOrderVoucherCodeDeposit struct {
	// MerchantAccessMode 是商户接入模式
	MerchantAccessMode *string `json:"merchant_access_mode,omitempty"`
	// OutBizNo 是外部业务单号，用作幂等控制。		幂等作用：	参数不变的情况下，再次请求返回与上一次相同的结果。		外部接入方需保证业务单号唯一。
	OutBizNo *string `json:"out_biz_no,omitempty"`
	// VoucherCodes 是券码的数量列表。接口参数为列表类型。		限制：	目前最大上传 1000 个。		单个code最长64位。		商户上传的券code列表，code允许包含的字符有0-9、a-z、A-Z、-、_、+、=、|。
	VoucherCodes *[]string `json:"voucher_codes,omitempty"`
}

// MarketingActivityOrderVoucherModify 是支付宝官方 V3 OpenAPI 模型。
type MarketingActivityOrderVoucherModify struct {
	ActivityBaseInfo *ActivityBaseInfo `json:"activity_base_info,omitempty"`
	// ActivityName 是活动名称。
	ActivityName  *string              `json:"activity_name,omitempty"`
	CustomerGuide *CustomerGuideModify `json:"customer_guide,omitempty"`
	// MerchantAccessMode 是商户接入模式
	MerchantAccessMode *string `json:"merchant_access_mode,omitempty"`
	// OutBizNo 是外部业务单号，用作幂等控制。	幂等作用：	参数不变的情况下，再次请求返回与上一次相同的结果。
	OutBizNo *string `json:"out_biz_no,omitempty"`
	// PublishEndTime 是券发放结束时间。该值为空表示不修改	格式为：yyyy-MM-dd HH:mm:ss
	PublishEndTime *string `json:"publish_end_time,omitempty"`
	// PublishStartTime 是券发放开始时间。 格式为：yyyy-MM-dd HH:mm:ss
	PublishStartTime          *string                    `json:"publish_start_time,omitempty"`
	VoucherAvailableScopeInfo *VoucherAvailableScopeInfo `json:"voucher_available_scope_info,omitempty"`
	// VoucherAvailableScopeModifyType 是可用范围修改类型。
	VoucherAvailableScopeModifyType *string                      `json:"voucher_available_scope_modify_type,omitempty"`
	VoucherCustomerGuideInfo        *VoucherCustomerGuideInfo    `json:"voucher_customer_guide_info,omitempty"`
	VoucherDisplayInfo              *VoucherDisplayInfoModify    `json:"voucher_display_info,omitempty"`
	VoucherDisplayPatternInfo       *VoucherDisplayPatternInfo   `json:"voucher_display_pattern_info,omitempty"`
	VoucherSendModeInfo             *VoucherSendModeInfo         `json:"voucher_send_mode_info,omitempty"`
	VoucherSendRule                 *VoucherSendRuleDetailModify `json:"voucher_send_rule,omitempty"`
	VoucherUseRule                  *VoucherUseRuleModify        `json:"voucher_use_rule,omitempty"`
	VoucherUseRuleInfo              *VoucherUseRuleInfo          `json:"voucher_use_rule_info,omitempty"`
}

// MarketingActivityOrderVoucherRefund 是支付宝官方 V3 OpenAPI 模型。
type MarketingActivityOrderVoucherRefund struct {
	// BizDt 是当前订单的退券时间。				格式为：yyyy-MM-dd HH:mm:ss
	BizDt *string `json:"biz_dt,omitempty"`
	// MerchantAccessMode 是商户接入模式
	MerchantAccessMode *string `json:"merchant_access_mode,omitempty"`
	// OutBizNo 是外部业务单号，用作幂等控制。				幂等作用：		参数不变的情况下，再次请求返回与上一次相同的结果。				外部接入方需保证业务单号唯一
	OutBizNo *string `json:"out_biz_no,omitempty"`
	// TotalFee 是当前订单的退款总金额(包含优惠券金额)，单位为元。
	TotalFee                *string                  `json:"total_fee,omitempty"`
	VoucherRefundDetailInfo *VoucherRefundDetailInfo `json:"voucher_refund_detail_info,omitempty"`
}

// MarketingActivityOrderVoucherStop 是支付宝官方 V3 OpenAPI 模型。
type MarketingActivityOrderVoucherStop struct {
	// MerchantAccessMode 是商户接入模式
	MerchantAccessMode *string `json:"merchant_access_mode,omitempty"`
	// OutBizNo 是外部业务单号，用作幂等控制。		幂等作用：	参数不变的情况下，再次请求返回与上一次相同的结果。		外部接入方需保证业务单号唯一。
	OutBizNo *string `json:"out_biz_no,omitempty"`
}

// MarketingActivityOrderVoucherUse 是支付宝官方 V3 OpenAPI 模型。
type MarketingActivityOrderVoucherUse struct {
	// BizDt 是外部优惠券的核销时间，用于流水记录。 格式为：yyyy-MM-dd HH:mm:ss
	BizDt *string `json:"biz_dt,omitempty"`
	// GoodsDetail 是商品明细信息
	GoodsDetail *[]GoodsDetailDTO `json:"goods_detail,omitempty"`
	// MerchantAccessMode 是商户接入模式
	MerchantAccessMode *string `json:"merchant_access_mode,omitempty"`
	// OutBizNo 是外部业务单号，用作幂等控制。				幂等作用：		参数不变的情况下，再次请求返回与上一次相同的结果。				外部接入方需保证业务单号唯一。
	OutBizNo *string `json:"out_biz_no,omitempty"`
	// RealShopId 是用户核销优惠券的代运营商业关系门店id，调用接口alipay.business.relation.shop.create创建门店返回的real_shop_id。	限制:	传入的门店id必须是创建商家券voucher_use_rule. voucher_available_scope.  order_voucher_available_shop对象中real_shop_ids中的门店id
	RealShopId *string `json:"real_shop_id,omitempty"`
	// StoreId 是用户核销优惠券的支付门店id,调用接口ant.merchant.expand.shop.create创建门店返回的门店id。	限制:	传入的门店id必须是创建商家券voucher_use_rule. voucher_available_scope.  order_voucher_available_shop对象中shop_ids中的门店id
	StoreId *string `json:"store_id,omitempty"`
	// TotalFee 是订单总金额
	TotalFee *string `json:"total_fee,omitempty"`
	// TradeChannel 是交易渠道。
	TradeChannel *string `json:"trade_channel,omitempty"`
	// TradeNo 是该交易在支付宝系统中的交易流水号。	注意：若用户使用支付宝为交易的渠道，本参数必填。
	TradeNo              *string               `json:"trade_no,omitempty"`
	VoucherUseDetailInfo *VoucherUseDetailInfo `json:"voucher_use_detail_info,omitempty"`
}

// AuditResultInfo 是支付宝官方 V3 OpenAPI 模型。
type AuditResultInfo struct {
	// AuditReason 是审核驳回的原因
	AuditReason *string `json:"audit_reason,omitempty"`
	// AuditResult 是审核结果	SUBMIT：审核中	APPROVE：审核通过	REJECT：审核驳回
	AuditResult *string `json:"audit_result,omitempty"`
	// AuditType 是审核类型	UNIFY_ACCESS_CHECK：统一接入平台审核
	AuditType *string `json:"audit_type,omitempty"`
}

// BelongMerchantInfo 是支付宝官方 V3 OpenAPI 模型。
type BelongMerchantInfo struct {
	// BusinessType 是合作业务类型,默认为ISV_FOR_MERCHANT
	BusinessType *string `json:"business_type,omitempty"`
	// MerchantId 是商户PID,默认为当前接口调用商户。
	MerchantId *string `json:"merchant_id,omitempty"`
	// MerchantIdType 是商户id类型。特别说明：如果merchant_id_type选择SMID。则表明当前商家券活动的归属者是该SMID所绑定的支付宝PID。因此要求该SMID必须绑定相应的支付宝PID。
	MerchantIdType *string `json:"merchant_id_type,omitempty"`
}

// BizfundSettleInfo 是支付宝官方 V3 OpenAPI 模型。
type BizfundSettleInfo struct {
	// SettleMode 是请款|账期 模式
	SettleMode *string `json:"settle_mode,omitempty"`
	// SettleTargetType 是结算的目标对象类型
	SettleTargetType *string `json:"settle_target_type,omitempty"`
}

// CustomerDefineDeductRule 是支付宝官方 V3 OpenAPI 模型。
type CustomerDefineDeductRule struct {
	// CustomerDefineDeductRuleDesc 是自定义优惠规则描述
	CustomerDefineDeductRuleDesc *string `json:"customer_define_deduct_rule_desc,omitempty"`
}

// CustomerDefineVoucherInfo 是支付宝官方 V3 OpenAPI 模型。
type CustomerDefineVoucherInfo struct {
	DeductInfo *DeductInfo `json:"deduct_info,omitempty"`
}

// CustomerGuide 是支付宝官方 V3 OpenAPI 模型。
type CustomerGuide struct {
	// MiniAppId 是券可用的小程序appId，卡包详情页可跳转到该appId
	MiniAppId *string `json:"mini_app_id,omitempty"`
	// MiniAppPath 是指定跳转到mini_app_id时的具体页面路径。 	限制： 	1、只有mini_app_id有值时该值传入才会有效	2、该小程序路径是相对路径。详情参见 小程序scheme链接介绍
	MiniAppPath *string `json:"mini_app_path,omitempty"`
	// RealShopIds 是代运营商业关系门店列表，列表中的门店id是调用接口alipay.business.relation.shop.create创建门店返回的real_shop_id。接口参数是列表类型。
	RealShopIds *[]string `json:"real_shop_ids,omitempty"`
	// ServiceCodes 是小程序服务编码，通过 alipay.open.app.appcontent.function.create(小程序服务创建)接口创建服务后获取。
	ServiceCodes *[]string `json:"service_codes,omitempty"`
	// ShopIds 是券可使用的门店列表。列表中的门店id是通过调用接口ant.merchant.expand.shop.create创建门店返回的支付宝门店id  接口参数是列表类型。
	ShopIds *[]string `json:"shop_ids,omitempty"`
	// StoreIds 是该字段后续废弃。券可使用的门店列表。列表中的门店id是通过调用接口ant.merchant.expand.shop.create创建门店返回的支付宝门店id。接口参数是列表类型。
	StoreIds         *[]string         `json:"store_ids,omitempty"`
	VoucherSendGuide *VoucherSendGuide `json:"voucher_send_guide,omitempty"`
	VoucherUseGuide  *VoucherUseGuide  `json:"voucher_use_guide,omitempty"`
}

// CustomerGuideModify 是支付宝官方 V3 OpenAPI 模型。
type CustomerGuideModify struct {
	// MiniAppId 是券可用的小程序 appId，卡包详情页可跳转到该 appId。
	MiniAppId *string `json:"mini_app_id,omitempty"`
	// MiniAppPath 是指定跳转到 mini_app_id 时的具体页面路径。
	MiniAppPath *string `json:"mini_app_path,omitempty"`
	// RealShopIds 是券可使用的商业关系门店列表,列表中的门店id是调用接口alipay.business.relation.shop.create创建门店返回的real_shop_id。
	RealShopIds *[]string `json:"real_shop_ids,omitempty"`
	// ServiceCodes 是小程序服务编码，通过 alipay.open.app.appcontent.function.create(小程序服务创建)接口创建服务后获取。
	ServiceCodes *[]string `json:"service_codes,omitempty"`
	// ShopIds 是券可使用的门店列表。列表中的门店id是通过调用接口ant.merchant.expand.shop.create创建门店返回的支付宝门店 id 接口参数是列表类型。 接口参数是列表类型。
	ShopIds *[]string `json:"shop_ids,omitempty"`
	// StoreIds 是后续废弃。券可使用的门店列表。列表中的门店id是通过调用接口ant.merchant.expand.shop.create创建门店返回的支付宝门店  id 接口参数是列表类型。	接口参数是列表类型。
	StoreIds         *[]string               `json:"store_ids,omitempty"`
	VoucherSendGuide *VoucherSendGuideModify `json:"voucher_send_guide,omitempty"`
	VoucherUseGuide  *VoucherUseGuideModify  `json:"voucher_use_guide,omitempty"`
}

// DateRangeInfo 是支付宝官方 V3 OpenAPI 模型。
type DateRangeInfo struct {
	// BeginDate 是开始日期  格式：yyyy-MM-dd
	BeginDate *string `json:"begin_date,omitempty"`
	// EndDate 是结束日期  格式：yyyy-MM-dd
	EndDate *string `json:"end_date,omitempty"`
}

// DateRuleInfo 是支付宝官方 V3 OpenAPI 模型。
type DateRuleInfo struct {
	DateRangeInfo *DateRangeInfo `json:"date_range_info,omitempty"`
	TimeRangeInfo *TimeRangeInfo `json:"time_range_info,omitempty"`
}

// DeductInfo 是支付宝官方 V3 OpenAPI 模型。
type DeductInfo struct {
	CustomerDefineDeductRule *CustomerDefineDeductRule `json:"customer_define_deduct_rule,omitempty"`
}

// DiscountVoucher 是支付宝官方 V3 OpenAPI 模型。
type DiscountVoucher struct {
	// CeilingAmount 是封顶金额。
	CeilingAmount *string `json:"ceiling_amount,omitempty"`
	// Discount 是折扣率。
	Discount *string `json:"discount,omitempty"`
	// FloorAmount 是门槛金额。说明：该字段可不填，认为无门槛;
	FloorAmount *string `json:"floor_amount,omitempty"`
	// GoodsName 是商品名称。
	GoodsName *string `json:"goods_name,omitempty"`
	// OriginAmount 是原价。说明：该字段可不填，填入商品名称goods_name则必填;
	OriginAmount *string `json:"origin_amount,omitempty"`
}

// DiscountVoucherInfo 是支付宝官方 V3 OpenAPI 模型。
type DiscountVoucherInfo struct {
	// CeilingAmount 是封顶金额。
	CeilingAmount *string `json:"ceiling_amount,omitempty"`
	// Discount 是折扣率。
	Discount *string `json:"discount,omitempty"`
	// FloorAmount 是门槛金额。说明：该字段可不填，认为无门槛;
	FloorAmount                *string                     `json:"floor_amount,omitempty"`
	VoucherDeductThresholdInfo *VoucherDeductThresholdInfo `json:"voucher_deduct_threshold_info,omitempty"`
}

// EndTimeInfo 是支付宝官方 V3 OpenAPI 模型。
type EndTimeInfo struct {
	// EndTime 是结束时间  格式：HH:mm:ss
	EndTime *string `json:"end_time,omitempty"`
	// EndTimeType 是结束时间类型
	EndTimeType *string `json:"end_time_type,omitempty"`
}

// ExchangeVoucher 是支付宝官方 V3 OpenAPI 模型。
type ExchangeVoucher struct {
	// Amount 是券的价值
	Amount *string `json:"amount,omitempty"`
	// BizType 是兑换券业务类型。		枚举值	团购券：GROUP_BUY_EXCHANGE_VOUCHER	代金券：FIX_EXCHANGE_VOUCHER	注意：兑换券通过大促活动权益报名能力推广至支付宝会场时，本参数必填。
	BizType *string `json:"biz_type,omitempty"`
	// CustomerServiceMobile 是客服电话
	CustomerServiceMobile *string `json:"customer_service_mobile,omitempty"`
	// CustomerServiceUrl 是客服链接
	CustomerServiceUrl *string `json:"customer_service_url,omitempty"`
	// FloorAmount 是门槛金额。说明：该字段可不填，认为无门槛;
	FloorAmount *string `json:"floor_amount,omitempty"`
	// OverdueRefundable 是是否支持优惠券过期后，自动退款给用户。	不填默认否，枚举值：	true：是	false：否
	OverdueRefundable *bool `json:"overdue_refundable,omitempty"`
	// PayeePid 是收款账号。目前的结算规则是，每核销一笔优惠券，支付宝会打款到该收款账户。
	PayeePid *string `json:"payee_pid,omitempty"`
	// Refundable 是购买的优惠券是否允许退款。不填默认否，枚举值：	true：是	false：否
	Refundable *bool `json:"refundable,omitempty"`
	// SaleAmount 是用户购买优惠券需要支付的金额
	SaleAmount *string `json:"sale_amount,omitempty"`
	// SettleType 是解决类型。
	SettleType *string `json:"settle_type,omitempty"`
	// VoucherDetailUrl 是领(购)券详情页链接，从支付宝公域跳转到服务商(商户)自定义领(购)券详情页。
	VoucherDetailUrl *string `json:"voucher_detail_url,omitempty"`
	// VoucherName 是对消费者展示的券(商品)名称。
	VoucherName *string `json:"voucher_name,omitempty"`
}

// ExchangeVoucherInfo 是支付宝官方 V3 OpenAPI 模型。
type ExchangeVoucherInfo struct {
	// Amount 是券的价值
	Amount *string `json:"amount,omitempty"`
	// BizType 是兑换券业务类型。 注意：兑换券通过大促活动权益报名能力推广至支付宝会场时，本参数必填。
	BizType           *string                   `json:"biz_type,omitempty"`
	ExchangeGoodsInfo *VoucherExchangeGoodsInfo `json:"exchange_goods_info,omitempty"`
	// FloorAmount 是优惠门槛金额，表示只有当订单金额大于等于门槛金额时券才能使用。该字段为空时表示无门槛。  门槛金额的校验由服务商(商户)核销时自行校验，支付宝侧只做展示使用。
	FloorAmount                *string                     `json:"floor_amount,omitempty"`
	VoucherDeductThresholdInfo *VoucherDeductThresholdInfo `json:"voucher_deduct_threshold_info,omitempty"`
}

// ExchangeVoucherModify 是支付宝官方 V3 OpenAPI 模型。
type ExchangeVoucherModify struct {
	// VoucherDetailUrl 是领(购)券详情页链接，从支付宝公域跳转到服务商(商户)自定义领(购)券详情页。
	VoucherDetailUrl *string `json:"voucher_detail_url,omitempty"`
	// VoucherName 是对消费者展示的券(商品)名称。
	VoucherName *string `json:"voucher_name,omitempty"`
}

// FailVoucherCodeDetail 是支付宝官方 V3 OpenAPI 模型。
type FailVoucherCodeDetail struct {
	// ErrorCode 是券码导入失败错误码
	ErrorCode *string `json:"error_code,omitempty"`
	// ErrorMsg 是券码导入失败错误原因描述
	ErrorMsg *string `json:"error_msg,omitempty"`
	// VoucherCode 是导入失败的券码
	VoucherCode *string `json:"voucher_code,omitempty"`
}

// FixVoucher 是支付宝官方 V3 OpenAPI 模型。
type FixVoucher struct {
	// Amount 是面额，每张代金券可以抵扣的金额。
	Amount *string `json:"amount,omitempty"`
	// FloorAmount 是门槛金额。该字段不填写，认为无门槛。
	FloorAmount *string `json:"floor_amount,omitempty"`
	// GoodsName 是商品名称。
	GoodsName *string `json:"goods_name,omitempty"`
}

// FixVoucherInfo 是支付宝官方 V3 OpenAPI 模型。
type FixVoucherInfo struct {
	// Amount 是面额，每张代金券可以抵扣的金额。
	Amount *string `json:"amount,omitempty"`
	// FloorAmount 是门槛金额。该字段不填写，认为无门槛。
	FloorAmount                *string                     `json:"floor_amount,omitempty"`
	VoucherDeductThresholdInfo *VoucherDeductThresholdInfo `json:"voucher_deduct_threshold_info,omitempty"`
}

// GoodsDetailDTO 是支付宝官方 V3 OpenAPI 模型。
type GoodsDetailDTO struct {
	// GoodsId 是商品编号
	GoodsId *string `json:"goods_id,omitempty"`
	// GoodsName 是商品名称
	GoodsName *string `json:"goods_name,omitempty"`
	// Price 是商品单价。单位 元		精确到小数点后2位。
	Price *string `json:"price,omitempty"`
	// Quantity 是商品数量
	Quantity *int64 `json:"quantity,omitempty"`
}

// HolidayRuleInfo 是支付宝官方 V3 OpenAPI 模型。
type HolidayRuleInfo struct {
	TimeRangeInfo *TimeRangeInfo `json:"time_range_info,omitempty"`
}

// OrderVoucherAvailableCityCode 是支付宝官方 V3 OpenAPI 模型。
type OrderVoucherAvailableCityCode struct {
	// AllCity 是是否全国。选择全国后，无须填写city_codes字段。系统默认填充全国全部城市信息。
	AllCity *bool `json:"all_city,omitempty"`
	// CityCodes 是城市编码。请按照表格中内容填写。 （参考资料）
	CityCodes *[]string `json:"city_codes,omitempty"`
}

// OrderVoucherAvailableCityCodeModify 是支付宝官方 V3 OpenAPI 模型。
type OrderVoucherAvailableCityCodeModify struct {
	// CityCodes 是城市编码。请按照https://gw.alipayobjects.com/os/basement_prod/253c4dcb-b8a4-4a1e-8be2-79e191a9b6db.xlsx 表格中内容填写。 （参考资料： http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/）
	CityCodes *[]string `json:"city_codes,omitempty"`
}

// OrderVoucherAvailableScopeResult 是支付宝官方 V3 OpenAPI 模型。
type OrderVoucherAvailableScopeResult struct {
	OrderVoucherAvailableShopResult *OrderVoucherAvailableShopResult `json:"order_voucher_available_shop_result,omitempty"`
	// VoucherAvailableType 是可用范围类型。
	VoucherAvailableType *string `json:"voucher_available_type,omitempty"`
}

// OrderVoucherAvailableShop 是支付宝官方 V3 OpenAPI 模型。
type OrderVoucherAvailableShop struct {
	OrderVoucherMerchantAllShop *OrderVoucherMerchantAllShop `json:"order_voucher_merchant_all_shop,omitempty"`
	// RealShopIds 是代运营商业关系门店列表，列表中的门店id是调用接口alipay.business.relation.shop.create创建门店返回的real_shop_id  接口参数是列表类型。
	RealShopIds *[]string `json:"real_shop_ids,omitempty"`
	// ShopIds 是券可使用的门店列表。列表中的门店id是通过调用接口ant.merchant.expand.shop.create创建门店返回的支付宝门店id  接口参数是列表类型。
	ShopIds *[]string `json:"shop_ids,omitempty"`
}

// OrderVoucherAvailableShopModify 是支付宝官方 V3 OpenAPI 模型。
type OrderVoucherAvailableShopModify struct {
	OrderVoucherMerchantAllShop *OrderVoucherMerchantAllShopModify `json:"order_voucher_merchant_all_shop,omitempty"`
	// RealShopIds 是券可使用的商业关系门店列表,列表中的门店id是调用接口alipay.business.relation.shop.create创建门店返回的real_shop_id。
	RealShopIds *[]string `json:"real_shop_ids,omitempty"`
	// ShopIds 是券可使用的门店列表。列表中的门店id是通过调用接口ant.merchant.expand.shop.create创建门店返回的支付宝门店id 接口参数是列表类型。接口参数是列表类型。
	ShopIds *[]string `json:"shop_ids,omitempty"`
}

// OrderVoucherAvailableShopResult 是支付宝官方 V3 OpenAPI 模型。
type OrderVoucherAvailableShopResult struct {
	OrderVoucherMerchantAllShopResult *OrderVoucherMerchantAllShopResult `json:"order_voucher_merchant_all_shop_result,omitempty"`
	// RealShopFailInfos 是请求失败的代运营商业关系门店详情
	RealShopFailInfos *[]OrderVoucherRealShopFailInfo `json:"real_shop_fail_infos,omitempty"`
	// ShopFailInfos 是请求失败的支付宝门店详情。
	ShopFailInfos *[]OrderVoucherShopFailInfo `json:"shop_fail_infos,omitempty"`
	// SuccessRealShopIds 是请求成功的代运营商业关系门店
	SuccessRealShopIds *[]string `json:"success_real_shop_ids,omitempty"`
	// SuccessShopIds 是请求成功的支付宝门店。
	SuccessShopIds *[]string `json:"success_shop_ids,omitempty"`
}

// OrderVoucherMerchantAllShop 是支付宝官方 V3 OpenAPI 模型。
type OrderVoucherMerchantAllShop struct {
	// ExcludeShopIds 是券不可使用的门店列表。指定商户全部门店可用时可通过该字段排除部分不可用门店。列表中的门店id是通过调用接口ant.merchant.expand.shop.create创建门店返回的支付宝门店id  接口参数是列表类型。
	ExcludeShopIds *[]string `json:"exclude_shop_ids,omitempty"`
}

// OrderVoucherMerchantAllShopModify 是支付宝官方 V3 OpenAPI 模型。
type OrderVoucherMerchantAllShopModify struct {
	// ExcludeShopIds 是券不可使用的门店列表。指定商户全部门店可用时可通过该字段排除部分不可用门店。列表中的门店id是通过调用接口ant.merchant.expand.shop.create创建门店返回的支付宝门店id  接口参数是列表类型。
	ExcludeShopIds *[]string `json:"exclude_shop_ids,omitempty"`
}

// OrderVoucherMerchantAllShopResult 是支付宝官方 V3 OpenAPI 模型。
type OrderVoucherMerchantAllShopResult struct {
	// ExcludeShopFailInfos 是请求失败的不可用门店详情。
	ExcludeShopFailInfos *[]OrderVoucherShopFailInfo `json:"exclude_shop_fail_infos,omitempty"`
	// SuccessExcludeShopIds 是请求成功的不可用门店列表。
	SuccessExcludeShopIds *[]string `json:"success_exclude_shop_ids,omitempty"`
}

// OrderVoucherRealShopFailInfo 是支付宝官方 V3 OpenAPI 模型。
type OrderVoucherRealShopFailInfo struct {
	// FailMessage 是请求失败提示信息。
	FailMessage *string `json:"fail_message,omitempty"`
	// FailReasons 是请求失败原因。
	FailReasons *[]string `json:"fail_reasons,omitempty"`
	// RealShopId 是请求失败的代运营商业关系门店 ID。
	RealShopId *string `json:"real_shop_id,omitempty"`
}

// OrderVoucherShopFailInfo 是支付宝官方 V3 OpenAPI 模型。
type OrderVoucherShopFailInfo struct {
	// FailMessage 是请求失败提示信息。
	FailMessage *string `json:"fail_message,omitempty"`
	// FailReasons 是请求失败的原因。
	FailReasons *[]string `json:"fail_reasons,omitempty"`
	// ShopId 是支付宝侧蚂蚁店铺 id。
	ShopId *string `json:"shop_id,omitempty"`
}

// OrderVoucherUseRuleResult 是支付宝官方 V3 OpenAPI 模型。
type OrderVoucherUseRuleResult struct {
	VoucherAvailableScopeResult *OrderVoucherAvailableScopeResult `json:"voucher_available_scope_result,omitempty"`
}

// SpecialVoucher 是支付宝官方 V3 OpenAPI 模型。
type SpecialVoucher struct {
	// FloorAmount 是门槛金额。说明：该字段可不填，认为无门槛;
	FloorAmount *string `json:"floor_amount,omitempty"`
	// GoodsName 是商品名称。
	GoodsName *string `json:"goods_name,omitempty"`
	// OriginAmount 是原价。	限制：	币种为人民币，单位为元。小数点以后最多保留两位;	取值范围:0.1<=x<=9999。
	OriginAmount *string `json:"origin_amount,omitempty"`
	// SpecialAmount 是特价，即：原价-特价=优惠金额。
	SpecialAmount *string `json:"special_amount,omitempty"`
}

// SpecialVoucherInfo 是支付宝官方 V3 OpenAPI 模型。
type SpecialVoucherInfo struct {
	// FloorAmount 是门槛金额。说明：该字段可不填，认为无门槛;
	FloorAmount *string `json:"floor_amount,omitempty"`
	// SpecialAmount 是特价，即：原价-特价=优惠金额。
	SpecialAmount              *string                     `json:"special_amount,omitempty"`
	VoucherDeductThresholdInfo *VoucherDeductThresholdInfo `json:"voucher_deduct_threshold_info,omitempty"`
}
