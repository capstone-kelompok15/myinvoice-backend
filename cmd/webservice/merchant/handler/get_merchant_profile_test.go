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
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/stringutils"
	"github.com/stretchr/testify/assert"
)

func (s merchantHandlerSuite) TestGetMerchantProfile() {
	resp := dto.MerchantProfileResponse{
		ID:                  1,
		Username:            "Valid",
		Email:               "Valid",
		MerchantName:        "Valid",
		DisplayProfileURL:   stringutils.MakePointerString("Valid"),
		MerchantPhoneNumber: stringutils.MakePointerString("Valid"),
		MerchantAddress:     stringutils.MakePointerString("Valid"),
	}

	tests := []struct {
		name              string
		merchantContext   *dto.AdminContext
		expectedCode      int
		expectedResponses dto.BaseResponse
		service           funcGetMerchantProfile
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
				Data:  resp,
			},
			service: func(ctx context.Context, merchantID int) (*dto.MerchantProfileResponse, error) {
				return &resp, nil
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
			service: func(ctx context.Context, merchantID int) (*dto.MerchantProfileResponse, error) {
				return &resp, nil
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
			service: func(ctx context.Context, merchantID int) (*dto.MerchantProfileResponse, error) {
				return nil, customerrors.ErrInternalServer
			},
		},
		{
			name: "Error because merchant not found",
			merchantContext: &dto.AdminContext{
				ID:           1,
				MerchantID:   1,
				MerchantName: "Valid",
			},
			expectedCode: http.StatusNotFound,
			expectedResponses: dto.BaseResponse{
				Error: &dto.ErrorBaseResponse{
					Message: customerrors.ErrRecordNotFound.Error(),
				},
				Data: nil,
			},
			service: func(ctx context.Context, merchantID int) (*dto.MerchantProfileResponse, error) {
				return nil, customerrors.ErrRecordNotFound
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
			funcGetMerchantProfile: test.service,
		}

		s.Suite.T().Run(test.name, func(t *testing.T) {
			if assert.NoError(t, s.handler.GetMerchantProfile()(ctx)) {
				assert.Equal(t, string(expectedResponse), res.Body.String()[:len(res.Body.String())-1])
				assert.Equal(t, test.expectedCode, res.Code)
			}
		})
	}
}
