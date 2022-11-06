package models

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
)

type Investor struct {
	ID    int    `json:"id"`
	Login string `json:"login"`
	// TODO
}

// GetInvestors ...
func GetInvestors(ctx context.Context, db *pgx.Conn, pid int) ([]Investor, error) {
	rows, err := db.Query(
		ctx,
		`
			SELECT u.id, u.login, d.val
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

func (inv Investor) Save(ctx context.Context, db *pgx.Conn) error {
	return nil
}
