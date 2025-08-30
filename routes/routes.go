package routes

import (
	"net/http"

	"github.com/stedankyi/church-management-system/controllers"
)

func SetupRoutes() {
	http.HandleFunc("/login", controllers.Login)
	http.HandleFunc("/users", controllers.GetUsers)
	http.HandleFunc("/users/:id", controllers.GetUser)
	http.HandleFunc("/users", controllers.CreateUser)
	http.HandleFunc("/users/:id", controllers.UpdateUser)
	http.HandleFunc("/users/:id", controllers.DeleteUser)
}
