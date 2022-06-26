package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/spandan114/golang-zero-to-pro/GOORM-CRUD/pkg/models"
	"github.com/spandan114/golang-zero-to-pro/GOORM-CRUD/pkg/utils"
	"gorm.io/gorm"
)

var DB *gorm.DB

func CreateAuthor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	authorData := &models.Author{}
	json.NewDecoder(r.Body).Decode(authorData)
	var newAuthor = models.CreateAuthor(authorData)
	json.NewEncoder(w).Encode(newAuthor)
}

func GetAuthors(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	authors := models.GetAuthors()
	json.NewEncoder(w).Encode(authors)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	bookData := &models.Book{}
	json.NewDecoder(r.Body).Decode(bookData)
	err := utils.IsValid(*bookData)

	if err != "" {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":     "success",
			"statusCode": 400,
			"data":       err,
		})
		return
	}
	fmt.Printf("Data : %v\n", bookData)
	var newBook = models.CreateBook(bookData)
	json.NewEncoder(w).Encode(newBook)
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	books := models.GetBook()
	json.NewEncoder(w).Encode(books)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 0, 0)
	if err != nil {
		json.NewEncoder(w).Encode("Error while parsing ID")
		return
	}
	book := models.GetUserById(id)
	json.NewEncoder(w).Encode(&book)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, error := strconv.ParseInt(params["id"], 0, 0)
	if error != nil {
		json.NewEncoder(w).Encode("Error while parsing ID")
		return
	}

	bookData := &models.Book{}
	json.NewDecoder(r.Body).Decode(bookData)
	err := utils.IsValid(*bookData)

	if err != "" {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":     "success",
			"statusCode": 400,
			"data":       err,
		})
		return
	}

	var book = models.GetUserById(id)
	book.Name = bookData.Name
	book.Author = bookData.Author
	book.Publication = bookData.Publication

	updatedBook := models.UpdateBook(book)
	json.NewEncoder(w).Encode(updatedBook)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 0, 0)
	if err != nil {
		json.NewEncoder(w).Encode("Error while parsing ID")
		return
	}
	book := models.DeleteUser(id)
	json.NewEncoder(w).Encode(&book)
}
