package core

import "context"

type Installable func(ctx context.Context) (context.Context, context.CancelFunc)

func Setup(ctx context.Context, plugins ...Installable) (context.Context, context.CancelFunc) {
	cleanup := make([]context.CancelFunc, 0, len(plugins))

	for _, install := range plugins {
		var cancel context.CancelFunc

		ctx, cancel = install(ctx)
		if cancel == nil {
			continue
		}
		cleanup = append(cleanup, cancel)
	}

	return ctx, func() {
		for i := len(cleanup) - 1; i >= 0; i-- {
			cleanup[i]()
		}
	}
}
