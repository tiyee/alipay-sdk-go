package box

// SearchBoxApplyResponse 是支付宝官方 V3 OpenAPI 模型。
type SearchBoxApplyResponse struct {
	// BoxId 是搜索直达配置id。
	BoxId *string `json:"box_id,omitempty"`
}

// SearchBoxBatchQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type SearchBoxBatchQueryResponse struct {
	// BoxInfos 是搜索直达基础信息。
	BoxInfos *[]SearchBoxBasicInfo `json:"box_infos,omitempty"`
	// PageNumber 是当前页码。
	PageNumber *int64 `json:"page_number,omitempty"`
	// PageSize 是每页记录数。
	PageSize *int64 `json:"page_size,omitempty"`
	// TotalCount 是总记录数。
	TotalCount *int64 `json:"total_count,omitempty"`
	// TotalPageCount 是总页数。
	TotalPageCount *int64 `json:"total_page_count,omitempty"`
}

// SearchBoxConsultResponse 是支付宝官方 V3 OpenAPI 模型。
type SearchBoxConsultResponse struct {
	// Access 是是否准入，false代表不准入。
	Access *bool `json:"access,omitempty"`
	// NotPassItems 是不准入项列表。不准入项用字符串描述，可参照接口文档和<a href="https://opendocs.alipay.com/pre-open/0311q3#%E5%87%86%E5%85%A5%E6%9D%A1%E4%BB%B6">准入条件</a>，进行调整以达到申请要求。
	NotPassItems *[]string `json:"not_pass_items,omitempty"`
}

// SearchBoxModifyResponse 是支付宝官方 V3 OpenAPI 模型。
type SearchBoxModifyResponse struct {
	// ModuleId 是模块配置ID。
	ModuleId *string `json:"module_id,omitempty"`
}

// SearchBoxOfflineResponse 是支付宝官方 V3 OpenAPI 模型。
type SearchBoxOfflineResponse struct {
}

// SearchBoxOnlineResponse 是支付宝官方 V3 OpenAPI 模型。
type SearchBoxOnlineResponse struct {
}

// SearchBoxQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type SearchBoxQueryResponse struct {
	AccountModule   *SearchBoxAccountModule   `json:"account_module,omitempty"`
	BasicInfoModule *SearchBoxBasicInfoModule `json:"basic_info_module,omitempty"`
	// BoxId 是搜索直达配置id。
	BoxId *string `json:"box_id,omitempty"`
	// BoxStatus 是搜索直达配置状态，AUDIT-审核中/ONLINE-已上架/REJECT-驳回/OFFLINE-已下架。
	BoxStatus              *string                    `json:"box_status,omitempty"`
	BusinessDistrictModule *BoxBusinessDistrictModule `json:"business_district_module,omitempty"`
	// DefaultKeywords 是搜索直达默认触发词，由系统生成，无法修改。
	DefaultKeywords  *[]string               `json:"default_keywords,omitempty"`
	KeywordModule    *SearchBoxKeyWordModule `json:"keyword_module,omitempty"`
	LatestAuditImage *SearchBoxImageModule   `json:"latest_audit_image,omitempty"`
	ServiceModule    *SearchBoxServiceModule `json:"service_module,omitempty"`
	ValidImage       *SearchBoxImageModule   `json:"valid_image,omitempty"`
}

// SearchBoxActivityApplyResponse 是支付宝官方 V3 OpenAPI 模型。
type SearchBoxActivityApplyResponse struct {
	// BoxActivityId 是搜索直达活动ID。
	BoxActivityId *string `json:"box_activity_id,omitempty"`
}

// SearchBoxActivityQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type SearchBoxActivityQueryResponse struct {
	// BackgroundWord 是底纹词。
	BackgroundWord     *string             `json:"background_word,omitempty"`
	BackgroundWordInfo *BackgroundWordInfo `json:"background_word_info,omitempty"`
	// BoxActivityId 是搜索直达活动id。
	BoxActivityId *string `json:"box_activity_id,omitempty"`
	// BoxId 是搜索直达id。
	BoxId *string `json:"box_id,omitempty"`
	// EndTime 是活动结束时间。
	EndTime *string `json:"end_time,omitempty"`
	// FailReason 是审核失败原因。
	FailReason *string `json:"fail_reason,omitempty"`
	// GmtModified 是更新时间。
	GmtModified *string `json:"gmt_modified,omitempty"`
	// MaterialType 是IMAGE-图片/VIDEO-视频。
	MaterialType *string `json:"material_type,omitempty"`
	// MaterialUrl 是当material_type="IMAGE"时，为图片url；当material_type="VIDEO"时，为视频url。
	MaterialUrl  *string                   `json:"material_url,omitempty"`
	OrientedRule *DeliveryOrientedRuleInfo `json:"oriented_rule,omitempty"`
	// RunStatus 是运行状态，INITIAL-初始/ONLINE-已上架/EXPIRE-已失效/OFFLINE-已下架。
	RunStatus *string `json:"run_status,omitempty"`
	// StartTime 是活动开始时间。
	StartTime *string `json:"start_time,omitempty"`
	// Status 是配置状态，INITIAL-初始/AUDIT-审核中/CANCEL-已取消/ONLINE-已上架/REJECT-驳回/OFFLINE-已下架/EXPIRE-已失效。
	Status *string `json:"status,omitempty"`
	// TargetAppid 是跳转应用ID。
	TargetAppid *string `json:"target_appid,omitempty"`
	// TargetAppname 是模板名称。
	TargetAppname *string `json:"target_appname,omitempty"`
	// TargetRegions 是投放目标区域。
	TargetRegions *[]DeliveryTargetRegion `json:"target_regions,omitempty"`
	// Title 是活动标题。
	Title     *string                     `json:"title,omitempty"`
	VideoInfo *SearchBoxActivityVideoInfo `json:"video_info,omitempty"`
}
