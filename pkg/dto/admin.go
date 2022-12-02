package dto

const (
	AccountCTXKey = "ACCOUNT-CTX-KEY"
)

type AdminLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AdminContext struct {
	ID           int    `json:"id"`
	MerchantID   int    `json:"merchant_id"`
	MerchantName string `json:"merchant_name"`
}

type AdminLoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
