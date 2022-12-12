package dto

type NotificationRequest struct {
	Page  uint64 `query:"page" validate:"required"`
	Limit uint64 `query:"limit" validate:"required"`
}

type CustomerNotificationDB struct {
	ID           int    `db:"id"`
	CustomerID   int    `db:"customer_id"`
	MerchantID   int    `db:"merchant_id"`
	InvoiceID    int    `db:"invoice_id"`
	Title        string `db:"title"`
	Type         string `db:"type"`
	IsRead       bool   `db:"is_read"`
	MerchantName string `db:"merchant_name"`
	CreatedAt    string `db:"created_at"`
}
type MerchantNotificationDB struct {
	ID           int    `db:"id"`
	CustomerID   int    `db:"customer_id"`
	MerchantID   int    `db:"merchant_id"`
	InvoiceID    int    `db:"invoice_id"`
	Title        string `db:"title"`
	Type         string `db:"type"`
	IsRead       bool   `db:"is_read"`
	CustomerName string `db:"customer_name"`
	CreatedAt    string `db:"created_at"`
}

type CreateNotification struct {
	CustomerID          int `json:"customer_id"`
	MerchantID          int `json:"merchant_id"`
	InvoiceID           int `json:"invoice_id"`
	NotificationTitleID int `json:"notification_title_id"`
}

type NotificationRespond struct {
	ID               int    `json:"id"`
	InvoiceID        int    `json:"invoice_id"`
	NotificationType string `json:"notification_type"`
	Title            string `json:"title"`
	Content          string `json:"content"`
	IsRead           bool   `json:"is_read"`
	CreatedAt        string `json:"created_at"`
}
