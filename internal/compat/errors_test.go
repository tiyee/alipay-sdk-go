package compat_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/tiyee/alipay-sdk-go/internal/compat"
)

type testError struct{}

func (*testError) Error() string { return "test" }

type declinedMatch struct{}

func (declinedMatch) Error() string { return "declined" }
func (declinedMatch) As(target any) bool {
	if dst, ok := target.(**testError); ok {
		*dst = &testError{}
	}
	return false
}

func TestAsType(t *testing.T) {
	want := &testError{}
	err := errors.Join(errors.New("other"), fmt.Errorf("wrapped: %w", want))
	if got, ok := compat.AsType[*testError](err); !ok || got != want {
		t.Fatal("lost wrapped error")
	}
	if got, ok := compat.AsType[*testError](declinedMatch{}); ok || got != nil {
		t.Fatal("unsuccessful match must return zero")
	}
}
