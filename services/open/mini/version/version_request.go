package version

// MiniVersionAuditedCancelRequest 是支付宝官方 V3 OpenAPI 模型。
type MiniVersionAuditedCancelRequest struct {
	// AppVersion 是商家小程序版本号。可将已过审、审核被驳回的小程序版本退回到开发中的状态。。
	AppVersion *string `json:"app_version,omitempty"`
	// BundleId 是小程序客户端类型。常见支持如下客户端： com.alipay.alipaywallet：支付宝端； com.alibaba.android.rimet：DINGDING端； com.amap.app：高德端； com.alibaba.ailabs.genie.webapps：天猫精灵端； com.alipay.iot.xpaas：支付宝IoT端。 如需更多端投放，请联系业务BD。。
	BundleId *string `json:"bundle_id,omitempty"`
}

// MiniVersionDeleteRequest 是支付宝官方 V3 OpenAPI 模型。
type MiniVersionDeleteRequest struct {
	// AppVersion 是小程序版本号。只有状态为开发中的版本才能被删除。。
	AppVersion *string `json:"app_version,omitempty"`
	// BundleId 是小程序客户端类型，默认为支付宝端。常见支持如下客户端： com.alipay.alipaywallet：支付宝端； com.alibaba.android.rimet：DINGDING端； com.amap.app：高德端； com.alibaba.ailabs.genie.webapps：天猫精灵端； com.alipay.iot.xpaas：支付宝IoT端。 如需更多端投放，请联系业务BD。。
	BundleId *string `json:"bundle_id,omitempty"`
}

// MiniVersionOfflineRequest 是支付宝官方 V3 OpenAPI 模型。
type MiniVersionOfflineRequest struct {
	// AppVersion 是商家小程序已上架版本号。。
	AppVersion *string `json:"app_version,omitempty"`
	// BundleId 是小程序客户端类型，默认为支付宝端。常见支持如下客户端： com.alipay.alipaywallet：支付宝端； com.alibaba.android.rimet：DINGDING端； com.amap.app：高德端； com.alibaba.ailabs.genie.webapps：天猫精灵端； com.alipay.iot.xpaas：支付宝IoT端。 如需更多端投放，请联系业务BD。。
	BundleId *string `json:"bundle_id,omitempty"`
}

// MiniVersionOnlineRequest 是支付宝官方 V3 OpenAPI 模型。
type MiniVersionOnlineRequest struct {
	// AppVersion 是商家小程序版本号，需为审核通过状态或灰度中版本。。
	AppVersion *string `json:"app_version,omitempty"`
	// BundleId 是小程序客户端类型，默认为支付宝端。常见支持如下客户端： com.alipay.alipaywallet：支付宝端； com.alibaba.android.rimet：DINGDING端； com.amap.app：高德端； com.alibaba.ailabs.genie.webapps：天猫精灵端； com.alipay.iot.xpaas：支付宝IoT端。 如需更多端投放，请联系业务BD。。
	BundleId *string `json:"bundle_id,omitempty"`
	// Downgrade 是true - 性能优先； false - 覆盖优先（默认）；试点开放，如需使用，请咨询在线客服。
	Downgrade *bool `json:"downgrade,omitempty"`
	// PermitRegistrationLimitRelease 是如果该字段为true，在首次发版的小程序在备案未完成的情况下，允许小程序上架，但小程序上架后只可以上架可扫码进行使用，但不可被搜索、限制访问人数（50人/每天）、限制公域（无法进行商品、服务、营促销）、限制上架后90天不完成备案自动下架。 默认为false。
	PermitRegistrationLimitRelease *bool `json:"permit_registration_limit_release,omitempty"`
}

// MiniVersionRollbackRequest 是支付宝官方 V3 OpenAPI 模型。
type MiniVersionRollbackRequest struct {
	// AppVersion 是商家小程序已上架版本。下架后将自动回滚至上一（已上架）小程序版本。 例如：商家小程序有 0.01（上一上架版本）、0.02（未上架版本）、0.03（当前上架版本） 三个版本，回滚时需传入商家小程序版本 0.03 表示将 0.03 版本回滚至上一个已上架版本即此处 0.01 版本，同时 0.03 版本将自动下架，0.01 版本自动上架（无需审核）。。
	AppVersion *string `json:"app_version,omitempty"`
	// BundleId 是小程序客户端类型，默认为支付宝端。常见支持如下客户端： com.alipay.alipaywallet：支付宝端； com.alibaba.android.rimet：DINGDING端； com.amap.app：高德端； com.alibaba.ailabs.genie.webapps：天猫精灵端； com.alipay.iot.xpaas：支付宝IoT端。 如需更多端投放，请联系业务BD。。
	BundleId *string `json:"bundle_id,omitempty"`
}

