# batch 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/trade/batch`。本目录提供 4 个支付宝开放平台 V3 REST 操作。

## 使用方式

先按[项目接入说明](../../../README.md)创建 `*core.Client`，再通过 `batch.NewService(client)` 创建 Service；所有网络调用以 `context.Context` 为首个参数，支持 `opts ...core.CallOption`。响应为 `*core.Response[T]`，业务数据位于 `resp.Data`。

请求、响应与嵌套类型的链接指向本目录源码；可参考本目录的 [API 示例](api_batch_example_test.go)。

示例文件中的 `TestSandboxService_*` 可使用沙箱平台证书执行真实调用。默认 `go test` 保持离线。

返回[交易领域目录](../README.md)。请求、响应和专属模型位于本包，跨包共享类型从 `trade` 包引用。

## 接口目录

| SDK 方法 | 官方接口 | 用途 | HTTP |
| --- | --- | --- | --- |
| [`BatchSettle`](#api-batchsettle) | [`alipay.trade.batch.settle`](https://opendocs.alipay.com/pre-apis/05vfpn) | 批量结算 | `POST` |
| [`BatchSettleQuery`](#api-batchsettlequery) | [`alipay.trade.batch.settle.query`](https://opendocs.alipay.com/pre-apis/05vel6) | 批量结算查询 | `POST` |
| [`BatchTransfer`](#api-batchtransfer) | [`alipay.trade.batch.transfer`](https://opendocs.alipay.com/pre-apis/05vd9z) | 请款分账 | `POST` |
| [`BatchTransferQuery`](#api-batchtransferquery) | [`alipay.trade.batch.transfer.query`](https://opendocs.alipay.com/pre-apis/05vsbp) | 请款分账查询接口 | `GET` |

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-batchsettle"></a>

### BatchSettle — 批量结算

- 官方接口：[`alipay.trade.batch.settle`](https://opendocs.alipay.com/pre-apis/05vfpn)
- 路由：`POST /v3/alipay/trade/batch/settle`
- 请求类型：[BatchSettleRequest](batch_request.go#L53)
- 响应类型：[BatchSettleResponse](batch_response.go#L32)

买家付款成功之后，资金进入二级商户待结算户，支付宝等待平台的结算指令进行结算。平台可根据平台规则，发起批量结算请求，将资金结算到二级商户的结算账号。请款金额只要不超过待结算的总余额即可。

请求约束：BatchSettleRequest 是 alipay.trade.batch.settle（批量结算）的请求。
out_request_no 与 settle_clauses 必选。写操作：重复的请求流水号
返回既有单据状态，SDK 不自动重试。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `out_request_no` | `OutRequestNo` | `string` | JSON body | 未标注 | 结算请求外部流水号，32个字符以内、可包含字母、数字、下划线；需保证在商户端不重复，如果重复则返回该流水号对应的结算单据的状态。<br>SDK 字段说明：结算请求外部流水号：32 字符以内，可包含字母、<br>数字、下划线，需保证在商户端不重复；重复时返回既有单据状态。 |
| `settle_clauses` | `SettleClauses` | [`[]SettleClause`](#model-settleclause) | JSON body | 未标注 | 结算明细条款 |
| `biz_product` | `BizProduct` | `*string` | JSON body | 未标注 | 收单产品码，商家和支付宝签约的产品码 |
| `settle_type` | `SettleType` | `*string` | JSON body | 未标注 | 结算方式，目前仅支持提前放款结算quickSettlement，提前放款结算需要商户开通快收服务，不填则为普通结算方式<br>SDK 字段说明：结算方式；目前仅支持提前放款结算 quickSettlement。 |
| `extend_params` | `ExtendParams` | `*string` | JSON body | 未标注 | 扩展参数 |

注意事项：

- 重复的 OutRequestNo 返回既有单据状态；受理结果以 ResultCode 为准，
处理结果以批量结算查询为准。

<a id="api-batchsettlequery"></a>

### BatchSettleQuery — 批量结算查询

- 官方接口：[`alipay.trade.batch.settle.query`](https://opendocs.alipay.com/pre-apis/05vel6)
- 路由：`POST /v3/alipay/trade/batch/settle/query`
- 请求类型：[BatchSettleQueryRequest](batch_request.go#L70)
- 响应类型：[BatchSettleQueryResponse](batch_response.go#L44)

商户请求批量结算结果查询

请求约束：BatchSettleQueryRequest 是 alipay.trade.batch.settle.query
（批量结算查询）的请求：settle_no 与 out_request_no 必填其一，
同时存在时优先取 settle_no。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `settle_no` | `SettleNo` | `*string` | JSON body | 未标注 | 结算单据号，与out_request_no二者必填其一 |
| `out_request_no` | `OutRequestNo` | `*string` | JSON body | 未标注 | 批量结算时传入的外部流水号，与settle_no二者必填其一。 out_request_no和settle_no如果同时存在，优先取settle_no |
| `extend_params` | `ExtendParams` | `*string` | JSON body | 未标注 | 扩展参数 |

<a id="api-batchtransfer"></a>

### BatchTransfer — 请款分账

- 官方接口：[`alipay.trade.batch.transfer`](https://opendocs.alipay.com/pre-apis/05vd9z)
- 路由：`POST /v3/alipay/trade/batch/transfer`
- 请求类型：[BatchTransferRequest](batch_request.go#L10)
- 响应类型：[BatchTransferResponse](batch_response.go#L8)

在红星美凯龙场景下，请款到不同主体的银行卡/账户中

请求约束：BatchTransferRequest 是 alipay.trade.batch.transfer（请款分账）的请求。
out_request_no 与 royalty_parameters 必选；多个分账条款要求出账方
只能有一个。写操作：重复请求流水号返回既有单据状态。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `out_request_no` | `OutRequestNo` | `string` | JSON body | 未标注 | 结算请求外部流水号，32个字符以内、可包含字母、数字、下划线；需保证在商户端不重复<br>SDK 字段说明：结算请求外部流水号：32 字符以内，可包含字母、<br>数字、下划线，需保证在商户端不重复。 |
| `royalty_parameters` | `RoyaltyParameters` | [`[]BatchRoyaltyClause`](#model-batchroyaltyclause) | JSON body | 未标注 | 分账条款明细 特别注意：最好一次只传一个分账条款。 如果传多个，则要求出账方只能有一个， 举例：A分给B+A分给C+A分给D=OK A分给B+C分给D=不行 |
| `extend_params` | `ExtendParams` | `*string` | JSON body | 未标注 | 业务扩展参数 |

注意事项：

- 多个分账条款要求出账方只能有一个；处理结果以查询接口为准。

<a id="api-batchtransferquery"></a>

### BatchTransferQuery — 请款分账查询接口

- 官方接口：[`alipay.trade.batch.transfer.query`](https://opendocs.alipay.com/pre-apis/05vsbp)
- 路由：`GET /v3/alipay/trade/batch/transfer/query`
- 请求类型：[BatchTransferQueryRequest](batch_request.go#L23)
- 响应类型：[BatchTransferQueryResponse](batch_response.go#L20)

针对请款分账接口的查询功能

请求约束：BatchTransferQueryRequest 是 alipay.trade.batch.transfer.query
（请款分账查询，V3 GET）的请求：settle_no 与 out_request_no
必填其一，同时存在时优先取 settle_no。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `settle_no` | `SettleNo` | `*string` | query | 否 | 结算单据号，与out_request_no二者必填其一 |
| `out_request_no` | `OutRequestNo` | `*string` | query | 否 | 批量结算时传入的外部流水号，与settle_no二者必填其一。 out_request_no和settle_no如果同时存在，优先取settle_no |
| `extend_params` | `ExtendParams` | `*string` | query | 否 | 扩展字段 |

## 嵌套请求参数

以下仅展开本目录请求中实际使用的对象；字段必填规则沿用官方快照，“未标注”须结合业务说明确认。响应模型请沿各接口的响应类型链接查看。

<a id="model-batchroyaltyclause"></a>

### BatchRoyaltyClause

类型定义：[BatchRoyaltyClause](batch_model.go#L21)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `scene` | `Scene` | `*string` | 未标注 | 分账场景，目前仅支持5中类型，RETURN_MONEY(垫资还款), CHARGE（费用）, REPLENISH（补贴），FUND_TRANS_IN_SWITCHING（存量资金搬迁），PUNISHMENT（处罚）<br>SDK 字段说明：分账场景：RETURN_MONEY（垫资还款）、CHARGE（费用）、<br>REPLENISH（补贴）、FUND_TRANS_IN_SWITCHING（存量资金搬迁）、<br>PUNISHMENT（处罚）。 |
| `amount` | `Amount` | `*string` | 未标注 | 分账金额，单位为元 |
| `desc` | `Desc` | `*string` | 未标注 | 分账描述 |
| `trans_in_account_type` | `TransInAccountType` | `*string` | 未标注 | 分账账户类型。 bankCard: 分账账户为银行卡； alipayBalance: 分账账户为支付宝余额户; defaultSettle: 按默认结算规则分账<br>SDK 字段说明：分账账户类型：bankCard、alipayBalance 或<br>defaultSettle。 |
| `trans_in_account_id_type` | `TransInAccountIDType` | `*string` | 未标注 | 分账转入账户id类型。 当trans_in_account_type 为bankCard时，本参数为cardSerialNo，表示分账账户id是银行卡编号; 当trans_in_account_type 为alipayBalance时，本参数为userId或者loginName或者openId，其中userId表示分账账户id是支付宝唯一用户号，loginName表示分账账户id是支付宝登录号，openId表示分账账户id是支付宝OpenId; 当 trans_in_account_type 为 defaultSettle 时，本参数必须为空。<br>SDK 字段说明：分账账户 id 类型。 |
| `trans_in_account_id` | `TransInAccountID` | `*string` | 未标注 | 分账转入账户id。当分账账户id类型是cardSerialNo时，本参数为用户在支付宝绑定的卡编号；当分账账户id类型是userId时，本参数为用户的支付宝账号对应的支付宝唯一用户号；当分账账户id类型是loginName时，本参数为用户的支付宝登录号;当分账账户id类型是openId时，本参数为用户的在该应用下的支付宝OpenId; 当 trans_in_account_type 为defaultSettle时，本参数必须为空<br>SDK 字段说明：分账转入账户 id。 |
| `trans_in_entity_type` | `TransInEntityType` | `*string` | 未标注 | 分账转入主体类型。 SecondMerchant：分账转入主体为二级商户 Store: 分账转入主体为门店 |
| `trans_in_entity_id` | `TransInEntityID` | `*string` | 未标注 | 分账转出主体账 号。<br> 当分账转出主体类型为SecondMerchant，本参数为二级商户的SecondMerchantID<br> 当分账转出类型为Store，本参数为StoreID<br>SDK 字段说明：分账转入主体账号。 |
| `trans_in_entity_biz_type` | `TransInEntityBizType` | `*string` | 未标注 | 分账账户业务类型，目前仅支持已结算类型 settled<br>SDK 字段说明：分账转入主体业务类型（目前仅支持 settled）。 |
| `trans_in_sub_merchant` | `TransInSubMerchant` | [`*trade.SubMerchant`](../trade_model.go#L628) | 未标注 | 分账转入方二级商户信息。 |
| `trans_out_entity_type` | `TransOutEntityType` | `*string` | 未标注 | 分账转出主体类型。 SecondMerchant：结算主体为二级商户 Store: 结算主体为门店 |
| `trans_out_entity_id` | `TransOutEntityID` | `*string` | 未标注 | 分账转出主体账。<br> 当分账转出主体类型为SecondMerchant，本参数为二级商户的SecondMerchantId<br> 当分账转出类型为Store，本参数为StoreID<br>SDK 字段说明：分账转出主体账号。 |
| `trans_out_entity_biz_type` | `TransOutEntityBizType` | `*string` | 未标注 | 分账转出方主体业务类型，目前仅支持已结算类型 settled |
| `trans_out_sub_merchant` | `TransOutSubMerchant` | [`*trade.SubMerchant`](../trade_model.go#L628) | 未标注 | 分账转出方二级商户信息。 |

<a id="model-settleclause"></a>

### SettleClause

类型定义：[SettleClause](batch_model.go#L90)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `settle_entity_type` | `SettleEntityType` | `*string` | 未标注 | 结算主体类型。 SecondMerchant：结算主体为二级商户 Store：结算主体为门店<br>SDK 字段说明：结算主体类型：SecondMerchant 或 Store。 |
| `settle_entity_id` | `SettleEntityID` | `*string` | 未标注 | 结算主体账号。 当结算主体类型为SecondMerchant，传二级商户ID(smid)；结算主体类型为Store时，传门店ID<br>SDK 字段说明：结算主体账号：二级商户 ID 或门店 ID。 |
| `settle_entity_biz_type` | `SettleEntityBizType` | `*string` | 未标注 | 已结算资产<br>SDK 字段说明：已结算资产类型。 |
| `settle_account_type` | `SettleAccountType` | `*string` | 未标注 | 结算账户类型。 bankCard: 结算账户为银行卡； alipayBalance: 结算账户为支付宝余额户； defaultSettle: 按默认结算规则结算<br>SDK 字段说明：结算账户类型：bankCard、alipayBalance 或<br>defaultSettle（按默认结算规则结算）。 |
| `settle_account_id_type` | `SettleAccountIDType` | `*string` | 未标注 | 结算账户id类型。当 settle_account_type 为 defaultSettle 时，本参数必须为空。 当settle_account_type 为bankCard时，本参数为cardSerialNo，表示结算账户id是银行卡编号; 当settle_account_type 为alipayBalance时，本参数为userId或者loginName，其中userId表示结算账户id是支付宝唯一用户号，loginName表示结算账户id是支付宝登录号。<br>SDK 字段说明：结算账户 id 类型；defaultSettle 时必须为空。 |
| `settle_account_id` | `SettleAccountID` | `*string` | 未标注 | 结算账户id。 当结算账户id类型是cardSerialNo时，本参数为用户在支付宝绑定的卡编号；当结算账户id类型是userId时，本参数为用户的支付宝账号对应的支付宝唯一用户号；当结算账户id类型是loginName时，本参数为用户的支付宝登录号; 当 settle_account_type 为defaultSettle时，本参数必须为空<br>SDK 字段说明：结算账户 id。 |
| `settle_account_entity` | `SettleAccountEntity` | `*string` | 未标注 | 结算账号所属于的结算主体类型。 secondMerchant 二级商户的结算账号 store 门店的结算账号 如果不填，默认为settleEntityType对应的类型<br>SDK 字段说明：结算账号所属的结算主体类型。 |
| `amount` | `Amount` | `*string` | 未标注 | 结算金额，单位为元 |
| `currency` | `Currency` | `*string` | 未标注 | 结算币种，支持人民币：CNY<br>SDK 字段说明：结算币种；目前支持人民币 CNY。 |
| `sub_merchant` | `SubMerchant` | [`*trade.SubMerchant`](../trade_model.go#L628) | 未标注 | 二级商户信息。 |
