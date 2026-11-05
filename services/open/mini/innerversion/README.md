# innerversion 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/open/mini/innerversion`。本包提供 13 个支付宝 V3 REST 操作。

通过 `innerversion.NewService(executor)` 复用 `*core.Client` 或最小 `core.Executor`。网络调用以 `context.Context` 为首参数，支持 `opts ...core.CallOption`，业务数据位于 `resp.Data`。

接口方法按单词边界使用大驼峰命名；请求、响应与专属模型位于对应业务包，共享模型从共同父包引用。

接入方式见[项目说明](../../../../README.md)。默认测试离线执行。

## 业务子包

| 导入子路径 | 操作数 | 方法 |
| --- | --- | --- |
| [`audit`](audit/README.md) | 2 | `MiniInnerVersionAuditCancel` / `MiniInnerVersionAuditSubmit` |
| [`beta`](beta/README.md) | 2 | `MiniInnerVersionBetaConfirm` / `MiniInnerVersionBetaInfoQuery` |
| [`gray`](gray/README.md) | 2 | `MiniInnerVersionGrayFinish` / `MiniInnerVersionGrayPublish` |
| [`upload`](upload/README.md) | 4 | `MiniInnerVersionNoBuildUpload` / `MiniInnerVersionTemplatedUpload` / `MiniInnerVersionUpload` / `MiniInnerVersionUploadStatusQuery` |

## 接口目录

