package options

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
)

// TLSOptions for TLS connections
type TLSOptions struct {
	CA   string `long:"tls-ca" description:"TLS Certificate Authority (CA) Certificate"`
	Cert string `long:"tls-cert" description:"TLS Client Certificate"`
	Key  string `long:"tls-key" description:"TLS Client Key"`
}

// ToTLSConfig generates a go tls.Config from a TLSOptions struct
func (o *TLSOptions) ToTLSConfig() (*tls.Config, error) {
	tlsConfig := tls.Config{}

	// Require both o.Key and o.Cert to be specified
	if (o.Cert == "" && o.Key != "") || (o.Cert != "" && o.Key == "") {
		return nil, fmt.Errorf("TLS requires both --tls-cert and --tls-key arguments to be specified")
	}

	// Load client certificate and key if specified
	if o.Cert != "" && o.Key != "" {
		// Load certificate and key
		cert, err := tls.LoadX509KeyPair(o.Cert, o.Key)
		if err != nil {
			return nil, err
		}

		tlsConfig.Certificates = []tls.Certificate{cert}
	}

	// Load certificate authority if provided
	if o.CA != "" {
		caPem, err := ioutil.ReadFile(o.CA)
		if err != nil {
			return nil, err
		}

		roots := x509.NewCertPool()
		roots.AppendCertsFromPEM(caPem)

		tlsConfig.RootCAs = roots
	}

	return &tlsConfig, nil
}
