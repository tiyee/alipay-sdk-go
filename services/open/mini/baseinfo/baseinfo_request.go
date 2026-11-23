// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package baseinfo

import (
	"github.com/tiyee/alipay-sdk-go/core"
)

// MiniBaseInfoModifyRequest 是支付宝官方 V3 OpenAPI 模型。
type MiniBaseInfoModifyRequest struct {
	AppLogo *core.MultipartFile `json:"app_logo,omitempty" multipart:"app_logo,file"`
	Data    *MiniBaseInfoModify `json:"data,omitempty" multipart:"data,json"`
}

// MiniBaseInfoNameCheckRequest is the request type for alipay.open.mini.baseinfo.name.check.
type MiniBaseInfoNameCheckRequest struct {
	// AppName 是小程序名字，长度限制 3~30 个字符，仅支持包含简体中文、数字、英文、下划线、+、-。
	AppName *string `json:"app_name,omitempty" query:"app_name,omitempty"`
}

// MiniBaseInfoQueryRequest is the request type for alipay.open.mini.baseinfo.query.
type MiniBaseInfoQueryRequest struct{}
