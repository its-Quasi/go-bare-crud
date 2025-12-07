package main

import (
	"bare-crud/internal/service"
	"bare-crud/internal/store"
	"bare-crud/internal/transport"
	"database/sql"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./books.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//Creation of necesary tables
	q := `
		CREATE TABLE IF NOT EXISTS books (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT NOT NULL,
			description TEXT NOT NULL
		)
	`

	if _, err := db.Exec(q); err != nil {
		log.Fatal(err)
	}

	//necesary dependencies

	bookStore := store.New(db)
	bookService := service.New(bookStore)
	bookHandler := transport.New(bookService)

	// config routes
	http.HandleFunc("/books", bookHandler.HandleBooks)

	// launch server
	log.Fatal(http.ListenAndServe(":8080", nil))

}
