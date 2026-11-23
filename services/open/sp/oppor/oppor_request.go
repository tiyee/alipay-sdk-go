package oppor

// SpOpporDetailQueryRequest is the request type for alipay.open.sp.oppor.detail.query.
type SpOpporDetailQueryRequest struct {
	// IsvPid 是服务商pid
	IsvPid *string `json:"isv_pid,omitempty" query:"isv_pid,omitempty"`
	// OpporId 是商机作业Id
	OpporId *string `json:"oppor_id,omitempty" query:"oppor_id,omitempty"`
}

// SpOpporPageQueryRequest is the request type for alipay.open.sp.oppor.page.query.
type SpOpporPageQueryRequest struct {
	// IsvPid 是服务商pid
	IsvPid *string `json:"isv_pid,omitempty" query:"isv_pid,omitempty"`
	// StatusList 是作业记录状态列表,以,拼接传入
	StatusList *string `json:"status_list,omitempty" query:"status_list,omitempty"`
	// PageSize 是页大小
	PageSize *string `json:"page_size,omitempty" query:"page_size,omitempty"`
	// PageNum 是页数
	PageNum *string `json:"page_num,omitempty" query:"page_num,omitempty"`
}
