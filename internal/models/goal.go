package models

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v4"
)

// Goal ...
type Goal struct {
	ID      int
	PID     int
	Target  int
	DueDate time.Time
}

// Save ...
func (g *Goal) Save(ctx context.Context, db *pgx.Conn) error {
	_, err := db.Exec(
		ctx,
		`
			INSERT INTO goals
				(pid, target, due_date)
				VALUES ($1, $2, $3);	
		`,
		g.PID, g.Target, g.DueDate,
	)
	if err != nil {
		err = fmt.Errorf("...: %w", err)
		return err
	}

	return nil
}

// GetGoals ...
func GetGoals(ctx context.Context, db *pgx.Conn, pid int) ([]Goal, error) {
	rows, err := db.Query(
		ctx,
		`
		SELECT pid, target, due_date
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

	var goals []Goal
	for rows.Next() {
		var goal Goal
		err = rows.Scan()
		if err != nil {
			err = fmt.Errorf("...: %w", err)
			return nil, err
		}

		goals = append(goals, goal)
	}

	return goals, nil
}
