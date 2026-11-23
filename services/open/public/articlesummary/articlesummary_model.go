// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package articlesummary

// ArticleSummaryAnalysisData 是支付宝官方 V3 OpenAPI 模型。
type ArticleSummaryAnalysisData struct {
	// AvgReadTime 是人均阅读时长。
	AvgReadTime *string `json:"avg_read_time,omitempty"`
	// Date 是日期。
	Date *string `json:"date,omitempty"`
	// DeliverUserCnt 是送达人数。
	DeliverUserCnt *int64 `json:"deliver_user_cnt,omitempty"`
	// ExposeUserCnt 是图文曝光人数。
	ExposeUserCnt *int64 `json:"expose_user_cnt,omitempty"`
	// PraiseUserCnt 是点赞数。
	PraiseUserCnt *int64 `json:"praise_user_cnt,omitempty"`
	// ReadUserCnt 是图文阅读人数。
	ReadUserCnt *int64 `json:"read_user_cnt,omitempty"`
	// ReplyUserCnt 是评论数。
	ReplyUserCnt *int64 `json:"reply_user_cnt,omitempty"`
	// ShareUserCnt 是分享人数。
	ShareUserCnt *int64 `json:"share_user_cnt,omitempty"`
}
