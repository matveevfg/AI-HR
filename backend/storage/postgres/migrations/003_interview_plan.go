package migrations

import (
	"context"

	"github.com/uptrace/bun"
)

func init() {
	MigrationSet.MustRegister(
		func(ctx context.Context, db *bun.DB) error {
			_, err := db.ExecContext(ctx, `
				CREATE TABLE IF NOT EXISTS plans (
					id UUID PRIMARY KEY,
					vacancy_id UUID REFERENCES vacancies(id) ,
					plan TEXT NOT NULL
				);
			`,
			)
			return err
		},
		func(ctx context.Context, db *bun.DB) error {
			_, err := db.ExecContext(ctx, ``)
			return err
		})
}
