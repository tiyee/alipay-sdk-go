package ad

// ConversionData 是支付宝官方 V3 OpenAPI 模型。
type ConversionData struct {
	// AdId 是广告创意id，来自：<a href='https://adpub.alipay.com/lark/adrlark/qm1v2gtpvcftele4'>监测上报</a>或<a href='https://adpub.alipay.com/lark/adrlark/gw3740rggm4n5gar'>落地页宏替换</a>配置的__AD_ID__宏参发生用户点击替换后的值。
	AdId *string `json:"ad_id,omitempty"`
	// AttributeList 是转化事件属性信息，用于转化事件类型相关属性规则上传。 可支持上传属性的转化事件类型及属性规则请参考该文档：<a href='https://adpub.alipay.com/adrlark/ivdktpyh511x9r6i'>转化事件类型属性规则</a>。
	AttributeList *[]ConversionProperty `json:"attribute_list,omitempty"`
	Attributes    *ConversionProperty   `json:"attributes,omitempty"`
	// BizNo 是转化流水号：由用户自定义，用于幂等。
	BizNo *string `json:"biz_no,omitempty"`
	// CallbackExtInfo 是来自：<a href='https://adpub.alipay.com/lark/adrlark/qm1v2gtpvcftele4'>监测上报</a>或<a href='https://adpub.alipay.com/lark/adrlark/gw3740rggm4n5gar'>落地页宏替换</a>配置的__CALLBACK_EXT_INFO__宏参发生用户点击替换后的值，需url decode后传回。。
	CallbackExtInfo *string `json:"callback_ext_info,omitempty"`
	// Cid 是用于精准归因，来自：<a href='https://adpub.alipay.com/lark/adrlark/qm1v2gtpvcftele4'>监测上报</a>或<a href='https://adpub.alipay.com/lark/adrlark/gw3740rggm4n5gar'>落地页宏替换</a>配置的__CID__宏参发生用户点击替换后的值。<br> 支持cid上报的转化事件类型请参见：<br> <a href='https://adpub.alipay.com/lark/adrlark/pg4vb1c0g0u6fx7b'>app推广类转化事件类型</a><br> <a href='https://adpub.alipay.com/lark/adrlark/rvzylhryxh37yplt'>通用类转化事件类型</a>。
	Cid *string `json:"cid,omitempty"`
	// ConversionAmount 是转化金额，单位分。 金额类转化事件必传。
	ConversionAmount *string `json:"conversion_amount,omitempty"`
	// ConversionId 是转化事件id，当source=COMMON_CONVERSION_ID需上传该字段，其他情况无需上传。。
	ConversionId *string `json:"conversion_id,omitempty"`
	// ConversionTime 是转化时间，UTC 时间戳，单位：秒。
	ConversionTime *int64 `json:"conversion_time,omitempty"`
	// ConversionType 是转化事件类型数字。
	ConversionType *string `json:"conversion_type,omitempty"`
	// CreativeId 是广告创意id，来自：<a href='https://adpub.alipay.com/lark/adrlark/qm1v2gtpvcftele4'>监测上报</a>或<a href='https://adpub.alipay.com/lark/adrlark/gw3740rggm4n5gar'>落地页宏替换</a>配置的__AD_ID__宏参发生用户点击替换后的值。
	CreativeId *string `json:"creative_id,omitempty"`
	// DataId 是当source=XLIGHT或者DATASET时使用： 数据集id(接口升级，该参数已不再使用，故废弃该参数，但不影响历史数据使用。)。
	DataId *string `json:"data_id,omitempty"`
	// DataSrcType 是当source=OTHER时使用: 数据类型： KR_MEMBER - 客如云入会 KR_TRADE - 客如云交易 TB_LIVE - 淘宝直播(接口升级，该参数已不再使用，故废弃该参数，但不影响历史数据使用。)。
	DataSrcType *string `json:"data_src_type,omitempty"`
	// GroupId 是广告单元id，来自：<a href='https://adpub.alipay.com/lark/adrlark/qm1v2gtpvcftele4'>监测上报</a>或<a href='https://adpub.alipay.com/lark/adrlark/gw3740rggm4n5gar'>落地页宏替换</a>配置的__GROUP_ID__宏参发生用户点击替换后的值。
	GroupId *string `json:"group_id,omitempty"`
	// JoinChannel 是XLIGHT - 灯火归因 MERCHANT -商家自主归因 该字段若为空则默认为XLIGHT。
	JoinChannel *string `json:"join_channel,omitempty"`
	// PlanId 是广告计划id，来自：<a href='https://adpub.alipay.com/lark/adrlark/qm1v2gtpvcftele4'>监测上报</a>或<a href='https://adpub.alipay.com/lark/adrlark/gw3740rggm4n5gar'>落地页宏替换</a>配置的__PLAN_ID__宏参发生用户点击替换后的值。
	PlanId *string `json:"plan_id,omitempty"`
	// PrincipalId 是商户在灯火pb端的id, 可代替principal_tag。。
	PrincipalId *string `json:"principal_id,omitempty"`
	// PrincipalTag 是商家标志，可代替principal_id。获取方法： <a href='https://adpub.alipay.com/lark/adrlark/sbk3vetg6ki5m4r8'>principal_tag获取sop</a><br>。
	PrincipalTag *string `json:"principal_tag,omitempty"`
	// PropertyList 是转化归因字段列表。不同转化事件类型回传property_list不同。请参见回传示例：<br> <a href="https://adpub.alipay.com/lark/adrlark/pvbgwf72ea090877">通用类转化事件回传示例</a><br> <a href="https://adpub.alipay.com/lark/adrlark/kgnnk94p0p1mwbgp">APP推广类转化事件回传示例</a><br>。
	PropertyList *[]ConversionProperty `json:"property_list,omitempty"`
	// Source 是<br> COMMON_TARGET：<a href='https://adpub.alipay.com/lark/adrlark/rvzylhryxh37yplt'>通用类转化事件类型</a>适用<br> CALLBACK：<a href='https://adpub.alipay.com/lark/adrlark/pg4vb1c0g0u6fx7b'>APP推广类转化事件类型</a>适用<br> 不同source的参数回传规则请参考： COMMON_TARGET：<a href='https://adpub.alipay.com/lark/adrlark/dxw7fkkdnhm45spm'>通用转化回传参数说明</a><br> CALLBACK： <a href='https://adpub.alipay.com/lark/adrlark/wxou6agg3z1o6wyn'>APP推广类转化回传参数说明</a>。
	Source *string `json:"source,omitempty"`
	// TargetId 是当source=OTHER时使用： 主体id，例如品牌id(接口升级，该参数已不再使用，故废弃该参数，但不影响历史数据使用。)。
	TargetId *string `json:"target_id,omitempty"`
	// TargetType 是当source=OTHER时使用： 主体类型： BRAND - 品牌 STORE - 店铺 LIVE - 直播 等等(接口升级，该参数已不再使用，故废弃该参数，但不影响历史数据使用。)。
	TargetType *string `json:"target_type,omitempty"`
	// Uuid 是发生转化用户的唯一标识。 uuid_type不同，uuid传值不同：<br> 1. 若uuid_type=【PID】：uuid为发生转化用户的2088id。若获取不到，可固定一个值后，上传callback_ext_info参数。<br> 2. 若uuid_type=【OAID_MD5】【IDFA_MD5】【IP_UA】【CAID】【MEI】: uuid传的值为<a href='https://adpub.alipay.com/lark/adrlark/qm1v2gtpvcftele4'>监测上报</a>或<a href='https://adpub.alipay.com/lark/adrlark/gw3740rggm4n5gar'>落地页宏替换</a>配置的对应宏参下发的值<br> 其中当uuid_type=IP_UA时：__IP__、__UA__宏参需用"_"拼接上传<br>。
	Uuid *string `json:"uuid,omitempty"`
	// UuidOpenId 是支付宝用户在应用维度下的唯一标识。
	UuidOpenId *string `json:"uuid_open_id,omitempty"`
	// UuidType 是<a href='https://adpub.alipay.com/lark/adrlark/rvzylhryxh37yplt'>通用类转化事件类型：</a><br> uuid_type固定为PID：表示转化用户唯一标识<br> <a href='https://adpub.alipay.com/lark/adrlark/pg4vb1c0g0u6fx7b'>APP推广类转化事件类型：</a><br> OAID_MD5：表示OAID原值MD5加密后的值<br> IDFA_MD5：表示IDFA原值MD5加密后的值<br> IP_UA：表示IP_UA拼接值，以下划线拼接<br> CAID：表示CAID原值<br> IMEI：表示IMEI原值<br>。
	UuidType *string `json:"uuid_type,omitempty"`
}

