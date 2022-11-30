package repo

import (
	"examples/identity/internal/entity"
	"fmt"

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
	if err := g.db.Preload("Users").First(&group).Error; err != nil {
		return nil, err
	}

	return group, nil
}

func (g *groupRepo) CreateOne(request *entity.Group) (uint32, error) {
	err := g.db.Create(&request).Error
	if err != nil {
		return 0, err
	}
	err = g.db.Create(&entity.GroupUser{
		GroupID: request.ID,
		UserID:  request.AdminID,
		Role:    entity.Owner,
	}).Error
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
func (g *groupRepo) JoinGroupByLink(userEmail string, groupCode string) (*entity.Group, error) {
	group := &entity.Group{}

	err := g.db.Where("invitation_link=?", groupCode).Preload("Users").First(group).Error
	if group.ID == 0 || err != nil {
		return nil, err
	}

	user := &entity.User{}
	err = g.db.Where("email=?", userEmail).First(user).Error
	if user.ID == 0 || err != nil {
		return nil, err
	}

	existedUser := &entity.GroupUser{}
	err = g.db.Where("email=?", userEmail).First(existedUser).Error
	if existedUser.UserID != 0 {
		return nil, err
	}

	groupUser := &entity.GroupUser{
		GroupID: group.ID,
		UserID:  user.ID,
		Role:    entity.Member,
	}

	if err := g.db.Model(groupUser).Create(groupUser).Error; err != nil {
		return nil, err
	}

	return group, nil
}

func (g *groupRepo) Invite(email_list []string, groupID uint32) error {
	users := []uint32{}
	fmt.Println("emaillist: ", email_list)
	for _, email := range email_list {
		user := entity.User{}
		err := g.db.Where("email=?", email).First(&user).Error
		if err != nil {
			continue
		}

		existed := &entity.GroupUser{}
		g.db.Where("user_id=?", user.ID).Where("group_id=?", groupID).First(existed)
		if existed.UserID != 0 {
			continue
		}
		users = append(users, user.ID)
	}
	fmt.Println("id_list", users)
	groupUsers := []*entity.GroupUser{}
	for _, userID := range users {
		groupUser := &entity.GroupUser{
			GroupID: groupID,
			UserID:  userID,
			Role:    entity.Member,
		}
		groupUsers = append(groupUsers, groupUser)
	}
	return g.db.Create(&groupUsers).Error
}

func (g *groupRepo) AssignRole(groupUser *entity.GroupUser, ownerEmail string) error {
	user := &entity.User{}

	if err := g.db.Where("email=?", ownerEmail).First(user).Error; err != nil {
		return err
	}

	owner := &entity.GroupUser{}

	if err := g.db.Where("user_id=?", user.ID).First(owner).Error; err != nil {
		return err
	}

	if owner.Role != entity.Owner || owner.UserID == groupUser.UserID {
		return fmt.Errorf("do not have permission")
	}

	return g.db.Model(groupUser).Updates(groupUser).Error
}
