package auth

import "github.com/golang-jwt/jwt/v5"

type Service interface {
	GenerateToken(user int) (string, error)
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
