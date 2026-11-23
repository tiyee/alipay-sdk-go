// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package merchantlist

// MerchantBaseEnterOpen 是支付宝官方 V3 OpenAPI 模型。
type MerchantBaseEnterOpen struct {
	// LoginId 是登陆账号
	LoginId *string `json:"login_id,omitempty"`
	// LogoInfo 是商户logo文件流信息，该字段为文件流经base64编码后得到的字符串。
	LogoInfo *string `json:"logo_info,omitempty"`
	// MName 是商户品牌全称。
	MName *string `json:"m_name,omitempty"`
	// MShortName 是商户品牌简称
	MShortName *string `json:"m_short_name,omitempty"`
}

// SubMerchantCommonEnterOpen 是支付宝官方 V3 OpenAPI 模型。
type SubMerchantCommonEnterOpen struct {
	// ExtendFields 是扩展字段，为json字符串格式；入驻支付即开票的产品该字段必填，用户白名单，由此字段传入；支付即开票商户配置维度，如果不传默认为SMID_MODE，目前仅支持SMID_MODE和PID_MODE
	ExtendFields *string `json:"extend_fields,omitempty"`
	// ProductCode 是商户门店入驻的产品码,STANDARD_INVOICE:扫码开票，INVOICE_RETURN:企业自建-发票回传，INVOICE_EXPENSE:发票报销，PAYMENT_OPEN:支付即开票，PAY_FEE_OPEN:缴费后开票。
	ProductCode *string `json:"product_code,omitempty"`
	// SShortName 是商户门店所属的服务方简称。
	SShortName *string `json:"s_short_name,omitempty"`
}

// SubMerchantEnterOpen 是支付宝官方 V3 OpenAPI 模型。
type SubMerchantEnterOpen struct {
	// ExtendFields 是扩展字段，为json字符串格式；入驻缴费后开票的产品如果需要传入自定义开票链接，由此字段传入。
	ExtendFields *string `json:"extend_fields,omitempty"`
	// Industry 是商户行业，枚举值如下：subway：地铁，etc：ETC，online_car_hailing：网约车，catering：餐饮，hotel：酒店，convenience_store：便利店，oil：加油，insurance：保险，service_provider：运营商，life_payment：生活缴费，political：政务，other：其他
	Industry *string `json:"industry,omitempty"`
	// Pid 是开票商户pid，入驻支付即开票场景的时候，非登录账号入驻模式，则该字段必传
	Pid *string `json:"pid,omitempty"`
	// RegisterNo 是商户门店税号。
	RegisterNo *string `json:"register_no,omitempty"`
	// SubMName 是商户门店全称。
	SubMName *string `json:"sub_m_name,omitempty"`
	// SubMShortName 是商户门店简称
	SubMShortName *string `json:"sub_m_short_name,omitempty"`
}
