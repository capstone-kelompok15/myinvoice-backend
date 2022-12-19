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

func (s customerHandlerSuite) TestGetAllNotificationCustomer() {
	tests := []struct {
		name   string
		params struct {
			page  string
			limit string
		}
		customerCtx       *dto.CustomerContext
		expectedCode      int
		expectedResponses dto.BaseResponse
		service           funcGetAllNotificationCustomer
	}{
		{
			name: "Valid",
			params: struct {
				page  string
				limit string
			}{
				page:  "1",
				limit: "10",
			},
			customerCtx: &dto.CustomerContext{
				ID:       1,
				DeviceID: "abcdf",
				FullName: "Logged in Customer",
			},
			expectedCode: http.StatusOK,
			expectedResponses: dto.BaseResponse{
				Error: nil,
				Data: []dto.NotificationRespond{
					{
						ID:               1,
						InvoiceID:        1,
						NotificationType: "Valid",
						Title:            "Valid",
						Content:          "Valid",
						IsRead:           true,
						CreatedAt:        "1970-01-01",
					},
				},
			},
			service: func(ctx context.Context, customerID int, req *dto.NotificationRequest) (*[]dto.NotificationRespond, error) {
				return &[]dto.NotificationRespond{
					{
						ID:               1,
						InvoiceID:        1,
						NotificationType: "Valid",
						Title:            "Valid",
						Content:          "Valid",
						IsRead:           true,
						CreatedAt:        "1970-01-01",
					},
				}, nil
			},
		},
		{
			name: "Error because customer context is nil",
			params: struct {
				page  string
				limit string
			}{
				page:  "1",
				limit: "10",
			},
			customerCtx:  nil,
			expectedCode: http.StatusInternalServerError,
			expectedResponses: dto.BaseResponse{
				Error: &dto.ErrorBaseResponse{
					Message: customerrors.ErrInternalServer.Error(),
				},
				Data: nil,
			},
			service: func(ctx context.Context, customerID int, req *dto.NotificationRequest) (*[]dto.NotificationRespond, error) {
				return &[]dto.NotificationRespond{
					{
						ID:               1,
						InvoiceID:        1,
						NotificationType: "Valid",
						Title:            "Valid",
						Content:          "Valid",
						IsRead:           true,
						CreatedAt:        "1970-01-01",
					},
				}, nil
			},
		},
		{
			name: "Error because request does not have query param",
			params: struct {
				page  string
				limit string
			}{},
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
						"Page is a required field",
						"Limit is a required field",
					},
				},
				Data: nil,
			},
			service: func(ctx context.Context, customerID int, req *dto.NotificationRequest) (*[]dto.NotificationRespond, error) {
				return &[]dto.NotificationRespond{
					{
						ID:               1,
						InvoiceID:        1,
						NotificationType: "Valid",
						Title:            "Valid",
						Content:          "Valid",
						IsRead:           true,
						CreatedAt:        "1970-01-01",
					},
				}, nil
			},
		},
		{
			name: "Error because customer context is nil",
			params: struct {
				page  string
				limit string
			}{
				page:  "1",
				limit: "10",
			},
			customerCtx:  nil,
			expectedCode: http.StatusInternalServerError,
			expectedResponses: dto.BaseResponse{
				Error: &dto.ErrorBaseResponse{
					Message: customerrors.ErrInternalServer.Error(),
				},
				Data: nil,
			},
			service: func(ctx context.Context, customerID int, req *dto.NotificationRequest) (*[]dto.NotificationRespond, error) {
				return &[]dto.NotificationRespond{
					{
						ID:               1,
						InvoiceID:        1,
						NotificationType: "Valid",
						Title:            "Valid",
						Content:          "Valid",
						IsRead:           true,
						CreatedAt:        "1970-01-01",
					},
				}, nil
			},
		},
		{
			name: "Error because service return error",
			params: struct {
				page  string
				limit string
			}{
				page:  "1",
				limit: "10",
			},
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
			service: func(ctx context.Context, customerID int, req *dto.NotificationRequest) (*[]dto.NotificationRespond, error) {
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
		query := req.URL.Query()
		query.Add("page", test.params.page)
		query.Add("limit", test.params.limit)
		req.URL.RawQuery = query.Encode()

		ctx := s.e.NewContext(req, res)
		ctx.Set(dto.CustomerCTXKey, test.customerCtx)

		s.handler.service = mockCustomerService{
			funcGetAllNotificationCustomer: test.service,
		}

		s.Suite.T().Run(test.name, func(t *testing.T) {
			if assert.NoError(t, s.handler.GetAllNotificationCustomer()(ctx)) {
				assert.Equal(t, string(expectedResponse), res.Body.String()[:len(res.Body.String())-1])
				assert.Equal(t, test.expectedCode, res.Code)
			}
		})
	}
}
