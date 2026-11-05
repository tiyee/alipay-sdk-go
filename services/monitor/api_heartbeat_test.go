package monitor_test

import (
	"context"
	"strings"
	"testing"

	"github.com/tiyee/alipay-sdk-go/core"
	"github.com/tiyee/alipay-sdk-go/services/monitor"
)

type executor struct{ request core.WireRequest }

func (e *executor) Do(_ context.Context, req core.WireRequest) (*core.APIResult, error) {
	e.request = req
	return &core.APIResult{HTTPStatus: 200, SignatureVerified: true, Node: []byte(`{"pid":"2088"}`)}, nil
}

func TestHeartbeatSyn(t *testing.T) {
	e := &executor{}
	s, err := monitor.NewService(e)
	if err != nil {
		t.Fatal(err)
	}
	content := "中文 + %"
	resp, err := s.HeartbeatSyn(context.Background(), monitor.HeartbeatSynRequest{BizContent: &content})
	if err != nil {
		t.Fatal(err)
	}
	if e.request.Spec.Method != "monitor.heartbeat.syn" || e.request.Spec.HTTPMethod != core.HTTPMethodPost || e.request.Spec.Path != "/v3/monitor/heartbeat/syn" {
		t.Fatalf("unexpected operation: %+v", e.request.Spec)
	}
	if !strings.Contains(string(e.request.Params.JSON), `"biz_content":"中文 + %"`) {
		t.Fatalf("unexpected body: %s", e.request.Params.JSON)
	}
	if resp.Data.PID == nil || *resp.Data.PID != "2088" {
		t.Fatalf("unexpected response: %+v", resp.Data)
	}
}

func TestNewServiceRejectsNil(t *testing.T) {
	if _, err := monitor.NewService(nil); err == nil {
		t.Fatal("expected error")
	}
}
