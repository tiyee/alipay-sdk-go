# iservice/ccm/servicerecord 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/iservice/ccm/servicerecord`。本目录提供 2 个支付宝开放平台 V3 REST 操作。

## 使用方式

先按[项目接入说明](../../../../README.md)创建 `*core.Client`，再通过 `servicerecord.NewService(client)` 创建 Service；所有网络调用以 `context.Context` 为首个参数，支持 `opts ...core.CallOption`。响应为 `*core.Response[T]`，业务数据位于 `resp.Data`。

请求、响应与嵌套类型的链接指向本目录源码；可参考本目录的 [API 示例](api_servicerecord_example_test.go)。

示例文件中的 `TestSandboxService_*` 可使用沙箱平台证书执行真实调用。默认 `go test` 保持离线。

## 接口目录

| SDK 方法 | 官方接口 | 用途 | HTTP |
| --- | --- | --- | --- |
| [`CcmServiceRecordCreate`](#api-ccmservicerecordcreate) | [`alipay.iservice.ccm.servicerecord.create`](https://opendocs.alipay.com/open-v3/44b1d498_alipay.iservice.ccm.servicerecord.create) | 创建服务记录 | `POST` |
| [`CcmServiceRecordModify`](#api-ccmservicerecordmodify) | [`alipay.iservice.ccm.servicerecord.modify`](https://opendocs.alipay.com/open-v3/fefe64f8_alipay.iservice.ccm.servicerecord.modify) | 变更服务记录 | `POST` |

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-ccmservicerecordcreate"></a>

### CcmServiceRecordCreate — 创建服务记录

- 官方接口：[`alipay.iservice.ccm.servicerecord.create`](https://opendocs.alipay.com/open-v3/44b1d498_alipay.iservice.ccm.servicerecord.create)
- 路由：`POST /v3/alipay/iservice/ccm/servicerecord/create`
- 请求类型：[CcmServiceRecordCreateRequest](servicerecord_request.go#L5)
- 响应类型：[CcmServiceRecordCreateResponse](servicerecord_response.go#L9)

提供给云客服租户自主创建服务记录

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `acid` | `Acid` | `*string` | JSON body | 未标注 | 来电唯一标识 |
| `aid` | `Aid` | `*string` | JSON body | 未标注 | 客服坐席号(软电话中定义的) |
| `ani` | `Ani` | `*string` | JSON body | 未标注 | 主叫号码 |
| `asid` | `Asid` | `*string` | JSON body | 未标注 | 恒生专用来电唯一标识 |
| `category_list` | `CategoryList` | [`*[]ServiceRecordCatRequest`](servicerecord_request.go#L120) | JSON body | 未标注 | 类目信息(详细信息请参考具体类型描述) |
| `chat_begin_time` | `ChatBeginTime` | `*string` | JSON body | 未标注 | 标准时间格式：yyyy-MM-dd HH:mm:ss |
| `chat_end_time` | `ChatEndTime` | `*string` | JSON body | 未标注 | 标准时间格式：yyyy-MM-dd HH:mm:ss |
| `city` | `City` | `*string` | JSON body | 未标注 | 来电地址 |
| `creator_id` | `CreatorId` | `*string` | JSON body | 未标注 | 创建者ID |
| `discon_symbol` | `DisconSymbol` | `*string` | JSON body | 未标注 | 先挂断方 AGENT:客服 CUSTOMER:客户 |
| `dnis` | `Dnis` | `*string` | JSON body | 未标注 | 被叫号码 |
| `file_size` | `FileSize` | `*int64` | JSON body | 未标注 | 文件大小 |
| `gmt_create` | `GmtCreate` | `*string` | JSON body | 未标注 | 创建时间(yyyy-MM-dd HH:mm:ss) |
| `gmt_modified` | `GmtModified` | `*string` | JSON body | 未标注 | 修改时间 |
| `memo` | `Memo` | `*string` | JSON body | 未标注 | 服务记录备注信息 |
| `modifier_id` | `ModifierId` | `*string` | JSON body | 未标注 | 修改人ID |
| `outbound_task_id` | `OutboundTaskId` | `*string` | JSON body | 未标注 | 外呼任务ID |
| `satisfaction` | `Satisfaction` | `*string` | JSON body | 未标注 | 满意度 0:非常满意 1:满意 2:一般 3:不满意 4:非常不满意 |
| `satisfaction_memo` | `SatisfactionMemo` | `*string` | JSON body | 未标注 | 满意度补充信息 |
| `service_source` | `ServiceSource` | `*string` | JSON body | 未标注 | 1:客户来电 4:电话回访 5:网点服务 6:在线服务 7:外呼任务 8:手动外呼 9:在线离线留言 |
| `service_time` | `ServiceTime` | `*int64` | JSON body | 未标注 | 服务时长(单位:秒) |
| `skillgroup_id` | `SkillgroupId` | `*string` | JSON body | 未标注 | 技能组ID |
| `skillgroup_name` | `SkillgroupName` | `*string` | JSON body | 未标注 | 技能组名称(技能组ID对应的技能组名称) |
| `tnt_inst_id` | `TntInstId` | `*string` | JSON body | 未标注 | 业务租户 |
| `user_id` | `UserId` | `*string` | JSON body | 未标注 | 来访用户唯一标识ID |
| `verify_result` | `VerifyResult` | `*string` | JSON body | 未标注 | IVR验证结果 ROLE_N:未定位到用户 CERTNO_Y_PWD_Y:证件号码+交易密码验证通过 CERTNO_Y_PWD_N:证件号码验证通过+交易密码不通过 CERTNO_Y_PWD_UNKNOW:证件号码验证通过 PHONE_Y_PWD_Y:来电手机+交易密码验证通过 PHONE_Y_PWD_N:来电手机+交易密码未通过 PHONE_Y_PWD_UNKNOW:来电手机账户 SEARCH_Y_PWD_Y:搜索定位用户+推屏交易密码通过 SEARCH_Y_PWD_N:搜索定位用户+推屏交易密码未通过 SEARCH_Y_PWD_UNKNOW:搜索定位用户 |

<a id="api-ccmservicerecordmodify"></a>

### CcmServiceRecordModify — 变更服务记录

- 官方接口：[`alipay.iservice.ccm.servicerecord.modify`](https://opendocs.alipay.com/open-v3/fefe64f8_alipay.iservice.ccm.servicerecord.modify)
- 路由：`POST /v3/alipay/iservice/ccm/servicerecord/modify`
- 请求类型：[CcmServiceRecordModifyRequest](servicerecord_request.go#L61)
- 响应类型：[CcmServiceRecordModifyResponse](servicerecord_response.go#L15)

提供给云客服租户自主变更服务记录

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `acid` | `Acid` | `*string` | JSON body | 未标注 | 来电唯一标识 |
| `aid` | `Aid` | `*string` | JSON body | 未标注 | 客服坐席号(软电话中定义的) |
| `ani` | `Ani` | `*string` | JSON body | 未标注 | 主叫号码 |
| `asid` | `Asid` | `*string` | JSON body | 未标注 | 恒生专用来电唯一标识 |
| `category_list` | `CategoryList` | [`*ServiceRecordCatRequest`](servicerecord_request.go#L120) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `chat_begin_time` | `ChatBeginTime` | `*string` | JSON body | 未标注 | 标准时间格式：yyyy-MM-dd HH:mm:ss |
| `chat_end_time` | `ChatEndTime` | `*string` | JSON body | 未标注 | 标准时间格式：yyyy-MM-dd HH:mm:ss |
| `city` | `City` | `*string` | JSON body | 未标注 | 来电地址 |
| `creator_id` | `CreatorId` | `*string` | JSON body | 未标注 | 创建者ID |
| `discon_symbol` | `DisconSymbol` | `*string` | JSON body | 未标注 | 先挂断方 AGENT:客服 CUSTOMER:客户 |
| `dnis` | `Dnis` | `*string` | JSON body | 未标注 | 被叫号码 |
| `file_size` | `FileSize` | `*int64` | JSON body | 未标注 | 文件大小 |
| `gmt_create` | `GmtCreate` | `*string` | JSON body | 未标注 | 创建时间(yyyy-MM-dd HH:mm:ss) |
| `gmt_modified` | `GmtModified` | `*string` | JSON body | 未标注 | 修改时间 |
| `id` | `Id` | `*string` | JSON body | 未标注 | 云客服服务记录ID |
| `memo` | `Memo` | `*string` | JSON body | 未标注 | 服务记录备注信息 |
| `modifier_id` | `ModifierId` | `*string` | JSON body | 未标注 | 修改人ID |
| `outbound_task_id` | `OutboundTaskId` | `*string` | JSON body | 未标注 | 外呼任务ID |
| `satisfaction` | `Satisfaction` | `*string` | JSON body | 未标注 | 满意度 0:非常满意 1:满意 2:一般 3:不满意 4:非常不满意 |
| `satisfaction_memo` | `SatisfactionMemo` | `*string` | JSON body | 未标注 | 满意度补充信息 |
| `service_source` | `ServiceSource` | `*string` | JSON body | 未标注 | 1:客户来电 4:电话回访 5:网点服务 6:在线服务 7:外呼任务 8:手动外呼 9:在线离线留言 |
| `service_time` | `ServiceTime` | `*int64` | JSON body | 未标注 | 服务时长(单位:秒) |
| `skillgroup_id` | `SkillgroupId` | `*string` | JSON body | 未标注 | 技能组ID |
| `skillgroup_name` | `SkillgroupName` | `*string` | JSON body | 未标注 | 技能组名称(技能组ID对应的技能组名称) |
| `status` | `Status` | `*string` | JSON body | 未标注 | 1:待处理 2:暂存 3:完结 4:废除 5:无效 6:未知 |
| `tnt_inst_id` | `TntInstId` | `*string` | JSON body | 未标注 | 业务租户 |
| `user_id` | `UserId` | `*string` | JSON body | 未标注 | 来访用户唯一标识ID |
| `verify_result` | `VerifyResult` | `*string` | JSON body | 未标注 | IVR验证结果 ROLE_N:未定位到用户 CERTNO_Y_PWD_Y:证件号码+交易密码验证通过 CERTNO_Y_PWD_N:证件号码验证通过+交易密码不通过 CERTNO_Y_PWD_UNKNOW:证件号码验证通过 PHONE_Y_PWD_Y:来电手机+交易密码验证通过 PHONE_Y_PWD_N:来电手机+交易密码未通过 PHONE_Y_PWD_UNKNOW:来电手机账户 SEARCH_Y_PWD_Y:搜索定位用户+推屏交易密码通过 SEARCH_Y_PWD_N:搜索定位用户+推屏交易密码未通过 SEARCH_Y_PWD_UNKNOW:搜索定位用户 |
