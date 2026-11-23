// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package baseinfo

// MiniBaseInfoModifyResponse 是支付宝官方 V3 OpenAPI 模型。
type MiniBaseInfoModifyResponse struct {
	// ModifyStatus 是小程序基础信息修改结果，无开发中版本或仅有开发中版本修改立即生效，有上架版本会启动小程序信息审核，审核通过后信息修改生效。。
	ModifyStatus *string `json:"modify_status,omitempty"`
}

// MiniBaseInfoNameCheckResponse 是支付宝官方 V3 OpenAPI 模型。
type MiniBaseInfoNameCheckResponse struct {
	// CheckCode 是校验结果code。
	CheckCode *string `json:"check_code,omitempty"`
	// CheckMemo 是名称校验结果详细说明。
	CheckMemo *string `json:"check_memo,omitempty"`
}

// MiniBaseInfoQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type MiniBaseInfoQueryResponse struct {
	// AppDesc 是小程序应用描述。
	AppDesc *string `json:"app_desc,omitempty"`
	// AppEnglishName 是小程序应用英文名称。
	AppEnglishName *string `json:"app_english_name,omitempty"`
	// AppLogo 是小程序应用logo图标。
	AppLogo *string `json:"app_logo,omitempty"`
	// AppName 是小程序应用名称。
	AppName *string `json:"app_name,omitempty"`
	// AppSlogan 是小程序应用简介，一句话描述小程序功能。
	AppSlogan *string `json:"app_slogan,omitempty"`
	// CategoryNames 是类目名称，格式为第一个一级类目_第一个二级类目;第二个一级类目_第二个二级类目;。
	CategoryNames *string `json:"category_names,omitempty"`
	// PackageNames 是功能包名称。
	PackageNames *[]string `json:"package_names,omitempty"`
	// RegistrationLimitRelease 是false：非受限发布；true：受限发布。 受限发布：是为了能让首次发版的小程序在备案未完成的情况下，进行上架的一种折中的操作方式，小程序可以上架可扫码进行使用，但不可被搜索、限制访问人数（50人/每天）、限制公域（无法进行商品、服务、营促销）、限制上架后90天不完成备案自动下架。。
	RegistrationLimitRelease *bool `json:"registration_limit_release,omitempty"`
	// SafeDomains 是域白名单。
	SafeDomains *[]string `json:"safe_domains,omitempty"`
	// ServiceEmail 是小程序客服邮箱。
	ServiceEmail *string `json:"service_email,omitempty"`
	// ServicePhone 是小程序客服电话。
	ServicePhone *string `json:"service_phone,omitempty"`
}
