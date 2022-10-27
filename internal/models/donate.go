package models

import (
	"context"

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
		return err
	}
	return nil
}

func GetDonates() ([]Donate, error) {
	return nil, nil
}
