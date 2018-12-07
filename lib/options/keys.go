package options

// RSAOptions for RSA key creation
type RSAOptions struct {
	KeyBits uint32 `long:"rsa-bits" description:"RSA private key length" default:"4096" yaml:",omitempty"`
}

// ECDSAOptions for ECDSA key creation
type ECDSAOptions struct {
	Curve string `long:"ecdsa-curve" description:"ECDSA curve for key generation" choice:"P224" choice:"P256" choice:"P384" choice:"P521" default:"P256" yaml:",omitempty"`
}

// CreateKeyOptions passed to CreateKey command
type CreateKeyOptions struct {
	Algorithm    string `long:"algorithm" description:"Public key cryptographic algorithm to use" choice:"rsa" choice:"ecdsa" default:"ecdsa"`
	RSAOptions   `yaml:",inline"`
	ECDSAOptions `yaml:",inline"`
}
