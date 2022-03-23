package sqlite

import (
	"context"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
	"lnpay/internal/config"
	"lnpay/internal/data"
)

type sqlite struct {
	db *sql.DB
}

func New(cfg config.SQLite) (data.Database, error) {
	db, err := sql.Open("sqlite3", cfg.Path)
	if err != nil {
		return nil, errors.Wrap(err, "failed to open sqlite database")
	}

	return &sqlite{db: db}, nil
}

func (s *sqlite) GetVersion(ctx context.Context) (int, error) {
	var version int
	err := s.db.QueryRowContext(ctx, "SELECT version FROM version").Scan(&version)
	if err != nil {
		return 0, errors.Wrap(err, "failed to get version")
	}

	return version, nil
}
