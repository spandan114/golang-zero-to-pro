package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spandan114/golang-zero-to-pro/29-JWT-Auth/helpers"
)

func Authenticate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		clientToken := ctx.Request.Header.Get("Token")
		if clientToken == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"err": "Token not found"})
			ctx.Abort()
			return
		}

		claims, err := helpers.ValidateToken(clientToken)
		if err != "" {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
			ctx.Abort()
			return
		}
		ctx.Set("email", claims.Email)
		ctx.Set("first_name", claims.FirstName)
		ctx.Set("last_name", claims.LastName)
		ctx.Set("uid", claims.Uid)
		ctx.Set("user_type", claims.UserType)
		ctx.Next()
	}
}
