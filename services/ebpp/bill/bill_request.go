// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package bill

// BillAddRequest 是支付宝官方 V3 OpenAPI 模型。
type BillAddRequest struct {
	// AuthToken 在 URL query 中传递。用户授权令牌
	AuthToken *string `query:"auth_token" json:"-"`
	// BankBillNo 是外部订单号
	BankBillNo *string `json:"bank_bill_no,omitempty"`
	// BillDate 是账单的账期，格式为 yyyyMMdd。例如：202012表示2020年12月的账单。
	BillDate *string `json:"bill_date,omitempty"`
	// BillKey 是账单单据号，例如水费单号，手机号，电费号，信用卡卡号。没有唯一性要求。
	BillKey *string `json:"bill_key,omitempty"`
	// ChargeInst 是支付宝给每个出账机构指定了一个对应的英文短名称来唯一表示该收费单位。
	ChargeInst *string `json:"charge_inst,omitempty"`
	// ExtendField 是扩展属性
	ExtendField *string `json:"extend_field,omitempty"`
	// MerchantOrderNo 是输出机构的业务流水号，需要保证唯一性
	MerchantOrderNo *string `json:"merchant_order_no,omitempty"`
	// Mobile 是用户的手机号
	Mobile *string `json:"mobile,omitempty"`
	// OrderType 是支付宝订单类型。枚举支持：	*JF：公共事业缴纳。	*WUYE：物业缴费。	*HK：信用卡还款。	*TX：通讯缴费。
	OrderType *string `json:"order_type,omitempty"`
	// OwnerName 是拥有该账单的用户姓名
	OwnerName *string `json:"owner_name,omitempty"`
	// PayAmount 是缴费金额。用户支付的总金额。单位为 元（人民币）。取值范围为[0.01，100000000.00]，精确到小数点后两位。
	PayAmount *string `json:"pay_amount,omitempty"`
	// ServiceAmount 是账单的服务费。
	ServiceAmount *string `json:"service_amount,omitempty"`
	// SubOrderType 是子业务类型是业务类型的下一级概念。枚举支持：	*WATER：缴水费。	*ELECTRIC：缴电费。	*GAS：缴燃气费。	*COMMUN：缴固话宽带费。	*CATV：缴有线电视费。	*WUYE：缴物业费。	*RC：定期还车贷代扣。	*RH：定期还房贷代扣。	*RR：定期还房租代扣。	*RN：定期还网贷代扣。	*CZ：手机充值代扣。	例如：WATER表示JF下面的水费。
	SubOrderType *string `json:"sub_order_type,omitempty"`
	// TrafficLocation 是交通违章地点，sub_order_type=TRAFFIC时填写。
	TrafficLocation *string `json:"traffic_location,omitempty"`
	// TrafficRegulations 是违章行为，sub_order_type=TRAFFIC时填写。
	TrafficRegulations *string `json:"traffic_regulations,omitempty"`
}

// BillGetRequest is the query request model for alipay.ebpp.bill.get.
type BillGetRequest struct {
	// AuthToken 是用户授权令牌
	AuthToken *string `json:"auth_token,omitempty" query:"auth_token,omitempty"`
	// OrderType 是支付宝订单类型。公共事业缴纳JF,信用卡还款HK
	OrderType *string `json:"order_type,omitempty" query:"order_type,omitempty"`
	// MerchantOrderNo 是输出机构的业务流水号，需要保证唯一性。
	MerchantOrderNo *string `json:"merchant_order_no,omitempty" query:"merchant_order_no,omitempty"`
}
