package router

import (
	"github.com/capstone-kelompok15/myinvoice-backend/cmd/webservice/apiversioning"
	"github.com/capstone-kelompok15/myinvoice-backend/cmd/webservice/merchant/handler"
	"github.com/capstone-kelompok15/myinvoice-backend/internal/merchant/service"
	custommiddleware "github.com/capstone-kelompok15/myinvoice-backend/pkg/middleware/service"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/validatorutils"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type MerchantRouterParams struct {
	E          *echo.Echo
	Log        *logrus.Entry
	Validator  *validatorutils.Validator
	Service    service.MerchantService
	Middleware custommiddleware.Middleware
}

func InitMerchantRouter(params *MerchantRouterParams) {
	merchantHandler := handler.NewMerchantHandler(&handler.MerchantHandler{
		Service:   params.Service,
		Log:       params.Log,
		Validator: params.Validator,
	})

	merchantV1Group := params.E.Group(apiversioning.APIVersionOne + "/merchants")
	merchantV1Group.GET("/dashboard", merchantHandler.GetDashboard(), params.Middleware.AdminMustAuthorized())
	merchantV1Group.GET("/:merchant_id/banks", merchantHandler.GetMerchantBank())
	merchantV1Group.PUT("/banks/:merchant_bank_id", merchantHandler.UpdateMerchantBank(), params.Middleware.AdminMustAuthorized())
}
