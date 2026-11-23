package generatedapi_test

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/tiyee/alipay-sdk-go/core"
	"github.com/tiyee/alipay-sdk-go/core/option"
	"github.com/tiyee/alipay-sdk-go/internal/testutil"
	antmerchant "github.com/tiyee/alipay-sdk-go/services/ant/merchant"
	commerceoperation "github.com/tiyee/alipay-sdk-go/services/commerce/operation"
	commercetransport "github.com/tiyee/alipay-sdk-go/services/commerce/transport"
	ebppbill "github.com/tiyee/alipay-sdk-go/services/ebpp/bill"
	ebppinvoiceenterprise "github.com/tiyee/alipay-sdk-go/services/ebpp/invoice/enterprise"
	ebppinvoiceinstitution "github.com/tiyee/alipay-sdk-go/services/ebpp/invoice/institution"
	ebppinvoicetaxno "github.com/tiyee/alipay-sdk-go/services/ebpp/invoice/taxno"
	ecocontract "github.com/tiyee/alipay-sdk-go/services/eco/contract"
	ecodoc "github.com/tiyee/alipay-sdk-go/services/eco/doc"
	ecosign "github.com/tiyee/alipay-sdk-go/services/eco/sign"
	fundenterprisepay "github.com/tiyee/alipay-sdk-go/services/fund/enterprisepay"
	fundusertrade "github.com/tiyee/alipay-sdk-go/services/fund/usertrade"
	iserviceccmagent "github.com/tiyee/alipay-sdk-go/services/iservice/ccm/agent"
	iserviceccmrole "github.com/tiyee/alipay-sdk-go/services/iservice/ccm/role"
	iserviceccmswarticle "github.com/tiyee/alipay-sdk-go/services/iservice/ccm/sw/article"
	iserviceccmswlibrary "github.com/tiyee/alipay-sdk-go/services/iservice/ccm/sw/library"
	iserviceccmswtree "github.com/tiyee/alipay-sdk-go/services/iservice/ccm/sw/tree"
	iserviceccmswtreecategory "github.com/tiyee/alipay-sdk-go/services/iservice/ccm/sw/treecategory"
	marketingactivityvoucherpackage "github.com/tiyee/alipay-sdk-go/services/marketing/activity/voucherpackage"
	marketingcampaign "github.com/tiyee/alipay-sdk-go/services/marketing/campaign"
	marketingcard "github.com/tiyee/alipay-sdk-go/services/marketing/card"
	marketingcardbenefit "github.com/tiyee/alipay-sdk-go/services/marketing/card/benefit"
	mobilepublictemplate "github.com/tiyee/alipay-sdk-go/services/mobile/public/template"
	openappmembers "github.com/tiyee/alipay-sdk-go/services/open/app/members"
	openappservice "github.com/tiyee/alipay-sdk-go/services/open/app/service"
	openminiinner "github.com/tiyee/alipay-sdk-go/services/open/mini/inner"
	openminiinnerversion "github.com/tiyee/alipay-sdk-go/services/open/mini/innerversion"
	openminiminiapp "github.com/tiyee/alipay-sdk-go/services/open/mini/miniapp"
	openminisafedomain "github.com/tiyee/alipay-sdk-go/services/open/mini/safedomain"
	openminitemplatemsg "github.com/tiyee/alipay-sdk-go/services/open/mini/templatemsg"
	openpublicaccount "github.com/tiyee/alipay-sdk-go/services/open/public/account"
	openpublicadvert "github.com/tiyee/alipay-sdk-go/services/open/public/advert"
	openpubliccontact "github.com/tiyee/alipay-sdk-go/services/open/public/contact"
	openpublicgroup "github.com/tiyee/alipay-sdk-go/services/open/public/group"
	openpubliclife "github.com/tiyee/alipay-sdk-go/services/open/public/life"
	openpublicmenu "github.com/tiyee/alipay-sdk-go/services/open/public/menu"
	openpublicpayee "github.com/tiyee/alipay-sdk-go/services/open/public/payee"
	openpublicpersonalized "github.com/tiyee/alipay-sdk-go/services/open/public/personalized"
	openpublictopic "github.com/tiyee/alipay-sdk-go/services/open/public/topic"
	security "github.com/tiyee/alipay-sdk-go/services/security"
	userinfo "github.com/tiyee/alipay-sdk-go/services/user/info"
	useruserinfo "github.com/tiyee/alipay-sdk-go/services/user/userinfo"
)

