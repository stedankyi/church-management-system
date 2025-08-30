package routes

import (
	"github.com/stedankyi/church-management-system/controllers"
)

func SetupRoutes() {
	GET("/users/:id", controllers.GetUser)
	GET("/users", controllers.GetUsers)
	POST("/users", controllers.CreateUser)
	PUT("/users/:id", controllers.UpdateUser)
	DELETE("/users/:id", controllers.DeleteUser)
}
