package postgres

import (
	"database/sql"
	"log"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPosgresRepository() (*PostgresRepository, error) {
	db, err := sql.Open("pgx", "")
	if err != nil {
		log.Printf("Error while connect to posgres %v\n", err)
		return nil, err
	}

	return &PostgresRepository{
		db: db,
	}, nil
}
