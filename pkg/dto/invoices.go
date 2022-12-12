package dto

import "time"

type CreateInvoiceRequest struct {
	CustomerID   int                          `json:"customer_id" validate:"required"`
	DueAtRequest string                       `json:"due_at" validate:"required"`
	Items        []CreateInvoiceDetailRequest `json:"items" validate:"required"`
	DueAt        time.Time                    `json:"-"`
}

type CreateInvoiceDetailRequest struct {
	Quantity int    `json:"quantity" validate:"required"`
	Price    int    `json:"price" validate:"required"`
	Product  string `json:"product" validate:"required"`
}
