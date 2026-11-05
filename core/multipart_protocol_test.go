package core_test

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/tiyee/alipay-sdk-go/core"
	"github.com/tiyee/alipay-sdk-go/internal/testutil"
	"github.com/tiyee/alipay-sdk-go/services/ant/merchant"
)

func TestMultipartSignsJSONPartAndSendsFile(t *testing.T) {
	gateway := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseMultipartForm(1024 * 1024); err != nil {
			t.Error(err)
			return
		}
		data := r.FormValue("data")
		testutil.VerifyV3Authorization(t, r, []byte(data), "")
		if data != `{"image_type":"test"}` {
			t.Errorf("data = %s", data)
		}
		file, header, err := r.FormFile("image_content")
		if err != nil {
			t.Error(err)
			return
		}
		defer file.Close()
		body, err := io.ReadAll(file)
		if err != nil || string(body) != "image bytes" || header.Filename != "test.png" {
			t.Errorf("file = %q, header = %+v, error = %v", body, header, err)
		}
		response := []byte(`{"image_id":"image-1"}`)
		for key, value := range testutil.V3ResponseHeaders(t, response, "1764000000000", "response-nonce", "") {
			w.Header()[key] = value
		}
		_, _ = w.Write(response)
	}))
	defer gateway.Close()
	client := newTestClient(t, gateway.URL)
	service, err := merchant.NewService(client)
	if err != nil {
		t.Fatal(err)
	}
	imageType := "test"
	result, err := service.ExpandIndirectImageUpload(context.Background(), merchant.ExpandIndirectImageUploadRequest{
		Data:         &merchant.ExpandIndirectImageUpload{ImageType: &imageType},
		ImageContent: &core.MultipartFile{Filename: "test.png", ContentType: "image/png", Content: []byte("image bytes")},
	})
	if err != nil {
		t.Fatal(err)
	}
	if result.Data.ImageId == nil || *result.Data.ImageId != "image-1" {
		t.Fatalf("response = %+v", result.Data)
	}
}
