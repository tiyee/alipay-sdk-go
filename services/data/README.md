# data 接口

导入路径：`github.com/tiyee/alipay-sdk-go/services/data`。本目录汇总业务子包；跨业务共享模型按需保留在本包。

接口方法按单词边界使用大驼峰命名；请求、响应与专属模型位于对应业务包，共享模型从共同父包引用。

接入方式见[项目说明](../../README.md)。默认测试离线执行。

## 业务子包

| 导入子路径 | 操作数 | 方法 |
| --- | --- | --- |
| [`bill`](bill/README.md) | 15 | `BillDownloadURLQuery` / `BillBalance` / `BillBalanceHis` / `BillAccountLog` / `BillAccountBookEReceiptApply` / `BillAccountBookEReceiptQuery` / `BillBailQuery` / `BillBizFundAgentQuery` / `BillBuyQuery` / `BillEReceiptApply` / `BillEReceiptQuery` / `BillEReceiptAgentApply` / `BillSellQuery` / `BillTransferQuery` / `BillTransferAccountBookQuery` |
| [`dataservice/ad`](dataservice/ad/README.md) | 3 | `DataServiceAdConversionUpload` / `DataServiceAdPromotePageBatchQuery` / `DataServiceAdPromotePageDownload` |

## 参数约定与注意事项

- 参数表依据 2026-09-23 版 OpenAPI 快照与当前 SDK 模型整理。描述文件的获取方式见支付宝开放平台的[SDK 与工具](https://opendocs.alipay.com/open-v3/065bsc)。
- 参数表仅列出当前 SDK 公开的请求字段。“必填（快照）”为快照明确标记的结果；“未标注”不代表可选，“否”也不排除特定场景条件必填。互斥、二选一、产品限制及必填条件须结合参数说明与具体产品文档确认。
- Go 指针用于表达未传值，`omitempty` 等标签仅控制编码；不能据此推断业务必填性。嵌套对象参数见文末，具名枚举链接指向源码定义。
- 公共签名与响应验签由 Client 处理；代调用授权令牌按需使用 `core.WithAppAuthToken` 传入。REST 调用不接受 `core.WithNotifyURL`，接口自有的通知字段按请求结构体传递。
- SDK 默认不自动重试。写操作超时后须按接口规则查询状态，并遵守业务请求号的幂等要求；HTTP 或调用成功仍需检查具体业务状态。金额按接口定义使用精确十进制字符串，避免用浮点数计算。
- 产品开通权限、授权及认证模式按具体接口确认。当前仓库尚未完成真实沙箱验收；生成接口尚未逐项验证业务条件。
