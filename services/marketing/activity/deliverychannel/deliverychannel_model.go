// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package deliverychannel

// DeliveryAgencyMerchantInfo 是支付宝官方 V3 OpenAPI 模型。
type DeliveryAgencyMerchantInfo struct {
	// BusinessType 是合作业务类型。
	BusinessType *string `json:"business_type,omitempty"`
	// MerchantId 是商户信息
	MerchantId *string `json:"merchant_id,omitempty"`
	// MerchantIdType 是商户id类型。 限制: 服务商代运营模式必传	如果merchant_id_type选择SMID。则表明当前商家券活动的归属者是该SMID所绑定的支付宝PID。 因此要求该SMID必须绑定相应的支付宝PID。
	MerchantIdType *string `json:"merchant_id_type,omitempty"`
}

// DeliveryChannelInfo 是支付宝官方 V3 OpenAPI 模型。
type DeliveryChannelInfo struct {
	// BoothCode 是展位码:boothCode表达的是具体渠道所属的展位码。例如：支付结果页PAY_RESULT
	BoothCode *string `json:"booth_code,omitempty"`
	// Channel 是可投放的渠道标识.		说明：	针对不同的boot_code，channel值各不相同.接口使用者可以认为channel是不同booth_code展位下的渠道的唯一标识。		例如：	booth_code为PAYMENT_RESULT。 channel为某个商户的pid。
	Channel *string `json:"channel,omitempty"`
	// ChannelName 是渠道的名称。		说明：	booth_code：PAYMENT_RESULT，channel_name为商户名称。
	ChannelName *string `json:"channel_name,omitempty"`
}
