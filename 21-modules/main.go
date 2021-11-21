package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hii")
	r := mux.NewRouter()
	r.HandleFunc("/", runServer).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func runServer(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello go</h1>"))
}
