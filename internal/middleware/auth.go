package middleware

import (
	"net/http"
	"strings"

	"backend/internal/lib"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Authorization header required",
			})
			c.Abort()
			return
		}

		splitToken := strings.Split(authHeader, " ")
		if len(splitToken) != 2 || splitToken[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid token format",
			})
			c.Abort()
			return
		}

		// 🔥 pakai VerifyToken dari lib
		claims, err := lib.VerifyToken(splitToken[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			c.Abort()
			return
		}

		// 🔥 langsung ambil Id (bukan user_id lagi)
		c.Set("user_id", claims.Id)
		println("TOKEN:", splitToken[1])
		println("USER ID FROM TOKEN:", claims.Id)

		c.Next()
	}
}