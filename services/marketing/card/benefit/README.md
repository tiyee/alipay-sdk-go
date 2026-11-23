# benefit 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/marketing/card/benefit`。本包提供 4 个支付宝 V3 REST 操作。

通过 `benefit.NewService(executor)` 复用 `*core.Client` 或最小 `core.Executor`。网络调用以 `context.Context` 为首参数，支持 `opts ...core.CallOption`，业务数据位于 `resp.Data`。

接口方法按单词边界使用大驼峰命名；请求、响应与专属模型位于对应业务包，共享模型从共同父包引用。

接入方式见[项目说明](../../../../README.md)。默认测试离线执行。

返回[领域目录](../README.md)。

## 接口目录

| SDK 方法 | 官方接口 |
| --- | --- |
| [`CardBenefitCreate`](#api-cardbenefitcreate) | `alipay.marketing.card.benefit.create` |
| [`CardBenefitDelete`](#api-cardbenefitdelete) | `alipay.marketing.card.benefit.delete` |
| [`CardBenefitModify`](#api-cardbenefitmodify) | `alipay.marketing.card.benefit.modify` |
| [`CardBenefitQuery`](#api-cardbenefitquery) | `alipay.marketing.card.benefit.query` |

## 使用示例

- [CardBenefitCreate](api_benefit_cardbenefitcreate_example_test.go)
- [CardBenefitDelete](api_benefit_cardbenefitdelete_example_test.go)
- [CardBenefitModify](api_benefit_cardbenefitmodify_example_test.go)
- [CardBenefitQuery](api_benefit_cardbenefitquery_example_test.go)

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-cardbenefitcreate"></a>

### CardBenefitCreate — 会员卡模板外部权益创建

- 官方接口：[`alipay.marketing.card.benefit.create`](https://opendocs.alipay.com/open-v3/2aa50202_alipay.marketing.card.benefit.create?scene=common&pathHash=c87e30d3)
- 路由：`POST /v3/alipay/marketing/card/benefit/create`
- 请求类型：[CardBenefitCreateRequest](benefit_request.go#L4)
- 响应类型：[CardBenefitCreateResponse](benefit_response.go#L4)

会员卡模板外部权益创建，创建的权益内容信息将展示在卡包中卡详情页中部区域。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `mcard_template_benefit` | `McardTemplateBenefit` | [`*McardTemplateBenefit`](#model-mcardtemplatebenefit) | JSON body | 未标注 | 见字段类型及官方接口文档。 |

<a id="api-cardbenefitdelete"></a>

### CardBenefitDelete — 会员卡模板外部权益删除

- 官方接口：[`alipay.marketing.card.benefit.delete`](https://opendocs.alipay.com/open-v3/85008a68_alipay.marketing.card.benefit.delete?pathHash=112339fb)
- 路由：`DELETE /v3/alipay/marketing/card/benefit/delete`
- 请求类型：[CardBenefitDeleteRequest](benefit_request.go#L16)
- 响应类型：[CardBenefitDeleteResponse](benefit_response.go#L10)

会员卡模板外部权益删除

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `template_id` | `TemplateId` | `*string` | query | 否 | 会员卡模板ID，通过 [alipay.marketing.card.template.create](https://opendocs.alipay.com/apis/api_5/alipay.marketing.card.template.create)（会员卡模板创建)接口创建会员卡模板获取。 |
| `benefit_id` | `BenefitId` | `*string` | query | 否 | 权益ID，通过 [alipay.marketing.card.benefit.create](https://opendocs.alipay.com/apis/api_5/alipay.marketing.card.benefit.create)(会员卡模板外部权益创建)接口创建获取。 |

<a id="api-cardbenefitmodify"></a>

### CardBenefitModify — 会员卡模板外部权益修改

- 官方接口：[`alipay.marketing.card.benefit.modify`](https://opendocs.alipay.com/open-v3/f8a32ad5_alipay.marketing.card.benefit.modify?pathHash=df326137)
- 路由：`POST /v3/alipay/marketing/card/benefit/modify`
- 请求类型：[CardBenefitModifyRequest](benefit_request.go#L9)
- 响应类型：[CardBenefitModifyResponse](benefit_response.go#L16)

会员卡模板外部权益修改

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `benefit_id` | `BenefitId` | `*string` | JSON body | 未标注 | 权益ID，通过 [alipay.marketing.card.benefit.create](https://opendocs.alipay.com/apis/api_5/alipay.marketing.card.benefit.create)(会员卡模板外部权益创建)接口创建获取。 |
| `mcard_template_benefit` | `McardTemplateBenefit` | [`*McardTemplateBenefit`](#model-mcardtemplatebenefit) | JSON body | 未标注 | 见字段类型及官方接口文档。 |

<a id="api-cardbenefitquery"></a>

### CardBenefitQuery — 会员卡模板外部权益查询

- 官方接口：[`alipay.marketing.card.benefit.query`](https://opendocs.alipay.com/open-v3/217a04aa_alipay.marketing.card.benefit.query?pathHash=c22b4480)
- 路由：`GET /v3/alipay/marketing/card/benefit/query`
- 请求类型：[CardBenefitQueryRequest](benefit_request.go#L24)
- 响应类型：[CardBenefitQueryResponse](benefit_response.go#L22)

会员卡模板外部权益查询

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `template_id` | `TemplateId` | `*string` | query | 否 | 会员卡模板ID，通过 [alipay.marketing.card.template.create](https://opendocs.alipay.com/apis/api_5/alipay.marketing.card.template.create)（会员卡模板创建)接口创建会员卡模板获取。 |
| `benefit_id` | `BenefitId` | `*string` | query | 否 | 权益ID，通过 [alipay.marketing.card.benefit.create](https://opendocs.alipay.com/apis/api_5/alipay.marketing.card.benefit.create)(会员卡模板外部权益创建)接口创建获取。 |

## 嵌套请求参数

<a id="model-mcardtemplatebenefit"></a>

### McardTemplateBenefit

类型定义：[McardTemplateBenefit](benefit_model.go#L4)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `benefit_desc` | `BenefitDesc` | `*[]string` | 未标注 | 权益描述信息 |
| `end_date` | `EndDate` | `*string` | 未标注 | 权益结束时间。 注：在权益开始时间和结束时间范围内的权益才会认为是有效权益进行展示。 |
| `ext_info` | `ExtInfo` | `*string` | 未标注 | 会员卡模板权益扩展信息：JSON格式; openUrl 说明：跳转到商户的优惠活动页面 |
| `start_date` | `StartDate` | `*string` | 未标注 | 权益开始时间 |
| `template_id` | `TemplateId` | `*string` | 未标注 | 会员卡模板ID |
| `title` | `Title` | `*string` | 未标注 | 权益标题 |
