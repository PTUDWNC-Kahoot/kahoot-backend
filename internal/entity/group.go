package entity

type Group struct {
	ID             int       `json:"id"`
	AdminID        int       `json:"adminId"`
	Name           string    `json:"name"`
	InvitationLink string    `json:"invitationLink"`
	Members        []*User   `json:"members"`
	Kahoots        []*Kahoot `json:"kahoots"`
}

type Topic struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	CoverImageUrl string `json:"coverImageUrl"`
}
