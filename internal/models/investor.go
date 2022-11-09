package models

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
)

type Investor struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Interests string    `json:"interests"`
	Projects  []Project `json:"projects"`
	Total     int       `json:"total"`
}

func (inv Investor) Save(ctx context.Context, db *pgx.Conn) error {
	_, err := db.Exec(
		ctx,
		`
		insert into investors
			(name, interests)
			values ($1, $2);	
		`,
		inv.Name, inv.Interests,
	)
	return err
}

func (inv Investor) SaveMock(ctx context.Context, db *pgx.Conn) error {
	return nil
}

// GetInvestors ...
func GetInvestors(ctx context.Context, db *pgx.Conn, pid int) ([]Investor, error) {
	rows, err := db.Query(
		ctx,
		`
			SELECT u.id, u.name, d.val
				FROM users AS u
				LEFT JOIN investments AS d
					ON u.id = d.uid
				WHERE d.pid = $1;
		`,
		pid,
	)
	if err != nil {
		err = fmt.Errorf("...: %w", err)
		return nil, err
	}
	defer rows.Close()

	var users []Investor
	for rows.Next() {
		var user Investor
		err = rows.Scan()
		if err != nil {
			err = fmt.Errorf("...: %w", err)
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func GetInvestorsMock(ctx context.Context, db *pgx.Conn, pid int) ([]Investor, error) {
	return []Investor{
		{
			ID:        0,
			Name:      "inv0",
			Interests: "it",
			Projects: []Project{
				{
					ID:   0,
					Name: "project 0",
				},
				{
					ID:   1,
					Name: "project 1",
				},
			},
			Total: 800,
		},
		{
			ID:        1,
			Name:      "inv1",
			Interests: "металлургия, агро",
			Projects: []Project{
				{
					ID:   2,
					Name: "project 2",
				},
				{
					ID:   3,
					Name: "project 3",
				},
				{
					ID:   4,
					Name: "project 4",
				},
				{
					ID:   5,
					Name: "project 5",
				},
				{
					ID:   6,
					Name: "project 6",
				},
			},
			Total: 1000234,
		},
		{
			ID:        2,
			Name:      "inv2",
			Interests: "",
			Projects:  []Project{},
		},
		{
			ID:        3,
			Name:      "inv3",
			Interests: "облачные технологии, IoT",
			Projects: []Project{
				{
					ID:   7,
					Name: "project 7",
				},
			},
			Total: 23421,
		},
		{
			ID:        4,
			Name:      "inv4",
			Interests: "заборы",
			Projects: []Project{
				{
					ID:   8,
					Name: "project 8",
				},
				{
					ID:   9,
					Name: "project 9",
				},
			},
			Total: 945645,
		},
	}, nil
}
