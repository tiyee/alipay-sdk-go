// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package edu

// EduKtBillingModifyRequest 是支付宝官方 V3 OpenAPI 模型。
type EduKtBillingModifyRequest struct {
	// BankSuccess 是成功Y，失败N
	BankSuccess *string `json:"bank_success,omitempty"`
	// BuyerLogonId 是退款时，支付宝返回的用户的登录id
	BuyerLogonId *string `json:"buyer_logon_id,omitempty"`
	// BuyerUserId 是支付宝返回的买家支付宝用户id
	BuyerUserId *string `json:"buyer_user_id,omitempty"`
	// BuyerUserOpenId 是支付宝返回的买家支付宝用户id加密后信息
	BuyerUserOpenId *string `json:"buyer_user_open_id,omitempty"`
	// FundChange 是本次退款是否发生了资金变化
	FundChange *string `json:"fund_change,omitempty"`
	// GmtRefund 是支付宝返回的退款时间，而不是商户退款申请的时间
	GmtRefund *string `json:"gmt_refund,omitempty"`
	// OutRequestNo 是标识一次退款请求，同一笔交易多次退款需要保证唯一，如需部分退款，则此参数必传。（若退款时填写，则同步退款状态时也必须填写）
	OutRequestNo *string `json:"out_request_no,omitempty"`
	// OutTradeNo 是isv系统的订单号
	OutTradeNo *string `json:"out_trade_no,omitempty"`
	// RefundAmount 是需要退款的金额，该金额不能大于订单金额,单位为元，支持两位小数
	RefundAmount *string `json:"refund_amount,omitempty"`
	// RefundDetailItemList 是支付宝返回的退款资金渠道，json格式字符串
	RefundDetailItemList *string `json:"refund_detail_item_list,omitempty"`
	// RefundReason 是退款原因，商家根据客户实际退款原因填写（若退款时填写，则同步退款状态时也必须填写）
	RefundReason *string `json:"refund_reason,omitempty"`
	// Status 是状态：1:缴费成功，2:关闭账单，3、退费 	如果为退款状态，需要填写fund_change,  	refund_amount, refund_reason, 	out_request_no, buyer_logon_id, 	gmt_refund, 	buyer_user_id, refund_detail_item_list; 	4、同步网商返回的状态,如果是网商银行的账单，bank_success这个字段必填
	Status *string `json:"status,omitempty"`
	// TradeNo 是支付宝返回的交易号
	TradeNo *string `json:"trade_no,omitempty"`
}

