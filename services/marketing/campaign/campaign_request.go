// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package campaign

// CampaignCashCreateRequest 是支付宝官方 V3 OpenAPI 模型。
type CampaignCashCreateRequest struct {
	// CouponName 是红包名称,商户在查询列表、详情看到的名字,同时也会显示在商户付款页面。
	CouponName *string `json:"coupon_name,omitempty"`
	// EndTime 是活动结束时间,必须大于活动开始时间, 基本格式:yyyy-MM-dd HH:mm:ss,活动有效时间最长为6个月，过期后需要创建新的活动。
	EndTime *string `json:"end_time,omitempty"`
	// MerchantLink 是商户打款后的跳转链接，从支付宝收银台打款成功后的跳转链接。不填时，打款后停留在支付宝支付成功页。商户实际跳转会自动添加crowdNo作为跳转参数。示例: http://www.yourhomepage.com?crowdNo=XXX
	MerchantLink *string `json:"merchant_link,omitempty"`
	// PrizeMsg 是活动文案，用户在账单、红包中看到的账单描述、红包描述。
	PrizeMsg *string `json:"prize_msg,omitempty"`
	// PrizeType 是现金红包的发放形式,。枚举支持：	*fixed：固定金额。	*random：随机金额。选择随机金额时，单个红包的金额在平均金额的0.5~1.5倍之间浮动。
	PrizeType *string `json:"prize_type,omitempty"`
	// SendFreqency 是用户在当前活动参与次数、频率限制。支持日(D)、周(W)、月(M)、终身(L)维度的限制，整个字段不填时默认值为L1（即用户终生仅能参与1次）。	注意：	 1. 终身(L)限制必选且最大值为 100。	 2. 日(D)、周(W)、月(M)限制可选且最多只能选择一个（即限制日领取次数就不能再限制月领取次数），频率设置必须小于等于终身(L)的次数。	 3. 多个配置之间使用"|"进行分隔。如 D3|L4（表示限制用户终生只能参与4次，单日只能参与3次）。	 4. 允许多次领取时，活动触发接口需要传入out_biz_no来配合。
	SendFreqency *string `json:"send_freqency,omitempty"`
	// StartTime 是活动开始时间，必须大于活动创建的时间。入参支持如下两种形式：	 1. 填入固定时间，时间格式为 yyyy-MM-dd HH:mm:ss，如 2020-12-10 22:28:30	 2. 填字符串 NowTime，表示创建即生效。
	StartTime *string `json:"start_time,omitempty"`
	// TotalMoney 是活动发放的现金总金额,最小金额1.00元,最大金额10000000.00元。每个红包的最大金额不允许超过200元,最小金额不得低于0.20元。 实际的金额限制可能会根据业务进行动态调整。
	TotalMoney *string `json:"total_money,omitempty"`
	// TotalNum 是红包发放个数，最小为 1 个，最大10000000个。	说明：不同的发放形式（即prize_type）会使得发放个数含义不同：	 1. 若prize_type 为 fixed（固定金额），则每个用户领取的红包金额为total_money除以total_num得到固定金额。	 2. 若prize_type为 random（随机金额），则每个用户领取的红包金额为total_money除以total_num得到的平均金额值的0.5~1.5倍。由于金额是随机的，在红包金额全部被领取完时，有可能total_num有所剩余、或者大于设置值的情况。
	TotalNum *string `json:"total_num,omitempty"`
}

// CampaignCashStatusModifyRequest 是支付宝官方 V3 OpenAPI 模型。
type CampaignCashStatusModifyRequest struct {
	// CampStatus 是金活动修改后的状态。支持修改为：	*PAUSE：活动暂停。	*READY：活动开始。	*CLOSED：活动结束。
	CampStatus *string `json:"camp_status,omitempty"`
	// CrowdNo 是现金活动号，通过alipay.marketing.campaign.cash.create(创建现金活动)接口创建现金活动获取。
	CrowdNo *string `json:"crowd_no,omitempty"`
}

