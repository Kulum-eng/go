package application

import "association/domain"

type GetAssociationUseCase struct {
	repo domain.AssociationRepository
}

func NewGetAssociationUseCase(repo domain.AssociationRepository) *GetAssociationUseCase {
	return &GetAssociationUseCase{repo: repo}
}

func (useCase *GetAssociationUseCase) Execute(id int) (*domain.Association, error) {
	return useCase.repo.GetByID(id)
}

func (useCase *GetAssociationUseCase) ExecuteAll() ([]domain.Association, error) {
	return useCase.repo.GetAll()
}
