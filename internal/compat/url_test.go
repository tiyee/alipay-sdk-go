package compat_test

import (
	"net/url"
	"reflect"
	"testing"

	"github.com/tiyee/alipay-sdk-go/internal/compat"
)

func TestCloneIsolation(t *testing.T) {
	if compat.CloneValues(nil) != nil || compat.CloneURL(nil) != nil {
		t.Fatal("nil clone")
	}
	src := url.Values{"multi": {"a", "b"}, "unset": nil, "empty": {}}
	dst := compat.CloneValues(src)
	if !reflect.DeepEqual(src, dst) {
		t.Fatal("clone lost empty/nil values")
	}
	dst["multi"][0] = "changed"
	delete(dst, "empty")
	if src.Get("multi") != "a" || !src.Has("empty") {
		t.Fatal("clone modified source")
	}
	u, err := url.Parse("https://user:example@example.test/path?q=test")
	if err != nil {
		t.Fatal(err)
	}
	cloned := compat.CloneURL(u)
	if u.String() != cloned.String() || u == cloned || u.User == cloned.User {
		t.Fatal("URL copy did not preserve and isolate its fields")
	}
}
