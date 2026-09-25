package main

import (
	"fmt"
	"log"
	"os"

	"booking-app/controllers"
	"booking-app/database"
	"booking-app/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Local development can use .env; containers receive values via --env-file.
	_ = godotenv.Load()
	database.InitDB()

	controllers.SyncTickets()

	r := gin.Default()

	routes.RegisterRoutes(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Printf("Server is running on http://localhost:%s...\n", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
