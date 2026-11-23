package customs

// CustomsDeclareResponse 是报关的业务响应。
type CustomsDeclareResponse struct {
	// AlipayDeclareNo 是支付宝报关流水号。
	AlipayDeclareNo *string `json:"alipay_declare_no,omitempty"`
	// TradeNo 是支付宝推送到海关的支付单据号。
	TradeNo *string `json:"trade_no,omitempty"`
	// OutTradeNo 是国际站外部订单号。
	OutTradeNo *string `json:"out_trade_no,omitempty"`
	// IdentityCheck 是订购人与支付人身份信息验证结果（T/F）；
	// 同一笔报关单只校验一次。
	IdentityCheck *string `json:"identity_check,omitempty"`
	// Currency 是币种。
	Currency *string `json:"currency,omitempty"`
	// TotalAmount 是交易总金额。
	TotalAmount *string `json:"total_amount,omitempty"`
	// PayCode 是支付机构注册号。
	PayCode *string `json:"pay_code,omitempty"`
	// PayTransactionID 是清算流水号。
	PayTransactionID *string `json:"pay_transaction_id,omitempty"`
	// VerDept 是清算机构标志。
	VerDept *string `json:"ver_dept,omitempty"`
}

// CustomsQueryResponse 是报关查询的业务响应。
type CustomsQueryResponse struct {
	// Records 是匹配到的报关记录。
	Records []CustomsDeclareRecordInfo `json:"records,omitempty"`
	// NotFound 是不存在记录的报关请求号，逗号分隔。
	NotFound *string `json:"not_found,omitempty"`
}
