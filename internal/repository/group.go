package repo

import (
	"examples/identity/internal/entity"

	"gorm.io/gorm"
)

type groupRepo struct {
	db *gorm.DB
}

func NewGroupRepo(db *gorm.DB) GroupRepo {
	return &groupRepo{
		db: db,
	}
}

func (g *groupRepo) Collection() ([]*entity.Group, error) {
	group := []*entity.Group{}
	err := g.db.Find(&group).Error
	if err != nil {
		return nil, err
	}
	return group, nil
}

func (g *groupRepo) GetOne(id uint32) (*entity.Group, error) {
	group := &entity.Group{ID: id}
	err := g.db.First(&group).Error
	if err != nil {
		return nil, err
	}
	return group, nil
}

func (g *groupRepo) CreateOne(request *entity.Group) (uint32, error) {
	err := g.db.Create(&request).Error
	if err != nil {
		return 0, err
	}
	return request.ID, nil
}

func (g *groupRepo) UpdateOne(request *entity.Group) error {
	return g.db.Updates(&request).Error
}

func (g *groupRepo) DeleteOne(id uint32) error {
	return g.db.Delete(&entity.Group{ID: id}).Error
}
func (g *groupRepo) JoinGroupByLink(groupCode string) (*entity.Group, error) {
	group := &entity.Group{}
	err := g.db.Where("invitation_link=?", groupCode).First(group).Error
	if group.ID == 0 || err != nil {
		return nil, err
	}
	return group, nil
}
