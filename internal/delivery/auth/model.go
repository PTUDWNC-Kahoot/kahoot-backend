package auth

type AuthenRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthenResponse struct {
	Token string `json:"token"`
}
