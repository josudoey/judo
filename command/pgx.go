package command

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/josudoey/judo/core"
	"go.uber.org/zap"
)

func newPgxConn(ctx context.Context, connString string) *pgx.Conn {
	conn, err := pgx.Connect(ctx, connString)

	if err != nil {
		logger := core.UseLogger(ctx)
		logger.Fatal("Failed to parse PostgreSQL URL", zap.Error(err))
	}

	go func() {
		<-ctx.Done()
		conn.Close(ctx)
	}()

	return conn
}
