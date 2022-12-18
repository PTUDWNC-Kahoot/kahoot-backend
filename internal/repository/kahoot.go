package repo

import (
	"examples/kahootee/internal/usecase"

	"gorm.io/gorm"
)

type kahootRepo struct {
	db *gorm.DB
}

func NewKahootRepo(db *gorm.DB) usecase.KahootRepo {
	return &kahootRepo{
		db: db,
	}
}
