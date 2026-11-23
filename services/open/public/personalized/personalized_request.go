// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package personalized

// PublicPersonalizedExtensionCreateRequest 是支付宝官方 V3 OpenAPI 模型。
type PublicPersonalizedExtensionCreateRequest struct {
	// Areas 是扩展区列表，最大条数为3。
	Areas *[]ExtensionArea `json:"areas,omitempty"`
	// LabelRule 是标签规则，目前限定只能传入1条，在扩展区上线后，满足该标签规则的用户进入生活号首页，将看到该套扩展区。生成标签及打标详情参见<a href="https://opendocs.alipay.com/fw/api/106877">标签管理</a>。。
	LabelRule *[]LabelRule `json:"label_rule,omitempty"`
}

// PublicPersonalizedMenuCreateRequest 是支付宝官方 V3 OpenAPI 模型。
type PublicPersonalizedMenuCreateRequest struct {
	// Button 是一级菜单列表。说明： * 如果是文本菜单，最多有4个一级菜单，若开发者在后台打开了"咨询反馈"的开关，则只能有3个一级菜单。 * 如果是 ICON 菜单信息，最多有80个一级菜单(忽略二级菜单)。。
	Button *[]ButtonObject `json:"button,omitempty"`
	// GroupId 是人群分组id。分组创建及管理接入详情参见 <a href="https://opendocs.alipay.com/fw/api/106931">分组管理</a>。 注意：group_id 与 label_rule（标签规则）不能同时为空。。
	GroupId *string `json:"group_id,omitempty"`
	// LabelRule 是标签规则，目前限定只能传入1条，在个性化菜单创建成功后，满足该标签规则的用户进入生活号首页，将看到该套菜单。生成标签及打标详情参见<a href="https://opendocs.alipay.com/fw/api/106877">标签管理</a>。 注意：group_id（人群分组id） 与 label_rule 不能同时为空。。
	LabelRule *[]LabelRule `json:"label_rule,omitempty"`
	// MobileClientType 是手机客户端类型，枚举支持： *iphone； *android； *wp； 说明：不填为不区分机型。。
	MobileClientType *string `json:"mobile_client_type,omitempty"`
	// Type 是菜单类型，不填时默认为 text（文本型菜单）。枚举值如下： * text：文本型菜单。 * icon：表示 icon 型菜单，当传值为"icon"时，菜单节点的icon字段必传。。
	Type *string `json:"type,omitempty"`
}

// PublicPersonalizedExtensionBatchQueryRequest is the request type for alipay.open.public.personalized.extension.batchquery.
type PublicPersonalizedExtensionBatchQueryRequest struct{}

// PublicPersonalizedExtensionDeleteRequest is the request type for alipay.open.public.personalized.extension.delete.
type PublicPersonalizedExtensionDeleteRequest struct {
	// ExtensionKey 是一套扩展区的key，删除默认扩展区时传入default ，查询扩展区列表可以获得每套扩展区的key
	ExtensionKey *string `json:"extension_key,omitempty" query:"extension_key,omitempty"`
}

// PublicPersonalizedMenuDeleteRequest is the request type for alipay.open.public.personalized.menu.delete.
type PublicPersonalizedMenuDeleteRequest struct {
	// MenuKey 是要删除的个性化菜单key。通过 alipay.open.public.personalized.menu.create(个性化菜单创建)接口创建个性化菜单后获取。
	MenuKey *string `json:"menu_key,omitempty" query:"menu_key,omitempty"`
}
