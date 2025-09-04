package migrations

import (
	"context"

	"github.com/uptrace/bun"
)

func init() {
	MigrationSet.MustRegister(
		func(ctx context.Context, db *bun.DB) error {
			_, err := db.ExecContext(ctx, `
				CREATE TABLE IF NOT EXISTS resumes (
					id UUID PRIMARY KEY,
					vacancy_id UUID REFERENCES vacancies(id) ,
					location TEXT,
					citizenship TEXT,
					relocation_willingness TEXT,
					business_trips_willingness TEXT,
					desired_position TEXT,
					specializations TEXT[],
					employment_types TEXT[],
					work_schedules TEXT[],
					max_commute_time TEXT,
					total_experience TEXT,
					education TEXT[],
					courses TEXT[],
					languages TEXT[],
					skills TEXT[],
					has_car BOOLEAN,
					license_category TEXT,
					about TEXT
				);
				
				CREATE TABLE IF NOT EXISTS work_places (
					id SERIAL PRIMARY KEY,
					resume_id UUID NOT NULL REFERENCES resumes(id),
					company TEXT,
					position TEXT,
					period TEXT,
					duration TEXT,
					responsibilities TEXT[]
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
