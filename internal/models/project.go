package models

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
)

// Project ...
type Project struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Desc   string `json:"desc"`
	Owner  string `json:"owner"`
	Sum    int    `json:"sum"`
	Target int    `json:"target"`
}

// Save ...
func (p *Project) Save(ctx context.Context, db *pgx.Conn) error {
	// _, err := db.Exec(
	// 	ctx,
	// 	`
	// 		INSERT INTO projects
	// 			(name, desc, owner_id)
	// 			VALUES ($1, $2, $3);
	// 	`,
	// 	p.Name, p.Desc, p.Owner,
	// )
	// if err != nil {
	// 	err = fmt.Errorf("...: %w", err)
	// 	return err
	// }

	return nil
}

// GetProjects ...
func GetProjects(ctx context.Context, db *pgx.Conn, page, limit int) ([]Project, int, error) {
	rows, err := db.Query(
		ctx,
		`
		SELECT id, name, desc, owner_id
			FROM projects;
		`,
	)
	if err != nil {
		err = fmt.Errorf("...: %w", err)
		return nil, 0, err
	}
	defer rows.Close()

	var projects []Project
	for rows.Next() {
		var project Project
		err = rows.Scan()
		if err != nil {
			err = fmt.Errorf("...: %w", err)
			return nil, 0, err
		}

		projects = append(projects, project)
	}

	return projects, 0, nil // TODO count
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

func GetProjectMock(ctx context.Context, db *pgx.Conn, pid int) (Project, error) {
	return Project{
		ID:    pid,
		Owner: "создатель проекта",
		Desc: `
		большое описание проекта большое описание проекта большое описание проекта большое описание проекта большое описание проекта
		большое описание проекта большое описание проекта большое описание проекта большое описание проекта большое описание проекта
		большое описание проекта большое описание проекта большое описание проекта большое описание проекта большое описание проекта
		большое описание проекта большое описание проекта большое описание проекта большое описание проекта большое описание проекта
		большое описание проекта большое описание проекта большое описание проекта большое описание проекта большое описание проекта
		большое описание проекта большое описание проекта большое описание проекта большое описание проекта большое описание проекта
		большое описание проекта большое описание проекта большое описание проекта большое описание проекта большое описание проекта
		большое описание проекта большое описание проекта большое описание проекта большое описание проекта большое описание проекта
		большое описание проекта большое описание проекта большое описание проекта большое описание проекта большое описание проекта
		`,
		Name:   "название проекта",
		Sum:    100000,
		Target: 400000,
	}, nil
}

func GetProjectsMock(ctx context.Context, db *pgx.Conn, page int, limit int) ([]Project, int, error) {
	return []Project{
		{
			ID:     1,
			Owner:  "владелец 1",
			Name:   "проект 1",
			Desc:   "описание проекта 1",
			Sum:    178,
			Target: 300,
		},
		{
			ID:     2,
			Owner:  "владелец 2",
			Name:   "проект 2",
			Desc:   "описание описание описание описание описание описание описание описание описание описание описание описание описание описание описание",
			Sum:    1780,
			Target: 3000,
		},
		{
			ID:     3,
			Owner:  "владелец 3",
			Name:   "проект 3",
			Sum:    17800,
			Target: 30000,
		},
		{
			ID:     4,
			Owner:  "владелец 4",
			Name:   "проект 4",
			Sum:    178000,
			Target: 300000,
		},
		{
			ID:     5,
			Owner:  "владелец 5",
			Name:   "проект 5",
			Sum:    1780000,
			Target: 3000000,
		},
	}, 5, nil
}
