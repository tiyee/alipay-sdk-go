package customs

// CustomsDeclareBuyerInfo 是报关订购人信息。
type CustomsDeclareBuyerInfo struct {
	// BuyerName 是订购人姓名。
	BuyerName *string `json:"buyer_name,omitempty"`
	// BuyerCertNo 是订购人身份证号。
	BuyerCertNo *string `json:"buyer_cert_no,omitempty"`
}

// CustomsDeclareRecordInfo 是一条报关记录。
type CustomsDeclareRecordInfo struct {
	// OutRequestNo 是报关请求号。
	OutRequestNo *string `json:"out_request_no,omitempty"`
	// TradeNo 是支付宝推送到海关的支付单据号。
	TradeNo *string `json:"trade_no,omitempty"`
	// AlipayDeclareNo 是支付宝报关流水号。
	AlipayDeclareNo *string `json:"alipay_declare_no,omitempty"`
	// Amount 是报关金额。
	Amount *string `json:"amount,omitempty"`
	// CustomsPlace 是本记录报关的海关编号。
	CustomsPlace *string `json:"customs_place,omitempty"`
	// Status 是报关单当前状态：ws（等待发送海关）、sending（已提交
	// 发送海关）、succ（海关返回受理成功）等，未知值原样保留。
	Status *string `json:"status,omitempty"`
	// CustomsResultCode 是海关回执结果码，原样透传。
	CustomsResultCode *string `json:"customs_result_code,omitempty"`
	// CustomsResultInfo 是海关回执结果描述。
	CustomsResultInfo *string `json:"customs_result_info,omitempty"`
	// CustomsResultReturnTime 是海关回执时间（yyyyMMddHHmmss）。
	CustomsResultReturnTime *string `json:"customs_result_return_time,omitempty"`
	// LastModifiedTime 是报关记录状态最后更新时间。
	LastModifiedTime *string `json:"last_modified_time,omitempty"`
	// IsSplit 标识是否拆单（T/F）。
	IsSplit *string `json:"is_split,omitempty"`
	// SubOutBizNo 是拆单子订单号。
	SubOutBizNo *string `json:"sub_out_biz_no,omitempty"`
	// MerchantCustomsCode 是商户海关备案编号。
	MerchantCustomsCode *string `json:"merchant_customs_code,omitempty"`
	// MerchantCustomsName 是商户海关备案名称。
	MerchantCustomsName *string `json:"merchant_customs_name,omitempty"`
	// Memo 是备注说明。
	Memo *string `json:"memo,omitempty"`
}
