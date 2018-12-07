// see:
//   - https://golang.org/src/crypto/tls/generate_cert.go
//   - https://github.com/ryankurte/evilproxy/blob/master/lib/ingress/tls.go

package certfriend

import (
	"bytes"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
	"math/big"

	"github.com/ryankurte/cert-friend/lib/options"
)

// Certificate object is a valid certificate
type Certificate struct {
	*x509.Certificate
	der []byte
}

// Generate a random 256 bit serial for the certificate
func serial() (*big.Int, error) {
	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 256)
	return rand.Int(rand.Reader, serialNumberLimit)
}

// Build a base template for all certificate types
func template(co options.CertificateOptions, so options.SubjectOptions) (*x509.Certificate, error) {
	sn, err := serial()
	if err != nil {
		return nil, err
	}

	return &x509.Certificate{
		SerialNumber: sn,
		Subject:      so.ToPkixName(),

		NotBefore: co.ValidFrom,
		NotAfter:  co.ValidFrom.Add(co.ValidFor),

		KeyUsage:    x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},

		DNSNames:    co.Hostnames,
		IPAddresses: co.IPs,

		BasicConstraintsValid: true,
		IsCA:                  false,
	}, nil
}

// LoadCertificate loads a certificate from a file
func LoadCertificate(filename string) (*Certificate, error) {
	// Read PEM encoded file
	certPEM, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// Search for certificate block
	var certDERBlock *pem.Block
	certPEMBlock := certPEM
	for {
		certDERBlock, certPEMBlock = pem.Decode(certPEMBlock)
		if certDERBlock.Type == "CERTIFICATE" {
			break
		}
		if certDERBlock == nil && len(certPEM) == 0 {
			break
		}
	}

	// Decode certificate block
	cert, err := x509.ParseCertificate(certDERBlock.Bytes)
	if err != nil {
		return nil, err
	}

	return &Certificate{Certificate: cert, der: certPEM}, nil
}

// Save Saves a certificate to a PEM encoded file
func (c *Certificate) Save(filename string) error {
	buff := bytes.NewBuffer([]byte{})

	// Encode certificate to PEM
	if err := pem.Encode(buff, &pem.Block{Type: "CERTIFICATE", Bytes: c.der}); err != nil {
		return err
	}

	// Write PEM to file
	return ioutil.WriteFile(filename, buff.Bytes(), 0666)
}

// NewCA Creates a new CA certificate
func NewCA(key Key, co options.CertificateOptions, so options.SubjectOptions) (*Certificate, error) {
	t, err := template(co, so)
	if err != nil {
		return nil, err
	}

	t.IsCA = true
	t.KeyUsage |= x509.KeyUsageCertSign | x509.KeyUsageCRLSign
	t.MaxPathLen = 2

	derBytes, err := x509.CreateCertificate(rand.Reader, t, t, key.PublicKey(), key.PrivateKey())
	if err != nil {
		return nil, err
	}

	return &Certificate{Certificate: t, der: derBytes}, nil
}

// NewIntermediate creates a new intermediate certificate TODO
func NewIntermediate(key Key, co options.CertificateOptions, so options.SubjectOptions) (*Certificate, error) {
	t, err := template(co, so)
	if err != nil {
		return nil, err
	}

	t.IsCA = true
	t.KeyUsage |= x509.KeyUsageCertSign | x509.KeyUsageCRLSign
	t.MaxPathLen = 1

	return nil, nil
}

// NewServer creates a new Server certificate TODO
func NewServer(key Key, co options.CertificateOptions, so options.SubjectOptions) (*Certificate, error) {
	t, err := template(co, so)
	if err != nil {
		return nil, err
	}

	t.IsCA = false
	t.MaxPathLen = 0
	t.KeyUsage |= x509.KeyUsageKeyAgreement
	t.ExtKeyUsage = []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth}

	return nil, nil
}

// NewClient creates a new client certificate TODO
func NewClient(key Key, co options.CertificateOptions, so options.SubjectOptions) (*Certificate, error) {
	t, err := template(co, so)
	if err != nil {
		return nil, err
	}

	t.IsCA = false
	t.MaxPathLen = 0
	t.ExtKeyUsage = []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth}

	return nil, nil
}
