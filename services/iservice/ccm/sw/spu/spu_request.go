package spu

// CcmSwSPUSyncRequest 是支付宝官方 V3 OpenAPI 模型。
type CcmSwSPUSyncRequest struct {
	// Attribute 是商品属性列表
	Attribute *[]SPUAttribute `json:"attribute,omitempty"`
	// Brand 是商品品牌
	Brand *string `json:"brand,omitempty"`
	// Category 是商品类目，用于商品库中分类
	Category *string `json:"category,omitempty"`
	// Description 是商品描述
	Description *string `json:"description,omitempty"`
	// Icon 是商品图片链接	特殊可选规则：is_delete=false的情况下必填
	Icon *string `json:"icon,omitempty"`
	// IsDelete 是默认为false，为true，表示删除知识点
	IsDelete *string `json:"is_delete,omitempty"`
	// LibraryId 是知识库id
	LibraryId *int64 `json:"library_id,omitempty"`
	// LibraryName 是知识库名称
	LibraryName *string `json:"library_name,omitempty"`
	// LinkUrl 是商品链接	特殊可选规则：is_delete=false的情况下必填
	LinkUrl *string `json:"link_url,omitempty"`
	// OriginalPrice 是商品原价
	OriginalPrice *string `json:"original_price,omitempty"`
	// Price 是商品售价	特殊可选规则：is_delete=false的情况下必填
	Price *string `json:"price,omitempty"`
	// SpuId 是商品ID
	SpuId *string `json:"spu_id,omitempty"`
	// Status 是商品状态：AVAILABLE 表示可售卖,SOLD_OUT为售罄/不可卖，默认可售卖
	Status *string `json:"status,omitempty"`
	// Title 是商品名称	特殊可选规则：is_delete=false的情况下必填
	Title *string `json:"title,omitempty"`
}
