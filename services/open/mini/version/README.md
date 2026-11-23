# version 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/open/mini/version`。本包提供 9 个支付宝 V3 REST 操作。

通过 `version.NewService(executor)` 复用 `*core.Client` 或最小 `core.Executor`。网络调用以 `context.Context` 为首参数，支持 `opts ...core.CallOption`，业务数据位于 `resp.Data`。

接口方法按单词边界使用大驼峰命名；请求、响应与专属模型位于对应业务包，共享模型从共同父包引用。

接入方式见[项目说明](../../../../README.md)。默认测试离线执行。

## 业务子包

| 导入子路径 | 操作数 | 方法 |
| --- | --- | --- |
| [`audit`](audit/README.md) | 2 | `MiniVersionAuditApply` / `MiniVersionAuditCancel` |
| [`gray`](gray/README.md) | 2 | `MiniVersionGrayCancel` / `MiniVersionGrayOnline` |

## 接口目录

| SDK 方法 | 官方接口 |
| --- | --- |
| [`MiniVersionAuditedCancel`](#api-miniversionauditedcancel) | `alipay.open.mini.version.audited.cancel` |
| [`MiniVersionBuildQuery`](#api-miniversionbuildquery) | `alipay.open.mini.version.build.query` |
| [`MiniVersionDelete`](#api-miniversiondelete) | `alipay.open.mini.version.delete` |
| [`MiniVersionDetailQuery`](#api-miniversiondetailquery) | `alipay.open.mini.version.detail.query` |
| [`MiniVersionListQuery`](#api-miniversionlistquery) | `alipay.open.mini.version.list.query` |
| [`MiniVersionOffline`](#api-miniversionoffline) | `alipay.open.mini.version.offline` |
| [`MiniVersionOnline`](#api-miniversiononline) | `alipay.open.mini.version.online` |
| [`MiniVersionRollback`](#api-miniversionrollback) | `alipay.open.mini.version.rollback` |
| [`MiniVersionUpload`](#api-miniversionupload) | `alipay.open.mini.version.upload` |

## 使用示例

- [MiniVersionAuditedCancel](api_version_miniversionauditedcancel_example_test.go)
- [MiniVersionBuildQuery](api_version_miniversionbuildquery_example_test.go)
- [MiniVersionDelete](api_version_miniversiondelete_example_test.go)
- [MiniVersionDetailQuery](api_version_miniversiondetailquery_example_test.go)
- [MiniVersionListQuery](api_version_miniversionlistquery_example_test.go)
- [MiniVersionOffline](api_version_miniversionoffline_example_test.go)
- [MiniVersionOnline](api_version_miniversiononline_example_test.go)
- [MiniVersionRollback](api_version_miniversionrollback_example_test.go)
- [MiniVersionUpload](api_version_miniversionupload_example_test.go)

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-miniversionauditedcancel"></a>

### MiniVersionAuditedCancel — 小程序退回开发

- 官方接口：[`alipay.open.mini.version.audited.cancel`](https://opendocs.alipay.com/open-v3/7a284018_alipay.open.mini.version.audited.cancel)
- 路由：`POST /v3/alipay/open/mini/version/audited/cancel`
- 请求类型：[MiniVersionAuditedCancelRequest](version_request.go#L4)
- 响应类型：[MiniVersionAuditedCancelResponse](version_response.go#L4)

小程序退回开发

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `app_version` | `AppVersion` | `*string` | JSON body | 未标注 | 商家小程序版本号。可将已过审、审核被驳回的小程序版本退回到开发中的状态。 |
| `bundle_id` | `BundleId` | `*string` | JSON body | 未标注 | 小程序客户端类型。常见支持如下客户端： com.alipay.alipaywallet：支付宝端； com.alibaba.android.rimet：DINGDING端； com.amap.app：高德端； com.alibaba.ailabs.genie.webapps：天猫精灵端； com.alipay.iot.xpaas：支付宝IoT端。 如需更多端投放，请联系业务BD。 |

<a id="api-miniversionbuildquery"></a>

### MiniVersionBuildQuery — 小程序查询版本构建状态

- 官方接口：[`alipay.open.mini.version.build.query`](https://opendocs.alipay.com/open-v3/dda696d3_alipay.open.mini.version.build.query)
- 路由：`GET /v3/alipay/open/mini/version/build/query`
- 请求类型：[MiniVersionBuildQueryRequest](version_request.go#L62)
- 响应类型：[MiniVersionBuildQueryResponse](version_response.go#L8)

小程序查询版本构建状态

请求约束：MiniVersionBuildQueryRequest is the request type for alipay.open.mini.version.build.query.

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `app_version` | `AppVersion` | `*string` | query | 否 | 小程序版本号 |
| `bundle_id` | `BundleId` | `*string` | query | 否 | 小程序客户端类型，默认为支付宝端。常见支持如下客户端： com.alipay.alipaywallet：支付宝端； com.alibaba.android.rimet：DINGDING端； com.amap.app：高德端； com.alibaba.ailabs.genie.webapps：天猫精灵端； com.alipay.iot.xpaas：支付宝IoT端。 如需更多端投放，请联系业务BD。 |

<a id="api-miniversiondelete"></a>

### MiniVersionDelete — 小程序删除版本

- 官方接口：[`alipay.open.mini.version.delete`](https://opendocs.alipay.com/open-v3/eef61c13_alipay.open.mini.version.delete)
- 路由：`POST /v3/alipay/open/mini/version/delete`
- 请求类型：[MiniVersionDeleteRequest](version_request.go#L12)
- 响应类型：[MiniVersionDeleteResponse](version_response.go#L20)

小程序删除版本

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `app_version` | `AppVersion` | `*string` | JSON body | 未标注 | 小程序版本号。只有状态为开发中的版本才能被删除。 |
| `bundle_id` | `BundleId` | `*string` | JSON body | 未标注 | 小程序客户端类型，默认为支付宝端。常见支持如下客户端： com.alipay.alipaywallet：支付宝端； com.alibaba.android.rimet：DINGDING端； com.amap.app：高德端； com.alibaba.ailabs.genie.webapps：天猫精灵端； com.alipay.iot.xpaas：支付宝IoT端。 如需更多端投放，请联系业务BD。 |

<a id="api-miniversiondetailquery"></a>

### MiniVersionDetailQuery — 小程序版本详情查询

- 官方接口：[`alipay.open.mini.version.detail.query`](https://opendocs.alipay.com/open-v3/048e5e71_alipay.open.mini.version.detail.query)
- 路由：`GET /v3/alipay/open/mini/version/detail/query`
- 请求类型：[MiniVersionDetailQueryRequest](version_request.go#L70)
- 响应类型：[MiniVersionDetailQueryResponse](version_response.go#L24)

小程序版本详情查询

请求约束：MiniVersionDetailQueryRequest is the request type for alipay.open.mini.version.detail.query.

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `app_version` | `AppVersion` | `*string` | query | 否 | 小程序版本号 |
| `bundle_id` | `BundleId` | `*string` | query | 否 | 小程序客户端类型，默认为支付宝端。常见支持如下客户端： com.alipay.alipaywallet：支付宝端； com.alibaba.android.rimet：DINGDING端； com.amap.app：高德端； com.alibaba.ailabs.genie.webapps：天猫精灵端； com.alipay.iot.xpaas：支付宝IoT端。 如需更多端投放，请联系业务BD。 |

<a id="api-miniversionlistquery"></a>

### MiniVersionListQuery — 小程序版本列表查询

- 官方接口：[`alipay.open.mini.version.list.query`](https://opendocs.alipay.com/open-v3/ac91bcd8_alipay.open.mini.version.list.query)
- 路由：`GET /v3/alipay/open/mini/version/list/query`
- 请求类型：[MiniVersionListQueryRequest](version_request.go#L78)
- 响应类型：[MiniVersionListQueryResponse](version_response.go#L86)

根据状态、端分页查询小程序版本基础信息列表，查询结果根据版本号大小降序排列，即0.0.2在前，0.0.1在后。

请求约束：MiniVersionListQueryRequest is the request type for alipay.open.mini.version.list.query.

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `bundle_id` | `BundleId` | `*string` | query | 否 | 端参数，用于查询多端版本，不传默认为支付宝端。如需查看小程序版本信息，则为必传。常见支持如下客户端： com.alipay.alipaywallet：支付宝端； com.alibaba.android.rimet：DINGDING端； com.amap.app：高德端； com.alibaba.ailabs.genie.webapps：天猫精灵端； com.alipay.iot.xpaas：支付宝IoT端。 如需更多端投放，请联系业务BD。 |
| `version_status` | `VersionStatus` | `*string` | query | 否 | 版本状态列表，用英文逗号","分割，可选；不填默认不返回。如需查看小程序版本信息，则为必传。 |

<a id="api-miniversionoffline"></a>

### MiniVersionOffline — 小程序下架

- 官方接口：[`alipay.open.mini.version.offline`](https://opendocs.alipay.com/open-v3/c5ffdf9c_alipay.open.mini.version.offline)
- 路由：`POST /v3/alipay/open/mini/version/offline`
- 请求类型：[MiniVersionOfflineRequest](version_request.go#L20)
- 响应类型：[MiniVersionOfflineResponse](version_response.go#L94)

小程序下架

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `app_version` | `AppVersion` | `*string` | JSON body | 未标注 | 商家小程序已上架版本号。 |
| `bundle_id` | `BundleId` | `*string` | JSON body | 未标注 | 小程序客户端类型，默认为支付宝端。常见支持如下客户端： com.alipay.alipaywallet：支付宝端； com.alibaba.android.rimet：DINGDING端； com.amap.app：高德端； com.alibaba.ailabs.genie.webapps：天猫精灵端； com.alipay.iot.xpaas：支付宝IoT端。 如需更多端投放，请联系业务BD。 |

<a id="api-miniversiononline"></a>

### MiniVersionOnline — 小程序上架

- 官方接口：[`alipay.open.mini.version.online`](https://opendocs.alipay.com/open-v3/4d9b2578_alipay.open.mini.version.online)
- 路由：`POST /v3/alipay/open/mini/version/online`
- 请求类型：[MiniVersionOnlineRequest](version_request.go#L28)
- 响应类型：[MiniVersionOnlineResponse](version_response.go#L98)

小程序上架

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `app_version` | `AppVersion` | `*string` | JSON body | 未标注 | 商家小程序版本号，需为审核通过状态或灰度中版本。 |
| `bundle_id` | `BundleId` | `*string` | JSON body | 未标注 | 小程序客户端类型，默认为支付宝端。常见支持如下客户端： com.alipay.alipaywallet：支付宝端； com.alibaba.android.rimet：DINGDING端； com.amap.app：高德端； com.alibaba.ailabs.genie.webapps：天猫精灵端； com.alipay.iot.xpaas：支付宝IoT端。 如需更多端投放，请联系业务BD。 |
| `downgrade` | `Downgrade` | `*bool` | JSON body | 未标注 | true - 性能优先； false - 覆盖优先（默认）；试点开放，如需使用，请咨询在线客服 |
| `permit_registration_limit_release` | `PermitRegistrationLimitRelease` | `*bool` | JSON body | 未标注 | 如果该字段为true，在首次发版的小程序在备案未完成的情况下，允许小程序上架，但小程序上架后只可以上架可扫码进行使用，但不可被搜索、限制访问人数（50人/每天）、限制公域（无法进行商品、服务、营促销）、限制上架后90天不完成备案自动下架。 默认为false |

<a id="api-miniversionrollback"></a>

### MiniVersionRollback — 小程序回滚

- 官方接口：[`alipay.open.mini.version.rollback`](https://opendocs.alipay.com/open-v3/f1c39e2b_alipay.open.mini.version.rollback)
- 路由：`POST /v3/alipay/open/mini/version/rollback`
- 请求类型：[MiniVersionRollbackRequest](version_request.go#L40)
- 响应类型：[MiniVersionRollbackResponse](version_response.go#L102)

小程序回滚

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `app_version` | `AppVersion` | `*string` | JSON body | 未标注 | 商家小程序已上架版本。下架后将自动回滚至上一（已上架）小程序版本。 例如：商家小程序有 0.01（上一上架版本）、0.02（未上架版本）、0.03（当前上架版本） 三个版本，回滚时需传入商家小程序版本 0.03 表示将 0.03 版本回滚至上一个已上架版本即此处 0.01 版本，同时 0.03 版本将自动下架，0.01 版本自动上架（无需审核）。 |
| `bundle_id` | `BundleId` | `*string` | JSON body | 未标注 | 小程序客户端类型，默认为支付宝端。常见支持如下客户端： com.alipay.alipaywallet：支付宝端； com.alibaba.android.rimet：DINGDING端； com.amap.app：高德端； com.alibaba.ailabs.genie.webapps：天猫精灵端； com.alipay.iot.xpaas：支付宝IoT端。 如需更多端投放，请联系业务BD。 |

<a id="api-miniversionupload"></a>

### MiniVersionUpload — 小程序基于模板上传版本

- 官方接口：[`alipay.open.mini.version.upload`](https://opendocs.alipay.com/open-v3/9d32147b_alipay.open.mini.version.upload)
- 路由：`POST /v3/alipay/open/mini/version/upload`
- 请求类型：[MiniVersionUploadRequest](version_request.go#L48)
- 响应类型：[MiniVersionUploadResponse](version_response.go#L106)

小程序基于模板上传版本

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `app_version` | `AppVersion` | `*string` | JSON body | 未标注 | 商家小程序版本号。版本号必须满足 x.y.z, 且均为数字。要求版本号比商户之前最新的版本号高。 |
| `bundle_id` | `BundleId` | `*string` | JSON body | 未标注 | 小程序客户端类型，默认为支付宝端。常见支持如下客户端： com.alipay.alipaywallet：支付宝端； com.alibaba.android.rimet：DINGDING端； com.amap.app：高德端； com.alibaba.ailabs.genie.webapps：天猫精灵端； com.alipay.iot.xpaas：支付宝IoT端。 如需更多端投放，请联系业务BD。 |
| `ext` | `Ext` | `*string` | JSON body | 未标注 | 自定义参数，用于配置商家小程序的个性化内容，具体使用方法参见示例代码。 在小程序中可以通过 [my.getExtConfig()](https://opendocs.alipay.com/mini/api/getExtConfig) 或 [my.getExtConfigSync()](https://opendocs.alipay.com/mini/api/getExtConfigSync) 获取 ext 参数内容（建议支付宝版本 &gt; 10.1.38，并使用 [my.canIUse](https://opendocs.alipay.com/mini/api/can-i-use) 做兼容处理）。 |
| `template_id` | `TemplateId` | `*string` | JSON body | 未标注 | 小程序模板 APPID。 |
| `template_version` | `TemplateVersion` | `*string` | JSON body | 未标注 | 小程序模板版本号。版本号必须满足 x.y.z, 且均为数字。不传默认使用最新在架（已上架）模板版本。 |

## 嵌套请求参数

<a id="model-regioninfo"></a>

### RegionInfo

类型定义：[RegionInfo](version_model.go#L70)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `area_code` | `AreaCode` | `*string` | 未标注 | 地址所属区代码 |
| `area_name` | `AreaName` | `*string` | 未标注 | 地址所属区名称 |
| `city_code` | `CityCode` | `*string` | 未标注 | 地址所属市代码 |
| `city_name` | `CityName` | `*string` | 未标注 | 地址所属市名称 |
| `province_code` | `ProvinceCode` | `*string` | 未标注 | 地址所属省份代码 |
| `province_name` | `ProvinceName` | `*string` | 未标注 | 地址所属省份名称 |
