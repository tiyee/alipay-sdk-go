// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package einvpackage

// ConsumeOutputInfo 是支付宝官方 V3 OpenAPI 模型。
type ConsumeOutputInfo struct {
	// BillNo 是支付宝交易号
	BillNo *string `json:"bill_no,omitempty"`
	// CategoryName 是账单分类
	CategoryName *string `json:"category_name,omitempty"`
	// ConsumeAmount 是金额，单位元
	ConsumeAmount *string `json:"consume_amount,omitempty"`
	// ConsumeDate 是日期
	ConsumeDate *string `json:"consume_date,omitempty"`
	// ConsumeTitle 是交易记录标题
	ConsumeTitle *string `json:"consume_title,omitempty"`
	// PayeeName 是商家名称
	PayeeName *string `json:"payee_name,omitempty"`
}

// InvoiceItemOpen 是支付宝官方 V3 OpenAPI 模型。
type InvoiceItemOpen struct {
	// ItemExTaxAmount 是明细行不含税金额，单位元，保留两位小数
	ItemExTaxAmount *string `json:"item_ex_tax_amount,omitempty"`
	// ItemName 是开票项目：货物或应税劳务、服务名称
	ItemName *string `json:"item_name,omitempty"`
	// ItemNo 是国税局制定的商品税收编码，必须是最末级
	ItemNo *string `json:"item_no,omitempty"`
	// ItemQuantity 是数量；	1.当row_type=0或2且item_unit_price为空，可空；	2.可精确到小数点后6位
	ItemQuantity *int64 `json:"item_quantity,omitempty"`
	// ItemSpec 是规格型号
	ItemSpec *string `json:"item_spec,omitempty"`
	// ItemSumAmount 是明细行价税合计，单位元，必须保证item_sum_amount=item_ex_tax_amount+item_tax_amount。
	ItemSumAmount *string `json:"item_sum_amount,omitempty"`
	// ItemTaxAmount 是明细行税额，单位元，保留两位小数，无税或者免税情况下输入：0.00。
	ItemTaxAmount *string `json:"item_tax_amount,omitempty"`
	// ItemTaxRate 是明细行税率，无税或者免税情况下输入：0.00。
	ItemTaxRate *string `json:"item_tax_rate,omitempty"`
	// ItemUnit 是单位
	ItemUnit *string `json:"item_unit,omitempty"`
	// ItemUnitPrice 是不含税单价（元）;	1.当row_type=0或2且item_quantity为空，可空	2.可精确到小数点后8位；
	ItemUnitPrice *string `json:"item_unit_price,omitempty"`
	// RowType 是发票行性质：0表示正常行，1表示折扣行，2表示被折扣行。
	RowType *string `json:"row_type,omitempty"`
}

// InvoiceOutputInfo 是支付宝官方 V3 OpenAPI 模型。
type InvoiceOutputInfo struct {
	// CheckCode 是发票防伪码，验证码
	CheckCode *string `json:"check_code,omitempty"`
	// Checker 是票面上复核人信息。
	Checker *string `json:"checker,omitempty"`
	// Clerk 是票面上开票员信息。
	Clerk *string `json:"clerk,omitempty"`
	// ExTaxAmount 是不含税金额，票面上的不含税金额，单位（元），保留两位小数。
	ExTaxAmount *string `json:"ex_tax_amount,omitempty"`
	// InvoiceCode 是发票代码
	InvoiceCode *string `json:"invoice_code,omitempty"`
	// InvoiceContent 是发票明细项。
	InvoiceContent *[]InvoiceItemOpen `json:"invoice_content,omitempty"`
	// InvoiceDate 是开票日期，格式如：YYYY-MM-DD。
	InvoiceDate *string `json:"invoice_date,omitempty"`
	// InvoiceImgUrl 是发票预览图片地址
	InvoiceImgUrl *string `json:"invoice_img_url,omitempty"`
	// InvoiceKind 是标识发票的类型，PLAIN:增值税电子普通发票;SPECIAL:增值税专用发票;PLAIN_INVOICE:增值税普通发票;SALSE_INVOICE:机动车销售统一发票 ;PAPER_INVOICE:增值税普通发票(卷式)。
	InvoiceKind *string `json:"invoice_kind,omitempty"`
	// InvoiceMemo 是票面上备注信息。
	InvoiceMemo *string `json:"invoice_memo,omitempty"`
	// InvoiceNo 是发票号码
	InvoiceNo    *string           `json:"invoice_no,omitempty"`
	InvoiceTitle *InvoiceTitleOpen `json:"invoice_title,omitempty"`
	// InvoiceType 是标识是红票还是蓝票，RED:红票;BLUE:蓝票
	InvoiceType *string `json:"invoice_type,omitempty"`
	// OriBlueInvCode 是红票情况下，必须填入原始蓝票的发票代码。
	OriBlueInvCode *string `json:"ori_blue_inv_code,omitempty"`
	// OriBlueInvNo 是红票情况下，必须填入原始蓝票的发票号码。
	OriBlueInvNo *string `json:"ori_blue_inv_no,omitempty"`
	// Payee 是票面上收款人信息。
	Payee *string `json:"payee,omitempty"`
	// PayeeAddressTel 是票面上销售方地址、电话信息。
	PayeeAddressTel *string `json:"payee_address_tel,omitempty"`
	// PayeeBankNameAccount 是票面上销售方开户行及账号。
	PayeeBankNameAccount *string `json:"payee_bank_name_account,omitempty"`
	// PayeeName 是票面上销售方企业名称。
	PayeeName *string `json:"payee_name,omitempty"`
	// PayeeRegisterNo 是票面上销售方税号信息。
	PayeeRegisterNo *string `json:"payee_register_no,omitempty"`
	// SumAmount 是价税合计，票面上的总开票金额，单位（元），保留两位小数，sum_amount=ex_tax_amount+tax_amount。
	SumAmount *string `json:"sum_amount,omitempty"`
	// TaxAmount 是合计税额，票面上的合计税额，单位（元），保留两位小数。
	TaxAmount *string `json:"tax_amount,omitempty"`
}

