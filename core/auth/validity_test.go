package auth

import (
	"crypto/x509"
	"errors"
	"testing"
	"time"
)

func TestCertificateValidityBoundariesAndSnapshot(t *testing.T) {
	start := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	end := start.Add(time.Hour)
	cert := &x509.Certificate{NotBefore: start, NotAfter: end}
	var validity certificateValidity
	validity.add("alipay_certificates[0]", cert)
	cert.NotBefore = start.Add(-time.Hour)
	cert.NotAfter = end.Add(time.Hour)
	for _, test := range []struct {
		name string
		now  time.Time
		fail bool
	}{
		{"before", start.Add(-time.Nanosecond), true},
		{"start", start, false},
		{"end", end, false},
		{"expired", end.Add(time.Nanosecond), true},
	} {
		t.Run(test.name, func(t *testing.T) {
			err := validity.check(test.now)
			if !test.fail {
				if err != nil {
					t.Fatal(err)
				}
				return
			}
			var invalid *CertValidityError
			if !errors.As(err, &invalid) || invalid.Field != "alipay_certificates[0]" || invalid.NotBefore != start || invalid.NotAfter != end {
				t.Fatalf("unexpected error: %v", err)
			}
			invalid.NotAfter = end.Add(time.Hour)
			if validity.check(test.now) == nil {
				t.Fatal("mutating an error changed the validity snapshot")
			}
		})
	}
}
