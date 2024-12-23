package cmd

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/josudoey/judo/core"
	"github.com/spf13/cobra"
)

type Runnable func(ctx context.Context, args []string) error

func Run(fn Runnable) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		ctx, cleanup := core.Setup(cmd.Context(), core.LoggerPlugin)
		go func() {
			<-ctx.Done()
			cleanup()
		}()

		return run(ctx, args, fn)
	}
}

func run(ctx context.Context, args []string, fn Runnable) error {
	ctx, cancel := context.WithCancel(ctx)
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-shutdown
		cancel()
	}()
	return fn(ctx, args)
}
