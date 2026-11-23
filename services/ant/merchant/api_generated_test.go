package merchant

import (
	"context"
	"github.com/tiyee/alipay-sdk-go/core"
	"testing"
)

type recordingExecutor struct{ req core.WireRequest }

func (e *recordingExecutor) Do(_ context.Context, r core.WireRequest) (*core.APIResult, error) {
	e.req = r
	return &core.APIResult{HTTPStatus: 200, SignatureVerified: true, Node: []byte(`{}`)}, nil
}
func TestOperationSpecs(t *testing.T) {
	cases := []struct {
		operation, method, path string
		mode                    core.ParameterMode
		call                    func(*Service) error
	}{
		{"ant.merchant.expand.apprecommend.account.create", "POST", "/v3/ant/merchant/expand/apprecommend/account/create", core.ParameterModeJSONBody, func(s *Service) error {
			_, err := s.ExpandAppRecommendAccountCreate(context.Background(), ExpandAppRecommendAccountCreateRequest{})
			return err
		}},
		{"ant.merchant.expand.apprecommend.account.delete", "DELETE", "/v3/ant/merchant/expand/apprecommend/account/delete", core.ParameterModeQuery, func(s *Service) error {
			_, err := s.ExpandAppRecommendAccountDelete(context.Background(), ExpandAppRecommendAccountDeleteRequest{})
			return err
		}},
		{"ant.merchant.expand.apprecommend.account.query", "GET", "/v3/ant/merchant/expand/apprecommend/account/query", core.ParameterModeQuery, func(s *Service) error {
			_, err := s.ExpandAppRecommendAccountQuery(context.Background(), ExpandAppRecommendAccountQueryRequest{})
			return err
		}},
		{"ant.merchant.expand.apprecommend.available.query", "GET", "/v3/ant/merchant/expand/apprecommend/available/query", core.ParameterModeQuery, func(s *Service) error {
			_, err := s.ExpandAppRecommendAvailableQuery(context.Background(), ExpandAppRecommendAvailableQueryRequest{})
			return err
		}},
		{"ant.merchant.expand.info.query", "GET", "/v3/ant/merchant/expand/info/query", core.ParameterModeQuery, func(s *Service) error {
			_, err := s.ExpandInfoQuery(context.Background(), ExpandInfoQueryRequest{})
			return err
		}},
		{"ant.merchant.expand.item.create", "POST", "/v3/ant/merchant/expand/item/create", core.ParameterModeJSONBody, func(s *Service) error {
			_, err := s.ExpandItemCreate(context.Background(), ExpandItemCreateRequest{})
			return err
		}},
		{"ant.merchant.expand.item.delete", "DELETE", "/v3/ant/merchant/expand/item/delete", core.ParameterModeQuery, func(s *Service) error {
			_, err := s.ExpandItemDelete(context.Background(), ExpandItemDeleteRequest{})
			return err
		}},
		{"ant.merchant.expand.item.modify", "POST", "/v3/ant/merchant/expand/item/modify", core.ParameterModeJSONBody, func(s *Service) error {
			_, err := s.ExpandItemModify(context.Background(), ExpandItemModifyRequest{})
			return err
		}},
		{"ant.merchant.expand.item.open.batchquery", "POST", "/v3/ant/merchant/expand/item/open/batchquery", core.ParameterModeJSONBody, func(s *Service) error {
			_, err := s.ExpandItemOpenBatchQuery(context.Background(), ExpandItemOpenBatchQueryRequest{})
			return err
		}},
		{"ant.merchant.expand.item.open.create", "POST", "/v3/ant/merchant/expand/item/open/create", core.ParameterModeJSONBody, func(s *Service) error {
			_, err := s.ExpandItemOpenCreate(context.Background(), ExpandItemOpenCreateRequest{})
			return err
		}},
		{"ant.merchant.expand.item.open.delete", "DELETE", "/v3/ant/merchant/expand/item/open/delete", core.ParameterModeQuery, func(s *Service) error {
			_, err := s.ExpandItemOpenDelete(context.Background(), ExpandItemOpenDeleteRequest{})
			return err
		}},
		{"ant.merchant.expand.item.open.modify", "POST", "/v3/ant/merchant/expand/item/open/modify", core.ParameterModeJSONBody, func(s *Service) error {
			_, err := s.ExpandItemOpenModify(context.Background(), ExpandItemOpenModifyRequest{})
			return err
		}},
		{"ant.merchant.expand.item.open.query", "GET", "/v3/ant/merchant/expand/item/open/query", core.ParameterModeQuery, func(s *Service) error {
			_, err := s.ExpandItemOpenQuery(context.Background(), ExpandItemOpenQueryRequest{})
			return err
		}},
		{"ant.merchant.expand.item.query", "GET", "/v3/ant/merchant/expand/item/query", core.ParameterModeQuery, func(s *Service) error {
			_, err := s.ExpandItemQuery(context.Background(), ExpandItemQueryRequest{})
			return err
		}},
		{"ant.merchant.expand.item.status.modify", "POST", "/v3/ant/merchant/expand/item/status/modify", core.ParameterModeJSONBody, func(s *Service) error {
			_, err := s.ExpandItemStatusModify(context.Background(), ExpandItemStatusModifyRequest{})
			return err
		}},
		{"ant.merchant.expand.shop.consult", "POST", "/v3/ant/merchant/expand/shop/consult", core.ParameterModeJSONBody, func(s *Service) error {
			_, err := s.ExpandShopConsult(context.Background(), ExpandShopConsultRequest{})
			return err
		}},
		{"ant.merchant.expand.shop.receiptaccount.save", "POST", "/v3/ant/merchant/expand/shop/receiptaccount/save", core.ParameterModeJSONBody, func(s *Service) error {
			_, err := s.ExpandShopReceiptAccountSave(context.Background(), ExpandShopReceiptAccountSaveRequest{})
			return err
		}},
		{"ant.merchant.expand.indirect.image.upload", "POST", "/v3/ant/merchant/indirect/image", core.ParameterModeMultipart, func(s *Service) error {
			_, err := s.ExpandIndirectImageUpload(context.Background(), ExpandIndirectImageUploadRequest{})
			return err
		}},
		{"ant.merchant.expand.mcc.query", "GET", "/v3/ant/merchant/mcc/query", core.ParameterModeQuery, func(s *Service) error {
			_, err := s.ExpandMCCQuery(context.Background(), ExpandMCCQueryRequest{})
			return err
		}},
		{"ant.merchant.expand.order.query", "GET", "/v3/ant/merchant/order/{order_id}", core.ParameterModeQuery, func(s *Service) error {
			_, err := s.ExpandOrderQuery(context.Background(), ExpandOrderQueryRequest{OrderId: "path-value"})
			return err
		}},
		{"ant.merchant.expand.shop.query", "GET", "/v3/ant/merchant/shop", core.ParameterModeQuery, func(s *Service) error {
			_, err := s.ExpandShopQuery(context.Background(), ExpandShopQueryRequest{})
			return err
		}},
		{"ant.merchant.expand.shop.modify", "PATCH", "/v3/ant/merchant/shop", core.ParameterModeJSONBody, func(s *Service) error {
			_, err := s.ExpandShopModify(context.Background(), ExpandShopModifyRequest{})
			return err
		}},
		{"ant.merchant.expand.shop.create", "POST", "/v3/ant/merchant/shop", core.ParameterModeJSONBody, func(s *Service) error {
			_, err := s.ExpandShopCreate(context.Background(), ExpandShopCreateRequest{})
			return err
		}},
		{"ant.merchant.expand.shop.close", "PATCH", "/v3/ant/merchant/shop/close", core.ParameterModeJSONBody, func(s *Service) error {
			_, err := s.ExpandShopClose(context.Background(), ExpandShopCloseRequest{})
			return err
		}},
		{"ant.merchant.expand.shop.page.query", "GET", "/v3/ant/merchant/shop/pagequery", core.ParameterModeQuery, func(s *Service) error {
			_, err := s.ExpandShopPageQuery(context.Background(), ExpandShopPageQueryRequest{})
			return err
		}}}
	for _, tc := range cases {
		t.Run(tc.operation, func(t *testing.T) {
			e := &recordingExecutor{}
			s, err := NewService(e)
			if err != nil {
				t.Fatal(err)
			}
			if err = tc.call(s); err != nil {
				t.Fatal(err)
			}
			if e.req.Spec.Method != tc.operation || e.req.Spec.HTTPMethod.String() != tc.method || e.req.Spec.Path != tc.path || e.req.Spec.ParameterMode != tc.mode || e.req.Params.Mode != tc.mode {
				t.Fatalf("unexpected spec: %+v", e.req.Spec)
			}
		})
	}
}
