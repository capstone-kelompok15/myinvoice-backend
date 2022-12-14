package websocketutils

type Message struct {
	MessageType    int    `json:"message_type"`
	Content        string `json:"content"`
	InvoiceID      int    `json:"invoice_id"`
	SendToCustomer int    `json:"send_to_customer"`
	SendToMerchant int    `json:"send_to_merchant"`
	PaymentTypeID  int    `json:"payment_type_id"`
	PaymentType    string `json:"payment_type"`
	Disconnect     bool   `json:"disconnect,omitempty"`
}

type MessageParams struct {
	Content        string
	SendToCustomer int
	SendToMerchant int
	InvoiceID      int
	PaymentTypeID  int
	PaymentType    string
}

func NewWebSocketMessage(params *MessageParams) *Message {
	return &Message{
		MessageType:    1,
		Content:        params.Content,
		InvoiceID:      params.InvoiceID,
		SendToCustomer: params.SendToCustomer,
		SendToMerchant: params.SendToMerchant,
		PaymentTypeID:  params.PaymentTypeID,
		PaymentType:    params.PaymentType,
		Disconnect:     false,
	}
}
