package app

import (
	"users-api/controllers"
)

func MapRoutes() {
	router.GET("/health-check", controllers.HealthCheck)
	router.GET("/ping", controllers.CreateUser)
}
