package models

import (
	"context"

	"github.com/jackc/pgx/v4"
)

type Innovator struct {
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	Projects []Project `json:"projects"`
}

func (inv Innovator) Save(ctx context.Context, db *pgx.Conn) error {
	_, err := db.Exec(
		ctx,
		`
		insert into innovators
			(name)
			values ($1);
		`,
		inv.Name,
	)
	return err
}

func (inv Innovator) SaveMock(ctx context.Context, db *pgx.Conn) error {
	return nil
}

func GetInnovators(ctx context.Context, db *pgx.Conn) ([]Innovator, error) {
	rows, err := db.Query(
		ctx,
		`
		
		`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var slc []Innovator
	for rows.Next() {
		var ob Innovator
		err = rows.Scan(
		// todo
		)
		if err != nil {
			return nil, err
		}

		slc = append(slc, ob)
	}

	return slc, nil
}

func GetInnovatorsMock(ctx context.Context, db *pgx.Conn) ([]Innovator, error) {
	return []Innovator{
		{
			ID:   0,
			Name: "innovator0",
			Projects: []Project{
				{
					ID:   0,
					Name: "project 0",
				},
			},
		},
		{
			ID:   1,
			Name: "innovator1",
			Projects: []Project{
				{
					ID:   1,
					Name: "project 1",
				},
				{
					ID:   2,
					Name: "project 2",
				},
				{
					ID:   3,
					Name: "project 3",
				},
			},
		},
		{
			ID:   2,
			Name: "innovator2",
		},
	}, nil
}
