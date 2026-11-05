package isv

// ISVAuthSceneInfo 是支付宝官方 V3 OpenAPI 模型。
type ISVAuthSceneInfo struct {
	// SceneCode 是运营场景编码 OPERATION_POINTS：管理运营积分 SHOP_MANAGE：管理门店信息 MINI_APP_OPER：运营支付宝小程序 PROMOTION_MANAGE：运营营销活动。
	SceneCode *string `json:"scene_code,omitempty"`
	// ScenePermissions 是运营场景下的权限编码，多个权限编码以,隔开 1、管理门店信息：SHOP_MANAGE；基础权限（升级）：SHOP_MANAGE_BASE 2、运营营销活动：PROMOTION_MANAGE ；基础权限（升级）：PROMOTION_MANAGE_BASE 3、运营支付宝小程序：MINI_APP_OPER；基础权限（升级）：MINI_APP_OPER_BASE 4、管理运营积分：OPERATION_POINTS；基础权限（升级）：OPERATION_POINTS_BASE。
	ScenePermissions *string `json:"scene_permissions,omitempty"`
}

// SubTaskInfo 是支付宝官方 V3 OpenAPI 模型。
type SubTaskInfo struct {
	// ExecuteResult 是子任务执行结果；执行成功为成功，失败则填充失败结果。
	ExecuteResult *string `json:"execute_result,omitempty"`
	// Status 是"子任务状态状态 INIT:初始化 FAILURE:失败 FINISH:成功"。
	Status *string `json:"status,omitempty"`
	// Type 是子任务类型 SIGN_TASK：签约子任务 OPE_TASK：代运营授权子任务 APP_TASK：应用授权子任务。
	Type *string `json:"type,omitempty"`
	// TypeDesc 是子任务类型描述。
	TypeDesc *string `json:"type_desc,omitempty"`
}
