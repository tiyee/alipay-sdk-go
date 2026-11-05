package sp

import (
	"github.com/tiyee/alipay-sdk-go/core"
)

// SpLeadsExpandCreateRequest 是支付宝官方 V3 OpenAPI 模型。
type SpLeadsExpandCreateRequest struct {
	// Address 是拓展助手商机地址。
	Address *string `json:"address,omitempty"`
	// CityCode 是拓展助手商机地址市code。
	CityCode *string `json:"city_code,omitempty"`
	// CityName 是拓展助手商机地址市名称。
	CityName *string `json:"city_name,omitempty"`
	// DistrictCode 是拓展助手商机地址区code。
	DistrictCode *string `json:"district_code,omitempty"`
	// DistrictName 是拓展助手商机地址区名称。
	DistrictName *string `json:"district_name,omitempty"`
	// Email 是拓展助手商机邮箱。
	Email *string `json:"email,omitempty"`
	// Latitude 是拓展助手商机地址纬度。
	Latitude *string `json:"latitude,omitempty"`
	// Longitude 是拓展助手商机地址经度。
	Longitude *string `json:"longitude,omitempty"`
	// Mcc 是mcc行业。
	Mcc *string `json:"mcc,omitempty"`
	// MerchantPid 是2088账号。
	MerchantPid *string `json:"merchant_pid,omitempty"`
	// Name 是拓展助手商机名称。
	Name *string `json:"name,omitempty"`
	// OutBizNo 是外部幂等唯一键。
	OutBizNo *string `json:"out_biz_no,omitempty"`
	// Phone 是拓展助手商机联系电话。
	Phone *string `json:"phone,omitempty"`
	// ProvinceCode 是拓展助手商机地址省code。
	ProvinceCode *string `json:"province_code,omitempty"`
	// ProvinceName 是拓展助手商机地址省名称。
	ProvinceName *string `json:"province_name,omitempty"`
	// Scene 是拓展助手场景，参数约定。
	Scene *string `json:"scene,omitempty"`
}

// SpMerchantInconsistentApproveRequest 是支付宝官方 V3 OpenAPI 模型。
type SpMerchantInconsistentApproveRequest struct {
	// BackCard 是身份证背面照，图片文件大小在 50K-5M 之间，不限制长宽，支持 png、bmp、gif、jpg、jpeg 格式。 请传入<a href="https://opendocs.alipay.com/apis/01ea4t">alipay.open.sp.image.upload</a>(图片上传接口) 返回的 image_id。。
	BackCard *string `json:"back_card,omitempty"`
	// CommitmentLetter 是实际经营人承诺函照片，要求证件文本信息清晰可见，图片文件大小在 50K-5M 之间，不限制长宽，支持 png、bmp、gif、jpg、jpeg 格式。 请传入<a href="https://opendocs.alipay.com/apis/01ea4t">alipay.open.sp.image.upload</a>(图片上传接口) 返回的 image_id。 该资质是否必传请参见报名资质要求。。
	CommitmentLetter *string `json:"commitment_letter,omitempty"`
	// FrontCard 是身份证正面照，要求证件文本信息清晰可见，图片文件大小在 50K-5M 之间，不限制长宽，支持 png、bmp、gif、jpg、jpeg 格式. 请传<a href="https://opendocs.alipay.com/apis/01ea4t">alipay.open.sp.image.upload</a>(图片上传接口) 返回的 image_id。。
	FrontCard *string `json:"front_card,omitempty"`
	// HandheldBusinessLicense 是手持营业执照合照，要求证件文本信息清晰可见，图片文件大小在 50K-5M 之间，不限制长宽，支持 png、bmp、gif、jpg、jpeg 格式。 请传入<a href="https://opendocs.alipay.com/apis/01ea4t">alipay.open.sp.image.upload</a>(图片上传接口) 返回的 image_id。 该资质是否必传请参见报名资质要求。。
	HandheldBusinessLicense *string `json:"handheld_business_license,omitempty"`
	// HandheldCard 是手持身份证合照，要求证件文本信息清晰可见，图片文件大小在 50K-5M 之间，不限制长宽，支持 png、bmp、gif、jpg、jpeg 格式。 请传入<a href="https://opendocs.alipay.com/apis/01ea4t">alipay.open.sp.image.upload</a>(图片上传接口) 返回的 image_id。 该资质是否必传请参见报名资质要求。。
	HandheldCard *string `json:"handheld_card,omitempty"`
	// ItemCode 是服务优选商品编码。
	ItemCode *string `json:"item_code,omitempty"`
	// MerchantPid 是商家pid。
	MerchantPid *string `json:"merchant_pid,omitempty"`
	// MiniAppid 是商家小程序appId。
	MiniAppid *string `json:"mini_appid,omitempty"`
	// OutBizNo 是外部业务号，需不重复。
	OutBizNo *string `json:"out_biz_no,omitempty"`
	// PromotorPid 是推广服务商(S2)pid。
	PromotorPid *string `json:"promotor_pid,omitempty"`
}

// SpImageUploadRequest 是支付宝官方 V3 OpenAPI 模型。
type SpImageUploadRequest struct {
	ImageContent *core.MultipartFile `json:"image_content,omitempty" multipart:"image_content,file"`
}
