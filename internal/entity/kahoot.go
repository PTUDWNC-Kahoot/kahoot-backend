package entity

type Kahoot struct {
	Id              int
	Account_id      int
	Title           string
	Description     string
	Cover_image_url string
	Visibility      bool
}

type Slide struct {
	Id             int
	Kahoot_id      int
	Type          string
	Order          int
	Question       string
	Time_limit     int
	Points         int
	Image_url      string
	Video_url      string
	Answer_options string
	Title          string
	Text           string
}

type Answer struct {
	Id         int
	Kahoot_id  int
	Image_url  string
	Color      string
	Content    string
	Is_correct bool
	Order      int
}
