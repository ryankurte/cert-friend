/**
 * cert-friend
 * Certificate generation options
 *
 * This Software is licensed under the GNU GPLv3.
 *
 * https://github.com/ryankurte/cert-friend
 * Copyright 2018 Ryan Kurte
 */

package options

import (
	"fmt"
	"net"
	"time"

	"github.com/Songmu/prompter"

	"github.com/araddon/dateparse"
)

// CertificateOptions are general certificate options
type CertificateOptions struct {
	Hostnames []string      `long:"hostname" description:"Hostnames for certificate" yaml:",omitempty"`
	IPs       []net.IP      `long:"ip" description:"IP addresses for certificate" yaml:",omitempty"`
	ValidFrom time.Time     `long:"valid-from" description:"Certificate initial validity" yaml:",omitempty"`
	ValidFor  time.Duration `long:"valid-for" description:"Length of time for which the certificate will be valid" yaml:",omitempty"`
}

// General Build general configuration with CLI interaction
func (o *CertificateOptions) General() error {
	validFromStr := prompter.Prompt("From when would you like certificates to be valid?", o.ValidFrom.String())
	validFromDate, err := dateparse.ParseStrict(validFromStr)
	if err != nil {
		return err
	}
	o.ValidFrom = validFromDate

	validForStr := prompter.Prompt("How long would you like certificates to be valid for?", o.ValidFor.String())
	validFor, err := time.ParseDuration(validForStr)
	if err != nil {
		return err
	}
	o.ValidFor = validFor

	return nil
}

// Specific build specific configuration with CLI interaction
func (o *CertificateOptions) Specific() error {
	hostnames := make([]string, 0)
	fmt.Println("Enter certificate hostnames.")
	for {
		hostnames = append(hostnames, prompter.Prompt("Hostname:", ""))
		if prompter.YesNo("Would you like to add another hostname?", false) {
			break
		}
	}
	o.Hostnames = hostnames
	return nil
}
