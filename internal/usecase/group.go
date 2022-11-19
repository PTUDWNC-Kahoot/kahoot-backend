package usecase

import repo "examples/identity/internal/repository"

type groupUsecase struct {
	repo repo.GroupRepo
}

func NewGroupUsecase(repo repo.KahootRepo) GroupUsecase {
	return &groupUsecase{
		repo: repo,
	}
}
