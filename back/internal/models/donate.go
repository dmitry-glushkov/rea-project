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
	fmt.Println("uid: ", d.UID, "; pid: ", d.PID, "; val: ", d.Val, ";")
	return nil // TODO
}

func GetDonates() ([]Donate, error) {
	return nil, nil
}
