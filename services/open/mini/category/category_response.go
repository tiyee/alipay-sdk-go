// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package category

// MiniCategoryQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type MiniCategoryQueryResponse struct {
	// CategoryList 是小程序类目列表。
	CategoryList *[]MiniAppCategory `json:"category_list,omitempty"`
	// MiniCategoryList 是小程序新类目列表。小程序新类目列表，为三级类目。创建、修改、提审传入mini_category_ids，资质相关请参考<a href="https://opendocs.alipay.com/mini/operation/material ">https://opendocs.alipay.com/mini/operation/material</a>。
	MiniCategoryList *[]MiniAppCategory `json:"mini_category_list,omitempty"`
}

// MiniCategoryRequireQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type MiniCategoryRequireQueryResponse struct {
	// CategoryRequireInfoList 是类目对应资质列表。
	CategoryRequireInfoList *[]CategoryRequireInfo `json:"category_require_info_list,omitempty"`
}
