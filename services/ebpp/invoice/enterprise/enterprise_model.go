// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package enterprise

// EcShopInfo 是支付宝官方 V3 OpenAPI 模型。
type EcShopInfo struct {
	// Address 是门店地址
	Address *string `json:"address,omitempty"`
	// CityId 是城市编码
	CityId *string `json:"city_id,omitempty"`
	// CityName 是城市名称
	CityName *string `json:"city_name,omitempty"`
	// DistrictId 是区编码
	DistrictId *string `json:"district_id,omitempty"`
	// DistrictName 是区名称
	DistrictName *string `json:"district_name,omitempty"`
	// Latitude 是纬度
	Latitude *string `json:"latitude,omitempty"`
	// Longitude 是经度
	Longitude *string `json:"longitude,omitempty"`
	// MccCode1 是一级mcc编码
	MccCode1 *string `json:"mcc_code_1,omitempty"`
	// MccCode2 是二级mcc编码
	MccCode2 *string `json:"mcc_code_2,omitempty"`
	// MccName1 是一级mcc名称
	MccName1 *string `json:"mcc_name_1,omitempty"`
	// MccName2 是二级mcc名称
	MccName2 *string `json:"mcc_name_2,omitempty"`
	// OutdoorImgUrl 是门头照地址
	OutdoorImgUrl *string `json:"outdoor_img_url,omitempty"`
	// PoiId 是高德poiId
	PoiId *string `json:"poi_id,omitempty"`
	// ProvinceId 是省份编码
	ProvinceId *string `json:"province_id,omitempty"`
	// ProvinceName 是省份名称
	ProvinceName *string `json:"province_name,omitempty"`
	// ShopName 是门店名称
	ShopName *string `json:"shop_name,omitempty"`
}
