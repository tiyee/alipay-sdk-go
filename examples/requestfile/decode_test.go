package requestfile

import (
	"encoding/json"
	"testing"

	"github.com/tiyee/alipay-sdk-go/core"
	"github.com/tiyee/alipay-sdk-go/services/eco/sign"
)

func TestStrictDecode(t *testing.T) {
	type payload struct {
		Name  string          `json:"name"`
		Child json.RawMessage `json:"child"`
	}
	for _, tt := range []struct {
		name, data string
		bad        bool
	}{
		{"valid", `{"name":"测试","child":{"amount":9007199254740993}}`, false},
		{"unknown", `{"typo":true}`, true},
		{"duplicate", `{"name":"one","name":"two"}`, true},
		{"escaped duplicate", `{"name":"one","na\u006de":"two"}`, true},
		{"nested duplicate", `{"child":[{"a":1,"a":2}]}`, true},
		{"independent objects", `{"child":[{"a":1},{"a":2}]}`, false},
		{"trailing", `{} {}`, true},
		{"malformed", `{"child":[}`, true},
	} {
		t.Run(tt.name, func(t *testing.T) {
			value, err := Decode[payload]([]byte(tt.data))
			if (err != nil) != tt.bad {
				t.Fatalf("error = %v; want bad=%v", err, tt.bad)
			}
			if tt.name == "valid" && string(value.Child) != `{"amount":9007199254740993}` {
				t.Fatalf("raw JSON changed: %s", value.Child)
			}
		})
	}
}

func TestMixedServiceRequestFile(t *testing.T) {
	req, err := Decode[sign.SignFlowCancelRequest]([]byte(`{"target_app_id":"app","flow_id":"flow"}`))
	if err != nil {
		t.Fatal(err)
	}
	query, err := core.QueryEncoder[sign.SignFlowCancelRequest]().Encode(req)
	if err != nil || query.Form.Get("target_app_id") != "app" {
		t.Fatalf("missing query field: %+v, %v", query.Form, err)
	}
	body, err := core.JSONBodyEncoder[sign.SignFlowCancelRequest]().Encode(req)
	if err != nil || string(body.JSON) != `{"flow_id":"flow"}` {
		t.Fatalf("wrong body: %s, %v", body.JSON, err)
	}
	for _, invalid := range []string{`null`, `{"target_app_id":"a","target_app_id":"b"}`, `{"TargetAppId":"a"}`, `{"target_app_id":"a","unknown":1}`, `{"target_app_id":true}`} {
		if req, err := Decode[sign.SignFlowCancelRequest]([]byte(invalid)); err == nil || req != (sign.SignFlowCancelRequest{}) {
			t.Errorf("accepted invalid request file %s", invalid)
		}
	}
}
