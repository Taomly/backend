package routers

import (
	"auth/handlers"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func SetupRouter(r *gin.Engine, db *pgxpool.Pool) {
	r.POST("signup/", handlers.Signup(db))
	r.POST("login/", handlers.Login())
}
