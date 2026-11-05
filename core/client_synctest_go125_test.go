//go:build go1.25

package core_test

import (
	"context"
	"errors"
	"net/http"
	"testing"
	"testing/synctest"
	"time"

	"github.com/tiyee/alipay-sdk-go/core"
	"github.com/tiyee/alipay-sdk-go/core/option"
)

func TestClientTimeoutWithVirtualTime(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		httpClient := &http.Client{Transport: contextTransport(func(r *http.Request) (*http.Response, error) {
			<-r.Context().Done()
			return nil, r.Context().Err()
		})}
		client := newTestClient(t, "http://alipay.test",
			option.WithHTTPClient(httpClient), option.WithRequestTimeout(time.Minute))
		start := time.Now()
		_, err := core.Execute(context.Background(), client, newQueryOperation(t), core.Request[queryRequest]{
			Body: queryRequest{OutTradeNo: "timeout-1"},
		})
		if !errors.Is(err, context.DeadlineExceeded) {
			t.Fatalf("error = %v, want DeadlineExceeded", err)
		}
		var transportErr *core.TransportError
		if !errors.As(err, &transportErr) {
			t.Fatalf("error type = %T, want *core.TransportError", err)
		}
		if elapsed := time.Since(start); elapsed < time.Minute {
			t.Fatalf("timeout returned too early: %v", elapsed)
		}
	})
}

func TestClientContextCancellation(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		httpClient := &http.Client{Transport: contextTransport(func(r *http.Request) (*http.Response, error) {
			<-r.Context().Done()
			return nil, r.Context().Err()
		})}
		client := newTestClient(t, "http://alipay.test",
			option.WithHTTPClient(httpClient), option.WithRequestTimeout(2*time.Hour))
		ctx, cancel := context.WithCancel(context.Background())
		go func() {
			time.Sleep(time.Second)
			cancel()
		}()
		start := time.Now()
		_, err := core.Execute(ctx, client, newQueryOperation(t), core.Request[queryRequest]{
			Body: queryRequest{OutTradeNo: "cancel-1"},
		})
		if !errors.Is(err, context.Canceled) {
			t.Fatalf("error = %v, want context.Canceled", err)
		}
		var transportErr *core.TransportError
		if !errors.As(err, &transportErr) || transportErr.Kind() != core.ErrorKindCanceled {
			t.Fatalf("error = %v, want canceled transport error", err)
		}
		if elapsed := time.Since(start); elapsed < time.Second {
			t.Fatalf("cancel returned too early: %v", elapsed)
		}
	})
}
