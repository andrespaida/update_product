package main

import (
	"log"
	"os"
	"update_product/database"
	"update_product/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.ConnectDB()

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "GET", "POST", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
	}))

	router.PUT("/products/:id", handlers.UpdateProduct)

	port := os.Getenv("PORT")
	if port == "" {
		port = "4003"
	}

	log.Println("🚀 Server running on port", port)
	router.Run(":" + port)
}