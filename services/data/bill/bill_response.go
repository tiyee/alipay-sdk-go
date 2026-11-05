package bill

// BillDownloadURLResponse 是对账单下载地址查询的业务响应。
// 下载地址需在获取后 30 秒内使用；链接缺失按业务状态处理，
// 以 bill_file_code 为准。
type BillDownloadURLResponse struct {
	// BillDownloadURL 是账单下载地址；获取链接后 30 秒未下载即失效。
	BillDownloadURL *string `json:"bill_download_url,omitempty"`
	// BillFileCode 描述本次申请的账单文件状态：
	// EMPTY_DATA_WITH_BILL_FILE（无数据但可获取空数据账单文件）等。
	BillFileCode *BillFileCode `json:"bill_file_code,omitempty"`
}

// BillBalanceResponse 是当前余额查询的业务响应。
type BillBalanceResponse struct {
	// TotalAmount 是支付宝账户余额。
	TotalAmount *string `json:"total_amount,omitempty"`
	// AvailableAmount 是账户可用余额。
	AvailableAmount *string `json:"available_amount,omitempty"`
	// FreezeAmount 是冻结金额。
	FreezeAmount *string `json:"freeze_amount,omitempty"`
	// SettleAmount 是待结算金额。
	SettleAmount *string `json:"settle_amount,omitempty"`
}

// BillBalanceHisResponse 是历史余额查询的业务响应。
type BillBalanceHisResponse struct {
	// BeginningBalance 是期初余额。
	BeginningBalance *string `json:"beginning_balance,omitempty"`
	// EndingBalance 是期末余额。
	EndingBalance *string `json:"ending_balance,omitempty"`
}

// BillAccountLogResponse 是账务明细查询的业务响应。
type BillAccountLogResponse struct {
	// DetailList 是账务明细；查询为空时不返回。
	DetailList []AccountLogItem `json:"detail_list,omitempty"`
	// TotalSize 是满足查询条件的明细总数。
	TotalSize *string `json:"total_size,omitempty"`
	// PageNo 是分页号。
	PageNo *string `json:"page_no,omitempty"`
	// PageSize 是分页大小。
	PageSize *string `json:"page_size,omitempty"`
}

// DataBillAccountBookEReceiptApplyResponse 是支付宝官方 V3 OpenAPI 模型。
type DataBillAccountBookEReceiptApplyResponse struct {
	// FileId 是文件申请号file_id信息。使用file_id可以查询处理状态。
	FileId *string `json:"file_id,omitempty"`
}

// DataBillAccountBookEReceiptQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type DataBillAccountBookEReceiptQueryResponse struct {
	// DownloadUrl 是下载链接。status为SUCCESS时返回。用户可以使用此http链接下载文件内容。有效时间20s。 生成的文件是zip格式。需要解压后获取电子回单pdf内容。
	DownloadUrl *string `json:"download_url,omitempty"`
	// ErrorMessage 是如果生成失败，则会返回失败原因。
	ErrorMessage *string `json:"error_message,omitempty"`
	// Status 是处理状态 INIT - 初始化 PROCESS - 处理中 SUCCESS - 成功 FAIL - 失败。
	Status *string `json:"status,omitempty"`
}

// DataBillBailQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type DataBillBailQueryResponse struct {
	// DetailList 是保证金明细列表，最多返回5000条。
	DetailList *[]BailDetailResult `json:"detail_list,omitempty"`
}

// DataBillBizFundAgentQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type DataBillBizFundAgentQueryResponse struct {
	// DetailList 是资金明细记录。
	DetailList *[]BizFundReportResult `json:"detail_list,omitempty"`
	// PageNo 是分页号，从1开始。
	PageNo *string `json:"page_no,omitempty"`
	// PageSize 是页面大小，最大2000条。
	PageSize *string `json:"page_size,omitempty"`
	// TotalSize 是数据分页总页数。
	TotalSize *string `json:"total_size,omitempty"`
}

// DataBillBuyQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type DataBillBuyQueryResponse struct {
	// DetailList 是交易流水详情。
	DetailList *[]TradeItemResult `json:"detail_list,omitempty"`
	// PageNo 是分页号，从1开始。
	PageNo *string `json:"page_no,omitempty"`
	// PageSize 是分页大小1000-2000。
	PageSize *string `json:"page_size,omitempty"`
	// TotalSize 是账务明细总数。返回满足查询条件的明细的数量。
	TotalSize *string `json:"total_size,omitempty"`
}

// DataBillEReceiptApplyResponse 是支付宝官方 V3 OpenAPI 模型。
type DataBillEReceiptApplyResponse struct {
	// FileId 是文件申请号file_id信息。 使用file_id可以查询处理状态，有效期：2天。
	FileId *string `json:"file_id,omitempty"`
}

// DataBillEReceiptQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type DataBillEReceiptQueryResponse struct {
	// DownloadUrl 是下载链接，status 为 SUCCESS 时返回。用户可以使用此http链接下载文件内容。有效时间为 30s。 生成文件为PDF格式，下载即可获取电子回单 PDF 内容。。
	DownloadUrl *string `json:"download_url,omitempty"`
	// ErrorMessage 是如果生成失败，则会返回失败原因。
	ErrorMessage *string `json:"error_message,omitempty"`
	// Status 是处理状态。枚举值如下： INIT：初始化。 PROCESS：处理中。 SUCCESS：成功。 FAIL：失败。。
	Status *string `json:"status,omitempty"`
}

// DataBillEReceiptAgentApplyResponse 是支付宝官方 V3 OpenAPI 模型。
type DataBillEReceiptAgentApplyResponse struct {
	// FileId 是文件申请号file_id信息。使用file_id可以查询处理状态。
	FileId *string `json:"file_id,omitempty"`
}

// DataBillSellQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type DataBillSellQueryResponse struct {
	// DetailList 是交易流水详情，如果不存在记录则为空。
	DetailList *[]TradeItemResult `json:"detail_list,omitempty"`
	// PageNo 是分页号，从1开始。
	PageNo *string `json:"page_no,omitempty"`
	// PageSize 是分页大小1000-2000。
	PageSize *string `json:"page_size,omitempty"`
	// TotalSize 是账务明细总数。返回满足查询条件的明细的数量。
	TotalSize *string `json:"total_size,omitempty"`
}

// DataBillTransferQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type DataBillTransferQueryResponse struct {
	// DetailList 是充值、提现、转账明细记录。
	DetailList *[]TransferDetailResult `json:"detail_list,omitempty"`
	// PageNo 是分页号，从1开始。
	PageNo *string `json:"page_no,omitempty"`
	// PageSize 是分页大小1000-2000。
	PageSize *string `json:"page_size,omitempty"`
	// TotalSize 是充值、提现、转账明细总数。返回满足查询条件的明细的数量。
	TotalSize *string `json:"total_size,omitempty"`
}

// DataBillTransferAccountBookQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type DataBillTransferAccountBookQueryResponse struct {
	// DetailList 是充值、提现、转账明细记录。
	DetailList *[]TransferAccountBookDetailResult `json:"detail_list,omitempty"`
	// PageNo 是分页号，从1开始。
	PageNo *string `json:"page_no,omitempty"`
	// PageSize 是分页大小1000-2000。
	PageSize *string `json:"page_size,omitempty"`
	// TotalSize 是充值、提现、转账明细总数。返回满足查询条件的明细的数量。
	TotalSize *string `json:"total_size,omitempty"`
}
