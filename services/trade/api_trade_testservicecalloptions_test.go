package trade_test

import (
	"context"
	"testing"

	"github.com/tiyee/alipay-sdk-go/core"
	"github.com/tiyee/alipay-sdk-go/services/trade"
)

func TestServiceCallOptions(t *testing.T) {
	svc, executor := newRecordingService(t)
	executor.result = verifiedResult(`{"code":"10000","msg":"Success"}`)
	// V3 REST 路径不提供 notify_url 调用选项：notify_url 属于请求模型的
	// 业务字段，只有页面产物构建路径消费该调用选项。
	_, err := svc.Query(context.Background(), trade.QueryRequest{TradeNo: "T1"},
		core.WithAppAuthToken("auth-token-1"))
	if err != nil {
		t.Fatalf("Query() error = %v", err)
	}
	if executor.req.Settings.AppAuthToken != "auth-token-1" {
		t.Errorf("AppAuthToken = %q", executor.req.Settings.AppAuthToken)
	}
}
