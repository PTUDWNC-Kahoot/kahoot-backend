package usecase

import "examples/kahootee/internal/entity"

type AuthUsecase interface {
	Login(request *entity.User) (*entity.User, []*entity.Group, []*entity.Kahoot, string, error)
	Register(request *entity.User) error
	CreateRegisterOrder(*entity.RegisterOrder) (uint32, error)
	VerifyEmail(string, int) bool
	CheckEmailExisted(string) bool
}
type KahootUsecase interface {
}
type GroupUsecase interface {
	GetGroups() ([]*entity.Group, error)
	Get(id uint32) (*entity.Group, error)
	Create(request *entity.Group) (uint32, error) //generate invitation link
	Update(request *entity.Group) error
	Delete(id uint32) error
	JoinGroupByLink(string, string) (*entity.Group, error)
	Invite([]string, uint32) error
	AssignRole(*entity.GroupUser, string) error
}


type AuthRepo interface {
	Login(request *entity.User) (*entity.User, []*entity.Group, []*entity.Kahoot, error)
	Register(*entity.User) error
	CreateRegisterOrder(*entity.RegisterOrder) (uint32, error)
	VerifyEmail(string, int) bool
	CheckEmailExisted(string) bool
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
	Invite([]string, uint32) error
	AssignRole(*entity.GroupUser, string) error
}