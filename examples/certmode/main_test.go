package main

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/tiyee/alipay-sdk-go/internal/testutil"
)

func TestCertificateFilesLoadPlatformChain(t *testing.T) {
	suite := testutil.NewCertSuite(t)
	leaf, intermediate := suite.NewIntermediateChain(t)
	dir := t.TempDir()
	first, second := filepath.Join(dir, "platform.crt"), filepath.Join(dir, "rotated.crt")
	if err := os.WriteFile(first, []byte(testutil.CertToPEM(t, leaf)+testutil.CertToPEM(t, intermediate)), 0600); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(second, []byte(suite.AlipayCertPEM), 0600); err != nil {
		t.Fatal(err)
	}
	pems, err := readAlipayCertificateFiles(first + ", " + second)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := buildCertificateCredentials(suite.AppCertPEM, testutil.AppPrivateKeyPEM, suite.RootPEM, pems); err != nil {
		t.Fatal(err)
	}
}
