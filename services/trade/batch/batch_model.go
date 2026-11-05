package batch

// BatchSettleStatus 是批量结算与请款分账明细的处理状态，
// 取值依据官方 V3 文档。
type BatchSettleStatus string

const (
	// BatchSettleStatusAcceptSuccess 表示受理成功。
	BatchSettleStatusAcceptSuccess BatchSettleStatus = "ACCEPT_SUCCESS"
	// BatchSettleStatusSuccess 表示结算成功。
	BatchSettleStatusSuccess BatchSettleStatus = "SUCCESS"
	// BatchSettleStatusFail 表示结算失败。
	BatchSettleStatusFail BatchSettleStatus = "FAIL"
	// BatchSettleStatusFailRetry 表示失败重试。
	BatchSettleStatusFailRetry BatchSettleStatus = "FAIL_RETRY"
)

// BatchRoyaltyClause 是请款分账的单条条款。
type BatchRoyaltyClause struct {
	// Scene 是分账场景：RETURN_MONEY（垫资还款）、CHARGE（费用）、
	// REPLENISH（补贴）、FUND_TRANS_IN_SWITCHING（存量资金搬迁）、
	// PUNISHMENT（处罚）。
	Scene *string `json:"scene,omitempty"`
	// Amount 是分账金额，单位为元。
	Amount *string `json:"amount,omitempty"`
	// Desc 是分账描述。
	Desc *string `json:"desc,omitempty"`
	// TransInAccountType 是分账账户类型：bankCard、alipayBalance 或
	// defaultSettle。
	TransInAccountType *string `json:"trans_in_account_type,omitempty"`
	// TransInAccountIDType 是分账账户 id 类型。
	TransInAccountIDType *string `json:"trans_in_account_id_type,omitempty"`
	// TransInAccountID 是分账转入账户 id。
	TransInAccountID *string `json:"trans_in_account_id,omitempty"`
	// TransInEntityType 是分账转入主体类型。
	TransInEntityType *string `json:"trans_in_entity_type,omitempty"`
	// TransInEntityID 是分账转入主体账号。
	TransInEntityID *string `json:"trans_in_entity_id,omitempty"`
	// TransInEntityBizType 是分账转入主体业务类型（目前仅支持 settled）。
	TransInEntityBizType *string `json:"trans_in_entity_biz_type,omitempty"`
	// TransInSubMerchant 是分账转入方二级商户信息。
	TransInSubMerchant *SubMerchant `json:"trans_in_sub_merchant,omitempty"`
	// TransOutEntityType 是分账转出主体类型。
	TransOutEntityType *string `json:"trans_out_entity_type,omitempty"`
	// TransOutEntityID 是分账转出主体账号。
	TransOutEntityID *string `json:"trans_out_entity_id,omitempty"`
	// TransOutEntityBizType 是分账转出方主体业务类型。
	TransOutEntityBizType *string `json:"trans_out_entity_biz_type,omitempty"`
	// TransOutSubMerchant 是分账转出方二级商户信息。
	TransOutSubMerchant *SubMerchant `json:"trans_out_sub_merchant,omitempty"`
}

// BatchRoyaltyDetail 是请款分账查询的单条明细。
type BatchRoyaltyDetail struct {
	// Status 是处理状态。
	Status *BatchSettleStatus `json:"status,omitempty"`
	// Amount 是分账金额，单位为元。
	Amount *string `json:"amount,omitempty"`
	// Desc 是分账描述。
	Desc *string `json:"desc,omitempty"`
	// TransInAccountType 是分账账户类型。
	TransInAccountType *string `json:"trans_in_account_type,omitempty"`
	// TransInAccountIDType 是分账账户 id 类型。
	TransInAccountIDType *string `json:"trans_in_account_id_type,omitempty"`
	// TransInAccountID 是分账转入账户 id。
	TransInAccountID *string `json:"trans_in_account_id,omitempty"`
	// TransInAccountOpenID 是分账转入方的 OpenId。
	TransInAccountOpenID *string `json:"trans_in_account_open_id,omitempty"`
	// TransInEntityType 是分账转入主体类型。
	TransInEntityType *string `json:"trans_in_entity_type,omitempty"`
	// TransInEntityID 是分账转入主体账号。
	TransInEntityID *string `json:"trans_in_entity_id,omitempty"`
	// TransInSubMerchant 是分账转入方二级商户信息。
	TransInSubMerchant *SubMerchant `json:"trans_in_sub_merchant,omitempty"`
	// TransOutEntityType 是分账转出主体类型。
	TransOutEntityType *string `json:"trans_out_entity_type,omitempty"`
	// TransOutEntityID 是分账转出主体账号。
	TransOutEntityID *string `json:"trans_out_entity_id,omitempty"`
	// TransOutSubMerchant 是分账转出方二级商户信息。
	TransOutSubMerchant *SubMerchant `json:"trans_out_sub_merchant,omitempty"`
	// ErrorCode 是错误编码。
	ErrorCode *string `json:"error_code,omitempty"`
	// ErrorDesc 是错误描述。
	ErrorDesc *string `json:"error_desc,omitempty"`
}

