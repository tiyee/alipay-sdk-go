# blueseaactivity 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/open/sp/blueseaactivity`。本包提供 3 个支付宝 V3 REST 操作。

通过 `blueseaactivity.NewService(executor)` 复用 `*core.Client` 或最小 `core.Executor`。网络调用以 `context.Context` 为首参数，支持 `opts ...core.CallOption`，业务数据位于 `resp.Data`。

接口方法按单词边界使用大驼峰命名；请求、响应与专属模型位于对应业务包，共享模型从共同父包引用。

接入方式见[项目说明](../../../../README.md)。默认测试离线执行。

返回[领域目录](../README.md)。

## 接口目录

| SDK 方法 | 官方接口 |
| --- | --- |
| [`SpBlueSeaActivityCreate`](#api-spblueseaactivitycreate) | `alipay.open.sp.blueseaactivity.create` |
| [`SpBlueSeaActivityModify`](#api-spblueseaactivitymodify) | `alipay.open.sp.blueseaactivity.modify` |
| [`SpBlueSeaActivityQuery`](#api-spblueseaactivityquery) | `alipay.open.sp.blueseaactivity.query` |

## 使用示例

- [SpBlueSeaActivityCreate](api_blueseaactivity_spblueseaactivitycreate_example_test.go)
- [SpBlueSeaActivityModify](api_blueseaactivity_spblueseaactivitymodify_example_test.go)
- [SpBlueSeaActivityQuery](api_blueseaactivity_spblueseaactivityquery_example_test.go)

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-spblueseaactivitycreate"></a>

### SpBlueSeaActivityCreate — 服务商返佣活动报名申请

- 官方接口：[`alipay.open.sp.blueseaactivity.create`](https://opendocs.alipay.com/open-v3/a974b7ae_alipay.open.sp.blueseaactivity.create)
- 路由：`POST /v3/alipay/open/sp/blueseaactivity/create`
- 请求类型：[SpBlueSeaActivityCreateRequest](blueseaactivity_request.go#L4)
- 响应类型：[SpBlueSeaActivityCreateResponse](blueseaactivity_response.go#L4)

服务商返佣活动报名

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `address` | `Address` | `*string` | JSON body | 未标注 | 详细地址。 该资质是否必传请参见[报名资质要求](https://opendocs.alipay.com/open/01hd83)。 |
| `biz_scene` | `BizScene` | `*string` | JSON body | 未标注 | 蓝海活动的场景，支持场景如下： * 直连餐饮（BLUE_SEA_FOOD_APPLY）； * 直连快消（BLUE_SEA_FMCG_APPLY）； * 间连餐饮（BLUE_SEA_FOOD_INDIRECT_APPLY）； * 间连快消（BLUE_SEA_FMCG_INDIRECT_APPLY）。 |
| `business_lic` | `BusinessLic` | `*string` | JSON body | 未标注 | 营业执照，要求证件文本信息清晰可见，图片文件大小在 50K-5M 之间，不限制长宽，支持 png、bmp、gif、jpg、jpeg 格式。 请传入[alipay.open.sp.image.upload](https://opendocs.alipay.com/apis/01ea4t)(图片上传接口) 返回的 image_id。 该资质是否必传请参见[报名资质要求](https://opendocs.alipay.com/open/01hd83)。 |
| `city_code` | `CityCode` | `*string` | JSON body | 未标注 | 城市编码。请按照[表格](https://gw.alipayobjects.com/os/basement_prod/253c4dcb-b8a4-4a1e-8be2-79e191a9b6db.xlsx) 中内容填写。 （请参考资料：[ 统计用区划和城乡划分代码](http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/)） |
| `district_code` | `DistrictCode` | `*string` | JSON body | 未标注 | 区县编码。请按照[表格](https://gw.alipayobjects.com/os/basement_prod/253c4dcb-b8a4-4a1e-8be2-79e191a9b6db.xlsx) 中内容填写。 （请参考资料：[ 统计用区划和城乡划分代码](http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/)） |
| `food_business_lic` | `FoodBusinessLic` | `*string` | JSON body | 未标注 | 食品经营许可证，要求证件文本信息清晰可见，图片文件大小在 50K-5M 之间，不限制长宽，支持 png、bmp、gif、jpg、jpeg 格式。 请传入[alipay.open.sp.image.upload](https://opendocs.alipay.com/apis/01ea4t)(图片上传接口) 返回的 image_id。 该资质是否必传请参见[报名资质要求](https://opendocs.alipay.com/open/01hd83)。 |
| `food_circulate_lic` | `FoodCirculateLic` | `*string` | JSON body | 未标注 | 食品流通许可证，要求证件文本信息清晰可见，图片文件大小在 50K-5M 之间，不限制长宽，支持 png、bmp、gif、jpg、jpeg 格式。 请传入[alipay.open.sp.image.upload](https://opendocs.alipay.com/apis/01ea4t)(图片上传接口) 返回的 image_id。 该资质是否必传请参见[报名资质要求](https://opendocs.alipay.com/open/01hd83)。 |
| `food_health_lic` | `FoodHealthLic` | `*string` | JSON body | 未标注 | 食品卫生许可证，要求证件文本信息清晰可见，图片文件大小在 50K-5M 之间，不限制长宽，支持 png、bmp、gif、jpg、jpeg 格式。 请传入[alipay.open.sp.image.upload](https://opendocs.alipay.com/apis/01ea4t)(图片上传接口) 返回的 image_id。 该资质是否必传请参见[报名资质要求](https://opendocs.alipay.com/open/01hd83)。 |
| `food_production_lic` | `FoodProductionLic` | `*string` | JSON body | 未标注 | 食品生产许可证，要求证件文本信息清晰可见，图片文件大小在 50K-5M 之间，不限制长宽，支持 png、bmp、gif、jpg、jpeg 格式。 请传入[alipay.open.sp.image.upload](https://opendocs.alipay.com/apis/01ea4t)(图片上传接口) 返回的 image_id。 该资质是否必传请参见[报名资质要求](https://opendocs.alipay.com/open/01hd83)。 |
| `food_service_lic` | `FoodServiceLic` | `*string` | JSON body | 未标注 | 餐饮服务许可证，要求证件文本信息清晰可见，图片文件大小在 50K-5M 之间，不限制长宽，支持 png、bmp、gif、jpg、jpeg 格式。 请传入[alipay.open.sp.image.upload](https://opendocs.alipay.com/apis/01ea4t)(图片上传接口) 返回的 image_id。 该资质是否必传请参见[报名资质要求](https://opendocs.alipay.com/open/01hd83)。 |
| `indoor_pic` | `IndoorPic` | `*string` | JSON body | 未标注 | 内景照，要求图片清晰可见，图片文件大小在 50K-5M 之间，不限制长宽，支持 png、bmp、gif、jpg、jpeg 格式。 请传入[alipay.open.sp.image.upload](https://opendocs.alipay.com/apis/01ea4t)(图片上传接口) 返回的 image_id。 该资质是否必传请参见[报名资质要求](https://opendocs.alipay.com/open/01hd83)。 |
| `merchant_logon` | `MerchantLogon` | `*string` | JSON body | 未标注 | 参与蓝海活动的商户支付宝账号，只有当参与直连蓝海活动场景（BLUE_SEA_FOOD_APPLY/BLUE_SEA_FMCG_APPLY）时必填，间连场景可空。 |
| `province_code` | `ProvinceCode` | `*string` | JSON body | 未标注 | 省份编码。请按照[表格](https://gw.alipayobjects.com/os/basement_prod/253c4dcb-b8a4-4a1e-8be2-79e191a9b6db.xlsx) 中内容填写。 （请参考资料：[ 统计用区划和城乡划分代码](http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/)） |
| `shop_entrance_pic` | `ShopEntrancePic` | `*string` | JSON body | 未标注 | 门头照，要求图片清晰可见，图片文件大小在 50K-5M 之间，不限制长宽，支持 png、bmp、gif、jpg、jpeg 格式。 请传入[alipay.open.sp.image.upload](https://opendocs.alipay.com/apis/01ea4t)(图片上传接口) 返回的 image_id。 该资质是否必传请参见[报名资质要求](https://opendocs.alipay.com/open/01hd83)。 |
| `sub_merchant_id` | `SubMerchantId` | `*string` | JSON body | 未标注 | 参与蓝海活动的间连商户账号，只有当参与间连蓝海活动场景（BLUE_SEA_FOOD_INDIRECT_APPLY/BLUE_SEA_FMCG_INDIRECT_APPLY）时必填，直连场景可空。 |
| `tobacco_lic` | `TobaccoLic` | `*string` | JSON body | 未标注 | 烟草专卖零售许可证，要求证件文本信息清晰可见，图片文件大小在 50K-5M 之间，不限制长宽，支持 png、bmp、gif、jpg、jpeg 格式。 请传入[alipay.open.sp.image.upload](https://opendocs.alipay.com/apis/01ea4t)(图片上传接口) 返回的 image_id。 该资质是否必传请参见[报名资质要求](https://opendocs.alipay.com/open/01hd83)。 |

<a id="api-spblueseaactivitymodify"></a>

### SpBlueSeaActivityModify — 服务商返佣活动申请单修改

- 官方接口：[`alipay.open.sp.blueseaactivity.modify`](https://opendocs.alipay.com/open-v3/8cbcb343_alipay.open.sp.blueseaactivity.modify)
- 路由：`POST /v3/alipay/open/sp/blueseaactivity/modify`
- 请求类型：[SpBlueSeaActivityModifyRequest](blueseaactivity_request.go#L40)
- 响应类型：[SpBlueSeaActivityModifyResponse](blueseaactivity_response.go#L10)

服务商返佣活动

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `address` | `Address` | `*string` | JSON body | 未标注 | 详细地址 |
| `business_lic` | `BusinessLic` | `*string` | JSON body | 未标注 | 营业执照，要求证件文本信息清晰可见。 请上传照片的image_id，传参明细请参见[报名资质要求](https://opendocs.alipay.com/open/01hd83) |
| `city_code` | `CityCode` | `*string` | JSON body | 未标注 | 城市编码。请按照https://gw.alipayobjects.com/os/basement_prod/253c4dcb-b8a4-4a1e-8be2-79e191a9b6db.xlsx 表格中内容填写。 （参考资料： http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/） |
| `district_code` | `DistrictCode` | `*string` | JSON body | 未标注 | 区县编码。请按照https://gw.alipayobjects.com/os/basement_prod/253c4dcb-b8a4-4a1e-8be2-79e191a9b6db.xlsx 表格中内容填写。 （参考资料： http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/） |
| `food_business_lic` | `FoodBusinessLic` | `*string` | JSON body | 未标注 | 食品经营许可证，要求证件文本信息清晰可见。请上传照片的image_id，传参明细请参见[报名资质要求](https://opendocs.alipay.com/open/01hd83) |
| `food_circulate_lic` | `FoodCirculateLic` | `*string` | JSON body | 未标注 | 食品流通许可证，要求证件文本信息清晰可见。请上传照片的image_id，传参明细请参见[报名资质要求](https://opendocs.alipay.com/open/01hd83) |
| `food_health_lic` | `FoodHealthLic` | `*string` | JSON body | 未标注 | 食品卫生许可证，要求证件文本信息清晰可见。请上传照片的image_id，传参明细请参见[报名资质要求](https://opendocs.alipay.com/open/01hd83) |
| `food_production_lic` | `FoodProductionLic` | `*string` | JSON body | 未标注 | 食品生产许可证，要求证件文本信息清晰可见。请上传照片的image_id，传参明细请参见[报名资质要求](https://opendocs.alipay.com/open/01hd83) |
| `food_service_lic` | `FoodServiceLic` | `*string` | JSON body | 未标注 | 餐饮服务许可证，要求证件文本信息清晰可见。请上传照片的image_id，传参明细请参见[报名资质要求](https://opendocs.alipay.com/open/01hd83) |
| `indoor_pic` | `IndoorPic` | `*string` | JSON body | 未标注 | 门头照，要求内景照片清晰可见。请上传照片的image_id，传参明细请参见[报名资质要求](https://opendocs.alipay.com/open/01hd83) |
| `order_id` | `OrderId` | `*string` | JSON body | 未标注 | 申请单 id。通过 [alipay.open.sp.blueseaactivity.create](https://opendocs.alipay.com/apis/01ebig)接口获取。 |
| `province_code` | `ProvinceCode` | `*string` | JSON body | 未标注 | 省份编码。请按照https://gw.alipayobjects.com/os/basement_prod/253c4dcb-b8a4-4a1e-8be2-79e191a9b6db.xlsx 表格中内容填写。 （参考资料： http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/） |
| `shop_entrance_pic` | `ShopEntrancePic` | `*string` | JSON body | 未标注 | 门头照，要求店铺外观照片清晰可见。请上传照片的image_id，传参明细请参见[报名资质要求](https://opendocs.alipay.com/open/01hd83) |
| `tobacco_lic` | `TobaccoLic` | `*string` | JSON body | 未标注 | 烟草专卖零售许可证，要求证件文本信息清晰可见。请上传照片的image_id，传参明细请参见[报名资质要求](https://opendocs.alipay.com/open/01hd83) |

<a id="api-spblueseaactivityquery"></a>

### SpBlueSeaActivityQuery — 服务商返佣活动申请单详情查询

- 官方接口：[`alipay.open.sp.blueseaactivity.query`](https://opendocs.alipay.com/open-v3/4ed0241a_alipay.open.sp.blueseaactivity.query)
- 路由：`GET /v3/alipay/open/sp/blueseaactivity/query`
- 请求类型：[SpBlueSeaActivityQueryRequest](blueseaactivity_request.go#L72)
- 响应类型：[SpBlueSeaActivityQueryResponse](blueseaactivity_response.go#L14)

服务商返佣活动

请求约束：SpBlueSeaActivityQueryRequest is the request type for alipay.open.sp.blueseaactivity.query.

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `order_id` | `OrderId` | `*string` | query | 否 | 申请单 id。通过 [alipay.open.sp.blueseaactivity.create](https://opendocs.alipay.com/apis/01ebig)接口获取 |
