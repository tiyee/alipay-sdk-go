package generatedapi_test

import (
	"context"
	"errors"
	"net/url"
	"testing"

	"github.com/tiyee/alipay-sdk-go/core"
	"github.com/tiyee/alipay-sdk-go/services/commerce/ec/employee"
	"github.com/tiyee/alipay-sdk-go/services/iservice/ccm/sw/treecategory"
	"github.com/tiyee/alipay-sdk-go/services/open/mini/category"
)

type queryRecorder struct{ form url.Values }

func (r *queryRecorder) Do(_ context.Context, req core.WireRequest) (*core.APIResult, error) {
	r.form = req.Params.Form
	return &core.APIResult{SignatureVerified: true, Node: []byte("{}")}, nil
}

func TestNumericAndBooleanServiceQueries(t *testing.T) {
	r := &queryRecorder{}
	employees, err := employee.NewService(r)
	if err != nil {
		t.Fatal(err)
	}
	page := int64(1)
	_, err = employees.EcEmployeeIDListQuery(t.Context(), employee.EcEmployeeIDListQueryRequest{PageNum: &page})
	var missing *core.ValidationError
	if !errors.As(err, &missing) || missing.Field != "page_size" || r.form != nil {
		t.Fatalf("missing size: %v, form %v", err, r.form)
	}
	size := int64(20)
	_, err = employees.EcEmployeeIDListQuery(t.Context(), employee.EcEmployeeIDListQueryRequest{PageNum: &page, PageSize: &size})
	if err != nil || r.form.Encode() != "page_num=1&page_size=20" {
		t.Fatalf("pagination: %v, %v", r.form, err)
	}
	cat, err := category.NewService(r)
	if err != nil {
		t.Fatal(err)
	}
	flag := false
	_, err = cat.MiniCategoryQuery(t.Context(), category.MiniCategoryQueryRequest{IsFilter: &flag})
	if err != nil || r.form.Encode() != "is_filter=false" {
		t.Fatalf("explicit false: %v, %v", r.form, err)
	}
}

func TestTreeCategoryRequiresPairedPageParameters(t *testing.T) {
	r := &queryRecorder{}
	tree, err := treecategory.NewService(r)
	if err != nil {
		t.Fatal(err)
	}
	v := int64(1)
	for _, req := range []treecategory.CcmSwTreeCategoryBatchQueryRequest{{PageNum: &v}, {PageSize: &v}} {
		_, err = tree.CcmSwTreeCategoryBatchQuery(t.Context(), req)
		var missing *core.ValidationError
		if !errors.As(err, &missing) || r.form != nil {
			t.Fatalf("unpaired pagination: %v, %v", err, r.form)
		}
	}
}
