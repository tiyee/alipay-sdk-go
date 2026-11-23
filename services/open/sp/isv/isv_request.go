package isv

// SpISVBillSyncRequest 是支付宝官方 V3 OpenAPI 模型。
type SpISVBillSyncRequest struct {
	// EndTime 是账单结束时间。
	EndTime *string `json:"end_time,omitempty"`
	// GenerateTime 是出账时间，回传出账信息时，必须传递。
	GenerateTime *string `json:"generate_time,omitempty"`
	// ItemCode 是服务优选商品编码。
	ItemCode *string `json:"item_code,omitempty"`
	// OutBizNo 是外部业务号，传isv系统生成的账单号，需要保证唯一。
	OutBizNo *string `json:"out_biz_no,omitempty"`
	// PayTime 是账单支付时间，回传支付数时，必须传递。
	PayTime *string `json:"pay_time,omitempty"`
	// Price 是账单金额，这里填写的是整数，单位为分，比如1元，那么输入100。
	Price *string `json:"price,omitempty"`
	// PromotorPid 是推广服务商(S2)pid。
	PromotorPid *string `json:"promotor_pid,omitempty"`
	// StartTime 是账单开始时间。
	StartTime *string `json:"start_time,omitempty"`
	// Status 是GENERATE_BILL代表”已出账“，PAID_BILL代表”已支付“，选择这两者之一。
	Status *string `json:"status,omitempty"`
	// SubPromotorPid 是S1回传的S2的操作员工子账号，如果是子账号操作，请回传。
	SubPromotorPid *string `json:"sub_promotor_pid,omitempty"`
}

// SpISVSignAuthCreateRequest 是支付宝官方 V3 OpenAPI 模型。
type SpISVSignAuthCreateRequest struct {
	// IsvAuthSceneInfos 是代运营授权场景信息。
	IsvAuthSceneInfos *[]ISVAuthSceneInfo `json:"isv_auth_scene_infos,omitempty"`
	// MerchantLogonId 是商户登录账号 支持手机号和邮箱账号；不支持pid。
	MerchantLogonId *string `json:"merchant_logon_id,omitempty"`
	// NeedAppAuth 是1表示需要，0表示不需要 不传参数默认是需要（1） 是否创建三方应用授权子任务，也就是是否需要给三方应用授权。
	NeedAppAuth *string `json:"need_app_auth,omitempty"`
	// SignOrderNo 是签约单号 可通过alipay.open.agent.confirm接口获取签约单号。
	SignOrderNo *string `json:"sign_order_no,omitempty"`
}

// SpISVSignAuthQueryRequest is the request type for alipay.open.sp.isv.signauth.query.
type SpISVSignAuthQueryRequest struct {
	// OrderId 是签约授权一体化订单ID	可通过alipay.open.sp.isv.signauth.create 接口获取
	OrderId *string `json:"order_id,omitempty" query:"order_id,omitempty"`
}
