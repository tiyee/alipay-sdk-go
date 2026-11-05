// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package category

// MiniCategoryRequireQueryRequest 是支付宝官方 V3 OpenAPI 模型。
type MiniCategoryRequireQueryRequest struct {
	// CategoryCodeList 是类目列表，不超过五个。
	CategoryCodeList *[]string `json:"category_code_list,omitempty"`
}

// MiniCategoryQueryRequest is the request type for alipay.open.mini.category.query.
type MiniCategoryQueryRequest struct {
	// IsFilter 是是否过滤小程序不可用类目。枚举值如下： true：表示过滤不可用类目。 false：表示不过滤不可用类目。
	IsFilter *string `json:"is_filter,omitempty" query:"is_filter,omitempty"`
}
