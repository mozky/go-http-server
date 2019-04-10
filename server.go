package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hi there!\n"))
}

func main() {
	r := mux.NewRouter()
	log.Print("Starting new server on :8000")

	r.HandleFunc("/", Handler)

	log.Fatal(http.ListenAndServe(":8000", r))
}
