package merchant

import (
	"context"
	"testing"
)

func TestQueryParametersAndResponse(t *testing.T) {
	e := &recordingExecutor{}
	s, err := NewService(e)
	if err != nil {
		t.Fatal(err)
	}
	flag := false
	_, err = s.ExpandItemQuery(context.Background(), ExpandItemQueryRequest{UndefinedCategory: &flag})
	if err != nil {
		t.Fatal(err)
	}
	if got := e.req.Params.Form.Encode(); got != "undefined_category=false" {
		t.Fatalf("query = %q", got)
	}

	zero := int64(0)
	_, err = s.ExpandAppRecommendAccountQuery(context.Background(), ExpandAppRecommendAccountQueryRequest{PageSize: &zero})
	if err != nil {
		t.Fatal(err)
	}
	if got := e.req.Params.Form.Encode(); got != "page_size=0" {
		t.Fatalf("query = %q", got)
	}

	source := "APLUS_SHOP"
	_, err = s.ExpandOrderQuery(context.Background(), ExpandOrderQueryRequest{OrderId: "order/1", BizSource: &source})
	if err != nil {
		t.Fatal(err)
	}
	if e.req.Params.PathParams["order_id"] != "order/1" || e.req.Params.Form.Get("biz_source") != source {
		t.Fatalf("parameters = %+v", e.req.Params)
	}
}
