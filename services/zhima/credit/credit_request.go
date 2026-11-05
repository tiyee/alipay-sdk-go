// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package credit

// PayAfterUseCreditAgreementTransferRequest is the official V3 ZhimaCreditPayafteruseCreditagreementTransferModel schema.
type PayAfterUseCreditAgreementTransferRequest struct {
	// CategoryId represents category_id. 芝麻外部类目
	CategoryId *string `json:"category_id,omitzero"`
	// CreditAgreementId represents credit_agreement_id. 芝麻开通/授权协议号，在用户开通芝麻先享服务后，由芝麻生成并返回给商户；只有当opertaion_type传入AGREEMENT_TRANSFER值是，该字段必传；其他场景不需要传入
	CreditAgreementId *string `json:"credit_agreement_id,omitzero"`
	// DeductAgreementNo represents deduct_agreement_no. 用户的代扣协议号
	DeductAgreementNo *string `json:"deduct_agreement_no,omitzero"`
	// ExtendParams represents extend_params. 业务扩展参数，Json格式
	ExtendParams *string `json:"extend_params,omitzero"`
	// OperationType represents operation_type. AGREEMENT_TRANSFER代表是存量先享协议的主体切换功能；其他场景不需要传入；默认为空；
	OperationType *string `json:"operation_type,omitzero"`
	// OutAgreementNo represents out_agreement_no. 商户外部协议号，不同的支付宝用户需要传递不同的外部单号
	OutAgreementNo *string `json:"out_agreement_no,omitzero"`
	// ProductCode represents product_code. 商户签约的芝麻产品的产品码
	ProductCode *string `json:"product_code,omitzero"`
	// ZmServiceId represents zm_service_id. 芝麻信用服务ID
	ZmServiceId *string `json:"zm_service_id,omitzero"`
}

// PayAfterUseCreditBizOrderFinishRequest is the official V3 ZhimaCreditPayafteruseCreditbizorderFinishModel schema.
type PayAfterUseCreditBizOrderFinishRequest struct {
	// CreditBizOrderId represents credit_biz_order_id. 信用服务订单号
	CreditBizOrderId *string `json:"credit_biz_order_id,omitzero"`
	// IsFulfilled represents is_fulfilled. 字符串类型，用户此订单是否守约。 传"true"时，用户在芝麻信用-守约记录中，该笔订单是已守约状态； 传"false"时，用户在芝麻信用-守约记录中，该笔订单是已取消状态。
	IsFulfilled *string `json:"is_fulfilled,omitzero"`
	// OutRequestNo represents out_request_no. 商户外部请求号
	OutRequestNo *string `json:"out_request_no,omitzero"`
	// Remark represents remark. 商户对本次操作的附言描述
	Remark *string `json:"remark,omitzero"`
}

// PayAfterUseCreditBizOrderOrderRequest is the official V3 ZhimaCreditPayafteruseCreditbizorderOrderModel schema.
type PayAfterUseCreditBizOrderOrderRequest struct {
	// AmountType represents amount_type. 只有当传递了order_amount时，该参数才有意义； 1）该参数不传时，默认为ORDER_AMOUNT。 2）传ORDER_AMOUNT时，表示order_amount传入的金额为后付金额，在发起扣款时，最大扣款支付金额为order_amount传入的值（取值单位为元）； 3）传RISK_AMOUNT时，表示order_amount传入的金额为风险预估金额，在发起扣款时，最大扣款支付金额为商户签约时约定的上限额度（取值单位为元）。
	AmountType *string `json:"amount_type,omitzero"`
	// Body represents body. 订单描述
	Body *string `json:"body,omitzero"`
	// CategoryId represents category_id. 芝麻外部类目，芝麻先享接入无差异化风控诉求可不传
	CategoryId *string `json:"category_id,omitzero"`
	// CommercialSubMode represents commercial_sub_mode. 业务子模式。默认的单次付模式无需传入，阶段付模式传入以区分是分次还是分期子模式。
	CommercialSubMode *string `json:"commercial_sub_mode,omitzero"`
	// CreditAgreementId represents credit_agreement_id. 芝麻开通协议号
	CreditAgreementId *string `json:"credit_agreement_id,omitzero"`
	// CreditCommercialMode represents credit_commercial_mode. 信用业务模式，不填默认为单次扣款模式。阶段付模式为STAGE_PAYMENT，其它模式请根据对应的技术支持文档传入
	CreditCommercialMode *string `json:"credit_commercial_mode,omitzero"`
	// ExtendParams represents extend_params. 业务扩展参数
	ExtendParams *string `json:"extend_params,omitzero"`
	// OrderAmount represents order_amount. 订单金额，该金额为当前订单扣款的累计最大额度。例如，下单时传递10.00，则扣款时最大支付金额为10元。若该参数不传，则默认使用商户签约时约定的上限额度。芝麻速办业务场景（极速回收、极速返利、极速退款等）商户请求时，order_amount必传，且amount_type类型需传递ORDER_AMOUNT。
	OrderAmount *string `json:"order_amount,omitzero"`
	// OutOrderNo represents out_order_no. 商户外部订单号，保证不重复
	OutOrderNo *string `json:"out_order_no,omitzero"`
	// PaymentTotalTimes represents payment_total_times. 多阶段订单次数，业务模式为阶段付模式下时需传入
	PaymentTotalTimes *string `json:"payment_total_times,omitzero"`
	// ProductCode represents product_code. 产品码，不传默认为CREDIT_PAY_AFTER_USE
	ProductCode *string `json:"product_code,omitzero"`
	// StagePeriodType represents stage_period_type. 阶段付分期类型。阶段付模式，且子业务模式为分期模式下需要传入，分次不需要
	StagePeriodType *string `json:"stage_period_type,omitzero"`
	// Subject represents subject. 订单标题。 注意：不可使用特殊字符，如 /，=，& 等。
	Subject *string `json:"subject,omitzero"`
}

