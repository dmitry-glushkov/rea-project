package models

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/lib/pq"
)

type Innovator struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Pids     pq.Int32Array
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
		select id, name, pids
			from innovators
		`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var slc []Innovator
	var pids pq.Int32Array
	for rows.Next() {
		var ob Innovator
		err = rows.Scan(
			&ob.ID,
			&ob.Name,
			&ob.Pids,
		)
		if err != nil {
			return nil, err
		}

		pids = append(pids, ob.Pids...)
		slc = append(slc, ob)
	}

	var rowsp pgx.Rows
	rowsp, err = db.Query(
		ctx,
		`
		select id, name
			from projects
			where id = any($1);	
		`,
		pids,
	)
	if err != nil {
		return nil, err
	}
	defer rowsp.Close()

	projects := make(map[int]Project)
	for rowsp.Next() {
		var p Project
		err = rowsp.Scan(
			&p.ID,
			&p.Name,
		)
		if err != nil {
			return nil, err
		}

		projects[p.ID] = p
	}

	for _, i := range slc {
		var prjs []Project
		for _, pid := range i.Pids {
			prjs = append(prjs, projects[int(pid)])
		}
		i.Projects = prjs
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
