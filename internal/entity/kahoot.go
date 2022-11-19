package entity

type Kahoot struct {
	ID            int    `json:"id"`
	AccountID     int    `json:"accountId"`
	Title         string `json:"title"`
	Description   string `json:"description"`
	CoverImageURL string `json:"coverImageUrl"`
	Visibility    bool   `json:"visibility"`
}

type Slide struct {
	ID            int    `json:"id"`
	KahootID      int    `json:"kahootId"`
	Type          string `json:"type"`
	Order         int    `json:"order"`
	Question      string `json:"question"`
	TimeLimit     int    `json:"timeLimit"`
	Points        int    `json:"points"`
	ImageURL      string `json:"imageUrl"`
	VideoURL      string `json:"videoUrl"`
	AnswerPptions string `json:"answerPptions"`
	Title         string `json:"title"`
	Text          string `json:"text"`
}

type Answer struct {
	ID        int    `json:"id"`
	KahootID  int    `json:"kahootId"`
	ImageURL  string `json:"imageUrl"`
	Color     string `json:"color"`
	Content   string `json:"content"`
	IsCorrect bool   `json:"isCorrect"`
	Order     int    `json:"order"`
}
