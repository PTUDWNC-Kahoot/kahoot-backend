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
func (g *groupRepo) JoinGroupByLink(userEmail string, groupCode string) (*entity.Group, error) {
	group := &entity.Group{}

	err := g.db.Where("invitation_link=?", groupCode).First(group).Error
	if group.ID == 0 || err != nil {
		return nil, err
	}

	user := &entity.User{}
	err = g.db.Where("email=?", userEmail).First(user).Error
	if user.ID == 0 || err != nil {
		return nil, err
	}

	existedMember := &entity.GroupMember{}
	err = g.db.Where("email=?", userEmail).First(existedMember).Error
	if existedMember.MemberID != 0 {
		return nil, err
	}

	groupMember := &entity.GroupMember{
		GroupID:  group.ID,
		MemberID: user.ID,
		Role:     entity.Member,
	}

	if err := g.db.Model(groupMember).Create(groupMember).Error; err != nil {
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

		existed := &entity.GroupMember{}
		g.db.Where("member_id=?", user.ID).Where("group_id=?", groupID).First(existed)
		if existed.MemberID != 0 {
			continue
		}
		users = append(users, user.ID)
	}
	fmt.Println("id_list", users)
	groupMembers := []*entity.GroupMember{}
	for _, userID := range users {
		groupMember := &entity.GroupMember{
			GroupID:  groupID,
			MemberID: userID,
			Role:     entity.Member,
		}
		groupMembers = append(groupMembers, groupMember)
	}
	return g.db.Create(&groupMembers).Error
}
