package controllers

import (
	"net/http"
	"users-api/domain"
	"users-api/presentation/response"
	"users-api/use_cases"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user domain.User

	if err := c.ShouldBindJSON(&user); err != nil {
		respose := response.BadRequestResponse("invalid params", err)
		c.JSON(http.StatusBadRequest, respose)
		return
	}

	result, err := use_cases.CreateUserUseCase(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func GetAll(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
		"status":  true,
	})
}

func FindUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
		"status":  true,
	})
}
