package options

import (
	"time"
)

// CertificateOptions are general certificate options
type CertificateOptions struct {
	Hostnames []string      `long:"hostname" description:"Hostname or IP for certificate" yaml:",omitempty"`
	ValidFrom time.Time     `long:"valid-from" description:"Certificate initial validity" yaml:",omitempty"`
	ValidFor  time.Duration `long:"valid-for" description:"Length of time for which the certificate will be valid" yaml:",omitempty"`
}
