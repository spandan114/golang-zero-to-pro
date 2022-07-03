package utils

import (
	"github.com/spandan114/golang-zero-to-pro/GOORM-CRUD/pkg/models"
)

func IsValid(b models.Book) string {

	// check if the name empty
	if b.Name == "" {
		return "The name is required!"
	}
	// check the name field is between 3 to 120 chars
	if len(b.Name) < 2 || len(b.Name) > 40 {
		return "The name field must be between 2-40 chars!"
	}
	// if b.Author == "" {
	// 	return "The Author field is required!"
	// }

	if b.Publication == "" {
		return "The Publication field is required!"

	}
	return ""
}
