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

type OverviewMerchantDashboard struct {
	PaymentReceivedQuantity int64 `json:"payment_received_quantity"`
	CustomerQuantity        int   `json:"customer_quantity"`
	InvoiceQuantity         int   `json:"invoice_quantity"`
	UnpaidInvoiceQuantity   int   `json:"unpaid_invoice_quantity"`
}

type RecentInvoiceMerchantDashboard struct {
	Price              int64  `json:"price" db:"price"`
	InvoiceID          int64  `json:"invoice_id" db:"invoice_id"`
	CustomerName       string `json:"customer_name" db:"customer_name"`
	InvoiceExpiredDate string `json:"invoice_expired_date" db:"invoice_expired_date"`
}

type RecentPaymentMerchantDashboard struct {
	RecentInvoiceMerchantDashboard
	PaymentType string `json:"payment_type" db:"payment_type"`
}

type MerchantDashboard struct {
	OverviewMerchantDashboard      `json:"overview"`
	RecentInvoiceMerchantDashboard []RecentInvoiceMerchantDashboard `json:"recent_invoice"`
	RecentPaymentMerchantDashboard []RecentPaymentMerchantDashboard `json:"recent_payment"`
}

type AdminResetPassword struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required,min=8"`
	Code     string `json:"code" validate:"required"`
}

type GetMerchantBankRequest struct {
	MerchantID int `param:"merchant_id"`
}

type GetMerchantBankResponse struct {
	BankName   string `json:"bank_name" db:"bank_name"`
	BankCode   string `json:"bank_code" db:"bank_code"`
	OnBehalfOf string `json:"on_behalf_of" db:"on_behalf_of"`
	BankNumber string `json:"bank_number" db:"bank_number"`
}

type MerchantBankData struct {
	BankID     int    `json:"bank_id" validate:"required"`
	OnBehalfOf string `json:"on_behalf_of" validate:"required"`
	BankNumber string `json:"bank_number" validate:"required"`
}

type UpdateMerchantBankDataRequest struct {
	MerchantBankData
	MerchantBankID int `param:"merchant_bank_id" validate:"required"`
	MerchantID     int
}
