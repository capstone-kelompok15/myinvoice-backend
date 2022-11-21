package tokenutils

import (
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"

	"github.com/golang-jwt/jwt"
)

func NewAccessToken(jwtSecret string, acc *dto.UserContext) (string, error) {
	claims := dto.NewUserClaims(acc)
	unsignedToken := jwt.NewWithClaims(jwt.SigningMethodHS256, *claims)
	accessToken, err := unsignedToken.SignedString([]byte(jwtSecret))
	return accessToken, err
}
