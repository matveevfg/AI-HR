package aiHr

import (
	"context"
	"encoding/json"
	"io"
	"mime/multipart"
	"os"

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

func (s *Service) SaveVacancy(ctx context.Context, files []*multipart.FileHeader) error {
	for _, file := range files {
		src, err := file.Open()
		if err != nil {
			return err
		}

		fileData, err := io.ReadAll(src)
		if err != nil {
			return err
		}

		tmp, err := os.Create(os.TempDir() + "\\" + file.Filename)
		if err != nil {
			return err
		}

		if err := os.WriteFile(tmp.Name(), fileData, 0666); err != nil {
			return err
		}

		text, err := extractTextFromDocx(tmp.Name())
		if err != nil {
			return err
		}

		vacancy, err := s.llmClient.VacancyToJSON(ctx, text)
		if err != nil {
			return err
		}

		vacansyJSON, err := json.Marshal(vacancy)
		if err != nil {
			return err
		}

		plan, err := s.llmClient.VacancyInterviewPlan(ctx, string(vacansyJSON))
		if err != nil {
			return err
		}

		vacancy.ID = uuid.New()
		vacancy.Status = "active"

		if err := s.storage.SaveVacancy(ctx, vacancy); err != nil {
			return err
		}

		vacancyPlan := &models.Plan{
			ID:        uuid.New(),
			VacancyID: vacancy.ID,
			Plan:      plan,
		}

		if err := s.storage.SavePlan(ctx, vacancyPlan); err != nil {
			return err
		}

		if err := src.Close(); err != nil {
			return err
		}

		name := tmp.Name()

		if err := tmp.Close(); err != nil {
			return err
		}

		if err := os.Remove(name); err != nil {
			return err
		}
	}

	return nil
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
