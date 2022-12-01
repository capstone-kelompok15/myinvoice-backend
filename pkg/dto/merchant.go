package dto

type MerchantRegisterRequest struct {
	Email               string                        `json:"email" validate:"required,email"`
	Password            string                        `json:"password" validate:"required,min=8"`
	Username            string                        `json:"username" validate:"required"`
	MerchantName        string                        `json:"merchant_name" validate:"required"`
	MerchantAddress     string                        `json:"merchant_address" validate:"required"`
	MerchantPhoneNumber string                        `json:"merchant_phone_number" validate:"required"`
	MerchantBank        []MerchantBankRegisterRequest `json:"merchant_banks" validate:"required,min=1"`
}

type MerchantBankRegisterRequest struct {
	BankID     int    `json:"bank_id" validate:"required"`
	OnBehalfOf string `json:"on_behalf_of" validate:"required"`
	BankNumber string `json:"bank_number" validate:"required"`
}
