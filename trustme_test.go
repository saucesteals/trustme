package trustme

import (
	"crypto/x509"
	"reflect"
	"testing"
)

func TestTrust(t *testing.T) {
	roots := x509.NewCertPool()

	if !roots.AppendCertsFromPEM(pem) {
		t.Fatal(errFailedToParse)
		return
	}

	system, err := x509.SystemCertPool()
	if err != nil {
		t.Fatal(err)
	}

	if reflect.DeepEqual(roots.Subjects(), system.Subjects()) {
		t.Fatal("CAs match")
	}

	if err := Trust(); err != nil {
		t.Fatal(err)
	}

	system, err = x509.SystemCertPool()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(roots.Subjects(), system.Subjects()) {
		t.Fatal("CAs do not match")
	}
}
