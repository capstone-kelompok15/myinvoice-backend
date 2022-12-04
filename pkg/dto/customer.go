package dto

type CustomerRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
	FullName string `json:"full_name" validate:"required"`
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
