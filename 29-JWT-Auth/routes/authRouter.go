package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/spandan114/golang-zero-to-pro/29-JWT-Auth/controllers"
)

func AuthRouter(r *gin.Engine) {
	r.POST("/api/login", controllers.Login())
	r.POST("/api/signup", controllers.Signup())
}
