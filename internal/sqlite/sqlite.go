package sqlite

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

func New(persistLocation string) (*sql.DB, error) {
	ploc := ":memory:"
	if persistLocation != "" {
		ploc = persistLocation
	}
	db, err := sql.Open("sqlite", ploc)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(1)

	return db, nil
}
