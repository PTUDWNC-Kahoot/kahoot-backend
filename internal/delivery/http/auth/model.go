package auth

import "examples/kahootee/internal/entity"

type AuthenRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterWithVerification struct {
	Email      string `json:"email"`
	Password   string `json:"password"`
	VerifyCode int    `json:"verifyCode"`
}

type AuthenResponse struct {
	Token         string           `json:"token"`
	ID            uint32           `json:"id"`
	Name          string           `json:"name"`
	Workplace     string           `json:"workplace"`
	Organization  string           `json:"organization"`
	CoverImageURL string           `json:"coverImageUrl"`
	Groups        []*entity.Group  `json:"groups"`
	Kahoots       []*entity.Kahoot `json:"kahootsList"`
}

type GoogleResponse struct {
	Email           string `json:"email"`
	Name            string `json:"name"`
	Picture         string `json:"picture"`
	IsEmailVerified bool   `json:"isEmailVerified"`
}

func (a AuthenRequest) Validate() bool {
	if a.Email == "" || a.Password == "" {
		return false
	}
	return true
}

func (a RegisterWithVerification) Validate() bool {
	if a.Email == "" || a.Password == "" {
		return false
	}
	return true
}
