// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package matchuser

// PublicMatchUserLabelCreateRequest 是支付宝官方 V3 OpenAPI 模型。
type PublicMatchUserLabelCreateRequest struct {
	// LabelId 是标签 id，只支持生活号自定义标签。通过 <a href="https://opendocs.alipay.com/apis/api_6/alipay.open.public.life.label.create">alipay.open.public.life.label.create</a>(创建标签接口)创建自定义标签后获取。。
	LabelId *string `json:"label_id,omitempty"`
	// LabelValue 是标签值，由开发者自主指定，标签值类型要满足 <a href="https://opendocs.alipay.com/apis/api_6/alipay.open.public.life.label.create">alipay.open.public.life.label.create</a>(创建标签接口)中data_type参数的限定。。
	LabelValue *string `json:"label_value,omitempty"`
	// Matchers 是支付宝用户匹配器列表，最多传入10条。
	Matchers *[]Matcher `json:"matchers,omitempty"`
}

// PublicMatchUserLabelDeleteRequest 是支付宝官方 V3 OpenAPI 模型。
type PublicMatchUserLabelDeleteRequest struct {
	// LabelId 是标签 id，只支持生活号自定义标签。通过 <a href="https://opendocs.alipay.com/apis/api_6/alipay.open.public.life.label.create">alipay.open.public.life.label.create</a>(创建标签接口)创建自定义标签后获取。。
	LabelId *string `json:"label_id,omitempty"`
	// Matchers 是支付宝用户匹配器列表，最多传入10条。
	Matchers *[]Matcher `json:"matchers,omitempty"`
}
