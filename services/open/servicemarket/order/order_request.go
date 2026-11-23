package order

import (
	"github.com/tiyee/alipay-sdk-go/core"
)

// ServiceMarketOrderAcceptRequest 是支付宝官方 V3 OpenAPI 模型。
type ServiceMarketOrderAcceptRequest struct {
	// CommodityOrderId 是服务商品订单ID。
	CommodityOrderId *string `json:"commodity_order_id,omitempty"`
}

// ServiceMarketOrderItemCancelRequest 是支付宝官方 V3 OpenAPI 模型。
type ServiceMarketOrderItemCancelRequest struct {
	// CancelReason 是当前门店区域不支持实施。
	CancelReason *string `json:"cancel_reason,omitempty"`
	// CommodityOrderId 是订购服务订单ID，通过支付宝发送的 <a href="https://opendocs.alipay.com/apis/api_3/alipay.open.servicemarket.order.notify">服务市场商户确认订购通知</a> 通知获取。。
	CommodityOrderId *string `json:"commodity_order_id,omitempty"`
	// MiniAppId 是商家订购服务选择的某一小程序的APPID，如果该订单为蚂蚁服务市场发布的小程序插件服务所产生的订单，则此字段必填。
	MiniAppId *string `json:"mini_app_id,omitempty"`
	// ShopId 是商家订购服务选择的某一门店的ID。若该订单： 为支付宝服务市场发布的服务所产生订单则该参数可选。。
	ShopId *string `json:"shop_id,omitempty"`
}

// ServiceMarketOrderItemCompleteRequest 是支付宝官方 V3 OpenAPI 模型。
type ServiceMarketOrderItemCompleteRequest struct {
	// CommodityOrderId 是订购服务插件订单号。
	CommodityOrderId *string `json:"commodity_order_id,omitempty"`
	// MiniAppId 是商家订购服务选择的某一小程序的APPID，如果该订单为蚂蚁服务市场发布的小程序插件服务所产生的订单，则此字段必填。
	MiniAppId *string `json:"mini_app_id,omitempty"`
	// ShopId 是商家订购服务选择的某一门店的ID。若该订单： 为支付宝服务市场发布的服务所产生订单则该参数可选。。
	ShopId *string `json:"shop_id,omitempty"`
}

// ServiceMarketOrderItemConfirmRequest 是支付宝官方 V3 OpenAPI 模型。
type ServiceMarketOrderItemConfirmRequest struct {
	// CommodityOrderId 是订购服务插件订单号。
	CommodityOrderId *string `json:"commodity_order_id,omitempty"`
	// MiniAppId 是商家订购服务选择的某一小程序的APPID，如果该订单为蚂蚁服务市场发布的小程序插件服务所产生的订单，则此字段必填。
	MiniAppId *string `json:"mini_app_id,omitempty"`
	// ShopId 是商家订购服务选择的某一门店的ID。若该订单： 为支付宝(蚂蚁)服务市场发布的服务所产生订单则该参数可选。。
	ShopId *string `json:"shop_id,omitempty"`
}

// ServiceMarketOrderRejectRequest 是支付宝官方 V3 OpenAPI 模型。
type ServiceMarketOrderRejectRequest struct {
	// CommodityOrderId 是订购服务商品订单ID。
	CommodityOrderId *string `json:"commodity_order_id,omitempty"`
	// RejectReason 是拒绝接单原因。
	RejectReason *string `json:"reject_reason,omitempty"`
}

// ServiceMarketOrderNotifyRequest 是支付宝官方 V3 OpenAPI 模型。
type ServiceMarketOrderNotifyRequest struct {
	// NotifyUrl 是通知地址。
	NotifyUrl *string `json:"notify_url,omitempty"`
}

// ServiceMarketOrderCreateRequest 是支付宝官方 V3 OpenAPI 模型。
type ServiceMarketOrderCreateRequest struct {
	AppLogo *core.MultipartFile       `json:"app_logo,omitempty" multipart:"app_logo,file"`
	Data    *ServiceMarketOrderCreate `json:"data,omitempty" multipart:"data,json"`
}

// ServiceMarketOrderQueryRequest is the request type for alipay.open.servicemarket.order.query.
type ServiceMarketOrderQueryRequest struct {
	// CommodityOrderId 是订购服务订单ID，通过支付宝发送的 服务市场商户确认订购通知 获取。
	CommodityOrderId *string `json:"commodity_order_id,omitempty" query:"commodity_order_id,omitempty"`
	// StartPage 是查询开始页码，1 为第一页。
	StartPage *string `json:"start_page,omitempty" query:"start_page,omitempty"`
}
