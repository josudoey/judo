package core

import (
	"context"

	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

func NewPgxConn(ctx context.Context, connString string) *pgx.Conn {
	conn, err := pgx.Connect(ctx, connString)

	if err != nil {
		logger := UseLogger(ctx)
		logger.Fatal("Failed to parse PostgreSQL URL", zap.Error(err))
	}

	go func() {
		<-ctx.Done()
		conn.Close(ctx)
	}()

	return conn
}
