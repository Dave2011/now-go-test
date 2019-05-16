package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", Handler).Methods("GET")
	log.Fatal(http.ListenAndServe(":3030", router))
}
func Handler(w http.ResponseWriter, r *http.Request) { fmt.Fprintf(w, "Wow, SUCH Test") }
