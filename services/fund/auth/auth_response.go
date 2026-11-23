package auth

// FundAuthOperationCancelResponse 是支付宝官方 V3 OpenAPI 模型。
type FundAuthOperationCancelResponse struct {
	// Action 是本次撤销触发的资金动作 close：关闭冻结明细，无资金解冻 unfreeze：产生了资金解冻。
	Action *string `json:"action,omitempty"`
	// AuthNo 是支付宝资金授权订单号。。
	AuthNo *string `json:"auth_no,omitempty"`
	// OperationId 是支付宝的冻结操作流水号。。
	OperationId *string `json:"operation_id,omitempty"`
	// OutOrderNo 是商户的授权资金订单号。。
	OutOrderNo *string `json:"out_order_no,omitempty"`
	// OutRequestNo 是商户的冻结操作流水号 。。
	OutRequestNo *string `json:"out_request_no,omitempty"`
}

// FundAuthOperationDetailQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type FundAuthOperationDetailQueryResponse struct {
	// Amount 是该笔资金操作流水operation_id对应的操作金额，单位为：元（人民币）。
	Amount *string `json:"amount,omitempty"`
	// AuthNo 是支付宝资金授权订单号。
	AuthNo *string `json:"auth_no,omitempty"`
	// CreditAmount 是该笔资金操作流水operation_id对应的操作信用金额。
	CreditAmount *string `json:"credit_amount,omitempty"`
	// CreditMerchantExt 是芝麻透出给商户的信息，具体内容由商户与芝麻约定后返回。
	CreditMerchantExt *string         `json:"credit_merchant_ext,omitempty"`
	CreditRiskInfo    *CreditRiskInfo `json:"credit_risk_info,omitempty"`
	// EnterprisePayInfo 是因公付业务信息。
	EnterprisePayInfo *string `json:"enterprise_pay_info,omitempty"`
	// ExtraParam 是商户请求创建预授权订单时传入的扩展参数，仅返回商户自定义的扩展信息（merchantExt）。
	ExtraParam *string `json:"extra_param,omitempty"`
	// FundAmount 是该笔资金操作流水operation_id对应的操作自有资金金额。
	FundAmount *string `json:"fund_amount,omitempty"`
	// GmtCreate 是资金授权单据操作流水创建时间， 格式：YYYY-MM-DD HH:MM:SS。
	GmtCreate *string `json:"gmt_create,omitempty"`
	// GmtTrans 是支付宝账务处理成功时间， 格式：YYYY-MM-DD HH:MM:SS。
	GmtTrans *string `json:"gmt_trans,omitempty"`
	// OperationId 是支付宝资金操作流水号。
	OperationId *string `json:"operation_id,omitempty"`
	// OperationType 是支付宝资金操作类型，表示当前查询到的这笔明细的操作类型。。
	OperationType *string `json:"operation_type,omitempty"`
	// OrderStatus 是授权单状态：INIT（初始状态：已创建未授权）、AUTHORIZED（已授权状态：授权成功，可以进行转支付或解冻操作）、FINISH（完成状态：转支付完成且无剩余冻结资金）、CLOSED（关闭状态：授权未完成超时关闭或冻结资金全额解冻）。
	OrderStatus *string `json:"order_status,omitempty"`
	// OrderTitle 是业务订单的简单描述，如商品名称等。
	OrderTitle *string `json:"order_title,omitempty"`
	// OutOrderNo 是商户的授权资金订单号。
	OutOrderNo *string `json:"out_order_no,omitempty"`
	// OutRequestNo 是商户资金操作的请求流水号。
	OutRequestNo *string `json:"out_request_no,omitempty"`
	// PayerLogonId 是付款方支付宝账号（Email或手机号），仅作展示使用，默认会加“*”号处理。
	PayerLogonId *string `json:"payer_logon_id,omitempty"`
	// PayerOpenId 是支付宝openId，用户（userId）在应用（appId）下的唯一标识。。
	PayerOpenId *string `json:"payer_open_id,omitempty"`
	// PayerUserId 是付款方支付宝账号对应的支付宝唯一用户号，以2088开头的16位纯数字组成。
	PayerUserId *string `json:"payer_user_id,omitempty"`
	// PreAuthType 是预授权类型，信用预授权情况下值为 CREDIT_AUTH，表示该笔预授权为信用预授权，实际没有冻结用户资金；其它情况均不返回该字段。。
	PreAuthType *string `json:"pre_auth_type,omitempty"`
	// Remark 是商户对本次操作的附言描述，长度不超过100个字母或50个汉字。
	Remark *string `json:"remark,omitempty"`
	// RestAmount 是订单当前剩余冻结金额，单位为：元（人民币）。。
	RestAmount *string `json:"rest_amount,omitempty"`
	// RestCreditAmount 是剩余冻结信用金额，单位为：元（人民币），精确到小数点后两位。
	RestCreditAmount *string `json:"rest_credit_amount,omitempty"`
	// RestFundAmount 是剩余冻结自有资金金额，单位为：元（人民币），精确到小数点后两位。
	RestFundAmount *string `json:"rest_fund_amount,omitempty"`
	// Status 是资金操作流水的状态， 目前支持： INIT：初始 SUCCESS：成功 CLOSED：关闭。
	Status *string `json:"status,omitempty"`
	// TotalFreezeAmount 是订单累计的冻结金额，单位为：元（人民币）。
	TotalFreezeAmount *string `json:"total_freeze_amount,omitempty"`
	// TotalFreezeCreditAmount 是累计冻结信用金额，单位为：元（人民币），精确到小数点后两位。
	TotalFreezeCreditAmount *string `json:"total_freeze_credit_amount,omitempty"`
	// TotalFreezeFundAmount 是累计冻结自有资金金额，单位为：元（人民币），精确到小数点后两位。
	TotalFreezeFundAmount *string `json:"total_freeze_fund_amount,omitempty"`
	// TotalPayAmount 是订单累计用于支付的金额，单位为：元（人民币）。
	TotalPayAmount *string `json:"total_pay_amount,omitempty"`
	// TotalPayCreditAmount 是累计支付信用金额，单位为：元（人民币），精确到小数点后两位。
	TotalPayCreditAmount *string `json:"total_pay_credit_amount,omitempty"`
	// TotalPayFundAmount 是累计支付自有资金金额，单位为：元（人民币），精确到小数点后两位。
	TotalPayFundAmount *string `json:"total_pay_fund_amount,omitempty"`
	// TransCurrency 是标价币种, amount 对应的币种单位。支持澳元：AUD, 新西兰元：NZD, 台币：TWD, 美元：USD, 欧元：EUR, 英镑：GBP。
	TransCurrency *string `json:"trans_currency,omitempty"`
}

