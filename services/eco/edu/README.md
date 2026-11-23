# eco/edu 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/eco/edu`。本目录提供 6 个支付宝开放平台 V3 REST 操作。

## 使用方式

先按[项目接入说明](../../../README.md)创建 `*core.Client`，再通过 `edu.NewService(client)` 创建 Service；所有网络调用以 `context.Context` 为首个参数，支持 `opts ...core.CallOption`。响应为 `*core.Response[T]`，业务数据位于 `resp.Data`。

请求、响应与嵌套类型的链接指向本目录源码；可参考本目录的 [API 示例](README.md#业务子包)。

示例文件中的 `TestSandboxService_*` 可使用沙箱平台证书执行真实调用。默认 `go test` 保持离线。

## 接口目录

| SDK 方法 | 官方接口 | 用途 | HTTP |
| --- | --- | --- | --- |
| [`EduKtBillingModify`](#api-eduktbillingmodify) | [`alipay.eco.edu.kt.billing.modify`](https://opendocs.alipay.com/open-v3/1406f415_alipay.eco.edu.kt.billing.modify) | 教育缴费账单状态同步接口 | `POST` |
| [`EduKtBillingQuery`](#api-eduktbillingquery) | [`alipay.eco.edu.kt.billing.query`](https://opendocs.alipay.com/open-v3/d9852f73_alipay.eco.edu.kt.billing.query) | 缴费账单查询 | `GET` |
| [`EduKtBillingSend`](#api-eduktbillingsend) | [`alipay.eco.edu.kt.billing.send`](https://opendocs.alipay.com/open-v3/7c14b5a1_alipay.eco.edu.kt.billing.send) | 教育缴费账单发送接口 | `POST` |
| [`EduKtParentQuery`](#api-eduktparentquery) | [`alipay.eco.edu.kt.parent.query`](https://opendocs.alipay.com/apis/05x0kr) | 查询学生家长状态接口 | `GET` |
| [`EduKtSchoolInfoModify`](#api-eduktschoolinfomodify) | [`alipay.eco.edu.kt.schoolinfo.modify`](https://opendocs.alipay.com/open-v3/a319c0c9_alipay.eco.edu.kt.schoolinfo.modify) | 教育缴费学校信息录入接口 | `POST` |
| [`EduKtStudentModify`](#api-eduktstudentmodify) | [`alipay.eco.edu.kt.student.modify`](https://opendocs.alipay.com/apis/05wun3) | 学生信息更新 | `POST` |

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-eduktbillingmodify"></a>

### EduKtBillingModify — 教育缴费账单状态同步接口

- 官方接口：[`alipay.eco.edu.kt.billing.modify`](https://opendocs.alipay.com/open-v3/1406f415_alipay.eco.edu.kt.billing.modify)
- 路由：`POST /v3/alipay/eco/edu/kt/billing/modify`
- 请求类型：[EduKtBillingModifyRequest](edu_request.go#L5)
- 响应类型：[EduKtBillingModifyResponse](edu_response.go#L5)

isv向支付宝－中小学－教育缴费发送教育缴费账单后，家长在支付宝－中小学－教育缴费中对账单完成支付操作且支付成功，ISV在自有系统内容对账，完成后通过此接口同步对账后的账单状态。状态分为：缴费成功、缴费失败。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `bank_success` | `BankSuccess` | `*string` | JSON body | 未标注 | 成功Y，失败N |
| `buyer_logon_id` | `BuyerLogonId` | `*string` | JSON body | 未标注 | 退款时，支付宝返回的用户的登录id |
| `buyer_user_id` | `BuyerUserId` | `*string` | JSON body | 未标注 | 支付宝返回的买家支付宝用户id |
| `buyer_user_open_id` | `BuyerUserOpenId` | `*string` | JSON body | 未标注 | 支付宝返回的买家支付宝用户id加密后信息 |
| `fund_change` | `FundChange` | `*string` | JSON body | 未标注 | 本次退款是否发生了资金变化 |
| `gmt_refund` | `GmtRefund` | `*string` | JSON body | 未标注 | 支付宝返回的退款时间，而不是商户退款申请的时间 |
| `out_request_no` | `OutRequestNo` | `*string` | JSON body | 未标注 | 标识一次退款请求，同一笔交易多次退款需要保证唯一，如需部分退款，则此参数必传。（若退款时填写，则同步退款状态时也必须填写） |
| `out_trade_no` | `OutTradeNo` | `*string` | JSON body | 未标注 | isv系统的订单号 |
| `refund_amount` | `RefundAmount` | `*string` | JSON body | 未标注 | 需要退款的金额，该金额不能大于订单金额,单位为元，支持两位小数 |
| `refund_detail_item_list` | `RefundDetailItemList` | `*string` | JSON body | 未标注 | 支付宝返回的退款资金渠道，json格式字符串 |
| `refund_reason` | `RefundReason` | `*string` | JSON body | 未标注 | 退款原因，商家根据客户实际退款原因填写（若退款时填写，则同步退款状态时也必须填写） |
| `status` | `Status` | `*string` | JSON body | 未标注 | 状态：1:缴费成功，2:关闭账单，3、退费<br> 如果为退款状态，需要填写fund_change, <br> refund_amount, refund_reason,<br> out_request_no, buyer_logon_id,<br> gmt_refund,<br> buyer_user_id, refund_detail_item_list;<br> 4、同步网商返回的状态,如果是网商银行的账单，bank_success这个字段必填 |
| `trade_no` | `TradeNo` | `*string` | JSON body | 未标注 | 支付宝返回的交易号 |

<a id="api-eduktbillingquery"></a>

### EduKtBillingQuery — 缴费账单查询

- 官方接口：[`alipay.eco.edu.kt.billing.query`](https://opendocs.alipay.com/open-v3/d9852f73_alipay.eco.edu.kt.billing.query)
- 路由：`GET /v3/alipay/eco/edu/kt/billing/query`
- 请求类型：[EduKtBillingQueryRequest](edu_request.go#L152)
- 响应类型：[EduKtBillingQueryResponse](edu_response.go#L11)

缴费账单查询

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `isv_pid` | `IsvPid` | `*string` | query | 否 | Isv pid |
| `school_pid` | `SchoolPid` | `*string` | query | 否 | 学校支付宝pid，直付通填写smid |
| `out_trade_no` | `OutTradeNo` | `*string` | query | 否 | ISV调用发送账单接口，返回给商户的order_no |

<a id="api-eduktbillingsend"></a>

### EduKtBillingSend — 教育缴费账单发送接口

- 官方接口：[`alipay.eco.edu.kt.billing.send`](https://opendocs.alipay.com/open-v3/7c14b5a1_alipay.eco.edu.kt.billing.send)
- 路由：`POST /v3/alipay/eco/edu/kt/billing/send`
- 请求类型：[EduKtBillingSendRequest](edu_request.go#L35)
- 响应类型：[EduKtBillingSendResponse](edu_response.go#L19)

商家发送教育缴费账单给孩子，家长在支付宝可以查看自己绑定的孩子的缴费账单。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `amount` | `Amount` | `*string` | JSON body | 未标注 | 总金额，单位为元，精确到小数点后两位，取值范围[0.01,100000000]，<br> 如果是非多选项，是要和缴费项的总和相同，多选模式不做验证 |
| `charge_bill_title` | `ChargeBillTitle` | `*string` | JSON body | 未标注 | 缴费账单名称 |
| `charge_item` | `ChargeItem` | [`*[]ChargeItems`](#model-chargeitems) | JSON body | 未标注 | 缴费详情：输入json格式字符串。Json定义：key填写缴费项名称，value填写缴费项金额，金额保留2位小数（单位：元） |
| `charge_type` | `ChargeType` | `*string` | JSON body | 未标注 | 缴费项模式：空或"N"，表示缴费项不可选，<br> "M"表示缴费项为可选 ，支持单选和多选。 |
| `child_name` | `ChildName` | `*string` | JSON body | 未标注 | 孩子名字 |
| `class_in` | `ClassIn` | `*string` | JSON body | 未标注 | 孩子所在班级 |
| `end_enable` | `EndEnable` | `*string` | JSON body | 未标注 | 截止日期是否生效，与gmt_end发布配合使用,N为gmt_end不生效，用户过期后仍可以缴费；Y为gmt_end生效，用户过期后，不能再缴费。 |
| `ext_info` | `ExtInfo` | [`*BillSendExtInfo`](#model-billsendextinfo) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `gmt_end` | `GmtEnd` | `*string` | JSON body | 未标注 | 缴费截止时间，格式"yyyy-MM-dd HH:mm:ss"，日期要大于当前时间。请注意，过期时间不宜设置过短。 |
| `grade` | `Grade` | `*string` | JSON body | 未标注 | 孩子所在年级 |
| `out_trade_no` | `OutTradeNo` | `*string` | JSON body | 未标注 | ISV端的缴费账单编号 |
| `partner_id` | `PartnerId` | `*string` | JSON body | 未标注 | Isv支付宝pid, 支付宝签约后，返回给ISV编号 |
| `school_no` | `SchoolNo` | `*string` | JSON body | 未标注 | 学校编码，录入学校接口返回的参数 |
| `school_pid` | `SchoolPid` | `*string` | JSON body | 未标注 | 学校支付宝pid,直付通填写smid |
| `student_code` | `StudentCode` | `*string` | JSON body | 未标注 | 学生的学号，只支持字母和数字类型，一般以教育局学号为准，作为学生的唯一标识。此字段与student_identify、家长user_mobile至少选一个 |
| `student_identify` | `StudentIdentify` | `*string` | JSON body | 未标注 | 学生的身份证号，如果ISV有学生身份证号，则同步身份证号作为学生唯一标识。此字段与student_code、家长user_mobile至少选一个。<br> 大陆身份证必须是18位 ， 其它地区或国家的身份证开头需要加"IC"开头区分并且不超过18位，但查询账单的时候不要带"IC" |
| `users` | `Users` | [`*[]UserDetails`](#model-userdetails) | JSON body | 未标注 | 孩子的家长信息，最多一次输入10个家长，此字段做为识别家长的孩子用，与student_identify、student_code至少选一个 |

<a id="api-eduktparentquery"></a>

### EduKtParentQuery — 查询学生家长状态接口

- 官方接口：[`alipay.eco.edu.kt.parent.query`](https://opendocs.alipay.com/apis/05x0kr)
- 路由：`GET /v3/alipay/eco/edu/kt/parent/query`
- 请求类型：[EduKtParentQueryRequest](edu_request.go#L162)
- 响应类型：[EduKtParentQueryResponse](edu_response.go#L25)

ISV发送账单后，用户在支付宝缴费成功。ISV可以通过此接口查询，用户缴费的时候，是否创建了缴费账户，如果已经添加，学校使用ISV的系统再次发账单就可以通过支付宝的通知触达用户；如果未添加，学校即可联系家长，督促家长在支付宝添加学生的缴费账户。
目前的出现的场景是家长第一次缴费成功后，学校使用ISV的系统再次发账单，认为家长应该收到通知，实际上只有建立缴费账户的家长才能收到，未建立缴费账户的是收不到通知，这部分家长有可能就缴费不及时或不通过支付宝缴费。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `school_pid` | `SchoolPid` | `*string` | query | 否 | 学校支付宝pid |
| `school_no` | `SchoolNo` | `*string` | query | 否 | 学校编码，录入学校接口返回的school_no参数 |
| `partner_id` | `PartnerId` | `*string` | query | 否 | Isv的支付宝pid |
| `child_name` | `ChildName` | `*string` | query | 否 | 孩子或学生姓名 |
| `user_mobile` | `UserMobile` | `*string` | query | 否 | 用户手机号，发账单时填写users数组中的一个手机号。结果返回用户是否通过此手机号添加此学生的缴费账户。user_mobile、student_code 、student_identify 必须并且只能填一项。 |
| `student_code` | `StudentCode` | `*string` | query | 否 | 学生的学号，一个学校的学号必须是唯一。结果返回用户是否通过此学号添加此学生的缴费账户。user_mobile、student_code 、student_identify 必须并且只能填一项。 |
| `student_identify` | `StudentIdentify` | `*string` | query | 否 | 学生的身份证号，使用身份证规则验证。结果返回用户是否通过此身份证号添加此学生的缴费账户。user_mobile、student_code 、student_identify 必须并且只能填一项。 |

<a id="api-eduktschoolinfomodify"></a>

### EduKtSchoolInfoModify — 教育缴费学校信息录入接口

- 官方接口：[`alipay.eco.edu.kt.schoolinfo.modify`](https://opendocs.alipay.com/open-v3/a319c0c9_alipay.eco.edu.kt.schoolinfo.modify)
- 路由：`POST /v3/alipay/eco/edu/kt/schoolinfo/modify`
- 请求类型：[EduKtSchoolInfoModifyRequest](edu_request.go#L72)
- 响应类型：[EduKtSchoolInfoModifyResponse](edu_response.go#L31)

ISV与支付宝签约后，通过此接口同步学校信息到中小学教育平台。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `bank_notify_url` | `BankNotifyUrl` | `*string` | JSON body | 未标注 | 与浙江网商交易见证平台有交互ISV输入网商交易异步通知回调URL，教育缴费同步账单信息给网商，网商会回调此url，ISV即可获取网商相关的参数，根据教育缴费平台账单发送接口返回的 order_no和网商返回的outer_trade_no来对应账单信息。 |
| `bank_partner_id` | `BankPartnerId` | `*string` | JSON body | 未标注 | 与浙江网商交易见证平台有交互的ISV,由交易见证平台分配给合作方业务平台 签约唯一ID ，由网商分配给ISV的渠道参数 |
| `bank_uid` | `BankUid` | `*string` | JSON body | 未标注 | 与浙江网商交易见证平台有交互的ISV在创建账户时的uid，也就是ISV平台上的用户ID（字母或数字） |
| `bankcard_no` | `BankcardNo` | `*string` | JSON body | 未标注 | 对应集团到卡模式的银行编号.学校与支付宝后台签约时，由学校提交给支付宝的编号 |
| `batch_no` | `BatchNo` | `*string` | JSON body | 未标注 | 集团收单模式：分账批次号，支付宝配置后提供的银行卡批次号 |
| `card_alias_no` | `CardAliasNo` | `*string` | JSON body | 未标注 | 学校开通直付通卡编号，smid与card_alias_no必须同时填写 |
| `city_code` | `CityCode` | `*string` | JSON body | 未标注 | 城市的国家编码。参见 [国家统计局出版的行政区划代码](http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/)。 |
| `city_name` | `CityName` | `*string` | JSON body | 未标注 | 城市名称 |
| `corporate_branch_pid` | `CorporateBranchPid` | `*string` | JSON body | 未标注 | 集团收单模式：BD批量上传银行卡信息后，支付宝系统分配给ISV的每个卡分配的唯一标识 |
| `district_code` | `DistrictCode` | `*string` | JSON body | 未标注 | 区县的国家编码。参见 [国家统计局出版的行政区划代码](http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/)。 |
| `district_name` | `DistrictName` | `*string` | JSON body | 未标注 | 区县名称 |
| `isv_name` | `IsvName` | `*string` | JSON body | 未标注 | ISV公司名称 ， 会在账单详情页面展示给用户 |
| `isv_no` | `IsvNo` | `*string` | JSON body | 未标注 | 注意：本参数于2017-08-03开始已经废弃，不再需要传递。 由支付宝提供的给已经签约的isv的编码，业务上一般直接采用isv的支付宝PID。 |
| `isv_notify_url` | `IsvNotifyUrl` | `*string` | JSON body | 未标注 | 此通知地址是为了保持教育缴费平台与ISV商户支付状态一致性。用户支付成功后，支付宝会根据本isv_notify_url(异步通知说明https://docs.open.alipay.com/203/105286/)，通过POST请求的形式将支付结果作为参数通知到商户系统，ISV商户可以根据返回的参数更新账单状态。 |
| `isv_phone` | `IsvPhone` | `*string` | JSON body | 未标注 | ISV的联系方式 ， 会在账单详情页面展示给用户，用户有问题可以直接联系此电话获取帮助 |
| `isv_pid` | `IsvPid` | `*string` | JSON body | 未标注 | 填写已经签约教育缴费的isv的支付宝PID |
| `province_code` | `ProvinceCode` | `*string` | JSON body | 未标注 | 省份的国家编码，参见 [国家统计局出版的行政区划代码](http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/)。 |
| `province_name` | `ProvinceName` | `*string` | JSON body | 未标注 | 省名称 |
| `school_icon` | `SchoolIcon` | `*string` | JSON body | 未标注 | 学校的校徽或logo，在学校展示页面作为学校的标识。该字段为图片的链接地址，只支持png或jpg图片格式，图片高度为108，宽度为108 ，不大于20k。<br> 注意：此链接要求公网可以访问，否则无法正常展示。 |
| `school_icon_type` | `SchoolIconType` | `*string` | JSON body | 未标注 | 如果填写了school_icon参数，则此字段不能为空。目前只支持png和jpg两种格式 |
| `school_name` | `SchoolName` | `*string` | JSON body | 未标注 | 学校名称 |
| `school_pid` | `SchoolPid` | `*string` | JSON body | 未标注 | 学校签约支付宝教育缴费支付宝pid,如果是直付通学校，填写直付通返回的smid |
| `school_stdcode` | `SchoolStdcode` | `*string` | JSON body | 未标注 | 学校(机构)标识码（由教育部按照国家标准及编码规则编制，可以在教育局官网查询） |
| `school_type` | `SchoolType` | `*string` | JSON body | 未标注 | 学校的类型。枚举值如下： *1：代表托儿所。 *2：代表幼儿园。 *3：代表小学。 *4：代表初中。 *5：代表高中。 注意：如果学校兼有多种属性，可以连写，比如：45：代表兼有初中部高中部；34：代表兼有小学部初中部。 |
| `smid` | `Smid` | `*string` | JSON body | 未标注 | 学校开通直付通返回的二级商户id，smid与card_alias_no必须同时填写 |
| `trans_in` | `TransIn` | `*string` | JSON body | 未标注 | 集团收单模式：分账批次号，支付宝配置后提供的银行卡批次号 |
| `white_channel_code` | `WhiteChannelCode` | `*string` | JSON body | 未标注 | 与浙江网商交易见证平台有交互的ISV,由网商分配给ISV的渠道参数 |

<a id="api-eduktstudentmodify"></a>

### EduKtStudentModify — 学生信息更新

- 官方接口：[`alipay.eco.edu.kt.student.modify`](https://opendocs.alipay.com/apis/05wun3)
- 路由：`POST /v3/alipay/eco/edu/kt/student/modify`
- 请求类型：[EduKtStudentModifyRequest](edu_request.go#L130)
- 响应类型：[EduKtStudentModifyResponse](edu_response.go#L39)

学生信息更新

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `child_name` | `ChildName` | `*string` | JSON body | 未标注 | 修改后的学生姓名<br> 本接口调用时，child_name、student_code、student_identify、users这几个参数至少需要填写其中一个，不能同时为空 |
| `isv_pid` | `IsvPid` | `*string` | JSON body | 未标注 | 已经签约教育缴费的isv的支付宝PID |
| `school_no` | `SchoolNo` | `*string` | JSON body | 未标注 | 学校编号，调用alipay.eco.edu.kt.schoolinfo.modify接口录入学校信息时，接口返回的编号 |
| `school_pid` | `SchoolPid` | `*string` | JSON body | 未标注 | 学校用来签约支付宝教育缴费的支付宝PID |
| `status` | `Status` | `*string` | JSON body | 未标注 | 区分ISV操作，“D”表示删除，“U”表示更新，区分大小写。<br> 如果为U，则学生名字，学号，身份证至少填写一项 |
| `student_code` | `StudentCode` | `*string` | JSON body | 未标注 | 修改后的学号<br> 本接口调用时，child_name、student_code、student_identify、users这几个参数至少需要填写其中一个，不能同时为空 |
| `student_identify` | `StudentIdentify` | `*string` | JSON body | 未标注 | 修改后的身份证号码<br> 本接口调用时，child_name、student_code、student_identify、users这几个参数至少需要填写其中一个，不能同时为空 |
| `student_no` | `StudentNo` | `*string` | JSON body | 未标注 | 支付宝-中小学-教育缴费生成的学生唯一编号 |
| `users` | `Users` | [`*[]UserDetails`](#model-userdetails) | JSON body | 未标注 | 孩子的家长信息，最多一次输入20个家长。如果输入的家长信息不存在，则该改学生增加家长信息<br> 本接口调用时，child_name、student_code、student_identify、users这几个参数至少需要填写其中一个，不能同时为空 |

## 嵌套请求参数

以下仅展开本目录请求中实际使用的对象；字段必填规则沿用官方快照，“未标注”须结合业务说明确认。响应模型请沿各接口的响应类型链接查看。

<a id="model-billsendextinfo"></a>

### BillSendExtInfo

类型定义：[BillSendExtInfo](edu_model.go#L5)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `ext_school_id` | `ExtSchoolId` | `*string` | 未标注 | 学校外标编号 |
| `order_pay_type` | `OrderPayType` | `*string` | 未标注 | 账单唤起收银台的方式： 1-手机wap;2-当面付。默认值为1（不传值或传非数字值） |
| `royalty_amount` | `RoyaltyAmount` | `*string` | 未标注 | 分账金额，单位：元 |
| `school_name` | `SchoolName` | `*string` | 未标注 | 学校名称 |
| `trans_in_pid` | `TransInPid` | `*string` | 未标注 | 分账收入户PID |

<a id="model-chargeitems"></a>

### ChargeItems

类型定义：[ChargeItems](edu_model.go#L19)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `item_mandatory` | `ItemMandatory` | `*string` | 未标注 | 缴费项是否必选<br> 如果缴费项是多选模式，此参数生效。 “Y”表示必填，“N”或空表示非必填。 |
| `item_maximum` | `ItemMaximum` | `*int64` | 未标注 | 缴费项最大可选数<br> 如果缴费项是多选模式，此参数生效，范围是1-9，如果为空，则最大项默认为9 |
| `item_name` | `ItemName` | `*string` | 未标注 | 缴费项名称 |
| `item_price` | `ItemPrice` | `*string` | 未标注 | 缴费项金额 |
| `item_serial_number` | `ItemSerialNumber` | `*int64` | 未标注 | 缴费项序号，如果缴费项是多选模式，此项为必填，建议从1开始的连续数字，<br> 用户支付成功后，通过passback_params参数带回已选择的缴费项。例如:orderNo=uoo234234&isvOrderNo=24werwe&items=1-2&#124;2-1&#124;3-5<br> 1-2&#124;2-1&#124;3-5 表示：缴费项序列号-缴费项数&#124;缴费项序列号-缴费项数 |

<a id="model-userdetails"></a>

### UserDetails

类型定义：[UserDetails](edu_model.go#L33)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `user_change_mobile` | `UserChangeMobile` | `*string` | 未标注 | 用户变更手机号，替换旧的手机号 |
| `user_mobile` | `UserMobile` | `*string` | 未标注 | 孩子家长的手机号 |
| `user_name` | `UserName` | `*string` | 未标注 | 张三 |
| `user_relation` | `UserRelation` | `*string` | 未标注 | 孩子与家长的关系： 1、爸爸 2、妈妈 3、爷爷 4、奶奶 5、外公 6、外婆 7、家长 |
