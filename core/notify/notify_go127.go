//go:build go1.27

package notify

import (
	"context"
	"net/http"
	"net/url"
)

// Parse 是 Go 1.27 及以上提供的方法入口，委托包级 Parse。
func (h *Handler) Parse[T any](ctx context.Context, values url.Values, decoder Decoder[T]) (Notification[T], error) {
	return Parse(ctx, h, values, decoder)
}

// ParseRequest 是 Go 1.27 及以上提供的方法入口，委托包级 ParseRequest。
func (h *Handler) ParseRequest[T any](ctx context.Context, r *http.Request, decoder Decoder[T]) (Notification[T], error) {
	return ParseRequest(ctx, h, r, decoder)
}
