package application

import "association/domain"

type UpdateAssociationUseCase struct {
	repo domain.AssociationRepository
}

func NewUpdateAssociationUseCase(repo domain.AssociationRepository) *UpdateAssociationUseCase {
	return &UpdateAssociationUseCase{repo: repo}
}

func (useCase *UpdateAssociationUseCase) Execute(association domain.Association) error {
	return useCase.repo.Update(association)
}
