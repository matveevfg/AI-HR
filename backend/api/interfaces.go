package api

import (
	"context"
	"mime/multipart"

	"github.com/google/uuid"

	"github.com/matveevfg/AI-HR/backend/api/requests"
	"github.com/matveevfg/AI-HR/backend/models"
)

type AiHrService interface {
	Vacancy(ctx context.Context, id uuid.UUID) (*models.Vacancy, error)
	Vacancies(ctx context.Context, filter requests.VacancyFilter) ([]*models.Vacancy, error)
	SaveVacancy(ctx context.Context, vacancy *models.Vacancy) (*uuid.UUID, error)
	DeleteVacancy(ctx context.Context, id uuid.UUID) error
	SetVacancyInactive(ctx context.Context, id uuid.UUID) error
	SetVacancyActive(ctx context.Context, id uuid.UUID) error

	SaveResume(ctx context.Context, file *multipart.FileHeader) error
}
