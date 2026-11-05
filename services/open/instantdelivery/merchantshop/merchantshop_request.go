package merchantshop

// InstantDeliveryMerchantShopBatchQueryRequest 是支付宝官方 V3 OpenAPI 模型。
type InstantDeliveryMerchantShopBatchQueryRequest struct {
	// PageNo 是当前页码。
	PageNo *int64 `json:"page_no,omitempty"`
	// PageSize 是分页数量, 最大50。。
	PageSize *int64 `json:"page_size,omitempty"`
	// ShopName 是门店名称。
	ShopName *string `json:"shop_name,omitempty"`
	// ShopNo 是商家门店编码。。
	ShopNo *string `json:"shop_no,omitempty"`
}

// InstantDeliveryMerchantShopCreateRequest 是支付宝官方 V3 OpenAPI 模型。
type InstantDeliveryMerchantShopCreateRequest struct {
	// ContactName 是联系人姓名。
	ContactName *string `json:"contact_name,omitempty"`
	// DetailAddress 是地址。商户详细经营地址或人员所在地点。
	DetailAddress *string `json:"detail_address,omitempty"`
	// EnterpriseCity 是城市编码。请按照https://gw.alipayobjects.com/os/basement_prod/253c4dcb-b8a4-4a1e-8be2-79e191a9b6db.xlsx 表格中内容填写。 （参考资料： http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/）。
	EnterpriseCity *string `json:"enterprise_city,omitempty"`
	// EnterpriseDistrict 是区县编码。请按照https://gw.alipayobjects.com/os/basement_prod/253c4dcb-b8a4-4a1e-8be2-79e191a9b6db.xlsx 表格中内容填写。 （参考资料： http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/）。
	EnterpriseDistrict *string `json:"enterprise_district,omitempty"`
	// EnterpriseProvince 是省份编码。请按照https://gw.alipayobjects.com/os/basement_prod/253c4dcb-b8a4-4a1e-8be2-79e191a9b6db.xlsx 表格中内容填写。 （参考资料： http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/）。
	EnterpriseProvince *string `json:"enterprise_province,omitempty"`
	// Latitude 是纬度，浮点型,小数点后最多保留6位 如需要录入经纬度，请以高德坐标系为准，录入时请确保经纬度参数准确。高德经纬度查询：http://lbs.amap.com/console/show/picker。
	Latitude *string `json:"latitude,omitempty"`
	// Longitude 是经度，浮点型, 小数点后最多保留6位。 如需要录入经纬度，请以高德坐标系为准，录入时请确保经纬度参数准确。高德经纬度查询：http://lbs.amap.com/console/show/picker。
	Longitude *string `json:"longitude,omitempty"`
	// OutBizNo 是外部业务号。
	OutBizNo *string `json:"out_biz_no,omitempty"`
	// Phone 是联系人电话/手机号。
	Phone *string `json:"phone,omitempty"`
	// Poiid 是高德poiid。
	Poiid *string `json:"poiid,omitempty"`
	// ShopCategory 是店铺类目，取值参见文件https://mif-pub.alipayobjects.com/ShopCategory.xlsx 中的三级门店类目。
	ShopCategory *string `json:"shop_category,omitempty"`
	// ShopName 是门店名称，最长不超过256个字符。名称+地址需要全局唯一。。
	ShopName *string `json:"shop_name,omitempty"`
	// ShopNo 是商家门店编码，可自定义，但必须唯一。
	ShopNo *string `json:"shop_no,omitempty"`
}

// InstantDeliveryMerchantShopModifyRequest 是支付宝官方 V3 OpenAPI 模型。
type InstantDeliveryMerchantShopModifyRequest struct {
	// ContactName 是联系人姓名。
	ContactName *string `json:"contact_name,omitempty"`
	// LogisticsCodes 是要更新的配送公司列表，必须包含已经创建过该门店的配送公司，可以新增开通了账户的配送公司。
	LogisticsCodes *[]string `json:"logistics_codes,omitempty"`
	// OutBizNo 是外部业务号。
	OutBizNo *string `json:"out_biz_no,omitempty"`
	// Phone 是联系人电话/手机号。
	Phone *string `json:"phone,omitempty"`
	// ShopNo 是商家门店编码。
	ShopNo *string `json:"shop_no,omitempty"`
}

// InstantDeliveryMerchantShopQueryRequest is the request type for alipay.open.instantdelivery.merchantshop.query.
type InstantDeliveryMerchantShopQueryRequest struct {
	// ShopNo 是商家门店编码。
	ShopNo *string `json:"shop_no,omitempty" query:"shop_no,omitempty"`
}
