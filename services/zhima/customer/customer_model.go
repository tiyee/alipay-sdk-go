// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package customer

// CloudResumeCertificateInfo is the official V3 CloudResumeCertificateInfo schema.
type CloudResumeCertificateInfo struct {
	// CertificateCertLevel represents certificate_cert_level. 证书等级
	CertificateCertLevel *string `json:"certificate_cert_level,omitzero"`
	// CertificateName represents certificate_name. 证书名称
	CertificateName *string `json:"certificate_name,omitzero"`
	// CertificateSendCardAddress represents certificate_send_card_address. 发证机构
	CertificateSendCardAddress *string `json:"certificate_send_card_address,omitzero"`
	// CertificateVerifyStatus represents certificate_verify_status. UNPROMISED=未承诺 PROMISED=已承诺 VERFIED=已验真 UNVERIFIED=验证未通过
	CertificateVerifyStatus *string `json:"certificate_verify_status,omitzero"`
	// InValidityPeriod represents in_validity_period. VALID=有效 INVALID=失效 UNKNOWN=未知
	InValidityPeriod *string `json:"in_validity_period,omitzero"`
}

// CloudResumeEducationExperience is the official V3 CloudResumeEducationExperience schema.
type CloudResumeEducationExperience struct {
	// Degree represents degree. 学历
	Degree *string `json:"degree,omitzero"`
	// EducationStatus represents education_status. 教育状态
	EducationStatus *string `json:"education_status,omitzero"`
	// FinishSchoolMonth represents finish_school_month. 毕业月份
	FinishSchoolMonth *string `json:"finish_school_month,omitzero"`
	// FinishSchoolYear represents finish_school_year. 毕业年份
	FinishSchoolYear *int64 `json:"finish_school_year,omitzero"`
	// InSchoolMonth represents in_school_month. 入学月份
	InSchoolMonth *int64 `json:"in_school_month,omitzero"`
	// InSchoolYear represents in_school_year. 入学年份
	InSchoolYear *int64 `json:"in_school_year,omitzero"`
	// Location represents location. 受教育地点
	Location *string `json:"location,omitzero"`
	// Major represents major. 上学专业
	Major *string `json:"major,omitzero"`
	// Month represents month. 月
	Month *int64 `json:"month,omitzero"`
	// SchoolName represents school_name. 学校名称
	SchoolName *string `json:"school_name,omitzero"`
	// TongZhao represents tong_zhao. 是否统招
	TongZhao *bool `json:"tong_zhao,omitzero"`
	// Year represents year. 年
	Year *int64 `json:"year,omitzero"`
}

// CloudResumeHeadPic is the official V3 CloudResumeHeadPic schema.
type CloudResumeHeadPic struct {
	// PicUrl represents pic_url. 头像连接url
	PicUrl *string `json:"pic_url,omitzero"`
}

// CloudResumePositionIntention is the official V3 CloudResumePositionIntention schema.
type CloudResumePositionIntention struct {
	// IntentionCity represents intention_city. 意向城市Code。具体地区编码参见https://lbs.amap.com/api/webservice/download 里面城市编码表
	IntentionCity *string `json:"intention_city,omitzero"`
	// JobId represents job_id. 职业id(这个字段在目前版本不对外暴露)
	JobId *string `json:"job_id,omitzero"`
	// JobName represents job_name. 职业名称
	JobName *string `json:"job_name,omitzero"`
	// ProfessionId represents profession_id. 行业id(这个字段在目前版本不对外暴露)
	ProfessionId *string `json:"profession_id,omitzero"`
	// ProfessionName represents profession_name. 行业名称
	ProfessionName *string `json:"profession_name,omitzero"`
	// SalaryMax represents salary_max. 最大工资
	SalaryMax *string `json:"salary_max,omitzero"`
	// SalaryMin represents salary_min. 最小薪资
	SalaryMin *string `json:"salary_min,omitzero"`
	// SalaryUnit represents salary_unit. 工资单位，/月 /日
	SalaryUnit *string `json:"salary_unit,omitzero"`
	// WorkProperty represents work_property. 工作属性
	WorkProperty *string `json:"work_property,omitzero"`
}

// CloudResumeSkillInfo is the official V3 CloudResumeSkillInfo schema.
type CloudResumeSkillInfo struct {
	// SkillName represents skill_name. 技能标签名字
	SkillName *string `json:"skill_name,omitzero"`
}

