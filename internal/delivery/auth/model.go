package auth

type AuthenRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthenResponse struct {
	Token string `json:"token"`
}

func (a AuthenRequest) Validate() bool {
	if a.Email == "" || a.Password == "" {
		return false
	}
	return true
}
