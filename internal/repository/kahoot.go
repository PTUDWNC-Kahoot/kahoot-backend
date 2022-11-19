package repo

import "gorm.io/gorm"

type kahootRepo struct {
	db *gorm.DB
}

func NewKahootRepo(db *gorm.DB) KahootRepo {
	return &kahootRepo{
		db: db,
	}
}
