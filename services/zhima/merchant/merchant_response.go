// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package merchant

// SubsidiariesApplyResponse is the official V3 ZhimaMerchantSubsidiariesApplyResponseModel schema.
type SubsidiariesApplyResponse struct {
	// BizErrorCode represents biz_error_code. 业务错误码
	BizErrorCode *string `json:"biz_error_code,omitzero"`
	// BizErrorMessage represents biz_error_message. 业务错误信息
	BizErrorMessage *string `json:"biz_error_message,omitzero"`
	// OrderNo represents order_no. 工单标识
	OrderNo *string `json:"order_no,omitzero"`
}

// SubsidiariesCloseResponse is the official V3 ZhimaMerchantSubsidiariesCloseResponseModel schema.
type SubsidiariesCloseResponse struct {
	// BizErrorCode represents biz_error_code. 业务错误码
	BizErrorCode *string `json:"biz_error_code,omitzero"`
	// BizErrorMessage represents biz_error_message. 业务错误信息
	BizErrorMessage *string `json:"biz_error_message,omitzero"`
	// OrderMemo represents order_memo. 工单审核备注
	OrderMemo *string `json:"order_memo,omitzero"`
	// OrderNo represents order_no. 工单标识，业务成功时返回工单号
	OrderNo *string `json:"order_no,omitzero"`
	// OrderStatus represents order_status. 工单状态
	OrderStatus *string `json:"order_status,omitzero"`
}

// ZmGoCumulateQueryResponse is the official V3 ZhimaMerchantZmgoCumulateQueryResponseModel schema.
type ZmGoCumulateQueryResponse struct {
	// AggrAmount represents aggr_amount. 总共累计金额
	AggrAmount *string `json:"aggr_amount,omitzero"`
	// AggrDiscountAmount represents aggr_discount_amount. 总共累计优惠金额
	AggrDiscountAmount *string `json:"aggr_discount_amount,omitzero"`
	// AggrTimes represents aggr_times. 总共累计次数
	AggrTimes *int64 `json:"aggr_times,omitzero"`
	// AgreementId represents agreement_id. 芝麻go协议号，唯一标识一个芝麻go协议。
	AgreementId *string `json:"agreement_id,omitzero"`
	// DetailList represents detail_list. 详情数据列表
	DetailList *[]CumulateDataDetail `json:"detail_list,omitzero"`
	// FailReason represents fail_reason. 查询失败原因
	FailReason *string `json:"fail_reason,omitzero"`
	// PageNo represents page_no. 分页参数，返回传入的值。
	PageNo *int64 `json:"page_no,omitzero"`
	// PageSize represents page_size. 分页参数，返回传入的值。
	PageSize *int64 `json:"page_size,omitzero"`
}

// ZmGoCumulateSyncResponse is the official V3 ZhimaMerchantZmgoCumulateSyncResponseModel schema.
type ZmGoCumulateSyncResponse struct {
	// AgreementId represents agreement_id. 芝麻go协议号，唯一标识一个芝麻go协议。
	AgreementId *string `json:"agreement_id,omitzero"`
	// FailReason represents fail_reason. 数据回传失败原因描述
	FailReason *string `json:"fail_reason,omitzero"`
	// OutBizNo represents out_biz_no. 外部业务号，唯一标识一笔回传数据。
	OutBizNo *string `json:"out_biz_no,omitzero"`
}

// ZmGoTemplateCreateResponse is the official V3 ZhimaMerchantZmgoTemplateCreateResponseModel schema.
type ZmGoTemplateCreateResponse struct {
	// TemplateNo represents template_no. 芝麻GO模板号
	TemplateNo *string `json:"template_no,omitzero"`
}

// ZmGoTemplateQueryResponse is the official V3 ZhimaMerchantZmgoTemplateQueryResponseModel schema.
type ZmGoTemplateQueryResponse struct {
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