// MiniVersionUploadRequest 是支付宝官方 V3 OpenAPI 模型。
type MiniVersionUploadRequest struct {
	// AppVersion 是商家小程序版本号。版本号必须满足 x.y.z, 且均为数字。要求版本号比商户之前最新的版本号高。。
	AppVersion *string `json:"app_version,omitempty"`
	// BundleId 是小程序客户端类型，默认为支付宝端。常见支持如下客户端： com.alipay.alipaywallet：支付宝端； com.alibaba.android.rimet：DINGDING端； com.amap.app：高德端； com.alibaba.ailabs.genie.webapps：天猫精灵端； com.alipay.iot.xpaas：支付宝IoT端。 如需更多端投放，请联系业务BD。。
	BundleId *string `json:"bundle_id,omitempty"`
	// Ext 是自定义参数，用于配置商家小程序的个性化内容，具体使用方法参见示例代码。 在小程序中可以通过 <a href="https://opendocs.alipay.com/mini/api/getExtConfig">my.getExtConfig()</a> 或 <a href="https://opendocs.alipay.com/mini/api/getExtConfigSync">my.getExtConfigSync()</a> 获取 ext 参数内容（建议支付宝版本 > 10.1.38，并使用 <a href="https://opendocs.alipay.com/mini/api/can-i-use">my.canIUse</a> 做兼容处理）。。
	Ext *string `json:"ext,omitempty"`
	// TemplateId 是小程序模板 APPID。。
	TemplateId *string `json:"template_id,omitempty"`
	// TemplateVersion 是小程序模板版本号。版本号必须满足 x.y.z, 且均为数字。不传默认使用最新在架（已上架）模板版本。。
	TemplateVersion *string `json:"template_version,omitempty"`
}

// MiniVersionBuildQueryRequest is the request type for alipay.open.mini.version.build.query.
type MiniVersionBuildQueryRequest struct {
	// AppVersion 是小程序版本号
	AppVersion *string `json:"app_version,omitempty" query:"app_version,omitempty"`
	// BundleId 是小程序客户端类型，默认为支付宝端。常见支持如下客户端：	com.alipay.alipaywallet：支付宝端；	com.alibaba.android.rimet：DINGDING端；	com.amap.app：高德端；	com.alibaba.ailabs.genie.webapps：天猫精灵端；	com.alipay.iot.xpaas：支付宝IoT端。	如需更多端投放，请联系业务BD。
	BundleId *string `json:"bundle_id,omitempty" query:"bundle_id,omitempty"`
}

// MiniVersionDetailQueryRequest is the request type for alipay.open.mini.version.detail.query.
type MiniVersionDetailQueryRequest struct {
	// AppVersion 是小程序版本号
	AppVersion *string `json:"app_version,omitempty" query:"app_version,omitempty"`
	// BundleId 是小程序客户端类型，默认为支付宝端。常见支持如下客户端：	com.alipay.alipaywallet：支付宝端；	com.alibaba.android.rimet：DINGDING端；	com.amap.app：高德端；	com.alibaba.ailabs.genie.webapps：天猫精灵端；	com.alipay.iot.xpaas：支付宝IoT端。	如需更多端投放，请联系业务BD。
	BundleId *string `json:"bundle_id,omitempty" query:"bundle_id,omitempty"`
}

// MiniVersionListQueryRequest is the request type for alipay.open.mini.version.list.query.
type MiniVersionListQueryRequest struct {
	// BundleId 是端参数，用于查询多端版本，不传默认为支付宝端。如需查看小程序版本信息，则为必传。常见支持如下客户端： com.alipay.alipaywallet：支付宝端； com.alibaba.android.rimet：DINGDING端； com.amap.app：高德端； com.alibaba.ailabs.genie.webapps：天猫精灵端； com.alipay.iot.xpaas：支付宝IoT端。 如需更多端投放，请联系业务BD。
	BundleId *string `json:"bundle_id,omitempty" query:"bundle_id,omitempty"`
	// VersionStatus 是版本状态列表，用英文逗号","分割，可选；不填默认不返回。如需查看小程序版本信息，则为必传。
	VersionStatus *string `json:"version_status,omitempty" query:"version_status,omitempty"`
}
