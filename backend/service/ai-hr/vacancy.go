package aiHr

import (
	"context"

	"github.com/google/uuid"
	"github.com/matveevfg/AI-HR/backend/api/requests"
	"github.com/matveevfg/AI-HR/backend/models"
)

func (s *Service) Vacancy(ctx context.Context, id uuid.UUID) (*models.Vacancy, error) {
	return s.storage.Vacancy(ctx, id)
}

func (s *Service) Vacancies(ctx context.Context, filter requests.VacancyFilter) ([]*models.Vacancy, error) {
	return s.storage.Vacancies(ctx, filter)
}

func (s *Service) SaveVacancy(ctx context.Context, vacancy *models.Vacancy) (*uuid.UUID, error) {
	if vacancy.ID == uuid.Nil {
		vacancy.ID = uuid.New()
	}

	if err := s.storage.SaveVacancy(ctx, vacancy); err != nil {
		return nil, err
	}

	return &vacancy.ID, nil
}

func (s *Service) DeleteVacancy(ctx context.Context, id uuid.UUID) error {
	return s.storage.DeleteVacancy(ctx, id)
}

func (s *Service) SetVacancyInactive(ctx context.Context, id uuid.UUID) error {
	return s.storage.SetVacancyInactive(ctx, id)
}

func (s *Service) SetVacancyActive(ctx context.Context, id uuid.UUID) error {
	return s.storage.SetVacancyActive(ctx, id)
}
