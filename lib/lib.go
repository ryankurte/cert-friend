package cafriend

import (
	"fmt"
	"log"
	"regexp"
	"strconv"

	"github.com/Songmu/prompter"
)

type CA struct {
}

func Configure(b *BaseOptions, o *ConfigOptions) Config {
	var c Config

	log.Printf("Alright, let's configure some certificate infrastructure!")
	log.Printf(
		"You're about to be bombarded with rather a lot of choices, we're going to try and inform you as to pros and cons of each so you can apply them to your own context.\n\n" +
			"If you're not sure, the defaults are fine, and if you'd rather skip this choose-your-own-adventure and edit a default configuration file, re-run this with the --non-interactive flag.\n\n")

	log.Printf(
		"First we have to decide what type of cryptographic algorithms you would like to use...\n" +
			"RSA is a reasonably old algorithm based on prime factorisation, keys should be > 2096 bits, and operations are relatively slow on embedded systems\n" +
			"	For more information see: https://en.wikipedia.org/wiki/RSA_(cryptosystem)\n" +
			"ECDSA is a more modern elliptic curve based algorithm, keys can be shorter (256-512bits) and performance on embedded systems is slightly better\n" +
			"	For more information see: https://en.wikipedia.org/wiki/Elliptic_Curve_Digital_Signature_Algorithm\n" +
			"A 256-bit ECDSA key offers equivalent cryptographic strength to a 3072-bit RSA key, so unless you have a specific reason not to, we recommend ECDSA.\n\n")
	c.Algorithm = prompter.Choose("Which cryptographic algorithm would you like to use?", []string{"rsa", "ecdsa"}, "ecdsa")

	switch c.Algorithm {
	case "rsa":
		log.Printf(
			"\n\nRSA Keys come in different sizes, in general longer keys are more secure but more difficult to compute (ie. take more time).\n" +
				"We recommend a key length of 4096 bits, and a minimum of 2048 bits should shorter keys be required.\n\n")
		str := prompter.Regexp("How many bits would you like your RSA key to be?", regexp.MustCompile(`(?m)[\d]+`), fmt.Sprintf("%d", o.KeyBits))
		keyBits, _ := strconv.Atoi(str)
		c.RSA.KeyBits = uint32(keyBits)
		if c.RSA.KeyBits < 2048 {
			log.Fatalf("RSA keys with length < 2048 bits are not supported (if you /really/ need this please open an issue)")
		}
	case "ecdsa":
		log.Printf(
			"\n\nECDSA supports a set of different curves for key generation, each corresponding to the key length and cryptographic strength.\n" +
				"We recommend a P256 for a reasonable balance between key length and cryptographic strength\n\n")
		c.ECDSA.Curve = prompter.Choose("How many bits would you like your RSA key to be?", []string{"P224", "P256", "P384", "P521"}, o.Curve)
	}

	// TODO: hash function

	// TODO: subject

	// Prompt to save
	save := prompter.YesNo(fmt.Sprintf("Would you like to save this configuration to: '%s'?", b.Config), true)
	exists := FileExists(b.Config)

	if (save && !exists) || o.Overwrite {
		SaveFile(b.Config, &c)
	} else if save && exists {
		overwrite := prompter.YesNo("\n[WARNING] configuration file already exists, would you like to overwrite?", false)
		if overwrite {
			SaveFile(b.Config, &c)
		}
	}
	return c
}

func CreateCA() CA {
	return CA{}
}

func (ca *CA) CreateServer() {

}

func (ca *CA) CreateClient() {

}

func (ca *CA) Revoke() {

}
