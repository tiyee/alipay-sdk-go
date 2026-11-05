// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package voucherpackage

// ActivityVoucherPackageBatchQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type ActivityVoucherPackageBatchQueryResponse struct {
	// PageNum 是分页参数，当前所在的页码
	PageNum *int64 `json:"page_num,omitempty"`
	// PageSize 是分页参数，每页记录数，最大不可超过30
	PageSize *int64 `json:"page_size,omitempty"`
	// TotalSize 是券包总数量
	TotalSize *int64 `json:"total_size,omitempty"`
	// VoucherPackageInfo 是券包信息
	VoucherPackageInfo *[]VoucherPackageInfo `json:"voucher_package_info,omitempty"`
}

// ActivityVoucherPackageConsultResponse 是支付宝官方 V3 OpenAPI 模型。
type ActivityVoucherPackageConsultResponse struct {
	// VoucherPackageConsultResult 是券包购买咨询结果
	VoucherPackageConsultResult *[]VoucherPackageConsultResult `json:"voucher_package_consult_result,omitempty"`
}

// ActivityVoucherPackageQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type ActivityVoucherPackageQueryResponse struct {
	// VoucherPackageActivityInfo 是券包活动信息
	VoucherPackageActivityInfo *[]VoucherPackageActivityInfo `json:"voucher_package_activity_info,omitempty"`
	VoucherPackageBaseInfo     *VoucherPackageBaseInfo       `json:"voucher_package_base_info,omitempty"`
	VoucherPackageSalesInfo    *VoucherPackageSalesInfo      `json:"voucher_package_sales_info,omitempty"`
	VoucherPackageUseRule      *VoucherPackageUseRule        `json:"voucher_package_use_rule,omitempty"`
}
