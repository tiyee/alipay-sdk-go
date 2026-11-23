package protocol

import "testing"

func TestCertParamsAreProtected(t *testing.T) {
	if !IsProtectedParam(ParamAppCertSN) {
		t.Error("app_cert_sn should be a protected parameter")
	}
	if !IsProtectedParam(ParamAlipayRootCertSN) {
		t.Error("alipay_root_cert_sn should be a protected parameter")
	}
}
