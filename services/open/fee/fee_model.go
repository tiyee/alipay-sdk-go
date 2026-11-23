// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package fee

// FeeAdjustApply 是支付宝官方 V3 OpenAPI 模型。
type FeeAdjustApply struct {
	// Account 是服务商代申请特殊费率的商家支付宝账户。
	Account *string `json:"account,omitempty"`
	// ApplicationFee 是服务商代商家申请特殊费率的费率值（%），如0.38表示0.38%。
	ApplicationFee *string `json:"application_fee,omitempty"`
	// CertDateLimitation 是证件有效期，格式：yyyy-MM-dd，长期有效传入：9999-12-31。
	CertDateLimitation *string `json:"cert_date_limitation,omitempty"`
	// CertNo 是对应证件类型（certType）的证件号，如营业执照为营业执照号。
	CertNo *string `json:"cert_no,omitempty"`
	// CertType 是证件类型，枚举值为 餐饮服务许可证 FOOD_SERVICE_LICENSE， 食品卫生许可证 FOOD_HEALTH_LICENSE， 食品经营许可证 FOOD_BUSINESS_LICENSE， 食品流通许可证 FOOD_CIRCULATE_LICENSE， 食品生产许可证 FOOD_PRODUCTION_LICENS， 营业执照 DEFAULT_LICENSE， 烟草资质 TOBACCO_PRODUCTION_LICENSE。
	CertType *string `json:"cert_type,omitempty"`
	// CityCode 是市。
	CityCode *string `json:"city_code,omitempty"`
	// DetailAddress 是详细地址。
	DetailAddress *string `json:"detail_address,omitempty"`
	// DistrictCode 是区/县。
	DistrictCode *string `json:"district_code,omitempty"`
	// ProductCode 是服务商代商家申请调整费率的产品码。
	ProductCode *string `json:"product_code,omitempty"`
	// ProvinceCode 是省份。
	ProvinceCode *string `json:"province_code,omitempty"`
}
