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
