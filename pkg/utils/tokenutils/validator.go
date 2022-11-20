package tokenutils

import (
	"github.com/capstone-kelompok15/myinvoice-backend/config"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/errors"
	"github.com/golang-jwt/jwt"
)

func ValidateAccessToken(conf *config.JWTConfig, accessToken string) (*dto.UserContext, error) {
	token, err := jwt.Parse(accessToken, func(t *jwt.Token) (interface{}, error) {
		if method, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.ErrUnauthorized
		} else if method != jwt.SigningMethodHS256 {
			return nil, errors.ErrUnauthorized
		}

		return []byte(conf.JWTSecretKey), nil
	})

	if err != nil {
		return nil, errors.ErrUnauthorized
	}

	mapClaims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.ErrUnauthorized
	}
	user := dto.UserContext{
		ID: int(mapClaims["id"].(float64)),
	}

	return &user, nil
}
