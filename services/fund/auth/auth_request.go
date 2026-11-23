package auth

// FundAuthOperationCancelRequest 是支付宝官方 V3 OpenAPI 模型。
type FundAuthOperationCancelRequest struct {
	// AuthNo 是支付宝授权资金订单号。 与商户的授权资金订单号不能同时为空，二者都传入时，以支付宝资金授权订单号为准，该参数与支付宝授权资金操作流水号配对使用。。
	AuthNo *string `json:"auth_no,omitempty"`
	// NotifyUrl 是通知地址。
	NotifyUrl *string `json:"notify_url,omitempty"`
	// OperationId 是支付宝的授权资金操作流水号。 与商户的授权资金操作流水号不能同时为空，二者都传入时，以支付宝的授权资金操作流水号为准，该参数与支付宝授权资金订单号配对使用。。
	OperationId *string `json:"operation_id,omitempty"`
	// OutOrderNo 是商户的授权资金订单号。 与支付宝的授权资金订单号不能同时为空，二者都传入时，以支付宝的授权资金订单号为准，该参数与商户的授权资金操作流水号配对使用。 该值与资金冻结时 out_order_no一致。。
	OutOrderNo *string `json:"out_order_no,omitempty"`
	// OutRequestNo 是商户的授权资金操作流水号。 与支付宝的授权资金操作流水号不能同时为空，二者都传入时，以支付宝的授权资金操作流水号为准，该参数与商户的授权资金订单号配对使用。 该值与资金冻结时out_request_no一致。
	OutRequestNo *string `json:"out_request_no,omitempty"`
	// Remark 是商户对本次撤销操作的附言描述。 长度不超过100个字母或50个汉字。。
	Remark *string `json:"remark,omitempty"`
}

// FundAuthOperationDetailQueryRequest 是支付宝官方 V3 OpenAPI 模型。
type FundAuthOperationDetailQueryRequest struct {
	// AuthNo 是支付宝授权资金订单号。 与商户的授权资金订单号不能同时为空，二者都传入时，以支付宝资金授权订单号为准，该参数与支付宝授权资金操作流水号配对使用。。
	AuthNo *string `json:"auth_no,omitempty"`
	// OperationId 是支付宝的授权资金操作流水号。 与商户的授权资金操作流水号不能同时为空，二者都传入时，以支付宝的授权资金操作流水号为准，该参数与支付宝授权资金订单号配对使用。。
	OperationId *string `json:"operation_id,omitempty"`
	// OperationType 是需要查询的授权资金操作类型。 未传入本参数时，如果仅查询出单笔明细则直接返回，如果查询出多笔则优先返回冻结明细、无冻结明细时返回解冻明细； 当传入本参数时，则严格按照该操作类型返回对应明细。
	OperationType *string `json:"operation_type,omitempty"`
	// OutOrderNo 是商户的授权资金订单号。 与支付宝的授权资金订单号不能同时为空，二者都传入时，以支付宝的授权资金订单号为准，该参数与商户的授权资金操作流水号配对使用。 该值与资金冻结时 out_order_no一致。。
	OutOrderNo *string `json:"out_order_no,omitempty"`
	// OutRequestNo 是商户的授权资金操作流水号。 与支付宝的授权资金操作流水号不能同时为空，二者都传入时，以支付宝的授权资金操作流水号为准，该参数与商户的授权资金订单号配对使用。 查询冻结明细时，该值与发起冻结操作时传入的out_request_no一致； 查询解冻明细时，该值与发起解冻操作时传入的out_request_no一致； 查询支付明细时，该值与发起转支付操作时传入的out_trade_no一致。。
	OutRequestNo *string `json:"out_request_no,omitempty"`
	// QueryOptions 是需要查询的额外信息。
	QueryOptions *[]string `json:"query_options,omitempty"`
}

