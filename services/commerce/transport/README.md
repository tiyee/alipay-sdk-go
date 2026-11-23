# commerce/transport 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/commerce/transport`。本目录提供 4 个支付宝开放平台 V3 REST 操作。

## 使用方式

先按[项目接入说明](../../../README.md)创建 `*core.Client`，再通过 `transport.NewService(client)` 创建 Service；所有网络调用以 `context.Context` 为首个参数，支持 `opts ...core.CallOption`。响应为 `*core.Response[T]`，业务数据位于 `resp.Data`。

请求、响应与嵌套类型的链接指向本目录源码；可参考本目录的 [API 示例](api_transport_example_test.go)。

示例文件中的 `TestSandboxService_*` 可使用沙箱平台证书执行真实调用。默认 `go test` 保持离线。

## 接口目录

| SDK 方法 | 官方接口 | 用途 | HTTP |
| --- | --- | --- | --- |
| [`TransportChargerChargerBindInfoSync`](#api-transportchargerchargerbindinfosync) | [`alipay.commerce.transport.charger.chargerbindinfo.sync`](https://opendocs.alipay.com/pre-apis/05xqw6) | 充电运营商同步用户绑定充电桩信息接口 | `POST` |
| [`TransportParkingEnterInfoSync`](#api-transportparkingenterinfosync) | [`alipay.commerce.transport.parking.enterinfo.sync`](https://opendocs.alipay.com/open-v3/e38b4d57_alipay.commerce.transport.parking.enterinfo.sync) | 停车入场模板消息配置 | `POST` |
| [`TransportParkingExitInfoSync`](#api-transportparkingexitinfosync) | [`alipay.commerce.transport.parking.exitinfo.sync`](https://opendocs.alipay.com/open-v3/9986eded_alipay.commerce.transport.parking.exitinfo.sync) | 停车离场消息模板配置 | `POST` |
| [`TransportParkingPaymentInfoSync`](#api-transportparkingpaymentinfosync) | [`alipay.commerce.transport.parking.paymentinfo.sync`](https://opendocs.alipay.com/open-v3/12f22c8f_alipay.commerce.transport.parking.paymentinfo.sync) | 停车支付模板消息配置 | `POST` |

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-transportchargerchargerbindinfosync"></a>

### TransportChargerChargerBindInfoSync — 充电运营商同步用户绑定充电桩信息接口

- 官方接口：[`alipay.commerce.transport.charger.chargerbindinfo.sync`](https://opendocs.alipay.com/pre-apis/05xqw6)
- 路由：`POST /v3/alipay/commerce/transport/charger/chargerbindinfo/sync`
- 请求类型：[TransportChargerChargerBindInfoSyncRequest](transport_request.go#L5)
- 响应类型：[TransportChargerChargerBindInfoSyncResponse](transport_response.go#L5)

设备运营商通过该接口同步自己阵地中用户绑定充电桩的信息，包含绑定信息，解绑信息，设备状态等等设备信息

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `auth_token` | `AuthToken` | `*string` | query | 否 | 用户授权令牌 |
| `bind_qrcode` | `BindQrcode` | `*string` | JSON body | 未标注 | 绑定充电桩二维码值，用户扫码绑桩扫描的二维码 |
| `bind_status` | `BindStatus` | `*string` | JSON body | 未标注 | 绑定状态 1:绑定; 2:解绑; 0:未知 |
| `bind_time` | `BindTime` | `*string` | JSON body | 未标注 | 绑定时间，绑定成功必传 |
| `connector_type` | `ConnectorType` | `*string` | JSON body | 未标注 | 1：家用插座 2、交流接口插座 3：交流接口插头 4、直流接口枪头 5、无线充电座 6、其他 |
| `current_output_type` | `CurrentOutputType` | `*string` | JSON body | 未标注 | 1：国标 2：欧标 3：美标 4：日标 5：其他 |
| `equip_id` | `EquipId` | `*string` | JSON body | 未标注 | 设备编号 |
| `equip_imei` | `EquipImei` | `*string` | JSON body | 未标注 | 设备IMEI号 |
| `equip_name` | `EquipName` | `*string` | JSON body | 未标注 | 设备名称，对用户展示，可以是设备名称，也可以是设备SN号，设备编号 |
| `equip_sn` | `EquipSn` | `*string` | JSON body | 未标注 | 设备SN号 |
| `equip_status` | `EquipStatus` | `*string` | JSON body | 未标注 | 设备状态，遵循中电联协议的设备状态。0：离网;1：空闲;2：占用未充电;3：占用充电中;4：占用（预约锁定）;255：故障 |
| `equip_type_no` | `EquipTypeNo` | `*string` | JSON body | 未标注 | 设备型号 |
| `network_data_begin_time` | `NetworkDataBeginTime` | `*string` | JSON body | 未标注 | 流量开始时间 |
| `network_data_end_time` | `NetworkDataEndTime` | `*string` | JSON body | 未标注 | 流量结束时间 |
| `operator_id` | `OperatorId` | `*string` | JSON body | 未标注 | 运营商编码，一般为企业组织机构代码 |
| `operator_uid` | `OperatorUid` | `*string` | JSON body | 未标注 | 商户内部唯一标识用户的用户标识 |
| `rated_current` | `RatedCurrent` | `*string` | JSON body | 未标注 | 额定电流,单位：A |
| `rated_power` | `RatedPower` | `*string` | JSON body | 未标注 | 额定功率，单位：kW |
| `rated_voltage` | `RatedVoltage` | `*string` | JSON body | 未标注 | 额定电压，单位：V |
| `s_brand_code` | `SBrandCode` | `*string` | JSON body | 未标注 | 设备有二级品牌商，需要传二级品牌编号 |
| `unbind_time` | `UnbindTime` | `*string` | JSON body | 未标注 | 解绑时间，解绑必传 |

<a id="api-transportparkingenterinfosync"></a>

### TransportParkingEnterInfoSync — 停车入场模板消息配置

- 官方接口：[`alipay.commerce.transport.parking.enterinfo.sync`](https://opendocs.alipay.com/open-v3/e38b4d57_alipay.commerce.transport.parking.enterinfo.sync)
- 路由：`POST /v3/alipay/commerce/transport/parking/enterinfo/sync`
- 请求类型：[TransportParkingEnterInfoSyncRequest](transport_request.go#L51)
- 响应类型：[TransportParkingEnterInfoSyncResponse](transport_response.go#L15)

车辆驶入停车场时，服务商通过此接口配置消息模板内容告知支付宝发送给用户

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `agreement_query` | `AgreementQuery` | `*bool` | JSON body | 未标注 | 是否启用车牌代扣状态查询功能，true为启用，false为停用 |
| `charge_info_list` | `ChargeInfoList` | [`*[]ParkingChargeInfo`](#model-parkingchargeinfo) | JSON body | 未标注 | 本次行程收费计划清单 |
| `entrance_name` | `EntranceName` | `*string` | JSON body | 未标注 | 车辆进入停车场的入口通道名称 |
| `entrance_number` | `EntranceNumber` | `*int64` | JSON body | 未标注 | 车辆进入停车场的入口通道编号；按线下实际编号为准； |
| `free_enter_minutes` | `FreeEnterMinutes` | `*string` | JSON body | 未标注 | 当前停车场的入场免费时长分钟数 |
| `in_time` | `InTime` | `*string` | JSON body | 未标注 | 车辆入场的时间，格式"YYYY-MM-DD HH:mm:ss"，24小时制，请保证服务器时间准确，入场时间不应晚于当前网络时间 |
| `is_encrypt_plate_no` | `IsEncryptPlateNo` | `*bool` | JSON body | 未标注 | 车牌是否加密，true为加密，false为不加密，默认为false |
| `membership_info` | `MembershipInfo` | [`*ParkingMembershipInfo`](#model-parkingmembershipinfo) | JSON body | 未标注 | 见字段类型及官方接口文档。 |
| `need_charge` | `NeedCharge` | `*bool` | JSON body | 未标注 | 当前行程是否需要计费。true：需要，false：不需要。不传默认为true。 |
| `open_appid` | `OpenAppid` | `*string` | JSON body | 未标注 | 蚂蚁会员统一ID对应的归属应用appid |
| `open_id` | `OpenId` | `*string` | JSON body | 未标注 | 蚂蚁会员统一ID |
| `out_serial_no` | `OutSerialNo` | `*string` | JSON body | 未标注 | 外部停车流水号(用于串通进场与出场信息) |
| `parking_id` | `ParkingId` | `*string` | JSON body | 未标注 | 支付宝停车平台ID，由支付宝定义的该停车场标识，同一个isv或商户范围内唯一。通过 alipay.eco.mycar.parking.parkinglotinfo.create (录入停车场信息)接口获取。 |
| `plate_color` | `PlateColor` | `*string` | JSON body | 未标注 | 车牌颜色，车牌颜色，枚举支持： *BLUE：蓝。 *GREEN：绿。 *YELLOW：黄。 *WHITE：白。 *BLACK：黑。 *LIMEGREEN：黄绿色。 |
| `plate_no` | `PlateNo` | `*string` | JSON body | 未标注 | 车牌号（支持加密格式） |
| `service_url` | `ServiceUrl` | `*string` | JSON body | 未标注 | 停车服务页面地址。 1、服务商停车服务页面地址必须是支付宝小程序URL（无需转换https），详见：https://opendocs.alipay.com/support/01rb18#URL%20%E6%A0%BC%E5%BC%8F 2、若服务商没有服务链接，可传输支付宝停车官方小程序的服务链接：alipays://platformapi/startapp?appId=2021001102642986&page=pages%2Fparking-fee%2Findex 3、若此次对接的是无感支付，则服务链接传输为：alipays://platformapi/startapp?appId=2021001102642986&page=%2Fpages%2Fparking-bill%2Findex |

<a id="api-transportparkingexitinfosync"></a>

### TransportParkingExitInfoSync — 停车离场消息模板配置

- 官方接口：[`alipay.commerce.transport.parking.exitinfo.sync`](https://opendocs.alipay.com/open-v3/9986eded_alipay.commerce.transport.parking.exitinfo.sync)
- 路由：`POST /v3/alipay/commerce/transport/parking/exitinfo/sync`
- 请求类型：[TransportParkingExitInfoSyncRequest](transport_request.go#L86)
- 响应类型：[TransportParkingExitInfoSyncResponse](transport_response.go#L27)

用户要驶离停车场时，服务商通过此接口配置消息模板内容告知支付宝发送给用户

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `exit_name` | `ExitName` | `*string` | JSON body | 未标注 | 车辆出场的出口通道名称；取实地通道名称； |
| `exit_number` | `ExitNumber` | `*int64` | JSON body | 未标注 | 车辆出场的出口通道编号；与线下实地编码一致； |
| `is_encrypt_plate_no` | `IsEncryptPlateNo` | `*bool` | JSON body | 未标注 | 车牌是否加密，true为加密，false为不加密，默认为false |
| `open_appid` | `OpenAppid` | `*string` | JSON body | 未标注 | 蚂蚁会员统一ID对应的归属应用appid |
| `open_id` | `OpenId` | `*string` | JSON body | 未标注 | 蚂蚁会员统一ID |
| `out_serial_no` | `OutSerialNo` | `*string` | JSON body | 未标注 | 外部停车流水号(用于串通进场与出场信息) |
| `out_time` | `OutTime` | `*string` | JSON body | 未标注 | 车辆出场的时间，格式"YYYY-MM-DD HH:mm:ss"，24小时制，请保证服务器时间准确，出场时间不应晚于当前网络时间，也不当早于入场时间。 |
| `plate_color` | `PlateColor` | `*string` | JSON body | 未标注 | 车牌颜色，车牌颜色，枚举支持： *BLUE：蓝。 *GREEN：绿。 *YELLOW：黄。 *WHITE：白。 *BLACK：黑。 *LIMEGREEN：黄绿色。 |
| `plate_no` | `PlateNo` | `*string` | JSON body | 未标注 | 车牌号（支持加密格式） |
| `service_url` | `ServiceUrl` | `*string` | JSON body | 未标注 | 停车服务页面地址。 1、服务商停车服务页面地址必须是支付宝小程序URL（无需转换https），详见：https://opendocs.alipay.com/support/01rb18#URL%20%E6%A0%BC%E5%BC%8F 2、若服务商没有服务链接，可传输支付宝停车官方小程序的服务链接：alipays://platformapi/startapp?appId=2021001102642986&page=pages%2Fparking-fee%2Findex 3、若此次对接的是无感支付，则服务链接传输为：alipays://platformapi/startapp?appId=2021001102642986&page=%2Fpages%2Fparking-bill%2Findex |

<a id="api-transportparkingpaymentinfosync"></a>

### TransportParkingPaymentInfoSync — 停车支付模板消息配置

- 官方接口：[`alipay.commerce.transport.parking.paymentinfo.sync`](https://opendocs.alipay.com/open-v3/12f22c8f_alipay.commerce.transport.parking.paymentinfo.sync)
- 路由：`POST /v3/alipay/commerce/transport/parking/paymentinfo/sync`
- 请求类型：[TransportParkingPaymentInfoSyncRequest](transport_request.go#L110)
- 响应类型：[TransportParkingPaymentInfoSyncResponse](transport_response.go#L35)

用户在场内支付时，服务商通过此接口配置消息模板内容告知支付宝发送给用户

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `alipay_trade_no` | `AlipayTradeNo` | `*string` | JSON body | 未标注 | 支付宝交易号； |
| `discount_amount` | `DiscountAmount` | `*string` | JSON body | 未标注 | 停车费折扣金额；单位元； |
| `discount_information` | `DiscountInformation` | [`*[]ParkingPaymentDiscountInfo`](#model-parkingpaymentdiscountinfo) | JSON body | 未标注 | 停车费缴费折扣信息； |
| `free_exit_minutes` | `FreeExitMinutes` | `*string` | JSON body | 未标注 | 停车场内缴费后，允许出场的免费时长分钟数 |
| `inactive_user` | `InactiveUser` | `*bool` | JSON body | 未标注 | 365天未使用支付宝付停车费用户 |
| `is_encrypt_plate_no` | `IsEncryptPlateNo` | `*bool` | JSON body | 未标注 | 车牌是否加密，true为加密，false为不加密，默认为false |
| `mobile_number` | `MobileNumber` | `*string` | JSON body | 未标注 | 缴费用户的手机号； |
| `node_type` | `NodeType` | `*string` | JSON body | 未标注 | 表示当前请求类型：用户缴费或用户在岗亭前缴费准备离场状态 |
| `open_appid` | `OpenAppid` | `*string` | JSON body | 未标注 | 蚂蚁会员统一ID对应的归属应用appid |
| `open_id` | `OpenId` | `*string` | JSON body | 未标注 | 蚂蚁会员统一ID |
| `order_create_time` | `OrderCreateTime` | `*string` | JSON body | 未标注 | 停车岗亭前订单创建时间，格式"YYYY-MM-DD HH:mm:ss"，24小时制，请保证服务器时间准确，创建时不应晚于当前网络时间 |
| `out_order_no` | `OutOrderNo` | `*string` | JSON body | 未标注 | 服务商停车费支付订单号； |
| `out_serial_no` | `OutSerialNo` | `*string` | JSON body | 未标注 | 外部停车流水号(用于串通进场与出场信息) |
| `pay_frequency` | `PayFrequency` | `*string` | JSON body | 未标注 | 用户在支付宝侧支付停车费次数 |
| `payment_amount` | `PaymentAmount` | `*string` | JSON body | 未标注 | 停车费实付金额；单位元； |
| `payment_time` | `PaymentTime` | `*string` | JSON body | 未标注 | 场内缴费时间，格式"YYYY-MM-DD HH:mm:ss"，24小时制，请保证服务器时间准确，场内缴费不应晚于当前网络时间 |
| `payment_type` | `PaymentType` | `*string` | JSON body | 未标注 | 支付方式； ALIPAY：支付宝 WECHAT：微信 CUQP：云闪付 CASH：现金 OTHER：其他渠道 FREE: 免费 |
| `payment_user_open_id` | `PaymentUserOpenId` | `*string` | JSON body | 未标注 | 停车缴费用户的OPENID； |
| `plate_color` | `PlateColor` | `*string` | JSON body | 未标注 | 车牌颜色，车牌颜色，枚举支持： *BLUE：蓝。 *GREEN：绿。 *YELLOW：黄。 *WHITE：白。 *BLACK：黑。 *LIMEGREEN：黄绿色。 |
| `plate_no` | `PlateNo` | `*string` | JSON body | 未标注 | 车牌号（支持加密格式） |
| `service_url` | `ServiceUrl` | `*string` | JSON body | 未标注 | 停车服务页面地址。 1、服务商停车服务页面地址必须是支付宝小程序URL（无需转换https），详见： [ https://opendocs.alipay.com/support/01rb18#URL%20%E6%A0%BC%E5%BC%8F ](https://opendocs.alipay.com/support/01rb18#URL%20%E6%A0%BC%E5%BC%8F ) 2、若服务链接可沿用入场时提供的服务链接，可以留空。 3、若此次对接的是无感支付，则服务链接传输为：alipays://platformapi/startapp?appId=2021001102642986&page=%2Fpages%2Fparking-bill%2Findex |
| `total_amount` | `TotalAmount` | `*string` | JSON body | 未标注 | 停车费总金额；单位元；理论应等于实付金额+折扣金额 |

## 嵌套请求参数

以下仅展开本目录请求中实际使用的对象；字段必填规则沿用官方快照，“未标注”须结合业务说明确认。响应模型请沿各接口的响应类型链接查看。

<a id="model-parkingchargeinfo"></a>

### ParkingChargeInfo

类型定义：[ParkingChargeInfo](transport_model.go#L5)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `charge_fee` | `ChargeFee` | `*string` | 未标注 | 计费金额；单位元； |
| `charge_node` | `ChargeNode` | `*string` | 未标注 | START_CHARGING：开始计费 FEE_INCREASE：费用增加 TO_TOP：日封顶 |
| `charge_time` | `ChargeTime` | `*string` | 未标注 | 开始计费或者费用增加的时间；传东八区时间； |

<a id="model-parkingmembershipinfo"></a>

### ParkingMembershipInfo

类型定义：[ParkingMembershipInfo](transport_model.go#L15)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `membership_balance` | `MembershipBalance` | `*string` | 未标注 | 周卡、月卡：剩余天数（包括今天），单位日； 次卡：剩余次数（包括本次），单位次； 储值：余额，单位元； 折扣：折扣力度，代表0-10折；0折即为免费； |
| `membership_type` | `MembershipType` | `*string` | 未标注 | 会员类型； WEEK_CARD：周卡 MONTH_CARD：月卡 SEASON_CARD：季卡 YEAR_CARD：年卡 MULTIPLE_PASS_CARD：次卡 MONEY_STORED_CARD：储值 DISCOUNT_CARD：折扣 |

<a id="model-parkingpaymentdiscountinfo"></a>

### ParkingPaymentDiscountInfo

类型定义：[ParkingPaymentDiscountInfo](transport_model.go#L23)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `discount_amount` | `DiscountAmount` | `*string` | 未标注 | 折扣金额；单位元； |
| `discount_type` | `DiscountType` | `*string` | 未标注 | 折扣类型； COMPREHENSIVE：综合； DISCOUNT：打折； VOUCHER：券； POINTS：积分； |
