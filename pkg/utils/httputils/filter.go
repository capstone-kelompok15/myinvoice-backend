package httputils

import (
	"strconv"
	"time"

	customerrors "github.com/capstone-kelompok15/myinvoice-backend/pkg/errors"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/dateutils"
	"github.com/labstack/echo/v4"
)

func GetPaginationMandatoryParams(ec echo.Context, required bool) (offset int, limit int, err error) {
	o := ec.QueryParam("offset")
	l := ec.QueryParam("limit")

	if required {
		if o == "" || l == "" {
			err = customerrors.ErrBadRequest
			return
		}
	}

	limit, err = strconv.Atoi(l)
	if err != nil {
		limit = 0
	}

	offset, err = strconv.Atoi(o)
	if err != nil {
		offset = 0
	}

	return
}

func GetDateQueryMandatoryParams(ec echo.Context) (dateStart time.Time, dateEnd time.Time) {

	start, err := dateutils.StringToDate(ec.QueryParam("start_date"))
	if err != nil {
		return time.Now(), time.Now()
	}

	end, err := dateutils.StringToDate(ec.QueryParam("end_date"))
	if err != nil {
		return *start, time.Now()
	}

	return *start, *end
}
