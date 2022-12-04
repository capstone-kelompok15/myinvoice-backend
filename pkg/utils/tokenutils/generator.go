package tokenutils

import (
	"time"

	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/randomutils"

	"github.com/golang-jwt/jwt"
)

type RefreshTokenParams struct {
	UserID int
}

func NewRefreshToken(params *RefreshTokenParams) *dto.AdminRefreshToken {
	generatedToken := randomutils.GenerateNRandomString(128)
	return &dto.AdminRefreshToken{
		Token:          generatedToken,
		AdminID:        params.UserID,
		IsValid:        true,
		ExpirationDate: time.Now().Add(time.Hour * 24 * 7),
	}
}

func NewAccessToken(jwtSecret string, acc *dto.AdminContext) (string, error) {
	claims := dto.NewUserClaims(acc)
	unsignedToken := jwt.NewWithClaims(jwt.SigningMethodHS256, *claims)
	accessToken, err := unsignedToken.SignedString([]byte(jwtSecret))
	return accessToken, err
}
