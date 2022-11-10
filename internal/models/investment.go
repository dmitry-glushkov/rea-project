package models

import (
	"context"

	"github.com/jackc/pgx/v4"
)

// Investment ...
type Investment struct {
	UIDname string `json:"uid"`
	UID     int    `json:"iid"`
	PID     int    `json:"pid"`
	Val     int    `json:"val"`
}

// Save ...
func (d *Investment) Save(ctx context.Context, db *pgx.Conn) error {
	_, err := db.Exec(
		ctx,
		`
			INSERT INTO investments
				(uid, pid, val)
				VALUES ((select id from investors where name = $1), $2, $3);
		`,
		d.UIDname, d.PID, d.Val,
	)
	if err != nil {
		return err
	}

	_, err = db.Exec(
		ctx,
		`
		update investors
			set pids = array_append(investors.pids, $1)
			where id = (select id from investors where name = $2);
		`,
		d.PID, d.UIDname,
	)

	return err
}

func (d *Investment) SaveMock(ctx context.Context, db *pgx.Conn) error {
	return nil
}

func GetProjectInvestments(ctx context.Context, db *pgx.Conn, pid int) ([]Investment, error) {
	rows, err := db.Query(
		ctx,
		`
		SELECT 
			(select name from investors as inv where inv.id = im.uid),
			coalesce(pid, 0), 
			coalesce(val, 0)
			FROM investments as im
			WHERE pid = $1;	
		`,
		pid,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var investments []Investment
	for rows.Next() {
		var investment Investment
		err = rows.Scan(
			&investment.UIDname,
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

func GetUserInvestments(ctx context.Context, db *pgx.Conn, uid int) ([]Investment, error) {
	rows, err := db.Query(
		ctx,
		`
		SELECT uid, pid, val
			FROM investments
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
			UID: 1,
			Val: 100,
		},
		{
			UID: 2,
			Val: 250,
		},
		{
			UID: 3,
			Val: 100,
		},
		{
			UID: 4,
			Val: 1000,
		},
		{
			UID: 5,
			Val: 50,
		},
	}, nil
}
