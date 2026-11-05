package innerversion

// MiniInnerVersionBackDevPublishRequest 是支付宝官方 V3 OpenAPI 模型。
type MiniInnerVersionBackDevPublishRequest struct {
	// AppOrigin 是业务来源。
	AppOrigin *string `json:"app_origin,omitempty"`
	// AppVersion 是小程序版本。
	AppVersion *string `json:"app_version,omitempty"`
	// BundleId 是端信息。
	BundleId *string `json:"bundle_id,omitempty"`
	// MiniAppId 是小程序ID，仅特殊场景使用，普通业务方无需关注该参数。
	MiniAppId *string `json:"mini_app_id,omitempty"`
	// Pid 是小程序所属PID。
	Pid *string `json:"pid,omitempty"`
}

// MiniInnerVersionConditionBatchQueryRequest 是支付宝官方 V3 OpenAPI 模型。
type MiniInnerVersionConditionBatchQueryRequest struct {
	// AppOrigin 是业务来源。
	AppOrigin *string `json:"app_origin,omitempty"`
	// BundleId 是com.alipay.alipaywallet:支付宝，com.amap.app:高德。
	BundleId *string `json:"bundle_id,omitempty"`
	// MiniAppId 是小程序ID，仅特殊场景使用，普通业务方无需关注该参数。。
	MiniAppId *string `json:"mini_app_id,omitempty"`
	// PageNum 是页码。
	PageNum *string `json:"page_num,omitempty"`
	// PageSize 是分页数。
	PageSize *string `json:"page_size,omitempty"`
	// VersionStatus 是INIT:开发中，RELEASE:上架，OFFLINE:下架，AUDITING:审核中，AUDIT_REJECT:审核驳回，WAIT_RELEASE:待上架，GRAY:灰度中，以,号隔开。
	VersionStatus *string `json:"version_status,omitempty"`
}

// MiniInnerVersionContentRollbackRequest 是支付宝官方 V3 OpenAPI 模型。
type MiniInnerVersionContentRollbackRequest struct {
	// AppOrigin 是业务场景来源。
	AppOrigin *string `json:"app_origin,omitempty"`
	// AppVersion 是小程序当前上架版本号。
	AppVersion *string `json:"app_version,omitempty"`
	// BundleId 是端。
	BundleId *string `json:"bundle_id,omitempty"`
	// MiniAppId 是小程序ID，仅特殊场景使用，普通业务方无需关注该参数。。
	MiniAppId *string `json:"mini_app_id,omitempty"`
	// Pid 是小程序所属主体。
	Pid *string `json:"pid,omitempty"`
}

// MiniInnerVersionOfflinePublishRequest 是支付宝官方 V3 OpenAPI 模型。
type MiniInnerVersionOfflinePublishRequest struct {
	// AppOrigin 是业务来源。
	AppOrigin *string `json:"app_origin,omitempty"`
	// AppVersion 是版本号。
	AppVersion *string `json:"app_version,omitempty"`
	// BundleId 是端信息。
	BundleId *string `json:"bundle_id,omitempty"`
	// MiniAppId 是小程序ID，仅特殊场景使用，普通业务方无需关注该参数。
	MiniAppId *string `json:"mini_app_id,omitempty"`
	// Pid 是小程序所属PID。
	Pid *string `json:"pid,omitempty"`
}

// MiniInnerVersionOnlineRequest 是支付宝官方 V3 OpenAPI 模型。
type MiniInnerVersionOnlineRequest struct {
	// AppOfflineVersion 是下架版本。
	AppOfflineVersion *string `json:"app_offline_version,omitempty"`
	// AppOnlineVersion 是上架版本。
	AppOnlineVersion *string `json:"app_online_version,omitempty"`
	// AppOrigin 是业务来源。
	AppOrigin *string `json:"app_origin,omitempty"`
	// BundleId 是一个端的标识，用于区分不同的客户端，每接入一个客户端，都需要向小程序应用中心申请bundelId入驻。
	BundleId *string `json:"bundle_id,omitempty"`
	// InstCode 是租户code，alipay or taobao。
	InstCode *string `json:"inst_code,omitempty"`
	// MiniAppId 是小程序ID。
	MiniAppId *string `json:"mini_app_id,omitempty"`
}

// MiniInnerVersionOnlinePublishRequest 是支付宝官方 V3 OpenAPI 模型。
type MiniInnerVersionOnlinePublishRequest struct {
	// AppOrigin 是业务来源。
	AppOrigin *string `json:"app_origin,omitempty"`
	// AppVersion 是小程序版本。
	AppVersion *string `json:"app_version,omitempty"`
	// BundleId 是端信息。
	BundleId *string `json:"bundle_id,omitempty"`
	// Downgrade 是否服务降级，如果降级，拉包策略如果端上有本地包，则使用本地包打开，同时异步下载新包，待下次打开则使用新包。
	Downgrade *bool `json:"downgrade,omitempty"`
	// MiniAppId 是小程序ID，仅特殊场景使用，普通业务方无需关注该参数。
	MiniAppId *string `json:"mini_app_id,omitempty"`
	// Pid 是小程序所属PID。
	Pid *string `json:"pid,omitempty"`
	// ReleaseType 是normal默认普通发布，high 高保发布。
	ReleaseType *string `json:"release_type,omitempty"`
}

