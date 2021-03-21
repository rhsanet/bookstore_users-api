package use_cases

import (
	"users-api/domain/users"
	"users-api/presentation/response"
)

func CreateUserUseCase(user users.User) (*users.User, *response.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	return &user, nil
}
