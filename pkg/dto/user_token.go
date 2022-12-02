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
			ExpiresAt: time.Now().Add(2 * time.Hour).Unix(),
		},
		AdminContext: AdminContext{
			ID:           acc.ID,
			MerchantID:   acc.MerchantID,
			MerchantName: acc.MerchantName,
		},
	}
}
