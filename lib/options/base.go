/**
 * cert-friend
 * Base Options shared between all subcommands
 *
 * This Software is licensed under the GNU GPLv3.
 *
 * https://github.com/ryankurte/cert-friend
 * Copyright 2018 Ryan Kurte
 */

package options

// BaseOptions share by all subcommands
type BaseOptions struct {
	Config         string `short:"c" long:"config" description:"Configuration file containing defaults and desired operations" default:"certfriend.yml"`
	Database       string `short:"d" long:"database" description:"Database file for managing issued certificates" default:"database.yml"`
	OutDir         string `short:"o" long:"out-dir" description:"Directory for generated outputs" default:"certs/"`
	NonInteractive bool   `long:"non-interactive" description:"Disable interactive mode prompts"`
	Verbose        bool   `short:"v" long:"verbose" description:"Enable verbose logging"`
}
