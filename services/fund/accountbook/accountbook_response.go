package accountbook

import (
	"github.com/tiyee/alipay-sdk-go/services/fund"
)

// FundAccountBookCreateResponse 是支付宝官方 V3 OpenAPI 模型。
type FundAccountBookCreateResponse struct {
	// AccountBookId 是开通的资金记账本id。
	AccountBookId *string           `json:"account_book_id,omitempty"`
	ExtCardInfo   *fund.ExtCardInfo `json:"ext_card_info,omitempty"`
}

// FundAccountBookNotifyQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type FundAccountBookNotifyQueryResponse struct {
	// AccountBookId 是记账本ID。
	AccountBookId *string `json:"account_book_id,omitempty"`
	// Code 是响应code。 成功为10000； 业务异常为40004；。
	Code *string `json:"code,omitempty"`
	// Msg 是响应描述。 成功为Success； 业务异常为Business Failed；。
	Msg *string `json:"msg,omitempty"`
	// SubscribeStatus 是订阅关系状态， VALID：有效的； INVALID：失效的；。
	SubscribeStatus *string `json:"subscribe_status,omitempty"`
}

// FundAccountBookNotifySubscribeResponse 是支付宝官方 V3 OpenAPI 模型。
type FundAccountBookNotifySubscribeResponse struct {
	// AccountBookId 是记账本ID。
	AccountBookId *string `json:"account_book_id,omitempty"`
	// Code 是响应code。 成功为10000； 业务异常为40004；。
	Code *string `json:"code,omitempty"`
	// Msg 是响应描述。 成功为Success； 业务异常为Business Failed；。
	Msg *string `json:"msg,omitempty"`
	// SubscribeStatus 是订阅关系状态， VALID：有效的； INVALID：失效的；。
	SubscribeStatus *string `json:"subscribe_status,omitempty"`
}

// FundAccountBookNotifyUnsubscribeResponse 是支付宝官方 V3 OpenAPI 模型。
type FundAccountBookNotifyUnsubscribeResponse struct {
	// AccountBookId 是记账本ID。
	AccountBookId *string `json:"account_book_id,omitempty"`
	// Code 是响应code。 成功为10000； 业务异常为40004；。
	Code *string `json:"code,omitempty"`
	// Msg 是响应描述。 成功为Success； 业务异常为Business Failed；。
	Msg *string `json:"msg,omitempty"`
	// SubscribeStatus 是订阅关系状态， VALID：有效的； INVALID：失效的；。
	SubscribeStatus *string `json:"subscribe_status,omitempty"`
}

// FundAccountBookQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type FundAccountBookQueryResponse struct {
	// AccountBookId 是记账账簿id。
	AccountBookId *string      `json:"account_book_id,omitempty"`
	AuthInfo      *AuthInfoDTO `json:"auth_info,omitempty"`
	// AvailableAmount 是可用余额。
	AvailableAmount *string           `json:"available_amount,omitempty"`
	ExtCardInfo     *fund.ExtCardInfo `json:"ext_card_info,omitempty"`
}
