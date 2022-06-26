package models

import (
	"fmt"

	"github.com/spandan114/golang-zero-to-pro/GOORM-CRUD/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	AuthorId    uint
	Author      Author `json:"author" gorm:"foreignkey:AuthorId"`
	Publication string `json:"publication"`
}

type Author struct {
	gorm.Model
	Name    string `json:"name"`
	Website string `json:"website"`
}

func init() {
	config.ConnectDB()
	db = config.GetDb()
	db.AutoMigrate(&Author{})
	db.AutoMigrate(&Book{})
}

func CreateAuthor(a *Author) *Author {
	result := db.Create(&a)
	if result.Error != nil {
		fmt.Printf("Error : %v \n", result.Error)
	}
	fmt.Printf("Row count : %v\n", a.ID)
	return a
}

func GetAuthors() []Author {
	authors := []Author{}
	db.Find(&authors)
	return authors
}

func CreateBook(b *Book) *Book {
	result := db.Create(&b)
	if result.Error != nil {
		fmt.Printf("Error : %v \n", result.Error)
	}
	fmt.Printf("Row count : %v\n", b.ID)
	return b
}

func GetBook() []Book {
	books := []Book{}
	db.Preload("Author").Find(&books)

	return books
}

func GetUserById(id int64) *Book {
	var book *Book
	db.Preload("Author").First(&book, id)
	return book
}

func DeleteUser(id int64) *Book {
	var book *Book
	db.Where("id=?", id).Delete(&book)
	return book
}

func UpdateBook(book *Book) *Book {
	db.Save(&book)
	return book
}
