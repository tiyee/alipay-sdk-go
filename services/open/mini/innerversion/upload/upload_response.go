package upload

// MiniInnerVersionNoBuildUploadResponse 是支付宝官方 V3 OpenAPI 模型。
type MiniInnerVersionNoBuildUploadResponse struct {
}

// MiniInnerVersionTemplatedUploadResponse 是支付宝官方 V3 OpenAPI 模型。
type MiniInnerVersionTemplatedUploadResponse struct {
}

// MiniInnerVersionUploadResponse 是支付宝官方 V3 OpenAPI 模型。
type MiniInnerVersionUploadResponse struct {
	// BuildInfo 是构建信息。
	BuildInfo *string `json:"build_info,omitempty"`
	// BuildPackageUrl 是包地址。
	BuildPackageUrl *string `json:"build_package_url,omitempty"`
	// BuildStatus 是构建状态。
	BuildStatus *string `json:"build_status,omitempty"`
	// BuildVersion 是构建版本。
	BuildVersion *string `json:"build_version,omitempty"`
	// NeedRotation 是是否需要轮询。
	NeedRotation *string `json:"need_rotation,omitempty"`
	// NewBuildPackageUrl 是appx构建返回值。
	NewBuildPackageUrl *string `json:"new_build_package_url,omitempty"`
	// PackageId 是包管理包ID。
	PackageId *string `json:"package_id,omitempty"`
	// VersionCreated 是版本创建结果。
	VersionCreated *string `json:"version_created,omitempty"`
}

// MiniInnerVersionUploadStatusQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type MiniInnerVersionUploadStatusQueryResponse struct {
	// BuildInfo 是构建信息。
	BuildInfo *string `json:"build_info,omitempty"`
	// BuildPackageUrl 是构建好的包地址。
	BuildPackageUrl *string `json:"build_package_url,omitempty"`
	// BuildStatus 是构建状态。
	BuildStatus *string `json:"build_status,omitempty"`
	// LogUrl 是构建日志地址。
	LogUrl *string `json:"log_url,omitempty"`
	// NeedRotation 是是否需要轮询。
	NeedRotation *string `json:"need_rotation,omitempty"`
	// NewBuildPackageUrl 是构建好的appx2.0包地址。
	NewBuildPackageUrl *string `json:"new_build_package_url,omitempty"`
	// NewResultUrl 是构建的结果地址。
	NewResultUrl *string `json:"new_result_url,omitempty"`
	// ResultUrl 是构建的结果地址。
	ResultUrl *string `json:"result_url,omitempty"`
	// VersionCreated 是创建版本结果。
	VersionCreated *string `json:"version_created,omitempty"`
}
