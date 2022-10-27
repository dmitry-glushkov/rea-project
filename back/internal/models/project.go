package models

import (
	"context"

	"github.com/jackc/pgx/v4"
)

// Project ...
type Project struct {
	ID      int
	Name    string
	Desc    string
	OwnerID int
}

// Save ...
func (p *Project) Save(ctx context.Context, db *pgx.Conn) error {
	return nil
}

// GetProject ...
func GetProjects(ctx context.Context, db *pgx.Conn, page, limit int) ([]Project, error) {
	return nil, nil
}

func GetProject(ctx context.Context, db *pgx.Conn, pid int) (Project, error) {
	return Project{}, nil
}
