package helpers

import (
	"errors"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func MatchUserTypeToUid(ctx *gin.Context, userId string) (err error) {
	userType := ctx.GetString("user_type")
	uId := ctx.GetString("uid")
	err = nil

	if userType == "USER" || uId != userId {
		err = errors.New("unauthorize to access this request")
		return err
	}
	return err
}

func VerifyPassword(userPassword string, providedPassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword))
	check := true
	msg := ""
	if err != nil {
		msg = "Password doesn't match !"
		check = false
	}
	return check, msg
}

func HashPassword(password string) string {
	return "Hii"
}