// EduKtBillingSendRequest 是支付宝官方 V3 OpenAPI 模型。
type EduKtBillingSendRequest struct {
	// Amount 是总金额，单位为元，精确到小数点后两位，取值范围[0.01,100000000]， 	如果是非多选项，是要和缴费项的总和相同，多选模式不做验证
	Amount *string `json:"amount,omitempty"`
	// ChargeBillTitle 是缴费账单名称
	ChargeBillTitle *string `json:"charge_bill_title,omitempty"`
	// ChargeItem 是缴费详情：输入json格式字符串。Json定义：key填写缴费项名称，value填写缴费项金额，金额保留2位小数（单位：元）
	ChargeItem *[]ChargeItems `json:"charge_item,omitempty"`
	// ChargeType 是缴费项模式：空或"N"，表示缴费项不可选， 	"M"表示缴费项为可选 ，支持单选和多选。
	ChargeType *string `json:"charge_type,omitempty"`
	// ChildName 是孩子名字
	ChildName *string `json:"child_name,omitempty"`
	// ClassIn 是孩子所在班级
	ClassIn *string `json:"class_in,omitempty"`
	// EndEnable 是截止日期是否生效，与gmt_end发布配合使用,N为gmt_end不生效，用户过期后仍可以缴费；Y为gmt_end生效，用户过期后，不能再缴费。
	EndEnable *string          `json:"end_enable,omitempty"`
	ExtInfo   *BillSendExtInfo `json:"ext_info,omitempty"`
	// GmtEnd 是缴费截止时间，格式"yyyy-MM-dd HH:mm:ss"，日期要大于当前时间。请注意，过期时间不宜设置过短。
	GmtEnd *string `json:"gmt_end,omitempty"`
	// Grade 是孩子所在年级
	Grade *string `json:"grade,omitempty"`
	// OutTradeNo 是ISV端的缴费账单编号
	OutTradeNo *string `json:"out_trade_no,omitempty"`
	// PartnerId 是Isv支付宝pid, 支付宝签约后，返回给ISV编号
	PartnerId *string `json:"partner_id,omitempty"`
	// SchoolNo 是学校编码，录入学校接口返回的参数
	SchoolNo *string `json:"school_no,omitempty"`
	// SchoolPid 是学校支付宝pid,直付通填写smid
	SchoolPid *string `json:"school_pid,omitempty"`
	// StudentCode 是学生的学号，只支持字母和数字类型，一般以教育局学号为准，作为学生的唯一标识。此字段与student_identify、家长user_mobile至少选一个
	StudentCode *string `json:"student_code,omitempty"`
	// StudentIdentify 是学生的身份证号，如果ISV有学生身份证号，则同步身份证号作为学生唯一标识。此字段与student_code、家长user_mobile至少选一个。 	大陆身份证必须是18位 ， 其它地区或国家的身份证开头需要加"IC"开头区分并且不超过18位，但查询账单的时候不要带"IC"
	StudentIdentify *string `json:"student_identify,omitempty"`
	// Users 是孩子的家长信息，最多一次输入10个家长，此字段做为识别家长的孩子用，与student_identify、student_code至少选一个
	Users *[]UserDetails `json:"users,omitempty"`
}

