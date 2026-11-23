# zoloz/authentication 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/zoloz/authentication`。本目录提供 5 个支付宝开放平台 V3 REST 操作。

## 使用方式

先按[项目接入说明](../../../README.md)创建 `*core.Client`，再通过 `authentication.NewService(client)` 创建 Service；所有网络调用以 `context.Context` 为首个参数，支持 `opts ...core.CallOption`。响应为 `*core.Response[T]`，业务数据位于 `resp.Data`。

请求、响应与嵌套类型的链接指向本目录源码；可参考本目录的 [API 示例](api_authentication_example_test.go)。

示例文件中的 `TestSandboxService_*` 可使用沙箱平台证书执行真实调用。默认 `go test` 保持离线。

## 接口目录

| SDK 方法 | 官方接口 | 用途 | HTTP |
| --- | --- | --- | --- |
| [`CustomerFaceManageCreate`](#api-customerfacemanagecreate) | [`zoloz.authentication.customer.facemanage.create`](https://opendocs.alipay.com/pre-apis/05vda2) | 热脸入库 | `POST` |
| [`CustomerFaceManageDelete`](#api-customerfacemanagedelete) | [`zoloz.authentication.customer.facemanage.delete`](https://opendocs.alipay.com/pre-apis/05vh16) | 热脸出库 | `POST` |
| [`CustomerFTokenQuery`](#api-customerftokenquery) | [`zoloz.authentication.customer.ftoken.query`](https://opendocs.alipay.com/open-v3/ebd5a135_zoloz.authentication.customer.ftoken.query) | 查询刷脸结果信息接口 | `POST` |
| [`CustomerSmilePayInitialize`](#api-customersmilepayinitialize) | [`zoloz.authentication.customer.smilepay.initialize`](https://opendocs.alipay.com/pre-apis/05vfpu) | 人脸刷脸付初始化 | `POST` |
| [`SmilePayInitialize`](#api-smilepayinitialize) | [`zoloz.authentication.smilepay.initialize`](https://opendocs.alipay.com/open-v3/1177dbeb_zoloz.authentication.smilepay.initialize) | 刷脸支付初始化 | `POST` |

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-customerfacemanagecreate"></a>

### CustomerFaceManageCreate — 热脸入库

- 官方接口：[`zoloz.authentication.customer.facemanage.create`](https://opendocs.alipay.com/pre-apis/05vda2)
- 路由：`POST /v3/zoloz/authentication/customer/facemanage/create`
- 请求类型：[CustomerFaceManageCreateRequest](authentication_request.go#L5)
- 响应类型：[CustomerFaceManageCreateResponse](authentication_response.go#L5)

热脸入库

请求约束：CustomerFaceManageCreateRequest is the official V3 ZolozAuthenticationCustomerFacemanageCreateModel schema.

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `areacode` | `Areacode` | `*string` | JSON body | 未标注 | 地域编码 |
| `biz_type` | `BizType` | `*string` | JSON body | 未标注 | 人脸产品能力，预热场景，需要传college，k12，scenic，subway，traffic。具体值视具体场景 |
| `bizscale` | `Bizscale` | `*string` | JSON body | 未标注 | 业务量规模 |
| `brandcode` | `Brandcode` | `*string` | JSON body | 未标注 | 商户品牌 |
| `devicenum` | `Devicenum` | `*string` | JSON body | 未标注 | 商户机具唯一编码，关键参数 人脸预热，传logicGroupId |
| `extinfo` | `Extinfo` | `*string` | JSON body | 未标注 | 拓展参数 |
| `facetype` | `Facetype` | `*string` | JSON body | 未标注 | 入库类型 IDCARD:身份证 ALIPAY_USER:支付宝用户id, ALIPAY_TEL:手机号入库 CUSTOMER:自定义 DIRECT_CLIENT_PRE：本地预写入 DIRECT_CLIENT：本地客户端库 ENCLOSED：封闭场景 |
| `faceval` | `Faceval` | `*string` | JSON body | 未标注 | 入库用户信息 人脸预热，alipayUid必填 |
| `group` | `Group` | `*string` | JSON body | 未标注 | 分组5 |
| `storecode` | `Storecode` | `*string` | JSON body | 未标注 | 门店编码 |
| `validtimes` | `Validtimes` | `*string` | JSON body | 未标注 | 有效期天数，如7天、30天、365天 |

<a id="api-customerfacemanagedelete"></a>

### CustomerFaceManageDelete — 热脸出库

- 官方接口：[`zoloz.authentication.customer.facemanage.delete`](https://opendocs.alipay.com/pre-apis/05vh16)
- 路由：`POST /v3/zoloz/authentication/customer/facemanage/delete`
- 请求类型：[CustomerFaceManageDeleteRequest](authentication_request.go#L31)
- 响应类型：[CustomerFaceManageDeleteResponse](authentication_response.go#L19)

出库逻辑

请求约束：CustomerFaceManageDeleteRequest is the official V3 ZolozAuthenticationCustomerFacemanageDeleteModel schema.

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `areacode` | `Areacode` | `*string` | JSON body | 未标注 | 地域编码 |
| `biz_type` | `BizType` | `*string` | JSON body | 未标注 | 人脸产品能力 |
| `bizscale` | `Bizscale` | `*string` | JSON body | 未标注 | 业务量 |
| `brandcode` | `Brandcode` | `*string` | JSON body | 未标注 | 商户品牌 |
| `devicenum` | `Devicenum` | `*string` | JSON body | 未标注 | 商户机具唯一编码，关键参数 |
| `extinfo` | `Extinfo` | `*string` | JSON body | 未标注 | 拓展参数 |
| `facetype` | `Facetype` | `*string` | JSON body | 未标注 | 入库类型 IDCARD:身份证 ALIPAY_USER:支付宝用户id, ALIPAY_TEL:手机号入库 CUSTOMER:自定义 |
| `faceval` | `Faceval` | `*string` | JSON body | 未标注 | 入库用户信息 |
| `group` | `Group` | `*string` | JSON body | 未标注 | 分组 |
| `storecode` | `Storecode` | `*string` | JSON body | 未标注 | 门店编码 |
| `validtimes` | `Validtimes` | `*string` | JSON body | 未标注 | 有效期天数，如7天、30天、365天 |

<a id="api-customerftokenquery"></a>

### CustomerFTokenQuery — 查询刷脸结果信息接口

- 官方接口：[`zoloz.authentication.customer.ftoken.query`](https://opendocs.alipay.com/open-v3/ebd5a135_zoloz.authentication.customer.ftoken.query)
- 路由：`POST /v3/zoloz/authentication/customer/ftoken/query`
- 请求类型：[CustomerFTokenQueryRequest](authentication_request.go#L57)
- 响应类型：[CustomerFTokenQueryResponse](authentication_response.go#L33)

查询刷脸结果信息接口

请求约束：CustomerFTokenQueryRequest is the official V3 ZolozAuthenticationCustomerFtokenQueryModel schema.

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `biz_type` | `BizType` | `*string` | JSON body | 未标注 | 1、1：1人脸验证能力<br> 2、1：n人脸搜索能力（支付宝uid入库）<br> 3、1：n人脸搜索能力（支付宝手机号入库）<br> 4、手机号和人脸识别综合能力 |
| `ext_info` | `ExtInfo` | [`*FaceExtInfo`](#model-faceextinfo) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `ftoken` | `Ftoken` | `*string` | JSON body | 未标注 | 人脸token |
| `zim_id` | `ZimId` | `*string` | JSON body | 未标注 | 刷脸初始化流程中产生的zimId值 |

<a id="api-customersmilepayinitialize"></a>

### CustomerSmilePayInitialize — 人脸刷脸付初始化

- 官方接口：[`zoloz.authentication.customer.smilepay.initialize`](https://opendocs.alipay.com/pre-apis/05vfpu)
- 路由：`POST /v3/zoloz/authentication/customer/smilepay/initialize`
- 请求类型：[CustomerSmilePayInitializeRequest](authentication_request.go#L69)
- 响应类型：[CustomerSmilePayInitializeResponse](authentication_response.go#L53)

人脸刷脸付初始化

请求约束：CustomerSmilePayInitializeRequest is the official V3 ZolozAuthenticationCustomerSmilepayInitializeModel schema.

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `service_id` | `ServiceId` | `*string` | JSON body | 未标注 | 刷脸服务Id |
| `service_params` | `ServiceParams` | `*string` | JSON body | 未标注 | 商户可传入刷脸服务所需的扩展参数，json格式 |
| `zimmetainfo` | `Zimmetainfo` | `*string` | JSON body | 未标注 | { "apdidToken": "设备指纹", "appName": "应用名称", "appVersion": "应用版本", "bioMetaInfo": "生物信息如2.3.0:3,-4" } |

<a id="api-smilepayinitialize"></a>

### SmilePayInitialize — 刷脸支付初始化

- 官方接口：[`zoloz.authentication.smilepay.initialize`](https://opendocs.alipay.com/open-v3/1177dbeb_zoloz.authentication.smilepay.initialize)
- 路由：`POST /v3/zoloz/authentication/smilepay/initialize`
- 请求类型：[SmilePayInitializeRequest](authentication_request.go#L79)
- 响应类型：[SmilePayInitializeResponse](authentication_response.go#L59)

通过该接口获取刷脸支付服务的初始化信息，详细字段描述请参考产品说明文档

请求约束：SmilePayInitializeRequest is the official V3 ZolozAuthenticationSmilepayInitializeModel schema.

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `apdid_token` | `ApdidToken` | `*string` | JSON body | 未标注 | 设备指纹，用于唯一标识一台设备 |
| `app_name` | `AppName` | `*string` | JSON body | 未标注 | 人脸识别应用名称 |
| `app_version` | `AppVersion` | `*string` | JSON body | 未标注 | 人脸识别应用版本号 |
| `base_ver` | `BaseVer` | `*string` | JSON body | 未标注 | 基础包版本号 |
| `bio_meta_info` | `BioMetaInfo` | `*string` | JSON body | 未标注 | 生物识别元信息 |
| `device_model` | `DeviceModel` | `*string` | JSON body | 未标注 | 设备型号 |
| `device_type` | `DeviceType` | `*string` | JSON body | 未标注 | 设备类型 |
| `ext_info` | `ExtInfo` | [`*FaceExtParams`](#model-faceextparams) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `machine_info` | `MachineInfo` | [`*FaceMachineInfo`](#model-facemachineinfo) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `merchant_info` | `MerchantInfo` | [`*FaceMerchantInfo`](#model-facemerchantinfo) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `os_version` | `OsVersion` | `*string` | JSON body | 未标注 | 操作系统版本号 |
| `remote_log_id` | `RemoteLogId` | `*string` | JSON body | 未标注 | 业务ID |
| `zim_ver` | `ZimVer` | `*string` | JSON body | 未标注 | ZIM版本号 |

## 嵌套请求参数

以下仅展开本目录请求中实际使用的对象；字段必填规则沿用官方快照，“未标注”须结合业务说明确认。响应模型请沿各接口的响应类型链接查看。

<a id="model-faceextinfo"></a>

### FaceExtInfo

类型定义：[FaceExtInfo](authentication_model.go#L5)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `max_age` | `MaxAge` | `*string` | 未标注 | 年龄区间判断的上限，闭区间 |
| `min_age` | `MinAge` | `*string` | 未标注 | 年龄区间判断的下限，闭区间 |
| `query_type` | `QueryType` | `*string` | 未标注 | query_type不填, 返回uid<br> query_type=1, 返回手机号<br> query_type=2, 返回图片 |

<a id="model-faceextparams"></a>

### FaceExtParams

类型定义：[FaceExtParams](authentication_model.go#L25)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `biz_type` | `BizType` | `*string` | 未标注 | 业务类型：7，基于1:N人脸搜索的刷脸支付场景；8，基于姓名和身份证号的刷脸支付场景。 |

<a id="model-facemachineinfo"></a>

### FaceMachineInfo

类型定义：[FaceMachineInfo](authentication_model.go#L31)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `camera_drive_ver` | `CameraDriveVer` | `*string` | 未标注 | 摄像头驱动版本号 |
| `camera_model` | `CameraModel` | `*string` | 未标注 | 摄像头型号 |
| `camera_name` | `CameraName` | `*string` | 未标注 | 摄像头名称 |
| `camera_ver` | `CameraVer` | `*string` | 未标注 | 摄像头版本号 |
| `ext` | `Ext` | `*string` | 未标注 | 扩展信息 |
| `machine_code` | `MachineCode` | `*string` | 未标注 | 机具编码 |
| `machine_model` | `MachineModel` | `*string` | 未标注 | 机具型号 |
| `machine_ver` | `MachineVer` | `*string` | 未标注 | 机具版本号 |

<a id="model-facemerchantinfo"></a>

### FaceMerchantInfo

类型定义：[FaceMerchantInfo](authentication_model.go#L51)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `area_code` | `AreaCode` | `*string` | 未标注 | 区域编码 |
| `brand_code` | `BrandCode` | `*string` | 未标注 | 品牌编码 |
| `device_mac` | `DeviceMac` | `*string` | 未标注 | 机具Mac地址 |
| `device_num` | `DeviceNum` | `*string` | 未标注 | 机具编码 |
| `geo` | `Geo` | `*string` | 未标注 | 经纬度 |
| `group` | `Group` | `*string` | 未标注 | 机具分组编码 |
| `merchant_id` | `MerchantId` | `*string` | 未标注 | 商户ID |
| `partner_id` | `PartnerId` | `*string` | 未标注 | ISV ID |
| `store_code` | `StoreCode` | `*string` | 未标注 | 门店编码 |
| `wifimac` | `Wifimac` | `*string` | 未标注 | WI-FI Mac地址 |
| `wifiname` | `Wifiname` | `*string` | 未标注 | WI-FI 名称 |
