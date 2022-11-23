package entity

import "gorm.io/gorm"

type Role int8

const (
	Owner Role = iota
	CoOwner
	Member
	KickedOut
)

type Group struct {
	ID             uint32    `json:"id"`
	AdminID        uint32    `json:"adminId"`
	Name           string    `json:"name"`
	CoverImageURL  string    `json:"coverImageUrl"`
	InvitationLink string    `json:"invitationLink"`
	Users          []*User   `json:"users" gorm:"many2many:group_users;ForeignKey:id;References:id"`
	Kahoots        []*Kahoot `json:"kahoots" gorm:"many2many:group_kahoots;"`
	gorm.Model
}

type Topic struct {
	ID            uint32 `json:"id"`
	Name          string `json:"name"`
	CoverImageUrl string `json:"coverImageUrl"`
}

type GroupMember struct {
	GroupID  uint32 `gorm:"primaryKey"`
	MemberID uint32 `gorm:"primaryKey"`
	Role     Role   `json:"role"`
}
