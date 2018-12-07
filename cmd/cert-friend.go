package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jessevdk/go-flags"

	certfriend "github.com/ryankurte/ca-friend/lib"
	"github.com/ryankurte/cert-friend/lib/options"
)

var version string = "NOT SET"

func main() {
	log.SetFlags(0)

	// Parse options
	c := options.Options{}
	p := flags.NewParser(&c, flags.Default)
	_, err := p.Parse()
	if err != nil {
		os.Exit(0)
	}

	// Try load config file
	var config certfriend.Config
	_, err = certfriend.LoadFile(c.BaseOptions.Config, &config)
	if err != nil {
		log.Fatalf("Error loading config: '%s'", err)
	}

	// Try to load database
	var database certfriend.Database
	_, err = certfriend.LoadFile(c.BaseOptions.Database, &database)
	if err != nil {
		log.Fatalf("Error loading database: '%s'", err)
	}

	switch p.Active.Name {
	case "configure":
		certfriend.Configure(config, database, c.BaseOptions, c.Config)

	default:
		fmt.Printf("Command '%s' not yet implemented", p.Active.Name)
		os.Exit(0)
	}
}
