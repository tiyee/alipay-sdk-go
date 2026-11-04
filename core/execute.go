package core

import "context"

// Execute 是包级泛型执行入口：编码请求、委托执行器、校验签名状态、
// 解码业务节点。响应类型从 operation 推导，调用方不能任意替换。
// 领域 Service 与自定义接口共用该流程；Client.Execute 仅委托本函数，
// 不存在第二套协议实现。
//
// 失败约定：
//   - 编码、配置或 Operation 无效时不发送请求，Result 为 nil；
//   - 执行器失败时 Result 可能携带已获得的元信息，Data 为 nil；
//   - 仅当返回 error 为 nil 时 Data 非 nil，调用方必须先检查 error。
func Execute[Req, Resp any](ctx context.Context, executor Executor,
	operation Operation[Req, Resp], req Request[Req]) (*Response[Resp], error) {
	if executor == nil {
		return nil, &ConfigError{Field: "executor", Reason: "executor is required"}
	}
	spec, err := operation.validate()
	if err != nil {
		return nil, err
	}
	params, err := operation.encoder.Encode(req.Body)
	if err != nil {
		return nil, &ProtocolError{Op: "encode", Reason: "encode request", Err: err}
	}
	if err := params.validate(); err != nil {
		return nil, err
	}
	result, err := executor.Do(ctx, WireRequest{
		Spec:     spec,
		Params:   params,
		Settings: req.Options,
	})
	if err != nil {
		return &Response[Resp]{Result: result}, err
	}
	if result == nil {
		return nil, &ProtocolError{Op: "execute",
			Reason: "executor returned nil result without error"}
	}
	if !result.SignatureVerified {
		return &Response[Resp]{Result: result}, &SignatureError{Stage: "response",
			Reason: "executor returned success without signature verification"}
	}
	data, err := operation.decoder.Decode(result.Node)
	if err != nil {
		return &Response[Resp]{Result: result}, &ProtocolError{Op: "decode",
			Reason: "decode response node", Err: err}
	}
	return &Response[Resp]{Data: &data, Result: result}, nil
}
