package core_test

import (
	"errors"
	"testing"

	"github.com/tiyee/alipay-sdk-go/core"
)

func TestQueryEncoderScalarValues(t *testing.T) {
	type request struct {
		Zero     *int64  `query:"zero,omitempty"`
		Negative *int64  `query:"negative"`
		Flag     *bool   `query:"flag,omitempty"`
		Empty    *string `query:"empty,omitempty"`
		Omitted  int     `query:"omitted,omitempty"`
		Missing  *bool   `query:"missing,omitempty"`
	}
	zero, negative, flag, empty := int64(0), int64(-7), false, ""
	encoded, err := core.QueryEncoder[request]().Encode(request{Zero: &zero, Negative: &negative, Flag: &flag, Empty: &empty})
	if err != nil {
		t.Fatal(err)
	}
	if got := encoded.Form.Encode(); got != "empty=&flag=false&negative=-7&zero=0" {
		t.Fatalf("encoded query = %q", got)
	}
}

func TestQueryEncoderRequiredFields(t *testing.T) {
	type pagination struct {
		Page     *int64 `query:"page,required"`
		Detailed *bool  `query:"detailed,required"`
	}
	page, detailed := int64(0), false
	enc := core.QueryEncoder[*pagination]()
	for _, tc := range []struct {
		request *pagination
		field   string
	}{
		{nil, "path_params"},
		{&pagination{}, "page"},
		{&pagination{Page: &page}, "detailed"},
	} {
		_, err := enc.Encode(tc.request)
		var missing *core.ValidationError
		if !errors.As(err, &missing) || missing.Field != tc.field {
			t.Fatalf("got %v, want missing %s", err, tc.field)
		}
	}
	encoded, err := enc.Encode(&pagination{Page: &page, Detailed: &detailed})
	if err != nil || encoded.Form.Encode() != "detailed=false&page=0" {
		t.Fatalf("explicit zero values: %v, %v", encoded.Form, err)
	}
}
