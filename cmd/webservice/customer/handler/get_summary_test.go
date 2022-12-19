package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/capstone-kelompok15/myinvoice-backend/cmd/webservice/apiversioning"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
	customerrors "github.com/capstone-kelompok15/myinvoice-backend/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func (s customerHandlerSuite) TestGetSummary() {
	tests := []struct {
		name              string
		customerCtx       *dto.CustomerContext
		expectedCode      int
		expectedResponses dto.BaseResponse
		service           funcGetSummary
	}{
		{
			name: "Valid",
			customerCtx: &dto.CustomerContext{
				ID:       1,
				DeviceID: "abcdf",
				FullName: "Logged in Customer",
			},
			expectedCode: http.StatusOK,
			expectedResponses: dto.BaseResponse{
				Error: nil,
				Data: dto.CustomerSummary{
					TotalPaid:   1000,
					TotalUnpaid: 1000,
				},
			},
			service: func(ctx context.Context, customerID int) (*dto.CustomerSummary, error) {
				return &dto.CustomerSummary{
					TotalPaid:   1000,
					TotalUnpaid: 1000,
				}, nil
			},
		},
		{
			name:         "Error because customer context is nil",
			customerCtx:  nil,
			expectedCode: http.StatusInternalServerError,
			expectedResponses: dto.BaseResponse{
				Error: &dto.ErrorBaseResponse{
					Message: customerrors.ErrInternalServer.Error(),
				},
				Data: nil,
			},
			service: func(ctx context.Context, customerID int) (*dto.CustomerSummary, error) {
				return &dto.CustomerSummary{
					TotalPaid:   1000,
					TotalUnpaid: 1000,
				}, nil
			},
		},
		{
			name: "Error because service return error",
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
			service: func(ctx context.Context, customerID int) (*dto.CustomerSummary, error) {
				return nil, customerrors.ErrInternalServer
			},
		},
	}

	for _, test := range tests {
		expectedResponse, err := json.Marshal(test.expectedResponses)
		if err != nil {
			log.Fatal(err.Error())
		}

		res := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, apiversioning.APIVersionOne+"/customers", nil)

		ctx := s.e.NewContext(req, res)
		ctx.Set(dto.CustomerCTXKey, test.customerCtx)

		s.handler.service = mockCustomerService{
			funcGetSummary: test.service,
		}

		s.Suite.T().Run(test.name, func(t *testing.T) {
			if assert.NoError(t, s.handler.GetSummary()(ctx)) {
				assert.Equal(t, string(expectedResponse), res.Body.String()[:len(res.Body.String())-1])
				assert.Equal(t, test.expectedCode, res.Code)
			}
		})
	}
}
