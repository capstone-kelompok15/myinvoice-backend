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

func (s *bankHandlerSuite) TestGetAllBank() {
	tests := []struct {
		name              string
		expectedCode      int
		expectedResponses dto.BaseResponse
		service           funcGetAllBank
	}{
		{
			name:         "Valid",
			expectedCode: http.StatusOK,
			expectedResponses: dto.BaseResponse{
				Error: nil,
				Data: []dto.BankResponse{
					{
						ID:       1,
						BankName: "BCA",
						Code:     "BCA",
					},
				},
			},
			service: func(ctx context.Context) (*[]dto.BankResponse, error) {
				return &[]dto.BankResponse{
					{
						ID:       1,
						BankName: "BCA",
						Code:     "BCA",
					},
				}, nil
			},
		},
		{
			name:         "Internal Server Error",
			expectedCode: http.StatusInternalServerError,
			expectedResponses: dto.BaseResponse{
				Error: &dto.ErrorBaseResponse{
					Message: customerrors.ErrInternalServer.Error(),
				},
				Data: nil,
			},
			service: func(ctx context.Context) (*[]dto.BankResponse, error) {
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
		req := httptest.NewRequest(http.MethodGet, apiversioning.APIVersionOne+"/banks", nil)

		ctx := s.e.NewContext(req, res)
		s.handler.service = mockBankService{
			funcGetAllBank: test.service,
		}

		s.Suite.T().Run(test.name, func(t *testing.T) {
			if assert.NoError(t, s.handler.GetAllBank()(ctx)) {
				assert.Equal(t, string(expectedResponse), res.Body.String()[:len(res.Body.String())-1])
				assert.Equal(t, test.expectedCode, res.Code)
			}
		})
	}
}
