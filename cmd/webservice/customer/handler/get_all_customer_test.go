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

func (s *customerHandlerSuite) TestGetAllCustomer() {
	tests := []struct {
		name   string
		params struct {
			page  string
			limit string
		}
		expectedCode      int
		expectedResponses dto.BaseResponse
		service           funcGetAllCustomer
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
			expectedCode: http.StatusOK,
			expectedResponses: dto.BaseResponse{
				Error: nil,
				Data: []dto.GetAllCustomerRespond{
					{
						ID:                1,
						FullName:          "valid test",
						Email:             "valid@gmail.com",
						DisplayProfileURL: stringutils.MakePointerString("https://valid.com"),
						Address:           stringutils.MakePointerString("valid"),
						CreatedAt:         "1970-01-01",
						UpdatedAt:         "1970-01-01",
					},
				},
			},
			service: func(ctx context.Context, req *dto.GetAllCustomerRequest) (*[]dto.GetAllCustomerRespond, error) {
				return &[]dto.GetAllCustomerRespond{
					{
						ID:                1,
						FullName:          "valid test",
						Email:             "valid@gmail.com",
						DisplayProfileURL: stringutils.MakePointerString("https://valid.com"),
						Address:           stringutils.MakePointerString("valid"),
						CreatedAt:         "1970-01-01",
						UpdatedAt:         "1970-01-01",
					},
				}, nil
			},
		},
		{
			name: "Bad request",
			params: struct {
				page  string
				limit string
			}{},
			expectedCode: http.StatusBadRequest,
			expectedResponses: dto.BaseResponse{
				Error: &dto.ErrorBaseResponse{
					Message: "bad request",
					Detail: []string{
						"Page is a required field",
						"Limit is a required field",
					},
				},
				Data: nil,
			},
			service: func(ctx context.Context, req *dto.GetAllCustomerRequest) (*[]dto.GetAllCustomerRespond, error) {
				return &[]dto.GetAllCustomerRespond{
					{
						ID:                1,
						FullName:          "valid test",
						Email:             "valid@gmail.com",
						DisplayProfileURL: stringutils.MakePointerString("https://valid.com"),
						Address:           stringutils.MakePointerString("valid"),
						CreatedAt:         "1970-01-01",
						UpdatedAt:         "1970-01-01",
					},
				}, nil
			},
		},
		{
			name: "Service return error",
			params: struct {
				page  string
				limit string
			}{
				page:  "1",
				limit: "10",
			},
			expectedCode: http.StatusInternalServerError,
			expectedResponses: dto.BaseResponse{
				Error: &dto.ErrorBaseResponse{
					Message: customerrors.ErrInternalServer.Error(),
				},
				Data: nil,
			},
			service: func(ctx context.Context, req *dto.GetAllCustomerRequest) (*[]dto.GetAllCustomerRespond, error) {
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

		s.handler.service = mockCustomerService{
			funcGetAllCustomer: test.service,
		}

		s.Suite.T().Run(test.name, func(t *testing.T) {
			if assert.NoError(t, s.handler.GetAllCustomer()(ctx)) {
				assert.Equal(t, string(expectedResponse), res.Body.String()[:len(res.Body.String())-1])
				assert.Equal(t, test.expectedCode, res.Code)
			}
		})
	}
}
