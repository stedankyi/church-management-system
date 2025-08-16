package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stedankyi/church-management-system/config"
	"github.com/stedankyi/church-management-system/routes"
)

func main() {
	fmt.Println("Hello Welcome to the Church Management System!")

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not found in the environment file!")
	}

	fmt.Println("Port:", portString)

	// Connect to database
	config.ConnectDatabase()

	// Create routers using gin
	router := gin.Default()
	routes.SetupRoutes(router)

	// Start webserver using gin
	router.Run(":" + portString)

}
