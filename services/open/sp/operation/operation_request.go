package operation

// SpOperationApplyRequest 是支付宝官方 V3 OpenAPI 模型。
type SpOperationApplyRequest struct {
	// AccessProductCode 是接入的产品编号。 枚举如下： * OPENAPI_BIND_DEFAULT：操作类型为账号绑定。 * OPENAPI_AUTH_DEFAULT：操作类型为代运营授权时。。
	AccessProductCode *string `json:"access_product_code,omitempty"`
	// AlipayAccount 是支付宝登录账号，通常为手机号或者邮箱。 间连场景必填。 直连场景选填，特别注意merchant_no和alipay_account不能同时为空，都有值优先取merchant_no。。
	AlipayAccount *string `json:"alipay_account,omitempty"`
	// IsvScenePermissions 是场景授权列表结构结构：场景codeA:权限code1,权限code2;场景codeB:权限code1,权限code2; 说明： * 本参数和access_product_code只需要传一个。 * 场景 + 权限Code含义： SHOP_MANAGE:SHOP_MANAGE_BASE：管理门店信息 MINI_APP_OPER:MINI_APP_OPER_BASE：运营支付宝小程序 PROMOTION_MANAGE:PROMOTION_MANAGE_BASE：运营营销活动 OPERATION_POINTS:OPERATION_POINTS_BASE：管理运营积分 INCENTIVE_POINT_MANAGE:INCENTIVE_POINT_MANAGE_BASE：管理激励点数。
	IsvScenePermissions *string `json:"isv_scene_permissions,omitempty"`
	// MerchantNo 是支付宝商户号。 间连场景，merchant_no必填，传入商户smid，特别注意仅支持2088开头的间连商户。 直连场景，merchant_no选填，传入商户支付宝pid，特别注意merchant_no和alipay_account不能同时为空，优先取merchant_no。。
	MerchantNo *string `json:"merchant_no,omitempty"`
	// OperateType 是代运营操作类型。枚举如下： * ACCOUNT_BIND：代表绑定支付宝账号，仅对于间连商户。 * OPERATION_AUTH：代表代运营授权，支持间连和直连商户，其中间连场景包含绑定支付宝账号。。
	OperateType *string `json:"operate_type,omitempty"`
	// OutBizNo 是外部操作流水，ISV自定义。每次操作需要确保唯一。。
	OutBizNo *string `json:"out_biz_no,omitempty"`
}

// SpOperationQRCodeQueryRequest is the request type for alipay.open.sp.operation.qrcode.query.
type SpOperationQRCodeQueryRequest struct {
	// OutBizNo 是外部操作流水，由服务商自定义，需确保每次操作唯一。仅支持数字、字母、下划线组合。
	OutBizNo *string `json:"out_biz_no,omitempty" query:"out_biz_no,omitempty"`
	// OperateType 是代运营操作类型。取值如下：	* ACCOUNT_BIND：账号绑定，仅支持间连商户。	* OPERATION_AUTH：代运营授权，支持间连、直连商户。
	OperateType *string `json:"operate_type,omitempty" query:"operate_type,omitempty"`
	// MerchantNo 是支付宝商户号。注意仅支持 2088 开头的间连商户。	若被代运营者是间连商户，则 merchant_no 必填；	若为直连商户，则 merchant_no 和 alipay_account 不能同时为空，都有值优先取 merchant_no。
	MerchantNo *string `json:"merchant_no,omitempty" query:"merchant_no,omitempty"`
	// AlipayAccount 是支付宝登录账号。通常为手机号或者邮箱。	若被代运营者是间连商户，该字段无需填写。	若为直连商户，则 merchant_no 和 alipay_account 不能同时为空，都有值优先取 merchant_no。
	AlipayAccount *string `json:"alipay_account,omitempty" query:"alipay_account,omitempty"`
	// AccessProductCode 是接入的产品编号。	操作类型为绑定的，填OPENAPI_BIND_DEFAULT;	操作类型为授权的，填OPENAPI_AUTH_DEFAULT
	AccessProductCode *string `json:"access_product_code,omitempty" query:"access_product_code,omitempty"`
	// IsvScenePermissions 是场景授权列表结构结构：场景codeA:权限code1,权限code2;场景codeB:权限code1,权限code2;
	IsvScenePermissions *string `json:"isv_scene_permissions,omitempty" query:"isv_scene_permissions,omitempty"`
}

// SpOperationResultQueryRequest is the request type for alipay.open.sp.operation.result.query.
type SpOperationResultQueryRequest struct {
	// OperateType 是代运营操作类型。取值如下 ACCOUNT_BIND：账号绑定； OPERATION_AUTH：代运营授权。
	OperateType *string `json:"operate_type,omitempty" query:"operate_type,omitempty"`
	// BatchNo 是支付宝操作批次号。	batch_no提供基于操作批次的精确查询，只查询本操作批次的代运营操作结果，不反映最终代运营操作结果。	注意：	1.access_product_code、merchant_no、alipay_account三个参数组合与传入 batch_no 为两种查询方式，不可同时为空且以 batch_no 优先。	2.access_product_code、merchant_no、alipay_account三个参数组合可以查询最终代运营操作结果。
	BatchNo *string `json:"batch_no,omitempty" query:"batch_no,omitempty"`
	// AccessProductCode 是接入的产品编号。未传batch_no时必填，枚举如下：	OPENAPI_BIND_DEFAULT：操作类型为账号绑定；	OPENAPI_AUTH_DEFAULT：操作类型为代运营授权。
	AccessProductCode *string `json:"access_product_code,omitempty" query:"access_product_code,omitempty"`
	// MerchantNo 是支付宝商户号。支持间连、直连商户，值为2088开头。	未传batch_no时需注意：	若查询间连商户则 merchant_no 必填； 	若查询直连商户，则merchant_no和alipay_account不能同时为空，都有值优先取merchant_no。
	MerchantNo *string `json:"merchant_no,omitempty" query:"merchant_no,omitempty"`
	// AlipayAccount 是支付宝登录账号，通常为手机号或者邮箱。未传batch_no时需注意：	1.当查询merchant_no为间连商户时：	  * 若 operate_type=ACCOUNT_BIND，则alipay_account必填。	  * 若 operate_type=OPERATION_AUTH，则alipay_account无需传入。	2.当查询merchant_no为直连商户时：	  * merchant_no和alipay_account不能同时为空，都有值优先取merchant_no。
	AlipayAccount *string `json:"alipay_account,omitempty" query:"alipay_account,omitempty"`
	// IsvScenePermissions 是场景授权列表结构结构：场景codeA:权限code1,权限code2;场景codeB:权限code1,权限code2;
	IsvScenePermissions *string `json:"isv_scene_permissions,omitempty" query:"isv_scene_permissions,omitempty"`
}
