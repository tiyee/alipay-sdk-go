package enterprise

import (
	"encoding/json"

	"github.com/tiyee/alipay-sdk-go/services/commerce/ec"
)

// EcEnterpriseAddressAddResponse 是支付宝官方 V3 OpenAPI 模型。
type EcEnterpriseAddressAddResponse struct {
	// AddressId 是地址id
	AddressId *string `json:"address_id,omitempty"`
}

// EcEnterpriseAddressModifyResponse 是支付宝官方 V3 OpenAPI 模型。
type EcEnterpriseAddressModifyResponse struct {
	// AddressId 是地址id
	AddressId *string `json:"address_id,omitempty"`
}

// EcEnterpriseAddressQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type EcEnterpriseAddressQueryResponse struct {
	// CurrentPage 是页码
	CurrentPage *int64 `json:"current_page,omitempty"`
	// EnterpriseAddressInfoList 是地址列表
	EnterpriseAddressInfoList *[]AddressInfoDTO `json:"enterprise_address_info_list,omitempty"`
	// TotalNum 是总数量
	TotalNum *int64 `json:"total_num,omitempty"`
	// TotalPages 是总页数
	TotalPages *int64 `json:"total_pages,omitempty"`
}

// EcEnterpriseAgreementQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type EcEnterpriseAgreementQueryResponse struct {
	EnterpriseAgreementDto *EnterpriseAgreementDTO `json:"enterprise_agreement_dto,omitempty"`
}

// EcEnterpriseCreateResponse 是支付宝官方 V3 OpenAPI 模型。
type EcEnterpriseCreateResponse struct {
	// EnterpriseId 是企业id
	EnterpriseId *string `json:"enterprise_id,omitempty"`
	// IotGroupId 是封闭场景（如班车）的人脸库id，如果入参create_iot_group传false，则不会创建企业人脸库，该字段为空。
	IotGroupId *string `json:"iot_group_id,omitempty"`
	// IotLogicGroupId 是开放场景（如团餐）的人脸库id，如果入参create_iot_group传false，则不会创建企业人脸库，该字段为空。
	IotLogicGroupId *string `json:"iot_logic_group_id,omitempty"`
	// SignUrl 是资金代付签约链接
	SignUrl *string `json:"sign_url,omitempty"`
}

// EcEnterpriseDeleteResponse 是支付宝官方 V3 OpenAPI 模型。
type EcEnterpriseDeleteResponse struct {
	// EnterpriseId 是企业ID(与入参相同，无意义的出参)
	EnterpriseId *string `json:"enterprise_id,omitempty"`
}

// EcEnterpriseInfoModifyResponse 是支付宝官方 V3 OpenAPI 模型。
type EcEnterpriseInfoModifyResponse = json.RawMessage

// EcEnterpriseInfoQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type EcEnterpriseInfoQueryResponse struct {
	EnterpriseInfo *ec.EnterpriseInfoDTO `json:"enterprise_info,omitempty"`
}

// EcEnterpriseUnsignResponse 是支付宝官方 V3 OpenAPI 模型。
type EcEnterpriseUnsignResponse struct {
	// AccountId 是入参传的共同账户id，出参也会返回共同账户id
	AccountId *string `json:"account_id,omitempty"`
	// EnterpriseId 是企业id
	EnterpriseId *string `json:"enterprise_id,omitempty"`
	// UnsignUrl 是解约链接
	UnsignUrl *string `json:"unsign_url,omitempty"`
}
