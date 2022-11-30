package entity

import "gorm.io/gorm"

type Role int8

const (
	Owner Role = iota + 1
	CoOwner
	Member
	KickedOut
)

type Group struct {
	ID             uint32       `json:"id"`
	AdminID        uint32       `json:"adminId"`
	Name           string       `json:"name"`
	CoverImageURL  string       `json:"coverImageUrl"`
	InvitationLink string       `json:"invitationLink"`
	Users          []*GroupUser `json:"users"  gorm:"-"`
	Kahoots        []*Kahoot    `json:"kahoots" gorm:"many2many:group_kahoots;"`
	gorm.Model
}

type Topic struct {
	ID            uint32 `json:"id"`
	Name          string `json:"name"`
	CoverImageUrl string `json:"coverImageUrl"`
}

type GroupUser struct {
	GroupID uint32 `json:"groupId" gorm:"primaryKey"`
	UserID  uint32 `json:"userId" gorm:"primaryKey"`
	Role    Role   `json:"role"`
	Name    string `json:"name"`
}
