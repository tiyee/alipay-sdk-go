// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package bill

// BillAddResponse 是支付宝官方 V3 OpenAPI 模型。
type BillAddResponse struct {
	// AlipayOrderNo 是支付宝的业务订单号，具有唯一性。
	AlipayOrderNo *string `json:"alipay_order_no,omitempty"`
	// BankBillNo 是外部订单号，由于对账时回传给外部商户
	BankBillNo *string `json:"bank_bill_no,omitempty"`
	// BillDate 是账单的账期，格式为 yyyyMMdd。例如：202012表示2020年12月的账单。
	BillDate *string `json:"bill_date,omitempty"`
	// BillKey 是账单单据号，例如水费单号，手机号，电费号，信用卡卡号。没有唯一性要求。
	BillKey *string `json:"bill_key,omitempty"`
	// ChargeInst 是支付宝给每个出账机构指定了一个对应的英文短名称来唯一表示该收费单位。
	ChargeInst *string `json:"charge_inst,omitempty"`
	// ChargeInstName 是出账机构中文名称。
	ChargeInstName *string `json:"charge_inst_name,omitempty"`
	// ExtendField 是扩展属性，该属性值现在用于确保只有一个人可以支付成功 	用法：多个人对同一笔外部欠费单创建多个账单时，确保该值不变
	ExtendField *string `json:"extend_field,omitempty"`
	// MerchantOrderNo 是输出机构的业务流水号，需要保证唯一性。
	MerchantOrderNo *string `json:"merchant_order_no,omitempty"`
	// OrderType 是支付宝订单类型。枚举支持：	*JF：公共事业缴纳。	*WUYE：物业缴费。	*HK：信用卡还款。	*TX：通讯缴费。
	OrderType *string `json:"order_type,omitempty"`
	// OwnerName 是拥有该账单的用户姓名
	OwnerName *string `json:"owner_name,omitempty"`
	// PayAmount 是缴费金额。用户支付的总金额。单位为 元（人民币）。取值范围为[0.01，100000000.00]，精确到小数点后两位。
	PayAmount *string `json:"pay_amount,omitempty"`
	// ServiceAmount 是账单的服务费
	ServiceAmount *string `json:"service_amount,omitempty"`
	// SubOrderType 是子业务类型是业务类型的下一级概念。枚举支持：	*WATER：缴水费。	*ELECTRIC：缴电费。	*GAS：缴燃气费。	*COMMUN：缴固话宽带费。	*CATV：缴有线电视费。	*WUYE：缴物业费。	*RC：定期还车贷代扣。	*RH：定期还房贷代扣。	*RR：定期还房租代扣。	*RN：定期还网贷代扣。	*CZ：手机充值代扣。	例如：WATER表示JF下面的水费。
	SubOrderType *string `json:"sub_order_type,omitempty"`
	// TradeNo 是50新链路订单号
	TradeNo *string `json:"trade_no,omitempty"`
}

// BillGetResponse 是支付宝官方 V3 OpenAPI 模型。
type BillGetResponse struct {
	// AlipayOrderNo 是支付宝的业务订单号，具有唯一性。
	AlipayOrderNo *string `json:"alipay_order_no,omitempty"`
	// BillDate 是账单的账期，例如201203表示2012年3月的账单。
	BillDate *string `json:"bill_date,omitempty"`
	// BillKey 是账单单据号，例如水费单号，手机号，电费号，信用卡卡号。没有唯一性要求。
	BillKey *string `json:"bill_key,omitempty"`
	// ChargeInst 是支付宝给每个出账机构指定了一个对应的英文短名称来唯一表示该收费单位。
	ChargeInst *string `json:"charge_inst,omitempty"`
	// ChargeInstName 是出账机构中文名称。
	ChargeInstName *string `json:"charge_inst_name,omitempty"`
	// MerchantOrderNo 是输出机构的业务流水号，需要保证唯一性。
	MerchantOrderNo *string `json:"merchant_order_no,omitempty"`
	// OrderStatus 是账单的状态。 INIT:等待付款，SUCCESS:成功,FAILED:失败。
	OrderStatus *string `json:"order_status,omitempty"`
	// OrderType 是支付宝订单类型。公共事业缴纳JF,信用卡还款HK
	OrderType *string `json:"order_type,omitempty"`
	// OwnerName 是拥有该账单的用户姓名
	OwnerName *string `json:"owner_name,omitempty"`
	// PayAmount 是缴费金额。用户支付的总金额。单位为：RMB Yuan。取值范围为[0.01，100000000.00]，精确到小数点后两位。
	PayAmount *string `json:"pay_amount,omitempty"`
	// PayTime 是付款时间
	PayTime *string `json:"pay_time,omitempty"`
	// ServiceAmount 是账单的服务费
	ServiceAmount *string `json:"service_amount,omitempty"`
	// SubOrderType 是子业务类型是业务类型的下一级概念，例如：WATER表示JF下面的水费，ELECTRIC表示JF下面的电费，GAS表示JF下面的燃气费。
	SubOrderType *string `json:"sub_order_type,omitempty"`
	// TrafficLocation 是交通违章地点，sub_order_type=TRAFFIC时有值
	TrafficLocation *string `json:"traffic_location,omitempty"`
	// TrafficRegulations 是违章行为，sub_order_type=TRAFFIC时有值。
	TrafficRegulations *string `json:"traffic_regulations,omitempty"`
}
