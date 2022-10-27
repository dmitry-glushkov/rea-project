package models

import (
	"context"
	"fmt"

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
	_, err := db.Exec(
		ctx,
		`
			INSERT INTO projects
				(name, desc, owner_id)
				VALUES ($1, $2, $3);	
		`,
		p.Name, p.Desc, p.OwnerID,
	)
	if err != nil {
		err = fmt.Errorf("...: %w", err)
		return err
	}

	return nil
}

// GetProject ...
func GetProjects(ctx context.Context, db *pgx.Conn, page, limit int) ([]Project, error) {
	rows, err := db.Query(
		ctx,
		`
		SELECT id, name, desc, owner_id
			FROM projects;
		`,
	)
	if err != nil {
		err = fmt.Errorf("...: %w", err)
		return nil, err
	}
	defer rows.Close()

	var projects []Project
	for rows.Next() {
		var project Project
		err = rows.Scan()
		if err != nil {
			err = fmt.Errorf("...: %w", err)
			return nil, err
		}

		projects = append(projects, project)
	}

	return projects, nil
}

func GetProject(ctx context.Context, db *pgx.Conn, pid int) (Project, error) {
	row := db.QueryRow(
		ctx,
		`
		SELECT id, name, desc, owner_id
			FROM projects
			WHERE id = $1;	
		`,
		pid,
	)

	var project Project
	err := row.Scan()
	if err != nil {
		err = fmt.Errorf("...: %w", err)
		return Project{}, err
	}

	return project, nil
}
