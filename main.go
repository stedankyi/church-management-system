package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/stedankyi/church-management-system/config"
	"github.com/stedankyi/church-management-system/routes"
)

type Login struct {
	HashPassword string
	SessionToken string
	CSRFToken    string
}

var users = map[string]Login{}

func main() {
	fmt.Println("Hello, Welcome to the Church Management System!")

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

	// Create routers
	routes.SetupRoutes()

	// Start webserver
	http.ListenAndServe(":"+portString, nil)
}
