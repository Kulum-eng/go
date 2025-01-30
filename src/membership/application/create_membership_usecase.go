package application

import (
	"membership/domain"
	"membership/domain/ports"

)

type CreateMembershipUseCase struct {
	repo ports.MembershipRepository
}

func NewCreateMembershipUseCase(repo ports.MembershipRepository) *CreateMembershipUseCase {
	return &CreateMembershipUseCase{repo: repo}
}

func (uc *CreateMembershipUseCase) Execute(membership domain.Membership) error {
	return uc.repo.CreateMembership(membership)
}
