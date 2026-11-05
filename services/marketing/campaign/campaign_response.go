// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package campaign

import (
	"encoding/json"
)

// CampaignCashCreateResponse 是支付宝官方 V3 OpenAPI 模型。
type CampaignCashCreateResponse struct {
	// CrowdNo 是生成的现金红包活动号
	CrowdNo *string `json:"crowd_no,omitempty"`
	// OriginCrowdNo 是原始活动号,商户排查问题时提供的活动依据
	OriginCrowdNo *string `json:"origin_crowd_no,omitempty"`
	// PayUrl 是活动创建后的付款链接，返回的是urlencode编码后的字符串。需要先进行urldecode解码，然后在浏览器中进行访问，会先进行支付宝登录引导，然后商户进行付款，付款有效期24小时。
	PayUrl *string `json:"pay_url,omitempty"`
}

// CampaignCashDetailQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type CampaignCashDetailQueryResponse struct {
	// CampStatus 是活动状态，CREATED: 已创建未打款	PAID:已打款	READY:活动已开始	PAUSE:活动已暂停	CLOSED:活动已结束	SETTLED:活动已清算
	CampStatus *string `json:"camp_status,omitempty"`
	// CouponName 是红包名称
	CouponName *string `json:"coupon_name,omitempty"`
	// CrowdNo 是活动号
	CrowdNo *string `json:"crowd_no,omitempty"`
	// EndTime 是活动结束时间
	EndTime *string `json:"end_time,omitempty"`
	// OriginCrowdNo 是原始活动号,商户排查问题时提供的活动依据
	OriginCrowdNo *string `json:"origin_crowd_no,omitempty"`
	// PrizeMsg 是活动文案,用户在账单中看到的红包描述
	PrizeMsg *string `json:"prize_msg,omitempty"`
	// PrizeType 是现金红包的发放形式, fixed为固定金额,random为随机金额
	PrizeType *string `json:"prize_type,omitempty"`
	// SendAmount 是活动已发放金额
	SendAmount *string `json:"send_amount,omitempty"`
	// StartTime 是活动开始时间
	StartTime *string `json:"start_time,omitempty"`
	// TotalAmount 是活动总金额
	TotalAmount *string `json:"total_amount,omitempty"`
	// TotalCount 是红包总个数
	TotalCount *int64 `json:"total_count,omitempty"`
	// TotalNum 是红包总个数(废弃)
	TotalNum *int64 `json:"total_num,omitempty"`
}

// CampaignCashListQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type CampaignCashListQueryResponse struct {
	// CampList 是活动列表
	CampList *[]CashCampaignInfo `json:"camp_list,omitempty"`
	// PageIndex 是分页的页码,起始从1开始
	PageIndex *string `json:"page_index,omitempty"`
	// PageSize 是分页每页大小
	PageSize *string `json:"page_size,omitempty"`
	// TotalSize 是活动总个数
	TotalSize *string `json:"total_size,omitempty"`
}

// CampaignCashStatusModifyResponse 是支付宝官方 V3 OpenAPI 模型。
type CampaignCashStatusModifyResponse = json.RawMessage

// CampaignCashTriggerResponse 是支付宝官方 V3 OpenAPI 模型。
type CampaignCashTriggerResponse struct {
	// BizNo 是支付宝业务号,发奖成功时返回,调用者可提供此字符串进行问题排查与核对等
	BizNo *string `json:"biz_no,omitempty"`
	// CouponName 是红包名称,创建活动时设置，用于账单列表和详情、红包列表和详情的展示
	CouponName *string `json:"coupon_name,omitempty"`
	// ErrorMsg 是用户领取失败的错误信息描述
	ErrorMsg *string `json:"error_msg,omitempty"`
	// MerchantLogo 是商户头像logo的图片url地址，用于账单和红包列表、详情的展示
	MerchantLogo *string `json:"merchant_logo,omitempty"`
	// OutBizNo 是外部业务号,回填请求中的out_biz_no,请求参数中传了out_biz_no就会回传回去，如果不传就回传默认的"default_out_biz_no"，请求者可用于日志记录与核对等
	OutBizNo *string `json:"out_biz_no,omitempty"`
	// PartnerId 是发送红包商户签约pid，发奖成功时非空
	PartnerId *string `json:"partner_id,omitempty"`
	// PrizeAmount 是现金红包金额，发奖成功时非空，单位为元，保留两位小数
	PrizeAmount *string `json:"prize_amount,omitempty"`
	// PrizeMsg 是活动文案,用于账单的备注展示、红包详情页的文案展示
	PrizeMsg *string `json:"prize_msg,omitempty"`
	// RepeatTriggerFlag 是用户是否重复领取，如果重复领取，返回的是之前的中奖结果，如果是首次领取，则走发奖流程
	RepeatTriggerFlag *string `json:"repeat_trigger_flag,omitempty"`
	// TriggerResult 是是否中奖结果状态，取值为true或false。如果为true表示发奖成功，这时返回的结果中的其他字段非空；如果为false表示发奖失败，这时返回的其他字段为空
	TriggerResult *string `json:"trigger_result,omitempty"`
}

// CampaignOrderVoucherConsultResponse 是支付宝官方 V3 OpenAPI 模型。
type CampaignOrderVoucherConsultResponse struct {
	// OptimalTotalPromoAmount 是最优券叠加总优惠金额，单位为元
	OptimalTotalPromoAmount *string `json:"optimal_total_promo_amount,omitempty"`
	// VoucherConsultList 是优惠券咨询结果列表
	VoucherConsultList *[]VoucherConsultInfo `json:"voucher_consult_list,omitempty"`
}
