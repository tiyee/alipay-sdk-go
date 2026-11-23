package twostage

// UserTwoStageCommonUseRequest 是支付宝官方 V3 OpenAPI 模型。
type UserTwoStageCommonUseRequest struct {
	// DynamicId 是商户扫描用户的付款码值。
	DynamicId *string `json:"dynamic_id,omitempty"`
	// PayCheckStrategy 是传1表示 校验pid(pay_pid)和来支付的时候的pid一致性；	传2表示校验pid(pay_pid)和来支付的时候的pid一致性、校验scene_no和来支付时DYNAMIC_TOKEN_OUT_BIZ_NO属性一致；                       	不传值和value=2的表现一致
	PayCheckStrategy *string `json:"pay_check_strategy,omitempty"`
	// PayPid 是商家进行二阶段支付的PID信息。
	PayPid *string `json:"pay_pid,omitempty"`
	// SenceNo 是外部业务号，用于标识这笔解码请求，对同一个码的重复解码请求，sence_no必须与上一次保持一致，每次请求的sence_no必须不一样，如alipay.user.twostage.common.use接口配合alipay.trade.pay（统一收单交易支付接口）一并使用时，alipay.trade.pay接口的extend_params属性中必须设置DYNAMIC_TOKEN_OUT_BIZ_NO，且值必须与sence_no保持一致。
	SenceNo *string `json:"sence_no,omitempty"`
}

// UserTwoStageIndirectUseRequest 是支付宝官方 V3 OpenAPI 模型。
type UserTwoStageIndirectUseRequest struct {
	// DynamicId 是商户扫描用户的付款码值。18~24位，25~30开头，例如28开头的18位的数字；或人脸支付的ftoken等。
	DynamicId *string `json:"dynamic_id,omitempty"`
	// OrgPid 是进件信息中，SMID对应的银行机构的PID信息，一般为2088开头的16位数字。
	OrgPid *string `json:"org_pid,omitempty"`
	// PaySmid 是进件信息中，二级商户ID（ sub_merchant_id)信息，一般为2088开头的16位数字。
	PaySmid *string `json:"pay_smid,omitempty"`
	// SenceNo 是外部业务号，用于标识这笔解码请求，对同一个码的重复解码请求，sence_no必须与上一次保持一致，每次请求的sence_no必须不一样，如alipay.user.twostage.common.use接口配合alipay.trade.pay（统一收单交易支付接口）一并使用时，alipay.trade.pay接口的extend_params属性中必须设置DYNAMIC_TOKEN_OUT_BIZ_NO，且值必须与sence_no保持一致
	SenceNo *string `json:"sence_no,omitempty"`
	// SourcePid 是进件信息中，SMID对应渠道的PID信息，一般为2088开头的16位数字。
	SourcePid *string `json:"source_pid,omitempty"`
}
