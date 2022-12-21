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

func (s authHandlerSuite) TestAdminLogin() {
	tests := []struct {
		name              string
		req               dto.AdminLoginRequest
		contentType       string
		expectedCode      int
		expectedResponses dto.BaseResponse
		service           funcLoginAdmin
	}{
		{
			name: "Valid",
			req: dto.AdminLoginRequest{
				Email:    "valid",
				Password: "valid123",
			},
			contentType:  echo.MIMEApplicationJSON,
			expectedCode: http.StatusOK,
			expectedResponses: dto.BaseResponse{
				Error: nil,
				Data: dto.AdminLoginResponse{
					AccessToken:  "123.123.123",
					RefreshToken: "123",
				},
			},
			service: func(ctx context.Context, req *dto.AdminLoginRequest) (*dto.AdminLoginResponse, error) {
				return &dto.AdminLoginResponse{
					AccessToken:  "123.123.123",
					RefreshToken: "123",
				}, nil
			},
		},
		{
			name: "Error because content type is not json",
			req: dto.AdminLoginRequest{
				Email:    "valid",
				Password: "valid123",
			},
			expectedCode: http.StatusBadRequest,
			expectedResponses: dto.BaseResponse{
				Error: &dto.ErrorBaseResponse{
					Message: customerrors.ErrBadRequest.Error(),
				},
				Data: nil,
			},
			service: func(ctx context.Context, req *dto.AdminLoginRequest) (*dto.AdminLoginResponse, error) {
				return nil, nil
			},
		},
		{
			name:         "Error because some require request field is empty",
			req:          dto.AdminLoginRequest{},
			contentType:  echo.MIMEApplicationJSON,
			expectedCode: http.StatusBadRequest,
			expectedResponses: dto.BaseResponse{
				Error: &dto.ErrorBaseResponse{
					Message: customerrors.ErrBadRequest.Error(),
					Detail: []string{
						"Email is a required field",
						"Password is a required field",
					},
				},
				Data: nil,
			},
			service: func(ctx context.Context, req *dto.AdminLoginRequest) (*dto.AdminLoginResponse, error) {
				return nil, nil
			},
		},
		{
			name: "Error because service return error",
			req: dto.AdminLoginRequest{
				Email:    "valid",
				Password: "valid123",
			},
			contentType:  echo.MIMEApplicationJSON,
			expectedCode: http.StatusInternalServerError,
			expectedResponses: dto.BaseResponse{
				Error: &dto.ErrorBaseResponse{
					Message: customerrors.ErrInternalServer.Error(),
				},
				Data: nil,
			},
			service: func(ctx context.Context, req *dto.AdminLoginRequest) (*dto.AdminLoginResponse, error) {
				return nil, customerrors.ErrInternalServer
			},
		},
		{
			name: "Error because service return unauthorized error",
			req: dto.AdminLoginRequest{
				Email:    "valid",
				Password: "valid123",
			},
			contentType:  echo.MIMEApplicationJSON,
			expectedCode: http.StatusUnauthorized,
			expectedResponses: dto.BaseResponse{
				Error: &dto.ErrorBaseResponse{
					Message: customerrors.ErrUnauthorized.Error(),
				},
				Data: nil,
			},
			service: func(ctx context.Context, req *dto.AdminLoginRequest) (*dto.AdminLoginResponse, error) {
				return nil, customerrors.ErrUnauthorized
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
			funcLoginAdmin: test.service,
		}

		s.Suite.T().Run(test.name, func(t *testing.T) {
			if assert.NoError(t, s.handler.AdminLogin()(ctx)) {
				assert.Equal(t, string(expectedResponse), res.Body.String()[:len(res.Body.String())-1])
				assert.Equal(t, test.expectedCode, res.Code)
			}
		})
	}
}
