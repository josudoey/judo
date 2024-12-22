package command

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type PGXScript struct {
	PostgresURL string `env:"POSTGRES_URL"`

	pgxConn *pgx.Conn
}

func NewPGXScript(ctx context.Context) *PGXScript {
	script := MustEnvParseAs[PGXScript]()
	script.pgxConn = newPgxConn(ctx, script.PostgresURL)
	return &script
}