// CampaignCashTriggerRequest 是支付宝官方 V3 OpenAPI 模型。
type CampaignCashTriggerRequest struct {
	// CrowdNo 是现金活动号，通过alipay.marketing.campaign.cash.create(创建现金活动)接口创建现金活动获取。
	CrowdNo *string `json:"crowd_no,omitempty"`
	// LoginId 是用户登录账号名，一般为邮箱或手机号。	注意：user_id与login_id二选一必填，同时传入时以user_id为准。
	LoginId *string `json:"login_id,omitempty"`
	// OpenId 是open_id
	OpenId *string `json:"open_id,omitempty"`
	// OrderPrice 是此字段如果传入金额，就忽略prize_type算法，按照传入的金额发奖。如果不传或者小于等于0，则按照活动创建时指定的prize_type为fixed或者random算法发奖
	OrderPrice *string `json:"order_price,omitempty"`
	// OutBizNo 是领取红包的外部业务号，只由可由字母、数字、下划线组成。同一个活动中不可重复，相同的外部业务号会被幂等并返回之前的结果。不填时，系统会生成一个默认固定的外部业务号。
	OutBizNo *string `json:"out_biz_no,omitempty"`
	// UserId 是用户支付宝唯一标识，2088开头。	注意：user_id与login_id二选一必填，同时传入时以user_id为准。
	UserId *string `json:"user_id,omitempty"`
}

// CampaignOrderVoucherConsultRequest 是支付宝官方 V3 OpenAPI 模型。
type CampaignOrderVoucherConsultRequest struct {
	// AuthToken 在 URL query 中传递。用户授权令牌
	AuthToken *string `query:"auth_token" json:"-"`
	// BusinessParam 是业务参数大字段，优惠咨询的控制参数，json格式；目前支持传入useBigAmountSkipOrderThold为N来控制不使用大金额跳过优惠的订单门槛检查；默认不传；
	BusinessParam *string `json:"business_param,omitempty"`
	// ItemConsultList 是商品咨询请求列表（当需要咨询单品券时必传，如果某商品不希望参与本次单品优惠咨询则不传递对应信息即可）
	ItemConsultList *[]ItemConsultRequest `json:"item_consult_list,omitempty"`
	// OrderAmount 是订单金额（如同时享受商户自有优惠请先扣除后传入），单位为元，最多2位小数
	OrderAmount *string `json:"order_amount,omitempty"`
	// SceneCode 是场景码：默认(DEFAULT)
	SceneCode *[]string `json:"scene_code,omitempty"`
	// SpecifiedAppId 是券指定的核销appid（如果配券时指定了核销范围为线上小程序及相应的appid则此处必传）
	SpecifiedAppId *string `json:"specified_app_id,omitempty"`
}

// ItemConsultRequest 是支付宝官方 V3 OpenAPI 模型。
type ItemConsultRequest struct {
	// ItemId 是商品id
	ItemId *string `json:"item_id,omitempty"`
	// Price 是商品单价，单位为元，最多2位小数
	Price *string `json:"price,omitempty"`
	// Quantity 是商品数量
	Quantity *string `json:"quantity,omitempty"`
}

// CampaignCashDetailQueryRequest is the query request model for alipay.marketing.campaign.cash.detail.query.
type CampaignCashDetailQueryRequest struct {
	// CrowdNo 是现金活动号，通过alipay.marketing.campaign.cash.create(创建现金活动)接口创建现金活动获取。
	CrowdNo *string `json:"crowd_no,omitempty" query:"crowd_no,omitempty"`
}

// CampaignCashListQueryRequest is the query request model for alipay.marketing.campaign.cash.list.query.
type CampaignCashListQueryRequest struct {
	// CampStatus 是要查询的活动状态,不填默认为 ALL 返回所有类型。枚举支持:	*ALL：所有类型的活动。	*CREATED：已创建未打款。	*PAID：已打款。	*READY：活动已开始。	*PAUSE：活动已暂停。	*CLOSED：活动已结束。	*SETTLE：活动已清算。
	CampStatus *string `json:"camp_status,omitempty" query:"camp_status,omitempty"`
	// PageSize 是分页查询时每页返回的列表大小，每页数据最大为 50。
	PageSize *string `json:"page_size,omitempty" query:"page_size,omitempty"`
	// PageIndex 是分页查询时的页码，从1开始
	PageIndex *string `json:"page_index,omitempty" query:"page_index,omitempty"`
}
