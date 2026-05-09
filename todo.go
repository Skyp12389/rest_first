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
	router.HandleFunc("/alltodo", func(w http.ResponseWriter, r *http.Request) { handlers.ShowAllTODO(DB, w, r) })

	log.Fatal(http.ListenAndServe(":8080", router))

	// for {
	// 	fmt.Println("=== TO DO MENU ===")
	// 	fmt.Println("1. Добавить задачу\n2. Удалить задачу\n3. Обновить задачу\n4. Отметить как выполненую\n5. Найти задачу по ID")

	// 	var choice int
	// 	fmt.Scan(&choice)
	// 	switch choice {
	// 	case 1:
	// 		{
	// 			fmt.Print("Введите задачу:")
	// 			var b bool
	// 			var a string
	// 			fmt.Scan(&a)
	// 			now := time.Now()
	// 			repotodo.InsertTODO(DB, &repotodo.TODO{Task: a, Done: b, CreatedAT: &now})
	// 		}
	// 	case 2:
	// 		{
	// 			var sad int
	// 			fmt.Scan(&sad)
	// 			repotodo.DeleteTODO(DB, sad)
	// 		}
	// 	case 3:
	// 		{
	// 			fmt.Println("Введите номер задачи для ее обновления: ")
	// 			fmt.Scan()
	// 		}
	// 		// case 4:
	// 		// 	{
	// 		// 		ReadyTODO()
	// 		// 	}
	// 		// case 5:
	// 		// 	{
	// 		// 		FindByIDTODO()
	// 		// 	}
	// 		// }
	// 	}
	// }
}
