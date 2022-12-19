package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/capstone-kelompok15/myinvoice-backend/cmd/webservice/apiversioning"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
	customerrors "github.com/capstone-kelompok15/myinvoice-backend/pkg/errors"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/stringutils"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func (s customerHandlerSuite) TestUpdateCustomerProfilePicture() {
	tests := []struct {
		name   string
		params struct {
			file *string
		}
		contentType       string
		customerCtx       *dto.CustomerContext
		expectedCode      int
		expectedResponses dto.BaseResponse
		service           funcUpdateProfilePicture
	}{
		{
			name: "Valid",
			params: struct{ file *string }{
				file: stringutils.MakePointerString("file"),
			},
			customerCtx: &dto.CustomerContext{
				ID:       1,
				DeviceID: "abcdf",
				FullName: "Logged in Customer",
			},
			expectedCode: http.StatusOK,
			expectedResponses: dto.BaseResponse{
				Error: nil,
				Data: dto.UpdateProfilePictureResponse{
					ImageURL: "url",
				},
			},
			service: func(ctx context.Context, userID *int, newProfilePictureURL *string) (*string, error) {
				return stringutils.MakePointerString("url"), nil
			},
		},
		{
			name: "Error because customer context is nil",
			params: struct{ file *string }{
				file: stringutils.MakePointerString("file"),
			},
			customerCtx:  nil,
			expectedCode: http.StatusInternalServerError,
			expectedResponses: dto.BaseResponse{
				Error: &dto.ErrorBaseResponse{
					Message: customerrors.ErrInternalServer.Error(),
				},
				Data: nil,
			},
			service: func(ctx context.Context, userID *int, newProfilePictureURL *string) (*string, error) {
				return stringutils.MakePointerString("url"), nil
			},
		},
		{
			name: "Error because form file nil",
			params: struct{ file *string }{
				file: nil,
			},
			customerCtx: &dto.CustomerContext{
				ID:       1,
				DeviceID: "abcdf",
				FullName: "Logged in Customer",
			},
			contentType:  echo.MIMEApplicationJSON,
			expectedCode: http.StatusBadRequest,
			expectedResponses: dto.BaseResponse{
				Error: &dto.ErrorBaseResponse{
					Message: customerrors.ErrBadRequest.Error(),
				},
				Data: nil,
			},
			service: func(ctx context.Context, userID *int, newProfilePictureURL *string) (*string, error) {
				return stringutils.MakePointerString("url"), nil
			},
		},
		{
			name: "Error because service return error",
			params: struct{ file *string }{
				file: stringutils.MakePointerString("file"),
			},
			customerCtx: &dto.CustomerContext{
				ID:       1,
				DeviceID: "abcdf",
				FullName: "Logged in Customer",
			},
			contentType:  echo.MIMEApplicationJSON,
			expectedCode: http.StatusInternalServerError,
			expectedResponses: dto.BaseResponse{
				Error: &dto.ErrorBaseResponse{
					Message: customerrors.ErrInternalServer.Error(),
				},
				Data: nil,
			},
			service: func(ctx context.Context, userID *int, newProfilePictureURL *string) (*string, error) {
				return nil, customerrors.ErrInternalServer
			},
		},
		{
			name: "Error because user not found",
			params: struct{ file *string }{
				file: stringutils.MakePointerString("file"),
			},
			customerCtx: &dto.CustomerContext{
				ID:       1,
				DeviceID: "abcdf",
				FullName: "Logged in Customer",
			},
			contentType:  echo.MIMEApplicationJSON,
			expectedCode: http.StatusNotFound,
			expectedResponses: dto.BaseResponse{
				Error: &dto.ErrorBaseResponse{
					Message: customerrors.ErrRecordNotFound.Error(),
				},
				Data: nil,
			},
			service: func(ctx context.Context, userID *int, newProfilePictureURL *string) (*string, error) {
				return nil, customerrors.ErrRecordNotFound
			},
		},
	}

	for _, test := range tests {
		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)

		if test.params.file != nil {
			writer.CreateFormFile("profile_picture", *test.params.file)
			test.contentType = writer.FormDataContentType()
		}

		err := writer.Close()
		if err != nil {
			log.Fatal(err.Error())
		}

		expectedResponse, err := json.Marshal(test.expectedResponses)
		if err != nil {
			log.Fatal(err.Error())
		}

		res := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, apiversioning.APIVersionOne+"/customers", body)
		req.Header.Set(echo.HeaderContentType, test.contentType)

		ctx := s.e.NewContext(req, res)
		ctx.Set(dto.CustomerCTXKey, test.customerCtx)

		s.handler.service = mockCustomerService{
			funcUpdateProfilePicture: test.service,
		}

		s.Suite.T().Run(test.name, func(t *testing.T) {
			if assert.NoError(t, s.handler.UpdateCustomerProfilePicture()(ctx)) {
				assert.Equal(t, string(expectedResponse), res.Body.String()[:len(res.Body.String())-1])
				assert.Equal(t, test.expectedCode, res.Code)
			}
		})
	}
}
