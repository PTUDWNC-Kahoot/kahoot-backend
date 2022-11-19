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

func (u *authUsecase) Login(request *entity.User) (string, error) {
	isUserAuthenticated := u.repo.Login(request)

	var token string
	var err error = nil

	if isUserAuthenticated {
		token, err = u.jwtService.GenerateJWT(request.Email)
		return token, err
	}
	return "", err
}
func (u *authUsecase) Register(request *entity.User) (string, error) {
	isUserAuthenticated := u.repo.Register(request)

	var token string
	var err error = nil

	if isUserAuthenticated {
		token, err = u.jwtService.GenerateJWT(request.Email)
		return token, err
	}
	return "", err
}
