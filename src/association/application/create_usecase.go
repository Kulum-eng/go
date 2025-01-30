package application

import (
	"api/src/association/domain"
	"api/src/association/domain/ports"
)

type CreateUseCase struct {
	repo ports.Repository
}

func NewCreateUseCase(repo ports.Repository) *CreateUseCase {
	return &CreateUseCase{repo: repo}
}

func (uc *CreateUseCase) Execute(association domain.Association) error {
	return uc.repo.Create(association)
}
