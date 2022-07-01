package helpers

import (
	"log"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	// "github.com/spandan114/golang-zero-to-pro/29-JWT-Auth/database"
	// "go.mongodb.org/mongo-driver/mongo"
)

type SignedDetails struct {
	Email     string
	FirstName string
	LastName  string
	Uid       string
	UserTpe   string
	jwt.StandardClaims
}

// var userCollection *mongo.Collection = database.OpenCollection(database.Client, "User")
var SECRET_KEY string = os.Getenv("JWT_SEC")

func GenerateAllTokens(email string, firstName string, lastName string, uid string, userType string) (signedToken string, refreshToken string, err error) {
	claims := &SignedDetails{
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
		Uid:       uid,
		UserTpe:   userType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	refreshClaims := &SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(),
		},
	}

	// token, err := jwt.NewWithClaims(jwt.SigningMethodES256, claims).SignedString([]byte("secretKey"))
	// refreshToken, err2 := jwt.NewWithClaims(jwt.SigningMethodES256, refreshClaim).SignedString([]byte("secretKey"))

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
	refreshToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(SECRET_KEY))

	if err != nil {
		log.Panic(err)
		return
	}
	return token, refreshToken, err
}
