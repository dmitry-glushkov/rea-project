package models

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
)

// Donate ...
type Donate struct {
	UID int
	PID int
	Val int
}

// Save ...
func (d *Donate) Save(ctx context.Context, db *pgx.Conn) error {
	_, err := db.Exec(
		ctx,
		`
			INSERT INTO donates
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

func GetProjectDonates(ctx context.Context, db *pgx.Conn, pid int) ([]Donate, error) {
	rows, err := db.Query(
		ctx,
		`
		SELECT uid, pid, val
			FROM goals
			WHERE pid = $1;	
		`,
		pid,
	)
	if err != nil {
		err = fmt.Errorf("...: %w", err)
		return nil, err
	}
	defer rows.Close()

	var donates []Donate
	for rows.Next() {
		var donate Donate
		err = rows.Scan()
		if err != nil {
			err = fmt.Errorf("...: %w", err)
			return nil, err
		}

		donates = append(donates, donate)
	}

	return donates, nil
}

func GetUserDonates(ctx context.Context, db *pgx.Conn, uid int) ([]Donate, error) {
	rows, err := db.Query(
		ctx,
		`
		SELECT uid, pid, val
			FROM goals
			WHERE uid = $1;	
		`,
		uid,
	)
	if err != nil {
		err = fmt.Errorf("...: %w", err)
		return nil, err
	}
	defer rows.Close()

	var donates []Donate
	for rows.Next() {
		var donate Donate
		err = rows.Scan()
		if err != nil {
			err = fmt.Errorf("...: %w", err)
			return nil, err
		}

		donates = append(donates, donate)
	}

	return donates, nil
}
