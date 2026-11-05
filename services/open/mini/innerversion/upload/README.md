# upload 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/open/mini/innerversion/upload`。本包提供 4 个支付宝 V3 REST 操作。

通过 `upload.NewService(executor)` 复用 `*core.Client` 或最小 `core.Executor`。网络调用以 `context.Context` 为首参数，支持 `opts ...core.CallOption`，业务数据位于 `resp.Data`。

接口方法按单词边界使用大驼峰命名；请求、响应与专属模型位于对应业务包，共享模型从共同父包引用。

接入方式见[项目说明](../../../../../README.md)。默认测试离线执行。

返回[领域目录](../README.md)。

## 接口目录

| SDK 方法 | 官方接口 |
| --- | --- |
| [`MiniInnerVersionNoBuildUpload`](#api-miniinnerversionnobuildupload) | `alipay.open.mini.innerversion.nobuild.upload` |
| [`MiniInnerVersionTemplatedUpload`](#api-miniinnerversiontemplatedupload) | `alipay.open.mini.innerversion.templated.upload` |
| [`MiniInnerVersionUpload`](#api-miniinnerversionupload) | `alipay.open.mini.innerversion.upload` |
| [`MiniInnerVersionUploadStatusQuery`](#api-miniinnerversionuploadstatusquery) | `alipay.open.mini.innerversion.uploadstatus.query` |

## 使用示例

- [MiniInnerVersionNoBuildUpload](api_upload_miniinnerversionnobuildupload_example_test.go)
- [MiniInnerVersionTemplatedUpload](api_upload_miniinnerversiontemplatedupload_example_test.go)
- [MiniInnerVersionUpload](api_upload_miniinnerversionupload_example_test.go)
- [MiniInnerVersionUploadStatusQuery](api_upload_miniinnerversionuploadstatusquery_example_test.go)

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-miniinnerversionnobuildupload"></a>

### MiniInnerVersionNoBuildUpload — 手淘小程序产物包免构建版本上传

- 官方接口：[`alipay.open.mini.innerversion.nobuild.upload`](https://opendocs.alipay.com/pre-apis/05xiqc)
- 路由：`POST /v3/alipay/open/mini/innerversion/nobuild/upload`
- 请求类型：[MiniInnerVersionNoBuildUploadRequest](upload_request.go#L4)
- 响应类型：[MiniInnerVersionNoBuildUploadResponse](upload_response.go#L4)

手淘小程序产物包免构建版本上传

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `build_extra_info` | `BuildExtraInfo` | `*string` | JSON body | 未标注 | 扩展信息，比如adaptorName、tinycliVersion、tinycliName、import-module、allowPrecompile、extJson、allowInstallDependency、aggregationMainAppId，如果没有特殊要求，tinycliVersion版本请用最新的： https://registry.npm.alibaba-inc.com/@alipay/tiny-cli/huoban-prod |
| `build_js_permission` | `BuildJsPermission` | `*string` | JSON body | 未标注 | jsapi权限文件 |
| `build_main_url` | `BuildMainUrl` | `*string` | JSON body | 未标注 | 小程序页面主入口。/index.html#page/component/index，该值需要和extendInfo中的page参数保持一致 |
| `build_qcloud_info` | `BuildQcloudInfo` | `*string` | JSON body | 未标注 | 打包平台扩展信息 |
| `build_version` | `BuildVersion` | `*string` | JSON body | 未标注 | 小程序版本 |
| `builded_package_size` | `BuildedPackageSize` | `*string` | JSON body | 未标注 | 已经构建过的amr包大小，单位是字节，建议如实填写 |
| `builded_package_url` | `BuildedPackageUrl` | `*string` | JSON body | 未标注 | 已经构建的包地址，目前主要自行构建的场景使用 |
| `builded_plugin_size` | `BuildedPluginSize` | `*string` | JSON body | 未标注 | 构建好的插件包amr大小 |
| `builded_plugin_url` | `BuildedPluginUrl` | `*string` | JSON body | 未标注 | 构建好的插件包地址 |
| `bundle_id` | `BundleId` | `*string` | JSON body | 未标注 | 一个端的标识，用于区分不同的客户端，每接入一个客户端，都需要向小程序应用中心申请bundleId入驻 |
| `components` | `Components` | `*string` | JSON body | 未标注 | 组件参数 |
| `inst_code` | `InstCode` | `*string` | JSON body | 未标注 | 上传调试版的接入租户类型 |
| `mini_app_id` | `MiniAppId` | `*string` | JSON body | 未标注 | 小程序id |
| `new_builded_package_size` | `NewBuildedPackageSize` | `*string` | JSON body | 未标注 | 2.0已经构建过的amr包大小，单位是字节，主要是给自行构建的场景使用 |
| `new_builded_package_url` | `NewBuildedPackageUrl` | `*string` | JSON body | 未标注 | 2.0产物包地址 |
| `new_builded_plugin_size` | `NewBuildedPluginSize` | `*string` | JSON body | 未标注 | 构建好的插件包amr大小 |
| `new_builded_plugin_url` | `NewBuildedPluginUrl` | `*string` | JSON body | 未标注 | 构建好的插件包地址 |
| `no_sign` | `NoSign` | `*bool` | JSON body | 未标注 | 否免加签, true的时候代表传过来的是加签后的包地址 |
| `plugin_refs` | `PluginRefs` | [`*[]MiniAppPluginReference`](#model-miniapppluginreference) | JSON body | 未标注 | 小程序代码中引用的插件列表，包含插件id和插件版本信息，业务方需要自行解析源码包的app.json里面的plugins信息 |
| `sub_packages` | `SubPackages` | [`*[]SubPackageInfo`](#model-subpackageinfo) | JSON body | 未标注 | 构建好的分包信息 |

<a id="api-miniinnerversiontemplatedupload"></a>

### MiniInnerVersionTemplatedUpload — 基于模板构建版本

- 官方接口：[`alipay.open.mini.innerversion.templated.upload`](https://opendocs.alipay.com/pre-apis/05xdzq)
- 路由：`POST /v3/alipay/open/mini/innerversion/templated/upload`
- 请求类型：[MiniInnerVersionTemplatedUploadRequest](upload_request.go#L48)
- 响应类型：[MiniInnerVersionTemplatedUploadResponse](upload_response.go#L8)

基于模板构建实例化小程序，同步接口，调用成功表示创建版本成功

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `app_origin` | `AppOrigin` | `*string` | JSON body | 未标注 | 来源类型，新接入方需要向支付宝申请专用来源，否则不予接入。 |
| `app_version` | `AppVersion` | `*string` | JSON body | 未标注 | 格式为: x.y.z，其中x、y、z均为整型数字，版本不能重复构建，且新版本号必须高于旧版本 |
| `bundle_id` | `BundleId` | `*string` | JSON body | 未标注 | 端id |
| `ext` | `Ext` | `*string` | JSON body | 未标注 | 模板的配置参数，完整内容请参见：https://docs.alipay.com/mini/isv/ortga1 |
| `mini_app_id` | `MiniAppId` | `*string` | JSON body | 未标注 | 小程序ID，特殊场景专用，普通业务方无需关注该参数。 |
| `template_id` | `TemplateId` | `*string` | JSON body | 未标注 | 模板id，模板必须已经处于上架/下架状态 |
| `template_version` | `TemplateVersion` | `*string` | JSON body | 未标注 | 模板版本，必须处于上架/下架状态。如果为空，默认取当前在架，或者最新下架的版本。 |

<a id="api-miniinnerversionupload"></a>

### MiniInnerVersionUpload — 内部小程序-打包构建

- 官方接口：[`alipay.open.mini.innerversion.upload`](https://opendocs.alipay.com/pre-apis/05xhot)
- 路由：`POST /v3/alipay/open/mini/innerversion/upload`
- 请求类型：[MiniInnerVersionUploadRequest](upload_request.go#L66)
- 响应类型：[MiniInnerVersionUploadResponse](upload_response.go#L12)

内部小程序，创建打包构建任务，异步返回，使用查询接口轮询构建结果

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `app_origin` | `AppOrigin` | `*string` | JSON body | 未标注 | 业务来源，新接入方需要向支付宝申请专用来源，否则不予接入。https://yuque.antfin-inc.com/tinyapp-all/qddncu/bidushixiang |
| `build_app_type` | `BuildAppType` | `*string` | JSON body | 未标注 | IDE开发打包类型，写死tinyApp |
| `build_extra_info` | `BuildExtraInfo` | `*string` | JSON body | 未标注 | 扩展信息，比如adaptorName、tinycliVersion、tinycliName、import-module、allowPrecompile、extJson、allowInstallDependency、aggregationMainAppId，如果没有特殊要求，tinycliVersion版本请用最新的： https://registry.npm.alibaba-inc.com/@alipay/tiny-cli/huoban-prod |
| `build_extra_mini_project_config` | `BuildExtraMiniProjectConfig` | `*string` | JSON body | 未标注 | 用于覆盖开发者配置字段，实现止血和灰度逻辑 |
| `build_js_permission` | `BuildJsPermission` | `*string` | JSON body | 未标注 | jsapi 权限文件,钉钉外请勿传 |
| `build_main_url` | `BuildMainUrl` | `*string` | JSON body | 未标注 | 小程序页面主入口。/index.html#page/component/index，该值需要和extendInfo中的page参数保持一致 |
| `build_max_android_client_version` | `BuildMaxAndroidClientVersion` | `*string` | JSON body | 未标注 | 最大Android客户端版本号，禁止使用，如果需要使用请联系小程序平台评估，否则端上会出现问题 |
| `build_max_ios_client_version` | `BuildMaxIosClientVersion` | `*string` | JSON body | 未标注 | 最大iOS客户单版本号，禁止使用，如果需要使用请联系小程序平台评估，否则端上会出现问题 |
| `build_min_android_client_version` | `BuildMinAndroidClientVersion` | `*string` | JSON body | 未标注 | 最小Android客户端版本号，禁止使用，如果需要使用请联系小程序平台评估，否则端上会出现问题 |
| `build_min_ios_client_version` | `BuildMinIosClientVersion` | `*string` | JSON body | 未标注 | 最小iOS客户单版本号，禁止使用，如果需要使用请联系小程序平台评估，否则端上会出现问题 |
| `build_package_md_5` | `BuildPackageMd5` | `*string` | JSON body | 未标注 | 源码包文件流的MD5字符串 |
| `build_package_name` | `BuildPackageName` | `*string` | JSON body | 未标注 | 包名称 |
| `build_package_stream` | `BuildPackageStream` | `*string` | JSON body | 未标注 | base64编码后的小程序源码包文件流 |
| `build_qcloud_info` | `BuildQcloudInfo` | `*string` | JSON body | 未标注 | 打包平台扩展信息 |
| `build_signed_pkg_url` | `BuildSignedPkgUrl` | `*string` | JSON body | 未标注 | 已经构建加签过的包地址 |
| `build_source_pkg_size` | `BuildSourcePkgSize` | `*string` | JSON body | 未标注 | 源码包大小 |
| `build_source_pkg_url` | `BuildSourcePkgUrl` | `*string` | JSON body | 未标注 | 源码包地址，源码包大小不能超过20M且地址必须永久有效，若超20M，将限制业务来源调用。将源码文件夹整个用zip格式压缩成后缀amr文件传至文件服务器 |
| `build_sub_url` | `BuildSubUrl` | `*string` | JSON body | 未标注 | 子入口 |
| `build_version` | `BuildVersion` | `*string` | JSON body | 未标注 | 小程序版本 |
| `builded_package_size` | `BuildedPackageSize` | `*string` | JSON body | 未标注 | 已经构建过的包大小，单位是字节，主要是给自行构建的场景使用 |
| `builded_package_url` | `BuildedPackageUrl` | `*string` | JSON body | 未标注 | 已经构建的包地址，目前主要自行构建的场景使用 |
| `bundle_id` | `BundleId` | `*string` | JSON body | 未标注 | 一个端的标识，用于区分不同的客户端，每接入一个客户端，都需要向小程序应用中心申请bundleId入驻 |
| `client_type` | `ClientType` | `*string` | JSON body | 未标注 | 多端类型，除了钉钉外其他业务方禁止使用 |
| `inst_code` | `InstCode` | `*string` | JSON body | 未标注 | 上传调试版的接入租户类型。 |
| `mini_app_id` | `MiniAppId` | `*string` | JSON body | 未标注 | 小程序ID |
| `plugin_refs` | `PluginRefs` | [`*[]MiniAppPluginReference`](#model-miniapppluginreference) | JSON body | 未标注 | 小程序代码中引用的插件列表，包含插件id和插件版本信息，业务方需要自行解析源码包的app.json里面的plugins信息 |

<a id="api-miniinnerversionuploadstatusquery"></a>

### MiniInnerVersionUploadStatusQuery — 多端小程序-查询构建状态

- 官方接口：[`alipay.open.mini.innerversion.uploadstatus.query`](https://opendocs.alipay.com/pre-apis/05xskn)
- 路由：`GET /v3/alipay/open/mini/innerversion/uploadstatus/query`
- 请求类型：[MiniInnerVersionUploadStatusQueryRequest](upload_request.go#L122)
- 响应类型：[MiniInnerVersionUploadStatusQueryResponse](upload_response.go#L32)

查询构建状态

请求约束：MiniInnerVersionUploadStatusQueryRequest is the request type for alipay.open.mini.innerversion.uploadstatus.query.

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `build_package_id` | `BuildPackageId` | `*string` | query | 否 | 包管理ID，用于查询打包状态 |
| `mini_app_id` | `MiniAppId` | `*string` | query | 否 | 小程序ID |
| `build_version` | `BuildVersion` | `*string` | query | 否 | 小程序版本 |
| `inst_code` | `InstCode` | `*string` | query | 否 | 租户id |
| `bundle_id` | `BundleId` | `*string` | query | 否 | 端信息 |

## 嵌套请求参数

<a id="model-miniapppluginreference"></a>

### MiniAppPluginReference

类型定义：[MiniAppPluginReference](upload_model.go#L4)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `lazy` | `Lazy` | `*bool` | 未标注 | 是否懒加载 |
| `plugin_id` | `PluginId` | `*string` | 未标注 | 插件id，即插件的唯一标识，形如小程序的appId |
| `plugin_lazy` | `PluginLazy` | `*bool` | 未标注 | 是否懒加载 |
| `plugin_version` | `PluginVersion` | `*string` | 未标注 | 小程序引用的插件版本，可以指定插件版本，如：1.0.11；也可以填*，表示引用最新的插件版本，此时，当插件发布新版本后，小程序内的插件会自动更新到最新的版本。 |

<a id="model-subpackageinfo"></a>

### SubPackageInfo

类型定义：[SubPackageInfo](upload_model.go#L16)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `builded_package_url` | `BuildedPackageUrl` | `*string` | 未标注 | 构建好的结果地址 |
| `new_builded_package_url` | `NewBuildedPackageUrl` | `*string` | 未标注 | 构建好的分包地址 |
| `new_size` | `NewSize` | `*string` | 未标注 | 2.0分包大小，单位字节 |
| `path` | `Path` | `*string` | 未标注 | 分包路径 |
| `size` | `Size` | `*string` | 未标注 | 分包大小，单位字节 |
| `source_url` | `SourceUrl` | `*string` | 未标注 | 源码地址 |
| `type` | `Type` | `*string` | 未标注 | 分包类型： MAIN ｜ SUB |
