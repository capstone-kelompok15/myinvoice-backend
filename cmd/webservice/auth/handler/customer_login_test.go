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

func (s authHandlerSuite) TestCustomerLogin() {
	tests := []struct {
		name              string
		req               dto.CustomerLoginRequest
		contentType       string
		expectedCode      int
		expectedResponses dto.BaseResponse
		service           funcCustomerLogin
	}{
		{
			name: "Valid",
			req: dto.CustomerLoginRequest{
				Email:    "valid@gmail.com",
				Password: "valid123",
				DeviceID: "valid",
			},
			contentType:  echo.MIMEApplicationJSON,
			expectedCode: http.StatusOK,
			expectedResponses: dto.BaseResponse{
				Error: nil,
				Data: dto.CustomerAccessToken{
					AccessToken: "123.123.123",
				},
			},
			service: func(ctx context.Context, req *dto.CustomerLoginRequest) (*dto.CustomerAccessToken, error) {
				return &dto.CustomerAccessToken{
					AccessToken: "123.123.123",
				}, nil
			},
		},
		{
			name: "Error because content type is not json",
			req: dto.CustomerLoginRequest{
				Email:    "valid@gmail.com",
				Password: "valid123",
				DeviceID: "valid",
			},
			expectedCode: http.StatusBadRequest,
			expectedResponses: dto.BaseResponse{
				Error: &dto.ErrorBaseResponse{
					Message: customerrors.ErrBadRequest.Error(),
				},
				Data: nil,
			},
			service: func(ctx context.Context, req *dto.CustomerLoginRequest) (*dto.CustomerAccessToken, error) {
				return &dto.CustomerAccessToken{
					AccessToken: "123.123.123",
				}, nil
			},
		},
		{
			name:         "Error because some required field is not filled",
			req:          dto.CustomerLoginRequest{},
			contentType:  echo.MIMEApplicationJSON,
			expectedCode: http.StatusBadRequest,
			expectedResponses: dto.BaseResponse{
				Error: &dto.ErrorBaseResponse{
					Message: customerrors.ErrBadRequest.Error(),
					Detail: []string{
						"Email is a required field",
						"Password is a required field",
						"DeviceID is a required field",
					},
				},
				Data: nil,
			},
			service: func(ctx context.Context, req *dto.CustomerLoginRequest) (*dto.CustomerAccessToken, error) {
				return &dto.CustomerAccessToken{
					AccessToken: "123.123.123",
				}, nil
			},
		},
		{
			name: "Error because email or password is invalid",
			req: dto.CustomerLoginRequest{
				Email:    "valid@gmail.com",
				Password: "valid123",
				DeviceID: "valid",
			},
			contentType:  echo.MIMEApplicationJSON,
			expectedCode: http.StatusUnauthorized,
			expectedResponses: dto.BaseResponse{
				Error: &dto.ErrorBaseResponse{
					Message: customerrors.ErrEmailPasswordIncorrect.Error(),
				},
				Data: nil,
			},
			service: func(ctx context.Context, req *dto.CustomerLoginRequest) (*dto.CustomerAccessToken, error) {
				return nil, customerrors.ErrEmailPasswordIncorrect
			},
		},
		{
			name: "Error because service return internal server error",
			req: dto.CustomerLoginRequest{
				Email:    "valid@gmail.com",
				Password: "valid123",
				DeviceID: "valid",
			},
			contentType:  echo.MIMEApplicationJSON,
			expectedCode: http.StatusInternalServerError,
			expectedResponses: dto.BaseResponse{
				Error: &dto.ErrorBaseResponse{
					Message: customerrors.ErrInternalServer.Error(),
				},
				Data: nil,
			},
			service: func(ctx context.Context, req *dto.CustomerLoginRequest) (*dto.CustomerAccessToken, error) {
				return nil, customerrors.ErrInternalServer
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
			funcCustomerLogin: test.service,
		}

		s.Suite.T().Run(test.name, func(t *testing.T) {
			if assert.NoError(t, s.handler.CustomerLogin()(ctx)) {
				assert.Equal(t, string(expectedResponse), res.Body.String()[:len(res.Body.String())-1])
				assert.Equal(t, test.expectedCode, res.Code)
			}
		})
	}
}
