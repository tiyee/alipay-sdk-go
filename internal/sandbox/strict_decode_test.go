package sandbox

import (
	"encoding/json"
	"testing"
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
			value, err := decodeRequest[payload]([]byte(tt.data))
			if (err != nil) != tt.bad {
				t.Fatalf("error = %v; want bad=%v", err, tt.bad)
			}
			if tt.name == "valid" && string(value.Child) != `{"amount":9007199254740993}` {
				t.Fatalf("raw JSON changed: %s", value.Child)
			}
		})
	}
}
