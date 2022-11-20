package router

import (
	"github.com/capstone-kelompok15/myinvoice-backend/cmd/webservice/apiversioning"
	"github.com/capstone-kelompok15/myinvoice-backend/cmd/webservice/ping/handler"
	"github.com/capstone-kelompok15/myinvoice-backend/internal/ping/service"
	"github.com/labstack/echo/v4"
)

type RouterParams struct {
	E           *echo.Echo
	PingService service.PingService
}

func InitPingRouter(params RouterParams) {
	pingV1Group := params.E.Group(apiversioning.APIVersionOne + "/ping")

	pingV1Group.GET("", handler.PingHandler(params.PingService))
}
