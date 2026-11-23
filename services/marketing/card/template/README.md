# template 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/marketing/card/template`。本包提供 4 个支付宝 V3 REST 操作。

通过 `template.NewService(executor)` 复用 `*core.Client` 或最小 `core.Executor`。网络调用以 `context.Context` 为首参数，支持 `opts ...core.CallOption`，业务数据位于 `resp.Data`。

接口方法按单词边界使用大驼峰命名；请求、响应与专属模型位于对应业务包，共享模型从共同父包引用。

接入方式见[项目说明](../../../../README.md)。默认测试离线执行。

返回[领域目录](../README.md)。

## 接口目录

| SDK 方法 | 官方接口 |
| --- | --- |
| [`CardTemplateBatchQuery`](#api-cardtemplatebatchquery) | `alipay.marketing.card.template.batchquery` |
| [`CardTemplateCreate`](#api-cardtemplatecreate) | `alipay.marketing.card.template.create` |
| [`CardTemplateModify`](#api-cardtemplatemodify) | `alipay.marketing.card.template.modify` |
| [`CardTemplateQuery`](#api-cardtemplatequery) | `alipay.marketing.card.template.query` |

## 使用示例

- [CardTemplateBatchQuery](api_template_cardtemplatebatchquery_example_test.go)
- [CardTemplateCreate](api_template_cardtemplatecreate_example_test.go)
- [CardTemplateModify](api_template_cardtemplatemodify_example_test.go)
- [CardTemplateQuery](api_template_cardtemplatequery_example_test.go)

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-cardtemplatebatchquery"></a>

### CardTemplateBatchQuery — 会员卡模板批量查询

- 官方接口：[`alipay.marketing.card.template.batchquery`](https://opendocs.alipay.com/open-v3/54dcbcfa_alipay.marketing.card.template.batchquery?pathHash=030a7dec)
- 路由：`POST /v3/alipay/marketing/card/template/batchquery`
- 请求类型：[CardTemplateBatchQueryRequest](template_request.go#L4)
- 响应类型：[CardTemplateBatchQueryResponse](template_response.go#L4)

会员卡模板批量查询

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `page_no` | `PageNo` | `*int64` | JSON body | 未标注 | 查询起始页码，如第1页，单位：/页 |
| `page_size` | `PageSize` | `*int64` | JSON body | 未标注 | 每页数据大小，默认 20条，最大支持20个，单位：/个 |

<a id="api-cardtemplatecreate"></a>

### CardTemplateCreate — 会员卡模板创建

- 官方接口：[`alipay.marketing.card.template.create`](https://opendocs.alipay.com/open-v3/18fdd0bd_alipay.marketing.card.template.create?scene=1c74e13b68cc4a58912fb6f0c40b62af%2C1529c3090a7b4db99164667a0b22dcaa&pathHash=d206dedb)
- 路由：`POST /v3/alipay/marketing/card/template/create`
- 请求类型：[CardTemplateCreateRequest](template_request.go#L12)
- 响应类型：[CardTemplateCreateResponse](template_response.go#L12)

会员卡模板创建

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `access_version` | `AccessVersion` | `*string` | JSON body | 未标注 | 接入版本 |
| `biz_no_prefix` | `BizNoPrefix` | `*string` | JSON body | 未标注 | 业务卡号前缀，由商户指定 支付宝业务卡号生成规则：biz_no_prefix(商户指定)卡号前缀 + biz_no_suffix(实时生成）卡号后缀 |
| `biz_no_suffix_len` | `BizNoSuffixLen` | `*string` | JSON body | 未标注 | 业务卡号后缀长度，与biz_no_prefix配合，扣除系统预留2位，剩下对应seq长度。在生成卡号时，若seq位数不足前置补0，若seq位数超出则以实际为准。举例：设为10，其中8位用于生成seq，可覆盖1亿用户，即使seq超过1亿，生成卡号也不报错，但总长度不得超过32位。建议按需设置合适的值，以获得长度一致的业务卡号，建议长度20，性能更好。 |
| `card_action_list` | `CardActionList` | [`*[]TemplateActionInfoDTO`](#model-templateactioninfodto) | JSON body | 未标注 | "卡行动点配置； 行动点，即用户可点击跳转的区块，类似按钮控件的交互； 单张卡最多定制3个行动点。如果各项的配置内容完全一样，展示时会进行去重。" |
| `card_level_conf` | `CardLevelConf` | [`*[]TemplateCardLevelConfDTO`](#model-templatecardlevelconfdto) | JSON body | 未标注 | 卡级别配置 |
| `card_spec_tag` | `CardSpecTag` | `*string` | JSON body | 未标注 | 卡特定标签，只供特定业务使用，通常接入无需关注 |
| `card_type` | `CardType` | `*string` | JSON body | 未标注 | 卡类型 |
| `column_info_list` | `ColumnInfoList` | [`*[]TemplateColumnInfoDTO`](#model-templatecolumninfodto) | JSON body | 未标注 | 栏位信息 |
| `field_rule_list` | `FieldRuleList` | [`*[]TemplateFieldRuleDTO`](#model-templatefieldruledto) | JSON body | 未标注 | "字段规则列表，会员卡开卡过程中，会员卡信息的生成规则， 例如：卡有效期为开卡后两年内有效，则设置为：DATE_IN_FUTURE" |
| `mdcode_notify_conf` | `MdcodeNotifyConf` | [`*TemplateMdcodeNotifyConfDTO`](#model-templatemdcodenotifyconfdto) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `notify_url` | `NotifyUrl` | `*string` | JSON body | 未标注 | 通知地址 |
| `open_card_conf` | `OpenCardConf` | [`*TemplateOpenCardConfDTO`](#model-templateopencardconfdto) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `paid_outer_card_conf` | `PaidOuterCardConf` | [`*PaidOuterCardTemplateConfDTO`](#model-paidoutercardtemplateconfdto) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `pub_channels` | `PubChannels` | [`*[]PubChannelDTO`](#model-pubchanneldto) | JSON body | 未标注 | 卡模板投放渠道 |
| `request_id` | `RequestId` | `*string` | JSON body | 未标注 | 请求ID，商家自定义且并保证唯一性。 |
| `service_label_list` | `ServiceLabelList` | `*[]string` | JSON body | 未标注 | 服务标签列表 |
| `shop_ids` | `ShopIds` | `*[]string` | JSON body | 未标注 | 会员卡上架门店id（支付宝门店id），即发放会员卡的商家门店id |
| `spi_app_id` | `SpiAppId` | `*string` | JSON body | 未标注 | spi应用id，为实现spi.alipay.user.opencard.get接口的app_id。若是第三方代理模式，可以设置成服务商的 APPID 或者是商家自己的 APPID； 若不是第三方代理模式，只能设置商家自己的 APPID。 |
| `template_benefit_info` | `TemplateBenefitInfo` | [`*[]TemplateBenefitInfoDTO`](#model-templatebenefitinfodto) | JSON body | 未标注 | "权益信息， 1、在卡包的卡详情页面会自动添加“会员权益”栏位，展现此字段配置的会员卡权益信息， 2、如果添加门店渠道，则可在门店页展现会员卡的权益" |
| `template_form_config` | `TemplateFormConfig` | [`*TemplateFormConfig`](#model-templateformconfig) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `template_style_info` | `TemplateStyleInfo` | [`*TemplateStyleInfoDTO`](#model-templatestyleinfodto) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `write_off_type` | `WriteOffType` | `*string` | JSON body | 未标注 | 卡包详情页面中展现出的卡码（可用于扫码核销） (1) 静态码 qrcode: 二维码，扫码得商户开卡传入的external_card_no barcode: 条形码，扫码得商户开卡传入的external_card_no (2) 动态码-支付宝生成码值(动态码会在2分钟左右后过期) dqrcode: 动态二维码，扫码得到的码值可配合会员卡查询接口使用 dbarcode: 动态条形码，扫码得到的码值可配合会员卡查询接口使用 (3) 动态码-商家自主生成码值（码值、时效性都由商户控制） mdqrcode: 商户动态二维码，扫码得商户自主传入的码值 mdbarcode: 商户动态条码，扫码得商户自主传入的码值 如需使用商户动态码，对接详情参见 [商户动态发码](https://opendocs.alipay.com/open/02y7gt)。 |

<a id="api-cardtemplatemodify"></a>

### CardTemplateModify — 会员卡模板修改

- 官方接口：[`alipay.marketing.card.template.modify`](https://opendocs.alipay.com/open-v3/fabfc518_alipay.marketing.card.template.modify?scene=692a39a5ad9e45d4899177a7ec9cf68e%2Cc76bbe273e3b471bbba690c624b9ed86&pathHash=e43d6caf)
- 路由：`POST /v3/alipay/marketing/card/template/modify`
- 请求类型：[CardTemplateModifyRequest](template_request.go#L55)
- 响应类型：[CardTemplateModifyResponse](template_response.go#L18)

会员卡模板修改

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `access_version` | `AccessVersion` | `*string` | JSON body | 未标注 | 接口接入版本 |
| `biz_no_prefix` | `BizNoPrefix` | `*string` | JSON body | 未标注 | 业务卡号前缀，由商户指定 支付宝业务卡号生成规则：biz_no_prefix(商户指定)卡号前缀 + biz_no_suffix(实时生成）卡号后缀 |
| `biz_no_suffix_len` | `BizNoSuffixLen` | `*string` | JSON body | 未标注 | 业务卡号后缀长度，与biz_no_prefix配合，扣除系统预留2位，剩下对应seq长度。在生成卡号时，若seq位数不足前置补0，若seq位数超出则以实际为准。举例：设为10，其中8位用于生成seq，可覆盖1亿用户，即使seq超过1亿，生成卡号也不报错，但总长度不得超过32位。建议按需设置合适的值，以获得长度一致的业务卡号，建议长度20，性能更好。单位：/位 |
| `card_action_list` | `CardActionList` | [`*[]TemplateActionInfoDTO`](#model-templateactioninfodto) | JSON body | 未标注 | "卡行动点配置； 行动点，即用户可点击跳转的区块，类似按钮控件的交互； 单张卡最多定制4个行动点。" |
| `card_level_conf` | `CardLevelConf` | [`*[]TemplateCardLevelConfDTO`](#model-templatecardlevelconfdto) | JSON body | 未标注 | 卡级别配置 |
| `card_spec_tag` | `CardSpecTag` | `*string` | JSON body | 未标注 | 卡特定标签，只供特定业务使用，通常接入无需关注 |
| `column_info_list` | `ColumnInfoList` | [`*[]TemplateColumnInfoDTO`](#model-templatecolumninfodto) | JSON body | 未标注 | 栏位信息（卡包详情页面展现的栏位） |
| `field_rule_list` | `FieldRuleList` | [`*[]TemplateFieldRuleDTO`](#model-templatefieldruledto) | JSON body | 未标注 | "字段规则列表，会员卡开卡过程中，会员卡信息的生成规则， 例如：卡有效期为开卡后两年内有效，则设置为：DATE_IN_FUTURE 注意：商家会员卡场景不支持修改该内容。 " |
| `mdcode_notify_conf` | `MdcodeNotifyConf` | [`*TemplateMdcodeNotifyConfDTO`](#model-templatemdcodenotifyconfdto) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `open_card_conf` | `OpenCardConf` | [`*TemplateOpenCardConfDTO`](#model-templateopencardconfdto) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `paid_outer_card_conf` | `PaidOuterCardConf` | [`*PaidOuterCardTemplateConfDTO`](#model-paidoutercardtemplateconfdto) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `pub_channels` | `PubChannels` | [`*[]PubChannelDTO`](#model-pubchanneldto) | JSON body | 未标注 | 卡模板投放渠道 |
| `request_id` | `RequestId` | `*string` | JSON body | 未标注 | 请求 ID，商户自定义且需确保每次请求唯一。 |
| `shop_ids` | `ShopIds` | `*[]string` | JSON body | 未标注 | 会员卡上架门店 id（支付宝门店id），即发放会员卡的商家门店id。不传则保持创建模板时信息。 注意：不支持删除已有shop_id，仅支持新增。 |
| `spi_app_id` | `SpiAppId` | `*string` | JSON body | 未标注 | spi应用id，为实现spi.alipay.user.opencard.get接口的app_id。若是第三方代理模式，可以设置成服务商的 APPID 或者是商家自己的 APPID； 若不是第三方代理模式，只能设置商家自己的 APPID。 |
| `template_benefit_info` | `TemplateBenefitInfo` | [`*[]TemplateBenefitInfoDTO`](#model-templatebenefitinfodto) | JSON body | 未标注 | "权益信息， 1、在卡包的卡详情页面会自动添加权益栏位，展现会员卡特权， 2、如果添加门店渠道，则可在门店页展现会员卡的权益" |
| `template_form_config` | `TemplateFormConfig` | [`*TemplateFormConfig`](#model-templateformconfig) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `template_id` | `TemplateId` | `*string` | JSON body | 未标注 | 会员卡模板id。调用alipay.marketing.card.template.create（会员卡模板创建接口）创建模板后同步返回。 |
| `template_style_info` | `TemplateStyleInfo` | [`*TemplateStyleInfoDTO`](#model-templatestyleinfodto) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `write_off_type` | `WriteOffType` | `*string` | JSON body | 未标注 | 卡包详情页面中展现出的卡码（可用于扫码核销） (1) 静态码 qrcode: 二维码，扫码得商户开卡传入的external_card_no barcode: 条形码，扫码得商户开卡传入的external_card_no (2) 动态码-支付宝生成码值(动态码会在2分钟左右后过期) dqrcode: 动态二维码，扫码得到的码值可配合会员卡查询接口使用 dbarcode: 动态条形码，扫码得到的码值可配合会员卡查询接口使用 (3) 动态码-商家自主生成码值（码值、时效性都由商户控制） mdqrcode: 商户动态二维码，扫码得商户自主传入的码值 mdbarcode: 商户动态条码，扫码得商户自主传入的码值。 如需使用商户动态码，请联系支付宝技术支持获取相关文档。 |

<a id="api-cardtemplatequery"></a>

### CardTemplateQuery — 会员卡模板查询接口

- 官方接口：[`alipay.marketing.card.template.query`](https://opendocs.alipay.com/open-v3/0bd1a249_alipay.marketing.card.template.query?scene=871b0111e48c4490ad0a7dbf2a465b36%2C1eca10e1486044f9b3a0d9171111f810&pathHash=84fe74c8)
- 路由：`GET /v3/alipay/marketing/card/template/query`
- 请求类型：[CardTemplateQueryRequest](template_request.go#L94)
- 响应类型：[CardTemplateQueryResponse](template_response.go#L24)

会员卡模板查询接口

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `template_id` | `TemplateId` | `*string` | query | 否 | 会员卡模板id。调用alipay.marketing.card.template.create（会员卡模板创建接口）创建模板后同步返回。 |

## 嵌套请求参数

<a id="model-moreinfodto"></a>

### MoreInfoDTO

类型定义：[MoreInfoDTO](template_model.go#L17)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `descs` | `Descs` | `*[]string` | 未标注 | 描述，当operate_type为openNative必填 |
| `params` | `Params` | `*string` | 未标注 | 扩展参数，需要URL地址回带的值，JSON格式(openweb时填) |
| `title` | `Title` | `*string` | 未标注 | 二级页面标题，只有当operate_type为openNative时有效。 |
| `url` | `Url` | `*string` | 未标注 | 超链接(选择openweb的时候必须填写url参数内容) |

<a id="model-paidoutercardcyclesellconfdto"></a>

### PaidOuterCardCycleSellConfDTO

类型定义：[PaidOuterCardCycleSellConfDTO](template_model.go#L29)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `cycle_selling_url` | `CycleSellingUrl` | `*string` | 未标注 | 用户开通连续购买地址 |
| `cycle_type` | `CycleType` | `*[]string` | 未标注 | 周期购买类型。 |
| `support_cycle_sell` | `SupportCycleSell` | `*bool` | 未标注 | 是否支持用户连续购买 |

<a id="model-paidoutercardmanageurlconfdto"></a>

### PaidOuterCardManageURLConfDTO

类型定义：[PaidOuterCardManageURLConfDTO](template_model.go#L39)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `cycle_manage_url` | `CycleManageUrl` | `*string` | 未标注 | 连续购买管理地址。包括查看当前用户连续购买详情，关闭连续购买等功能 |
| `downgrade_url` | `DowngradeUrl` | `*string` | 未标注 | 付费外卡降级地址 |
| `refund_url` | `RefundUrl` | `*string` | 未标注 | 续费外卡退款地址 |
| `renew_url` | `RenewUrl` | `*string` | 未标注 | 付费外卡续费地址 |
| `upgrade_url` | `UpgradeUrl` | `*string` | 未标注 | 付费外卡升级地址 |

<a id="model-paidoutercardpricedetaildto"></a>

### PaidOuterCardPriceDetailDTO

类型定义：[PaidOuterCardPriceDetailDTO](template_model.go#L53)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `desc` | `Desc` | `*string` | 未标注 | 方案描述 |
| `price` | `Price` | `*string` | 未标注 | 售卖具体金额。单位为元。支持两位小数。 |
| `price_type` | `PriceType` | `*string` | 未标注 | 售卖金额类型 |
| `worth` | `Worth` | `*string` | 未标注 | 价值。单位为元。支持两位小数 |

<a id="model-paidoutercardsellingconfdto"></a>

### PaidOuterCardSellingConfDTO

类型定义：[PaidOuterCardSellingConfDTO](template_model.go#L65)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `end_date` | `EndDate` | `*string` | 未标注 | 售卖结束时间。 格式：yyyy-MM-dd HH:mm:ss 如果永久有效则不传。 |
| `price_detail` | `PriceDetail` | [`*[]PaidOuterCardPriceDetailDTO`](#model-paidoutercardpricedetaildto) | 未标注 | 售卖方案列表 |
| `selling_url` | `SellingUrl` | `*string` | 未标注 | 售卖地址 |
| `start_date` | `StartDate` | `*string` | 未标注 | 售卖开始时间。格式：yyyy-MM-dd HH:mm:ss |

<a id="model-paidoutercardtemplateconfdto"></a>

### PaidOuterCardTemplateConfDTO

类型定义：[PaidOuterCardTemplateConfDTO](template_model.go#L77)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `cycle_selling_conf` | `CycleSellingConf` | [`*PaidOuterCardCycleSellConfDTO`](#model-paidoutercardcyclesellconfdto) | 未标注 | 见字段类型及官方接口文档。 |
| `manage_url_conf` | `ManageUrlConf` | [`*PaidOuterCardManageURLConfDTO`](#model-paidoutercardmanageurlconfdto) | 未标注 | 见字段类型及官方接口文档。 |
| `open_selling_conf` | `OpenSellingConf` | [`*PaidOuterCardSellingConfDTO`](#model-paidoutercardsellingconfdto) | 未标注 | 见字段类型及官方接口文档。 |

<a id="model-pubchanneldto"></a>

### PubChannelDTO

类型定义：[PubChannelDTO](template_model.go#L84)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `ext_info` | `ExtInfo` | `*string` | 未标注 | 扩展信息，无需配置 |
| `pub_channel` | `PubChannel` | `*string` | 未标注 | 投放渠道类型（选择需要推广投放的渠道） SHOP_DETAIL：店铺详情页 |

<a id="model-templateactioninfodto"></a>

### TemplateActionInfoDTO

类型定义：[TemplateActionInfoDTO](template_model.go#L92)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `code` | `Code` | `*string` | 未标注 | 行动点业务CODE，商户自定义 |
| `mini_app_url` | `MiniAppUrl` | [`*TemplateActionMiniAppURLDTO`](#model-templateactionminiappurldto) | 未标注 | 见字段类型及官方接口文档。 |
| `text` | `Text` | `*string` | 未标注 | 行动点展示文案 |
| `url` | `Url` | `*string` | 未标注 | 行动点跳转链接，当url_type填"url"或不填时必填，支持http(s)和支付宝schema地址等 |
| `url_type` | `UrlType` | `*string` | 未标注 | 跳转链接类型，不填则默认为url类型： url：对应填写url参数 miniAppUrl: 对应填写mini_app_url参数，跳转至指定的支付宝小程序页面 |

<a id="model-templateactionminiappurldto"></a>

### TemplateActionMiniAppURLDTO

类型定义：[TemplateActionMiniAppURLDTO](template_model.go#L105)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `display_on_list` | `DisplayOnList` | `*string` | 未标注 | 行动点按钮是否在列表页展示，true 或 false，不填则默认false； 列表页可以配置展示最多2项小程序跳转行动点。 |
| `mini_app_id` | `MiniAppId` | `*string` | 未标注 | 小程序appId |
| `mini_page_param` | `MiniPageParam` | `*string` | 未标注 | 指定小程序页面，不填则默认跳转至对应小程序首页 |
| `mini_query_param` | `MiniQueryParam` | `*string` | 未标注 | query参数，商户自定义传参，填入值将以"query"为参数名提交至指定小程序页面 |

<a id="model-templatebenefitinfodto"></a>

### TemplateBenefitInfoDTO

类型定义：[TemplateBenefitInfoDTO](template_model.go#L117)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `benefit_desc` | `BenefitDesc` | `*[]string` | 未标注 | 权益描述信息 |
| `end_date` | `EndDate` | `*string` | 未标注 | 权益结束时间 |
| `start_date` | `StartDate` | `*string` | 未标注 | 权益开始时间 |
| `title` | `Title` | `*string` | 未标注 | 权益描述 |

<a id="model-templatecardlevelconfdto"></a>

### TemplateCardLevelConfDTO

类型定义：[TemplateCardLevelConfDTO](template_model.go#L129)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `level` | `Level` | `*string` | 未标注 | 会员级别 该级别和开卡接口中的level要一致 |
| `level_desc` | `LevelDesc` | `*string` | 未标注 | 会员级别描述 |
| `level_icon` | `LevelIcon` | `*string` | 未标注 | 会员级别对应icon图片， 会员信息预留字段，暂不在用户端展示； 通过接口（alipay.offline.material.image.upload）上传图片 |
| `level_show_name` | `LevelShowName` | `*string` | 未标注 | 会员级别显示名称 |

<a id="model-templatecolumninfodto"></a>

### TemplateColumnInfoDTO

类型定义：[TemplateColumnInfoDTO](template_model.go#L141)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `code` | `Code` | `*string` | 未标注 | 标准栏位：行为由支付宝统一定，同时已经分配标准Code BALANCE：会员卡余额 POINT：积分 LEVEL：等级 TELEPHONE：联系方式 自定义栏位：行为由商户定义，自定义Code码（只要无重复） |
| `group_title` | `GroupTitle` | `*string` | 未标注 | 若template_style_info.column_info_layout 的值为grid，此项为宫格项所属分组标题。可空。如果需要展示该项，还需支付宝内部进行特殊配置。 |
| `icon_id` | `IconId` | `*string` | 未标注 | 当template_style_info.column_info_layout 的值为grid时，此参数必填。此项为宫格项的展示icon。通过接口（alipay.offline.material.image.upload）上传图片。 |
| `more_info` | `MoreInfo` | [`*MoreInfoDTO`](#model-moreinfodto) | 未标注 | 见字段类型及官方接口文档。 |
| `operate_type` | `OperateType` | `*string` | 未标注 | 1、openNative：打开二级页面，展现 more中descs 2、openWeb：打开URL 3、staticinfo：静态信息 注意： 不填则默认staticinfo； 标准code尽量使用staticinfo，例如TELEPHONE商家电话栏位就只支持staticinfo； |
| `tag` | `Tag` | `*string` | 未标注 | 只有当template_style_info.column_info_layout 的值为grid时，此参数有效。此项为宫格项标签，最多只会展示一个标签。 |
| `title` | `Title` | `*string` | 未标注 | 栏目标题 |
| `value` | `Value` | `*string` | 未标注 | 卡包详情页面，卡栏位右边展现的值 TELEPHONE栏位的商家联系电话号码由此value字段传入 |

<a id="model-templatefieldruledto"></a>

### TemplateFieldRuleDTO

类型定义：[TemplateFieldRuleDTO](template_model.go#L160)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `field_name` | `FieldName` | `*string` | 未标注 | 模板字段规则 字段名称(用于定义会员卡开卡接口时卡相应的参数) 余额：Balance 积分：Point 等级：Level 开卡日期：OpenDate 过期日期：ValidDate |
| `rule_name` | `RuleName` | `*string` | 未标注 | 规则名 1、ASSIGN_FROM_REQUEST: 以rule_value为key值，表示该栏位的值从会员卡开卡接口中获取，会员卡开卡接口的card_info中获取对应参数值 2、DATE_IN_FUTURE: 生成一个未来的日期（格式YYYY-MM-DD)，当选择DATE_IN_FUTURE的时候，field_name 必须是OpenDate或ValidDate， 值为(10m或10d 分别表示10个月或10天) 3、CONST: 常量，会员卡开卡接口进行开卡的时候使用模板创建时候设置的值，即取rule_value的值 |
| `rule_value` | `RuleValue` | `*string` | 未标注 | 规则值，根据rule_name，采取相应取值策略 CONST：直接取rule_value作为卡属性值 DATE_IN_FUTURE：10m或10d 分别表示10个月或10天 ASSIGN_FROM_REQUEST：在开卡Reuqest请求中按rule_value取值，现在和field_name对应的为（OpenDate、ValidDate、Level、Point、Balance） |

<a id="model-templateformconfig"></a>

### TemplateFormConfig

类型定义：[TemplateFormConfig](template_model.go#L170)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `fields` | `Fields` | [`*TemplateFormFields`](#model-templateformfields) | 未标注 | 见字段类型及官方接口文档。 |
| `open_card_mini_app_id` | `OpenCardMiniAppId` | `*string` | 未标注 | 入会时承接入会弹层的小程序appId |

<a id="model-templateformfields"></a>

### TemplateFormFields

类型定义：[TemplateFormFields](template_model.go#L177)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `optional` | `Optional` | `*[]string` | 未标注 | 开卡表单可填字段，可选字段选项不能与必填字段required重复 |
| `required` | `Required` | `*[]string` | 未标注 | 开卡表单必填字段，必填字段必须有一个值，并且不能与可选字段optional里面重复 |

<a id="model-templatemdcodenotifyconfdto"></a>

### TemplateMdcodeNotifyConfDTO

类型定义：[TemplateMdcodeNotifyConfDTO](template_model.go#L185)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `ext_params` | `ExtParams` | `*string` | 未标注 | 扩展参数信息； 格式为key-value键值对； 支付宝POST请求指定url时，除BizCardNo等固定参数外，将带上ext_params中配置的所有key-value参数。 |
| `url` | `Url` | `*string` | 未标注 | 商户接收发码通知的地址链接； 只支持https地址； 用户打开会员卡时，支付宝提交POST请求此url地址，通知商户发码。 |

<a id="model-templateopencardconfdto"></a>

### TemplateOpenCardConfDTO

类型定义：[TemplateOpenCardConfDTO](template_model.go#L193)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `card_rights` | `CardRights` | [`*[]TemplateRightsContentDTO`](#model-templaterightscontentdto) | 未标注 | 领卡权益信息 |
| `conf` | `Conf` | `*string` | 未标注 | 配置，预留字段，暂时不用 |
| `open_card_source_type` | `OpenCardSourceType` | `*string` | 未标注 | 开卡渠道类型 外部系统：ISV （代理运营） 商户直连对接：MER （商家自运营） |
| `open_card_url` | `OpenCardUrl` | `*string` | 未标注 | 开卡链接，必须http、https开头 |
| `source_app_id` | `SourceAppId` | `*string` | 未标注 | 渠道APPID，提供领卡页面的服务提供方 |

<a id="model-templaterightscontentdto"></a>

### TemplateRightsContentDTO

类型定义：[TemplateRightsContentDTO](template_model.go#L207)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `detail` | `Detail` | `*string` | 未标注 | 权益内容详情 |
| `logo_id` | `LogoId` | `*string` | 未标注 | 权益icon图片Id（目前用于开卡授权页使用），通过接口（alipay.offline.material.image.upload）上传图片 |
| `title` | `Title` | `*string` | 未标注 | 权益内容标题 |

<a id="model-templatestyleinfodto"></a>

### TemplateStyleInfoDTO

类型定义：[TemplateStyleInfoDTO](template_model.go#L217)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `background_id` | `BackgroundId` | `*string` | 未标注 | 背景图片Id，通过接口（alipay.offline.material.image.upload）上传图片 图片说明：2M以内，格式：bmp、png、jpeg、jpg、gif； 尺寸不小于1020*643px； 图片不得有圆角，不得拉伸变形 |
| `banner_img_id` | `BannerImgId` | `*string` | 未标注 | banner图片id。 通过接口（alipay.offline.material.image.upload）上传图片。 |
| `banner_url` | `BannerUrl` | `*string` | 未标注 | banner跳转地址。 |
| `bg_color` | `BgColor` | `*string` | 未标注 | 字体颜色（非背景色），只影响卡详情中部信息区域字体颜色 |
| `brand_name` | `BrandName` | `*string` | 未标注 | 品牌商名称 |
| `card_show_name` | `CardShowName` | `*string` | 未标注 | 钱包端显示名称 |
| `color` | `Color` | `*string` | 未标注 | 注意：此字段已废弃。 |
| `column_info_layout` | `ColumnInfoLayout` | `*string` | 未标注 | 如果为空则默认为list。 |
| `feature_descriptions` | `FeatureDescriptions` | `*[]string` | 未标注 | 特色信息，用于领卡预览 |
| `front_image_enable` | `FrontImageEnable` | `*bool` | 未标注 | 设置是否在卡面展示（个人头像）图片信息，默认不展示； 当前仅用于身份验证信息类型的个人头像图片； 图片id随创建卡/更新卡时传入； 详见会员卡产品文档。 |
| `front_text_list_enable` | `FrontTextListEnable` | `*bool` | 未标注 | 设置是否在卡面展示文案信息，默认不展示； 文案信息分行展示，最多展示4行文案，每行文案分为label和value两部分； 文案实际内容随创建卡/更新卡时传入； 详见会员卡产品说明文档。 |
| `logo_id` | `LogoId` | `*string` | 未标注 | logo的图片ID，通过接口（alipay.offline.material.image.upload）上传图片 图片说明：1M以内，格式bmp、png、jpeg、jpg、gif； 尺寸不小于500*500px的正方形； 请优先使用商家LOGO； |
| `slogan` | `Slogan` | `*string` | 未标注 | 标语 |
| `slogan_img_id` | `SloganImgId` | `*string` | 未标注 | 标语图片id， 通过接口（alipay.offline.material.image.upload）上传图片 |
