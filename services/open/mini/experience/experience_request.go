// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package experience

// MiniExperienceCancelRequest 是支付宝官方 V3 OpenAPI 模型。
type MiniExperienceCancelRequest struct {
	// AppVersion 是小程序版本号。
	AppVersion *string `json:"app_version,omitempty"`
	// BundleId 是小程序客户端类型，默认为支付宝端。常见支持如下客户端： com.alipay.alipaywallet：支付宝端； com.alibaba.android.rimet：DINGDING端； com.amap.app：高德端； com.alibaba.ailabs.genie.webapps：天猫精灵端； com.alipay.iot.xpaas：支付宝IoT端。 如需更多端投放，请联系业务BD。。
	BundleId *string `json:"bundle_id,omitempty"`
}

// MiniExperienceCreateRequest 是支付宝官方 V3 OpenAPI 模型。
type MiniExperienceCreateRequest struct {
	// AppVersion 是商家小程序版本号。。
	AppVersion *string `json:"app_version,omitempty"`
	// BundleId 是小程序客户端类型，默认为支付宝端。常见支持如下客户端： com.alipay.alipaywallet：支付宝端； com.alibaba.android.rimet：DINGDING端； com.amap.app：高德端； com.alibaba.ailabs.genie.webapps：天猫精灵端； com.alipay.iot.xpaas：支付宝IoT端。 如需更多端投放，请联系业务BD。。
	BundleId *string `json:"bundle_id,omitempty"`
}

// MiniExperienceQueryRequest is the request type for alipay.open.mini.experience.query.
type MiniExperienceQueryRequest struct {
	// AppVersion 是商家小程序版本号。
	AppVersion *string `json:"app_version,omitempty" query:"app_version,omitempty"`
	// BundleId 是小程序客户端类型，默认为支付宝端。常见支持如下客户端：	com.alipay.alipaywallet：支付宝端；	com.alibaba.android.rimet：DINGDING端；	com.amap.app：高德端；	com.alibaba.ailabs.genie.webapps：天猫精灵端；	com.alipay.iot.xpaas：支付宝IoT端。	如需更多端投放，请联系业务BD。
	BundleId *string `json:"bundle_id,omitempty" query:"bundle_id,omitempty"`
}
