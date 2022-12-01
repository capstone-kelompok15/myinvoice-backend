package config

type CustomerToken struct {
	SecretKey string `validate:"required,min=64"`
}
