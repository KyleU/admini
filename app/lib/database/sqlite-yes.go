//go:build darwin || (!android && linux && 386) || (!android && linux && amd64) || (!android && linux && arm) || (!android && linux && arm64) || (!android && linux && riscv64) || (windows && amd64)

// Package database - Content managed by Project Forge, see [projectforge.md] for details.
package database

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	_ "modernc.org/sqlite" // load sqlite driver.

	"admini.dev/admini/app/lib/telemetry"
	"admini.dev/admini/app/util"
)

const SQLiteEnabled = true

var typeSQLite = &DBType{Key: "sqlite", Title: "SQLite", Quote: `"`, Placeholder: "$", SupportsReturning: true}

func OpenSQLiteDatabase(ctx context.Context, key string, params *SQLiteParams, logger util.Logger) (*Service, error) {
	_, span, logger := telemetry.StartSpan(ctx, "database:open", logger)
	defer span.Complete()
	if params.File == "" {
		return nil, errors.New("need filename for SQLite database")
	}
	u := fmt.Sprintf("file:%s?cache=shared&_journal=WAL&_timeout=10000&_foreign_keys=on", params.File)
	db, err := sqlx.Open("sqlite", u)
	if err != nil {
		return nil, errors.Wrap(err, "error opening database")
	}
	_, err = db.Exec("PRAGMA foreign_keys = ON")
	if err != nil {
		return nil, errors.Wrap(err, "error enabling foreign keys")
	}
	return NewService(typeSQLite, key, key, params.Schema, "sqlite", params.Debug, db, logger)
}
