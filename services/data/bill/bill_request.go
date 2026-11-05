package bill

import (
	"github.com/tiyee/alipay-sdk-go/core"
)

// BillDownloadURLRequest 是 alipay.data.dataservice.bill.downloadurl.query
// （对账单下载地址查询，V3 GET）的请求。
type BillDownloadURLRequest struct {
	// BillType 是账单类型（如 trade、signcustomer）；
	// 取值依据各产品对账单说明。
	BillType string `query:"bill_type"`
	// BillDate 是账单时间：日账单 yyyy-MM-dd（T+1，不支持当日），
	// 月账单 yyyy-MM（次月 3 日生成）。
	BillDate string `query:"bill_date"`
	// SMID 是二级商户 smid；仅 bill_type=trade_zft_merchant 时可用。
	SMID *string `query:"smid,omitempty"`
	// Secure 标识返回 https 安全链接；非 true 使用 http 链接。
	Secure *string `query:"secure,omitempty"`
}

// BillBalanceRequest 是 alipay.data.bill.balance.query（当前余额查询，
// V3 GET）的请求；查询自身账户时无需传 bill_user_id。
type BillBalanceRequest struct {
	// BillUserID 是目标查询账户；仅支持部分场景。
	BillUserID *string `query:"bill_user_id,omitempty"`
}

// BillBalanceHisRequest 是 alipay.data.bill.balancehis.query
// （历史余额查询，V3 GET）的请求：biz_date 与 biz_month 互斥，
// 传入其一。
type BillBalanceHisRequest struct {
	// BizDate 是查询指定自然日（yyyy-MM-dd）的期初与期末余额。
	BizDate *string `query:"biz_date,omitempty"`
	// BizMonth 是查询指定自然月（yyyy-MM）的期初与期末余额。
	BizMonth *string `query:"biz_month,omitempty"`
}

// BillAccountLogRequest 是 alipay.data.bill.accountlog.query
// （账务明细查询，V3 GET）的请求：时间范围与订单号二选一使用，
// 支付宝订单号与商户订单号互斥。
type BillAccountLogRequest struct {
	// StartTime 是账务流水创建时间的起始范围；只能查询一年内的记录。
	StartTime *string `query:"start_time,omitempty"`
	// EndTime 是账务流水创建时间的结束范围；与起始时间间隔不超过 31 天，
	// 查询结果为左闭右开区间。
	EndTime *string `query:"end_time,omitempty"`
	// AlipayOrderNo 是支付宝订单号精确查询；与 MerchantOrderNo 互斥。
	AlipayOrderNo *string `query:"alipay_order_no,omitempty"`
	// MerchantOrderNo 是商户订单号精确查询；与 AlipayOrderNo 互斥。
	MerchantOrderNo *string `query:"merchant_order_no,omitempty"`
	// PageNo 是分页号，从 1 开始。
	PageNo *string `query:"page_no,omitempty"`
	// PageSize 是分页大小（1000-2000），默认 2000。
	PageSize *string `query:"page_size,omitempty"`
	// TransCode 是账务的类型代码；特殊场景下使用。
	TransCode *string `query:"trans_code,omitempty"`
	// AgreementNo 是协议授权码；特殊场景下使用。
	AgreementNo *string `query:"agreement_no,omitempty"`
	// AgreementProductCode 是协议产品码；特殊场景下使用。
	AgreementProductCode *string `query:"agreement_product_code,omitempty"`
	// BillUserID 是指定用户做账单查询。
	BillUserID *string `query:"bill_user_id,omitempty"`
	// OpenID 是用户在该应用下的唯一标识。
	OpenID *string `query:"open_id,omitempty"`
}

func (r BillDownloadURLRequest) validate() error {
	if r.BillType == "" {
		return &core.ValidationError{Field: "bill_type", Reason: "bill_type is required"}
	}
	if r.BillDate == "" {
		return &core.ValidationError{Field: "bill_date", Reason: "bill_date is required"}
	}
	return nil
}

func (r BillBalanceHisRequest) validate() error {
	hasDate := r.BizDate != nil && *r.BizDate != ""
	hasMonth := r.BizMonth != nil && *r.BizMonth != ""
	switch {
	case hasDate && hasMonth:
		return &core.ValidationError{Field: "biz_date",
			Reason: "biz_date and biz_month are mutually exclusive"}
	case !hasDate && !hasMonth:
		return &core.ValidationError{Field: "biz_date",
			Reason: "one of biz_date or biz_month is required"}
	}
	return nil
}

