package info

// UserDeliverAddress 是支付宝官方 V3 OpenAPI 模型。
type UserDeliverAddress struct {
	// Address 是地址
	Address *string `json:"address,omitempty"`
	// AddressCode 是区域编码
	AddressCode *string `json:"address_code,omitempty"`
	// DefaultDeliverAddress 是是否默认收货地址
	DefaultDeliverAddress *string `json:"default_deliver_address,omitempty"`
	// DeliverArea 是区/县
	DeliverArea *string `json:"deliver_area,omitempty"`
	// DeliverCity 是市
	DeliverCity *string `json:"deliver_city,omitempty"`
	// DeliverFullname 是收货人全名
	DeliverFullname *string `json:"deliver_fullname,omitempty"`
	// DeliverMobile 是收货地址的联系人移动电话
	DeliverMobile *string `json:"deliver_mobile,omitempty"`
	// DeliverPhone 是收货地址的联系人固定电话
	DeliverPhone *string `json:"deliver_phone,omitempty"`
	// DeliverProvince 是省
	DeliverProvince *string `json:"deliver_province,omitempty"`
	// Zip 是邮编
	Zip *string `json:"zip,omitempty"`
}

// UserPicture 是支付宝官方 V3 OpenAPI 模型。
type UserPicture struct {
	// PictureType 是图片类型，包括身份证正反面、营业执照等
	PictureType *string `json:"picture_type,omitempty"`
	// PictureUrl 是用于调用alipay.user.certify.image.fetch接口，获取图片资源
	PictureUrl *string `json:"picture_url,omitempty"`
}