// FundAuthOrderAppFreezeRequest 是支付宝官方 V3 OpenAPI 模型。
type FundAuthOrderAppFreezeRequest struct {
	// Amount 是需要冻结的金额，单位为：元（人民币），精确到小数点后两位。 取值范围：[0.01,100000000.00]。
	Amount *string `json:"amount,omitempty"`
	// BusinessParams 是业务参数，如风控参数outRiskInfo等。。
	BusinessParams *string `json:"business_params,omitempty"`
	// DepositProductMode 是免押受理台模式，使用免押产品必传该字段。根据免押不同业务模式将开通受理台区分三种模式，商家可根据调用预授权冻结接口传入的参数决定该笔免押订单选择哪种受理台模式。不同受理台模式需要传入不同参数，其中：POSTPAY 表示后付金额已知，POSTPAY_UNCERTAIN 表示后付金额未知，DEPOSIT_ONLY 表示纯免押。 具体规则参考文档：<a href="https://opendocs.alipay.com/b/08tf3t?pathHash=d67d7545">https://opendocs.alipay.com/b/08tf3t?pathHash=d67d7545</a>。
	DepositProductMode *string `json:"deposit_product_mode,omitempty"`
	// DisablePayChannels 是无特殊需要请勿传入；商户可用该参数禁用支付渠道。 传入后用户不可使用列表中的渠道进行支付，目前支持两种禁用渠道：信用卡快捷（OPTIMIZED_MOTO）、信用卡卡通（BIGAMOUNT_CREDIT_CARTOON）。与可用支付渠道不能同时传入。
	DisablePayChannels *string `json:"disable_pay_channels,omitempty"`
	// EnablePayChannels 是无特殊需要请勿传入；商户可用该参数指定支付渠道。 传入后用户仅能使用列表中的渠道进行支付，目前支持三种渠道，余额宝（MONEY_FUND）、花呗（PCREDIT_PAY）以及芝麻信用（CREDITZHIMA）。与禁用支付渠道不可同时传入。
	EnablePayChannels *string `json:"enable_pay_channels,omitempty"`
	// ExtraParam 是业务扩展参数，用于特定业务信息的传递，json格式。 1、category，信用类目，信用预授权场景必传，具体类目信息见<a href="https://opendocs.alipay.com/open/10719">https://opendocs.alipay.com/open/10719</a>； 2、serviceId，信用服务ID：信用预授权场景必传。需要商家在 <a href="https://b.alipay.com/page/zmgaia/pre-auth/index">开放平台-芝麻免押-信用服务管理</a> 创建信用服务获取，详情可查看 <a href="https://opendocs.alipay.com/open/03w0a7?pathHash=51f6b4f2&ref=api#%E7%94%B3%E8%AF%B7%E4%BF%A1%E7%94%A8%E6%9C%8D%E5%8A%A1">创建信用服务</a>。在创建过程中如果有其它疑问，可以咨询芝麻客服小二（0571-88158055 转 2）； 3、creditExtInfo，信用参数，可选，如有需要请与芝麻约定后传入，信用服务说明见<a href="https://opendocs.alipay.com/open/11157/qlsxya">https://opendocs.alipay.com/open/11157/qlsxya</a>。
	ExtraParam *string `json:"extra_param,omitempty"`
	// IdentityParams 是无特殊需要请勿传入；买家实名信息。 传入后支付宝会比对买家在支付宝端的实名信息。包含两个可选key： 1.identity_hash，买家姓名拼接身份证号后，使用SHA256摘要方式与UTF8编码后的hash值，返回的十六进制字符串，例如 “张三4566498798498498498498”对应的identity_hash应为“acc2b92ffc5ed9b472faa19748f10045c30434132784f774b00216a56b8841c6” 2.alipay_user_id，买家uid。
	IdentityParams *string `json:"identity_params,omitempty"`
	// NotifyUrl 是通知地址。
	NotifyUrl *string `json:"notify_url,omitempty"`
	// OrderTitle 是订单标题。 业务订单的简单描述，如商品名称等。
	OrderTitle *string `json:"order_title,omitempty"`
	// OutOrderNo 是商户授权资金订单号。 商家自定义需保证在商户端不重复。仅支持字母、数字、下划线。。
	OutOrderNo *string `json:"out_order_no,omitempty"`
	// OutRequestNo 是商户本次资金操作的请求流水号，用于标示请求流水的唯一性。 可与out_order_no相同，仅支持字母、数字、下划线。。
	OutRequestNo *string `json:"out_request_no,omitempty"`
	// PayTimeout 是预授权订单相对超时时间。从商户客户端请求时间开始计算。 预授权订单允许的最晚授权时间，逾期将关闭该笔订单。取值范围：1m～15d。m-分钟，h-小时，d-天。 该参数数值不接受小数点， 如 1.5h，可转换为90m。 默认为15m。。
	PayTimeout *string `json:"pay_timeout,omitempty"`
	// PayeeLogonId 是收款账户的支付宝登录号（email或手机号）。 如果传入则会校验该登录号对应的账号是否具备当前商户收款权限，如果商户希望用户能够使用花呗，则用户号(payee_user_id)和登录号(payee_logon_id)两者必须传入其一。
	PayeeLogonId *string `json:"payee_logon_id,omitempty"`
	// PayeeUserId 是收款账户的支付宝用户号。 以2088开头的16位纯数字，如果传入则会校验该账号是否具备当前商户收款权限，如果商户希望用户能够使用花呗，则用户号(payee_user_id)和登录号(payee_logon_id)两者必须传入其一。
	PayeeUserId *string `json:"payee_user_id,omitempty"`
	// PostPayments 是后付费项目， 有付费项目时需要传入该字段。不同受理台模式需要传入不同参数，后付费项目名称和计费说明需要通过校验规则，同时计费说明将展示在开通受理台上。当受理台模式（deposit_product_mode）传入POSTPAY 时，后付费项目名称（name）、金额（amount）必传，计费说明（description）选传；当传入 POSTPAY_UNCERTAIN 时，后付费项目名称（name）、计费说明（description）必传，金额（amount）不传。 具体规则参考文档： <a href="https://opendocs.alipay.com/b/08tf3t?pathHash=d67d7545">https://opendocs.alipay.com/b/08tf3t?pathHash=d67d7545</a>。
	PostPayments *[]PostPayment `json:"post_payments,omitempty"`
	// ProductCode 是销售产品码。 预授权支付产品固定为 PREAUTH_PAY。
	ProductCode *string `json:"product_code,omitempty"`
	// SceneCode 是场景码，用于区分预授权不同业务场景。如：当面预授权通用场景（O2O_AUTH_COMMON_SCENE）、支付宝预授权通用场景（ONLINE_AUTH_COMMON_SCENE）、境外当面预授权通用场景（OVERSEAS_O2O_AUTH_COMMON_SCENE）、境外支付预授权通用场景（OVERSEAS_ONLINE_AUTH_COMMON_SCENE）等。
	SceneCode *string `json:"scene_code,omitempty"`
	// SettleCurrency 是商户指定的结算币种。支持澳元：AUD, 新西兰元：NZD, 台币：TWD, 美元：USD, 欧元：EUR, 英镑：GBP。
	SettleCurrency *string `json:"settle_currency,omitempty"`
	// TimeoutExpress 是预授权订单相对超时时间。从商户客户端请求时间开始计算。 预授权订单允许的最晚授权时间，逾期将关闭该笔订单。取值范围：1m～15d。m-分钟，h-小时，d-天。 该参数数值不接受小数点， 如 1.5h，可转换为90m。 默认为15m。。
	TimeoutExpress *string `json:"timeout_express,omitempty"`
	// TransCurrency 是标价币种, amount 对应的币种单位。支持澳元：AUD, 新西兰元：NZD, 台币：TWD, 美元：USD, 欧元：EUR, 英镑：GBP, 人民币：CNY。
	TransCurrency *string `json:"trans_currency,omitempty"`
}

