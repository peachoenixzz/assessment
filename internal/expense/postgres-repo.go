package expense

import (
	"database/sql"
	"fmt"

	"github.com/lib/pq"
	"github.com/peachoenixz/assessment/pkg/log"
)

func NewPostgres(client *sql.DB) Repo {
	createTable(client)
	return &PostgresRepo{Client: client}
}

func NewPostgresMock(Client *sql.DB) Repo {
	return &PostgresRepo{Client: Client}
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

func (r PostgresRepo) InsertExpense(req Request) (Response, error) {
	var res Response
	err := r.Client.QueryRow("INSERT INTO expenses (title,amount,note,tags) values ($1,$2,$3,$4) RETURNING title,amount,note,tags,id", req.Title, req.Amount, req.Note, pq.Array(req.Tags)).
		Scan(&res.Title, &res.Amount, &res.Note, pq.Array(&res.Tags), &res.ID)
	if err != nil {
		return Response{}, err
	}
	return res, nil
}

func (r PostgresRepo) GetExpenseByID(id string) (Response, error) {
	var res Response
	stmt, err := r.Client.Prepare("SELECT id,title,amount,note,tags FROM expenses WHERE id=$1")
	if err != nil {
		return Response{}, err
	}
	err = stmt.QueryRow(id).Scan(&res.ID, &res.Title, &res.Amount, &res.Note, pq.Array(&res.Tags))
	if err != nil {
		return Response{}, err
	}
	return res, nil
}

func (r PostgresRepo) GetExpense() ([]Response, error) {
	var responses []Response
	stmt, err := r.Client.Prepare("SELECT id,title,amount,note,tags FROM expenses ORDER BY id ASC")
	if err != nil {
		return []Response{}, err
	}

	rows, err := stmt.Query()
	if err != nil {
		return []Response{}, err
	}

	for rows.Next() {
		var res Response
		err = rows.Scan(&res.ID, &res.Title, &res.Amount, &res.Note, pq.Array(&res.Tags))
		if err != nil {
			return []Response{}, err
		}
		responses = append(responses, res)
	}
	fmt.Println(responses)
	return responses, nil
}

func (r PostgresRepo) UpdateExpenseByID(req Request, id string) (Response, error) {
	var res Response
	stmt, err := r.Client.Prepare("UPDATE expenses SET title=$1,amount=$2,note=$3,tags=$4 WHERE id=$5 RETURNING title,amount,note,tags,id")
	if err != nil {
		log.ErrorLog(err.Error(), "DATABASE")
		return Response{}, err
	}
	err = stmt.QueryRow(req.Title, req.Amount, req.Note, pq.Array(&req.Tags), id).Scan(&res.Title, &res.Amount, &res.Note, pq.Array(&res.Tags), &res.ID)
	if err != nil {
		log.ErrorLog(err.Error(), "DATABASE QUERY ROW")
		return Response{}, err
	}
	return res, nil
}
