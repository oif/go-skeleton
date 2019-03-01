package echo

import (
	"flag"

	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	command := NewCommand()
	command.Flags().AddGoFlagSet(flag.CommandLine)

	return command
}

func NewCommand() *cobra.Command {
	opt := NewDefaultOption()

	cmd := &cobra.Command{
		Use:   "echo",
		Short: "Echo it",
		RunE: func(c *cobra.Command, args []string) error {
			if err := opt.Run(); err != nil {
				return err
			}
			return nil
		},
	}
	return cmd
}