// MiniInnerVersionPackageCreateRequest 是支付宝官方 V3 OpenAPI 模型。
type MiniInnerVersionPackageCreateRequest struct {
	// AppOrigin 是业务来源。
	AppOrigin *string `json:"app_origin,omitempty"`
	// AppVersion 是小程序版本。
	AppVersion *string `json:"app_version,omitempty"`
	// BundleId 是端信息。
	BundleId *string `json:"bundle_id,omitempty"`
	// MiniAppId 是小程序ID，仅特殊场景使用，普通业务方无需关注该参数。。
	MiniAppId *string `json:"mini_app_id,omitempty"`
	// PackageType 是打包类型。预发包和覆盖率包。
	PackageType *string `json:"package_type,omitempty"`
	// Pid 是pid. 伙伴链路（app_origin='HUOBAN'）时必填。
	Pid *string `json:"pid,omitempty"`
}

// MiniInnerVersionSyncRequest 是支付宝官方 V3 OpenAPI 模型。
type MiniInnerVersionSyncRequest struct {
	// AppOrigin 是业务参数来源。
	AppOrigin *string `json:"app_origin,omitempty"`
	// AppVersion 是推送的小程序版本号。
	AppVersion *string `json:"app_version,omitempty"`
	// BundleId 是端信息。
	BundleId *string `json:"bundle_id,omitempty"`
	// DevId 是操作人ID。
	DevId *string `json:"dev_id,omitempty"`
	// MiniAppId 是小程序ID，仅特殊场景使用，普通业务方无需关注该参数。。
	MiniAppId *string `json:"mini_app_id,omitempty"`
	// SyncType 是同送方式。
	SyncType *string `json:"sync_type,omitempty"`
}

// MiniInnerVersionInfoDeleteRequest is the request type for alipay.open.mini.innerversion.info.delete.
type MiniInnerVersionInfoDeleteRequest struct {
	// MiniAppId 是小程序ID
	MiniAppId *string `json:"mini_app_id,omitempty" query:"mini_app_id,omitempty"`
	// AppOrigin 是业务场景来源
	AppOrigin *string `json:"app_origin,omitempty" query:"app_origin,omitempty"`
	// BundleId 是端信息
	BundleId *string `json:"bundle_id,omitempty" query:"bundle_id,omitempty"`
	// AppVersion 是小程序版本号
	AppVersion *string `json:"app_version,omitempty" query:"app_version,omitempty"`
	// Pid 是操作者ID，支付宝租户必填
	Pid *string `json:"pid,omitempty" query:"pid,omitempty"`
	// InstCode 是租户信息
	InstCode *string `json:"inst_code,omitempty" query:"inst_code,omitempty"`
}

// MiniInnerVersionInfoQueryRequest is the request type for alipay.open.mini.innerversion.info.query.
type MiniInnerVersionInfoQueryRequest struct {
	// BundleId 是端信息
	BundleId *string `json:"bundle_id,omitempty" query:"bundle_id,omitempty"`
	// MiniAppId 是小程序ID
	MiniAppId *string `json:"mini_app_id,omitempty" query:"mini_app_id,omitempty"`
	// AppVersion 是查询版本号
	AppVersion *string `json:"app_version,omitempty" query:"app_version,omitempty"`
	// AppOrigin 是业务来源，新接入方需要向支付宝申请专用来源，否则不予接入。https://yuque.antfin-inc.com/tinyapp-all/qddncu/bidushixiang
	AppOrigin *string `json:"app_origin,omitempty" query:"app_origin,omitempty"`
}

// MiniInnerVersionInstantiationQueryRequest is the request type for alipay.open.mini.innerversion.instantiation.query.
type MiniInnerVersionInstantiationQueryRequest struct {
	// AppOrigin 是来源类型，新接入方需要向支付宝申请专用来源，否则不予接入，申请方式请参见接入手册。
	AppOrigin *string `json:"app_origin,omitempty" query:"app_origin,omitempty"`
	// BundleId 是端id
	BundleId *string `json:"bundle_id,omitempty" query:"bundle_id,omitempty"`
	// MiniAppId 是小程序ID，特殊场景专用，普通业务方无需关注该参数。
	MiniAppId *string `json:"mini_app_id,omitempty" query:"mini_app_id,omitempty"`
	// AppVersion 是实例化小程序版本号
	AppVersion *string `json:"app_version,omitempty" query:"app_version,omitempty"`
}

// MiniInnerVersionLastQueryRequest is the request type for alipay.open.mini.innerversion.last.query.
type MiniInnerVersionLastQueryRequest struct {
	// MiniAppId 是小程序ID，仅特殊场景使用，普通业务方无需关注该参数。
	MiniAppId *string `json:"mini_app_id,omitempty" query:"mini_app_id,omitempty"`
	// BundleId 是端信息
	BundleId *string `json:"bundle_id,omitempty" query:"bundle_id,omitempty"`
	// AppOrigin 是业务来源场景
	AppOrigin *string `json:"app_origin,omitempty" query:"app_origin,omitempty"`
	// Pid 是小程序主体
	Pid *string `json:"pid,omitempty" query:"pid,omitempty"`
}

// MiniInnerVersionModelForOfflineQueryRequest is the request type for alipay.open.mini.innerversion.modelforoffline.query.
type MiniInnerVersionModelForOfflineQueryRequest struct {
	// MiniAppId 是小程序ID
	MiniAppId *string `json:"mini_app_id,omitempty" query:"mini_app_id,omitempty"`
	// AppVersion 是版本号
	AppVersion *string `json:"app_version,omitempty" query:"app_version,omitempty"`
	// BundleId 是com.alipay.alipaywallet:支付包
	BundleId *string `json:"bundle_id,omitempty" query:"bundle_id,omitempty"`
	// InstCode 是alipay:支付宝
	InstCode *string `json:"inst_code,omitempty" query:"inst_code,omitempty"`
}
