// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package category

// CategoryRequireInfo 是支付宝官方 V3 OpenAPI 模型。
type CategoryRequireInfo struct {
	// BusinessLicenceRequired 是营业执照页面是否必填。
	BusinessLicenceRequired *bool `json:"business_licence_required,omitempty"`
	// CategoryCode 是类目code（各级类目code下划线"_"拼接）。类目信息参考alipay.open.mini.category.query。
	CategoryCode *string `json:"category_code,omitempty"`
	// CategoryName 是类目名称（各级类目名称下划线"_"拼接）。
	CategoryName *string `json:"category_name,omitempty"`
	// CategoryRequirements 是类目要求原始描述信息。
	CategoryRequirements *string `json:"category_requirements,omitempty"`
	// DoorPhotoRequired 是门头照页面是否必填。
	DoorPhotoRequired *bool `json:"door_photo_required,omitempty"`
	// SpecialLicenceRequired 是特殊资质页面是否必填。
	SpecialLicenceRequired *bool `json:"special_licence_required,omitempty"`
}

// MiniAppCategory 是支付宝官方 V3 OpenAPI 模型。
type MiniAppCategory struct {
	// CategoryId 是类目id。
	CategoryId *string `json:"category_id,omitempty"`
	// CategoryName 是类目名称。
	CategoryName *string `json:"category_name,omitempty"`
	// HasChild 是是否有子类目。
	HasChild *bool `json:"has_child,omitempty"`
	// NeedLicense 是是否需要营业执照。
	NeedLicense *bool `json:"need_license,omitempty"`
	// NeedOutDoorPic 是是否需要门头照。
	NeedOutDoorPic *bool `json:"need_out_door_pic,omitempty"`
	// NeedSpecialLicense 是是否需要特许营业执照。
	NeedSpecialLicense *bool `json:"need_special_license,omitempty"`
	// ParentCategoryId 是父类目id。
	ParentCategoryId *string `json:"parent_category_id,omitempty"`
}
