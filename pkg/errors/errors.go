package errors

import (
	"errors"
	"net/http"

	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
)

var (
	ErrInternalServer    = errors.New("internal server error")
	ErrBadRequest        = errors.New("bad request")
	ErrRecordNotFound    = errors.New("record not found")
	ErrNotFound          = errors.New("not found")
	ErrUnauthorized      = errors.New("unauthorized")
	ErrAuthTokenExpired  = errors.New("token expired")
	ErrAccountDuplicated = errors.New("email or username has already exists")
	ErrUniqueRecord      = errors.New("record must be unique")
)

var errMap map[error]dto.ErrorResponse = map[error]dto.ErrorResponse{
	ErrInternalServer:    {HTTPErrorCode: http.StatusInternalServerError, Message: ErrInternalServer.Error()},
	ErrBadRequest:        {HTTPErrorCode: http.StatusBadRequest, Message: ErrBadRequest.Error()},
	ErrRecordNotFound:    {HTTPErrorCode: http.StatusNotFound, Message: ErrRecordNotFound.Error()},
	ErrNotFound:          {HTTPErrorCode: http.StatusNotFound, Message: ErrNotFound.Error()},
	ErrUnauthorized:      {HTTPErrorCode: http.StatusUnauthorized, Message: ErrUnauthorized.Error()},
	ErrAuthTokenExpired:  {HTTPErrorCode: http.StatusUnauthorized, Message: ErrAuthTokenExpired.Error()},
	ErrAccountDuplicated: {HTTPErrorCode: http.StatusBadRequest, Message: ErrAccountDuplicated.Error()},
	ErrUniqueRecord:      {HTTPErrorCode: http.StatusBadRequest, Message: ErrUniqueRecord.Error()},
}

func GetErr(param error) dto.ErrorResponse {
	res, exists := errMap[param]
	if !exists {
		return errMap[ErrInternalServer]
	}
	return res
}
