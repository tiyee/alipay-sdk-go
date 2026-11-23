// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package expensescene

// InvoiceExpenseSceneMerchantQueryRequest is the query request model for alipay.ebpp.invoice.expensescene.merchant.query.
type InvoiceExpenseSceneMerchantQueryRequest struct {
	// AccountId 是共同账户ID
	AccountId *string `json:"account_id,omitempty" query:"account_id,omitempty"`
	// AgreementNo 是授权签约协议号
	AgreementNo *string `json:"agreement_no,omitempty" query:"agreement_no,omitempty"`
	// EnterpriseId 是企业ID
	EnterpriseId *string `json:"enterprise_id,omitempty" query:"enterprise_id,omitempty"`
	// RoleId 是商户ID
	RoleId *string `json:"role_id,omitempty" query:"role_id,omitempty"`
	// Pid 是商户ID
	Pid *string `json:"pid,omitempty" query:"pid,omitempty"`
	// ShopId 是门店ID
	ShopId *string `json:"shop_id,omitempty" query:"shop_id,omitempty"`
	// ExpenseType 是费用类型	MEAL - 工作餐
	ExpenseType *string `json:"expense_type,omitempty" query:"expense_type,omitempty"`
	// PageNum 是页码
	PageNum *int64 `json:"page_num,omitempty" query:"page_num,required"`
	// PageSize 是每页行数（最大100）
	PageSize *int64 `json:"page_size,omitempty" query:"page_size,required"`
}
