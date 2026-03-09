package routes

import (
	container "backend/internal/di"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func SetupRoutes(r *gin.Engine, db *pgx.Conn) {

	c := container.NewContainer(db)

	authHandler := c.AuthHandler()

	r.POST("/auth", authHandler.AuthLogin)
	r.POST("/auth/new", authHandler.AuthRegister)
}