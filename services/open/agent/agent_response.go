package agent

// AgentCancelResponse 是支付宝官方 V3 OpenAPI 模型。
type AgentCancelResponse struct {
}

// AgentConfirmResponse 是支付宝官方 V3 OpenAPI 模型。
type AgentConfirmResponse struct {
	// AppAuthToken 是应用授权令牌。
	AppAuthToken *string `json:"app_auth_token,omitempty"`
	// AppRefreshToken 是刷新令牌。
	AppRefreshToken *string `json:"app_refresh_token,omitempty"`
	// AuthAppId 是授权商户的appid。
	AuthAppId *string `json:"auth_app_id,omitempty"`
	// ExpiresIn 是应用授权令牌的有效时间（从接口调用时间作为起始时间），单位到秒。
	ExpiresIn *string `json:"expires_in,omitempty"`
	// OpenId 是授权商户的user_id。
	OpenId *string `json:"open_id,omitempty"`
	// OrderNo 是签约单号，如果是签约必返回。
	OrderNo *string `json:"order_no,omitempty"`
	// ReExpiresIn 是刷新令牌的有效时间（从接口调用时间作为起始时间），单位到秒。
	ReExpiresIn *string `json:"re_expires_in,omitempty"`
	// UserId 是授权商户的user_id。
	UserId *string `json:"user_id,omitempty"`
}

// AgentCreateResponse 是支付宝官方 V3 OpenAPI 模型。
type AgentCreateResponse struct {
	// BatchNo 是本次代商户操作的全局唯一事务编号，后续代商户创建小程序、代签约当面付等产品、提交事务等接口都需要传递该batch_no值，且要确认只有 init 状态的batch_no才能发起调用。。
	BatchNo *string `json:"batch_no,omitempty"`
	// BatchStatus 是ISV 代商户操作事务状态，事务状态包括： init=初始状态，本接口alipay.open.agent.create返回 init 状态，只有init状态允许进行各种业务接口调用； submit=提交状态，事务已经到达终态，调用alipay.open.agent.confirm接口可以提交init状态的事务 cancel=取消状态，事务已经到达终态，调用alipay.open.agent.cancel接口可以取消init状态的事务 timeout=超时状态，事务已经到达终态，init状态的事务，在【1个小时】后会自动超时 注意：只有 init 状态才允许进行接口调用，其它状态都是终态，不允许继续进行接口调用。。
	BatchStatus *string `json:"batch_status,omitempty"`
}

// AgentMiniCreateResponse 是支付宝官方 V3 OpenAPI 模型。
type AgentMiniCreateResponse struct {
}

// AgentOrderQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type AgentOrderQueryResponse struct {
	// AgentAppId 是代理创建的应用ID，如果有代理商户创建应用，商户确认成功后，才返回应用ID，否则不返回。。
	AgentAppId *string `json:"agent_app_id,omitempty"`
	// ConfirmUrl 是只有申请单状态在MERCHANT_CONFIRM状态下，才会返回商户确认签约链接。
	ConfirmUrl *string `json:"confirm_url,omitempty"`
	// MerchantPid 是商户pid。
	MerchantPid *string `json:"merchant_pid,omitempty"`
	// OrderNo 是签约单号。
	OrderNo *string `json:"order_no,omitempty"`
	// OrderStatus 是支付宝商户入驻申请单状态，申请单状态包括： MERCHANT_INFO_HOLD=暂存，提交事务出现业务校验异常时，会暂存申请单信息，可以调用业务接口修正参数，并重新提交 MERCHANT_AUDITING=审核中，申请信息正在人工审核中 MERCHANT_CONFIRM=待商户确认，申请信息审核通过，等待联系人确认签约或授权 MERCHANT_CONFIRM_SUCCESS=商户确认成功，商户同意签约或授权 MERCHANT_CONFIRM_TIME_OUT=商户超时未确认，如果商户受到确认信息15天内未确认，则需要重新提交申请信息 MERCHANT_APPLY_ORDER_CANCELED=审核失败或商户拒绝，申请信息审核被驳回，或者商户选择拒绝签约或授权。
	OrderStatus *string `json:"order_status,omitempty"`
	// ProductAgentStatusInfos 是申请单中每个产品的签约状态。
	ProductAgentStatusInfos *[]ProductAgentStatusInfo `json:"product_agent_status_infos,omitempty"`
	// RejectReason 是审核失败的拒绝原因，只有审核失败才会返回该值。
	RejectReason *string `json:"reject_reason,omitempty"`
	// RestrictInfos 是受限信息。
	RestrictInfos *[]SignRestrictInfo `json:"restrict_infos,omitempty"`
}
