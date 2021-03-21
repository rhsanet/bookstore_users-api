package controllers

import (
	"net/http"
	"users-api/domain/users"
	"users-api/presentation/response"
	"users-api/use_cases"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		errMsg := response.BadRequestResponse("invalid params")
		c.JSON(http.StatusBadRequest, errMsg)
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
	userId := c.Param("user_id")
	if userId == "" {
		err := response.BadRequestResponse("Invalid user id")
		c.JSON(err.StatusCode, err)
		return
	}

	user, err := use_cases.FindUserById(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(200, user)
}
