package models

import (
	"context"

	"github.com/jackc/pgx/v4"
)

type Expertise struct {
	ID      int    `json:"id"`
	Pid     int    `json:"pid"`
	Content string `json:"content"`
}

func (exp Expertise) Save(ctx context.Context, db *pgx.Conn) error {
	_, err := db.Exec(
		ctx,
		`
		insert into expertises
			(pid, content)
			values ($1, $2);	
		`,
		exp.Pid, exp.Content,
	)
	return err
}

func (exp Expertise) SaveMock(ctx context.Context, db *pgx.Conn) error {
	return nil
}

func GetExpertises(ctx context.Context, db *pgx.Conn, pid int) ([]Expertise, error) {
	rows, err := db.Query(
		ctx,
		`
		select id, pid, content
			from expertises
			where pid = $1;
		`,
		pid,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var slc []Expertise
	for rows.Next() {
		var ob Expertise
		err = rows.Scan(
			&ob.ID,
			&ob.Pid,
			&ob.Content,
		)
		if err != nil {
			return nil, err
		}

		slc = append(slc, ob)
	}

	return slc, nil
}

func GetExpertisesMock(ctx context.Context, db *pgx.Conn, pid int) ([]Expertise, error) {
	return []Expertise{
		{
			ID:      0,
			Pid:     pid,
			Content: "эксперт экспертиза эксперт экспертиза эксперт экспертиза эксперт экспертиза эксперт экспертиза эксперт экспертиза эксперт экспертиза эксперт экспертиза эксперт экспертиза эксперт экспертиза ",
		},
	}, nil
}
