package repo

import "gorm.io/gorm"

type groupRepo struct {
	db *gorm.DB
}

func NewGroupRepo(db *gorm.DB) GroupRepo {
	return &groupRepo{
		db: db,
	}
}
