package article

// CcmSwArticleBatchQueryRequest 是支付宝官方 V3 OpenAPI 模型。
type CcmSwArticleBatchQueryRequest struct {
	// CategoryId 是所属类目ID，如果search_all_category为true则不用填
	CategoryId *int64 `json:"category_id,omitempty"`
	// CcsInstanceId 是子部门ID，不传为默认部门
	CcsInstanceId *string `json:"ccs_instance_id,omitempty"`
	// EndTime 是结束时间，并且开始时间不能为空
	EndTime *string `json:"end_time,omitempty"`
	// Ids 是文章ID集合
	Ids *[]int64 `json:"ids,omitempty"`
	// Keyword 是关键字
	Keyword *string `json:"keyword,omitempty"`
	// Keywords 是标签
	Keywords *[]string `json:"keywords,omitempty"`
	// LibraryId 是知识库ID
	LibraryId *int64 `json:"library_id,omitempty"`
	// PageNum 是页数，page_size不能为空
	PageNum *int64 `json:"page_num,omitempty"`
	// PageSize 是页大小，page_num不能为空
	PageSize *int64 `json:"page_size,omitempty"`
	// SearchAllCategory 是是否搜索所有类目，如果为true则不用填写category_id值
	SearchAllCategory *bool `json:"search_all_category,omitempty"`
	// SearchCategoryType 是Current（搜索当前节点）；	Children（搜索当前节点以及子节点）
	SearchCategoryType *string `json:"search_category_type,omitempty"`
	// StartTime 是开始时间，并且结束时间不能为空
	StartTime *string `json:"start_time,omitempty"`
	// Status 是状态，PUBLISHED（已发布），UNPUBLISH（未发布），EXPIRED（失效），DELETED（已删除）
	Status *string `json:"status,omitempty"`
}

// CcmSwArticleCreateRequest 是支付宝官方 V3 OpenAPI 模型。
type CcmSwArticleCreateRequest struct {
	// Attachments 是附件列表
	Attachments *[]ArticleAttachmentInfo `json:"attachments,omitempty"`
	// CategoryId 是所属类目ID，如果search_all_category为true则不用填
	CategoryId *int64 `json:"category_id,omitempty"`
	// CcsInstanceId 是子部门ID，不传为默认部门
	CcsInstanceId *string `json:"ccs_instance_id,omitempty"`
	// Content 是内容
	Content *string `json:"content,omitempty"`
	// ExtendTitles 是扩展标题
	ExtendTitles *[]string `json:"extend_titles,omitempty"`
	// Keywords 是标签
	Keywords *[]string `json:"keywords,omitempty"`
	// LibraryId 是知识库ID
	LibraryId *string `json:"library_id,omitempty"`
	// PublishEnd 是默认为 2999-12-31 23:59:59
	PublishEnd *string `json:"publish_end,omitempty"`
	// PublishStart 是默认为当前时间
	PublishStart *string `json:"publish_start,omitempty"`
	// RelatedArticles 是关联知识点
	RelatedArticles *[]BaseArticleInfo `json:"related_articles,omitempty"`
	// SceneCodes 是场景ID。KNOWLEDGE（内部知识库）；ROBOT（机器人）;HELP（帮助中心）；WHELP（无线帮助中心）
	SceneCodes *[]string `json:"scene_codes,omitempty"`
	// Title 是标题
	Title *string `json:"title,omitempty"`
}

// CcmSwArticleModifyRequest 是支付宝官方 V3 OpenAPI 模型。
type CcmSwArticleModifyRequest struct {
	// Attachments 是附件列表
	Attachments *[]ArticleAttachmentInfo `json:"attachments,omitempty"`
	// CategoryId 是所属类目ID，如果search_all_category为true则不用填
	CategoryId *int64 `json:"category_id,omitempty"`
	// CcsInstanceId 是子部门ID，不传为默认部门
	CcsInstanceId *string `json:"ccs_instance_id,omitempty"`
	// Content 是内容
	Content *string `json:"content,omitempty"`
	// ExtendTitles 是扩展标题
	ExtendTitles *[]string `json:"extend_titles,omitempty"`
	// Id 是文章ID
	Id *int64 `json:"id,omitempty"`
	// Keywords 是标签
	Keywords *[]string `json:"keywords,omitempty"`
	// PublishEnd 是有效期结束时间
	PublishEnd *string `json:"publish_end,omitempty"`
	// PublishStart 是有效期开始时间
	PublishStart *string `json:"publish_start,omitempty"`
	// RelatedArticles 是关联知识点列表
	RelatedArticles *[]BaseArticleInfo `json:"related_articles,omitempty"`
	// SceneCodes 是场景ID。1（内部知识库）；2（机器人）;3（帮助中心）；4（无线帮助中心）
	SceneCodes *[]string `json:"scene_codes,omitempty"`
	// Title 是标题
	Title *string `json:"title,omitempty"`
}

// CcmSwArticleDeleteRequest is the request model for alipay.iservice.ccm.sw.article.delete.
type CcmSwArticleDeleteRequest struct{}

// CcmSwArticleGetRequest is the query request model for alipay.iservice.ccm.sw.article.get.
type CcmSwArticleGetRequest struct {
	// CcsInstanceId 是子部门ID，不传为默认部门
	CcsInstanceId *string `json:"ccs_instance_id,omitempty" query:"ccs_instance_id,omitempty"`
	// Id 是文章ID
	Id *int64 `json:"id,omitempty" query:"id,required"`
}
