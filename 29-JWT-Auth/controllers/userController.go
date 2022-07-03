package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/spandan114/golang-zero-to-pro/29-JWT-Auth/database"
	"github.com/spandan114/golang-zero-to-pro/29-JWT-Auth/helpers"
	"github.com/spandan114/golang-zero-to-pro/29-JWT-Auth/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "User")

var validate = validator.New()

func Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user models.User
		var foundUser models.User
		if err := ctx.BindJSON(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err := userCollection.FindOne(context.Background(), bson.M{"email": user.Email}).Decode(&foundUser)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Email or Password is incorrect !"})
			return
		}

		passwordIsValid, msg := helpers.VerifyPassword(user.Password, foundUser.Password)

		if passwordIsValid != true {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": msg})
		}

		token, refreshToken, _ := helpers.GenerateAllTokens(foundUser.Email, foundUser.FirstName, foundUser.LastName, foundUser.UserID, foundUser.UserType)
		helpers.UpdateAllTokens(token, refreshToken, foundUser.UserID)

		errMsg := userCollection.FindOne(context.Background(), bson.M{"user_id": foundUser.UserID}).Decode(&foundUser)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": errMsg})
		}
		ctx.JSON(http.StatusOK, foundUser)
	}
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
		emailCount, err := userCollection.CountDocuments(context.Background(), bson.M{"email": user.Email})

		if err != nil {
			log.Panic(err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error occur while counting user email"})
			return
		}

		numCount, err := userCollection.CountDocuments(context.Background(), bson.M{"phone": user.PhoneNumber})
		if err != nil {
			log.Panic(err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error occur while counting user phone number"})
			return
		}

		if emailCount > 0 || numCount > 0 {
			ctx.JSON(http.StatusExpectationFailed, gin.H{"error": "This email or phone number alredy exist"})
			return
		}

		user.CtratedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.ID = primitive.NewObjectID()
		user.UserID = user.ID.Hex()
		token, refreshToken, _ := helpers.GenerateAllTokens(user.Email, user.FirstName, user.LastName, user.UserID, user.UserType)
		user.Token = token
		user.RefreshToken = refreshToken
		user.Password = helpers.HashPassword(user.Password)

		res, err := userCollection.InsertOne(context.Background(), user)

		if err != nil {
			msg := fmt.Sprint("User not created yet")
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		ctx.JSON(http.StatusOK, res)
	}
}

func GetUsers() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if err := helpers.CheckUserType(ctx, "ADMIN"); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		cursor, err := userCollection.Find(context.Background(), bson.D{{}})

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		var allusers []bson.M
		if err = cursor.All(ctx, &allusers); err != nil {
			log.Fatal(err)
		}

		ctx.JSON(http.StatusOK, allusers)
	}
}

func GetUserById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId := ctx.Param("user_id")
		if err := helpers.MatchUserTypeToUid(ctx, userId); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		fmt.Printf("%v \n", userId)

		var user models.User
		err := userCollection.FindOne(context.Background(), bson.M{"userid": userId}).Decode(&user)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, user)

	}
}
