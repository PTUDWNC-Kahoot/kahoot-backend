package usecase

import (
	"examples/kahootee/internal/entity"
)

type AuthUsecase interface { //nolint:dupl
	Login(request *entity.User) (*entity.User, []*entity.Group, []*entity.Presentation, string, error)
	Register(request *entity.User) error
	CreateRegisterOrder(*entity.RegisterOrder) (uint32, error)
	VerifyEmail(string, int) bool
	CheckEmailExisted(string) bool
}

type AuthRepo interface { //nolint:dupl
	Login(request *entity.User) (*entity.User, []*entity.Group, []*entity.Presentation, error)
	Register(*entity.User) error
	CreateRegisterOrder(*entity.RegisterOrder) (uint32, error)
	VerifyEmail(string, int) bool
	CheckEmailExisted(string) bool
}

type KahootUsecase interface {
}
type GroupUsecase interface { //nolint:dupl
	GetGroups() ([]*entity.Group, error)
	Get(id uint32) (*entity.Group, error)
	Create(request *entity.Group) (uint32, error)
	Update(request *entity.Group) error
	Delete(id uint32) error
	JoinGroupByLink(string, string) (*entity.Group, error)
	Invite([]string, uint32) error
	AssignRole(*entity.GroupUser, string) error
}

type GroupRepo interface { //nolint:dupl
	Collection() ([]*entity.Group, error)
	GetOne(id uint32) (*entity.Group, error)
	CreateOne(request *entity.Group) (uint32, error)
	UpdateOne(request *entity.Group) error
	DeleteOne(id uint32) error
	JoinGroupByLink(string, string) (*entity.Group, error)
	Invite([]string, uint32) error
	AssignRole(*entity.GroupUser, string) error
}

type KahootRepo interface {
}

type User interface {
	GetProfile(id uint32) (*entity.User, error)
	UpdateProfile(user *entity.User) error
	DeleteProfile(id uint32) error
}

type UserRepo interface {
	GetProfile(id uint32) (*entity.User, error)
	UpdateProfile(user *entity.User) error
	DeleteProfile(id uint32) error
}