// InvoiceTitleOpen 是支付宝官方 V3 OpenAPI 模型。
type InvoiceTitleOpen struct {
	// PayerAddressTel 是票面上的购买方地址、电话
	PayerAddressTel *string `json:"payer_address_tel,omitempty"`
	// PayerBankNameAccount 是票面上的购买方开户行及账户
	PayerBankNameAccount *string `json:"payer_bank_name_account,omitempty"`
	// PayerRegisterNo 是票面上的购买方纳税人识别号
	PayerRegisterNo *string `json:"payer_register_no,omitempty"`
	// TitleName 是票面上的购买方名称
	TitleName *string `json:"title_name,omitempty"`
}

// OCRNormalScanInfo 是支付宝官方 V3 OpenAPI 模型。
type OCRNormalScanInfo struct {
	// InvoiceCode 是发票代码
	InvoiceCode *string `json:"invoice_code,omitempty"`
	// InvoiceDate 是开票时间
	InvoiceDate *string `json:"invoice_date,omitempty"`
	// InvoiceNo 是发票号码
	InvoiceNo *string `json:"invoice_no,omitempty"`
	// Price 是金额（元）
	Price *string `json:"price,omitempty"`
	// Remark 是明细事由
	Remark *string `json:"remark,omitempty"`
}

// OCRPlaneScanInfo 是支付宝官方 V3 OpenAPI 模型。
type OCRPlaneScanInfo struct {
	// Destination 是目的地
	Destination *string `json:"destination,omitempty"`
	// FlightNo 是航班号
	FlightNo *string `json:"flight_no,omitempty"`
	// InvoiceDate 是乘机日期
	InvoiceDate *string `json:"invoice_date,omitempty"`
	// Origin 是出发地
	Origin *string `json:"origin,omitempty"`
	// Passenger 是乘客
	Passenger *string `json:"passenger,omitempty"`
	// Price 是金额（元）
	Price *string `json:"price,omitempty"`
	// Remark 是明细事由
	Remark *string `json:"remark,omitempty"`
	// SeatClass 是飞机舱位
	SeatClass *string `json:"seat_class,omitempty"`
}

// OCRTaxiScanInfo 是支付宝官方 V3 OpenAPI 模型。
type OCRTaxiScanInfo struct {
	// GetOffTime 是下车时间
	GetOffTime *string `json:"get_off_time,omitempty"`
	// GetOnTime 是上车时间
	GetOnTime *string `json:"get_on_time,omitempty"`
	// InvoiceCode 是发票代码
	InvoiceCode *string `json:"invoice_code,omitempty"`
	// InvoiceDate 是开票时间
	InvoiceDate *string `json:"invoice_date,omitempty"`
	// InvoiceNo 是发票号码
	InvoiceNo *string `json:"invoice_no,omitempty"`
	// Passenger 是乘客人
	Passenger *string `json:"passenger,omitempty"`
	// Price 是金额（元）
	Price *string `json:"price,omitempty"`
	// TravelDist 是乘车距离
	TravelDist *string `json:"travel_dist,omitempty"`
}

// OCRTrainScanInfo 是支付宝官方 V3 OpenAPI 模型。
type OCRTrainScanInfo struct {
	// Destination 是目的地
	Destination *string `json:"destination,omitempty"`
	// InvoiceDate 是乘车日期
	InvoiceDate *string `json:"invoice_date,omitempty"`
	// Origin 是出发地
	Origin *string `json:"origin,omitempty"`
	// Passenger 是乘客
	Passenger *string `json:"passenger,omitempty"`
	// Price 是金额（元）
	Price *string `json:"price,omitempty"`
	// Remark 是明细事由
	Remark *string `json:"remark,omitempty"`
	// SeatClass 是座次
	SeatClass *string `json:"seat_class,omitempty"`
	// TrainNo 是车次
	TrainNo *string `json:"train_no,omitempty"`
}

// PackageItemOpenInfo 是支付宝官方 V3 OpenAPI 模型。
type PackageItemOpenInfo struct {
	ConsumeOutputInfo *ConsumeOutputInfo `json:"consume_output_info,omitempty"`
	// FileDownloadUrl 是发票文件下载链接
	FileDownloadUrl *string `json:"file_download_url,omitempty"`
	// FileType 是文件类型，如pdf，ofd，jpg
	FileType          *string            `json:"file_type,omitempty"`
	InvoiceOutputInfo *InvoiceOutputInfo `json:"invoice_output_info,omitempty"`
	// ItemSource 是项目来源，如：ocr\einvoice\consume
	ItemSource *string `json:"item_source,omitempty"`
	// ItemSubType 是项目子类，如：出租车、火车、飞机等
	ItemSubType *string `json:"item_sub_type,omitempty"`
	// ItemType 是项目大类，如：餐饮美食、酒店住宿、差旅出行等
	ItemType          *string            `json:"item_type,omitempty"`
	OCRNormalScanInfo *OCRNormalScanInfo `json:"ocr_normal_scan_info,omitempty"`
	OCRPlaneScanInfo  *OCRPlaneScanInfo  `json:"ocr_plane_scan_info,omitempty"`
	OCRTaxiScanInfo   *OCRTaxiScanInfo   `json:"ocr_taxi_scan_info,omitempty"`
	OCRTrainScanInfo  *OCRTrainScanInfo  `json:"ocr_train_scan_info,omitempty"`
}
