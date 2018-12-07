package options

import (
	"fmt"
	"log"
	"regexp"
	"strconv"

	"github.com/Songmu/prompter"
)

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

// Interactive Build configuration with CLI interaction
func (o *CreateKeyOptions) Interactive() {
	log.Printf(
		"We have a choice between RSA and ECDSA:\n" +
			"RSA is a reasonably old algorithm based on prime factorisation, keys should be > 2096 bits, and operations are relatively slow on embedded systems\n" +
			"	For more information see: https://en.wikipedia.org/wiki/RSA_(cryptosystem)\n" +
			"ECDSA is a more modern elliptic curve based algorithm, keys can be shorter (256-512bits) and performance on embedded systems is slightly better\n" +
			"	For more information see: https://en.wikipedia.org/wiki/Elliptic_Curve_Digital_Signature_Algorithm\n" +
			"A 256-bit ECDSA key offers equivalent cryptographic strength to a 3072-bit RSA key. Unless you have a specific reason not to do so, we recommend ECDSA.\n\n")
	o.Algorithm = prompter.Choose("Which cryptographic algorithm would you like to use?", []string{"rsa", "ecdsa"}, o.Algorithm)

	switch o.Algorithm {
	case "rsa":
		log.Printf(
			"\n\nRSA Keys come in different sizes, in general longer keys are more secure but more difficult to compute (ie. take more time).\n" +
				"We recommend a key length of 4096 bits, and a minimum of 2048 bits should shorter keys be required.\n\n")
		str := prompter.Regexp("How many bits would you like your RSA key to be?", regexp.MustCompile(`(?m)[\d]+`), fmt.Sprintf("%d", o.KeyBits))
		keyBits, _ := strconv.Atoi(str)
		o.KeyBits = uint32(keyBits)
		if o.KeyBits < 2048 {
			log.Fatalf("RSA keys with length < 2048 bits are not supported (if you /really/ need this please open an issue)")
		}
		o.Curve = ""
	case "ecdsa":
		log.Printf(
			"\n\nECDSA supports a set of different curves for key generation, each corresponding to the key length and cryptographic strength.\n" +
				"We recommend a P256 for a reasonable balance between key length and cryptographic strength\n\n")
		o.Curve = prompter.Choose("How many bits would you like your RSA key to be?", []string{"P224", "P256", "P384", "P521"}, o.Curve)
		o.KeyBits = 0
	}
}
