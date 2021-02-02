package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	m "../models"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const connStr = "user= dbname= password= host=localhost sslmode=disable"

func Index(w http.ResponseWriter, r *http.Request) {

	books := []m.BookSub{}

	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*") // erm this might just be used for local

	db.Select(&books, "select title, author, link, total, uniq from books limit 100")

	json.NewEncoder(w).Encode(books)

}

func GetTop(w http.ResponseWriter, r *http.Request) {

	books := []m.BookSub{}

	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		fmt.Println(err)
	}

	params := mux.Vars(r)

	query := "select title, author, link, total, uniq from books order by " + params["col"] + " desc limit 100"

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*") // erm this might just be used for local

	db.Select(&books, query)

	json.NewEncoder(w).Encode(books)

}

func GetLow(w http.ResponseWriter, r *http.Request) {

	books := []m.BookSub{}

	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		fmt.Println(err)
	}

	params := mux.Vars(r)

	query := "select title, author, link, total, uniq from books order by " + params["col"] + " asc limit 100"

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*") // erm this might just be used for local

	db.Select(&books, query)

	json.NewEncoder(w).Encode(books)

}

func Search(w http.ResponseWriter, r *http.Request) {

	books := []m.BookSub{}

	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		fmt.Println(err)
	}

	params := mux.Vars(r)
	search := params["search"]

	query := "select title, author, link, total, uniq from books where title like '%" + search + "%' or author like '%" + search + "%' order by total desc limit 100"

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*") // erm this might just be used for local

	db.Select(&books, query)

	json.NewEncoder(w).Encode(books)
}
