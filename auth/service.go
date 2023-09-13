package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Service interface {
	GenerateToken(userID int) (string, error)
}

type jwtService struct {

}

func NewSevice() *jwtService {
	return &jwtService{}
}

func (s *jwtService) GenerateToken(userID int) (string, error){
	claim := jwt.MapClaims{}
	claim["user_id"] = userID
	claim["exp"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedToken, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return signedToken, err
	}
	return signedToken, nil
}