package impl

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
	customerrors "github.com/capstone-kelompok15/myinvoice-backend/pkg/errors"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/httputils"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/passwordutils"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/tokenutils"
	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
)

func (m *middleware) CustomerMustAuthorized() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authorization := c.Request().Header.Get("Authorization")
			if authorization == "" {
				return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
					Err:    customerrors.ErrBadRequest,
					Detail: []string{"Authorization header value couldn't be empty"},
				})
			}

			splitted := strings.SplitAfter(authorization, "Bearer ")
			if len(splitted) != 2 {
				return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
					Err:    customerrors.ErrBadRequest,
					Detail: []string{"Bearer format is not valid"},
				})
			}

			if splitted[1] == "" {
				return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
					Err:    customerrors.ErrBadRequest,
					Detail: []string{"Bearer value is couldn't empty"},
				})
			}

			accessToken := splitted[1]
			splittedToken := strings.Split(accessToken, ".")
			encodedHeader, encodedPayload, encodedSignature := splittedToken[0], splittedToken[1], splittedToken[2]

			validate := validateCustomerToken(m.config.SecretKey, encodedHeader, encodedPayload, encodedSignature)
			if !validate {
				return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
					Err:    customerrors.ErrUnauthorized,
					Detail: []string{"Token not valid"},
				})
			}

			headerStr, err := tokenutils.DecodeBase64String(encodedHeader)
			if err != nil {
				m.log.Warningln("[CustomerMustAuthorized] Failed on decode header:", err.Error())
				return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
					Err: customerrors.ErrInternalServer,
				})
			}

			payloadStr, err := tokenutils.DecodeBase64String(encodedPayload)
			if err != nil {
				m.log.Warningln("[CustomerMustAuthorized] Failed on decode payload", err.Error())
				return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
					Err: customerrors.ErrInternalServer,
				})
			}

			var header dto.HeaderCustomerTokenPart
			var payload dto.PayloadCustomerTokenPart

			err = json.Unmarshal([]byte(*headerStr), &header)
			if err != nil {
				m.log.Warningln("[CustomerMustAuthorized] Failed on unmarshalling header to json", err.Error())
				return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
					Err: customerrors.ErrInternalServer,
				})
			}

			err = json.Unmarshal([]byte(*payloadStr), &payload)
			if err != nil {
				m.log.Warningln("[CustomerMustAuthorized] Failed on unmarshalling payload to json", err.Error())
				return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
					Err: customerrors.ErrInternalServer,
				})
			}

			var customer *dto.CustomerContext

			redisKey := fmt.Sprintf("customer-access-token:%d", payload.ID)
			res := m.redis.Get(c.Request().Context(), redisKey)
			if res.Err() == redis.Nil {
				deviceID := passwordutils.HashPassword(header.DeviceID)
				customer, err = m.middlewareRepo.ValidateCustomer(c.Request().Context(), &deviceID, &payload.ID)
				if err != nil {
					if err == customerrors.ErrRecordNotFound {
						return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
							Err:    customerrors.ErrUnauthorized,
							Detail: []string{"Account not exist"},
						})
					}
					m.log.Warningln("[CustomerMustAuthorized] Failed on running validate customer service:", err.Error())
					return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
						Err: customerrors.ErrInternalServer,
					})
				}
				customer.DeviceID = header.DeviceID
				customerJSONByte, err := json.Marshal(customer)
				if err != nil {
					m.log.Warningln("[CustomerMustAuthorized] Failed on marshalling customer context:", err.Error())
					return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
						Err: customerrors.ErrInternalServer,
					})
				}
				redisErr := m.redis.Set(c.Request().Context(), redisKey, string(customerJSONByte), 1*time.Hour)
				if redisErr.Err() != nil {
					m.log.Warningln("[CustomerMustAuthorized] Failed on set value to the redis", err.Error())
					return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
						Err: customerrors.ErrInternalServer,
					})
				}
			} else {
				err = json.Unmarshal([]byte(res.Val()), &customer)
				if err != nil {
					m.log.Warningln("[CustomerMustAuthorized] Failed on unmarshalling the redis value to the customer context", err.Error())
					return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
						Err: customerrors.ErrInternalServer,
					})
				}

				if customer.DeviceID != header.DeviceID {
					return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
						Err:    customerrors.ErrUnauthorized,
						Detail: []string{"Device ID invalid"},
					})
				}
			}

			c.Set(dto.CustomerCTXKey, customer)
			return next(c)
		}
	}
}

func validateCustomerToken(secretKey string, headerPart string, payloadPart string, signaturePartRequest string) bool {
	var accessToken strings.Builder

	accessToken.Write([]byte(headerPart + "."))
	accessToken.Write([]byte(payloadPart))

	signaturePartExpectedHMAC := tokenutils.CreateSignatureToken(accessToken, secretKey)
	signaturePartExpectedStr := tokenutils.EncodeBase64String(signaturePartExpectedHMAC)

	return signaturePartRequest == *signaturePartExpectedStr
}
