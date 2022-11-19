package service

import (
	"examples/identity/config"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
)

type JWTHelper interface {
	GenerateJWT(email string) (string, error)
	ValidateJWT(encodedToken string) (*jwt.Token, error)
}
type jwtService struct {
	secretKey string
}

type customClaims struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func getSecretKey(s *config.Specification) string {
	secret := s.SecretKey
	if secret == "" {
		secret = "secret"
	}
	return secret
}
func (service *jwtService) GenerateJWT(username string) (string, error) {
	claims := customClaims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 30).Unix(),
			Issuer:    "nameOfWebsiteHere",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(service.secretKey))
	if err != nil {
		return "", err
	}
	return tokenString, err
}
func (service *jwtService) ValidateJWT(encodedToken string) (*jwt.Token, error) {

	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, nil
		}
		return []byte(service.secretKey), nil
	})
}

func NewJWTService(s *config.Specification) JWTHelper {
	return &jwtService{
		secretKey: getSecretKey(s),
	}
}
