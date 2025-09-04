package migrations

import (
	"context"

	"github.com/uptrace/bun"
)

func init() {
	MigrationSet.MustRegister(
		func(ctx context.Context, db *bun.DB) error {
			_, err := db.ExecContext(ctx, `
				CREATE TABLE IF NOT EXISTS vacancies (
					id UUID PRIMARY KEY,
					status VARCHAR(255) NOT NULL,
					name VARCHAR(255) NOT NULL,
					region VARCHAR(255) NOT NULL,
					city VARCHAR(255) NOT NULL,
					address TEXT NOT NULL,
					work_type VARCHAR(255) NOT NULL,
					contract_type VARCHAR(255) NOT NULL,
					employment_type VARCHAR(255) NOT NULL,
					work_schedule VARCHAR(255) NOT NULL,
					income INTEGER NULL,
					salary_max INTEGER NULL,
					salary_min INTEGER NULL,
					pronounces TEXT NOT NULL,
					responsibilities TEXT NOT NULL,
					requirements TEXT NOT NULL,
					education TEXT NOT NULL,
					experience TEXT NOT NULL,
					special_programs BOOLEAN,
					computer_skills BOOLEAN,
					foreign_languages BOOLEAN,
					language_level VARCHAR(255),
					has_business_trips BOOLEAN,
					additional_information TEXT,
					created_at TIMESTAMPTZ,
					updated_at TIMESTAMPTZ
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
