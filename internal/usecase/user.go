package usecase

import (
	"examples/kahootee/internal/entity"
)

type user struct {
	repo UserRepo
}

func NewUser(repo UserRepo) User {
	return &user{
		repo: repo,
	}
}

func (u *user) GetProfile(id uint32) (*entity.User, error) {
	return u.repo.GetProfile(id)
}
func (u *user) UpdateProfile(user *entity.User) error {
	return u.repo.UpdateProfile(user)
}
func (u *user) DeleteProfile(id uint32) error {
	return u.repo.DeleteProfile(id)
}
