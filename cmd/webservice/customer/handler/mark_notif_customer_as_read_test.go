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

func (s customerHandlerSuite) TestMarkNotifCustomerAsRead() {
	tests := []struct {
		name   string
		params struct {
			notifId string
		}
		customerCtx       *dto.CustomerContext
		expectedCode      int
		expectedResponses dto.BaseResponse
		service           funcMarkNotifCustomerAsRead
	}{
		{
			name:   "Valid",
			params: struct{ notifId string }{notifId: "1"},
			customerCtx: &dto.CustomerContext{
				ID:       1,
				DeviceID: "abcdf",
				FullName: "Logged in Customer",
			},
			expectedCode: http.StatusOK,
			expectedResponses: dto.BaseResponse{
				Error: nil,
				Data:  "Mark notif customer as read success",
			},
			service: func(ctx context.Context, NotifID, CustomerID int) error {
				return nil
			},
		},
		{
			name:         "Error because customer context is nil",
			params:       struct{ notifId string }{notifId: "1"},
			customerCtx:  nil,
			expectedCode: http.StatusInternalServerError,
			expectedResponses: dto.BaseResponse{
				Error: &dto.ErrorBaseResponse{
					Message: customerrors.ErrInternalServer.Error(),
				},
				Data: nil,
			},
			service: func(ctx context.Context, NotifID, CustomerID int) error {
				return nil
			},
		},
		{
			name:   "Error because service return error",
			params: struct{ notifId string }{notifId: "1"},
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
			service: func(ctx context.Context, NotifID, CustomerID int) error {
				return customerrors.ErrInternalServer
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
		ctx.SetParamNames("id")
		ctx.SetParamValues(test.params.notifId)

		s.handler.service = mockCustomerService{
			funcMarkNotifCustomerAsRead: test.service,
		}

		s.Suite.T().Run(test.name, func(t *testing.T) {
			if assert.NoError(t, s.handler.MarkNotifCustomerAsRead()(ctx)) {
				assert.Equal(t, string(expectedResponse), res.Body.String()[:len(res.Body.String())-1])
				assert.Equal(t, test.expectedCode, res.Code)
			}
		})
	}
}
