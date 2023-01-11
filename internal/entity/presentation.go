package entity

type SlideType int8

const (
	SlideTypeMultipleChoice    SlideType = iota + 1 // Multiplechoice: Question, Options, Points,Image
	SlideTypeMultipleParagraph                      // paragraph: heading, paragraph, image
	SildeTypeHeading                                // heading: heading, subheading, image
)

type Presentation struct {
	ID            uint32          `json:"id"`
	GroupID       uint32          `json:"-"`
	Owner         uint32          `json:"-"`
	Title         string          `json:"title"`
	Description   string          `json:"description"`
	CoverImageURL string          `json:"coverImageUrl"`
	Visibility    bool            `json:"visibility"`
	Slides        []*Slide        `json:"slides"`
	Collaborators []*Collaborator `json:"collaborators"`
}

type Slide struct {
	ID             uint32    `json:"id"`
	PresentationID uint32    `json:"presentationId"`
	Type           SlideType `json:"type"`
	Question       string    `json:"question"`
	Options        []*Option `json:"options"`
	Heading        string    `json:"heading"`
	SubHeading     string    `json:"subHeading"`
	Paragraph      string    `json:"paragraph"`
	ImageURL       string    `json:"imageUrl"`
}

type Option struct {
	ID        uint32 `json:"id"`
	SlideID   uint32 `json:"slideID"`
	ImageURL  string `json:"imageUrl"`
	Color     string `json:"color"`
	Content   string `json:"content"`
	IsCorrect bool   `json:"isCorrect"`
}

type Collaborator struct {
	ID             uint32 `json:"id"`
	UserID         uint32 `json:"userId"`
	Name           string `json:"name"`
	PresentationID uint32 `json:"presentationId"`
}
