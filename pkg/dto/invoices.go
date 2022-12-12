package dto

import "time"

type CreateInvoiceRequest struct {
	CustomerID     int                          `json:"customer_id"`
	PaymentType    int                          `json:"payment_type"`
	MerchantBank   int                          `json:"merchant_bank"`
	DueAtRequest   string                       `json:"due_at"`
	InvoiceDetails []CreateInvoiceDetailRequest `json:"invoice_details"`
	DueAt          time.Time                    `json:"-"`
}

type CreateInvoiceDetailRequest struct {
	Quantity int    `json:"quantity"`
	Price    int    `json:"price"`
	Product  string `json:"product"`
}
