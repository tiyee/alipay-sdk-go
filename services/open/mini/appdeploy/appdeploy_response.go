// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package appdeploy

// MiniAppDeployByAppVersionQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type MiniAppDeployByAppVersionQueryResponse struct {
	// Deploys 是发布信息。
	Deploys *[]MiniAppDeployResponse `json:"deploys,omitempty"`
}

// MiniAppDeployByDeployVersionQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type MiniAppDeployByDeployVersionQueryResponse struct {
	// Deploys 是小程序发布信息。
	Deploys *[]MiniAppDeployResponse `json:"deploys,omitempty"`
}

// MiniAppDeployResponse 是支付宝官方 V3 OpenAPI 模型。
type MiniAppDeployResponse struct {
	// AndroidClientMax 是安卓最大版本号。
	AndroidClientMax *string `json:"android_client_max,omitempty"`
	// AndroidClientMin 是安卓最低版本号。
	AndroidClientMin *string `json:"android_client_min,omitempty"`
	// AppVersion 是小程序版本号。
	AppVersion *string `json:"app_version,omitempty"`
	// BundleId 是端标识。
	BundleId *string `json:"bundle_id,omitempty"`
	// DeployVersion 是发布版本。
	DeployVersion *string `json:"deploy_version,omitempty"`
	// GmtCreate 是创建时间。
	GmtCreate *string `json:"gmt_create,omitempty"`
	// GmtModified 是修改时间。
	GmtModified *string `json:"gmt_modified,omitempty"`
	// IosClientMax 是iOS最大版本号。
	IosClientMax *string `json:"ios_client_max,omitempty"`
	// IosClientMin 是iOS最小版本号。
	IosClientMin *string `json:"ios_client_min,omitempty"`
	// MiniAppId 是小程序ID。
	MiniAppId *string `json:"mini_app_id,omitempty"`
	// PackageSize 是包大小。
	PackageSize *string `json:"package_size,omitempty"`
	// Status 是发布状态 INVALID("INVALID", "非法"), INIT("INIT", "初始化状态"), PRE_ONLINE("PRE_ONLINE", "预发"); GRAY("GRAY", "灰度"), ONLINE("ONLINE", "上架"), ABANDON("ABANDON", "废弃"),。
	Status *string `json:"status,omitempty"`
}
