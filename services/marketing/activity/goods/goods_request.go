// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package goods

// ActivityGoodsBatchQueryRequest is the query request model for alipay.marketing.activity.goods.batchquery.
type ActivityGoodsBatchQueryRequest struct {
	// MerchantId 是商户PID,默认为当前接口调用商户。
	MerchantId *string `json:"merchant_id,omitempty" query:"merchant_id,omitempty"`
	// ActivityId 是活动id
	ActivityId string `json:"-" path:"activity_id"`
	// GoodsUseType 是活动单品类型。
	GoodsUseType *string `json:"goods_use_type,omitempty" query:"goods_use_type,omitempty"`
	// PageNum 是分页查询页码。	限制:	必须为大于0的整数
	PageNum *int64 `json:"page_num,omitempty" query:"page_num,required"`
	// PageSize 是分页查询单页数据条数。	限制:	1.必须为大于0的整数	2.每页最大值为20
	PageSize *int64 `json:"page_size,omitempty" query:"page_size,required"`
	// ProductVersion 是版本号		枚举值:	2.0.0
	ProductVersion *string `json:"product_version,omitempty" query:"product_version,omitempty"`
	// MerchantAccessMode 是商户接入模式
	MerchantAccessMode *string `json:"merchant_access_mode,omitempty" query:"merchant_access_mode,omitempty"`
}
