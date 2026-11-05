package trans

// OrderQueryResponse 是转账订单查询的业务响应。
type OrderQueryResponse struct {
	// OrderID 是支付宝转账单据号。
	OrderID *string `json:"order_id,omitempty"`
	// OutBizNo 是商户转账唯一订单号。
	OutBizNo *string `json:"out_biz_no,omitempty"`
	// Status 是转账单据状态；SUCCESS 配合单笔转账到银行卡产品时
	// 仍可能变为退票状态，需持续关注。
	Status *TransStatus `json:"status,omitempty"`
	// PayDate 是支付时间；转账失败不返回。
	PayDate *string `json:"pay_date,omitempty"`
	// ArrivalTimeEnd 是预计到账时间；仅供参考，不能作为实际到账时间。
	ArrivalTimeEnd *string `json:"arrival_time_end,omitempty"`
	// OrderFee 是预计收费金额；转账到银行卡专用。
	OrderFee *string `json:"order_fee,omitempty"`
	// ErrorCode 是查询到失败状态时的错误代码。
	ErrorCode *string `json:"error_code,omitempty"`
	// FailReason 是失败或退票的具体原因。
	FailReason *string `json:"fail_reason,omitempty"`
}

// CommonQueryResponse 是转账单据通用查询的业务响应。
type CommonQueryResponse struct {
	// OrderID 是支付宝转账单据号。
	OrderID *string `json:"order_id,omitempty"`
	// OutBizNo 是商户订单号。
	OutBizNo *string `json:"out_biz_no,omitempty"`
	// Status 是转账单据状态。
	Status *TransStatus `json:"status,omitempty"`
	// PayDate 是支付时间；转账失败不返回。
	PayDate *string `json:"pay_date,omitempty"`
	// TransAmount 是付款金额；订单状态为 SUCCESS 才返回。
	TransAmount *string `json:"trans_amount,omitempty"`
	// ReceiverUserID 是收款方支付宝账号 uid。
	ReceiverUserID *string `json:"receiver_user_id,omitempty"`
	// ReceiverOpenID 是收款方支付宝账号 openid。
	ReceiverOpenID *string `json:"receiver_open_id,omitempty"`
	// ErrorCode 是失败或退票时的错误代码。
	ErrorCode *string `json:"error_code,omitempty"`
	// FailReason 是失败或退票的具体原因。
	FailReason *string `json:"fail_reason,omitempty"`
	// SubStatus 是转账单据子状态（二阶段转账领取/退款）。
	SubStatus *string `json:"sub_status,omitempty"`
	// PayFundOrderID 是支付宝支付资金流水号。
	PayFundOrderID *string `json:"pay_fund_order_id,omitempty"`
	// PassbackParams 是商户回传参数。
	PassbackParams *string `json:"passback_params,omitempty"`
}

// PayeeBindQueryResponse 是收款账号绑定查询的业务响应。
type PayeeBindQueryResponse struct {
	// Bind 标识是否绑定收款账号（"true"/"false"）。
	Bind *string `json:"bind,omitempty"`
}

// UniTransferResponse 是单笔转账的业务响应。
// status 为 SUCCESS 才表示转账成功；DEALING/REFUND 等状态需结合
// 查询接口与异步通知确认。
type UniTransferResponse struct {
	// OrderID 是支付宝转账订单号。
	OrderID *string `json:"order_id,omitempty"`
	// OutBizNo 是商户订单号。
	OutBizNo *string `json:"out_biz_no,omitempty"`
	// PayFundOrderID 是支付宝支付资金流水号。
	PayFundOrderID *string `json:"pay_fund_order_id,omitempty"`
	// Status 是转账单据状态：SUCCESS、FAIL、DEALING、REFUND 等。
	Status *TransStatus `json:"status,omitempty"`
	// TransDate 是订单支付时间（yyyy-MM-dd HH:mm:ss）。
	TransDate *string `json:"trans_date,omitempty"`
	// Amount 是转账金额，特殊场景提供。
	Amount *string `json:"amount,omitempty"`
	// SubStatus 是转账单据子状态（二阶段转账领取/退款）。
	SubStatus *string `json:"sub_status,omitempty"`
	// SubPayFundOrderID 是支付宝支付子单资金流水号；二阶段主子单
	// 场景使用该单号下载电子回单。
	SubPayFundOrderID *string `json:"sub_pay_fund_order_id,omitempty"`
	// Link 是访问链接；多阶段转账时后续阶段跳转支付宝端内领取使用。
	Link *string `json:"link,omitempty"`
	// SettleSerialNo 是金融机构业务流水号。
	SettleSerialNo *string `json:"settle_serial_no,omitempty"`
}

// ToAccountTransferResponse 是单笔转账到支付宝账户的业务响应。
type ToAccountTransferResponse struct {
	// OrderID 是支付宝转账单据号；成功一定返回，失败可能不返回。
	OrderID *string `json:"order_id,omitempty"`
	// OutBizNo 是商户转账唯一订单号，原样返回。
	OutBizNo *string `json:"out_biz_no,omitempty"`
	// PayDate 是支付时间；仅转账成功返回。
	PayDate *string `json:"pay_date,omitempty"`
}

// PayResponse 是资金支付的业务响应。
// status 为 UNKNOWN 时必须通过查询接口确认最新状态，不能视为成功。
type PayResponse struct {
	// OrderID 是该笔转账在支付宝系统内部的单据 ID。
	OrderID *string `json:"order_id,omitempty"`
	// OutBizNo 是商户端的唯一订单号。
	OutBizNo *string `json:"out_biz_no,omitempty"`
	// PayFundOrderID 是支付宝支付资金流水号。
	PayFundOrderID *string `json:"pay_fund_order_id,omitempty"`
	// Status 是支付状态：SUCCESS、FAIL、UNKNOWN。
	Status *TransStatus `json:"status,omitempty"`
	// TransPayTime 是订单支付时间。
	TransPayTime *string `json:"trans_pay_time,omitempty"`
}

// RefundResponse 是资金退回的业务响应。
type RefundResponse struct {
	// OrderID 是原转账的支付宝订单号。
	OrderID *string `json:"order_id,omitempty"`
	// OutRequestNo 是资金退回请求号。
	OutRequestNo *string `json:"out_request_no,omitempty"`
	// RefundOrderID 是退款的支付宝系统内部单据 id。
	RefundOrderID *string `json:"refund_order_id,omitempty"`
	// RefundAmount 是本次退款的金额。
	RefundAmount *string `json:"refund_amount,omitempty"`
	// RefundDate 是退款资金退回时间。
	RefundDate *string `json:"refund_date,omitempty"`
	// Status 是退款状态；SUCCESS 表示退款成功。
	Status *TransStatus `json:"status,omitempty"`
}
