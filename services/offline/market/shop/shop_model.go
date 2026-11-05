package shop

// ShopCategoryConfigInfo 是支付宝官方 V3 OpenAPI 模型。
type ShopCategoryConfigInfo struct {
	// Id 是类目ID
	Id *string `json:"id,omitempty"`
	// IsLeaf 是是否是叶子节点
	IsLeaf *string `json:"is_leaf,omitempty"`
	// Level 是类目层级
	Level *string `json:"level,omitempty"`
	// Link 是类目层级路径
	Link *string `json:"link,omitempty"`
	// Nm 是类目名称
	Nm *string `json:"nm,omitempty"`
}

// ShopCategoryInfo 是支付宝官方 V3 OpenAPI 模型。
type ShopCategoryInfo struct {
	// CategoryId 是类目编号
	CategoryId *string `json:"category_id,omitempty"`
	// CategoryLevel 是类目层级,目前最多支持1、2、3三级
	CategoryLevel *string `json:"category_level,omitempty"`
	// CategoryName 是类目名称
	CategoryName *string `json:"category_name,omitempty"`
}

// ShopCommentInfo 是支付宝官方 V3 OpenAPI 模型。
type ShopCommentInfo struct {
	// AvgPopularity 是店铺30天平均人气值
	AvgPopularity *string `json:"avg_popularity,omitempty"`
	// AvgPopularityName 是avg_popularity字段展示文案
	AvgPopularityName *string `json:"avg_popularity_name,omitempty"`
	// Score 是门店评分，评分越高，门店质量越高
	Score *string `json:"score,omitempty"`
	// Star 是门店星级，0~5,越高，店铺质量越高
	Star *string `json:"star,omitempty"`
}
