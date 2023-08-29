package command

import "github.com/spf13/cobra"

// New returns the new root command.
func New() *cobra.Command {
	command := cobra.Command{
		Use:   "fiddle",
		Short: "the missing RSS manager",
		Long:  "🎻 The missing RSS manager.",
		Args:  cobra.NoArgs,

		SilenceErrors: false,
		SilenceUsage:  true,
	}

	command.AddCommand(Read(), Watch())

	return &command
}
