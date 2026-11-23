// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package baseinfo

// MiniBaseInfoModify 是支付宝官方 V3 OpenAPI 模型。
type MiniBaseInfoModify struct {
	// AppCategoryIds 是11_12;12_13。小程序类目，格式为 第一个一级类目_第一个二级类目;第二个一级类目_第二个二级类目，详细类目可以参考<a href="https://docs.open.alipay.com/api_49/alipay.open.mini.category.query">https://docs.open.alipay.com/api_49/alipay.open.mini.category.query</a>接口查询mini_category_list 。。
	AppCategoryIds *string `json:"app_category_ids,omitempty"`
	// AppDesc 是小程序描述，长度限制 20~400 个字符。 如果前期已经设置过该信息，本次可不填，平台将会为你默认上传该信息。如果前期没有设置过该信息，则本次为必填。可通过<a href="https://opendocs.alipay.com/mini/03l21r">https://opendocs.alipay.com/mini/03l21r</a>查询当前小程序信息。
	AppDesc *string `json:"app_desc,omitempty"`
	// AppEnglishName 是小程序应用英文名称。长度限制 3~30 个字符，仅支持包含中文、数字、英文、下划线、+、-。 如果前期已经设置过该信息，本次可不填，平台将会为你默认上传该信息。如果前期没有设置过该信息，则本次为必填。可通过<a href="https://opendocs.alipay.com/mini/03l21r?pathHash=36d314be&scene=common">alipay.open.mini.baseinfo.query</a>查询当前小程序信息。
	AppEnglishName *string `json:"app_english_name,omitempty"`
	// AppName 是小程序名称。长度限制 3~30 个字符，仅支持包含中文、数字、英文、下划线、+、-。 如果前期已经设置过该信息，本次可不填，平台将会为你默认上传该信息。如果前期没有设置过该信息，则本次为必填。可通过<a href="https://opendocs.alipay.com/mini/03l21r?pathHash=36d314be&scene=common ">alipay.open.mini.baseinfo.query</a>查询当前小程序信息。
	AppName *string `json:"app_name,omitempty"`
	// AppSlogan 是小程序简介，一句话描述小程序功能，长度限制 10~32个字符。 如果前期已经设置过该信息，本次可不填，平台将会为你默认上传该信息。如果前期没有设置过该信息，则本次为必填。可通过<a href="https://opendocs.alipay.com/mini/03l21r?pathHash=36d314be&scene=common">alipay.open.mini.baseinfo.query</a>查询当前小程序信息。
	AppSlogan *string `json:"app_slogan,omitempty"`
	// MiniCategoryIds 是新小程序前台类目，格式为 第一个一级类目_第一个二级类目;第二个一级类目_第二个二级类目_第二个三级类目，详细类目可以通过 <a href="https://docs.open.alipay.com/api_49/alipay.open.mini.category.query">https://docs.open.alipay.com/api_49/alipay.open.mini.category.query</a> 接口查询mini_category_list。 如果前期已经设置过该信息，本次可不填，平台将会为你默认上传该信息。如果前期没有设置过该信息，则本次为必填。可通过<a href="https://opendocs.alipay.com/mini/03l21r">https://opendocs.alipay.com/mini/03l21r</a>查询当前小程序信息 注意：个人开发者不得使用企业类目。。
	MiniCategoryIds *string `json:"mini_category_ids,omitempty"`
	// ServiceEmail 是小程序客服邮箱，如果不填默认采用当前小程序的应用客服邮箱，小程序客服电话和邮箱至少输入一个。如果前期已经设置过该信息，本次可不填，平台将会为你默认上传该信息。 注意：2021年7月1日后，该字段将逐步灰度为可选字段，请按可选开发。。
	ServiceEmail *string `json:"service_email,omitempty"`
	// ServicePhone 是小程序客服电话，长度限制5~30个字符，仅支持包含数字和-。如果不填默认采用当前小程序的应用客服电话，小程序客服电话和邮箱至少输入一个。如果前期已经设置过该信息，本次可不填，平台将会为你默认上传该信息。 注意：2021年7月1日后，该字段将逐步灰度为必填字段，请按必填开发。。
	ServicePhone *string `json:"service_phone,omitempty"`
}
