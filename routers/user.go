package routers

import (
	"auth/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	r.POST("signup/", handlers.Signup())
	r.POST("login/", handlers.Login())
}