// PeZmGoAgreementUnsignRequest is the official V3 ZhimaCreditPeZmgoAgreementUnsignModel schema.
type PeZmGoAgreementUnsignRequest struct {
	// AgreementId represents agreement_id. 支付宝系统中用以唯一标识用户签约记录的编号。
	AgreementId *string `json:"agreement_id,omitzero"`
	// AlipayUserId represents alipay_user_id. 支付宝的用户id
	AlipayUserId *string `json:"alipay_user_id,omitzero"`
	// OpenId represents open_id. 支付宝的用户id
	OpenId *string `json:"open_id,omitzero"`
	// PartnerId represents partner_id. 商户ID
	PartnerId *string `json:"partner_id,omitzero"`
	// QuitType represents quit_type. quit_type为USER_CANCEL_QUIT或者SETTLE_APPLY_QUIT
	QuitType *string `json:"quit_type,omitzero"`
}

// PeZmGoBizOptCloseRequest is the official V3 ZhimaCreditPeZmgoBizoptCloseModel schema.
type PeZmGoBizOptCloseRequest struct {
	// AlipayUserId represents alipay_user_id. 支付宝用户Id
	AlipayUserId *string `json:"alipay_user_id,omitzero"`
	// OpenId represents open_id. 支付宝用户Id
	OpenId *string `json:"open_id,omitzero"`
	// OutRequestNo represents out_request_no. 商户本次操作的请求流水号，由商户生成，请确保唯一性
	OutRequestNo *string `json:"out_request_no,omitzero"`
	// PartnerId represents partner_id. 商户Id
	PartnerId *string `json:"partner_id,omitzero"`
	// TemplateId represents template_id. 芝麻GO配置模板Id
	TemplateId *string `json:"template_id,omitzero"`
}

