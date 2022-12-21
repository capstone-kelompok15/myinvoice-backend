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

func (s authHandlerSuite) TestCustomerEmailVerification() {
	tests := []struct {
		name              string
		req               dto.CustomerEmailVerification
		contentType       string
		expectedCode      int
		expectedResponses dto.BaseResponse
		service           funcCustomerEmailVerification
	}{
		{
			name: "Valid",
			req: dto.CustomerEmailVerification{
				Email: "valid@gmail.com",
				Code:  "abcd",
			},
			contentType:  echo.MIMEApplicationJSON,
			expectedCode: http.StatusOK,
			expectedResponses: dto.BaseResponse{
				Error: nil,
				Data:  "Email Verified",
			},
			service: func(ctx context.Context, req *dto.CustomerEmailVerification) error {
				return nil
			},
		},
		{
			name: "Error because content type is not json",
			req: dto.CustomerEmailVerification{
				Email: "valid@gmail.com",
				Code:  "abcd",
			},
			expectedCode: http.StatusBadRequest,
			expectedResponses: dto.BaseResponse{
				Error: &dto.ErrorBaseResponse{
					Message: customerrors.ErrBadRequest.Error(),
				},
				Data: nil,
			},
			service: func(ctx context.Context, req *dto.CustomerEmailVerification) error {
				return nil
			},
		},
		{
			name:         "Error because some required field is not filled",
			req:          dto.CustomerEmailVerification{},
			contentType:  echo.MIMEApplicationJSON,
			expectedCode: http.StatusBadRequest,
			expectedResponses: dto.BaseResponse{
				Error: &dto.ErrorBaseResponse{
					Message: customerrors.ErrBadRequest.Error(),
					Detail: []string{
						"Email is a required field",
						"Code is a required field",
					},
				},
				Data: nil,
			},
			service: func(ctx context.Context, req *dto.CustomerEmailVerification) error {
				return nil
			},
		},
		{
			name: "Error because service return unauthorized error",
			req: dto.CustomerEmailVerification{
				Email: "valid@gmail.com",
				Code:  "abcd",
			},
			contentType:  echo.MIMEApplicationJSON,
			expectedCode: http.StatusUnauthorized,
			expectedResponses: dto.BaseResponse{
				Error: &dto.ErrorBaseResponse{
					Message: customerrors.ErrUnauthorized.Error(),
				},
				Data: nil,
			},
			service: func(ctx context.Context, req *dto.CustomerEmailVerification) error {
				return customerrors.ErrUnauthorized
			},
		},
		{
			name: "Error because service return unauthorized error",
			req: dto.CustomerEmailVerification{
				Email: "valid@gmail.com",
				Code:  "abcd",
			},
			contentType:  echo.MIMEApplicationJSON,
			expectedCode: http.StatusInternalServerError,
			expectedResponses: dto.BaseResponse{
				Error: &dto.ErrorBaseResponse{
					Message: customerrors.ErrInternalServer.Error(),
				},
				Data: nil,
			},
			service: func(ctx context.Context, req *dto.CustomerEmailVerification) error {
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
			funcCustomerEmailVerification: test.service,
		}

		s.Suite.T().Run(test.name, func(t *testing.T) {
			if assert.NoError(t, s.handler.CustomerEmailVerification()(ctx)) {
				assert.Equal(t, string(expectedResponse), res.Body.String()[:len(res.Body.String())-1])
				assert.Equal(t, test.expectedCode, res.Code)
			}
		})
	}
}
