package options

// BaseOptions share by all subcommands
type BaseOptions struct {
	Config         string `short:"c" long:"config" description:"Configuration file containing defaults and desired operations" default:"certfriend.yml"`
	Database       string `short:"d" long:"database" description:"Database file for managing issued certificates" default:"database.yml"`
	NonInteractive bool   `long:"non-interactive" description:"Disable interactive mode prompts"`
	Verbose        bool   `short:"v" long:"verbose" description:"Enable verbose logging"`
}
