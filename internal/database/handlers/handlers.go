package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	repotodo "todo/RepoTODO"

	"github.com/gorilla/mux"
)

func HiHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Дороу залупа, это работает"))
}

func GetAllTODOHandler(DB *sql.DB, w http.ResponseWriter, r *http.Request) ([]repotodo.TODO, error) {
	todoList, err := repotodo.GetALLTODO(DB)
	if err != nil {
		http.Error(w, "Ошибка БД", 500)
		return nil, err
	}
	err = json.NewEncoder(w).Encode(todoList)
	if err != nil {
		w.WriteHeader(500)
		msg := fmt.Sprintf("Дороу залупа, это не работает %s", err)
		w.Write([]byte(msg))
		return nil, err
	}

	return todoList, nil
}

func GetTODOByIDHandler(DB *sql.DB, w http.ResponseWriter, r *http.Request) (*repotodo.TODO, error) {
	vars := mux.Vars(r)
	idstr := vars["id"]

	if idstr == "" {
		http.Error(w, "НЕТУ АЙДИ", 400)
		return nil, fmt.Errorf("no id")
	}

	id, err := strconv.Atoi(idstr)
	if err != nil {
		http.Error(w, "ХУЙНЯ", 400)
		return nil, err
	}

	singleTodo, err := repotodo.GetTODOByID(DB, id)
	if err != nil {
		w.WriteHeader(500)
		msg := fmt.Sprintf("Такого айди нет: %s", err)
		w.Write([]byte(msg))
	}

	err = json.NewEncoder(w).Encode(singleTodo)

	return singleTodo, nil
}

func DeleteTODOByIDHandler(DB *sql.DB, w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	idstr := vars["id"]

	if idstr == "" {
		http.Error(w, "НЕТУ АЙДИ", 400)
		return fmt.Errorf("no id")
	}

	id, err := strconv.Atoi(idstr)
	if err != nil {
		http.Error(w, "Кривой АЙДИ", 400)
		return err
	}

	err = repotodo.DeleteTODO(DB, id)
	if err != nil {
		http.Error(w, err.Error(), 404)
		return err
	}
	return nil
}

func SaveTODOHandler(DB *sql.DB, w http.ResponseWriter, r *http.Request) (*repotodo.TODO, error) {
	var newTODO repotodo.TODO

	err := json.NewDecoder(r.Body).Decode(&newTODO)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	_, err = repotodo.InsertTODO(DB, &newTODO)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &newTODO, nil
}

func UpdateTODOByIDHandler(DB *sql.DB, w http.ResponseWriter, r *http.Request) (*repotodo.TODO, error) {
	var updatedTodo repotodo.TODO

	err := json.NewDecoder(r.Body).Decode(&updatedTodo)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	_, err = repotodo.UpdateTODOByID(DB, &updatedTodo)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &updatedTodo, nil
}
