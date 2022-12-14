package dto

import "time"

const (
	AdminCTXKey = "ADMIN-CTX-KEY"
)

type AdminLoginRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
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

type AdminEmailVerification struct {
	Email string `json:"email" validate:"required,email"`
	Code  string `json:"code" validate:"required,min=32,max=32"`
}

type AdminRefreshEmailVerificationCode struct {
	Email string `json:"email" validate:"required,email"`
}

type AdminResetPasswordRequest struct {
	Email string `json:"email" validate:"required,email"`
}
