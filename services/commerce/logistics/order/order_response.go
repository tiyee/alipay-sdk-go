package order

// LogisticsOrderInstantDeliveryCancelResponse 是支付宝官方 V3 OpenAPI 模型。
type LogisticsOrderInstantDeliveryCancelResponse struct {
	// CancelFee 是取消金额
	CancelFee *string `json:"cancel_fee,omitempty"`
}

// LogisticsOrderInstantDeliveryCreateResponse 是支付宝官方 V3 OpenAPI 模型。
type LogisticsOrderInstantDeliveryCreateResponse struct {
	// CouponFee 是优惠券费用，单位元
	CouponFee *string `json:"coupon_fee,omitempty"`
	// DeliverFee 是运费，单位元
	DeliverFee *string `json:"deliver_fee,omitempty"`
	// DispatchDuration 是预计骑手还剩多久接单，单位：秒
	DispatchDuration *int64 `json:"dispatch_duration,omitempty"`
	// Distance 是配送距离
	Distance *int64 `json:"distance,omitempty"`
	// Fee 是实际运费，单位元
	Fee *string `json:"fee,omitempty"`
	// FinishCode 是收货码，骑手必须输入收货码才能完成订单妥投
	FinishCode *string `json:"finish_code,omitempty"`
	// InsuranceFee 是保价费用，单位元
	InsuranceFee *string `json:"insurance_fee,omitempty"`
	// OrderNo 是支付宝订单流水号
	OrderNo *string `json:"order_no,omitempty"`
	// PayAmount 是支付金额, 实际扣减的费用以此字段为准，单位元
	PayAmount *string `json:"pay_amount,omitempty"`
	// PickupCode 是取货码, 骑手必须输入取货码才能从商家取货
	PickupCode *string `json:"pickup_code,omitempty"`
	// Status 是即时配送运单状态
	Status *string `json:"status,omitempty"`
	// WaybillNo 是即时配送运单编号
	WaybillNo *string `json:"waybill_no,omitempty"`
}

// LogisticsOrderInstantDeliveryPreCreateResponse 是支付宝官方 V3 OpenAPI 模型。
type LogisticsOrderInstantDeliveryPreCreateResponse struct {
	// Waybills 是即时配送运单列表
	Waybills *[]PreCreateWaybillIstd `json:"waybills,omitempty"`
}

// LogisticsOrderIstdCancelPreConsultResponse 是支付宝官方 V3 OpenAPI 模型。
type LogisticsOrderIstdCancelPreConsultResponse struct {
	// AllowCancel 是是否允许取消
	AllowCancel *bool `json:"allow_cancel,omitempty"`
	// CancelFee 是取消金额
	CancelFee *string `json:"cancel_fee,omitempty"`
}

// LogisticsOrderIstdRetryCreateResponse 是支付宝官方 V3 OpenAPI 模型。
type LogisticsOrderIstdRetryCreateResponse struct {
	// CouponFee 是优惠券费用
	CouponFee *string `json:"coupon_fee,omitempty"`
	// DeliverFee 是运费
	DeliverFee *string `json:"deliver_fee,omitempty"`
	// DispatchDuration 是预计骑手还剩多久接单，单位：秒
	DispatchDuration *int64 `json:"dispatch_duration,omitempty"`
	// Distance 是配送距离(单位：米)
	Distance *int64 `json:"distance,omitempty"`
	// Fee 是实际运费
	Fee *string `json:"fee,omitempty"`
	// FinishCode 是收货码，骑手必须输入收货码才能完成订单妥投
	FinishCode *string `json:"finish_code,omitempty"`
	// InsuranceFee 是保价费用
	InsuranceFee *string `json:"insurance_fee,omitempty"`
	// OrderNo 是支付宝订单流水号
	OrderNo *string `json:"order_no,omitempty"`
	// PayAmount 是支付金额, 实际扣减的费用以此字段为准
	PayAmount *string `json:"pay_amount,omitempty"`
	// PickupCode 是取货码, 骑手必须输入取货码才能从商家取货
	PickupCode *string `json:"pickup_code,omitempty"`
	// Status 是即时配送运单状态
	Status *string `json:"status,omitempty"`
	// WaybillNo 是即时配送运单编号
	WaybillNo *string `json:"waybill_no,omitempty"`
}