// FundAuthOrderFreezeRequest 是支付宝官方 V3 OpenAPI 模型。
type FundAuthOrderFreezeRequest struct {
	// Amount 是需要冻结的金额，单位为：元（人民币），精确到小数点后两位。 取值范围：[0.01,100000000.00]。
	Amount *string `json:"amount,omitempty"`
	// AuthCode 是用户付款码。 1.条码场景：25~30开头的长度为16~24位的数字，实际字符串长度以开发者获取的付款码长度为准； 2.刷脸场景： 1）fp开头的35位字符串； 2）300-700字符的随机字符串； 注：刷脸场景考虑到未来可能拓展更多格式，建议外围不必做规则拦截，由支付宝统一做有效性校验。
	AuthCode *string `json:"auth_code,omitempty"`
	// AuthCodeType 是付款码类型。 1.条码场景：bar_code 2.刷脸场景：security_code。
	AuthCodeType *string `json:"auth_code_type,omitempty"`
	// BusinessParams 是业务参数，如风控参数outRiskInfo等。。
	BusinessParams *string `json:"business_params,omitempty"`
	// DepositProductMode 是免押受理台模式，使用免押产品必传该字段。根据免押不同业务模式将开通受理台区分三种模式，商家可根据调用预授权冻结接口传入的参数决定该笔免押订单选择哪种受理台模式。不同受理台模式需要传入不同参数，其中：POSTPAY 表示后付金额已知，POSTPAY_UNCERTAIN 表示后付金额未知，DEPOSIT_ONLY 表示纯免押。 具体规则参考文档：<a href="https://opendocs.alipay.com/b/08tf3t?pathHash=d67d7545">https://opendocs.alipay.com/b/08tf3t?pathHash=d67d7545</a>。
	DepositProductMode *string `json:"deposit_product_mode,omitempty"`
	// DisablePayChannels 是无特殊需要请勿传入；商户可用该参数禁用支付渠道。 传入后用户不可使用列表中的渠道进行支付，目前支持两种禁用渠道：信用卡快捷（OPTIMIZED_MOTO）、信用卡卡通（BIGAMOUNT_CREDIT_CARTOON）。与可用支付渠道不能同时传入。
	DisablePayChannels *string `json:"disable_pay_channels,omitempty"`
	// EnablePayChannels 是无特殊需要请勿传入；商户可用该参数指定支付渠道。 传入后用户仅能使用列表中的渠道进行支付，目前支持三种渠道，余额宝（MONEY_FUND）、花呗（PCREDIT_PAY）以及芝麻信用（CREDITZHIMA）。与禁用支付渠道不可同时传入。
	EnablePayChannels *string `json:"enable_pay_channels,omitempty"`
	// ExtraParam 是业务扩展参数，用于特定业务信息的传递，json格式。 1、category，信用类目，信用预授权场景必传，具体类目信息见<a href="https://opendocs.alipay.com/open/10719">https://opendocs.alipay.com/open/10719</a>； 2、serviceId，信用服务ID：信用预授权场景必传。需要商家在 <a href="https://b.alipay.com/page/zmgaia/pre-auth/index">开放平台-芝麻免押-信用服务管理</a> 创建信用服务获取，详情可查看 <a href="https://opendocs.alipay.com/open/03w0a7?pathHash=51f6b4f2&ref=api#%E7%94%B3%E8%AF%B7%E4%BF%A1%E7%94%A8%E6%9C%8D%E5%8A%A1">创建信用服务</a>。在创建过程中如果有其它疑问，可以咨询芝麻客服小二（0571-88158055 转 2）； 3、creditExtInfo，信用参数，可选，如有需要请与芝麻约定后传入，信用服务说明见<a href="https://opendocs.alipay.com/open/11157/qlsxya">https://opendocs.alipay.com/open/11157/qlsxya</a>。
	ExtraParam *string `json:"extra_param,omitempty"`
	// IdentityParams 是用户实名信息参数，包含：姓名+身份证号的hash值、指定用户的uid。商户传入用户实名信息参数，支付宝会对比用户在支付宝端的实名信息。 姓名+身份证号hash值使用SHA256摘要方式与UTF8编码,返回十六进制的字符串。 identity_hash和alipay_user_id都是可选的，如果两个都传，则会先校验identity_hash，然后校验alipay_user_id。其中identity_hash的待加密字样如"张三4566498798498498498498"。
	IdentityParams *string `json:"identity_params,omitempty"`
	// NotifyUrl 是通知地址。
	NotifyUrl *string `json:"notify_url,omitempty"`
	// OrderTitle 是订单标题。 业务订单的简单描述，如商品名称等。
	OrderTitle *string `json:"order_title,omitempty"`
	// OutOrderNo 是商户授权资金订单号。 商家自定义需保证在商户端不重复。仅支持字母、数字、下划线。。
	OutOrderNo *string `json:"out_order_no,omitempty"`
	// OutRequestNo 是商户本次资金操作的请求流水号，用于标示请求流水的唯一性。 可与out_order_no相同，仅支持字母、数字、下划线。。
	OutRequestNo *string `json:"out_request_no,omitempty"`
	// PayTimeout 是预授权订单相对超时时间，从商户请求时间开始计算。 预授权订单允许的最晚授权时间，逾期将关闭该笔订单。取值范围：1m～15d。m-分钟，h-小时，d-天。 该参数数值不接受小数点， 如 1.5h，可转换为90m。 默认为15m。。
	PayTimeout *string `json:"pay_timeout,omitempty"`
	// PayeeLogonId 是收款账户的支付宝登录号（email或手机号）。 如果传入则会校验该登录号对应的账号是否具备当前商户收款权限，如果商户希望用户能够使用花呗，则用户号(payee_user_id)和登录号(payee_logon_id)两者必须传入其一。
	PayeeLogonId *string `json:"payee_logon_id,omitempty"`
	// PayeeUserId 是收款账户的支付宝用户号。 以2088开头的16位纯数字，如果传入则会校验该账号是否具备当前商户收款权限，如果商户希望用户能够使用花呗，则用户号(payee_user_id)和登录号(payee_logon_id)两者必须传入其一。
	PayeeUserId *string `json:"payee_user_id,omitempty"`
	// PostPayments 是后付费项目， 有付费项目时需要传入该字段。不同受理台模式需要传入不同参数，后付费项目名称和计费说明需要通过校验规则，同时计费说明将展示在开通受理台上。当受理台模式（deposit_product_mode）传入POSTPAY 时，后付费项目名称（name）、金额（amount）必传，计费说明（description）选传；当传入 POSTPAY_UNCERTAIN 时，后付费项目名称（name）、计费说明（description）必传，金额（amount）不传。 具体规则参考文档： <a href="https://opendocs.alipay.com/b/08tf3t?pathHash=d67d7545">https://opendocs.alipay.com/b/08tf3t?pathHash=d67d7545</a>。
	PostPayments *[]PostPayment `json:"post_payments,omitempty"`
	// ProductCode 是销售产品码。 当面资金预授权固定为 PRE_AUTH。
	ProductCode *string `json:"product_code,omitempty"`
	// SceneCode 是场景码。 刷脸场景下传入HOTEL，其他情况下无需传入。
	SceneCode *string `json:"scene_code,omitempty"`
	// SettleCurrency 是商户指定的结算币种。支持澳元：AUD, 新西兰元：NZD, 台币：TWD, 美元：USD, 欧元：EUR, 英镑：GBP。
	SettleCurrency *string `json:"settle_currency,omitempty"`
	// TerminalParams 是机具管控sdk加签参数，参数示例 "terminal_params":"{"terminalType":"IOT","signature":"QIIAX8DqbFbNf2oe97FI1RSLAycC/tU4GVjer3bN8K4qLtAB","apdidToken":"xPA3ptuArwYc3F6Va_pjVwv7Qx7Tg5TJdrA_Jb_moYte9AqGZgEAAA==","hardToken":"","time":"1539847253","bizCode":"11000200040004000121","bizTid":"010100F01i1XyacMgpOinHerfdBw1xA9dNDocctlnqhLD8lfODr1A7Q","signedKeys":"authCode,totalAmount,apdidToken,hardToken,time,bizCode,bizTid"}"。
	TerminalParams *string `json:"terminal_params,omitempty"`
	// TimeoutExpress 是预授权订单相对超时时间，从商户请求时间开始计算。 预授权订单允许的最晚授权时间，逾期将关闭该笔订单。取值范围：1m～15d。m-分钟，h-小时，d-天。 该参数数值不接受小数点， 如 1.5h，可转换为90m。 默认为15m。。
	TimeoutExpress *string `json:"timeout_express,omitempty"`
	// TransCurrency 是标价币种, amount 对应的币种单位。支持澳元：AUD, 新西兰元：NZD, 台币：TWD, 美元：USD, 欧元：EUR, 英镑：GBP。
	TransCurrency *string `json:"trans_currency,omitempty"`
}

