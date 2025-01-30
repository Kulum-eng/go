package application

import (
	"user/domain"
	"user/domain/ports"

)

type CreateUserUseCase struct {
	repo ports.UserRepository
}

func NewCreateUserUseCase(repo ports.UserRepository) *CreateUserUseCase {
	return &CreateUserUseCase{repo: repo}
}

func (uc *CreateUserUseCase) Execute(user domain.User) error {
	return uc.repo.CreateUser(user)
}
