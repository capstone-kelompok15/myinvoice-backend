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

func (s merchantHandlerSuite) TestMarkNotifMerchantAsRead() {
	tests := []struct {
		name   string
		params struct {
			notifID string
		}
		merchantContext   *dto.AdminContext
		expectedCode      int
		expectedResponses dto.BaseResponse
		service           funcMarkNotifMerchantAsRead
	}{
		{
			name: "Valid",
			params: struct{ notifID string }{
				notifID: "1",
			},
			merchantContext: &dto.AdminContext{
				ID:           1,
				MerchantID:   1,
				MerchantName: "Valid",
			},
			expectedCode: http.StatusOK,
			expectedResponses: dto.BaseResponse{
				Error: nil,
				Data:  "Mark notif merchant as read success",
			},
			service: func(ctx context.Context, NotifID, MerchantID int) error {
				return nil
			},
		},
		{
			name: "Error because merchant context is nil",
			params: struct{ notifID string }{
				notifID: "1",
			},
			merchantContext: nil,
			expectedCode:    http.StatusInternalServerError,
			expectedResponses: dto.BaseResponse{
				Error: &dto.ErrorBaseResponse{
					Message: customerrors.ErrInternalServer.Error(),
				},
				Data: nil,
			},
			service: func(ctx context.Context, NotifID, MerchantID int) error {
				return nil
			},
		},
		{
			name:   "Error because parameter is empty",
			params: struct{ notifID string }{},
			merchantContext: &dto.AdminContext{
				ID:           1,
				MerchantID:   1,
				MerchantName: "Valid",
			},
			expectedCode: http.StatusBadRequest,
			expectedResponses: dto.BaseResponse{
				Error: &dto.ErrorBaseResponse{
					Message: customerrors.ErrBadRequest.Error(),
				},
				Data: nil,
			},
			service: func(ctx context.Context, NotifID, MerchantID int) error {
				return nil
			},
		},
		{
			name: "Error because service return error",
			params: struct{ notifID string }{
				notifID: "1",
			},
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
			service: func(ctx context.Context, NotifID, MerchantID int) error {
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
		req := httptest.NewRequest(http.MethodGet, apiversioning.APIVersionOne+"/merchants", nil)

		ctx := s.e.NewContext(req, res)
		ctx.Set(dto.AdminCTXKey, test.merchantContext)
		ctx.SetParamNames("id")
		ctx.SetParamValues(test.params.notifID)

		s.handler.service = mockMerchantService{
			funcMarkNotifMerchantAsRead: test.service,
		}

		s.Suite.T().Run(test.name, func(t *testing.T) {
			if assert.NoError(t, s.handler.MarkNotifMerchantAsRead()(ctx)) {
				assert.Equal(t, string(expectedResponse), res.Body.String()[:len(res.Body.String())-1])
				assert.Equal(t, test.expectedCode, res.Code)
			}
		})
	}
}
