# 支付宝接口目录

整理日期：2026-11-05。此目录收录支付宝官方 V3 OpenAPI 快照中的全部 746 个 REST 操作、已核验的 3 个页面和唤端接口，以及本次从官方资料补录的 35 个快照外接口。补录接口的 V3 契约尚未确认，不计入已确认的 V3 接口数量。此目录仍不是支付宝全部接口的完整清单。

接口名、HTTP 方法和路径来自官方 V3 OpenAPI 快照（`info.version=2026-09-23`，提交 `02177d34fce069f6c20d83c13f963150bb45467b`）。支付宝在[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)中提供描述文件。接口名称链接到支付宝文档中心对应的 V3 接口页，其中 `pre-apis` 为定向开放接口文档。产品权限以应用实际签约情况为准。

快照外接口的来源是官方[文档索引](https://opendocs.alipay.com/common/llms.txt)、[文档摘要](https://opendocs.alipay.com/common/llms-full.txt)指向的接口全文及[资金操作类敏感接口列表](https://opendocs.alipay.com/common/02kkuw.md)。下方补录部分分别说明已取得接口全文和仅有清单列名的核验状态；清单列名不代表接口目前仍可开通或调用。

目录只表示官方资料中存在，不表示本 SDK 已实现、经过沙箱验证或可用于生产。当前实现及验证状态见[接口覆盖记录](api-coverage.md)，支持范围与契约清点结论见[接口版本说明](api-versions.md)。实现具体接口前应查阅其官方参数表、授权和协议要求。

## 页面与唤端接口

以下接口使用官方 V3 文档和 SDK 的页面或订单串构建方式，不属于下方 REST 清单，不能推导 REST 路由。

| 接口 | 用途 | 官方文档 |
| --- | --- | --- |
| `alipay.trade.page.pay` | 电脑网站支付 | [接口文档](https://opendocs.alipay.com/open-v3/2423fad5_alipay.trade.page.pay) |
| `alipay.trade.wap.pay` | 手机网站支付 | [接口文档](https://opendocs.alipay.com/open-v3/1a957be0_alipay.trade.wap.pay) |
| `alipay.trade.app.pay` | App 支付 | [接口文档](https://opendocs.alipay.com/open-v3/429e4d75_alipay.trade.app.pay)、[支付并签约场景](https://opendocs.alipay.com/open-v3/0jf33m) |

## V3 REST 接口

按接口名的前两段分组。HTTP 方法与路径直接取自官方 OpenAPI；单击接口名称可查看支付宝开放平台的接口文档。

### `alios.open`（1）

| 接口 | 名称 | 方法 | 路径 |
| --- | --- | --- | --- |
| [`alios.open.auto.info.query`](https://opendocs.alipay.com/pre-apis/05igwf) | 查询阿里车的车辆信息 | `GET` | `/v3/alios/open/auto/info/query` |

### `alipay.boss`（8）

| 接口 | 名称 | 方法 | 路径 |
| --- | --- | --- | --- |
| [`alipay.boss.fnc.invoice.apply`](https://opendocs.alipay.com/pre-apis/05xo5f) | 应收开票申请(根据账单申请) | `POST` | `/v3/alipay/boss/fnc/invoice/apply` |
| [`alipay.boss.fnc.invoice.batchquery`](https://opendocs.alipay.com/pre-apis/05xdyx) | 应收发票综合查询 | `POST` | `/v3/alipay/boss/fnc/invoice/batchquery` |
| [`alipay.boss.fnc.invoice.query`](https://opendocs.alipay.com/pre-apis/05xhqn) | 根据发票ID查询发票信息 | `GET` | `/v3/alipay/boss/fnc/invoice/query` |
| [`alipay.boss.fnc.invoicereceipt.batchquery`](https://opendocs.alipay.com/pre-apis/05xiqt) | 根据月账单号集合和来源查询开票单据及金额汇总 | `POST` | `/v3/alipay/boss/fnc/invoicereceipt/batchquery` |
| [`alipay.boss.fnc.invoicereceipt.create`](https://opendocs.alipay.com/pre-apis/05xnh4) | 账单变更同步接口月帐单创建消息同步 | `POST` | `/v3/alipay/boss/fnc/invoicereceipt/create` |
| [`alipay.boss.fnc.userinvoiceinfo.create`](https://opendocs.alipay.com/pre-apis/05xgfe) | 新增客户开票资料 | `POST` | `/v3/alipay/boss/fnc/userinvoiceinfo/create` |
| [`alipay.boss.fnc.userinvoiceinfo.modify`](https://opendocs.alipay.com/pre-apis/05xshh) | 修改用户开票资料 | `POST` | `/v3/alipay/boss/fnc/userinvoiceinfo/modify` |
| [`alipay.boss.fnc.userinvoiceinfo.query`](https://opendocs.alipay.com/pre-apis/05xmie) | 根据PID获取相关的开票资料 | `GET` | `/v3/alipay/boss/fnc/userinvoiceinfo/query` |

### `alipay.commerce`（55）

| 接口 | 名称 | 方法 | 路径 |
| --- | --- | --- | --- |
| [`alipay.commerce.cityfacilitator.station.query`](https://opendocs.alipay.com/apis/05wvuo) | 地铁购票站点数据查询 | `GET` | `/v3/alipay/commerce/cityfacilitator/station/query` |
| [`alipay.commerce.cityfacilitator.voucher.batchquery`](https://opendocs.alipay.com/apis/05wol9) | 地铁购票订单批量查询 | `POST` | `/v3/alipay/commerce/cityfacilitator/voucher/batchquery` |
| [`alipay.commerce.cityfacilitator.voucher.generate`](https://opendocs.alipay.com/apis/05wz9v) | 地铁购票核销码发码 | `POST` | `/v3/alipay/commerce/cityfacilitator/voucher/generate` |
| [`alipay.commerce.cityfacilitator.voucher.refund`](https://opendocs.alipay.com/apis/05wujr) | 地铁购票发码退款 | `POST` | `/v3/alipay/commerce/cityfacilitator/voucher/refund` |
| [`alipay.commerce.ec.balance.downloadurl.query`](https://opendocs.alipay.com/apis/05j1bx) | 对账单文件下载接口 | `GET` | `/v3/alipay/commerce/ec/balance/downloadurl/query` |
| [`alipay.commerce.ec.balance.period.modify`](https://opendocs.alipay.com/apis/05j82h) | 企业码月对账单账期修改 | `POST` | `/v3/alipay/commerce/ec/balance/period/modify` |
| [`alipay.commerce.ec.consume.detail.query`](https://opendocs.alipay.com/apis/05jgvi) | 企业码账单详情查询 | `POST` | `/v3/alipay/commerce/ec/consume/detail/query` |
| [`alipay.commerce.ec.department.create`](https://opendocs.alipay.com/apis/05itkw) | 企业部门创建 | `POST` | `/v3/alipay/commerce/ec/department` |
| [`alipay.commerce.ec.department.delete`](https://opendocs.alipay.com/apis/05jgvm) | 企业部门删除 | `POST` | `/v3/alipay/commerce/ec/department/delete` |
| [`alipay.commerce.ec.department.info.modify`](https://opendocs.alipay.com/apis/05iwe0) | 企业部门信息修改 | `PUT` | `/v3/alipay/commerce/ec/department` |
| [`alipay.commerce.ec.department.info.query`](https://opendocs.alipay.com/apis/05uggp) | 查询部门详情 | `GET` | `/v3/alipay/commerce/ec/department/info/query` |
| [`alipay.commerce.ec.department.sublist.query`](https://opendocs.alipay.com/apis/05uipk) | 查询子部门id列表 | `GET` | `/v3/alipay/commerce/ec/department/sublist/query` |
| [`alipay.commerce.ec.employee.add`](https://opendocs.alipay.com/apis/05lo9i) | 添加员工 | `POST` | `/v3/alipay/commerce/ec/employee/add` |
| [`alipay.commerce.ec.employee.delete`](https://opendocs.alipay.com/apis/05j06m) | 删除员工 | `POST` | `/v3/alipay/commerce/ec/employee/delete` |
| [`alipay.commerce.ec.employee.idlist.query`](https://opendocs.alipay.com/apis/05udr4) | 查询部门下员工id列表 | `GET` | `/v3/alipay/commerce/ec/employee/idlist/query` |
| [`alipay.commerce.ec.employee.info.modify`](https://opendocs.alipay.com/apis/05j1c0) | 修改员工基础信息 | `PUT` | `/v3/alipay/commerce/ec/employee/info` |
| [`alipay.commerce.ec.employee.info.query`](https://opendocs.alipay.com/apis/05uhw6) | 查询员工详情 | `GET` | `/v3/alipay/commerce/ec/employee/info/query` |
| [`alipay.commerce.ec.employee.invite.query`](https://opendocs.alipay.com/apis/05j9je) | 获取员工签约激活链接 | `GET` | `/v3/alipay/commerce/ec/employee/invite` |
| [`alipay.commerce.ec.employee.title.create`](https://opendocs.alipay.com/apis/05ip0u) | 员工抬头：新增员工抬头关系 | `POST` | `/v3/alipay/commerce/ec/employee/title` |
| [`alipay.commerce.ec.employee.title.delete`](https://opendocs.alipay.com/apis/05imon) | 员工抬头：失效员工抬头关系 | `POST` | `/v3/alipay/commerce/ec/employee/title/delete` |
| [`alipay.commerce.ec.employee.title.modify`](https://opendocs.alipay.com/apis/05j4cg) | 修改员工抬头信息 | `PUT` | `/v3/alipay/commerce/ec/employee/title` |
| [`alipay.commerce.ec.enterprise.address.add`](https://opendocs.alipay.com/apis/05j04z) | 企业地址添加 | `POST` | `/v3/alipay/commerce/ec/enterprise/address` |
| [`alipay.commerce.ec.enterprise.address.modify`](https://opendocs.alipay.com/apis/05j1bb) | 企业地址修改 | `PUT` | `/v3/alipay/commerce/ec/enterprise/address` |
| [`alipay.commerce.ec.enterprise.address.query`](https://opendocs.alipay.com/apis/05iz2p) | 企业地址查询 | `GET` | `/v3/alipay/commerce/ec/enterprise/address` |
| [`alipay.commerce.ec.enterprise.agreement.query`](https://opendocs.alipay.com/apis/05j2cw) | 查询企业签约企业码信息 | `GET` | `/v3/alipay/commerce/ec/enterprise/agreement` |
| [`alipay.commerce.ec.enterprise.create`](https://opendocs.alipay.com/pre-apis/05l7n7) | 企业签约注册-接口已过时 | `POST` | `/v3/alipay/commerce/ec/enterprise/create` |
| [`alipay.commerce.ec.enterprise.delete`](https://opendocs.alipay.com/apis/05j82m) | 企业注销 | `POST` | `/v3/alipay/commerce/ec/enterprise/delete` |
| [`alipay.commerce.ec.enterprise.info.modify`](https://opendocs.alipay.com/apis/05ip0v) | 修改企业基础信息 | `PUT` | `/v3/alipay/commerce/ec/enterprise` |
| [`alipay.commerce.ec.enterprise.info.query`](https://opendocs.alipay.com/apis/05iwc6) | 查询企业详情 | `GET` | `/v3/alipay/commerce/ec/enterprise/info` |
| [`alipay.commerce.ec.enterprise.unsign`](https://opendocs.alipay.com/apis/05jach) | 企业解约 | `GET` | `/v3/alipay/commerce/ec/enterprise/unsignurl` |
| [`alipay.commerce.ec.jointaccountbill.detail.batchquery`](https://opendocs.alipay.com/apis/05j4cf) | 企业账单批量分页查询 | `GET` | `/v3/alipay/commerce/ec/jointaccountbill/detail/batchquery` |
| [`alipay.commerce.ec.user.enterprise.query`](https://opendocs.alipay.com/apis/05j82l) | 用户企业列表 | `GET` | `/v3/alipay/commerce/ec/user/enterprise` |
| [`alipay.commerce.educate.campus.institutions.add`](https://opendocs.alipay.com/open-v3/2f573ce0_alipay.commerce.educate.campus.institutions.add) | 入驻学校信息 | `POST` | `/v3/alipay/commerce/educate/campus/institutions/add` |
| [`alipay.commerce.educate.campus.institutions.query`](https://opendocs.alipay.com/open-v3/b340d586_alipay.commerce.educate.campus.institutions.query) | 查询学校信息 | `GET` | `/v3/alipay/commerce/educate/campus/institutions/query` |
| [`alipay.commerce.educate.schoolcard.order.sync`](https://opendocs.alipay.com/open-v3/b4497dd8_alipay.commerce.educate.schoolcard.order.sync) | 校园卡支付订单数据同步 | `POST` | `/v3/alipay/commerce/educate/schoolcard/order/sync` |
| [`alipay.commerce.iot.device.report.upload`](https://opendocs.alipay.com/pre-apis/05xk46) | 设备检测报告上传 | `POST` | `/v3/alipay/commerce/iot/device/report/upload` |
| [`alipay.commerce.iot.sdarttool.message.query`](https://opendocs.alipay.com/pre-apis/06bpqm) | 消息查询 | `GET` | `/v3/alipay/commerce/iot/sdarttool/message/query` |
| [`alipay.commerce.iot.sdarttool.message.send`](https://opendocs.alipay.com/pre-apis/06bzwt) | IOT消息服务 | `POST` | `/v3/alipay/commerce/iot/sdarttool/message/send` |
| [`alipay.commerce.logistics.invoice.istdwaybill.create`](https://opendocs.alipay.com/pre-apis/068d50) | 开即时配送订单发票 | `POST` | `/v3/alipay/commerce/logistics/invoice/istdwaybill/create` |
| [`alipay.commerce.logistics.invoice.istdwaybill.query`](https://opendocs.alipay.com/pre-apis/068lff) | 查询即时配送订单的开票结果 | `GET` | `/v3/alipay/commerce/logistics/invoice/istdwaybill/query` |
| [`alipay.commerce.logistics.logisticscompany.instantdelivery.query`](https://opendocs.alipay.com/pre-apis/068msd) | 查询支付宝支持的即时配送公司 | `GET` | `/v3/alipay/commerce/logistics/logisticscompany/instantdelivery/query` |
| [`alipay.commerce.logistics.order.instantdelivery.cancel`](https://opendocs.alipay.com/pre-apis/068eft) | 取消即时配送订单 | `POST` | `/v3/alipay/commerce/logistics/order/instantdelivery/cancel` |
| [`alipay.commerce.logistics.order.instantdelivery.create`](https://opendocs.alipay.com/pre-apis/05lo9w) | 下即时配送订单 | `POST` | `/v3/alipay/commerce/logistics/order/instantdelivery/create` |
| [`alipay.commerce.logistics.order.instantdelivery.precreate`](https://opendocs.alipay.com/pre-apis/05l560) | 预下即时配送订单 | `POST` | `/v3/alipay/commerce/logistics/order/instantdelivery/precreate` |
| [`alipay.commerce.logistics.order.istdcancel.preconsult`](https://opendocs.alipay.com/pre-apis/068d52) | 预取消即时配送订单 | `POST` | `/v3/alipay/commerce/logistics/order/istdcancel/preconsult` |
| [`alipay.commerce.logistics.order.istdretry.create`](https://opendocs.alipay.com/pre-apis/05l561) | 重下即时配送订单 | `POST` | `/v3/alipay/commerce/logistics/order/istdretry/create` |
| [`alipay.commerce.logistics.waybill.istddetail.query`](https://opendocs.alipay.com/pre-apis/068fh3) | 查询即时配送运单详情 | `GET` | `/v3/alipay/commerce/logistics/waybill/istddetail/query` |
| [`alipay.commerce.operation.activity.merchant.modify`](https://opendocs.alipay.com/pre-apis/063rs2) | 商户签约信息修改 | `POST` | `/v3/alipay/commerce/operation/activity/merchant/modify` |
| [`alipay.commerce.operation.activity.merchant.sign`](https://opendocs.alipay.com/pre-apis/06347n) | 商户签约活动 | `POST` | `/v3/alipay/commerce/operation/activity/merchant/sign` |
| [`alipay.commerce.operation.activity.merchant.unsign`](https://opendocs.alipay.com/pre-apis/06336j) | 活动商户解约 | `POST` | `/v3/alipay/commerce/operation/activity/merchant/unsign` |
| [`alipay.commerce.operation.bank.activity.query`](https://opendocs.alipay.com/pre-apis/063ecd) | 银行活动信息查询 | `GET` | `/v3/alipay/commerce/operation/bank/activity/query` |
| [`alipay.commerce.transport.charger.chargerbindinfo.sync`](https://opendocs.alipay.com/pre-apis/05xqw6) | 充电运营商同步用户绑定充电桩信息接口 | `POST` | `/v3/alipay/commerce/transport/charger/chargerbindinfo/sync` |
| [`alipay.commerce.transport.parking.enterinfo.sync`](https://opendocs.alipay.com/open-v3/e38b4d57_alipay.commerce.transport.parking.enterinfo.sync) | 停车入场模板消息配置 | `POST` | `/v3/alipay/commerce/transport/parking/enterinfo/sync` |
| [`alipay.commerce.transport.parking.exitinfo.sync`](https://opendocs.alipay.com/open-v3/9986eded_alipay.commerce.transport.parking.exitinfo.sync) | 停车离场消息模板配置 | `POST` | `/v3/alipay/commerce/transport/parking/exitinfo/sync` |
| [`alipay.commerce.transport.parking.paymentinfo.sync`](https://opendocs.alipay.com/open-v3/12f22c8f_alipay.commerce.transport.parking.paymentinfo.sync) | 停车支付模板消息配置 | `POST` | `/v3/alipay/commerce/transport/parking/paymentinfo/sync` |

### `alipay.data`（18）

| 接口 | 名称 | 方法 | 路径 |
| --- | --- | --- | --- |
| [`alipay.data.bill.accountbookereceipt.apply`](https://opendocs.alipay.com/pre-apis/05wyde) | 申请子账本电子回单(incubating) | `POST` | `/v3/alipay/data/bill/accountbookereceipt/apply` |
| [`alipay.data.bill.accountbookereceipt.query`](https://opendocs.alipay.com/open-v3/6ce2ea1b_alipay.data.bill.accountbookereceipt.query) | 查询子账本电子回单状态(incubating) | `GET` | `/v3/alipay/data/bill/accountbookereceipt/query` |
| [`alipay.data.bill.accountlog.query`](https://opendocs.alipay.com/open-v3/4cce9615_alipay.data.bill.accountlog.query) | 支付宝商家账户账务明细查询 | `GET` | `/v3/alipay/data/bill/accountlog/query` |
| [`alipay.data.bill.bail.query`](https://opendocs.alipay.com/open-v3/4421b78a_alipay.data.bill.bail.query) | 支付宝商家账户保证金查询 | `GET` | `/v3/alipay/data/bill/bail/query` |
| [`alipay.data.bill.balance.query`](https://opendocs.alipay.com/open-v3/9b66d9d8_alipay.data.bill.balance.query) | 支付宝商家账户当前余额查询 | `GET` | `/v3/alipay/data/bill/balance/query` |
| [`alipay.data.bill.balancehis.query`](https://opendocs.alipay.com/open-v3/029a9c24_alipay.data.bill.balancehis.query) | 支付宝商家账户历史余额查询 | `GET` | `/v3/alipay/data/bill/balancehis/query` |
| [`alipay.data.bill.bizfundagent.query`](https://opendocs.alipay.com/open-v3/9b60baa5_alipay.data.bill.bizfundagent.query) | ISV代理商户资金业务账单查询 | `GET` | `/v3/alipay/data/bill/bizfundagent/query` |
| [`alipay.data.bill.buy.query`](https://opendocs.alipay.com/open-v3/83504d51_alipay.data.bill.buy.query) | 支付宝商家账户买入交易查询 | `GET` | `/v3/alipay/data/bill/buy/query` |
| [`alipay.data.bill.ereceipt.apply`](https://opendocs.alipay.com/open-v3/46460043_alipay.data.bill.ereceipt.apply) | 申请电子回单(incubating) | `POST` | `/v3/alipay/data/bill/ereceipt/apply` |
| [`alipay.data.bill.ereceipt.query`](https://opendocs.alipay.com/open-v3/65fde737_alipay.data.bill.ereceipt.query) | 查询电子回单状态(incubating) | `GET` | `/v3/alipay/data/bill/ereceipt/query` |
| [`alipay.data.bill.ereceiptagent.apply`](https://opendocs.alipay.com/open-v3/d8c0f195_alipay.data.bill.ereceiptagent.apply) | ISV代理商户申请电子回单 | `POST` | `/v3/alipay/data/bill/ereceiptagent/apply` |
| [`alipay.data.bill.sell.query`](https://opendocs.alipay.com/open-v3/3c1a5835_alipay.data.bill.sell.query) | 支付宝商家账户卖出交易查询 | `GET` | `/v3/alipay/data/bill/sell/query` |
| [`alipay.data.bill.transfer.query`](https://opendocs.alipay.com/open-v3/8ddb5fed_alipay.data.bill.transfer.query) | 支付宝商家账户充值，转账，提现查询 | `GET` | `/v3/alipay/data/bill/transfer/query` |
| [`alipay.data.bill.transferaccountbook.query`](https://opendocs.alipay.com/pre-apis/05wycr) | 子账本充提转账单查询(incubating) | `GET` | `/v3/alipay/data/bill/transferaccountbook/query` |
| [`alipay.data.dataservice.ad.conversion.upload`](https://opendocs.alipay.com/open-v3/32a2992f_alipay.data.dataservice.ad.conversion.upload) | 转化数据回传 | `POST` | `/v3/alipay/data/dataservice/ad/conversion/upload` |
| [`alipay.data.dataservice.ad.promotepage.batchquery`](https://opendocs.alipay.com/open-v3/17d7cd6f_alipay.data.dataservice.ad.promotepage.batchquery) | 自建推广页列表批量查询 | `GET` | `/v3/alipay/data/dataservice/ad/promotepage/batchquery` |
| [`alipay.data.dataservice.ad.promotepage.download`](https://opendocs.alipay.com/open-v3/2887027f_alipay.data.dataservice.ad.promotepage.download) | 自建推广页留资数据查询 | `GET` | `/v3/alipay/data/dataservice/ad/promotepage/download` |
| [`alipay.data.dataservice.bill.downloadurl.query`](https://opendocs.alipay.com/open-v3/c8d608d7_alipay.data.dataservice.bill.downloadurl.query?scene=common&pathHash=de817b49) | 查询对账单下载地址 | `GET` | `/v3/alipay/data/dataservice/bill/downloadurl/query` |

### `alipay.ebpp`（77）

| 接口 | 名称 | 方法 | 路径 |
| --- | --- | --- | --- |
| [`alipay.ebpp.bill.add`](https://opendocs.alipay.com/open-v3/b9117832_alipay.ebpp.bill.add) | 创建账单 | `POST` | `/v3/alipay/ebpp/bill/add` |
| [`alipay.ebpp.bill.get`](https://opendocs.alipay.com/open-v3/b6ddabc9_alipay.ebpp.bill.get) | 查询账单 | `GET` | `/v3/alipay/ebpp/bill/get` |
| [`alipay.ebpp.invoice.apply.inv.upload`](https://opendocs.alipay.com/apis/05ujox) | 发票回传 | `PUT` | `/v3/alipay/ebpp/invoice/apply/inv/upload` |
| [`alipay.ebpp.invoice.apply.result.sync`](https://opendocs.alipay.com/apis/05vwzk) | ISV向支付宝同步发票申请结果 | `POST` | `/v3/alipay/ebpp/invoice/apply/result/sync` |
| [`alipay.ebpp.invoice.apply.status.notify`](https://opendocs.alipay.com/apis/05uo9u) | 发票申请状态变更 | `PUT` | `/v3/alipay/ebpp/invoice/apply/status/notify` |
| [`alipay.ebpp.invoice.applystatus.query`](https://opendocs.alipay.com/apis/05jve5) | 根据外部订单号查询开票状态 | `POST` | `/v3/alipay/ebpp/invoice/applystatus/query` |
| [`alipay.ebpp.invoice.auth.sign`](https://opendocs.alipay.com/pre-apis/05jtxa) | 发票授权关系签约 | `POST` | `/v3/alipay/ebpp/invoice/auth/sign` |
| [`alipay.ebpp.invoice.auth.unsign`](https://opendocs.alipay.com/apis/05jwd5) | 发票授权关系解约 | `POST` | `/v3/alipay/ebpp/invoice/auth/unsign` |
| [`alipay.ebpp.invoice.detail.output.query`](https://opendocs.alipay.com/apis/05jxvp) | 获取发票详细信息 | `POST` | `/v3/alipay/ebpp/invoice/detail/output/query` |
| [`alipay.ebpp.invoice.ecorder.order.query`](https://opendocs.alipay.com/apis/05imol) | 企业码订单查询 | `GET` | `/v3/alipay/ebpp/invoice/ecorder/order/query` |
| [`alipay.ebpp.invoice.einvpackage.query`](https://opendocs.alipay.com/apis/05wvv7) | 查询用户发票包的明细信息 | `GET` | `/v3/alipay/ebpp/invoice/einvpackage/query` |
| [`alipay.ebpp.invoice.enterprise.merchantrelation.create`](https://opendocs.alipay.com/apis/05psjg) | 创建企业商户关系 | `POST` | `/v3/alipay/ebpp/invoice/enterprise/merchantrelation/create` |
| [`alipay.ebpp.invoice.enterpriseconsume.consume.batchquery`](https://opendocs.alipay.com/apis/05jgvh) | 企业因公账单查询 | `POST` | `/v3/alipay/ebpp/invoice/enterpriseconsume/consume/batchquery` |
| [`alipay.ebpp.invoice.enterpriseconsume.detail.query`](https://opendocs.alipay.com/apis/05j4ce) | 凭证详情查询接口 | `GET` | `/v3/alipay/ebpp/invoice/enterpriseconsume/detail/query` |
| [`alipay.ebpp.invoice.enterpriseconsume.enterpriseopenrule.create`](https://opendocs.alipay.com/apis/05wvsl) | 新增开票规则 | `POST` | `/v3/alipay/ebpp/invoice/enterpriseconsume/enterpriseopenrule/create` |
| [`alipay.ebpp.invoice.enterpriseconsume.enterpriseopenrule.modify`](https://opendocs.alipay.com/apis/05wzal) | 修改开票规则 | `POST` | `/v3/alipay/ebpp/invoice/enterpriseconsume/enterpriseopenrule/modify` |
| [`alipay.ebpp.invoice.enterpriseconsume.enterpriseopenrule.query`](https://opendocs.alipay.com/apis/05wye3) | 查询开票规则 | `GET` | `/v3/alipay/ebpp/invoice/enterpriseconsume/enterpriseopenrule/query` |
| [`alipay.ebpp.invoice.enterpriseconsume.openrulerelation.create`](https://opendocs.alipay.com/apis/05ucs6) | 新增开票规则并绑定费控规则 | `POST` | `/v3/alipay/ebpp/invoice/enterpriseconsume/openrulerelation/create` |
| [`alipay.ebpp.invoice.enterpriseconsume.relatedetail.query`](https://opendocs.alipay.com/apis/05ip0t) | 合花账单关联详情查询 | `GET` | `/v3/alipay/ebpp/invoice/enterpriseconsume/relatedetail/query` |
| [`alipay.ebpp.invoice.enterpriseconsume.summaryinvoicedetail.query`](https://opendocs.alipay.com/apis/05imom) | 查询汇总发票详情 | `GET` | `/v3/alipay/ebpp/invoice/enterpriseconsume/summaryinvoicedetail/query` |
| [`alipay.ebpp.invoice.enterpriseexctrl.employertitle.batchquery`](https://opendocs.alipay.com/apis/05wvun) | 企业抬头：分页查询企业抬头信息 | `POST` | `/v3/alipay/ebpp/invoice/enterpriseexctrl/employertitle/batchquery` |
| [`alipay.ebpp.invoice.enterpriseexctrl.employertitle.create`](https://opendocs.alipay.com/apis/05wydt) | 企业抬头：新增企业抬头 | `POST` | `/v3/alipay/ebpp/invoice/enterpriseexctrl/employertitle/create` |
| [`alipay.ebpp.invoice.enterpriseexctrl.employertitle.modify`](https://opendocs.alipay.com/apis/05lhsq) | 企业抬头：修改企业抬头 | `PUT` | `/v3/alipay/ebpp/invoice/enterpriseexctrl/employertitle` |
| [`alipay.ebpp.invoice.enterpriseexctrl.employertitle.query`](https://opendocs.alipay.com/apis/05x1fx) | 企业抬头：查询企业抬头 | `GET` | `/v3/alipay/ebpp/invoice/enterpriseexctrl/employertitle/query` |
| [`alipay.ebpp.invoice.expense.progress.sync`](https://opendocs.alipay.com/apis/05wx2n) | 发票报销处理进度同步 | `POST` | `/v3/alipay/ebpp/invoice/expense/progress/sync` |
| [`alipay.ebpp.invoice.expensecomsue.outsource.notify`](https://opendocs.alipay.com/apis/05jgvn) | 同步外部额度变更（交易形式） | `POST` | `/v3/alipay/ebpp/invoice/expensecomsue/outsource/notify` |
| [`alipay.ebpp.invoice.expensecontrol.aggregation.create`](https://opendocs.alipay.com/apis/05xsf0) | 创建费控规则聚合关系 | `POST` | `/v3/alipay/ebpp/invoice/expensecontrol/aggregation/create` |
| [`alipay.ebpp.invoice.expensecontrol.issuebatch.cancel`](https://opendocs.alipay.com/apis/05wumt) | 作废额度发放 | `POST` | `/v3/alipay/ebpp/invoice/expensecontrol/issuebatch/cancel` |
| [`alipay.ebpp.invoice.expensecontrol.issuebatch.create`](https://opendocs.alipay.com/apis/05j9jf) | 手动发放额度接口 | `POST` | `/v3/alipay/ebpp/invoice/expensecontrol/issuebatch/create` |
| [`alipay.ebpp.invoice.expensecontrol.quota.create`](https://opendocs.alipay.com/apis/05iz32) | 创建余额/点券 | `POST` | `/v3/alipay/ebpp/invoice/expensecontrol/quota/create` |
| [`alipay.ebpp.invoice.expensecontrol.quota.modify`](https://opendocs.alipay.com/apis/05wuoa) | 修改余额/点券 | `POST` | `/v3/alipay/ebpp/invoice/expensecontrol/quota/modify` |
| [`alipay.ebpp.invoice.expensecontrol.quota.query`](https://opendocs.alipay.com/apis/05imop) | 查询余额/点券 | `POST` | `/v3/alipay/ebpp/invoice/expensecontrol/quota/query` |
| [`alipay.ebpp.invoice.expenserules.employee.query`](https://opendocs.alipay.com/apis/05j4cj) | 查询费控规则员工 | `GET` | `/v3/alipay/ebpp/invoice/expenserules/employee/query` |
| [`alipay.ebpp.invoice.expenserules.employeerules.query`](https://opendocs.alipay.com/apis/05j4ci) | 查询员工绑定的所有费控规则 | `GET` | `/v3/alipay/ebpp/invoice/expenserules/employeerules/query` |
| [`alipay.ebpp.invoice.expenserules.groupemployee.modify`](https://opendocs.alipay.com/apis/05j82o) | 修改费控规则员工 | `POST` | `/v3/alipay/ebpp/invoice/expenserules/groupemployee/modify` |
| [`alipay.ebpp.invoice.expenserules.projectemployee.modify`](https://opendocs.alipay.com/apis/05wvx5) | 修改项目的员工列表 | `POST` | `/v3/alipay/ebpp/invoice/expenserules/projectemployee/modify` |
| [`alipay.ebpp.invoice.expenserules.projectinfo.modify`](https://opendocs.alipay.com/apis/05wuq6) | 修改项目基础信息 | `POST` | `/v3/alipay/ebpp/invoice/expenserules/projectinfo/modify` |
| [`alipay.ebpp.invoice.expenserules.projectrule.create`](https://opendocs.alipay.com/apis/05isio) | 创建项目费控规则 | `POST` | `/v3/alipay/ebpp/invoice/expenserules/projectrule/create` |
| [`alipay.ebpp.invoice.expenserules.projectrule.query`](https://opendocs.alipay.com/apis/05ip0x) | 查询项目费控规则 | `GET` | `/v3/alipay/ebpp/invoice/expenserules/projectrule/query` |
| [`alipay.ebpp.invoice.expenserules.projectrules.modify`](https://opendocs.alipay.com/apis/05xnct) | 修改项目费控规则 | `POST` | `/v3/alipay/ebpp/invoice/expenserules/projectrules/modify` |
| [`alipay.ebpp.invoice.expenserules.scenerule.create`](https://opendocs.alipay.com/apis/05j82p) | 创建费控规则 | `POST` | `/v3/alipay/ebpp/invoice/expenserules/scenerule/create` |
| [`alipay.ebpp.invoice.expenserules.scenerule.modify`](https://opendocs.alipay.com/apis/05xk1w) | 修改费控规则 | `POST` | `/v3/alipay/ebpp/invoice/expenserules/scenerule/modify` |
| [`alipay.ebpp.invoice.expenserules.scenerule.query`](https://opendocs.alipay.com/apis/05x1fs) | 查询费控规则信息 | `GET` | `/v3/alipay/ebpp/invoice/expenserules/scenerule/query` |
| [`alipay.ebpp.invoice.expensescene.merchant.query`](https://opendocs.alipay.com/apis/05pe0a) | 查询企业可用商户信息 | `GET` | `/v3/alipay/ebpp/invoice/expensescene/merchant/query` |
| [`alipay.ebpp.invoice.file.output.query`](https://opendocs.alipay.com/apis/05l1fc) | 获取电子发票原文件 | `GET` | `/v3/alipay/ebpp/invoice/file/output/query` |
| [`alipay.ebpp.invoice.info.send`](https://opendocs.alipay.com/apis/05l6nx) | 发票信息回传接口（新版） | `POST` | `/v3/alipay/ebpp/invoice/info/send` |
| [`alipay.ebpp.invoice.institution.create`](https://opendocs.alipay.com/apis/05wvvq) | 制度新增 | `POST` | `/v3/alipay/ebpp/invoice/institution/create` |
| [`alipay.ebpp.invoice.institution.delete`](https://opendocs.alipay.com/apis/05wunr) | 制度删除 | `DELETE` | `/v3/alipay/ebpp/invoice/institution/delete` |
| [`alipay.ebpp.invoice.institution.detailinfo.query`](https://opendocs.alipay.com/apis/05j9jg) | 制度详情查询 | `GET` | `/v3/alipay/ebpp/invoice/institution/detailinfo/query` |
| [`alipay.ebpp.invoice.institution.expenserule.create`](https://opendocs.alipay.com/apis/05waoz) | 创建费控使用规则 | `POST` | `/v3/alipay/ebpp/invoice/institution/expenserule/create` |
| [`alipay.ebpp.invoice.institution.expenserule.delete`](https://opendocs.alipay.com/apis/05xa73) | 删除使用规则 | `POST` | `/v3/alipay/ebpp/invoice/institution/expenserule/delete` |
| [`alipay.ebpp.invoice.institution.expenserule.modify`](https://opendocs.alipay.com/apis/05wuq0) | 编辑使用规则 | `POST` | `/v3/alipay/ebpp/invoice/institution/expenserule/modify` |
| [`alipay.ebpp.invoice.institution.modify`](https://opendocs.alipay.com/apis/05wvqu) | 制度编辑 | `POST` | `/v3/alipay/ebpp/invoice/institution/modify` |
| [`alipay.ebpp.invoice.institution.pageinfo.query`](https://opendocs.alipay.com/apis/05xo4v) | 费控制度分页查询 | `GET` | `/v3/alipay/ebpp/invoice/institution/pageinfo/query` |
| [`alipay.ebpp.invoice.institution.scope.modify`](https://opendocs.alipay.com/apis/05ip0w) | 修改制度下适配范围 | `POST` | `/v3/alipay/ebpp/invoice/institution/scope/modify` |
| [`alipay.ebpp.invoice.institution.scopepageinfo.query`](https://opendocs.alipay.com/apis/05j1c2) | 分页查询制度下人员 | `GET` | `/v3/alipay/ebpp/invoice/institution/scopepageinfo/query` |
| [`alipay.ebpp.invoice.issuerule.create`](https://opendocs.alipay.com/apis/05xoz2) | 新增发放规则 | `POST` | `/v3/alipay/ebpp/invoice/issuerule/create` |
| [`alipay.ebpp.invoice.issuerule.delete`](https://opendocs.alipay.com/apis/05wol0) | 删除发放规则 | `POST` | `/v3/alipay/ebpp/invoice/issuerule/delete` |
| [`alipay.ebpp.invoice.issuerule.modify`](https://opendocs.alipay.com/apis/05wyj4) | 编辑发放规则 | `POST` | `/v3/alipay/ebpp/invoice/issuerule/modify` |
| [`alipay.ebpp.invoice.isvtoken.reim.apply`](https://opendocs.alipay.com/apis/05xa7p) | 获取报销授权令牌 | `POST` | `/v3/alipay/ebpp/invoice/isvtoken/reim/apply` |
| [`alipay.ebpp.invoice.list.expense.sync`](https://opendocs.alipay.com/apis/05jve4) | 报销状态批量同步开放接口 | `POST` | `/v3/alipay/ebpp/invoice/list/expense/sync` |
| [`alipay.ebpp.invoice.merchant.enterstatus.query`](https://opendocs.alipay.com/apis/06x4dr) | 商户入驻流程状态查询 | `GET` | `/v3/alipay/ebpp/invoice/merchant/enterstatus/query` |
| [`alipay.ebpp.invoice.merchantlist.enter.apply`](https://opendocs.alipay.com/apis/05vxtl) | 商户批量入驻申请接口 | `POST` | `/v3/alipay/ebpp/invoice/merchantlist/enter/apply` |
| [`alipay.ebpp.invoice.order.query`](https://opendocs.alipay.com/pre-apis/05v5uk) | 根据外部订单号查询发票信息 | `GET` | `/v3/alipay/ebpp/invoice/order/query` |
| [`alipay.ebpp.invoice.sycn`](https://opendocs.alipay.com/apis/05la3z) | 发票信息同步至支付宝 | `POST` | `/v3/alipay/ebpp/invoice/sycn` |
| [`alipay.ebpp.invoice.sync.simple.send`](https://opendocs.alipay.com/apis/05kxx8) | 简单模式发票回传接口 | `POST` | `/v3/alipay/ebpp/invoice/sync/simple/send` |
| [`alipay.ebpp.invoice.taxno.batchquery`](https://opendocs.alipay.com/apis/05jrk9) | 获取指定企业的发票要素列表 | `POST` | `/v3/alipay/ebpp/invoice/taxno/batchquery` |
| [`alipay.ebpp.invoice.title.dynamic.get`](https://opendocs.alipay.com/apis/05v8k3) | 根据动态码查询发票抬头 | `GET` | `/v3/alipay/ebpp/invoice/title/dynamic/get` |
| [`alipay.ebpp.invoice.title.list.get`](https://opendocs.alipay.com/apis/06864n) | 获取指定用户所有的有效抬头列表 | `GET` | `/v3/alipay/ebpp/invoice/title/list/get` |
| [`alipay.ebpp.invoice.token.batchquery`](https://opendocs.alipay.com/apis/05khyy) | 根据查询令牌获取发票要素列表 | `GET` | `/v3/alipay/ebpp/invoice/token/batchquery` |
| [`alipay.ebpp.pdeduct.async.pay`](https://opendocs.alipay.com/pre-apis/05xmda) | 公共事业缴费直连代扣异步扣款支付接口 | `POST` | `/v3/alipay/ebpp/pdeduct/async/pay` |
| [`alipay.ebpp.pdeduct.bill.pay.status`](https://opendocs.alipay.com/open-v3/c7416677_alipay.ebpp.pdeduct.bill.pay.status) | 缴费直连代扣订单支付状态查询 | `GET` | `/v3/alipay/ebpp/pdeduct/bill/pay/status` |
| [`alipay.ebpp.pdeduct.pay`](https://opendocs.alipay.com/open-v3/fea30a1b_alipay.ebpp.pdeduct.pay) | 公共事业缴费直连代扣扣款支付接口 | `POST` | `/v3/alipay/ebpp/pdeduct/pay` |
| [`alipay.ebpp.pdeduct.sign.add`](https://opendocs.alipay.com/open-v3/084db95a_alipay.ebpp.pdeduct.sign.add) | 缴费直连代扣签约 | `POST` | `/v3/alipay/ebpp/pdeduct/sign/add` |
| [`alipay.ebpp.pdeduct.sign.cancel`](https://opendocs.alipay.com/open-v3/2ad13974_alipay.ebpp.pdeduct.sign.cancel) | 缴费直连代扣取消签约 | `POST` | `/v3/alipay/ebpp/pdeduct/sign/cancel` |
| [`alipay.ebpp.pdeduct.sign.query`](https://opendocs.alipay.com/apis/05x9ew) | 直连代扣协议查询接口 | `GET` | `/v3/alipay/ebpp/pdeduct/sign/query` |
| [`alipay.ebpp.pdeduct.sign.validate`](https://opendocs.alipay.com/apis/05wukm) | 缴费直连代扣签约前置校验 | `POST` | `/v3/alipay/ebpp/pdeduct/sign/validate` |

### `alipay.eco`（32）

| 接口 | 名称 | 方法 | 路径 |
| --- | --- | --- | --- |
| [`alipay.eco.contract.signflows.create`](https://opendocs.alipay.com/pre-apis/0681tj) | 创建电子合同签署流程(E签宝) | `POST` | `/v3/alipay/eco/contract/signflows/create` |
| [`alipay.eco.doc.template.create`](https://opendocs.alipay.com/pre-apis/05whak) | 创建合同模板（E签宝） | `POST` | `/v3/alipay/eco/doc/template/create` |
| [`alipay.eco.doctemplate.settingurl.query`](https://opendocs.alipay.com/pre-apis/05wf9z) | 获取合同模板设置地址（E签宝） | `GET` | `/v3/alipay/eco/doctemplate/settingurl/query` |
| [`alipay.eco.edu.kt.billing.modify`](https://opendocs.alipay.com/open-v3/1406f415_alipay.eco.edu.kt.billing.modify) | 教育缴费账单状态同步接口 | `POST` | `/v3/alipay/eco/edu/kt/billing/modify` |
| [`alipay.eco.edu.kt.billing.query`](https://opendocs.alipay.com/open-v3/d9852f73_alipay.eco.edu.kt.billing.query) | 缴费账单查询 | `GET` | `/v3/alipay/eco/edu/kt/billing/query` |
| [`alipay.eco.edu.kt.billing.send`](https://opendocs.alipay.com/open-v3/7c14b5a1_alipay.eco.edu.kt.billing.send) | 教育缴费账单发送接口 | `POST` | `/v3/alipay/eco/edu/kt/billing/send` |
| [`alipay.eco.edu.kt.parent.query`](https://opendocs.alipay.com/apis/05x0kr) | 查询学生家长状态接口 | `GET` | `/v3/alipay/eco/edu/kt/parent/query` |
| [`alipay.eco.edu.kt.schoolinfo.modify`](https://opendocs.alipay.com/open-v3/a319c0c9_alipay.eco.edu.kt.schoolinfo.modify) | 教育缴费学校信息录入接口 | `POST` | `/v3/alipay/eco/edu/kt/schoolinfo/modify` |
| [`alipay.eco.edu.kt.student.modify`](https://opendocs.alipay.com/apis/05wun3) | 学生信息更新 | `POST` | `/v3/alipay/eco/edu/kt/student/modify` |
| [`alipay.eco.file.path.query`](https://opendocs.alipay.com/pre-apis/05wghs) | 获取文件直传地址（E签宝） | `GET` | `/v3/alipay/eco/file/path/query` |
| [`alipay.eco.mycar.parking.chargeinfo.sync`](https://opendocs.alipay.com/open-v3/18d2ca24_alipay.eco.mycar.parking.chargeinfo.sync) | 停车场价格信息同步 | `POST` | `/v3/alipay/eco/mycar/parking/chargeinfo/sync` |
| [`alipay.eco.mycar.parking.charginginfo.sync`](https://opendocs.alipay.com/open-v3/91cababc_alipay.eco.mycar.parking.charginginfo.sync) | 车辆停车计费接口 | `POST` | `/v3/alipay/eco/mycar/parking/charginginfo/sync` |
| [`alipay.eco.mycar.parking.config.query`](https://opendocs.alipay.com/open-v3/735dfdcb_alipay.eco.mycar.parking.config.query) | ISV系统配置查询接口 | `GET` | `/v3/alipay/eco/mycar/parking/config/query` |
| [`alipay.eco.mycar.parking.config.set`](https://opendocs.alipay.com/open-v3/17f00a51_alipay.eco.mycar.parking.config.set) | 停车ISV系统配置接口 | `POST` | `/v3/alipay/eco/mycar/parking/config/set` |
| [`alipay.eco.mycar.parking.enterinfo.sync`](https://opendocs.alipay.com/open-v3/eb761e07_alipay.eco.mycar.parking.enterinfo.sync) | 车辆驶入接口 | `POST` | `/v3/alipay/eco/mycar/parking/enterinfo/sync` |
| [`alipay.eco.mycar.parking.exitinfo.sync`](https://opendocs.alipay.com/open-v3/4b64c632_alipay.eco.mycar.parking.exitinfo.sync) | 车辆驶出接口 | `POST` | `/v3/alipay/eco/mycar/parking/exitinfo/sync` |
| [`alipay.eco.mycar.parking.order.sync`](https://opendocs.alipay.com/open-v3/e22974fc_alipay.eco.mycar.parking.order.sync) | 订单同步接口 | `POST` | `/v3/alipay/eco/mycar/parking/order/sync` |
| [`alipay.eco.mycar.parking.order.update`](https://opendocs.alipay.com/open-v3/fe91ddf6_alipay.eco.mycar.parking.order.update) | 订单更新接口 | `POST` | `/v3/alipay/eco/mycar/parking/order/update` |
| [`alipay.eco.mycar.parking.overtimecharginginfo.sync`](https://opendocs.alipay.com/open-v3/3e99daa8_alipay.eco.mycar.parking.overtimecharginginfo.sync) | 车辆停车超时重计费接口 | `POST` | `/v3/alipay/eco/mycar/parking/overtimecharginginfo/sync` |
| [`alipay.eco.mycar.parking.parkinglotbiz.transfer`](https://opendocs.alipay.com/open-v3/cb924074_alipay.eco.mycar.parking.parkinglotbiz.transfer) | 车场业务归属配置 | `POST` | `/v3/alipay/eco/mycar/parking/parkinglotbiz/transfer` |
| [`alipay.eco.mycar.parking.parkinglotinfo.create`](https://opendocs.alipay.com/open-v3/fc05d317_alipay.eco.mycar.parking.parkinglotinfo.create) | 录入停车场信息 | `POST` | `/v3/alipay/eco/mycar/parking/parkinglotinfo/create` |
| [`alipay.eco.mycar.parking.parkinglotinfo.query`](https://opendocs.alipay.com/open-v3/4772d6e6_alipay.eco.mycar.parking.parkinglotinfo.query) | 停车场信息查询 | `GET` | `/v3/alipay/eco/mycar/parking/parkinglotinfo/query` |
| [`alipay.eco.mycar.parking.parkinglotinfo.update`](https://opendocs.alipay.com/open-v3/223553d2_alipay.eco.mycar.parking.parkinglotinfo.update) | 修改停车场信息 | `POST` | `/v3/alipay/eco/mycar/parking/parkinglotinfo/update` |
| [`alipay.eco.mycar.parking.paymentinfo.sync`](https://opendocs.alipay.com/open-v3/def30621_alipay.eco.mycar.parking.paymentinfo.sync) | 车辆停车缴费接口 | `POST` | `/v3/alipay/eco/mycar/parking/paymentinfo/sync` |
| [`alipay.eco.mycar.parking.spaceinfo.sync`](https://opendocs.alipay.com/open-v3/cab695b7_alipay.eco.mycar.parking.spaceinfo.sync) | 停车场车位信息同步 | `POST` | `/v3/alipay/eco/mycar/parking/spaceinfo/sync` |
| [`alipay.eco.mycar.parking.vehicle.query`](https://opendocs.alipay.com/open-v3/c04327bf_alipay.eco.mycar.parking.vehicle.query) | 车牌查询接口 | `GET` | `/v3/alipay/eco/mycar/parking/vehicle/query` |
| [`alipay.eco.sign.flow.cancel`](https://opendocs.alipay.com/pre-apis/05wvs6) | 签署流程撤销（E签宝） | `POST` | `/v3/alipay/eco/sign/flow/cancel` |
| [`alipay.eco.sign.flow.create`](https://opendocs.alipay.com/pre-apis/05xsf2) | 创建流程（E签宝） | `POST` | `/v3/alipay/eco/sign/flow/create` |
| [`alipay.eco.sign.flow.finish`](https://opendocs.alipay.com/pre-apis/05wuhl) | 签署流程归档 | `POST` | `/v3/alipay/eco/sign/flow/finish` |
| [`alipay.eco.sign.flow.query`](https://opendocs.alipay.com/pre-apis/05wvxl) | 签署流程查询（E签宝） | `GET` | `/v3/alipay/eco/sign/flow/query` |
| [`alipay.eco.signflows.detail.query`](https://opendocs.alipay.com/pre-apis/05wylm) | 获取签署流程合同与附件的下载地址（E签宝） | `GET` | `/v3/alipay/eco/signflows/detail/query` |
| [`alipay.eco.signflows.url.query`](https://opendocs.alipay.com/pre-apis/05wyc8) | 获取签署地址 | `GET` | `/v3/alipay/eco/signflows/url/query` |

### `alipay.fund`（42）

| 接口 | 名称 | 方法 | 路径 |
| --- | --- | --- | --- |
| [`alipay.fund.account.query`](https://opendocs.alipay.com/open-v3/77e2b925_alipay.fund.account.query) | 支付宝资金账户资产查询接口 | `GET` | `/v3/alipay/fund/account/query` |
| [`alipay.fund.accountbook.create`](https://opendocs.alipay.com/apis/0cjl4u) | 资金记账本开通 | `POST` | `/v3/alipay/fund/accountbook/create` |
| [`alipay.fund.accountbook.notify.query`](https://opendocs.alipay.com/open-v3/495b0416_alipay.fund.accountbook.notify.query) | 记账本通知订阅关系查询接口 | `GET` | `/v3/alipay/fund/accountbook/notify/query` |
| [`alipay.fund.accountbook.notify.subscribe`](https://opendocs.alipay.com/open-v3/37df5c87_alipay.fund.accountbook.notify.subscribe) | 记账本通知订阅接口 | `POST` | `/v3/alipay/fund/accountbook/notify/subscribe` |
| [`alipay.fund.accountbook.notify.unsubscribe`](https://opendocs.alipay.com/open-v3/1923216a_alipay.fund.accountbook.notify.unsubscribe) | 记账本通知取消订阅接口 | `POST` | `/v3/alipay/fund/accountbook/notify/unsubscribe` |
| [`alipay.fund.accountbook.query`](https://opendocs.alipay.com/open-v3/f072b3ba_alipay.fund.accountbook.query) | 资金记账本的信息查询 | `GET` | `/v3/alipay/fund/accountbook/query` |
| [`alipay.fund.agreement.quota.modify`](https://opendocs.alipay.com/apis/05xo4q) | 额度设置接口 | `POST` | `/v3/alipay/fund/agreement/quota/modify` |
| [`alipay.fund.agreement.quota.query`](https://opendocs.alipay.com/apis/05ximq) | 额度查询接口 | `POST` | `/v3/alipay/fund/agreement/quota/query` |
| [`alipay.fund.auth.operation.cancel`](https://opendocs.alipay.com/open-v3/2fecf986_alipay.fund.auth.operation.cancel) | 资金授权撤销接口 | `POST` | `/v3/alipay/fund/auth/operation/cancel` |
| [`alipay.fund.auth.operation.detail.query`](https://opendocs.alipay.com/open-v3/984254a0_alipay.fund.auth.operation.detail.query) | 资金授权操作查询接口 | `POST` | `/v3/alipay/fund/auth/operation/detail/query` |
| [`alipay.fund.auth.order.app.freeze`](https://opendocs.alipay.com/open-v3/46bf87fe_alipay.fund.auth.order.app.freeze) | 线上资金授权冻结接口（文档演示 App 唤端，右列 REST 路由来自快照） | `POST` | `/v3/alipay/fund/auth/order/app/freeze` |
| [`alipay.fund.auth.order.freeze`](https://opendocs.alipay.com/open-v3/1ffb104d_alipay.fund.auth.order.freeze) | 资金授权冻结接口 | `POST` | `/v3/alipay/fund/auth/order/freeze` |
| [`alipay.fund.auth.order.unfreeze`](https://opendocs.alipay.com/open-v3/cb3b0ae2_alipay.fund.auth.order.unfreeze) | 资金授权解冻接口 | `POST` | `/v3/alipay/fund/auth/order/unfreeze` |
| [`alipay.fund.auth.order.voucher.create`](https://opendocs.alipay.com/open-v3/6f5ccbfe_alipay.fund.auth.order.voucher.create) | 资金授权发码接口 | `POST` | `/v3/alipay/fund/auth/order/voucher/create` |
| [`alipay.fund.enterprisepay.group.add`](https://opendocs.alipay.com/apis/05wmyu) | 因公付新增账户下群组 | `POST` | `/v3/alipay/fund/enterprisepay/group/add` |
| [`alipay.fund.enterprisepay.group.delete`](https://opendocs.alipay.com/apis/05wvyb) | 因公付删除账户下群组 | `DELETE` | `/v3/alipay/fund/enterprisepay/group/delete` |
| [`alipay.fund.enterprisepay.group.modify`](https://opendocs.alipay.com/apis/05wmzb) | 因公付更新账户下群组 | `POST` | `/v3/alipay/fund/enterprisepay/group/modify` |
| [`alipay.fund.enterprisepay.group.query`](https://opendocs.alipay.com/apis/05wcpq) | 因公付查询账户下群组 | `GET` | `/v3/alipay/fund/enterprisepay/group/query` |
| [`alipay.fund.enterprisepay.member.modify`](https://opendocs.alipay.com/apis/05wvv1) | 因公付更新员工资金协议 | `POST` | `/v3/alipay/fund/enterprisepay/member/modify` |
| [`alipay.fund.enterprisepay.sign`](https://opendocs.alipay.com/apis/05udr5) | 因公付企业签约 | `POST` | `/v3/alipay/fund/enterprisepay/sign` |
| [`alipay.fund.enterprisepay.unsign`](https://opendocs.alipay.com/apis/05wbw2) | 因公付企业解约 | `POST` | `/v3/alipay/fund/enterprisepay/unsign` |
| [`alipay.fund.jointaccount.bill.query`](https://opendocs.alipay.com/apis/05wuhp) | 因公付账单查询 | `GET` | `/v3/alipay/fund/jointaccount/bill/query` |
| [`alipay.fund.jointaccount.detail.query`](https://opendocs.alipay.com/apis/05imoi) | 共享资金-查询账户详情 | `GET` | `/v3/alipay/fund/jointaccount/detail/query` |
| [`alipay.fund.jointaccount.list.query`](https://opendocs.alipay.com/apis/05wcqw) | 企业查询代付账户列表 | `GET` | `/v3/alipay/fund/jointaccount/list/query` |
| [`alipay.fund.jointaccount.member.batchquery`](https://opendocs.alipay.com/apis/05wfae) | 企业批量查询员工信息（分页） | `POST` | `/v3/alipay/fund/jointaccount/member/batchquery` |
| [`alipay.fund.jointaccount.member.bind`](https://opendocs.alipay.com/apis/05wvt3) | 因公付邀请员工 | `POST` | `/v3/alipay/fund/jointaccount/member/bind` |
| [`alipay.fund.jointaccount.member.consult`](https://opendocs.alipay.com/apis/05wupc) | 员企关系咨询 | `POST` | `/v3/alipay/fund/jointaccount/member/consult` |
| [`alipay.fund.jointaccount.member.query`](https://opendocs.alipay.com/apis/05wujm) | 企业查询员工列表 | `GET` | `/v3/alipay/fund/jointaccount/member/query` |
| [`alipay.fund.jointaccount.member.unbind`](https://opendocs.alipay.com/apis/05x1fk) | 企业删除员工 | `POST` | `/v3/alipay/fund/jointaccount/member/unbind` |
| [`alipay.fund.jointaccount.memberrule.modify`](https://opendocs.alipay.com/apis/05wygg) | 企业设置员工规则 | `POST` | `/v3/alipay/fund/jointaccount/memberrule/modify` |
| [`alipay.fund.jointaccount.quota.query`](https://opendocs.alipay.com/apis/05wuhk) | 查询账户额度详情 | `GET` | `/v3/alipay/fund/jointaccount/quota/query` |
| [`alipay.fund.jointaccount.rule.modify`](https://opendocs.alipay.com/apis/05wghc) | 企业账户规则设置 | `POST` | `/v3/alipay/fund/jointaccount/rule/modify` |
| [`alipay.fund.jointaccount.sign.query`](https://opendocs.alipay.com/apis/05wulj) | 企业签约结果查询 | `GET` | `/v3/alipay/fund/jointaccount/sign/query` |
| [`alipay.fund.jointaccount.trade.query`](https://opendocs.alipay.com/apis/05wykp) | 共同账户交易查询 | `GET` | `/v3/alipay/fund/jointaccount/trade/query` |
| [`alipay.fund.trans.common.query`](https://opendocs.alipay.com/open-v3/bcbe6351_alipay.fund.trans.common.query?scene=common&pathHash=36f664d9) | 转账业务单据查询接口 | `GET` | `/v3/alipay/fund/trans/common/query` |
| [`alipay.fund.trans.order.query`](https://opendocs.alipay.com/open-v3/f00381dd_alipay.fund.trans.order.query) | 查询转账订单接口 | `GET` | `/v3/alipay/fund/trans/order/query` |
| [`alipay.fund.trans.pay`](https://opendocs.alipay.com/pre-apis/05xggg) | 资金支付接口 | `POST` | `/v3/alipay/fund/trans/pay` |
| [`alipay.fund.trans.payee.bind.query`](https://opendocs.alipay.com/open-v3/03fe4a0d_alipay.fund.trans.payee.bind.query?scene=common&pathHash=377d1f29) | 资金收款账号绑定关系查询 | `GET` | `/v3/alipay/fund/trans/payee/bind/query` |
| [`alipay.fund.trans.refund`](https://opendocs.alipay.com/open-v3/0c6ff10a_alipay.fund.trans.refund) | 资金退回接口 | `POST` | `/v3/alipay/fund/trans/refund` |
| [`alipay.fund.trans.toaccount.transfer`](https://opendocs.alipay.com/open-v3/53a4ce8e_alipay.fund.trans.toaccount.transfer) | 单笔转账到支付宝账户接口 | `POST` | `/v3/alipay/fund/trans/toaccount/transfer` |
| [`alipay.fund.trans.uni.transfer`](https://opendocs.alipay.com/open-v3/26a22ad7_alipay.fund.trans.uni.transfer?scene=common%2Cca56bca529e64125a2786703c6192d41%2C66dd06f5a923403393b85de68d3c0055&pathHash=9b52afe2) | 单笔转账接口 | `POST` | `/v3/alipay/fund/trans/uni/transfer` |
| [`alipay.fund.usertrade.batchquery`](https://opendocs.alipay.com/pre-apis/05xggy) | 用户交易数据查询 | `POST` | `/v3/alipay/fund/usertrade/batchquery` |

### `alipay.iservice`（45）

| 接口 | 名称 | 方法 | 路径 |
| --- | --- | --- | --- |
| [`alipay.iservice.ccm.agent.create`](https://opendocs.alipay.com/open-v3/d09eb15f_alipay.iservice.ccm.agent.create) | 新增客服接口 | `POST` | `/v3/alipay/iservice/ccm/agent/create` |
| [`alipay.iservice.ccm.agent.delete`](https://opendocs.alipay.com/open-v3/dc156fee_alipay.iservice.ccm.agent.delete) | 删除客服接口 | `DELETE` | `/v3/alipay/iservice/ccm/agent/delete` |
| [`alipay.iservice.ccm.agent.get`](https://opendocs.alipay.com/open-v3/4a14056c_alipay.iservice.ccm.agent.get) | 查询客服详情 | `GET` | `/v3/alipay/iservice/ccm/agent/get` |
| [`alipay.iservice.ccm.agent.modify`](https://opendocs.alipay.com/open-v3/e5964f19_alipay.iservice.ccm.agent.modify) | 更新客服接口 | `POST` | `/v3/alipay/iservice/ccm/agent/modify` |
| [`alipay.iservice.ccm.agent.query`](https://opendocs.alipay.com/open-v3/992aed93_alipay.iservice.ccm.agent.query) | 客服查询接口 | `POST` | `/v3/alipay/iservice/ccm/agent/query` |
| [`alipay.iservice.ccm.agent.schedulelog.query`](https://opendocs.alipay.com/open-v3/be76c956_alipay.iservice.ccm.agent.schedulelog.query) | 查询客服状态变更流水日志 | `POST` | `/v3/alipay/iservice/ccm/agent/schedulelog/query` |
| [`alipay.iservice.ccm.function.query`](https://opendocs.alipay.com/open-v3/45b96512_alipay.iservice.ccm.function.query) | 列出所有的功能点 | `GET` | `/v3/alipay/iservice/ccm/function/query` |
| [`alipay.iservice.ccm.instance.create`](https://opendocs.alipay.com/open-v3/d2fd8b94_alipay.iservice.ccm.instance.create) | 创建租户实例 | `POST` | `/v3/alipay/iservice/ccm/instance/create` |
| [`alipay.iservice.ccm.instance.get`](https://opendocs.alipay.com/open-v3/e02604d3_alipay.iservice.ccm.instance.get) | 查询单个租户实例（数据权限） | `GET` | `/v3/alipay/iservice/ccm/instance/get` |
| [`alipay.iservice.ccm.instance.query`](https://opendocs.alipay.com/open-v3/eb6f9aca_alipay.iservice.ccm.instance.query) | 列出所有的租户实例（数据权限） | `GET` | `/v3/alipay/iservice/ccm/instance/query` |
| [`alipay.iservice.ccm.isv.initialize`](https://opendocs.alipay.com/open-v3/eb2e1712_alipay.iservice.ccm.isv.initialize?scene=common&pathHash=da5a3102) | ccm isv注册接口 | `POST` | `/v3/alipay/iservice/ccm/isv/initialize` |
| [`alipay.iservice.ccm.ols.chatrecord.query`](https://opendocs.alipay.com/apis/05wokm) | 查询某通在线服务的聊天记录 | `GET` | `/v3/alipay/iservice/ccm/ols/chatrecord/query` |
| [`alipay.iservice.ccm.robot.avatarbase.query`](https://opendocs.alipay.com/open-v3/6d82b0cb_alipay.iservice.ccm.robot.avatarbase.query) | 虚拟数字人开放接口调用 | `POST` | `/v3/alipay/iservice/ccm/robot/avatarbase/query` |
| [`alipay.iservice.ccm.role.create`](https://opendocs.alipay.com/open-v3/632638c3_alipay.iservice.ccm.role.create) | 创建角色 | `POST` | `/v3/alipay/iservice/ccm/role/create` |
| [`alipay.iservice.ccm.role.delete`](https://opendocs.alipay.com/open-v3/8a5f6702_alipay.iservice.ccm.role.delete) | 删除角色 | `DELETE` | `/v3/alipay/iservice/ccm/role/delete` |
| [`alipay.iservice.ccm.role.get`](https://opendocs.alipay.com/apis/05wvxp) | 获取角色详情 | `GET` | `/v3/alipay/iservice/ccm/role/get` |
| [`alipay.iservice.ccm.role.modify`](https://opendocs.alipay.com/open-v3/8e38792c_alipay.iservice.ccm.role.modify) | 更新角色 | `POST` | `/v3/alipay/iservice/ccm/role/modify` |
| [`alipay.iservice.ccm.role.page.query`](https://opendocs.alipay.com/open-v3/0c530d84_alipay.iservice.ccm.role.page.query) | 分页查询角色 | `GET` | `/v3/alipay/iservice/ccm/role/page/query` |
| [`alipay.iservice.ccm.service.buy`](https://opendocs.alipay.com/open-v3/705fa9ad_alipay.iservice.ccm.service.buy?scene=common&pathHash=42f013fd) | 租户开通isv服务接口 | `POST` | `/v3/alipay/iservice/ccm/service/buy` |
| [`alipay.iservice.ccm.service.close`](https://opendocs.alipay.com/open-v3/2202303e_alipay.iservice.ccm.service.close?scene=common&pathHash=4e8520fc) | 关闭租户开通的isv服务接口 | `POST` | `/v3/alipay/iservice/ccm/service/close` |
| [`alipay.iservice.ccm.service.initialize`](https://opendocs.alipay.com/open-v3/ab388310_alipay.iservice.ccm.service.initialize?scene=common&pathHash=3dce471e) | isv服务初始化接口 | `POST` | `/v3/alipay/iservice/ccm/service/initialize` |
| [`alipay.iservice.ccm.servicerecord.create`](https://opendocs.alipay.com/open-v3/44b1d498_alipay.iservice.ccm.servicerecord.create) | 创建服务记录 | `POST` | `/v3/alipay/iservice/ccm/servicerecord/create` |
| [`alipay.iservice.ccm.servicerecord.modify`](https://opendocs.alipay.com/open-v3/fefe64f8_alipay.iservice.ccm.servicerecord.modify) | 变更服务记录 | `POST` | `/v3/alipay/iservice/ccm/servicerecord/modify` |
| [`alipay.iservice.ccm.sw.article.batchquery`](https://opendocs.alipay.com/open-v3/031db27a_alipay.iservice.ccm.sw.article.batchquery) | 知识库-问答-批量查询 | `POST` | `/v3/alipay/iservice/ccm/sw/article/batchquery` |
| [`alipay.iservice.ccm.sw.article.create`](https://opendocs.alipay.com/open-v3/7c58db8d_alipay.iservice.ccm.sw.article.create) | 知识库-问答-创建 | `PUT` | `/v3/alipay/iservice/ccm/sw/article/create` |
| [`alipay.iservice.ccm.sw.article.delete`](https://opendocs.alipay.com/open-v3/d2475255_alipay.iservice.ccm.sw.article.delete) | 知识库-问答-删除 | `DELETE` | `/v3/alipay/iservice/ccm/sw/article/delete` |
| [`alipay.iservice.ccm.sw.article.get`](https://opendocs.alipay.com/open-v3/c2addca8_alipay.iservice.ccm.sw.article.get) | 知识库-问答-获取 | `GET` | `/v3/alipay/iservice/ccm/sw/article/get` |
| [`alipay.iservice.ccm.sw.article.modify`](https://opendocs.alipay.com/open-v3/cc0b0135_alipay.iservice.ccm.sw.article.modify) | 知识库-问答-修改 | `PATCH` | `/v3/alipay/iservice/ccm/sw/article/modify` |
| [`alipay.iservice.ccm.sw.knowledge.modify`](https://opendocs.alipay.com/open-v3/c77df3ee_alipay.iservice.ccm.sw.knowledge.modify) | 知识库问答管理 | `POST` | `/v3/alipay/iservice/ccm/sw/knowledge/modify` |
| [`alipay.iservice.ccm.sw.library.batchquery`](https://opendocs.alipay.com/open-v3/04e0096f_alipay.iservice.ccm.sw.library.batchquery) | 知识库-库-批量查询 | `GET` | `/v3/alipay/iservice/ccm/sw/library/batchquery` |
| [`alipay.iservice.ccm.sw.library.create`](https://opendocs.alipay.com/open-v3/404aca95_alipay.iservice.ccm.sw.library.create) | 知识库-库-创建 | `PUT` | `/v3/alipay/iservice/ccm/sw/library/create` |
| [`alipay.iservice.ccm.sw.library.delete`](https://opendocs.alipay.com/open-v3/cd16a3aa_alipay.iservice.ccm.sw.library.delete) | 知识库-库-删除 | `DELETE` | `/v3/alipay/iservice/ccm/sw/library/delete` |
| [`alipay.iservice.ccm.sw.library.modify`](https://opendocs.alipay.com/open-v3/10893431_alipay.iservice.ccm.sw.library.modify) | 知识库-库-修改 | `PATCH` | `/v3/alipay/iservice/ccm/sw/library/modify` |
| [`alipay.iservice.ccm.sw.order.sync`](https://opendocs.alipay.com/apis/0630z3) | 商品订单同步 | `POST` | `/v3/alipay/iservice/ccm/sw/order/sync` |
| [`alipay.iservice.ccm.sw.spu.sync`](https://opendocs.alipay.com/apis/063syw) | 商品库-商品信息同步 | `POST` | `/v3/alipay/iservice/ccm/sw/spu/sync` |
| [`alipay.iservice.ccm.sw.tree.batchquery`](https://opendocs.alipay.com/open-v3/7ce6cb6e_alipay.iservice.ccm.sw.tree.batchquery) | 知识库-类目-批量查询 | `GET` | `/v3/alipay/iservice/ccm/sw/tree/batchquery` |
| [`alipay.iservice.ccm.sw.tree.create`](https://opendocs.alipay.com/open-v3/c2bb833d_alipay.iservice.ccm.sw.tree.create) | 知识库-类目-创建 | `PUT` | `/v3/alipay/iservice/ccm/sw/tree/create` |
| [`alipay.iservice.ccm.sw.tree.delete`](https://opendocs.alipay.com/open-v3/10547684_alipay.iservice.ccm.sw.tree.delete) | 知识库-类目-删除 | `DELETE` | `/v3/alipay/iservice/ccm/sw/tree/delete` |
| [`alipay.iservice.ccm.sw.tree.get`](https://opendocs.alipay.com/open-v3/930c8404_alipay.iservice.ccm.sw.tree.get) | 知识库-类目-获取 | `GET` | `/v3/alipay/iservice/ccm/sw/tree/get` |
| [`alipay.iservice.ccm.sw.tree.modify`](https://opendocs.alipay.com/open-v3/b89ce966_alipay.iservice.ccm.sw.tree.modify) | 知识库-类目-修改 | `PATCH` | `/v3/alipay/iservice/ccm/sw/tree/modify` |
| [`alipay.iservice.ccm.sw.tree.publish`](https://opendocs.alipay.com/open-v3/3aa3bbe6_alipay.iservice.ccm.sw.tree.publish) | 知识库-类目-发布 | `PUT` | `/v3/alipay/iservice/ccm/sw/tree/publish` |
| [`alipay.iservice.ccm.sw.treecategory.batchquery`](https://opendocs.alipay.com/open-v3/24848e83_alipay.iservice.ccm.sw.treecategory.batchquery) | 知识库-节点-批量查询 | `GET` | `/v3/alipay/iservice/ccm/sw/treecategory/batchquery` |
| [`alipay.iservice.ccm.sw.treecategory.create`](https://opendocs.alipay.com/open-v3/08f8ea8c_alipay.iservice.ccm.sw.treecategory.create) | 知识库-节点-创建 | `PUT` | `/v3/alipay/iservice/ccm/sw/treecategory/create` |
| [`alipay.iservice.ccm.sw.treecategory.delete`](https://opendocs.alipay.com/open-v3/cf84306f_alipay.iservice.ccm.sw.treecategory.delete) | 知识库-节点-删除 | `DELETE` | `/v3/alipay/iservice/ccm/sw/treecategory/delete` |
| [`alipay.iservice.ccm.sw.treecategory.modify`](https://opendocs.alipay.com/open-v3/6fea3439_alipay.iservice.ccm.sw.treecategory.modify) | 知识库-节点-修改 | `PATCH` | `/v3/alipay/iservice/ccm/sw/treecategory/modify` |

### `alipay.marketing`（64）

| 接口 | 名称 | 方法 | 路径 |
| --- | --- | --- | --- |
| [`alipay.marketing.activity.app.batchquery`](https://opendocs.alipay.com/open-v3/c8f333a8_alipay.marketing.activity.app.batchquery) | 查询活动可用小程序 | `GET` | `/v3/alipay/marketing/activity/{activity_id}/app/batchquery` |
| [`alipay.marketing.activity.batchquery`](https://opendocs.alipay.com/open-v3/05355ed4_alipay.marketing.activity.batchquery) | 条件查询活动列表 | `POST` | `/v3/alipay/marketing/activity/batchquery` |
| [`alipay.marketing.activity.consult`](https://opendocs.alipay.com/open-v3/3c187ec5_alipay.marketing.activity.consult) | 活动领取咨询接口 | `POST` | `/v3/alipay/marketing/activity/consult` |
| [`alipay.marketing.activity.delivery.create`](https://opendocs.alipay.com/open-v3/43363a7a_alipay.marketing.activity.delivery.create) | 创建推广计划 | `POST` | `/v3/alipay/marketing/delivery` |
| [`alipay.marketing.activity.delivery.query`](https://opendocs.alipay.com/open-v3/cab7f6c3_alipay.marketing.activity.delivery.query) | 查询推广计划 | `POST` | `/v3/alipay/marketing/delivery/{delivery_id}/query` |
| [`alipay.marketing.activity.delivery.stop`](https://opendocs.alipay.com/open-v3/dbc77a2d_alipay.marketing.activity.delivery.stop) | 停止推广计划 | `PATCH` | `/v3/alipay/marketing/delivery/{delivery_id}/stop` |
| [`alipay.marketing.activity.deliverychannel.query`](https://opendocs.alipay.com/apis/05v7cm) | 查询可推广渠道 | `POST` | `/v3/alipay/marketing/activity/deliverychannel/query` |
| [`alipay.marketing.activity.goods.batchquery`](https://opendocs.alipay.com/open-v3/d196d1f1_alipay.marketing.activity.goods.batchquery) | 查询活动适用商品 | `GET` | `/v3/alipay/marketing/activity/{activity_id}/goods/batchquery` |
| [`alipay.marketing.activity.merchant.batchquery`](https://opendocs.alipay.com/open-v3/30e4ba35_alipay.marketing.activity.merchant.batchquery) | 查询活动可用商户 | `GET` | `/v3/alipay/marketing/activity/{activity_id}/merchant/batchquery` |
| [`alipay.marketing.activity.ordervoucher.append`](https://opendocs.alipay.com/apis/05vjnq) | 修改商家券活动发券数量上限 | `PATCH` | `/v3/alipay/marketing/ordervoucher/activity/{activity_id}/append` |
| [`alipay.marketing.activity.ordervoucher.associate`](https://opendocs.alipay.com/apis/05v8k7) | 关联订单信息 | `POST` | `/v3/alipay/marketing/activity/ordervoucher/associate` |
| [`alipay.marketing.activity.ordervoucher.codecount`](https://opendocs.alipay.com/open-v3/1b98a84b_alipay.marketing.activity.ordervoucher.codecount) | 统计商家券券码数量 | `GET` | `/v3/alipay/marketing/ordervoucher/activity/{activity_id}/voucher/codecount` |
| [`alipay.marketing.activity.ordervoucher.codedeposit`](https://opendocs.alipay.com/apis/05u431) | 同步商家券券码 | `POST` | `/v3/alipay/marketing/ordervoucher/activity/{activity_id}/voucher/codedeposit` |
| [`alipay.marketing.activity.ordervoucher.create`](https://opendocs.alipay.com/apis/05qgfh) | 创建商家券活动 | `POST` | `/v3/alipay/marketing/ordervoucher/activity` |
| [`alipay.marketing.activity.ordervoucher.disassociate`](https://opendocs.alipay.com/apis/05vluj) | 取消关联订单信息 | `POST` | `/v3/alipay/marketing/activity/ordervoucher/disassociate` |
| [`alipay.marketing.activity.ordervoucher.modify`](https://opendocs.alipay.com/apis/05vfpx) | 修改商家券活动基本信息 | `PATCH` | `/v3/alipay/marketing/ordervoucher/activity/{activity_id}` |
| [`alipay.marketing.activity.ordervoucher.query`](https://opendocs.alipay.com/open-v3/ab283f3f_alipay.marketing.activity.ordervoucher.query) | 查询商家券活动 | `GET` | `/v3/alipay/marketing/activity/ordervoucher/query` |
| [`alipay.marketing.activity.ordervoucher.refund`](https://opendocs.alipay.com/apis/0636fj) | 取消券核销状态 | `POST` | `/v3/alipay/marketing/ordervoucher/activity/{activity_id}/voucher/{voucher_code}/refund` |
| [`alipay.marketing.activity.ordervoucher.stop`](https://opendocs.alipay.com/apis/0630yz) | 停止商家券活动 | `PATCH` | `/v3/alipay/marketing/ordervoucher/activity/{activity_id}/stop` |
| [`alipay.marketing.activity.ordervoucher.use`](https://opendocs.alipay.com/apis/0639qt) | 同步券核销状态 | `POST` | `/v3/alipay/marketing/ordervoucher/activity/{activity_id}/voucher/{voucher_code}/use` |
| [`alipay.marketing.activity.query`](https://opendocs.alipay.com/open-v3/168d324a_alipay.marketing.activity.query) | 查询活动详情 | `GET` | `/v3/alipay/marketing/activity/{activity_id}` |
| [`alipay.marketing.activity.shop.batchquery`](https://opendocs.alipay.com/open-v3/ebd50e50_alipay.marketing.activity.shop.batchquery) | 查询活动可用门店 | `GET` | `/v3/alipay/marketing/activity/{activity_id}/shop/batchquery` |
| [`alipay.marketing.activity.user.batchqueryvoucher`](https://opendocs.alipay.com/open-v3/63264abf_alipay.marketing.activity.user.batchqueryvoucher) | 条件查询用户券 | `GET` | `/v3/alipay/marketing/activity/ordervoucher/user/batchqueryvoucher` |
| [`alipay.marketing.activity.user.queryvoucher`](https://opendocs.alipay.com/open-v3/e75be5aa_alipay.marketing.activity.user.queryvoucher) | 查询用户券详情 | `GET` | `/v3/alipay/marketing/activity/ordervoucher/user/voucher` |
| [`alipay.marketing.activity.voucher.append`](https://opendocs.alipay.com/apis/05vsc4) | 追加支付券预算 | `POST` | `/v3/alipay/marketing/activity/voucher/append` |
| [`alipay.marketing.activity.voucher.create`](https://opendocs.alipay.com/apis/05vb41) | 创建支付券 | `POST` | `/v3/alipay/marketing/activity/voucher/create` |
| [`alipay.marketing.activity.voucher.modify`](https://opendocs.alipay.com/apis/05v8k9) | 修改支付券基本信息 | `POST` | `/v3/alipay/marketing/activity/voucher/modify` |
| [`alipay.marketing.activity.voucher.publish`](https://opendocs.alipay.com/apis/05vsc3) | 激活支付券 | `POST` | `/v3/alipay/marketing/activity/voucher/publish` |
| [`alipay.marketing.activity.voucher.query`](https://opendocs.alipay.com/apis/05veld) | 查询支付券详情 | `GET` | `/v3/alipay/marketing/activity/voucher/query` |
| [`alipay.marketing.activity.voucher.send`](https://opendocs.alipay.com/apis/05vidi) | 发放支付券 | `POST` | `/v3/alipay/marketing/activity/voucher/send` |
| [`alipay.marketing.activity.voucher.stop`](https://opendocs.alipay.com/apis/05vc1k) | 停止支付券 | `POST` | `/v3/alipay/marketing/activity/voucher/stop` |
| [`alipay.marketing.activity.voucherpackage.batchquery`](https://opendocs.alipay.com/apis/05velf) | 券包批量查询 | `POST` | `/v3/alipay/marketing/activity/voucherpackage/batchquery` |
| [`alipay.marketing.activity.voucherpackage.consult`](https://opendocs.alipay.com/apis/05vc1l) | 券包购买咨询 | `POST` | `/v3/alipay/marketing/activity/voucherpackage/consult` |
| [`alipay.marketing.activity.voucherpackage.query`](https://opendocs.alipay.com/apis/05v9l4) | 券包详情查询 | `GET` | `/v3/alipay/marketing/activity/voucherpackage/query` |
| [`alipay.marketing.campaign.cash.create`](https://opendocs.alipay.com/open-v3/e572d9c1_alipay.marketing.campaign.cash.create?scene=common&pathHash=5197b580) | 创建现金活动 | `POST` | `/v3/alipay/marketing/campaign/cash/create` |
| [`alipay.marketing.campaign.cash.detail.query`](https://opendocs.alipay.com/open-v3/7fede5af_alipay.marketing.campaign.cash.detail.query?scene=common&pathHash=430db29a) | 现金活动详情查询 | `GET` | `/v3/alipay/marketing/campaign/cash/detail/query` |
| [`alipay.marketing.campaign.cash.list.query`](https://opendocs.alipay.com/open-v3/b0871b97_alipay.marketing.campaign.cash.list.query?scene=common&pathHash=e06176c5) | 现金活动列表查询 | `GET` | `/v3/alipay/marketing/campaign/cash/list/query` |
| [`alipay.marketing.campaign.cash.status.modify`](https://opendocs.alipay.com/open-v3/4f41b60d_alipay.marketing.campaign.cash.status.modify?scene=common&pathHash=7b368c26) | 更改现金活动状态 | `POST` | `/v3/alipay/marketing/campaign/cash/status/modify` |
| [`alipay.marketing.campaign.cash.trigger`](https://opendocs.alipay.com/open-v3/06823eee_alipay.marketing.campaign.cash.trigger?scene=common&pathHash=e770d208) | 触发现金红包活动 | `POST` | `/v3/alipay/marketing/campaign/cash/trigger` |
| [`alipay.marketing.campaign.order.voucher.consult`](https://opendocs.alipay.com/open-v3/bb374dd7_alipay.marketing.campaign.order.voucher.consult) | 订单优惠前置咨询 | `POST` | `/v3/alipay/marketing/campaign/order/voucher/consult` |
| [`alipay.marketing.card.activateform.query`](https://opendocs.alipay.com/open-v3/a195d7c2_alipay.marketing.card.activateform.query) | 查询用户提交的会员卡表单信息 | `GET` | `/v3/alipay/marketing/card/activateform/query` |
| [`alipay.marketing.card.activateurl.apply`](https://opendocs.alipay.com/open-v3/85f9492f_alipay.marketing.card.activateurl.apply) | 获取会员卡领卡投放链接 | `POST` | `/v3/alipay/marketing/card/activateurl/apply` |
| [`alipay.marketing.card.benefit.create`](https://opendocs.alipay.com/open-v3/2aa50202_alipay.marketing.card.benefit.create?scene=common&pathHash=c87e30d3) | 会员卡模板外部权益创建 | `POST` | `/v3/alipay/marketing/card/benefit/create` |
| [`alipay.marketing.card.benefit.delete`](https://opendocs.alipay.com/open-v3/85008a68_alipay.marketing.card.benefit.delete?pathHash=112339fb) | 会员卡模板外部权益删除 | `DELETE` | `/v3/alipay/marketing/card/benefit/delete` |
| [`alipay.marketing.card.benefit.modify`](https://opendocs.alipay.com/open-v3/f8a32ad5_alipay.marketing.card.benefit.modify?pathHash=df326137) | 会员卡模板外部权益修改 | `POST` | `/v3/alipay/marketing/card/benefit/modify` |
| [`alipay.marketing.card.benefit.query`](https://opendocs.alipay.com/open-v3/217a04aa_alipay.marketing.card.benefit.query?pathHash=c22b4480) | 会员卡模板外部权益查询 | `GET` | `/v3/alipay/marketing/card/benefit/query` |
| [`alipay.marketing.card.consume.sync`](https://opendocs.alipay.com/open-v3/e0f362b2_alipay.marketing.card.consume.sync) | 会员卡消费记录同步 | `POST` | `/v3/alipay/marketing/card/consume/sync` |
| [`alipay.marketing.card.delete`](https://opendocs.alipay.com/open-v3/fa4f5056_alipay.marketing.card.delete?scene=common&pathHash=b9e0e504) | 会员卡删卡 | `DELETE` | `/v3/alipay/marketing/card/delete` |
| [`alipay.marketing.card.formtemplate.set`](https://opendocs.alipay.com/open-v3/c7503ca0_alipay.marketing.card.formtemplate.set?scene=common&pathHash=36809f41) | 会员卡开卡表单模板配置 | `POST` | `/v3/alipay/marketing/card/formtemplate/set` |
| [`alipay.marketing.card.message.notify`](https://opendocs.alipay.com/open-v3/b710cee9_alipay.marketing.card.message.notify?scene=common&pathHash=9eb33531) | 会员卡消息通知 | `POST` | `/v3/alipay/marketing/card/message/notify` |
| [`alipay.marketing.card.open`](https://opendocs.alipay.com/open-v3/8e9fe25b_alipay.marketing.card.open) | 会员卡开卡 | `POST` | `/v3/alipay/marketing/card/open` |
| [`alipay.marketing.card.query`](https://opendocs.alipay.com/open-v3/b2494988_alipay.marketing.card.query?scene=common&pathHash=842cfce1) | 会员卡查询 | `POST` | `/v3/alipay/marketing/card/query` |
| [`alipay.marketing.card.template.batchquery`](https://opendocs.alipay.com/open-v3/54dcbcfa_alipay.marketing.card.template.batchquery?pathHash=030a7dec) | 会员卡模板批量查询 | `POST` | `/v3/alipay/marketing/card/template/batchquery` |
| [`alipay.marketing.card.template.create`](https://opendocs.alipay.com/open-v3/18fdd0bd_alipay.marketing.card.template.create?scene=1c74e13b68cc4a58912fb6f0c40b62af%2C1529c3090a7b4db99164667a0b22dcaa&pathHash=d206dedb) | 会员卡模板创建 | `POST` | `/v3/alipay/marketing/card/template/create` |
| [`alipay.marketing.card.template.modify`](https://opendocs.alipay.com/open-v3/fabfc518_alipay.marketing.card.template.modify?scene=692a39a5ad9e45d4899177a7ec9cf68e%2Cc76bbe273e3b471bbba690c624b9ed86&pathHash=e43d6caf) | 会员卡模板修改 | `POST` | `/v3/alipay/marketing/card/template/modify` |
| [`alipay.marketing.card.template.query`](https://opendocs.alipay.com/open-v3/0bd1a249_alipay.marketing.card.template.query?scene=871b0111e48c4490ad0a7dbf2a465b36%2C1eca10e1486044f9b3a0d9171111f810&pathHash=84fe74c8) | 会员卡模板查询接口 | `GET` | `/v3/alipay/marketing/card/template/query` |
| [`alipay.marketing.card.update`](https://opendocs.alipay.com/open-v3/c83a3e12_alipay.marketing.card.update?scene=common&pathHash=dc8a305e) | 会员卡更新 | `POST` | `/v3/alipay/marketing/card/update` |
| [`alipay.marketing.material.create`](https://opendocs.alipay.com/apis/05xk1p) | 创建支付宝素材 | `POST` | `/v3/alipay/marketing/material/create` |
| [`alipay.marketing.material.image.upload`](https://opendocs.alipay.com/open-v3/574939dd_alipay.marketing.material.image.upload) | 营销图片资源上传接口 | `POST` | `/v3/alipay/marketing/material/image` |
| [`alipay.marketing.recruit.enroll.close`](https://opendocs.alipay.com/pre-apis/05xe2w) | 下线报名 | `POST` | `/v3/alipay/marketing/recruit/enroll/close` |
| [`alipay.marketing.recruit.enroll.create`](https://opendocs.alipay.com/pre-apis/05xqxy) | 招商报名提交 | `POST` | `/v3/alipay/marketing/recruit/enroll/create` |
| [`alipay.marketing.recruit.enroll.query`](https://opendocs.alipay.com/pre-apis/05xk7o) | 报名详情查询 | `GET` | `/v3/alipay/marketing/recruit/enroll/query` |
| [`alipay.marketing.recruit.plan.query`](https://opendocs.alipay.com/pre-apis/05xhns) | 招商方案详情查询 | `POST` | `/v3/alipay/marketing/recruit/plan/query` |
| [`alipay.marketing.recruit.planlist.query`](https://opendocs.alipay.com/pre-apis/05xk3f) | 查询可报名的方案列表 | `POST` | `/v3/alipay/marketing/recruit/planlist/query` |

### `alipay.merchant`（18）

| 接口 | 名称 | 方法 | 路径 |
| --- | --- | --- | --- |
| [`alipay.merchant.anchor.info.query`](https://opendocs.alipay.com/pre-apis/05yqsw) | 查询主播信息 | `GET` | `/v3/alipay/merchant/anchor/query` |
| [`alipay.merchant.image.upload`](https://opendocs.alipay.com/open-v3/132c6b5b_alipay.merchant.image.upload) | 商户上传处理图片 | `POST` | `/v3/alipay/merchant/image/upload` |
| [`alipay.merchant.indirect.authorder.close`](https://opendocs.alipay.com/apis/063knt) | 商家认证申请单撤销 | `PUT` | `/v3/alipay/merchant/indirect/authorder/close` |
| [`alipay.merchant.indirect.authorder.create`](https://opendocs.alipay.com/apis/05gz1v) | 商家认证申请单提交 | `POST` | `/v3/alipay/merchant/indirect/authorder/create` |
| [`alipay.merchant.indirect.authorder.querystatus`](https://opendocs.alipay.com/apis/063luk) | 查询商家认证申请单状态 | `POST` | `/v3/alipay/merchant/indirect/authorder/querystatus` |
| [`alipay.merchant.indirect.smidbind.query`](https://opendocs.alipay.com/apis/063ec9) | 获取商家认证状态 | `GET` | `/v3/alipay/merchant/indirect/smidbind/query` |
| [`alipay.merchant.iot.device.bind`](https://opendocs.alipay.com/open-v3/c82d5f5f_alipay.merchant.iot.device.bind?scene=common&pathHash=0fb151a0) | IoT设备绑定门店 | `POST` | `/v3/alipay/merchant/iot/device/bind` |
| [`alipay.merchant.iot.device.query`](https://opendocs.alipay.com/open-v3/9b618feb_alipay.merchant.iot.device.query?scene=common&pathHash=9f128d89) | IoT设备绑定关系查询 | `GET` | `/v3/alipay/merchant/iot/device/query` |
| [`alipay.merchant.iot.device.verify`](https://opendocs.alipay.com/open-v3/4e0eced6_alipay.merchant.iot.device.verify?scene=common&pathHash=39ad307c) | IoT设备绑定校验 | `POST` | `/v3/alipay/merchant/iot/device/verify` |
| [`alipay.merchant.item.file.upload`](https://opendocs.alipay.com/open-v3/1f740bc5_alipay.merchant.item.file.upload) | 商品文件上传接口 | `POST` | `/v3/alipay/merchant/item/file/upload` |
| [`alipay.merchant.live.channel.query`](https://opendocs.alipay.com/apis/069r05) | 渠道信息查询 | `GET` | `/v3/alipay/merchant/live/channel/query` |
| [`alipay.merchant.live.itemplan.modify`](https://opendocs.alipay.com/apis/069o8w) | 支付宝货品推广计划创建 | `POST` | `/v3/alipay/merchant/live/itemplan/modify` |
| [`alipay.merchant.order.sync`](https://opendocs.alipay.com/open-v3/be2462f6_alipay.merchant.order.sync) | 订单数据同步接口 | `POST` | `/v3/alipay/merchant/order/sync` |
| [`alipay.merchant.tradecomplain.batchquery`](https://opendocs.alipay.com/open-v3/152a9215_alipay.merchant.tradecomplain.batchquery) | 查询交易投诉列表 | `POST` | `/v3/alipay/merchant/tradecomplain/batchquery` |
| [`alipay.merchant.tradecomplain.feedback.submit`](https://opendocs.alipay.com/open-v3/da1eaf4f_alipay.merchant.tradecomplain.feedback.submit) | 商家处理交易投诉 | `POST` | `/v3/alipay/merchant/tradecomplain/feedback` |
| [`alipay.merchant.tradecomplain.query`](https://opendocs.alipay.com/open-v3/cecf0505_alipay.merchant.tradecomplain.query) | 查询单条交易投诉详情 | `GET` | `/v3/alipay/merchant/tradecomplain/{complain_event_id}` |
| [`alipay.merchant.tradecomplain.reply.submit`](https://opendocs.alipay.com/open-v3/5851abbd_alipay.merchant.tradecomplain.reply.submit) | 商家留言回复 | `POST` | `/v3/alipay/merchant/tradecomplain/reply` |
| [`alipay.merchant.tradecomplain.supplement.submit`](https://opendocs.alipay.com/open-v3/bb0a25ee_alipay.merchant.tradecomplain.supplement.submit) | 商家补充凭证 | `POST` | `/v3/alipay/merchant/tradecomplain/supplement` |

### `alipay.mobile`（14）

| 接口 | 名称 | 方法 | 路径 |
| --- | --- | --- | --- |
| [`alipay.mobile.public.follow.list`](https://opendocs.alipay.com/apis/05x9eb) | 获取关注者列表 | `GET` | `/v3/alipay/mobile/public/follow/list` |
| [`alipay.mobile.public.info.query`](https://opendocs.alipay.com/apis/05wuox) | 服务窗基础信息查询接口 | `GET` | `/v3/alipay/mobile/public/info/query` |
| [`alipay.mobile.public.menu.add`](https://opendocs.alipay.com/pre-apis/05wuom) | 创建菜单 | `POST` | `/v3/alipay/mobile/public/menu/add` |
| [`alipay.mobile.public.menu.get`](https://opendocs.alipay.com/pre-apis/05wyfs) | 查询菜单 | `GET` | `/v3/alipay/mobile/public/menu/get` |
| [`alipay.mobile.public.menu.update`](https://opendocs.alipay.com/pre-apis/05x1fv) | 更新菜单 | `POST` | `/v3/alipay/mobile/public/menu/update` |
| [`alipay.mobile.public.message.custom.send`](https://opendocs.alipay.com/apis/05wuk7) | 异步单发消息 | `POST` | `/v3/alipay/mobile/public/message/custom/send` |
| [`alipay.mobile.public.message.single.send`](https://opendocs.alipay.com/apis/06bf6f) | 单发模板消息 | `POST` | `/v3/alipay/mobile/public/message/single/send` |
| [`alipay.mobile.public.message.total.send`](https://opendocs.alipay.com/pre-apis/05wyco) | 群发消息 | `POST` | `/v3/alipay/mobile/public/message/total/send` |
| [`alipay.mobile.public.qrcode.create`](https://opendocs.alipay.com/apis/05wumi) | 生成带参数的二维码 | `POST` | `/v3/alipay/mobile/public/qrcode/create` |
| [`alipay.mobile.public.shortlink.create`](https://opendocs.alipay.com/apis/0i05hc) | 服务窗短链自主生成接口 | `POST` | `/v3/alipay/mobile/public/shortlink/create` |
| [`alipay.mobile.public.template.message.delete`](https://opendocs.alipay.com/pre-apis/05xsen) | 模板消息删除接口 | `POST` | `/v3/alipay/mobile/public/template/message/delete` |
| [`alipay.mobile.public.template.message.get`](https://opendocs.alipay.com/apis/05wh9v) | 模板消息领取接口 | `GET` | `/v3/alipay/mobile/public/template/message/get` |
| [`alipay.mobile.public.template.message.query`](https://opendocs.alipay.com/pre-apis/05xdxr) | 模板消息查询接口 | `POST` | `/v3/alipay/mobile/public/template/message/query` |
| [`alipay.mobile.std.public.message.custom.send`](https://opendocs.alipay.com/pre-apis/05wult) | 异步单发消息 | `POST` | `/v3/alipay/mobile/std/public/message/custom/send` |

### `alipay.offline`（8）

| 接口 | 名称 | 方法 | 路径 |
| --- | --- | --- | --- |
| [`alipay.offline.market.applyorder.batchquery`](https://opendocs.alipay.com/apis/05xk26) | 业务流水批量查询接口 | `POST` | `/v3/alipay/offline/market/applyorder/batchquery` |
| [`alipay.offline.market.shop.batchquery`](https://opendocs.alipay.com/apis/05xf6f) | 查询商户的门店编号列表 | `POST` | `/v3/alipay/offline/market/shop/batchquery` |
| [`alipay.offline.market.shop.category.query`](https://opendocs.alipay.com/apis/05xdx8) | 门店类目配置查询接口 | `GET` | `/v3/alipay/offline/market/shop/category/query` |
| [`alipay.offline.market.shop.create`](https://opendocs.alipay.com/apis/05woku) | 创建门店信息 | `POST` | `/v3/alipay/offline/market/shop/create` |
| [`alipay.offline.market.shop.modify`](https://opendocs.alipay.com/apis/05wx1o) | 修改门店信息 | `POST` | `/v3/alipay/offline/market/shop/modify` |
| [`alipay.offline.market.shop.querydetail`](https://opendocs.alipay.com/apis/05wuly) | 查询单个门店信息接口 | `GET` | `/v3/alipay/offline/market/shop/querydetail` |
| [`alipay.offline.market.shop.summary.batchquery`](https://opendocs.alipay.com/apis/05wx2m) | 门店摘要信息批量查询接口 | `POST` | `/v3/alipay/offline/market/shop/summary/batchquery` |
| [`alipay.offline.material.image.upload`](https://opendocs.alipay.com/open-v3/89e1d8f6_alipay.offline.material.image.upload?scene=common&pathHash=92bedf84) | 上传门店照片和视频接口 | `POST` | `/v3/alipay/offline/material/image/upload` |

### `alipay.open`（241）

| 接口 | 名称 | 方法 | 路径 |
| --- | --- | --- | --- |
| [`alipay.open.agent.cancel`](https://opendocs.alipay.com/open-v3/8e66406f_alipay.open.agent.cancel) | 取消代商户签约、创建应用事务 | `POST` | `/v3/alipay/open/agent/cancel` |
| [`alipay.open.agent.common.sign`](https://opendocs.alipay.com/open-v3/f0ab46ad_alipay.open.agent.common.sign) | 代签约产品通用接口 | `POST` | `/v3/alipay/open/agent/common/sign` |
| [`alipay.open.agent.commonsign.confirm`](https://opendocs.alipay.com/open-v3/2f809b83_alipay.open.agent.commonsign.confirm) | 代商户签约，提交信息确认接口 | `POST` | `/v3/alipay/open/agent/commonsign/confirm` |
| [`alipay.open.agent.confirm`](https://opendocs.alipay.com/open-v3/e853dfa1_alipay.open.agent.confirm) | 提交代商户签约、创建应用事务 | `POST` | `/v3/alipay/open/agent/confirm` |
| [`alipay.open.agent.create`](https://opendocs.alipay.com/open-v3/b4c46445_alipay.open.agent.create) | 开启代商户签约、创建应用事务 | `POST` | `/v3/alipay/open/agent/create` |
| [`alipay.open.agent.facetoface.sign`](https://opendocs.alipay.com/open-v3/e1eb4236_alipay.open.agent.facetoface.sign) | 代签约当面付产品 | `POST` | `/v3/alipay/open/agent/facetoface/sign` |
| [`alipay.open.agent.mini.create`](https://opendocs.alipay.com/open-v3/e39a93e8_alipay.open.agent.mini.create) | 代商家创建小程序应用 | `POST` | `/v3/alipay/open/agent/mini/create` |
| [`alipay.open.agent.mobilepay.sign`](https://opendocs.alipay.com/open-v3/db21f6cb_alipay.open.agent.mobilepay.sign) | 代签约APP支付产品 | `POST` | `/v3/alipay/open/agent/mobilepay/sign` |
| [`alipay.open.agent.offlinepayment.sign`](https://opendocs.alipay.com/open-v3/8b80b506_alipay.open.agent.offlinepayment.sign?scene=common&pathHash=255833d8) | 代签约当面付快捷版产品 | `POST` | `/v3/alipay/open/agent/offlinepayment/sign` |
| [`alipay.open.agent.order.query`](https://opendocs.alipay.com/open-v3/ceae4d66_alipay.open.agent.order.query) | 查询申请单状态 | `GET` | `/v3/alipay/open/agent/order/query` |
| [`alipay.open.agent.signstatus.query`](https://opendocs.alipay.com/open-v3/031e92b9_alipay.open.agent.signstatus.query) | 查询商户某个产品的签约状态 | `POST` | `/v3/alipay/open/agent/signstatus/query` |
| [`alipay.open.app.api.field.apply`](https://opendocs.alipay.com/open-v3/db6692ea_alipay.open.app.api.field.apply) | 申请获取接口用户敏感信息字段 | `POST` | `/v3/alipay/open/app/api/field/apply` |
| [`alipay.open.app.api.field.query`](https://opendocs.alipay.com/open-v3/0cfbef83_alipay.open.app.api.field.query) | 用户信息申请记录查询 | `GET` | `/v3/alipay/open/app/api/field/query` |
| [`alipay.open.app.api.query`](https://opendocs.alipay.com/open-v3/d091ad0b_alipay.open.app.api.query) | 查询应用可申请的接口出参敏感字段列表 | `GET` | `/v3/alipay/open/app/api/query` |
| [`alipay.open.app.api.scene.query`](https://opendocs.alipay.com/open-v3/2bc4666c_alipay.open.app.api.scene.query) | 查询接口字段使用场景 | `GET` | `/v3/alipay/open/app/api/scene/query` |
| [`alipay.open.app.members.create`](https://opendocs.alipay.com/open-v3/8b6e0190_alipay.open.app.members.create) | 应用添加成员 | `POST` | `/v3/alipay/open/app/members/create` |
| [`alipay.open.app.members.delete`](https://opendocs.alipay.com/open-v3/26ef7be4_alipay.open.app.members.delete) | 应用删除成员 | `DELETE` | `/v3/alipay/open/app/members/delete` |
| [`alipay.open.app.members.query`](https://opendocs.alipay.com/open-v3/6a8b8874_alipay.open.app.members.query) | 应用查询成员列表 | `GET` | `/v3/alipay/open/app/members/query` |
| [`alipay.open.app.messagetemplate.subscribe.query`](https://opendocs.alipay.com/open-v3/af42509c_alipay.open.app.messagetemplate.subscribe.query) | 模板订阅关系查询 | `POST` | `/v3/alipay/open/app/messagetemplate/subscribe/query` |
| [`alipay.open.app.mini.templatemessage.send`](https://opendocs.alipay.com/open-v3/0136e987_alipay.open.app.mini.templatemessage.send) | 小程序发送模板消息 | `POST` | `/v3/alipay/open/app/mini/templatemessage/send` |
| [`alipay.open.app.qrcode.create`](https://opendocs.alipay.com/open-v3/76c2627e_alipay.open.app.qrcode.create) | 小程序生成推广二维码接口 | `POST` | `/v3/alipay/open/app/qrcode/create` |
| [`alipay.open.app.service.apply`](https://opendocs.alipay.com/open-v3/8b18b82d_alipay.open.app.service.apply) | 服务提报申请 | `POST` | `/v3/alipay/open/app/service/apply` |
| [`alipay.open.app.service.delete`](https://opendocs.alipay.com/open-v3/5b7ade20_alipay.open.app.service.delete) | 服务失效接口 | `DELETE` | `/v3/alipay/open/app/service/delete` |
| [`alipay.open.app.service.list.query`](https://opendocs.alipay.com/open-v3/77d4e708_alipay.open.app.service.list.query) | 服务批量查询 | `GET` | `/v3/alipay/open/app/service/list/query` |
| [`alipay.open.app.service.query`](https://opendocs.alipay.com/open-v3/86469046_alipay.open.app.service.query) | 服务查询 | `GET` | `/v3/alipay/open/app/service/query` |
| [`alipay.open.app.service.schema.query`](https://opendocs.alipay.com/open-v3/afe544b5_alipay.open.app.service.schema.query) | 服务schema下发 | `GET` | `/v3/alipay/open/app/service/schema/query` |
| [`alipay.open.auth.app.aes.get`](https://opendocs.alipay.com/open-v3/4d82ebbf_alipay.open.auth.app.aes.get) | 授权应用aes密钥查询 | `GET` | `/v3/alipay/open/auth/app/aes/get` |
| [`alipay.open.auth.app.aes.set`](https://opendocs.alipay.com/open-v3/306fd1a6_alipay.open.auth.app.aes.set) | 授权应用aes密钥设置 | `POST` | `/v3/alipay/open/auth/app/aes/set` |
| [`alipay.open.auth.token.app`](https://opendocs.alipay.com/open-v3/4b00cc9c_alipay.open.auth.token.app) | 换取应用授权令牌 | `POST` | `/v3/alipay/open/auth/token/app` |
| [`alipay.open.auth.token.app.query`](https://opendocs.alipay.com/open-v3/b6e7c6b3_alipay.open.auth.token.app.query) | 查询某个应用授权AppAuthToken的授权信息 | `GET` | `/v3/alipay/open/auth/token/app/query` |
| [`alipay.open.auth.userauth.relationship.query`](https://opendocs.alipay.com/open-v3/5ea1017e_alipay.open.auth.userauth.relationship.query) | 用户授权关系查询 | `GET` | `/v3/alipay/open/auth/userauth/relationship/query` |
| [`alipay.open.biz.create`](https://opendocs.alipay.com/apis/0i01un) | bizcontent测试12 | `POST` | `/v3/alipay/open/biz/create` |
| [`alipay.open.fee.adjust.apply`](https://opendocs.alipay.com/open-v3/fe77ed62_alipay.open.fee.adjust.apply) | 特殊费率申请 | `POST` | `/v3/alipay/open/fee/adjust/apply` |
| [`alipay.open.file.upload`](https://opendocs.alipay.com/open-v3/5aa91070_alipay.open.file.upload?scene=common&pathHash=e40a9831) | 支付宝文件上传接口 | `POST` | `/v3/alipay/open/file/upload` |
| [`alipay.open.instantdelivery.account.create`](https://opendocs.alipay.com/pre-apis/05vlui) | 即时配送商家账户创建 | `POST` | `/v3/alipay/open/instantdelivery/account/create` |
| [`alipay.open.instantdelivery.account.query`](https://opendocs.alipay.com/pre-apis/05vda7) | 即时配送商家账户查询 | `GET` | `/v3/alipay/open/instantdelivery/account/query` |
| [`alipay.open.instantdelivery.merchantagreement.query`](https://opendocs.alipay.com/pre-apis/05vlug) | 协议签署查询api | `GET` | `/v3/alipay/open/instantdelivery/merchantagreement/query` |
| [`alipay.open.instantdelivery.merchantshop.batchquery`](https://opendocs.alipay.com/pre-apis/05vn66) | 即时配送商家门店分页查询 | `POST` | `/v3/alipay/open/instantdelivery/merchantshop/batchquery` |
| [`alipay.open.instantdelivery.merchantshop.create`](https://opendocs.alipay.com/pre-apis/05vluh) | 即时配送商家门店创建 | `POST` | `/v3/alipay/open/instantdelivery/merchantshop/create` |
| [`alipay.open.instantdelivery.merchantshop.modify`](https://opendocs.alipay.com/pre-apis/068gkz) | 即时配送商家门店更新 | `POST` | `/v3/alipay/open/instantdelivery/merchantshop/modify` |
| [`alipay.open.instantdelivery.merchantshop.query`](https://opendocs.alipay.com/pre-apis/068ja3) | 即时配送商家门店详情查询 | `GET` | `/v3/alipay/open/instantdelivery/merchantshop/query` |
| [`alipay.open.invite.order.query`](https://opendocs.alipay.com/open-v3/a5f3aaab_alipay.open.invite.order.query) | 查询签约申请单状态 | `GET` | `/v3/alipay/open/invite/order/query` |
| [`alipay.open.mini.appdeploy.byappversion.query`](https://opendocs.alipay.com/pre-apis/05xo5a) | 通过版本查询小程序发布 | `GET` | `/v3/alipay/open/mini/appdeploy/byappversion/query` |
| [`alipay.open.mini.appdeploy.bydeployversion.query`](https://opendocs.alipay.com/pre-apis/05xaay) | 根据发布版本查询小程序发布信息 | `GET` | `/v3/alipay/open/mini/appdeploy/bydeployversion/query` |
| [`alipay.open.mini.baseinfo.modify`](https://opendocs.alipay.com/open-v3/762743d3_alipay.open.mini.baseinfo.modify) | 小程序修改基础信息 | `POST` | `/v3/alipay/open/mini/baseinfo/modify` |
| [`alipay.open.mini.baseinfo.name.check`](https://opendocs.alipay.com/open-v3/bf9816ac_alipay.open.mini.baseinfo.name.check) | 小程序名称检测 | `GET` | `/v3/alipay/open/mini/baseinfo/name/check` |
| [`alipay.open.mini.baseinfo.query`](https://opendocs.alipay.com/open-v3/c5995fc4_alipay.open.mini.baseinfo.query) | 查询小程序基础信息 | `GET` | `/v3/alipay/open/mini/baseinfo/query` |
| [`alipay.open.mini.category.query`](https://opendocs.alipay.com/open-v3/bb00b1a6_alipay.open.mini.category.query) | 小程序类目树查询 | `GET` | `/v3/alipay/open/mini/category/query` |
| [`alipay.open.mini.category.require.query`](https://opendocs.alipay.com/open-v3/7f3621b8_alipay.open.mini.category.require.query) | 查询类目所需资质信息 | `POST` | `/v3/alipay/open/mini/category/require/query` |
| [`alipay.open.mini.experience.cancel`](https://opendocs.alipay.com/open-v3/d4e75c29_alipay.open.mini.experience.cancel) | 小程序取消体验版 | `POST` | `/v3/alipay/open/mini/experience/cancel` |
| [`alipay.open.mini.experience.create`](https://opendocs.alipay.com/open-v3/287a2040_alipay.open.mini.experience.create) | 小程序生成体验版 | `POST` | `/v3/alipay/open/mini/experience/create` |
| [`alipay.open.mini.experience.query`](https://opendocs.alipay.com/open-v3/99c5f4ae_alipay.open.mini.experience.query) | 小程序体验版状态查询接口 | `GET` | `/v3/alipay/open/mini/experience/query` |
| [`alipay.open.mini.individual.business.certify`](https://opendocs.alipay.com/open-v3/7ba79277_alipay.open.mini.individual.business.certify) | 个人账户升级为个体工商户 | `POST` | `/v3/alipay/open/mini/individual/business/certify` |
| [`alipay.open.mini.inner.experience.cancel`](https://opendocs.alipay.com/pre-apis/05xe27) | 内部链路取消体验版 | `POST` | `/v3/alipay/open/mini/inner/experience/cancel` |
| [`alipay.open.mini.inner.members.add`](https://opendocs.alipay.com/pre-apis/05xhp2) | 内部链路添加开发成员 | `POST` | `/v3/alipay/open/mini/inner/members/add` |
| [`alipay.open.mini.inner.members.delete`](https://opendocs.alipay.com/pre-apis/05xqzc) | 内部链路删除开发成员 | `DELETE` | `/v3/alipay/open/mini/inner/members/delete` |
| [`alipay.open.mini.inner.members.query`](https://opendocs.alipay.com/pre-apis/05xaad) | 内部链路查询开发成员 | `GET` | `/v3/alipay/open/mini/inner/members/query` |
| [`alipay.open.mini.innerapp.create`](https://opendocs.alipay.com/pre-apis/05x9i7) | 内部链路小程序创建 | `POST` | `/v3/alipay/open/mini/innerapp/create` |
| [`alipay.open.mini.innerapp.plugin.order`](https://opendocs.alipay.com/pre-apis/05xng4) | 小程序代商户订购插件 | `POST` | `/v3/alipay/open/mini/innerapp/plugin/order` |
| [`alipay.open.mini.innerapp.pluginrelation.query`](https://opendocs.alipay.com/pre-apis/05xhsh) | 查询插件引用关系 | `POST` | `/v3/alipay/open/mini/innerapp/pluginrelation/query` |
| [`alipay.open.mini.innerapp.pluginservice.publish`](https://opendocs.alipay.com/pre-apis/05xe0v) | 插件发布能力中心 | `POST` | `/v3/alipay/open/mini/innerapp/pluginservice/publish` |
| [`alipay.open.mini.innerapp.pluginsyncmode.modify`](https://opendocs.alipay.com/pre-apis/05xo5c) | 修改插件继承模式 | `POST` | `/v3/alipay/open/mini/innerapp/pluginsyncmode/modify` |
| [`alipay.open.mini.innerapp.pluginuseconfig.cancel`](https://opendocs.alipay.com/pre-apis/05xgfv) | 撤销插件端版本配置灰度 | `POST` | `/v3/alipay/open/mini/innerapp/pluginuseconfig/cancel` |
| [`alipay.open.mini.innerapp.pluginuseconfig.online`](https://opendocs.alipay.com/pre-apis/05xacn) | 上线插件使用端版本配置 | `POST` | `/v3/alipay/open/mini/innerapp/pluginuseconfig/online` |
| [`alipay.open.mini.innerapp.pluginuseconfig.upgrade`](https://opendocs.alipay.com/pre-apis/05xo8d) | 灰度端版本配置 | `POST` | `/v3/alipay/open/mini/innerapp/pluginuseconfig/upgrade` |
| [`alipay.open.mini.innerapp.service.publish`](https://opendocs.alipay.com/pre-apis/05xndu) | 小程序模板发布服务市场 | `POST` | `/v3/alipay/open/mini/innerapp/service/publish` |
| [`alipay.open.mini.innerapp.service.query`](https://opendocs.alipay.com/pre-apis/05xk3b) | 小程序服务市场搜索 | `GET` | `/v3/alipay/open/mini/innerapp/service/query` |
| [`alipay.open.mini.innerbaseinfo.applogo.upload`](https://opendocs.alipay.com/pre-apis/05xdzx) | 内部小程序-小程序logo图片上传 | `POST` | `/v3/alipay/open/mini/innerbaseinfo/applogo/upload` |
| [`alipay.open.mini.innerbaseinfo.query`](https://opendocs.alipay.com/pre-apis/05xiqj) | 内部小程序-应用信息查询 | `GET` | `/v3/alipay/open/mini/innerbaseinfo/query` |
| [`alipay.open.mini.innerversion.audit.cancel`](https://opendocs.alipay.com/pre-apis/05xaab) | 内部链路撤销审核 | `POST` | `/v3/alipay/open/mini/innerversion/audit/cancel` |
| [`alipay.open.mini.innerversion.audit.submit`](https://opendocs.alipay.com/pre-apis/05xnee) | 小程序版本提交审核 | `POST` | `/v3/alipay/open/mini/innerversion/audit/submit` |
| [`alipay.open.mini.innerversion.backdev.publish`](https://opendocs.alipay.com/pre-apis/05xqxi) | 内部链路小程序退回开发 | `POST` | `/v3/alipay/open/mini/innerversion/backdev/publish` |
| [`alipay.open.mini.innerversion.beta.confirm`](https://opendocs.alipay.com/pre-apis/05x9jh) | 确认邀测结果 | `POST` | `/v3/alipay/open/mini/innerversion/beta/confirm` |
| [`alipay.open.mini.innerversion.betainfo.query`](https://opendocs.alipay.com/pre-apis/05xp3o) | 查看可邀测状态接口 | `GET` | `/v3/alipay/open/mini/innerversion/betainfo/query` |
| [`alipay.open.mini.innerversion.condition.batchquery`](https://opendocs.alipay.com/pre-apis/05xk66) | 根据状态批量查询版本信息 | `POST` | `/v3/alipay/open/mini/innerversion/condition/batchquery` |
| [`alipay.open.mini.innerversion.content.rollback`](https://opendocs.alipay.com/pre-apis/05xp3p) | 内部链路回滚版本 | `POST` | `/v3/alipay/open/mini/innerversion/content/rollback` |
| [`alipay.open.mini.innerversion.gray.finish`](https://opendocs.alipay.com/pre-apis/05xge5) | 内部链路结束灰度 | `POST` | `/v3/alipay/open/mini/innerversion/gray/finish` |
| [`alipay.open.mini.innerversion.gray.publish`](https://opendocs.alipay.com/pre-apis/05xk4w) | 内部链路小程序灰度上架 | `POST` | `/v3/alipay/open/mini/innerversion/gray/publish` |
| [`alipay.open.mini.innerversion.info.delete`](https://opendocs.alipay.com/pre-apis/05xe18) | 内部链路删除版本 | `DELETE` | `/v3/alipay/open/mini/innerversion/info/delete` |
| [`alipay.open.mini.innerversion.info.query`](https://opendocs.alipay.com/pre-apis/05xdzi) | 内部链路小程序版本信息查询 | `GET` | `/v3/alipay/open/mini/innerversion/info/query` |
| [`alipay.open.mini.innerversion.instantiation.query`](https://opendocs.alipay.com/pre-apis/05xdyz) | 小程序实例化关系查询 | `GET` | `/v3/alipay/open/mini/innerversion/instantiation/query` |
| [`alipay.open.mini.innerversion.last.query`](https://opendocs.alipay.com/pre-apis/05xhs8) | 内部链路查询最新版本号 | `GET` | `/v3/alipay/open/mini/innerversion/last/query` |
| [`alipay.open.mini.innerversion.modelforoffline.query`](https://opendocs.alipay.com/pre-apis/05xqwp) | 小程序线上版本包数据同步线下场景查询 | `GET` | `/v3/alipay/open/mini/innerversion/modelforoffline/query` |
| [`alipay.open.mini.innerversion.nobuild.upload`](https://opendocs.alipay.com/pre-apis/05xiqc) | 手淘小程序产物包免构建版本上传 | `POST` | `/v3/alipay/open/mini/innerversion/nobuild/upload` |
| [`alipay.open.mini.innerversion.offline.publish`](https://opendocs.alipay.com/pre-apis/05xo7b) | 内部链路小程序下架 | `POST` | `/v3/alipay/open/mini/innerversion/offline/publish` |
| [`alipay.open.mini.innerversion.online`](https://opendocs.alipay.com/pre-apis/05xhq6) | 内部小程序-上架 | `POST` | `/v3/alipay/open/mini/innerversion/online` |
| [`alipay.open.mini.innerversion.online.publish`](https://opendocs.alipay.com/pre-apis/05xgf5) | 内部链路小程序上架 | `POST` | `/v3/alipay/open/mini/innerversion/online/publish` |
| [`alipay.open.mini.innerversion.package.create`](https://opendocs.alipay.com/pre-apis/05x9fo) | 单独构建包信息 | `POST` | `/v3/alipay/open/mini/innerversion/package/create` |
| [`alipay.open.mini.innerversion.sync`](https://opendocs.alipay.com/pre-apis/05xlnt) | 小程序包信息推送 | `POST` | `/v3/alipay/open/mini/innerversion/sync` |
| [`alipay.open.mini.innerversion.templated.upload`](https://opendocs.alipay.com/pre-apis/05xdzq) | 基于模板构建版本 | `POST` | `/v3/alipay/open/mini/innerversion/templated/upload` |
| [`alipay.open.mini.innerversion.upload`](https://opendocs.alipay.com/pre-apis/05xhot) | 内部小程序-打包构建 | `POST` | `/v3/alipay/open/mini/innerversion/upload` |
| [`alipay.open.mini.innerversion.uploadstatus.query`](https://opendocs.alipay.com/pre-apis/05xskn) | 多端小程序-查询构建状态 | `GET` | `/v3/alipay/open/mini/innerversion/uploadstatus/query` |
| [`alipay.open.mini.isv.create`](https://opendocs.alipay.com/open-v3/719ad502_alipay.open.mini.isv.create) | isv服务商代商户创建小程序 | `POST` | `/v3/alipay/open/mini/isv/create` |
| [`alipay.open.mini.isv.fastregister.create`](https://opendocs.alipay.com/open-v3/58d4f210_alipay.open.mini.isv.fastregister.create) | ISV代创建试用小程序 | `POST` | `/v3/alipay/open/mini/isv/fastregister/create` |
| [`alipay.open.mini.isv.fastregister.query`](https://opendocs.alipay.com/open-v3/033e1882_alipay.open.mini.isv.fastregister.query) | 查询试用小程序 | `GET` | `/v3/alipay/open/mini/isv/fastregister/query` |
| [`alipay.open.mini.isv.query`](https://opendocs.alipay.com/open-v3/f6bafc97_alipay.open.mini.isv.query) | isv查询代商家创建小程序记录 | `GET` | `/v3/alipay/open/mini/isv/query` |
| [`alipay.open.mini.miniapp.brand.cancel`](https://opendocs.alipay.com/open-v3/ac02eafe_alipay.open.mini.miniapp.brand.cancel) | 小程序品牌提交认证后取消品牌认证 | `POST` | `/v3/alipay/open/mini/miniapp/brand/cancel` |
| [`alipay.open.mini.miniapp.brand.query`](https://opendocs.alipay.com/open-v3/6187351c_alipay.open.mini.miniapp.brand.query) | 小程序品牌提交认证后查询品牌审核结果以及商户已有品牌 | `GET` | `/v3/alipay/open/mini/miniapp/brand/query` |
| [`alipay.open.mini.miniapp.brand.submit`](https://opendocs.alipay.com/open-v3/24509d98_alipay.open.mini.miniapp.brand.submit) | 小程序品牌提交认证 | `POST` | `/v3/alipay/open/mini/miniapp/brand/submit` |
| [`alipay.open.mini.miniapp.brand.upload`](https://opendocs.alipay.com/open-v3/e77e0097_alipay.open.mini.miniapp.brand.upload) | 小程序品牌提交认证时上传用户资质以及身份信息 | `POST` | `/v3/alipay/open/mini/miniapp/brand/upload` |
| [`alipay.open.mini.miniapp.serviceconfig.modify`](https://opendocs.alipay.com/open-v3/95e9abe6_alipay.open.mini.miniapp.serviceconfig.modify) | 小程序设置客服方式 | `POST` | `/v3/alipay/open/mini/miniapp/serviceconfig/modify` |
| [`alipay.open.mini.qrcode.bind`](https://opendocs.alipay.com/open-v3/01ebf6a3_alipay.open.mini.qrcode.bind) | 关联普通二维码 | `POST` | `/v3/alipay/open/mini/qrcode/bind` |
| [`alipay.open.mini.qrcode.query`](https://opendocs.alipay.com/open-v3/c89a2391_alipay.open.mini.qrcode.query) | 查询普通二维码 | `GET` | `/v3/alipay/open/mini/qrcode/query` |
| [`alipay.open.mini.qrcode.unbind`](https://opendocs.alipay.com/open-v3/8905c9d2_alipay.open.mini.qrcode.unbind) | 删除已关联普通二维码 | `POST` | `/v3/alipay/open/mini/qrcode/unbind` |
| [`alipay.open.mini.resource.promotionsource.notify`](https://opendocs.alipay.com/pre-apis/05xhr9) | 提供给淘联盟推广位信息接口 | `POST` | `/v3/alipay/open/mini/resource/promotionsource/notify` |
| [`alipay.open.mini.resource.record.notify`](https://opendocs.alipay.com/pre-apis/05xqyn) | 提供给淘联盟备案信息接口 | `POST` | `/v3/alipay/open/mini/resource/record/notify` |
| [`alipay.open.mini.safedomain.create`](https://opendocs.alipay.com/open-v3/27ccb222_alipay.open.mini.safedomain.create) | 小程序添加域白名单 | `POST` | `/v3/alipay/open/mini/safedomain/create` |
| [`alipay.open.mini.safedomain.delete`](https://opendocs.alipay.com/open-v3/d96d3a21_alipay.open.mini.safedomain.delete) | 小程序删除域白名单 | `DELETE` | `/v3/alipay/open/mini/safedomain/delete` |
| [`alipay.open.mini.template.marketing.create`](https://opendocs.alipay.com/open-v3/5c89642c_alipay.open.mini.template.marketing.create) | 小程序消息运营位创建 | `POST` | `/v3/alipay/open/mini/template/marketing/create` |
| [`alipay.open.mini.template.usage.query`](https://opendocs.alipay.com/open-v3/883d84eb_alipay.open.mini.template.usage.query) | 查询使用模板的小程序列表 | `GET` | `/v3/alipay/open/mini/template/usage/query` |
| [`alipay.open.mini.templatemsg.maketing.batchquery`](https://opendocs.alipay.com/open-v3/59c80879_alipay.open.mini.templatemsg.maketing.batchquery) | 小程序模板消息营销投放详情分页批量查询 | `POST` | `/v3/alipay/open/mini/templatemsg/maketing/batchquery` |
| [`alipay.open.mini.templatemsg.tinypayswitch.confirm`](https://opendocs.alipay.com/open-v3/653c0f4a_alipay.open.mini.templatemsg.tinypayswitch.confirm) | 小程序支付消息确认接口 | `POST` | `/v3/alipay/open/mini/templatemsg/tinypayswitch/confirm` |
| [`alipay.open.mini.tips.delivery.batchquery`](https://opendocs.alipay.com/open-v3/75394e37_alipay.open.mini.tips.delivery.batchquery) | 小程序收藏引导投放活动配置批量查询 | `POST` | `/v3/alipay/open/mini/tips/delivery/batchquery` |
| [`alipay.open.mini.tips.delivery.create`](https://opendocs.alipay.com/open-v3/ca9b74a9_alipay.open.mini.tips.delivery.create) | 小程序收藏引导投放活动配置创建 | `POST` | `/v3/alipay/open/mini/tips/delivery/create` |
| [`alipay.open.mini.tips.delivery.modify`](https://opendocs.alipay.com/open-v3/a4b55863_alipay.open.mini.tips.delivery.modify) | 小程序收藏引导投放活动修改 | `POST` | `/v3/alipay/open/mini/tips/delivery/modify` |
| [`alipay.open.mini.tips.delivery.query`](https://opendocs.alipay.com/open-v3/f59fb71e_alipay.open.mini.tips.delivery.query) | 小程序收藏引导投放活动详情查询 | `GET` | `/v3/alipay/open/mini/tips/delivery/query` |
| [`alipay.open.mini.tips.statistic.query`](https://opendocs.alipay.com/open-v3/87364a7f_alipay.open.mini.tips.statistic.query) | 小程序收藏引导汇总数据查询 | `GET` | `/v3/alipay/open/mini/tips/statistic/query` |
| [`alipay.open.mini.version.audit.apply`](https://opendocs.alipay.com/open-v3/dd9959d2_alipay.open.mini.version.audit.apply) | 小程序提交审核 | `POST` | `/v3/alipay/open/mini/version/audit/apply` |
| [`alipay.open.mini.version.audit.cancel`](https://opendocs.alipay.com/open-v3/9b50e79c_alipay.open.mini.version.audit.cancel) | 小程序撤销审核 | `POST` | `/v3/alipay/open/mini/version/audit/cancel` |
| [`alipay.open.mini.version.audited.cancel`](https://opendocs.alipay.com/open-v3/7a284018_alipay.open.mini.version.audited.cancel) | 小程序退回开发 | `POST` | `/v3/alipay/open/mini/version/audited/cancel` |
| [`alipay.open.mini.version.build.query`](https://opendocs.alipay.com/open-v3/dda696d3_alipay.open.mini.version.build.query) | 小程序查询版本构建状态 | `GET` | `/v3/alipay/open/mini/version/build/query` |
| [`alipay.open.mini.version.delete`](https://opendocs.alipay.com/open-v3/eef61c13_alipay.open.mini.version.delete) | 小程序删除版本 | `POST` | `/v3/alipay/open/mini/version/delete` |
| [`alipay.open.mini.version.detail.query`](https://opendocs.alipay.com/open-v3/048e5e71_alipay.open.mini.version.detail.query) | 小程序版本详情查询 | `GET` | `/v3/alipay/open/mini/version/detail/query` |
| [`alipay.open.mini.version.gray.cancel`](https://opendocs.alipay.com/open-v3/d539c567_alipay.open.mini.version.gray.cancel) | 小程序结束灰度 | `POST` | `/v3/alipay/open/mini/version/gray/cancel` |
| [`alipay.open.mini.version.gray.online`](https://opendocs.alipay.com/open-v3/6ab395d9_alipay.open.mini.version.gray.online) | 小程序灰度上架 | `POST` | `/v3/alipay/open/mini/version/gray/online` |
| [`alipay.open.mini.version.list.query`](https://opendocs.alipay.com/open-v3/ac91bcd8_alipay.open.mini.version.list.query) | 小程序版本列表查询 | `GET` | `/v3/alipay/open/mini/version/list/query` |
| [`alipay.open.mini.version.offline`](https://opendocs.alipay.com/open-v3/c5ffdf9c_alipay.open.mini.version.offline) | 小程序下架 | `POST` | `/v3/alipay/open/mini/version/offline` |
| [`alipay.open.mini.version.online`](https://opendocs.alipay.com/open-v3/4d9b2578_alipay.open.mini.version.online) | 小程序上架 | `POST` | `/v3/alipay/open/mini/version/online` |
| [`alipay.open.mini.version.rollback`](https://opendocs.alipay.com/open-v3/f1c39e2b_alipay.open.mini.version.rollback) | 小程序回滚 | `POST` | `/v3/alipay/open/mini/version/rollback` |
| [`alipay.open.mini.version.upload`](https://opendocs.alipay.com/open-v3/9d32147b_alipay.open.mini.version.upload) | 小程序基于模板上传版本 | `POST` | `/v3/alipay/open/mini/version/upload` |
| [`alipay.open.public.account.create`](https://opendocs.alipay.com/apis/05wyjj) | 添加绑定商户会员号 | `POST` | `/v3/alipay/open/public/account/create` |
| [`alipay.open.public.account.delete`](https://opendocs.alipay.com/apis/05wuq3) | 解除绑定商户会员号 | `DELETE` | `/v3/alipay/open/public/account/delete` |
| [`alipay.open.public.account.query`](https://opendocs.alipay.com/apis/05wupq) | 查询绑定商户会员号 | `GET` | `/v3/alipay/open/public/account/query` |
| [`alipay.open.public.advert.batchquery`](https://opendocs.alipay.com/apis/05wggg) | 生活号广告位查询接口 | `POST` | `/v3/alipay/open/public/advert/batchquery` |
| [`alipay.open.public.advert.create`](https://opendocs.alipay.com/apis/05xmcx) | 生活号广告位添加接口 | `POST` | `/v3/alipay/open/public/advert/create` |
| [`alipay.open.public.advert.delete`](https://opendocs.alipay.com/apis/05wvvm) | 生活号广告位删除接口 | `DELETE` | `/v3/alipay/open/public/advert/delete` |
| [`alipay.open.public.advert.modify`](https://opendocs.alipay.com/apis/05x1fw) | 生活号广告位修改接口 | `POST` | `/v3/alipay/open/public/advert/modify` |
| [`alipay.open.public.articlesummary.data.batchquery`](https://opendocs.alipay.com/apis/05wz9t) | 图文分析-按时间查询数据接口 | `POST` | `/v3/alipay/open/public/articlesummary/data/batchquery` |
| [`alipay.open.public.contact.follow.batchquery`](https://opendocs.alipay.com/apis/05wnsz) | 查询服务窗联系人关注列表 | `POST` | `/v3/alipay/open/public/contact/follow/batchquery` |
| [`alipay.open.public.default.extension.create`](https://opendocs.alipay.com/apis/05wvst) | 默认扩展区创建接口 | `POST` | `/v3/alipay/open/public/default/extension/create` |
| [`alipay.open.public.follow.batchquery`](https://opendocs.alipay.com/apis/05wcpt) | 获取关注者列表 | `POST` | `/v3/alipay/open/public/follow/batchquery` |
| [`alipay.open.public.group.batchquery`](https://opendocs.alipay.com/apis/05wydu) | 查询用户分组列表 | `POST` | `/v3/alipay/open/public/group/batchquery` |
| [`alipay.open.public.group.create`](https://opendocs.alipay.com/apis/05wukz) | 用户分组创建接口 | `POST` | `/v3/alipay/open/public/group/create` |
| [`alipay.open.public.group.crowd.query`](https://opendocs.alipay.com/apis/05wyfg) | 人群数量查询 | `POST` | `/v3/alipay/open/public/group/crowd/query` |
| [`alipay.open.public.group.delete`](https://opendocs.alipay.com/apis/05wuht) | 用户分组删除接口 | `DELETE` | `/v3/alipay/open/public/group/delete` |
| [`alipay.open.public.group.modify`](https://opendocs.alipay.com/apis/05wcqq) | 用户分组修改接口 | `POST` | `/v3/alipay/open/public/group/modify` |
| [`alipay.open.public.info.modify`](https://opendocs.alipay.com/apis/05wz9l) | 生活号基础信息修改接口 | `POST` | `/v3/alipay/open/public/info/modify` |
| [`alipay.open.public.info.query`](https://opendocs.alipay.com/apis/05wygh) | 生活号基础信息查询接口 | `GET` | `/v3/alipay/open/public/info/query` |
| [`alipay.open.public.label.create`](https://opendocs.alipay.com/pre-apis/05wupo) | 公众号标签管理-添加标签 | `POST` | `/v3/alipay/open/public/label/create` |
| [`alipay.open.public.label.query`](https://opendocs.alipay.com/pre-apis/05xq3d) | 公众号标签管理-查询标签 | `GET` | `/v3/alipay/open/public/label/query` |
| [`alipay.open.public.label.user.query`](https://opendocs.alipay.com/pre-apis/05wum4) | 公众号标签管理-查询用户标签 | `GET` | `/v3/alipay/open/public/label/user/query` |
| [`alipay.open.public.life.aboard.apply`](https://opendocs.alipay.com/apis/05wyga) | 上架生活号 | `POST` | `/v3/alipay/open/public/life/aboard/apply` |
| [`alipay.open.public.life.agentcreate.query`](https://opendocs.alipay.com/apis/05wuo2) | isv代创建生活号申请状态查询接口 | `GET` | `/v3/alipay/open/public/life/agentcreate/query` |
| [`alipay.open.public.life.debark.apply`](https://opendocs.alipay.com/apis/05wnsi) | 下架生活号 | `POST` | `/v3/alipay/open/public/life/debark/apply` |
| [`alipay.open.public.life.label.batchquery`](https://opendocs.alipay.com/apis/05wui2) | 标签批量查询接口 | `POST` | `/v3/alipay/open/public/life/label/batchquery` |
| [`alipay.open.public.life.label.create`](https://opendocs.alipay.com/apis/05wh9j) | 创建标签接口 | `POST` | `/v3/alipay/open/public/life/label/create` |
| [`alipay.open.public.life.label.delete`](https://opendocs.alipay.com/apis/05wvsg) | 标签删除接口 | `DELETE` | `/v3/alipay/open/public/life/label/delete` |
| [`alipay.open.public.life.label.modify`](https://opendocs.alipay.com/apis/05xhmq) | 标签修改接口 | `POST` | `/v3/alipay/open/public/life/label/modify` |
| [`alipay.open.public.life.msg.recall`](https://opendocs.alipay.com/apis/05wieh) | 生活号消息撤回接口 | `POST` | `/v3/alipay/open/public/life/msg/recall` |
| [`alipay.open.public.life.msg.send`](https://opendocs.alipay.com/pre-apis/05wulv) | 生活号广播消息发送接口 | `POST` | `/v3/alipay/open/public/life/msg/send` |
| [`alipay.open.public.matchuser.label.create`](https://opendocs.alipay.com/apis/05wycu) | 用户打标接口 | `POST` | `/v3/alipay/open/public/matchuser/label/create` |
| [`alipay.open.public.matchuser.label.delete`](https://opendocs.alipay.com/apis/05wz8y) | 用户取消标签接口 | `POST` | `/v3/alipay/open/public/matchuser/label/delete` |
| [`alipay.open.public.menu.batchquery`](https://opendocs.alipay.com/apis/05wol5) | 菜单列表查询接口 | `POST` | `/v3/alipay/open/public/menu/batchquery` |
| [`alipay.open.public.menu.create`](https://opendocs.alipay.com/apis/05wvxt) | 默认菜单创建接口 | `POST` | `/v3/alipay/open/public/menu/create` |
| [`alipay.open.public.menu.data.batchquery`](https://opendocs.alipay.com/apis/05wyh5) | 菜单分析数据查询接口 | `POST` | `/v3/alipay/open/public/menu/data/batchquery` |
| [`alipay.open.public.menu.modify`](https://opendocs.alipay.com/apis/05xgci) | 默认菜单更新接口 | `POST` | `/v3/alipay/open/public/menu/modify` |
| [`alipay.open.public.menu.query`](https://opendocs.alipay.com/pre-apis/05wuni) | 公众服务平台查询菜单 | `GET` | `/v3/alipay/open/public/menu/query` |
| [`alipay.open.public.message.content.create`](https://opendocs.alipay.com/apis/05xhmr) | 创建图文消息内容接口 | `POST` | `/v3/alipay/open/public/message/content/create` |
| [`alipay.open.public.message.content.modify`](https://opendocs.alipay.com/apis/05wuj5) | 更新图文消息内容接口 | `POST` | `/v3/alipay/open/public/message/content/modify` |
| [`alipay.open.public.message.custom.send`](https://opendocs.alipay.com/apis/05wcpx) | 异步单发消息 | `POST` | `/v3/alipay/open/public/message/custom/send` |
| [`alipay.open.public.message.group.send`](https://opendocs.alipay.com/apis/05wyfq) | 分组消息发送接口 | `POST` | `/v3/alipay/open/public/message/group/send` |
| [`alipay.open.public.message.label.send`](https://opendocs.alipay.com/pre-apis/05xlk0) | 根据标签组发消息接口 | `POST` | `/v3/alipay/open/public/message/label/send` |
| [`alipay.open.public.message.single.send`](https://opendocs.alipay.com/apis/06bpqn) | 单发模板消息 | `POST` | `/v3/alipay/open/public/message/single/send` |
| [`alipay.open.public.message.total.send`](https://opendocs.alipay.com/apis/05wnte) | 群发消息 | `POST` | `/v3/alipay/open/public/message/total/send` |
| [`alipay.open.public.payee.bind.create`](https://opendocs.alipay.com/apis/05wvrh) | 添加收款账号接口 | `POST` | `/v3/alipay/open/public/payee/bind/create` |
| [`alipay.open.public.payee.bind.delete`](https://opendocs.alipay.com/apis/05xq3f) | 解绑收款账号接口 | `DELETE` | `/v3/alipay/open/public/payee/bind/delete` |
| [`alipay.open.public.personalized.extension.batchquery`](https://opendocs.alipay.com/apis/05wcq6) | 扩展区批量查询接口 | `POST` | `/v3/alipay/open/public/personalized/extension/batchquery` |
| [`alipay.open.public.personalized.extension.create`](https://opendocs.alipay.com/apis/05wcpn) | 个性化扩展区创建接口 | `POST` | `/v3/alipay/open/public/personalized/extension/create` |
| [`alipay.open.public.personalized.extension.delete`](https://opendocs.alipay.com/apis/05wydg) | 扩展区删除接口 | `DELETE` | `/v3/alipay/open/public/personalized/extension/delete` |
| [`alipay.open.public.personalized.menu.create`](https://opendocs.alipay.com/apis/05x9eg) | 个性化菜单创建 | `POST` | `/v3/alipay/open/public/personalized/menu/create` |
| [`alipay.open.public.personalized.menu.delete`](https://opendocs.alipay.com/apis/05xa6t) | 个性化菜单删除 | `DELETE` | `/v3/alipay/open/public/personalized/menu/delete` |
| [`alipay.open.public.qrcode.create`](https://opendocs.alipay.com/apis/05wuhi) | 带参推广二维码接口 | `POST` | `/v3/alipay/open/public/qrcode/create` |
| [`alipay.open.public.setting.category.query`](https://opendocs.alipay.com/apis/05xdxp) | 生活号查询行业设置接口 | `GET` | `/v3/alipay/open/public/setting/category/query` |
| [`alipay.open.public.shortlink.create`](https://opendocs.alipay.com/apis/05xin2) | 服务窗短链自主生成接口 | `POST` | `/v3/alipay/open/public/shortlink/create` |
| [`alipay.open.public.singlearticle.data.batchquery`](https://opendocs.alipay.com/apis/05wuky) | 图文分析-按文章查询数据接口 | `POST` | `/v3/alipay/open/public/singlearticle/data/batchquery` |
| [`alipay.open.public.template.message.get`](https://opendocs.alipay.com/apis/05wbw6) | 消息模板领取接口 | `GET` | `/v3/alipay/open/public/template/message/get` |
| [`alipay.open.public.template.message.industry.modify`](https://opendocs.alipay.com/apis/05wuhh) | 模板消息行业设置修改接口 | `POST` | `/v3/alipay/open/public/template/message/industry/modify` |
| [`alipay.open.public.topic.batchquery`](https://opendocs.alipay.com/apis/05wvr7) | 营销位查询接口 | `POST` | `/v3/alipay/open/public/topic/batchquery` |
| [`alipay.open.public.topic.create`](https://opendocs.alipay.com/apis/05wuj6) | 营销位添加接口 | `POST` | `/v3/alipay/open/public/topic/create` |
| [`alipay.open.public.topic.delete`](https://opendocs.alipay.com/apis/05wyj8) | 营销位删除接口 | `DELETE` | `/v3/alipay/open/public/topic/delete` |
| [`alipay.open.public.topic.modify`](https://opendocs.alipay.com/apis/05wuk1) | 营销位修改接口 | `POST` | `/v3/alipay/open/public/topic/modify` |
| [`alipay.open.public.user.data.batchquery`](https://opendocs.alipay.com/apis/05wum9) | 用户分析数据查询接口 | `POST` | `/v3/alipay/open/public/user/data/batchquery` |
| [`alipay.open.public.user.follow.query`](https://opendocs.alipay.com/apis/05wcq8) | 查询用户是否已经关注了生活号 | `GET` | `/v3/alipay/open/public/user/follow/query` |
| [`alipay.open.search.ability.batchquery`](https://opendocs.alipay.com/pre-apis/05wggi) | 搜索运营列表分页查询 | `POST` | `/v3/alipay/open/search/ability/batchquery` |
| [`alipay.open.search.appkeyword.apply`](https://opendocs.alipay.com/open-v3/5ed663e2_alipay.open.search.appkeyword.apply) | 提报搜索关键词 | `POST` | `/v3/alipay/open/search/appkeyword/apply` |
| [`alipay.open.search.appkeyword.batchquery`](https://opendocs.alipay.com/open-v3/0e3cdc41_alipay.open.search.appkeyword.batchquery) | 查询小程序已配置关键词 | `GET` | `/v3/alipay/open/search/appkeyword/batchquery` |
| [`alipay.open.search.appkeyword.delete`](https://opendocs.alipay.com/open-v3/e7b09aa2_alipay.open.search.appkeyword.delete) | 删除搜索关键词 | `POST` | `/v3/alipay/open/search/appkeyword/delete` |
| [`alipay.open.search.appkeyword.querystatus`](https://opendocs.alipay.com/open-v3/2028e38b_alipay.open.search.appkeyword.querystatus) | 查询小程序搜索关键词的审核工单的状态 | `GET` | `/v3/alipay/open/search/appkeyword/querystatus` |
| [`alipay.open.search.appkeywordquota.query`](https://opendocs.alipay.com/open-v3/764cf4b6_alipay.open.search.appkeywordquota.query) | 查询小程序可配置关键词数 | `GET` | `/v3/alipay/open/search/appkeywordquota/query` |
| [`alipay.open.search.baseorder.modify`](https://opendocs.alipay.com/open-v3/abbe8486_alipay.open.search.baseorder.modify) | 搜索运营提报基础信息工单 | `POST` | `/v3/alipay/open/search/baseorder/modify` |
| [`alipay.open.search.box.apply`](https://opendocs.alipay.com/open-v3/622c6b11_alipay.open.search.box.apply?scene=common&pathHash=d1a9edf5) | 申请创建搜索直达配置 | `POST` | `/v3/alipay/open/search/box/apply` |
| [`alipay.open.search.box.batchquery`](https://opendocs.alipay.com/open-v3/fd1480ed_alipay.open.search.box.batchquery?scene=common&pathHash=4ca4364b) | 批量查询搜索直达配置列表 | `POST` | `/v3/alipay/open/search/box/batchquery` |
| [`alipay.open.search.box.consult`](https://opendocs.alipay.com/open-v3/28d19e54_alipay.open.search.box.consult) | 搜索直达创建预校验 | `POST` | `/v3/alipay/open/search/box/consult` |
| [`alipay.open.search.box.modify`](https://opendocs.alipay.com/open-v3/28bd3996_alipay.open.search.box.modify?scene=common&pathHash=f07276ab) | 修改搜索直达配置 | `POST` | `/v3/alipay/open/search/box/modify` |
| [`alipay.open.search.box.offline`](https://opendocs.alipay.com/open-v3/4d9cc1dc_alipay.open.search.box.offline) | 下架搜索直达 | `POST` | `/v3/alipay/open/search/box/offline` |
| [`alipay.open.search.box.online`](https://opendocs.alipay.com/open-v3/00d18795_alipay.open.search.box.online) | 上架搜索直达 | `POST` | `/v3/alipay/open/search/box/online` |
| [`alipay.open.search.box.query`](https://opendocs.alipay.com/open-v3/97b73d37_alipay.open.search.box.query?scene=common&pathHash=6f2e052c) | 查询搜索直达配置详情 | `GET` | `/v3/alipay/open/search/box/query` |
| [`alipay.open.search.boxactivity.apply`](https://opendocs.alipay.com/open-v3/33107e1b_alipay.open.search.boxactivity.apply) | 申请创建搜索直达活动配置 | `POST` | `/v3/alipay/open/search/boxactivity/apply` |
| [`alipay.open.search.boxactivity.query`](https://opendocs.alipay.com/open-v3/5b195961_alipay.open.search.boxactivity.query) | 查询搜索直达活动配置详情 | `GET` | `/v3/alipay/open/search/boxactivity/query` |
| [`alipay.open.search.orderdetail.query`](https://opendocs.alipay.com/open-v3/d55fc0f0_alipay.open.search.orderdetail.query) | 查询搜索服务工单的详细信息 | `POST` | `/v3/alipay/open/search/orderdetail/query` |
| [`alipay.open.search.serviceorder.batchquery`](https://opendocs.alipay.com/open-v3/13830d89_alipay.open.search.serviceorder.batchquery) | 搜索运营服务查询接口 | `POST` | `/v3/alipay/open/search/serviceorder/batchquery` |
| [`alipay.open.search.subservicekeyword.apply`](https://opendocs.alipay.com/open-v3/183cd891_alipay.open.search.subservicekeyword.apply) | 提报服务关键词 | `POST` | `/v3/alipay/open/search/subservicekeyword/apply` |
| [`alipay.open.search.subservicekeyword.batchquery`](https://opendocs.alipay.com/open-v3/6ee3720b_alipay.open.search.subservicekeyword.batchquery) | 服务关键词批量查询接口 | `GET` | `/v3/alipay/open/search/subservicekeyword/batchquery` |
| [`alipay.open.search.subservicekeyword.delete`](https://opendocs.alipay.com/open-v3/bf575dfc_alipay.open.search.subservicekeyword.delete) | 删除服务关键词 | `POST` | `/v3/alipay/open/search/subservicekeyword/delete` |
| [`alipay.open.search.subservicekeyword.querystatus`](https://opendocs.alipay.com/open-v3/7acef1a0_alipay.open.search.subservicekeyword.querystatus) | 查询小程序服务关键词的审核工单的状态 | `GET` | `/v3/alipay/open/search/subservicekeyword/querystatus` |
| [`alipay.open.servicemarket.commodity.shop.offline`](https://opendocs.alipay.com/open-v3/f56986eb_alipay.open.servicemarket.commodity.shop.offline) | 门店插件下架操作 | `POST` | `/v3/alipay/open/servicemarket/commodity/shop/offline` |
| [`alipay.open.servicemarket.order.accept`](https://opendocs.alipay.com/open-v3/added906_alipay.open.servicemarket.order.accept) | 服务商接单操作 | `POST` | `/v3/alipay/open/servicemarket/order/accept` |
| [`alipay.open.servicemarket.order.create`](https://opendocs.alipay.com/pre-apis/05xo6s) | 服务市场代订购小程序 | `POST` | `/v3/alipay/open/servicemarket/order/create` |
| [`alipay.open.servicemarket.order.item.cancel`](https://opendocs.alipay.com/open-v3/97757082_alipay.open.servicemarket.order.item.cancel) | 服务订单明细实施项单项取消 | `POST` | `/v3/alipay/open/servicemarket/order/item/cancel` |
| [`alipay.open.servicemarket.order.item.complete`](https://opendocs.alipay.com/open-v3/e31fe80d_alipay.open.servicemarket.order.item.complete) | 服务商完成订单内单个明细实施项 | `POST` | `/v3/alipay/open/servicemarket/order/item/complete` |
| [`alipay.open.servicemarket.order.item.confirm`](https://opendocs.alipay.com/open-v3/13c55349_alipay.open.servicemarket.order.item.confirm) | 服务商代商家确认实施完成 | `POST` | `/v3/alipay/open/servicemarket/order/item/confirm` |
| [`alipay.open.servicemarket.order.notify`](https://opendocs.alipay.com/open-v3/2233a7ba_alipay.open.servicemarket.order.notify) | 服务市场商户确认订购通知 | `POST` | `/v3/alipay/open/servicemarket/order/notify` |
| [`alipay.open.servicemarket.order.query`](https://opendocs.alipay.com/open-v3/a59abd06_alipay.open.servicemarket.order.query) | 服务市场订单明细查询 | `GET` | `/v3/alipay/open/servicemarket/order/query` |
| [`alipay.open.servicemarket.order.reject`](https://opendocs.alipay.com/open-v3/f5d452ab_alipay.open.servicemarket.order.reject) | 服务商拒绝接单 | `POST` | `/v3/alipay/open/servicemarket/order/reject` |
| [`alipay.open.servicemarket.promotiontask.relation.query`](https://opendocs.alipay.com/open-v3/76898459_alipay.open.servicemarket.promotiontask.relation.query) | 服务商推广关系查询 | `GET` | `/v3/alipay/open/servicemarket/promotiontask/relation/query` |
| [`alipay.open.sp.blueseaactivity.create`](https://opendocs.alipay.com/open-v3/a974b7ae_alipay.open.sp.blueseaactivity.create) | 服务商返佣活动报名申请 | `POST` | `/v3/alipay/open/sp/blueseaactivity/create` |
| [`alipay.open.sp.blueseaactivity.modify`](https://opendocs.alipay.com/open-v3/8cbcb343_alipay.open.sp.blueseaactivity.modify) | 服务商返佣活动申请单修改 | `POST` | `/v3/alipay/open/sp/blueseaactivity/modify` |
| [`alipay.open.sp.blueseaactivity.query`](https://opendocs.alipay.com/open-v3/4ed0241a_alipay.open.sp.blueseaactivity.query) | 服务商返佣活动申请单详情查询 | `GET` | `/v3/alipay/open/sp/blueseaactivity/query` |
| [`alipay.open.sp.image.upload`](https://opendocs.alipay.com/open-v3/59220073_alipay.open.sp.image.upload) | 图片上传接口 | `POST` | `/v3/alipay/open/sp/image/upload` |
| [`alipay.open.sp.isv.bill.sync`](https://opendocs.alipay.com/apis/05vidn) | 服务商账单数据回传 | `POST` | `/v3/alipay/open/sp/isv/bill/sync` |
| [`alipay.open.sp.isv.signauth.create`](https://opendocs.alipay.com/open-v3/e4f4b975_alipay.open.sp.isv.signauth.create) | 创建签约授权一体化订单 | `POST` | `/v3/alipay/open/sp/isv/signauth/create` |
| [`alipay.open.sp.isv.signauth.query`](https://opendocs.alipay.com/open-v3/a322ed32_alipay.open.sp.isv.signauth.query) | 查询签约授权一体化订单状态详情 | `GET` | `/v3/alipay/open/sp/isv/signauth/query` |
| [`alipay.open.sp.leads.expand.create`](https://opendocs.alipay.com/apis/05vh1e) | 在线写入商机库 | `POST` | `/v3/alipay/open/sp/leads/expand/create` |
| [`alipay.open.sp.merchant.inconsistent.approve`](https://opendocs.alipay.com/apis/05v8ke) | 非同人商家进件资料审批 | `POST` | `/v3/alipay/open/sp/merchant/inconsistent/approve` |
| [`alipay.open.sp.operation.apply`](https://opendocs.alipay.com/open-v3/a8869af2_alipay.open.sp.operation.apply?scene=common&pathHash=6c69600a) | 向商户发起代运营操作 | `POST` | `/v3/alipay/open/sp/operation/apply` |
| [`alipay.open.sp.operation.qrcode.query`](https://opendocs.alipay.com/open-v3/8245f76c_alipay.open.sp.operation.qrcode.query?scene=common&pathHash=c82f7840) | 查询代运营授权或者账号绑定二维码 | `GET` | `/v3/alipay/open/sp/operation/qrcode/query` |
| [`alipay.open.sp.operation.result.query`](https://opendocs.alipay.com/open-v3/229831e2_alipay.open.sp.operation.result.query?scene=common&pathHash=4c2b6906) | 查询服务商代运营操作结果 | `GET` | `/v3/alipay/open/sp/operation/result/query` |
| [`alipay.open.sp.oppor.detail.query`](https://opendocs.alipay.com/apis/05vn6l) | 商机作业详情查询 | `GET` | `/v3/alipay/open/sp/oppor/detail/query` |
| [`alipay.open.sp.oppor.page.query`](https://opendocs.alipay.com/apis/05vdai) | 商机作业列表查询 | `GET` | `/v3/alipay/open/sp/oppor/page/query` |
| [`alipay.open.violation.violationdetail.query`](https://opendocs.alipay.com/open-v3/cd81d87b_alipay.open.violation.violationdetail.query?scene=common&pathHash=e3dfda29) | 违规记录详情查询 | `GET` | `/v3/alipay/open/violation/violationdetail/query` |
| [`alipay.open.violation.violationevent.batchquery`](https://opendocs.alipay.com/open-v3/529a02eb_alipay.open.violation.violationevent.batchquery?scene=common&pathHash=67cab0ec) | 违规记录列表查询 | `POST` | `/v3/alipay/open/violation/violationevent/batchquery` |

### `alipay.pcredit`（6）

| 接口 | 名称 | 方法 | 路径 |
| --- | --- | --- | --- |
| [`alipay.pcredit.huabei.auth.accumulation.query`](https://opendocs.alipay.com/pre-apis/05ge5u) | 花芝轻会员周期累计数据查询 | `GET` | `/v3/alipay/pcredit/huabei/auth/accumulation/query` |
| [`alipay.pcredit.huabei.auth.agreement.close`](https://opendocs.alipay.com/pre-apis/05gb9o) | 花芝协议关闭接口 | `PUT` | `/v3/alipay/pcredit/huabei/auth/agreement/close` |
| [`alipay.pcredit.huabei.auth.agreement.query`](https://opendocs.alipay.com/pre-apis/05fign) | 花芝轻会员协议查询接口 | `GET` | `/v3/alipay/pcredit/huabei/auth/agreement/query` |
| [`alipay.pcredit.huabei.auth.order.query`](https://opendocs.alipay.com/pre-apis/05h84d) | 花呗先享订单查询接口 | `GET` | `/v3/alipay/pcredit/huabei/auth/order/query` |
| [`alipay.pcredit.huabei.auth.order.unfreeze`](https://opendocs.alipay.com/pre-apis/05gcoh) | 花呗先享解冻或解约接口 | `PUT` | `/v3/alipay/pcredit/huabei/auth/order/unfreeze` |
| [`alipay.pcredit.huabei.auth.settle.apply`](https://opendocs.alipay.com/apis/05ghz8) | 花芝轻会员结算申请 | `PUT` | `/v3/alipay/pcredit/huabei/auth/settle/apply` |

### `alipay.security`（3）

| 接口 | 名称 | 方法 | 路径 |
| --- | --- | --- | --- |
| [`alipay.security.prod.sss.query`](https://opendocs.alipay.com/apis/0557cf) | 金柚测试3.0接口 | `POST` | `/v3/alipay/security/prod/sss/query/{bbb}` |
| [`alipay.security.risk.customerrisk.send`](https://opendocs.alipay.com/open-v3/cfacc9de_alipay.security.risk.customerrisk.send) | 商户数据同步 | `POST` | `/v3/alipay/security/risk/customerrisk/send` |
| [`alipay.security.risk.verifyidentity.miniapp.confirm`](https://opendocs.alipay.com/apis/05vidu) | 小程序核验服务结果确认接口 | `POST` | `/v3/alipay/security/risk/verifyidentity/miniapp/confirm` |

### `alipay.social`（1）

| 接口 | 名称 | 方法 | 路径 |
| --- | --- | --- | --- |
| [`alipay.social.base.contentlib.standardcontent.batchquery`](https://opendocs.alipay.com/apis/05wye1) | 内容批量查询接口 | `POST` | `/v3/alipay/social/base/contentlib/standardcontent/batchquery` |

### `alipay.system`（1）

| 接口 | 名称 | 方法 | 路径 |
| --- | --- | --- | --- |
| [`alipay.system.oauth.token`](https://opendocs.alipay.com/open-v3/a9cc8a2c_alipay.system.oauth.token) | 换取授权访问令牌 | `POST` | `/v3/alipay/system/oauth/token` |

### `alipay.trade`（27）

SDK 按基础交易与 `order`、`refund`、`batch`、`royalty`、`settle`、`customs` 子包组织；包路径、方法与类型见[交易目录](../services/trade/README.md)。

| 接口 | 名称 | 方法 | 路径 |
| --- | --- | --- | --- |
| [`alipay.trade.advance.consult`](https://opendocs.alipay.com/open-v3/11cb1f6e_alipay.trade.advance.consult) | 订单咨询服务 | `GET` | `/v3/alipay/trade/advance/consult` |
| [`alipay.trade.batch.settle`](https://opendocs.alipay.com/pre-apis/05vfpn) | 批量结算 | `POST` | `/v3/alipay/trade/batch/settle` |
| [`alipay.trade.batch.settle.query`](https://opendocs.alipay.com/pre-apis/05vel6) | 批量结算查询 | `POST` | `/v3/alipay/trade/batch/settle/query` |
| [`alipay.trade.batch.transfer`](https://opendocs.alipay.com/pre-apis/05vd9z) | 请款分账 | `POST` | `/v3/alipay/trade/batch/transfer` |
| [`alipay.trade.batch.transfer.query`](https://opendocs.alipay.com/pre-apis/05vsbp) | 请款分账查询接口 | `GET` | `/v3/alipay/trade/batch/transfer/query` |
| [`alipay.trade.cancel`](https://opendocs.alipay.com/open-v3/cd7d54d2_alipay.trade.cancel?scene=common&pathHash=cb83add6) | 统一收单交易撤销接口 | `POST` | `/v3/alipay/trade/cancel` |
| [`alipay.trade.close`](https://opendocs.alipay.com/open-v3/48ea518b_alipay.trade.close?scene=common&pathHash=2c6cdbea) | 统一收单交易关闭接口 | `POST` | `/v3/alipay/trade/close` |
| [`alipay.trade.create`](https://opendocs.alipay.com/open-v3/bf4eae5e_alipay.trade.create) | 统一收单交易创建接口 | `POST` | `/v3/alipay/trade/create` |
| [`alipay.trade.customs.declare`](https://opendocs.alipay.com/pre-apis/05wyhb) | 统一收单报关接口 | `POST` | `/v3/alipay/trade/customs/declare` |
| [`alipay.trade.customs.query`](https://opendocs.alipay.com/pre-apis/05wyjm) | 查询报关详细信息 | `GET` | `/v3/alipay/trade/customs/query` |
| [`alipay.trade.fastpay.refund.query`](https://opendocs.alipay.com/open-v3/b9ef37bd_alipay.trade.fastpay.refund.query?scene=common&pathHash=f071344e) | 统一收单交易退款查询 | `POST` | `/v3/alipay/trade/fastpay/refund/query` |
| [`alipay.trade.order.pay`](https://opendocs.alipay.com/open-v3/c2751e3b_alipay.trade.order.pay) | 统一收单交易订单支付接口 | `POST` | `/v3/alipay/trade/order/pay` |
| [`alipay.trade.order.settle`](https://opendocs.alipay.com/open-v3/34402227_alipay.trade.order.settle) | 统一收单交易结算接口 | `POST` | `/v3/alipay/trade/order/settle` |
| [`alipay.trade.order.settle.query`](https://opendocs.alipay.com/open-v3/3af46ab7_alipay.trade.order.settle.query) | 交易分账查询接口 | `GET` | `/v3/alipay/trade/order/settle/query` |
| [`alipay.trade.orderinfo.sync`](https://opendocs.alipay.com/open-v3/f8b82729_alipay.trade.orderinfo.sync) | 支付宝订单信息同步接口 | `POST` | `/v3/alipay/trade/orderinfo/sync` |
| [`alipay.trade.overseas.settle`](https://opendocs.alipay.com/apis/05v7c4) | 统一境内收单跨境结算接口 | `POST` | `/v3/alipay/trade/overseas/settle` |
| [`alipay.trade.pay`](https://opendocs.alipay.com/open-v3/08c7f9f8_alipay.trade.pay) | 统一收单交易支付接口 | `POST` | `/v3/alipay/trade/pay` |
| [`alipay.trade.precreate`](https://opendocs.alipay.com/open-v3/fa0c2141_alipay.trade.precreate) | 统一收单线下交易预创建 | `POST` | `/v3/alipay/trade/precreate` |
| [`alipay.trade.query`](https://opendocs.alipay.com/open-v3/d33bf036_alipay.trade.query) | 统一收单交易查询 | `POST` | `/v3/alipay/trade/query` |
| [`alipay.trade.refund`](https://opendocs.alipay.com/open-v3/6b16d4a2_alipay.trade.refund?scene=common&pathHash=00ab12b7) | 统一收单交易退款接口 | `POST` | `/v3/alipay/trade/refund` |
| [`alipay.trade.refund.apply`](https://opendocs.alipay.com/apis/05wbvf) | 异步退款受理接口 | `POST` | `/v3/alipay/trade/refund/apply` |
| [`alipay.trade.royalty.rate.query`](https://opendocs.alipay.com/open-v3/d9aece9f_alipay.trade.royalty.rate.query) | 分账比例查询 | `GET` | `/v3/alipay/trade/royalty/rate/query` |
| [`alipay.trade.royalty.relation.batchquery`](https://opendocs.alipay.com/open-v3/1a11475c_alipay.trade.royalty.relation.batchquery) | 分账关系查询 | `POST` | `/v3/alipay/trade/royalty/relation/batchquery` |
| [`alipay.trade.royalty.relation.bind`](https://opendocs.alipay.com/open-v3/e701637d_alipay.trade.royalty.relation.bind) | 分账关系绑定 | `POST` | `/v3/alipay/trade/royalty/relation/bind` |
| [`alipay.trade.royalty.relation.unbind`](https://opendocs.alipay.com/open-v3/302cd1f8_alipay.trade.royalty.relation.unbind) | 分账关系解绑 | `POST` | `/v3/alipay/trade/royalty/relation/unbind` |
| [`alipay.trade.settle.confirm`](https://opendocs.alipay.com/apis/0hp5md) | 统一收单确认结算接口 | `POST` | `/v3/alipay/trade/settle/confirm` |
| [`alipay.trade.settle.receivables.query`](https://opendocs.alipay.com/pre-apis/05vlu7) | 待结算查询 | `POST` | `/v3/alipay/trade/settle/receivables/query` |

### `alipay.user`（18）

| 接口 | 名称 | 方法 | 路径 |
| --- | --- | --- | --- |
| [`alipay.user.agreement.executionplan.modify`](https://opendocs.alipay.com/open-v3/21227669_alipay.user.agreement.executionplan.modify?scene=common&pathHash=89395d7c) | 周期性扣款协议执行计划修改接口 | `POST` | `/v3/alipay/user/agreement/executionplan/modify` |
| [`alipay.user.agreement.migrate`](https://opendocs.alipay.com/open-v3/a611f0a3_alipay.user.agreement.migrate) | 代扣协议迁移 | `POST` | `/v3/alipay/user/agreement/migrate` |
| [`alipay.user.agreement.permission.create`](https://opendocs.alipay.com/open-v3/87401d30_alipay.user.agreement.permission.create) | 支付宝协议支付许可创建、服务变更接口 | `POST` | `/v3/alipay/user/agreement/permission/create` |
| [`alipay.user.agreement.query`](https://opendocs.alipay.com/open-v3/4dd54af2_alipay.user.agreement.query) | 支付宝个人代扣协议查询接口 | `GET` | `/v3/alipay/user/agreement/query` |
| [`alipay.user.agreement.transfer`](https://opendocs.alipay.com/open-v3/360a356d_alipay.user.agreement.transfer) | 协议由普通通用代扣协议产品转移到周期扣协议产品 | `POST` | `/v3/alipay/user/agreement/transfer` |
| [`alipay.user.agreement.unsign`](https://opendocs.alipay.com/open-v3/b16bea7c_alipay.user.agreement.unsign) | 支付宝个人代扣协议解约接口 | `POST` | `/v3/alipay/user/agreement/unsign` |
| [`alipay.user.agreement.userverify.apply`](https://opendocs.alipay.com/open-v3/df28f556_alipay.user.agreement.userverify.apply) | 协议用户一致性验证 | `POST` | `/v3/alipay/user/agreement/userverify/apply` |
| [`alipay.user.agreement.userverify.query`](https://opendocs.alipay.com/open-v3/a514c337_alipay.user.agreement.userverify.query) | 用户协议鉴权查询结果 | `GET` | `/v3/alipay/user/agreement/userverify/query` |
| [`alipay.user.alipaypoint.budgetlib.query`](https://opendocs.alipay.com/apis/05wmz4) | 查询集分宝预算库详情 | `GET` | `/v3/alipay/user/alipaypoint/budgetlib/query` |
| [`alipay.user.alipaypoint.send`](https://opendocs.alipay.com/apis/05wvr1) | 集分宝发放接口 | `POST` | `/v3/alipay/user/alipaypoint/send` |
| [`alipay.user.certdoc.certverify.consult`](https://opendocs.alipay.com/apis/05wvw5) | 实名证件信息比对验证咨询 | `GET` | `/v3/alipay/user/certdoc/certverify/consult` |
| [`alipay.user.certdoc.certverify.preconsult`](https://opendocs.alipay.com/apis/05wfau) | 实名证件信息比对验证预咨询 | `POST` | `/v3/alipay/user/certdoc/certverify/preconsult` |
| [`alipay.user.certify.open.initialize`](https://opendocs.alipay.com/open-v3/013ee3e1_alipay.user.certify.open.initialize) | 身份认证初始化服务 | `POST` | `/v3/alipay/user/certify/open/initialize` |
| [`alipay.user.certify.open.query`](https://opendocs.alipay.com/open-v3/c166c117_alipay.user.certify.open.query) | 身份认证记录查询 | `GET` | `/v3/alipay/user/certify/open/query` |
| [`alipay.user.info.share`](https://opendocs.alipay.com/open-v3/2f4a8991_alipay.user.info.share) | 支付宝会员授权信息查询接口 | `POST` | `/v3/alipay/user/info/share` |
| [`alipay.user.twostage.common.use`](https://opendocs.alipay.com/apis/05xim7) | 通用当面付二阶段接口 | `POST` | `/v3/alipay/user/twostage/common/use` |
| [`alipay.user.twostage.indirect.use`](https://opendocs.alipay.com/pre-apis/05x9jk) | 双联二阶段 | `POST` | `/v3/alipay/user/twostage/indirect/use` |
| [`alipay.user.userinfo.share`](https://opendocs.alipay.com/open-v3/a7417208_alipay.user.userinfo.share) | 支付宝钱包用户信息共享 | `POST` | `/v3/alipay/user/userinfo/share` |

### `ant.merchant`（25）

| 接口 | 名称 | 方法 | 路径 |
| --- | --- | --- | --- |
| [`ant.merchant.expand.apprecommend.account.create`](https://opendocs.alipay.com/open-v3/32d136f0_ant.merchant.expand.apprecommend.account.create) | 关联账号和小程序 | `POST` | `/v3/ant/merchant/expand/apprecommend/account/create` |
| [`ant.merchant.expand.apprecommend.account.delete`](https://opendocs.alipay.com/open-v3/130a15de_ant.merchant.expand.apprecommend.account.delete) | 取消关联账号和小程序 | `DELETE` | `/v3/ant/merchant/expand/apprecommend/account/delete` |
| [`ant.merchant.expand.apprecommend.account.query`](https://opendocs.alipay.com/open-v3/cb07737c_ant.merchant.expand.apprecommend.account.query) | 查询已关联指定APP的账号列表 | `GET` | `/v3/ant/merchant/expand/apprecommend/account/query` |
| [`ant.merchant.expand.apprecommend.available.query`](https://opendocs.alipay.com/open-v3/fc63d6c8_ant.merchant.expand.apprecommend.available.query) | 查询可关联APP的账号列表 | `GET` | `/v3/ant/merchant/expand/apprecommend/available/query` |
| [`ant.merchant.expand.indirect.image.upload`](https://opendocs.alipay.com/open-v3/af8e3a42_ant.merchant.expand.indirect.image.upload) | 图片上传 | `POST` | `/v3/ant/merchant/indirect/image` |
| [`ant.merchant.expand.info.query`](https://opendocs.alipay.com/apis/05vb3x) | 蚂蚁商户信息查询接口 | `GET` | `/v3/ant/merchant/expand/info/query` |
| [`ant.merchant.expand.item.create`](https://opendocs.alipay.com/open-v3/bf0fb8ad_ant.merchant.expand.item.create) | 商品创建接口 | `POST` | `/v3/ant/merchant/expand/item/create` |
| [`ant.merchant.expand.item.delete`](https://opendocs.alipay.com/open-v3/ac2de50f_ant.merchant.expand.item.delete) | 商品删除接口 | `DELETE` | `/v3/ant/merchant/expand/item/delete` |
| [`ant.merchant.expand.item.modify`](https://opendocs.alipay.com/open-v3/60d35f53_ant.merchant.expand.item.modify) | 商品修改接口 | `POST` | `/v3/ant/merchant/expand/item/modify` |
| [`ant.merchant.expand.item.open.batchquery`](https://opendocs.alipay.com/pre-apis/05vjno) | 批量查询商品接口 | `POST` | `/v3/ant/merchant/expand/item/open/batchquery` |
| [`ant.merchant.expand.item.open.create`](https://opendocs.alipay.com/pre-apis/05vn67) | 创建商品接口 | `POST` | `/v3/ant/merchant/expand/item/open/create` |
| [`ant.merchant.expand.item.open.delete`](https://opendocs.alipay.com/pre-apis/05v7ca) | 删除商品接口 | `DELETE` | `/v3/ant/merchant/expand/item/open/delete` |
| [`ant.merchant.expand.item.open.modify`](https://opendocs.alipay.com/pre-apis/05vda6) | 修改商品接口 | `POST` | `/v3/ant/merchant/expand/item/open/modify` |
| [`ant.merchant.expand.item.open.query`](https://opendocs.alipay.com/pre-apis/05vlue) | 查询商品接口 | `GET` | `/v3/ant/merchant/expand/item/open/query` |
| [`ant.merchant.expand.item.query`](https://opendocs.alipay.com/open-v3/aad82813_ant.merchant.expand.item.query) | 商品查询接口 | `GET` | `/v3/ant/merchant/expand/item/query` |
| [`ant.merchant.expand.item.status.modify`](https://opendocs.alipay.com/open-v3/5cb8e04e_ant.merchant.expand.item.status.modify) | 商品状态修改接口 | `POST` | `/v3/ant/merchant/expand/item/status/modify` |
| [`ant.merchant.expand.mcc.query`](https://opendocs.alipay.com/open-v3/b00ab2d2_ant.merchant.expand.mcc.query) | 商户mcc信息查询 | `GET` | `/v3/ant/merchant/mcc/query` |
| [`ant.merchant.expand.order.query`](https://opendocs.alipay.com/open-v3/dd6e9adb_ant.merchant.expand.order.query) | 商户申请单查询 | `GET` | `/v3/ant/merchant/order/{order_id}` |
| [`ant.merchant.expand.shop.close`](https://opendocs.alipay.com/open-v3/4950631b_ant.merchant.expand.shop.close) | 蚂蚁门店关闭 | `PATCH` | `/v3/ant/merchant/shop/close` |
| [`ant.merchant.expand.shop.consult`](https://opendocs.alipay.com/apis/05vda1) | 蚂蚁店铺创建咨询 | `POST` | `/v3/ant/merchant/expand/shop/consult` |
| [`ant.merchant.expand.shop.create`](https://opendocs.alipay.com/open-v3/99b91969_ant.merchant.expand.shop.create) | 蚂蚁门店创建 | `POST` | `/v3/ant/merchant/shop` |
| [`ant.merchant.expand.shop.modify`](https://opendocs.alipay.com/open-v3/b1a37213_ant.merchant.expand.shop.modify) | 修改蚂蚁门店 | `PATCH` | `/v3/ant/merchant/shop` |
| [`ant.merchant.expand.shop.page.query`](https://opendocs.alipay.com/open-v3/fe443839_ant.merchant.expand.shop.page.query) | 店铺分页查询接口 | `GET` | `/v3/ant/merchant/shop/pagequery` |
| [`ant.merchant.expand.shop.query`](https://opendocs.alipay.com/open-v3/b80b5948_ant.merchant.expand.shop.query) | 门店查询接口 | `GET` | `/v3/ant/merchant/shop` |
| [`ant.merchant.expand.shop.receiptaccount.save`](https://opendocs.alipay.com/open-v3/74728d59_ant.merchant.expand.shop.receiptaccount.save) | 店铺增加收单账号 | `POST` | `/v3/ant/merchant/expand/shop/receiptaccount/save` |

### `datadigital.fincloud`（10）

| 接口 | 名称 | 方法 | 路径 |
| --- | --- | --- | --- |
| [`datadigital.fincloud.generalsaas.face.certify.initialize`](https://opendocs.alipay.com/open-v3/645bdc9c_datadigital.fincloud.generalsaas.face.certify.initialize) | 跳转支付宝人脸核身初始化 | `POST` | `/v3/datadigital/fincloud/generalsaas/face/certify/initialize` |
| [`datadigital.fincloud.generalsaas.face.certify.query`](https://opendocs.alipay.com/open-v3/10a1ccfd_datadigital.fincloud.generalsaas.face.certify.query) | 跳转支付宝人脸核身查询记录 | `GET` | `/v3/datadigital/fincloud/generalsaas/face/certify/query` |
| [`datadigital.fincloud.generalsaas.face.certify.verify`](https://opendocs.alipay.com/open-v3/d3f37a88_datadigital.fincloud.generalsaas.face.certify.verify) | 跳转支付宝人脸核身开始认证 | `POST` | `/v3/datadigital/fincloud/generalsaas/face/certify/verify` |
| [`datadigital.fincloud.generalsaas.face.check.initialize`](https://opendocs.alipay.com/open-v3/503aa13e_datadigital.fincloud.generalsaas.face.check.initialize) | 活体检测初始化 | `POST` | `/v3/datadigital/fincloud/generalsaas/face/check/initialize` |
| [`datadigital.fincloud.generalsaas.face.check.query`](https://opendocs.alipay.com/open-v3/4b93abbc_datadigital.fincloud.generalsaas.face.check.query) | 活体检测结果查询 | `GET` | `/v3/datadigital/fincloud/generalsaas/face/check/query` |
| [`datadigital.fincloud.generalsaas.face.source.certify`](https://opendocs.alipay.com/open-v3/21280890_datadigital.fincloud.generalsaas.face.source.certify) | 纯服务端人脸核身 | `POST` | `/v3/datadigital/fincloud/generalsaas/face/source/certify` |
| [`datadigital.fincloud.generalsaas.face.verification.initialize`](https://opendocs.alipay.com/open-v3/34fb02dc_datadigital.fincloud.generalsaas.face.verification.initialize) | 人脸核身初始化 | `POST` | `/v3/datadigital/fincloud/generalsaas/face/verification/initialize` |
| [`datadigital.fincloud.generalsaas.face.verification.query`](https://opendocs.alipay.com/open-v3/1ef1e2ed_datadigital.fincloud.generalsaas.face.verification.query) | 人脸核身结果查询 | `GET` | `/v3/datadigital/fincloud/generalsaas/face/verification/query` |
| [`datadigital.fincloud.generalsaas.ocr.mobile.initialize`](https://opendocs.alipay.com/open-v3/74ea0784_datadigital.fincloud.generalsaas.ocr.mobile.initialize) | App端OCR初始化 | `POST` | `/v3/datadigital/fincloud/generalsaas/ocr/mobile/initialize` |
| [`datadigital.fincloud.generalsaas.ocr.server.detect`](https://opendocs.alipay.com/open-v3/a6f20b9d_datadigital.fincloud.generalsaas.ocr.server.detect) | 服务端OCR | `POST` | `/v3/datadigital/fincloud/generalsaas/ocr/server/detect` |

### `grandsecurity.bizrisks`（1）

| 接口 | 名称 | 方法 | 路径 |
| --- | --- | --- | --- |
| [`grandsecurity.bizrisks.fact.check`](https://opendocs.alipay.com/pre-apis/0jjgu4) | 事实核查服务 | `POST` | `/v3/stream/grandsecurity/bizrisks/fact/check` |

### `monitor.heartbeat`（1）

| 接口 | 名称 | 方法 | 路径 |
| --- | --- | --- | --- |
| [`monitor.heartbeat.syn`](https://opendocs.alipay.com/apis/05vc16) | 验签接口 | `POST` | `/v3/monitor/heartbeat/syn` |

### `zhima.credit`（13）

| 接口 | 名称 | 方法 | 路径 |
| --- | --- | --- | --- |
| [`zhima.credit.payafteruse.creditagreement.query`](https://opendocs.alipay.com/open-v3/eaff9140_zhima.credit.payafteruse.creditagreement.query) | 查询服务开通/授权信息 | `GET` | `/v3/zhima/credit/payafteruse/creditagreement/query` |
| [`zhima.credit.payafteruse.creditagreement.transfer`](https://opendocs.alipay.com/apis/05vdam) | 芝麻信用服务授权 | `POST` | `/v3/zhima/credit/payafteruse/creditagreement/transfer` |
| [`zhima.credit.payafteruse.creditbizorder.finish`](https://opendocs.alipay.com/open-v3/7972514c_zhima.credit.payafteruse.creditbizorder.finish?scene=common&pathHash=ba647372) | 结束信用服务订单 | `POST` | `/v3/zhima/credit/payafteruse/creditbizorder/finish` |
| [`zhima.credit.payafteruse.creditbizorder.order`](https://opendocs.alipay.com/open-v3/51ba4ec1_zhima.credit.payafteruse.creditbizorder.order) | 芝麻信用服务下单（免用户确认场景） | `POST` | `/v3/zhima/credit/payafteruse/creditbizorder/order` |
| [`zhima.credit.payafteruse.creditbizorder.query`](https://opendocs.alipay.com/open-v3/0f11f105_zhima.credit.payafteruse.creditbizorder.query?scene=common&pathHash=203d9491) | 信用服务订单查询 | `GET` | `/v3/zhima/credit/payafteruse/creditbizorder/query` |
| [`zhima.credit.pe.zmgo.agreement.query`](https://opendocs.alipay.com/open-v3/849a96d9_zhima.credit.pe.zmgo.agreement.query) | 芝麻Go协议查询接口 | `GET` | `/v3/zhima/credit/pe/zmgo/agreement/query` |
| [`zhima.credit.pe.zmgo.agreement.unsign`](https://opendocs.alipay.com/open-v3/4d77e884_zhima.credit.pe.zmgo.agreement.unsign) | 芝麻GO协议解约 | `POST` | `/v3/zhima/credit/pe/zmgo/agreement/unsign` |
| [`zhima.credit.pe.zmgo.bizopt.close`](https://opendocs.alipay.com/apis/05vluc) | 芝麻GO签约关单 | `POST` | `/v3/zhima/credit/pe/zmgo/bizopt/close` |
| [`zhima.credit.pe.zmgo.cumulation.sync`](https://opendocs.alipay.com/apis/05v7c8) | 芝麻go用户数据回传 | `POST` | `/v3/zhima/credit/pe/zmgo/cumulation/sync` |
| [`zhima.credit.pe.zmgo.preorder.create`](https://opendocs.alipay.com/open-v3/3a385c23_zhima.credit.pe.zmgo.preorder.create) | 芝麻GO签约预创单 | `PUT` | `/v3/zhima/credit/pe/zmgo/preorder/create` |
| [`zhima.credit.pe.zmgo.settle.apply`](https://opendocs.alipay.com/open-v3/3a0505ce_zhima.credit.pe.zmgo.settle.apply) | 芝麻GO结算申请 | `POST` | `/v3/zhima/credit/pe/zmgo/settle/apply` |
| [`zhima.credit.pe.zmgo.settle.refund`](https://opendocs.alipay.com/open-v3/a66ab3b0_zhima.credit.pe.zmgo.settle.refund) | 芝麻GO结算退款接口 | `POST` | `/v3/zhima/credit/pe/zmgo/settle/refund` |
| [`zhima.credit.pe.zmgo.settle.unfreeze`](https://opendocs.alipay.com/apis/05gzu9) | 芝麻GO解冻接口 | `PUT` | `/v3/zhima/credit/pe/zmgo/settle/unfreeze` |

### `zhima.customer`（6）

| 接口 | 名称 | 方法 | 路径 |
| --- | --- | --- | --- |
| [`zhima.customer.jobworth.authentication.query`](https://opendocs.alipay.com/open-v3/d25db704_zhima.customer.jobworth.authentication.query) | 职得身份认证查询接口 | `GET` | `/v3/zhima/customer/jobworth/authentication` |
| [`zhima.customer.jobworth.cloudresume.query`](https://opendocs.alipay.com/apis/05wukc) | 智能简历基本数据查询 | `GET` | `/v3/zhima/customer/jobworth/cloudresume/query` |
| [`zhima.customer.jobworth.info.query`](https://opendocs.alipay.com/apis/05ttl7) | 职得工作证信息查询 | `GET` | `/v3/zhima/customer/jobworth/info/query` |
| [`zhima.customer.jobworth.jobdata.add`](https://opendocs.alipay.com/apis/05wyeb) | 工作证职业数据回流接口 | `POST` | `/v3/zhima/customer/jobworth/jobdata/add` |
| [`zhima.customer.jobworth.picture.upload`](https://opendocs.alipay.com/apis/05wulq) | 职得工作证图片上传 | `POST` | `/v3/zhima/customer/jobworth/picture/upload` |
| [`zhima.customer.jobworth.position.add`](https://opendocs.alipay.com/apis/05xa7n) | 职得工作证岗位收集 | `POST` | `/v3/zhima/customer/jobworth/position/add` |

### `zhima.merchant`（6）

| 接口 | 名称 | 方法 | 路径 |
| --- | --- | --- | --- |
| [`zhima.merchant.subsidiaries.apply`](https://opendocs.alipay.com/apis/05ygh0) | 子商户信息添加 | `POST` | `/v3/zhima/merchant/subsidiaries/apply` |
| [`zhima.merchant.subsidiaries.close`](https://opendocs.alipay.com/apis/05yopn) | 关闭子商户信息工单 | `POST` | `/v3/zhima/merchant/subsidiaries/close` |
| [`zhima.merchant.zmgo.cumulate.query`](https://opendocs.alipay.com/open-v3/62aec983_zhima.merchant.zmgo.cumulate.query) | 商家芝麻GO累计数据查询接口 | `GET` | `/v3/zhima/merchant/zmgo/cumulate/query` |
| [`zhima.merchant.zmgo.cumulate.sync`](https://opendocs.alipay.com/open-v3/153c048c_zhima.merchant.zmgo.cumulate.sync) | 商家芝麻GO累计数据回传接口 | `POST` | `/v3/zhima/merchant/zmgo/cumulate/sync` |
| [`zhima.merchant.zmgo.template.create`](https://opendocs.alipay.com/open-v3/b0eba80a_zhima.merchant.zmgo.template.create) | 商户创建芝麻GO模板接口 | `POST` | `/v3/zhima/merchant/zmgo/template/create` |
| [`zhima.merchant.zmgo.template.query`](https://opendocs.alipay.com/open-v3/1cbec0d3_zhima.merchant.zmgo.template.query) | 芝麻GO模板查询 | `GET` | `/v3/zhima/merchant/zmgo/template/query` |

### `zoloz.authentication`（5）

| 接口 | 名称 | 方法 | 路径 |
| --- | --- | --- | --- |
| [`zoloz.authentication.customer.facemanage.create`](https://opendocs.alipay.com/pre-apis/05vda2) | 热脸入库 | `POST` | `/v3/zoloz/authentication/customer/facemanage/create` |
| [`zoloz.authentication.customer.facemanage.delete`](https://opendocs.alipay.com/pre-apis/05vh16) | 热脸出库 | `POST` | `/v3/zoloz/authentication/customer/facemanage/delete` |
| [`zoloz.authentication.customer.ftoken.query`](https://opendocs.alipay.com/open-v3/ebd5a135_zoloz.authentication.customer.ftoken.query) | 查询刷脸结果信息接口 | `POST` | `/v3/zoloz/authentication/customer/ftoken/query` |
| [`zoloz.authentication.customer.smilepay.initialize`](https://opendocs.alipay.com/pre-apis/05vfpu) | 人脸刷脸付初始化 | `POST` | `/v3/zoloz/authentication/customer/smilepay/initialize` |
| [`zoloz.authentication.smilepay.initialize`](https://opendocs.alipay.com/open-v3/1177dbeb_zoloz.authentication.smilepay.initialize) | 刷脸支付初始化 | `POST` | `/v3/zoloz/authentication/smilepay/initialize` |

## 快照外接口（35 项，V3 契约未确认）

核验日期：2026-11-05。以下接口未出现在本目录引用的 V3 OpenAPI 快照中，也未计入上方 REST 和页面/唤端接口统计。本次仅补充目录记录，SDK 未实现这些接口；不能按接口名推导 HTTP 方法或 `/v3/...` 路由，也不能以传统网关服务端调用代替 V3。

### OpenID 转换（6 项，已查阅接口全文）

以下官方全文的公共参数包含 `method`、`version=1.0`、`biz_content`，请求示例使用 `https://openapi.alipay.com/gateway.do`。这些依据确认的是文档中的传统网关调用方式，不足以确认 V3 支持；不能仅凭 `version=1.0` 或未命中快照断言官方没有其他调用方式。

| 接口 | 名称 | 核验状态 |
| --- | --- | --- |
| [`alipay.open.app.openid.applyorder.create`](https://opendocs.alipay.com/common/0arwjk.md) | 创建openid批量转换工单 | 全文已核验；V3 契约未确认 |
| [`alipay.open.app.openid.applyorder.upload`](https://opendocs.alipay.com/common/0ar4ll.md) | 批量上传用户userid到openid转换工单 | 全文已核验；V3 契约未确认 |
| [`alipay.open.app.openid.applyorder.submit`](https://opendocs.alipay.com/common/0arxyn.md) | openid转换工单提交审核 | 全文已核验；V3 契约未确认 |
| [`alipay.open.app.openid.batchquery`](https://opendocs.alipay.com/common/0are0o.md) | 根据支付宝用户userid批量获取用户openid和unionid | 全文已核验；V3 契约未确认 |
| [`alipay.open.app.openid.applyorder.checkavailable`](https://opendocs.alipay.com/common/0arju6.md) | 检查转换工单审核状态 | 全文已核验；V3 契约未确认 |
| [`alipay.open.app.openid.applyorder.cancel`](https://opendocs.alipay.com/common/0are0p.md) | 取消openid转换工单 | 全文已核验；V3 契约未确认 |

### 官方敏感接口清单补录（29 项，仅核验清单列名）

以下名称与接口文档链接逐项取自官方[资金操作类敏感接口列表](https://opendocs.alipay.com/common/02kkuw.md)，该清单不标注各接口的 V3 支持情况。本次访问其接口文档链接未取得可核验的正文（空正文或站点 HTML 页面），因此只确认官方清单曾列出这些接口；现行可用性、参数、协议版本与 V3 契约均待确认。原始链接保留供后续核验，不据此提供调用实现。

#### 支付宝 `alipay.*`（16 项）

| 接口 | 名称 |
| --- | --- |
| [`alipay.commerce.lottery.present.send`](https://opendocs.alipay.com/apis/009zvg.md) | 商家给用户赠送彩票 |
| [`alipay.commerce.transport.nfccard.send`](https://docs.open.alipay.com/api_1/alipay.commerce.transport.nfccard.send) | NFC用户卡信息同步 |
| [`alipay.fund.coupon.order.agreement.pay`](https://docs.open.alipay.com/api_5/alipay.fund.coupon.order.agreement.pay) | 红包协议支付接口 |
| [`alipay.fund.coupon.order.app.pay`](https://docs.open.alipay.com/api_5/alipay.fund.coupon.order.app.pay) | 红包无线支付接口 |
| [`alipay.fund.coupon.order.disburse`](https://docs.open.alipay.com/api_5/alipay.fund.coupon.order.disburse) | 红包打款接口 |
| [`alipay.fund.coupon.order.refund`](https://docs.open.alipay.com/api_5/alipay.fund.coupon.order.refund) | 红包退回接口 |
| [`alipay.marketing.campaign.discount.budget.append`](https://opendocs.alipay.com/apis/009zvw.md) | 营销立减活动预算追加 |
| [`alipay.marketing.campaign.discount.budget.create`](https://docs.open.alipay.com/api_5/alipay.marketing.campaign.discount.budget.create) | 营销立减活动预算创建 |
| [`alipay.marketing.cashlessvoucher.template.create`](https://opendocs.alipay.com/open/02aamd.md) | 无资金券模板创建接口 |
| [`alipay.marketing.cashlessvoucher.template.modify`](https://opendocs.alipay.com/open/02a92y.md) | 无资金券模板修改接口 |
| [`alipay.marketing.cashvoucher.template.create`](https://opendocs.alipay.com/open/02aamp.md) | 创建资金券模板 |
| [`alipay.marketing.coupon.template.create`](https://docs.open.alipay.com/api_5/alipay.marketing.coupon.template.create) | 创建红包模板 |
| [`alipay.marketing.voucher.send`](https://opendocs.alipay.com/open/02a6z4.md) | 发券接口 |
| [`alipay.pass.instance.add`](https://opendocs.alipay.com/open/02ailb.md) | 卡券实例发放接口 |
| [`alipay.pass.template.add`](https://opendocs.alipay.com/open/02aila.md) | 卡券模板创建接口 |
| [`alipay.pass.template.update`](https://opendocs.alipay.com/open/02ahjr.md) | 卡券模板更新接口 |

#### 口碑 `koubei.*`（10 项）

| 接口 | 名称 |
| --- | --- |
| [`koubei.marketing.campaign.activity.create`](https://docs.open.alipay.com/api_5/koubei.marketing.campaign.activity.create) | 活动创建接口 |
| [`koubei.marketing.tool.prizesend.auth`](https://docs.open.alipay.com/api_24/koubei.marketing.tool.prizesend.auth) | 发券授权 |
| [`koubei.trade.itemorder.buy`](https://docs.open.alipay.com/api_1/koubei.trade.itemorder.buy) | 口碑商品交易购买接口 |
| [`koubei.trade.itemorder.query`](https://docs.open.alipay.com/api_1/koubei.trade.itemorder.query) | 口碑商品交易查询接口 |
| [`koubei.trade.itemorder.refund`](https://docs.open.alipay.com/api_1/koubei.trade.itemorder.refund) | 口碑商品交易退货接口 |
| [`koubei.trade.order.precreate`](https://docs.open.alipay.com/api_1/koubei.trade.order.precreate) | 口碑订单预下单 |
| [`koubei.trade.ticket.ticketcode.cancel`](https://docs.open.alipay.com/api_1/koubei.trade.ticket.ticketcode.cancel) | 口碑凭证码撤销核销 |
| [`koubei.trade.ticket.ticketcode.delay`](https://docs.open.alipay.com/api_1/koubei.trade.ticket.ticketcode.delay) | 口碑凭证延期接口 |
| [`koubei.trade.ticket.ticketcode.query`](https://opendocs.alipay.com/apis/009zik.md) | 口碑凭证码查询 |
| [`koubei.trade.ticket.ticketcode.send`](https://docs.open.alipay.com/api_1/koubei.trade.ticket.ticketcode.send) | 码商发码成功回调接口 |

#### 网商银行 `mybank.*`（3 项）

| 接口 | 名称 |
| --- | --- |
| [`mybank.finance.yulibao.capital.purchase`](https://docs.open.alipay.com/api_38/mybank.finance.yulibao.capital.purchase) | 余利宝申购 |
| [`mybank.finance.yulibao.capital.ransom`](https://docs.open.alipay.com/api_38/mybank.finance.yulibao.capital.ransom) | 余利宝赎回 |
| [`mybank.payment.trade.order.create`](https://docs.open.alipay.com/api_1/mybank.payment.trade.order.create) | 网商银行全渠道收单业务订单创建 |
