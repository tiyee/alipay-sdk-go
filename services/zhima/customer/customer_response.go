// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package customer

// JobWorthAuthenticationQueryResponse is the official V3 ZhimaCustomerJobworthAuthenticationQueryResponseModel schema.
type JobWorthAuthenticationQueryResponse struct {
	// IdentityResult represents identity_result. SUCCESS代表可信，FAILED代表不可信，WAITING_VERIFY代表未知需要等待结果
	IdentityResult *string `json:"identity_result,omitzero"`
	// IdentityResultSkipUrl represents identity_result_skip_url. 身份认证查询结果页面的小程序链接，有芝麻提供，用户在端内可直接访问
	IdentityResultSkipUrl *string `json:"identity_result_skip_url,omitzero"`
	// MerchantIdentityId represents merchant_identity_id. 表示在商户下已经发起过验证的用户唯一标志，如果用户已经发起过验证，那么这个标志会相同返回
	MerchantIdentityId *string `json:"merchant_identity_id,omitzero"`
	// OpenId represents open_id. 蚂蚁统一会员ID
	OpenId *string `json:"open_id,omitzero"`
	// TokenStatus represents token_status. 代表订单状态
	TokenStatus *string `json:"token_status,omitzero"`
	// UserId represents user_id. 蚂蚁统一会员ID
	UserId *string `json:"user_id,omitzero"`
}

// JobWorthCloudResumeQueryResponse is the official V3 ZhimaCustomerJobworthCloudresumeQueryResponseModel schema.
type JobWorthCloudResumeQueryResponse struct {
	// Birthday represents birthday. 出生日期
	Birthday *string `json:"birthday,omitzero"`
	// Certificates represents certificates. 证书名称
	Certificates *[]CloudResumeCertificateInfo `json:"certificates,omitzero"`
	// EducationExperiences represents education_experiences. 教育经历
	EducationExperiences *[]CloudResumeEducationExperience `json:"education_experiences,omitzero"`
	// Email represents email. 用户在芝麻工作证填写的邮箱
	Email *string `json:"email,omitzero"`
	// Gender represents gender. 性别
	Gender *string `json:"gender,omitzero"`
	// IntentionStatus represents intention_status. 用户的求职状态
	IntentionStatus *string `json:"intention_status,omitzero"`
	// PersonDesc represents person_desc. 用户在芝麻工作证填写的自我介绍
	PersonDesc *string `json:"person_desc,omitzero"`
	// Phone represents phone. 手机号码
	Phone *string `json:"phone,omitzero"`
	// PicUrl represents pic_url.
	PicUrl *CloudResumeHeadPic `json:"pic_url,omitzero"`
	// PositionIntentions represents position_intentions. 求职期望
	PositionIntentions *[]CloudResumePositionIntention `json:"position_intentions,omitzero"`
	// PositionType represents position_type. 职业身份
	PositionType *string `json:"position_type,omitzero"`
	// ResidentialArea represents residential_area. 常驻地址
	ResidentialArea *string `json:"residential_area,omitzero"`
	// Skills represents skills. 技能信息
	Skills *[]CloudResumeSkillInfo `json:"skills,omitzero"`
	// UserName represents user_name. 姓名
	UserName *string `json:"user_name,omitzero"`
	// WorkExperiences represents work_experiences. 工作经历
	WorkExperiences *[]CloudResumeWorkExperience `json:"work_experiences,omitzero"`
	// WorkStartDate represents work_start_date. 工作开始日期（时间戳）
	WorkStartDate *int64 `json:"work_start_date,omitzero"`
}

// JobWorthInfoQueryResponse is the official V3 ZhimaCustomerJobworthInfoQueryResponseModel schema.
type JobWorthInfoQueryResponse struct {
	// AcceptanceId represents acceptance_id. 受理台单号作为jsapi的一个参数。如果不使用jsapi可以忽略
	AcceptanceId *string `json:"acceptance_id,omitzero"`
	// AuthToken represents auth_token. 用于授权校验,授权之前需要通过token校验
	AuthToken *string `json:"auth_token,omitzero"`
	// HtmlUrl represents html_url. h5url
	HtmlUrl *string `json:"html_url,omitzero"`
	// PictureUrl represents picture_url. 职得工作证图片url,可以直接用于展示
	PictureUrl *string `json:"picture_url,omitzero"`
	// SceneResult represents scene_result. 场景授权结果
	SceneResult *bool `json:"scene_result,omitzero"`
	// SubCode represents sub_code. 业务编码
	SubCode *string `json:"sub_code,omitzero"`
	// SubMsg represents sub_msg. 业务中文结果信息
	SubMsg *string `json:"sub_msg,omitzero"`
	// UpdateUrl represents update_url. 工作证图片更新的url,需要跳转到工作证小程序中进行更新
	UpdateUrl *string `json:"update_url,omitzero"`
	// Url represents url. 职得工作证跳转小程序链接
	Url *string `json:"url,omitzero"`
}

// JobWorthJobDataAddResponse is the official V3 ZhimaCustomerJobworthJobdataAddResponseModel schema.
type JobWorthJobDataAddResponse struct {
	// AcceptanceId represents acceptance_id. 受理台单号，用来调用工作证受理台
	AcceptanceId *string `json:"acceptance_id,omitzero"`
	// SubCode represents sub_code. 错误码
	SubCode *string `json:"sub_code,omitzero"`
	// SubMsg represents sub_msg. 业务中文结果信息
	SubMsg *string `json:"sub_msg,omitzero"`
	// Url represents url. 职得工作证跳转小程序链接
	Url *string `json:"url,omitzero"`
}

// JobWorthPictureUploadResponse is the official V3 ZhimaCustomerJobworthPictureUploadResponseModel schema.
type JobWorthPictureUploadResponse struct {
	// PicId represents pic_id. 图片返回ID
	PicId *string `json:"pic_id,omitzero"`
	// SubCode represents sub_code. 错误码
	SubCode *string `json:"sub_code,omitzero"`
	// SubMsg represents sub_msg. 错误原因
	SubMsg *string `json:"sub_msg,omitzero"`
}

// JobWorthPositionAddResponse is the official V3 ZhimaCustomerJobworthPositionAddResponseModel schema.
type JobWorthPositionAddResponse struct {
	// PositionId represents position_id. 岗位信息ID
	PositionId *string `json:"position_id,omitzero"`
	// SubCode represents sub_code. 业务错误
	SubCode *string `json:"sub_code,omitzero"`
	// SubMsg represents sub_msg. 业务错误信息
	SubMsg *string `json:"sub_msg,omitzero"`
}
