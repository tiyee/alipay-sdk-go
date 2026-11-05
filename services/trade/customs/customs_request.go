package customs

import (
	"strings"

	"github.com/tiyee/alipay-sdk-go/core"
	"github.com/tiyee/alipay-sdk-go/internal/tradevalidate"
)

// CustomsDeclareRequest 是 alipay.trade.customs.declare（统一收单报关）
// 的请求。out_request_no、amount、customs_place、merchant_customs_code、
// merchant_customs_name 与交易号必填。
type CustomsDeclareRequest struct {
	// OutRequestNo 是报关流水号：商户生成的唯一业务编号，6-64 位，
	// 允许数字、字母、下划线与短横线。
	OutRequestNo string `json:"out_request_no"`
	// TradeNo 是支付宝交易号（与 OutTradeNo 二选一）。
	TradeNo string `json:"trade_no,omitempty"`
	// OutTradeNo 是国际站外部订单号。
	OutTradeNo string `json:"out_trade_no,omitempty"`
	// Amount 是报关金额，单位为人民币元，精确到小数点后 2 位。
	Amount string `json:"amount"`
	// CustomsPlace 是海关编号，取值见官方海关编号列表。
	CustomsPlace string `json:"customs_place"`
	// MerchantCustomsCode 是商户在海关备案的编号。
	MerchantCustomsCode string `json:"merchant_customs_code"`
	// MerchantCustomsName 是商户海关备案名称。
	MerchantCustomsName string `json:"merchant_customs_name"`
	// IsSplit 控制本单是否拆单；仅传 T 或 t 时触发拆单。
	IsSplit *string `json:"is_split,omitempty"`
	// SubOutBizNo 是拆单报关的商户子订单号；拆单时必须传入。
	SubOutBizNo *string `json:"sub_out_biz_no,omitempty"`
	// BuyerInfo 是订购人信息。
	BuyerInfo *CustomsDeclareBuyerInfo `json:"buyer_info,omitempty"`
	// DeclareMode 是报关模式：1 表示需要强校验买家和支付人的身份信息。
	DeclareMode *int64 `json:"declare_mode,omitempty"`
	// DepositTradeNo 是定金交易流水号；“定金+尾款”报关场景传入。
	DepositTradeNo *string `json:"deposit_trade_no,omitempty"`
	// MerchantID 是国际站商户标识。
	MerchantID *string `json:"merchant_id,omitempty"`
}

// CustomsQueryRequest 是 alipay.trade.customs.query（报关查询，V3 GET）
// 的请求：报关请求号逗号分隔，单次最多 10 个。
type CustomsQueryRequest struct {
	// OutRequestNos 是需要查询的商户端报关请求号，多个值用英文半角
	// 逗号分隔，单次最多 10 个。
	OutRequestNos string `query:"out_request_nos"`
}

func (r CustomsDeclareRequest) validate() error {
	if r.OutRequestNo == "" {
		return &core.ValidationError{Field: "out_request_no", Reason: "out_request_no is required"}
	}
	if err := tradevalidate.Identifier(r.TradeNo, r.OutTradeNo, "out_trade_no"); err != nil {
		return err
	}
	if err := tradevalidate.Amount(r.Amount, "amount"); err != nil {
		return err
	}
	if r.CustomsPlace == "" {
		return &core.ValidationError{Field: "customs_place", Reason: "customs_place is required"}
	}
	if r.MerchantCustomsCode == "" {
		return &core.ValidationError{Field: "merchant_customs_code", Reason: "merchant_customs_code is required"}
	}
	if r.MerchantCustomsName == "" {
		return &core.ValidationError{Field: "merchant_customs_name", Reason: "merchant_customs_name is required"}
	}
	return nil
}

func (r CustomsQueryRequest) validate() error {
	if r.OutRequestNos == "" {
		return &core.ValidationError{Field: "out_request_nos", Reason: "out_request_nos is required"}
	}
	if count := len(strings.Split(r.OutRequestNos, ",")); count > 10 {
		return &core.ValidationError{Field: "out_request_nos",
			Reason: "out_request_nos supports at most 10 values per request"}
	}
	return nil
}
