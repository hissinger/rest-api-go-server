package main

import (
	"net/http"
	"rest-api-server/api"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/users/", api.GetUsers).Methods("GET")
	r.HandleFunc("/users/{id}", api.GetUser).Methods("GET")
	r.HandleFunc("/users/", api.AddUser).Methods("POST")
	r.HandleFunc("/users/{id}", api.DelUser).Methods("DELETE")

	http.ListenAndServe(":3000", r)
}