func (r BillAccountLogRequest) validate() error {
	hasRange := (r.StartTime != nil && *r.StartTime != "") || (r.EndTime != nil && *r.EndTime != "")
	hasOrder := (r.AlipayOrderNo != nil && *r.AlipayOrderNo != "") || (r.MerchantOrderNo != nil && *r.MerchantOrderNo != "")
	if !hasRange && !hasOrder {
		return &core.ValidationError{Field: "start_time",
			Reason: "a time range or an exact order number is required"}
	}
	if r.AlipayOrderNo != nil && *r.AlipayOrderNo != "" && r.MerchantOrderNo != nil && *r.MerchantOrderNo != "" {
		return &core.ValidationError{Field: "alipay_order_no",
			Reason: "alipay_order_no and merchant_order_no are mutually exclusive"}
	}
	return nil
}

// DataBillAccountBookEReceiptApplyRequest 是支付宝官方 V3 OpenAPI 模型。
type DataBillAccountBookEReceiptApplyRequest struct {
	// AgreementNo 是协议号。
	AgreementNo *string `json:"agreement_no,omitempty"`
	// Key 是明细凭证，传入流水号（转账接口查询结果）。汇总凭证，传入起止时间，格式yyyy-MM-dd_yyyy-MM-dd。。
	Key *string `json:"key,omitempty"`
	// StoreNo 是子账本号，或者子账本名称。模糊查询。
	StoreNo *string `json:"store_no,omitempty"`
	// Type 是申请的类型。可传入：FUND_PLATFORM_DETAIL - 资金明细证明，FUND_PLATFORM_SUM - 资金汇总证明。
	Type *string `json:"type,omitempty"`
}

// DataBillEReceiptApplyRequest 是支付宝官方 V3 OpenAPI 模型。
type DataBillEReceiptApplyRequest struct {
	// BillUserId 是授权商户pid。需调用方获取商户授权，才可使用。。
	BillUserId *string `json:"bill_user_id,omitempty"`
	// Key 是根据不同业务类型，传入不同参数。 格式如下： * ACCOUNT_LOG_DETAIL - 传入账务流水号，示例：117007*********9151 * ACCOUNT_FLOW_DETAIL - 传入流水日期区间示例：20220301-20220331 * FUND_DETAIL - 传入资金单据号，示例：2022051511007***************0731 * ACCOUNT_LOG_SUM_DAILY - 传入日期，示例：2022-01-01 * ACCOUNT_LOG_SUM_MONTHLY - 传入月份，示例：2022-01。
	Key *string `json:"key,omitempty"`
	// Type 是申请的类型。可传入： * ACCOUNT_LOG_DETAIL - 余额收支证明 * ACCOUNT_FLOW_DETAIL - 余额收支流水证明 * FUND_DETAIL - 转入转出收支证明 * ACCOUNT_LOG_SUM_DAILY - 收支汇总证明（日汇总） * ACCOUNT_LOG_SUM_MONTHLY - 收支汇总证明（月汇总）。
	Type *string `json:"type,omitempty"`
}

// DataBillEReceiptAgentApplyRequest 是支付宝官方 V3 OpenAPI 模型。
type DataBillEReceiptAgentApplyRequest struct {
	// AgreementNo 是ISV与商户签约授权协议号。
	AgreementNo *string `json:"agreement_no,omitempty"`
	// AgreementType 是根据不同业务协议类型，传入不同参数。传入协议产品码（personal_product_code，通过协议查询接口、协议签约通知响应参数获取），不填的话默认按照示例值传入。 安全发协议：FUND_SAFT_SIGN_WITHHOLDING_P 专用金协议：FLEXIBLE_EMPLOYMENT_SERVICE_P 批量有密：FUND_AUTHORIZE_FLOW。
	AgreementType *string `json:"agreement_type,omitempty"`
	// EndTime 是查询时间区间的结束时间，属于闭区间。 [start_time,end_time]。
	EndTime *string `json:"end_time,omitempty"`
	// Key 是根据不同业务类型，传入不同参数。格式如下： * FUND_DETAIL - 资金业务回单。传入资金单号(fundOrderId，可通过充值、转账、提现接口获取)。
	Key *string `json:"key,omitempty"`
	// StartTime 是搜索范围的起始时间。
	StartTime *string `json:"start_time,omitempty"`
	// Type 是申请的类型（目前仅支持资金业务回单）。可传入： * FUND_DETAIL - 资金业务回单。。
	Type *string `json:"type,omitempty"`
}

