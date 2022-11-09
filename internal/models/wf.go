package models

import (
	"context"

	"github.com/jackc/pgx/v4"
)

type WorkFee struct {
	Id  int `json:"id"`
	Pid int `json:"pid"`
	Cid int `json:"cid"`
	Sum int `json:"sum"`
	Dcm Doc `json:"doc"`
}

func (wf WorkFee) Save(ctx context.Context, db *pgx.Conn) error {
	_, err := db.Exec(
		ctx,
		`
		insert into wfs
			(pid, cid, sum)
			values ($1, $2, $3);		
		`,
		wf.Pid, wf.Cid, wf.Sum,
	)
	return err
}

func (wf WorkFee) SaveMock(ctx context.Context, db *pgx.Conn) error {
	return nil
}

// если pid, то все переводы исоплнителям по проекту
// если cid, то все переводы исполнителю по всем проектам
func GetWFs(ctx context.Context, db *pgx.Conn, pid, cid int) ([]WorkFee, error) {
	rows, err := db.Query(
		ctx,
		`
		
		`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var slc []WorkFee
	for rows.Next() {
		var ob WorkFee
		err = rows.Scan(
		// todo
		)
		if err != nil {
			return nil, err
		}

		slc = append(slc, ob)
	}

	return slc, nil
}

func GetWFsMock(ctx context.Context, db *pgx.Conn, pid, cid int) ([]WorkFee, error) {
	return []WorkFee{
		{
			Id:  0,
			Pid: pid,
			Cid: 0,
			Sum: 100,
			Dcm: Doc{
				ID:    0,
				Title: "doc 0",
			},
		},
		{
			Id:  1,
			Pid: pid,
			Cid: 1,
			Sum: 200,
			Dcm: Doc{
				ID:    1,
				Title: "doc 1",
			},
		},
		{
			Id:  2,
			Pid: pid,
			Cid: 2,
			Sum: 400,
			Dcm: Doc{
				ID:    2,
				Title: "doc 2",
			},
		},
	}, nil
}
