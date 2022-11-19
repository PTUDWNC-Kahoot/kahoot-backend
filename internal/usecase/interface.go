package usecase

import "examples/identity/internal/entity"

type AuthUsecase interface {
	Login(request *entity.User) (string, error)
	Register(request *entity.User) (string, error)
}
type KahootUsecase interface {
}
type GroupUsecase interface {
}
