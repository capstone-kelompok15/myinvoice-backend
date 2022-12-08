package tokenutils

import "encoding/base64"

func EncodeBase64String(text string) *string {
	encode64Text := base64.RawURLEncoding.EncodeToString([]byte(text))
	return &encode64Text
}

func DecodeBase64String(encodeText string) (*string, error) {
	decoded64Byte, err := base64.RawURLEncoding.DecodeString(encodeText)
	if err != nil {
		return nil, err
	}
	decoded64Str := string(decoded64Byte)
	return &decoded64Str, nil
}
