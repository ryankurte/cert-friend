package cafriend

// Options is a go-flags compatible options structure for passing command line options.
type Options struct {
	BaseOptions

	CreateKey    CreateKeyOptions    `command:"key" description:"Generate a new public/private key pair"`
	CreateCA     CreateCAOptions     `command:"ca" description:"Generate a new Certificate Authority (CA)"`
	CreateServer CreateServerOptions `command:"server" description:"Generate a new Server Certificate"`
	CreateClient CreateClientOptions `command:"client" description:"Generate a new Client Certificate"`
	Revoke       RevokeOptions       `command:"revoke" description:"Revoke a certificate"`

	Sync SyncOptions `command:"sync" description:"Synchronize database with the provided configuration file"`

	CRL CRLOptions `command:"crl" description:"Generate a Certificate Revocation List (CRL)"`
	CT  CTOptions  `command:"ct" description:"Generate a Certificate Transparency Log (CTL)"`

	Config ConfigOptions `command:"configure" description:"Generate configuration files"`

	Version Version `command:"version" description:"Show version and exit"`
}

type BaseOptions struct {
	Config         string `short:"c" long:"config" description:"Configuration file containing defaults and desired operations" default:"certfriend.yml"`
	Database       string `short:"d" long:"database" description:"Database file for managing issued certificates" default:"database.yml"`
	NonInteractive bool   `long:"non-interactive" description:"Disable interactive mode prompts"`
	Verbose        bool   `short:"v" long:"verbose" description:"Enable verbose logging"`
}

type CreateKeyOptions struct {
	RSA   RSAOptions   `command:"rsa" description:"Generate ECDSA keypair"`
	ECDSA ECDSAOptions `command:"ecdsa" description:"Generate ECDSA keypair"`
}

type RSAOptions struct {
	KeyBits uint32 `long:"rsa-bits" description:"RSA private key length" default:"4096"`
}

type ECDSAOptions struct {
	Curve string `long:"ecdsa-curve" description:"ECDSA curve for key generation" choice:"P224" choice:"P256" choice:"P384" choice:"P521" default:"P256"`
}

type Subject struct {
}

type ConfigOptions struct {
	RSAOptions
	ECDSAOptions
	Overwrite bool `long:"overwrite" description:"Force overwriting of existing configuration"`
}

type CreateCAOptions struct {
}

type CreateServerOptions struct {
}

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

var version string = "NOT SET"

type General struct {
	KeyType string `short:"t" long:"key-type" description:"Public key type" option:"rsa" option:"ecdsa"`
	KeySize uint   `short:"s" long:"key-size" description:"Private key length`
}
