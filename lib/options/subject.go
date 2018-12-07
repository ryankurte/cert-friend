package options

import (
	"crypto/x509/pkix"

	"github.com/Songmu/prompter"
)

// SubjectOptions encodes certificate subject information
type SubjectOptions struct {
	CommonName string `long:"commonname" description:"Common Name"`
	Serial     string `long:"serial" description:"Serial string"`

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

// General Build general configuration with CLI interaction
func (so *SubjectOptions) General() {
	so.Country = prompter.Prompt("Enter your ISO 3166 country code", so.Country)
	so.Org = prompter.Prompt("Enter the name of your organisation", so.Org)
	so.OrgUnit = prompter.Prompt("Enter your organisational unit", so.OrgUnit)
}

// Specific Build certificate specific configuration with CLI interaction
func (so *SubjectOptions) Specific() {
	so.CommonName = prompter.Prompt("Enter a Common Name for your certificate", so.CommonName)
	so.Serial = prompter.Prompt("Enter a certificate serial number", so.Serial)
}
