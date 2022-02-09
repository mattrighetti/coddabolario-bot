package database

import (
	"github.com/jmoiron/sqlx"
)

type (
	WordStorage interface {
		Create(word, definition string) error
		GetWords(word string) ([]Word, error)
	}

	Words struct {
		*sqlx.DB
	}

	Word struct {
		ID         int32  `db:"id"`
		Word       string `db:"word"`
		Definition string `db:"definition"`
	}
)

func (db *Words) Create(word, definition string) error {
	const q = `INSERT INTO coddabolario`
	_, err := db.Exec(q, word, definition)
	return err
}

func (db *Words) GetWords(word string) ([]Word, error) {
	var words = []Word{}
	const q = `SELECT id, word, definition FROM coddabolario WHERE word LIKE ? LIMIT 10`
	err := db.Select(&words, q, word+"%")
	return words, err
}