// These non-GET operations declare query parameters or omit the request body in Alipay OpenAPI 2026-09-23.
// Alipay documents how to obtain the OpenAPI description at https://opendocs.alipay.com/open-v3/065bsc
func TestNonGETParameterLocations(t *testing.T) {
	value, bodyValue := "a/b+&", "body"
	t.Run("alipay.commerce.operation.activity.merchant.unsign", func(t *testing.T) {
		verifyRequest(t, commerceoperation.NewService, (*commerceoperation.Service).OperationActivityMerchantUnsign, commerceoperation.OperationActivityMerchantUnsignRequest{},
			"POST", "/v3/alipay/commerce/operation/activity/merchant/unsign", "", ``)
	})
	t.Run("alipay.commerce.transport.charger.chargerbindinfo.sync", func(t *testing.T) {
		verifyRequest(t, commercetransport.NewService, (*commercetransport.Service).TransportChargerChargerBindInfoSync, commercetransport.TransportChargerChargerBindInfoSyncRequest{AuthToken: &value, BindQrcode: &bodyValue},
			"POST", "/v3/alipay/commerce/transport/charger/chargerbindinfo/sync", "auth_token", `{"bind_qrcode":"body"}`)
	})
	t.Run("alipay.ebpp.bill.add", func(t *testing.T) {
		verifyRequest(t, ebppbill.NewService, (*ebppbill.Service).BillAdd, ebppbill.BillAddRequest{AuthToken: &value, BankBillNo: &bodyValue},
			"POST", "/v3/alipay/ebpp/bill/add", "auth_token", `{"bank_bill_no":"body"}`)
	})
	t.Run("alipay.ebpp.invoice.enterprise.merchantrelation.create", func(t *testing.T) {
		verifyRequest(t, ebppinvoiceenterprise.NewService, (*ebppinvoiceenterprise.Service).InvoiceEnterpriseMerchantRelationCreate, ebppinvoiceenterprise.InvoiceEnterpriseMerchantRelationCreateRequest{EnterpriseId: &value, AccountId: &bodyValue},
			"POST", "/v3/alipay/ebpp/invoice/enterprise/merchantrelation/create", "enterprise_id", `{"account_id":"body"}`)
	})
	t.Run("alipay.ebpp.invoice.institution.delete", func(t *testing.T) {
		verifyRequest(t, ebppinvoiceinstitution.NewService, (*ebppinvoiceinstitution.Service).InvoiceInstitutionDelete, ebppinvoiceinstitution.InvoiceInstitutionDeleteRequest{AccountId: &value, AgreementNo: &value, EnterpriseId: &value, InstitutionId: &value},
			"DELETE", "/v3/alipay/ebpp/invoice/institution/delete", "account_id agreement_no enterprise_id institution_id", ``)
	})
	t.Run("alipay.ebpp.invoice.taxno.batchquery", func(t *testing.T) {
		verifyRequest(t, ebppinvoicetaxno.NewService, (*ebppinvoicetaxno.Service).InvoiceTaxNoBatchQuery, ebppinvoicetaxno.InvoiceTaxNoBatchQueryRequest{AuthToken: &value, EnableTradeOut: &bodyValue},
			"POST", "/v3/alipay/ebpp/invoice/taxno/batchquery", "auth_token", `{"enable_trade_out":"body"}`)
	})
	t.Run("alipay.eco.contract.signflows.create", func(t *testing.T) {
		verifyRequest(t, ecocontract.NewService, (*ecocontract.Service).ContractSignFlowsCreate, ecocontract.ContractSignFlowsCreateRequest{TargetAppId: &value, BusinessScene: &bodyValue},
			"POST", "/v3/alipay/eco/contract/signflows/create", "target_app_id", `{"business_scene":"body"}`)
	})
	t.Run("alipay.eco.doc.template.create", func(t *testing.T) {
		verifyRequest(t, ecodoc.NewService, (*ecodoc.Service).DocTemplateCreate, ecodoc.DocTemplateCreateRequest{TargetAppId: &value, ContentMd5: &bodyValue},
			"POST", "/v3/alipay/eco/doc/template/create", "target_app_id", `{"content_md_5":"body"}`)
	})
	t.Run("alipay.eco.sign.flow.cancel", func(t *testing.T) {
		verifyRequest(t, ecosign.NewService, (*ecosign.Service).SignFlowCancel, ecosign.SignFlowCancelRequest{TargetAppId: &value, FlowId: &bodyValue},
			"POST", "/v3/alipay/eco/sign/flow/cancel", "target_app_id", `{"flow_id":"body"}`)
	})
	t.Run("alipay.eco.sign.flow.create", func(t *testing.T) {
		verifyRequest(t, ecosign.NewService, (*ecosign.Service).SignFlowCreate, ecosign.SignFlowCreateRequest{TargetAppId: &value, BusinessScene: &bodyValue},
			"POST", "/v3/alipay/eco/sign/flow/create", "target_app_id", `{"business_scene":"body"}`)
	})
	t.Run("alipay.eco.sign.flow.finish", func(t *testing.T) {
		verifyRequest(t, ecosign.NewService, (*ecosign.Service).SignFlowFinish, ecosign.SignFlowFinishRequest{TargetAppId: &value, FlowId: &bodyValue},
			"POST", "/v3/alipay/eco/sign/flow/finish", "target_app_id", `{"flow_id":"body"}`)
	})
	t.Run("alipay.fund.enterprisepay.group.delete", func(t *testing.T) {
		verifyRequest(t, fundenterprisepay.NewService, (*fundenterprisepay.Service).EnterprisePayGroupDelete, fundenterprisepay.FundEnterprisePayGroupDeleteRequest{ProductCode: &value, BizScene: &value, AccountId: &value, AgreementNo: &value, OutBizNo: &value},
			"DELETE", "/v3/alipay/fund/enterprisepay/group/delete", "account_id agreement_no biz_scene out_biz_no product_code", ``)
	})
	t.Run("alipay.fund.usertrade.batchquery", func(t *testing.T) {
		verifyRequest(t, fundusertrade.NewService, (*fundusertrade.Service).UserTradeBatchQuery, fundusertrade.FundUserTradeBatchQueryRequest{AuthToken: &value, BizScene: &bodyValue},
			"POST", "/v3/alipay/fund/usertrade/batchquery", "auth_token", `{"biz_scene":"body"}`)
	})
	t.Run("alipay.iservice.ccm.agent.delete", func(t *testing.T) {
		verifyRequest(t, iserviceccmagent.NewService, (*iserviceccmagent.Service).CcmAgentDelete, iserviceccmagent.CcmAgentDeleteRequest{Id: &value},
			"DELETE", "/v3/alipay/iservice/ccm/agent/delete", "id", ``)
	})
	t.Run("alipay.iservice.ccm.role.delete", func(t *testing.T) {
		verifyRequest(t, iserviceccmrole.NewService, (*iserviceccmrole.Service).CcmRoleDelete, iserviceccmrole.CcmRoleDeleteRequest{Id: &value, CcsInstanceId: &value},
			"DELETE", "/v3/alipay/iservice/ccm/role/delete", "ccs_instance_id id", ``)
	})
	t.Run("alipay.iservice.ccm.sw.article.delete", func(t *testing.T) {
		verifyRequest(t, iserviceccmswarticle.NewService, (*iserviceccmswarticle.Service).CcmSwArticleDelete, iserviceccmswarticle.CcmSwArticleDeleteRequest{},
			"DELETE", "/v3/alipay/iservice/ccm/sw/article/delete", "", ``)
	})
	t.Run("alipay.iservice.ccm.sw.library.delete", func(t *testing.T) {
		verifyRequest(t, iserviceccmswlibrary.NewService, (*iserviceccmswlibrary.Service).CcmSwLibraryDelete, iserviceccmswlibrary.CcmSwLibraryDeleteRequest{},
			"DELETE", "/v3/alipay/iservice/ccm/sw/library/delete", "", ``)
	})
	t.Run("alipay.iservice.ccm.sw.tree.delete", func(t *testing.T) {
		verifyRequest(t, iserviceccmswtree.NewService, (*iserviceccmswtree.Service).CcmSwTreeDelete, iserviceccmswtree.CcmSwTreeDeleteRequest{},
			"DELETE", "/v3/alipay/iservice/ccm/sw/tree/delete", "", ``)
	})
	t.Run("alipay.iservice.ccm.sw.treecategory.delete", func(t *testing.T) {
		verifyRequest(t, iserviceccmswtreecategory.NewService, (*iserviceccmswtreecategory.Service).CcmSwTreeCategoryDelete, iserviceccmswtreecategory.CcmSwTreeCategoryDeleteRequest{},
			"DELETE", "/v3/alipay/iservice/ccm/sw/treecategory/delete", "", ``)
	})
	t.Run("alipay.marketing.activity.voucherpackage.consult", func(t *testing.T) {
		verifyRequest(t, marketingactivityvoucherpackage.NewService, (*marketingactivityvoucherpackage.Service).ActivityVoucherPackageConsult, marketingactivityvoucherpackage.ActivityVoucherPackageConsultRequest{AuthToken: &value},
			"POST", "/v3/alipay/marketing/activity/voucherpackage/consult", "auth_token", `{}`)
	})
	t.Run("alipay.marketing.campaign.order.voucher.consult", func(t *testing.T) {
		verifyRequest(t, marketingcampaign.NewService, (*marketingcampaign.Service).CampaignOrderVoucherConsult, marketingcampaign.CampaignOrderVoucherConsultRequest{AuthToken: &value, BusinessParam: &bodyValue},
			"POST", "/v3/alipay/marketing/campaign/order/voucher/consult", "auth_token", `{"business_param":"body"}`)
	})
	t.Run("alipay.marketing.card.benefit.delete", func(t *testing.T) {
		verifyRequest(t, marketingcardbenefit.NewService, (*marketingcardbenefit.Service).CardBenefitDelete, marketingcardbenefit.CardBenefitDeleteRequest{TemplateId: &value, BenefitId: &value},
			"DELETE", "/v3/alipay/marketing/card/benefit/delete", "benefit_id template_id", ``)
	})
	t.Run("alipay.marketing.card.delete", func(t *testing.T) {
		verifyRequest(t, marketingcard.NewService, (*marketingcard.Service).CardDelete, marketingcard.CardDeleteRequest{OutSerialNo: &value, TargetCardNo: &value, TargetCardNoType: &value, ReasonCode: &value, ExtInfo: &value},
			"DELETE", "/v3/alipay/marketing/card/delete", "ext_info out_serial_no reason_code target_card_no target_card_no_type", ``)
	})
	t.Run("alipay.marketing.card.open", func(t *testing.T) {
		verifyRequest(t, marketingcard.NewService, (*marketingcard.Service).CardOpen, marketingcard.CardOpenRequest{AuthToken: &value, CardTemplateId: &bodyValue},
			"POST", "/v3/alipay/marketing/card/open", "auth_token", `{"card_template_id":"body"}`)
	})
	t.Run("alipay.mobile.public.template.message.query", func(t *testing.T) {
		verifyRequest(t, mobilepublictemplate.NewService, (*mobilepublictemplate.Service).PublicTemplateMessageQuery, mobilepublictemplate.MobilePublicTemplateMessageQueryRequest{AuthToken: &value, Template: &bodyValue},
			"POST", "/v3/alipay/mobile/public/template/message/query", "auth_token", `{"template":"body"}`)
	})
	t.Run("alipay.open.app.members.delete", func(t *testing.T) {
		verifyRequest(t, openappmembers.NewService, (*openappmembers.Service).AppMembersDelete, openappmembers.AppMembersDeleteRequest{UserId: &value, OpenId: &value, Role: &value},
			"DELETE", "/v3/alipay/open/app/members/delete", "open_id role user_id", ``)
	})
	t.Run("alipay.open.app.service.delete", func(t *testing.T) {
		verifyRequest(t, openappservice.NewService, (*openappservice.Service).AppServiceDelete, openappservice.AppServiceDeleteRequest{ServiceCode: &value},
			"DELETE", "/v3/alipay/open/app/service/delete", "service_code", ``)
	})
	t.Run("alipay.open.mini.inner.members.delete", func(t *testing.T) {
		verifyRequest(t, openminiinner.NewService, (*openminiinner.Service).MiniInnerMembersDelete, openminiinner.MiniInnerMembersDeleteRequest{MiniAppId: &value, AppOrigin: &value, DomainAccount: &value, LoginId: &value, Type: &value},
			"DELETE", "/v3/alipay/open/mini/inner/members/delete", "app_origin domain_account login_id mini_app_id type", ``)
	})
	t.Run("alipay.open.mini.innerversion.info.delete", func(t *testing.T) {
		verifyRequest(t, openminiinnerversion.NewService, (*openminiinnerversion.Service).MiniInnerVersionInfoDelete, openminiinnerversion.MiniInnerVersionInfoDeleteRequest{MiniAppId: &value, AppOrigin: &value, BundleId: &value, AppVersion: &value, Pid: &value, InstCode: &value},
			"DELETE", "/v3/alipay/open/mini/innerversion/info/delete", "app_origin app_version bundle_id inst_code mini_app_id pid", ``)
	})
	t.Run("alipay.open.mini.miniapp.brand.cancel", func(t *testing.T) {
		verifyRequest(t, openminiminiapp.NewService, (*openminiminiapp.Service).MiniMiniAppBrandCancel, openminiminiapp.MiniMiniAppBrandCancelRequest{},
			"POST", "/v3/alipay/open/mini/miniapp/brand/cancel", "", ``)
	})
	t.Run("alipay.open.mini.safedomain.delete", func(t *testing.T) {
		verifyRequest(t, openminisafedomain.NewService, (*openminisafedomain.Service).MiniSafeDomainDelete, openminisafedomain.MiniSafeDomainDeleteRequest{SafeDomain: &value},
			"DELETE", "/v3/alipay/open/mini/safedomain/delete", "safe_domain", ``)
	})
	t.Run("alipay.open.mini.templatemsg.tinypayswitch.confirm", func(t *testing.T) {
		verifyRequest(t, openminitemplatemsg.NewService, (*openminitemplatemsg.Service).MiniTemplateMsgTinyPaySwitchConfirm, openminitemplatemsg.MiniTemplateMsgTinyPaySwitchConfirmRequest{},
			"POST", "/v3/alipay/open/mini/templatemsg/tinypayswitch/confirm", "", ``)
	})
	t.Run("alipay.open.public.account.delete", func(t *testing.T) {
		verifyRequest(t, openpublicaccount.NewService, (*openpublicaccount.Service).PublicAccountDelete, openpublicaccount.PublicAccountDeleteRequest{AgreementId: &value, BindAccountNo: &value, FromUserId: &value, OpenId: &value},
			"DELETE", "/v3/alipay/open/public/account/delete", "agreement_id bind_account_no from_user_id open_id", ``)
	})
	t.Run("alipay.open.public.advert.batchquery", func(t *testing.T) {
		verifyRequest(t, openpublicadvert.NewService, (*openpublicadvert.Service).PublicAdvertBatchQuery, openpublicadvert.PublicAdvertBatchQueryRequest{},
			"POST", "/v3/alipay/open/public/advert/batchquery", "", ``)
	})
	t.Run("alipay.open.public.advert.delete", func(t *testing.T) {
		verifyRequest(t, openpublicadvert.NewService, (*openpublicadvert.Service).PublicAdvertDelete, openpublicadvert.PublicAdvertDeleteRequest{AdvertId: &value, AdvertGroup: &value},
			"DELETE", "/v3/alipay/open/public/advert/delete", "advert_group advert_id", ``)
	})
	t.Run("alipay.open.public.contact.follow.batchquery", func(t *testing.T) {
		verifyRequest(t, openpubliccontact.NewService, (*openpubliccontact.Service).PublicContactFollowBatchQuery, openpubliccontact.PublicContactFollowBatchQueryRequest{AuthToken: &value},
			"POST", "/v3/alipay/open/public/contact/follow/batchquery", "auth_token", ``)
	})
	t.Run("alipay.open.public.group.batchquery", func(t *testing.T) {
		verifyRequest(t, openpublicgroup.NewService, (*openpublicgroup.Service).PublicGroupBatchQuery, openpublicgroup.PublicGroupBatchQueryRequest{},
			"POST", "/v3/alipay/open/public/group/batchquery", "", ``)
	})
	t.Run("alipay.open.public.group.delete", func(t *testing.T) {
		verifyRequest(t, openpublicgroup.NewService, (*openpublicgroup.Service).PublicGroupDelete, openpublicgroup.PublicGroupDeleteRequest{GroupId: &value},
			"DELETE", "/v3/alipay/open/public/group/delete", "group_id", ``)
	})
	t.Run("alipay.open.public.life.aboard.apply", func(t *testing.T) {
		verifyRequest(t, openpubliclife.NewService, (*openpubliclife.Service).PublicLifeAboardApply, openpubliclife.PublicLifeAboardApplyRequest{},
			"POST", "/v3/alipay/open/public/life/aboard/apply", "", ``)
	})
	t.Run("alipay.open.public.life.debark.apply", func(t *testing.T) {
		verifyRequest(t, openpubliclife.NewService, (*openpubliclife.Service).PublicLifeDebarkApply, openpubliclife.PublicLifeDebarkApplyRequest{},
			"POST", "/v3/alipay/open/public/life/debark/apply", "", ``)
	})
	t.Run("alipay.open.public.life.label.batchquery", func(t *testing.T) {
		verifyRequest(t, openpubliclife.NewService, (*openpubliclife.Service).PublicLifeLabelBatchQuery, openpubliclife.PublicLifeLabelBatchQueryRequest{},
			"POST", "/v3/alipay/open/public/life/label/batchquery", "", ``)
	})
	t.Run("alipay.open.public.life.label.delete", func(t *testing.T) {
		verifyRequest(t, openpubliclife.NewService, (*openpubliclife.Service).PublicLifeLabelDelete, openpubliclife.PublicLifeLabelDeleteRequest{LabelId: &value},
			"DELETE", "/v3/alipay/open/public/life/label/delete", "label_id", ``)
	})
	t.Run("alipay.open.public.menu.batchquery", func(t *testing.T) {
		verifyRequest(t, openpublicmenu.NewService, (*openpublicmenu.Service).PublicMenuBatchQuery, openpublicmenu.PublicMenuBatchQueryRequest{},
			"POST", "/v3/alipay/open/public/menu/batchquery", "", ``)
	})
	t.Run("alipay.open.public.payee.bind.delete", func(t *testing.T) {
		verifyRequest(t, openpublicpayee.NewService, (*openpublicpayee.Service).PublicPayeeBindDelete, openpublicpayee.PublicPayeeBindDeleteRequest{LoginId: &value, Pid: &value},
			"DELETE", "/v3/alipay/open/public/payee/bind/delete", "login_id pid", ``)
	})
	t.Run("alipay.open.public.personalized.extension.batchquery", func(t *testing.T) {
		verifyRequest(t, openpublicpersonalized.NewService, (*openpublicpersonalized.Service).PublicPersonalizedExtensionBatchQuery, openpublicpersonalized.PublicPersonalizedExtensionBatchQueryRequest{},
			"POST", "/v3/alipay/open/public/personalized/extension/batchquery", "", ``)
	})
	t.Run("alipay.open.public.personalized.extension.delete", func(t *testing.T) {
		verifyRequest(t, openpublicpersonalized.NewService, (*openpublicpersonalized.Service).PublicPersonalizedExtensionDelete, openpublicpersonalized.PublicPersonalizedExtensionDeleteRequest{ExtensionKey: &value},
			"DELETE", "/v3/alipay/open/public/personalized/extension/delete", "extension_key", ``)
	})
	t.Run("alipay.open.public.personalized.menu.delete", func(t *testing.T) {
		verifyRequest(t, openpublicpersonalized.NewService, (*openpublicpersonalized.Service).PublicPersonalizedMenuDelete, openpublicpersonalized.PublicPersonalizedMenuDeleteRequest{MenuKey: &value},
			"DELETE", "/v3/alipay/open/public/personalized/menu/delete", "menu_key", ``)
	})
	t.Run("alipay.open.public.topic.batchquery", func(t *testing.T) {
		verifyRequest(t, openpublictopic.NewService, (*openpublictopic.Service).PublicTopicBatchQuery, openpublictopic.PublicTopicBatchQueryRequest{},
			"POST", "/v3/alipay/open/public/topic/batchquery", "", ``)
	})
	t.Run("alipay.open.public.topic.delete", func(t *testing.T) {
		verifyRequest(t, openpublictopic.NewService, (*openpublictopic.Service).PublicTopicDelete, openpublictopic.PublicTopicDeleteRequest{TopicId: &value},
			"DELETE", "/v3/alipay/open/public/topic/delete", "topic_id", ``)
	})
	t.Run("alipay.security.prod.sss.query", func(t *testing.T) {
		verifyRequest(t, security.NewService, (*security.Service).ProdSssQuery, security.SecurityProdSssQueryRequest{Aaa: &value, AaaOpenId: &value, Bbb: "segment/1"},
			"POST", "/v3/alipay/security/prod/sss/query/segment%2F1", "aaa aaa_open_id", `{}`)
	})
	t.Run("alipay.user.info.share", func(t *testing.T) {
		verifyRequest(t, userinfo.NewService, (*userinfo.Service).InfoShare, userinfo.UserInfoShareRequest{AuthToken: &value},
			"POST", "/v3/alipay/user/info/share", "auth_token", ``)
	})
	t.Run("alipay.user.userinfo.share", func(t *testing.T) {
		verifyRequest(t, useruserinfo.NewService, (*useruserinfo.Service).UserInfoShare, useruserinfo.UserUserInfoShareRequest{},
			"POST", "/v3/alipay/user/userinfo/share", "", ``)
	})
	t.Run("ant.merchant.expand.apprecommend.account.delete", func(t *testing.T) {
		verifyRequest(t, antmerchant.NewService, (*antmerchant.Service).ExpandAppRecommendAccountDelete, antmerchant.ExpandAppRecommendAccountDeleteRequest{AppNo: &value, AccNo: &value},
			"DELETE", "/v3/ant/merchant/expand/apprecommend/account/delete", "acc_no app_no", ``)
	})
	t.Run("ant.merchant.expand.item.delete", func(t *testing.T) {
		verifyRequest(t, antmerchant.NewService, (*antmerchant.Service).ExpandItemDelete, antmerchant.ExpandItemDeleteRequest{ItemId: &value, ExternalItemId: &value},
			"DELETE", "/v3/ant/merchant/expand/item/delete", "external_item_id item_id", ``)
	})
	t.Run("ant.merchant.expand.item.open.delete", func(t *testing.T) {
		verifyRequest(t, antmerchant.NewService, (*antmerchant.Service).ExpandItemOpenDelete, antmerchant.ExpandItemOpenDeleteRequest{ItemId: &value},
			"DELETE", "/v3/ant/merchant/expand/item/open/delete", "item_id", ``)
	})
}

