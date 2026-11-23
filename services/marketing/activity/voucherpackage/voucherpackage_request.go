// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package voucherpackage

// ActivityVoucherPackageBatchQueryRequest 是支付宝官方 V3 OpenAPI 模型。
type ActivityVoucherPackageBatchQueryRequest struct {
	// PageNum 是分页参数，需要查询的页码，起始页是1
	PageNum *int64 `json:"page_num,omitempty"`
	// PageSize 是分页参数，每页记录数，最大不可超过30
	PageSize *int64 `json:"page_size,omitempty"`
	// VoucherPackagePurchaseStartTime 是券包购买开始时间。 格式为：yyyy-MM-dd HH:mm:ss
	VoucherPackagePurchaseStartTime *string `json:"voucher_package_purchase_start_time,omitempty"`
	// VoucherPackageStatus 是券包状态
	VoucherPackageStatus *string `json:"voucher_package_status,omitempty"`
}

// ActivityVoucherPackageConsultRequest 是支付宝官方 V3 OpenAPI 模型。
type ActivityVoucherPackageConsultRequest struct {
	// AuthToken 在 URL query 中传递。用户授权令牌
	AuthToken *string `query:"auth_token" json:"-"`
	// VoucherPackageIdList 是券包id列表
	VoucherPackageIdList *[]string `json:"voucher_package_id_list,omitempty"`
}

// ActivityVoucherPackageQueryRequest is the query request model for alipay.marketing.activity.voucherpackage.query.
type ActivityVoucherPackageQueryRequest struct {
	// VoucherPackageId 是券包id
	VoucherPackageId *string `json:"voucher_package_id,omitempty" query:"voucher_package_id,omitempty"`
}
