package models

import (
	"context"

	"github.com/jackc/pgx/v4"
)

type Doc struct {
	ID     int    `json:"id"`
	Pid    int    `json:"pid"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Dcm    string `json:"dcm"`
	Cid    int    `json:"cid"`
}

func (d Doc) Save(ctx context.Context, db *pgx.Conn) error {
	_, err := db.Exec(
		ctx,
		`
		insert into docs
			(title, author, document, pid, cid)
			values ($1, $2, $3, $4, $5);	
		`,
		d.Title, d.Author, d.Dcm, d.Pid, d.Cid,
	)
	if err != nil {
		return err
	}

	_, err = db.Exec(
		ctx,
		`
		update contracotrs
			set pids = array_append(contractors.pids, $1)
			where id = $2;
		`,
		d.Pid, d.Cid,
	)
	return err
}

func (d Doc) SaveMock(ctx context.Context, db *pgx.Conn) error {
	return nil
}

func GetDocs(ctx context.Context, db *pgx.Conn, pid int) ([]Doc, error) {
	rows, err := db.Query(
		ctx,
		`
			select id, pid, title, author, document, cid
				from docs
				where (pid = $1 or $1 = 0);
		`,
		pid,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var slc []Doc
	for rows.Next() {
		var ob Doc
		err = rows.Scan(
			&ob.ID,
			&ob.Pid,
			&ob.Title,
			&ob.Author,
			&ob.Dcm,
			&ob.Cid,
		)
		if err != nil {
			return nil, err
		}

		slc = append(slc, ob)
	}

	return slc, nil
}

func GetDocsMock(ctx context.Context, db *pgx.Conn, pid int) ([]Doc, error) {
	return []Doc{
		{
			ID:     0,
			Pid:    pid,
			Title:  "title 0",
			Author: "author",
			Dcm:    "текст документа",
			Cid:    1,
		},
		{
			ID:     1,
			Pid:    pid,
			Title:  "title 1",
			Author: "author",
			Dcm:    "текст документа",
			Cid:    1,
		},
		{
			ID:     2,
			Pid:    pid,
			Title:  "title 2",
			Author: "author",
			Dcm:    "текст документа текст документа текст документа текст документа текст документа текст документа",
			Cid:    4,
		},
	}, nil
}
