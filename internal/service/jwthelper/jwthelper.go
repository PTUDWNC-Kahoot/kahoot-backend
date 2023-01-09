package service

import (
	"errors"
	"fmt"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"

	"examples/kahootee/config"
)

type JWTHelper interface {
	GenerateJWT(email string, id uint32) (string, error)
	ValidateJWT(encodedToken string) (*customClaims, error)
}
type jwtService struct {
	secretKey string
}

var ErrExpiredToken = errors.New("token has expired")

type customClaims struct {
	ID    uint32 `json:"id"`
	Email string `json:"email"`
	jwt.StandardClaims
}

func getSecretKey(s *config.Specification) string {
	secret := s.SecretKey
	if secret == "" {
		secret = "secret"
	}
	return secret
}
func (service *jwtService) GenerateJWT(email string, id uint32) (string, error) {
	claims := customClaims{
		ID:    id,
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			Issuer:    "whatisit",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(service.secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, err
}

func (service *jwtService) ValidateJWT(token string) (*customClaims, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("Unexpected signing method")
		}
		return []byte(service.secretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &customClaims{}, keyFunc)
	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, ErrExpiredToken) {
			return nil, errors.New("token expired")
		}
		return nil, errors.New("token invalid")
	}

	payload, ok := jwtToken.Claims.(*customClaims)
	if !ok {
		return nil, errors.New("token invalid")
	}
	fmt.Println("PAYLOAD", payload)
	return payload, nil
}

func NewJWTService(s *config.Specification) JWTHelper {
	return &jwtService{
		secretKey: getSecretKey(s),
	}
}