// FundAuthOrderUnfreezeRequest 是支付宝官方 V3 OpenAPI 模型。
type FundAuthOrderUnfreezeRequest struct {
	// Amount 是本次操作解冻的金额，单位为：元（人民币），精确到小数点后两位。 取值范围：[0.01,100000000.00]。
	Amount *string `json:"amount,omitempty"`
	// AuthNo 是支付宝资金授权订单号。。
	AuthNo *string `json:"auth_no,omitempty"`
	// ExtraParam 是解冻扩展信息。map<string, string>的json格式，目前支持如下key： unfreezeBizInfo：由芝麻消费，当前支持value如下： "bizComplete":"true"——标识本次解冻用户已履约，true表示信用单履约完结。
	ExtraParam *string `json:"extra_param,omitempty"`
	// NotifyUrl 是通知地址。
	NotifyUrl *string `json:"notify_url,omitempty"`
	// OutRequestNo 是解冻请求流水号。 如果是针对同一笔授权单不同的解冻请求，如第一次解冻1元，第二次解冻2元，则解冻请求流水号必须不重复； 如果是针对同一笔解冻请求的多次发起，则需要保证每次发起，解冻请求流水号和解冻金额都相同。
	OutRequestNo *string `json:"out_request_no,omitempty"`
	// Remark 是商户对本次解冻操作的附言描述。
	Remark *string `json:"remark,omitempty"`
}

