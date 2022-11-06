package models

import (
	"context"

	"github.com/jackc/pgx/v4"
)

type Innovator struct {
	ID   int    `json:"id"`
	Name string `json:"login"`
}

func GetInnovators(ctx context.Context, db *pgx.Conn) ([]Innovator, error) {
	// TODO
	return nil, nil
}
