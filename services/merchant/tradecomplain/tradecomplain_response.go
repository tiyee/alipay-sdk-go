package tradecomplain

import (
	"encoding/json"
)

// MerchantTradeComplainBatchQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type MerchantTradeComplainBatchQueryResponse struct {
	// PageNum 是当前页码数
	PageNum *int64 `json:"page_num,omitempty"`
	// PageSize 是每页记录数
	PageSize *int64 `json:"page_size,omitempty"`
	// TotalNum 是总条数
	TotalNum *int64 `json:"total_num,omitempty"`
	// TotalPageNum 是总页码数
	TotalPageNum *int64 `json:"total_page_num,omitempty"`
	// TradeComplainInfos 是交易纠纷工单列表信息
	TradeComplainInfos *[]TradeComplainQueryResponse `json:"trade_complain_infos,omitempty"`
}

// MerchantTradeComplainFeedbackSubmitResponse 是支付宝官方 V3 OpenAPI 模型。
type MerchantTradeComplainFeedbackSubmitResponse = json.RawMessage

// MerchantTradeComplainQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type MerchantTradeComplainQueryResponse struct {
	// ComplainEventId 是支付宝侧投诉单号
	ComplainEventId *string `json:"complain_event_id,omitempty"`
	// ComplainReason 是用户投诉原因
	ComplainReason *string `json:"complain_reason,omitempty"`
	// Content 是用户投诉内容
	Content *string `json:"content,omitempty"`
	// GmtCreate 是投诉单创建时间
	GmtCreate *string `json:"gmt_create,omitempty"`
	// GmtFinished 是投诉单完结时间
	GmtFinished *string `json:"gmt_finished,omitempty"`
	// GmtModified 是投诉单修改时间
	GmtModified *string `json:"gmt_modified,omitempty"`
	// Images 是投诉图片
	Images *[]string `json:"images,omitempty"`
	// LeafCategoryName 是用户投诉诉求
	LeafCategoryName *string `json:"leaf_category_name,omitempty"`
	// MerchantOrderNo 是商家订单号
	MerchantOrderNo *string `json:"merchant_order_no,omitempty"`
	// PhoneNo 是投诉人电话号码
	PhoneNo *string `json:"phone_no,omitempty"`
	// ReplyDetailInfos 是用户与商家之间的协商记录
	ReplyDetailInfos *[]ReplyRecordResponse `json:"reply_detail_infos,omitempty"`
	// Status 是状态
	Status *string `json:"status,omitempty"`
	// TargetId 是应用id，例如小程序id、生活号id、商家pid
	TargetId *string `json:"target_id,omitempty"`
	// TargetType 是投诉对象类型
	TargetType *string `json:"target_type,omitempty"`
	// TradeAmount 是交易金额，单位元
	TradeAmount *string `json:"trade_amount,omitempty"`
	// TradeNo 是支付宝交易号
	TradeNo *string `json:"trade_no,omitempty"`
}

// MerchantTradeComplainReplySubmitResponse 是支付宝官方 V3 OpenAPI 模型。
type MerchantTradeComplainReplySubmitResponse = json.RawMessage

// MerchantTradeComplainSupplementSubmitResponse 是支付宝官方 V3 OpenAPI 模型。
type MerchantTradeComplainSupplementSubmitResponse = json.RawMessage

// ReplyRecordResponse 是支付宝官方 V3 OpenAPI 模型。
type ReplyRecordResponse struct {
	// Content 是回复内容
	Content *string `json:"content,omitempty"`
	// GmtCreate 是回复时间
	GmtCreate *string `json:"gmt_create,omitempty"`
	// Images 是回复图片
	Images *[]string `json:"images,omitempty"`
	// ReplierName 是回复人名称
	ReplierName *string `json:"replier_name,omitempty"`
	// ReplierRole 是回复人角色
	ReplierRole *string `json:"replier_role,omitempty"`
}

// TradeComplainQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type TradeComplainQueryResponse struct {
	// ComplainEventId 是支付宝侧投诉单号
	ComplainEventId *string `json:"complain_event_id,omitempty"`
	// ComplainReason 是投诉原因
	ComplainReason *string `json:"complain_reason,omitempty"`
	// Content 是已经扣除了我的钱，为什么再次扣款？？？
	Content *string `json:"content,omitempty"`
	// GmtCreate 是投诉单创建时间
	GmtCreate *string `json:"gmt_create,omitempty"`
	// GmtFinished 是投诉单结束时间
	GmtFinished *string `json:"gmt_finished,omitempty"`
	// GmtModified 是投诉单修改时间
	GmtModified *string `json:"gmt_modified,omitempty"`
	// Images 是投诉图片
	Images *[]string `json:"images,omitempty"`
	// LeafCategoryName 是投诉诉求
	LeafCategoryName *string `json:"leaf_category_name,omitempty"`
	// MerchantOrderNo 是商家订单号
	MerchantOrderNo *string `json:"merchant_order_no,omitempty"`
	// PhoneNo 是投诉人电话号码
	PhoneNo *string `json:"phone_no,omitempty"`
	// Status 是状态	商家处理中：MERCHANT_PROCESSING	商家已反馈：MERCHANT_FEEDBACKED	投诉已完结：FINISHED	投诉已撤销：CANCELLED	平台处理中：PLATFORM_PROCESSING	平台处理完结：PLATFORM_FINISH	系统关闭：CLOSED
	Status *string `json:"status,omitempty"`
	// TargetId 是应用id，例如小程序id、生活号id、商家pid
	TargetId *string `json:"target_id,omitempty"`
	// TargetType 是应用类型	小程序为：APPID	生活号为：PUBLICID	商家为：PID
	TargetType *string `json:"target_type,omitempty"`
	// TradeNo 是支付宝交易号
	TradeNo *string `json:"trade_no,omitempty"`
}
