package dto

type BankResponse struct {
	ID       int    `json:"id" db:"id"`
	BankName string `json:"bank_name" db:"bank_name"`
	Code     int    `json:"code" db:"code"`
}
