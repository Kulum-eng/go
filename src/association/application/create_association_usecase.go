package application

import "association/domain"

type CreateAssociationUseCase struct {
	repo domain.AssociationRepository
}

func NewCreateAssociationUseCase(repo domain.AssociationRepository) *CreateAssociationUseCase {
	return &CreateAssociationUseCase{repo: repo}
}

func (useCase *CreateAssociationUseCase) Execute(association domain.Association) error {
	return useCase.repo.Create(association)
}
