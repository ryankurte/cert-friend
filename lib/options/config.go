/**
 * cert-friend
 * Configuration Generation options
 *
 * This Software is licensed under the GNU GPLv3.
 *
 * https://github.com/ryankurte/cert-friend
 * Copyright 2018 Ryan Kurte
 */
package options

// ConfigOptions passed to Configuration subcommand
type ConfigOptions struct {
	CreateKeyOptions `yaml:",inline"`

	Overwrite bool `long:"overwrite" description:"Force overwriting of existing configuration" yaml:"-"`
}
