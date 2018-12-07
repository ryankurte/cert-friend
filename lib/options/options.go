package options

// Options is a go-flags compatible options structure for passing command line options.
type Options struct {
	BaseOptions

	Config       ConfigOptions       `command:"configure" description:"Configure new certificate Infrastructure"`
	CreateKey    CreateKeyOptions    `command:"key" description:"Generate a new public/private key pair"`
	CreateCA     CreateCAOptions     `command:"ca" description:"Generate a new Certificate Authority (CA)"`
	CreateServer CreateServerOptions `command:"server" description:"Generate a new Server Certificate"`
	CreateClient CreateClientOptions `command:"client" description:"Generate a new Client Certificate"`
	Revoke       RevokeOptions       `command:"revoke" description:"Revoke a certificate"`

	Sync SyncOptions `command:"sync" description:"Synchronize database with the provided configuration file"`

	CRL CRLOptions `command:"crl" description:"Generate a Certificate Revocation List (CRL)"`
	CT  CTOptions  `command:"ct" description:"Generate a Certificate Transparency Log (CTL)"`

	Version Version `command:"version" description:"Show version and exit"`
}

// CreateCAOptions passed to CreateCA subcommand
type CreateCAOptions struct {
}

// CreateServerOptions passed to CreateServer subcommand
type CreateServerOptions struct {
}

// CreateClientOptions passed to CreateClient subcommand
type CreateClientOptions struct {
}

type RevokeOptions struct {
}

type SyncOptions struct {
}

type CRLOptions struct {
}

type CTOptions struct {
}

type Version struct{}
