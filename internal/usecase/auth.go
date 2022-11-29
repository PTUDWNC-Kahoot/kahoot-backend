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

func (u *authUsecase) Register(request *entity.User) error {
	return u.repo.Register(request)
}

func (u *authUsecase) CreateRegisterOrder(request *entity.RegisterOrder) (uint32, error) {
	id, err := u.repo.CreateRegisterOrder(request)
	if err != nil || id == 0 {
		return 0, err
	}
	return id, nil
}

func (u *authUsecase) VerifyEmail(email string, verifyCode int) bool {
	return u.repo.VerifyEmail(email, verifyCode)
}

func (u *authUsecase) CheckEmailExisted(email string) bool {
	return u.repo.CheckEmailExisted(email)
}
