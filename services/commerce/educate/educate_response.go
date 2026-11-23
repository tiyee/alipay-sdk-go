// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package educate

// EducateCampusInstitutionsAddResponse 是支付宝官方 V3 OpenAPI 模型。
type EducateCampusInstitutionsAddResponse struct {
	// InstId 是学校内标。如果入驻成功，返回学校内标；入驻失败则不返回学校内标。
	InstId *string `json:"inst_id,omitempty"`
	// InstStdCode 是学校外标。如果入驻成功，返回学校外标；入驻失败则不返回学校外标。
	InstStdCode *string `json:"inst_std_code,omitempty"`
	// Reason 是入驻失败原因
	Reason *string `json:"reason,omitempty"`
}

// EducateCampusInstitutionsQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type EducateCampusInstitutionsQueryResponse struct {
	SchoolInfo *SchoolSimpleInfo `json:"school_info,omitempty"`
	// SchoolInfoList 是学校基础信息列表
	SchoolInfoList *[]SchoolBaseInfo `json:"school_info_list,omitempty"`
}

// EducateSchoolCardOrderSyncResponse 是支付宝官方 V3 OpenAPI 模型。
type EducateSchoolCardOrderSyncResponse struct {
	// OrderId 是订单号
	OrderId *string `json:"order_id,omitempty"`
}
