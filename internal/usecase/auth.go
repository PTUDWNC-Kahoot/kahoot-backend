package usecase

import (
	"examples/identity/internal/entity"
	repo "examples/identity/internal/repository"
	service "examples/identity/internal/service/jwthelper"
)

type authUsecase struct {
	repo       repo.AuthRepo
	jwtService service.JWTHelper
}

func NewAuthUsecase(repo repo.AuthRepo, jwtService service.JWTHelper) AuthUsecase {
	return &authUsecase{
		repo:       repo,
		jwtService: jwtService,
	}
}

func (u *authUsecase) Login(request *entity.User) (*entity.User, []*entity.Group, []*entity.Kahoot, string, error) {
	user, groups, kahoots, err := u.repo.Login(request)
	if err != nil || user.ID == 0 {
		return nil, nil, nil, "", err
	}

	var token string

	token, err = u.jwtService.GenerateJWT(request.Email)
	if err != nil {
		return nil, nil, nil, "", err
	}
	return user, groups, kahoots, token, nil
}

func (u *authUsecase) Register(request *entity.User) (uint32, string, error) {
	id, err := u.repo.Register(request)
	if err != nil || id == 0 {
		return 0, "", err
	}

	var token string

	token, err = u.jwtService.GenerateJWT(request.Email)
	if err != nil {
		return 0, "", err
	}
	return id, token, nil
}
