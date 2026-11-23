package user

import (
	"github.com/tiyee/alipay-sdk-go/services/commerce/ec"
)

// EcUserEnterpriseQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type EcUserEnterpriseQueryResponse struct {
	// EnterpriseInfoList 是用户所属企业列表
	EnterpriseInfoList *[]ec.EnterpriseInfoDTO `json:"enterprise_info_list,omitempty"`
}
