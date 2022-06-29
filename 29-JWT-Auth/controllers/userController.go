package controllers

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/spandan114/golang-zero-to-pro/29-JWT-Auth/database"
	"github.com/spandan114/golang-zero-to-pro/29-JWT-Auth/helpers"
	"github.com/spandan114/golang-zero-to-pro/29-JWT-Auth/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "User")

var validate = validator.New()

func Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}

func Signup() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user models.User
		if err := ctx.BindJSON(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := validate.Struct(&user)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		count, err := userCollection.CountDocuments(context.Background(), bson.M{"email": user.Email})

		if err != nil {
			log.Panic(err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error occur while counting user email"})
		}

		count, err := userCollection.CountDocuments(context.Background(), bson.M{"phone": user.PhoneNumber})
		if err != nil {
			log.Panic(err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error occur while counting user phone number"})
		}

		if count > 0 {
			ctx.JSON(http.StatusExpectationFailed, gin.H{"error": "This email or phone number alredy exist"})
		}

	}
}

func GetUsers() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func GetUserById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId := ctx.Param("user_id")
		if err := helpers.MatchUserTypeToUid(ctx, userId); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var user models.User
		err := userCollection.FindOne(context.Background(), bson.M{"user_id": userId}).Decode(&user)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, user)

	}
}