func verifyRequest[S, Req, Resp any](t *testing.T, newService func(core.Executor) (*S, error), call func(*S, context.Context, Req, ...core.CallOption) (*core.Response[Resp], error), req Req, method, path, queryKeys, wantBody string) {
	t.Helper()
	wantQuery := url.Values{}
	for _, key := range strings.Fields(queryKeys) {
		wantQuery.Set(key, "a/b+&")
	}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Error(err)
			return
		}
		if r.Method != method || r.URL.EscapedPath() != path || r.URL.RawQuery != wantQuery.Encode() || string(body) != wantBody {
			t.Errorf("got %s %s with body %q, want %s %s?%s with body %q", r.Method, r.RequestURI, body, method, path, wantQuery.Encode(), wantBody)
		}
		testutil.VerifyV3Authorization(t, r, body, "")
		response := []byte("{}")
		for k, v := range testutil.V3ResponseHeaders(t, response, "1764000000000", "nonce", "") {
			w.Header()[k] = v
		}
		_, _ = w.Write(response)
	}))
	defer server.Close()
	client, err := core.NewClient(testutil.NewCredentials(t), option.WithGateway(server.URL), option.WithAppID("offline"))
	if err != nil {
		t.Fatal(err)
	}
	svc, err := newService(client)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := call(svc, t.Context(), req); err != nil {
		t.Fatal(err)
	}
}
