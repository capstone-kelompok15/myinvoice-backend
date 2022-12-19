package authutils

import (
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"

	"github.com/labstack/echo/v4"
)

func AdminContextFromRequestContext(ec echo.Context) *dto.AdminContext {
	extractedAccount, ok := ec.Get(dto.AdminCTXKey).(*dto.AdminContext)
	if !ok {
		return nil
	}

	return extractedAccount
}

func CustomerFromRequestContext(ec echo.Context) *dto.CustomerContext {
	extractedAccount, ok := ec.Get(dto.CustomerCTXKey).(*dto.CustomerContext)
	if !ok {
		return nil
	}

	return extractedAccount
}
