package main

import (
	"log"
	"os"
	"update_product/database"
	"update_product/handlers"

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

	// Ruta para actualizar producto
	router.PUT("/products/:id", handlers.UpdateProduct)

	port := os.Getenv("PORT")
	if port == "" {
		port = "4002"
	}

	log.Println("🚀 Server running on port", port)
	router.Run(":" + port)
}