// DataBillAccountBookEReceiptQueryRequest 是 alipay.data.bill.accountbookereceipt.query 的 V3 query 请求参数。
type DataBillAccountBookEReceiptQueryRequest struct {
	// FileId 是根据申请id查询状态。申请接口可以参考alipay.data.bill.ereceiptagent.apply。
	FileId *string `query:"file_id,omitempty"`
	// AgreementNo 是协议号，根据不同业务协议类型，传入对应类型的协议号，用于isv授权检查并获取商户信息。如果业务类型未指定，则使用默认类型对应的协议号。。
	AgreementNo *string `query:"agreement_no,omitempty"`
	// AgreementType 是根据不同业务协议类型，传入不同参数。传入协议产品码（personal_product_code，通过协议查询接口、协议签约通知响应参数获取），不填的话默认按照示例值传入。
	AgreementType *string `query:"agreement_type,omitempty"`
	// Secure 是true表示使用安全链接，即返回的下载链接为https。 非true值的情况下，统一为false，使用http链接。。
	Secure *string `query:"secure,omitempty"`
}

// DataBillBailQueryRequest 是 alipay.data.bill.bail.query 的 V3 query 请求参数。
type DataBillBailQueryRequest struct {
	// StartTime 是保证金流水创建时间的起始范围。
	StartTime *string `query:"start_time,omitempty"`
	// EndTime 是保证金流水创建时间的结束范围。与起始时间间隔不超过31天。查询结果为起始时间至结束时间的左闭右开区间。
	EndTime *string `query:"end_time,omitempty"`
	// BailType 是保证金类型，目前支持TMALL_S_BAIL-天猫保证金，TAOBAO_BAIL-淘宝保证金。
	BailType *string `query:"bail_type,omitempty"`
	// TransLogId 是保证金流水号。如果查询参数中指定流水号，则只查询流水号相关的记录。
	TransLogId *string `query:"trans_log_id,omitempty"`
	// BizOrigNo 是业务基础订单号。如果查询参数中指定订单号，则只查询相关的记录。
	BizOrigNo *string `query:"biz_orig_no,omitempty"`
}

// DataBillBizFundAgentQueryRequest 是 alipay.data.bill.bizfundagent.query 的 V3 query 请求参数。
type DataBillBizFundAgentQueryRequest struct {
	// StartTime 是账单查询时间范围 -- 起始时间。
	StartTime *string `query:"start_time,omitempty"`
	// EndTime 是账单查询时间范围 -- 结束范围。时间范围最大不超过31天。。
	EndTime *string `query:"end_time,omitempty"`
	// PageNo 是页码，从1开始。
	PageNo *string `query:"page_no,omitempty"`
	// PageSize 是分页大小1000-2000，默认2000。
	PageSize *string `query:"page_size,omitempty"`
	// AgreementType 是根据不同业务协议类型，传入不同参数。传入协议产品码（personal_product_code，通过协议查询接口、协议签约通知响应参数获取），不填的话默认按照示例值传入。 安全发协议：FUND_SAFT_SIGN_WITHHOLDING_P 专用金协议：FLEXIBLE_EMPLOYMENT_SERVICE_P。
	AgreementType *string `query:"agreement_type,omitempty"`
	// AgreementNo 是ISV与商户授权协议号。
	AgreementNo *string `query:"agreement_no,omitempty"`
	// AccountBookId 是记账本id，服务商在帮助商户开通时候获取的唯一身份号。
	AccountBookId *string `query:"account_book_id,omitempty"`
}

