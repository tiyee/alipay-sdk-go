// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package einvpackage

// InvoiceEInvPackageQueryRequest is the query request model for alipay.ebpp.invoice.einvpackage.query.
type InvoiceEInvPackageQueryRequest struct {
	// EncryptedUid 是加密过的uid。获取详情参见 "推"模式发票报销。
	EncryptedUid *string `json:"encrypted_uid,omitempty" query:"encrypted_uid,omitempty"`
	// PackageId 是发票包id。获取详情参见 "推"模式发票报销。
	PackageId *string `json:"package_id,omitempty" query:"package_id,omitempty"`
}