// EduKtSchoolInfoModifyRequest 是支付宝官方 V3 OpenAPI 模型。
type EduKtSchoolInfoModifyRequest struct {
	// BankNotifyUrl 是与浙江网商交易见证平台有交互ISV输入网商交易异步通知回调URL，教育缴费同步账单信息给网商，网商会回调此url，ISV即可获取网商相关的参数，根据教育缴费平台账单发送接口返回的 order_no和网商返回的outer_trade_no来对应账单信息。
	BankNotifyUrl *string `json:"bank_notify_url,omitempty"`
	// BankPartnerId 是与浙江网商交易见证平台有交互的ISV,由交易见证平台分配给合作方业务平台 签约唯一ID ，由网商分配给ISV的渠道参数
	BankPartnerId *string `json:"bank_partner_id,omitempty"`
	// BankUid 是与浙江网商交易见证平台有交互的ISV在创建账户时的uid，也就是ISV平台上的用户ID（字母或数字）
	BankUid *string `json:"bank_uid,omitempty"`
	// BankcardNo 是对应集团到卡模式的银行编号.学校与支付宝后台签约时，由学校提交给支付宝的编号
	BankcardNo *string `json:"bankcard_no,omitempty"`
	// BatchNo 是集团收单模式：分账批次号，支付宝配置后提供的银行卡批次号
	BatchNo *string `json:"batch_no,omitempty"`
	// CardAliasNo 是学校开通直付通卡编号，smid与card_alias_no必须同时填写
	CardAliasNo *string `json:"card_alias_no,omitempty"`
	// CityCode 是城市的国家编码。参见 国家统计局出版的行政区划代码。
	CityCode *string `json:"city_code,omitempty"`
	// CityName 是城市名称
	CityName *string `json:"city_name,omitempty"`
	// CorporateBranchPid 是集团收单模式：BD批量上传银行卡信息后，支付宝系统分配给ISV的每个卡分配的唯一标识
	CorporateBranchPid *string `json:"corporate_branch_pid,omitempty"`
	// DistrictCode 是区县的国家编码。参见 国家统计局出版的行政区划代码。
	DistrictCode *string `json:"district_code,omitempty"`
	// DistrictName 是区县名称
	DistrictName *string `json:"district_name,omitempty"`
	// IsvName 是ISV公司名称 ， 会在账单详情页面展示给用户
	IsvName *string `json:"isv_name,omitempty"`
	// IsvNo 是注意：本参数于2017-08-03开始已经废弃，不再需要传递。	由支付宝提供的给已经签约的isv的编码，业务上一般直接采用isv的支付宝PID。
	IsvNo *string `json:"isv_no,omitempty"`
	// IsvNotifyUrl 是此通知地址是为了保持教育缴费平台与ISV商户支付状态一致性。用户支付成功后，支付宝会根据本isv_notify_url(异步通知说明https://docs.open.alipay.com/203/105286/)，通过POST请求的形式将支付结果作为参数通知到商户系统，ISV商户可以根据返回的参数更新账单状态。
	IsvNotifyUrl *string `json:"isv_notify_url,omitempty"`
	// IsvPhone 是ISV的联系方式 ， 会在账单详情页面展示给用户，用户有问题可以直接联系此电话获取帮助
	IsvPhone *string `json:"isv_phone,omitempty"`
	// IsvPid 是填写已经签约教育缴费的isv的支付宝PID
	IsvPid *string `json:"isv_pid,omitempty"`
	// ProvinceCode 是省份的国家编码，参见 国家统计局出版的行政区划代码。
	ProvinceCode *string `json:"province_code,omitempty"`
	// ProvinceName 是省名称
	ProvinceName *string `json:"province_name,omitempty"`
	// SchoolIcon 是学校的校徽或logo，在学校展示页面作为学校的标识。该字段为图片的链接地址，只支持png或jpg图片格式，图片高度为108，宽度为108 ，不大于20k。 	注意：此链接要求公网可以访问，否则无法正常展示。
	SchoolIcon *string `json:"school_icon,omitempty"`
	// SchoolIconType 是如果填写了school_icon参数，则此字段不能为空。目前只支持png和jpg两种格式
	SchoolIconType *string `json:"school_icon_type,omitempty"`
	// SchoolName 是学校名称
	SchoolName *string `json:"school_name,omitempty"`
	// SchoolPid 是学校签约支付宝教育缴费支付宝pid,如果是直付通学校，填写直付通返回的smid
	SchoolPid *string `json:"school_pid,omitempty"`
	// SchoolStdcode 是学校(机构)标识码（由教育部按照国家标准及编码规则编制，可以在教育局官网查询）
	SchoolStdcode *string `json:"school_stdcode,omitempty"`
	// SchoolType 是学校的类型。枚举值如下：	*1：代表托儿所。	*2：代表幼儿园。	*3：代表小学。	*4：代表初中。	*5：代表高中。	注意：如果学校兼有多种属性，可以连写，比如：45：代表兼有初中部高中部；34：代表兼有小学部初中部。
	SchoolType *string `json:"school_type,omitempty"`
	// Smid 是学校开通直付通返回的二级商户id，smid与card_alias_no必须同时填写
	Smid *string `json:"smid,omitempty"`
	// TransIn 是集团收单模式：分账批次号，支付宝配置后提供的银行卡批次号
	TransIn *string `json:"trans_in,omitempty"`
	// WhiteChannelCode 是与浙江网商交易见证平台有交互的ISV,由网商分配给ISV的渠道参数
	WhiteChannelCode *string `json:"white_channel_code,omitempty"`
}

