package authutils

import (
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"

	"github.com/labstack/echo/v4"
)

func UserFromRequestContext(ec echo.Context) *dto.UserContext {
	extractedAccount, ok := ec.Get(dto.AccountCTXKey).(*dto.UserContext)
	if !ok {
		return nil
	}

	return extractedAccount
}