// FundAuthOrderAppFreezeResponse 是支付宝官方 V3 OpenAPI 模型。
type FundAuthOrderAppFreezeResponse struct {
	// Amount 是本次操作冻结的金额，单位为：元（人民币），精确到小数点后两位。
	Amount *string `json:"amount,omitempty"`
	// AuthNo 是支付宝的资金授权订单号。
	AuthNo *string `json:"auth_no,omitempty"`
	// CreditAmount 是本次冻结操作中信用冻结金额，单位为：元（人民币），精确到小数点后两位。
	CreditAmount *string `json:"credit_amount,omitempty"`
	// FundAmount 是本次冻结操作中自有资金冻结金额，单位为：元（人民币），精确到小数点后两位。
	FundAmount *string `json:"fund_amount,omitempty"`
	// GmtTrans 是资金授权成功时间 格式：YYYY-MM-DD HH:MM:SS。
	GmtTrans *string `json:"gmt_trans,omitempty"`
	// OperationId 是支付宝的资金操作流水号。
	OperationId *string `json:"operation_id,omitempty"`
	// OrderStr 是获取签名后的业务数据。
	OrderStr *string `json:"orderStr,omitempty"`
	// OutOrderNo 是商户的授权资金订单号。
	OutOrderNo *string `json:"out_order_no,omitempty"`
	// OutRequestNo 是商户本次资金操作的请求流水号。
	OutRequestNo *string `json:"out_request_no,omitempty"`
	// PayerOpenId 是支付宝openId，用户（userId）在应用（appId）下的唯一用户标识。。
	PayerOpenId *string `json:"payer_open_id,omitempty"`
	// PayerUserId 是付款方支付宝用户号。
	PayerUserId *string `json:"payer_user_id,omitempty"`
	// PreAuthType 是预授权类型，目前支持 CREDIT_AUTH(信用预授权); 商户可根据该标识来判断该笔预授权的类型，当返回值为"CREDIT_AUTH"表明该笔预授权为信用预授权，没有真实冻结资金；当返回值为空或者不为"CREDIT_AUTH"则表明该笔预授权为普通资金预授权，会冻结用户资金。。
	PreAuthType *string `json:"pre_auth_type,omitempty"`
	// Status 是资金预授权明细的状态 目前支持： INIT：初始 SUCCESS: 成功 CLOSED：关闭。
	Status *string `json:"status,omitempty"`
	// TransCurrency 是标价币种, amount 对应的币种单位。支持澳元：AUD, 新西兰元：NZD, 台币：TWD, 美元：USD, 欧元：EUR, 英镑：GBP, 人民币：CNY。
	TransCurrency *string `json:"trans_currency,omitempty"`
}

