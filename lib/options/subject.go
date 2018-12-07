package options

import (
	"crypto/x509/pkix"
)

// SubjectOptions encodes certificate subject information
type SubjectOptions struct {
	CommonName string `long:"commonname" description:"Common Name"`
	Serial     string `long:"serial" description:"Serial"`

	Country string `long:"country" description:"ISO 3166 country code" yaml:",omitempty"`
	Org     string `long:"org" description:"Organisation name" yaml:",omitempty"`
	OrgUnit string `long:"orgunit" description:"Organisational Unit" yaml:",omitempty"`
}

// ToPkixName converts a SubjectOptions object to a go pkix.Name object
func (so *SubjectOptions) ToPkixName() pkix.Name {
	name := pkix.Name{
		CommonName:   so.CommonName,
		SerialNumber: so.Serial,
	}

	if so.Country != "" {
		name.Country = []string{so.Country}
	}
	if so.Org != "" {
		name.Organization = []string{so.Org}
	}
	if so.OrgUnit != "" {
		name.OrganizationalUnit = []string{so.OrgUnit}
	}

	return name
}