// PeZmGoCumulationSyncRequest is the official V3 ZhimaCreditPeZmgoCumulationSyncModel schema.
type PeZmGoCumulationSyncRequest struct {
	// AgreementNo represents agreement_no. 芝麻go协议号，唯一标识一个芝麻go协议。为必传入参。
	AgreementNo *string `json:"agreement_no,omitzero"`
	// AmountTypeData represents amount_type_data.
	AmountTypeData *AmountTypeData `json:"amount_type_data,omitzero"`
	// BizAction represents biz_action. 数据回传的动作类型，为枚举值，为必传入参。 可取值： ("ORDER_PAID", "正向支付")，("ORDER_REFUND", "逆向退款")。 其中ORDER_PAID表示回传为正向的，即订单金额累加，优惠累加，任务次数累加；而ORDER_REFUND表示回传为逆向的，即订单金额退款，优惠退款，任务次数回退。
	BizAction *string `json:"biz_action,omitzero"`
	// BizTime represents biz_time. 交易发生时间，包括：订单交易时间，优惠核销时间，以及任务完成时间。为必传入参。 如：用户在【2019-03-08 00:00:00】核销了一个红包，但是【2019-03-10 00:00:00】才进行数据回传，此时biz_time取值红包核销的时间，也就是【2019-03-08 00:00:00】。
	BizTime *string `json:"biz_time,omitzero"`
	// CumulateDataType represents cumulate_data_type. 累计类型。为枚举值，为必传入参。 可取值：("AMOUNT", "金额类型")，("TASK", "任务类型")，("AMOUNT_TASK", "额度和任务类型")。 其中，交易解耦模式仅允许AMOUNT或TASK。而交易绑定模式下，仅允许AMOUNT_TASK。
	CumulateDataType *string `json:"cumulate_data_type,omitzero"`
	// ExtInfo represents ext_info. 扩展字段，格式为Json格式的字符串，非必传。
	ExtInfo *string `json:"ext_info,omitzero"`
	// HasAlipayTrade represents has_alipay_trade. 是否为支付宝交易。仅在累计模式为【交易绑定模式】下被消费。当芝麻Go模板配置的累计模式为交易绑定模式时，此时需传入扩展参数是否为支付宝交易。若是支付宝交易，传值true，否则为false。若不传此字段，默认为false。 若为true，则会对交易号trade_no进行校验，校验是否为支付宝交易，以及用户userId和商户partnerId是否一致，若校验不通过，累计失败。
	HasAlipayTrade *bool `json:"has_alipay_trade,omitzero"`
	// OpenId represents open_id. 用户userId，蚂蚁统一会员ID，为必传入参。
	OpenId *string `json:"open_id,omitzero"`
	// OutBizNo represents out_biz_no. 外部业务号。需保证唯一，为必传入参。 若接口调用失败，提示信息为更换幂等号，商户更换此字段重新发起调用。
	OutBizNo *string `json:"out_biz_no,omitzero"`
	// PartnerId represents partner_id. 商户ID，为必传入参。
	PartnerId *string `json:"partner_id,omitzero"`
	// PayOutBizNo represents pay_out_biz_no. 逆向对应的正向外部业务号。逆向场景为必传入参。也就是biz_action为ORDER_REFUND时为必传入参。 逆向时需通过此字段明确对应的正向的外部业务号。
	PayOutBizNo *string `json:"pay_out_biz_no,omitzero"`
	// RequestFrom represents request_from. 取固定值ExternalMerchantSource，为必传入参。表示数据回传来源为外部商户，数据回传内部逻辑需感知。
	RequestFrom *string `json:"request_from,omitzero"`
	// TaskTypeData represents task_type_data.
	TaskTypeData *TaskTypeData `json:"task_type_data,omitzero"`
	// UserId represents user_id. 用户userId，蚂蚁统一会员ID，为必传入参。
	UserId *string `json:"user_id,omitzero"`
}

// PeZmGoPreOrderCreateRequest is the official V3 ZhimaCreditPeZmgoPreorderCreateModel schema.
type PeZmGoPreOrderCreateRequest struct {
	// AlipayOpenId represents alipay_open_id. 支付宝用户ID。同一个外部请求号，调用方要确保支付宝userId取值不变
	AlipayOpenId *string `json:"alipay_open_id,omitzero"`
	// AlipayUserId represents alipay_user_id. 支付宝用户ID。同一个外部请求号，调用方要确保支付宝userId取值不变
	AlipayUserId *string `json:"alipay_user_id,omitzero"`
	// BizTime represents biz_time. 业务发生时间，如2016-07-06 00:00:02 签约超时关单时间依赖这个时间 这个时间会做幂等校验，如果创单幂等，但时间不同会失败
	BizTime *string `json:"biz_time,omitzero"`
	// CustomTempConf represents custom_temp_conf. 废弃字段，勿用
	CustomTempConf *string `json:"custom_temp_conf,omitzero"`
	// ExpireAisleData represents expire_aisle_data. 协议到期后通知
	ExpireAisleData *string `json:"expire_aisle_data,omitzero"`
	// ExtTemplateConf represents ext_template_conf.
	ExtTemplateConf *ExtTemplateConf `json:"ext_template_conf,omitzero"`
	// ExtendParams represents extend_params.
	ExtendParams *PreOrderExtInfo `json:"extend_params,omitzero"`
	// FreezeAmount represents freeze_amount. 冻结金额，单位：元
	FreezeAmount *string `json:"freeze_amount,omitzero"`
	// IsvPid represents isv_pid. ISV商户ID
	IsvPid *string `json:"isv_pid,omitzero"`
	// OutRequestNo represents out_request_no. 外部请求号，唯一
	OutRequestNo *string `json:"out_request_no,omitzero"`
	// PartnerId represents partner_id. 商户ID
	PartnerId *string `json:"partner_id,omitzero"`
	// PartnerUserIdentifier represents partner_user_identifier. 商户侧用户唯一标识。ex.商户侧uid
	PartnerUserIdentifier *string `json:"partner_user_identifier,omitzero"`
	// PayAisleData represents pay_aisle_data. 结算成功后通知
	PayAisleData *string `json:"pay_aisle_data,omitzero"`
	// SignAisleData represents sign_aisle_data. 签约成功后通知
	SignAisleData *string `json:"sign_aisle_data,omitzero"`
	// TemplateId represents template_id. 签约芝麻GO的模板ID
	TemplateId *string `json:"template_id,omitzero"`
	// TimeoutExpress represents timeout_express. 超时关单时间 默认1H。1m～15d。m表示分钟，h表示小时，d表示天。该参数数值不接受小数点， 如1.5h，可转换为90m 签约超时关单时间依赖这个时间
	TimeoutExpress *string `json:"timeout_express,omitzero"`
}

