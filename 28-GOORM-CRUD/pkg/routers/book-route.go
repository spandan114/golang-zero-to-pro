package routers

import (
	"github.com/gorilla/mux"
	"github.com/spandan114/golang-zero-to-pro/GOORM-CRUD/pkg/controllers"
)

var RegisterBookstoreRoutes = func(r *mux.Router) {

	r.HandleFunc("/api/author", controllers.CreateAuthor).Methods("POST")
	r.HandleFunc("/api/author", controllers.GetAuthors).Methods("GET")

	r.HandleFunc("/api/book", controllers.CreateBook).Methods("POST")
	r.HandleFunc("/api/book", controllers.GetBooks).Methods("GET")
	r.HandleFunc("/api/book/{id}", controllers.GetBookById).Methods("GET")
	r.HandleFunc("/api/book/{id}", controllers.UpdateBook).Methods("PUT")
	r.HandleFunc("/api/book/{id}", controllers.DeleteBook).Methods("DELETE")
}
