package core

import "context"

// ExecuteJSON 使用默认编解码器构造 V3 Operation 并执行一次调用。
// 请求编码按 spec.ParameterMode 选择 JSON、query 或 multipart，响应按 JSON 解码。
// spec 必须显式声明认证要求；规格无效时不调用 executor。
// 执行、验签状态检查、错误与响应语义均与 Execute 相同。
// 需要复用 Operation 或自定义编解码器时，使用 NewOperation 和 Execute。
func ExecuteJSON[Req, Resp any](ctx context.Context, executor Executor,
	spec OperationSpec, req Request[Req]) (*Response[Resp], error) {
	operation, err := NewJSONOperation[Req, Resp](spec)
	if err != nil {
		return nil, err
	}
	return Execute(ctx, executor, operation, req)
}

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
	if err := params.validate(spec.ParameterMode); err != nil {
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
