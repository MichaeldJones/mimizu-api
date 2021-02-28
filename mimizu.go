package main

import (
	"net/http"

	c "./controllers"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/api/books", c.Index).Methods("GET")

	http.ListenAndServe(":80", r)
}
