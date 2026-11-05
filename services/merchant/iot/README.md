# iot 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/merchant/iot`。本包提供 3 个支付宝 V3 REST 操作。

通过 `iot.NewService(executor)` 复用 `*core.Client` 或最小 `core.Executor`。网络调用以 `context.Context` 为首参数，支持 `opts ...core.CallOption`，业务数据位于 `resp.Data`。

接口方法按单词边界使用大驼峰命名；请求、响应与专属模型位于对应业务包，共享模型从共同父包引用。

接入方式见[项目说明](../../../README.md)。默认测试离线执行。

返回[领域目录](../README.md)。

## 接口目录

| SDK 方法 | 官方接口 |
| --- | --- |
| [`IoTDeviceBind`](#api-iotdevicebind) | `alipay.merchant.iot.device.bind` |
| [`IoTDeviceQuery`](#api-iotdevicequery) | `alipay.merchant.iot.device.query` |
| [`IoTDeviceVerify`](#api-iotdeviceverify) | `alipay.merchant.iot.device.verify` |

## 使用示例

- [IoTDeviceBind](api_iot_iotdevicebind_example_test.go)
- [IoTDeviceQuery](api_iot_iotdevicequery_example_test.go)
- [IoTDeviceVerify](api_iot_iotdeviceverify_example_test.go)

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-iotdevicebind"></a>

### IoTDeviceBind — IoT设备绑定门店

- 官方接口：[`alipay.merchant.iot.device.bind`](https://opendocs.alipay.com/open-v3/c82d5f5f_alipay.merchant.iot.device.bind?scene=common&pathHash=0fb151a0)
- 路由：`POST /v3/alipay/merchant/iot/device/bind`
- 请求类型：[MerchantIoTDeviceBindRequest](iot_request.go#L4)
- 响应类型：[MerchantIoTDeviceBindResponse](iot_response.go#L8)

同步IoT设备、商户和门店的绑定关系

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `biz_tid` | `BizTid` | `*string` | JSON body | 未标注 | 设备 ID ，仅device_id_type 为 ID 时填写。 |
| `device_id_type` | `DeviceIdType` | `*string` | JSON body | 未标注 | 可选方式。枚举支持： * ID：表示使用 biztid 作为设备唯一识别标识， * SN：表示使用 supplier_id、device_sn联合作为设备唯一识别标识。 注意：由于不同机型的supplier_id不同，推荐使用 ID 。 |
| `device_sn` | `DeviceSn` | `*string` | JSON body | 未标注 | 设备序列号 ，device_id_type 为 SN 时填写。需配合supplier_id使用。 |
| `merchant_type` | `MerchantType` | `*string` | JSON body | 未标注 | 商户类型。枚举支持： * direct：直连商户； * indirect：间连商户。 |
| `pid` | `Pid` | `*string` | JSON body | 未标注 | 直连场景填写商户收单pid。间连场景填写smid绑定的pid（可通过代运营授权API获取）。 |
| `shop_id` | `ShopId` | `*string` | JSON body | 未标注 | 商户pid名下的门店 ID。可通过门店管理API获取，需注意请勿使用smid创建门店 ID。 |
| `smid` | `Smid` | `*string` | JSON body | 未标注 | 直连场景不填，间连场景填写商户收单smid |
| `supplier_id` | `SupplierId` | `*string` | JSON body | 未标注 | 设备供应商ID ，device_id_type 为 SN 时填写。需注意不同机型的供应商ID可能不同。 |

<a id="api-iotdevicequery"></a>

### IoTDeviceQuery — IoT设备绑定关系查询

- 官方接口：[`alipay.merchant.iot.device.query`](https://opendocs.alipay.com/open-v3/9b618feb_alipay.merchant.iot.device.query?scene=common&pathHash=9f128d89)
- 路由：`GET /v3/alipay/merchant/iot/device/query`
- 请求类型：[MerchantIoTDeviceQueryRequest](iot_request.go#L42)
- 响应类型：[MerchantIoTDeviceQueryResponse](iot_response.go#L11)

通过设备唯一标识查询设备绑定关系

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `device_id_type` | `DeviceIdType` | `*string` | query | 否 | 可选方式 [ID,SN]。ID-使用biztid作为设备唯一识别标识；SN-使用supplier_id、device_sn联合作为设备唯一识别标识。由于不同机型的supplier_id不同，推荐使用 ID 。 |
| `biz_tid` | `BizTid` | `*string` | query | 否 | 设备 ID ，device_id_type 为 ID 时填写。 |
| `supplier_id` | `SupplierId` | `*string` | query | 否 | 设备供应商ID ，device_id_type 为 SN 时填写。需注意不同机型的供应商ID可能不同。 |
| `device_sn` | `DeviceSn` | `*string` | query | 否 | 设备序列号 ，device_id_type 为 SN 时填写。需配合supplier_id使用。 |

<a id="api-iotdeviceverify"></a>

### IoTDeviceVerify — IoT设备绑定校验

- 官方接口：[`alipay.merchant.iot.device.verify`](https://opendocs.alipay.com/open-v3/4e0eced6_alipay.merchant.iot.device.verify?scene=common&pathHash=39ad307c)
- 路由：`POST /v3/alipay/merchant/iot/device/verify`
- 请求类型：[MerchantIoTDeviceVerifyRequest](iot_request.go#L24)
- 响应类型：[MerchantIoTDeviceVerifyResponse](iot_response.go#L23)

校验IoT设备和商户收单账号的对应关系

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `biz_tid` | `BizTid` | `*string` | JSON body | 未标注 | 设备 ID ，device_id_type 为 ID 时填写。 |
| `device_id_type` | `DeviceIdType` | `*string` | JSON body | 未标注 | 可选方式 [ID,SN]。ID-使用biztid作为设备唯一识别标识；SN-使用supplier_id、device_sn联合作为设备唯一识别标识。由于不同机型的supplier_id不同，推荐使用 ID 。 |
| `device_sn` | `DeviceSn` | `*string` | JSON body | 未标注 | 设备序列号 ，device_id_type 为 SN 时填写。需配合supplier_id使用。 |
| `merchant_type` | `MerchantType` | `*string` | JSON body | 未标注 | 商户类型，直连商户填写direct，间连商户填写indirect |
| `pid` | `Pid` | `*string` | JSON body | 未标注 | 直连场景填写商户收单pid，间连场景不填 |
| `smid` | `Smid` | `*string` | JSON body | 未标注 | 直连场景不填，间连场景填写商户收单smid |
| `supplier_id` | `SupplierId` | `*string` | JSON body | 未标注 | 设备供应商ID ，device_id_type 为 SN 时填写。需注意不同机型的供应商ID可能不同。 |
