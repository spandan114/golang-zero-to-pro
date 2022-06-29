package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id"`
	FirstName    string             `json:"first_name" validation:"required"`
	LastName     string             `json:"last_name" validation:"required"`
	Password     string             `json:"password" validation:"required, min=6"`
	PhoneNumber  string             `json:"phone" validation:"phone, required"`
	Email        string             `json:"email" validation:"email, required"`
	UserID       string             `json:"user_id"`
	Token        string             `json:"token"`
	RefreshToken string             `json:"refresh_token"`
	UserType     string             `json:"type" validation:"required, eq=ADMIN|eq=USER"`
	CtratedAt    time.Time          `json:"created_at"`
	UpdatedAt    time.Time          `json:"updated_at"`
}
