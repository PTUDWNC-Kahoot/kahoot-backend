package repo

import (
	"examples/kahootee/internal/entity"
	"examples/kahootee/internal/usecase"

	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) usecase.UserRepo {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) GetProfile(id uint32) (*entity.User, error) {
	user := &entity.User{}
	if err := r.db.Debug().First(user, id).Error; err != nil {
		return nil, err
	}

	return user, nil
}
func (r *userRepo) UpdateProfile(user *entity.User) error {
	return r.db.Updates(user).Error
}
func (r *userRepo) DeleteProfile(id uint32) error {
	return r.db.Delete(&entity.User{}, id).Error
}
