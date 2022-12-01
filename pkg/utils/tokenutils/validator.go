package tokenutils

import (
	"github.com/capstone-kelompok15/myinvoice-backend/config"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
	customerrors "github.com/capstone-kelompok15/myinvoice-backend/pkg/errors"
	"github.com/golang-jwt/jwt"
)

func ValidateAccessToken(conf *config.JWTConfig, accessToken string) (*dto.AdminContext, error) {
	token, err := jwt.Parse(accessToken, func(t *jwt.Token) (interface{}, error) {
		if method, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, customerrors.ErrUnauthorized
		} else if method != jwt.SigningMethodHS256 {
			return nil, customerrors.ErrUnauthorized
		}

		return []byte(conf.JWTSecretKey), nil
	})

	if err != nil {
		return nil, customerrors.ErrUnauthorized
	}

	mapClaims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, customerrors.ErrUnauthorized
	}
	user := dto.AdminContext{
		ID: int(mapClaims["id"].(float64)),
	}

	return &user, nil
}
