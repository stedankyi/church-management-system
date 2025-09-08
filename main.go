package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"github.com/stedankyi/church-management-system/config"
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
	SetupRoutes()

	// Start webserver
	http.ListenAndServe(":"+portString, nil)
}
