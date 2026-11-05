// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package menu

// ButtonObject 是支付宝官方 V3 OpenAPI 模型。
type ButtonObject struct {
	// ActionParam 是当actionType为link时，该参数为url链接； 当actionType为out时，该参数为用户自定义参数； 当actionType为tel时，该参数为电话号码。 当action_type为map时，该参数为查看地图的关键字。 当action_type为consumption时，该参数可不传。 该参数最长255个字符，不允许冒号等特殊字符。。
	ActionParam *string `json:"action_param,omitempty"`
	// ActionType 是菜单类型： out——事件型菜单； link——链接型菜单； tel——点击拨打电话； map——点击查看地图； consumption——点击查看用户与生活号管理员账号之间的消费记录。
	ActionType *string `json:"action_type,omitempty"`
	// Icon 是icon图片url，必须是http协议的url，尺寸为60X60，最大不超过5M，请先调用<a href="https://docs.open.alipay.com/api_3/alipay.offline.material.image.upload"> 图片上传接口</a>获得图片url。
	Icon *string `json:"icon,omitempty"`
	// Name 是菜单名称，icon菜单名称不超过5个汉字，文本菜单名称不超过9个汉字，编码格式为GBK。
	Name *string `json:"name,omitempty"`
	// SubButton 是二级菜单数组，若sub_button为空，则一级菜单必须指定action_type和action_param的值，二级菜单个数可以为1~5个。。
	SubButton *[]SubButton `json:"sub_button,omitempty"`
}

// MenuAnalysisData 是支付宝官方 V3 OpenAPI 模型。
type MenuAnalysisData struct {
	// AvgClickUserCnt 是人均点击次数。
	AvgClickUserCnt *string `json:"avg_click_user_cnt,omitempty"`
	// ClickCnt 是菜单点击次数。
	ClickCnt *int64 `json:"click_cnt,omitempty"`
	// ClickUserCnt 是菜单点击人数。
	ClickUserCnt *int64 `json:"click_user_cnt,omitempty"`
	// Date 是日期。
	Date *string `json:"date,omitempty"`
	// MenuType 是菜单类型 ，iconDefault ：图标菜单、default：文字菜单。
	MenuType *string `json:"menu_type,omitempty"`
	// Name 是菜单名称。
	Name *string `json:"name,omitempty"`
	// SubName 是子菜单名称，文字菜单才有。
	SubName *string `json:"sub_name,omitempty"`
}

// QueryLabelRule 是支付宝官方 V3 OpenAPI 模型。
type QueryLabelRule struct {
	// LabelId 是标签id。
	LabelId *string `json:"label_id,omitempty"`
	// LabelName 是标签名。
	LabelName *string `json:"label_name,omitempty"`
	// LabelValue 是标签值，多值会用英文逗号分隔。
	LabelValue *string `json:"label_value,omitempty"`
	// Operator 是运算符。
	Operator *string `json:"operator,omitempty"`
}

// QueryMenu 是支付宝官方 V3 OpenAPI 模型。
type QueryMenu struct {
	// Button 是一级菜单列表。
	Button *[]ButtonObject `json:"button,omitempty"`
	// LabelRule 是标签规则项列表。
	LabelRule *[]QueryLabelRule `json:"label_rule,omitempty"`
	// MenuKey 是菜单唯一id。
	MenuKey *string `json:"menu_key,omitempty"`
	// Type 是菜单类型，icon：icon型菜单，text：文本型菜单。
	Type *string `json:"type,omitempty"`
}

// SubButton 是支付宝官方 V3 OpenAPI 模型。
type SubButton struct {
	// ActionParam 是当actionType为link时，该参数为url链接； 当actionType为out时，该参数为用户自定义参数； 当actionType为tel时，该参数为电话号码。 当action_type为map时，该参数为查看地图的关键字。 当action_type为consumption时，该参数可不传。 该参数最长255个字符，不允许冒号等特殊字符。。
	ActionParam *string `json:"action_param,omitempty"`
	// ActionType 是菜单类型： out——事件型菜单； link——链接型菜单； tel——点击拨打电话； map——点击查看地图； consumption——点击查看用户与生活号管理员账号之间的消费记录。
	ActionType *string `json:"action_type,omitempty"`
	// Icon 是icon图片url，必须是http协议的url，尺寸为60X60，最大不超过5M，请先调用<a href="https://docs.open.alipay.com/api_3/alipay.offline.material.image.upload"> 图片上传接口</a>获得图片url。
	Icon *string `json:"icon,omitempty"`
	// Name 是菜单名称，icon菜单名称不超过5个汉字，文本菜单名称不超过9个汉字，编码格式为GBK。
	Name *string `json:"name,omitempty"`
}
