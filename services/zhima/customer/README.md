# zhima/customer 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/zhima/customer`。本目录提供 6 个支付宝开放平台 V3 REST 操作。

## 使用方式

先按[项目接入说明](../../../README.md)创建 `*core.Client`，再通过 `customer.NewService(client)` 创建 Service；所有网络调用以 `context.Context` 为首个参数，支持 `opts ...core.CallOption`。响应为 `*core.Response[T]`，业务数据位于 `resp.Data`。

请求、响应与嵌套类型的链接指向本目录源码；可参考本目录的 [API 示例](api_customer_example_test.go)。

示例文件中的 `TestSandboxService_*` 可使用沙箱平台证书执行真实调用。默认 `go test` 保持离线。

## 接口目录

| SDK 方法 | 官方接口 | 用途 | HTTP |
| --- | --- | --- | --- |
| [`JobWorthAuthenticationQuery`](#api-jobworthauthenticationquery) | [`zhima.customer.jobworth.authentication.query`](https://opendocs.alipay.com/open-v3/d25db704_zhima.customer.jobworth.authentication.query) | 职得身份认证查询接口 | `GET` |
| [`JobWorthCloudResumeQuery`](#api-jobworthcloudresumequery) | [`zhima.customer.jobworth.cloudresume.query`](https://opendocs.alipay.com/apis/05wukc) | 智能简历基本数据查询 | `GET` |
| [`JobWorthInfoQuery`](#api-jobworthinfoquery) | [`zhima.customer.jobworth.info.query`](https://opendocs.alipay.com/apis/05ttl7) | 职得工作证信息查询 | `GET` |
| [`JobWorthJobDataAdd`](#api-jobworthjobdataadd) | [`zhima.customer.jobworth.jobdata.add`](https://opendocs.alipay.com/apis/05wyeb) | 工作证职业数据回流接口 | `POST` |
| [`JobWorthPictureUpload`](#api-jobworthpictureupload) | [`zhima.customer.jobworth.picture.upload`](https://opendocs.alipay.com/apis/05wulq) | 职得工作证图片上传 | `POST` |
| [`JobWorthPositionAdd`](#api-jobworthpositionadd) | [`zhima.customer.jobworth.position.add`](https://opendocs.alipay.com/apis/05xa7n) | 职得工作证岗位收集 | `POST` |

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。

## 接口说明与请求参数

<a id="api-jobworthauthenticationquery"></a>

### JobWorthAuthenticationQuery — 职得身份认证查询接口

- 官方接口：[`zhima.customer.jobworth.authentication.query`](https://opendocs.alipay.com/open-v3/d25db704_zhima.customer.jobworth.authentication.query)
- 路由：`GET /v3/zhima/customer/jobworth/authentication`
- 请求类型：[JobWorthAuthenticationQueryRequest](customer_request.go#L37)
- 响应类型：[JobWorthAuthenticationQueryResponse](customer_response.go#L5)

根据一次性token查询用户身份认证结果

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `conn_key` | `ConnKey` | `*string` | query | 否 | 调用第一步创建接口的out_agreement_no，由外部传入 |
| `once_token` | `OnceToken` | `*string` | query | 否 | 查询身份认证结果的凭证，在认证结果到达终态后只允许查询成功一次，若需要再次查询，则需要调用第一个接口再拉起受理台 |
| `query_type` | `QueryType` | `*string` | query | 否 | exact:精确查询,dim:模糊查询 |
| `identity_type` | `IdentityType` | `*string` | query | 否 | 04:教师，其它身份类型联系芝麻技术支持确认 |
| `service_id` | `ServiceId` | `*string` | query | 否 | 信用服务id，商家接入产品功能包时有运营同学分配 |
| `user_id` | `UserId` | `*string` | query | 否 | 蚂蚁统一会员ID 和身份证任选一个传入 |
| `open_id` | `OpenId` | `*string` | query | 否 | 蚂蚁统一会员ID 和身份证任选一个传入 |
| `cert_no` | `CertNo` | `*string` | query | 否 | 证件号 |
| `cert_type` | `CertType` | `*string` | query | 否 | "身份证 : 0"，"护照: 1"，"台湾往来通行证：4","港澳往来通行证：3", "回乡证: 2"；"港澳证件：5" |

<a id="api-jobworthcloudresumequery"></a>

### JobWorthCloudResumeQuery — 智能简历基本数据查询

- 官方接口：[`zhima.customer.jobworth.cloudresume.query`](https://opendocs.alipay.com/apis/05wukc)
- 路由：`GET /v3/zhima/customer/jobworth/cloudresume/query`
- 请求类型：[JobWorthCloudResumeQueryRequest](customer_request.go#L59)
- 响应类型：[JobWorthCloudResumeQueryResponse](customer_response.go#L21)

用于查询工作证基本信息

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `conn_key` | `ConnKey` | `*string` | query | 否 | 长效的认证令牌：zhima.credit.payafteruse.creditagreement.sign 的out_agreement_no |
| `once_token` | `OnceToken` | `*string` | query | 否 | 交换云简历的单次token：zhima.credit.payafteruse.creditagreement.sign 的 extra_param入参onceToken（授权后5分钟内，单次有效 |

<a id="api-jobworthinfoquery"></a>

### JobWorthInfoQuery — 职得工作证信息查询

- 官方接口：[`zhima.customer.jobworth.info.query`](https://opendocs.alipay.com/apis/05ttl7)
- 路由：`GET /v3/zhima/customer/jobworth/info/query`
- 请求类型：[JobWorthInfoQueryRequest](customer_request.go#L67)
- 响应类型：[JobWorthInfoQueryResponse](customer_response.go#L57)

查询用户职得工作证信息
查询方式3选1，都设置时uid优先
1、支付宝uid查询：user_id
2、身份证+姓名查询：cert_type+cert_no
3、智能简历场景：conn_key

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `conn_key` | `ConnKey` | `*string` | query | 否 | 外部订单号，zhima.credit.payafteruse.creditagreement.sign的入参 out_agreement_no智能简历场景支持只传该值 |
| `user_name` | `UserName` | `*string` | query | 否 | 姓名 |
| `cert_no` | `CertNo` | `*string` | query | 否 | 证件号 ，根据cert_type类型设置对应证件号码，选择身份证校验时必传 |
| `user_id` | `UserId` | `*string` | query | 否 | 蚂蚁统一会员ID ，为2088开头的唯一标识 选择支付宝uid检验时必传，支付宝uid或身份证+身份证类型 查询2选1，都设置时uid优先 |
| `open_id` | `OpenId` | `*string` | query | 否 | 蚂蚁统一会员ID ，为2088开头的唯一标识 选择支付宝uid检验时必传，支付宝uid或身份证+身份证类型 查询2选1，都设置时uid优先 |
| `back_url` | `BackUrl` | `*string` | query | 否 | 若用户没有开通工作证或者芝麻，成功开通后回跳的链接 ，支持http或schema地址 |
| `cert_type` | `CertType` | `*string` | query | 否 | 证件类型 ，枚举值： 0:身份证1:护照2:台湾同胞回乡证3:港澳居民往来内地通行证4:台湾居民通行证5:港澳居民通行证 选择身份证校验时必传，支付宝uid或身份证+身份证类型 查询2选1，都设置时uid优先 |
| `ka_visitor_id` | `KaVisitorId` | `*string` | query | 否 | 商户内部的唯一id如果是需要返回h5页面必填 |
| `visitor_type` | `VisitorType` | `*string` | query | 否 | 访问类型 C = company or P = person 如果是需要返回h5页面必填 |
| `visitor_name` | `VisitorName` | `*string` | query | 否 | 商户侧配置访客名称 如果是需要返回h5页面必填 |
| `company_certificate` | `CompanyCertificate` | `*string` | query | 否 | 公司社会统一代码h5页面访客展示使用 |
| `visitor_url` | `VisitorUrl` | `*string` | query | 否 | 上传图片接口返回的id |
| `has_button` | `HasButton` | `*bool` | query | 否 | 在h5页面是否展示底部跳转按钮 |
| `has_number` | `HasNumber` | `*bool` | query | 否 | true or false 是否展示 |
| `has_qr_code` | `HasQrCode` | `*bool` | query | 否 | 是否展示二维码 |
| `has_html` | `HasHtml` | `*bool` | query | 否 | 是否需要返回h5页面 |
| `job_id` | `JobId` | `*string` | query | 否 | 对外使用的jobid |
| `industry_id` | `IndustryId` | `*string` | query | 否 | 行业描述ID |
| `cloud_resume_scene` | `CloudResumeScene` | `*string` | query | 否 | 子场景 |

<a id="api-jobworthjobdataadd"></a>

### JobWorthJobDataAdd — 工作证职业数据回流接口

- 官方接口：[`zhima.customer.jobworth.jobdata.add`](https://opendocs.alipay.com/apis/05wyeb)
- 路由：`POST /v3/zhima/customer/jobworth/jobdata/add`
- 请求类型：[JobWorthJobDataAddRequest](customer_request.go#L5)
- 响应类型：[JobWorthJobDataAddResponse](customer_response.go#L79)

合作的三方伙伴通过该接口，将数据回传到工作证

请求约束：JobWorthJobDataAddRequest is the official V3 ZhimaCustomerJobworthJobdataAddModel schema.

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `cert_no` | `CertNo` | `*string` | JSON body | 未标注 | 身份证证件号 |
| `cert_type` | `CertType` | `*string` | JSON body | 未标注 | 0:身份证1:护照2:台湾同胞回乡证3:港澳居民往来内地通行证4:台湾居民通行证5:港澳居民通行证 |
| `conn_key` | `ConnKey` | `*string` | JSON body | 未标注 | 关联key，受理台返回 |
| `job_data_list` | `JobDataList` | [`*[]JobWorthJobData`](#model-jobworthjobdata) | JSON body | 未标注 | 回传信息 |
| `open_id` | `OpenId` | `*string` | JSON body | 未标注 | 支付宝uid |
| `user_id` | `UserId` | `*string` | JSON body | 未标注 | 支付宝uid |
| `user_name` | `UserName` | `*string` | JSON body | 未标注 | 用户中文名 |

<a id="api-jobworthpictureupload"></a>

### JobWorthPictureUpload — 职得工作证图片上传

- 官方接口：[`zhima.customer.jobworth.picture.upload`](https://opendocs.alipay.com/apis/05wulq)
- 路由：`POST /v3/zhima/customer/jobworth/picture/upload`
- 请求类型：[JobWorthPictureUploadRequest](customer_request.go#L23)
- 响应类型：[JobWorthPictureUploadResponse](customer_response.go#L91)

工作证图片上传通用接口

请求约束：JobWorthPictureUploadRequest is the official V3 ZhimaCustomerJobworthPictureUploadModel schema.

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `content` | `Content` | `*string` | JSON body | 未标注 | 图片对应的base64字符串，支持jpg和png，1M之内 |
| `pic_type` | `PicType` | `*string` | JSON body | 未标注 | 只能传入指定类型，具体类型看文档 |

<a id="api-jobworthpositionadd"></a>

### JobWorthPositionAdd — 职得工作证岗位收集

- 官方接口：[`zhima.customer.jobworth.position.add`](https://opendocs.alipay.com/apis/05xa7n)
- 路由：`POST /v3/zhima/customer/jobworth/position/add`
- 请求类型：[JobWorthPositionAddRequest](customer_request.go#L31)
- 响应类型：[JobWorthPositionAddResponse](customer_response.go#L101)

保存商户传入岗位信息，当接口调用成功，业务问题时返回subcode

请求约束：JobWorthPositionAddRequest is the official V3 ZhimaCustomerJobworthPositionAddModel schema.

| 协议字段 | Go 字段 | Go 类型 | 位置 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- | --- |
| `jobworth_position_info` | `JobworthPositionInfo` | [`*JobWorthPositionInfo`](#model-jobworthpositioninfo) | JSON body | 未标注 | 见字段类型及官方接口文档。 |

## 嵌套请求参数

以下仅展开本目录请求中实际使用的对象；字段必填规则沿用官方快照，“未标注”须结合业务说明确认。响应模型请沿各接口的响应类型链接查看。

<a id="model-jobworthjobdata"></a>

### JobWorthJobData

类型定义：[JobWorthJobData](customer_model.go#L103)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `certificate_grant_institution` | `CertificateGrantInstitution` | `*string` | 未标注 | 证书发证机构 |
| `certificate_id` | `CertificateId` | `*string` | 未标注 | 证书编号 |
| `certificate_level` | `CertificateLevel` | `*string` | 未标注 | 证书等级：中文 |
| `certificate_name` | `CertificateName` | `*string` | 未标注 | 证书名称 |
| `certificate_pic_id` | `CertificatePicId` | `*string` | 未标注 | 工作证图片上传接口返回的id |
| `company_name` | `CompanyName` | `*string` | 未标注 | 公司名称 |
| `degree` | `Degree` | `*string` | 未标注 | 学历，可以用以下字段HIGHSCHOOL_AND_BELOW（高中及以下），POLYTECHNIC（中专），COLLEGE（大专），BACHELOR（本科），MASTER（硕士），DOCTOR_AND_ABOVE（博士及以上） |
| `delivery_position_id` | `DeliveryPositionId` | `*string` | 未标注 | 商户侧投递岗位id |
| `delivery_time` | `DeliveryTime` | `*int64` | 未标注 | 投递时间戳 |
| `education_status` | `EducationStatus` | `*string` | 未标注 | 教育状态只有两种，GRADUATE(毕业)，STUDY（未毕业） |
| `entry_no` | `EntryNo` | `*string` | 未标注 | 报名流水号 |
| `exam_score` | `ExamScore` | `*string` | 未标注 | 考试分数 |
| `head_pic_id` | `HeadPicId` | `*string` | 未标注 | 工作证图片上传接口返回的id |
| `intention_city` | `IntentionCity` | `*string` | 未标注 | 求职意向，城市的adccode 例如110000（北京市） |
| `intention_city_name` | `IntentionCityName` | `*string` | 未标注 | 意向城市名，例如北京市，传入北京即可，大理白族自治州，传入大理即可 |
| `issue_date` | `IssueDate` | `*string` | 未标注 | 发证日期，ms时间戳 |
| `job_id` | `JobId` | `*string` | 未标注 | 岗位对应的职业id，字段参考https://www.yuque.com/wx3dkp/gwckr4/bs03t3 |
| `job_name` | `JobName` | `*string` | 未标注 | 职业名称 |
| `ka_job_id` | `KaJobId` | `*string` | 未标注 | 商户职业Id |
| `ka_job_name` | `KaJobName` | `*string` | 未标注 | 商户职业名称 |
| `ka_profession_id` | `KaProfessionId` | `*string` | 未标注 | 商户行业Id |
| `ka_profession_name` | `KaProfessionName` | `*string` | 未标注 | 商户行业名称 |
| `location` | `Location` | `*string` | 未标注 | 省份 |
| `month` | `Month` | `*string` | 未标注 | 毕业的月 |
| `once_token` | `OnceToken` | `*string` | 未标注 | 对应类型相关接口生成 |
| `profession_id` | `ProfessionId` | `*string` | 未标注 | 岗位对应的行业id，字段参考https://www.yuque.com/wx3dkp/gwckr4/bs03t3 |
| `profession_name` | `ProfessionName` | `*string` | 未标注 | 行业名称 |
| `salary_max` | `SalaryMax` | `*string` | 未标注 | 求职意向，薪水范围，不能低于salary_min |
| `salary_min` | `SalaryMin` | `*string` | 未标注 | 求职意向，薪水范围，不能高于salary_max |
| `salary_unit` | `SalaryUnit` | `*string` | 未标注 | 薪资单位，元/日(DAY)、元/次(TIME)、元/月(MONTH)、元/小时(HOUR)、元/件 (NUM)、元/周 (WEEK)、其他 (OTHER) |
| `school_name` | `SchoolName` | `*string` | 未标注 | 学校 |
| `skill_name` | `SkillName` | `*string` | 未标注 | 用英文逗号分隔，最多6项 |
| `start_time` | `StartTime` | `*string` | 未标注 | 参加工作的时间，格式形如 2011-07 |
| `type` | `Type` | `*string` | 未标注 | 信息类型 |
| `valid_date_end` | `ValidDateEnd` | `*string` | 未标注 | 有效期：开始结束 ms 时间戳 |
| `valid_date_start` | `ValidDateStart` | `*string` | 未标注 | 有效期：开始时间 ms时间戳 |
| `verify_status` | `VerifyStatus` | `*string` | 未标注 | 校验状态 |
| `work_desc` | `WorkDesc` | `*string` | 未标注 | 工作描述 |
| `work_end_time` | `WorkEndTime` | `*string` | 未标注 | 工作结束时间，毫秒级时间戳 |
| `work_place` | `WorkPlace` | `*string` | 未标注 | 工作经历工作地 |
| `work_property` | `WorkProperty` | `*string` | 未标注 | 岗位属性：全职(FULL_TIME)或者兼职(PART_TIME) |
| `work_start_time` | `WorkStartTime` | `*string` | 未标注 | 工作开始时间，毫秒级时间戳 |
| `year` | `Year` | `*string` | 未标注 | 毕业的年 |

<a id="model-jobworthpositioninfo"></a>

### JobWorthPositionInfo

类型定义：[JobWorthPositionInfo](customer_model.go#L193)。

| 协议字段 | Go 字段 | Go 类型 | 必填（快照） | 说明 |
| --- | --- | --- | --- | --- |
| `age` | `Age` | `*string` | 未标注 | 职位要求-年龄要求 无要求（不传）、28岁以上(OVER28)、40岁以下(BELOW40)、25岁～35岁(25TO35) |
| `benefit` | `Benefit` | `*string` | 未标注 | 福利，现有标签，最多5个标签，使用英文,分割 参考 https://www.yuque.com/docs/share/805e9840-ae9d-499c-b566-46b3b8c9fb2d?# 《4.1 服务端接入----职位传入接口》的1.3.1.1 |
| `certifications` | `Certifications` | `*string` | 未标注 | 技能、证书等要求，比如学历、健康证、普通话或其他技能证书等 |
| `company_certificate` | `CompanyCertificate` | `*string` | 未标注 | 发布企业统一社会信用代码 |
| `company_logo_afts_id` | `CompanyLogoAftsId` | `*string` | 未标注 | 上传的aftsid |
| `company_name` | `CompanyName` | `*string` | 未标注 | 企业工商全称 |
| `count` | `Count` | `*int64` | 未标注 | 招聘人数 |
| `education` | `Education` | `*string` | 未标注 | 学历要求,HIGHSCHOOL_AND_BELOW（高中及以下），POLYTECHNIC（中专），COLLEGE（大专），BACHELOR（本科），MASTER（硕士），DOCTOR_AND_ABOVE（博士及以上） |
| `gender` | `Gender` | `*string` | 未标注 | 无要求（不传）、男或者女 |
| `ka_position_id` | `KaPositionId` | `*string` | 未标注 | 合作方岗位id自定义保持唯一性 |
| `position_desc` | `PositionDesc` | `*string` | 未标注 | 长文本描述 |
| `position_id` | `PositionId` | `*string` | 未标注 | 工作证岗位库的职位主键，只有更新职位信息时需要传入。第一次传入岗位后返回 |
| `position_job_id` | `PositionJobId` | `*string` | 未标注 | 岗位对应的行业id，字段参考https://www.yuque.com/docs/share/805e9840-ae9d-499c-b566-46b3b8c9fb2d?# 《4.1 服务端接入----职位传入接口》的1.3.1.2 |
| `position_job_name` | `PositionJobName` | `*string` | 未标注 | 岗位对应的职业名称 |
| `position_profession_id` | `PositionProfessionId` | `*string` | 未标注 | 岗位对应的职业id，字段参考https://www.yuque.com/docs/share/805e9840-ae9d-499c-b566-46b3b8c9fb2d?# 《4.1 服务端接入----职位传入接口》的1.3.1.2 |
| `position_property` | `PositionProperty` | `*string` | 未标注 | 岗位属性：全职(FULL_TIME)或者兼职(PART_TIME) |
| `position_status` | `PositionStatus` | `*string` | 未标注 | 职位的状态，只有上架(UNCHECK)和下架(OFFLINE)两种 |
| `position_title` | `PositionTitle` | `*string` | 未标注 | 短文本描述 |
| `salary_max` | `SalaryMax` | `*string` | 未标注 | 薪水范围，不能低于salary_min |
| `salary_min` | `SalaryMin` | `*string` | 未标注 | 薪水范围，不能高于salary_max |
| `salary_type` | `SalaryType` | `*string` | 未标注 | 薪资类型：日结(DAY)、月结(MONTH)、周结(WEEK)、完工结(DONE)、其他(OTHER) |
| `salary_unit` | `SalaryUnit` | `*string` | 未标注 | 薪资单位，元/日(DAY)、元/次(TIME)、元/月(MONTH)、元/小时(HOUR)、元/件 (NUM)、元/周 (WEEK)、其他 (OTHER) |
| `skill_tag` | `SkillTag` | `*string` | 未标注 | 岗位需要的技能标签，最多5个，使用英文逗号分割 |
| `skip_url` | `SkipUrl` | `*string` | 未标注 | 职位跳转链接，当前只支持支付宝小程序，也就是alipays://platformapi/ 开头 |
| `work_city` | `WorkCity` | `*string` | 未标注 | 高德city code，例：北京010，线上工作9999 参考https://www.yuque.com/docs/share/805e9840-ae9d-499c-b566-46b3b8c9fb2d?# 《4.1 服务端接入----职位传入接口》的1.3.1.3 |
| `work_longitude` | `WorkLongitude` | `*string` | 未标注 | 经纬度，前面是经度，后面是纬度，使用英文逗号隔开 |
| `work_region` | `WorkRegion` | `*string` | 未标注 | 工作地所在的区县,使用高德的adcode，例：朝阳区110105 参考https://www.yuque.com/docs/share/805e9840-ae9d-499c-b566-46b3b8c9fb2d?# 《4.1 服务端接入----职位传入接口》的1.3.1.2 |
| `work_year` | `WorkYear` | `*string` | 未标注 | 工作年限，1年以下（ONE）、1～3年（THREE）、3～5年（FIVE）、5～10年（TEN）、10年以上（OVER_TEN） |
