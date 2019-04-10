package main

import (
	"flag"
	"fmt"
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

	ipPtr := flag.String("ip", "127.0.0.1", "Ip that webserver binds to")
	portPtr := flag.Int("port", 8000, "Port that webserver listens to")

	flag.Parse()

	log.Printf("Starting new server on %s:%d 😀", *ipPtr, *portPtr)

	r.HandleFunc("/", Handler)
	r.HandleFunc("/user/{username}", UserHandler)

	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", *ipPtr, *portPtr), r))
}
