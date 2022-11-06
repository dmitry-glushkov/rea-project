package db

import (
	"context"

	"github.com/jackc/pgx/v4"
)

// NewDB ...
func NewDB(ctx context.Context) (*pgx.Conn, error) {
	connString := "postgres://dmglushkov:3822@localhost:5432/postgres"

	return pgx.Connect(ctx, connString)
}
