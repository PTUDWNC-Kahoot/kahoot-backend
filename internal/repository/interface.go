package repo

import "examples/identity/internal/entity"

type AuthRepo interface {
	Login(request *entity.User) bool
	Register(*entity.User) bool
}

type KahootRepo interface {
}