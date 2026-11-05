# 接口参数核查记录

2026-11-06 对当前工作树的 746 个 V3 REST 操作进行了字段核查和修复。744 个操作的请求和成功响应模型通过快照比对；另外两个操作各省略一个已废弃的令牌响应字段。

## 依据与范围

字段基准为支付宝发布的 [V3 OpenAPI 文件](https://github.com/alipay/alipay-sdk-java-all/blob/0552690f689f0ca46e34ad8af68ed897f65ee59c/v3/api/openapi.yaml)，版本为 2026-09-23，Git blob 为 `967e291810c3fa0e2247d31d9663c3b2311ebc93`。通过[官方索引](https://opendocs.alipay.com/common/llms.txt)和[接口目录](apis.md)定位具体产品文档。另核对了修复前存在差异的 49 个 REST 操作及三个支付产物的在线参数表。

从 Service 源码提取操作名、方法、路径、参数模式和 Go 类型，用编译后的 Go 类型递归对照快照的 query、路径、请求与成功响应 schema。处理指针、数组、内嵌结构、跨包共享类型及 multipart 文件。逐项区分快照字段覆盖和具体产品文档的业务条件。快照覆盖 740 条路由、746 个操作，含 22 个 multipart 操作的 88 个文件参数、17 个含路径参数操作以及 15 个 JSON 与 query 组合操作。

## 修复与取舍

| 项目 | 修复前的证据 | 当前行为 |
| --- | --- | --- |
| 结算时间 | 交易支付、创建、预创建和确认结算共用模型，SDK 曾把账期时间放进明细项。初次核查漏计了确认结算。 | 四个 REST 请求统一编码为 `settle_info.settle_period_time`。PC/WAP 支付产物采用各自文档确认的同一层级。[预创建](https://opendocs.alipay.com/open-v3/fa0c2141_alipay.trade.precreate) · [确认结算](https://opendocs.alipay.com/apis/0hp5md) |
| 指定支付方式 | 预创建曾将 `payment_type` 嵌在间联机构信息内。 | `PreCreateRequest.PaymentType` 在请求体顶层发送，支持官方值 `agent_pay`。[预创建](https://opendocs.alipay.com/open-v3/fa0c2141_alipay.trade.precreate) |
| 资金退回 | 内置请求曾发送 `refund_reason`，遗漏快照的 `business_params`，并允许省略场景码。 | 通过 `remark` 传备注，`biz_scene` 依据在线参数表校验。交易退款自己的 `refund_reason` 仍由交易退款接口定义。[资金退回](https://opendocs.alipay.com/open-v3/0c6ff10a_alipay.fund.trans.refund) |
| 历史余额 | SDK 曾向历史余额接口发送当前余额接口的 `bill_user_id`。 | 历史余额支持 `biz_date`、`biz_month`。当前余额按自身契约支持 `bill_user_id`。[历史余额](https://opendocs.alipay.com/open-v3/029a9c24_alipay.data.bill.balancehis.query) · [当前余额](https://opendocs.alipay.com/open-v3/9b66d9d8_alipay.data.bill.balance.query) |
| Query 类型 | 31 个操作的 54 处整数和 3 处布尔字段原为字符串指针。 | 分别使用 `*int64` 和 `*bool`；在线文档标必填的 41 处会在执行器接收请求前检查。显式零值和 false 会发送。树类目分页参数按文档成对传递。[企业员工分页](https://opendocs.alipay.com/apis/05udr4) |
| 业务字段 | 原有 17 个操作存在字段缺口，其中 12 个仅有缺失。交易退款 21 处缺口分属请求 10 处、响应 11 处。 | 已补齐交易、退款、转账和账单模型，包括小程序经营主体 `op_app_id`、请求和响应的退款币种、付款方与收款方的嵌套信息。具体权限和条件必填项以对应接口页为准。[交易创建](https://opendocs.alipay.com/open-v3/bf4eae5e_alipay.trade.create) · [交易退款](https://opendocs.alipay.com/open-v3/6b16d4a2_alipay.trade.refund) |
| 授权令牌 | `alipay.open.auth.token.app.expires_in` 的说明明确标为作废，`alipay.system.oauth.token.alipay_user_id` 已废弃。 | 两个字段保留为记录在案的快照差异。用户授权令牌自己的 `expires_in` 继续由 `system.OAuthTokenResponse` 提供。[应用授权令牌](https://opendocs.alipay.com/open-v3/4b00cc9c_alipay.open.auth.token.app) · [用户授权令牌](https://opendocs.alipay.com/open-v3/a9cc8a2c_alipay.system.oauth.token) |

## 页面与唤端产物

这三个操作单独核对在线参数表。PC 支付支持二维码模式 4 的 `qrcode_width`、二级商户、门店和开票信息，WAP 支付将 `quit_url` 放在业务 JSON 中。`return_url` 是两者的公共参数。页面业务扩展信息 `business_params` 按文档作为 JSON 文本发送。[电脑网站支付](https://opendocs.alipay.com/open-v3/2423fad5_alipay.trade.page.pay) · [手机网站支付](https://opendocs.alipay.com/open-v3/1a957be0_alipay.trade.wap.pay)

App 普通支付可省略产品码，周期扣款需传签约参数。SDK 通过 `AppPayRequest.AgreementSignParams` 生成 `biz_content.agreement_sign_params`。离线测试对订单串解码并独立复算 RSA2 签名。周期规则的 `is_trial_mode` 在官方参数表标为必填，同页说明又明确允许省略并按无试用期处理。SDK 依据说明接受省略，显式值限于 `Y` 和 `N`。本次修复覆盖上文列出的产物字段，其他可选能力以对应业务包的字段表为准。[App 支付](https://opendocs.alipay.com/open-v3/429e4d75_alipay.trade.app.pay) · [支付并签约场景](https://opendocs.alipay.com/open-v3/0jf33m)

## REST 逐接口结果（746 项）

接口的官方文档链接见[接口目录](apis.md)。✅ 表示本次核查的方法、路由、参数位置、字段名和基础类型与固定快照匹配。

### services/alios/open（1 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alios.open.auto.info.query` | GET `/v3/alios/open/auto/info/query` | ✅ |  |

### services/boss/fnc/invoice（3 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.boss.fnc.invoice.apply` | POST `/v3/alipay/boss/fnc/invoice/apply` | ✅ |  |
| `alipay.boss.fnc.invoice.batchquery` | POST `/v3/alipay/boss/fnc/invoice/batchquery` | ✅ |  |
| `alipay.boss.fnc.invoice.query` | GET `/v3/alipay/boss/fnc/invoice/query` | ✅ |  |

### services/boss/fnc/invoicereceipt（2 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.boss.fnc.invoicereceipt.batchquery` | POST `/v3/alipay/boss/fnc/invoicereceipt/batchquery` | ✅ |  |
| `alipay.boss.fnc.invoicereceipt.create` | POST `/v3/alipay/boss/fnc/invoicereceipt/create` | ✅ |  |

### services/boss/fnc/userinvoiceinfo（3 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.boss.fnc.userinvoiceinfo.create` | POST `/v3/alipay/boss/fnc/userinvoiceinfo/create` | ✅ |  |
| `alipay.boss.fnc.userinvoiceinfo.modify` | POST `/v3/alipay/boss/fnc/userinvoiceinfo/modify` | ✅ |  |
| `alipay.boss.fnc.userinvoiceinfo.query` | GET `/v3/alipay/boss/fnc/userinvoiceinfo/query` | ✅ |  |

### services/commerce/cityfacilitator（4 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.commerce.cityfacilitator.station.query` | GET `/v3/alipay/commerce/cityfacilitator/station/query` | ✅ |  |
| `alipay.commerce.cityfacilitator.voucher.batchquery` | POST `/v3/alipay/commerce/cityfacilitator/voucher/batchquery` | ✅ |  |
| `alipay.commerce.cityfacilitator.voucher.generate` | POST `/v3/alipay/commerce/cityfacilitator/voucher/generate` | ✅ |  |
| `alipay.commerce.cityfacilitator.voucher.refund` | POST `/v3/alipay/commerce/cityfacilitator/voucher/refund` | ✅ |  |

### services/commerce/ec/balance（2 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.commerce.ec.balance.downloadurl.query` | GET `/v3/alipay/commerce/ec/balance/downloadurl/query` | ✅ |  |
| `alipay.commerce.ec.balance.period.modify` | POST `/v3/alipay/commerce/ec/balance/period/modify` | ✅ |  |

### services/commerce/ec/consume（1 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.commerce.ec.consume.detail.query` | POST `/v3/alipay/commerce/ec/consume/detail/query` | ✅ |  |

### services/commerce/ec/department（5 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.commerce.ec.department.create` | POST `/v3/alipay/commerce/ec/department` | ✅ |  |
| `alipay.commerce.ec.department.delete` | POST `/v3/alipay/commerce/ec/department/delete` | ✅ |  |
| `alipay.commerce.ec.department.info.modify` | PUT `/v3/alipay/commerce/ec/department` | ✅ |  |
| `alipay.commerce.ec.department.info.query` | GET `/v3/alipay/commerce/ec/department/info/query` | ✅ |  |
| `alipay.commerce.ec.department.sublist.query` | GET `/v3/alipay/commerce/ec/department/sublist/query` | ✅ |  |

### services/commerce/ec/employee（9 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.commerce.ec.employee.add` | POST `/v3/alipay/commerce/ec/employee/add` | ✅ |  |
| `alipay.commerce.ec.employee.delete` | POST `/v3/alipay/commerce/ec/employee/delete` | ✅ |  |
| `alipay.commerce.ec.employee.idlist.query` | GET `/v3/alipay/commerce/ec/employee/idlist/query` | ✅ |  |
| `alipay.commerce.ec.employee.info.modify` | PUT `/v3/alipay/commerce/ec/employee/info` | ✅ |  |
| `alipay.commerce.ec.employee.info.query` | GET `/v3/alipay/commerce/ec/employee/info/query` | ✅ |  |
| `alipay.commerce.ec.employee.invite.query` | GET `/v3/alipay/commerce/ec/employee/invite` | ✅ |  |
| `alipay.commerce.ec.employee.title.create` | POST `/v3/alipay/commerce/ec/employee/title` | ✅ |  |
| `alipay.commerce.ec.employee.title.delete` | POST `/v3/alipay/commerce/ec/employee/title/delete` | ✅ |  |
| `alipay.commerce.ec.employee.title.modify` | PUT `/v3/alipay/commerce/ec/employee/title` | ✅ |  |

### services/commerce/ec/enterprise（9 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.commerce.ec.enterprise.address.add` | POST `/v3/alipay/commerce/ec/enterprise/address` | ✅ |  |
| `alipay.commerce.ec.enterprise.address.modify` | PUT `/v3/alipay/commerce/ec/enterprise/address` | ✅ |  |
| `alipay.commerce.ec.enterprise.address.query` | GET `/v3/alipay/commerce/ec/enterprise/address` | ✅ |  |
| `alipay.commerce.ec.enterprise.agreement.query` | GET `/v3/alipay/commerce/ec/enterprise/agreement` | ✅ |  |
| `alipay.commerce.ec.enterprise.create` | POST `/v3/alipay/commerce/ec/enterprise/create` | ✅ |  |
| `alipay.commerce.ec.enterprise.delete` | POST `/v3/alipay/commerce/ec/enterprise/delete` | ✅ |  |
| `alipay.commerce.ec.enterprise.info.modify` | PUT `/v3/alipay/commerce/ec/enterprise` | ✅ |  |
| `alipay.commerce.ec.enterprise.info.query` | GET `/v3/alipay/commerce/ec/enterprise/info` | ✅ |  |
| `alipay.commerce.ec.enterprise.unsign` | GET `/v3/alipay/commerce/ec/enterprise/unsignurl` | ✅ |  |

### services/commerce/ec/jointaccountbill（1 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.commerce.ec.jointaccountbill.detail.batchquery` | GET `/v3/alipay/commerce/ec/jointaccountbill/detail/batchquery` | ✅ |  |

### services/commerce/ec/user（1 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.commerce.ec.user.enterprise.query` | GET `/v3/alipay/commerce/ec/user/enterprise` | ✅ |  |

### services/commerce/educate（3 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.commerce.educate.campus.institutions.add` | POST `/v3/alipay/commerce/educate/campus/institutions/add` | ✅ |  |
| `alipay.commerce.educate.campus.institutions.query` | GET `/v3/alipay/commerce/educate/campus/institutions/query` | ✅ |  |
| `alipay.commerce.educate.schoolcard.order.sync` | POST `/v3/alipay/commerce/educate/schoolcard/order/sync` | ✅ |  |

### services/commerce/iot（3 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.commerce.iot.device.report.upload` | POST `/v3/alipay/commerce/iot/device/report/upload` | ✅ |  |
| `alipay.commerce.iot.sdarttool.message.query` | GET `/v3/alipay/commerce/iot/sdarttool/message/query` | ✅ |  |
| `alipay.commerce.iot.sdarttool.message.send` | POST `/v3/alipay/commerce/iot/sdarttool/message/send` | ✅ |  |

### services/commerce/logistics/invoice（2 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.commerce.logistics.invoice.istdwaybill.create` | POST `/v3/alipay/commerce/logistics/invoice/istdwaybill/create` | ✅ |  |
| `alipay.commerce.logistics.invoice.istdwaybill.query` | GET `/v3/alipay/commerce/logistics/invoice/istdwaybill/query` | ✅ |  |

### services/commerce/logistics/logisticscompany（1 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.commerce.logistics.logisticscompany.instantdelivery.query` | GET `/v3/alipay/commerce/logistics/logisticscompany/instantdelivery/query` | ✅ |  |

### services/commerce/logistics/order（5 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.commerce.logistics.order.instantdelivery.cancel` | POST `/v3/alipay/commerce/logistics/order/instantdelivery/cancel` | ✅ |  |
| `alipay.commerce.logistics.order.instantdelivery.create` | POST `/v3/alipay/commerce/logistics/order/instantdelivery/create` | ✅ |  |
| `alipay.commerce.logistics.order.instantdelivery.precreate` | POST `/v3/alipay/commerce/logistics/order/instantdelivery/precreate` | ✅ |  |
| `alipay.commerce.logistics.order.istdcancel.preconsult` | POST `/v3/alipay/commerce/logistics/order/istdcancel/preconsult` | ✅ |  |
| `alipay.commerce.logistics.order.istdretry.create` | POST `/v3/alipay/commerce/logistics/order/istdretry/create` | ✅ |  |

### services/commerce/logistics/waybill（1 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.commerce.logistics.waybill.istddetail.query` | GET `/v3/alipay/commerce/logistics/waybill/istddetail/query` | ✅ |  |

### services/commerce/operation（4 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.commerce.operation.activity.merchant.modify` | POST `/v3/alipay/commerce/operation/activity/merchant/modify` | ✅ |  |
| `alipay.commerce.operation.activity.merchant.sign` | POST `/v3/alipay/commerce/operation/activity/merchant/sign` | ✅ |  |
| `alipay.commerce.operation.activity.merchant.unsign` | POST `/v3/alipay/commerce/operation/activity/merchant/unsign` | ✅ |  |
| `alipay.commerce.operation.bank.activity.query` | GET `/v3/alipay/commerce/operation/bank/activity/query` | ✅ |  |

### services/commerce/transport（4 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.commerce.transport.charger.chargerbindinfo.sync` | POST `/v3/alipay/commerce/transport/charger/chargerbindinfo/sync` | ✅ |  |
| `alipay.commerce.transport.parking.enterinfo.sync` | POST `/v3/alipay/commerce/transport/parking/enterinfo/sync` | ✅ |  |
| `alipay.commerce.transport.parking.exitinfo.sync` | POST `/v3/alipay/commerce/transport/parking/exitinfo/sync` | ✅ |  |
| `alipay.commerce.transport.parking.paymentinfo.sync` | POST `/v3/alipay/commerce/transport/parking/paymentinfo/sync` | ✅ |  |

### services/data/bill（15 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.data.bill.accountbookereceipt.apply` | POST `/v3/alipay/data/bill/accountbookereceipt/apply` | ✅ |  |
| `alipay.data.bill.accountbookereceipt.query` | GET `/v3/alipay/data/bill/accountbookereceipt/query` | ✅ |  |
| `alipay.data.bill.accountlog.query` | GET `/v3/alipay/data/bill/accountlog/query` | ✅ |  |
| `alipay.data.bill.bail.query` | GET `/v3/alipay/data/bill/bail/query` | ✅ |  |
| `alipay.data.bill.balance.query` | GET `/v3/alipay/data/bill/balance/query` | ✅ |  |
| `alipay.data.bill.balancehis.query` | GET `/v3/alipay/data/bill/balancehis/query` | ✅ |  |
| `alipay.data.bill.bizfundagent.query` | GET `/v3/alipay/data/bill/bizfundagent/query` | ✅ |  |
| `alipay.data.bill.buy.query` | GET `/v3/alipay/data/bill/buy/query` | ✅ |  |
| `alipay.data.bill.ereceipt.apply` | POST `/v3/alipay/data/bill/ereceipt/apply` | ✅ |  |
| `alipay.data.bill.ereceipt.query` | GET `/v3/alipay/data/bill/ereceipt/query` | ✅ |  |
| `alipay.data.bill.ereceiptagent.apply` | POST `/v3/alipay/data/bill/ereceiptagent/apply` | ✅ |  |
| `alipay.data.bill.sell.query` | GET `/v3/alipay/data/bill/sell/query` | ✅ |  |
| `alipay.data.bill.transfer.query` | GET `/v3/alipay/data/bill/transfer/query` | ✅ |  |
| `alipay.data.bill.transferaccountbook.query` | GET `/v3/alipay/data/bill/transferaccountbook/query` | ✅ |  |
| `alipay.data.dataservice.bill.downloadurl.query` | GET `/v3/alipay/data/dataservice/bill/downloadurl/query` | ✅ |  |

### services/data/dataservice/ad（3 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.data.dataservice.ad.conversion.upload` | POST `/v3/alipay/data/dataservice/ad/conversion/upload` | ✅ |  |
| `alipay.data.dataservice.ad.promotepage.batchquery` | GET `/v3/alipay/data/dataservice/ad/promotepage/batchquery` | ✅ |  |
| `alipay.data.dataservice.ad.promotepage.download` | GET `/v3/alipay/data/dataservice/ad/promotepage/download` | ✅ |  |

### services/ebpp/bill（2 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.ebpp.bill.add` | POST `/v3/alipay/ebpp/bill/add` | ✅ |  |
| `alipay.ebpp.bill.get` | GET `/v3/alipay/ebpp/bill/get` | ✅ |  |

### services/ebpp/invoice/apply（3 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.ebpp.invoice.apply.inv.upload` | PUT `/v3/alipay/ebpp/invoice/apply/inv/upload` | ✅ |  |
| `alipay.ebpp.invoice.apply.result.sync` | POST `/v3/alipay/ebpp/invoice/apply/result/sync` | ✅ |  |
| `alipay.ebpp.invoice.apply.status.notify` | PUT `/v3/alipay/ebpp/invoice/apply/status/notify` | ✅ |  |

### services/ebpp/invoice/applystatus（1 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.ebpp.invoice.applystatus.query` | POST `/v3/alipay/ebpp/invoice/applystatus/query` | ✅ |  |

### services/ebpp/invoice/auth（2 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.ebpp.invoice.auth.sign` | POST `/v3/alipay/ebpp/invoice/auth/sign` | ✅ |  |
| `alipay.ebpp.invoice.auth.unsign` | POST `/v3/alipay/ebpp/invoice/auth/unsign` | ✅ |  |

### services/ebpp/invoice/detail（1 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.ebpp.invoice.detail.output.query` | POST `/v3/alipay/ebpp/invoice/detail/output/query` | ✅ |  |

### services/ebpp/invoice/ecorder（1 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.ebpp.invoice.ecorder.order.query` | GET `/v3/alipay/ebpp/invoice/ecorder/order/query` | ✅ |  |

### services/ebpp/invoice/einvpackage（1 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.ebpp.invoice.einvpackage.query` | GET `/v3/alipay/ebpp/invoice/einvpackage/query` | ✅ |  |

### services/ebpp/invoice/enterprise（1 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.ebpp.invoice.enterprise.merchantrelation.create` | POST `/v3/alipay/ebpp/invoice/enterprise/merchantrelation/create` | ✅ |  |

### services/ebpp/invoice/enterpriseconsume（8 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.ebpp.invoice.enterpriseconsume.consume.batchquery` | POST `/v3/alipay/ebpp/invoice/enterpriseconsume/consume/batchquery` | ✅ |  |
| `alipay.ebpp.invoice.enterpriseconsume.detail.query` | GET `/v3/alipay/ebpp/invoice/enterpriseconsume/detail/query` | ✅ |  |
| `alipay.ebpp.invoice.enterpriseconsume.enterpriseopenrule.create` | POST `/v3/alipay/ebpp/invoice/enterpriseconsume/enterpriseopenrule/create` | ✅ |  |
| `alipay.ebpp.invoice.enterpriseconsume.enterpriseopenrule.modify` | POST `/v3/alipay/ebpp/invoice/enterpriseconsume/enterpriseopenrule/modify` | ✅ |  |
| `alipay.ebpp.invoice.enterpriseconsume.enterpriseopenrule.query` | GET `/v3/alipay/ebpp/invoice/enterpriseconsume/enterpriseopenrule/query` | ✅ |  |
| `alipay.ebpp.invoice.enterpriseconsume.openrulerelation.create` | POST `/v3/alipay/ebpp/invoice/enterpriseconsume/openrulerelation/create` | ✅ |  |
| `alipay.ebpp.invoice.enterpriseconsume.relatedetail.query` | GET `/v3/alipay/ebpp/invoice/enterpriseconsume/relatedetail/query` | ✅ |  |
| `alipay.ebpp.invoice.enterpriseconsume.summaryinvoicedetail.query` | GET `/v3/alipay/ebpp/invoice/enterpriseconsume/summaryinvoicedetail/query` | ✅ |  |

### services/ebpp/invoice/enterpriseexctrl（4 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.ebpp.invoice.enterpriseexctrl.employertitle.batchquery` | POST `/v3/alipay/ebpp/invoice/enterpriseexctrl/employertitle/batchquery` | ✅ |  |
| `alipay.ebpp.invoice.enterpriseexctrl.employertitle.create` | POST `/v3/alipay/ebpp/invoice/enterpriseexctrl/employertitle/create` | ✅ |  |
| `alipay.ebpp.invoice.enterpriseexctrl.employertitle.modify` | PUT `/v3/alipay/ebpp/invoice/enterpriseexctrl/employertitle` | ✅ |  |
| `alipay.ebpp.invoice.enterpriseexctrl.employertitle.query` | GET `/v3/alipay/ebpp/invoice/enterpriseexctrl/employertitle/query` | ✅ |  |

### services/ebpp/invoice/expense（1 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.ebpp.invoice.expense.progress.sync` | POST `/v3/alipay/ebpp/invoice/expense/progress/sync` | ✅ |  |

### services/ebpp/invoice/expensecomsue（1 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.ebpp.invoice.expensecomsue.outsource.notify` | POST `/v3/alipay/ebpp/invoice/expensecomsue/outsource/notify` | ✅ |  |

### services/ebpp/invoice/expensecontrol（6 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.ebpp.invoice.expensecontrol.aggregation.create` | POST `/v3/alipay/ebpp/invoice/expensecontrol/aggregation/create` | ✅ |  |
| `alipay.ebpp.invoice.expensecontrol.issuebatch.cancel` | POST `/v3/alipay/ebpp/invoice/expensecontrol/issuebatch/cancel` | ✅ |  |
| `alipay.ebpp.invoice.expensecontrol.issuebatch.create` | POST `/v3/alipay/ebpp/invoice/expensecontrol/issuebatch/create` | ✅ |  |
| `alipay.ebpp.invoice.expensecontrol.quota.create` | POST `/v3/alipay/ebpp/invoice/expensecontrol/quota/create` | ✅ |  |
| `alipay.ebpp.invoice.expensecontrol.quota.modify` | POST `/v3/alipay/ebpp/invoice/expensecontrol/quota/modify` | ✅ |  |
| `alipay.ebpp.invoice.expensecontrol.quota.query` | POST `/v3/alipay/ebpp/invoice/expensecontrol/quota/query` | ✅ |  |

### services/ebpp/invoice/expenserules/employee（3 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.ebpp.invoice.expenserules.employee.query` | GET `/v3/alipay/ebpp/invoice/expenserules/employee/query` | ✅ |  |
| `alipay.ebpp.invoice.expenserules.employeerules.query` | GET `/v3/alipay/ebpp/invoice/expenserules/employeerules/query` | ✅ |  |
| `alipay.ebpp.invoice.expenserules.groupemployee.modify` | POST `/v3/alipay/ebpp/invoice/expenserules/groupemployee/modify` | ✅ |  |

### services/ebpp/invoice/expenserules/project（5 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.ebpp.invoice.expenserules.projectemployee.modify` | POST `/v3/alipay/ebpp/invoice/expenserules/projectemployee/modify` | ✅ |  |
| `alipay.ebpp.invoice.expenserules.projectinfo.modify` | POST `/v3/alipay/ebpp/invoice/expenserules/projectinfo/modify` | ✅ |  |
| `alipay.ebpp.invoice.expenserules.projectrule.create` | POST `/v3/alipay/ebpp/invoice/expenserules/projectrule/create` | ✅ |  |
| `alipay.ebpp.invoice.expenserules.projectrule.query` | GET `/v3/alipay/ebpp/invoice/expenserules/projectrule/query` | ✅ |  |
| `alipay.ebpp.invoice.expenserules.projectrules.modify` | POST `/v3/alipay/ebpp/invoice/expenserules/projectrules/modify` | ✅ |  |

### services/ebpp/invoice/expenserules/scene（3 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.ebpp.invoice.expenserules.scenerule.create` | POST `/v3/alipay/ebpp/invoice/expenserules/scenerule/create` | ✅ |  |
| `alipay.ebpp.invoice.expenserules.scenerule.modify` | POST `/v3/alipay/ebpp/invoice/expenserules/scenerule/modify` | ✅ |  |
| `alipay.ebpp.invoice.expenserules.scenerule.query` | GET `/v3/alipay/ebpp/invoice/expenserules/scenerule/query` | ✅ |  |

### services/ebpp/invoice/expensescene（1 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.ebpp.invoice.expensescene.merchant.query` | GET `/v3/alipay/ebpp/invoice/expensescene/merchant/query` | ✅ |  |

### services/ebpp/invoice/file（1 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.ebpp.invoice.file.output.query` | GET `/v3/alipay/ebpp/invoice/file/output/query` | ✅ |  |

### services/ebpp/invoice/info（1 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.ebpp.invoice.info.send` | POST `/v3/alipay/ebpp/invoice/info/send` | ✅ |  |

### services/ebpp/invoice/institution（10 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.ebpp.invoice.institution.create` | POST `/v3/alipay/ebpp/invoice/institution/create` | ✅ |  |
| `alipay.ebpp.invoice.institution.delete` | DELETE `/v3/alipay/ebpp/invoice/institution/delete` | ✅ |  |
| `alipay.ebpp.invoice.institution.detailinfo.query` | GET `/v3/alipay/ebpp/invoice/institution/detailinfo/query` | ✅ |  |
| `alipay.ebpp.invoice.institution.expenserule.create` | POST `/v3/alipay/ebpp/invoice/institution/expenserule/create` | ✅ |  |
| `alipay.ebpp.invoice.institution.expenserule.delete` | POST `/v3/alipay/ebpp/invoice/institution/expenserule/delete` | ✅ |  |
| `alipay.ebpp.invoice.institution.expenserule.modify` | POST `/v3/alipay/ebpp/invoice/institution/expenserule/modify` | ✅ |  |
| `alipay.ebpp.invoice.institution.modify` | POST `/v3/alipay/ebpp/invoice/institution/modify` | ✅ |  |
| `alipay.ebpp.invoice.institution.pageinfo.query` | GET `/v3/alipay/ebpp/invoice/institution/pageinfo/query` | ✅ |  |
| `alipay.ebpp.invoice.institution.scope.modify` | POST `/v3/alipay/ebpp/invoice/institution/scope/modify` | ✅ |  |
| `alipay.ebpp.invoice.institution.scopepageinfo.query` | GET `/v3/alipay/ebpp/invoice/institution/scopepageinfo/query` | ✅ |  |

### services/ebpp/invoice/issuerule（3 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.ebpp.invoice.issuerule.create` | POST `/v3/alipay/ebpp/invoice/issuerule/create` | ✅ |  |
| `alipay.ebpp.invoice.issuerule.delete` | POST `/v3/alipay/ebpp/invoice/issuerule/delete` | ✅ |  |
| `alipay.ebpp.invoice.issuerule.modify` | POST `/v3/alipay/ebpp/invoice/issuerule/modify` | ✅ |  |

### services/ebpp/invoice/isvtoken（1 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.ebpp.invoice.isvtoken.reim.apply` | POST `/v3/alipay/ebpp/invoice/isvtoken/reim/apply` | ✅ |  |

### services/ebpp/invoice/list（1 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.ebpp.invoice.list.expense.sync` | POST `/v3/alipay/ebpp/invoice/list/expense/sync` | ✅ |  |

### services/ebpp/invoice/merchant（1 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.ebpp.invoice.merchant.enterstatus.query` | GET `/v3/alipay/ebpp/invoice/merchant/enterstatus/query` | ✅ |  |

### services/ebpp/invoice/merchantlist（1 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.ebpp.invoice.merchantlist.enter.apply` | POST `/v3/alipay/ebpp/invoice/merchantlist/enter/apply` | ✅ |  |

### services/ebpp/invoice/order（1 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.ebpp.invoice.order.query` | GET `/v3/alipay/ebpp/invoice/order/query` | ✅ |  |

### services/ebpp/invoice/sycn（1 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.ebpp.invoice.sycn` | POST `/v3/alipay/ebpp/invoice/sycn` | ✅ |  |

### services/ebpp/invoice/sync（1 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.ebpp.invoice.sync.simple.send` | POST `/v3/alipay/ebpp/invoice/sync/simple/send` | ✅ |  |

### services/ebpp/invoice/taxno（1 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.ebpp.invoice.taxno.batchquery` | POST `/v3/alipay/ebpp/invoice/taxno/batchquery` | ✅ |  |

### services/ebpp/invoice/title（2 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.ebpp.invoice.title.dynamic.get` | GET `/v3/alipay/ebpp/invoice/title/dynamic/get` | ✅ |  |
| `alipay.ebpp.invoice.title.list.get` | GET `/v3/alipay/ebpp/invoice/title/list/get` | ✅ |  |

### services/ebpp/invoice/token（1 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.ebpp.invoice.token.batchquery` | GET `/v3/alipay/ebpp/invoice/token/batchquery` | ✅ |  |

### services/ebpp/pdeduct（3 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.ebpp.pdeduct.async.pay` | POST `/v3/alipay/ebpp/pdeduct/async/pay` | ✅ |  |
| `alipay.ebpp.pdeduct.bill.pay.status` | GET `/v3/alipay/ebpp/pdeduct/bill/pay/status` | ✅ |  |
| `alipay.ebpp.pdeduct.pay` | POST `/v3/alipay/ebpp/pdeduct/pay` | ✅ |  |

### services/ebpp/pdeduct/sign（4 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.ebpp.pdeduct.sign.add` | POST `/v3/alipay/ebpp/pdeduct/sign/add` | ✅ |  |
| `alipay.ebpp.pdeduct.sign.cancel` | POST `/v3/alipay/ebpp/pdeduct/sign/cancel` | ✅ |  |
| `alipay.ebpp.pdeduct.sign.query` | GET `/v3/alipay/ebpp/pdeduct/sign/query` | ✅ |  |
| `alipay.ebpp.pdeduct.sign.validate` | POST `/v3/alipay/ebpp/pdeduct/sign/validate` | ✅ |  |

### services/eco/contract（1 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.eco.contract.signflows.create` | POST `/v3/alipay/eco/contract/signflows/create` | ✅ |  |

### services/eco/doc（1 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.eco.doc.template.create` | POST `/v3/alipay/eco/doc/template/create` | ✅ |  |

### services/eco/doctemplate（1 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.eco.doctemplate.settingurl.query` | GET `/v3/alipay/eco/doctemplate/settingurl/query` | ✅ |  |

### services/eco/edu（6 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.eco.edu.kt.billing.modify` | POST `/v3/alipay/eco/edu/kt/billing/modify` | ✅ |  |
| `alipay.eco.edu.kt.billing.query` | GET `/v3/alipay/eco/edu/kt/billing/query` | ✅ |  |
| `alipay.eco.edu.kt.billing.send` | POST `/v3/alipay/eco/edu/kt/billing/send` | ✅ |  |
| `alipay.eco.edu.kt.parent.query` | GET `/v3/alipay/eco/edu/kt/parent/query` | ✅ |  |
| `alipay.eco.edu.kt.schoolinfo.modify` | POST `/v3/alipay/eco/edu/kt/schoolinfo/modify` | ✅ |  |
| `alipay.eco.edu.kt.student.modify` | POST `/v3/alipay/eco/edu/kt/student/modify` | ✅ |  |

### services/eco/file（1 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.eco.file.path.query` | GET `/v3/alipay/eco/file/path/query` | ✅ |  |

### services/eco/mycar/parking（16 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.eco.mycar.parking.chargeinfo.sync` | POST `/v3/alipay/eco/mycar/parking/chargeinfo/sync` | ✅ |  |
| `alipay.eco.mycar.parking.charginginfo.sync` | POST `/v3/alipay/eco/mycar/parking/charginginfo/sync` | ✅ |  |
| `alipay.eco.mycar.parking.config.query` | GET `/v3/alipay/eco/mycar/parking/config/query` | ✅ |  |
| `alipay.eco.mycar.parking.config.set` | POST `/v3/alipay/eco/mycar/parking/config/set` | ✅ |  |
| `alipay.eco.mycar.parking.enterinfo.sync` | POST `/v3/alipay/eco/mycar/parking/enterinfo/sync` | ✅ |  |
| `alipay.eco.mycar.parking.exitinfo.sync` | POST `/v3/alipay/eco/mycar/parking/exitinfo/sync` | ✅ |  |
| `alipay.eco.mycar.parking.order.sync` | POST `/v3/alipay/eco/mycar/parking/order/sync` | ✅ |  |
| `alipay.eco.mycar.parking.order.update` | POST `/v3/alipay/eco/mycar/parking/order/update` | ✅ |  |
| `alipay.eco.mycar.parking.overtimecharginginfo.sync` | POST `/v3/alipay/eco/mycar/parking/overtimecharginginfo/sync` | ✅ |  |
| `alipay.eco.mycar.parking.parkinglotbiz.transfer` | POST `/v3/alipay/eco/mycar/parking/parkinglotbiz/transfer` | ✅ |  |
| `alipay.eco.mycar.parking.parkinglotinfo.create` | POST `/v3/alipay/eco/mycar/parking/parkinglotinfo/create` | ✅ |  |
| `alipay.eco.mycar.parking.parkinglotinfo.query` | GET `/v3/alipay/eco/mycar/parking/parkinglotinfo/query` | ✅ |  |
| `alipay.eco.mycar.parking.parkinglotinfo.update` | POST `/v3/alipay/eco/mycar/parking/parkinglotinfo/update` | ✅ |  |
| `alipay.eco.mycar.parking.paymentinfo.sync` | POST `/v3/alipay/eco/mycar/parking/paymentinfo/sync` | ✅ |  |
| `alipay.eco.mycar.parking.spaceinfo.sync` | POST `/v3/alipay/eco/mycar/parking/spaceinfo/sync` | ✅ |  |
| `alipay.eco.mycar.parking.vehicle.query` | GET `/v3/alipay/eco/mycar/parking/vehicle/query` | ✅ |  |

### services/eco/sign（4 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.eco.sign.flow.cancel` | POST `/v3/alipay/eco/sign/flow/cancel` | ✅ |  |
| `alipay.eco.sign.flow.create` | POST `/v3/alipay/eco/sign/flow/create` | ✅ |  |
| `alipay.eco.sign.flow.finish` | POST `/v3/alipay/eco/sign/flow/finish` | ✅ |  |
| `alipay.eco.sign.flow.query` | GET `/v3/alipay/eco/sign/flow/query` | ✅ |  |

### services/eco/signflows（2 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.eco.signflows.detail.query` | GET `/v3/alipay/eco/signflows/detail/query` | ✅ |  |
| `alipay.eco.signflows.url.query` | GET `/v3/alipay/eco/signflows/url/query` | ✅ |  |

### services/fund/account（1 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.fund.account.query` | GET `/v3/alipay/fund/account/query` | ✅ |  |

### services/fund/accountbook（5 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.fund.accountbook.create` | POST `/v3/alipay/fund/accountbook/create` | ✅ |  |
| `alipay.fund.accountbook.notify.query` | GET `/v3/alipay/fund/accountbook/notify/query` | ✅ |  |
| `alipay.fund.accountbook.notify.subscribe` | POST `/v3/alipay/fund/accountbook/notify/subscribe` | ✅ |  |
| `alipay.fund.accountbook.notify.unsubscribe` | POST `/v3/alipay/fund/accountbook/notify/unsubscribe` | ✅ |  |
| `alipay.fund.accountbook.query` | GET `/v3/alipay/fund/accountbook/query` | ✅ |  |

### services/fund/agreement（2 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.fund.agreement.quota.modify` | POST `/v3/alipay/fund/agreement/quota/modify` | ✅ |  |
| `alipay.fund.agreement.quota.query` | POST `/v3/alipay/fund/agreement/quota/query` | ✅ |  |

### services/fund/auth（6 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.fund.auth.operation.cancel` | POST `/v3/alipay/fund/auth/operation/cancel` | ✅ |  |
| `alipay.fund.auth.operation.detail.query` | POST `/v3/alipay/fund/auth/operation/detail/query` | ✅ |  |
| `alipay.fund.auth.order.app.freeze` | POST `/v3/alipay/fund/auth/order/app/freeze` | ✅ |  |
| `alipay.fund.auth.order.freeze` | POST `/v3/alipay/fund/auth/order/freeze` | ✅ |  |
| `alipay.fund.auth.order.unfreeze` | POST `/v3/alipay/fund/auth/order/unfreeze` | ✅ |  |
| `alipay.fund.auth.order.voucher.create` | POST `/v3/alipay/fund/auth/order/voucher/create` | ✅ |  |

### services/fund/enterprisepay（7 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.fund.enterprisepay.group.add` | POST `/v3/alipay/fund/enterprisepay/group/add` | ✅ |  |
| `alipay.fund.enterprisepay.group.delete` | DELETE `/v3/alipay/fund/enterprisepay/group/delete` | ✅ |  |
| `alipay.fund.enterprisepay.group.modify` | POST `/v3/alipay/fund/enterprisepay/group/modify` | ✅ |  |
| `alipay.fund.enterprisepay.group.query` | GET `/v3/alipay/fund/enterprisepay/group/query` | ✅ |  |
| `alipay.fund.enterprisepay.member.modify` | POST `/v3/alipay/fund/enterprisepay/member/modify` | ✅ |  |
| `alipay.fund.enterprisepay.sign` | POST `/v3/alipay/fund/enterprisepay/sign` | ✅ |  |
| `alipay.fund.enterprisepay.unsign` | POST `/v3/alipay/fund/enterprisepay/unsign` | ✅ |  |

### services/fund/jointaccount（13 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.fund.jointaccount.bill.query` | GET `/v3/alipay/fund/jointaccount/bill/query` | ✅ |  |
| `alipay.fund.jointaccount.detail.query` | GET `/v3/alipay/fund/jointaccount/detail/query` | ✅ |  |
| `alipay.fund.jointaccount.list.query` | GET `/v3/alipay/fund/jointaccount/list/query` | ✅ |  |
| `alipay.fund.jointaccount.member.batchquery` | POST `/v3/alipay/fund/jointaccount/member/batchquery` | ✅ |  |
| `alipay.fund.jointaccount.member.bind` | POST `/v3/alipay/fund/jointaccount/member/bind` | ✅ |  |
| `alipay.fund.jointaccount.member.consult` | POST `/v3/alipay/fund/jointaccount/member/consult` | ✅ |  |
| `alipay.fund.jointaccount.member.query` | GET `/v3/alipay/fund/jointaccount/member/query` | ✅ |  |
| `alipay.fund.jointaccount.member.unbind` | POST `/v3/alipay/fund/jointaccount/member/unbind` | ✅ |  |
| `alipay.fund.jointaccount.memberrule.modify` | POST `/v3/alipay/fund/jointaccount/memberrule/modify` | ✅ |  |
| `alipay.fund.jointaccount.quota.query` | GET `/v3/alipay/fund/jointaccount/quota/query` | ✅ |  |
| `alipay.fund.jointaccount.rule.modify` | POST `/v3/alipay/fund/jointaccount/rule/modify` | ✅ |  |
| `alipay.fund.jointaccount.sign.query` | GET `/v3/alipay/fund/jointaccount/sign/query` | ✅ |  |
| `alipay.fund.jointaccount.trade.query` | GET `/v3/alipay/fund/jointaccount/trade/query` | ✅ |  |

### services/fund/trans（7 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.fund.trans.common.query` | GET `/v3/alipay/fund/trans/common/query` | ✅ |  |
| `alipay.fund.trans.order.query` | GET `/v3/alipay/fund/trans/order/query` | ✅ |  |
| `alipay.fund.trans.pay` | POST `/v3/alipay/fund/trans/pay` | ✅ |  |
| `alipay.fund.trans.payee.bind.query` | GET `/v3/alipay/fund/trans/payee/bind/query` | ✅ |  |
| `alipay.fund.trans.refund` | POST `/v3/alipay/fund/trans/refund` | ✅ |  |
| `alipay.fund.trans.toaccount.transfer` | POST `/v3/alipay/fund/trans/toaccount/transfer` | ✅ |  |
| `alipay.fund.trans.uni.transfer` | POST `/v3/alipay/fund/trans/uni/transfer` | ✅ |  |

### services/fund/usertrade（1 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.fund.usertrade.batchquery` | POST `/v3/alipay/fund/usertrade/batchquery` | ✅ |  |

### services/iservice/ccm/agent（6 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.iservice.ccm.agent.create` | POST `/v3/alipay/iservice/ccm/agent/create` | ✅ |  |
| `alipay.iservice.ccm.agent.delete` | DELETE `/v3/alipay/iservice/ccm/agent/delete` | ✅ |  |
| `alipay.iservice.ccm.agent.get` | GET `/v3/alipay/iservice/ccm/agent/get` | ✅ |  |
| `alipay.iservice.ccm.agent.modify` | POST `/v3/alipay/iservice/ccm/agent/modify` | ✅ |  |
| `alipay.iservice.ccm.agent.query` | POST `/v3/alipay/iservice/ccm/agent/query` | ✅ |  |
| `alipay.iservice.ccm.agent.schedulelog.query` | POST `/v3/alipay/iservice/ccm/agent/schedulelog/query` | ✅ |  |

### services/iservice/ccm/function（1 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.iservice.ccm.function.query` | GET `/v3/alipay/iservice/ccm/function/query` | ✅ |  |

### services/iservice/ccm/instance（3 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.iservice.ccm.instance.create` | POST `/v3/alipay/iservice/ccm/instance/create` | ✅ |  |
| `alipay.iservice.ccm.instance.get` | GET `/v3/alipay/iservice/ccm/instance/get` | ✅ |  |
| `alipay.iservice.ccm.instance.query` | GET `/v3/alipay/iservice/ccm/instance/query` | ✅ |  |

### services/iservice/ccm/isv（1 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.iservice.ccm.isv.initialize` | POST `/v3/alipay/iservice/ccm/isv/initialize` | ✅ |  |

### services/iservice/ccm/ols（1 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.iservice.ccm.ols.chatrecord.query` | GET `/v3/alipay/iservice/ccm/ols/chatrecord/query` | ✅ |  |

### services/iservice/ccm/robot（1 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.iservice.ccm.robot.avatarbase.query` | POST `/v3/alipay/iservice/ccm/robot/avatarbase/query` | ✅ |  |

### services/iservice/ccm/role（5 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.iservice.ccm.role.create` | POST `/v3/alipay/iservice/ccm/role/create` | ✅ |  |
| `alipay.iservice.ccm.role.delete` | DELETE `/v3/alipay/iservice/ccm/role/delete` | ✅ |  |
| `alipay.iservice.ccm.role.get` | GET `/v3/alipay/iservice/ccm/role/get` | ✅ |  |
| `alipay.iservice.ccm.role.modify` | POST `/v3/alipay/iservice/ccm/role/modify` | ✅ |  |
| `alipay.iservice.ccm.role.page.query` | GET `/v3/alipay/iservice/ccm/role/page/query` | ✅ |  |

### services/iservice/ccm/service（3 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.iservice.ccm.service.buy` | POST `/v3/alipay/iservice/ccm/service/buy` | ✅ |  |
| `alipay.iservice.ccm.service.close` | POST `/v3/alipay/iservice/ccm/service/close` | ✅ |  |
| `alipay.iservice.ccm.service.initialize` | POST `/v3/alipay/iservice/ccm/service/initialize` | ✅ |  |

### services/iservice/ccm/servicerecord（2 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.iservice.ccm.servicerecord.create` | POST `/v3/alipay/iservice/ccm/servicerecord/create` | ✅ |  |
| `alipay.iservice.ccm.servicerecord.modify` | POST `/v3/alipay/iservice/ccm/servicerecord/modify` | ✅ |  |

### services/iservice/ccm/sw/article（5 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.iservice.ccm.sw.article.batchquery` | POST `/v3/alipay/iservice/ccm/sw/article/batchquery` | ✅ |  |
| `alipay.iservice.ccm.sw.article.create` | PUT `/v3/alipay/iservice/ccm/sw/article/create` | ✅ |  |
| `alipay.iservice.ccm.sw.article.delete` | DELETE `/v3/alipay/iservice/ccm/sw/article/delete` | ✅ |  |
| `alipay.iservice.ccm.sw.article.get` | GET `/v3/alipay/iservice/ccm/sw/article/get` | ✅ |  |
| `alipay.iservice.ccm.sw.article.modify` | PATCH `/v3/alipay/iservice/ccm/sw/article/modify` | ✅ |  |

### services/iservice/ccm/sw/knowledge（1 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.iservice.ccm.sw.knowledge.modify` | POST `/v3/alipay/iservice/ccm/sw/knowledge/modify` | ✅ |  |

### services/iservice/ccm/sw/library（4 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.iservice.ccm.sw.library.batchquery` | GET `/v3/alipay/iservice/ccm/sw/library/batchquery` | ✅ |  |
| `alipay.iservice.ccm.sw.library.create` | PUT `/v3/alipay/iservice/ccm/sw/library/create` | ✅ |  |
| `alipay.iservice.ccm.sw.library.delete` | DELETE `/v3/alipay/iservice/ccm/sw/library/delete` | ✅ |  |
| `alipay.iservice.ccm.sw.library.modify` | PATCH `/v3/alipay/iservice/ccm/sw/library/modify` | ✅ |  |

### services/iservice/ccm/sw/order（1 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.iservice.ccm.sw.order.sync` | POST `/v3/alipay/iservice/ccm/sw/order/sync` | ✅ |  |

### services/iservice/ccm/sw/spu（1 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.iservice.ccm.sw.spu.sync` | POST `/v3/alipay/iservice/ccm/sw/spu/sync` | ✅ |  |

### services/iservice/ccm/sw/tree（6 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.iservice.ccm.sw.tree.batchquery` | GET `/v3/alipay/iservice/ccm/sw/tree/batchquery` | ✅ |  |
| `alipay.iservice.ccm.sw.tree.create` | PUT `/v3/alipay/iservice/ccm/sw/tree/create` | ✅ |  |
| `alipay.iservice.ccm.sw.tree.delete` | DELETE `/v3/alipay/iservice/ccm/sw/tree/delete` | ✅ |  |
| `alipay.iservice.ccm.sw.tree.get` | GET `/v3/alipay/iservice/ccm/sw/tree/get` | ✅ |  |
| `alipay.iservice.ccm.sw.tree.modify` | PATCH `/v3/alipay/iservice/ccm/sw/tree/modify` | ✅ |  |
| `alipay.iservice.ccm.sw.tree.publish` | PUT `/v3/alipay/iservice/ccm/sw/tree/publish` | ✅ |  |

### services/iservice/ccm/sw/treecategory（4 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.iservice.ccm.sw.treecategory.batchquery` | GET `/v3/alipay/iservice/ccm/sw/treecategory/batchquery` | ✅ |  |
| `alipay.iservice.ccm.sw.treecategory.create` | PUT `/v3/alipay/iservice/ccm/sw/treecategory/create` | ✅ |  |
| `alipay.iservice.ccm.sw.treecategory.delete` | DELETE `/v3/alipay/iservice/ccm/sw/treecategory/delete` | ✅ |  |
| `alipay.iservice.ccm.sw.treecategory.modify` | PATCH `/v3/alipay/iservice/ccm/sw/treecategory/modify` | ✅ |  |

### services/marketing/activity/app（1 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.marketing.activity.app.batchquery` | GET `/v3/alipay/marketing/activity/{activity_id}/app/batchquery` | ✅ |  |

### services/marketing/activity/batchquery（1 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.marketing.activity.batchquery` | POST `/v3/alipay/marketing/activity/batchquery` | ✅ |  |

### services/marketing/activity/consult（1 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.marketing.activity.consult` | POST `/v3/alipay/marketing/activity/consult` | ✅ |  |

### services/marketing/activity/delivery（3 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.marketing.activity.delivery.create` | POST `/v3/alipay/marketing/delivery` | ✅ |  |
| `alipay.marketing.activity.delivery.query` | POST `/v3/alipay/marketing/delivery/{delivery_id}/query` | ✅ |  |
| `alipay.marketing.activity.delivery.stop` | PATCH `/v3/alipay/marketing/delivery/{delivery_id}/stop` | ✅ |  |

### services/marketing/activity/deliverychannel（1 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.marketing.activity.deliverychannel.query` | POST `/v3/alipay/marketing/activity/deliverychannel/query` | ✅ |  |

### services/marketing/activity/goods（1 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.marketing.activity.goods.batchquery` | GET `/v3/alipay/marketing/activity/{activity_id}/goods/batchquery` | ✅ |  |

### services/marketing/activity/merchant（1 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.marketing.activity.merchant.batchquery` | GET `/v3/alipay/marketing/activity/{activity_id}/merchant/batchquery` | ✅ |  |

### services/marketing/activity/ordervoucher（11 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.marketing.activity.ordervoucher.append` | PATCH `/v3/alipay/marketing/ordervoucher/activity/{activity_id}/append` | ✅ |  |
| `alipay.marketing.activity.ordervoucher.associate` | POST `/v3/alipay/marketing/activity/ordervoucher/associate` | ✅ |  |
| `alipay.marketing.activity.ordervoucher.codecount` | GET `/v3/alipay/marketing/ordervoucher/activity/{activity_id}/voucher/codecount` | ✅ |  |
| `alipay.marketing.activity.ordervoucher.codedeposit` | POST `/v3/alipay/marketing/ordervoucher/activity/{activity_id}/voucher/codedeposit` | ✅ |  |
| `alipay.marketing.activity.ordervoucher.create` | POST `/v3/alipay/marketing/ordervoucher/activity` | ✅ |  |
| `alipay.marketing.activity.ordervoucher.disassociate` | POST `/v3/alipay/marketing/activity/ordervoucher/disassociate` | ✅ |  |
| `alipay.marketing.activity.ordervoucher.modify` | PATCH `/v3/alipay/marketing/ordervoucher/activity/{activity_id}` | ✅ |  |
| `alipay.marketing.activity.ordervoucher.query` | GET `/v3/alipay/marketing/activity/ordervoucher/query` | ✅ |  |
| `alipay.marketing.activity.ordervoucher.refund` | POST `/v3/alipay/marketing/ordervoucher/activity/{activity_id}/voucher/{voucher_code}/refund` | ✅ |  |
| `alipay.marketing.activity.ordervoucher.stop` | PATCH `/v3/alipay/marketing/ordervoucher/activity/{activity_id}/stop` | ✅ |  |
| `alipay.marketing.activity.ordervoucher.use` | POST `/v3/alipay/marketing/ordervoucher/activity/{activity_id}/voucher/{voucher_code}/use` | ✅ |  |

### services/marketing/activity/query（1 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.marketing.activity.query` | GET `/v3/alipay/marketing/activity/{activity_id}` | ✅ |  |

### services/marketing/activity/shop（1 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.marketing.activity.shop.batchquery` | GET `/v3/alipay/marketing/activity/{activity_id}/shop/batchquery` | ✅ |  |

### services/marketing/activity/user（2 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.marketing.activity.user.batchqueryvoucher` | GET `/v3/alipay/marketing/activity/ordervoucher/user/batchqueryvoucher` | ✅ |  |
| `alipay.marketing.activity.user.queryvoucher` | GET `/v3/alipay/marketing/activity/ordervoucher/user/voucher` | ✅ |  |

### services/marketing/activity/voucher（7 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.marketing.activity.voucher.append` | POST `/v3/alipay/marketing/activity/voucher/append` | ✅ |  |
| `alipay.marketing.activity.voucher.create` | POST `/v3/alipay/marketing/activity/voucher/create` | ✅ |  |
| `alipay.marketing.activity.voucher.modify` | POST `/v3/alipay/marketing/activity/voucher/modify` | ✅ |  |
| `alipay.marketing.activity.voucher.publish` | POST `/v3/alipay/marketing/activity/voucher/publish` | ✅ |  |
| `alipay.marketing.activity.voucher.query` | GET `/v3/alipay/marketing/activity/voucher/query` | ✅ |  |
| `alipay.marketing.activity.voucher.send` | POST `/v3/alipay/marketing/activity/voucher/send` | ✅ |  |
| `alipay.marketing.activity.voucher.stop` | POST `/v3/alipay/marketing/activity/voucher/stop` | ✅ |  |

### services/marketing/activity/voucherpackage（3 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.marketing.activity.voucherpackage.batchquery` | POST `/v3/alipay/marketing/activity/voucherpackage/batchquery` | ✅ |  |
| `alipay.marketing.activity.voucherpackage.consult` | POST `/v3/alipay/marketing/activity/voucherpackage/consult` | ✅ |  |
| `alipay.marketing.activity.voucherpackage.query` | GET `/v3/alipay/marketing/activity/voucherpackage/query` | ✅ |  |

### services/marketing/campaign（6 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.marketing.campaign.cash.create` | POST `/v3/alipay/marketing/campaign/cash/create` | ✅ |  |
| `alipay.marketing.campaign.cash.detail.query` | GET `/v3/alipay/marketing/campaign/cash/detail/query` | ✅ |  |
| `alipay.marketing.campaign.cash.list.query` | GET `/v3/alipay/marketing/campaign/cash/list/query` | ✅ |  |
| `alipay.marketing.campaign.cash.status.modify` | POST `/v3/alipay/marketing/campaign/cash/status/modify` | ✅ |  |
| `alipay.marketing.campaign.cash.trigger` | POST `/v3/alipay/marketing/campaign/cash/trigger` | ✅ |  |
| `alipay.marketing.campaign.order.voucher.consult` | POST `/v3/alipay/marketing/campaign/order/voucher/consult` | ✅ |  |

### services/marketing/card（9 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.marketing.card.activateform.query` | GET `/v3/alipay/marketing/card/activateform/query` | ✅ |  |
| `alipay.marketing.card.activateurl.apply` | POST `/v3/alipay/marketing/card/activateurl/apply` | ✅ |  |
| `alipay.marketing.card.consume.sync` | POST `/v3/alipay/marketing/card/consume/sync` | ✅ |  |
| `alipay.marketing.card.delete` | DELETE `/v3/alipay/marketing/card/delete` | ✅ |  |
| `alipay.marketing.card.formtemplate.set` | POST `/v3/alipay/marketing/card/formtemplate/set` | ✅ |  |
| `alipay.marketing.card.message.notify` | POST `/v3/alipay/marketing/card/message/notify` | ✅ |  |
| `alipay.marketing.card.open` | POST `/v3/alipay/marketing/card/open` | ✅ |  |
| `alipay.marketing.card.query` | POST `/v3/alipay/marketing/card/query` | ✅ |  |
| `alipay.marketing.card.update` | POST `/v3/alipay/marketing/card/update` | ✅ |  |

### services/marketing/card/benefit（4 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.marketing.card.benefit.create` | POST `/v3/alipay/marketing/card/benefit/create` | ✅ |  |
| `alipay.marketing.card.benefit.delete` | DELETE `/v3/alipay/marketing/card/benefit/delete` | ✅ |  |
| `alipay.marketing.card.benefit.modify` | POST `/v3/alipay/marketing/card/benefit/modify` | ✅ |  |
| `alipay.marketing.card.benefit.query` | GET `/v3/alipay/marketing/card/benefit/query` | ✅ |  |

### services/marketing/card/template（4 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.marketing.card.template.batchquery` | POST `/v3/alipay/marketing/card/template/batchquery` | ✅ |  |
| `alipay.marketing.card.template.create` | POST `/v3/alipay/marketing/card/template/create` | ✅ |  |
| `alipay.marketing.card.template.modify` | POST `/v3/alipay/marketing/card/template/modify` | ✅ |  |
| `alipay.marketing.card.template.query` | GET `/v3/alipay/marketing/card/template/query` | ✅ |  |

### services/marketing/material（2 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.marketing.material.create` | POST `/v3/alipay/marketing/material/create` | ✅ |  |
| `alipay.marketing.material.image.upload` | POST `/v3/alipay/marketing/material/image` | ✅ |  |

### services/marketing/recruit（5 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.marketing.recruit.enroll.close` | POST `/v3/alipay/marketing/recruit/enroll/close` | ✅ |  |
| `alipay.marketing.recruit.enroll.create` | POST `/v3/alipay/marketing/recruit/enroll/create` | ✅ |  |
| `alipay.marketing.recruit.enroll.query` | GET `/v3/alipay/marketing/recruit/enroll/query` | ✅ |  |
| `alipay.marketing.recruit.plan.query` | POST `/v3/alipay/marketing/recruit/plan/query` | ✅ |  |
| `alipay.marketing.recruit.planlist.query` | POST `/v3/alipay/marketing/recruit/planlist/query` | ✅ |  |

### services/merchant/anchor（1 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.merchant.anchor.info.query` | GET `/v3/alipay/merchant/anchor/query` | ✅ |  |

### services/merchant/image（1 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.merchant.image.upload` | POST `/v3/alipay/merchant/image/upload` | ✅ |  |

### services/merchant/indirect（4 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.merchant.indirect.authorder.close` | PUT `/v3/alipay/merchant/indirect/authorder/close` | ✅ |  |
| `alipay.merchant.indirect.authorder.create` | POST `/v3/alipay/merchant/indirect/authorder/create` | ✅ |  |
| `alipay.merchant.indirect.authorder.querystatus` | POST `/v3/alipay/merchant/indirect/authorder/querystatus` | ✅ |  |
| `alipay.merchant.indirect.smidbind.query` | GET `/v3/alipay/merchant/indirect/smidbind/query` | ✅ |  |

### services/merchant/iot（3 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.merchant.iot.device.bind` | POST `/v3/alipay/merchant/iot/device/bind` | ✅ |  |
| `alipay.merchant.iot.device.query` | GET `/v3/alipay/merchant/iot/device/query` | ✅ |  |
| `alipay.merchant.iot.device.verify` | POST `/v3/alipay/merchant/iot/device/verify` | ✅ |  |

### services/merchant/item（1 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.merchant.item.file.upload` | POST `/v3/alipay/merchant/item/file/upload` | ✅ |  |

### services/merchant/live（2 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.merchant.live.channel.query` | GET `/v3/alipay/merchant/live/channel/query` | ✅ |  |
| `alipay.merchant.live.itemplan.modify` | POST `/v3/alipay/merchant/live/itemplan/modify` | ✅ |  |

### services/merchant/order（1 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.merchant.order.sync` | POST `/v3/alipay/merchant/order/sync` | ✅ |  |

### services/merchant/tradecomplain（5 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.merchant.tradecomplain.batchquery` | POST `/v3/alipay/merchant/tradecomplain/batchquery` | ✅ |  |
| `alipay.merchant.tradecomplain.feedback.submit` | POST `/v3/alipay/merchant/tradecomplain/feedback` | ✅ |  |
| `alipay.merchant.tradecomplain.query` | GET `/v3/alipay/merchant/tradecomplain/{complain_event_id}` | ✅ |  |
| `alipay.merchant.tradecomplain.reply.submit` | POST `/v3/alipay/merchant/tradecomplain/reply` | ✅ |  |
| `alipay.merchant.tradecomplain.supplement.submit` | POST `/v3/alipay/merchant/tradecomplain/supplement` | ✅ |  |

### services/mobile/public/follow（1 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.mobile.public.follow.list` | GET `/v3/alipay/mobile/public/follow/list` | ✅ |  |

### services/mobile/public/info（1 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.mobile.public.info.query` | GET `/v3/alipay/mobile/public/info/query` | ✅ |  |

### services/mobile/public/menu（3 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.mobile.public.menu.add` | POST `/v3/alipay/mobile/public/menu/add` | ✅ |  |
| `alipay.mobile.public.menu.get` | GET `/v3/alipay/mobile/public/menu/get` | ✅ |  |
| `alipay.mobile.public.menu.update` | POST `/v3/alipay/mobile/public/menu/update` | ✅ |  |

### services/mobile/public/message（3 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.mobile.public.message.custom.send` | POST `/v3/alipay/mobile/public/message/custom/send` | ✅ |  |
| `alipay.mobile.public.message.single.send` | POST `/v3/alipay/mobile/public/message/single/send` | ✅ |  |
| `alipay.mobile.public.message.total.send` | POST `/v3/alipay/mobile/public/message/total/send` | ✅ |  |

### services/mobile/public/qrcode（1 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.mobile.public.qrcode.create` | POST `/v3/alipay/mobile/public/qrcode/create` | ✅ |  |

### services/mobile/public/shortlink（1 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.mobile.public.shortlink.create` | POST `/v3/alipay/mobile/public/shortlink/create` | ✅ |  |

### services/mobile/public/template（3 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.mobile.public.template.message.delete` | POST `/v3/alipay/mobile/public/template/message/delete` | ✅ |  |
| `alipay.mobile.public.template.message.get` | GET `/v3/alipay/mobile/public/template/message/get` | ✅ |  |
| `alipay.mobile.public.template.message.query` | POST `/v3/alipay/mobile/public/template/message/query` | ✅ |  |

### services/mobile/std/public/message（1 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.mobile.std.public.message.custom.send` | POST `/v3/alipay/mobile/std/public/message/custom/send` | ✅ |  |

### services/offline/market/applyorder（1 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.offline.market.applyorder.batchquery` | POST `/v3/alipay/offline/market/applyorder/batchquery` | ✅ |  |

### services/offline/market/shop（6 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.offline.market.shop.batchquery` | POST `/v3/alipay/offline/market/shop/batchquery` | ✅ |  |
| `alipay.offline.market.shop.category.query` | GET `/v3/alipay/offline/market/shop/category/query` | ✅ |  |
| `alipay.offline.market.shop.create` | POST `/v3/alipay/offline/market/shop/create` | ✅ |  |
| `alipay.offline.market.shop.modify` | POST `/v3/alipay/offline/market/shop/modify` | ✅ |  |
| `alipay.offline.market.shop.querydetail` | GET `/v3/alipay/offline/market/shop/querydetail` | ✅ |  |
| `alipay.offline.market.shop.summary.batchquery` | POST `/v3/alipay/offline/market/shop/summary/batchquery` | ✅ |  |

### services/offline/material/image（1 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.offline.material.image.upload` | POST `/v3/alipay/offline/material/image/upload` | ✅ |  |

### services/open/agent（5 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.open.agent.cancel` | POST `/v3/alipay/open/agent/cancel` | ✅ |  |
| `alipay.open.agent.confirm` | POST `/v3/alipay/open/agent/confirm` | ✅ |  |
| `alipay.open.agent.create` | POST `/v3/alipay/open/agent/create` | ✅ |  |
| `alipay.open.agent.mini.create` | POST `/v3/alipay/open/agent/mini/create` | ✅ |  |
| `alipay.open.agent.order.query` | GET `/v3/alipay/open/agent/order/query` | ✅ |  |

### services/open/agent/sign（6 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.open.agent.common.sign` | POST `/v3/alipay/open/agent/common/sign` | ✅ |  |
| `alipay.open.agent.commonsign.confirm` | POST `/v3/alipay/open/agent/commonsign/confirm` | ✅ |  |
| `alipay.open.agent.facetoface.sign` | POST `/v3/alipay/open/agent/facetoface/sign` | ✅ |  |
| `alipay.open.agent.mobilepay.sign` | POST `/v3/alipay/open/agent/mobilepay/sign` | ✅ |  |
| `alipay.open.agent.offlinepayment.sign` | POST `/v3/alipay/open/agent/offlinepayment/sign` | ✅ |  |
| `alipay.open.agent.signstatus.query` | POST `/v3/alipay/open/agent/signstatus/query` | ✅ |  |

### services/open/app/api（4 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.open.app.api.field.apply` | POST `/v3/alipay/open/app/api/field/apply` | ✅ |  |
| `alipay.open.app.api.field.query` | GET `/v3/alipay/open/app/api/field/query` | ✅ |  |
| `alipay.open.app.api.query` | GET `/v3/alipay/open/app/api/query` | ✅ |  |
| `alipay.open.app.api.scene.query` | GET `/v3/alipay/open/app/api/scene/query` | ✅ |  |

### services/open/app/members（3 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.open.app.members.create` | POST `/v3/alipay/open/app/members/create` | ✅ |  |
| `alipay.open.app.members.delete` | DELETE `/v3/alipay/open/app/members/delete` | ✅ |  |
| `alipay.open.app.members.query` | GET `/v3/alipay/open/app/members/query` | ✅ |  |

### services/open/app（3 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.open.app.messagetemplate.subscribe.query` | POST `/v3/alipay/open/app/messagetemplate/subscribe/query` | ✅ |  |
| `alipay.open.app.mini.templatemessage.send` | POST `/v3/alipay/open/app/mini/templatemessage/send` | ✅ |  |
| `alipay.open.app.qrcode.create` | POST `/v3/alipay/open/app/qrcode/create` | ✅ |  |

### services/open/app/service（5 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.open.app.service.apply` | POST `/v3/alipay/open/app/service/apply` | ✅ |  |
| `alipay.open.app.service.delete` | DELETE `/v3/alipay/open/app/service/delete` | ✅ |  |
| `alipay.open.app.service.list.query` | GET `/v3/alipay/open/app/service/list/query` | ✅ |  |
| `alipay.open.app.service.query` | GET `/v3/alipay/open/app/service/query` | ✅ |  |
| `alipay.open.app.service.schema.query` | GET `/v3/alipay/open/app/service/schema/query` | ✅ |  |

### services/open/auth（4 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.open.auth.app.aes.get` | GET `/v3/alipay/open/auth/app/aes/get` | ✅ |  |
| `alipay.open.auth.app.aes.set` | POST `/v3/alipay/open/auth/app/aes/set` | ✅ |  |
| `alipay.open.auth.token.app.query` | GET `/v3/alipay/open/auth/token/app/query` | ✅ |  |
| `alipay.open.auth.userauth.relationship.query` | GET `/v3/alipay/open/auth/userauth/relationship/query` | ✅ |  |

### services/open（1 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.open.auth.token.app` | POST `/v3/alipay/open/auth/token/app` | 子集 | 响应 expires_in 已作废；官方说明应用令牌长期有效。 |

### services/open/biz（1 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.open.biz.create` | POST `/v3/alipay/open/biz/create` | ✅ |  |

### services/open/fee（1 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.open.fee.adjust.apply` | POST `/v3/alipay/open/fee/adjust/apply` | ✅ |  |

### services/open/file（1 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.open.file.upload` | POST `/v3/alipay/open/file/upload` | ✅ |  |

### services/open/instantdelivery/account（2 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.open.instantdelivery.account.create` | POST `/v3/alipay/open/instantdelivery/account/create` | ✅ |  |
| `alipay.open.instantdelivery.account.query` | GET `/v3/alipay/open/instantdelivery/account/query` | ✅ |  |

### services/open/instantdelivery/merchantagreement（1 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.open.instantdelivery.merchantagreement.query` | GET `/v3/alipay/open/instantdelivery/merchantagreement/query` | ✅ |  |

### services/open/instantdelivery/merchantshop（4 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.open.instantdelivery.merchantshop.batchquery` | POST `/v3/alipay/open/instantdelivery/merchantshop/batchquery` | ✅ |  |
| `alipay.open.instantdelivery.merchantshop.create` | POST `/v3/alipay/open/instantdelivery/merchantshop/create` | ✅ |  |
| `alipay.open.instantdelivery.merchantshop.modify` | POST `/v3/alipay/open/instantdelivery/merchantshop/modify` | ✅ |  |
| `alipay.open.instantdelivery.merchantshop.query` | GET `/v3/alipay/open/instantdelivery/merchantshop/query` | ✅ |  |

### services/open/invite（1 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.open.invite.order.query` | GET `/v3/alipay/open/invite/order/query` | ✅ |  |

### services/open/mini/appdeploy（2 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.open.mini.appdeploy.byappversion.query` | GET `/v3/alipay/open/mini/appdeploy/byappversion/query` | ✅ |  |
| `alipay.open.mini.appdeploy.bydeployversion.query` | GET `/v3/alipay/open/mini/appdeploy/bydeployversion/query` | ✅ |  |

### services/open/mini/baseinfo（3 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.open.mini.baseinfo.modify` | POST `/v3/alipay/open/mini/baseinfo/modify` | ✅ |  |
| `alipay.open.mini.baseinfo.name.check` | GET `/v3/alipay/open/mini/baseinfo/name/check` | ✅ |  |
| `alipay.open.mini.baseinfo.query` | GET `/v3/alipay/open/mini/baseinfo/query` | ✅ |  |

### services/open/mini/category（2 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.open.mini.category.query` | GET `/v3/alipay/open/mini/category/query` | ✅ |  |
| `alipay.open.mini.category.require.query` | POST `/v3/alipay/open/mini/category/require/query` | ✅ |  |

### services/open/mini/experience（3 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.open.mini.experience.cancel` | POST `/v3/alipay/open/mini/experience/cancel` | ✅ |  |
| `alipay.open.mini.experience.create` | POST `/v3/alipay/open/mini/experience/create` | ✅ |  |
| `alipay.open.mini.experience.query` | GET `/v3/alipay/open/mini/experience/query` | ✅ |  |

### services/open/mini/individual（1 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.open.mini.individual.business.certify` | POST `/v3/alipay/open/mini/individual/business/certify` | ✅ |  |

### services/open/mini/inner（4 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.open.mini.inner.experience.cancel` | POST `/v3/alipay/open/mini/inner/experience/cancel` | ✅ |  |
| `alipay.open.mini.inner.members.add` | POST `/v3/alipay/open/mini/inner/members/add` | ✅ |  |
| `alipay.open.mini.inner.members.delete` | DELETE `/v3/alipay/open/mini/inner/members/delete` | ✅ |  |
| `alipay.open.mini.inner.members.query` | GET `/v3/alipay/open/mini/inner/members/query` | ✅ |  |

### services/open/mini/innerapp（10 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.open.mini.innerapp.create` | POST `/v3/alipay/open/mini/innerapp/create` | ✅ |  |
| `alipay.open.mini.innerapp.plugin.order` | POST `/v3/alipay/open/mini/innerapp/plugin/order` | ✅ |  |
| `alipay.open.mini.innerapp.pluginrelation.query` | POST `/v3/alipay/open/mini/innerapp/pluginrelation/query` | ✅ |  |
| `alipay.open.mini.innerapp.pluginservice.publish` | POST `/v3/alipay/open/mini/innerapp/pluginservice/publish` | ✅ |  |
| `alipay.open.mini.innerapp.pluginsyncmode.modify` | POST `/v3/alipay/open/mini/innerapp/pluginsyncmode/modify` | ✅ |  |
| `alipay.open.mini.innerapp.pluginuseconfig.cancel` | POST `/v3/alipay/open/mini/innerapp/pluginuseconfig/cancel` | ✅ |  |
| `alipay.open.mini.innerapp.pluginuseconfig.online` | POST `/v3/alipay/open/mini/innerapp/pluginuseconfig/online` | ✅ |  |
| `alipay.open.mini.innerapp.pluginuseconfig.upgrade` | POST `/v3/alipay/open/mini/innerapp/pluginuseconfig/upgrade` | ✅ |  |
| `alipay.open.mini.innerapp.service.publish` | POST `/v3/alipay/open/mini/innerapp/service/publish` | ✅ |  |
| `alipay.open.mini.innerapp.service.query` | GET `/v3/alipay/open/mini/innerapp/service/query` | ✅ |  |

### services/open/mini/innerbaseinfo（2 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.open.mini.innerbaseinfo.applogo.upload` | POST `/v3/alipay/open/mini/innerbaseinfo/applogo/upload` | ✅ |  |
| `alipay.open.mini.innerbaseinfo.query` | GET `/v3/alipay/open/mini/innerbaseinfo/query` | ✅ |  |

### services/open/mini/innerversion/audit（2 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.open.mini.innerversion.audit.cancel` | POST `/v3/alipay/open/mini/innerversion/audit/cancel` | ✅ |  |
| `alipay.open.mini.innerversion.audit.submit` | POST `/v3/alipay/open/mini/innerversion/audit/submit` | ✅ |  |

### services/open/mini/innerversion（13 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.open.mini.innerversion.backdev.publish` | POST `/v3/alipay/open/mini/innerversion/backdev/publish` | ✅ |  |
| `alipay.open.mini.innerversion.condition.batchquery` | POST `/v3/alipay/open/mini/innerversion/condition/batchquery` | ✅ |  |
| `alipay.open.mini.innerversion.content.rollback` | POST `/v3/alipay/open/mini/innerversion/content/rollback` | ✅ |  |
| `alipay.open.mini.innerversion.info.delete` | DELETE `/v3/alipay/open/mini/innerversion/info/delete` | ✅ |  |
| `alipay.open.mini.innerversion.info.query` | GET `/v3/alipay/open/mini/innerversion/info/query` | ✅ |  |
| `alipay.open.mini.innerversion.instantiation.query` | GET `/v3/alipay/open/mini/innerversion/instantiation/query` | ✅ |  |
| `alipay.open.mini.innerversion.last.query` | GET `/v3/alipay/open/mini/innerversion/last/query` | ✅ |  |
| `alipay.open.mini.innerversion.modelforoffline.query` | GET `/v3/alipay/open/mini/innerversion/modelforoffline/query` | ✅ |  |
| `alipay.open.mini.innerversion.offline.publish` | POST `/v3/alipay/open/mini/innerversion/offline/publish` | ✅ |  |
| `alipay.open.mini.innerversion.online` | POST `/v3/alipay/open/mini/innerversion/online` | ✅ |  |
| `alipay.open.mini.innerversion.online.publish` | POST `/v3/alipay/open/mini/innerversion/online/publish` | ✅ |  |
| `alipay.open.mini.innerversion.package.create` | POST `/v3/alipay/open/mini/innerversion/package/create` | ✅ |  |
| `alipay.open.mini.innerversion.sync` | POST `/v3/alipay/open/mini/innerversion/sync` | ✅ |  |

### services/open/mini/innerversion/beta（2 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.open.mini.innerversion.beta.confirm` | POST `/v3/alipay/open/mini/innerversion/beta/confirm` | ✅ |  |
| `alipay.open.mini.innerversion.betainfo.query` | GET `/v3/alipay/open/mini/innerversion/betainfo/query` | ✅ |  |

### services/open/mini/innerversion/gray（2 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.open.mini.innerversion.gray.finish` | POST `/v3/alipay/open/mini/innerversion/gray/finish` | ✅ |  |
| `alipay.open.mini.innerversion.gray.publish` | POST `/v3/alipay/open/mini/innerversion/gray/publish` | ✅ |  |

### services/open/mini/innerversion/upload（4 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.open.mini.innerversion.nobuild.upload` | POST `/v3/alipay/open/mini/innerversion/nobuild/upload` | ✅ |  |
| `alipay.open.mini.innerversion.templated.upload` | POST `/v3/alipay/open/mini/innerversion/templated/upload` | ✅ |  |
| `alipay.open.mini.innerversion.upload` | POST `/v3/alipay/open/mini/innerversion/upload` | ✅ |  |
| `alipay.open.mini.innerversion.uploadstatus.query` | GET `/v3/alipay/open/mini/innerversion/uploadstatus/query` | ✅ |  |

### services/open/mini/isv（4 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.open.mini.isv.create` | POST `/v3/alipay/open/mini/isv/create` | ✅ |  |
| `alipay.open.mini.isv.fastregister.create` | POST `/v3/alipay/open/mini/isv/fastregister/create` | ✅ |  |
| `alipay.open.mini.isv.fastregister.query` | GET `/v3/alipay/open/mini/isv/fastregister/query` | ✅ |  |
| `alipay.open.mini.isv.query` | GET `/v3/alipay/open/mini/isv/query` | ✅ |  |

### services/open/mini/miniapp（5 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.open.mini.miniapp.brand.cancel` | POST `/v3/alipay/open/mini/miniapp/brand/cancel` | ✅ |  |
| `alipay.open.mini.miniapp.brand.query` | GET `/v3/alipay/open/mini/miniapp/brand/query` | ✅ |  |
| `alipay.open.mini.miniapp.brand.submit` | POST `/v3/alipay/open/mini/miniapp/brand/submit` | ✅ |  |
| `alipay.open.mini.miniapp.brand.upload` | POST `/v3/alipay/open/mini/miniapp/brand/upload` | ✅ |  |
| `alipay.open.mini.miniapp.serviceconfig.modify` | POST `/v3/alipay/open/mini/miniapp/serviceconfig/modify` | ✅ |  |

### services/open/mini/qrcode（3 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.open.mini.qrcode.bind` | POST `/v3/alipay/open/mini/qrcode/bind` | ✅ |  |
| `alipay.open.mini.qrcode.query` | GET `/v3/alipay/open/mini/qrcode/query` | ✅ |  |
| `alipay.open.mini.qrcode.unbind` | POST `/v3/alipay/open/mini/qrcode/unbind` | ✅ |  |

### services/open/mini/resource（2 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.open.mini.resource.promotionsource.notify` | POST `/v3/alipay/open/mini/resource/promotionsource/notify` | ✅ |  |
| `alipay.open.mini.resource.record.notify` | POST `/v3/alipay/open/mini/resource/record/notify` | ✅ |  |

### services/open/mini/safedomain（2 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.open.mini.safedomain.create` | POST `/v3/alipay/open/mini/safedomain/create` | ✅ |  |
| `alipay.open.mini.safedomain.delete` | DELETE `/v3/alipay/open/mini/safedomain/delete` | ✅ |  |

### services/open/mini/template（2 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.open.mini.template.marketing.create` | POST `/v3/alipay/open/mini/template/marketing/create` | ✅ |  |
| `alipay.open.mini.template.usage.query` | GET `/v3/alipay/open/mini/template/usage/query` | ✅ |  |

### services/open/mini/templatemsg（2 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.open.mini.templatemsg.maketing.batchquery` | POST `/v3/alipay/open/mini/templatemsg/maketing/batchquery` | ✅ |  |
| `alipay.open.mini.templatemsg.tinypayswitch.confirm` | POST `/v3/alipay/open/mini/templatemsg/tinypayswitch/confirm` | ✅ |  |

### services/open/mini/tips（5 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.open.mini.tips.delivery.batchquery` | POST `/v3/alipay/open/mini/tips/delivery/batchquery` | ✅ |  |
| `alipay.open.mini.tips.delivery.create` | POST `/v3/alipay/open/mini/tips/delivery/create` | ✅ |  |
| `alipay.open.mini.tips.delivery.modify` | POST `/v3/alipay/open/mini/tips/delivery/modify` | ✅ |  |
| `alipay.open.mini.tips.delivery.query` | GET `/v3/alipay/open/mini/tips/delivery/query` | ✅ |  |
| `alipay.open.mini.tips.statistic.query` | GET `/v3/alipay/open/mini/tips/statistic/query` | ✅ |  |

### services/open/mini/version/audit（2 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.open.mini.version.audit.apply` | POST `/v3/alipay/open/mini/version/audit/apply` | ✅ |  |
| `alipay.open.mini.version.audit.cancel` | POST `/v3/alipay/open/mini/version/audit/cancel` | ✅ |  |

### services/open/mini/version（9 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.open.mini.version.audited.cancel` | POST `/v3/alipay/open/mini/version/audited/cancel` | ✅ |  |
| `alipay.open.mini.version.build.query` | GET `/v3/alipay/open/mini/version/build/query` | ✅ |  |
| `alipay.open.mini.version.delete` | POST `/v3/alipay/open/mini/version/delete` | ✅ |  |
| `alipay.open.mini.version.detail.query` | GET `/v3/alipay/open/mini/version/detail/query` | ✅ |  |
| `alipay.open.mini.version.list.query` | GET `/v3/alipay/open/mini/version/list/query` | ✅ |  |
| `alipay.open.mini.version.offline` | POST `/v3/alipay/open/mini/version/offline` | ✅ |  |
| `alipay.open.mini.version.online` | POST `/v3/alipay/open/mini/version/online` | ✅ |  |
| `alipay.open.mini.version.rollback` | POST `/v3/alipay/open/mini/version/rollback` | ✅ |  |
| `alipay.open.mini.version.upload` | POST `/v3/alipay/open/mini/version/upload` | ✅ |  |

### services/open/mini/version/gray（2 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.open.mini.version.gray.cancel` | POST `/v3/alipay/open/mini/version/gray/cancel` | ✅ |  |
| `alipay.open.mini.version.gray.online` | POST `/v3/alipay/open/mini/version/gray/online` | ✅ |  |

### services/open/public/account（3 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.open.public.account.create` | POST `/v3/alipay/open/public/account/create` | ✅ |  |
| `alipay.open.public.account.delete` | DELETE `/v3/alipay/open/public/account/delete` | ✅ |  |
| `alipay.open.public.account.query` | GET `/v3/alipay/open/public/account/query` | ✅ |  |

### services/open/public/advert（4 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.open.public.advert.batchquery` | POST `/v3/alipay/open/public/advert/batchquery` | ✅ |  |
| `alipay.open.public.advert.create` | POST `/v3/alipay/open/public/advert/create` | ✅ |  |
| `alipay.open.public.advert.delete` | DELETE `/v3/alipay/open/public/advert/delete` | ✅ |  |
| `alipay.open.public.advert.modify` | POST `/v3/alipay/open/public/advert/modify` | ✅ |  |

### services/open/public/articlesummary（1 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.open.public.articlesummary.data.batchquery` | POST `/v3/alipay/open/public/articlesummary/data/batchquery` | ✅ |  |

### services/open/public/contact（1 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.open.public.contact.follow.batchquery` | POST `/v3/alipay/open/public/contact/follow/batchquery` | ✅ |  |

### services/open/public/default（1 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.open.public.default.extension.create` | POST `/v3/alipay/open/public/default/extension/create` | ✅ |  |

### services/open/public/follow（1 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.open.public.follow.batchquery` | POST `/v3/alipay/open/public/follow/batchquery` | ✅ |  |

### services/open/public/group（5 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.open.public.group.batchquery` | POST `/v3/alipay/open/public/group/batchquery` | ✅ |  |
| `alipay.open.public.group.create` | POST `/v3/alipay/open/public/group/create` | ✅ |  |
| `alipay.open.public.group.crowd.query` | POST `/v3/alipay/open/public/group/crowd/query` | ✅ |  |
| `alipay.open.public.group.delete` | DELETE `/v3/alipay/open/public/group/delete` | ✅ |  |
| `alipay.open.public.group.modify` | POST `/v3/alipay/open/public/group/modify` | ✅ |  |

### services/open/public/info（2 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.open.public.info.modify` | POST `/v3/alipay/open/public/info/modify` | ✅ |  |
| `alipay.open.public.info.query` | GET `/v3/alipay/open/public/info/query` | ✅ |  |

### services/open/public/label（3 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.open.public.label.create` | POST `/v3/alipay/open/public/label/create` | ✅ |  |
| `alipay.open.public.label.query` | GET `/v3/alipay/open/public/label/query` | ✅ |  |
| `alipay.open.public.label.user.query` | GET `/v3/alipay/open/public/label/user/query` | ✅ |  |

### services/open/public/life（9 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.open.public.life.aboard.apply` | POST `/v3/alipay/open/public/life/aboard/apply` | ✅ |  |
| `alipay.open.public.life.agentcreate.query` | GET `/v3/alipay/open/public/life/agentcreate/query` | ✅ |  |
| `alipay.open.public.life.debark.apply` | POST `/v3/alipay/open/public/life/debark/apply` | ✅ |  |
| `alipay.open.public.life.label.batchquery` | POST `/v3/alipay/open/public/life/label/batchquery` | ✅ |  |
| `alipay.open.public.life.label.create` | POST `/v3/alipay/open/public/life/label/create` | ✅ |  |
| `alipay.open.public.life.label.delete` | DELETE `/v3/alipay/open/public/life/label/delete` | ✅ |  |
| `alipay.open.public.life.label.modify` | POST `/v3/alipay/open/public/life/label/modify` | ✅ |  |
| `alipay.open.public.life.msg.recall` | POST `/v3/alipay/open/public/life/msg/recall` | ✅ |  |
| `alipay.open.public.life.msg.send` | POST `/v3/alipay/open/public/life/msg/send` | ✅ |  |

### services/open/public/matchuser（2 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.open.public.matchuser.label.create` | POST `/v3/alipay/open/public/matchuser/label/create` | ✅ |  |
| `alipay.open.public.matchuser.label.delete` | POST `/v3/alipay/open/public/matchuser/label/delete` | ✅ |  |

### services/open/public/menu（5 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.open.public.menu.batchquery` | POST `/v3/alipay/open/public/menu/batchquery` | ✅ |  |
| `alipay.open.public.menu.create` | POST `/v3/alipay/open/public/menu/create` | ✅ |  |
| `alipay.open.public.menu.data.batchquery` | POST `/v3/alipay/open/public/menu/data/batchquery` | ✅ |  |
| `alipay.open.public.menu.modify` | POST `/v3/alipay/open/public/menu/modify` | ✅ |  |
| `alipay.open.public.menu.query` | GET `/v3/alipay/open/public/menu/query` | ✅ |  |

### services/open/public/message（7 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.open.public.message.content.create` | POST `/v3/alipay/open/public/message/content/create` | ✅ |  |
| `alipay.open.public.message.content.modify` | POST `/v3/alipay/open/public/message/content/modify` | ✅ |  |
| `alipay.open.public.message.custom.send` | POST `/v3/alipay/open/public/message/custom/send` | ✅ |  |
| `alipay.open.public.message.group.send` | POST `/v3/alipay/open/public/message/group/send` | ✅ |  |
| `alipay.open.public.message.label.send` | POST `/v3/alipay/open/public/message/label/send` | ✅ |  |
| `alipay.open.public.message.single.send` | POST `/v3/alipay/open/public/message/single/send` | ✅ |  |
| `alipay.open.public.message.total.send` | POST `/v3/alipay/open/public/message/total/send` | ✅ |  |

### services/open/public/payee（2 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.open.public.payee.bind.create` | POST `/v3/alipay/open/public/payee/bind/create` | ✅ |  |
| `alipay.open.public.payee.bind.delete` | DELETE `/v3/alipay/open/public/payee/bind/delete` | ✅ |  |

### services/open/public/personalized（5 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.open.public.personalized.extension.batchquery` | POST `/v3/alipay/open/public/personalized/extension/batchquery` | ✅ |  |
| `alipay.open.public.personalized.extension.create` | POST `/v3/alipay/open/public/personalized/extension/create` | ✅ |  |
| `alipay.open.public.personalized.extension.delete` | DELETE `/v3/alipay/open/public/personalized/extension/delete` | ✅ |  |
| `alipay.open.public.personalized.menu.create` | POST `/v3/alipay/open/public/personalized/menu/create` | ✅ |  |
| `alipay.open.public.personalized.menu.delete` | DELETE `/v3/alipay/open/public/personalized/menu/delete` | ✅ |  |

### services/open/public/qrcode（1 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.open.public.qrcode.create` | POST `/v3/alipay/open/public/qrcode/create` | ✅ |  |

### services/open/public/setting（1 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.open.public.setting.category.query` | GET `/v3/alipay/open/public/setting/category/query` | ✅ |  |

### services/open/public/shortlink（1 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.open.public.shortlink.create` | POST `/v3/alipay/open/public/shortlink/create` | ✅ |  |

### services/open/public/singlearticle（1 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.open.public.singlearticle.data.batchquery` | POST `/v3/alipay/open/public/singlearticle/data/batchquery` | ✅ |  |

### services/open/public/template（2 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.open.public.template.message.get` | GET `/v3/alipay/open/public/template/message/get` | ✅ |  |
| `alipay.open.public.template.message.industry.modify` | POST `/v3/alipay/open/public/template/message/industry/modify` | ✅ |  |

### services/open/public/topic（4 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.open.public.topic.batchquery` | POST `/v3/alipay/open/public/topic/batchquery` | ✅ |  |
| `alipay.open.public.topic.create` | POST `/v3/alipay/open/public/topic/create` | ✅ |  |
| `alipay.open.public.topic.delete` | DELETE `/v3/alipay/open/public/topic/delete` | ✅ |  |
| `alipay.open.public.topic.modify` | POST `/v3/alipay/open/public/topic/modify` | ✅ |  |

### services/open/public/user（2 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.open.public.user.data.batchquery` | POST `/v3/alipay/open/public/user/data/batchquery` | ✅ |  |
| `alipay.open.public.user.follow.query` | GET `/v3/alipay/open/public/user/follow/query` | ✅ |  |

### services/open/search（1 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.open.search.ability.batchquery` | POST `/v3/alipay/open/search/ability/batchquery` | ✅ |  |

### services/open/search/appkeyword（5 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.open.search.appkeyword.apply` | POST `/v3/alipay/open/search/appkeyword/apply` | ✅ |  |
| `alipay.open.search.appkeyword.batchquery` | GET `/v3/alipay/open/search/appkeyword/batchquery` | ✅ |  |
| `alipay.open.search.appkeyword.delete` | POST `/v3/alipay/open/search/appkeyword/delete` | ✅ |  |
| `alipay.open.search.appkeyword.querystatus` | GET `/v3/alipay/open/search/appkeyword/querystatus` | ✅ |  |
| `alipay.open.search.appkeywordquota.query` | GET `/v3/alipay/open/search/appkeywordquota/query` | ✅ |  |

### services/open/search/order（3 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.open.search.baseorder.modify` | POST `/v3/alipay/open/search/baseorder/modify` | ✅ |  |
| `alipay.open.search.orderdetail.query` | POST `/v3/alipay/open/search/orderdetail/query` | ✅ |  |
| `alipay.open.search.serviceorder.batchquery` | POST `/v3/alipay/open/search/serviceorder/batchquery` | ✅ |  |

### services/open/search/box（9 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.open.search.box.apply` | POST `/v3/alipay/open/search/box/apply` | ✅ |  |
| `alipay.open.search.box.batchquery` | POST `/v3/alipay/open/search/box/batchquery` | ✅ |  |
| `alipay.open.search.box.consult` | POST `/v3/alipay/open/search/box/consult` | ✅ |  |
| `alipay.open.search.box.modify` | POST `/v3/alipay/open/search/box/modify` | ✅ |  |
| `alipay.open.search.box.offline` | POST `/v3/alipay/open/search/box/offline` | ✅ |  |
| `alipay.open.search.box.online` | POST `/v3/alipay/open/search/box/online` | ✅ |  |
| `alipay.open.search.box.query` | GET `/v3/alipay/open/search/box/query` | ✅ |  |
| `alipay.open.search.boxactivity.apply` | POST `/v3/alipay/open/search/boxactivity/apply` | ✅ |  |
| `alipay.open.search.boxactivity.query` | GET `/v3/alipay/open/search/boxactivity/query` | ✅ |  |

### services/open/search/subservicekeyword（4 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.open.search.subservicekeyword.apply` | POST `/v3/alipay/open/search/subservicekeyword/apply` | ✅ |  |
| `alipay.open.search.subservicekeyword.batchquery` | GET `/v3/alipay/open/search/subservicekeyword/batchquery` | ✅ |  |
| `alipay.open.search.subservicekeyword.delete` | POST `/v3/alipay/open/search/subservicekeyword/delete` | ✅ |  |
| `alipay.open.search.subservicekeyword.querystatus` | GET `/v3/alipay/open/search/subservicekeyword/querystatus` | ✅ |  |

### services/open/servicemarket（2 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.open.servicemarket.commodity.shop.offline` | POST `/v3/alipay/open/servicemarket/commodity/shop/offline` | ✅ |  |
| `alipay.open.servicemarket.promotiontask.relation.query` | GET `/v3/alipay/open/servicemarket/promotiontask/relation/query` | ✅ |  |

### services/open/servicemarket/order（8 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.open.servicemarket.order.accept` | POST `/v3/alipay/open/servicemarket/order/accept` | ✅ |  |
| `alipay.open.servicemarket.order.create` | POST `/v3/alipay/open/servicemarket/order/create` | ✅ |  |
| `alipay.open.servicemarket.order.item.cancel` | POST `/v3/alipay/open/servicemarket/order/item/cancel` | ✅ |  |
| `alipay.open.servicemarket.order.item.complete` | POST `/v3/alipay/open/servicemarket/order/item/complete` | ✅ |  |
| `alipay.open.servicemarket.order.item.confirm` | POST `/v3/alipay/open/servicemarket/order/item/confirm` | ✅ |  |
| `alipay.open.servicemarket.order.notify` | POST `/v3/alipay/open/servicemarket/order/notify` | ✅ |  |
| `alipay.open.servicemarket.order.query` | GET `/v3/alipay/open/servicemarket/order/query` | ✅ |  |
| `alipay.open.servicemarket.order.reject` | POST `/v3/alipay/open/servicemarket/order/reject` | ✅ |  |

### services/open/sp/blueseaactivity（3 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.open.sp.blueseaactivity.create` | POST `/v3/alipay/open/sp/blueseaactivity/create` | ✅ |  |
| `alipay.open.sp.blueseaactivity.modify` | POST `/v3/alipay/open/sp/blueseaactivity/modify` | ✅ |  |
| `alipay.open.sp.blueseaactivity.query` | GET `/v3/alipay/open/sp/blueseaactivity/query` | ✅ |  |

### services/open/sp（3 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.open.sp.image.upload` | POST `/v3/alipay/open/sp/image/upload` | ✅ |  |
| `alipay.open.sp.leads.expand.create` | POST `/v3/alipay/open/sp/leads/expand/create` | ✅ |  |
| `alipay.open.sp.merchant.inconsistent.approve` | POST `/v3/alipay/open/sp/merchant/inconsistent/approve` | ✅ |  |

### services/open/sp/isv（3 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.open.sp.isv.bill.sync` | POST `/v3/alipay/open/sp/isv/bill/sync` | ✅ |  |
| `alipay.open.sp.isv.signauth.create` | POST `/v3/alipay/open/sp/isv/signauth/create` | ✅ |  |
| `alipay.open.sp.isv.signauth.query` | GET `/v3/alipay/open/sp/isv/signauth/query` | ✅ |  |

### services/open/sp/operation（3 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.open.sp.operation.apply` | POST `/v3/alipay/open/sp/operation/apply` | ✅ |  |
| `alipay.open.sp.operation.qrcode.query` | GET `/v3/alipay/open/sp/operation/qrcode/query` | ✅ |  |
| `alipay.open.sp.operation.result.query` | GET `/v3/alipay/open/sp/operation/result/query` | ✅ |  |

### services/open/sp/oppor（2 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.open.sp.oppor.detail.query` | GET `/v3/alipay/open/sp/oppor/detail/query` | ✅ |  |
| `alipay.open.sp.oppor.page.query` | GET `/v3/alipay/open/sp/oppor/page/query` | ✅ |  |

### services/open/violation（2 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.open.violation.violationdetail.query` | GET `/v3/alipay/open/violation/violationdetail/query` | ✅ |  |
| `alipay.open.violation.violationevent.batchquery` | POST `/v3/alipay/open/violation/violationevent/batchquery` | ✅ |  |

### services/pcredit（6 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.pcredit.huabei.auth.accumulation.query` | GET `/v3/alipay/pcredit/huabei/auth/accumulation/query` | ✅ |  |
| `alipay.pcredit.huabei.auth.agreement.close` | PUT `/v3/alipay/pcredit/huabei/auth/agreement/close` | ✅ |  |
| `alipay.pcredit.huabei.auth.agreement.query` | GET `/v3/alipay/pcredit/huabei/auth/agreement/query` | ✅ |  |
| `alipay.pcredit.huabei.auth.order.query` | GET `/v3/alipay/pcredit/huabei/auth/order/query` | ✅ |  |
| `alipay.pcredit.huabei.auth.order.unfreeze` | PUT `/v3/alipay/pcredit/huabei/auth/order/unfreeze` | ✅ |  |
| `alipay.pcredit.huabei.auth.settle.apply` | PUT `/v3/alipay/pcredit/huabei/auth/settle/apply` | ✅ |  |

### services/security（3 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.security.prod.sss.query` | POST `/v3/alipay/security/prod/sss/query/{bbb}` | ✅ |  |
| `alipay.security.risk.customerrisk.send` | POST `/v3/alipay/security/risk/customerrisk/send` | ✅ |  |
| `alipay.security.risk.verifyidentity.miniapp.confirm` | POST `/v3/alipay/security/risk/verifyidentity/miniapp/confirm` | ✅ |  |

### services/social（1 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.social.base.contentlib.standardcontent.batchquery` | POST `/v3/alipay/social/base/contentlib/standardcontent/batchquery` | ✅ |  |

### services/system（1 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.system.oauth.token` | POST `/v3/alipay/system/oauth/token` | 子集 | 响应 alipay_user_id 已废弃；按接口场景使用公开身份字段。 |

### services/trade/advance（1 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.trade.advance.consult` | GET `/v3/alipay/trade/advance/consult` | ✅ |  |

### services/trade/batch（4 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.trade.batch.settle` | POST `/v3/alipay/trade/batch/settle` | ✅ |  |
| `alipay.trade.batch.settle.query` | POST `/v3/alipay/trade/batch/settle/query` | ✅ |  |
| `alipay.trade.batch.transfer` | POST `/v3/alipay/trade/batch/transfer` | ✅ |  |
| `alipay.trade.batch.transfer.query` | GET `/v3/alipay/trade/batch/transfer/query` | ✅ |  |

### services/trade（6 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.trade.cancel` | POST `/v3/alipay/trade/cancel` | ✅ |  |
| `alipay.trade.close` | POST `/v3/alipay/trade/close` | ✅ |  |
| `alipay.trade.create` | POST `/v3/alipay/trade/create` | ✅ |  |
| `alipay.trade.pay` | POST `/v3/alipay/trade/pay` | ✅ |  |
| `alipay.trade.precreate` | POST `/v3/alipay/trade/precreate` | ✅ |  |
| `alipay.trade.query` | POST `/v3/alipay/trade/query` | ✅ |  |

### services/trade/customs（2 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.trade.customs.declare` | POST `/v3/alipay/trade/customs/declare` | ✅ |  |
| `alipay.trade.customs.query` | GET `/v3/alipay/trade/customs/query` | ✅ |  |

### services/trade/refund（3 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.trade.fastpay.refund.query` | POST `/v3/alipay/trade/fastpay/refund/query` | ✅ |  |
| `alipay.trade.refund` | POST `/v3/alipay/trade/refund` | ✅ |  |
| `alipay.trade.refund.apply` | POST `/v3/alipay/trade/refund/apply` | ✅ |  |

### services/trade/order（3 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.trade.order.pay` | POST `/v3/alipay/trade/order/pay` | ✅ |  |
| `alipay.trade.order.settle` | POST `/v3/alipay/trade/order/settle` | ✅ |  |
| `alipay.trade.order.settle.query` | GET `/v3/alipay/trade/order/settle/query` | ✅ |  |

### services/trade/orderinfo（1 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.trade.orderinfo.sync` | POST `/v3/alipay/trade/orderinfo/sync` | ✅ |  |

### services/trade/overseas（1 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.trade.overseas.settle` | POST `/v3/alipay/trade/overseas/settle` | ✅ |  |

### services/trade/royalty（4 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.trade.royalty.rate.query` | GET `/v3/alipay/trade/royalty/rate/query` | ✅ |  |
| `alipay.trade.royalty.relation.batchquery` | POST `/v3/alipay/trade/royalty/relation/batchquery` | ✅ |  |
| `alipay.trade.royalty.relation.bind` | POST `/v3/alipay/trade/royalty/relation/bind` | ✅ |  |
| `alipay.trade.royalty.relation.unbind` | POST `/v3/alipay/trade/royalty/relation/unbind` | ✅ |  |

### services/trade/settle（2 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.trade.settle.confirm` | POST `/v3/alipay/trade/settle/confirm` | ✅ | settle_info 的字段层级与快照一致；在线接口页未列账期时间。 |
| `alipay.trade.settle.receivables.query` | POST `/v3/alipay/trade/settle/receivables/query` | ✅ |  |

### services/user/agreement（8 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.user.agreement.executionplan.modify` | POST `/v3/alipay/user/agreement/executionplan/modify` | ✅ |  |
| `alipay.user.agreement.migrate` | POST `/v3/alipay/user/agreement/migrate` | ✅ |  |
| `alipay.user.agreement.permission.create` | POST `/v3/alipay/user/agreement/permission/create` | ✅ |  |
| `alipay.user.agreement.query` | GET `/v3/alipay/user/agreement/query` | ✅ |  |
| `alipay.user.agreement.transfer` | POST `/v3/alipay/user/agreement/transfer` | ✅ |  |
| `alipay.user.agreement.unsign` | POST `/v3/alipay/user/agreement/unsign` | ✅ |  |
| `alipay.user.agreement.userverify.apply` | POST `/v3/alipay/user/agreement/userverify/apply` | ✅ |  |
| `alipay.user.agreement.userverify.query` | GET `/v3/alipay/user/agreement/userverify/query` | ✅ |  |

### services/user/alipaypoint（2 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.user.alipaypoint.budgetlib.query` | GET `/v3/alipay/user/alipaypoint/budgetlib/query` | ✅ |  |
| `alipay.user.alipaypoint.send` | POST `/v3/alipay/user/alipaypoint/send` | ✅ |  |

### services/user/certdoc（2 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.user.certdoc.certverify.consult` | GET `/v3/alipay/user/certdoc/certverify/consult` | ✅ |  |
| `alipay.user.certdoc.certverify.preconsult` | POST `/v3/alipay/user/certdoc/certverify/preconsult` | ✅ |  |

### services/user/certify（2 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.user.certify.open.initialize` | POST `/v3/alipay/user/certify/open/initialize` | ✅ |  |
| `alipay.user.certify.open.query` | GET `/v3/alipay/user/certify/open/query` | ✅ |  |

### services/user/info（1 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.user.info.share` | POST `/v3/alipay/user/info/share` | ✅ |  |

### services/user/twostage（2 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.user.twostage.common.use` | POST `/v3/alipay/user/twostage/common/use` | ✅ |  |
| `alipay.user.twostage.indirect.use` | POST `/v3/alipay/user/twostage/indirect/use` | ✅ |  |

### services/user/userinfo（1 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `alipay.user.userinfo.share` | POST `/v3/alipay/user/userinfo/share` | ✅ |  |

### services/ant/merchant（25 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `ant.merchant.expand.apprecommend.account.create` | POST `/v3/ant/merchant/expand/apprecommend/account/create` | ✅ |  |
| `ant.merchant.expand.apprecommend.account.delete` | DELETE `/v3/ant/merchant/expand/apprecommend/account/delete` | ✅ |  |
| `ant.merchant.expand.apprecommend.account.query` | GET `/v3/ant/merchant/expand/apprecommend/account/query` | ✅ |  |
| `ant.merchant.expand.apprecommend.available.query` | GET `/v3/ant/merchant/expand/apprecommend/available/query` | ✅ |  |
| `ant.merchant.expand.indirect.image.upload` | POST `/v3/ant/merchant/indirect/image` | ✅ |  |
| `ant.merchant.expand.info.query` | GET `/v3/ant/merchant/expand/info/query` | ✅ |  |
| `ant.merchant.expand.item.create` | POST `/v3/ant/merchant/expand/item/create` | ✅ |  |
| `ant.merchant.expand.item.delete` | DELETE `/v3/ant/merchant/expand/item/delete` | ✅ |  |
| `ant.merchant.expand.item.modify` | POST `/v3/ant/merchant/expand/item/modify` | ✅ |  |
| `ant.merchant.expand.item.open.batchquery` | POST `/v3/ant/merchant/expand/item/open/batchquery` | ✅ |  |
| `ant.merchant.expand.item.open.create` | POST `/v3/ant/merchant/expand/item/open/create` | ✅ |  |
| `ant.merchant.expand.item.open.delete` | DELETE `/v3/ant/merchant/expand/item/open/delete` | ✅ |  |
| `ant.merchant.expand.item.open.modify` | POST `/v3/ant/merchant/expand/item/open/modify` | ✅ |  |
| `ant.merchant.expand.item.open.query` | GET `/v3/ant/merchant/expand/item/open/query` | ✅ |  |
| `ant.merchant.expand.item.query` | GET `/v3/ant/merchant/expand/item/query` | ✅ |  |
| `ant.merchant.expand.item.status.modify` | POST `/v3/ant/merchant/expand/item/status/modify` | ✅ |  |
| `ant.merchant.expand.mcc.query` | GET `/v3/ant/merchant/mcc/query` | ✅ |  |
| `ant.merchant.expand.order.query` | GET `/v3/ant/merchant/order/{order_id}` | ✅ |  |
| `ant.merchant.expand.shop.close` | PATCH `/v3/ant/merchant/shop/close` | ✅ |  |
| `ant.merchant.expand.shop.consult` | POST `/v3/ant/merchant/expand/shop/consult` | ✅ |  |
| `ant.merchant.expand.shop.create` | POST `/v3/ant/merchant/shop` | ✅ |  |
| `ant.merchant.expand.shop.modify` | PATCH `/v3/ant/merchant/shop` | ✅ |  |
| `ant.merchant.expand.shop.page.query` | GET `/v3/ant/merchant/shop/pagequery` | ✅ |  |
| `ant.merchant.expand.shop.query` | GET `/v3/ant/merchant/shop` | ✅ |  |
| `ant.merchant.expand.shop.receiptaccount.save` | POST `/v3/ant/merchant/expand/shop/receiptaccount/save` | ✅ |  |

### services/datadigital/fincloud（10 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `datadigital.fincloud.generalsaas.face.certify.initialize` | POST `/v3/datadigital/fincloud/generalsaas/face/certify/initialize` | ✅ |  |
| `datadigital.fincloud.generalsaas.face.certify.query` | GET `/v3/datadigital/fincloud/generalsaas/face/certify/query` | ✅ |  |
| `datadigital.fincloud.generalsaas.face.certify.verify` | POST `/v3/datadigital/fincloud/generalsaas/face/certify/verify` | ✅ |  |
| `datadigital.fincloud.generalsaas.face.check.initialize` | POST `/v3/datadigital/fincloud/generalsaas/face/check/initialize` | ✅ |  |
| `datadigital.fincloud.generalsaas.face.check.query` | GET `/v3/datadigital/fincloud/generalsaas/face/check/query` | ✅ |  |
| `datadigital.fincloud.generalsaas.face.source.certify` | POST `/v3/datadigital/fincloud/generalsaas/face/source/certify` | ✅ |  |
| `datadigital.fincloud.generalsaas.face.verification.initialize` | POST `/v3/datadigital/fincloud/generalsaas/face/verification/initialize` | ✅ |  |
| `datadigital.fincloud.generalsaas.face.verification.query` | GET `/v3/datadigital/fincloud/generalsaas/face/verification/query` | ✅ |  |
| `datadigital.fincloud.generalsaas.ocr.mobile.initialize` | POST `/v3/datadigital/fincloud/generalsaas/ocr/mobile/initialize` | ✅ |  |
| `datadigital.fincloud.generalsaas.ocr.server.detect` | POST `/v3/datadigital/fincloud/generalsaas/ocr/server/detect` | ✅ |  |

### services/grandsecurity/bizrisks（1 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `grandsecurity.bizrisks.fact.check` | POST `/v3/stream/grandsecurity/bizrisks/fact/check` | ✅ |  |

### services/monitor（1 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `monitor.heartbeat.syn` | POST `/v3/monitor/heartbeat/syn` | ✅ |  |

### services/zhima/credit（13 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `zhima.credit.payafteruse.creditagreement.query` | GET `/v3/zhima/credit/payafteruse/creditagreement/query` | ✅ |  |
| `zhima.credit.payafteruse.creditagreement.transfer` | POST `/v3/zhima/credit/payafteruse/creditagreement/transfer` | ✅ |  |
| `zhima.credit.payafteruse.creditbizorder.finish` | POST `/v3/zhima/credit/payafteruse/creditbizorder/finish` | ✅ |  |
| `zhima.credit.payafteruse.creditbizorder.order` | POST `/v3/zhima/credit/payafteruse/creditbizorder/order` | ✅ |  |
| `zhima.credit.payafteruse.creditbizorder.query` | GET `/v3/zhima/credit/payafteruse/creditbizorder/query` | ✅ |  |
| `zhima.credit.pe.zmgo.agreement.query` | GET `/v3/zhima/credit/pe/zmgo/agreement/query` | ✅ |  |
| `zhima.credit.pe.zmgo.agreement.unsign` | POST `/v3/zhima/credit/pe/zmgo/agreement/unsign` | ✅ |  |
| `zhima.credit.pe.zmgo.bizopt.close` | POST `/v3/zhima/credit/pe/zmgo/bizopt/close` | ✅ |  |
| `zhima.credit.pe.zmgo.cumulation.sync` | POST `/v3/zhima/credit/pe/zmgo/cumulation/sync` | ✅ |  |
| `zhima.credit.pe.zmgo.preorder.create` | PUT `/v3/zhima/credit/pe/zmgo/preorder/create` | ✅ |  |
| `zhima.credit.pe.zmgo.settle.apply` | POST `/v3/zhima/credit/pe/zmgo/settle/apply` | ✅ |  |
| `zhima.credit.pe.zmgo.settle.refund` | POST `/v3/zhima/credit/pe/zmgo/settle/refund` | ✅ |  |
| `zhima.credit.pe.zmgo.settle.unfreeze` | PUT `/v3/zhima/credit/pe/zmgo/settle/unfreeze` | ✅ |  |

### services/zhima/customer（6 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `zhima.customer.jobworth.authentication.query` | GET `/v3/zhima/customer/jobworth/authentication` | ✅ |  |
| `zhima.customer.jobworth.cloudresume.query` | GET `/v3/zhima/customer/jobworth/cloudresume/query` | ✅ |  |
| `zhima.customer.jobworth.info.query` | GET `/v3/zhima/customer/jobworth/info/query` | ✅ |  |
| `zhima.customer.jobworth.jobdata.add` | POST `/v3/zhima/customer/jobworth/jobdata/add` | ✅ |  |
| `zhima.customer.jobworth.picture.upload` | POST `/v3/zhima/customer/jobworth/picture/upload` | ✅ |  |
| `zhima.customer.jobworth.position.add` | POST `/v3/zhima/customer/jobworth/position/add` | ✅ |  |

### services/zhima/merchant（6 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `zhima.merchant.subsidiaries.apply` | POST `/v3/zhima/merchant/subsidiaries/apply` | ✅ |  |
| `zhima.merchant.subsidiaries.close` | POST `/v3/zhima/merchant/subsidiaries/close` | ✅ |  |
| `zhima.merchant.zmgo.cumulate.query` | GET `/v3/zhima/merchant/zmgo/cumulate/query` | ✅ |  |
| `zhima.merchant.zmgo.cumulate.sync` | POST `/v3/zhima/merchant/zmgo/cumulate/sync` | ✅ |  |
| `zhima.merchant.zmgo.template.create` | POST `/v3/zhima/merchant/zmgo/template/create` | ✅ |  |
| `zhima.merchant.zmgo.template.query` | GET `/v3/zhima/merchant/zmgo/template/query` | ✅ |  |

### services/zoloz/authentication（5 项）

| 接口 | 方法与路径 | 结果 | 说明 |
| --- | --- | --- | --- |
| `zoloz.authentication.customer.facemanage.create` | POST `/v3/zoloz/authentication/customer/facemanage/create` | ✅ |  |
| `zoloz.authentication.customer.facemanage.delete` | POST `/v3/zoloz/authentication/customer/facemanage/delete` | ✅ |  |
| `zoloz.authentication.customer.ftoken.query` | POST `/v3/zoloz/authentication/customer/ftoken/query` | ✅ |  |
| `zoloz.authentication.customer.smilepay.initialize` | POST `/v3/zoloz/authentication/customer/smilepay/initialize` | ✅ |  |
| `zoloz.authentication.smilepay.initialize` | POST `/v3/zoloz/authentication/smilepay/initialize` | ✅ |  |

## 验证边界

逐接口明细针对上述固定版本的字段契约。网络、商户权限、条件必填、实际资金处理和通知语义需分别验证。本次运行离线测试，未进行沙箱或生产交易。
