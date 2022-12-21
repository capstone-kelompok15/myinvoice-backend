package handler

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/capstone-kelompok15/myinvoice-backend/cmd/webservice/apiversioning"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
	customerrors "github.com/capstone-kelompok15/myinvoice-backend/pkg/errors"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func (s authHandlerSuite) TestRefreshAdminEmailVerificationCode() {
	tests := []struct {
		name              string
		req               dto.AdminRefreshEmailVerificationCode
		contentType       string
		expectedCode      int
		expectedResponses dto.BaseResponse
		service           funcRefreshAdminEmailVerificationCode
	}{
		{
			name: "Valid",
			req: dto.AdminRefreshEmailVerificationCode{
				Email: "valid@gmail.com",
			},
			contentType:  echo.MIMEApplicationJSON,
			expectedCode: http.StatusCreated,
			expectedResponses: dto.BaseResponse{
				Error: nil,
				Data:  "New code sent!",
			},
			service: func(ctx context.Context, email string) error {
				return nil
			},
		},
		{
			name: "Error because content type is not json",
			req: dto.AdminRefreshEmailVerificationCode{
				Email: "valid@gmail.com",
			},
			expectedCode: http.StatusBadRequest,
			expectedResponses: dto.BaseResponse{
				Error: &dto.ErrorBaseResponse{
					Message: customerrors.ErrBadRequest.Error(),
				},
				Data: nil,
			},
			service: func(ctx context.Context, email string) error {
				return nil
			},
		},
		{
			name:         "Error because some required field is empty",
			req:          dto.AdminRefreshEmailVerificationCode{},
			contentType:  echo.MIMEApplicationJSON,
			expectedCode: http.StatusBadRequest,
			expectedResponses: dto.BaseResponse{
				Error: &dto.ErrorBaseResponse{
					Message: customerrors.ErrBadRequest.Error(),
					Detail: []string{
						"Email is a required field",
					},
				},
				Data: nil,
			},
			service: func(ctx context.Context, email string) error {
				return nil
			},
		},
		{
			name: "Error because service return error",
			req: dto.AdminRefreshEmailVerificationCode{
				Email: "valid@gmail.com",
			},
			contentType:  echo.MIMEApplicationJSON,
			expectedCode: http.StatusInternalServerError,
			expectedResponses: dto.BaseResponse{
				Error: &dto.ErrorBaseResponse{
					Message: customerrors.ErrInternalServer.Error(),
				},
				Data: nil,
			},
			service: func(ctx context.Context, email string) error {
				return customerrors.ErrInternalServer
			},
		},
		{
			name: "Error because service return unauthorized error",
			req: dto.AdminRefreshEmailVerificationCode{
				Email: "valid@gmail.com",
			},
			contentType:  echo.MIMEApplicationJSON,
			expectedCode: http.StatusNotFound,
			expectedResponses: dto.BaseResponse{
				Error: &dto.ErrorBaseResponse{
					Message: customerrors.ErrRecordNotFound.Error(),
				},
				Data: nil,
			},
			service: func(ctx context.Context, email string) error {
				return sql.ErrNoRows
			},
		},
	}

	for _, test := range tests {
		request, err := json.Marshal(test.req)
		if err != nil {
			log.Fatal(err.Error())
		}

		expectedResponse, err := json.Marshal(test.expectedResponses)
		if err != nil {
			log.Fatal(err.Error())
		}

		res := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, apiversioning.APIVersionOne+"/auth", strings.NewReader(string(request)))
		req.Header.Add(echo.HeaderContentType, test.contentType)

		ctx := s.e.NewContext(req, res)
		s.handler.service = mockAuthService{
			funcRefreshAdminEmailVerificationCode: test.service,
		}

		s.Suite.T().Run(test.name, func(t *testing.T) {
			if assert.NoError(t, s.handler.RefreshAdminEmailVerificationCode()(ctx)) {
				assert.Equal(t, string(expectedResponse), res.Body.String()[:len(res.Body.String())-1])
				assert.Equal(t, test.expectedCode, res.Code)
			}
		})
	}
}
