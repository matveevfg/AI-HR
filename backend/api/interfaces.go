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
	SaveVacancy(ctx context.Context, file []*multipart.FileHeader) error
	DeleteVacancy(ctx context.Context, id uuid.UUID) error
	SetVacancyInactive(ctx context.Context, id uuid.UUID) error
	SetVacancyActive(ctx context.Context, id uuid.UUID) error

	SaveResume(ctx context.Context, file []*multipart.FileHeader, vacancyID uuid.UUID) error
	Resumes(ctx context.Context, vacancyID uuid.UUID) ([]*models.Resume, error)
}
