package core_test

import (
	"context"
	"errors"
	"testing"

	"github.com/tiyee/go-alipay/core"
)

// 验证错误契约：可通过泛型 errors.AsType[T] 匹配，通过 errors.Is 保留原因链。
func TestErrorsAsTypeMatching(t *testing.T) {
	apiErr := &core.APIError{Code: core.ResponseCode("40004"),
		Msg: "Business Failed", SubCode: "ACQ.TRADE_NOT_EXIST", Verified: true}
	var got *core.APIError
	if matched, ok := errors.AsType[*core.APIError](apiErr); !ok || matched.Code != apiErr.Code {
		t.Fatalf("errors.AsType[*core.APIError] = %v, %v", matched, ok)
	}
	_ = got

	// 包装后的 APIError 仍可识别。
	wrapped := &core.ProtocolError{Op: "test", Reason: "wrapped", Err: apiErr}
	if matched, ok := errors.AsType[*core.APIError](wrapped); !ok || matched.SubCode != "ACQ.TRADE_NOT_EXIST" {
		t.Fatalf("errors.AsType through wrap failed: %v, %v", matched, ok)
	}

	// Kind 分类。
	if apiErr.Kind() != core.ErrorKindAPI {
		t.Fatalf("Kind() = %v", apiErr.Kind())
	}
	cancelErr := &core.TransportError{Op: "send", Err: context.Canceled}
	if cancelErr.Kind() != core.ErrorKindCanceled || !errors.Is(cancelErr, context.Canceled) {
		t.Fatalf("canceled classification failed: %v", cancelErr)
	}
	deadlineErr := &core.TransportError{Op: "send", Err: context.DeadlineExceeded}
	if deadlineErr.Kind() != core.ErrorKindTransport || !errors.Is(deadlineErr, context.DeadlineExceeded) {
		t.Fatalf("deadline classification failed: %v", deadlineErr)
	}

	// 错误文本不输出平台自由文本与报文。
	if text := apiErr.Error(); text == "" || containsAny(text, "Business Failed", "交易不存在") {
		t.Fatalf("APIError.Error() leaks platform text: %q", text)
	}
}

func containsAny(s string, subs ...string) bool {
	for _, sub := range subs {
		if len(s) >= len(sub) && (s == sub || indexOf(s, sub) >= 0) {
			return true
		}
	}
	return false
}

func indexOf(s, sub string) int {
	for i := 0; i+len(sub) <= len(s); i++ {
		if s[i:i+len(sub)] == sub {
			return i
		}
	}
	return -1
}

// 验证参数模式与响应码枚举的类型边界。
func TestProtocolEnumBoundaries(t *testing.T) {
	if core.ParameterModeBizJSON.String() != "biz_json" ||
		core.ParameterModeForm.String() != "form" {
		t.Fatalf("parameter mode strings: %v / %v",
			core.ParameterModeBizJSON, core.ParameterModeForm)
	}
	var unset core.ParameterMode
	if unset.String() != "ParameterMode(0)" {
		t.Fatalf("unset mode string = %q", unset.String())
	}
	if core.ResponseCodeSuccess.String() != "10000" {
		t.Fatalf("ResponseCodeSuccess = %q", core.ResponseCodeSuccess)
	}
}
