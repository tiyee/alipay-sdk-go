// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package message

// PublicMessageContentCreateRequest 是支付宝官方 V3 OpenAPI 模型。
type PublicMessageContentCreateRequest struct {
	// Benefit 是活动利益点，最多10个字符。仅 ctype 为 activity 类型时才需传入。。
	Benefit *string `json:"benefit,omitempty"`
	// Content 是消息正文（支持富文本）。
	Content *string `json:"content,omitempty"`
	// CouldComment 是是否允许评论。枚举支持： *T：允许。 *F：不允许，默认不允许。。
	CouldComment *string `json:"could_comment,omitempty"`
	// Cover 是封面图url，需传入 <a href="https://opendocs.alipay.com/apis/api_3/alipay.offline.material.image.upload">图片上传接口</a>上传图片获取的 image_url。 注意：尺寸为 996*450，最大不超过3M，支持格式：jpg、.png 。。
	Cover *string `json:"cover,omitempty"`
	// Ctype 是图文类型，不填默认普通图文。还支持 activity（活动图文）。。
	Ctype *string `json:"ctype,omitempty"`
	// ExtTags 是关键词列表，英文逗号分隔，最多不超过5个。
	ExtTags *string `json:"ext_tags,omitempty"`
	// LoginIds 是可预览支付宝账号列表，需要预览时才填写， 英文逗号分隔，最多不超过10个。
	LoginIds *string `json:"login_ids,omitempty"`
	// Title 是标题。
	Title *string `json:"title,omitempty"`
}

// PublicMessageContentModifyRequest 是支付宝官方 V3 OpenAPI 模型。
type PublicMessageContentModifyRequest struct {
	// Benefit 是活动利益点，图文类型ctype为activity类型时才需要传，最多10个字符。
	Benefit *string `json:"benefit,omitempty"`
	// Content 是消息正文（支持富文本）。
	Content *string `json:"content,omitempty"`
	// ContentId 是内容id。
	ContentId *string `json:"content_id,omitempty"`
	// CouldComment 是是否允许评论 T:允许 F:不允许，默认不允许。
	CouldComment *string `json:"could_comment,omitempty"`
	// Cover 是封面图url, 尺寸为996*450，最大不超过3M，支持格式:.jpg、.png ，请先调用<ahref="https://docs.open.alipay.com/api_3/alipay.offline.material.image.upload"> 图片上传接口</a>获得图片url。。
	Cover *string `json:"cover,omitempty"`
	// Ctype 是图文类型 activity: 活动图文，不填默认普通图文。
	Ctype *string `json:"ctype,omitempty"`
	// ExtTags 是关键词列表，英文逗号分隔，最多不超过5个。
	ExtTags *string `json:"ext_tags,omitempty"`
	// LoginIds 是可预览支付宝账号列表，需要预览时才填写， 英文逗号分隔，最多不超过10个。
	LoginIds *string `json:"login_ids,omitempty"`
	// Title 是标题。
	Title *string `json:"title,omitempty"`
}

// PublicMessageCustomSendRequest 是支付宝官方 V3 OpenAPI 模型。
type PublicMessageCustomSendRequest struct {
	// Articles 是msg_type为image-text，本参数必填。。
	Articles *[]Article `json:"articles,omitempty"`
	// Chat 是是否为聊天消息。枚举支持： *0：代表非聊天消息，消息显示在生活号主页，默认为 0。 *1：代表是聊天消息，消息显示在咨询反馈列表页。。
	Chat *string `json:"chat,omitempty"`
	// EventType 是触发消息的事件类型，默认为空。代表商户未改造。如果是follow，代表关注消息。click代表菜单点击，enter_ppchat代表进入事件；请注意事件类型的大小写。
	EventType *string `json:"event_type,omitempty"`
	// MsgType 是消息类型。枚举支持： * text：表示文本消息； * image-text：表示图文消息。。
	MsgType *string `json:"msg_type,omitempty"`
	// OpenId 是支付宝用户的唯一标识。
	OpenId *string `json:"open_id,omitempty"`
	Text   *Text   `json:"text,omitempty"`
	// ToUserId 是消息接收用户的userid。
	ToUserId *string `json:"to_user_id,omitempty"`
}

// PublicMessageGroupSendRequest 是支付宝官方 V3 OpenAPI 模型。
type PublicMessageGroupSendRequest struct {
	// Articles 是图文消息，当 msg_type为image-text，本参数必填。。
	Articles *[]Article `json:"articles,omitempty"`
	// GroupId 是分组ID，通过<a href="https://opendocs.alipay.com/apis/api_6/alipay.open.public.group.create">alipay.open.public.group.create</a>（用户分组创建接口) 创建用户分组获取。。
	GroupId *string `json:"group_id,omitempty"`
	Image   *Image  `json:"image,omitempty"`
	// MsgType 是消息类型。枚举支持： * text：表示文本消息； * image-text：表示图文消息。。
	MsgType *string `json:"msg_type,omitempty"`
	Text    *Text   `json:"text,omitempty"`
}

// PublicMessageLabelSendRequest 是支付宝官方 V3 OpenAPI 模型。
type PublicMessageLabelSendRequest struct {
	Filter   *Filter   `json:"filter,omitempty"`
	Material *Material `json:"material,omitempty"`
}

// PublicMessageSingleSendRequest 是支付宝官方 V3 OpenAPI 模型。
type PublicMessageSingleSendRequest struct {
	Template *Template `json:"template,omitempty"`
	// ToOpenId 是支付宝openId，用于支付宝用户在当前应用下的用户标识。。
	ToOpenId *string `json:"to_open_id,omitempty"`
	// ToUserId 是消息接收用户的支付宝用户id，用户在支付宝的唯一标识，以 2088 开头的 16 位纯数字组成。。
	ToUserId *string `json:"to_user_id,omitempty"`
}

// PublicMessageTotalSendRequest 是支付宝官方 V3 OpenAPI 模型。
type PublicMessageTotalSendRequest struct {
	// Articles 是图文消息，当 msg_type为image-text，本参数必填。。
	Articles *[]Article `json:"articles,omitempty"`
	// MsgType 是消息类型。枚举支持： * text：表示文本消息； * image-text：表示图文消息。。
	MsgType *string `json:"msg_type,omitempty"`
	Text    *Text   `json:"text,omitempty"`
}