// PeZmGoSettleApplyRequest is the official V3 ZhimaCreditPeZmgoSettleApplyModel schema.
type PeZmGoSettleApplyRequest struct {
	// ActionType represents action_type. DEFAULT_SETTLE 默认结算模式; PAY_TO_ZERO 支付转0元；
	ActionType *string `json:"action_type,omitzero"`
	// AgreementId represents agreement_id. 支付宝系统中用以唯一标识用户签约记录的编号。
	AgreementId *string `json:"agreement_id,omitzero"`
	// AlipayOpenId represents alipay_open_id. 买家在支付宝的用户id
	AlipayOpenId *string `json:"alipay_open_id,omitzero"`
	// AlipayUserId represents alipay_user_id. 买家在支付宝的用户id
	AlipayUserId *string `json:"alipay_user_id,omitzero"`
	// OutRequestNo represents out_request_no. 商户本次操作的请求流水号，用于标示请求流水的唯一性，不能包含除中文、英文、数字以外的字符，需要保证在商户端不重复。
	OutRequestNo *string `json:"out_request_no,omitzero"`
	// PartnerId represents partner_id. 商户ID
	PartnerId *string `json:"partner_id,omitzero"`
	// PayAmount represents pay_amount. 需要支付的金额，单位为：元（人民币），精确到小数点后两位
	PayAmount *string `json:"pay_amount,omitzero"`
	// SettleExtendParams represents settle_extend_params.
	SettleExtendParams *SettleExtraParams `json:"settle_extend_params,omitzero"`
	// TotalDiscountAmount represents total_discount_amount. 芝麻GO任务开始到到期过程中和任务相关享受的总优惠金额，单位为：元（人民币），精确到小数点后两位。
	TotalDiscountAmount *string `json:"total_discount_amount,omitzero"`
	// TotalRealPayAmount represents total_real_pay_amount. 芝麻GO任务开始到到期过程中和任务相关支付的总金额，单位为：元（人民币），精确到小数点后两位。
	TotalRealPayAmount *string `json:"total_real_pay_amount,omitzero"`
	// TotalTaskCount represents total_task_count. 芝麻GO任务开始到到期过程中和任务相关完成的任务次数
	TotalTaskCount *string `json:"total_task_count,omitzero"`
	// WithholdPlanNo represents withhold_plan_no. 芝麻GO结算时，需要传入的扣款单据号。来源于协议到期通知消息，外部商户接入时，该参数必填
	WithholdPlanNo *string `json:"withhold_plan_no,omitzero"`
}

// PeZmGoSettleRefundRequest is the official V3 ZhimaCreditPeZmgoSettleRefundModel schema.
type PeZmGoSettleRefundRequest struct {
	// AgreementId represents agreement_id. 签约芝麻GO的业务协议号，在签约成功、结算成功等通知中回回传给商户。
	AgreementId *string `json:"agreement_id,omitzero"`
	// AlipayOpenId represents alipay_open_id. 支付宝用户ID
	AlipayOpenId *string `json:"alipay_open_id,omitzero"`
	// AlipayUserId represents alipay_user_id. 支付宝用户ID
	AlipayUserId *string `json:"alipay_user_id,omitzero"`
	// Memo represents memo. 本次退款说明
	Memo *string `json:"memo,omitzero"`
	// OutRequestNo represents out_request_no. 外部请求号
	OutRequestNo *string `json:"out_request_no,omitzero"`
	// PartnerId represents partner_id. 商户ID
	PartnerId *string `json:"partner_id,omitzero"`
	// RefundAmount represents refund_amount. 退款金额，单位为：元（人民币），精确到小数点后两位
	RefundAmount *string `json:"refund_amount,omitzero"`
	// RefundType represents refund_type. 退款类型： MEMBER_FEE_REFUND 会员费结算退款 ，WITHHOLD_FEE_REFUND 周期扣款金额退款。 不填写默认为 MEMBER_FEE_REFUND
	RefundType *string `json:"refund_type,omitzero"`
	// WithholdPlanNo represents withhold_plan_no. 芝麻侧扣款计划单据号，在结算扣款成功的通知消息中会回传。
	WithholdPlanNo *string `json:"withhold_plan_no,omitzero"`
}

