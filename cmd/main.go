package cmd

import (
	"github.com/spf13/cobra"
)

var cmds []*cobra.Command

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd: true,
			DisableNoDescFlag: true,
		},
		SilenceErrors: true,
		SilenceUsage:  true,
	}
	cmd.AddCommand(cmds...)
	return cmd
}

func AddCommand(cmd *cobra.Command) error {
	cmds = append(cmds, cmd)
	return nil
}

func SetupCommand(getCmds ...func() *cobra.Command) error {
	for _, getCmd := range getCmds {
		AddCommand(getCmd())
	}

	return nil
}
