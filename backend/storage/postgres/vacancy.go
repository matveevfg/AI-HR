package postgres

import (
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/matveevfg/AI-HR/backend/api/requests"
	"github.com/matveevfg/AI-HR/backend/models"
	"github.com/uptrace/bun"

	"context"
)

func (p *Postgres) Vacancies(ctx context.Context, filter requests.VacancyFilter) ([]*models.Vacancy, error) {
	vacancies := make([]*models.Vacancy, 0)

	q := p.d.NewSelect().Model(&vacancies)

	if filter.Status != nil {
		q.Where("status = ?", *filter.Status)
	}
	if filter.Name != nil {
		q.Where("name LIKE '%?%'", bun.Safe(*filter.Name))
	}

	if err := q.Scan(ctx); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return vacancies, nil
		}

		return nil, err
	}

	return vacancies, nil
}

func (p *Postgres) Vacancy(ctx context.Context, id uuid.UUID) (*models.Vacancy, error) {
	var vacancy models.Vacancy

	if err := p.d.NewSelect().Model(&vacancy).Where("id = ?", id).Scan(ctx); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return &vacancy, nil
}

func (p *Postgres) SaveVacancy(ctx context.Context, vacancy *models.Vacancy) error {
	_, err := p.d.NewInsert().
		Model(vacancy).
		On("CONFLICT (id) DO UPDATE").
		Set("status = EXCLUDED.status").
		Set("name = EXCLUDED.name").
		Set("region = EXCLUDED.region").
		Set("city = EXCLUDED.city").
		Set("address = EXCLUDED.address").
		Set("contract_type = EXCLUDED.contract_type").
		Set("employment_type = EXCLUDED.employment_type").
		Set("work_schedule = EXCLUDED.work_schedule").
		Set("income = EXCLUDED.income").
		Set("salary_max = EXCLUDED.salary_max").
		Set("salary_min = EXCLUDED.salary_min").
		Set("pronounces = EXCLUDED.pronounces").
		Set("responsibilities = EXCLUDED.responsibilities").
		Set("requirements = EXCLUDED.requirements").
		Set("education = EXCLUDED.education").
		Set("experience = EXCLUDED.experience").
		Set("special_programs = EXCLUDED.special_programs").
		Set("computer_skills = EXCLUDED.computer_skills").
		Set("foreign_languages = EXCLUDED.foreign_languages").
		Set("language_level = EXCLUDED.language_level").
		Set("has_business_trips = EXCLUDED.has_business_trips").
		Set("additional_information = EXCLUDED.additional_information").
		Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (p *Postgres) DeleteVacancy(ctx context.Context, id uuid.UUID) error {
	vacancy := models.Vacancy{
		ID: id,
	}

	_, err := p.d.NewDelete().Model(&vacancy).WherePK().Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (p *Postgres) SetVacancyInactive(ctx context.Context, id uuid.UUID) error {
	var vacancy models.Vacancy

	_, err := p.d.NewUpdate().Model(&vacancy).Where("id = ?", id).Set("status = 'inactive'").Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (p *Postgres) SetVacancyActive(ctx context.Context, id uuid.UUID) error {
	var vacancy models.Vacancy

	_, err := p.d.NewUpdate().Model(&vacancy).Where("id = ?", id).Set("status = 'active'").Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}
