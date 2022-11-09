package models

import (
	"context"

	"github.com/jackc/pgx/v4"
)

type Risk struct {
	PID  int    `json:"pid"`
	Rsk  string `json:"risk"`
	Plan string `json:"plan"`
	Sum  int    `json:"sum"`
}

func (r Risk) Save(ctx context.Context, db *pgx.Conn) error {
	_, err := db.Exec(
		ctx,
		`
		insert into risks
			(risk, plan, pid, sum)	
			values ($1, $2, $3, $4);
		`,
		r.Rsk, r.Plan, r.PID, r.Sum,
	)
	return err
}

func (r Risk) SaveMock(ctx context.Context, db *pgx.Conn) error {
	return nil
}

func GetRisks(ctx context.Context, db *pgx.Conn, pid int) ([]Risk, error) {
	rows, err := db.Query(
		ctx,
		`
		
		`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var slc []Risk
	for rows.Next() {
		var ob Risk
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

func GetRisksMock(ctx context.Context, db *pgx.Conn, pid int) ([]Risk, error) {
	return []Risk{
		{
			PID:  pid,
			Rsk:  "потеряны все полимеры",
			Plan: "всех призвать, наказать кого надо, все вернуть",
			Sum:  1000000,
		},
		{
			PID:  pid,
			Rsk:  "качество продукта страдает",
			Plan: "выявить дефекты, проработать пути решения",
			Sum:  200000,
		},
	}, nil
}
