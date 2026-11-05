package version

// AppVersionInfo 是支付宝官方 V3 OpenAPI 模型。
type AppVersionInfo struct {
	// AppVersion 是版本号，一般为x.y.x格式，其中x、y、z都是数字。
	AppVersion *string `json:"app_version,omitempty"`
	// BaseAudit 是准入审核结果 （提审时 audit_rule 传 BASE_PROMOTE 时有该字段）。
	BaseAudit *string `json:"base_audit,omitempty"`
	// BundleId 是小程序客户端类型，默认为支付宝端。常见支持如下客户端： com.alipay.alipaywallet：支付宝端； com.alibaba.android.rimet：DINGDING端； com.amap.app：高德端； com.alibaba.ailabs.genie.webapps：天猫精灵端； com.alipay.iot.xpaas：支付宝IoT端。 如需更多端投放，请联系业务BD。。
	BundleId *string `json:"bundle_id,omitempty"`
	// CanRelease 是是否可上架 true：可上架 false：不可上架 （version_status 为PROMOTE_AUDIT_REJECT有值） （提审时 audit_rule 传 BASE_PROMOTE 时有该字段）。
	CanRelease *string `json:"can_release,omitempty"`
	// CreateTime 是版本创建时间，精确到秒。
	CreateTime *string `json:"create_time,omitempty"`
	// PromoteAudit 是营销审核结果： PASS：通过 REJECT：驳回 （提审时 audit_rule 传 BASE_PROMOTE 时有该字段）。
	PromoteAudit *string `json:"promote_audit,omitempty"`
	// VersionDescription 是小程序版本描述，介绍此版本的主要变更和功能，5-500个字符。。
	VersionDescription *string `json:"version_description,omitempty"`
	// VersionStatus 是版本状态，可选值为：INIT: 开发中, AUDITING: 审核中, AUDIT_REJECT: 审核驳回, WAIT_RELEASE: 待上架, BASE_AUDIT_PASS: 准入不可营销, GRAY: 灰度中, RELEASE: 已上架, OFFLINE: 已下架, AUDIT_OFFLINE: 被强制下架;。
	VersionStatus *string `json:"version_status,omitempty"`
}

// MiniAppAuditReason 是支付宝官方 V3 OpenAPI 模型。
type MiniAppAuditReason struct {
	// AuditImages 是审核图片列表。
	AuditImages *[]string `json:"audit_images,omitempty"`
	// Memos 是驳回原因列表。
	Memos *[]MiniAppAuditReasonMemo `json:"memos,omitempty"`
}

// MiniAppAuditReasonMemo 是支付宝官方 V3 OpenAPI 模型。
type MiniAppAuditReasonMemo struct {
	// Memo 是驳回原因。
	Memo *string `json:"memo,omitempty"`
	// MemoImageList 是规则图片链接。
	MemoImageList *[]string `json:"memo_image_list,omitempty"`
}

// MiniAppCategoryInfo 是支付宝官方 V3 OpenAPI 模型。
type MiniAppCategoryInfo struct {
	// FirstCategoryId 是一级类目id。
	FirstCategoryId *string `json:"first_category_id,omitempty"`
	// FirstCategoryName 是一级类目名称。
	FirstCategoryName *string `json:"first_category_name,omitempty"`
	// SecondCategoryId 是二级类目id。
	SecondCategoryId *string `json:"second_category_id,omitempty"`
	// SecondCategoryName 是二级类目名称。
	SecondCategoryName *string `json:"second_category_name,omitempty"`
	// ThirdCategoryId 是三级类目id，可空。
	ThirdCategoryId *string `json:"third_category_id,omitempty"`
	// ThirdCategoryName 是三级类目名称，可空。
	ThirdCategoryName *string `json:"third_category_name,omitempty"`
}

// MiniPackageInfo 是支付宝官方 V3 OpenAPI 模型。
type MiniPackageInfo struct {
	// DocUrl 是功能包文档地址。
	DocUrl *string `json:"doc_url,omitempty"`
	// PackageDesc 是功能包描述。
	PackageDesc *string `json:"package_desc,omitempty"`
	// PackageName 是功能包名称。
	PackageName *string `json:"package_name,omitempty"`
	// PackageOpenType 是功能包开通方式。
	PackageOpenType *string `json:"package_open_type,omitempty"`
	// Status 是功能包签约状态。
	Status *string `json:"status,omitempty"`
}

// RegionInfo 是支付宝官方 V3 OpenAPI 模型。
type RegionInfo struct {
	// AreaCode 是地址所属区代码。
	AreaCode *string `json:"area_code,omitempty"`
	// AreaName 是地址所属区名称。
	AreaName *string `json:"area_name,omitempty"`
	// CityCode 是地址所属市代码。
	CityCode *string `json:"city_code,omitempty"`
	// CityName 是地址所属市名称。
	CityName *string `json:"city_name,omitempty"`
	// ProvinceCode 是地址所属省份代码。
	ProvinceCode *string `json:"province_code,omitempty"`
	// ProvinceName 是地址所属省份名称。
	ProvinceName *string `json:"province_name,omitempty"`
}
