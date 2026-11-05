// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package deliverychannel

// ActivityDeliveryChannelQueryRequest 是支付宝官方 V3 OpenAPI 模型。
type ActivityDeliveryChannelQueryRequest struct {
	BelongMerchantInfo *DeliveryAgencyMerchantInfo `json:"belong_merchant_info,omitempty"`
	// BoothCode 是展位码。		boothCode含义：boothCode表达的是某个渠道可以投放的展位码。例如：支付结果页PAY_RESULT		枚举值：	PAY_RESULT：支付结果页。		后续新增可投放的展位后，会在文档中新增描述信息
	BoothCode *string `json:"booth_code,omitempty"`
	// PageNum 是查询的页码。		特别说明：	页码从1开始。
	PageNum *int64 `json:"page_num,omitempty"`
	// PageSize 是每页查询的个数，最大值为 100。
	PageSize *int64 `json:"page_size,omitempty"`
}
