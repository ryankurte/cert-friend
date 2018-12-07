/**
 * cert-friend
 * Certificate Database for managing an infrastructure instance
 *
 * This Software is licensed under the GNU GPLv3.
 *
 * https://github.com/ryankurte/cert-friend
 * Copyright 2018 Ryan Kurte
 */

package certfriend

import "time"

// DatabaseCertificate is a certificate line in the database
type DatabaseCertificate struct {
	CommonName string
	Serial     string
	Hash       string
	Issued     time.Time
	Revoked    bool `yaml:",omitempty"`
}

// Database contains a list of certificates for a given CA
type Database struct {
	root          DatabaseCertificate   `yaml:",omitempty"` // CA Root Certificate
	intermediates []DatabaseCertificate `yaml:",omitempty"` // CA Intermediate Certificates
	servers       []DatabaseCertificate `yaml:",omitempty"` // CA Server Certificates
	clients       []DatabaseCertificate `yaml:",omitempty"` // CA Client Certificates
}
