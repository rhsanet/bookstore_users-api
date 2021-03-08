package app

import (
	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func BootstrapApp() {
	MapRoutes()
	router.Run()
}
