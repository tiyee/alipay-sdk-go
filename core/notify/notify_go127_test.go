//go:build go1.27

package notify_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/tiyee/alipay-sdk-go/core/notify"
)

func TestGenericMethodsMatchPackageFunctions(t *testing.T) {
	ctx := context.Background()
	handler := newHandler(t)
	form := signedForm(t, baseForm())
	viaMethod, err := handler.Parse(ctx, form, testDecoder)
	if err != nil {
		t.Fatal(err)
	}
	viaPackage, err := notify.Parse(ctx, handler, form, testDecoder)
	if err != nil {
		t.Fatal(err)
	}
	if *viaMethod.Data != *viaPackage.Data {
		t.Fatal("Parse entry points diverge")
	}
	request := func() *http.Request {
		return httptest.NewRequest(http.MethodPost, "/notify", strings.NewReader(form.Encode()))
	}
	viaMethod, err = handler.ParseRequest(ctx, request(), testDecoder)
	if err != nil {
		t.Fatal(err)
	}
	viaPackage, err = notify.ParseRequest(ctx, handler, request(), testDecoder)
	if err != nil {
		t.Fatal(err)
	}
	if *viaMethod.Data != *viaPackage.Data {
		t.Fatal("ParseRequest entry points diverge")
	}
}
