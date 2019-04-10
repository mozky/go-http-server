package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hi there!\n"))
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hi " + username))
	log.Printf("Request for user %v\n", username)
}

func main() {
	r := mux.NewRouter()
	log.Print("Starting new server on :8000 😀")

	r.HandleFunc("/", Handler)
	r.HandleFunc("/user/{username}", UserHandler)

	log.Fatal(http.ListenAndServe(":8000", r))
}
