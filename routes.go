package main

import (
	"net/http"
)

func SetupRoutes() {
	http.HandleFunc("/login", login)
	http.HandleFunc("/users", getUsers)
	http.HandleFunc("/users/:id", getUser)
	http.HandleFunc("/users", createUser)
	http.HandleFunc("/users/:id", updateUser)
	http.HandleFunc("/users/:id", deleteUser)
}
