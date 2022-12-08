package dto

type BankResponse struct {
	ID       int    `json:"id" db:"id"`
	BankName string `json:"bank_name" db:"bank_name"`
	Code     string `json:"code" db:"code"`
}
