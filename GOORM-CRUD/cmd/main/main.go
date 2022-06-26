package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/spandan114/golang-zero-to-pro/GOORM-CRUD/pkg/routers"
)

func main() {
	r := mux.NewRouter()
	routers.RegisterBookstoreRoutes(r)
	http.Handle("/", r)
	fmt.Println("Server is running on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