// FundAuthOrderFreezeResponse 是支付宝官方 V3 OpenAPI 模型。
type FundAuthOrderFreezeResponse struct {
	// Amount 是本次操作冻结的金额，单位为：元（人民币），精确到小数点后两位。
	Amount *string `json:"amount,omitempty"`
	// AuthNo 是支付宝的资金授权订单号。
	AuthNo *string `json:"auth_no,omitempty"`
	// CreditAmount 是本次冻结操作中信用冻结金额，单位为：元（人民币），精确到小数点后两位。
	CreditAmount *string `json:"credit_amount,omitempty"`
	// FundAmount 是本次冻结操作中自有资金冻结金额，单位为：元（人民币），精确到小数点后两位。
	FundAmount *string `json:"fund_amount,omitempty"`
	// GmtTrans 是资金授权成功时间 格式：YYYY-MM-DD HH:MM:SS。
	GmtTrans *string `json:"gmt_trans,omitempty"`
	// OperationId 是支付宝的资金操作流水号。
	OperationId *string `json:"operation_id,omitempty"`
	// OutOrderNo 是商户的授权资金订单号。
	OutOrderNo *string `json:"out_order_no,omitempty"`
	// OutRequestNo 是商户本次资金操作的请求流水号。
	OutRequestNo *string `json:"out_request_no,omitempty"`
	// PayerLogonId 是付款方支付宝账号（Email或手机号）。
	PayerLogonId *string `json:"payer_logon_id,omitempty"`
	// PayerOpenId 是支付宝openId，用户（userId）在应用（appId）下的唯一用户标识。。
	PayerOpenId *string `json:"payer_open_id,omitempty"`
	// PayerUserId 是付款方支付宝用户号。
	PayerUserId *string `json:"payer_user_id,omitempty"`
	// PreAuthType 是预授权类型，目前支持 CREDIT_AUTH(信用预授权); 商户可根据该标识来判断该笔预授权的类型，当返回值为"CREDIT_AUTH"表明该笔预授权为信用预授权，没有真实冻结资金；当返回值为空或者不为"CREDIT_AUTH"则表明该笔预授权为普通资金预授权，会冻结用户资金。。
	PreAuthType *string `json:"pre_auth_type,omitempty"`
	// Status 是资金预授权明细的状态 目前支持： INIT：初始 SUCCESS: 成功 CLOSED：关闭。
	Status *string `json:"status,omitempty"`
	// TransCurrency 是标价币种, amount 对应的币种单位。支持澳元：AUD, 新西兰元：NZD, 台币：TWD, 美元：USD, 欧元：EUR, 英镑：GBP。
	TransCurrency *string `json:"trans_currency,omitempty"`
}

// FundAuthOrderUnfreezeResponse 是支付宝官方 V3 OpenAPI 模型。
type FundAuthOrderUnfreezeResponse struct {
	// Amount 是本次操作解冻的金额，单位为：元（人民币），精确到小数点后两位，取值范围：[0.01,100000000.00]。
	Amount *string `json:"amount,omitempty"`
	// AuthNo 是支付宝资金授权订单号。
	AuthNo *string `json:"auth_no,omitempty"`
	// CreditAmount 是本次解冻操作中信用解冻金额，单位为：元（人民币），精确到小数点后两位。
	CreditAmount *string `json:"credit_amount,omitempty"`
	// FundAmount 是本次解冻操作中自有资金解冻金额，单位为：元（人民币），精确到小数点后两位。
	FundAmount *string `json:"fund_amount,omitempty"`
	// GmtTrans 是授权资金解冻成功时间，格式：YYYY-MM-DD HH:MM:SS。
	GmtTrans *string `json:"gmt_trans,omitempty"`
	// OperationId 是支付宝资金操作流水号。
	OperationId *string `json:"operation_id,omitempty"`
	// OutOrderNo 是商户的授权资金订单号。
	OutOrderNo *string `json:"out_order_no,omitempty"`
	// OutRequestNo 是商户本次资金操作的请求流水号。
	OutRequestNo *string `json:"out_request_no,omitempty"`
	// Status 是资金操作流水的状态 目前支持：SUCCESS：成功。
	Status *string `json:"status,omitempty"`
}

// FundAuthOrderVoucherCreateResponse 是支付宝官方 V3 OpenAPI 模型。
type FundAuthOrderVoucherCreateResponse struct {
	// CodeType 是码类型，分为 barCode：条形码 (一维码) 和 qrCode:二维码(qrCode) ； 目前发码只支持 qrCode。
	CodeType *string `json:"code_type,omitempty"`
	// CodeUrl 是生成的带有支付宝logo的二维码地址，如：<a href="http://mobilecodec.alipay.com/show.htm?code=aeparsv2dknkqf3018556a">http://mobilecodec.alipay.com/show.htm?code=aeparsv2dknkqf3018556a</a>；商户端通过在末尾追加picSize来指定要显示的图片大小，如 显示1280大小的URL:<a href="http://mobilecodec.alipay.com/show.htm?code=aeparsv2dknkqf3018556a&picSize=1280">http://mobilecodec.alipay.com/show.htm?code=aeparsv2dknkqf3018556a&picSize=1280</a>；目前支持的大小有：256, 227, 270, 344, 430, 512, 570, 860, 1280, 1546；。
	CodeUrl *string `json:"code_url,omitempty"`
	// CodeValue 是当前发码请求生成的二维码码串，商户端可以利用二维码生成工具根据该码串值生成对应的二维码。
	CodeValue *string `json:"code_value,omitempty"`
	// OutOrderNo 是商户的授权资金订单号。
	OutOrderNo *string `json:"out_order_no,omitempty"`
	// OutRequestNo 是商户本次资金操作的请求流水号。
	OutRequestNo *string `json:"out_request_no,omitempty"`
}
