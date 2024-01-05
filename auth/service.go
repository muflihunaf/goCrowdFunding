package auth

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

type Service interface {
	GenerateToken(user int) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtService struct {
}

func NewService() *jwtService {
	return &jwtService{}
}

var SECRET_KEY = []byte("firmancuy_s3cr3t_k3y")

func (s *jwtService) GenerateToken(userId int) (string, error) {
	claim := jwt.MapClaims{}
	claim["user_id"] = userId

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return signedToken, err
	}
	return signedToken, nil
}

func (s *jwtService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		_, oke := token.Method.(*jwt.SigningMethodHMAC)
		if !oke {
			return nil, errors.New("invalid token")
		}

		return []byte(SECRET_KEY), nil
	})
	if err != nil {
		return token, err
	}
	return token, nil
}
