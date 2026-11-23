// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package group

// ComplexLabelRule 是支付宝官方 V3 OpenAPI 模型。
type ComplexLabelRule struct {
	// LabelId 是标签id。
	LabelId *string `json:"label_id,omitempty"`
	// LabelValue 是标签取值，当有多个取值时用英文","分隔（比如使用in操作符时）；不允许传入下划线"_"、竖线"|"或者空格" "。
	LabelValue *string `json:"label_value,omitempty"`
	// Operator 是目前支持EQ（等于）、NEQ（不等于）、LT（小于），GT（大于）、LTEQ（小于等于）、GTEQ（大于等于）、LIKE（匹配）、BETWEEN（范围）、IN（包含）、NOTIN（不包含）操作。
	Operator *string `json:"operator,omitempty"`
}

// QueryComplexLabelRule 是支付宝官方 V3 OpenAPI 模型。
type QueryComplexLabelRule struct {
	// LabelId 是标签id。
	LabelId *string `json:"label_id,omitempty"`
	// LabelName 是标签名。
	LabelName *string `json:"label_name,omitempty"`
	// LabelValue 是当有多个取值时用英文","分隔，不允许传入下划线"_"、竖线"|"或者空格" "。
	LabelValue *string `json:"label_value,omitempty"`
	// Operator 是运算符。
	Operator *string `json:"operator,omitempty"`
}

// QueryGroup 是支付宝官方 V3 OpenAPI 模型。
type QueryGroup struct {
	// Count 是人群中包含人数。
	Count *int64 `json:"count,omitempty"`
	// Id 是分组id。
	Id *string `json:"id,omitempty"`
	// LabelRule 是分组中的圈人规则。
	LabelRule *[]QueryComplexLabelRule `json:"label_rule,omitempty"`
	// Name 是用户分组名称。
	Name *string `json:"name,omitempty"`
}
