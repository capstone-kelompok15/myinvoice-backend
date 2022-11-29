package config

import "github.com/mailgun/mailgun-go/v4"

type Mailgun struct {
	PrivateApiKey string
	PublicApiKey  string
	Domain        string
	SenderEmail   string
}

func InitMailgun(params *Mailgun) *mailgun.MailgunImpl {
	client := mailgun.NewMailgun(params.Domain, params.PrivateApiKey)
	return client
}
