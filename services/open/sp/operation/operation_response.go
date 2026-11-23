package operation

// SpOperationApplyResponse 是支付宝官方 V3 OpenAPI 模型。
type SpOperationApplyResponse struct {
	// BatchNo 是支付宝操作批次号。
	BatchNo     *string              `json:"batch_no,omitempty"`
	BindAccount *RecommendAccountDTO `json:"bind_account,omitempty"`
	// RecommendAccounts 是只针对服务商代间连商家发起代运营绑定、授权时，如果传递的alipay_account不符合绑定或授权要求，这个字段会返回推荐的商家支付宝账号列表，包括：支付宝账号和名称，为保护商家信息，账号和名称都按照规范脱敏。。
	RecommendAccounts *[]RecommendAccountDTO `json:"recommend_accounts,omitempty"`
}

// SpOperationQRCodeQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type SpOperationQRCodeQueryResponse struct {
	// BatchNo 是支付宝操作批次号。
	BatchNo     *string              `json:"batch_no,omitempty"`
	BindAccount *RecommendAccountDTO `json:"bind_account,omitempty"`
	// QrCodeUrl 是二维码图片地址，urlEncode处理过。 注意： 1.有效期为24h; 2.用户扫码确认授权操作（同意/拒绝）后，支付宝将通过<a href="https://opendocs.alipay.com/apis/01lc2d">alipay.open.sp.operation.result.notify</a>接口发送异步通知给服务商。。
	QrCodeUrl *string `json:"qr_code_url,omitempty"`
	// RecommendAccountList 是只针对服务商代间连商家发起代运营绑定、授权时，如果传递的alipay_account不符合绑定或授权要求，这个字段会返回推荐的商家支付宝账号列表，包括：支付宝账号和名称，为保护商家信息，账号和名称都按照规范脱敏。。
	RecommendAccountList *[]RecommendAccountDTO `json:"recommend_account_list,omitempty"`
	RecommendAccounts    *RecommendAccountDTO   `json:"recommend_accounts,omitempty"`
}

// SpOperationResultQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type SpOperationResultQueryResponse struct {
	// BindUserId 是商户支付宝pid。仅间连场景，且存在绑定关系时返回值。。
	BindUserId *string `json:"bind_user_id,omitempty"`
	// HandleStatus 是代运营操作结果。 SUCCESS：代表成功。 PROCESS：待商家确认中。 NO_PERMISSION：表示当前商家支付宝账号无权限操作。需要提醒商家切换成发起授权时指定的支付宝账号。 NONE：表示不存在代运营绑定或授权关系。 NONE_ACCOUNT：间连商家推荐支付宝账号列表为空。。
	HandleStatus *string `json:"handle_status,omitempty"`
	// MerchantNo 是支付宝商户号。间连场景为商户smid，直连场景为商户支付宝pid。
	MerchantNo *string `json:"merchant_no,omitempty"`
}
