package aiHr

import (
	"context"

	"github.com/google/uuid"
	"github.com/matveevfg/AI-HR/backend/api/requests"
	"github.com/matveevfg/AI-HR/backend/models"
)

type storage interface {
	Vacancy(ctx context.Context, id uuid.UUID) (*models.Vacancy, error)
	Vacancies(ctx context.Context, filter requests.VacancyFilter) ([]*models.Vacancy, error)
	SaveVacancy(ctx context.Context, vacancy *models.Vacancy) error
	DeleteVacancy(ctx context.Context, id uuid.UUID) error
	SetVacancyInactive(ctx context.Context, id uuid.UUID) error
	SetVacancyActive(ctx context.Context, id uuid.UUID) error
}
