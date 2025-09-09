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

	SavePlan(ctx context.Context, plan *models.Plan) error

	SaveResume(ctx context.Context, resume *models.Resume) error
	SaveWorkPlaces(ctx context.Context, workPlaces []*models.WorkPlace) error
	Resumes(ctx context.Context, vacancyID uuid.UUID) ([]*models.Resume, error)

	CtxWithTx(ctx context.Context) (context.Context, error)
	TxCommit(ctx context.Context) error
	TxRollback(ctx context.Context) error
}

type llmClient interface {
	ResumeToJSON(ctx context.Context, resumeText string) (*models.Resume, error)
	VacancyToJSON(ctx context.Context, vacancyText string) (*models.Vacancy, error)
	VacancyInterviewPlan(ctx context.Context, vacancy string) (string, error)
}

type transcriptionClient interface {
}
