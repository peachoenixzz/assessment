package expense

import (
	"database/sql"
	"github.com/lib/pq"
	"github.com/peachoenixz/assessment/pkg/log"
)

func NewPostgres(client *sql.DB) Repo {
	createTable(client)
	return &PostgresRepo{Client: client}
}

type PostgresRepo struct {
	Client *sql.DB
}

func createTable(client *sql.DB) {

	createTb := `
	CREATE TABLE IF NOT EXISTS expenses (
		id SERIAL PRIMARY KEY,
		title TEXT,
		amount FLOAT,
		note TEXT,
		tags TEXT[]
	);
	`
	_, err := client.Exec(createTb)

	if err != nil {
		log.ErrorLog(err, "Database Postgres")
	}

	log.InfoLog("success create table expenses or exists", "Database Postgres")
}

func (r PostgresRepo) InsertExpense(req Request) (int, error) {
	var id int
	err := r.Client.QueryRow("INSERT INTO expenses (title,amount,note,tags) values ($1,$2,$3,$4) RETURNING id", req.Title, req.Amount, req.Note, pq.Array(req.Tags)).Scan(&id)
	if err != nil {
		return id, err
	}
	return id, nil
}
