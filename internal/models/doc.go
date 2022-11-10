package models

import (
	"context"

	"github.com/jackc/pgx/v4"
)

type Doc struct {
	ID    int    `json:"id"`
	Pid   int    `json:"pid"`
	Title string `json:"title"`
	Dcm   string `json:"dcm"`
	Cid   string `json:"cid"`
}

func (d Doc) Save(ctx context.Context, db *pgx.Conn) error {
	_, err := db.Exec(
		ctx,
		`
		insert into docs
			(title, document, pid, cid)
			values ($1, $2, $3, (select id from contractors where name = $4));	
		`,
		d.Title, d.Dcm, d.Pid, d.Cid,
	)
	if err != nil {
		return err
	}

	_, err = db.Exec(
		ctx,
		`
		update contractors
			set pids = array_append(contractors.pids, $1)
			where name = $2;
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
			select d.id, d.pid, d.title, d.document, (select name from contractors as c where c.id = d.cid)
				from docs as d
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
			ID:    0,
			Pid:   pid,
			Title: "title 0",
			Dcm:   "текст документа",
			Cid:   "isp1",
		},
		{
			ID:    1,
			Pid:   pid,
			Title: "title 1",
			Dcm:   "текст документа",
			Cid:   "isp2",
		},
		{
			ID:    2,
			Pid:   pid,
			Title: "title 2",
			Dcm:   "текст документа текст документа текст документа текст документа текст документа текст документа",
			Cid:   "isp3",
		},
	}, nil
}
