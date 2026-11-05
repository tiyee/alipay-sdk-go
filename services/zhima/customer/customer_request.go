// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package customer

// JobWorthJobDataAddRequest is the official V3 ZhimaCustomerJobworthJobdataAddModel schema.
type JobWorthJobDataAddRequest struct {
	// CertNo represents cert_no. 身份证证件号
	CertNo *string `json:"cert_no,omitzero"`
	// CertType represents cert_type. 0:身份证1:护照2:台湾同胞回乡证3:港澳居民往来内地通行证4:台湾居民通行证5:港澳居民通行证
	CertType *string `json:"cert_type,omitzero"`
	// ConnKey represents conn_key. 关联key，受理台返回
	ConnKey *string `json:"conn_key,omitzero"`
	// JobDataList represents job_data_list. 回传信息
	JobDataList *[]JobWorthJobData `json:"job_data_list,omitzero"`
	// OpenId represents open_id. 支付宝uid
	OpenId *string `json:"open_id,omitzero"`
	// UserId represents user_id. 支付宝uid
	UserId *string `json:"user_id,omitzero"`
	// UserName represents user_name. 用户中文名
	UserName *string `json:"user_name,omitzero"`
}

// JobWorthPictureUploadRequest is the official V3 ZhimaCustomerJobworthPictureUploadModel schema.
type JobWorthPictureUploadRequest struct {
	// Content represents content. 图片对应的base64字符串，支持jpg和png，1M之内
	Content *string `json:"content,omitzero"`
	// PicType represents pic_type. 只能传入指定类型，具体类型看文档
	PicType *string `json:"pic_type,omitzero"`
}

// JobWorthPositionAddRequest is the official V3 ZhimaCustomerJobworthPositionAddModel schema.
type JobWorthPositionAddRequest struct {
	// JobworthPositionInfo represents jobworth_position_info.
	JobworthPositionInfo *JobWorthPositionInfo `json:"jobworth_position_info,omitzero"`
}

// JobWorthAuthenticationQueryRequest contains parameters for zhima.customer.jobworth.authentication.query.
type JobWorthAuthenticationQueryRequest struct {
	// ConnKey represents conn_key. 调用第一步创建接口的out_agreement_no，由外部传入
	ConnKey *string `query:"conn_key"`
	// OnceToken represents once_token. 查询身份认证结果的凭证，在认证结果到达终态后只允许查询成功一次，若需要再次查询，则需要调用第一个接口再拉起受理台
	OnceToken *string `query:"once_token"`
	// QueryType represents query_type. exact:精确查询,dim:模糊查询
	QueryType *string `query:"query_type"`
	// IdentityType represents identity_type. 04:教师，其它身份类型联系芝麻技术支持确认
	IdentityType *string `query:"identity_type"`
	// ServiceId represents service_id. 信用服务id，商家接入产品功能包时有运营同学分配
	ServiceId *string `query:"service_id"`
	// UserId represents user_id. 蚂蚁统一会员ID 和身份证任选一个传入
	UserId *string `query:"user_id"`
	// OpenId represents open_id. 蚂蚁统一会员ID 和身份证任选一个传入
	OpenId *string `query:"open_id"`
	// CertNo represents cert_no. 证件号
	CertNo *string `query:"cert_no"`
	// CertType represents cert_type. "身份证 : 0"，"护照: 1"，"台湾往来通行证：4","港澳往来通行证：3", "回乡证: 2"；"港澳证件：5"
	CertType *string `query:"cert_type"`
}

// JobWorthCloudResumeQueryRequest contains parameters for zhima.customer.jobworth.cloudresume.query.
type JobWorthCloudResumeQueryRequest struct {
	// ConnKey represents conn_key. 长效的认证令牌：zhima.credit.payafteruse.creditagreement.sign 的out_agreement_no
	ConnKey *string `query:"conn_key"`
	// OnceToken represents once_token. 交换云简历的单次token：zhima.credit.payafteruse.creditagreement.sign 的 extra_param入参onceToken（授权后5分钟内，单次有效
	OnceToken *string `query:"once_token"`
}

// JobWorthInfoQueryRequest contains parameters for zhima.customer.jobworth.info.query.
type JobWorthInfoQueryRequest struct {
	// ConnKey represents conn_key. 外部订单号，zhima.credit.payafteruse.creditagreement.sign的入参 out_agreement_no智能简历场景支持只传该值
	ConnKey *string `query:"conn_key"`
	// UserName represents user_name. 姓名
	UserName *string `query:"user_name"`
	// CertNo represents cert_no. 证件号 ，根据cert_type类型设置对应证件号码，选择身份证校验时必传
	CertNo *string `query:"cert_no"`
	// UserId represents user_id. 蚂蚁统一会员ID ，为2088开头的唯一标识 选择支付宝uid检验时必传，支付宝uid或身份证+身份证类型 查询2选1，都设置时uid优先
	UserId *string `query:"user_id"`
	// OpenId represents open_id. 蚂蚁统一会员ID ，为2088开头的唯一标识 选择支付宝uid检验时必传，支付宝uid或身份证+身份证类型 查询2选1，都设置时uid优先
	OpenId *string `query:"open_id"`
	// BackUrl represents back_url. 若用户没有开通工作证或者芝麻，成功开通后回跳的链接 ，支持http或schema地址
	BackUrl *string `query:"back_url"`
	// CertType represents cert_type. 证件类型 ，枚举值： 0:身份证1:护照2:台湾同胞回乡证3:港澳居民往来内地通行证4:台湾居民通行证5:港澳居民通行证 选择身份证校验时必传，支付宝uid或身份证+身份证类型 查询2选1，都设置时uid优先
	CertType *string `query:"cert_type"`
	// KaVisitorId represents ka_visitor_id. 商户内部的唯一id如果是需要返回h5页面必填
	KaVisitorId *string `query:"ka_visitor_id"`
	// VisitorType represents visitor_type. 访问类型 C = company or P = person 如果是需要返回h5页面必填
	VisitorType *string `query:"visitor_type"`
	// VisitorName represents visitor_name. 商户侧配置访客名称 如果是需要返回h5页面必填
	VisitorName *string `query:"visitor_name"`
	// CompanyCertificate represents company_certificate. 公司社会统一代码h5页面访客展示使用
	CompanyCertificate *string `query:"company_certificate"`
	// VisitorUrl represents visitor_url. 上传图片接口返回的id
	VisitorUrl *string `query:"visitor_url"`
	// HasButton represents has_button. 在h5页面是否展示底部跳转按钮
	HasButton *bool `query:"has_button"`
	// HasNumber represents has_number. true or false 是否展示
	HasNumber *bool `query:"has_number"`
	// HasQrCode represents has_qr_code. 是否展示二维码
	HasQrCode *bool `query:"has_qr_code"`
	// HasHtml represents has_html. 是否需要返回h5页面
	HasHtml *bool `query:"has_html"`
	// JobId represents job_id. 对外使用的jobid
	JobId *string `query:"job_id"`
	// IndustryId represents industry_id. 行业描述ID
	IndustryId *string `query:"industry_id"`
	// CloudResumeScene represents cloud_resume_scene. 子场景
	CloudResumeScene *string `query:"cloud_resume_scene"`
}
