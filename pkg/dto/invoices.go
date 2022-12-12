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

type GetAllInvoicesParam struct {
	// General Merchant
	*DateFilter
	*PaginationFilter
	PaymentStatusID int

	// Merchant
	MerchantID               int
	MerchantFilterCustomerID int `query:"customer_id"`

	// Customer
	CustomerID int
}

type GetInvoiceResponse struct {
	InvoiceID int `json:"invoice_id" db:"invoice_id"`
	// MerchantID        int                `json:"merchant_id" db:"merchant_id"`
	// CustomerID        int     `json:"customer_id" db:"customer_id"`
	CustomerName      string  `json:"customer_name" db:"customer_name"`
	PaymentStatusID   int     `json:"payment_status_id" db:"payment_status_id"`
	PaymentStatusName string  `json:"payment_status_name" db:"payment_status_name"`
	PaymentTypeID     *int    `json:"payment_type_id" db:"payment_type_id"`
	PaymentTypeName   *string `json:"payment_type_name" db:"payment_type_name"`
	// MerchantBankID    *int               `json:"merchant_bank_id" db:"merchant_bank_id"`
	TotalPrice int64 `json:"total_price" db:"total_price"`
	// ProductQuantity   int                `json:"product_quantity" db:"product_quantity"`
	DueAt     string `json:"due_at" db:"due_at"`
	CreatedAt string `json:"created_at" db:"created_at"`
	UpdatedAt string `json:"updated_at" db:"updated_at"`
	// InvoiceQuantity int    `json:"invoice_quantity" db:"invoice_quantity"`
	// InvoiceDetail     []GetDetailInvoice `json:"invoice_detail"`
}

type GetDetailInvoice struct {
	InvoiceDetailID int    `json:"invoice_detail_id" db:"invoice_detail_id"`
	Product         string `json:"product" db:"product"`
	Quantity        int    `json:"quantity" db:"quantity"`
	Price           int64  `json:"price" db:"price"`
	CreatedAt       string `json:"created_at" db:"created_at"`
	UpdatedAt       string `json:"updated_at" db:"updated_at"`
}
