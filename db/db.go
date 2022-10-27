package db

import (
	"context"

	"github.com/jackc/pgx/v4"
)

// NewDB ...
func NewDB(ctx context.Context) (*pgx.Conn, error) {
	connString := "postgres://dmglushkov:3822@localhost:5432/postgres"
	conn, err := pgx.Connect(ctx, connString)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
