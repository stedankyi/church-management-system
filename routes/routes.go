package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/stedankyi/church-management-system/controllers"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/users/:id", controllers.GetUser)
	router.GET("/users", controllers.GetUsers)
	router.POST("/users", controllers.CreateUser)
	router.PUT("/users/:id", controllers.UpdateUser)
	router.DELETE("/users/:id", controllers.DeleteUser)
}
