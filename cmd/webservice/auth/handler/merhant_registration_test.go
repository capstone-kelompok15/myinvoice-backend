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

func (s authHandlerSuite) TestRegisterMerchant() {
	req := dto.MerchantRegisterRequest{
		Email:               "valid@gmail.com",
		Password:            "valid123",
		Username:            "valid",
		MerchantName:        "valid",
		MerchantAddress:     "valid",
		MerchantPhoneNumber: "valid",
		MerchantBank: []dto.MerchantBankRegisterRequest{
			{
				BankID:     1,
				OnBehalfOf: "valid",
				BankNumber: "valid",
			},
		},
	}

	tests := []struct {
		name              string
		req               dto.MerchantRegisterRequest
		contentType       string
		expectedCode      int
		expectedResponses dto.BaseResponse
		service           funcMerchantRegistration
	}{
		{
			name:         "Valid",
			req:          req,
			contentType:  echo.MIMEApplicationJSON,
			expectedCode: http.StatusCreated,
			expectedResponses: dto.BaseResponse{
				Error: nil,
				Data:  "Merchant Created!",
			},
			service: func(ctx context.Context, req *dto.MerchantRegisterRequest) error {
				return nil
			},
		},
		{
			name:         "Error because content type is not json",
			req:          req,
			expectedCode: http.StatusBadRequest,
			expectedResponses: dto.BaseResponse{
				Error: &dto.ErrorBaseResponse{
					Message: customerrors.ErrBadRequest.Error(),
				},
				Data: nil,
			},
			service: func(ctx context.Context, req *dto.MerchantRegisterRequest) error {
				return nil
			},
		},
		{
			name:         "Error because some required field is not filled",
			req:          dto.MerchantRegisterRequest{},
			contentType:  echo.MIMEApplicationJSON,
			expectedCode: http.StatusBadRequest,
			expectedResponses: dto.BaseResponse{
				Error: &dto.ErrorBaseResponse{
					Message: customerrors.ErrBadRequest.Error(),
					Detail: []string{
						"Email is a required field",
						"Password is a required field",
						"Username is a required field",
						"MerchantName is a required field",
						"MerchantAddress is a required field",
						"MerchantBank is a required field",
					},
				},
				Data: nil,
			},
			service: func(ctx context.Context, req *dto.MerchantRegisterRequest) error {
				return nil
			},
		},
		{
			name: "Error because some required field is not filled",
			req: dto.MerchantRegisterRequest{
				Email:               "valid@gmail.com",
				Password:            "valid123",
				Username:            "valid",
				MerchantName:        "valid",
				MerchantAddress:     "valid",
				MerchantPhoneNumber: "valid",
				MerchantBank: []dto.MerchantBankRegisterRequest{
					{
						BankID:     0,
						OnBehalfOf: "",
						BankNumber: "",
					},
				},
			},
			contentType:  echo.MIMEApplicationJSON,
			expectedCode: http.StatusBadRequest,
			expectedResponses: dto.BaseResponse{
				Error: &dto.ErrorBaseResponse{
					Message: customerrors.ErrBadRequest.Error(),
					Detail: []map[string]interface{}{
						{
							"bank_index": 0,
							"error_detail": []string{
								"BankID is a required field",
								"OnBehalfOf is a required field",
								"BankNumber is a required field",
							},
						},
					},
				},
				Data: nil,
			},
			service: func(ctx context.Context, req *dto.MerchantRegisterRequest) error {
				return nil
			},
		},
		{
			name:         "Error because service return error",
			req:          req,
			contentType:  echo.MIMEApplicationJSON,
			expectedCode: http.StatusInternalServerError,
			expectedResponses: dto.BaseResponse{
				Error: &dto.ErrorBaseResponse{
					Message: customerrors.ErrInternalServer.Error(),
				},
				Data: nil,
			},
			service: func(ctx context.Context, req *dto.MerchantRegisterRequest) error {
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
			funcMerchantRegistration: test.service,
		}

		s.Suite.T().Run(test.name, func(t *testing.T) {
			if assert.NoError(t, s.handler.RegisterMerchant()(ctx)) {
				assert.Equal(t, string(expectedResponse), res.Body.String()[:len(res.Body.String())-1])
				assert.Equal(t, test.expectedCode, res.Code)
			}
		})
	}
}
