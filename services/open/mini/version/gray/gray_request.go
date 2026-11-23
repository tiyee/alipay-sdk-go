package gray

// MiniVersionGrayCancelRequest 是支付宝官方 V3 OpenAPI 模型。
type MiniVersionGrayCancelRequest struct {
	// AppVersion 是商家灰度中的小程序版本号，结束灰度后小程序将回到审核通过状态。。
	AppVersion *string `json:"app_version,omitempty"`
	// BundleId 是小程序客户端类型，默认为支付宝端。常见支持如下客户端： com.alipay.alipaywallet：支付宝端； com.alibaba.android.rimet：DINGDING端； com.amap.app：高德端； com.alibaba.ailabs.genie.webapps：天猫精灵端； com.alipay.iot.xpaas：支付宝IoT端。 如需更多端投放，请联系业务BD。。
	BundleId *string `json:"bundle_id,omitempty"`
}

// MiniVersionGrayOnlineRequest 是支付宝官方 V3 OpenAPI 模型。
type MiniVersionGrayOnlineRequest struct {
	// AppVersion 是商家小程序版本号，需为审核通过状态的小程序版本。。
	AppVersion *string `json:"app_version,omitempty"`
	// BundleId 是小程序客户端类型，默认为支付宝端。常见支持如下客户端： com.alipay.alipaywallet：支付宝端； com.alibaba.android.rimet：DINGDING端； com.amap.app：高德端； com.alibaba.ailabs.genie.webapps：天猫精灵端； com.alipay.iot.xpaas：支付宝IoT端。 如需更多端投放，请联系业务BD。。
	BundleId *string `json:"bundle_id,omitempty"`
	// GrayStrategy 是小程序灰度策略值。
	GrayStrategy *string `json:"gray_strategy,omitempty"`
}
