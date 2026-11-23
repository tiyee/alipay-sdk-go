# parking 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/eco/mycar/parking`。本包提供 16 个支付宝 V3 REST 操作。

通过 `parking.NewService(executor)` 复用 `*core.Client` 或最小 `core.Executor`。网络调用以 `context.Context` 为首参数，支持 `opts ...core.CallOption`，业务数据位于 `resp.Data`。

接口方法按单词边界使用大驼峰命名；请求、响应与专属模型位于对应业务包，共享模型从共同父包引用。

接入方式见[项目说明](../../../../README.md)。默认测试离线执行。

返回[领域目录](../README.md)。

## 接口目录

| SDK 方法 | 官方接口 |
| --- | --- |
| [`MyCarParkingChargeInfoSync`](#api-mycarparkingchargeinfosync) | `alipay.eco.mycar.parking.chargeinfo.sync` |
| [`MyCarParkingChargingInfoSync`](#api-mycarparkingcharginginfosync) | `alipay.eco.mycar.parking.charginginfo.sync` |
| [`MyCarParkingConfigQuery`](#api-mycarparkingconfigquery) | `alipay.eco.mycar.parking.config.query` |
| [`MyCarParkingConfigSet`](#api-mycarparkingconfigset) | `alipay.eco.mycar.parking.config.set` |
| [`MyCarParkingEnterInfoSync`](#api-mycarparkingenterinfosync) | `alipay.eco.mycar.parking.enterinfo.sync` |
| [`MyCarParkingExitInfoSync`](#api-mycarparkingexitinfosync) | `alipay.eco.mycar.parking.exitinfo.sync` |
| [`MyCarParkingOrderSync`](#api-mycarparkingordersync) | `alipay.eco.mycar.parking.order.sync` |
| [`MyCarParkingOrderUpdate`](#api-mycarparkingorderupdate) | `alipay.eco.mycar.parking.order.update` |
| [`MyCarParkingOvertimeChargingInfoSync`](#api-mycarparkingovertimecharginginfosync) | `alipay.eco.mycar.parking.overtimecharginginfo.sync` |
| [`MyCarParkingParkingLotBizTransfer`](#api-mycarparkingparkinglotbiztransfer) | `alipay.eco.mycar.parking.parkinglotbiz.transfer` |
| [`MyCarParkingParkingLotInfoCreate`](#api-mycarparkingparkinglotinfocreate) | `alipay.eco.mycar.parking.parkinglotinfo.create` |
| [`MyCarParkingParkingLotInfoQuery`](#api-mycarparkingparkinglotinfoquery) | `alipay.eco.mycar.parking.parkinglotinfo.query` |
| [`MyCarParkingParkingLotInfoUpdate`](#api-mycarparkingparkinglotinfoupdate) | `alipay.eco.mycar.parking.parkinglotinfo.update` |
| [`MyCarParkingPaymentInfoSync`](#api-mycarparkingpaymentinfosync) | `alipay.eco.mycar.parking.paymentinfo.sync` |
| [`MyCarParkingSpaceInfoSync`](#api-mycarparkingspaceinfosync) | `alipay.eco.mycar.parking.spaceinfo.sync` |
| [`MyCarParkingVehicleQuery`](#api-mycarparkingvehiclequery) | `alipay.eco.mycar.parking.vehicle.query` |

## 使用示例

- [MyCarParkingChargeInfoSync](api_parking_mycarparkingchargeinfosync_example_test.go)
- [MyCarParkingChargingInfoSync](api_parking_mycarparkingcharginginfosync_example_test.go)
- [MyCarParkingConfigQuery](api_parking_mycarparkingconfigquery_example_test.go)
- [MyCarParkingConfigSet](api_parking_mycarparkingconfigset_example_test.go)
- [MyCarParkingEnterInfoSync](api_parking_mycarparkingenterinfosync_example_test.go)
- [MyCarParkingExitInfoSync](api_parking_mycarparkingexitinfosync_example_test.go)
- [MyCarParkingOrderSync](api_parking_mycarparkingordersync_example_test.go)
- [MyCarParkingOrderUpdate](api_parking_mycarparkingorderupdate_example_test.go)
- [MyCarParkingOvertimeChargingInfoSync](api_parking_mycarparkingovertimecharginginfosync_example_test.go)
- [MyCarParkingParkingLotBizTransfer](api_parking_mycarparkingparkinglotbiztransfer_example_test.go)
- [MyCarParkingParkingLotInfoCreate](api_parking_mycarparkingparkinglotinfocreate_example_test.go)
- [MyCarParkingParkingLotInfoQuery](api_parking_mycarparkingparkinglotinfoquery_example_test.go)
- [MyCarParkingParkingLotInfoUpdate](api_parking_mycarparkingparkinglotinfoupdate_example_test.go)
- [MyCarParkingPaymentInfoSync](api_parking_mycarparkingpaymentinfosync_example_test.go)
- [MyCarParkingSpaceInfoSync](api_parking_mycarparkingspaceinfosync_example_test.go)
- [MyCarParkingVehicleQuery](api_parking_mycarparkingvehiclequery_example_test.go)

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-mycarparkingchargeinfosync"></a>

### MyCarParkingChargeInfoSync — 停车场价格信息同步

- 官方接口：[`alipay.eco.mycar.parking.chargeinfo.sync`](https://opendocs.alipay.com/open-v3/18d2ca24_alipay.eco.mycar.parking.chargeinfo.sync)
- 路由：`POST /v3/alipay/eco/mycar/parking/chargeinfo/sync`
- 请求类型：[MyCarParkingChargeInfoSyncRequest](parking_request.go#L4)
- 响应类型：[MyCarParkingChargeInfoSyncResponse](parking_response.go#L8)

停车场价格信息同步

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `business_hours` | `BusinessHours` | `*string` | JSON body | 未标注 | 营业时间 |
| `car_daylight_advanced_price` | `CarDaylightAdvancedPrice` | `*string` | JSON body | 未标注 | 小车-白天 非首个计费单元价格 |
| `car_daylight_base_price` | `CarDaylightBasePrice` | `*string` | JSON body | 未标注 | 小车-白天 首个计费单元价格 |
| `car_night_advanced_price` | `CarNightAdvancedPrice` | `*string` | JSON body | 未标注 | 小车-夜间 非首个计费单元价格 |
| `car_night_base_price` | `CarNightBasePrice` | `*string` | JSON body | 未标注 | 小车-夜间 首个计费单元价格 |
| `car_onetime_price` | `CarOnetimePrice` | `*string` | JSON body | 未标注 | 小车 计次收费价格<br> xx元 / 次，按次收费的车场填写 |
| `daily_price_upperbound` | `DailyPriceUpperbound` | `*string` | JSON body | 未标注 | 单日停车封顶价格：具体值或“不封顶” |
| `daylight_business_hours` | `DaylightBusinessHours` | `*string` | JSON body | 未标注 | 车场白天（收费）营业时间段 |
| `free_period` | `FreePeriod` | `*string` | JSON body | 未标注 | 单次停车免费时长 |
| `has_charging_pile` | `HasChargingPile` | `*string` | JSON body | 未标注 | 是否有充电桩: -1 未知,0 否,1 是 |
| `is_charge` | `IsCharge` | `*string` | JSON body | 未标注 | 是否收费：-1 未知,0 否,1 是 |
| `night_business_hours` | `NightBusinessHours` | `*string` | JSON body | 未标注 | 车场夜间（收费）营业时间段 |
| `parking_id` | `ParkingId` | `*string` | JSON body | 未标注 | 停车场ID |
| `parking_space_count` | `ParkingSpaceCount` | `*int64` | JSON body | 未标注 | 总车位数 |
| `remark` | `Remark` | `*string` | JSON body | 未标注 | 复杂计费补充描述<br> 收费复杂无法格式化的车场填写，以文本形式表述 |
| `truck_daylight_advanced_price` | `TruckDaylightAdvancedPrice` | `*string` | JSON body | 未标注 | 大车-白天 非首个计费单元价格 |
| `truck_daylight_base_price` | `TruckDaylightBasePrice` | `*string` | JSON body | 未标注 | 大车-白天 首个计费单元价格 |
| `truck_night_advanced_price` | `TruckNightAdvancedPrice` | `*string` | JSON body | 未标注 | 大车-夜间 非首个计费单元价格 |
| `truck_night_base_price` | `TruckNightBasePrice` | `*string` | JSON body | 未标注 | 大车-夜间 首个计费单元价格 |
| `truck_onetime_price` | `TruckOnetimePrice` | `*string` | JSON body | 未标注 | 大车 计次收费价格<br> xx元 / 次，按次收费的车场填写 |

<a id="api-mycarparkingcharginginfosync"></a>

### MyCarParkingChargingInfoSync — 车辆停车计费接口

- 官方接口：[`alipay.eco.mycar.parking.charginginfo.sync`](https://opendocs.alipay.com/open-v3/91cababc_alipay.eco.mycar.parking.charginginfo.sync)
- 路由：`POST /v3/alipay/eco/mycar/parking/charginginfo/sync`
- 请求类型：[MyCarParkingChargingInfoSyncRequest](parking_request.go#L48)
- 响应类型：[MyCarParkingChargingInfoSyncResponse](parking_response.go#L14)

上传车辆在停车场的计费信息，将停车动作推进到计费状态，上传信息通过该接口提交到支付宝，支付宝返回对应的信息

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `car_number` | `CarNumber` | `*string` | JSON body | 未标注 | 车牌号 |
| `isv_url` | `IsvUrl` | `*string` | JSON body | 未标注 | 智能助理当前的跳转链接 |
| `parking_id` | `ParkingId` | `*string` | JSON body | 未标注 | 支付宝停车平台ID，由支付宝定义的该停车场标识，同一个isv或商户范围内唯一。通过 alipay.eco.mycar.parking.parkinglotinfo.create(录入停车场信息)接口获取。 |
| `serial_no` | `SerialNo` | `*string` | JSON body | 未标注 | 支付宝业务流水号，用于记录车辆从驶入到驶出的全流程。通过 alipay.eco.mycar.parking.enterinfo.sync (车辆驶入接口)接口获取。 |

<a id="api-mycarparkingconfigquery"></a>

### MyCarParkingConfigQuery — ISV系统配置查询接口

- 官方接口：[`alipay.eco.mycar.parking.config.query`](https://opendocs.alipay.com/open-v3/735dfdcb_alipay.eco.mycar.parking.config.query)
- 路由：`GET /v3/alipay/eco/mycar/parking/config/query`
- 请求类型：[MyCarParkingConfigQueryRequest](parking_request.go#L418)
- 响应类型：[MyCarParkingConfigQueryResponse](parking_response.go#L17)

ISV通过该接口，查询ISV已注册到车主平台停车业务系统配置信息。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `interface_name` | `InterfaceName` | `*string` | query | 否 | 接口名称。H5传入参数固定值：alipay.eco.mycar.parking.userpage.query； 小程序传入参数固定值：alipay.eco.mycar.parking.isv.homepage； 协议状态变更通知传入参数固定值： alipay.eco.mycar.parking.agreement.notify。 |
| `interface_type` | `InterfaceType` | `*string` | query | 否 | 接口类型。H5、小程序传入参数固定值：interface_page； 通知接口传入参数固定值：interface_service。 |

<a id="api-mycarparkingconfigset"></a>

### MyCarParkingConfigSet — 停车ISV系统配置接口

- 官方接口：[`alipay.eco.mycar.parking.config.set`](https://opendocs.alipay.com/open-v3/17f00a51_alipay.eco.mycar.parking.config.set)
- 路由：`POST /v3/alipay/eco/mycar/parking/config/set`
- 请求类型：[MyCarParkingConfigSetRequest](parking_request.go#L60)
- 响应类型：[MyCarParkingConfigSetResponse](parking_response.go#L30)

ISV通过该接口，配置车主平台各业务所需要的ISV的系统配置信息。如果接口配置存在则覆盖原有的接口信息。通过该接口设置的配置信息是立刻生效的，在调用该接口修改配置信息时请先评估对线上业务的影响，并做好处理方案。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `account_no` | `AccountNo` | `*string` | JSON body | 未标注 | 签约支付宝账号 |
| `interface_info_list` | `InterfaceInfoList` | [`*[]InterfaceInfoList`](#model-interfaceinfolist) | JSON body | 未标注 | 接口信息列表，停车业务需要配置的接口列表，该值为JSON数据格式的LIST对象，现阶段只需要配置一个页面接口即可 。每次请将所有的接口配置信息都传入，未传的接口信息将会被置空。 |
| `merchant_logo` | `MerchantLogo` | `*string` | JSON body | 未标注 | 商户在停车平台首页露出的LOGO。 注意： *该图片为PNG格式，内容为BASE64的字符串，建议图片尺寸27px*27px，大小不要超过60K。 *若为空则停车平台首页将不露出商户LOGO。 |
| `merchant_name` | `MerchantName` | `*string` | JSON body | 未标注 | 商户简称，由开发者提供 |
| `merchant_service_phone` | `MerchantServicePhone` | `*string` | JSON body | 未标注 | 商户客服电话 |

<a id="api-mycarparkingenterinfosync"></a>

### MyCarParkingEnterInfoSync — 车辆驶入接口

- 官方接口：[`alipay.eco.mycar.parking.enterinfo.sync`](https://opendocs.alipay.com/open-v3/eb761e07_alipay.eco.mycar.parking.enterinfo.sync)
- 路由：`POST /v3/alipay/eco/mycar/parking/enterinfo/sync`
- 请求类型：[MyCarParkingEnterInfoSyncRequest](parking_request.go#L74)
- 响应类型：[MyCarParkingEnterInfoSyncResponse](parking_response.go#L33)

上传车辆驶入信息，上传信息通过该接口提交到支付宝，支付宝返回对应的信息

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `agreement_query` | `AgreementQuery` | `*bool` | JSON body | 未标注 | 是否启用车牌代扣状态查询功能，true为启用，false为停用 |
| `car_color` | `CarColor` | `*string` | JSON body | 未标注 | 车牌颜色，车牌颜色，枚举支持： *BLUE：蓝。 *GREEN：绿。 *YELLOW：黄。 *WHITE：白。 *BLACK：黑。 *LIMEGREEN：黄绿色。 |
| `car_number` | `CarNumber` | `*string` | JSON body | 未标注 | 车牌号 |
| `entrance_id` | `EntranceId` | `*string` | JSON body | 未标注 | 用于识别车辆进口，多出入口车场适用 |
| `free_minutes` | `FreeMinutes` | `*int64` | JSON body | 未标注 | 当前停车场的入场免费时长分钟数 |
| `in_time` | `InTime` | `*string` | JSON body | 未标注 | 车辆入场的时间，格式"YYYY-MM-DD HH:mm:ss"，24小时制，请保证服务器时间准确，入场时间不应晚于当前网络时间 |
| `is_encrypt_car_number` | `IsEncryptCarNumber` | `*bool` | JSON body | 未标注 | 是否加密，默认为false |
| `isv_url` | `IsvUrl` | `*string` | JSON body | 未标注 | 智能助理当前的跳转链接。为ISV/[商家小程序scheme地址](https://opendocs.alipay.com/support/01rb18)或H5页面地址。 |
| `out_serial_no` | `OutSerialNo` | `*string` | JSON body | 未标注 | 外部停车流水号(用于串通进场与出场信息) |
| `parking_id` | `ParkingId` | `*string` | JSON body | 未标注 | 支付宝停车平台ID，由支付宝定义的该停车场标识，同一个isv或商户范围内唯一。通过 [alipay.eco.mycar.parking.parkinglotinfo.create](https://opendocs.alipay.com/apis/api_19/alipay.eco.mycar.parking.parkinglotinfo.create)(录入停车场信息)接口获取。 |
| `space_number` | `SpaceNumber` | `*string` | JSON body | 未标注 | 用于标识车辆停放车位编号，路侧适用 |
| `store_id` | `StoreId` | `*string` | JSON body | 未标注 | 与parking_id对应的外部停车场编号，在门店中维护 |

<a id="api-mycarparkingexitinfosync"></a>

### MyCarParkingExitInfoSync — 车辆驶出接口

- 官方接口：[`alipay.eco.mycar.parking.exitinfo.sync`](https://opendocs.alipay.com/open-v3/4b64c632_alipay.eco.mycar.parking.exitinfo.sync)
- 路由：`POST /v3/alipay/eco/mycar/parking/exitinfo/sync`
- 请求类型：[MyCarParkingExitInfoSyncRequest](parking_request.go#L102)
- 响应类型：[MyCarParkingExitInfoSyncResponse](parking_response.go#L47)

上传车辆驶出信息，上传信息通过该接口提交到支付宝，支付宝返回对应的信息

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `actual_amount` | `ActualAmount` | `*string` | JSON body | 未标注 | 实际金额 |
| `car_color` | `CarColor` | `*string` | JSON body | 未标注 | 车牌颜色，枚举支持： *BLUE：蓝。 *GREEN：绿。 *YELLOW：黄。 *WHITE：白。 *BLACK：黑。 *LIMEGREEN：黄绿色。 |
| `car_number` | `CarNumber` | `*string` | JSON body | 未标注 | 车牌号 |
| `discount_amount` | `DiscountAmount` | `*string` | JSON body | 未标注 | 折扣金额 |
| `exit_id` | `ExitId` | `*string` | JSON body | 未标注 | 用于识别车辆出口，多出入口车场适用 |
| `is_encrypt_car_number` | `IsEncryptCarNumber` | `*bool` | JSON body | 未标注 | 是否加密，默认为false |
| `isv_url` | `IsvUrl` | `*string` | JSON body | 未标注 | 智能助理当前的跳转链接 |
| `order_amount` | `OrderAmount` | `*string` | JSON body | 未标注 | 订单总金额 |
| `out_serial_no` | `OutSerialNo` | `*string` | JSON body | 未标注 | 外部停车流水号(用于串通进场与出场信息) |
| `out_time` | `OutTime` | `*string` | JSON body | 未标注 | 车辆离场时间，格式"YYYY-MM-DD HH:mm:ss"，24小时制 |
| `parking_id` | `ParkingId` | `*string` | JSON body | 未标注 | 支付宝停车平台ID，由支付宝定义的该停车场标识，同一个ISV或商户范围内唯一。通过 [alipay.eco.mycar.parking.parkinglotinfo.create](https://opendocs.alipay.com/apis/api_19/alipay.eco.mycar.parking.parkinglotinfo.create)(录入停车场信息)接口获取。 |
| `serial_no` | `SerialNo` | `*string` | JSON body | 未标注 | 支付宝业务流水号，用于记录车辆从驶入到驶出的全流程。通过 [alipay.eco.mycar.parking.enterinfo.sync](https://opendocs.alipay.com/apis/api_19/alipay.eco.mycar.parking.enterinfo.sync)(车辆驶入接口)接口获取。 |
| `space_number` | `SpaceNumber` | `*string` | JSON body | 未标注 | 用于标识车辆停放车位编号，路侧适用 |
| `store_id` | `StoreId` | `*string` | JSON body | 未标注 | 与parking_id对应的外部停车场编号，在门店中维护 |

<a id="api-mycarparkingordersync"></a>

### MyCarParkingOrderSync — 订单同步接口

- 官方接口：[`alipay.eco.mycar.parking.order.sync`](https://opendocs.alipay.com/open-v3/e22974fc_alipay.eco.mycar.parking.order.sync)
- 路由：`POST /v3/alipay/eco/mycar/parking/order/sync`
- 请求类型：[MyCarParkingOrderSyncRequest](parking_request.go#L134)
- 响应类型：[MyCarParkingOrderSyncResponse](parking_response.go#L50)

商户通过接口调用，回传订单信息给停车平台

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `agent_pid` | `AgentPid` | `*string` | JSON body | 未标注 | 该笔停车交易需要返佣的对象实体PID，可能是商户，可能是ISV；只做下沉，用于离线表层面对账。 |
| `car_number` | `CarNumber` | `*string` | JSON body | 未标注 | 车牌 |
| `card_number` | `CardNumber` | `*string` | JSON body | 未标注 | 如果是停车卡缴费，则填入停车卡卡号，否则为'*' |
| `in_duration` | `InDuration` | `*string` | JSON body | 未标注 | 停车时长（以分为单位） |
| `in_time` | `InTime` | `*string` | JSON body | 未标注 | 入场时间，格式"YYYY-MM-DD HH:mm:ss"，24小时制 |
| `open_id` | `OpenId` | `*string` | JSON body | 未标注 | 应用用户ID，唯一标识 |
| `order_no` | `OrderNo` | `*string` | JSON body | 未标注 | 支付宝支付流水，系统唯一 |
| `order_status` | `OrderStatus` | `*string` | JSON body | 未标注 | 设备商订单状态，枚举支持： * 0：成功。 * 1：失败。 |
| `order_time` | `OrderTime` | `*string` | JSON body | 未标注 | 订单创建时间，格式"YYYY-MM-DD HH:mm:ss"，24小时制 |
| `out_order_no` | `OutOrderNo` | `*string` | JSON body | 未标注 | 设备商订单号，由ISV系统生成 |
| `out_parking_id` | `OutParkingId` | `*string` | JSON body | 未标注 | ISV停车场ID，由ISV定义的停车场标识，同一个isv或商户范围内唯一。需与 [alipay.eco.mycar.parking.parkinglotinfo.create](https://opendocs.alipay.com/apis/api_19/alipay.eco.mycar.parking.parkinglotinfo.create)(录入停车场信息)接口传入值一致。 |
| `out_time` | `OutTime` | `*string` | JSON body | 未标注 | 出场时间，格式"YYYY-MM-DD HH:mm:ss"，24小时制 |
| `parking_id` | `ParkingId` | `*string` | JSON body | 未标注 | 支付宝停车平台ID，由支付宝定义的该停车场标识，同一个isv或商户范围内唯一。通过 [alipay.eco.mycar.parking.parkinglotinfo.create](https://opendocs.alipay.com/apis/api_19/alipay.eco.mycar.parking.parkinglotinfo.create)(录入停车场信息)接口获取。 |
| `parking_name` | `ParkingName` | `*string` | JSON body | 未标注 | 停车场名称，需与[alipay.eco.mycar.parking.parkinglotinfo.create](https://opendocs.alipay.com/apis/api_19/alipay.eco.mycar.parking.parkinglotinfo.create)(录入停车场信息)接口传入值一致。 |
| `parking_record_id` | `ParkingRecordId` | `*string` | JSON body | 未标注 | 支付宝业务流水号，用于记录车辆从驶入到驶出的全流程。通过 alipay.eco.mycar.parking.enterinfo.sync(车辆驶入接口)接口获取。 |
| `pay_money` | `PayMoney` | `*string` | JSON body | 未标注 | 缴费金额，保留小数点后两位 |
| `pay_scene` | `PayScene` | `*string` | JSON body | 未标注 | 支付场景：INPARKINGLOT_PAY：场内在线缴费；ENTRANCE_EXIT_PAY：出入口缴费、VEH_DEVICE_PAY：车机缴费；会根据场景判断是否发放能量，当前只有场内支付场景会发能量，需要能量发放请与服务接入支持同学提前沟通。 |
| `pay_time` | `PayTime` | `*string` | JSON body | 未标注 | 缴费时间, 格式"YYYYMM-DD HH:mm:ss"，24小时制 |
| `pay_type` | `PayType` | `*string` | JSON body | 未标注 | 付款方式，枚举支持： *1：支付宝在线缴费。 |
| `smid` | `Smid` | `*string` | JSON body | 未标注 | 间连商户ID |
| `user_id` | `UserId` | `*string` | JSON body | 未标注 | 停车缴费用户 id，用户在支付宝的唯一标识，以 2088 开头的 16 位纯数字组成。 注意：ISV需保证用户 id的正确性，以免导致用户在停车平台查询不到相关的订单信息。 |

<a id="api-mycarparkingorderupdate"></a>

### MyCarParkingOrderUpdate — 订单更新接口

- 官方接口：[`alipay.eco.mycar.parking.order.update`](https://opendocs.alipay.com/open-v3/fe91ddf6_alipay.eco.mycar.parking.order.update)
- 路由：`POST /v3/alipay/eco/mycar/parking/order/update`
- 请求类型：[MyCarParkingOrderUpdateRequest](parking_request.go#L180)
- 响应类型：[MyCarParkingOrderUpdateResponse](parking_response.go#L53)

商户通过接口调用，回传订单状态给停车平台

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `open_id` | `OpenId` | `*string` | JSON body | 未标注 | 应用用户ID，唯一标识 |
| `order_no` | `OrderNo` | `*string` | JSON body | 未标注 | 支付宝支付流水号，系统唯一 |
| `order_status` | `OrderStatus` | `*string` | JSON body | 未标注 | 用户停车订单状态，枚举支持： *0：成功。 *1：失败。 |
| `user_id` | `UserId` | `*string` | JSON body | 未标注 | 停车缴费用户 id，用户在支付宝的唯一标识，以 2088 开头的 16 位纯数字组成。 注意：ISV需保证用户 id 的正确性，以免导致用户在停车平台查询不到相关的订单信息。 |

<a id="api-mycarparkingovertimecharginginfosync"></a>

### MyCarParkingOvertimeChargingInfoSync — 车辆停车超时重计费接口

- 官方接口：[`alipay.eco.mycar.parking.overtimecharginginfo.sync`](https://opendocs.alipay.com/open-v3/3e99daa8_alipay.eco.mycar.parking.overtimecharginginfo.sync)
- 路由：`POST /v3/alipay/eco/mycar/parking/overtimecharginginfo/sync`
- 请求类型：[MyCarParkingOvertimeChargingInfoSyncRequest](parking_request.go#L192)
- 响应类型：[MyCarParkingOvertimeChargingInfoSyncResponse](parking_response.go#L56)

上传车辆在停车场超时出场后的重新计费信息，将停车动作推进到计费状态，上传信息通过该接口提交到支付宝，支付宝返回对应的信息

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `car_number` | `CarNumber` | `*string` | JSON body | 未标注 | 车牌号 |
| `isv_url` | `IsvUrl` | `*string` | JSON body | 未标注 | 智能助理当前的跳转链接 |
| `parking_id` | `ParkingId` | `*string` | JSON body | 未标注 | 支付宝停车平台ID，由支付宝定义的该停车场标识，同一个isv或商户范围内唯一。通过 alipay.eco.mycar.parking.parkinglotinfo.create(录入停车场信息)接口获取。 |
| `serial_no` | `SerialNo` | `*string` | JSON body | 未标注 | 支付宝业务流水号，用于记录车辆从驶入到驶出的全流程 |

<a id="api-mycarparkingparkinglotbiztransfer"></a>

### MyCarParkingParkingLotBizTransfer — 车场业务归属配置

- 官方接口：[`alipay.eco.mycar.parking.parkinglotbiz.transfer`](https://opendocs.alipay.com/open-v3/cb924074_alipay.eco.mycar.parking.parkinglotbiz.transfer)
- 路由：`POST /v3/alipay/eco/mycar/parking/parkinglotbiz/transfer`
- 请求类型：[MyCarParkingParkingLotBizTransferRequest](parking_request.go#L204)
- 响应类型：[MyCarParkingParkingLotBizTransferResponse](parking_response.go#L59)

解决车场多isv接管问题，可以通过归属灵活控制权限。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `business_isv` | `BusinessIsv` | [`*[]BusinessItem`](#model-businessitem) | JSON body | 未标注 | 车场业务归属关系列表 |
| `parking_id` | `ParkingId` | `*string` | JSON body | 未标注 | 车场创建返回的parking_id字段 |

<a id="api-mycarparkingparkinglotinfocreate"></a>

### MyCarParkingParkingLotInfoCreate — 录入停车场信息

- 官方接口：[`alipay.eco.mycar.parking.parkinglotinfo.create`](https://opendocs.alipay.com/open-v3/fc05d317_alipay.eco.mycar.parking.parkinglotinfo.create)
- 路由：`POST /v3/alipay/eco/mycar/parking/parkinglotinfo/create`
- 请求类型：[MyCarParkingParkingLotInfoCreateRequest](parking_request.go#L212)
- 响应类型：[MyCarParkingParkingLotInfoCreateResponse](parking_response.go#L62)

录入停车场信息，内容信息通过该接口提交到支付宝，支付宝会生成支付宝这边停车场ID反馈给开发者，用户后续更新和上送车辆信息。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `agent_id` | `AgentId` | `*string` | JSON body | 未标注 | 服务商的支付宝账号ID（2088开头的16位纯数字，[获取方式](https://opendocs.alipay.com/common/02ncut)），由服务商提供给ISV |
| `business_isv` | `BusinessIsv` | [`*[]BusinessItem`](#model-businessitem) | JSON body | 未标注 | 车场业务归属列表 |
| `charge_sum_space` | `ChargeSumSpace` | `*string` | JSON body | 未标注 | 停车场充电车位数 单位:个 |
| `charging_rule` | `ChargingRule` | [`*[]ParkingLotChargingRuleInfo`](#model-parkinglotchargingruleinfo) | JSON body | 未标注 | 停车场详细计费规则 |
| `city_id` | `CityId` | `*string` | JSON body | 未标注 | 该参数废弃 |
| `contact_alipay` | `ContactAlipay` | `*string` | JSON body | 未标注 | 该参数废弃 |
| `contact_emali` | `ContactEmali` | `*string` | JSON body | 未标注 | 该参数废弃 |
| `contact_mobile` | `ContactMobile` | `*string` | JSON body | 未标注 | 该参数废弃 |
| `contact_name` | `ContactName` | `*string` | JSON body | 未标注 | 该参数废弃 |
| `contact_tel` | `ContactTel` | `*string` | JSON body | 未标注 | 该参数废弃 |
| `contact_weixin` | `ContactWeixin` | `*string` | JSON body | 未标注 | 该参数废弃 |
| `equipment_name` | `EquipmentName` | `*string` | JSON body | 未标注 | 该参数废弃 |
| `is_support_invoice` | `IsSupportInvoice` | `*string` | JSON body | 未标注 | 是否支持电子发票。枚举支持： 0：表示不支持。 1：表示支持。 |
| `isv_mobile` | `IsvMobile` | `*string` | JSON body | 未标注 | ISV电话，传入original_isv_pid、original_isv_appid时，此处为服务商电话 |
| `latitude` | `Latitude` | `*string` | JSON body | 未标注 | 纬度 单位:度 |
| `longitude` | `Longitude` | `*string` | JSON body | 未标注 | 经度 单位:度 |
| `mchnt_id` | `MchntId` | `*string` | JSON body | 未标注 | 收款方的支付宝账号ID（2088开头的16位纯数字，[获取方式](https://opendocs.alipay.com/common/02ncut)），由停车场收款的业主方提供给ISV，该字段暂用于机具和物料申领 |
| `original_isv_appid` | `OriginalIsvAppid` | `*string` | JSON body | 未标注 | ISV的APPID,服务商调用必传，由ISV提供给服务商 |
| `original_isv_mobile` | `OriginalIsvMobile` | `*string` | JSON body | 未标注 | ISV电话，此值与isv_mobile组合使用 |
| `original_isv_name` | `OriginalIsvName` | `*string` | JSON body | 未标注 | ISV的名称，服务商调用必传，由ISV提供给服务商 |
| `original_isv_pid` | `OriginalIsvPid` | `*string` | JSON body | 未标注 | ISV的PID，服务商调用必传,由ISV提供给服务商 |
| `out_parking_id` | `OutParkingId` | `*string` | JSON body | 未标注 | ISV停车场ID，由ISV 自定义，同一个isv或商户范围内唯一 |
| `parking_address` | `ParkingAddress` | `*string` | JSON body | 未标注 | 停车场地址 |
| `parking_end_time` | `ParkingEndTime` | `*string` | JSON body | 未标注 | 该参数废弃 |
| `parking_fee_description` | `ParkingFeeDescription` | `*string` | JSON body | 未标注 | 收费说明 |
| `parking_fee_description_img` | `ParkingFeeDescriptionImg` | `*string` | JSON body | 未标注 | 商户在停车平台露出的停车价格图片；注意：该图片为PNG格式内容为BASE64的字符串，大小不要超过1MB |
| `parking_lot_type` | `ParkingLotType` | `*string` | JSON body | 未标注 | 停车场类型，COMMUNITY为居民小区、BUSINESS_AREA为商圈停车场（购物中心商业广场商场等）、ROADSIDE为路侧停车、PARK_SCENIC为公园景点（景点乐园公园老街古镇等）、OFFICE_BUILDING为商务楼宇（酒店写字楼商务楼园区等）、OTHER为其他、TRANSPORTATION为交通枢纽（机场火车站汽车站码头港口等）、PUBLIC_FACILITIES为市政设施（体育场博物图书馆医院学校等）、TERRITORY独立园区（办公工业物流园区等）、BUSINESS_PLACE经营场所（4S店、门市餐饮等 |
| `parking_mobile` | `ParkingMobile` | `*string` | JSON body | 未标注 | 停车场客服电话 |
| `parking_name` | `ParkingName` | `*string` | JSON body | 未标注 | 停车场名称，由ISV定义，尽量与高德地图上的一致。 |
| `parking_number` | `ParkingNumber` | `*string` | JSON body | 未标注 | 该参数废弃 |
| `parking_poiid` | `ParkingPoiid` | `*string` | JSON body | 未标注 | 如何获取parking_poiid（高德地图唯一标标识）参考文档 [https://opendocs.alipay.com/support/01rghx](https://opendocs.alipay.com/support/01rghx) ；若无法成功获取高德POI时，也可选用经纬度的方式注册。 |
| `parking_start_time` | `ParkingStartTime` | `*string` | JSON body | 未标注 | 该参数废弃 |
| `parking_type` | `ParkingType` | `*string` | JSON body | 未标注 | 该参数废弃 |
| `pay_type` | `PayType` | `*string` | JSON body | 未标注 | 支付方式。枚举支持： 1：表示支付宝在线缴费。 2：表示支付宝代扣缴费。 3：表示当面付。 4: 现金 说明：如支持多种方式以 ',' 进行分隔。 |
| `payment_mode` | `PaymentMode` | `*string` | JSON body | 未标注 | 该参数废弃 |
| `serivce_url` | `SerivceUrl` | `*string` | JSON body | 未标注 | 提供用户查看停车状态、缴费功能 |
| `service_list` | `ServiceList` | [`*[]ParkingLotServiceInfo`](#model-parkinglotserviceinfo) | JSON body | 未标注 | 停车场服务信息列表 |
| `shopingmall_id` | `ShopingmallId` | `*string` | JSON body | 未标注 | 商圈id |
| `sum_space` | `SumSpace` | `*string` | JSON body | 未标注 | 停车场车位数 单位:个 |
| `support_after_pay` | `SupportAfterPay` | `*string` | JSON body | 未标注 | 是否支持先离后付。枚举支持： Y：支持。 N：不支持。 说明：默认为空不支持，此参数适用于签约信用代扣的商户。 |
| `time_out` | `TimeOut` | `*string` | JSON body | 未标注 | 用户支付未离场的超时时间(以分钟为单位) |

<a id="api-mycarparkingparkinglotinfoquery"></a>

### MyCarParkingParkingLotInfoQuery — 停车场信息查询

- 官方接口：[`alipay.eco.mycar.parking.parkinglotinfo.query`](https://opendocs.alipay.com/open-v3/4772d6e6_alipay.eco.mycar.parking.parkinglotinfo.query)
- 路由：`GET /v3/alipay/eco/mycar/parking/parkinglotinfo/query`
- 请求类型：[MyCarParkingParkingLotInfoQueryRequest](parking_request.go#L426)
- 响应类型：[MyCarParkingParkingLotInfoQueryResponse](parking_response.go#L72)

停车场信息查询，通过停车场id或者ISV停车场ID查询停车场信息。只能查询正在调用接口商户自己创建的停车场，限制在归属PID层面上。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `parking_id` | `ParkingId` | `*string` | query | 否 | 支付宝停车平台ID，由支付宝定义的该停车场标识，同一个ISV或商户范围内唯一。通过 [alipay.eco.mycar.parking.parkinglotinfo.create](https://opendocs.alipay.com/apis/api_19/alipay.eco.mycar.parking.parkinglotinfo.create)(录入停车场信息)接口获取。 注意：parking_id和out_parking_id不能同时为空。 |
| `out_parking_id` | `OutParkingId` | `*string` | query | 否 | ISV停车场ID，由ISV定义的停车场标识，同一个ISV或商户范围内唯一。需与 [alipay.eco.mycar.parking.parkinglotinfo.create](https://opendocs.alipay.com/apis/api_19/alipay.eco.mycar.parking.parkinglotinfo.create)(录入停车场信息)接口传入值一致。 注意：parking_id和out_parking_id不能同时为空。 |

<a id="api-mycarparkingparkinglotinfoupdate"></a>

### MyCarParkingParkingLotInfoUpdate — 修改停车场信息

- 官方接口：[`alipay.eco.mycar.parking.parkinglotinfo.update`](https://opendocs.alipay.com/open-v3/223553d2_alipay.eco.mycar.parking.parkinglotinfo.update)
- 路由：`POST /v3/alipay/eco/mycar/parking/parkinglotinfo/update`
- 请求类型：[MyCarParkingParkingLotInfoUpdateRequest](parking_request.go#L298)
- 响应类型：[MyCarParkingParkingLotInfoUpdateResponse](parking_response.go#L128)

录入停车场信息，内容信息通过该接口提交到支付宝，支付宝会生成支付宝这边停车场ID反馈给开发者，用于后续更新和上送车辆信息，停车场名称不允许修改。

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `agent_id` | `AgentId` | `*string` | JSON body | 未标注 | 服务商ID（2088开头的16位纯数字），由服务商提供给ISV |
| `business_isv` | `BusinessIsv` | [`*[]BusinessItem`](#model-businessitem) | JSON body | 未标注 | 车场业务归属列表 |
| `charge_sum_space` | `ChargeSumSpace` | `*string` | JSON body | 未标注 | 停车场充电车位数 单位:个 |
| `charging_rule` | `ChargingRule` | [`*[]ParkingLotChargingRuleInfo`](#model-parkinglotchargingruleinfo) | JSON body | 未标注 | 停车场详细计费规则详细说明 |
| `city_id` | `CityId` | `*string` | JSON body | 未标注 | 该参数废弃 |
| `contact_alipay` | `ContactAlipay` | `*string` | JSON body | 未标注 | 该参数废弃 |
| `contact_email` | `ContactEmail` | `*string` | JSON body | 未标注 | 该参数废弃 |
| `contact_mobile` | `ContactMobile` | `*string` | JSON body | 未标注 | 该参数废弃 |
| `contact_name` | `ContactName` | `*string` | JSON body | 未标注 | 该参数废弃 |
| `contact_tel` | `ContactTel` | `*string` | JSON body | 未标注 | 该参数废弃 |
| `contact_weixin` | `ContactWeixin` | `*string` | JSON body | 未标注 | 该参数废弃 |
| `equipment_name` | `EquipmentName` | `*string` | JSON body | 未标注 | 该参数废弃 |
| `is_support_invoice` | `IsSupportInvoice` | `*string` | JSON body | 未标注 | 是否支持电子发票。枚举支持： *0：表示不支持。 *1：表示支持。 |
| `isv_mobile` | `IsvMobile` | `*string` | JSON body | 未标注 | ISV电话，传入original_isv_pid、original_isv_appid时，此处为服务商电话 |
| `latitude` | `Latitude` | `*string` | JSON body | 未标注 | 纬度，单位:度 |
| `longitude` | `Longitude` | `*string` | JSON body | 未标注 | 经度 单位:度 |
| `mchnt_id` | `MchntId` | `*string` | JSON body | 未标注 | 收款方ID（2088开头的16位纯数字），由停车场收款的业主方提供给ISV，该字段暂用于机具和物料申领，更新时不影响交易接口中的收款方值。 |
| `original_isv_appid` | `OriginalIsvAppid` | `*string` | JSON body | 未标注 | ISV的APPID,服务商调用必传，由ISV提供给服务商 |
| `original_isv_mobile` | `OriginalIsvMobile` | `*string` | JSON body | 未标注 | ISV电话，此值与isv_mobile组合使用 |
| `original_isv_name` | `OriginalIsvName` | `*string` | JSON body | 未标注 | ISV的名称，服务商调用必传，由ISV提供给服务商 |
| `original_isv_pid` | `OriginalIsvPid` | `*string` | JSON body | 未标注 | ISV的PID，服务商调用必传,由ISV提供给服务商 |
| `out_parking_id` | `OutParkingId` | `*string` | JSON body | 未标注 | 该参数废弃 |
| `parking_address` | `ParkingAddress` | `*string` | JSON body | 未标注 | 停车场地址 |
| `parking_end_time` | `ParkingEndTime` | `*string` | JSON body | 未标注 | 该参数废弃 |
| `parking_fee_description` | `ParkingFeeDescription` | `*string` | JSON body | 未标注 | 收费说明 |
| `parking_fee_description_img` | `ParkingFeeDescriptionImg` | `*string` | JSON body | 未标注 | 商户在停车平台露出的停车价格图片；注意：该图片为PNG格式内容为BASE64的字符串，大小不要超过1MB |
| `parking_id` | `ParkingId` | `*string` | JSON body | 未标注 | 支付宝返回停车场id，系统唯一 |
| `parking_lot_type` | `ParkingLotType` | `*string` | JSON body | 未标注 | 停车场类型，COMMUNITY为居民小区、BUSINESS_AREA为商圈停车场（购物中心商业广场商场等）、ROADSIDE为路侧停车、PARK_SCENIC为公园景点（景点乐园公园老街古镇等）、OFFICE_BUILDING为商务楼宇（酒店写字楼商务楼园区等）、OTHER为其他、TRANSPORTATION为交通枢纽（机场火车站汽车站码头港口等）、PUBLIC_FACILITIES为市政设施（体育场博物图书馆医院学校等）、TERRITORY独立园区（办公工业物流园区等）、BUSINESS_PLACE经营场所（4S店、门市餐饮等 |
| `parking_mobile` | `ParkingMobile` | `*string` | JSON body | 未标注 | 停车场客服电话 |
| `parking_name` | `ParkingName` | `*string` | JSON body | 未标注 | 停车场名称，由ISV定义，尽量与高德地图上的一致 |
| `parking_number` | `ParkingNumber` | `*string` | JSON body | 未标注 | 该参数废弃 |
| `parking_poiid` | `ParkingPoiid` | `*string` | JSON body | 未标注 | 如何获取parking_poiid（高德地图唯一标标识）参考文档 [https://opendocs.alipay.com/support/01rghx](https://opendocs.alipay.com/support/01rghx) ；若无法成功获取高德POI时，也可选用经纬度的方式注册。 |
| `parking_start_time` | `ParkingStartTime` | `*string` | JSON body | 未标注 | 该参数废弃 |
| `parking_type` | `ParkingType` | `*string` | JSON body | 未标注 | 该参数废弃 |
| `pay_type` | `PayType` | `*string` | JSON body | 未标注 | 支付方式。枚举支持： *1：表示支付宝在线缴费。 *2：表示支付宝代扣缴费。 *3：表示当面付。 *4：表示现金。 说明：如支持多种方式以 ',' 进行分隔。 |
| `payment_mode` | `PaymentMode` | `*string` | JSON body | 未标注 | 该参数废弃 |
| `remain_space` | `RemainSpace` | `*string` | JSON body | 未标注 | 剩余车位数 单位：个 |
| `serivce_url` | `SerivceUrl` | `*string` | JSON body | 未标注 | 提供用户查看停车状态、缴费功能 |
| `service_list` | `ServiceList` | [`*[]ParkingLotServiceInfo`](#model-parkinglotserviceinfo) | JSON body | 未标注 | 停车场场内服务信息 |
| `shopingmall_id` | `ShopingmallId` | `*string` | JSON body | 未标注 | 商圈id |
| `sum_space` | `SumSpace` | `*string` | JSON body | 未标注 | 停车场车位数 ，单位:个 |
| `support_after_pay` | `SupportAfterPay` | `*string` | JSON body | 未标注 | 是否支持先离后付。枚举支持： *Y：支持。 *N：不支持。 说明：默认为空不支持，此参数适用于签约信用代扣的商户。 |
| `time_out` | `TimeOut` | `*string` | JSON body | 未标注 | 用户支付未离场的超时时间(以分钟为单位) |

<a id="api-mycarparkingpaymentinfosync"></a>

### MyCarParkingPaymentInfoSync — 车辆停车缴费接口

- 官方接口：[`alipay.eco.mycar.parking.paymentinfo.sync`](https://opendocs.alipay.com/open-v3/def30621_alipay.eco.mycar.parking.paymentinfo.sync)
- 路由：`POST /v3/alipay/eco/mycar/parking/paymentinfo/sync`
- 请求类型：[MyCarParkingPaymentInfoSyncRequest](parking_request.go#L388)
- 响应类型：[MyCarParkingPaymentInfoSyncResponse](parking_response.go#L136)

上传车辆在停车场的缴费信息，将停车动作推进到计费状态，上传信息通过该接口提交到支付宝，支付宝返回对应的信息

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `car_number` | `CarNumber` | `*string` | JSON body | 未标注 | 车牌号 |
| `isv_url` | `IsvUrl` | `*string` | JSON body | 未标注 | 停车消息透出后，消息详情的落地承接页链接，可配置ISV维度的统一值 |
| `out_serial_no` | `OutSerialNo` | `*string` | JSON body | 未标注 | 外部业务停车流水号(用于串通进场与出场信息)。如果 out_serial_no与serial_no都传递，则一起校验。 ● serial_no与out_serial_no二选一必填。 ● 入参有serial_no和out_serial_no，则优先根据serial_no和out_serial_no一起匹配进出场事件 ● 入参仅有serial_no，则优先根据serial_no匹配进出场事件 ● 入参仅有out_serial_no，则根据out_serial_no匹配进出场事件 |
| `parking_id` | `ParkingId` | `*string` | JSON body | 未标注 | 支付宝停车平台ID，由支付宝定义的该停车场标识，同一个isv或商户范围内唯一。通过 alipay.eco.mycar.parking.parkinglotinfo.create(录入停车场信息)接口获取。 |
| `payment_free_minutes` | `PaymentFreeMinutes` | `*int64` | JSON body | 未标注 | 停车场内缴费后，允许出场的免费时长分钟数 |
| `payment_time` | `PaymentTime` | `*string` | JSON body | 未标注 | 场内缴费时间，格式"YYYY-MM-DD HH:mm:ss"，24小时制，请保证服务器时间准确，场内缴费不应晚于当前网络时间 |
| `serial_no` | `SerialNo` | `*string` | JSON body | 未标注 | 支付宝业务流水号(用于串通进场与出场信息)。可通过 alipay.eco.mycar.parking.enterinfo.sync (车辆驶入接口)接口获取。 ● serial_no与out_serial_no二选一必填。 ● 入参有serial_no和out_serial_no，则优先根据serial_no和out_serial_no一起匹配进出场事件 ● 入参仅有serial_no，则优先根据serial_no匹配进出场事件 ● 入参仅有out_serial_no，则根据out_serial_no匹配进出场事件 |

<a id="api-mycarparkingspaceinfosync"></a>

### MyCarParkingSpaceInfoSync — 停车场车位信息同步

- 官方接口：[`alipay.eco.mycar.parking.spaceinfo.sync`](https://opendocs.alipay.com/open-v3/cab695b7_alipay.eco.mycar.parking.spaceinfo.sync)
- 路由：`POST /v3/alipay/eco/mycar/parking/spaceinfo/sync`
- 请求类型：[MyCarParkingSpaceInfoSyncRequest](parking_request.go#L406)
- 响应类型：[MyCarParkingSpaceInfoSyncResponse](parking_response.go#L139)

停车场位置信息同步

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `free_charging_pile` | `FreeChargingPile` | `*int64` | JSON body | 未标注 | 空闲充电桩车位数 |
| `free_parking_space` | `FreeParkingSpace` | `*int64` | JSON body | 未标注 | 车场空闲可用的车位数 |
| `parking_id` | `ParkingId` | `*string` | JSON body | 未标注 | 停车场ID |
| `parking_space_status` | `ParkingSpaceStatus` | `*string` | JSON body | 未标注 | 停车场实时状态: 0 拥堵,1 正常，2 空闲 |

<a id="api-mycarparkingvehiclequery"></a>

### MyCarParkingVehicleQuery — 车牌查询接口

- 官方接口：[`alipay.eco.mycar.parking.vehicle.query`](https://opendocs.alipay.com/open-v3/c04327bf_alipay.eco.mycar.parking.vehicle.query)
- 路由：`GET /v3/alipay/eco/mycar/parking/vehicle/query`
- 请求类型：[MyCarParkingVehicleQueryRequest](parking_request.go#L434)
- 响应类型：[MyCarParkingVehicleQueryResponse](parking_response.go#L145)

商户通过接口调用，获取用户车牌信息

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `auth_token` | `AuthToken` | `*string` | query | 否 | 用户授权令牌 |
| `car_id` | `CarId` | `*string` | query | 否 | 支付宝用户车辆ID，系统唯一。（该参数会在停车平台用户点击查询缴费，跳转到ISV停车缴费查询页面时，从请求中传递） |

## 嵌套请求参数

<a id="model-businessitem"></a>

### BusinessItem

类型定义：[BusinessItem](parking_model.go#L4)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `business_appid` | `BusinessAppid` | `*string` | 未标注 | 业务归属appid |
| `business_payee_id` | `BusinessPayeeId` | `*string` | 未标注 | 业务收款pid，business_type为AGREEMENT_PAY时，此参数必传 |
| `business_pid` | `BusinessPid` | `*string` | 未标注 | 业务归属pid |
| `business_type` | `BusinessType` | `*string` | 未标注 | 业务类型ONLINE_PAY(在线缴费) 、AGREEMENT_PAY(无感停车) DEVICE_ONLINE_PAY(车机在线缴费) |

<a id="model-interfaceinfolist"></a>

### InterfaceInfoList

类型定义：[InterfaceInfoList](parking_model.go#L16)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `interface_name` | `InterfaceName` | `*string` | 未标注 | H5传入参数固定值：alipay.eco.mycar.parking.userpage.query; 小程序传入参数固定值：alipay.eco.mycar.parking.isv.homepage; 协议状态变更通知传入参数固定值： alipay.eco.mycar.parking.agreement.notify; |
| `interface_type` | `InterfaceType` | `*string` | 未标注 | H5、小程序传入参数固定值:interface_page； 通知接口传入参数固定值:interface_service; |
| `interface_url` | `InterfaceUrl` | `*string` | 未标注 | H5 SPI接口的调用地址url，协议必须为https，对整个url字符串必须进行UrlEncode编码。编码为UTF-8; 小程序SPI接口的调用地址ur格式：小程序&#124;小程序缴费页面地址; 协议状态变更通知SPI接口的调用地址url，同H5格式; |

<a id="model-parkinglotchargingruleinfo"></a>

### ParkingLotChargingRuleInfo

类型定义：[ParkingLotChargingRuleInfo](parking_model.go#L26)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `charging_period` | `ChargingPeriod` | `*string` | 未标注 | 停车收费时间段，输入24小时制的小时时间段格式: 开始时间-结束时间, 结束时间必须大于开始时间 时间取值范围：0到24 |
| `charging_strategy` | `ChargingStrategy` | `*string` | 未标注 | 计费策略； UP：向上取整 DOMN：向下取整 |
| `charging_unit_fee` | `ChargingUnitFee` | `*string` | 未标注 | 跳费金额 单位元 |
| `charging_unit_minutes` | `ChargingUnitMinutes` | `*string` | 未标注 | 跳费间隔分钟数 |
| `enable_rollover_charge_period` | `EnableRolloverChargePeriod` | `*bool` | 未标注 | 超时是否顺延到下个计费区间 |
| `first_charging_span_minutes` | `FirstChargingSpanMinutes` | `*string` | 未标注 | 首段计费周期分钟数 |
| `first_charging_unit_fee` | `FirstChargingUnitFee` | `*string` | 未标注 | 首段计费周期跳费金额 单位: 元 |
| `first_charging_unit_minutes` | `FirstChargingUnitMinutes` | `*string` | 未标注 | 首段计费周期跳费间隔 单位分钟 |
| `free_enter_minutes` | `FreeEnterMinutes` | `*string` | 未标注 | 入场免费时间 单位:分钟 |
| `free_exit_minutes` | `FreeExitMinutes` | `*string` | 未标注 | 离场免费时间 单位:分钟 |
| `max_fee_per_day` | `MaxFeePerDay` | `*string` | 未标注 | 日封顶金额 单位: 元 |

<a id="model-parkinglotserviceinfo"></a>

### ParkingLotServiceInfo

类型定义：[ParkingLotServiceInfo](parking_model.go#L52)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `service_name` | `ServiceName` | `*string` | 未标注 | 在线缴费服务名称 |
| `service_type` | `ServiceType` | `*string` | 未标注 | 停车场可提供的服务类型 发票服务：INVOICE_SERVICE 优惠套餐: DISCOUNT_PACKAGE 预约停车: RESERVED_PARKING |
| `service_url` | `ServiceUrl` | `*string` | 未标注 | 商户提供服务对应的跳转链接 |
