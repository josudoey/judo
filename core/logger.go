package core

import (
	"context"
	"log"

	"go.uber.org/zap"
)

type loggerContextKey struct{}

func LoggerPlugin(ctx context.Context) (context.Context, context.CancelFunc) {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatal(err)
	}

	ctx = context.WithValue(ctx, loggerContextKey{}, logger)
	return ctx, func() {
		logger.Sync()
	}
}

func UseLogger(ctx context.Context) *zap.Logger {
	logger, ok := ctx.Value(loggerContextKey{}).(*zap.Logger)
	if !ok {
		return nil
	}

	return logger
}
