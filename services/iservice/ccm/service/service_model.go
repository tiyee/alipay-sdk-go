// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package service

// ISVSpiDefinition 是支付宝官方 V3 OpenAPI 模型。
type ISVSpiDefinition struct {
	// BizCode 是ISV自定义的标识功能的业务代码，不可重复
	BizCode *string `json:"biz_code,omitempty"`
	// Description 是spi功能描述
	Description *string `json:"description,omitempty"`
	// Icon 是功能图标文件url
	Icon *string `json:"icon,omitempty"`
	// SpiEndpoint 是spi接口服务地址
	SpiEndpoint *string `json:"spi_endpoint,omitempty"`
	// SpiExtProperty 是spi接口扩展参数，json格式字符串
	SpiExtProperty *string `json:"spi_ext_property,omitempty"`
	// SpiKey 是CCM预先定义的spi key，与插件位置有关
	SpiKey *string `json:"spi_key,omitempty"`
	// SpiName 是SPI 名称
	SpiName *string `json:"spi_name,omitempty"`
}

// SpiResult 是支付宝官方 V3 OpenAPI 模型。
type SpiResult struct {
	// BizCode 是isv自定义业务代码
	BizCode *string `json:"biz_code,omitempty"`
	// SpiId 是spi唯一id
	SpiId *int64 `json:"spi_id,omitempty"`
}
