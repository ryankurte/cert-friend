/**
 * cert-friend
 * Certificate Infrastructure Configuration Objects
 *
 * This Software is licensed under the GNU GPLv3.
 *
 * https://github.com/ryankurte/cert-friend
 * Copyright 2018 Ryan Kurte
 */

package certfriend

import (
	"github.com/ryankurte/cert-friend/lib/options"
)

// Config is the configuration for a certificate infrastructure instance
type Config struct {
	options.ConfigOptions      `yaml:",inline"` // General configuration options
	options.CertificateOptions `yaml:",inline"` // General certificate options
	options.SubjectOptions     `yaml:",inline"` // General subject options
	Certs                      []CertConfig     // List of issued certificates
}

// CertConfig is a configuration entry for a single certificate
type CertConfig struct {
	options.CertificateOptions `yaml:",inline"` // Certificate instance options (overrides global config)
	options.SubjectOptions     `yaml:",inline"` // Certificate subject options (overrides global config)
}