// DataBillBuyQueryRequest 是 alipay.data.bill.buy.query 的 V3 query 请求参数。
type DataBillBuyQueryRequest struct {
	// StartTime 是交易流水创建时间的起始范围。
	StartTime *string `query:"start_time,omitempty"`
	// EndTime 是交易流水创建时间的结束范围。与起始时间间隔不超过31天。查询结果为起始时间至结束时间的左闭右开区间。
	EndTime *string `query:"end_time,omitempty"`
	// AlipayOrderNo 是支付宝交易流水号。如果查询参数中指定流水号，则只查询流水号相关的记录。
	AlipayOrderNo *string `query:"alipay_order_no,omitempty"`
	// MerchantOrderNo 是商户交易号。如果查询参数中指定交易号，则只查询相关的记录。
	MerchantOrderNo *string `query:"merchant_order_no,omitempty"`
	// StoreNo 是门店编号，模糊搜索。
	StoreNo *string `query:"store_no,omitempty"`
	// PageNo 是分页号，从1开始。
	PageNo *string `query:"page_no,omitempty"`
	// PageSize 是分页大小1000-2000，默认2000。
	PageSize *string `query:"page_size,omitempty"`
}

// DataBillEReceiptQueryRequest 是 alipay.data.bill.ereceipt.query 的 V3 query 请求参数。
type DataBillEReceiptQueryRequest struct {
	// FileId 是根据申请id查询状态。通过 <a href="https://opendocs.alipay.com/apis/api_15/alipay.data.bill.ereceipt.apply">alipay.data.bill.ereceipt.apply(申请电子回单(incubating))</a> 接口同步响应获取。。
	FileId *string `query:"file_id,omitempty"`
}

// DataBillSellQueryRequest 是 alipay.data.bill.sell.query 的 V3 query 请求参数。
type DataBillSellQueryRequest struct {
	// StartTime 是卖出交易流水创建时间的起始范围。
	StartTime *string `query:"start_time,omitempty"`
	// EndTime 是卖出交易流水创建时间的结束范围。与起始时间间隔不超过31天。查询结果为起始时间至结束时间的左闭右开区间。
	EndTime *string `query:"end_time,omitempty"`
	// AlipayOrderNo 是支付宝交易流水号。如果查询参数中指定流水号，则只查询流水号相关的记录。
	AlipayOrderNo *string `query:"alipay_order_no,omitempty"`
	// MerchantOrderNo 是商户交易号。如果查询参数中指定交易号，则只查询相关的记录。
	MerchantOrderNo *string `query:"merchant_order_no,omitempty"`
	// StoreNo 是门店编号，模糊搜索。
	StoreNo *string `query:"store_no,omitempty"`
	// PageNo 是分页号，从1开始。
	PageNo *string `query:"page_no,omitempty"`
	// PageSize 是分页大小1000-2000，默认2000。
	PageSize *string `query:"page_size,omitempty"`
}

// DataBillTransferQueryRequest 是 alipay.data.bill.transfer.query 的 V3 query 请求参数。
type DataBillTransferQueryRequest struct {
	// StartTime 是充值、转账、提现流水业务时间的起始范围。
	StartTime *string `query:"start_time,omitempty"`
	// EndTime 是充值、转账、提现流水业务时间的结束范围。与起始时间间隔不超过31天。查询结果为起始时间至结束时间的左闭右开区间。
	EndTime *string `query:"end_time,omitempty"`
	// Type 是转账类型：充值-DEPOSIT，提现-WITHDRAW，转账-TRANSFER。。
	Type *string `query:"type,omitempty"`
	// PageNo 是分页号，从1开始。
	PageNo *string `query:"page_no,omitempty"`
	// PageSize 是分页大小1000-2000，默认2000。
	PageSize *string `query:"page_size,omitempty"`
}

// DataBillTransferAccountBookQueryRequest 是 alipay.data.bill.transferaccountbook.query 的 V3 query 请求参数。
type DataBillTransferAccountBookQueryRequest struct {
	// StartTime 是充值、转账、提现流水业务时间的起始范围。
	StartTime *string `query:"start_time,omitempty"`
	// EndTime 是充值、转账、提现流水业务时间的结束范围。与起始时间间隔不超过31天。查询结果为起始时间至结束时间的左闭右开区间。
	EndTime *string `query:"end_time,omitempty"`
	// Type 是转账类型：充值-DEPOSIT，提现-WITHDRAW，转账-TRANSFER。。
	Type *string `query:"type,omitempty"`
	// AgreementNo 是协议号。
	AgreementNo *string `query:"agreement_no,omitempty"`
	// StoreNo 是子账本号，或者子账本名称。模糊查询。
	StoreNo *string `query:"store_no,omitempty"`
	// PageNo 是分页号，从1开始。
	PageNo *string `query:"page_no,omitempty"`
	// PageSize 是分页大小1000-2000，默认2000。
	PageSize *string `query:"page_size,omitempty"`
}
