// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package biz

// Gavintest 是支付宝官方 V3 OpenAPI 模型。
type Gavintest struct {
	// Newid 是测试。
	Newid *int64 `json:"newid,omitempty"`
}

// GavintestNewLeveaOne 是支付宝官方 V3 OpenAPI 模型。
type GavintestNewLeveaOne struct {
	// Boolen 是2。
	Boolen *bool `json:"boolen,omitempty"`
	// Ces 是21。
	Ces *[]int64 `json:"ces,omitempty"`
	// Des 是12。
	Des *string `json:"des,omitempty"`
	// Str 是1。
	Str *[]string `json:"str,omitempty"`
}

// GavintestNewonline 是支付宝官方 V3 OpenAPI 模型。
type GavintestNewonline struct {
	// Dem 是1。
	Dem    *[]string             `json:"dem,omitempty"`
	String *GavintestNewLeveaOne `json:"string,omitempty"`
}
