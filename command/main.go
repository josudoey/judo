package command

import (
	"github.com/josudoey/judo/core"
	"github.com/spf13/cobra"
)

var cmds []*cobra.Command

func New() *cobra.Command {
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

func addCommand(cmd *cobra.Command) error {
	cmds = append(cmds, cmd)
	return nil
}

func setupCommand(getCmds ...func() *cobra.Command) error {
	for _, getCmd := range getCmds {
		addCommand(getCmd())
	}

	return nil
}

func commandRun(fn Runnable) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		ctx, cleanup := core.Setup(cmd.Context(), core.LoggerPlugin)
		go func() {
			<-ctx.Done()
			cleanup()
		}()

		return run(ctx, args, fn)
	}
}
