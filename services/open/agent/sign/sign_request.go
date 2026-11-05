package sign

import (
	"github.com/tiyee/alipay-sdk-go/core"
)

// AgentCommonSignConfirmRequest 是支付宝官方 V3 OpenAPI 模型。
type AgentCommonSignConfirmRequest struct {
	// BatchNo 是ISV 代商户操作事务编号，通过事务开启接口alipay.open.agent.create调用返回。。
	BatchNo *string `json:"batch_no,omitempty"`
}

// AgentSignStatusQueryRequest 是支付宝官方 V3 OpenAPI 模型。
type AgentSignStatusQueryRequest struct {
	// Pid 是isv要查询签约状态的商户账号，可以是支付宝账号pid（2088开头），也可以是商户的登录账号（邮箱或手机号）。
	Pid *string `json:"pid,omitempty"`
	// ProductCodes 是isv要查询商户签约状态的产品码，产品码是支付宝内部对产品的唯一标识。目前支持查询产品及产品码详情参见 <a href="https://opendocs.alipay.com/isv/04eh3g">获取商家产品开通状态</a>。。
	ProductCodes *[]string `json:"product_codes,omitempty"`
}

// AgentCommonSignRequest 是支付宝官方 V3 OpenAPI 模型。
type AgentCommonSignRequest struct {
	AppAuthPic             *core.MultipartFile `json:"app_auth_pic,omitempty" multipart:"app_auth_pic,file"`
	AppDemo                *core.MultipartFile `json:"app_demo,omitempty" multipart:"app_demo,file"`
	AppHomeScreenshot      *core.MultipartFile `json:"app_home_screenshot,omitempty" multipart:"app_home_screenshot,file"`
	AppItemScreenshot      *core.MultipartFile `json:"app_item_screenshot,omitempty" multipart:"app_item_screenshot,file"`
	AppPayScreenshot       *core.MultipartFile `json:"app_pay_screenshot,omitempty" multipart:"app_pay_screenshot,file"`
	BusinessLicenseAuthPic *core.MultipartFile `json:"business_license_auth_pic,omitempty" multipart:"business_license_auth_pic,file"`
	BusinessLicensePic     *core.MultipartFile `json:"business_license_pic,omitempty" multipart:"business_license_pic,file"`
	Data                   *AgentCommonSign    `json:"data,omitempty" multipart:"data,json"`
	H5ExtraPic             *core.MultipartFile `json:"h_5_extra_pic,omitempty" multipart:"h_5_extra_pic,file"`
	H5HomeScreenshot       *core.MultipartFile `json:"h_5_home_screenshot,omitempty" multipart:"h_5_home_screenshot,file"`
	H5ItemScreenshot       *core.MultipartFile `json:"h_5_item_screenshot,omitempty" multipart:"h_5_item_screenshot,file"`
	H5PayScreenshot        *core.MultipartFile `json:"h_5_pay_screenshot,omitempty" multipart:"h_5_pay_screenshot,file"`
	H5SitesLoa             *core.MultipartFile `json:"h_5_sites_loa,omitempty" multipart:"h_5_sites_loa,file"`
	MiniAppScreenshot      *core.MultipartFile `json:"mini_app_screenshot,omitempty" multipart:"mini_app_screenshot,file"`
	ShopScenePic           *core.MultipartFile `json:"shop_scene_pic,omitempty" multipart:"shop_scene_pic,file"`
	ShopSignBoardPic       *core.MultipartFile `json:"shop_sign_board_pic,omitempty" multipart:"shop_sign_board_pic,file"`
	SpecialLicensePic      *core.MultipartFile `json:"special_license_pic,omitempty" multipart:"special_license_pic,file"`
	WebHomeScreenshot      *core.MultipartFile `json:"web_home_screenshot,omitempty" multipart:"web_home_screenshot,file"`
	WebItemScreenshot      *core.MultipartFile `json:"web_item_screenshot,omitempty" multipart:"web_item_screenshot,file"`
	WebPayScreenshot       *core.MultipartFile `json:"web_pay_screenshot,omitempty" multipart:"web_pay_screenshot,file"`
	WebSitesLoa            *core.MultipartFile `json:"web_sites_loa,omitempty" multipart:"web_sites_loa,file"`
}

