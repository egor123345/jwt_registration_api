package composites

import (
	"database/sql"
	postgres_client "jwt_registration_api/pkg/client/postgres"
)

type PostgresComposite struct {
	Db *sql.DB
}

func NewPostgresComposite(connString string) (*PostgresComposite, error) {
	db, err := postgres_client.NewClient(connString)
	if err != nil {
		return nil, err
	}

	return &PostgresComposite{Db: db}, nil
}
