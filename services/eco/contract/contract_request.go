// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package contract

// ContractSignFlowsCreateRequest 是支付宝官方 V3 OpenAPI 模型。
type ContractSignFlowsCreateRequest struct {
	// TargetAppId 在 URL query 中传递。目标isv应用ID
	TargetAppId *string `query:"target_app_id" json:"-"`
	// Attachments 是附件信息（附件可作为合同的辅助证明材料，无需签署。需通过文件直传接口上传附件，获取文件id后添加）。注意事项参见 文件流上传方法。
	Attachments *[]Attachment `json:"attachments,omitempty"`
	// BusinessScene 是流程主题（简要概括合同的签署内容）
	BusinessScene *string         `json:"business_scene,omitempty"`
	ConfigInfo    *ConfigInfoBean `json:"config_info,omitempty"`
	// TemplateInfos 是模板信息（基于合同模板填充内容生成待签文件，并指定签署人）
	TemplateInfos *[]TemplateInfoBean `json:"template_infos,omitempty"`
}