// CloudResumeWorkExperience is the official V3 CloudResumeWorkExperience schema.
type CloudResumeWorkExperience struct {
	// CompanyName represents company_name. 公司名称
	CompanyName *string `json:"company_name,omitzero"`
	// JobId represents job_id. 职业id
	JobId *string `json:"job_id,omitzero"`
	// JobName represents job_name. 职业
	JobName *string `json:"job_name,omitzero"`
	// PositionName represents position_name. 用户在求职意向页面自主录入
	PositionName *string `json:"position_name,omitzero"`
	// ProfessionId represents profession_id. 行业id
	ProfessionId *string `json:"profession_id,omitzero"`
	// ProfessionName represents profession_name. 行业名称
	ProfessionName *string `json:"profession_name,omitzero"`
	// WorkDesc represents work_desc. 工作描述
	WorkDesc *string `json:"work_desc,omitzero"`
	// WorkEndTime represents work_end_time. 工作结束时间
	WorkEndTime *int64 `json:"work_end_time,omitzero"`
	// WorkStartTime represents work_start_time. 工作开始日期
	WorkStartTime *int64 `json:"work_start_time,omitzero"`
}

// JobWorthJobData is the official V3 JobWorthJobData schema.
type JobWorthJobData struct {
	// CertificateGrantInstitution represents certificate_grant_institution. 证书发证机构
	CertificateGrantInstitution *string `json:"certificate_grant_institution,omitzero"`
	// CertificateId represents certificate_id. 证书编号
	CertificateId *string `json:"certificate_id,omitzero"`
	// CertificateLevel represents certificate_level. 证书等级：中文
	CertificateLevel *string `json:"certificate_level,omitzero"`
	// CertificateName represents certificate_name. 证书名称
	CertificateName *string `json:"certificate_name,omitzero"`
	// CertificatePicId represents certificate_pic_id. 工作证图片上传接口返回的id
	CertificatePicId *string `json:"certificate_pic_id,omitzero"`
	// CompanyName represents company_name. 公司名称
	CompanyName *string `json:"company_name,omitzero"`
	// Degree represents degree. 学历，可以用以下字段HIGHSCHOOL_AND_BELOW（高中及以下），POLYTECHNIC（中专），COLLEGE（大专），BACHELOR（本科），MASTER（硕士），DOCTOR_AND_ABOVE（博士及以上）
	Degree *string `json:"degree,omitzero"`
	// DeliveryPositionId represents delivery_position_id. 商户侧投递岗位id
	DeliveryPositionId *string `json:"delivery_position_id,omitzero"`
	// DeliveryTime represents delivery_time. 投递时间戳
	DeliveryTime *int64 `json:"delivery_time,omitzero"`
	// EducationStatus represents education_status. 教育状态只有两种，GRADUATE(毕业)，STUDY（未毕业）
	EducationStatus *string `json:"education_status,omitzero"`
	// EntryNo represents entry_no. 报名流水号
	EntryNo *string `json:"entry_no,omitzero"`
	// ExamScore represents exam_score. 考试分数
	ExamScore *string `json:"exam_score,omitzero"`
	// HeadPicId represents head_pic_id. 工作证图片上传接口返回的id
	HeadPicId *string `json:"head_pic_id,omitzero"`
	// IntentionCity represents intention_city. 求职意向，城市的adccode 例如110000（北京市）
	IntentionCity *string `json:"intention_city,omitzero"`
	// IntentionCityName represents intention_city_name. 意向城市名，例如北京市，传入北京即可，大理白族自治州，传入大理即可
	IntentionCityName *string `json:"intention_city_name,omitzero"`
	// IssueDate represents issue_date. 发证日期，ms时间戳
	IssueDate *string `json:"issue_date,omitzero"`
	// JobId represents job_id. 岗位对应的职业id，字段参考https://www.yuque.com/wx3dkp/gwckr4/bs03t3
	JobId *string `json:"job_id,omitzero"`
	// JobName represents job_name. 职业名称
	JobName *string `json:"job_name,omitzero"`
	// KaJobId represents ka_job_id. 商户职业Id
	KaJobId *string `json:"ka_job_id,omitzero"`
	// KaJobName represents ka_job_name. 商户职业名称
	KaJobName *string `json:"ka_job_name,omitzero"`
	// KaProfessionId represents ka_profession_id. 商户行业Id
	KaProfessionId *string `json:"ka_profession_id,omitzero"`
	// KaProfessionName represents ka_profession_name. 商户行业名称
	KaProfessionName *string `json:"ka_profession_name,omitzero"`
	// Location represents location. 省份
	Location *string `json:"location,omitzero"`
	// Month represents month. 毕业的月
	Month *string `json:"month,omitzero"`
	// OnceToken represents once_token. 对应类型相关接口生成
	OnceToken *string `json:"once_token,omitzero"`
	// ProfessionId represents profession_id. 岗位对应的行业id，字段参考https://www.yuque.com/wx3dkp/gwckr4/bs03t3
	ProfessionId *string `json:"profession_id,omitzero"`
	// ProfessionName represents profession_name. 行业名称
	ProfessionName *string `json:"profession_name,omitzero"`
	// SalaryMax represents salary_max. 求职意向，薪水范围，不能低于salary_min
	SalaryMax *string `json:"salary_max,omitzero"`
	// SalaryMin represents salary_min. 求职意向，薪水范围，不能高于salary_max
	SalaryMin *string `json:"salary_min,omitzero"`
	// SalaryUnit represents salary_unit. 薪资单位，元/日(DAY)、元/次(TIME)、元/月(MONTH)、元/小时(HOUR)、元/件 (NUM)、元/周 (WEEK)、其他 (OTHER)
	SalaryUnit *string `json:"salary_unit,omitzero"`
	// SchoolName represents school_name. 学校
	SchoolName *string `json:"school_name,omitzero"`
	// SkillName represents skill_name. 用英文逗号分隔，最多6项
	SkillName *string `json:"skill_name,omitzero"`
	// StartTime represents start_time. 参加工作的时间，格式形如 2011-07
	StartTime *string `json:"start_time,omitzero"`
	// Type represents type. 信息类型
	Type *string `json:"type,omitzero"`
	// ValidDateEnd represents valid_date_end. 有效期：开始结束 ms 时间戳
	ValidDateEnd *string `json:"valid_date_end,omitzero"`
	// ValidDateStart represents valid_date_start. 有效期：开始时间 ms时间戳
	ValidDateStart *string `json:"valid_date_start,omitzero"`
	// VerifyStatus represents verify_status. 校验状态
	VerifyStatus *string `json:"verify_status,omitzero"`
	// WorkDesc represents work_desc. 工作描述
	WorkDesc *string `json:"work_desc,omitzero"`
	// WorkEndTime represents work_end_time. 工作结束时间，毫秒级时间戳
	WorkEndTime *string `json:"work_end_time,omitzero"`
	// WorkPlace represents work_place. 工作经历工作地
	WorkPlace *string `json:"work_place,omitzero"`
	// WorkProperty represents work_property. 岗位属性：全职(FULL_TIME)或者兼职(PART_TIME)
	WorkProperty *string `json:"work_property,omitzero"`
	// WorkStartTime represents work_start_time. 工作开始时间，毫秒级时间戳
	WorkStartTime *string `json:"work_start_time,omitzero"`
	// Year represents year. 毕业的年
	Year *string `json:"year,omitzero"`
}

