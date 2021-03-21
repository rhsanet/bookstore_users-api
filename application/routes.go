package app

import (
	"users-api/controllers"
)

func MapRoutes() {
	router.GET("/health-check", controllers.HealthCheck)
	router.POST("/users", controllers.CreateUser)
	router.GET("/users", controllers.GetAll)
	router.GET("/users/:user_id", controllers.FindUser)
}
