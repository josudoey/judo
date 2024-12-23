package script

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/josudoey/judo/cmd"
	"github.com/josudoey/judo/core"
)

type PgxScript struct {
	PostgresURL string `env:"POSTGRES_URL"`

	PgxConn *pgx.Conn
}

func NewPGXScript(ctx context.Context) *PgxScript {
	script := cmd.MustEnvParseAs[PgxScript]()
	script.PgxConn = core.NewPgxConn(ctx, script.PostgresURL)
	return &script
}
