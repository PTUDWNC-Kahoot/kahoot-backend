package entity

type Group struct {
	Id              int
	Admin_id        int
	Name            string
	Invitation_code string
	Members         []*User
	Kahoots         []*Kahoot
}

type Topic struct {
	Id              int
	Name            string
	Cover_image_url string
}
