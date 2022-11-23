package auth

import "examples/identity/internal/entity"

type AuthenRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthenResponse struct {
	Token string `json:"token"`

	ID            uint32           `json:"id"`
	Name          string           `json:"name"`
	Workplace     string           `json:"workplace"`
	Organization  string           `json:"organization"`
	CoverImageURL string           `json:"coverImageUrl"`
	Groups        []*entity.Group  `json:"groups"`
	Kahoots       []*entity.Kahoot `json:"kahoots"`
}

func (a AuthenRequest) Validate() bool {
	if a.Email == "" || a.Password == "" {
		return false
	}
	return true
}
