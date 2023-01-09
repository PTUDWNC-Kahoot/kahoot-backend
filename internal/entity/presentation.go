package entity

type Presentation struct {
	ID            uint32  `json:"id"`
	GroupID       uint32  `json:"-"`
	UserID        uint32  `json:"-"`
	Title         string  `json:"title"`
	Description   string  `json:"description"`
	CoverImageURL string  `json:"coverImageUrl"`
	Visibility    bool    `json:"visibility"`
	Slides        []*Slide `json:"slides"`
}

type Slide struct {
	ID             uint32 `json:"id"`
	PresentationID uint32 `json:"presentationId"`
	Type           string `json:"type"`
	Order          int8   `json:"order"`
	Question       string `json:"question"`
	TimeLimit      int8   `json:"timeLimit"`
	Points         int8   `json:"points"`
	ImageURL       string `json:"imageUrl"`
	VideoURL       string `json:"videoUrl"`
	AnswerOptions  string `json:"answerOptions"`
	Title          string `json:"title"`
	Text           string `json:"text"`
}

type Answer struct {
	ID        uint32 `json:"id"`
	KahootID  uint32 `json:"kahootId"`
	ImageURL  string `json:"imageUrl"`
	Color     string `json:"color"`
	Content   string `json:"content"`
	IsCorrect bool   `json:"isCorrect"`
	Order     int8   `json:"order"`
}
