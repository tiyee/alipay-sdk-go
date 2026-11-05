// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package message

// Article 是支付宝官方 V3 OpenAPI 模型。
type Article struct {
	// ActionName 是链接文字。
	ActionName *string `json:"action_name,omitempty"`
	// Desc 是图文消息内容。
	Desc *string `json:"desc,omitempty"`
	// ImageUrl 是图片链接，对于多条图文消息的第一条消息，该字段不能为空; 请先调用<a href="https://docs.open.alipay.com/api_3/alipay.offline.material.image.upload"> 图片上传接口</a>获得图片url。
	ImageUrl *string `json:"image_url,omitempty"`
	// Title 是图文消息标题。
	Title *string `json:"title,omitempty"`
	// Url 是点击图文消息跳转的链接。
	Url *string `json:"url,omitempty"`
}

// Context 是支付宝官方 V3 OpenAPI 模型。
type Context struct {
	// ActionName 是底部链接描述文字，如“查看详情”，最多能传8个汉字或16个英文字符，长度超出会报错。
	ActionName *string  `json:"action_name,omitempty"`
	First      *Keyword `json:"first,omitempty"`
	// HeadColor 是顶部色条的色值。
	HeadColor *string  `json:"head_color,omitempty"`
	Keyword1  *Keyword `json:"keyword1,omitempty"`
	Keyword2  *Keyword `json:"keyword2,omitempty"`
	Remark    *Keyword `json:"remark,omitempty"`
	// Url 是点击消息后承接页的地址。
	Url *string `json:"url,omitempty"`
}

// Filter 是支付宝官方 V3 OpenAPI 模型。
type Filter struct {
	Context *LabelContext `json:"context,omitempty"`
	// Template 是过滤器模板，${a}是一个变量，会被context参数中的a参数替换，从而展开为最终的表达式，template最多支持两个参数，支持and及or连接符。 and：同时满足条件； or：只需满足其中一个条件。
	Template *string `json:"template,omitempty"`
}

// Image 是支付宝官方 V3 OpenAPI 模型。
type Image struct {
	// Url 是图片url，请先调用alipay.offline.material.image.upload 图片上传接口获得图片url。
	Url *string `json:"url,omitempty"`
}

// Keyword 是支付宝官方 V3 OpenAPI 模型。
type Keyword struct {
	// Color 是当前文字颜色。
	Color *string `json:"color,omitempty"`
	// Value 是模板中占位符的值。
	Value *string `json:"value,omitempty"`
}

// LabelContext 是支付宝官方 V3 OpenAPI 模型。
type LabelContext struct {
	A *LabelFilter `json:"a,omitempty"`
}

// LabelFilter 是支付宝官方 V3 OpenAPI 模型。
type LabelFilter struct {
	// ColumnName 是标签组名，商户自定义的标签固定为label_id_list，支付宝开放的标签详见<a href="http://aopsdkdownload.cn-hangzhou.alipay-pub.aliyun-inc.com/doc/FirstPartOfTagsFromAlipay.xlsx">支付宝开放标签</a>。
	ColumnName *string `json:"column_name,omitempty"`
	// Op 是操作符，支持=、!=、in三个操作符；其中in表示是某几个标签中的一个。
	Op *string `json:"op,omitempty"`
	// Values 是标签数组，用于组装最后的表达式。
	Values *[]string `json:"values,omitempty"`
}

// Material 是支付宝官方 V3 OpenAPI 模型。
type Material struct {
	// Articles 是图文消息子消息项集合，单条消息最多6个子项，否则会发送失败。
	Articles *[]Article `json:"articles,omitempty"`
	// MsgType 是消息类型，text：文本类型，image-text：图文类型。当消息类型为text时，text参数必传，当消息类型为image-text时，articles参数必传。
	MsgType *string `json:"msg_type,omitempty"`
	Text    *Text   `json:"text,omitempty"`
}

// Template 是支付宝官方 V3 OpenAPI 模型。
type Template struct {
	Context *Context `json:"context,omitempty"`
	// TemplateId 是消息模板ID。
	TemplateId *string `json:"template_id,omitempty"`
}

// Text 是支付宝官方 V3 OpenAPI 模型。
type Text struct {
	// Content 是文本消息的内容。
	Content *string `json:"content,omitempty"`
	// Title 是文本消息的标题。
	Title *string `json:"title,omitempty"`
}
