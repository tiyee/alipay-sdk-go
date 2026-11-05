# ad 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/data/dataservice/ad`。本包提供 3 个支付宝 V3 REST 操作。

通过 `ad.NewService(executor)` 复用 `*core.Client` 或最小 `core.Executor`。网络调用以 `context.Context` 为首参数，支持 `opts ...core.CallOption`，业务数据位于 `resp.Data`。

接口方法按单词边界使用大驼峰命名；请求、响应与专属模型位于对应业务包，共享模型从共同父包引用。

接入方式见[项目说明](../../../../README.md)。默认测试离线执行。

返回[领域目录](../../README.md)。

## 接口目录

| SDK 方法 | 官方接口 |
| --- | --- |
| [`DataServiceAdConversionUpload`](#api-dataserviceadconversionupload) | `alipay.data.dataservice.ad.conversion.upload` |
| [`DataServiceAdPromotePageBatchQuery`](#api-dataserviceadpromotepagebatchquery) | `alipay.data.dataservice.ad.promotepage.batchquery` |
| [`DataServiceAdPromotePageDownload`](#api-dataserviceadpromotepagedownload) | `alipay.data.dataservice.ad.promotepage.download` |

## 使用示例

- [DataServiceAdConversionUpload](api_ad_dataserviceadconversionupload_example_test.go)
- [DataServiceAdPromotePageBatchQuery](api_ad_dataserviceadpromotepagebatchquery_example_test.go)
- [DataServiceAdPromotePageDownload](api_ad_dataserviceadpromotepagedownload_example_test.go)

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-dataserviceadconversionupload"></a>

### DataServiceAdConversionUpload — 转化数据回传

- 官方接口：[`alipay.data.dataservice.ad.conversion.upload`](https://opendocs.alipay.com/open-v3/32a2992f_alipay.data.dataservice.ad.conversion.upload)
- 路由：`POST /v3/alipay/data/dataservice/ad/conversion/upload`
- 请求类型：[DataDataServiceAdConversionUploadRequest](ad_request.go#L4)
- 响应类型：[DataDataServiceAdConversionUploadResponse](ad_response.go#L4)

供广告商家通过OpenAPI上报转化明细数据，用于广告投放效果的追踪和广告投放的优化。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `biz_token` | `BizToken` | `*string` | JSON body | 未标注 | 代理商访问灯火平台的token |
| `conversion_data_list` | `ConversionDataList` | [`*[]ConversionData`](#model-conversiondata) | JSON body | 未标注 | 转化数据列表, 最多1000条 |

<a id="api-dataserviceadpromotepagebatchquery"></a>

### DataServiceAdPromotePageBatchQuery — 自建推广页列表批量查询

- 官方接口：[`alipay.data.dataservice.ad.promotepage.batchquery`](https://opendocs.alipay.com/open-v3/17d7cd6f_alipay.data.dataservice.ad.promotepage.batchquery)
- 路由：`GET /v3/alipay/data/dataservice/ad/promotepage/batchquery`
- 请求类型：[DataDataServiceAdPromotePageBatchQueryRequest](ad_request.go#L12)
- 响应类型：[DataDataServiceAdPromotePageBatchQueryResponse](ad_response.go#L8)

用于获取指定商家的自建推广页列表

请求约束：DataDataServiceAdPromotePageBatchQueryRequest 是 alipay.data.dataservice.ad.promotepage.batchquery 的 V3 query 请求参数。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `biz_token` | `BizToken` | `*string` | query | 否 | 代理商访问灯火平台的token |
| `principal_tag` | `PrincipalTag` | `*string` | query | 否 | 商家标志 |
| `type` | `Type` | `*string` | query | 否 | 推广页类型：COLLECT_INFO - 免费留资；TRADE - 付费留资； OPERATION_PAID - 运营商付费留资；待扩展 |
| `page_no` | `PageNo` | `*string` | query | 否 | 分页参数之页数，从1开始 |
| `page_size` | `PageSize` | `*string` | query | 否 | 分页参数之每页大小，最大1000 |

<a id="api-dataserviceadpromotepagedownload"></a>

### DataServiceAdPromotePageDownload — 自建推广页留资数据查询

- 官方接口：[`alipay.data.dataservice.ad.promotepage.download`](https://opendocs.alipay.com/open-v3/2887027f_alipay.data.dataservice.ad.promotepage.download)
- 路由：`GET /v3/alipay/data/dataservice/ad/promotepage/download`
- 请求类型：[DataDataServiceAdPromotePageDownloadRequest](ad_request.go#L26)
- 响应类型：[DataDataServiceAdPromotePageDownloadResponse](ad_response.go#L20)

用于获取指定商家指定推广页的留资数据，注意调用频次小于20qps

请求约束：DataDataServiceAdPromotePageDownloadRequest 是 alipay.data.dataservice.ad.promotepage.download 的 V3 query 请求参数。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `start_date` | `StartDate` | `*string` | query | 否 | 留资开始日期，格式：yyyy-mm-dd 或 yyyy-mm-dd hh:mm:ss;不能早于30天前 |
| `end_date` | `EndDate` | `*string` | query | 否 | 留资结束日期，格式：yyyy-mm-dd 或 yyyy-mm-dd hh:mm:ss;不能晚于当天 |
| `page_no` | `PageNo` | `*string` | query | 否 | 分页参数之页数，从1开始 |
| `page_size` | `PageSize` | `*string` | query | 否 | 分页参数之每页大小，最大1000 |
| `biz_token` | `BizToken` | `*string` | query | 否 | 代理商访问灯火平台的token |
| `principal_tag` | `PrincipalTag` | `*string` | query | 否 | 商家标志 |
| `promote_page_id` | `PromotePageId` | `*string` | query | 否 | 推广页id |

## 嵌套请求参数

<a id="model-conversiondata"></a>

### ConversionData

类型定义：[ConversionData](ad_model.go#L4)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `ad_id` | `AdId` | `*string` | 未标注 | 广告创意id，来自：[监测上报](https://adpub.alipay.com/lark/adrlark/qm1v2gtpvcftele4)或[落地页宏替换](https://adpub.alipay.com/lark/adrlark/gw3740rggm4n5gar)配置的__AD_ID__宏参发生用户点击替换后的值 |
| `attribute_list` | `AttributeList` | [`*[]ConversionProperty`](#model-conversionproperty) | 未标注 | 转化事件属性信息，用于转化事件类型相关属性规则上传。 可支持上传属性的转化事件类型及属性规则请参考该文档：[转化事件类型属性规则](https://adpub.alipay.com/adrlark/ivdktpyh511x9r6i) |
| `attributes` | `Attributes` | [`*ConversionProperty`](#model-conversionproperty) | 未标注 | 见字段类型及官方接口文档。 |
| `biz_no` | `BizNo` | `*string` | 未标注 | 转化流水号：由用户自定义，用于幂等 |
| `callback_ext_info` | `CallbackExtInfo` | `*string` | 未标注 | 来自：[监测上报](https://adpub.alipay.com/lark/adrlark/qm1v2gtpvcftele4)或[落地页宏替换](https://adpub.alipay.com/lark/adrlark/gw3740rggm4n5gar)配置的__CALLBACK_EXT_INFO__宏参发生用户点击替换后的值，需url decode后传回。 |
| `cid` | `Cid` | `*string` | 未标注 | 用于精准归因，来自：[监测上报](https://adpub.alipay.com/lark/adrlark/qm1v2gtpvcftele4)或[落地页宏替换](https://adpub.alipay.com/lark/adrlark/gw3740rggm4n5gar)配置的__CID__宏参发生用户点击替换后的值。<br> 支持cid上报的转化事件类型请参见：<br> [app推广类转化事件类型](https://adpub.alipay.com/lark/adrlark/pg4vb1c0g0u6fx7b)<br> [通用类转化事件类型](https://adpub.alipay.com/lark/adrlark/rvzylhryxh37yplt) |
| `conversion_amount` | `ConversionAmount` | `*string` | 未标注 | 转化金额，单位分。 金额类转化事件必传 |
| `conversion_id` | `ConversionId` | `*string` | 未标注 | 转化事件id，当source=COMMON_CONVERSION_ID需上传该字段，其他情况无需上传。 |
| `conversion_time` | `ConversionTime` | `*int64` | 未标注 | 转化时间，UTC 时间戳，单位：秒 |
| `conversion_type` | `ConversionType` | `*string` | 未标注 | 转化事件类型数字 |
| `creative_id` | `CreativeId` | `*string` | 未标注 | 广告创意id，来自：[监测上报](https://adpub.alipay.com/lark/adrlark/qm1v2gtpvcftele4)或[落地页宏替换](https://adpub.alipay.com/lark/adrlark/gw3740rggm4n5gar)配置的__AD_ID__宏参发生用户点击替换后的值 |
| `data_id` | `DataId` | `*string` | 未标注 | 当source=XLIGHT或者DATASET时使用： 数据集id(接口升级，该参数已不再使用，故废弃该参数，但不影响历史数据使用。) |
| `data_src_type` | `DataSrcType` | `*string` | 未标注 | 当source=OTHER时使用: 数据类型： KR_MEMBER - 客如云入会 KR_TRADE - 客如云交易 TB_LIVE - 淘宝直播(接口升级，该参数已不再使用，故废弃该参数，但不影响历史数据使用。) |
| `group_id` | `GroupId` | `*string` | 未标注 | 广告单元id，来自：[监测上报](https://adpub.alipay.com/lark/adrlark/qm1v2gtpvcftele4)或[落地页宏替换](https://adpub.alipay.com/lark/adrlark/gw3740rggm4n5gar)配置的__GROUP_ID__宏参发生用户点击替换后的值 |
| `join_channel` | `JoinChannel` | `*string` | 未标注 | XLIGHT - 灯火归因 MERCHANT -商家自主归因 该字段若为空则默认为XLIGHT |
| `plan_id` | `PlanId` | `*string` | 未标注 | 广告计划id，来自：[监测上报](https://adpub.alipay.com/lark/adrlark/qm1v2gtpvcftele4)或[落地页宏替换](https://adpub.alipay.com/lark/adrlark/gw3740rggm4n5gar)配置的__PLAN_ID__宏参发生用户点击替换后的值 |
| `principal_id` | `PrincipalId` | `*string` | 未标注 | 商户在灯火pb端的id, 可代替principal_tag。 |
| `principal_tag` | `PrincipalTag` | `*string` | 未标注 | 商家标志，可代替principal_id。获取方法： [principal_tag获取sop](https://adpub.alipay.com/lark/adrlark/sbk3vetg6ki5m4r8)<br>SDK 字段说明：商家标志，可代替principal_id。获取方法： [principal_tag获取sop](https://adpub.alipay.com/lark/adrlark/sbk3vetg6ki5m4r8)<br>。 |
| `property_list` | `PropertyList` | [`*[]ConversionProperty`](#model-conversionproperty) | 未标注 | 转化归因字段列表。不同转化事件类型回传property_list不同。请参见回传示例：<br> [通用类转化事件回传示例](https://adpub.alipay.com/lark/adrlark/pvbgwf72ea090877)<br> [APP推广类转化事件回传示例](https://adpub.alipay.com/lark/adrlark/kgnnk94p0p1mwbgp)<br>SDK 字段说明：转化归因字段列表。不同转化事件类型回传property_list不同。请参见回传示例：<br> [通用类转化事件回传示例](https://adpub.alipay.com/lark/adrlark/pvbgwf72ea090877)<br> [APP推广类转化事件回传示例](https://adpub.alipay.com/lark/adrlark/kgnnk94p0p1mwbgp)<br>。 |
| `source` | `Source` | `*string` | 未标注 | COMMON_TARGET：[通用类转化事件类型](https://adpub.alipay.com/lark/adrlark/rvzylhryxh37yplt)适用<br> CALLBACK：[APP推广类转化事件类型](https://adpub.alipay.com/lark/adrlark/pg4vb1c0g0u6fx7b)适用<br> 不同source的参数回传规则请参考： COMMON_TARGET：[通用转化回传参数说明](https://adpub.alipay.com/lark/adrlark/dxw7fkkdnhm45spm)<br> CALLBACK： [APP推广类转化回传参数说明](https://adpub.alipay.com/lark/adrlark/wxou6agg3z1o6wyn) |
| `target_id` | `TargetId` | `*string` | 未标注 | 当source=OTHER时使用： 主体id，例如品牌id(接口升级，该参数已不再使用，故废弃该参数，但不影响历史数据使用。) |
| `target_type` | `TargetType` | `*string` | 未标注 | 当source=OTHER时使用： 主体类型： BRAND - 品牌 STORE - 店铺 LIVE - 直播 等等(接口升级，该参数已不再使用，故废弃该参数，但不影响历史数据使用。) |
| `uuid` | `Uuid` | `*string` | 未标注 | 发生转化用户的唯一标识。 uuid_type不同，uuid传值不同：<br> 1. 若uuid_type=【PID】：uuid为发生转化用户的2088id。若获取不到，可固定一个值后，上传callback_ext_info参数。<br> 2. 若uuid_type=【OAID_MD5】【IDFA_MD5】【IP_UA】【CAID】【MEI】: uuid传的值为[监测上报](https://adpub.alipay.com/lark/adrlark/qm1v2gtpvcftele4)或[落地页宏替换](https://adpub.alipay.com/lark/adrlark/gw3740rggm4n5gar)配置的对应宏参下发的值<br> 其中当uuid_type=IP_UA时：__IP__、__UA__宏参需用"_"拼接上传<br>SDK 字段说明：发生转化用户的唯一标识。 uuid_type不同，uuid传值不同：<br> 1. 若uuid_type=【PID】：uuid为发生转化用户的2088id。若获取不到，可固定一个值后，上传callback_ext_info参数。<br> 2. 若uuid_type=【OAID_MD5】【IDFA_MD5】【IP_UA】【CAID】【MEI】: uuid传的值为[监测上报](https://adpub.alipay.com/lark/adrlark/qm1v2gtpvcftele4)或[落地页宏替换](https://adpub.alipay.com/lark/adrlark/gw3740rggm4n5gar)配置的对应宏参下发的值<br> 其中当uuid_type=IP_UA时：__IP__、__UA__宏参需用"_"拼接上传<br>。 |
| `uuid_open_id` | `UuidOpenId` | `*string` | 未标注 | 支付宝用户在应用维度下的唯一标识 |
| `uuid_type` | `UuidType` | `*string` | 未标注 | [通用类转化事件类型：](https://adpub.alipay.com/lark/adrlark/rvzylhryxh37yplt)<br> uuid_type固定为PID：表示转化用户唯一标识<br> [APP推广类转化事件类型：](https://adpub.alipay.com/lark/adrlark/pg4vb1c0g0u6fx7b)<br> OAID_MD5：表示OAID原值MD5加密后的值<br> IDFA_MD5：表示IDFA原值MD5加密后的值<br> IP_UA：表示IP_UA拼接值，以下划线拼接<br> CAID：表示CAID原值<br> IMEI：表示IMEI原值<br>SDK 字段说明：[通用类转化事件类型：](https://adpub.alipay.com/lark/adrlark/rvzylhryxh37yplt)<br> uuid_type固定为PID：表示转化用户唯一标识<br> [APP推广类转化事件类型：](https://adpub.alipay.com/lark/adrlark/pg4vb1c0g0u6fx7b)<br> OAID_MD5：表示OAID原值MD5加密后的值<br> IDFA_MD5：表示IDFA原值MD5加密后的值<br> IP_UA：表示IP_UA拼接值，以下划线拼接<br> CAID：表示CAID原值<br> IMEI：表示IMEI原值<br>。 |

<a id="model-conversionproperty"></a>

### ConversionProperty

类型定义：[ConversionProperty](ad_model.go#L57)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `key` | `Key` | `*string` | 未标注 | 业务转化数据属性 |
| `name` | `Name` | `*string` | 未标注 | 转化属性名称 |
| `value` | `Value` | `*string` | 未标注 | 业务转化属性实例值， property_list的value值需与转化管理里创建转化事件时填写的归因口径值保持一致<br> attrbute_list的value值请参考该文档：[转化事件类型属性规则](https://adpub.alipay.com/adrlark/ivdktpyh511x9r6i) |
