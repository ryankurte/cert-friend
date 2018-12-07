package options

// ConfigOptions passed to Configuration subcommand
type ConfigOptions struct {
	CreateKeyOptions `yaml:",inline"`

	Overwrite bool `long:"overwrite" description:"Force overwriting of existing configuration" yaml:"-"`
}
