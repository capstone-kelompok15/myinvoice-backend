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

func (s authHandlerSuite) TestRegisterCustomer() {
	tests := []struct {
		name              string
		req               dto.CustomerRequest
		contentType       string
		expectedCode      int
		expectedResponses dto.BaseResponse
		service           funcCustomerRegistration
	}{
		{
			name: "Valid",
			req: dto.CustomerRequest{
				Email:    "valid@gmail.com",
				Password: "valid123",
				FullName: "valid",
			},
			contentType:  echo.MIMEApplicationJSON,
			expectedCode: http.StatusCreated,
			expectedResponses: dto.BaseResponse{
				Error: nil,
				Data:  "User Created!",
			},
			service: func(ctx context.Context, params *dto.CustomerRequest) error {
				return nil
			},
		},
		{
			name: "Error because content type is not json",
			req: dto.CustomerRequest{
				Email:    "valid@gmail.com",
				Password: "valid123",
				FullName: "valid",
			},
			expectedCode: http.StatusBadRequest,
			expectedResponses: dto.BaseResponse{
				Error: &dto.ErrorBaseResponse{
					Message: customerrors.ErrBadRequest.Error(),
				},
				Data: nil,
			},
			service: func(ctx context.Context, params *dto.CustomerRequest) error {
				return nil
			},
		},
		{
			name:         "Error because content type is not json",
			req:          dto.CustomerRequest{},
			contentType:  echo.MIMEApplicationJSON,
			expectedCode: http.StatusBadRequest,
			expectedResponses: dto.BaseResponse{
				Error: &dto.ErrorBaseResponse{
					Message: customerrors.ErrBadRequest.Error(),
					Detail: []string{
						"Email is a required field",
						"Password is a required field",
						"FullName is a required field",
					},
				},
				Data: nil,
			},
			service: func(ctx context.Context, params *dto.CustomerRequest) error {
				return nil
			},
		},
		{
			name: "Error because account duplicated",
			req: dto.CustomerRequest{
				Email:    "valid@gmail.com",
				Password: "valid123",
				FullName: "valid",
			},
			contentType:  echo.MIMEApplicationJSON,
			expectedCode: http.StatusBadRequest,
			expectedResponses: dto.BaseResponse{
				Error: &dto.ErrorBaseResponse{
					Message: customerrors.ErrAccountDuplicated.Error(),
				},
				Data: nil,
			},
			service: func(ctx context.Context, params *dto.CustomerRequest) error {
				return customerrors.ErrAccountDuplicated
			},
		},
		{
			name: "Error because internal server error",
			req: dto.CustomerRequest{
				Email:    "valid@gmail.com",
				Password: "valid123",
				FullName: "valid",
			},
			contentType:  echo.MIMEApplicationJSON,
			expectedCode: http.StatusInternalServerError,
			expectedResponses: dto.BaseResponse{
				Error: &dto.ErrorBaseResponse{
					Message: customerrors.ErrInternalServer.Error(),
				},
				Data: nil,
			},
			service: func(ctx context.Context, params *dto.CustomerRequest) error {
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
			funcCustomerRegistration: test.service,
		}

		s.Suite.T().Run(test.name, func(t *testing.T) {
			if assert.NoError(t, s.handler.RegisterCustomer()(ctx)) {
				assert.Equal(t, string(expectedResponse), res.Body.String()[:len(res.Body.String())-1])
				assert.Equal(t, test.expectedCode, res.Code)
			}
		})
	}
}