// FundAuthOrderVoucherCreateRequest 是支付宝官方 V3 OpenAPI 模型。
type FundAuthOrderVoucherCreateRequest struct {
	// Amount 是需要冻结的金额，单位为：元（人民币），精确到小数点后两位。 取值范围：[0.01,100000000.00]。
	Amount *string `json:"amount,omitempty"`
	// BusinessParams 是业务参数，如风控参数outRiskInfo等。。
	BusinessParams *string `json:"business_params,omitempty"`
	// DepositProductMode 是免押受理台模式，使用免押产品必传该字段。根据免押不同业务模式将开通受理台区分三种模式，商家可根据调用预授权冻结接口传入的参数决定该笔免押订单选择哪种受理台模式。不同受理台模式需要传入不同参数，其中：POSTPAY 表示后付金额已知，POSTPAY_UNCERTAIN 表示后付金额未知，DEPOSIT_ONLY 表示纯免押。 具体规则参考文档：<a href="https://opendocs.alipay.com/b/08tf3t?pathHash=d67d7545">https://opendocs.alipay.com/b/08tf3t?pathHash=d67d7545</a>。
	DepositProductMode *string `json:"deposit_product_mode,omitempty"`
	// DisablePayChannels 是无特殊需要请勿传入；商户可用该参数禁用支付渠道。 传入后用户不可使用列表中的渠道进行支付，目前支持两种禁用渠道：信用卡快捷（OPTIMIZED_MOTO）、信用卡卡通（BIGAMOUNT_CREDIT_CARTOON）。与可用支付渠道不能同时传入。
	DisablePayChannels *string `json:"disable_pay_channels,omitempty"`
	// EnablePayChannels 是无特殊需要请勿传入；商户可用该参数指定支付渠道。 传入后用户仅能使用列表中的渠道进行支付，目前支持三种渠道，余额宝（MONEY_FUND）、花呗（PCREDIT_PAY）以及芝麻信用（CREDITZHIMA）。与禁用支付渠道不可同时传入。
	EnablePayChannels *string `json:"enable_pay_channels,omitempty"`
	// ExtraParam 是业务扩展参数，用于特定业务信息的传递，json格式。 1、category，信用类目，信用预授权场景必传，具体类目信息见<a href="https://opendocs.alipay.com/open/10719">https://opendocs.alipay.com/open/10719</a>； 2、serviceId，信用服务ID：信用预授权场景必传。需要商家在 <a href="https://b.alipay.com/page/zmgaia/pre-auth/index">开放平台-芝麻免押-信用服务管理</a> 创建信用服务获取，详情可查看 <a href="https://opendocs.alipay.com/open/03w0a7?pathHash=51f6b4f2&ref=api#%E7%94%B3%E8%AF%B7%E4%BF%A1%E7%94%A8%E6%9C%8D%E5%8A%A1">创建信用服务</a>。在创建过程中如果有其它疑问，可以咨询芝麻客服小二（0571-88158055 转 2）； 3、creditExtInfo，信用参数，可选，如有需要请与芝麻约定后传入，信用服务说明见<a href="https://opendocs.alipay.com/open/11157/qlsxya">https://opendocs.alipay.com/open/11157/qlsxya</a>。
	ExtraParam *string `json:"extra_param,omitempty"`
	// IdentityParams 是无特殊需要请勿传入；买家实名信息。 传入后支付宝会比对买家在支付宝端的实名信息。包含两个可选key： 1.identity_hash，买家姓名拼接身份证号后，使用SHA256摘要方式与UTF8编码后的hash值，返回的十六进制字符串，例如 “张三4566498798498498498498”对应的identity_hash应为“acc2b92ffc5ed9b472faa19748f10045c30434132784f774b00216a56b8841c6” 2.alipay_open_id，买家支付宝openId 3.alipay_user_id，买家支付宝uid。
	IdentityParams *string `json:"identity_params,omitempty"`
	// NotifyUrl 是通知地址。
	NotifyUrl *string `json:"notify_url,omitempty"`
	// OrderTitle 是订单标题。 业务订单的简单描述，如商品名称等。
	OrderTitle *string `json:"order_title,omitempty"`
	// OutOrderNo 是商户授权资金订单号。 商家自定义需保证在商户端不重复。仅支持字母、数字、下划线。。
	OutOrderNo *string `json:"out_order_no,omitempty"`
	// OutRequestNo 是商户本次资金操作的请求流水号，用于标示请求流水的唯一性。 可与out_order_no相同，仅支持字母、数字、下划线。。
	OutRequestNo *string `json:"out_request_no,omitempty"`
	// PayTimeout 是预授权订单相对超时时间，从商户请求时间开始计算。 预授权订单允许的最晚授权时间，逾期将关闭该笔订单。取值范围：1m～15d。m-分钟，h-小时，d-天。 该参数数值不接受小数点， 如 1.5h，可转换为90m。 默认为15m。。
	PayTimeout *string `json:"pay_timeout,omitempty"`
	// PayeeLogonId 是收款账户的支付宝登录号（email或手机号）。 如果传入则会校验该登录号对应的账号是否具备当前商户收款权限，如果商户希望用户能够使用花呗，则用户号(payee_user_id)和登录号(payee_logon_id)两者必须传入其一，二者无需同时传入。
	PayeeLogonId *string `json:"payee_logon_id,omitempty"`
	// PayeeUserId 是收款账户的支付宝用户号。 以2088开头的16位纯数字，如果传入则会校验该账号是否具备当前商户收款权限，如果商户希望用户能够使用花呗，则用户号(payee_user_id)和登录号(payee_logon_id)两者必须传入其一，二者无需同时传入。
	PayeeUserId *string `json:"payee_user_id,omitempty"`
	// PostPayments 是后付费项目，有付费项目时需要传入该字段。不同受理台模式需要传入不同参数，后付费项目名称和计费说明需要通过校验规则，同时计费说明将展示在开通受理台上。当受理台模式（deposit_product_mode）传入POSTPAY 时，后付费项目名称（name）、金额（amount）必传，计费说明（description）选传；当传入 POSTPAY_UNCERTAIN 时，后付费项目名称（name）、计费说明（description）必传，金额（amount）不传。 具体规则参考文档：<a href="https://opendocs.alipay.com/b/08tf3t?pathHash=d67d7545">https://opendocs.alipay.com/b/08tf3t?pathHash=d67d7545</a>。
	PostPayments *[]PostPayment `json:"post_payments,omitempty"`
	// ProductCode 是销售产品码。 当面资金预授权固定为 PRE_AUTH。
	ProductCode *string `json:"product_code,omitempty"`
	// SettleCurrency 是商户指定的结算币种。支持澳元：AUD, 新西兰元：NZD, 台币：TWD, 美元：USD, 欧元：EUR, 英镑：GBP。
	SettleCurrency *string `json:"settle_currency,omitempty"`
	// TimeoutExpress 是预授权订单相对超时时间，从商户请求时间开始计算。 预授权订单允许的最晚授权时间，逾期将关闭该笔订单。取值范围：1m～15d。m-分钟，h-小时，d-天。 该参数数值不接受小数点， 如 1.5h，可转换为90m。 默认为15m。。
	TimeoutExpress *string `json:"timeout_express,omitempty"`
	// TransCurrency 是标价币种, amount 对应的币种单位。支持澳元：AUD, 新西兰元：NZD, 台币：TWD, 美元：USD, 欧元：EUR, 英镑：GBP。
	TransCurrency *string `json:"trans_currency,omitempty"`
}
