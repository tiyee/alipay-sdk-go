// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package open

// AutoInfoQueryRequest contains parameters for alios.open.auto.info.query.
type AutoInfoQueryRequest struct {
	// AuthToken represents auth_token. 用户授权令牌
	AuthToken *string `query:"auth_token"`
	// UserId represents user_id. 蚂蚁统一会员ID
	UserId *string `query:"user_id"`
	// OpenId represents open_id. 经度
	OpenId *string `query:"open_id"`
	// DeviceToken represents device_token. 设备token
	DeviceToken *string `query:"device_token"`
}