// EduKtStudentModifyRequest 是支付宝官方 V3 OpenAPI 模型。
type EduKtStudentModifyRequest struct {
	// ChildName 是修改后的学生姓名 	本接口调用时，child_name、student_code、student_identify、users这几个参数至少需要填写其中一个，不能同时为空
	ChildName *string `json:"child_name,omitempty"`
	// IsvPid 是已经签约教育缴费的isv的支付宝PID
	IsvPid *string `json:"isv_pid,omitempty"`
	// SchoolNo 是学校编号，调用alipay.eco.edu.kt.schoolinfo.modify接口录入学校信息时，接口返回的编号
	SchoolNo *string `json:"school_no,omitempty"`
	// SchoolPid 是学校用来签约支付宝教育缴费的支付宝PID
	SchoolPid *string `json:"school_pid,omitempty"`
	// Status 是区分ISV操作，“D”表示删除，“U”表示更新，区分大小写。 	如果为U，则学生名字，学号，身份证至少填写一项
	Status *string `json:"status,omitempty"`
	// StudentCode 是修改后的学号 	本接口调用时，child_name、student_code、student_identify、users这几个参数至少需要填写其中一个，不能同时为空
	StudentCode *string `json:"student_code,omitempty"`
	// StudentIdentify 是修改后的身份证号码 	本接口调用时，child_name、student_code、student_identify、users这几个参数至少需要填写其中一个，不能同时为空
	StudentIdentify *string `json:"student_identify,omitempty"`
	// StudentNo 是支付宝-中小学-教育缴费生成的学生唯一编号
	StudentNo *string `json:"student_no,omitempty"`
	// Users 是孩子的家长信息，最多一次输入20个家长。如果输入的家长信息不存在，则该改学生增加家长信息 	本接口调用时，child_name、student_code、student_identify、users这几个参数至少需要填写其中一个，不能同时为空
	Users *[]UserDetails `json:"users,omitempty"`
}

// EduKtBillingQueryRequest is the query request model for alipay.eco.edu.kt.billing.query.
type EduKtBillingQueryRequest struct {
	// IsvPid 是Isv pid
	IsvPid *string `json:"isv_pid,omitempty" query:"isv_pid,omitempty"`
	// SchoolPid 是学校支付宝pid，直付通填写smid
	SchoolPid *string `json:"school_pid,omitempty" query:"school_pid,omitempty"`
	// OutTradeNo 是ISV调用发送账单接口，返回给商户的order_no
	OutTradeNo *string `json:"out_trade_no,omitempty" query:"out_trade_no,omitempty"`
}

// EduKtParentQueryRequest is the query request model for alipay.eco.edu.kt.parent.query.
type EduKtParentQueryRequest struct {
	// SchoolPid 是学校支付宝pid
	SchoolPid *string `json:"school_pid,omitempty" query:"school_pid,omitempty"`
	// SchoolNo 是学校编码，录入学校接口返回的school_no参数
	SchoolNo *string `json:"school_no,omitempty" query:"school_no,omitempty"`
	// PartnerId 是Isv的支付宝pid
	PartnerId *string `json:"partner_id,omitempty" query:"partner_id,omitempty"`
	// ChildName 是孩子或学生姓名
	ChildName *string `json:"child_name,omitempty" query:"child_name,omitempty"`
	// UserMobile 是用户手机号，发账单时填写users数组中的一个手机号。结果返回用户是否通过此手机号添加此学生的缴费账户。user_mobile、student_code 、student_identify 必须并且只能填一项。
	UserMobile *string `json:"user_mobile,omitempty" query:"user_mobile,omitempty"`
	// StudentCode 是学生的学号，一个学校的学号必须是唯一。结果返回用户是否通过此学号添加此学生的缴费账户。user_mobile、student_code 、student_identify 必须并且只能填一项。
	StudentCode *string `json:"student_code,omitempty" query:"student_code,omitempty"`
	// StudentIdentify 是学生的身份证号，使用身份证规则验证。结果返回用户是否通过此身份证号添加此学生的缴费账户。user_mobile、student_code 、student_identify 必须并且只能填一项。
	StudentIdentify *string `json:"student_identify,omitempty" query:"student_identify,omitempty"`
}
