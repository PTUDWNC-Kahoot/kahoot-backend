package repo

import "examples/identity/internal/entity"

type AuthRepo interface {
	Login(request *entity.User) (*entity.User, []*entity.Group, []*entity.Kahoot, error)
	Register(*entity.User) (uint32, error)
}

type KahootRepo interface {
}

type GroupRepo interface {
	Collection() ([]*entity.Group, error)
	GetOne(id uint32) (*entity.Group, error)
	CreateOne(request *entity.Group) (uint32, error) //generate invitation link
	UpdateOne(request *entity.Group) error
	DeleteOne(id uint32) error
	JoinGroupByLink(string, string) (*entity.Group, error)
}
