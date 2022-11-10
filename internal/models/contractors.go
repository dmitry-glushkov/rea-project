package models

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/lib/pq"
)

type Contractor struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Interest string `json:"interests"`
	Pids     pq.Int32Array
	Projects []Project `json:"projects"`
}

func (c Contractor) Save(ctx context.Context, db *pgx.Conn) error {
	_, err := db.Exec(
		ctx,
		`
		insert into contractors
			(name, interests)
			values ($1, $2);	
		`,
		c.Name, c.Interest,
	)
	return err
}

func (c Contractor) SaveMock(ctx context.Context, db *pgx.Conn) error {
	return nil
}

func GetContractors(ctx context.Context, db *pgx.Conn, pid int) ([]Contractor, error) {
	rowsC, err := db.Query(
		ctx,
		`
			select id, name, interests, pids
				from contractors
				where ($1 = any(pids) or $1 = 0);
		`,
		pid,
	)
	if err != nil {
		return nil, err
	}
	defer rowsC.Close()

	var slc []*Contractor
	var pids pq.Int32Array
	for rowsC.Next() {
		var ob Contractor
		err = rowsC.Scan(
			&ob.ID,
			&ob.Name,
			&ob.Interest,
			&ob.Pids,
		)
		if err != nil {
			return nil, err
		}

		pids = append(pids, ob.Pids...)

		slc = append(slc, &ob)
	}

	var rowsP pgx.Rows
	rowsP, err = db.Query(
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
	defer rowsP.Close()

	projects := make(map[int]Project)
	for rowsP.Next() {
		var p Project
		err = rowsP.Scan(
			&p.ID,
			&p.Name,
		)
		if err != nil {
			return nil, err
		}

		projects[p.ID] = p
	}

	for _, c := range slc {
		var prjs []Project
		for _, pid := range c.Pids {
			prjs = append(prjs, projects[int(pid)])
		}
		c.Projects = prjs
	}

	var resp []Contractor
	for _, s := range slc {
		resp = append(resp, *s)
	}

	return resp, nil
}

func GetContractorsMock(ctx context.Context, db *pgx.Conn, pid int) ([]Contractor, error) {
	return []Contractor{
		{
			ID:       0,
			Name:     "contractor1",
			Interest: "заборы",
			Projects: []Project{
				{
					ID:   0,
					Name: "project 1",
				},
				{
					ID:   1,
					Name: "project 2",
				},
			},
		},
		{
			ID:       1,
			Name:     "contractor2",
			Interest: "загоны",
			Projects: []Project{
				{
					ID:   2,
					Name: "project 3",
				},
				{
					ID:   3,
					Name: "project 4",
				},
			},
		},
		{
			ID:       2,
			Name:     "contractor3",
			Interest: "ограды",
			Projects: []Project{
				{
					ID:   4,
					Name: "project 5",
				},
				{
					ID:   5,
					Name: "project 6",
				},
			},
		},
	}, nil
}
