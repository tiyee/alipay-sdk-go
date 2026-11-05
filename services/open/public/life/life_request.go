// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package life

import (
	"github.com/tiyee/alipay-sdk-go/core"
)

// PublicLifeLabelCreateRequest 是支付宝官方 V3 OpenAPI 模型。
type PublicLifeLabelCreateRequest struct {
	// DataType 是标签值类型，不填默认为 string 类型。 注意：目前只支持 string（字符串类型）。。
	DataType *string `json:"data_type,omitempty"`
	// LabelName 是自定义标签名。 注意：每个生活号最多创建 100 个自定义标签。。
	LabelName *string `json:"label_name,omitempty"`
}

// PublicLifeLabelModifyRequest 是支付宝官方 V3 OpenAPI 模型。
type PublicLifeLabelModifyRequest struct {
	// LabelId 是标签 id，只支持生活号自定义标签。通过 <a href="https://opendocs.alipay.com/apis/api_6/alipay.open.public.life.label.create">alipay.open.public.life.label.create</a>(创建标签接口)创建自定义标签后获取。。
	LabelId *string `json:"label_id,omitempty"`
	// LabelName 是标签名。
	LabelName *string `json:"label_name,omitempty"`
}

// PublicLifeMsgRecallRequest 是支付宝官方 V3 OpenAPI 模型。
type PublicLifeMsgRecallRequest struct {
	// MessageId 是消息id，可在调用<a href="https://opendocs.alipay.com/apis/api_6/alipay.open.public.message.total.send">群发消息</a>、<a href="https://opendocs.alipay.com/apis/api_6/alipay.open.public.message.group.send">组发消息</a>接口时会返回，请注意保存留用。。
	MessageId *string `json:"message_id,omitempty"`
}

// PublicLifeMsgSendRequest 是支付宝官方 V3 OpenAPI 模型。
type PublicLifeMsgSendRequest struct {
	Cover *core.MultipartFile `json:"cover,omitempty" multipart:"cover,file"`
	Data  *PublicLifeMsgSend  `json:"data,omitempty" multipart:"data,json"`
}

// PublicLifeAboardApplyRequest is the request type for alipay.open.public.life.aboard.apply.
type PublicLifeAboardApplyRequest struct{}

// PublicLifeAgentCreateQueryRequest is the request type for alipay.open.public.life.agentcreate.query.
type PublicLifeAgentCreateQueryRequest struct {
	// OutBizNo 是由开发者创建的外部入驻申请单据号，需与alipay.open.public.life.agent.create(isv代创建生活号接口)的out_biz_no一致。
	OutBizNo *string `json:"out_biz_no,omitempty" query:"out_biz_no,omitempty"`
}

// PublicLifeDebarkApplyRequest is the request type for alipay.open.public.life.debark.apply.
type PublicLifeDebarkApplyRequest struct{}

// PublicLifeLabelBatchQueryRequest is the request type for alipay.open.public.life.label.batchquery.
type PublicLifeLabelBatchQueryRequest struct{}

// PublicLifeLabelDeleteRequest is the request type for alipay.open.public.life.label.delete.
type PublicLifeLabelDeleteRequest struct {
	// LabelId 是标签 id，只支持生活号自定义标签。通过 alipay.open.public.life.label.create(创建标签接口)创建自定义标签后获取。
	LabelId *string `json:"label_id,omitempty" query:"label_id,omitempty"`
}
