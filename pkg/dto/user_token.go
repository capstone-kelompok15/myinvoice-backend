package dto

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type UserClaims struct {
	jwt.StandardClaims
	UserContext
}

func NewUserClaims(acc *UserContext) *UserClaims {
	return &UserClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(1 * time.Hour).Unix(),
		},
		UserContext: UserContext{
			ID: acc.ID,
		},
	}
}
