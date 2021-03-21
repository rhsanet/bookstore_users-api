package users

import (
	"strings"
	"users-api/presentation/response"
)

type User struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

func (user *User) Validate() *response.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return response.BadRequestResponse("Invalid email address")
	}
	return nil
}
