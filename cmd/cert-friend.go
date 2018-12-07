package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jessevdk/go-flags"

	"github.com/ryankurte/ca-friend/lib"
)

func main() {
	log.SetFlags(0)

	// Parse options
	c := cafriend.Options{}
	p := flags.NewParser(&c, flags.Default)
	_, err := p.Parse()
	if err != nil {
		os.Exit(0)
	}

	// Try load config file
	var config cafriend.Config
	_, err = cafriend.LoadFile(c.BaseOptions.Config, &config)
	if err != nil {
		log.Fatalf("Error loading config: '%s'", err)
	}

	// Try to load database
	var database cafriend.Database
	_, err = cafriend.LoadFile(c.BaseOptions.Database, &database)
	if err != nil {
		log.Fatalf("Error loading database: '%s'", err)
	}

	switch p.Active.Name {
	case "configure":
		cafriend.Configure(&c.BaseOptions, &c.Config)

	default:
		fmt.Printf("Command '%s' not yet implemented", p.Active.Name)
		os.Exit(0)
	}
}
