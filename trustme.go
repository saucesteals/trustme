package trustme

import (
	"crypto/tls"
	"crypto/x509"
	_ "embed"
	"net/http"
	_ "unsafe"
)

var (
	//go:linkname systemRoots crypto/x509.systemRoots
	systemRoots *x509.CertPool

	trustedRoots *x509.CertPool

	//go:embed cacert.pem
	pem []byte
)

func init() {
	trustedRoots = x509.NewCertPool()
	trustedRoots.AppendCertsFromPEM(pem)
}

// TrustedCertPool returns the trusted cert pool.
func TrustedCertPool() *x509.CertPool {
	return trustedRoots
}

// ReplaceSystemCertPool replaces the system CAs from crypto/x509.systemRoots with trustme.trustedRoots.
func ReplaceSystemCertPool() {
	x509.SystemCertPool() // fulfill once.Do(initSystemRoots) from crypto/x509.systemRootsPool
	systemRoots = trustedRoots
}

// ReplaceDefaultHTTPCertPool replaces net/http.DefaultTransport's cert pool with trustme.trustedRoots.
// You must apply
func ReplaceDefaultHTTPCertPool() {
	transport, ok := http.DefaultTransport.(*http.Transport)
	if !ok {
		return // ignore silently
	}

	transport.TLSClientConfig = TrustTLSConfig(transport.TLSClientConfig)
}

// TrustTLSConfig sets the tls.Config's RootCAs to trustme.trustedRoots.
// It returns the same tls.Config or a new one if the one provided was nil.
func TrustTLSConfig(config *tls.Config) *tls.Config {
	if config == nil {
		config = &tls.Config{}
	}
	config.RootCAs = trustedRoots
	return config
}

// TrustHTTPTransport sets the http.Transport's RootCAs to trustme.trustedRoots.
// It returns the same http.Transports or a new one if the one provided was nil.
func TrustHTTPTransport(transport *http.Transport) *http.Transport {
	if transport == nil {
		transport = &http.Transport{}
	}
	transport.TLSClientConfig = TrustTLSConfig(transport.TLSClientConfig)
	return transport
}

// Trust replaces the system CAs from crypto/x509.systemRoots with trustme.trustedRoots where applicable.
func Trust() {
	ReplaceSystemCertPool()
	ReplaceDefaultHTTPCertPool()
}
