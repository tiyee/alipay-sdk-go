package account

// InstantDeliveryAccountCreateRequest 是支付宝官方 V3 OpenAPI 模型。
type InstantDeliveryAccountCreateRequest struct {
	// LogisticsCodes 是配送公司编码。
	LogisticsCodes *[]string `json:"logistics_codes,omitempty"`
	// OutBizNo 是外部业务号。
	OutBizNo *string `json:"out_biz_no,omitempty"`
}

// InstantDeliveryAccountQueryRequest is the request type for alipay.open.instantdelivery.account.query.
type InstantDeliveryAccountQueryRequest struct{}
