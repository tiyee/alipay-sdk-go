// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package info

// PublicInfoModifyRequest 是支付宝官方 V3 OpenAPI 模型。
type PublicInfoModifyRequest struct {
	// AppName 是生活号名称，2-20个字之间。注意： * 不得含有违反法律法规和公序良俗的相关信息； * 不得侵害他人名誉权、知识产权、商业秘密等合法权利； * 不得以太过广泛的、或产品、行业词组来命名，如：女装、皮革批发； * 不得以实名认证的媒体资质账号创建服务窗，或媒体相关名称命名服务窗，如：XX电视台、XX杂志等。。
	AppName *string `json:"app_name,omitempty"`
	// AuthPic 是授权运营书图片地址，企业商户若为被经营方授权，需上传加盖公章的扫描件。支持 .jpg、 .jpeg、 .png 格式，需小于1M。使用 <a href="https://opendocs.alipay.com/apis/api_3/alipay.offline.material.image.upload">alipay.offline.material.image.upload</a>(上传门店照片和视频接口)上传图片后，将得到的image_url回填与此处。。
	AuthPic *string `json:"auth_pic,omitempty"`
	// BackgroundUrl 是背景图片地址，建议尺寸 1600 x 1000px，支持.jpg .jpeg .png格式，小于1M。使用 <a href="https://opendocs.alipay.com/apis/api_3/alipay.offline.material.image.upload">alipay.offline.material.image.upload</a>(上传门店照片和视频接口)上传图片后，将得到的image_url回填与此处。。
	BackgroundUrl *string `json:"background_url,omitempty"`
	// Introduction 是简介。
	Introduction *string `json:"introduction,omitempty"`
	// LicenseUrl 是营业执照地址，建议尺寸 320*320 px，支持 .jpg、 .jpeg、 .png 格式，需小于1M。 使用 <a href="https://opendocs.alipay.com/apis/api_3/alipay.offline.material.image.upload">alipay.offline.material.image.upload</a>(上传门店照片和视频接口)上传图片后，将得到的image_url回填与此处。。
	LicenseUrl *string `json:"license_url,omitempty"`
	// LogoUrl 是生活号头像地址，建议尺寸 320*320 px，支持 .jpg、 .jpeg、 .png 格式，需小于1M。 使用 <a href="https://opendocs.alipay.com/apis/api_3/alipay.offline.material.image.upload">alipay.offline.material.image.upload</a>(上传门店照片和视频接口)上传图片后，将得到的image_url回填与此处。。
	LogoUrl *string `json:"logo_url,omitempty"`
	// NotifyUrl 是通知地址。
	NotifyUrl *string `json:"notify_url,omitempty"`
	// PublicGreeting 是欢迎语。
	PublicGreeting *string `json:"public_greeting,omitempty"`
	// ShopPics 是门店照片地址，支持 .jpg、 .jpeg、 .png 格式，需小于1M。 使用 <a href="https://opendocs.alipay.com/apis/api_3/alipay.offline.material.image.upload">alipay.offline.material.image.upload</a>(上传门店照片和视频接口)上传图片后，将得到的image_url回填与此处。。
	ShopPics *[]string `json:"shop_pics,omitempty"`
}

// PublicInfoQueryRequest is the request type for alipay.open.public.info.query.
type PublicInfoQueryRequest struct{}
