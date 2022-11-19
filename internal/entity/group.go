package entity

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
	InvitationLink string    `json:"invitationLink"`
	Members        []*User   `json:"members"`
	Kahoots        []*Kahoot `json:"kahoots"`
}

type Topic struct {
	ID            uint32 `json:"id"`
	Name          string `json:"name"`
	CoverImageUrl string `json:"coverImageUrl"`
}

type GroupMember struct {
	GroupID  uint32 `json:"groupId"`
	MemberID uint32 `json:"memberId"`
	Role     Role   `json:"role"`
}
