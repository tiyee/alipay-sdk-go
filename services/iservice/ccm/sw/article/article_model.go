package article

import (
	"github.com/tiyee/alipay-sdk-go/services/iservice/ccm/sw"
)

// ArticleAttachmentInfo 是支付宝官方 V3 OpenAPI 模型。
type ArticleAttachmentInfo struct {
	// FileName 是文章附件
	FileName *string `json:"file_name,omitempty"`
	// Id 是附件id
	Id *int64 `json:"id,omitempty"`
	// Key 是附件完整Key
	Key *string `json:"key,omitempty"`
	// Url 是附件的url
	Url *string `json:"url,omitempty"`
}

// ArticleInfo 是支付宝官方 V3 OpenAPI 模型。
type ArticleInfo struct {
	// Attachments 是文章对应附件集合
	Attachments *[]ArticleAttachmentInfo `json:"attachments,omitempty"`
	// CategoryId 是所属类目ID
	CategoryId *int64 `json:"category_id,omitempty"`
	// CategoryNamePath 是类目名称路径
	CategoryNamePath *string `json:"category_name_path,omitempty"`
	// CategoryPath 是类目路径
	CategoryPath *[]sw.ArticleCategoryInfo `json:"category_path,omitempty"`
	// Content 是内容
	Content *string `json:"content,omitempty"`
	// CreateTime 是创建时间
	CreateTime *string `json:"create_time,omitempty"`
	// CreatorId 是创建人
	CreatorId *string `json:"creator_id,omitempty"`
	// ExtendTitles 是扩展标题
	ExtendTitles *[]string `json:"extend_titles,omitempty"`
	// Id 是文章ID
	Id *int64 `json:"id,omitempty"`
	// Keywords 是标签
	Keywords *[]string `json:"keywords,omitempty"`
	// LibraryId 是知识库ID
	LibraryId *int64 `json:"library_id,omitempty"`
	// OrderNo 是排序值
	OrderNo *int64 `json:"order_no,omitempty"`
	// Pictures 是文章对应图片集合
	Pictures *[]ArticleAttachmentInfo `json:"pictures,omitempty"`
	// PublishEnd 是有效期止
	PublishEnd *string `json:"publish_end,omitempty"`
	// PublishStart 是有效期起始
	PublishStart *string `json:"publish_start,omitempty"`
	// SceneCodes 是场景ID。1（内部知识库）；2（机器人）;3（帮助中心）；4（无线帮助中心）
	SceneCodes *[]string `json:"scene_codes,omitempty"`
	// Source 是来源
	Source *string `json:"source,omitempty"`
	// Status 是文章状态
	Status *string `json:"status,omitempty"`
	// StatusCode 是文章状态码
	StatusCode *string `json:"status_code,omitempty"`
	// Title 是标题
	Title *string `json:"title,omitempty"`
	// UpdateTime 是修改时间
	UpdateTime *string `json:"update_time,omitempty"`
	// UpdaterId 是修改人
	UpdaterId *string `json:"updater_id,omitempty"`
	// UpdaterName 是修改人名称
	UpdaterName *string `json:"updater_name,omitempty"`
}

// BaseArticleInfo 是支付宝官方 V3 OpenAPI 模型。
type BaseArticleInfo struct {
	// Id 是知识点id
	Id *int64 `json:"id,omitempty"`
	// Title 是知识点标题
	Title *string `json:"title,omitempty"`
}
