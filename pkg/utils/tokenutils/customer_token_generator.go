package tokenutils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"strings"

	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/dateutils"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/stringutils"
)

const delimiter = "."

func GenerateCustomerAccessToken(params *dto.CustomerAccessTokenParams) (*string, error) {
	deviceInformation, userInformation, err := marshalParams(params)
	if err != nil {
		return nil, err
	}

	var accessToken strings.Builder

	headPart := base64.RawURLEncoding.EncodeToString(deviceInformation)
	accessToken.Write([]byte(headPart + delimiter))

	payloadPart := base64.RawURLEncoding.EncodeToString(userInformation)
	accessToken.Write([]byte(payloadPart))

	signaturePart := CreateSignatureToken(accessToken, params.Config.SecretKey)

	hmacVerifyPartString := base64.RawURLEncoding.EncodeToString([]byte(signaturePart))
	accessToken.Write([]byte(delimiter + hmacVerifyPartString))

	return stringutils.MakePointerString(accessToken.String()), nil
}

func CreateSignatureToken(unsignedAccessToken strings.Builder, secretKey string) string {
	hmacVerifyPart := hmac.New(sha256.New, []byte(secretKey))
	hmacVerifyPart.Write([]byte(unsignedAccessToken.String()))
	signaturePart := string(hmacVerifyPart.Sum(nil))
	return signaturePart
}

func marshalParams(params *dto.CustomerAccessTokenParams) (deviceInformation, userInformation []byte, err error) {
	deviceInformationStruct := dto.HeaderCustomerTokenPart{
		DeviceID: params.DeviceInformation,
		Date:     dateutils.NowNanoTimeStamp(),
	}

	deviceInformation, err = json.Marshal(deviceInformationStruct)
	if err != nil {
		return nil, nil, err
	}

	userInformationStruct := dto.PayloadCustomerTokenPart{
		ID:       params.UserInformation.ID,
		FullName: params.UserInformation.FullName,
	}

	userInformation, err = json.Marshal(userInformationStruct)
	if err != nil {
		return nil, nil, err
	}

	return deviceInformation, userInformation, nil
}
