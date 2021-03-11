package use_cases

import (
	"users-api/domain"
	"users-api/presentation/response"
)

func CreateUserUseCase(user domain.User) (*domain.User, *response.ErrorMessage) {
	return &user, nil
}
