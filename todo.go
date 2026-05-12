package main

import (
	"database/sql"
	"log"
	"net/http"
	"todo/internal/database/handlers"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	connStr := "postgres://postgres:1@localhost:5432/EQ?sslmode=disable"
	DB, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer DB.Close()

	router := mux.NewRouter()

	router.HandleFunc("/", handlers.HiHandler).Methods("GET")
	router.HandleFunc("/todo", func(w http.ResponseWriter, r *http.Request) { handlers.GetAllTODOHandler(DB, w, r) }).Methods("GET")             // Получить все туду
	router.HandleFunc("/todo/{id}", func(w http.ResponseWriter, r *http.Request) { handlers.GetTODOByIDHandler(DB, w, r) }).Methods("GET")       // Получить туду по айди
	router.HandleFunc("/todo/{id}", func(w http.ResponseWriter, r *http.Request) { handlers.DeleteTODOByIDHandler(DB, w, r) }).Methods("DELETE") // Удалить туду по айди
	router.HandleFunc("/todo", func(w http.ResponseWriter, r *http.Request) { handlers.SaveTODOHandler(DB, w, r) }).Methods("POST")              // Добавить новое туду
	router.HandleFunc("/todo/{id}", func(w http.ResponseWriter, r *http.Request) { handlers.UpdateTODOByIDHandler(DB, w, r) }).Methods("PATCH")  // Обновить туду по айди

	router.PathPrefix("/").Handler(http.FileServer(http.Dir(".")))

	log.Fatal(http.ListenAndServe(":8080", router))

}
