package main

import (
	"os"

	"github.com/jessevdk/go-flags"

	"github.com/ryankurte/ca-friend/lib"
)

func main() {
	// Parse options
	c := cafriend.Options{}
	p := flags.NewParser(&c, flags.Default)
	_, err := p.Parse()
	if err != nil {
		os.Exit(0)
	}

}
