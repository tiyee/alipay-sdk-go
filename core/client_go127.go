//go:build go1.27

package core

import "context"

// Execute 实现 Go 1.27 泛型方法形式的调用入口，委托包级 Execute。
func (c *Client) Execute[Req, Resp any](ctx context.Context,
	operation Operation[Req, Resp], req Request[Req]) (*Response[Resp], error) {
	return Execute(ctx, c, operation, req)
}
