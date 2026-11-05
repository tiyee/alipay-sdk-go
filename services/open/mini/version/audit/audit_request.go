package audit

import (
	"github.com/tiyee/alipay-sdk-go/core"
)

// MiniVersionAuditCancelRequest 是支付宝官方 V3 OpenAPI 模型。
type MiniVersionAuditCancelRequest struct {
	// AppVersion 是商家小程序审核中的版本号, 不传默认撤消正在审核中的版本。。
	AppVersion *string `json:"app_version,omitempty"`
	// BundleId 是小程序客户端类型，默认为支付宝端。常见支持如下客户端： com.alipay.alipaywallet：支付宝端； com.alibaba.android.rimet：DINGDING端； com.amap.app：高德端； com.alibaba.ailabs.genie.webapps：天猫精灵端； com.alipay.iot.xpaas：支付宝IoT端。 如需更多端投放，请联系业务BD。。
	BundleId *string `json:"bundle_id,omitempty"`
}

// MiniVersionAuditApplyRequest 是支付宝官方 V3 OpenAPI 模型。
type MiniVersionAuditApplyRequest struct {
	AppLogo                  *core.MultipartFile    `json:"app_logo,omitempty" multipart:"app_logo,file"`
	Data                     *MiniVersionAuditApply `json:"data,omitempty" multipart:"data,json"`
	EighthSpecialLicensePic  *core.MultipartFile    `json:"eighth_special_license_pic,omitempty" multipart:"eighth_special_license_pic,file"`
	FifthLicensePic          *core.MultipartFile    `json:"fifth_license_pic,omitempty" multipart:"fifth_license_pic,file"`
	FifthScreenShot          *core.MultipartFile    `json:"fifth_screen_shot,omitempty" multipart:"fifth_screen_shot,file"`
	FifthSpecialLicensePic   *core.MultipartFile    `json:"fifth_special_license_pic,omitempty" multipart:"fifth_special_license_pic,file"`
	FirstLicensePic          *core.MultipartFile    `json:"first_license_pic,omitempty" multipart:"first_license_pic,file"`
	FirstScreenShot          *core.MultipartFile    `json:"first_screen_shot,omitempty" multipart:"first_screen_shot,file"`
	FirstSpecialLicensePic   *core.MultipartFile    `json:"first_special_license_pic,omitempty" multipart:"first_special_license_pic,file"`
	FourthLicensePic         *core.MultipartFile    `json:"fourth_license_pic,omitempty" multipart:"fourth_license_pic,file"`
	FourthScreenShot         *core.MultipartFile    `json:"fourth_screen_shot,omitempty" multipart:"fourth_screen_shot,file"`
	FourthSpecialLicensePic  *core.MultipartFile    `json:"fourth_special_license_pic,omitempty" multipart:"fourth_special_license_pic,file"`
	NinthSpecialLicensePic   *core.MultipartFile    `json:"ninth_special_license_pic,omitempty" multipart:"ninth_special_license_pic,file"`
	OutDoorPic               *core.MultipartFile    `json:"out_door_pic,omitempty" multipart:"out_door_pic,file"`
	SecondLicensePic         *core.MultipartFile    `json:"second_license_pic,omitempty" multipart:"second_license_pic,file"`
	SecondScreenShot         *core.MultipartFile    `json:"second_screen_shot,omitempty" multipart:"second_screen_shot,file"`
	SecondSpecialLicensePic  *core.MultipartFile    `json:"second_special_license_pic,omitempty" multipart:"second_special_license_pic,file"`
	SeventhSpecialLicensePic *core.MultipartFile    `json:"seventh_special_license_pic,omitempty" multipart:"seventh_special_license_pic,file"`
	SixthSpecialLicensePic   *core.MultipartFile    `json:"sixth_special_license_pic,omitempty" multipart:"sixth_special_license_pic,file"`
	TenthSpecialLicensePic   *core.MultipartFile    `json:"tenth_special_license_pic,omitempty" multipart:"tenth_special_license_pic,file"`
	TestFileName             *core.MultipartFile    `json:"test_file_name,omitempty" multipart:"test_file_name,file"`
	ThirdLicensePic          *core.MultipartFile    `json:"third_license_pic,omitempty" multipart:"third_license_pic,file"`
	ThirdScreenShot          *core.MultipartFile    `json:"third_screen_shot,omitempty" multipart:"third_screen_shot,file"`
	ThirdSpecialLicensePic   *core.MultipartFile    `json:"third_special_license_pic,omitempty" multipart:"third_special_license_pic,file"`
}
