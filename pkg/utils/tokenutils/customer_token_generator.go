package tokenutils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"strings"

	"github.com/capstone-kelompok15/myinvoice-backend/config"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/stringutils"
)

type CustomerAccessTokenParams struct {
	DeviceInformation string
	UserInformation   *dto.CustomerContext
	Config            *config.CustomerToken
}

const delimiter = "."

func GenerateCustomerAccessToken(params *CustomerAccessTokenParams) (*string, error) {
	deviceInformation, userInformation, err := marshalParams(params)
	if err != nil {
		return nil, err
	}

	var accessToken strings.Builder

	headPart := base64.RawStdEncoding.EncodeToString(deviceInformation)
	accessToken.Write([]byte(headPart + delimiter))

	payloadPart := base64.RawStdEncoding.EncodeToString(userInformation)
	accessToken.Write([]byte(payloadPart))

	hmacVerifyPart := hmac.New(sha256.New, []byte(params.Config.SecretKey))
	hmacVerifyPart.Write([]byte(accessToken.String()))

	hmacVerifyPartString := base64.RawStdEncoding.EncodeToString(hmacVerifyPart.Sum(nil))
	accessToken.Write([]byte(delimiter + hmacVerifyPartString))

	return stringutils.MakePointerString(accessToken.String()), nil
}

func marshalParams(params *CustomerAccessTokenParams) (deviceInformation, userInformation []byte, err error) {
	deviceInformationStruct := struct {
		DeviceID string `json:"device_id"`
	}{
		DeviceID: params.DeviceInformation,
	}

	deviceInformation, err = json.Marshal(deviceInformationStruct)
	if err != nil {
		return nil, nil, err
	}

	userInformationStruct := struct {
		ID       int    `json:"id"`
		FullName string `json:"full_name"`
	}{
		ID:       params.UserInformation.ID,
		FullName: params.UserInformation.FullName,
	}

	userInformation, err = json.Marshal(userInformationStruct)
	if err != nil {
		return nil, nil, err
	}

	return deviceInformation, userInformation, nil
}
