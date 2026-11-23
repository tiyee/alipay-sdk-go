// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package singlearticle

// SingleArticleAnalysisData 是支付宝官方 V3 OpenAPI 模型。
type SingleArticleAnalysisData struct {
	// AvgReadTime 是人均阅读时长。
	AvgReadTime *string `json:"avg_read_time,omitempty"`
	// Date 是文章发布日期。
	Date *string `json:"date,omitempty"`
	// DeliverUserCnt 是送达人数。
	DeliverUserCnt *int64 `json:"deliver_user_cnt,omitempty"`
	// ExposeUserCnt 是曝光人数。
	ExposeUserCnt *int64 `json:"expose_user_cnt,omitempty"`
	// PraiseUserCnt 是点赞数。
	PraiseUserCnt *int64 `json:"praise_user_cnt,omitempty"`
	// ReadUserCnt 是阅读人数。
	ReadUserCnt *int64 `json:"read_user_cnt,omitempty"`
	// ReplyUserCnt 是评论数。
	ReplyUserCnt *int64 `json:"reply_user_cnt,omitempty"`
	// ShareUserCnt 是分享人数。
	ShareUserCnt *int64 `json:"share_user_cnt,omitempty"`
	// Title 是文章标题。
	Title *string `json:"title,omitempty"`
}
