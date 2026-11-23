//go:build go1.25

package utils_test

import (
	"crypto/x509"
	"github.com/tiyee/alipay-sdk-go/core/auth"
	"github.com/tiyee/alipay-sdk-go/internal/testutil"
	"github.com/tiyee/alipay-sdk-go/utils"
	"os"
	"testing"
	"testing/synctest"
	"time"
)

func TestMixedRootBundleTrust(t *testing.T) {
	data, err := os.ReadFile("testdata/alipay-root.crt")
	if err != nil {
		t.Fatal(err)
	}
	roots, err := utils.ParseRootCertificatesPEM(string(data))
	if err != nil {
		t.Fatal(err)
	}

	// Verify against the original certificate dates even when this fixture ages out.
	synctest.Test(t, func(t *testing.T) {
		time.Sleep(time.Until(time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)))
		suite := testutil.NewCertSuite(t)
		config := auth.CertificateConfig{
			AppSigner: testutil.NewAppCryptoSigner(t), AppCertificate: suite.AppCert,
			AlipayCertificates: []*x509.Certificate{suite.AlipayCert}, RootCertificates: roots,
		}
		if _, err := auth.NewCertificateCredentials(config); err == nil {
			t.Fatal("untrusted leaf accepted")
		}
		config.RootCertificates = append(roots, suite.Root)
		if _, err := auth.NewCertificateCredentials(config); err != nil {
			t.Fatal(err)
		}
	})
}
