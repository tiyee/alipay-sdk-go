// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package educate

// CampusInfo 是支付宝官方 V3 OpenAPI 模型。
type CampusInfo struct {
	// Address 是地址
	Address *string `json:"address,omitempty"`
	// CampusId 是校区id
	CampusId *string `json:"campus_id,omitempty"`
	// CampusName 是校区名称
	CampusName *string `json:"campus_name,omitempty"`
	// InstName 是学校名称
	InstName *string `json:"inst_name,omitempty"`
}

// GoodsOrder 是支付宝官方 V3 OpenAPI 模型。
type GoodsOrder struct {
	// GoodsName 是商品名称
	GoodsName *string `json:"goods_name,omitempty"`
	// GoodsPictureId 是商品图片素材Id
	GoodsPictureId *string `json:"goods_picture_id,omitempty"`
}

// SchoolBaseInfo 是支付宝官方 V3 OpenAPI 模型。
type SchoolBaseInfo struct {
	// CampusInfo 是校区信息
	CampusInfo *[]CampusInfo `json:"campus_info,omitempty"`
	// CityCode 是城市编码
	CityCode *string `json:"city_code,omitempty"`
	// InstId 是学校内标
	InstId *string `json:"inst_id,omitempty"`
	// InstName 是学校名称
	InstName *string `json:"inst_name,omitempty"`
	// InstStdCode 是学校外标
	InstStdCode *string `json:"inst_std_code,omitempty"`
	// ProvinceCode 是省份编码
	ProvinceCode *string `json:"province_code,omitempty"`
}

// SchoolSimpleInfo 是支付宝官方 V3 OpenAPI 模型。
type SchoolSimpleInfo struct {
	CampusInfo *CampusInfo `json:"campus_info,omitempty"`
	// CityCode 是城市编码
	CityCode *string `json:"city_code,omitempty"`
	// InstId 是学校内标
	InstId *string `json:"inst_id,omitempty"`
	// InstName 是学校名称
	InstName *string `json:"inst_name,omitempty"`
	// InstStdCode 是学校外标
	InstStdCode *string `json:"inst_std_code,omitempty"`
	// ProvinceCode 是省份编码
	ProvinceCode *string `json:"province_code,omitempty"`
}
