package userinfo

// DeliverAddress 是支付宝官方 V3 OpenAPI 模型。
type DeliverAddress struct {
	// Address 是地址
	Address *string `json:"address,omitempty"`
	// AddressCode 是区域编码
	AddressCode *string `json:"address_code,omitempty"`
	// DefaultDeliverAddress 是是否默认收货地址
	DefaultDeliverAddress *string `json:"default_deliver_address,omitempty"`
	// DeliverArea 是收货人所在区县
	DeliverArea *string `json:"deliver_area,omitempty"`
	// DeliverCity 是收货人所在城市
	DeliverCity *string `json:"deliver_city,omitempty"`
	// DeliverFullname 是收货人全名
	DeliverFullname *string `json:"deliver_fullname,omitempty"`
	// DeliverMobile 是收货地址的联系人移动电话
	DeliverMobile *string `json:"deliver_mobile,omitempty"`
	// DeliverPhone 是收货地址的联系人固定电话
	DeliverPhone *string `json:"deliver_phone,omitempty"`
	// DeliverProvince 是收货人所在省份
	DeliverProvince *string `json:"deliver_province,omitempty"`
	// Zip 是邮政编码
	Zip *string `json:"zip,omitempty"`
}
