package app

import (
	"users-api/controllers"
)

func MapRoutes() {
	router.GET("/health-check", controllers.HealthCheck)
	router.POST("/", controllers.CreateUser)
	router.GET("/", controllers.GetAll)
	// router.GET("/:user_id", controllers.FindUser)
}
