# sign 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/open/agent/sign`。本包提供 6 个支付宝 V3 REST 操作。

通过 `sign.NewService(executor)` 复用 `*core.Client` 或最小 `core.Executor`。网络调用以 `context.Context` 为首参数，支持 `opts ...core.CallOption`，业务数据位于 `resp.Data`。

接口方法按单词边界使用大驼峰命名；请求、响应与专属模型位于对应业务包，共享模型从共同父包引用。

接入方式见[项目说明](../../../../README.md)。默认测试离线执行。

返回[领域目录](../README.md)。

## 接口目录

| SDK 方法 | 官方接口 |
| --- | --- |
| [`AgentCommonSign`](#api-agentcommonsign) | `alipay.open.agent.common.sign` |
| [`AgentCommonSignConfirm`](#api-agentcommonsignconfirm) | `alipay.open.agent.commonsign.confirm` |
| [`AgentFaceToFaceSign`](#api-agentfacetofacesign) | `alipay.open.agent.facetoface.sign` |
| [`AgentMobilePaySign`](#api-agentmobilepaysign) | `alipay.open.agent.mobilepay.sign` |
| [`AgentOfflinePaymentSign`](#api-agentofflinepaymentsign) | `alipay.open.agent.offlinepayment.sign` |
| [`AgentSignStatusQuery`](#api-agentsignstatusquery) | `alipay.open.agent.signstatus.query` |

## 使用示例

- [AgentCommonSign](api_sign_agentcommonsign_example_test.go)
- [AgentCommonSignConfirm](api_sign_agentcommonsignconfirm_example_test.go)
- [AgentFaceToFaceSign](api_sign_agentfacetofacesign_example_test.go)
- [AgentMobilePaySign](api_sign_agentmobilepaysign_example_test.go)
- [AgentOfflinePaymentSign](api_sign_agentofflinepaymentsign_example_test.go)
- [AgentSignStatusQuery](api_sign_agentsignstatusquery_example_test.go)

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-agentcommonsign"></a>

### AgentCommonSign — 代签约产品通用接口

- 官方接口：[`alipay.open.agent.common.sign`](https://opendocs.alipay.com/open-v3/f0ab46ad_alipay.open.agent.common.sign)
- 路由：`POST /v3/alipay/open/agent/common/sign`
- 请求类型：[AgentCommonSignRequest](sign_request.go#L22)
- 响应类型：[AgentCommonSignResponse](sign_response.go#L4)

三方应用代理签约产品，需要配合开启事务接口使用

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `app_auth_pic` | `AppAuthPic` | [`*core.MultipartFile`](../../../../core/multipart.go) | multipart | 未标注 | 上传的文件；使用 core.MultipartFile 指定文件名、内容类型和内容字节。 |
| `app_demo` | `AppDemo` | [`*core.MultipartFile`](../../../../core/multipart.go) | multipart | 未标注 | 上传的文件；使用 core.MultipartFile 指定文件名、内容类型和内容字节。 |
| `app_home_screenshot` | `AppHomeScreenshot` | [`*core.MultipartFile`](../../../../core/multipart.go) | multipart | 未标注 | 上传的文件；使用 core.MultipartFile 指定文件名、内容类型和内容字节。 |
| `app_item_screenshot` | `AppItemScreenshot` | [`*core.MultipartFile`](../../../../core/multipart.go) | multipart | 未标注 | 上传的文件；使用 core.MultipartFile 指定文件名、内容类型和内容字节。 |
| `app_pay_screenshot` | `AppPayScreenshot` | [`*core.MultipartFile`](../../../../core/multipart.go) | multipart | 未标注 | 上传的文件；使用 core.MultipartFile 指定文件名、内容类型和内容字节。 |
| `business_license_auth_pic` | `BusinessLicenseAuthPic` | [`*core.MultipartFile`](../../../../core/multipart.go) | multipart | 未标注 | 上传的文件；使用 core.MultipartFile 指定文件名、内容类型和内容字节。 |
| `business_license_pic` | `BusinessLicensePic` | [`*core.MultipartFile`](../../../../core/multipart.go) | multipart | 未标注 | 上传的文件；使用 core.MultipartFile 指定文件名、内容类型和内容字节。 |
| `data` | `Data` | [`*AgentCommonSign`](#model-agentcommonsign) | multipart | 未标注 | 见字段类型及官方接口文档。 |
| `h_5_extra_pic` | `H5ExtraPic` | [`*core.MultipartFile`](../../../../core/multipart.go) | multipart | 未标注 | 上传的文件；使用 core.MultipartFile 指定文件名、内容类型和内容字节。 |
| `h_5_home_screenshot` | `H5HomeScreenshot` | [`*core.MultipartFile`](../../../../core/multipart.go) | multipart | 未标注 | 上传的文件；使用 core.MultipartFile 指定文件名、内容类型和内容字节。 |
| `h_5_item_screenshot` | `H5ItemScreenshot` | [`*core.MultipartFile`](../../../../core/multipart.go) | multipart | 未标注 | 上传的文件；使用 core.MultipartFile 指定文件名、内容类型和内容字节。 |
| `h_5_pay_screenshot` | `H5PayScreenshot` | [`*core.MultipartFile`](../../../../core/multipart.go) | multipart | 未标注 | 上传的文件；使用 core.MultipartFile 指定文件名、内容类型和内容字节。 |
| `h_5_sites_loa` | `H5SitesLoa` | [`*core.MultipartFile`](../../../../core/multipart.go) | multipart | 未标注 | 上传的文件；使用 core.MultipartFile 指定文件名、内容类型和内容字节。 |
| `mini_app_screenshot` | `MiniAppScreenshot` | [`*core.MultipartFile`](../../../../core/multipart.go) | multipart | 未标注 | 上传的文件；使用 core.MultipartFile 指定文件名、内容类型和内容字节。 |
| `shop_scene_pic` | `ShopScenePic` | [`*core.MultipartFile`](../../../../core/multipart.go) | multipart | 未标注 | 上传的文件；使用 core.MultipartFile 指定文件名、内容类型和内容字节。 |
| `shop_sign_board_pic` | `ShopSignBoardPic` | [`*core.MultipartFile`](../../../../core/multipart.go) | multipart | 未标注 | 上传的文件；使用 core.MultipartFile 指定文件名、内容类型和内容字节。 |
| `special_license_pic` | `SpecialLicensePic` | [`*core.MultipartFile`](../../../../core/multipart.go) | multipart | 未标注 | 上传的文件；使用 core.MultipartFile 指定文件名、内容类型和内容字节。 |
| `web_home_screenshot` | `WebHomeScreenshot` | [`*core.MultipartFile`](../../../../core/multipart.go) | multipart | 未标注 | 上传的文件；使用 core.MultipartFile 指定文件名、内容类型和内容字节。 |
| `web_item_screenshot` | `WebItemScreenshot` | [`*core.MultipartFile`](../../../../core/multipart.go) | multipart | 未标注 | 上传的文件；使用 core.MultipartFile 指定文件名、内容类型和内容字节。 |
| `web_pay_screenshot` | `WebPayScreenshot` | [`*core.MultipartFile`](../../../../core/multipart.go) | multipart | 未标注 | 上传的文件；使用 core.MultipartFile 指定文件名、内容类型和内容字节。 |
| `web_sites_loa` | `WebSitesLoa` | [`*core.MultipartFile`](../../../../core/multipart.go) | multipart | 未标注 | 上传的文件；使用 core.MultipartFile 指定文件名、内容类型和内容字节。 |

注意事项：

- 文件使用 core.MultipartFile；data 部分使用表中具名结构体。SDK 负责 multipart 编码，不能将整个请求改为普通 JSON。

<a id="api-agentcommonsignconfirm"></a>

### AgentCommonSignConfirm — 代商户签约，提交信息确认接口

- 官方接口：[`alipay.open.agent.commonsign.confirm`](https://opendocs.alipay.com/open-v3/2f809b83_alipay.open.agent.commonsign.confirm)
- 路由：`POST /v3/alipay/open/agent/commonsign/confirm`
- 请求类型：[AgentCommonSignConfirmRequest](sign_request.go#L8)
- 响应类型：[AgentCommonSignConfirmResponse](sign_response.go#L8)

三方应用代理签约产品，上传商户信息后，确认提交信息。可以配合开启事务接口使用。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `batch_no` | `BatchNo` | `*string` | JSON body | 未标注 | ISV 代商户操作事务编号，通过事务开启接口alipay.open.agent.create调用返回。 |

<a id="api-agentfacetofacesign"></a>

### AgentFaceToFaceSign — 代签约当面付产品

- 官方接口：[`alipay.open.agent.facetoface.sign`](https://opendocs.alipay.com/open-v3/e1eb4236_alipay.open.agent.facetoface.sign)
- 路由：`POST /v3/alipay/open/agent/facetoface/sign`
- 请求类型：[AgentFaceToFaceSignRequest](sign_request.go#L47)
- 响应类型：[AgentFaceToFaceSignResponse](sign_response.go#L14)

三方应用代理签约当面付产品，需要配合开启事务接口使用

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `business_license_auth_pic` | `BusinessLicenseAuthPic` | [`*core.MultipartFile`](../../../../core/multipart.go) | multipart | 未标注 | 上传的文件；使用 core.MultipartFile 指定文件名、内容类型和内容字节。 |
| `business_license_pic` | `BusinessLicensePic` | [`*core.MultipartFile`](../../../../core/multipart.go) | multipart | 未标注 | 上传的文件；使用 core.MultipartFile 指定文件名、内容类型和内容字节。 |
| `data` | `Data` | [`*AgentFaceToFaceSign`](#model-agentfacetofacesign) | multipart | 未标注 | 见字段类型及官方接口文档。 |
| `shop_scene_pic` | `ShopScenePic` | [`*core.MultipartFile`](../../../../core/multipart.go) | multipart | 未标注 | 上传的文件；使用 core.MultipartFile 指定文件名、内容类型和内容字节。 |
| `shop_sign_board_pic` | `ShopSignBoardPic` | [`*core.MultipartFile`](../../../../core/multipart.go) | multipart | 未标注 | 上传的文件；使用 core.MultipartFile 指定文件名、内容类型和内容字节。 |
| `special_license_pic` | `SpecialLicensePic` | [`*core.MultipartFile`](../../../../core/multipart.go) | multipart | 未标注 | 上传的文件；使用 core.MultipartFile 指定文件名、内容类型和内容字节。 |

注意事项：

- 文件使用 core.MultipartFile；data 部分使用表中具名结构体。SDK 负责 multipart 编码，不能将整个请求改为普通 JSON。

<a id="api-agentmobilepaysign"></a>

### AgentMobilePaySign — 代签约APP支付产品

- 官方接口：[`alipay.open.agent.mobilepay.sign`](https://opendocs.alipay.com/open-v3/db21f6cb_alipay.open.agent.mobilepay.sign)
- 路由：`POST /v3/alipay/open/agent/mobilepay/sign`
- 请求类型：[AgentMobilePaySignRequest](sign_request.go#L57)
- 响应类型：[AgentMobilePaySignResponse](sign_response.go#L18)

三方应用代理签约APP支付产品，需要配合开启事务接口使用

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `app_auth_pic` | `AppAuthPic` | [`*core.MultipartFile`](../../../../core/multipart.go) | multipart | 未标注 | 上传的文件；使用 core.MultipartFile 指定文件名、内容类型和内容字节。 |
| `app_demo` | `AppDemo` | [`*core.MultipartFile`](../../../../core/multipart.go) | multipart | 未标注 | 上传的文件；使用 core.MultipartFile 指定文件名、内容类型和内容字节。 |
| `app_home_screenshot` | `AppHomeScreenshot` | [`*core.MultipartFile`](../../../../core/multipart.go) | multipart | 未标注 | 上传的文件；使用 core.MultipartFile 指定文件名、内容类型和内容字节。 |
| `app_item_screenshot` | `AppItemScreenshot` | [`*core.MultipartFile`](../../../../core/multipart.go) | multipart | 未标注 | 上传的文件；使用 core.MultipartFile 指定文件名、内容类型和内容字节。 |
| `app_pay_screenshot` | `AppPayScreenshot` | [`*core.MultipartFile`](../../../../core/multipart.go) | multipart | 未标注 | 上传的文件；使用 core.MultipartFile 指定文件名、内容类型和内容字节。 |
| `business_license_auth_pic` | `BusinessLicenseAuthPic` | [`*core.MultipartFile`](../../../../core/multipart.go) | multipart | 未标注 | 上传的文件；使用 core.MultipartFile 指定文件名、内容类型和内容字节。 |
| `business_license_pic` | `BusinessLicensePic` | [`*core.MultipartFile`](../../../../core/multipart.go) | multipart | 未标注 | 上传的文件；使用 core.MultipartFile 指定文件名、内容类型和内容字节。 |
| `data` | `Data` | [`*AgentMobilePaySign`](#model-agentmobilepaysign) | multipart | 未标注 | 见字段类型及官方接口文档。 |
| `home_screenshot` | `HomeScreenshot` | [`*core.MultipartFile`](../../../../core/multipart.go) | multipart | 未标注 | 上传的文件；使用 core.MultipartFile 指定文件名、内容类型和内容字节。 |
| `in_app_screenshot` | `InAppScreenshot` | [`*core.MultipartFile`](../../../../core/multipart.go) | multipart | 未标注 | 上传的文件；使用 core.MultipartFile 指定文件名、内容类型和内容字节。 |
| `pay_screenshot` | `PayScreenshot` | [`*core.MultipartFile`](../../../../core/multipart.go) | multipart | 未标注 | 上传的文件；使用 core.MultipartFile 指定文件名、内容类型和内容字节。 |
| `special_license_pic` | `SpecialLicensePic` | [`*core.MultipartFile`](../../../../core/multipart.go) | multipart | 未标注 | 上传的文件；使用 core.MultipartFile 指定文件名、内容类型和内容字节。 |

注意事项：

- 文件使用 core.MultipartFile；data 部分使用表中具名结构体。SDK 负责 multipart 编码，不能将整个请求改为普通 JSON。

<a id="api-agentofflinepaymentsign"></a>

### AgentOfflinePaymentSign — 代签约当面付快捷版产品

- 官方接口：[`alipay.open.agent.offlinepayment.sign`](https://opendocs.alipay.com/open-v3/8b80b506_alipay.open.agent.offlinepayment.sign?scene=common&pathHash=255833d8)
- 路由：`POST /v3/alipay/open/agent/offlinepayment/sign`
- 请求类型：[AgentOfflinePaymentSignRequest](sign_request.go#L73)
- 响应类型：[AgentOfflinePaymentSignResponse](sign_response.go#L22)

三方应用代理签约当面付快捷版产品，需要配合开启事务接口使用

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `business_license_pic` | `BusinessLicensePic` | [`*core.MultipartFile`](../../../../core/multipart.go) | multipart | 未标注 | 上传的文件；使用 core.MultipartFile 指定文件名、内容类型和内容字节。 |
| `data` | `Data` | [`*AgentOfflinePaymentSign`](#model-agentofflinepaymentsign) | multipart | 未标注 | 见字段类型及官方接口文档。 |
| `shop_scene_pic` | `ShopScenePic` | [`*core.MultipartFile`](../../../../core/multipart.go) | multipart | 未标注 | 上传的文件；使用 core.MultipartFile 指定文件名、内容类型和内容字节。 |
| `shop_sign_board_pic` | `ShopSignBoardPic` | [`*core.MultipartFile`](../../../../core/multipart.go) | multipart | 未标注 | 上传的文件；使用 core.MultipartFile 指定文件名、内容类型和内容字节。 |
| `special_license_pic` | `SpecialLicensePic` | [`*core.MultipartFile`](../../../../core/multipart.go) | multipart | 未标注 | 上传的文件；使用 core.MultipartFile 指定文件名、内容类型和内容字节。 |

注意事项：

- 文件使用 core.MultipartFile；data 部分使用表中具名结构体。SDK 负责 multipart 编码，不能将整个请求改为普通 JSON。

<a id="api-agentsignstatusquery"></a>

### AgentSignStatusQuery — 查询商户某个产品的签约状态

- 官方接口：[`alipay.open.agent.signstatus.query`](https://opendocs.alipay.com/open-v3/031e92b9_alipay.open.agent.signstatus.query)
- 路由：`POST /v3/alipay/open/agent/signstatus/query`
- 请求类型：[AgentSignStatusQueryRequest](sign_request.go#L14)
- 响应类型：[AgentSignStatusQueryResponse](sign_response.go#L26)

主要用在isv代理商户签约场景，用来查询商户签约某个产品的签约状态，可以配合"第三方接入核心API"功能包里的接口使用

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `pid` | `Pid` | `*string` | JSON body | 未标注 | isv要查询签约状态的商户账号，可以是支付宝账号pid（2088开头），也可以是商户的登录账号（邮箱或手机号） |
| `product_codes` | `ProductCodes` | `*[]string` | JSON body | 未标注 | isv要查询商户签约状态的产品码，产品码是支付宝内部对产品的唯一标识。目前支持查询产品及产品码详情参见 [获取商家产品开通状态](https://opendocs.alipay.com/isv/04eh3g)。 |

## 嵌套请求参数

<a id="model-agentcommonsign"></a>

### AgentCommonSign

类型定义：[AgentCommonSign](sign_model.go#L4)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `alipay_life_name` | `AlipayLifeName` | `*string` | 未标注 | 支付宝生活号(原服务窗)名称（1 app_name、app_demo；2 web_sites；3 alipay_life_name。1、2、3少选择一个填写） |
| `app_market` | `AppMarket` | `*string` | 未标注 | 投放应用市场 |
| `app_name` | `AppName` | `*string` | 未标注 | 商户的APP应用名称（1 app_name、app_demo；2 web_sites；3 alipay_life_name；4 wechat_official_account_name。1、2、3、4至少选择一个填写） |
| `app_status` | `AppStatus` | `*string` | 未标注 | 应用状态 |
| `app_test_account` | `AppTestAccount` | `*string` | 未标注 | 可以登录该应用的测试账户 |
| `app_test_account_password` | `AppTestAccountPassword` | `*string` | 未标注 | 可以登录此应用的账户的密码。对app_test_account的登录密码。 |
| `batch_no` | `BatchNo` | `*string` | 未标注 | 代商户操作事务编号，通过alipay.open.agent.create接口进行创建。 |
| `business_license_no` | `BusinessLicenseNo` | `*string` | 未标注 | 营业执照号码 |
| `date_limitation` | `DateLimitation` | `*string` | 未标注 | 营业期限 |
| `download_link` | `DownloadLink` | `*string` | 未标注 | 应用下载链接 |
| `h_5_memo` | `H5Memo` | `*string` | 未标注 | 网站扩展信息文本 |
| `h_5_sites` | `H5Sites` | `*string` | 未标注 | H5网站地址 |
| `h_5_status` | `H5Status` | `*string` | 未标注 | H5网站状态 |
| `life_app_id` | `LifeAppId` | `*string` | 未标注 | 请输入被签约商户的支付宝生活号ID编号 |
| `long_term` | `LongTerm` | `*bool` | 未标注 | 营业期限是否长期有效 |
| `mcc_code` | `MccCode` | `*string` | 未标注 | [商家经营类目2.0](https://opendocs.alipay.com/common/02khjv) 中的“经营类目编码” |
| `mini_app_id` | `MiniAppId` | `*string` | 未标注 | 请输入被签约商家下任意小程序APPID ID编号 |
| `mini_app_name` | `MiniAppName` | `*string` | 未标注 | 小程序名称 |
| `mobile_type` | `MobileType` | `*string` | 未标注 | 移动设备类型，当开通电脑网站支付产品时不需赋值；当开通手机网站支付时，传H5；当开通JSAPI时，传APP；当开通预授权支付产品，且交易场景选择“mobile”(移动设备)时，需要从APP,H5中至少选择一个，可多选，字符串英文逗号分隔。 |
| `product_code` | `ProductCode` | `*string` | 未标注 | isv要代商户签约产品码，产品码是支付宝内部对产品的唯一标识 |
| `rate` | `Rate` | `*string` | 未标注 | 服务费率（%），例如：请输入0.38~0.6 之间（小数点后两位，可取0.38%及0.6%），直播买卖通涉及JSAPI支付产品，交易手续费为 0.38%-0.6%，后续费率变化可参考支付宝标准服务费。 当签约且授权标识 sign_and_auth=true 时，该费率信息必填。 |
| `shop_city_code` | `ShopCityCode` | `*string` | 未标注 | 门店所在城市。请按照[地区码表格](https://mdn.alipayobjects.com/huamei_sm7gq8/afts/file/A*blT9RqSR9_gAAAAAAAAAAAAADuKQAQ/2022%E8%9A%82%E8%9A%81%E9%87%91%E6%9C%8D%E5%9C%B0%E5%8C%BA%E7%A0%81.xlsx) 中内容填写。 |
| `shop_detail_address` | `ShopDetailAddress` | `*string` | 未标注 | 门店详细地址 |
| `shop_district_code` | `ShopDistrictCode` | `*string` | 未标注 | 门店所在区/县。请按照[地区码表格](https://mdn.alipayobjects.com/huamei_sm7gq8/afts/file/A*blT9RqSR9_gAAAAAAAAAAAAADuKQAQ/2022%E8%9A%82%E8%9A%81%E9%87%91%E6%9C%8D%E5%9C%B0%E5%8C%BA%E7%A0%81.xlsx) 中内容填写。 |
| `shop_name` | `ShopName` | `*string` | 未标注 | 门店名称 |
| `shop_province_code` | `ShopProvinceCode` | `*string` | 未标注 | 门店所在省份。请按照[地区码表格](https://mdn.alipayobjects.com/huamei_sm7gq8/afts/file/A*blT9RqSR9_gAAAAAAAAAAAAADuKQAQ/2022%E8%9A%82%E8%9A%81%E9%87%91%E6%9C%8D%E5%9C%B0%E5%8C%BA%E7%A0%81.xlsx) 中内容填写。 |
| `special_license_type` | `SpecialLicenseType` | `*string` | 未标注 | 合约特殊资质的类型，不传默认99 |
| `trade_scene` | `TradeScene` | `*string` | 未标注 | 交易场景，当开通产品为电脑网站支付时，传pc；当开通产品为手机网站支付时，传mobile；当开通产品为JSAPI时，传tiny_app；当开通产品为预授权支付时，下面括号中三个至少传一个（tiny_app,mobile,offline_store），可多选，字符串用英文逗号分隔 |
| `web_sites` | `WebSites` | `*[]string` | 未标注 | 接入网址信息（1 app_name、app_demo；2 web_sites；3 alipay_life_name；4 wechat_official_account_name。1、2、3、4至少选择一个填写） |
| `web_status` | `WebStatus` | `*string` | 未标注 | 网站状态，枚举值为：ONLINE，OFFLINE |
| `web_test_account` | `WebTestAccount` | `*string` | 未标注 | 可以登录此网站的测试账户。 |
| `web_test_account_password` | `WebTestAccountPassword` | `*string` | 未标注 | 可以登录此网站的账户的密码。对应web_test_account的登录密码 |
| `wechat_official_account_name` | `WechatOfficialAccountName` | `*string` | 未标注 | 1 app_name、app_demo；2 web_sites；3 alipay_life_name；4 wechat_official_account_name。1、2、3、4至少选择一个填写(该字段已经不在使用，废弃，不用填写) |

<a id="model-agentfacetofacesign"></a>

### AgentFaceToFaceSign

类型定义：[AgentFaceToFaceSign](sign_model.go#L74)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `batch_no` | `BatchNo` | `*string` | 未标注 | 代商户操作事务编号，通过[接口地址](https://opendocs.alipay.com/apis/api_50/alipay.open.agent.create) (开启代商户签约、创建应用事务)接口进行事务创建后获取。 |
| `business_license_mobile` | `BusinessLicenseMobile` | `*string` | 未标注 | 上传非经营者本人（签约账号）的营业执照时必填，填写营业执照法人支付宝账号对应的手机号。 |
| `business_license_no` | `BusinessLicenseNo` | `*string` | 未标注 | 营业执照号码 |
| `date_limitation` | `DateLimitation` | `*string` | 未标注 | 营业期限 |
| `long_term` | `LongTerm` | `*bool` | 未标注 | 营业期限是否长期有效 |
| `mcc_code` | `MccCode` | `*string` | 未标注 | 商家经营类目编码。参见 [商家经营类目2.0](https://opendocs.alipay.com/open/01n22g#%E5%95%86%E5%AE%B6%E7%BB%8F%E8%90%A5%E7%B1%BB%E7%9B%AE) 中的“一级类目code_二级类目code”。 |
| `rate` | `Rate` | `*string` | 未标注 | 服务费率（%），0.38~0.6 之间（小数点后两位，可取0.38%及0.6%）。 当签约且授权标识 sign_and_auth=true 时，该费率信息必填。 |
| `shop_address` | `ShopAddress` | [`*SignAddressInfo`](#model-signaddressinfo) | 未标注 | 见字段类型及官方接口文档。 |
| `shop_name` | `ShopName` | `*string` | 未标注 | 店铺名称 |
| `sign_and_auth` | `SignAndAuth` | `*bool` | 未标注 | 签约且授权标识，默认为false，只进行签约操作； 如果设置为true，则表示签约成功后，会自动进行应用授权操作。 |

<a id="model-agentmobilepaysign"></a>

### AgentMobilePaySign

类型定义：[AgentMobilePaySign](sign_model.go#L97)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `app_market` | `AppMarket` | `*[]string` | 未标注 | 应用在哪些市场上架，枚举值为：苹果,应用宝,华为,360,小米,豌豆荚,其他 |
| `app_name` | `AppName` | `*string` | 未标注 | 商户的APP应用名称 |
| `app_status` | `AppStatus` | `*string` | 未标注 | 应用上架状态，枚举值为 已上线，未上线 |
| `app_test_account` | `AppTestAccount` | `*string` | 未标注 | 可以登录该应用的测试账户 |
| `app_test_account_password` | `AppTestAccountPassword` | `*string` | 未标注 | 可以登录此应用的账户的密码。对应app_test_account的登录密码 |
| `app_type` | `AppType` | `*[]string` | 未标注 | 应用类型，枚举值为：IOS，安卓 |
| `batch_no` | `BatchNo` | `*string` | 未标注 | 代商户操作事务编号，通过alipay.open.isv.agent.create接口进行创建。 |
| `business_license_mobile` | `BusinessLicenseMobile` | `*string` | 未标注 | 营业执照法人手机号码 |
| `business_license_no` | `BusinessLicenseNo` | `*string` | 未标注 | 营业执照号码。 |
| `date_limitation` | `DateLimitation` | `*string` | 未标注 | 营业期限 |
| `download_link` | `DownloadLink` | `*string` | 未标注 | 应用下载链接 |
| `long_term` | `LongTerm` | `*bool` | 未标注 | 营业期限是否长期有效 |
| `mcc_code` | `MccCode` | `*string` | 未标注 | 所属MCCCode，详情可参考 [商家经营类目](https://opendocs.alipay.com/open/01n22g#%E5%95%86%E5%AE%B6%E7%BB%8F%E8%90%A5%E7%B1%BB%E7%9B%AE) 中的“经营类目编码” |
| `mobile_type` | `MobileType` | `*string` | 未标注 | 传参：APP，代表设备类型是APP |
| `trade_scene` | `TradeScene` | `*string` | 未标注 | 传参：mobile，代表交易场景是移动设备 |

<a id="model-agentofflinepaymentsign"></a>

### AgentOfflinePaymentSign

类型定义：[AgentOfflinePaymentSign](sign_model.go#L131)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `batch_no` | `BatchNo` | `*string` | 未标注 | 代商户操作事务编号，通过alipay.open.agent.create接口进行创建。 |
| `business_license_mobile` | `BusinessLicenseMobile` | `*string` | 未标注 | 被邀请授权的营业执照法人手机号码，上传非同人营业执照时必填 |
| `business_license_no` | `BusinessLicenseNo` | `*string` | 未标注 | 营业执照号码。若填写，请与以下营业执照图片、期限、一起提供。 |
| `date_limitation` | `DateLimitation` | `*string` | 未标注 | 营业期限 |
| `delivery_address` | `DeliveryAddress` | [`*SignAddressInfo`](#model-signaddressinfo) | 未标注 | 见字段类型及官方接口文档。 |
| `long_term` | `LongTerm` | `*bool` | 未标注 | 营业期限是否长期有效 |
| `mcc_code` | `MccCode` | `*string` | 未标注 | 所属MCCCode，可参考 [商家经营类目](https://opendocs.alipay.com/open/01n22g#%E5%95%86%E5%AE%B6%E7%BB%8F%E8%90%A5%E7%B1%BB%E7%9B%AE) 中的“经营类目编码” |
| `rate` | `Rate` | `*string` | 未标注 | 服务费率（%），0.38~3之间，精确到0.01 |
| `shop_address` | `ShopAddress` | [`*SignAddressInfo`](#model-signaddressinfo) | 未标注 | 见字段类型及官方接口文档。 |
| `shop_name` | `ShopName` | `*string` | 未标注 | 店铺名称 |

<a id="model-signaddressinfo"></a>

### SignAddressInfo

类型定义：[SignAddressInfo](sign_model.go#L163)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `city_code` | `CityCode` | `*string` | 未标注 | 城市编码。请按照[地区码表格](https://mdn.alipayobjects.com/huamei_sm7gq8/afts/file/A*blT9RqSR9_gAAAAAAAAAAAAADuKQAQ/2022%E8%9A%82%E8%9A%81%E9%87%91%E6%9C%8D%E5%9C%B0%E5%8C%BA%E7%A0%81.xlsx) 中内容填写。 （参考资料： [参考资料](https://www.mca.gov.cn/n156/n186/index.html)） |
| `country_code` | `CountryCode` | `*string` | 未标注 | 国家编码，中国默认：156 |
| `detail_address` | `DetailAddress` | `*string` | 未标注 | 详细地址 |
| `district_code` | `DistrictCode` | `*string` | 未标注 | 区县编码。请按照[地区码表格](https://mdn.alipayobjects.com/huamei_sm7gq8/afts/file/A*blT9RqSR9_gAAAAAAAAAAAAADuKQAQ/2022%E8%9A%82%E8%9A%81%E9%87%91%E6%9C%8D%E5%9C%B0%E5%8C%BA%E7%A0%81.xlsx) 中内容填写。 （参考资料： [参考资料](https://www.mca.gov.cn/n156/n186/index.html)） |
| `latitude` | `Latitude` | `*string` | 未标注 | 纬度，浮点型,小数点后最多保留6位 如需要录入经纬度，请以高德坐标系为准，录入时请确保经纬度参数准确。高德经纬度查询：[查询地址](http://lbs.amap.com/console/show/picker) |
| `longitude` | `Longitude` | `*string` | 未标注 | 经度，浮点型, 小数点后最多保留6位。 如需要录入经纬度，请以高德坐标系为准，录入时请确保经纬度参数准确。高德经纬度查询：[查询地址](http://lbs.amap.com/console/show/picker) |
| `province_code` | `ProvinceCode` | `*string` | 未标注 | 省份编码。请按照[地区码表格](https://mdn.alipayobjects.com/huamei_sm7gq8/afts/file/A*blT9RqSR9_gAAAAAAAAAAAAADuKQAQ/2022%E8%9A%82%E8%9A%81%E9%87%91%E6%9C%8D%E5%9C%B0%E5%8C%BA%E7%A0%81.xlsx) 中内容填写。 （参考资料： [参考资料](https://www.mca.gov.cn/n156/n186/index.html)） |
