package main

import (
	con "backend/internal/di"
	routes "backend/internal/router"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	db, err := con.Connect()

	if err != nil {
		panic(err)
	}

	r := gin.Default()

	routes.SetupRoutes(r, db)

	port := os.Getenv("PORT")

	r.Run(":" + port)
}


