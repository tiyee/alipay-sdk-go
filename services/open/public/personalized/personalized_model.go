// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package personalized

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

// ExtensionArea 是支付宝官方 V3 OpenAPI 模型。
type ExtensionArea struct {
	// GotoUrl 是跳转链接，当type为"image"时必传，必须是https或alipays开头的url链接。已废弃。
	GotoUrl *string `json:"goto_url,omitempty"`
	// Height 是扩展区高度，当type值为"h5"时必填，取值范围为200-500的整数,单位是px。
	Height *int64 `json:"height,omitempty"`
	// Name 是扩展区名字。
	Name *string `json:"name,omitempty"`
	// Type 是扩展区类型，当前支持的值为h5，表示扩展区内容类型为h5。。
	Type *string `json:"type,omitempty"`
	// Url 是扩展区url，传入h5页面url，必须是https开头的链接。
	Url *string `json:"url,omitempty"`
}

// LabelRule 是支付宝官方 V3 OpenAPI 模型。
type LabelRule struct {
	// LabelId 是标签id。
	LabelId *string `json:"label_id,omitempty"`
	// LabelValue 是标签值，当有多个取值时用英文","分隔，不允许传入下划线"_"、竖线"|"或者空格" "和方括号"["、"]"。
	LabelValue *string `json:"label_value,omitempty"`
	// Operator 是目前支持EQ（等于）、BETWEEN（范围）、IN（包含）三种操作符；每个标签支持的运算符可以通过<a href="https://docs.open.alipay.com/api_6/alipay.open.public.life.label.batchquery#sintq">标签列表查询接口</a>获得。该字段允许为空，默认运算符为IN。
	Operator *string `json:"operator,omitempty"`
}

// QueryExtension 是支付宝官方 V3 OpenAPI 模型。
type QueryExtension struct {
	// Areas 是扩展区列表。
	Areas *[]ExtensionArea `json:"areas,omitempty"`
	// ExtensionKey 是扩展区套id。
	ExtensionKey *string `json:"extension_key,omitempty"`
	// LabelRules 是标签规则列表。
	LabelRules *[]QueryLabelRule `json:"label_rules,omitempty"`
	// Status 是扩展区状态，"ON"代表上线，"OFF"代表下线，只有上线的扩展区才能被用户看到。
	Status *string `json:"status,omitempty"`
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
