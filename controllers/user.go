package controllers

import (
	"encoding/json"
	"io/ioutil"
	"users-api/domain"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user domain.User
	data, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		// todo: handle error
		return
	}

	if err := json.Unmarshal(data, &user); err != nil {
		// todo: handle error
		return
	}

	c.JSON(200, gin.H{
		"message": "ok",
		"user":    user,
	})
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
