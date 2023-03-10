package database

import (
	"database/sql"
	"errors"
	"os"

	_ "github.com/lib/pq"
)

type PostgresDatabase struct {
	Client *sql.DB
}

func NewPostgres() (*PostgresDatabase, error) {
	url := os.Getenv("DATABASE_URL")
	if url == "" {
		return nil, errors.New("invalid database URL")
	}
	pg, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	return &PostgresDatabase{pg}, nil
}
