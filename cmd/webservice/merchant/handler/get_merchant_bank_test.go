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

func (s merchantHandlerSuite) TestGetMerchantBank() {
	resp := []dto.GetMerchantBankResponse{
		{
			ID:         1,
			BankName:   "BCA",
			BankCode:   "BCA",
			OnBehalfOf: "Valid",
			BankNumber: "123",
		},
	}

	tests := []struct {
		name   string
		params struct {
			merchantID string
		}
		expectedCode      int
		expectedResponses dto.BaseResponse
		service           funcGetMerchantBank
	}{
		{
			name: "Valid",
			params: struct{ merchantID string }{
				merchantID: "1",
			},
			expectedCode: http.StatusOK,
			expectedResponses: dto.BaseResponse{
				Error: nil,
				Data:  resp,
			},
			service: func(ctx context.Context, req *dto.GetMerchantBankRequest) (*[]dto.GetMerchantBankResponse, error) {
				return &resp, nil
			},
		},
		{
			name:         "Error because path parameter is empty",
			params:       struct{ merchantID string }{},
			expectedCode: http.StatusBadRequest,
			expectedResponses: dto.BaseResponse{
				Error: &dto.ErrorBaseResponse{
					Message: customerrors.ErrBadRequest.Error(),
					Detail: []string{
						"MerchantID is a required field",
					},
				},
				Data: nil,
			},
			service: func(ctx context.Context, req *dto.GetMerchantBankRequest) (*[]dto.GetMerchantBankResponse, error) {
				return &resp, nil
			},
		},
		{
			name: "Service return error",
			params: struct{ merchantID string }{
				merchantID: "1",
			},
			expectedCode: http.StatusInternalServerError,
			expectedResponses: dto.BaseResponse{
				Error: &dto.ErrorBaseResponse{
					Message: customerrors.ErrInternalServer.Error(),
				},
				Data: nil,
			},
			service: func(ctx context.Context, req *dto.GetMerchantBankRequest) (*[]dto.GetMerchantBankResponse, error) {
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
		ctx.SetParamNames("merchant_id")
		ctx.SetParamValues(test.params.merchantID)

		s.handler.service = mockMerchantService{
			funcGetMerchantBank: test.service,
		}

		s.Suite.T().Run(test.name, func(t *testing.T) {
			if assert.NoError(t, s.handler.GetMerchantBank()(ctx)) {
				assert.Equal(t, string(expectedResponse), res.Body.String()[:len(res.Body.String())-1])
				assert.Equal(t, test.expectedCode, res.Code)
			}
		})
	}
}
