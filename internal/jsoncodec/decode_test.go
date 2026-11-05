package jsoncodec_test

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/tiyee/alipay-sdk-go/internal/jsoncodec"
)

type details struct {
	Status string `json:"status"`
}

type response struct {
	details
	Items []details           `json:"items"`
	ByID  map[string]*details `json:"by_id"`
	Raw   json.RawMessage     `json:"raw"`
}

func TestDecodeContracts(t *testing.T) {
	const raw = `{ "value": 12345678901234567890, "STATUS":"verbatim" }`
	data := []byte(`{"status":"paid","STATUS":"wrong","future":1,"items":[{"Status":"wrong"}],"by_id":{"A":{"status":"ok","sTatus":"wrong"}},"raw":` + raw + `}`)
	original := bytes.Clone(data)
	var got response
	if err := jsoncodec.Unmarshal(data, &got); err != nil {
		t.Fatal(err)
	}
	if got.Status != "paid" || len(got.Items) != 1 || got.Items[0].Status != "" || got.ByID["A"] == nil || got.ByID["A"].Status != "ok" || string(got.Raw) != raw || !bytes.Equal(data, original) {
		t.Fatalf("unexpected decoded data or input mutation: %+v", got)
	}
	if err := jsoncodec.UnmarshalStrict(data, &got); err == nil {
		t.Fatal("accepted an unknown request field")
	}
	if err := jsoncodec.UnmarshalStrict([]byte(`{"items":[{"Status":"wrong"}]}`), &got); err == nil {
		t.Fatal("accepted a case variant in a request")
	}
	for _, input := range []string{
		`{"status":"ok","\u0073tatus":"bad"}`,
		`{"raw":{"x":1,"x":2}}`,
		`{"ignored":[{"x":1,"x":2}]}`,
		`{} []`,
		"{\"status\":\"\xff\"}",
		`{"status":"\ud800"}`,
	} {
		if err := jsoncodec.Unmarshal([]byte(input), &got); err == nil {
			t.Errorf("accepted %q", input)
		}
	}
	if err := jsoncodec.UnmarshalStrict([]byte(`{"status":"\ud83d\ude00"}`), &got); err != nil || got.Status != "😀" {
		t.Fatalf("valid Unicode pair: %q, %v", got.Status, err)
	}
}

func BenchmarkDecodeResponse(b *testing.B) {
	data := []byte(`{"status":"paid","items":[{"status":"settled"}],"by_id":{"order-1":{"status":"ok"}},"raw":{"amount":"88.88"}}`)
	b.ReportAllocs()
	for b.Loop() {
		var result response
		if err := jsoncodec.Unmarshal(data, &result); err != nil {
			b.Fatal(err)
		}
	}
}

func TestConcurrentDecode(t *testing.T) {
	type transaction struct {
		Amount string `json:"amount"`
	}
	type envelope struct {
		Data *transaction `json:"data"`
	}
	input := []byte(`{"Data":{"amount":"999"},"data":{"amount":"1.00"}}`)
	for i := 0; i < 8; i++ {
		t.Run("reader", func(t *testing.T) {
			t.Parallel()
			var v envelope
			if err := jsoncodec.Unmarshal(input, &v); err != nil || v.Data == nil || v.Data.Amount != "1.00" {
				t.Fatalf("unexpected response: %+v, %v", v, err)
			}
		})
	}
}
