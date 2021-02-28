package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	m "../models"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const connStr = "user=postgres dbname=postgres password=postgres host=localhost sslmode=disable"

func Index(w http.ResponseWriter, r *http.Request) {

	books := []m.BookSub{}

	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")

	db.Select(&books, "select title, author, link, total, uniq from books")

	json.NewEncoder(w).Encode(books)

	db.Close()
}
