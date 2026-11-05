// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package appdeploy

// MiniAppDeployByAppVersionQueryRequest is the request type for alipay.open.mini.appdeploy.byappversion.query.
type MiniAppDeployByAppVersionQueryRequest struct {
	// MiniAppId 是小程序ID
	MiniAppId *string `json:"mini_app_id,omitempty" query:"mini_app_id,omitempty"`
	// BundleId 是端标识
	BundleId *string `json:"bundle_id,omitempty" query:"bundle_id,omitempty"`
	// InstCode 是租户
	InstCode *string `json:"inst_code,omitempty" query:"inst_code,omitempty"`
	// AppVersion 是小程序版本
	AppVersion *string `json:"app_version,omitempty" query:"app_version,omitempty"`
}

// MiniAppDeployByDeployVersionQueryRequest is the request type for alipay.open.mini.appdeploy.bydeployversion.query.
type MiniAppDeployByDeployVersionQueryRequest struct {
	// BundleId 是客户端标识
	BundleId *string `json:"bundle_id,omitempty" query:"bundle_id,omitempty"`
	// MiniAppId 是小程序ID
	MiniAppId *string `json:"mini_app_id,omitempty" query:"mini_app_id,omitempty"`
	// InstCode 是租户
	InstCode *string `json:"inst_code,omitempty" query:"inst_code,omitempty"`
	// DeployVersion 是发布标识
	DeployVersion *string `json:"deploy_version,omitempty" query:"deploy_version,omitempty"`
}
