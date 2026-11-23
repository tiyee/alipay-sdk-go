// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package merchant

// SubsidiariesApplyRequest is the official V3 ZhimaMerchantSubsidiariesApplyModel schema.
type SubsidiariesApplyRequest struct {
	// Pid represents pid. 商户pid
	Pid *string `json:"pid,omitzero"`
	// ServiceId represents service_id. 信用服务id
	ServiceId *string `json:"service_id,omitzero"`
	// Smid represents smid. 二级商户进件id
	Smid *string `json:"smid,omitzero"`
	// SubMerchantContactNumber represents sub_merchant_contact_number. 子商户联系电话
	SubMerchantContactNumber *string `json:"sub_merchant_contact_number,omitzero"`
	// SubMerchantJumpLink represents sub_merchant_jump_link. 子商户跳转链接
	SubMerchantJumpLink *string `json:"sub_merchant_jump_link,omitzero"`
	// SubMerchantLogoUrl represents sub_merchant_logo_url. 子商户logo地址
	SubMerchantLogoUrl *string `json:"sub_merchant_logo_url,omitzero"`
	// SubMerchantName represents sub_merchant_name. 子商户名
	SubMerchantName *string `json:"sub_merchant_name,omitzero"`
	// SubPid represents sub_pid. 子商户id
	SubPid *string `json:"sub_pid,omitzero"`
}

// SubsidiariesCloseRequest is the official V3 ZhimaMerchantSubsidiariesCloseModel schema.
type SubsidiariesCloseRequest struct {
	// OrderNo represents order_no. 工单标识
	OrderNo *string `json:"order_no,omitzero"`
}

// ZmGoCumulateSyncRequest is the official V3 ZhimaMerchantZmgoCumulateSyncModel schema.
type ZmGoCumulateSyncRequest struct {
	// AgreementId represents agreement_id. 芝麻go协议号，唯一标识一个芝麻go协议。
	AgreementId *string `json:"agreement_id,omitzero"`
	// AmountTypeSyncData represents amount_type_sync_data.
	AmountTypeSyncData *AmountTypeSyncData `json:"amount_type_sync_data,omitzero"`
	// BizAction represents biz_action. 数据回传的动作类型。枚举如下： * POSITIVE：正向回传，即订单金额累加，优惠累加，任务次数累加。 * REVERSE：逆向回传，即订单金额退款，优惠退款，任务次数回退。
	BizAction *string `json:"biz_action,omitzero"`
	// BizTime represents biz_time. 回传数据发生的实际时间，包括：订单交易时间，优惠核销时间，以及任务完成时间。为必传入参。 如：用户在【2019-03-08 00:00:00】核销了一个红包，但是【2019-03-10 00:00:00】才进行数据回传，此时biz_time取值红包核销的时间，也就是【2019-03-08 00:00:00】。
	BizTime *string `json:"biz_time,omitzero"`
	// DataType represents data_type. 回传数据类型，枚举： TASK - 进度型，影响进度 DISCOUNT - 优惠型
	DataType *string `json:"data_type,omitzero"`
	// DiscountTypeSyncData represents discount_type_sync_data.
	DiscountTypeSyncData *DiscountTypeSyncData `json:"discount_type_sync_data,omitzero"`
	// OpenId represents open_id. 用户 id，用户在支付宝的唯一标识，以 2088 开头的 16 位纯数字组成。
	OpenId *string `json:"open_id,omitzero"`
	// OutBizNo represents out_biz_no. 外部业务号，唯一标识一笔回传数据。
	OutBizNo *string `json:"out_biz_no,omitzero"`
	// ProviderPid represents provider_pid. 数据回传的商户 ID，即和用户发生业务来往的 PID 主体。商户账号在支付宝的唯一标识，以 2088 开头的 16 位纯数字组成。
	ProviderPid *string `json:"provider_pid,omitzero"`
	// ReferOutBizNo represents refer_out_biz_no. 逆向对应的正向外部业务号。逆向场景为必传入参。也就是biz_action为REVERSE时为必传入参。 逆向时需通过此字段明确对应的正向的外部业务号。
	ReferOutBizNo *string `json:"refer_out_biz_no,omitzero"`
	// SubBizAction represents sub_biz_action. 数据回传动作子类型，枚举值，为必传入参。 ADD - 新增数据 UPDATE - 修改数据 备注：UPDATE只能更新金额信息。
	SubBizAction *string `json:"sub_biz_action,omitzero"`
	// TimesTypeSyncData represents times_type_sync_data.
	TimesTypeSyncData *TimesTypeSyncData `json:"times_type_sync_data,omitzero"`
	// UserId represents user_id. 用户 id，用户在支付宝的唯一标识，以 2088 开头的 16 位纯数字组成。
	UserId *string `json:"user_id,omitzero"`
}

// ZmGoTemplateCreateRequest is the official V3 ZhimaMerchantZmgoTemplateCreateModel schema.
type ZmGoTemplateCreateRequest struct {
	// BasicConfig represents basic_config.
	BasicConfig *ZMGOBasicConfig `json:"basic_config,omitzero"`
	// ExtConfig represents ext_config.
	ExtConfig *ZMGOExtConfig `json:"ext_config,omitzero"`
	// ObligationConfig represents obligation_config.
	ObligationConfig *ZMGOObligationConfig `json:"obligation_config,omitzero"`
	// OpenConfig represents open_config.
	OpenConfig *ZMGOOpenConfig `json:"open_config,omitzero"`
	// QuitConfig represents quit_config.
	QuitConfig *ZMGOQuitConfig `json:"quit_config,omitzero"`
	// RightConfig represents right_config.
	RightConfig *ZMGORightConfig `json:"right_config,omitzero"`
	// SettlementConfig represents settlement_config.
	SettlementConfig *ZMGOSettlementConfig `json:"settlement_config,omitzero"`
}

// ZmGoCumulateQueryRequest contains parameters for zhima.merchant.zmgo.cumulate.query.
type ZmGoCumulateQueryRequest struct {
	// AgreementId represents agreement_id. 芝麻go协议号，唯一标识一个芝麻go协议。
	AgreementId *string `query:"agreement_id"`
	// UserId represents user_id. 用户 id，用户在支付宝的唯一标识，以 2088 开头的 16 位纯数字组成。
	UserId *string `query:"user_id"`
	// OpenId represents open_id. 用户 id，用户在支付宝的唯一标识，以 2088 开头的 16 位纯数字组成。
	OpenId *string `query:"open_id"`
	// ProviderPid represents provider_pid. 数据回传的商户 ID，即和用户发生业务来往的 PID 主体。商户账号在支付宝的唯一标识，以 2088 开头的 16 位纯数字组成。
	ProviderPid *string `query:"provider_pid"`
	// NeedDetail represents need_detail. 是否需要回传明细。枚举支持： * true：需要回传明细。 * false：不需要回传明细。
	NeedDetail *bool `query:"need_detail"`
	// PageNo represents page_no. 分页参数，当need_detail=true时需要，代表当前页数，从1开始。
	PageNo *int64 `query:"page_no"`
	// PageSize represents page_size. 分页参数，当need_detail=true时需要，代表每页大小，不能大于20。
	PageSize *int64 `query:"page_size"`
}

// ZmGoTemplateQueryRequest contains parameters for zhima.merchant.zmgo.template.query.
type ZmGoTemplateQueryRequest struct {
	// TemplateNo represents template_no. 模板ID，由创建模板接口返回结果提供
	TemplateNo *string `query:"template_no"`
	// PartnerId represents partner_id. 商户ID
	PartnerId *string `query:"partner_id"`
}
