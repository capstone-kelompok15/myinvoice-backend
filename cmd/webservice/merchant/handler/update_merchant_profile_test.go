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

func (s merchantHandlerSuite) TestUpdateMerchantProfile() {
	req := dto.UpdateMerchantProfileRequest{
		Username:            "Valid",
		MerchantName:        "Valid",
		MerchantPhoneNumber: "Valid",
		MerchantAddress:     "Valid",
	}

	tests := []struct {
		name              string
		req               dto.UpdateMerchantProfileRequest
		contentType       string
		merchantContext   *dto.AdminContext
		expectedCode      int
		expectedResponses dto.BaseResponse
		service           funcUpdateMerchantProfile
	}{
		{
			name:        "Valid",
			req:         req,
			contentType: echo.MIMEApplicationJSON,
			merchantContext: &dto.AdminContext{
				ID:           1,
				MerchantID:   1,
				MerchantName: "Valid",
			},
			expectedCode: http.StatusOK,
			expectedResponses: dto.BaseResponse{
				Error: nil,
				Data:  "Update Merchant Success!",
			},
			service: func(ctx context.Context, merchantID *int, req *dto.UpdateMerchantProfileRequest) error {
				return nil
			},
		},
		{
			name:            "Error because merchant context is nil",
			req:             req,
			contentType:     echo.MIMEApplicationJSON,
			merchantContext: nil,
			expectedCode:    http.StatusInternalServerError,
			expectedResponses: dto.BaseResponse{
				Error: &dto.ErrorBaseResponse{
					Message: customerrors.ErrInternalServer.Error(),
				},
				Data: nil,
			},
			service: func(ctx context.Context, merchantID *int, req *dto.UpdateMerchantProfileRequest) error {
				return nil
			},
		},
		{
			name:        "Error because some payload is empty",
			req:         dto.UpdateMerchantProfileRequest{},
			contentType: echo.MIMEApplicationJSON,
			merchantContext: &dto.AdminContext{
				ID:           1,
				MerchantID:   1,
				MerchantName: "Valid",
			},
			expectedCode: http.StatusBadRequest,
			expectedResponses: dto.BaseResponse{
				Error: &dto.ErrorBaseResponse{
					Message: customerrors.ErrBadRequest.Error(),
					Detail: []string{
						"Username is a required field",
						"MerchantName is a required field",
						"MerchantPhoneNumber is a required field",
						"MerchantAddress is a required field",
					},
				},
				Data: nil,
			},
			service: func(ctx context.Context, merchantID *int, req *dto.UpdateMerchantProfileRequest) error {
				return nil
			},
		},
		{
			name:        "Error because some payload is empty",
			req:         req,
			contentType: echo.MIMEApplicationForm,
			merchantContext: &dto.AdminContext{
				ID:           1,
				MerchantID:   1,
				MerchantName: "Valid",
			},
			expectedCode: http.StatusBadRequest,
			expectedResponses: dto.BaseResponse{
				Error: &dto.ErrorBaseResponse{
					Message: customerrors.ErrBadRequest.Error(),
					Detail: []string{
						"Username is a required field",
						"MerchantName is a required field",
						"MerchantPhoneNumber is a required field",
						"MerchantAddress is a required field",
					},
				},
				Data: nil,
			},
			service: func(ctx context.Context, merchantID *int, req *dto.UpdateMerchantProfileRequest) error {
				return nil
			},
		},
		{
			name:        "Error because service return error",
			req:         req,
			contentType: echo.MIMEApplicationJSON,
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
			service: func(ctx context.Context, merchantID *int, req *dto.UpdateMerchantProfileRequest) error {
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
		req := httptest.NewRequest(http.MethodGet, apiversioning.APIVersionOne+"/merchants", strings.NewReader(string(request)))
		req.Header.Set(echo.HeaderContentType, test.contentType)

		ctx := s.e.NewContext(req, res)
		ctx.Set(dto.AdminCTXKey, test.merchantContext)

		s.handler.service = mockMerchantService{
			funcUpdateMerchantProfile: test.service,
		}

		s.Suite.T().Run(test.name, func(t *testing.T) {
			if assert.NoError(t, s.handler.UpdateMerchantProfile()(ctx)) {
				assert.Equal(t, string(expectedResponse), res.Body.String()[:len(res.Body.String())-1])
				assert.Equal(t, test.expectedCode, res.Code)
			}
		})
	}
}
