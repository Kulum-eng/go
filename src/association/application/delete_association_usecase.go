package application

import "association/domain"

type DeleteAssociationUseCase struct {
	repo domain.AssociationRepository
}

func NewDeleteAssociationUseCase(repo domain.AssociationRepository) *DeleteAssociationUseCase {
	return &DeleteAssociationUseCase{repo: repo}
}

func (useCase *DeleteAssociationUseCase) Execute(id int) error {
	return useCase.repo.Delete(id)
}
