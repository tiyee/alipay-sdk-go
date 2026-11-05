package audit

import (
	"github.com/tiyee/alipay-sdk-go/services/open/mini/version"
)

// MiniVersionAuditApply 是支付宝官方 V3 OpenAPI 模型。
type MiniVersionAuditApply struct {
	// AppCategoryIds 是小程序类目，格式为 第一个一级类目_第一个二级类目;第二个一级类目_第二个二级类目，详细类目可以通过 <a href="https://docs.open.alipay.com/api_49/alipay.open.mini.category.query">https://docs.open.alipay.com/api_49/alipay.open.mini.category.query</a>接口查询，如果不填默认采用当前小程序应用类目。使用默认应用类目后不需要再次上传营业执照号、营业执照名、营业执照截图、营业执照有效期。。
	AppCategoryIds *string `json:"app_category_ids,omitempty"`
	// AppDesc 是小程序描述，长度限制 20~400 个字符（中文字符按2个字符计算）。。
	AppDesc *string `json:"app_desc,omitempty"`
	// AppEnglishName 是小程序应用英文名称，长度限制 3～30 个字符（中文字符按2个字符计算）；仅支持英文和数字。。
	AppEnglishName *string `json:"app_english_name,omitempty"`
	// AppName 是小程序名称，长度限制 3~40 个字符（中文字符按2个字符计算），仅支持包含中文、数字、英文、下划线、+、-。 如果不填默认采用当前小程序应用名称，如无默认值则必填，不可为空。。
	AppName *string `json:"app_name,omitempty"`
	// AppSlogan 是小程序简介，一句话描述小程序功能，长度限制 10~32个字符（中文字符按2个字符计算）。。
	AppSlogan *string `json:"app_slogan,omitempty"`
	// AppVersion 是小程序版本号，请选择开发版本执行提交审核操作。。
	AppVersion *string `json:"app_version,omitempty"`
	// AuditRule 是审核类型，建议设置成“BASE_PROMOTE”，小程序将会更容易上架。将小程序发布审核拆分成“准入”、“营销”。原先小程序需要两个环节都审核通过才能上架，现在小程序只需要准入通过即可上架，但是无法在营销场景（支付宝泛搜、支付宝首页推荐等）展现给用户，小程序可通过精搜或者扫码等方式使用。如果小程序需要进行营销，可发布新版本，当准入&营销都通过后，小程序也可被营销。点击查看支付宝小程序营销规范：<a href="https://opendocs.alipay.com/b/03al2q">https://opendocs.alipay.com/b/03al2q</a>。
	AuditRule *string `json:"audit_rule,omitempty"`
	// AutoOnline 是审核通过后是否自动上架， 自动上架：true；不自动上架（默认）：false。 如果小程序上架前无需特殊处理，建议设置成true，平台将会在小程序审核通过时自动将其上架。 注意：若审核通过时，小程序仍未完成备案，则将会进行自动受限发布上架（发布上线后小程序将无法被搜索，限制小程序用户数为 50 人/天，无法进行相关业务推广，只可通过扫码使用，完成备案即可解除限制，上架后90天内不未完成备案，小程序将无法打开。）若审核通过时小程序也完成备案，则会正常上架。。
	AutoOnline *string `json:"auto_online,omitempty"`
	// BundleId 是小程序投放的端参数。例如投放到支付宝钱包是支付宝端。默认支付宝端。支持： com.alipay.alipaywallet:支付宝端； com.alipay.iot.xpaas：支付宝IoT端。。
	BundleId *string `json:"bundle_id,omitempty"`
	// LicenseName 是营业执照名称，参照<a href="https://opendocs.alipay.com/mini/operation/material">https://opendocs.alipay.com/mini/operation/material</a>,要求填写营业执照信息。。
	LicenseName *string `json:"license_name,omitempty"`
	// LicenseNo 是营业执照证件号，部分小程序类目需要提交，可通过<a href="https://opendocs.alipay.com/mini/03l8c8 ">https://opendocs.alipay.com/mini/03l8c8</a>查询类目是否需要上传营业执照信息。参照<a href="https://opendocs.alipay.com/mini/operation/material ">https://opendocs.alipay.com/mini/operation/material </a>要求填写营业执照信息。。
	LicenseNo *string `json:"license_no,omitempty"`
	// LicenseValidDate 是营业执照有效期，格式为 yyyy-MM-dd，9999-12-31表示长期有效。。
	LicenseValidDate *string `json:"license_valid_date,omitempty"`
	// Memo 是小程序备注，小程序备注最多500字符（中文字符按2个字符计算）。。
	Memo *string `json:"memo,omitempty"`
	// MiniCategoryIds 是新小程序前台类目，格式为 第一个一级类目_第一个二级类目;第二个一级类目_第二个二级类目_第二个三级类目，详细类目可以通过 <a href="https://docs.open.alipay.com/api_49/alipay.open.mini.category.query">https://docs.open.alipay.com/api_49/alipay.open.mini.category.query</a> 接口查询mini_category_list。。
	MiniCategoryIds *string `json:"mini_category_ids,omitempty"`
	// RegionType 是小程序服务区域类型。
	RegionType *string `json:"region_type,omitempty"`
	// ServiceEmail 是小程序客服邮箱，如果不填默认采用当前小程序的应用客服邮箱，小程序客服电话和邮箱至少输入一个。如果前期已经设置过该信息，本次可不填，平台将会为你默认上传该信息。。
	ServiceEmail *string `json:"service_email,omitempty"`
	// ServicePhone 是小程序客服电话，长度限制5~30个字符，仅支持包含数字和-。如果不填默认采用当前小程序的应用客服电话，小程序客服电话和邮箱至少输入一个。如果前期已经设置过该信息，本次可不填，平台将会为你默认上传该信息。。
	ServicePhone *string `json:"service_phone,omitempty"`
	// ServiceRegionInfo 是省市区信息。当region_type为LOCATION或传入city_code时，province_code不能为空；填写area_code时，province_code和city_code不能为空。只填province_code则全选该省；填写province_code和city_code则全选该市，以此类推。省市区code参见<a href="https://gw.alipayobjects.com/os/bmw-prod/0aab0319-13de-42b9-85cf-13877a5f78ed.xlsx">https://gw.alipayobjects.com/os/bmw-prod/0aab0319-13de-42b9-85cf-13877a5f78ed.xlsx</a>。
	ServiceRegionInfo *[]version.RegionInfo `json:"service_region_info,omitempty"`
	// SpeedUp 是如果有审核加急权益，是否使用审核加急权益加速审核： 加速（默认）：true 不加速：false。
	SpeedUp *string `json:"speed_up,omitempty"`
	// TestAccout 是测试账号，是否需要填写请参见<a href="https://opendocs.alipay.com/mini/operation/standard/case/akxg6r#3.%20%E6%B5%8B%E8%AF%95%E5%86%85%E5%AE%B9%E6%8F%90%E4%BA%A4%E4%B8%8D%E5%AE%8C%E6%95%B4">https://opendocs.alipay.com/mini/operation/standard/case/akxg6r#3.%20%E6%B5%8B%E8%AF%95%E5%86%85%E5%AE%B9%E6%8F%90%E4%BA%A4%E4%B8%8D%E5%AE%8C%E6%95%B4</a>。
	TestAccout *string `json:"test_accout,omitempty"`
	// TestPassword 是测试账号密码。
	TestPassword *string `json:"test_password,omitempty"`
	// VersionDesc 是小程序版本描述，介绍此版本的主要变更和功能，30-500个字符（中文字符按2个字符计算）。。
	VersionDesc *string `json:"version_desc,omitempty"`
}
