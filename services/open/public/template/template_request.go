// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package template

// PublicTemplateMessageIndustryModifyRequest 是支付宝官方 V3 OpenAPI 模型。
type PublicTemplateMessageIndustryModifyRequest struct {
	// PrimaryIndustryCode 是服务窗消息模板所属主行业一/二级编码，参见 <a href="https://alipay.open.taobao.com/doc2/detail?treeId=197&docType=1&articleId=105043">查看行业信息</a>。。
	PrimaryIndustryCode *string `json:"primary_industry_code,omitempty"`
	// PrimaryIndustryName 是服务窗消息模板所属主行业一/二级名称，参见 <a href="https://alipay.open.taobao.com/doc2/detail?treeId=197&docType=1&articleId=105043">查看行业信息</a>。。
	PrimaryIndustryName *string `json:"primary_industry_name,omitempty"`
	// SecondaryIndustryCode 是服务窗消息模板所属副行业一/二级编码，参见 <a href="https://alipay.open.taobao.com/doc2/detail?treeId=197&docType=1&articleId=105043">查看行业信息</a>。。
	SecondaryIndustryCode *string `json:"secondary_industry_code,omitempty"`
	// SecondaryIndustryName 是服务窗消息模板所属副行业一/二级名称，参见 <a href="https://alipay.open.taobao.com/doc2/detail?treeId=197&docType=1&articleId=105043">查看行业信息</a>。。
	SecondaryIndustryName *string `json:"secondary_industry_name,omitempty"`
}

// PublicTemplateMessageGetRequest is the request type for alipay.open.public.template.message.get.
type PublicTemplateMessageGetRequest struct {
	// TemplateId 是消息母板id，登陆生活号后台(fuwu.alipay.com)，点击菜单“模板消息”，点击“模板库”，即可看到相应模板的消息母板id
	TemplateId *string `json:"template_id,omitempty" query:"template_id,omitempty"`
}
