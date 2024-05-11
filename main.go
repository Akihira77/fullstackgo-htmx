package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/Akihira77/fullstack-go-htmx/handlers"
	"github.com/Akihira77/fullstack-go-htmx/store"
	"github.com/go-sql-driver/mysql"
)

func main() {
	cfg := mysql.Config{
		User:                 store.Envs.DBUser,
		Passwd:               store.Envs.DBPassword,
		Addr:                 store.Envs.DBAddress,
		DBName:               store.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	db, err := store.NewMySQLStorage(cfg)
	if err != nil {
		log.Fatal(err)
	}

	store := store.NewStore(db)

	initStorage(db)

	router := http.NewServeMux()

	handler := handlers.New(store)

	router.HandleFunc("GET /", handler.HandleHome)
	router.HandleFunc("GET /cars", handler.HandleCarsList)
	router.HandleFunc("GET /cars/search", handler.HandleCarSearch)
	router.HandleFunc("DELETE /cars/{id}", handler.HandleDeleteCar)

	// serve files in public
	router.Handle("GET /public", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	fmt.Printf("Listening on %v\n", "localhost:8080")
	http.ListenAndServe(":8080", router)
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to Database!")
}
