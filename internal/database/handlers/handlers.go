package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	repotodo "todo/RepoTODO"
)

func HiHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Дороу залупа, это работает"))
}

func ShowAllTODO(DB *sql.DB, w http.ResponseWriter, r *http.Request) ([]repotodo.TODO, error) {
	todoList, _ := repotodo.ShowALLTODO(DB)
	errs := json.NewEncoder(w).Encode(todoList)
	if errs != nil {
		w.WriteHeader(500)
		msg := fmt.Sprintf("Дороу залупа, это не работает %s", errs)
		w.Write([]byte(msg))
	}

	return todoList, nil
}

func deleteTODO(DB *sql.DB)
