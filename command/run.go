package command

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

type Runnable func(ctx context.Context, args []string) error

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
