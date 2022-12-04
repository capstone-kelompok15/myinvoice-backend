package dto

import "time"

const (
	AccountCTXKey = "ACCOUNT-CTX-KEY"
)

type AdminLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AdminContext struct {
	ID           int    `json:"id" db:"id"`
	MerchantID   int    `json:"merchant_id" db:"merchant_id"`
	MerchantName string `json:"merchant_name" db:"merchant_name"`
}

type AdminLoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type AdminRefreshToken struct {
	ID             int       `db:"id"`
	Token          string    `db:"token"`
	AdminID        int       `db:"admin_id"`
	IsValid        bool      `db:"is_valid"`
	ExpirationDate time.Time `db:"expired_date"`
}

type AdminRefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}