// ConversionProperty 是支付宝官方 V3 OpenAPI 模型。
type ConversionProperty struct {
	// Key 是业务转化数据属性。
	Key *string `json:"key,omitempty"`
	// Name 是转化属性名称。
	Name *string `json:"name,omitempty"`
	// Value 是业务转化属性实例值， property_list的value值需与转化管理里创建转化事件时填写的归因口径值保持一致<br> attrbute_list的value值请参考该文档：<a href='https://adpub.alipay.com/adrlark/ivdktpyh511x9r6i'>转化事件类型属性规则</a>。
	Value *string `json:"value,omitempty"`
}

// PromotePageData 是支付宝官方 V3 OpenAPI 模型。
type PromotePageData struct {
	// BizNo 是留资流水号。
	BizNo *string `json:"biz_no,omitempty"`
	// EncryptUid 是加密uid，用于深转数据回传。
	EncryptUid *string `json:"encrypt_uid,omitempty"`
	// PropertyList 是留资属性实例数据列表。
	PropertyList *[]PromotePagePropertyInstance `json:"property_list,omitempty"`
}

// PromotePageDetail 是支付宝官方 V3 OpenAPI 模型。
type PromotePageDetail struct {
	// GmtCreate 是推广页创建时间，格式：yyyy-mm-dd hh:mm:ss。
	GmtCreate *string `json:"gmt_create,omitempty"`
	// Id 是推广页id。
	Id *int64 `json:"id,omitempty"`
	// Name 是推广页名称。
	Name *string `json:"name,omitempty"`
	// PropertyList 是推广留资schema。
	PropertyList *[]PromotePageProperty `json:"property_list,omitempty"`
	// Type 是推广页类型：COLLECT_INFO - 普通留资；TRADE - 交易留资。
	Type *string `json:"type,omitempty"`
}

