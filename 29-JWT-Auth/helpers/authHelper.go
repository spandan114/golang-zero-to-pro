package helpers

import (
	"errors"

	"github.com/gin-gonic/gin"
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
