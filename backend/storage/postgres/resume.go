package postgres

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"

	"github.com/matveevfg/AI-HR/backend/models"
)

func (p *Postgres) SaveResume(ctx context.Context, resume *models.Resume) error {
	tx, ok := txFromCtx(ctx)
	if !ok {
		return errors.New("could not get tx from context")
	}

	_, err := tx.NewInsert().Model(resume).Ignore().Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (p *Postgres) SaveWorkPlaces(ctx context.Context, workPlaces []*models.WorkPlace) error {
	tx, ok := txFromCtx(ctx)
	if !ok {
		return errors.New("could not get tx from context")
	}

	_, err := tx.NewInsert().Model(workPlaces).ExcludeColumn("id").Ignore().Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (p *Postgres) Resumes(ctx context.Context, vacancyID uuid.UUID) ([]*models.Resume, error) {
	var resumes []*models.Resume

	if err := p.d.NewSelect().Model(&resumes).Where("vacancy_id = ?", vacancyID).Scan(ctx); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return resumes, nil
		}

		return nil, err
	}

	return resumes, nil
}
