package cafriend

import (
	"fmt"
	"log"

	"github.com/Songmu/prompter"
	"github.com/jinzhu/copier"

	"github.com/ryankurte/ca-friend/lib/options"
)

type CA struct {
}

func Configure(c Config, d Database, b options.BaseOptions, o options.ConfigOptions) Config {
	// Load options over config
	copier.Copy(&c.ConfigOptions, &o)

	log.Printf("Alright, let's configure some certificate infrastructure!")
	log.Printf(
		"You're about to be bombarded with rather a lot of choices, we're going to try and inform you as to pros and cons of each so you can apply them to your own context.\n\n" +
			"If you're not sure, the defaults are fine, and if you'd rather skip this choose-your-own-adventure and edit a default configuration file, re-run this with the --non-interactive flag.\n\n")

	// Load base configuration interactively
	c.ConfigOptions.Interactive()

	// TODO: subject

	// TODO: hash function

	// Save configuration
	exists := FileExists(b.Config)
	if o.Overwrite {
		log.Printf("Saving configuration to: '%s'", b.Config)
		SaveFile(b.Config, &o)
	} else {
		save := prompter.YesNo(fmt.Sprintf("Would you like to save this configuration to: '%s'?", b.Config), true)
		if save && !exists {
			log.Printf("Saving new configuration to: '%s'", b.Config)
			SaveFile(b.Config, &o)
		} else if save && exists {
			overwrite := prompter.YesNo("\n[WARNING] configuration file already exists, would you like to overwrite?", false)
			if overwrite {
				log.Printf("Overwriting configuration: '%s'", b.Config)
				SaveFile(b.Config, &o)
			} else {
				log.Printf("[WARNING] config file exists, use --overwrite to overwrite the existing file")
			}
		}
	}

	return Config{ConfigOptions: o}
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
