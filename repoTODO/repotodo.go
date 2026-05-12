package repotodo

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

type TODO struct {
	ID          int        `json:"id", db:"id"`
	Task        string     `json:"task", db:"task"`
	Done        bool       `json:"done", db:"done"`
	CreatedAT   *time.Time `json:"created_at", db:"created_at"`
	UpdatedAT   *time.Time `json:"updated_at", db:"updated_at"`
	Description *string    `json:"description", db:"description"`
}

func InsertTODO(db *sql.DB, todo *TODO) (int, error) {
	var newIDtask int
	query := `insert into todo(task, done, created_at, description) values ($1, $2, now(), $3) returning id`
	err := db.QueryRow(query, todo.Task, todo.Done, todo.Description).Scan(&newIDtask)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return newIDtask, nil
}

func DeleteTODO(db *sql.DB, id int) error {
	res, err := db.Exec("DELETE FROM todo WHERE id = $1", id)
	if err != nil {
		log.Println(err)
		return err
	}

	rowAffected, err := res.RowsAffected()
	if err != nil {
		log.Println(err)
		return err
	}

	if rowAffected == 0 {
		return fmt.Errorf("No rows deleted by id %d", id)
	}
	return nil
}

func UpdateTODOByID(db *sql.DB, todo *TODO) (*TODO, error) {
	var updatedTODO TODO

	query := `update todo set task = $2, done = $3, description = $4, updated_at = now() where id = $1`

	err := db.QueryRow(query, todo.ID, todo.Task, todo.Done, todo.Description).Scan(&updatedTODO)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &updatedTODO, nil
}

func GetALLTODO(db *sql.DB) ([]TODO, error) {
	query := `select id, task, done, created_at, updated_at, description from todo`

	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
	}

	todoslice := []TODO{}

	for rows.Next() {
		var todo TODO
		err := rows.Scan(&todo.ID, &todo.Task, &todo.Done, &todo.CreatedAT, &todo.UpdatedAT, &todo.Description)
		if err != nil {

			log.Println(err)
			return nil, err
		}

		todoslice = append(todoslice, todo)

	}
	return todoslice, nil
}

func GetTODOByID(db *sql.DB, id int) (*TODO, error) {
	var task TODO

	query := `select * from todo where id = $1`
	err := db.QueryRow(query, id).Scan(&task.ID, &task.Task, &task.Done, &task.CreatedAT, &task.UpdatedAT, &task.Description)
	if err != nil {
		log.Println(err, "ОШИБКА функции? СТРОКА 88 в репотуду")
		return nil, err
	}
	return &task, nil
}
