package migrations

import "github.com/uptrace/bun/migrate"

var MigrationSet = migrate.NewMigrations(migrate.WithMigrationsDirectory("migrations"))
