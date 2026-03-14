package routes

import (
	container "backend/internal/di"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func SetupRoutes(r *gin.Engine, db *pgx.Conn) {

	c := container.NewContainer(db)

	authHandler := c.AuthHandler()
	productHandler := c.ProductHandler()
	forgotHandler := c.ForgotPasswordHandler()

	r.POST("/auth", authHandler.AuthLogin)
	r.POST("/auth/new", authHandler.AuthRegister)
	
	r.POST("/auth/forgot-password", forgotHandler.RequestForgotPassword)
	r.PATCH("/reset-password", forgotHandler.ResetPassword)
	
	r.PATCH("/users/profile", authHandler.AuthProfile)

	// r.PATCH("/users/uploads", authHandler.AuthProfile)

	// halaman home
	r.GET("/recommended-products", productHandler.GetRecommendedProduct)
	r.GET("/reviews", productHandler.GetProductReview)

	// halaman product
	r.GET("/product", productHandler.GetProductList)
	r.GET("/product/:id", productHandler.GetProductDetail)
}