package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spandan114/golang-zero-to-pro/29-JWT-Auth/routes"
)

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	router := gin.New()
	router.Use(gin.Logger())

	routes.AuthRouter(router)
	routes.UserRouter(router)

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.Run(":" + port)
}
