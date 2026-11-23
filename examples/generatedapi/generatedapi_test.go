package generatedapi_test

import (
	"context"
	"fmt"

	"github.com/tiyee/alipay-sdk-go/core"
	activityquery "github.com/tiyee/alipay-sdk-go/services/marketing/activity/query"
	miniinner "github.com/tiyee/alipay-sdk-go/services/open/mini/inner"
)

type exampleExecutor struct{}

func (exampleExecutor) Do(_ context.Context, req core.WireRequest) (*core.APIResult, error) {
	pathValue := req.Params.PathParams["activity_id"]
	if pathValue != "" {
		fmt.Printf("%s path activity_id=%s\n", req.Spec.Method, pathValue)
	} else {
		fmt.Printf("%s %s\n", req.Spec.Method, req.Spec.Path)
	}
	return &core.APIResult{SignatureVerified: true, Node: []byte(`{}`)}, nil
}

func Example_generatedServices() {
	ctx := context.Background()
	innerService, _ := miniinner.NewService(exampleExecutor{})
	_, _ = innerService.MiniInnerMembersQuery(ctx, miniinner.MiniInnerMembersQueryRequest{})

	activityService, _ := activityquery.NewService(exampleExecutor{})
	_, _ = activityService.ActivityQuery(ctx, activityquery.ActivityQueryRequest{ActivityId: "activity-123"})

	// Output:
	// alipay.open.mini.inner.members.query /v3/alipay/open/mini/inner/members/query
	// alipay.marketing.activity.query path activity_id=activity-123
}
