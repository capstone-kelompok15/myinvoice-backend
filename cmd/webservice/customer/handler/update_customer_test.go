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

func (s customerHandlerSuite) TestUpdateCustomer() {
	tests := []struct {
		name              string
		request           dto.CustomerUpdateRequest
		contentType       string
		customerCtx       *dto.CustomerContext
		expectedCode      int
		expectedResponses dto.BaseResponse
		service           funcUpdateCustomer
	}{
		{
			name: "Valid",
			request: dto.CustomerUpdateRequest{
				FullName: "Valid",
				Address:  "Valid Street",
			},
			contentType: echo.MIMEApplicationJSON,
			customerCtx: &dto.CustomerContext{
				ID:       1,
				DeviceID: "abcdf",
				FullName: "Logged in Customer",
			},
			expectedCode: http.StatusOK,
			expectedResponses: dto.BaseResponse{
				Error: nil,
				Data:  "Update customer success!",
			},
			service: func(ctx context.Context, customerID *int, newData *dto.CustomerUpdateRequest) error {
				return nil
			},
		},
		{
			name: "Error because customer context nil",
			request: dto.CustomerUpdateRequest{
				FullName: "Valid",
				Address:  "Valid Street",
			},
			contentType:  echo.MIMEApplicationJSON,
			customerCtx:  nil,
			expectedCode: http.StatusInternalServerError,
			expectedResponses: dto.BaseResponse{
				Error: &dto.ErrorBaseResponse{
					Message: customerrors.ErrInternalServer.Error(),
				},
				Data: nil,
			},
			service: func(ctx context.Context, customerID *int, newData *dto.CustomerUpdateRequest) error {
				return nil
			},
		},
		{
			name: "Error because content type is not json",
			request: dto.CustomerUpdateRequest{
				FullName: "Valid",
				Address:  "Valid Street",
			},
			contentType: echo.MIMEApplicationForm,
			customerCtx: &dto.CustomerContext{
				ID:       1,
				DeviceID: "abcdf",
				FullName: "Logged in Customer",
			},
			expectedCode: http.StatusBadRequest,
			expectedResponses: dto.BaseResponse{
				Error: &dto.ErrorBaseResponse{
					Message: customerrors.ErrBadRequest.Error(),
					Detail: []string{
						"FullName is a required field",
						"Address is a required field",
					},
				},
				Data: nil,
			},
			service: func(ctx context.Context, customerID *int, newData *dto.CustomerUpdateRequest) error {
				return nil
			},
		},
		{
			name:        "Error because payload request is nil",
			request:     dto.CustomerUpdateRequest{},
			contentType: echo.MIMEApplicationJSON,
			customerCtx: &dto.CustomerContext{
				ID:       1,
				DeviceID: "abcdf",
				FullName: "Logged in Customer",
			},
			expectedCode: http.StatusBadRequest,
			expectedResponses: dto.BaseResponse{
				Error: &dto.ErrorBaseResponse{
					Message: customerrors.ErrBadRequest.Error(),
					Detail: []string{
						"FullName is a required field",
						"Address is a required field",
					},
				},
				Data: nil,
			},
			service: func(ctx context.Context, customerID *int, newData *dto.CustomerUpdateRequest) error {
				return nil
			},
		},
		{
			name: "Error because service return error",
			request: dto.CustomerUpdateRequest{
				FullName: "Valid",
				Address:  "Valid Street",
			},
			contentType: echo.MIMEApplicationJSON,
			customerCtx: &dto.CustomerContext{
				ID:       1,
				DeviceID: "abcdf",
				FullName: "Logged in Customer",
			},
			expectedCode: http.StatusInternalServerError,
			expectedResponses: dto.BaseResponse{
				Error: &dto.ErrorBaseResponse{
					Message: customerrors.ErrInternalServer.Error(),
				},
				Data: nil,
			},
			service: func(ctx context.Context, customerID *int, newData *dto.CustomerUpdateRequest) error {
				return customerrors.ErrInternalServer
			},
		},
		{
			name: "Error because record not found",
			request: dto.CustomerUpdateRequest{
				FullName: "Valid",
				Address:  "Valid Street",
			},
			contentType: echo.MIMEApplicationJSON,
			customerCtx: &dto.CustomerContext{
				ID:       1,
				DeviceID: "abcdf",
				FullName: "Logged in Customer",
			},
			expectedCode: http.StatusNotFound,
			expectedResponses: dto.BaseResponse{
				Error: &dto.ErrorBaseResponse{
					Message: customerrors.ErrRecordNotFound.Error(),
				},
				Data: nil,
			},
			service: func(ctx context.Context, customerID *int, newData *dto.CustomerUpdateRequest) error {
				return customerrors.ErrRecordNotFound
			},
		},
	}

	for _, test := range tests {
		requestJSON, err := json.Marshal(test.request)
		if err != nil {
			log.Fatal(err.Error())
		}

		expectedResponse, err := json.Marshal(test.expectedResponses)
		if err != nil {
			log.Fatal(err.Error())
		}

		res := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, apiversioning.APIVersionOne+"/customers", strings.NewReader(string(requestJSON)))
		req.Header.Set(echo.HeaderContentType, test.contentType)

		ctx := s.e.NewContext(req, res)
		ctx.Set(dto.CustomerCTXKey, test.customerCtx)

		s.handler.service = mockCustomerService{
			funcUpdateCustomer: test.service,
		}

		s.Suite.T().Run(test.name, func(t *testing.T) {
			if assert.NoError(t, s.handler.UpdateCustomer()(ctx)) {
				assert.Equal(t, string(expectedResponse), res.Body.String()[:len(res.Body.String())-1])
				assert.Equal(t, test.expectedCode, res.Code)
			}
		})
	}
}
