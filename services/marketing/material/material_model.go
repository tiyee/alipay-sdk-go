// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package material

// MarketingMaterialImageUpload 是支付宝官方 V3 OpenAPI 模型。
type MarketingMaterialImageUpload struct {
	BelongMerchantInfo *BelongMerchantInfo `json:"belong_merchant_info,omitempty"`
	// FileKey 是文件业务标识。		枚举值	alipay.marketing.activity.delivery.create接口中	delivery_base_info.delivery_material.delivery_single_material.delivery_image	当delivery_booth_code=PUBLIC_UNION，上传图片接口需指定file_key=PUBLIC_UNION_CHANNEL_PIC。上传图片尺寸600*600，支持格式：png、jpg、jpeg、bmp，大小不超过200kb；	当delivery_booth_code=PAYMENT_RESULT，上传图片接口需指定file_key=DELIVERY_CHANNEL_PIC。上传图片尺寸600*600，支持格式：png、jpg、jpeg、bmp，大小不超过200kb。	上传图片更多要求参考文档： https://render.alipay.com/p/c/18tpirlg12e8?operateFrom=BALIPAY		alipay.marketing.activity.ordervoucher.create接口中	voucher_display_info.brand_logo字段,file_key=PROMO_BRAND_LOGO，上传图片尺寸600*600，支持格式：png、jpg、jpeg、bmp，大小不超过2MB	voucher_display_info.voucher_image字段,file_key=PROMO_VOUCHER_IMAGE,上传图片尺寸600*600，支持格式：png、jpg、jpeg、bmp，大小不超过2MB
	FileKey *string `json:"file_key,omitempty"`
	// MerchantAccessMode 是商户接入模式
	MerchantAccessMode *string `json:"merchant_access_mode,omitempty"`
}

// BelongMerchantInfo 是支付宝官方 V3 OpenAPI 模型。
type BelongMerchantInfo struct {
	// BusinessType 是合作业务类型,默认为ISV_FOR_MERCHANT
	BusinessType *string `json:"business_type,omitempty"`
	// MerchantId 是商户PID,默认为当前接口调用商户。
	MerchantId *string `json:"merchant_id,omitempty"`
	// MerchantIdType 是商户id类型。特别说明：如果merchant_id_type选择SMID。则表明当前商家券活动的归属者是该SMID所绑定的支付宝PID。因此要求该SMID必须绑定相应的支付宝PID。
	MerchantIdType *string `json:"merchant_id_type,omitempty"`
}

// MaterialField 是支付宝官方 V3 OpenAPI 模型。
type MaterialField struct {
	// FieldName 是素材字段名称，来源于素材规范定义的字段名称。注意一次素材提报中字段名称不能重复。
	FieldName *string `json:"field_name,omitempty"`
	// FieldValue 是素材字段值，素材字段提报的实际值，支持多值，具体的要求请查看素材规范或对应的产品文档。
	FieldValue *[]string `json:"field_value,omitempty"`
}