// AgentFaceToFaceSignRequest 是支付宝官方 V3 OpenAPI 模型。
type AgentFaceToFaceSignRequest struct {
	BusinessLicenseAuthPic *core.MultipartFile  `json:"business_license_auth_pic,omitempty" multipart:"business_license_auth_pic,file"`
	BusinessLicensePic     *core.MultipartFile  `json:"business_license_pic,omitempty" multipart:"business_license_pic,file"`
	Data                   *AgentFaceToFaceSign `json:"data,omitempty" multipart:"data,json"`
	ShopScenePic           *core.MultipartFile  `json:"shop_scene_pic,omitempty" multipart:"shop_scene_pic,file"`
	ShopSignBoardPic       *core.MultipartFile  `json:"shop_sign_board_pic,omitempty" multipart:"shop_sign_board_pic,file"`
	SpecialLicensePic      *core.MultipartFile  `json:"special_license_pic,omitempty" multipart:"special_license_pic,file"`
}

// AgentMobilePaySignRequest 是支付宝官方 V3 OpenAPI 模型。
type AgentMobilePaySignRequest struct {
	AppAuthPic             *core.MultipartFile `json:"app_auth_pic,omitempty" multipart:"app_auth_pic,file"`
	AppDemo                *core.MultipartFile `json:"app_demo,omitempty" multipart:"app_demo,file"`
	AppHomeScreenshot      *core.MultipartFile `json:"app_home_screenshot,omitempty" multipart:"app_home_screenshot,file"`
	AppItemScreenshot      *core.MultipartFile `json:"app_item_screenshot,omitempty" multipart:"app_item_screenshot,file"`
	AppPayScreenshot       *core.MultipartFile `json:"app_pay_screenshot,omitempty" multipart:"app_pay_screenshot,file"`
	BusinessLicenseAuthPic *core.MultipartFile `json:"business_license_auth_pic,omitempty" multipart:"business_license_auth_pic,file"`
	BusinessLicensePic     *core.MultipartFile `json:"business_license_pic,omitempty" multipart:"business_license_pic,file"`
	Data                   *AgentMobilePaySign `json:"data,omitempty" multipart:"data,json"`
	HomeScreenshot         *core.MultipartFile `json:"home_screenshot,omitempty" multipart:"home_screenshot,file"`
	InAppScreenshot        *core.MultipartFile `json:"in_app_screenshot,omitempty" multipart:"in_app_screenshot,file"`
	PayScreenshot          *core.MultipartFile `json:"pay_screenshot,omitempty" multipart:"pay_screenshot,file"`
	SpecialLicensePic      *core.MultipartFile `json:"special_license_pic,omitempty" multipart:"special_license_pic,file"`
}

// AgentOfflinePaymentSignRequest 是支付宝官方 V3 OpenAPI 模型。
type AgentOfflinePaymentSignRequest struct {
	BusinessLicensePic *core.MultipartFile      `json:"business_license_pic,omitempty" multipart:"business_license_pic,file"`
	Data               *AgentOfflinePaymentSign `json:"data,omitempty" multipart:"data,json"`
	ShopScenePic       *core.MultipartFile      `json:"shop_scene_pic,omitempty" multipart:"shop_scene_pic,file"`
	ShopSignBoardPic   *core.MultipartFile      `json:"shop_sign_board_pic,omitempty" multipart:"shop_sign_board_pic,file"`
	SpecialLicensePic  *core.MultipartFile      `json:"special_license_pic,omitempty" multipart:"special_license_pic,file"`
}
