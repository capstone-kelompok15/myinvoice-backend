package passwordutils

import (
	"crypto/md5"
	"encoding/hex"
)

func HashPassword(str string) string {
	md5Hex := md5.Sum([]byte(str))
	return hex.EncodeToString(md5Hex[:])
}
