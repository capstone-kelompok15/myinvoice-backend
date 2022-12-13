package dto

import "github.com/capstone-kelompok15/myinvoice-backend/config"

const (
	CustomerCTXKey = "CUSTOMER-CTX-KEY"
)

type CustomerUpdateRequest struct {
	FullName string `json:"full_name" validate:"required"`
	Address  string `json:"address" validate:"required"`
}
type CustomerRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
	FullName string `json:"full_name" validate:"required"`
}
type GetAllCustomerRequest struct {
	Page  uint64 `query:"page" validate:"required"`
	Limit uint64 `query:"limit" validate:"required"`
	Name  string `query:"name"`
	Email string `query:"email"`
}

type GetAllCustomerRespond struct {
	ID                int     `json:"id" db:"id"`
	FullName          string  `json:"full_name" db:"full_name"`
	Email             string  `json:"email" db:"email"`
	DisplayProfileURL *string `json:"display_profile_url" db:"display_profile_url"`
	Address           string  `json:"address" db:"address"`
	CreatedAt         string  `json:"created_at" db:"created_at"`
	UpdatedAt         string  `json:"updated_at" db:"updated_at"`
}

type CustomerEmailVerification struct {
	Email string `json:"email" validate:"required,email"`
	Code  string `json:"code" validate:"required,min=4,max=4"`
}

type CustomerRefreshEmailVerificationcode struct {
	Email string `json:"email" validate:"required,email"`
}

type CustomerLoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
	DeviceID string `json:"device_id" validate:"required"`
}

type CustomerContext struct {
	ID       int    `json:"id" db:"id"`
	DeviceID string `json:"device_id"`
	FullName string `json:"full_name" db:"full_name"`
}

type CustomerAccessToken struct {
	AccessToken string `json:"access_token"`
}

type CustomerResetPasswordRequest struct {
	Email string `json:"email"`
}

type CustomerResetPassword struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Code     string `json:"code"`
}

type CustomerDetails struct {
	ID                       int     `json:"id" db:"id"`
	Email                    string  `json:"email" db:"email"`
	FullName                 string  `json:"full_name" db:"full_name"`
	DisplayProfilePictureURL *string `json:"display_profile_picture_url" db:"display_profile_url"`
}

type CustomerAccessTokenParams struct {
	DeviceInformation string
	UserInformation   *CustomerContext
	Config            *config.CustomerToken
}

type HeaderCustomerTokenPart struct {
	DeviceID string `json:"device_id"`
	Date     int64  `json:"issuer_at"`
}

type PayloadCustomerTokenPart struct {
	ID       int    `json:"id"`
	FullName string `json:"full_name"`
}

type UpdateProfilePictureResponse struct {
	ImageURL string `json:"image_url"`
}

type BriefCustomer struct {
	ID       int    `json:"id" db:"id"`
	FullName string `json:"full_name" db:"full_name"`
	Email    string `json:"email" db:"email"`
}

type CustomerSummary struct {
	TotalPaid   int64 `json:"total_paid" db:"total_paid"`
	TotalUnpaid int64 `json:"total_unpaid" db:"total_unpaid"`
}
