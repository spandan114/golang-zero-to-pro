package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/spandan114/golang-zero-to-pro/29-JWT-Auth/controllers"
	"github.com/spandan114/golang-zero-to-pro/29-JWT-Auth/middleware"
)

func UserRouter(r *gin.Engine) {
	r.Use(middleware.Authenticate())
	r.GET("/api/user", controllers.GetUsers())
	r.GET("/api/user/:user_id", controllers.GetUserById())
}