// SettleClause 是批量结算的明细条款。
type SettleClause struct {
	// SettleEntityType 是结算主体类型：SecondMerchant 或 Store。
	SettleEntityType *string `json:"settle_entity_type,omitempty"`
	// SettleEntityID 是结算主体账号：二级商户 ID 或门店 ID。
	SettleEntityID *string `json:"settle_entity_id,omitempty"`
	// SettleEntityBizType 是已结算资产类型。
	SettleEntityBizType *string `json:"settle_entity_biz_type,omitempty"`
	// SettleAccountType 是结算账户类型：bankCard、alipayBalance 或
	// defaultSettle（按默认结算规则结算）。
	SettleAccountType *string `json:"settle_account_type,omitempty"`
	// SettleAccountIDType 是结算账户 id 类型；defaultSettle 时必须为空。
	SettleAccountIDType *string `json:"settle_account_id_type,omitempty"`
	// SettleAccountID 是结算账户 id。
	SettleAccountID *string `json:"settle_account_id,omitempty"`
	// SettleAccountEntity 是结算账号所属的结算主体类型。
	SettleAccountEntity *string `json:"settle_account_entity,omitempty"`
	// Amount 是结算金额，单位为元。
	Amount *string `json:"amount,omitempty"`
	// Currency 是结算币种；目前支持人民币 CNY。
	Currency *string `json:"currency,omitempty"`
	// SubMerchant 是二级商户信息。
	SubMerchant *SubMerchant `json:"sub_merchant,omitempty"`
}

// SubMerchant 是支付宝二级商户信息。
type SubMerchant struct {
	// MerchantID 是支付宝二级商户编号；直付通和机构间连场景必传。
	MerchantID *string `json:"merchant_id,omitempty"`
	// MerchantType 是二级商户编号类型；目前仅支持 alipay。
	MerchantType *string `json:"merchant_type,omitempty"`
}

// BatchSettleDetail 是批量结算详情。
type BatchSettleDetail struct {
	// SettleEntityType 是结算主体类型。
	SettleEntityType *string `json:"settle_entity_type,omitempty"`
	// SettleEntityID 是结算主体账号。
	SettleEntityID *string `json:"settle_entity_id,omitempty"`
	// SettleAccountType 是结算账户类型。
	SettleAccountType *string `json:"settle_account_type,omitempty"`
	// SettleAccountIDType 是结算账户 id 类型。
	SettleAccountIDType *string `json:"settle_account_id_type,omitempty"`
	// SettleAccountID 是结算账户 id。
	SettleAccountID *string `json:"settle_account_id,omitempty"`
	// SettleAccountOpenID 是结算账户的 OpenId。
	SettleAccountOpenID *string `json:"settle_account_open_id,omitempty"`
	// Amount 是结算金额，单位为元。
	Amount *string `json:"amount,omitempty"`
	// Currency 是结算币种。
	Currency *string `json:"currency,omitempty"`
	// Status 是结算状态。
	Status *BatchSettleStatus `json:"status,omitempty"`
	// ErrorCode 是错误编码。
	ErrorCode *string `json:"error_code,omitempty"`
	// ErrorDesc 是错误描述。
	ErrorDesc *string `json:"error_desc,omitempty"`
	// SubMerchant 是二级商户信息。
	SubMerchant *SubMerchant `json:"sub_merchant,omitempty"`
}
