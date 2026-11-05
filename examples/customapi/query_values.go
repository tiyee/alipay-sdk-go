package customapi

import (
	"net/url"

	"github.com/tiyee/alipay-sdk-go/core"
)

// DirectParamsRequest 是机制演示用的直接参数请求模型：
// 不声明 `query` 标签，参数名与取值由 DirectQueryEncoder 自行映射。
type DirectParamsRequest struct {
	OutTradeNo string
	Extra      string
}

// DirectParamsResponse 是机制演示用的直接参数响应模型。
type DirectParamsResponse struct {
	TradeNo string `json:"trade_no"`
}

// NewDirectParamsOperation 用 core.NewOperation 接入自定义编码器，
// 演示外部包自实现编码逻辑的接入方式。method 与 path 是占位值，
// 仅演示编解码机制，不是真实支付宝接口。
func NewDirectParamsOperation() (core.Operation[DirectParamsRequest, DirectParamsResponse], error) {
	return core.NewOperation[DirectParamsRequest, DirectParamsResponse](
		core.OperationSpec{
			Method:          "custom.demo.query.api",
			HTTPMethod:      core.HTTPMethodGet,
			Path:            "/v3/custom/demo/query",
			ParameterMode:   core.ParameterModeQuery,
			AuthRequirement: core.AuthRequirementAny,
		}, DirectQueryEncoder{}, core.JSONDecoder[DirectParamsResponse]())
}

// DirectQueryEncoder 实现 core.Encoder：把具体 struct 编码为 V3
// query 参数。编码器只生成业务参数——V3 的公共身份参数（app_id 等）
// 由 Client 写入 Authorization 签名头，编码器不得生成公共参数；
// 输出必须确定，同一输入产生的编码结果同时用于签名与传输。
type DirectQueryEncoder struct{}

// Encode 实现 core.Encoder：先做必填校验（返回 *core.ValidationError，
// 调用方可按类型匹配），再按自定义映射输出 query 参数；空值参数
// 不进入编码结果。
func (DirectQueryEncoder) Encode(value DirectParamsRequest) (core.EncodedParams, error) {
	if value.OutTradeNo == "" {
		return core.EncodedParams{}, &core.ValidationError{
			Field: "out_trade_no", Reason: "out_trade_no is required"}
	}
	form := url.Values{}
	form.Set("out_trade_no", value.OutTradeNo)
	if value.Extra != "" {
		form.Set("extra", value.Extra)
	}
	return core.EncodedParams{Mode: core.ParameterModeQuery, Form: form}, nil
}
