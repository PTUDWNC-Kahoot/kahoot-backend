package usecase

import (
	"examples/kahootee/internal/entity"
	service "examples/kahootee/internal/service/jwthelper"
	"examples/kahootee/pkg/errors"
)

type authUsecase struct {
	repo       AuthRepo
	jwtService service.JWTHelper
}

func NewAuthUsecase(repo AuthRepo, jwtService service.JWTHelper) Auth {
	return &authUsecase{
		repo:       repo,
		jwtService: jwtService,
	}
}

func (u *authUsecase) Login(request *entity.User) (*entity.User, []*entity.Group, []*entity.Presentation, string, error) {
	user, groups, kahoots, err := u.repo.Login(request)
	if err != nil {
		return nil, nil, nil, "", err
	}

	var token string
	if user.ID == 0 {
		return nil, nil, nil, "", errors.ErrGeneral
	}

	token, err = u.jwtService.GenerateJWT(request.Email, user.ID)
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
