package models

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
)

// Investment ...
type Investment struct {
	UID int `json:"uid"`
	PID int `json:"pid"`
	Val int `json:"val"`
}

// Save ...
func (d *Investment) Save(ctx context.Context, db *pgx.Conn) error {
	_, err := db.Exec(
		ctx,
		`
			INSERT INTO investments
				(uid, pid, val)
				VALUES ($1, $2, $3);
		`,
		d.UID, d.PID, d.Val,
	)
	if err != nil {
		err = fmt.Errorf("...: %w", err)
		return err
	}

	return nil
}

func (d *Investment) SaveMock(ctx context.Context, db *pgx.Conn) error {
	return nil
}

func GetProjectInvestments(ctx context.Context, db *pgx.Conn, pid int) ([]Investment, error) {
	rows, err := db.Query(
		ctx,
		`
		SELECT uid, pid, val
			FROM stages
			WHERE pid = $1;	
		`,
		pid,
	)
	if err != nil {
		err = fmt.Errorf("...: %w", err)
		return nil, err
	}
	defer rows.Close()

	var investments []Investment
	for rows.Next() {
		var investment Investment
		err = rows.Scan()
		if err != nil {
			err = fmt.Errorf("...: %w", err)
			return nil, err
		}

		investments = append(investments, investment)
	}

	return investments, nil
}

func GetUserInvestments(ctx context.Context, db *pgx.Conn, uid int) ([]Investment, error) {
	rows, err := db.Query(
		ctx,
		`
		SELECT uid, pid, val
			FROM stages
			WHERE uid = $1;	
		`,
		uid,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var investments []Investment
	for rows.Next() {
		var investment Investment
		err = rows.Scan(
			&investment.UID,
			&investment.PID,
			&investment.Val,
		)
		if err != nil {
			return nil, err
		}

		investments = append(investments, investment)
	}

	return investments, nil
}

func GetInvestmentsMock(ctx context.Context, db *pgx.Conn, pid int) ([]Investment, error) {
	return []Investment{
		{
			UID: 0,
			Val: 100,
		},
		{
			UID: 1,
			Val: 250,
		},
		{
			UID: 2,
			Val: 100,
		},
		{
			UID: 3,
			Val: 1000,
		},
		{
			UID: 4,
			Val: 50,
		},
	}, nil
}
