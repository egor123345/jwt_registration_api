package postgres_client

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func NewClient(connString string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
