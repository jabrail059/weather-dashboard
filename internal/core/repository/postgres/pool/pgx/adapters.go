package core_pgx_pool

import (
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type PgxRows struct {
	pgx.Rows
}

type PgxRow struct {
	pgx.Row
}

type CommandTag struct {
	pgconn.CommandTag
}
