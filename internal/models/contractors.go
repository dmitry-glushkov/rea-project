package models

import (
	"context"

	"github.com/jackc/pgx/v4"
)

type Contractor struct {
}

func GetContractors(ctx context.Context, db *pgx.Conn) ([]Contractor, error) {
	// TODO
	return nil, nil
}