// PromotePageProperty 是支付宝官方 V3 OpenAPI 模型。
type PromotePageProperty struct {
	// Key 是留资属性key, 如: 1. principal_name: 商家名称 2. market_target_name: 营销目标 3. plan_name: 计划名称 4. plan_id: 计划id 5. group_name: 单元名称 6. group_id: 单元id 7. ad_name: 创意名称 8. creative_id: 创意id 9. click_id: 点击id 等。
	Key *string `json:"key,omitempty"`
	// Name 是留资属性名称。
	Name *string `json:"name,omitempty"`
	// Type 是自建站留资字段类型。
	Type *string `json:"type,omitempty"`
}

// PromotePagePropertyInstance 是支付宝官方 V3 OpenAPI 模型。
type PromotePagePropertyInstance struct {
	// Key 是留资属性key, 如: 1. principal_name: 商家名称 2. market_target_name: 营销目标 3. plan_name: 计划名称 4. plan_id: 计划id 5. group_name: 单元名称 6. group_id: 单元id 7. ad_name: 创意名称 8. creative_id: 创意id 9. click_id: 点击id 等。
	Key *string `json:"key,omitempty"`
	// Name 是留资属性名称。
	Name *string `json:"name,omitempty"`
	// Type 是自建站留资字段的类型。
	Type *string `json:"type,omitempty"`
	// Value 是留资属性实例值。
	Value *string `json:"value,omitempty"`
}
