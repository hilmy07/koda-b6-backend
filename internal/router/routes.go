package routes

import (
	container "backend/internal/di"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func corsMiddleware() gin.HandlerFunc {
	// godotenv.Load()
	return func(ctx *gin.Context) {
		ctx.Header("Access-Control-Allow-Origin", os.Getenv("FRONTEND_URL"))
		ctx.Header("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
		ctx.Header("Access-Control-Allow-Headers", "content-type,authorization")
		
		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(http.StatusOK)
		} else {
			ctx.Next()
		}
	}
}

func SetupRoutes(r *gin.Engine, db *pgx.Conn) {

	c := container.NewContainer(db)

	authHandler := c.AuthHandler()
	userHandler := c.UserHandler()
	productHandler := c.ProductHandler()
	forgotHandler := c.ForgotPasswordHandler()

	r.Use(corsMiddleware())

	r.POST("/auth", authHandler.AuthLogin)
	r.POST("/auth/new", authHandler.AuthRegister)
	r.PATCH("/users/profile", authHandler.AuthProfile)
	r.GET("/users", userHandler.GetUser)
	r.DELETE("/users/:id", userHandler.DeleteUser)
	
	r.POST("/auth/forgot-password", forgotHandler.RequestForgotPassword)
	r.PATCH("/reset-password", forgotHandler.ResetPassword)
	

	// r.PATCH("/users/uploads", authHandler.AuthProfile)

	// halaman home
	r.GET("/recommended-products", productHandler.GetRecommendedProduct)
	r.GET("/reviews", func (ctx *gin.Context) {
		ctx.Header("Access-Control-Allow-Origin", "http://localhost:5173")
		ctx.Header("Access-Control-Allow-Headers", "content-type")
		ctx.Data(http.StatusOK, "", []byte(""))
	}, productHandler.GetProductReview)

	// halaman product
	r.POST("/product/create", productHandler.CreateProduct)
	r.DELETE("/product/:id", productHandler.DeleteProduct)
	r.GET("/product", productHandler.GetProductList)
	r.GET("/product/:id", productHandler.GetProductDetail)
}


