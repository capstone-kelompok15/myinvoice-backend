package dto

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type UserClaims struct {
	jwt.StandardClaims
	AdminContext
}

func NewUserClaims(acc *AdminContext) *UserClaims {
	return &UserClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(1 * time.Hour).Unix(),
		},
		AdminContext: AdminContext{
			ID: acc.ID,
		},
	}
}
