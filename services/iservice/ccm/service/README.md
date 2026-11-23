# iservice/ccm/service 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/iservice/ccm/service`。本目录提供 3 个支付宝开放平台 V3 REST 操作。

## 使用方式

先按[项目接入说明](../../../../README.md)创建 `*core.Client`，再通过 `service.NewService(client)` 创建 Service；所有网络调用以 `context.Context` 为首个参数，支持 `opts ...core.CallOption`。响应为 `*core.Response[T]`，业务数据位于 `resp.Data`。

请求、响应与嵌套类型的链接指向本目录源码；可参考本目录的 [API 示例](api_service_example_test.go)。

示例文件中的 `TestSandboxService_*` 可使用沙箱平台证书执行真实调用。默认 `go test` 保持离线。

## 接口目录

| SDK 方法 | 官方接口 | 用途 | HTTP |
| --- | --- | --- | --- |
| [`CcmServiceBuy`](#api-ccmservicebuy) | [`alipay.iservice.ccm.service.buy`](https://opendocs.alipay.com/open-v3/705fa9ad_alipay.iservice.ccm.service.buy?scene=common&pathHash=42f013fd) | 租户开通isv服务接口 | `POST` |
| [`CcmServiceClose`](#api-ccmserviceclose) | [`alipay.iservice.ccm.service.close`](https://opendocs.alipay.com/open-v3/2202303e_alipay.iservice.ccm.service.close?scene=common&pathHash=4e8520fc) | 关闭租户开通的isv服务接口 | `POST` |
| [`CcmServiceInitialize`](#api-ccmserviceinitialize) | [`alipay.iservice.ccm.service.initialize`](https://opendocs.alipay.com/open-v3/ab388310_alipay.iservice.ccm.service.initialize?scene=common&pathHash=3dce471e) | isv服务初始化接口 | `POST` |

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-ccmservicebuy"></a>

### CcmServiceBuy — 租户开通isv服务接口

- 官方接口：[`alipay.iservice.ccm.service.buy`](https://opendocs.alipay.com/open-v3/705fa9ad_alipay.iservice.ccm.service.buy?scene=common&pathHash=42f013fd)
- 路由：`POST /v3/alipay/iservice/ccm/service/buy`
- 请求类型：[CcmServiceBuyRequest](service_request.go#L5)
- 响应类型：[CcmServiceBuyResponse](service_response.go#L9)

租户购买了isv服务后，通知ccm开通服务

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `service_code` | `ServiceCode` | `*string` | JSON body | 未标注 | 服务代码。通过[alipay.iservice.ccm.service.initialize](https://opendocs.alipay.com/apis/01dmgh)(isv服务初始化接口) 获取。 |

<a id="api-ccmserviceclose"></a>

### CcmServiceClose — 关闭租户开通的isv服务接口

- 官方接口：[`alipay.iservice.ccm.service.close`](https://opendocs.alipay.com/open-v3/2202303e_alipay.iservice.ccm.service.close?scene=common&pathHash=4e8520fc)
- 路由：`POST /v3/alipay/iservice/ccm/service/close`
- 请求类型：[CcmServiceCloseRequest](service_request.go#L11)
- 响应类型：[CcmServiceCloseResponse](service_response.go#L17)

租户购买的ISV服务到期后，isv通知ccm关闭服务

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `service_code` | `ServiceCode` | `*string` | JSON body | 未标注 | 服务代码。通过[alipay.iservice.ccm.service.initialize](https://opendocs.alipay.com/apis/01dmgh)(isv服务初始化接口) 获取。 |

<a id="api-ccmserviceinitialize"></a>

### CcmServiceInitialize — isv服务初始化接口

- 官方接口：[`alipay.iservice.ccm.service.initialize`](https://opendocs.alipay.com/open-v3/ab388310_alipay.iservice.ccm.service.initialize?scene=common&pathHash=3dce471e)
- 路由：`POST /v3/alipay/iservice/ccm/service/initialize`
- 请求类型：[CcmServiceInitializeRequest](service_request.go#L17)
- 响应类型：[CcmServiceInitializeResponse](service_response.go#L20)

ISV调用此接口对接入CCM的服务代码、名称、描述等信息做初始化

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `description` | `Description` | `*string` | JSON body | 未标注 | 服务描述 |
| `icon` | `Icon` | `*string` | JSON body | 未标注 | 服务图标url。支持 http(s) 链接。 |
| `service_code` | `ServiceCode` | `*string` | JSON body | 未标注 | 服务编码。可在服务上架后，在服务市场上查询此信息。 |
| `service_name` | `ServiceName` | `*string` | JSON body | 未标注 | 服务名称。可在服务上架后，在服务市场上查询此信息。 |
| `service_order_url` | `ServiceOrderUrl` | `*string` | JSON body | 未标注 | 服务订购二维码链接 |
| `spis` | `Spis` | [`*[]ISVSpiDefinition`](#model-isvspidefinition) | JSON body | 未标注 | spi定义列表 |

## 嵌套请求参数

以下仅展开本目录请求中实际使用的对象；字段必填规则沿用官方快照，“未标注”须结合业务说明确认。响应模型请沿各接口的响应类型链接查看。

<a id="model-isvspidefinition"></a>

### ISVSpiDefinition

类型定义：[ISVSpiDefinition](service_model.go#L5)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `biz_code` | `BizCode` | `*string` | 未标注 | ISV自定义的标识功能的业务代码，不可重复 |
| `description` | `Description` | `*string` | 未标注 | spi功能描述 |
| `icon` | `Icon` | `*string` | 未标注 | 功能图标文件url |
| `spi_endpoint` | `SpiEndpoint` | `*string` | 未标注 | spi接口服务地址 |
| `spi_ext_property` | `SpiExtProperty` | `*string` | 未标注 | spi接口扩展参数，json格式字符串 |
| `spi_key` | `SpiKey` | `*string` | 未标注 | CCM预先定义的spi key，与插件位置有关 |
| `spi_name` | `SpiName` | `*string` | 未标注 | SPI 名称 |
