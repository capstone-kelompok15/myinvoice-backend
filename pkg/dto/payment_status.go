package dto

type PaymentStatus struct {
	ID         int    `json:"id" db:"id"`
	StatusName string `json:"status_name" db:"status_name"`
}
