/**
 * cert-friend
 * Public Key Crypto Components
 *
 * This Software is licensed under the GNU GPLv3.
 *
 * https://github.com/ryankurte/cert-friend
 * Copyright 2018 Ryan Kurte
 */

package certfriend

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"

	"github.com/ryankurte/cert-friend/lib/options"
)

// Key is an RSA or ECDSA key
type Key struct {
	key   interface{}
	pkcs8 bool
}

// PublicKey fetches a public key for a given certificate
func (c *Key) PublicKey() interface{} {
	switch k := c.key.(type) {
	case *rsa.PrivateKey:
		return &k.PublicKey
	case *ecdsa.PrivateKey:
		return &k.PublicKey
	default:
		return nil
	}
}

// PrivateKey fetches the private key instance
func (c *Key) PrivateKey() interface{} {
	return c.key
}

// PEMBlock fetches a PEM encoded block containing the DER encoded private key
func (c *Key) PEMBlock() (*pem.Block, error) {
	// Encode to PKCS8 if specified
	if c.pkcs8 {
		b, err := x509.MarshalPKCS8PrivateKey(c.key)
		if err != nil {
			return nil, err
		}
		return &pem.Block{Type: "PRIVATE KEY", Bytes: b}, nil
	}

	// Otherwise use PKCS1 with appropriate internal block
	switch k := c.key.(type) {
	case *rsa.PrivateKey:
		b := x509.MarshalPKCS1PrivateKey(k)
		return &pem.Block{Type: "RSA PRIVATE KEY", Bytes: b}, nil
	case *ecdsa.PrivateKey:
		b, err := x509.MarshalECPrivateKey(k)
		if err != nil {
			return nil, err
		}
		return &pem.Block{Type: "EC PRIVATE KEY", Bytes: b}, nil

	default:
		return nil, fmt.Errorf("Invalid key type")
	}
}

// Save writes the private key to a PEM encoded file
func (c *Key) Save(filename string) error {
	buff := bytes.NewBuffer([]byte{})

	// Generate PEM block
	block, err := c.PEMBlock()
	if err != nil {
		return err
	}

	// Encode to PEM buffer
	if err := pem.Encode(buff, block); err != nil {
		return err
	}

	// Write to file
	return ioutil.WriteFile(filename, buff.Bytes(), 0600)
}

// LoadKey loads a key from a file
func LoadKey(filename string) (*Key, error) {
	// Load PEM data
	keyPEM, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// Decode PEM to DER block
	keyDERBlock, _ := pem.Decode(keyPEM)
	if err != nil {
		return nil, err
	}

	var key interface{}

	// Decode key DER block
	switch keyDERBlock.Type {
	case "RSA PRIVATE KEY":
		key, err = x509.ParsePKCS1PrivateKey(keyDERBlock.Bytes)
	case "EC PRIVATE KEY":
		key, err = x509.ParseECPrivateKey(keyDERBlock.Bytes)
	case "PRIVATE KEY":
		key, err = x509.ParsePKIXPublicKey(keyDERBlock.Bytes)
	}
	if err != nil {
		return nil, err
	}

	return &Key{key: key}, nil
}

// NewKey creates a new key with the provided options
func NewKey(ko options.CreateKeyOptions) (*Key, error) {
	var key interface{}
	var err error

	switch ko.Algorithm {
	case "rsa":
		key, err = NewRSA(ko.KeyBits)
	case "ecdsa":
		key, err = NewECDSA(ko.Curve)
	default:
		err = fmt.Errorf("Unsupported public key algorithm: '%s'", ko.Algorithm)
	}

	return &Key{key: key, pkcs8: ko.UsePKCS8}, err
}

// NewRSA generates a new RSA key with the provided size
func NewRSA(bits int) (interface{}, error) {
	return rsa.GenerateKey(rand.Reader, bits)
}

// NewECDSA generates a new ECDSA key on the provided curve
func NewECDSA(curve string) (interface{}, error) {
	var key interface{}
	var err error

	switch curve {
	case "P224":
		key, err = ecdsa.GenerateKey(elliptic.P224(), rand.Reader)
	case "P256":
		key, err = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	case "P384":
		key, err = ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	case "P521":
		key, err = ecdsa.GenerateKey(elliptic.P521(), rand.Reader)
	default:
		err = fmt.Errorf("Unsupported ECDSA curve: '%s'", curve)
	}

	return key, err
}
