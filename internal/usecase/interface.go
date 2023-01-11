package usecase

import (
	"examples/kahootee/internal/entity"
)

type Auth interface { //nolint:dupl
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

type Group interface { //nolint:dupl
	GetGroups(userid uint32) ([]*entity.Group, error)
	Get(id uint32) (*entity.Group, error)
	Create(request *entity.Group, user *entity.User) (uint32, error)
	Update(request *entity.Group) error
	Delete(id uint32) error
	JoinGroupByLink(string, string) (*entity.Group, error)
	Invite([]string, uint32) error
	AssignRole(*entity.GroupUser, string) error
}

type GroupRepo interface { //nolint:dupl
	Collection(userId uint32) ([]*entity.Group, error)
	GetOne(id uint32) (*entity.Group, error)
	CreateOne(request *entity.Group, user *entity.User) (uint32, error)
	UpdateOne(request *entity.Group) error
	DeleteOne(id uint32) error
	JoinGroupByLink(string, string) (*entity.Group, error)
	Invite([]string, uint32) error
	AssignRole(*entity.GroupUser, string) error
}

type User interface {
	GetProfile(id uint32) (*entity.User, error)
	UpdateProfile(user *entity.User) error
	DeleteProfile(id uint32) error
	GetSite(email string) (*entity.User, error)
}

type UserRepo interface {
	GetProfile(id uint32) (*entity.User, error)
	UpdateProfile(user *entity.User) error
	DeleteProfile(id uint32) error
	GetSite(email string) (*entity.User, error)
}

type Presentation interface {
	CreatePresentation(request *entity.Presentation) error
	GetPresentation(id uint32) (*entity.Presentation, error)
	GetPresentationByCode(code string) (*entity.Presentation, error)
	GroupCollection(groupId uint32) ([]*entity.Presentation, error)
	MyCollection(userId uint32) ([]*entity.Presentation, error)
	UpdatePresentation(request *entity.Presentation) error
	DeletePresentation(id uint32, userId uint32) error

	CreateSlide(slide *entity.Slide) error
	UpdateSlide(slide *entity.Slide) error
	DeleteSlide(id uint32) error

	AddCollaborator([]string, uint32) error
	GetCollaborators(presentationID uint32) ([]*entity.Collaborator, error)
	RemoveCollaborator(userID, presentationID uint32) error
}
type PresentationRepo interface {
	CreatePresentation(request *entity.Presentation) error
	GetPresentation(id uint32) (*entity.Presentation, error)
	GetPresentationByCode(code string) (*entity.Presentation, error)
	GroupCollection(groupId uint32) ([]*entity.Presentation, error)
	MyCollection(userId uint32) ([]*entity.Presentation, error)
	UpdatePresentation(request *entity.Presentation) error
	DeletePresentation(id uint32, userId uint32) error

	CreateSlide(slide *entity.Slide) error
	UpdateSlide(slide *entity.Slide) error
	DeleteSlide(id uint32) error

	AddCollaborator([]string, uint32) error
	GetCollaborators(presentationID uint32) ([]*entity.Collaborator, error)
	RemoveCollaborator(userID, presentationID uint32) error
}
