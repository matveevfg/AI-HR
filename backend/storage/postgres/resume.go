package postgres

import (
	"context"
	"errors"

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
