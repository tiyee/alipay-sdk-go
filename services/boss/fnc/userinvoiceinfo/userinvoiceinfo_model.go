package userinvoiceinfo

// UserMailInfoOrder 是支付宝官方 V3 OpenAPI 模型。
type UserMailInfoOrder struct {
	// City 是联系所在城市
	City *string `json:"city,omitempty"`
	// Country 是联系人国家
	Country *string `json:"country,omitempty"`
	// CountyDistrict 是联系人所在县/区
	CountyDistrict *string `json:"county_district,omitempty"`
	// DetailAddress 是联系所在详细地址
	DetailAddress *string `json:"detail_address,omitempty"`
	// Email 是电子邮箱
	Email *string `json:"email,omitempty"`
	// IpRoleId 是商户ipRole(pid)
	IpRoleId *string `json:"ip_role_id,omitempty"`
	// Name 是联系人名字
	Name *string `json:"name,omitempty"`
	// Province 是联系人省份
	Province *string `json:"province,omitempty"`
	// Street 是联系人所在街道
	Street *string `json:"street,omitempty"`
	// Telephone 是联系人电话
	Telephone *string `json:"telephone,omitempty"`
}

// UserMailInfoVO 是支付宝官方 V3 OpenAPI 模型。
type UserMailInfoVO struct {
	// Address 是收件人地址
	Address *string `json:"address,omitempty"`
	// City 是城市
	City *string `json:"city,omitempty"`
	// Country 是国家
	Country *string `json:"country,omitempty"`
	// CountyDistrict 是县区
	CountyDistrict *string `json:"county_district,omitempty"`
	// DetailAddress 是详细地址
	DetailAddress *string `json:"detail_address,omitempty"`
	// Email 是邮件人的电子邮箱
	Email *string `json:"email,omitempty"`
	// GmtCreate 是创建时间
	GmtCreate *string `json:"gmt_create,omitempty"`
	// GmtModified 是修改时间
	GmtModified *string `json:"gmt_modified,omitempty"`
	// Id 是收件人信息ID
	Id *string `json:"id,omitempty"`
	// Name 是收件人姓名
	Name *string `json:"name,omitempty"`
	// Pid 是pid用于指定该地址归属的PID，应对集团模式的商户所设定,发票邮寄地址根据PID确定
	Pid *string `json:"pid,omitempty"`
	// Province 是省份
	Province *string `json:"province,omitempty"`
	// Street 是街道
	Street *string `json:"street,omitempty"`
	// Telephone 是收件人电话
	Telephone *string `json:"telephone,omitempty"`
}
