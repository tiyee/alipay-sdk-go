// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package edu

// BillSendExtInfo 是支付宝官方 V3 OpenAPI 模型。
type BillSendExtInfo struct {
	// ExtSchoolId 是学校外标编号
	ExtSchoolId *string `json:"ext_school_id,omitempty"`
	// OrderPayType 是账单唤起收银台的方式：	1-手机wap;2-当面付。默认值为1（不传值或传非数字值）
	OrderPayType *string `json:"order_pay_type,omitempty"`
	// RoyaltyAmount 是分账金额，单位：元
	RoyaltyAmount *string `json:"royalty_amount,omitempty"`
	// SchoolName 是学校名称
	SchoolName *string `json:"school_name,omitempty"`
	// TransInPid 是分账收入户PID
	TransInPid *string `json:"trans_in_pid,omitempty"`
}

// ChargeItems 是支付宝官方 V3 OpenAPI 模型。
type ChargeItems struct {
	// ItemMandatory 是缴费项是否必选 	如果缴费项是多选模式，此参数生效。 “Y”表示必填，“N”或空表示非必填。
	ItemMandatory *string `json:"item_mandatory,omitempty"`
	// ItemMaximum 是缴费项最大可选数 	如果缴费项是多选模式，此参数生效，范围是1-9，如果为空，则最大项默认为9
	ItemMaximum *int64 `json:"item_maximum,omitempty"`
	// ItemName 是缴费项名称
	ItemName *string `json:"item_name,omitempty"`
	// ItemPrice 是缴费项金额
	ItemPrice *string `json:"item_price,omitempty"`
	// ItemSerialNumber 是缴费项序号，如果缴费项是多选模式，此项为必填，建议从1开始的连续数字， 	用户支付成功后，通过passback_params参数带回已选择的缴费项。例如:orderNo=uoo234234&isvOrderNo=24werwe&items=1-2|2-1|3-5 	1-2|2-1|3-5 表示：缴费项序列号-缴费项数|缴费项序列号-缴费项数
	ItemSerialNumber *int64 `json:"item_serial_number,omitempty"`
}

// UserDetails 是支付宝官方 V3 OpenAPI 模型。
type UserDetails struct {
	// UserChangeMobile 是用户变更手机号，替换旧的手机号
	UserChangeMobile *string `json:"user_change_mobile,omitempty"`
	// UserMobile 是孩子家长的手机号
	UserMobile *string `json:"user_mobile,omitempty"`
	// UserName 是张三
	UserName *string `json:"user_name,omitempty"`
	// UserRelation 是孩子与家长的关系： 1、爸爸  2、妈妈  3、爷爷  4、奶奶  5、外公  6、外婆  7、家长
	UserRelation *string `json:"user_relation,omitempty"`
}
