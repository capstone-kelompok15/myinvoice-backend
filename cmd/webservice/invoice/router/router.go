package router

import (
	"github.com/capstone-kelompok15/myinvoice-backend/cmd/webservice/apiversioning"
	"github.com/capstone-kelompok15/myinvoice-backend/cmd/webservice/invoice/handler"
	"github.com/capstone-kelompok15/myinvoice-backend/internal/invoice/service"
	custommiddleware "github.com/capstone-kelompok15/myinvoice-backend/pkg/middleware/service"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/validatorutils"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type InvoiceRouterParams struct {
	E          *echo.Echo
	Log        *logrus.Entry
	Validator  *validatorutils.Validator
	Service    service.InvoiceService
	Middleware custommiddleware.Middleware
}

func InitInvoiceRouter(params *InvoiceRouterParams) {
	invoiceHandler := handler.NewInvoiceHandler(&handler.InvoiceHandlerParams{
		Service:   params.Service,
		Log:       params.Log,
		Validator: params.Validator,
	})

	invoiceV1Group := params.E.Group(apiversioning.APIVersionOne + "/invoices")
	invoiceV1Group.POST("", invoiceHandler.CreateInvoice(), params.Middleware.AdminMustAuthorized())
	invoiceV1Group.GET("/merchants", invoiceHandler.MerchantGetAllInvoices(), params.Middleware.AdminMustAuthorized())
	invoiceV1Group.GET("/customers", invoiceHandler.CustomerGetAllInvoices(), params.Middleware.CustomerMustAuthorized())
	invoiceV1Group.GET("/:invoice_id/merchants", invoiceHandler.MerchantGetDetailInvoiceByID(), params.Middleware.AdminMustAuthorized())
	invoiceV1Group.GET("/:invoice_id/customers", invoiceHandler.CustomerGetDetailInvoiceByID(), params.Middleware.CustomerMustAuthorized())
	invoiceV1Group.GET("/merchants/customers", invoiceHandler.GetCustomers(), params.Middleware.AdminMustAuthorized())
	invoiceV1Group.GET("/:invoice_id/payments/upload", invoiceHandler.UploadPayment(), params.Middleware.CustomerMustAuthorized())
}
