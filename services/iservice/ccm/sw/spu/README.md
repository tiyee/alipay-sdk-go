# spu 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/iservice/ccm/sw/spu`。本包提供 1 个支付宝 V3 REST 操作。

通过 `spu.NewService(executor)` 复用 `*core.Client` 或最小 `core.Executor`。网络调用以 `context.Context` 为首参数，支持 `opts ...core.CallOption`，业务数据位于 `resp.Data`。

接口方法按单词边界使用大驼峰命名；请求、响应与专属模型位于对应业务包，共享模型从共同父包引用。

接入方式见[项目说明](../../../../../README.md)。默认测试离线执行。

返回[领域目录](../README.md)。

## 接口目录

| SDK 方法 | 官方接口 |
| --- | --- |
| [`CcmSwSPUSync`](#api-ccmswspusync) | `alipay.iservice.ccm.sw.spu.sync` |

## 使用示例

- [CcmSwSPUSync](api_spu_ccmswspusync_example_test.go)

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-ccmswspusync"></a>

### CcmSwSPUSync — 商品库-商品信息同步

- 官方接口：[`alipay.iservice.ccm.sw.spu.sync`](https://opendocs.alipay.com/apis/063syw)
- 路由：`POST /v3/alipay/iservice/ccm/sw/spu/sync`
- 请求类型：[CcmSwSPUSyncRequest](spu_request.go#L4)
- 响应类型：[CcmSwSPUSyncResponse](spu_response.go#L4)

商品库-商品信息同步

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `attribute` | `Attribute` | [`*[]SPUAttribute`](#model-spuattribute) | JSON body | 未标注 | 商品属性列表 |
| `brand` | `Brand` | `*string` | JSON body | 未标注 | 商品品牌 |
| `category` | `Category` | `*string` | JSON body | 未标注 | 商品类目，用于商品库中分类 |
| `description` | `Description` | `*string` | JSON body | 未标注 | 商品描述 |
| `icon` | `Icon` | `*string` | JSON body | 未标注 | 商品图片链接 特殊可选规则：is_delete=false的情况下必填 |
| `is_delete` | `IsDelete` | `*string` | JSON body | 未标注 | 默认为false，为true，表示删除知识点 |
| `library_id` | `LibraryId` | `*int64` | JSON body | 未标注 | 知识库id |
| `library_name` | `LibraryName` | `*string` | JSON body | 未标注 | 知识库名称 |
| `link_url` | `LinkUrl` | `*string` | JSON body | 未标注 | 商品链接 特殊可选规则：is_delete=false的情况下必填 |
| `original_price` | `OriginalPrice` | `*string` | JSON body | 未标注 | 商品原价 |
| `price` | `Price` | `*string` | JSON body | 未标注 | 商品售价 特殊可选规则：is_delete=false的情况下必填 |
| `spu_id` | `SpuId` | `*string` | JSON body | 未标注 | 商品ID |
| `status` | `Status` | `*string` | JSON body | 未标注 | 商品状态：AVAILABLE 表示可售卖,SOLD_OUT为售罄/不可卖，默认可售卖 |
| `title` | `Title` | `*string` | JSON body | 未标注 | 商品名称 特殊可选规则：is_delete=false的情况下必填 |

## 嵌套请求参数

<a id="model-spuattribute"></a>

### SPUAttribute

类型定义：[SPUAttribute](spu_model.go#L4)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `name` | `Name` | `*string` | 未标注 | 商品属性名 |
| `value` | `Value` | `*[]string` | 未标注 | 商品属性值 |
