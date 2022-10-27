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
		err = fmt.Errorf("...: %w", err)
		return err
	}

	return nil
}

// GetUsers ...
func GetUsers(ctx context.Context, db *pgx.Conn) ([]User, error) {
	rows, err := db.Query(
		ctx,
		`
			SELECT id, login, role FROM users;
		`,
	)
	if err != nil {
		err = fmt.Errorf("...: %w", err)
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err = rows.Scan()
		if err != nil {
			err = fmt.Errorf("...: %w", err)
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

// GetDonators ...
func GetDonators(ctx context.Context, db *pgx.Conn, pid int) ([]User, error) {
	rows, err := db.Query(
		ctx,
		`
			SELECT u.id, u.login, d.val
				FROM users AS u
				LEFT JOIN donates AS d
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

	var users []User
	for rows.Next() {
		var user User
		err = rows.Scan()
		if err != nil {
			err = fmt.Errorf("...: %w", err)
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}
