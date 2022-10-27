package models

import (
	"context"

	"github.com/jackc/pgx/v4"
)

// Goal ...
type Goal struct {
	ID      int
	PID     int
	Target  int
	DueDate int
}

// Save ...
func (g *Goal) Save(ctx context.Context, db *pgx.Conn) error {
	return nil
}

// GetGoals ...
func GetGoals(ctx context.Context, db *pgx.Conn, pid int) ([]Goal, error) {
	return nil, nil
}
