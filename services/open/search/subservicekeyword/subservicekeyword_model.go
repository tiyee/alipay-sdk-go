package subservicekeyword

import (
	"github.com/tiyee/alipay-sdk-go/services/open/search"
)

// SubServiceKeyWordInfo 是支付宝官方 V3 OpenAPI 模型。
type SubServiceKeyWordInfo struct {
	// AppName 是小程序名称。
	AppName *string `json:"app_name,omitempty"`
	// KeywordInfo 是此服务下的关键词。
	KeywordInfo *[]search.KeyWordInfo `json:"keyword_info,omitempty"`
	// MiniAppId 是小程序ID。
	MiniAppId *string `json:"mini_app_id,omitempty"`
	// SubServiceCode 是子服务code。
	SubServiceCode *string `json:"sub_service_code,omitempty"`
	// SubServiceDesc 是子服务简单描述。
	SubServiceDesc *string `json:"sub_service_desc,omitempty"`
	// SubServiceName 是子服务名。
	SubServiceName *string `json:"sub_service_name,omitempty"`
	// SubServiceStatus 是子服务状态， INITIAL：初始化：VALID:已生效；INVALID：已失效。
	SubServiceStatus *string `json:"sub_service_status,omitempty"`
}
