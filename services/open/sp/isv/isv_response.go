package isv

// SpISVBillSyncResponse 是支付宝官方 V3 OpenAPI 模型。
type SpISVBillSyncResponse struct {
	// Success 是返回true或者flase，ture代表账单回传成功，false表示账单回传失败。
	Success *bool `json:"success,omitempty"`
}

// SpISVSignAuthCreateResponse 是支付宝官方 V3 OpenAPI 模型。
type SpISVSignAuthCreateResponse struct {
	// OrderId 是签约授权一体化主单ID。
	OrderId *string `json:"order_id,omitempty"`
}

// SpISVSignAuthQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type SpISVSignAuthQueryResponse struct {
	// OrderId 是签约授权一体化订单ID。
	OrderId *string `json:"order_id,omitempty"`
	// Status 是签约授权一体化订单状态 INIT:初始化 UNCONFIRM:待确认 FAILURE:失败 FINISH:成功。
	Status *string `json:"status,omitempty"`
	// TaskInfos 是签约授权一体化订单子任务单信息。
	TaskInfos *[]SubTaskInfo `json:"task_infos,omitempty"`
}
