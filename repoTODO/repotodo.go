package repotodo

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

type TODO struct {
	ID        int        `json:"id", db:"id"`
	Task      string     `json:"task", db:"task"`
	Done      bool       `json:"done", db:"done"`
	CreatedAT *time.Time `json:"created_at", db:"created_at"`
	UpdatedAT *time.Time `json:"updated_at", db:"updated_at"`
}

func InsertTODO(db *sql.DB, todo *TODO) (int, error) {
	var newIDtask int
	query := `insert into todo(task, done, created_at) values ($1, $2, now()) returning id`
	err := db.QueryRow(query, todo.Task, todo.Done).Scan(&newIDtask)
	if err != nil {
		log.Println(err)

	}
	return newIDtask, nil
}

func DeleteTODO(db *sql.DB, id int) {
	res, err := db.Exec("DELETE FROM todo WHERE id = $1", id)
	fmt.Println(res)
	if err != nil {
		log.Println(err)
	}
}

func UpdateTODO(db *sql.DB, todo *TODO) {
	query := `update todo set task = $2, updated_at = now() where id = $1`
	err := db.QueryRow(query, todo.ID, todo.Task).Err()
	if err != nil {
		log.Println(err)
	}
}

func ReadyTODO(db *sql.DB, todo *TODO) {
	query := `update todo set done = $2 where id = $1`
	err := db.QueryRow(query, todo.ID, todo.Done).Err()
	if err != nil {
		log.Println(err)
	}
}

func ShowALLTODO(db *sql.DB) ([]TODO, error) {
	query := `select id, task, done, created_at, updated_at from todo`

	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
	}

	todoslice := []TODO{}

	for rows.Next() {
		var todo TODO
		err := rows.Scan(&todo.ID, &todo.Task, &todo.Done, &todo.CreatedAT, &todo.UpdatedAT)

		todoslice = append(todoslice, todo)
		if err != nil {

			log.Println(err)
			return nil, err
		}

	}
	return todoslice, nil
}
