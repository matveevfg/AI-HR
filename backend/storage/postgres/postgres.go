package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/matveevfg/AI-HR/backend/storage/postgres/migrations"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
	"github.com/uptrace/bun/migrate"
)

const (
	defaultTimeout = 30 * time.Second
)

type Postgres struct {
	d *bun.DB
}

func New(host, port, database, user, pass, envName string) (*Postgres, error) {
	sqldb := sql.OpenDB(pgdriver.NewConnector(
		pgdriver.WithAddr(host+":"+port),
		pgdriver.WithUser(user),
		pgdriver.WithInsecure(true),
		pgdriver.WithPassword(pass),
		pgdriver.WithDatabase(database),
		pgdriver.WithApplicationName("service"),
		pgdriver.WithTimeout(defaultTimeout),
	))

	db := bun.NewDB(sqldb, pgdialect.New(), bun.WithDiscardUnknownColumns())
	db.SetMaxOpenConns(10)
	if envName == "dev" {
		db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))
	}

	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("ping: %w", err)
	}

	return &Postgres{d: db}, nil
}

func (p *Postgres) Migrate() error {
	if p.d == nil {
		return fmt.Errorf("database is not initialized")
	}

	migrator, err := runInitCommand(p.d)
	if err != nil {
		return err
	}

	_, err = migrator.Migrate(context.Background())
	return err
}

// runInitCommand initializes new migrator.
func runInitCommand(db *bun.DB) (*migrate.Migrator, error) {
	migrator := migrate.NewMigrator(db, migrations.MigrationSet, migrate.WithMarkAppliedOnSuccess(true))
	if err := migrator.Init(context.Background()); err != nil {
		return nil, err
	}

	return migrator, nil
}

func (p *Postgres) Close() error {
	if err := p.d.Close(); err != nil {
		return err
	}
	return nil
}
