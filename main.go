package main

import (
	"net/http"
	"os"

	c "./controllers"

	"github.com/gorilla/mux"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	r := mux.NewRouter()

	r.HandleFunc("/api/books", c.Index).Methods("GET")
	r.HandleFunc("/api/books/top/{col}", c.GetTop).Methods("GET")
	r.HandleFunc("/api/books/low/{col}", c.GetLow).Methods("GET")
	r.HandleFunc("/api/books/search/{search}", c.Search).Methods("GET")

	r.PathPrefix("/css/").Handler(
		http.StripPrefix("/css/", http.FileServer(http.Dir("views/css/"))),
	)

	r.PathPrefix("/javascript/").Handler(
		http.StripPrefix("/javascript/", http.FileServer(http.Dir("views/javascript/"))),
	)

	var fs http.Handler = http.FileServer(http.Dir("./views"))

	r.Handle("/", fs)

	http.ListenAndServe(":"+port, r)
}
