// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package social

// SourceContentInfo 是支付宝官方 V3 OpenAPI 模型。
type SourceContentInfo struct {
	// ContentId 是内容ID
	ContentId *string `json:"content_id,omitempty"`
	// Link 是支付宝内容链接
	Link *string `json:"link,omitempty"`
	// PermissionStatus 是内容分发范围。0 - 仅创作者自己可见，1 - 所有人可见，2 - 分发范围受限。
	PermissionStatus *string `json:"permission_status,omitempty"`
	// SourceAuthor 是作者
	SourceAuthor *string `json:"source_author,omitempty"`
	// SourceLink 是来源文章的原始链接地址
	SourceLink *string `json:"source_link,omitempty"`
	// SourceMediaInfos 是素材列表
	SourceMediaInfos *[]SourceMediaInfo `json:"source_media_infos,omitempty"`
	// SourceOffers 是关联服务
	SourceOffers *[]SourceOffer `json:"source_offers,omitempty"`
	// SourcePublishDate 是文章发布时间
	SourcePublishDate *string `json:"source_publish_date,omitempty"`
	// SourceStatus 是内容状态	0-审核中	1-成功发布	2-审核不通过	3-己删除
	SourceStatus *string `json:"source_status,omitempty"`
	// SourceSummary 是文章的摘要
	SourceSummary *string `json:"source_summary,omitempty"`
	// SourceTitle 是内容标题
	SourceTitle *string `json:"source_title,omitempty"`
	// SourceType 是内容类型（0-长图文 1-短图文 2-视频）
	SourceType *string `json:"source_type,omitempty"`
	// SpecialTags 是平台加工后的内容标签，例如	HIGH_QUALITY: 优质
	SpecialTags *[]string `json:"special_tags,omitempty"`
}

// SourceMediaInfo 是支付宝官方 V3 OpenAPI 模型。
type SourceMediaInfo struct {
	// MediaId 是素材ID，对应“支付宝文件上传接口”获取的file_id
	MediaId *string `json:"media_id,omitempty"`
	// MediaType 是素材类型。	 image：图片	 video：视频	cover_static：静态封面
	MediaType *string `json:"media_type,omitempty"`
}

// SourceOffer 是支付宝官方 V3 OpenAPI 模型。
type SourceOffer struct {
	// OfferId 是offerId，根据offer类型对应不同ID。支持小程序的APPID/小程序服务的服务编码/优惠券的活动ID
	OfferId *string `json:"offer_id,omitempty"`
	// OfferType 是offer类型。目前仅支持小程序（mini_app/service）、小程序服务（subservice）、营销活动（voucher）、商品(goods)、小游戏(game)。	注意：其中“小程序”offer类型，在发布接口请用mini_app，查询接口返回为service
	OfferType *string `json:"offer_type,omitempty"`
}
