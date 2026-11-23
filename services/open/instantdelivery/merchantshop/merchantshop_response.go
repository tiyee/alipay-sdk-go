package merchantshop

// InstantDeliveryMerchantShopBatchQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type InstantDeliveryMerchantShopBatchQueryResponse struct {
	// PageNo 是当前页码。
	PageNo *int64 `json:"page_no,omitempty"`
	// PageSize 是分页数量, 最大50。。
	PageSize *int64 `json:"page_size,omitempty"`
	// Records 是门店列表数据。。
	Records *[]MerchantShopDTO `json:"records,omitempty"`
	// TotalNo 是总数量。
	TotalNo *int64 `json:"total_no,omitempty"`
}

// InstantDeliveryMerchantShopCreateResponse 是支付宝官方 V3 OpenAPI 模型。
type InstantDeliveryMerchantShopCreateResponse struct {
	// LogisticsShopStatus 是门店创建返回的结果。。
	LogisticsShopStatus *[]LogisticsShopStatusDTO `json:"logistics_shop_status,omitempty"`
}

// InstantDeliveryMerchantShopModifyResponse 是支付宝官方 V3 OpenAPI 模型。
type InstantDeliveryMerchantShopModifyResponse struct {
	// LogisticsShopStatus 是门店创建返回的结果。
	LogisticsShopStatus *[]LogisticsShopStatusDTO `json:"logistics_shop_status,omitempty"`
}

// InstantDeliveryMerchantShopQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type InstantDeliveryMerchantShopQueryResponse struct {
	// ContactName 是联系人姓名。
	ContactName *string `json:"contact_name,omitempty"`
	// DetailAddress 是地址。商户详细经营地址或人员所在地点。。
	DetailAddress *string `json:"detail_address,omitempty"`
	// EnterpriseCity 是城市编码。请按照https://gw.alipayobjects.com/os/basement_prod/253c4dcb-b8a4-4a1e-8be2-79e191a9b6db.xlsx 表格中内容填写。 （参考资料： http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/）。
	EnterpriseCity *string `json:"enterprise_city,omitempty"`
	// EnterpriseDistrict 是区县编码。请按照https://gw.alipayobjects.com/os/basement_prod/253c4dcb-b8a4-4a1e-8be2-79e191a9b6db.xlsx 表格中内容填写。 （参考资料： http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/）。
	EnterpriseDistrict *string `json:"enterprise_district,omitempty"`
	// EnterpriseProvince 是省份编码。请按照https://gw.alipayobjects.com/os/basement_prod/253c4dcb-b8a4-4a1e-8be2-79e191a9b6db.xlsx 表格中内容填写。 （参考资料： http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/）。
	EnterpriseProvince *string `json:"enterprise_province,omitempty"`
	// Latitude 是纬度，浮点型,小数点后最多保留6位 如需要录入经纬度，请以高德坐标系为准，录入时请确保经纬度参数准确。高德经纬度查询：http://lbs.amap.com/console/show/picker。
	Latitude *string `json:"latitude,omitempty"`
	// LogisticsShopStatus 是在配送公司的门店状态。。
	LogisticsShopStatus *[]LogisticsShopStatusDTO `json:"logistics_shop_status,omitempty"`
	// Longitude 是经度，浮点型, 小数点后最多保留6位。 如需要录入经纬度，请以高德坐标系为准，录入时请确保经纬度参数准确。高德经纬度查询：http://lbs.amap.com/console/show/picker。
	Longitude *string `json:"longitude,omitempty"`
	// Phone 是联系人电话/手机号。。
	Phone *string `json:"phone,omitempty"`
	// Poiid 是高德poiid。
	Poiid *string `json:"poiid,omitempty"`
	// ShopCategory 是店铺类目，取值参见文件https://mif-pub.alipayobjects.com/ShopCategory.xlsx 中的三级门店类目。
	ShopCategory *string `json:"shop_category,omitempty"`
	// ShopName 是门店名称，最长不超过256个字符。
	ShopName *string `json:"shop_name,omitempty"`
	// ShopNo 是商家门店编码。。
	ShopNo *string `json:"shop_no,omitempty"`
}
