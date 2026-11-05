// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package life

// PublicLifeMsgSend 是支付宝官方 V3 OpenAPI 模型。
type PublicLifeMsgSend struct {
	// Category 是消息分类，请传入对应分类编码值。
	Category *string `json:"category,omitempty"`
	// Content 是消息正文，html原文或纯文本。
	Content *string `json:"content,omitempty"`
	// Desc 是消息概述。
	Desc *string `json:"desc,omitempty"`
	// MsgType 是媒体资讯类生活号消息类型。
	MsgType *string `json:"msg_type,omitempty"`
	// SourceExtInfo 是消息来源方附属信息，供搜索、推荐使用 publish_time（int）：消息发布时间，单位秒 keyword_list（String）: 文章的标签列表，英文逗号分隔 comment（int）：消息来源处评论次数 reward（int）：消息来源处打赏次数 is_recommended（boolean）：消息在来源处是否被推荐 is_news（boolean）：消息是否实时性内容 read（int）：消息在来源处被阅读次数 like（int）：消息在来源处被点赞次数 is_hot（boolean）：消息在来源平台是否是热门内容 share（int）：文章在来源平台的分享次数 deadline（int）：文章的失效时间，单位秒。
	SourceExtInfo *string `json:"source_ext_info,omitempty"`
	// Title 是消息标题。
	Title *string `json:"title,omitempty"`
	// UniqueMsgId 是来源方消息唯一标识；若不为空，根据此id和生活号id对消息去重；若为空，不去重。
	UniqueMsgId *string `json:"unique_msg_id,omitempty"`
	// VideoLength 是生活号消息视频时长，单位：秒（视频类消息必填）。
	VideoLength *string `json:"video_length,omitempty"`
	// VideoSamples 是视频类型消息中视频抽样关键帧截图，视频类消息选填。
	VideoSamples *[]string `json:"video_samples,omitempty"`
	// VideoSize 是视频大小，单位：KB（视频类消息必填）。
	VideoSize *string `json:"video_size,omitempty"`
	// VideoSource 是视频资源来源id（视频类消息必填），取值限定youku, miaopai, taobao, sina中的一个。
	VideoSource *string `json:"video_source,omitempty"`
	// VideoTemporaryUrl 是视频的临时链接（优酷来源的视频消息，该字段不能为空）。
	VideoTemporaryUrl *string `json:"video_temporary_url,omitempty"`
	// VideoUrl 是生活号视频类消息视频id或url（视频类消息必填，根据来源区分）。
	VideoUrl *string `json:"video_url,omitempty"`
}

// LifeLabel 是支付宝官方 V3 OpenAPI 模型。
type LifeLabel struct {
	// Biz 是该标签支持的业务列表，menu表示个性化菜单，extension表示个性化扩展区，message表示消息触达。
	Biz *string `json:"biz,omitempty"`
	// Category 是标签类目。
	Category *string `json:"category,omitempty"`
	// DataType 是标签值数据类型。
	DataType *string `json:"data_type,omitempty"`
	// LabelCode 是标签英文代码。
	LabelCode *string `json:"label_code,omitempty"`
	// LabelId 是标签id，唯一标识一个标签。
	LabelId *string `json:"label_id,omitempty"`
	// LabelName 是标签名。
	LabelName *string `json:"label_name,omitempty"`
	// Operator 是该标签支持的运算符。
	Operator *string `json:"operator,omitempty"`
	// Options 是每个取值的业务含义。
	Options *[]Option `json:"options,omitempty"`
	// Type 是标签类型，目前分为common（通用标签）、custom（生活号自定义标签）、cloud（云实验室标签）。
	Type *string `json:"type,omitempty"`
}

// Option 是支付宝官方 V3 OpenAPI 模型。
type Option struct {
	// Text 是文本，通常用于理解对应的取值。
	Text *string `json:"text,omitempty"`
	// Value 是取值，通常使用简单的数字或字符串。
	Value *string `json:"value,omitempty"`
}
