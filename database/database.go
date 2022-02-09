package database

import (
	"github.com/jmoiron/sqlx"
	_ "modernc.org/sqlite"
)

type DB struct {
	*sqlx.DB
	Words WordStorage
}

func Open(filepath string) (*DB, error) {
	db, err := sqlx.Open("sqlite", filepath)
	if err != nil {
		return nil, err
	}

	return &DB{
		DB:    db,
		Words: &Words{DB: db},
	}, nil
}
