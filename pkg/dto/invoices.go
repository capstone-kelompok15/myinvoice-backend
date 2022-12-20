package dto

import "time"

type CreateInvoiceRequest struct {
	CustomerID   int                          `json:"customer_id" validate:"required"`
	DueAtRequest string                       `json:"due_at" validate:"required"`
	Note         string                       `json:"note"`
	Items        []CreateInvoiceDetailRequest `json:"items" validate:"required"`
	DueAt        time.Time                    `json:"-"`
}

type CreateInvoiceDetailRequest struct {
	Quantity int    `json:"quantity" validate:"required,min=1"`
	Price    int    `json:"price" validate:"required,min=0"`
	Product  string `json:"product" validate:"required"`
}

type GetAllInvoicesParam struct {
	// General Merchant
	*DateFilter
	*PaginationFilter
	PaymentStatusID int `query:"payment_status_id"`

	// Merchant
	MerchantID               int
	MerchantFilterCustomerID int `query:"customer_id"`

	// Customer
	CustomerID int
}

type GetDetailsInvoicesRequest struct {
	InvoiceID  int `param:"invoice_id" validate:"required"`
	CustomerID int
	MerchantID int
}

type GetInvoiceResponse struct {
	InvoiceID         int     `json:"invoice_id" db:"invoice_id"`
	MerchantID        int     `json:"merchant_id" db:"merchant_id"`
	MerchantName      string  `json:"merchant_name" db:"merchant_name"`
	CustomerName      string  `json:"customer_name" db:"customer_name"`
	PaymentStatusID   int     `json:"payment_status_id" db:"payment_status_id"`
	PaymentStatusName string  `json:"payment_status_name" db:"payment_status_name"`
	PaymentTypeID     *int    `json:"payment_type_id" db:"payment_type_id"`
	PaymentTypeName   *string `json:"payment_type_name" db:"payment_type_name"`
	TotalPrice        int64   `json:"total_price" db:"total_price"`
	DueAt             string  `json:"due_at" db:"due_at"`
	CreatedAt         string  `json:"created_at" db:"created_at"`
	UpdatedAt         string  `json:"updated_at" db:"updated_at"`
}

type GetInvoiceDetailsByIDResponse struct {
	InvoiceID           int                `json:"invoice_id" db:"invoice_id"`
	MerchantID          int                `json:"merchant_id" db:"merchant_id"`
	MerchantName        string             `json:"merchant_name" db:"merchant_name"`
	MerchantAddress     *string            `json:"merchant_address" db:"merchant_address"`
	CustomerID          int                `json:"customer_id" db:"customer_id"`
	CustomerName        string             `json:"customer_name" db:"customer_name"`
	CustomerEmail       string             `json:"customer_email" db:"customer_email"`
	CustomerAddress     *string            `json:"customer_address" db:"customer_address"`
	ApprovalDocumentURL *string            `json:"approval_document_url" db:"approval_document_url"`
	PaymentStatusID     int                `json:"payment_status_id" db:"payment_status_id"`
	PaymentStatusName   string             `json:"payment_status_name" db:"payment_status_name"`
	PaymentTypeID       *int               `json:"payment_type_id" db:"payment_type_id"`
	PaymentTypeName     *string            `json:"payment_type_name" db:"payment_type_name"`
	MerchantBankID      *int               `json:"merchant_bank_id" db:"merchant_bank_id"`
	TotalPrice          int64              `json:"total_price" db:"total_price"`
	ProductQuantity     int                `json:"product_quantity" db:"product_quantity"`
	Note                *string            `json:"note" db:"note"`
	Message             *string            `json:"message" db:"message"`
	DueAt               string             `json:"due_at" db:"due_at"`
	CreatedAt           string             `json:"created_at" db:"created_at"`
	UpdatedAt           string             `json:"updated_at" db:"updated_at"`
	InvoiceDetail       []GetInvoiceDetail `json:"invoice_detail"`
}

type GetInvoiceDetail struct {
	InvoiceDetailID int    `json:"invoice_detail_id" db:"invoice_detail_id"`
	Product         string `json:"product" db:"product"`
	Quantity        int    `json:"quantity" db:"quantity"`
	Price           int64  `json:"price" db:"price"`
	CreatedAt       string `json:"created_at" db:"created_at"`
	UpdatedAt       string `json:"updated_at" db:"updated_at"`
}

type GetMerchantCustomerList struct {
	*PaginationFilter
	MerchantID int
}
type GetInvoiceByID struct {
	CustomerID int `db:"customer_id"`
	MerchantID int `db:"merchant_id"`
}

type ReportDateInt struct {
	StartDate int
	EndDate   int
}

type ReportDateStr struct {
	StartDate string
	EndDate   string
}

type ReportParams struct {
	ReportDaysInt []ReportDateInt
	ReportDaysStr []ReportDateStr
	CustomerID    int

	PaymentStatus int    `query:"payment_status" validate:"required"`
	DateFilter    string `query:"date_filter" validate:"required,oneof='1 Week' '1 Month' '3 Month' '1 Year'"`
}

type ReportTransaction struct {
	TransactionQuantity int   `json:"transaction_quantity" db:"transaction_quantity"`
	TransactionTotal    int64 `json:"transaction_total" db:"transaction_total"`
}

type ReportDate struct {
	Date  string `json:"date"`
	Value int    `json:"value"`
}

type ReportResponse struct {
	Reports []ReportDate `json:"reports"`
	ReportTransaction
}

type DeleteInvoice struct {
	InvoiceID       int `param:"invoice_id"`
	InvoiceDetailID int `param:"invoice_detail_id"`
	CustomerID      int
}
