package main

import (
	"os"

	"github.com/jessevdk/go-flags"
)

type Options struct {
	NewCA NewCA `command:"new-ca" description:"Create a new CA"`

	Version Version `command:"version" description:"Show version and exit"`
	Verbose bool    `short:"v" long:"verbose" description:"Enable verbose logging"`
}

type NewCA struct {
}

type Version struct{}

var version string = "NOT SET"

func main() {
	// Parse options
	c := Options{}
	p := flags.NewParser(&c, flags.Default)
	_, err := p.Parse()
	if err != nil {
		os.Exit(0)
	}

	i

}
