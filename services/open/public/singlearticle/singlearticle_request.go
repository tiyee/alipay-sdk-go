// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package singlearticle

// PublicSingleArticleDataBatchQueryRequest 是支付宝官方 V3 OpenAPI 模型。
type PublicSingleArticleDataBatchQueryRequest struct {
	// BeginDate 是数据开始日期，时间格式为 "yyyyMMdd" 。。
	BeginDate *string `json:"begin_date,omitempty"`
	// EndDate 是数据结束日期，时间格式为"yyyyMMdd"。查询数据开始日期/结束日期时间跨度最大30天。。
	EndDate *string `json:"end_date,omitempty"`
}
