package twostage

// HSchoolInfo 是支付宝官方 V3 OpenAPI 模型。
type HSchoolInfo struct {
	// CampusNo 是学号
	CampusNo *string `json:"campus_no,omitempty"`
	// SchoolStdCode 是学校外标
	SchoolStdCode *string `json:"school_std_code,omitempty"`
}

// UserIdentityInfo 是支付宝官方 V3 OpenAPI 模型。
type UserIdentityInfo struct {
	HSchoolInfo *HSchoolInfo `json:"h_school_info,omitempty"`
}
