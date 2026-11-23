package bill

import (
	"context"

	"github.com/tiyee/alipay-sdk-go/core"
)

// Service 封装同一业务组的支付宝 V3 接口，复用注入执行器的通信与鉴权能力。
// 零值不可用；请通过 NewService 构建。
type Service struct {
	executor        core.Executor
	billDownloadurl core.Operation[BillDownloadURLRequest, BillDownloadURLResponse]
	billBalance     core.Operation[BillBalanceRequest, BillBalanceResponse]
	billBalancehis  core.Operation[BillBalanceHisRequest, BillBalanceHisResponse]
	billAccountlog  core.Operation[BillAccountLogRequest, BillAccountLogResponse]
}

// NewService 使用共享的 V3 执行器构建 Service；executor 为 nil 时返回错误。
func NewService(executor core.Executor) (*Service, error) {
	if executor == nil {
		return nil, &core.ValidationError{Field: "executor", Reason: "executor is required"}
	}
	billDownloadurl, err := core.NewJSONOperation[BillDownloadURLRequest, BillDownloadURLResponse](core.OperationSpec{
		Method: "alipay.data.dataservice.bill.downloadurl.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/data/dataservice/bill/downloadurl/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	})
	if err != nil {
		return nil, err
	}
	billBalance, err := core.NewJSONOperation[BillBalanceRequest, BillBalanceResponse](core.OperationSpec{
		Method: "alipay.data.bill.balance.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/data/bill/balance/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	})
	if err != nil {
		return nil, err
	}
	billBalancehis, err := core.NewJSONOperation[BillBalanceHisRequest, BillBalanceHisResponse](core.OperationSpec{
		Method: "alipay.data.bill.balancehis.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/data/bill/balancehis/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	})
	if err != nil {
		return nil, err
	}
	billAccountlog, err := core.NewJSONOperation[BillAccountLogRequest, BillAccountLogResponse](core.OperationSpec{
		Method: "alipay.data.bill.accountlog.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/data/bill/accountlog/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	})
	if err != nil {
		return nil, err
	}
	return &Service{
		executor:        executor,
		billDownloadurl: billDownloadurl,
		billBalance:     billBalance,
		billBalancehis:  billBalancehis,
		billAccountlog:  billAccountlog,
	}, nil
}

// BillDownloadURLQuery 调用 alipay.data.dataservice.bill.downloadurl.query
// 获取对账单下载地址（V3 GET /v3/alipay/data/dataservice/bill/downloadurl/query）。
//
// 契约：官方 V3 文档
// https://opendocs.alipay.com/open-v3/c8d608d7_alipay.data.dataservice.bill.downloadurl.query?scene=common&pathHash=de817b49
// 下载地址 30 秒失效；获取后应立即下载并校验文件内容。
func (s *Service) BillDownloadURLQuery(ctx context.Context, req BillDownloadURLRequest,
	opts ...core.CallOption) (*core.Response[BillDownloadURLResponse], error) {
	if err := req.validate(); err != nil {
		return nil, err
	}
	return core.Execute(ctx, s.executor, s.billDownloadurl,
		core.Request[BillDownloadURLRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// BillBalance 调用 alipay.data.bill.balance.query 查询当前余额
// （V3 GET /v3/alipay/data/bill/balance/query）。
//
// 契约：官方 V3 文档
// https://opendocs.alipay.com/open-v3/9b66d9d8_alipay.data.bill.balance.query
func (s *Service) BillBalance(ctx context.Context, req BillBalanceRequest,
	opts ...core.CallOption) (*core.Response[BillBalanceResponse], error) {
	return core.Execute(ctx, s.executor, s.billBalance,
		core.Request[BillBalanceRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// BillBalanceHis 调用 alipay.data.bill.balancehis.query 查询历史余额
// （V3 GET /v3/alipay/data/bill/balancehis/query）。
//
// 契约：官方 V3 文档
// https://opendocs.alipay.com/open-v3/029a9c24_alipay.data.bill.balancehis.query
// biz_date 与 biz_month 互斥，必须传入其一。
func (s *Service) BillBalanceHis(ctx context.Context, req BillBalanceHisRequest,
	opts ...core.CallOption) (*core.Response[BillBalanceHisResponse], error) {
	if err := req.validate(); err != nil {
		return nil, err
	}
	return core.Execute(ctx, s.executor, s.billBalancehis,
		core.Request[BillBalanceHisRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// BillAccountLog 调用 alipay.data.bill.accountlog.query 查询账务明细
// （V3 GET /v3/alipay/data/bill/accountlog/query）。
//
// 契约：官方 V3 文档
// https://opendocs.alipay.com/open-v3/4cce9615_alipay.data.bill.accountlog.query
// 时间范围查询间隔不超过 31 天；精确查询使用订单号（互斥）。
func (s *Service) BillAccountLog(ctx context.Context, req BillAccountLogRequest,
	opts ...core.CallOption) (*core.Response[BillAccountLogResponse], error) {
	if err := req.validate(); err != nil {
		return nil, err
	}
	return core.Execute(ctx, s.executor, s.billAccountlog,
		core.Request[BillAccountLogRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// BillAccountBookEReceiptApply calls alipay.data.bill.accountbookereceipt.apply (POST /v3/alipay/data/bill/accountbookereceipt/apply).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/05wyde
func (s *Service) BillAccountBookEReceiptApply(ctx context.Context, req DataBillAccountBookEReceiptApplyRequest, opts ...core.CallOption) (*core.Response[DataBillAccountBookEReceiptApplyResponse], error) {
	return core.ExecuteJSON[DataBillAccountBookEReceiptApplyRequest, DataBillAccountBookEReceiptApplyResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.data.bill.accountbookereceipt.apply", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/data/bill/accountbookereceipt/apply",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[DataBillAccountBookEReceiptApplyRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// BillAccountBookEReceiptQuery calls alipay.data.bill.accountbookereceipt.query (GET /v3/alipay/data/bill/accountbookereceipt/query).
// Official V3 contract: https://opendocs.alipay.com/open-v3/6ce2ea1b_alipay.data.bill.accountbookereceipt.query
func (s *Service) BillAccountBookEReceiptQuery(ctx context.Context, req DataBillAccountBookEReceiptQueryRequest, opts ...core.CallOption) (*core.Response[DataBillAccountBookEReceiptQueryResponse], error) {
	return core.ExecuteJSON[DataBillAccountBookEReceiptQueryRequest, DataBillAccountBookEReceiptQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.data.bill.accountbookereceipt.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/data/bill/accountbookereceipt/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[DataBillAccountBookEReceiptQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// BillBailQuery calls alipay.data.bill.bail.query (GET /v3/alipay/data/bill/bail/query).
// Official V3 contract: https://opendocs.alipay.com/open-v3/4421b78a_alipay.data.bill.bail.query
func (s *Service) BillBailQuery(ctx context.Context, req DataBillBailQueryRequest, opts ...core.CallOption) (*core.Response[DataBillBailQueryResponse], error) {
	return core.ExecuteJSON[DataBillBailQueryRequest, DataBillBailQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.data.bill.bail.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/data/bill/bail/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[DataBillBailQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// BillBizFundAgentQuery calls alipay.data.bill.bizfundagent.query (GET /v3/alipay/data/bill/bizfundagent/query).
// Official V3 contract: https://opendocs.alipay.com/open-v3/9b60baa5_alipay.data.bill.bizfundagent.query
func (s *Service) BillBizFundAgentQuery(ctx context.Context, req DataBillBizFundAgentQueryRequest, opts ...core.CallOption) (*core.Response[DataBillBizFundAgentQueryResponse], error) {
	return core.ExecuteJSON[DataBillBizFundAgentQueryRequest, DataBillBizFundAgentQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.data.bill.bizfundagent.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/data/bill/bizfundagent/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[DataBillBizFundAgentQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// BillBuyQuery calls alipay.data.bill.buy.query (GET /v3/alipay/data/bill/buy/query).
// Official V3 contract: https://opendocs.alipay.com/open-v3/83504d51_alipay.data.bill.buy.query
func (s *Service) BillBuyQuery(ctx context.Context, req DataBillBuyQueryRequest, opts ...core.CallOption) (*core.Response[DataBillBuyQueryResponse], error) {
	return core.ExecuteJSON[DataBillBuyQueryRequest, DataBillBuyQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.data.bill.buy.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/data/bill/buy/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[DataBillBuyQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// BillEReceiptApply calls alipay.data.bill.ereceipt.apply (POST /v3/alipay/data/bill/ereceipt/apply).
// Official V3 contract: https://opendocs.alipay.com/open-v3/46460043_alipay.data.bill.ereceipt.apply
func (s *Service) BillEReceiptApply(ctx context.Context, req DataBillEReceiptApplyRequest, opts ...core.CallOption) (*core.Response[DataBillEReceiptApplyResponse], error) {
	return core.ExecuteJSON[DataBillEReceiptApplyRequest, DataBillEReceiptApplyResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.data.bill.ereceipt.apply", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/data/bill/ereceipt/apply",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[DataBillEReceiptApplyRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// BillEReceiptQuery calls alipay.data.bill.ereceipt.query (GET /v3/alipay/data/bill/ereceipt/query).
// Official V3 contract: https://opendocs.alipay.com/open-v3/65fde737_alipay.data.bill.ereceipt.query
func (s *Service) BillEReceiptQuery(ctx context.Context, req DataBillEReceiptQueryRequest, opts ...core.CallOption) (*core.Response[DataBillEReceiptQueryResponse], error) {
	return core.ExecuteJSON[DataBillEReceiptQueryRequest, DataBillEReceiptQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.data.bill.ereceipt.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/data/bill/ereceipt/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[DataBillEReceiptQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// BillEReceiptAgentApply calls alipay.data.bill.ereceiptagent.apply (POST /v3/alipay/data/bill/ereceiptagent/apply).
// Official V3 contract: https://opendocs.alipay.com/open-v3/d8c0f195_alipay.data.bill.ereceiptagent.apply
func (s *Service) BillEReceiptAgentApply(ctx context.Context, req DataBillEReceiptAgentApplyRequest, opts ...core.CallOption) (*core.Response[DataBillEReceiptAgentApplyResponse], error) {
	return core.ExecuteJSON[DataBillEReceiptAgentApplyRequest, DataBillEReceiptAgentApplyResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.data.bill.ereceiptagent.apply", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/data/bill/ereceiptagent/apply",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[DataBillEReceiptAgentApplyRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// BillSellQuery calls alipay.data.bill.sell.query (GET /v3/alipay/data/bill/sell/query).
// Official V3 contract: https://opendocs.alipay.com/open-v3/3c1a5835_alipay.data.bill.sell.query
func (s *Service) BillSellQuery(ctx context.Context, req DataBillSellQueryRequest, opts ...core.CallOption) (*core.Response[DataBillSellQueryResponse], error) {
	return core.ExecuteJSON[DataBillSellQueryRequest, DataBillSellQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.data.bill.sell.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/data/bill/sell/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[DataBillSellQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// BillTransferQuery calls alipay.data.bill.transfer.query (GET /v3/alipay/data/bill/transfer/query).
// Official V3 contract: https://opendocs.alipay.com/open-v3/8ddb5fed_alipay.data.bill.transfer.query
func (s *Service) BillTransferQuery(ctx context.Context, req DataBillTransferQueryRequest, opts ...core.CallOption) (*core.Response[DataBillTransferQueryResponse], error) {
	return core.ExecuteJSON[DataBillTransferQueryRequest, DataBillTransferQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.data.bill.transfer.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/data/bill/transfer/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[DataBillTransferQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// BillTransferAccountBookQuery calls alipay.data.bill.transferaccountbook.query (GET /v3/alipay/data/bill/transferaccountbook/query).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/05wycr
func (s *Service) BillTransferAccountBookQuery(ctx context.Context, req DataBillTransferAccountBookQueryRequest, opts ...core.CallOption) (*core.Response[DataBillTransferAccountBookQueryResponse], error) {
	return core.ExecuteJSON[DataBillTransferAccountBookQueryRequest, DataBillTransferAccountBookQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.data.bill.transferaccountbook.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/data/bill/transferaccountbook/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[DataBillTransferAccountBookQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
