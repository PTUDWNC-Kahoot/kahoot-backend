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
	isUserAuthenticated := u.repo.Login(&entity.User{Id: request.Id, Email: request.Email, Password: request.Password})

	var token string
	var err error = nil

	if isUserAuthenticated {
		token, err = u.jwtService.GenerateJWT(request.Email)
		return token, err
	}
	return "", err
}
func (u *authUsecase) Register(request *entity.User) (string, error) {
	isUserAuthenticated := u.repo.Register(&entity.User{Id: request.Id, Email: request.Email, Password: request.Password})

	var token string
	var err error = nil

	if isUserAuthenticated {
		token, err = u.jwtService.GenerateJWT(request.Email)
		return token, err
	}
	return "", err
}
