package usecase

import repo "examples/identity/internal/repository"

type kahootUsecase struct {
	repo repo.KahootRepo
}


func NewKahootUsecase(repo repo.KahootRepo) KahootUsecase {
	return &kahootUsecase{
		repo: repo,
	}

}
