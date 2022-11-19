package v1

type IdentityRequest struct {
	TokenString string `json:"token"`
}
type IdentityResponse struct {
	IsValid bool
}