| SDK 方法 | 官方接口 |
| --- | --- |
| [`MiniInnerVersionBackDevPublish`](#api-miniinnerversionbackdevpublish) | `alipay.open.mini.innerversion.backdev.publish` |
| [`MiniInnerVersionConditionBatchQuery`](#api-miniinnerversionconditionbatchquery) | `alipay.open.mini.innerversion.condition.batchquery` |
| [`MiniInnerVersionContentRollback`](#api-miniinnerversioncontentrollback) | `alipay.open.mini.innerversion.content.rollback` |
| [`MiniInnerVersionInfoDelete`](#api-miniinnerversioninfodelete) | `alipay.open.mini.innerversion.info.delete` |
| [`MiniInnerVersionInfoQuery`](#api-miniinnerversioninfoquery) | `alipay.open.mini.innerversion.info.query` |
| [`MiniInnerVersionInstantiationQuery`](#api-miniinnerversioninstantiationquery) | `alipay.open.mini.innerversion.instantiation.query` |
| [`MiniInnerVersionLastQuery`](#api-miniinnerversionlastquery) | `alipay.open.mini.innerversion.last.query` |
| [`MiniInnerVersionModelForOfflineQuery`](#api-miniinnerversionmodelforofflinequery) | `alipay.open.mini.innerversion.modelforoffline.query` |
| [`MiniInnerVersionOfflinePublish`](#api-miniinnerversionofflinepublish) | `alipay.open.mini.innerversion.offline.publish` |
| [`MiniInnerVersionOnline`](#api-miniinnerversiononline) | `alipay.open.mini.innerversion.online` |
| [`MiniInnerVersionOnlinePublish`](#api-miniinnerversiononlinepublish) | `alipay.open.mini.innerversion.online.publish` |
| [`MiniInnerVersionPackageCreate`](#api-miniinnerversionpackagecreate) | `alipay.open.mini.innerversion.package.create` |
| [`MiniInnerVersionSync`](#api-miniinnerversionsync) | `alipay.open.mini.innerversion.sync` |

## 使用示例

- [MiniInnerVersionBackDevPublish](api_innerversion_miniinnerversionbackdevpublish_example_test.go)
- [MiniInnerVersionConditionBatchQuery](api_innerversion_miniinnerversionconditionbatchquery_example_test.go)
- [MiniInnerVersionContentRollback](api_innerversion_miniinnerversioncontentrollback_example_test.go)
- [MiniInnerVersionInfoDelete](api_innerversion_miniinnerversioninfodelete_example_test.go)
- [MiniInnerVersionInfoQuery](api_innerversion_miniinnerversioninfoquery_example_test.go)
- [MiniInnerVersionInstantiationQuery](api_innerversion_miniinnerversioninstantiationquery_example_test.go)
- [MiniInnerVersionLastQuery](api_innerversion_miniinnerversionlastquery_example_test.go)
- [MiniInnerVersionModelForOfflineQuery](api_innerversion_miniinnerversionmodelforofflinequery_example_test.go)
- [MiniInnerVersionOfflinePublish](api_innerversion_miniinnerversionofflinepublish_example_test.go)
- [MiniInnerVersionOnline](api_innerversion_miniinnerversiononline_example_test.go)
- [MiniInnerVersionOnlinePublish](api_innerversion_miniinnerversiononlinepublish_example_test.go)
- [MiniInnerVersionPackageCreate](api_innerversion_miniinnerversionpackagecreate_example_test.go)
- [MiniInnerVersionSync](api_innerversion_miniinnerversionsync_example_test.go)

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-miniinnerversionbackdevpublish"></a>

### MiniInnerVersionBackDevPublish — 内部链路小程序退回开发

- 官方接口：[`alipay.open.mini.innerversion.backdev.publish`](https://opendocs.alipay.com/pre-apis/05xqxi)
- 路由：`POST /v3/alipay/open/mini/innerversion/backdev/publish`
- 请求类型：[MiniInnerVersionBackDevPublishRequest](innerversion_request.go#L4)
- 响应类型：[MiniInnerVersionBackDevPublishResponse](innerversion_response.go#L4)

小程序提供一套针对一二方场景的开发者提供一套统一的内部能力开放体系，定向输出赋能

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `app_origin` | `AppOrigin` | `*string` | JSON body | 未标注 | 业务来源 |
| `app_version` | `AppVersion` | `*string` | JSON body | 未标注 | 小程序版本 |
| `bundle_id` | `BundleId` | `*string` | JSON body | 未标注 | 端信息 |
| `mini_app_id` | `MiniAppId` | `*string` | JSON body | 未标注 | 小程序ID，仅特殊场景使用，普通业务方无需关注该参数 |
| `pid` | `Pid` | `*string` | JSON body | 未标注 | 小程序所属PID |

<a id="api-miniinnerversionconditionbatchquery"></a>

### MiniInnerVersionConditionBatchQuery — 根据状态批量查询版本信息

- 官方接口：[`alipay.open.mini.innerversion.condition.batchquery`](https://opendocs.alipay.com/pre-apis/05xk66)
- 路由：`POST /v3/alipay/open/mini/innerversion/condition/batchquery`
- 请求类型：[MiniInnerVersionConditionBatchQueryRequest](innerversion_request.go#L18)
- 响应类型：[MiniInnerVersionConditionBatchQueryResponse](innerversion_response.go#L8)

小程序提供一套针对一二方场景的开发者提供一套统一的内部能力开放体系，定向输出赋能

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `app_origin` | `AppOrigin` | `*string` | JSON body | 未标注 | 业务来源 |
| `bundle_id` | `BundleId` | `*string` | JSON body | 未标注 | com.alipay.alipaywallet:支付宝，com.amap.app:高德 |
| `mini_app_id` | `MiniAppId` | `*string` | JSON body | 未标注 | 小程序ID，仅特殊场景使用，普通业务方无需关注该参数。 |
| `page_num` | `PageNum` | `*string` | JSON body | 未标注 | 页码 |
| `page_size` | `PageSize` | `*string` | JSON body | 未标注 | 分页数 |
| `version_status` | `VersionStatus` | `*string` | JSON body | 未标注 | INIT:开发中，RELEASE:上架，OFFLINE:下架，AUDITING:审核中，AUDIT_REJECT:审核驳回，WAIT_RELEASE:待上架，GRAY:灰度中，以,号隔开 |

<a id="api-miniinnerversioncontentrollback"></a>

### MiniInnerVersionContentRollback — 内部链路回滚版本

- 官方接口：[`alipay.open.mini.innerversion.content.rollback`](https://opendocs.alipay.com/pre-apis/05xp3p)
- 路由：`POST /v3/alipay/open/mini/innerversion/content/rollback`
- 请求类型：[MiniInnerVersionContentRollbackRequest](innerversion_request.go#L34)
- 响应类型：[MiniInnerVersionContentRollbackResponse](innerversion_response.go#L14)

针对一二方场景的开发者，小程序提供一套统一的内部能力体系，定向输出赋能，本接口主要用于回滚版本

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `app_origin` | `AppOrigin` | `*string` | JSON body | 未标注 | 业务场景来源 |
| `app_version` | `AppVersion` | `*string` | JSON body | 未标注 | 小程序当前上架版本号 |
| `bundle_id` | `BundleId` | `*string` | JSON body | 未标注 | 端 |
| `mini_app_id` | `MiniAppId` | `*string` | JSON body | 未标注 | 小程序ID，仅特殊场景使用，普通业务方无需关注该参数。 |
| `pid` | `Pid` | `*string` | JSON body | 未标注 | 小程序所属主体 |

<a id="api-miniinnerversioninfodelete"></a>

### MiniInnerVersionInfoDelete — 内部链路删除版本

- 官方接口：[`alipay.open.mini.innerversion.info.delete`](https://opendocs.alipay.com/pre-apis/05xe18)
- 路由：`DELETE /v3/alipay/open/mini/innerversion/info/delete`
- 请求类型：[MiniInnerVersionInfoDeleteRequest](innerversion_request.go#L128)
- 响应类型：[MiniInnerVersionInfoDeleteResponse](innerversion_response.go#L18)

针对一二方场景的开发者，小程序提供一套统一的内部能力体系，定向输出赋能，本接口主要用于删除版本

请求约束：MiniInnerVersionInfoDeleteRequest is the request type for alipay.open.mini.innerversion.info.delete.

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `mini_app_id` | `MiniAppId` | `*string` | query | 否 | 小程序ID |
| `app_origin` | `AppOrigin` | `*string` | query | 否 | 业务场景来源 |
| `bundle_id` | `BundleId` | `*string` | query | 否 | 端信息 |
| `app_version` | `AppVersion` | `*string` | query | 否 | 小程序版本号 |
| `pid` | `Pid` | `*string` | query | 否 | 操作者ID，支付宝租户必填 |
| `inst_code` | `InstCode` | `*string` | query | 否 | 租户信息 |

<a id="api-miniinnerversioninfoquery"></a>

### MiniInnerVersionInfoQuery — 内部链路小程序版本信息查询

- 官方接口：[`alipay.open.mini.innerversion.info.query`](https://opendocs.alipay.com/pre-apis/05xdzi)
- 路由：`GET /v3/alipay/open/mini/innerversion/info/query`
- 请求类型：[MiniInnerVersionInfoQueryRequest](innerversion_request.go#L144)
- 响应类型：[MiniInnerVersionInfoQueryResponse](innerversion_response.go#L22)

小程序提供一套针对一二方场景的开发者提供一套统一的内部能力开放体系，定向输出赋能

请求约束：MiniInnerVersionInfoQueryRequest is the request type for alipay.open.mini.innerversion.info.query.

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `bundle_id` | `BundleId` | `*string` | query | 否 | 端信息 |
| `mini_app_id` | `MiniAppId` | `*string` | query | 否 | 小程序ID |
| `app_version` | `AppVersion` | `*string` | query | 否 | 查询版本号 |
| `app_origin` | `AppOrigin` | `*string` | query | 否 | 业务来源，新接入方需要向支付宝申请专用来源，否则不予接入。https://yuque.antfin-inc.com/tinyapp-all/qddncu/bidushixiang |

<a id="api-miniinnerversioninstantiationquery"></a>

### MiniInnerVersionInstantiationQuery — 小程序实例化关系查询

- 官方接口：[`alipay.open.mini.innerversion.instantiation.query`](https://opendocs.alipay.com/pre-apis/05xdyz)
- 路由：`GET /v3/alipay/open/mini/innerversion/instantiation/query`
- 请求类型：[MiniInnerVersionInstantiationQueryRequest](innerversion_request.go#L156)
- 响应类型：[MiniInnerVersionInstantiationQueryResponse](innerversion_response.go#L72)

用于通过模板实例化的小程序版本查询实例化关系

请求约束：MiniInnerVersionInstantiationQueryRequest is the request type for alipay.open.mini.innerversion.instantiation.query.

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `app_origin` | `AppOrigin` | `*string` | query | 否 | 来源类型，新接入方需要向支付宝申请专用来源，否则不予接入，申请方式请参见接入手册。 |
| `bundle_id` | `BundleId` | `*string` | query | 否 | 端id |
| `mini_app_id` | `MiniAppId` | `*string` | query | 否 | 小程序ID，特殊场景专用，普通业务方无需关注该参数。 |
| `app_version` | `AppVersion` | `*string` | query | 否 | 实例化小程序版本号 |

<a id="api-miniinnerversionlastquery"></a>

### MiniInnerVersionLastQuery — 内部链路查询最新版本号

- 官方接口：[`alipay.open.mini.innerversion.last.query`](https://opendocs.alipay.com/pre-apis/05xhs8)
- 路由：`GET /v3/alipay/open/mini/innerversion/last/query`
- 请求类型：[MiniInnerVersionLastQueryRequest](innerversion_request.go#L168)
- 响应类型：[MiniInnerVersionLastQueryResponse](innerversion_response.go#L86)

针对一二方场景的开发者，小程序提供一套统一的内部能力体系，定向输出赋能，本接口主要用于查询最新版本号

请求约束：MiniInnerVersionLastQueryRequest is the request type for alipay.open.mini.innerversion.last.query.

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `mini_app_id` | `MiniAppId` | `*string` | query | 否 | 小程序ID，仅特殊场景使用，普通业务方无需关注该参数。 |
| `bundle_id` | `BundleId` | `*string` | query | 否 | 端信息 |
| `app_origin` | `AppOrigin` | `*string` | query | 否 | 业务来源场景 |
| `pid` | `Pid` | `*string` | query | 否 | 小程序主体 |

<a id="api-miniinnerversionmodelforofflinequery"></a>

### MiniInnerVersionModelForOfflineQuery — 小程序线上版本包数据同步线下场景查询

- 官方接口：[`alipay.open.mini.innerversion.modelforoffline.query`](https://opendocs.alipay.com/pre-apis/05xqwp)
- 路由：`GET /v3/alipay/open/mini/innerversion/modelforoffline/query`
- 请求类型：[MiniInnerVersionModelForOfflineQueryRequest](innerversion_request.go#L180)
- 响应类型：[MiniInnerVersionModelForOfflineQueryResponse](innerversion_response.go#L114)

用于查询线上的版本包数据，同步到线下

请求约束：MiniInnerVersionModelForOfflineQueryRequest is the request type for alipay.open.mini.innerversion.modelforoffline.query.

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `mini_app_id` | `MiniAppId` | `*string` | query | 否 | 小程序ID |
| `app_version` | `AppVersion` | `*string` | query | 否 | 版本号 |
| `bundle_id` | `BundleId` | `*string` | query | 否 | com.alipay.alipaywallet:支付包 |
| `inst_code` | `InstCode` | `*string` | query | 否 | alipay:支付宝 |

<a id="api-miniinnerversionofflinepublish"></a>

### MiniInnerVersionOfflinePublish — 内部链路小程序下架

- 官方接口：[`alipay.open.mini.innerversion.offline.publish`](https://opendocs.alipay.com/pre-apis/05xo7b)
- 路由：`POST /v3/alipay/open/mini/innerversion/offline/publish`
- 请求类型：[MiniInnerVersionOfflinePublishRequest](innerversion_request.go#L48)
- 响应类型：[MiniInnerVersionOfflinePublishResponse](innerversion_response.go#L122)

小程序提供一套针对一二方场景的开发者提供一套统一的内部能力开放体系，定向输出赋能

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `app_origin` | `AppOrigin` | `*string` | JSON body | 未标注 | 业务来源 |
| `app_version` | `AppVersion` | `*string` | JSON body | 未标注 | 版本号 |
| `bundle_id` | `BundleId` | `*string` | JSON body | 未标注 | 端信息 |
| `mini_app_id` | `MiniAppId` | `*string` | JSON body | 未标注 | 小程序ID，仅特殊场景使用，普通业务方无需关注该参数 |
| `pid` | `Pid` | `*string` | JSON body | 未标注 | 小程序所属PID |

<a id="api-miniinnerversiononline"></a>

### MiniInnerVersionOnline — 内部小程序-上架

- 官方接口：[`alipay.open.mini.innerversion.online`](https://opendocs.alipay.com/pre-apis/05xhq6)
- 路由：`POST /v3/alipay/open/mini/innerversion/online`
- 请求类型：[MiniInnerVersionOnlineRequest](innerversion_request.go#L62)
- 响应类型：[MiniInnerVersionOnlineResponse](innerversion_response.go#L130)

内部小程序上架接口

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `app_offline_version` | `AppOfflineVersion` | `*string` | JSON body | 未标注 | 下架版本 |
| `app_online_version` | `AppOnlineVersion` | `*string` | JSON body | 未标注 | 上架版本 |
| `app_origin` | `AppOrigin` | `*string` | JSON body | 未标注 | 业务来源 |
| `bundle_id` | `BundleId` | `*string` | JSON body | 未标注 | 一个端的标识，用于区分不同的客户端，每接入一个客户端，都需要向小程序应用中心申请bundelId入驻 |
| `inst_code` | `InstCode` | `*string` | JSON body | 未标注 | 租户code，alipay or taobao |
| `mini_app_id` | `MiniAppId` | `*string` | JSON body | 未标注 | 小程序ID |

<a id="api-miniinnerversiononlinepublish"></a>

### MiniInnerVersionOnlinePublish — 内部链路小程序上架

- 官方接口：[`alipay.open.mini.innerversion.online.publish`](https://opendocs.alipay.com/pre-apis/05xgf5)
- 路由：`POST /v3/alipay/open/mini/innerversion/online/publish`
- 请求类型：[MiniInnerVersionOnlinePublishRequest](innerversion_request.go#L78)
- 响应类型：[MiniInnerVersionOnlinePublishResponse](innerversion_response.go#L126)

小程序提供一套针对一二方场景的开发者提供一套统一的内部能力开放体系，定向输出赋能

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `app_origin` | `AppOrigin` | `*string` | JSON body | 未标注 | 业务来源 |
| `app_version` | `AppVersion` | `*string` | JSON body | 未标注 | 小程序版本 |
| `bundle_id` | `BundleId` | `*string` | JSON body | 未标注 | 端信息 |
| `downgrade` | `Downgrade` | `*bool` | JSON body | 未标注 | 否服务降级，如果降级，拉包策略如果端上有本地包，则使用本地包打开，同时异步下载新包，待下次打开则使用新包 |
| `mini_app_id` | `MiniAppId` | `*string` | JSON body | 未标注 | 小程序ID，仅特殊场景使用，普通业务方无需关注该参数 |
| `pid` | `Pid` | `*string` | JSON body | 未标注 | 小程序所属PID |
| `release_type` | `ReleaseType` | `*string` | JSON body | 未标注 | normal默认普通发布，high 高保发布 |

<a id="api-miniinnerversionpackagecreate"></a>

### MiniInnerVersionPackageCreate — 单独构建包信息

- 官方接口：[`alipay.open.mini.innerversion.package.create`](https://opendocs.alipay.com/pre-apis/05x9fo)
- 路由：`POST /v3/alipay/open/mini/innerversion/package/create`
- 请求类型：[MiniInnerVersionPackageCreateRequest](innerversion_request.go#L96)
- 响应类型：[MiniInnerVersionPackageCreateResponse](innerversion_response.go#L134)

内部OPENAPI 单独构建包信息。将构建链路中打预发包逻辑抽离出来，成为单独的能力，让伙伴等业务自己来调用

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `app_origin` | `AppOrigin` | `*string` | JSON body | 未标注 | 业务来源 |
| `app_version` | `AppVersion` | `*string` | JSON body | 未标注 | 小程序版本 |
| `bundle_id` | `BundleId` | `*string` | JSON body | 未标注 | 端信息 |
| `mini_app_id` | `MiniAppId` | `*string` | JSON body | 未标注 | 小程序ID，仅特殊场景使用，普通业务方无需关注该参数。 |
| `package_type` | `PackageType` | `*string` | JSON body | 未标注 | 打包类型。预发包和覆盖率包 |
| `pid` | `Pid` | `*string` | JSON body | 未标注 | pid. 伙伴链路（app_origin='HUOBAN'）时必填 |

<a id="api-miniinnerversionsync"></a>

### MiniInnerVersionSync — 小程序包信息推送

- 官方接口：[`alipay.open.mini.innerversion.sync`](https://opendocs.alipay.com/pre-apis/05xlnt)
- 路由：`POST /v3/alipay/open/mini/innerversion/sync`
- 请求类型：[MiniInnerVersionSyncRequest](innerversion_request.go#L112)
- 响应类型：[MiniInnerVersionSyncResponse](innerversion_response.go#L138)

小程序上架、回滚，因拉包机制，用户无法及时生效，拉到原版本的用户，过期才会失效，该接口能力可以推送新版本到用户

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `app_origin` | `AppOrigin` | `*string` | JSON body | 未标注 | 业务参数来源 |
| `app_version` | `AppVersion` | `*string` | JSON body | 未标注 | 推送的小程序版本号 |
| `bundle_id` | `BundleId` | `*string` | JSON body | 未标注 | 端信息 |
| `dev_id` | `DevId` | `*string` | JSON body | 未标注 | 操作人ID |
| `mini_app_id` | `MiniAppId` | `*string` | JSON body | 未标注 | 小程序ID，仅特殊场景使用，普通业务方无需关注该参数。 |
| `sync_type` | `SyncType` | `*string` | JSON body | 未标注 | 同送方式 |
