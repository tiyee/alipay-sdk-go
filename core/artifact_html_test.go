package core_test

import (
	"encoding/xml"
	"html"
	"io"
	"net/url"
	"strings"
	"testing"

	"github.com/tiyee/alipay-sdk-go/core"
	"github.com/tiyee/alipay-sdk-go/core/option"
	"github.com/tiyee/alipay-sdk-go/internal/testutil"
)

func TestFormHTMLSubmitsSignedValues(t *testing.T) {
	b, err := core.NewArtifactBuilder(testutil.NewCredentials(t), option.WithAppID("offline"), option.WithEnvironment(core.EnvironmentSandbox))
	if err != nil {
		t.Fatal(err)
	}
	const field = `a"><script>marker</script><input name="b`
	fields := url.Values{field: {"a'&<b>"}, "return_url": {"https://example.test/return?a=1&amp;b=2"}}
	a, err := b.BuildArtifact(t.Context(), "example.page", fields, core.CallSettings{})
	if err != nil {
		t.Fatal(err)
	}
	document := a.FormHTML()
	if strings.Contains(document, "<script>marker") {
		t.Fatal("field name produced markup")
	}
	decoder := xml.NewDecoder(strings.NewReader(document))
	decoder.Strict = false
	values := url.Values{}
	for {
		token, err := decoder.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatal(err)
		}
		el, ok := token.(xml.StartElement)
		if !ok {
			continue
		}
		attrs := map[string]string{}
		for _, a := range el.Attr {
			attrs[a.Name.Local] = a.Value
		}
		switch el.Name.Local {
		case "form":
			u, err := url.Parse(attrs["action"])
			if err != nil {
				t.Fatal(err)
			}
			values = u.Query()
		case "input":
			if attrs["type"] == "hidden" {
				values.Add(attrs["name"], attrs["value"])
			}
		}
	}
	if values.Encode() != a.Values.Encode() {
		t.Fatal("HTML submission differs from signed values")
	}
	if !strings.Contains(document, html.EscapeString(field)) {
		t.Fatal("missing escaped field")
	}
}
