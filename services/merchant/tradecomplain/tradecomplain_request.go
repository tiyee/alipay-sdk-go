package tradecomplain

// MerchantTradeComplainBatchQueryRequest 是支付宝官方 V3 OpenAPI 模型。
type MerchantTradeComplainBatchQueryRequest struct {
	// BeginTime 是查询开始时间	时间格式：yyyy-MM-dd HH:mm:ss	若不填写，则默认当前时间30天前。	注意：begin_time和end_time时间跨度最大一年
	BeginTime *string `json:"begin_time,omitempty"`
	// EndTime 是查询结束时间	时间格式：yyyy-MM-dd HH:mm:ss	若不填写，则默认当前时间。	注意：begin_time和end_time时间跨度最大一年
	EndTime *string `json:"end_time,omitempty"`
	// PageNum 是当前页，从1开始
	PageNum *int64 `json:"page_num,omitempty"`
	// PageSize 是每页条数，默认10条，最多支持20条
	PageSize *int64 `json:"page_size,omitempty"`
	// Status 是状态如下	待处理：MERCHANT_PROCESSING 	已处理：MERCHANT_FEEDBACKED 	投诉完结：FINISHED 	投诉关闭：CANCELLED 	客服处理中：PLATFORM_PROCESSING 	客服处理完结：PLATFORM_FINISH 	投诉关闭：CLOSED
	Status *string `json:"status,omitempty"`
	// TargetInfos 是应用信息
	TargetInfos *[]TargetInfo `json:"target_infos,omitempty"`
}

// MerchantTradeComplainFeedbackSubmitRequest 是支付宝官方 V3 OpenAPI 模型。
type MerchantTradeComplainFeedbackSubmitRequest struct {
	// ComplainEventId 是支付宝侧投诉单号
	ComplainEventId *string `json:"complain_event_id,omitempty"`
	// FeedbackCode 是小程序交易反馈类目ID：	02:通过其他方式退款;	03:已发货;	04:其他;	05:已完成售后服务;	06:非我方责任范围;		安心付/先享后付反馈类目ID：	194:反馈处理结果;
	FeedbackCode *string `json:"feedback_code,omitempty"`
	// FeedbackContent 是反馈内容，字数不超过200个字
	FeedbackContent *string `json:"feedback_content,omitempty"`
	// FeedbackImages 是商家处理投诉时反馈凭证的图片id，多个逗号隔开（图片id可以通过"商户上传处理图片"接口获取）
	FeedbackImages *string `json:"feedback_images,omitempty"`
	// Operator 是处理投诉人，字数不超过6个字
	Operator *string `json:"operator,omitempty"`
}

// MerchantTradeComplainReplySubmitRequest 是支付宝官方 V3 OpenAPI 模型。
type MerchantTradeComplainReplySubmitRequest struct {
	// ComplainEventId 是支付宝侧投诉单号
	ComplainEventId *string `json:"complain_event_id,omitempty"`
	// ReplyContent 是回复留言内容
	ReplyContent *string `json:"reply_content,omitempty"`
	// ReplyImages 是商家处理问题反馈时的回复凭证的图片id，多个逗号隔开（图片id可以通过"商户上传处理图片"接口获取）
	ReplyImages *string `json:"reply_images,omitempty"`
}

// MerchantTradeComplainSupplementSubmitRequest 是支付宝官方 V3 OpenAPI 模型。
type MerchantTradeComplainSupplementSubmitRequest struct {
	// ComplainEventId 是支付宝侧投诉单号
	ComplainEventId *string `json:"complain_event_id,omitempty"`
	// SupplementContent 是回复内容，最多不超过200个字
	SupplementContent *string `json:"supplement_content,omitempty"`
	// SupplementImages 是商家补充凭证时的图片id，多个逗号隔开（图片id可以通过"商户上传处理图片"接口获取）
	SupplementImages *string `json:"supplement_images,omitempty"`
}

// MerchantTradeComplainQueryRequest is the query request model for alipay.merchant.tradecomplain.query.
type MerchantTradeComplainQueryRequest struct {
	// ComplainEventId 是支付宝侧投诉单号
	ComplainEventId string `json:"-" path:"complain_event_id"`
}
