package config

type JWTConfig struct {
	JWTSecretKey string `validate:"required"`
}
