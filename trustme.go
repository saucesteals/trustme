package trustme

import (
	"crypto/x509"
	_ "embed"
	"errors"
	_ "unsafe"
)

var (
	errFailedToParse = errors.New("trustme: failed to parse cacert.pem")

	//go:linkname systemRoots crypto/x509.systemRoots
	systemRoots *x509.CertPool

	//go:embed cacert.pem
	pem []byte
)

func Trust() error {
	x509.SystemCertPool()
	roots := x509.NewCertPool()
	if !roots.AppendCertsFromPEM(pem) {
		return errFailedToParse
	}
	systemRoots = roots
	return nil
}
