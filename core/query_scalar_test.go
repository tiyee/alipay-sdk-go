package core_test

import (
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
