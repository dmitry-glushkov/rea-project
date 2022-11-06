package models

import (
	"context"

	"github.com/jackc/pgx/v4"
)

type Innovator struct {
}

func GetInnovators(ctx context.Context, db *pgx.Conn) ([]Innovator, error) {
	// TODO
	return nil, nil
}
