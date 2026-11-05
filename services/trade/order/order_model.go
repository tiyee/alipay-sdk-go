package order

// RoyaltyDetailState 是分账明细的执行状态，取值依据官方 V3 文档。
type RoyaltyDetailState string

const (
	// RoyaltyDetailStateSuccess 表示分账成功。
	RoyaltyDetailStateSuccess RoyaltyDetailState = "SUCCESS"
	// RoyaltyDetailStateFail 表示分账失败。
	RoyaltyDetailStateFail RoyaltyDetailState = "FAIL"
	// RoyaltyDetailStateProcessing 表示分账处理中。
	RoyaltyDetailStateProcessing RoyaltyDetailState = "PROCESSING"
)

// RoyaltyMode 是 alipay.trade.order.settle 的分账执行模式，
// 取值依据官方 V3 文档。
type RoyaltyMode string

const (
	// RoyaltyModeSync 表示分账同步执行。
	RoyaltyModeSync RoyaltyMode = "sync"
	// RoyaltyModeAsync 表示分账异步执行。
	RoyaltyModeAsync RoyaltyMode = "async"
)

// IsKnown 报告分账模式是否属于文档已知值。
func (m RoyaltyMode) IsKnown() bool {
	return m == RoyaltyModeSync || m == RoyaltyModeAsync
}

// BuyerPayDetail 是买家支付明细。
type BuyerPayDetail struct {
	// PayChannel 是支付渠道（如 offline_pos）。
	PayChannel string `json:"pay_channel"`
	// Amount 是买家支付金额，单位为元。
	Amount string `json:"amount"`
}

// OrderSettleExtendParams 是分账扩展参数。
type OrderSettleExtendParams struct {
	// RoyaltyFinish 标识该交易分账是否完结（true/false）；
	// true 表示本次分账处理完成后解冻剩余冻结金额。
	RoyaltyFinish *string `json:"royalty_finish,omitempty"`
	// RoyaltyFinishAmount 是部分分账完结的金额；传入时 RoyaltyFinish 须为 true。
	RoyaltyFinishAmount *string `json:"royalty_finish_amount,omitempty"`
}

// RoyaltyDetail 是分账明细的查询结果。
type RoyaltyDetail struct {
	// DetailID 是支付宝分账明细单号。
	DetailID *string `json:"detail_id,omitempty"`
	// OperationType 是分账操作类型：replenish（补差）、
	// replenish_refund（退补差）、transfer（分账）、transfer_refund（退分账）。
	OperationType *string `json:"operation_type,omitempty"`
	// State 是分账状态：SUCCESS、FAIL、PROCESSING。
	State *RoyaltyDetailState `json:"state,omitempty"`
	// Amount 是分账金额。
	Amount *string `json:"amount,omitempty"`
	// TransIn 是分账转入账号；仅补差退回与分账类型返回。
	TransIn *string `json:"trans_in,omitempty"`
	// TransInOpenID 是分账转入方的 OpenId。
	TransInOpenID *string `json:"trans_in_open_id,omitempty"`
	// TransInType 是分账转入账号类型。
	TransInType *string `json:"trans_in_type,omitempty"`
	// TransOut 是分账转出账号；仅补差与退分账类型返回。
	TransOut *string `json:"trans_out,omitempty"`
	// TransOutOpenID 是分账转出方的 OpenId。
	TransOutOpenID *string `json:"trans_out_open_id,omitempty"`
	// TransOutType 是分账转出账号类型。
	TransOutType *string `json:"trans_out_type,omitempty"`
	// ErrorCode 是分账失败错误码；只在分账失败时返回。
	ErrorCode *string `json:"error_code,omitempty"`
	// ErrorDesc 是分账错误描述。
	ErrorDesc *string `json:"error_desc,omitempty"`
	// ExecuteDt 是分账执行时间。
	ExecuteDt *string `json:"execute_dt,omitempty"`
}