// JobWorthPositionInfo is the official V3 JobWorthPositionInfo schema.
type JobWorthPositionInfo struct {
	// Age represents age. 职位要求-年龄要求 无要求（不传）、28岁以上(OVER28)、40岁以下(BELOW40)、25岁～35岁(25TO35)
	Age *string `json:"age,omitzero"`
	// Benefit represents benefit. 福利，现有标签，最多5个标签，使用英文,分割 参考 https://www.yuque.com/docs/share/805e9840-ae9d-499c-b566-46b3b8c9fb2d?# 《4.1 服务端接入----职位传入接口》的1.3.1.1
	Benefit *string `json:"benefit,omitzero"`
	// Certifications represents certifications. 技能、证书等要求，比如学历、健康证、普通话或其他技能证书等
	Certifications *string `json:"certifications,omitzero"`
	// CompanyCertificate represents company_certificate. 发布企业统一社会信用代码
	CompanyCertificate *string `json:"company_certificate,omitzero"`
	// CompanyLogoAftsId represents company_logo_afts_id. 上传的aftsid
	CompanyLogoAftsId *string `json:"company_logo_afts_id,omitzero"`
	// CompanyName represents company_name. 企业工商全称
	CompanyName *string `json:"company_name,omitzero"`
	// Count represents count. 招聘人数
	Count *int64 `json:"count,omitzero"`
	// Education represents education. 学历要求,HIGHSCHOOL_AND_BELOW（高中及以下），POLYTECHNIC（中专），COLLEGE（大专），BACHELOR（本科），MASTER（硕士），DOCTOR_AND_ABOVE（博士及以上）
	Education *string `json:"education,omitzero"`
	// Gender represents gender. 无要求（不传）、男或者女
	Gender *string `json:"gender,omitzero"`
	// KaPositionId represents ka_position_id. 合作方岗位id自定义保持唯一性
	KaPositionId *string `json:"ka_position_id,omitzero"`
	// PositionDesc represents position_desc. 长文本描述
	PositionDesc *string `json:"position_desc,omitzero"`
	// PositionId represents position_id. 工作证岗位库的职位主键，只有更新职位信息时需要传入。第一次传入岗位后返回
	PositionId *string `json:"position_id,omitzero"`
	// PositionJobId represents position_job_id. 岗位对应的行业id，字段参考https://www.yuque.com/docs/share/805e9840-ae9d-499c-b566-46b3b8c9fb2d?# 《4.1 服务端接入----职位传入接口》的1.3.1.2
	PositionJobId *string `json:"position_job_id,omitzero"`
	// PositionJobName represents position_job_name. 岗位对应的职业名称
	PositionJobName *string `json:"position_job_name,omitzero"`
	// PositionProfessionId represents position_profession_id. 岗位对应的职业id，字段参考https://www.yuque.com/docs/share/805e9840-ae9d-499c-b566-46b3b8c9fb2d?# 《4.1 服务端接入----职位传入接口》的1.3.1.2
	PositionProfessionId *string `json:"position_profession_id,omitzero"`
	// PositionProperty represents position_property. 岗位属性：全职(FULL_TIME)或者兼职(PART_TIME)
	PositionProperty *string `json:"position_property,omitzero"`
	// PositionStatus represents position_status. 职位的状态，只有上架(UNCHECK)和下架(OFFLINE)两种
	PositionStatus *string `json:"position_status,omitzero"`
	// PositionTitle represents position_title. 短文本描述
	PositionTitle *string `json:"position_title,omitzero"`
	// SalaryMax represents salary_max. 薪水范围，不能低于salary_min
	SalaryMax *string `json:"salary_max,omitzero"`
	// SalaryMin represents salary_min. 薪水范围，不能高于salary_max
	SalaryMin *string `json:"salary_min,omitzero"`
	// SalaryType represents salary_type. 薪资类型：日结(DAY)、月结(MONTH)、周结(WEEK)、完工结(DONE)、其他(OTHER)
	SalaryType *string `json:"salary_type,omitzero"`
	// SalaryUnit represents salary_unit. 薪资单位，元/日(DAY)、元/次(TIME)、元/月(MONTH)、元/小时(HOUR)、元/件 (NUM)、元/周 (WEEK)、其他 (OTHER)
	SalaryUnit *string `json:"salary_unit,omitzero"`
	// SkillTag represents skill_tag. 岗位需要的技能标签，最多5个，使用英文逗号分割
	SkillTag *string `json:"skill_tag,omitzero"`
	// SkipUrl represents skip_url. 职位跳转链接，当前只支持支付宝小程序，也就是alipays://platformapi/ 开头
	SkipUrl *string `json:"skip_url,omitzero"`
	// WorkCity represents work_city. 高德city code，例：北京010，线上工作9999 参考https://www.yuque.com/docs/share/805e9840-ae9d-499c-b566-46b3b8c9fb2d?# 《4.1 服务端接入----职位传入接口》的1.3.1.3
	WorkCity *string `json:"work_city,omitzero"`
	// WorkLongitude represents work_longitude. 经纬度，前面是经度，后面是纬度，使用英文逗号隔开
	WorkLongitude *string `json:"work_longitude,omitzero"`
	// WorkRegion represents work_region. 工作地所在的区县,使用高德的adcode，例：朝阳区110105 参考https://www.yuque.com/docs/share/805e9840-ae9d-499c-b566-46b3b8c9fb2d?# 《4.1 服务端接入----职位传入接口》的1.3.1.2
	WorkRegion *string `json:"work_region,omitzero"`
	// WorkYear represents work_year. 工作年限，1年以下（ONE）、1～3年（THREE）、3～5年（FIVE）、5～10年（TEN）、10年以上（OVER_TEN）
	WorkYear *string `json:"work_year,omitzero"`
}
