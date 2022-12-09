package router

import (
	"github.com/capstone-kelompok15/myinvoice-backend/cmd/webservice/apiversioning"
	"github.com/capstone-kelompok15/myinvoice-backend/cmd/webservice/bank/handler"
	"github.com/capstone-kelompok15/myinvoice-backend/internal/bank/service"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/validatorutils"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type BankRouterParams struct {
	E         *echo.Echo
	Log       *logrus.Entry
	Validator *validatorutils.Validator
	Service   service.BankService
}

func InitBankRouter(params *BankRouterParams) {
	bankHandler := handler.NewBankHandler(&handler.BankHandler{
		Service:   params.Service,
		Log:       params.Log,
		Validator: params.Validator,
	})

	bankV1Group := params.E.Group(apiversioning.APIVersionOne + "/banks")
	bankV1Group.GET("", bankHandler.GetAllBank())

}
