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

func (s merchantHandlerSuite) TestGetDashboard() {
	response := dto.MerchantDashboard{
		OverviewMerchantDashboard: dto.OverviewMerchantDashboard{
			PaymentReceivedQuantity: 1000,
			CustomerQuantity:        1000,
			InvoiceQuantity:         1000,
			UnpaidInvoiceQuantity:   1000,
		},
		RecentInvoiceMerchantDashboard: []dto.RecentInvoiceMerchantDashboard{
			{
				Price:              1000,
				InvoiceID:          1,
				CustomerName:       "Valid",
				InvoiceExpiredDate: "1970-01-01",
			},
		},
		RecentPaymentMerchantDashboard: []dto.RecentPaymentMerchantDashboard{
			{
				RecentInvoiceMerchantDashboard: dto.RecentInvoiceMerchantDashboard{
					Price:              1000,
					InvoiceID:          1,
					CustomerName:       "Valid",
					InvoiceExpiredDate: "1970-01-01",
				},
				PaymentType: "Manual",
			},
		},
	}

	tests := []struct {
		name              string
		merchantContext   *dto.AdminContext
		expectedCode      int
		expectedResponses dto.BaseResponse
		service           funcGetDashboard
	}{
		{
			name: "Valid",
			merchantContext: &dto.AdminContext{
				ID:           1,
				MerchantID:   1,
				MerchantName: "Valid",
			},
			expectedCode: http.StatusOK,
			expectedResponses: dto.BaseResponse{
				Error: nil,
				Data:  response,
			},
			service: func(ctx context.Context, merchantID int) (*dto.MerchantDashboard, error) {
				return &response, nil
			},
		},
		{
			name:            "Error because merchant context is nil",
			merchantContext: nil,
			expectedCode:    http.StatusInternalServerError,
			expectedResponses: dto.BaseResponse{
				Error: &dto.ErrorBaseResponse{
					Message: customerrors.ErrInternalServer.Error(),
				},
				Data: nil,
			},
			service: func(ctx context.Context, merchantID int) (*dto.MerchantDashboard, error) {
				return &response, nil
			},
		},
		{
			name: "Error because service return error",
			merchantContext: &dto.AdminContext{
				ID:           1,
				MerchantID:   1,
				MerchantName: "Valid",
			},
			expectedCode: http.StatusInternalServerError,
			expectedResponses: dto.BaseResponse{
				Error: &dto.ErrorBaseResponse{
					Message: customerrors.ErrInternalServer.Error(),
				},
				Data: nil,
			},
			service: func(ctx context.Context, merchantID int) (*dto.MerchantDashboard, error) {
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
		req := httptest.NewRequest(http.MethodGet, apiversioning.APIVersionOne+"/merchants", nil)

		ctx := s.e.NewContext(req, res)
		ctx.Set(dto.AdminCTXKey, test.merchantContext)

		s.handler.service = mockMerchantService{
			funcGetDashboard: test.service,
		}

		s.Suite.T().Run(test.name, func(t *testing.T) {
			if assert.NoError(t, s.handler.GetDashboard()(ctx)) {
				assert.Equal(t, string(expectedResponse), res.Body.String()[:len(res.Body.String())-1])
				assert.Equal(t, test.expectedCode, res.Code)
			}
		})
	}
}
