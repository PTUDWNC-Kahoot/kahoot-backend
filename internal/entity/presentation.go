package entity

type SlideType int8

const (
	SlideTypeGeneral SlideType = iota + 1
	SlideTypeMultipleChoice
)

type Presentation struct {
	ID            uint32   `json:"id"`
	GroupID       uint32   `json:"-"`
	UserID        uint32   `json:"-"`
	Title         string   `json:"title"`
	Description   string   `json:"description"`
	CoverImageURL string   `json:"coverImageUrl"`
	Visibility    bool     `json:"visibility"`
	Slides        []*Slide `json:"slides"`
}

type Slide struct {
	ID             uint32    `json:"id"`
	PresentationID uint32    `json:"presentationId"`
	Type           SlideType `json:"type"`
	Question       string    `json:"question"`
	TimeLimit      int8      `json:"timeLimit"`
	Points         int8      `json:"points"`
	ImageURL       string    `json:"imageUrl"`
	VideoURL       string    `json:"videoUrl"`
	Options        []*Option `json:"options"`
	Title          string    `json:"title"`
	Text           string    `json:"text"`
}

type Option struct {
	ID        uint32 `json:"id"`
	SlideID   uint32 `json:"slideID"`
	ImageURL  string `json:"imageUrl"`
	Color     string `json:"color"`
	Content   string `json:"content"`
	IsCorrect bool   `json:"isCorrect"`
}
