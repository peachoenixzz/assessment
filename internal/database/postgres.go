package database

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
	"github.com/peachoenixz/assessment/pkg/log"
)

type PostgresDatabase struct {
	Client *sql.DB
}

func NewPostgres() *PostgresDatabase {
	url := os.Getenv("DATABASE_URL")
	pg, err := sql.Open("postgres", url)
	if err != nil {
		log.ErrorLog("Error connect", "Database postgres")
	}
	return &PostgresDatabase{pg}
}
