package handler

import (
	"context"
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

func (s authHandlerSuite) TestCustomerResetPassword() {
	tests := []struct {
		name              string
		req               dto.CustomerResetPassword
		contentType       string
		expectedCode      int
		expectedResponses dto.BaseResponse
		service           funcResetPassword
	}{
		{
			name: "Valid",
			req: dto.CustomerResetPassword{
				Email:    "valid@gmail.com",
				Password: "valid123",
				Code:     "valid code",
			},
			contentType:  echo.MIMEApplicationJSON,
			expectedCode: http.StatusOK,
			expectedResponses: dto.BaseResponse{
				Error: nil,
				Data:  "Password was reset",
			},
			service: func(ctx context.Context, req *dto.CustomerResetPassword) error {
				return nil
			},
		},
		{
			name: "Content type is not json",
			req: dto.CustomerResetPassword{
				Email:    "valid@gmail.com",
				Password: "valid123",
				Code:     "valid code",
			},
			expectedCode: http.StatusBadRequest,
			expectedResponses: dto.BaseResponse{
				Error: &dto.ErrorBaseResponse{
					Message: customerrors.ErrBadRequest.Error(),
				},
				Data: nil,
			},
			service: func(ctx context.Context, req *dto.CustomerResetPassword) error {
				return nil
			},
		},
		{
			name:         "Error because some required field is not filled",
			req:          dto.CustomerResetPassword{},
			contentType:  echo.MIMEApplicationJSON,
			expectedCode: http.StatusBadRequest,
			expectedResponses: dto.BaseResponse{
				Error: &dto.ErrorBaseResponse{
					Message: customerrors.ErrBadRequest.Error(),
					Detail: []string{
						"Email is a required field",
						"Password is a required field",
						"Code is a required field",
					},
				},
				Data: nil,
			},
			service: func(ctx context.Context, req *dto.CustomerResetPassword) error {
				return nil
			},
		},
		{
			name: "Error because service return not found error",
			req: dto.CustomerResetPassword{
				Email:    "valid@gmail.com",
				Password: "valid123",
				Code:     "valid code",
			},
			contentType:  echo.MIMEApplicationJSON,
			expectedCode: http.StatusNotFound,
			expectedResponses: dto.BaseResponse{
				Error: &dto.ErrorBaseResponse{
					Message: customerrors.ErrNotFound.Error(),
				},
				Data: nil,
			},
			service: func(ctx context.Context, req *dto.CustomerResetPassword) error {
				return customerrors.ErrNotFound
			},
		},
		{
			name: "Error because service return internal server error",
			req: dto.CustomerResetPassword{
				Email:    "valid@gmail.com",
				Password: "valid123",
				Code:     "valid code",
			},
			contentType:  echo.MIMEApplicationJSON,
			expectedCode: http.StatusInternalServerError,
			expectedResponses: dto.BaseResponse{
				Error: &dto.ErrorBaseResponse{
					Message: customerrors.ErrInternalServer.Error(),
				},
				Data: nil,
			},
			service: func(ctx context.Context, req *dto.CustomerResetPassword) error {
				return customerrors.ErrInternalServer
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
			funcResetPassword: test.service,
		}

		s.Suite.T().Run(test.name, func(t *testing.T) {
			if assert.NoError(t, s.handler.CustomerResetPassword()(ctx)) {
				assert.Equal(t, string(expectedResponse), res.Body.String()[:len(res.Body.String())-1])
				assert.Equal(t, test.expectedCode, res.Code)
			}
		})
	}
}
