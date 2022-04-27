package sqlite

import (
	"database/sql"
	"embed"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
	"lnpay/internal/config"
	"lnpay/internal/data"
)

type sqlite struct {
	db *sql.DB
}

//go:embed migrations/*.sql
var fs embed.FS

func New(cfg config.SQLite) (data.Database, error) {
	db, err := sql.Open("sqlite3", cfg.Path)
	if err != nil {
		return nil, errors.Wrap(err, "failed to open sqlite database")
	}

	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		return nil, errors.Wrap(err, "failed to create sqlite driver")
	}

	d, err := iofs.New(fs, "migrations")
	if err != nil {
		return nil, errors.Wrap(err, "failed to embed migrations")
	}

	m, err := migrate.NewWithInstance(
		"iofs", d,
		"sqlite3", driver)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create migrations")
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return nil, errors.Wrap(err, "failed to migrate database")
	}

	return &sqlite{db: db}, nil
}