// PeZmGoSettleUnfreezeRequest is the official V3 ZhimaCreditPeZmgoSettleUnfreezeModel schema.
type PeZmGoSettleUnfreezeRequest struct {
	// AgreementId represents agreement_id. 支付宝系统中用以唯一标识用户签约记录的编号，即花呗先享协议号
	AgreementId *string `json:"agreement_id,omitzero"`
	// AlipayOpenId represents alipay_open_id. 买家在支付宝的用户id
	AlipayOpenId *string `json:"alipay_open_id,omitzero"`
	// AlipayUserId represents alipay_user_id. 买家在支付宝的用户id
	AlipayUserId *string `json:"alipay_user_id,omitzero"`
	// BizTime represents biz_time. 解冻成功时间
	BizTime *string `json:"biz_time,omitzero"`
	// OrderTitle represents order_title. 解冻的描述
	OrderTitle *string `json:"order_title,omitzero"`
	// OutRequestNo represents out_request_no. 商户本次操作的请求流水号，用于标识请求流水的唯一性，不能包含除中文、英文、数字以外的字符，需要保证在商户端不重复。由商户传入，最终返回给商户。
	OutRequestNo *string `json:"out_request_no,omitzero"`
	// PartnerId represents partner_id. 商户ID
	PartnerId *string `json:"partner_id,omitzero"`
	// UnfreezeAmount represents unfreeze_amount. 解冻金额
	UnfreezeAmount *string `json:"unfreeze_amount,omitzero"`
	// UnfreezeExtendParams represents unfreeze_extend_params.
	UnfreezeExtendParams *UnfreezeExtendParams `json:"unfreeze_extend_params,omitzero"`
}

// PayAfterUseCreditAgreementQueryRequest contains parameters for zhima.credit.payafteruse.creditagreement.query.
type PayAfterUseCreditAgreementQueryRequest struct {
	// OutAgreementNo represents out_agreement_no. 商户外部协议号
	OutAgreementNo *string `query:"out_agreement_no"`
	// CreditAgreementId represents credit_agreement_id. 芝麻开通/授权协议号，out_agreement_no与credit_agreement_id必填一个，推荐使用credit_agreement_id。
	CreditAgreementId *string `query:"credit_agreement_id"`
	// ProductCode represents product_code. 当产品是芝麻守约时必填，需传入CREDIT_AGREEMENT
	ProductCode *string `query:"product_code"`
	// ServiceId represents service_id. 芝麻信用服务id，上线前提前分配给商户
	ServiceId *string `query:"service_id"`
	// ExtraParam represents extra_param. 扩展参数
	ExtraParam *string `query:"extra_param"`
}

// PayAfterUseCreditBizOrderQueryRequest contains parameters for zhima.credit.payafteruse.creditbizorder.query.
type PayAfterUseCreditBizOrderQueryRequest struct {
	// CreditBizOrderId represents credit_biz_order_id. 信用服务订单号，out_order_no与credit_biz_order_id至少传一个
	CreditBizOrderId *string `query:"credit_biz_order_id"`
	// OutOrderNo represents out_order_no. 商户外部单号，out_order_no与credit_biz_order_id至少传一个
	OutOrderNo *string `query:"out_order_no"`
}

// PeZmGoAgreementQueryRequest contains parameters for zhima.credit.pe.zmgo.agreement.query.
type PeZmGoAgreementQueryRequest struct {
	// AgreementId represents agreement_id. 支付宝系统中用以唯一标识用户签约记录的编号，即花芝轻会员。传入该参数，会忽略其它所有参数。
	AgreementId *string `query:"agreement_id"`
	// AlipayUserId represents alipay_user_id. 买家在支付宝的用户id
	AlipayUserId *string `query:"alipay_user_id"`
	// OpenId represents open_id. 买家在支付宝的用户id
	OpenId *string `query:"open_id"`
}
