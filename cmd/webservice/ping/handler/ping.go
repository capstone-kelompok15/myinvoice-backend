package handler

import (
	"github.com/capstone-kelompok15/myinvoice-backend/internal/ping/service"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/httputils"
	"github.com/labstack/echo/v4"
)

func PingHandler(service service.PingService) echo.HandlerFunc {
	return func(c echo.Context) error {
		pingResponse := service.Ping()
		return httputils.WriteResponse(c, httputils.SuccessResponseParams{
			Data: pingResponse,
		})
	}
}
