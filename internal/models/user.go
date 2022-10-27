package models

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
)

// User ...
type User struct {
	ID    int
	Login string
	Pass  string
	Role  string
}

// Save ...
func (u *User) Save(ctx context.Context, db *pgx.Conn) error {
	_, err := db.Exec(
		ctx,
		`
		INSERT INTO users 
			(login, pass, role)
			VALUES ($1, $2, $3);
		`,
		u.Login, u.Pass, u.Role,
	)
	if err != nil {
		return fmt.Errorf("") // TODO ...
	}
	return nil
}

// GetUsers ...
func GetUsers(ctx context.Context, db *pgx.Conn) ([]User, error) {
	return nil, nil
}

// GetDonators ...
func GetDonators(ctx context.Context, db *pgx.Conn, pid int) ([]User, error) {
	return nil, nil
}
