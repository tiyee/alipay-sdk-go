package core

import (
	"testing"
)

func TestQueryEncoderSeparatesPathAndQueryParameters(t *testing.T) {
	type query struct {
		ActivityID string  `json:"-" path:"activity_id"`
		Page       *string `json:"page,omitempty" query:"page,omitempty"`
	}

	page := "2"
	encoded, err := QueryEncoder[query]().Encode(query{ActivityID: "activity/one", Page: &page})
	if err != nil {
		t.Fatal(err)
	}
	if got := encoded.PathParams["activity_id"]; got != "activity/one" {
		t.Fatalf("path parameter = %q", got)
	}
	if got := encoded.Form.Get("page"); got != "2" {
		t.Fatalf("query parameter = %q", got)
	}
	got, err := expandPathParams("/v3/alipay/marketing/activity/{activity_id}", encoded.PathParams)
	if err != nil {
		t.Fatal(err)
	}
	if got != "/v3/alipay/marketing/activity/activity%2Fone" {
		t.Fatalf("expanded path = %q", got)
	}
}

func TestJSONBodyEncoderKeepsPathOutOfBody(t *testing.T) {
	type body struct {
		Name string `json:"name"`
	}
	type request struct {
		body `json:",inline"`
		ID   string `json:"-" path:"delivery_id"`
	}

	encoded, err := JSONBodyEncoder[request]().Encode(request{body: body{Name: "delivery"}, ID: "d-1"})
	if err != nil {
		t.Fatal(err)
	}
	if got, want := string(encoded.JSON), `{"name":"delivery"}`; got != want {
		t.Fatalf("body = %s, want %s", got, want)
	}
	if got := encoded.PathParams["delivery_id"]; got != "d-1" {
		t.Fatalf("path parameter = %q", got)
	}
}

func TestExpandPathParamsRejectsMissingAndUnusedValues(t *testing.T) {
	for _, tc := range []struct {
		name   string
		path   string
		params map[string]string
	}{
		{name: "missing", path: "/v3/items/{item_id}"},
		{name: "unused", path: "/v3/items", params: map[string]string{"item_id": "1"}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if _, err := expandPathParams(tc.path, tc.params); err == nil {
				t.Fatal("expected path parameter error")
			}
		})
	}
}
